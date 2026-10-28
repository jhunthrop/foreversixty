package main

import (
	"context"
	"encoding/json"
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

// Problem is one validator finding, kept for the report. Value is model text, redacted on write.
type Problem struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

// Result is one fight x model.
type Result struct {
	Label     string  `json:"label"`
	Model     string  `json:"model"`
	Spec      string  `json:"spec"`
	Role      string  `json:"role"`
	Size      int     `json:"size"`
	Status    string  `json:"pack_status"`
	Outcome   Outcome `json:"outcome"`
	Note      string  `json:"note,omitempty"`
	Calls     int     `json:"calls"`
	TokensIn  int     `json:"tokens_in"`
	TokensOut int     `json:"tokens_out"`
	// Reasoning is the level the model ran at; ReasoningTokens is the part of TokensOut spent on it.
	Reasoning       string  `json:"reasoning"`
	ReasoningTokens int     `json:"reasoning_tokens"`
	CostUSD         float64 `json:"cost_usd"`
	CostEstimated   bool    `json:"cost_estimated"`
	LatencyMS       int64   `json:"latency_ms"`
	// Problems are the last answered call's validator findings; empty for a first-try pass.
	Problems []Problem `json:"problems,omitempty"`
	Pack     any       `json:"pack,omitempty"`
}

// Model is one model under test: the client that calls it and the slug it is recorded under.
type Model struct {
	Slug        string
	Client      llm.Client
	Temperature float64
	Reasoning   string // the level the client was built with; reported only
}

// Runner runs a built pack through a model with the generator's flow: the answer is parsed
// and validated, one retry lists the problems back, and nothing is stored or counted.
type Runner struct {
	Names *coach.NameDict
	Now   func() time.Time
}

// Run asks the model for b's review through coach.Generator.Attempt, the service's own flow
// (prompt, parse, validate, one retry), and maps what it made to a Result. The model's raw
// text (Attempt.Outputs) is never copied: only counts, problem kinds and validator quotes.
func (r Runner) Run(ctx context.Context, b Built, m Model) Result {
	res := baseResult(b, m)
	if b.Pack.Status != coach.StatusReady {
		res.Outcome, res.Note = OutcomeSkipped, "pack is "+string(b.Pack.Status)
		return res
	}
	res.Pack = b.Pack
	g := coach.Generator{LLM: m.Client, Model: m.Slug, Temperature: m.Temperature,
		Validator: coach.Validator{Names: r.Names, Roster: b.Roster}}
	start := r.Now()
	att, err := g.Attempt(ctx, b.Pack, b.Persona)
	res.LatencyMS = r.Now().Sub(start).Milliseconds()
	res.Calls, res.TokensIn, res.TokensOut = att.Calls, att.TokensIn, att.TokensOut
	res.ReasoningTokens = att.TokensReasoning
	res.CostUSD, res.CostEstimated = att.CostUSD, att.CostEstimated
	res.Problems = problemsOf(att.Validation)
	if err != nil {
		res.Outcome, res.Note = OutcomeModelError, err.Error()
		return res
	}
	res.Outcome = outcomeOf(att)
	return res
}

func outcomeOf(a coach.Attempt) Outcome {
	switch a.PassedOn {
	case 1:
		return OutcomePassFirst
	case 2:
		return OutcomePassRetry
	}
	return OutcomeFallback
}

func baseResult(b Built, m Model) Result {
	row := b.Target.Row
	return Result{Label: b.Target.Label, Model: m.Slug, Spec: row.Class + " " + row.Spec, Role: row.Role,
		Size: b.Target.Fight.Size, Status: string(b.Pack.Status), Reasoning: m.Reasoning}
}

func modelError(res Result, err error) Result {
	res.Outcome, res.Note = OutcomeModelError, err.Error()
	return res
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
