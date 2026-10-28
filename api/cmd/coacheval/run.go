package main

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

// Outcome is how one fight went through one model.
type Outcome string

const (
	OutcomePassFirst  Outcome = "pass_first"  // the first answer validated
	OutcomePassRetry  Outcome = "pass_retry"  // the retry validated
	OutcomeFallback   Outcome = "fallback"    // both answers failed validation: the service serves the rules
	OutcomeModelError Outcome = "model_error" // a call failed: the service serves the rules
	OutcomeSkipped    Outcome = "skipped"     // no model call: not reviewable, or no data
)

const (
	// evalMaxTokens mirrors the generator's per-answer bound (coach/generate.go reviewMaxTokens).
	evalMaxTokens = 1200
	roleAssistant = "assistant"
	roleUser      = "user"
)

// Problem is one validator finding, kept for the report. Value is model text, redacted on write.
type Problem struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

// Result is one fight x model.
type Result struct {
	Label         string    `json:"label"`
	Model         string    `json:"model"`
	Spec          string    `json:"spec"`
	Role          string    `json:"role"`
	Size          int       `json:"size"`
	Status        string    `json:"pack_status"`
	Outcome       Outcome   `json:"outcome"`
	Note          string    `json:"note,omitempty"`
	Calls         int       `json:"calls"`
	TokensIn      int       `json:"tokens_in"`
	TokensOut     int       `json:"tokens_out"`
	CostUSD       float64   `json:"cost_usd"`
	CostEstimated bool      `json:"cost_estimated"`
	LatencyMS     int64     `json:"latency_ms"`
	FirstProblems []Problem `json:"first_problems,omitempty"`
	LastProblems  []Problem `json:"last_problems,omitempty"`
	Pack          any       `json:"pack,omitempty"`
	Outputs       []string  `json:"outputs,omitempty"`
}

// Model is one model under test: the client that calls it and the slug it is recorded under.
type Model struct {
	Slug        string
	Client      llm.Client
	Temperature float64
}

// Runner runs a built pack through a model with the generator's flow: the answer is parsed
// and validated, one retry lists the problems back, and nothing is stored or counted.
type Runner struct {
	Names *coach.NameDict
	Now   func() time.Time
}

type answer struct {
	resp    llm.Response
	problem []coach.Problem
}

func (r Runner) Run(ctx context.Context, b Built, m Model) Result {
	res := baseResult(b, m)
	if b.Pack.Status != coach.StatusReady {
		res.Outcome, res.Note = OutcomeSkipped, "pack is "+string(b.Pack.Status)
		return res
	}
	res.Pack = b.Pack
	v := coach.Validator{Names: r.Names, Roster: b.Roster}
	msgs := coach.BuildPrompt(b.Pack, b.Persona)
	first, err := r.ask(ctx, m, msgs, b, v, &res)
	if err != nil {
		return modelError(res, err)
	}
	res.FirstProblems = problemsOf(first.problem)
	if len(first.problem) == 0 {
		res.Outcome = OutcomePassFirst
		return res
	}
	retry := append(slices.Clone(msgs),
		llm.Message{Role: roleAssistant, Content: first.resp.Content},
		llm.Message{Role: roleUser, Content: coach.RetryMessage(first.problem)})
	second, err := r.ask(ctx, m, retry, b, v, &res)
	if err != nil {
		return modelError(res, err)
	}
	res.LastProblems = problemsOf(second.problem)
	res.Outcome = OutcomeFallback
	if len(second.problem) == 0 {
		res.Outcome = OutcomePassRetry
	}
	return res
}

func baseResult(b Built, m Model) Result {
	row := b.Target.Row
	return Result{Label: b.Target.Label, Model: m.Slug, Spec: row.Class + " " + row.Spec, Role: row.Role,
		Size: b.Target.Fight.Size, Status: string(b.Pack.Status)}
}

func modelError(res Result, err error) Result {
	res.Outcome, res.Note = OutcomeModelError, err.Error()
	return res
}

// ask makes one call, adds its usage to res and validates the answer. A parse failure is a
// shape problem, as in the generator.
func (r Runner) ask(ctx context.Context, m Model, msgs []llm.Message, b Built, v coach.Validator, res *Result) (answer, error) {
	start := r.Now()
	resp, err := m.Client.Complete(ctx, llm.Request{Messages: msgs, JSONSchema: coach.OutputSchema(),
		Temperature: m.Temperature, MaxTokens: evalMaxTokens})
	res.LatencyMS += r.Now().Sub(start).Milliseconds()
	if err != nil {
		return answer{}, fmt.Errorf("model call: %w", err)
	}
	res.Calls++
	res.TokensIn += resp.Usage.PromptTokens
	res.TokensOut += resp.Usage.CompletionTokens
	res.CostUSD += resp.CostUSD
	res.CostEstimated = res.CostEstimated || resp.CostEstimated
	res.Outputs = append(res.Outputs, resp.Content)
	doc, err := coach.ParseOutput(resp.Content)
	if err != nil {
		return answer{resp: resp, problem: []coach.Problem{{Kind: coach.ProblemShape, Value: err.Error()}}}, nil
	}
	return answer{resp: resp, problem: v.Validate(doc, b.Pack, b.Persona)}, nil
}

func problemsOf(in []coach.Problem) []Problem {
	out := make([]Problem, len(in))
	for i, p := range in {
		out[i] = Problem{Kind: p.Kind, Value: p.Value}
	}
	return out
}

// packJSON is the pack as the model sees it, for the identifier assertion and the report.
func packJSON(p coach.Pack) string {
	b, _ := json.Marshal(p)
	return string(b)
}
