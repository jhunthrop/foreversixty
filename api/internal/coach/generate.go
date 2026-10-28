package coach

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

const (
	// ReviewMaxTokens bounds one model answer (§9.3, §15.3).
	ReviewMaxTokens = 1200
	roleAssistant   = "assistant"

	reasonValidation = "validation"
	reasonModelError = "model_error"
	// reasonPrivacy: the prompt pack carried a roster name or realm, so no call was made.
	reasonPrivacy = "privacy"

	// redactedPlayer replaces a roster name in a logged model output.
	redactedPlayer = "[player]"

	retryFactsLead = "These were not in the evidence pack: "
	retryOtherLead = "Fix these: "
	retryTail      = "Use only the pack."
)

var errNoClient = errors.New("coach: a generator needs a model client")

// Generator writes one review per key: caps, the model, validation, one retry with the
// problems listed back, then the rule-only rendering (§9, §10, §15). Validator carries the
// process-wide name dictionary (built once) and the fight's roster. Attempt needs only LLM
// and Validator; Generate also needs Caps, Store and Now.
type Generator struct {
	LLM       llm.Client
	Caps      *Caps
	Store     *Store
	Validator Validator
	// Realms is the fight's realm strings. With Validator.Roster they are what the prompt
	// pack must never carry (§19): a pack holding one is served rule-only, with no call.
	Realms      []string
	Model       string // the gateway slug, for ModelSlug when a response names none
	Temperature float64
	Log         *slog.Logger
	Now         func() time.Time
}

// Attempt is what the model flow made of one pack: the model's review when a try passed
// validation, else the rule-only rendering with the reason. Tokens and cost are summed over
// every answered call; ModelSlug and RequestID are the last answered call's.
type Attempt struct {
	Doc       ReviewDoc
	Source    string // SourceModel or SourceRules
	Reason    string // "" for a model review; reasonValidation or reasonModelError
	Retryable bool
	PassedOn  int // 1 or 2: the try that passed; 0 for the rule-only rendering
	Calls     int // model calls that answered
	TokensIn  int
	TokensOut int
	// TokensReasoning is the part of TokensOut the model spent reasoning.
	TokensReasoning int
	CostUSD         float64
	CostEstimated   bool
	ModelSlug       string
	RequestID       string
	Validation      []Problem // the last answered call's problems
	Outputs         []string  // each answered call's raw content, in order; never stored
}

// call is one model call and what validation made of it.
type call struct {
	content string
	resp    llm.Response
	probs   []Problem
}

// Generate writes one review for key: caps, then the model, then validation, one retry
// with the problems listed, then the rule-only rendering. It always stores a row and
// returns it; the error is a missing dependency or a failed save.
func (g *Generator) Generate(ctx context.Context, key ReviewKey, p Pack, persona Persona, accountID int64) (StoredReview, error) {
	if err := g.ready(); err != nil {
		return StoredReview{}, err
	}
	var res Attempt
	if g.withheld(p) {
		res = fallback(Attempt{}, p, persona, reasonPrivacy, false)
	} else if refusal := g.admit(ctx, accountID); refusal != RefusalNone {
		res = fallback(Attempt{}, p, persona, string(refusal), true)
	} else {
		// A model error is logged by g.call and is in res as model_error, retryable.
		res, _ = g.run(ctx, p, persona, g.recordSpend)
	}
	return g.save(ctx, storedOf(key, accountID, g.Now(), res))
}

// Attempt runs the model flow for one pack with no database: the call, validation, one
// retry with the problems listed, then the rule-only rendering. Nothing is stored and no
// spend is recorded. A model error is returned beside the filled rule-only Attempt.
func (g *Generator) Attempt(ctx context.Context, p Pack, persona Persona) (Attempt, error) {
	return g.run(ctx, p, persona, nil)
}

