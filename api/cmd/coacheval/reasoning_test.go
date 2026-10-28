package main

import (
	"io"
	"strings"
	"testing"
)

func TestParseOptions_Reasoning(t *testing.T) {
	env := func(string) string { return "x" }
	base := []string{"-corpus", "c", "-out", "o", "-model", "a", "-model", "b"}
	o, err := ParseOptions(append([]string{"-reasoning", "b=low"}, base...), env, io.Discard)
	if err != nil || o.ReasoningOf("a") != "off" || o.ReasoningOf("b") != "low" {
		t.Fatalf("%+v %v", o.Reasoning, err)
	}
	for name, arg := range map[string]string{"bad level": "b=xhigh", "no level": "b", "empty": "=low"} {
		if _, err := ParseOptions(append([]string{"-reasoning", arg}, base...), env, io.Discard); err == nil {
			t.Fatalf("%s: accepted", name)
		}
	}
}

func TestReport_ReasoningColumnAndNotes(t *testing.T) {
	results := []Result{
		{Label: "f01", Model: "a", Reasoning: "low", Outcome: OutcomePassFirst, TokensIn: 2000, TokensOut: 600, ReasoningTokens: 200},
		{Label: "f02", Model: "a", Reasoning: "low", Outcome: OutcomePassFirst, TokensIn: 2000, TokensOut: 800, ReasoningTokens: 400},
		{Label: "f01", Model: "b", Outcome: OutcomePassFirst, TokensIn: 2000, TokensOut: 100},
	}
	md := RenderReport(results, RunMeta{Planned: 2})
	for _, want := range []string{
		"| Evaluation: a (reasoning low) pass rate, mean tokens in/out |",
		"| Evaluation: b (reasoning off) pass rate, mean tokens in/out |",
		"| Model | Reasoning | Fights |",
		"| a | low | 2 (0 skipped) |",
		"| b | off | 1 (0 skipped) |",
		"700 (300) / 800",
		"100 (0) / 100",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("missing %q in\n%s", want, md)
		}
	}
}
