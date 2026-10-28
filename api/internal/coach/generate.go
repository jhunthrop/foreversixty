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
	// reviewMaxTokens bounds one model answer (§9.3, §15.3).
	reviewMaxTokens = 1200
	roleAssistant   = "assistant"

	reasonValidation = "validation"
	reasonModelError = "model_error"

	// redactedPlayer replaces a roster name in a logged model output.
	redactedPlayer = "[player]"

	retryFactsLead = "These were not in the evidence pack: "
	retryOtherLead = "Fix these: "
	retryTail      = "Use only the pack."
)

// Generator writes one review per key: caps, the model, validation, one retry with the
// problems listed back, then the rule-only rendering (§9, §10, §15). Validator carries the
// process-wide name dictionary (built once) and the fight's roster.
type Generator struct {
	LLM         llm.Client
	Caps        *Caps
	Store       *Store
	Validator   Validator
	Model       string // the gateway slug, for ModelSlug when a response names none
	Temperature float64
	Log         *slog.Logger
	Now         func() time.Time
}

// attempt is one model call and what validation made of it.
type attempt struct {
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
	row := StoredReview{Key: key, AccountID: accountID, GeneratedAt: g.Now()}
	if refusal := g.admit(ctx, accountID); refusal != RefusalNone {
		return g.save(ctx, fallback(row, p, persona, string(refusal), true))
	}
	msgs := BuildPrompt(p, persona)
	first, err := g.attempt(ctx, msgs, p, persona)
	if err != nil {
		return g.save(ctx, fallback(row, p, persona, reasonModelError, true))
	}
	row = g.withAttempt(row, first)
	if len(first.probs) == 0 {
		return g.save(ctx, passed(row, first))
	}
	second, err := g.attempt(ctx, retryMessages(msgs, first), p, persona)
	if err != nil {
		return g.save(ctx, fallback(row, p, persona, reasonModelError, true))
	}
	row = g.withAttempt(row, second)
	if len(second.probs) == 0 {
		return g.save(ctx, passed(row, second))
	}
	g.logFailure(p, first, second)
	return g.save(ctx, fallback(row, p, persona, reasonValidation, false))
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

// attempt makes one call, records its spend, then parses and validates the answer. A
// parse failure is a shape problem, so it is retried like any other.
func (g *Generator) attempt(ctx context.Context, msgs []llm.Message, p Pack, persona Persona) (attempt, error) {
	resp, err := g.LLM.Complete(ctx, llm.Request{
		Messages: msgs, JSONSchema: OutputSchema(), Temperature: g.Temperature, MaxTokens: reviewMaxTokens,
	})
	if err != nil {
		g.logger().Warn("coach: model call failed", slog.String("error", err.Error()))
		return attempt{}, err
	}
	if err := g.Caps.RecordSpend(ctx, resp); err != nil {
		g.logger().Error("coach: spend not recorded", slog.String("error", err.Error()), slog.Float64("usd", resp.CostUSD))
	}
	a := attempt{content: resp.Content, resp: resp}
	doc, err := ParseOutput(resp.Content)
	if err != nil {
		a.probs = []Problem{{Kind: ProblemShape, Value: err.Error()}}
		return a, nil
	}
	a.probs = g.Validator.Validate(doc, p, persona)
	return a, nil
}

// withAttempt adds an attempt's tokens and cost to the row and takes its slug, request id
// and problems as the last attempt's.
func (g *Generator) withAttempt(row StoredReview, a attempt) StoredReview {
	row.TokensIn += a.resp.Usage.PromptTokens
	row.TokensOut += a.resp.Usage.CompletionTokens
	row.CostUSD += a.resp.CostUSD
	row.CostEstimated = row.CostEstimated || a.resp.CostEstimated
	row.ModelSlug = cmpOr(a.resp.Model, g.Model)
	row.RequestID = a.resp.RequestID
	row.Validation = slices.Clone(a.probs)
	return row
}

func cmpOr(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

// passed is the row for an answer that validated; it parsed, so ParseOutput cannot fail.
func passed(row StoredReview, a attempt) StoredReview {
	doc, _ := ParseOutput(a.content)
	row.Doc, row.Source, row.Reason, row.Retryable = doc, SourceModel, "", false
	return row
}

// fallback is the rule-only rendering with why it was served.
func fallback(row StoredReview, p Pack, persona Persona, reason string, retryable bool) StoredReview {
	row.Doc = RenderRules(p, persona.Drills)
	row.Source, row.Reason, row.Retryable = SourceRules, reason, retryable
	return row
}

func (g *Generator) save(ctx context.Context, row StoredReview) (StoredReview, error) {
	if err := g.Store.SaveReview(ctx, row); err != nil {
		return row, fmt.Errorf("coach: store the generated review: %w", err)
	}
	return row, nil
}

// retryMessages is the first two messages, the model's answer and the problems listed back.
func retryMessages(msgs []llm.Message, first attempt) []llm.Message {
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
func (g *Generator) logFailure(p Pack, first, second attempt) {
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