// run is the one model flow; onCall, when set, sees every answered call's response.
func (g *Generator) run(ctx context.Context, p Pack, persona Persona, onCall func(context.Context, llm.Response)) (Attempt, error) {
	if g.LLM == nil {
		return Attempt{}, errNoClient
	}
	if g.withheld(p) {
		return fallback(Attempt{}, p, persona, reasonPrivacy, false), nil
	}
	msgs := BuildPrompt(p, persona)
	var res Attempt
	first, err := g.call(ctx, msgs, p, persona, onCall)
	if err != nil {
		return fallback(withCost(res, first.resp), p, persona, reasonModelError, true), err
	}
	res = g.withCall(res, first)
	if len(first.probs) == 0 {
		return passed(res, first, 1), nil
	}
	second, err := g.call(ctx, retryMessages(msgs, first), p, persona, onCall)
	if err != nil {
		return fallback(withCost(res, second.resp), p, persona, reasonModelError, true), err
	}
	res = g.withCall(res, second)
	if len(second.probs) == 0 {
		return passed(res, second, 2), nil
	}
	g.logFailure(p, first, second)
	return fallback(res, p, persona, reasonValidation, false), nil
}

func (g *Generator) ready() error {
	if g.LLM == nil || g.Caps == nil || g.Store == nil || g.Now == nil {
		return errors.New("coach: a generator needs a model client, caps, a store and a clock")
	}
	return nil
}

func (g *Generator) logger() *slog.Logger {
	if g.Log != nil {
		return g.Log
	}
	return slog.Default()
}

// admit asks the caps before any call. A caps error fails closed (RefusalError).
func (g *Generator) admit(ctx context.Context, accountID int64) Refusal {
	refusal, err := g.Caps.Admit(ctx, accountID)
	if err != nil {
		g.logger().Error("coach: caps check failed; serving the rule-only review", slog.String("error", err.Error()))
		return RefusalError
	}
	return refusal
}

// recordSpend adds one call's cost to the day's spend; a failure is logged, not fatal.
func (g *Generator) recordSpend(ctx context.Context, resp llm.Response) {
	if err := g.Caps.RecordSpend(ctx, resp); err != nil {
		g.logger().Error("coach: spend not recorded", slog.String("error", err.Error()), slog.Float64("usd", resp.CostUSD))
	}
}

// call makes one model call, hands it to onCall, then parses and validates the answer. A
// parse failure is a shape problem, so it is retried like any other. A failed call that
// still carries a cost (attempts the provider may have billed) goes to onCall too, so the
// spend cap counts it; the returned call then holds only that response.
func (g *Generator) call(ctx context.Context, msgs []llm.Message, p Pack, persona Persona, onCall func(context.Context, llm.Response)) (call, error) {
	resp, err := g.LLM.Complete(ctx, llm.Request{
		Messages: msgs, JSONSchema: OutputSchema(), Temperature: g.Temperature, MaxTokens: ReviewMaxTokens,
	})
	if err != nil {
		g.logger().Warn("coach: model call failed", slog.String("error", err.Error()), slog.Float64("usd", resp.CostUSD))
		if onCall != nil && resp.CostUSD > 0 {
			onCall(ctx, resp)
		}
		return call{resp: resp}, err
	}
	if onCall != nil {
		onCall(ctx, resp)
	}
	c := call{content: resp.Content, resp: resp}
	doc, err := ParseOutput(resp.Content)
	if err != nil {
		c.probs = []Problem{{Kind: ProblemShape, Value: err.Error()}}
		return c, nil
	}
	c.probs = g.Validator.Validate(doc, p, persona)
	return c, nil
}

// withCall adds a call's tokens and cost and takes its slug, request id and problems as
// the last call's.
func (g *Generator) withCall(res Attempt, c call) Attempt {
	res.Calls++
	res.TokensIn += c.resp.Usage.PromptTokens
	res.TokensOut += c.resp.Usage.CompletionTokens
	res.TokensReasoning += c.resp.Usage.ReasoningTokens
	res.CostUSD += c.resp.CostUSD
	res.CostEstimated = res.CostEstimated || c.resp.CostEstimated
	res.ModelSlug = cmpOr(c.resp.Model, g.Model)
	res.RequestID = c.resp.RequestID
	res.Validation = slices.Clone(c.probs)
	res.Outputs = append(slices.Clone(res.Outputs), c.content)
	return res
}

// withCost adds a failed call's billed cost to the attempt.
func withCost(res Attempt, resp llm.Response) Attempt {
	if resp.CostUSD > 0 {
		res.CostUSD += resp.CostUSD
		res.CostEstimated = res.CostEstimated || resp.CostEstimated
	}
	return res
}

func cmpOr(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

// passed is the Attempt for an answer that validated; it parsed, so ParseOutput cannot fail.
func passed(res Attempt, c call, try int) Attempt {
	doc, _ := ParseOutput(c.content)
	res.Doc, res.Source, res.Reason, res.Retryable, res.PassedOn = doc, SourceModel, "", false, try
	return res
}

// fallback is the rule-only rendering with why it was served.
func fallback(res Attempt, p Pack, persona Persona, reason string, retryable bool) Attempt {
	res.Doc = RenderRules(p, persona.Drills)
	res.Source, res.Reason, res.Retryable, res.PassedOn = SourceRules, reason, retryable, 0
	return res
}

// storedOf is the coach_reviews row for an Attempt.
func storedOf(key ReviewKey, accountID int64, at time.Time, a Attempt) StoredReview {
	return StoredReview{
		Key: key, Doc: a.Doc, Source: a.Source, Reason: a.Reason, Retryable: a.Retryable,
		TokensIn: a.TokensIn, TokensOut: a.TokensOut, CostUSD: a.CostUSD, CostEstimated: a.CostEstimated,
		ModelSlug: a.ModelSlug, RequestID: a.RequestID, AccountID: accountID, GeneratedAt: at, Validation: a.Validation,
	}
}

func (g *Generator) save(ctx context.Context, row StoredReview) (StoredReview, error) {
	if err := g.Store.SaveReview(ctx, row); err != nil {
		return row, fmt.Errorf("coach: store the generated review: %w", err)
	}
	return row, nil
}

// retryMessages is the first two messages, the model's answer and the problems listed back.
func retryMessages(msgs []llm.Message, first call) []llm.Message {
	return append(slices.Clone(msgs),
		llm.Message{Role: roleAssistant, Content: first.content},
		llm.Message{Role: roleUser, Content: RetryMessage(first.probs)})
}

// RetryMessage lists the problems back to the model: the facts the pack does not hold
// first, quoted, then every other problem, then the instruction.
// `These were not in the evidence pack: "12 casts", "Whirlwind". Use only the pack.`
func RetryMessage(probs []Problem) string {
	var facts, other []string
	for _, pr := range probs {
		if isFactProblem(pr) {
			facts = appendUnique(facts, fmt.Sprintf("%q", pr.Value))
		} else {
			other = appendUnique(other, pr.Value)
		}
	}
	var parts []string
	if len(facts) > 0 {
		parts = append(parts, retryFactsLead+strings.Join(facts, ", ")+".")
	}
	if len(other) > 0 {
		parts = append(parts, retryOtherLead+strings.Join(other, "; ")+".")
	}
	return strings.Join(append(parts, retryTail), " ")
}

func isFactProblem(p Problem) bool {
	return p.Kind != ProblemShape && p.Kind != ProblemLink && p.Kind != ProblemTone
}

func appendUnique(list []string, s string) []string {
	if slices.Contains(list, s) {
		return list
	}
	return append(list, s)
}

// logFailure writes one entry for a review that failed validation twice: the pack (which
// carries no names), both outputs and the problems, with every roster name redacted. It
// never carries the review key or the gateway key.
func (g *Generator) logFailure(p Pack, first, second call) {
	// promptPack and Problem hold only strings, numbers and slices: Marshal cannot fail.
	pack, _ := json.Marshal(promptPackOf(p))
	probs, _ := json.Marshal(struct{ First, Second []Problem }{first.probs, second.probs})
	redact := g.redactor()
	g.logger().Warn("coach: review failed validation twice; serving the rule-only review",
		slog.String("pack_version", p.Version),
		slog.String("pack", string(pack)),
		slog.String("first_output", redact(first.content)),
		slog.String("second_output", redact(second.content)),
		slog.String("problems", redact(string(probs))))
}

// redactor replaces each roster name, in any case, with redactedPlayer.
func (g *Generator) redactor() func(string) string {
	var res []*regexp.Regexp
	for _, name := range g.Validator.Roster {
		if name = strings.TrimSpace(name); name != "" {
			res = append(res, regexp.MustCompile(`(?i)`+regexp.QuoteMeta(name)))
		}
	}
	return func(s string) string {
		for _, re := range res {
			s = re.ReplaceAllLiteralString(s, redactedPlayer)
		}
		return s
	}
}
