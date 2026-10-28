package main

import (
	"fmt"
	"sort"
	"strings"
)

// RunMeta is what the report says about the run beside the results.
type RunMeta struct {
	Source  string // "corpus dir" or "database", never a URL
	Fake    bool
	Planned int // targets chosen
}

// RenderReport is the markdown report: the build-notes rows ready to paste, the per-model
// table, the problem kinds and one line per fight and model.
func RenderReport(results []Result, meta RunMeta) string {
	stats := statsByModel(results)
	var b strings.Builder
	b.WriteString("# Log Coach pre-launch evaluation\n\n")
	if meta.Fake {
		b.WriteString("**Fake run: llm.Fake answers, not a real model. Do not read these numbers.**\n\n")
	}
	fmt.Fprintf(&b, "Source: %s. Fights chosen: %d.\n\n", meta.Source, meta.Planned)
	b.WriteString("## Build-notes rows\n\n| Item | Result |\n|---|---|\n")
	for _, s := range stats {
		fmt.Fprintf(&b, "| Evaluation: %s pass rate, mean tokens in/out | %s |\n", s.Model, buildNotesCell(s))
	}
	b.WriteString("\n## Per model\n\n")
	b.WriteString("| Model | Fights | Pass first | Pass after retry | Fallback | Validator pass | Tokens in mean/p90 | Tokens out mean/p90 | Cost total | Cost per fight | Latency mean/p90 |\n")
	b.WriteString("|---|---|---|---|---|---|---|---|---|---|---|\n")
	for _, s := range stats {
		b.WriteString(statsRow(s))
	}
	b.WriteString("\nCost marked ~ is estimated from tokens and the configured prices; the gateway dashboard holds the billed figure.\n")
	writeProblemKinds(&b, stats)
	writeFightLines(&b, results)
	return b.String()
}

func buildNotesCell(s ModelStats) string {
	return fmt.Sprintf("%s (%d/%d), %s in / %s out", pct(s.PassRate()), s.Passed(), s.Fights,
		thousands(s.MeanIn), thousands(s.MeanOut))
}

func statsRow(s ModelStats) string {
	return fmt.Sprintf("| %s | %d (%d skipped) | %s | %s | %s | %s | %s / %s | %s / %s | %s | %s | %.1fs / %.1fs |\n",
		s.Model, s.Fights, s.Skipped, pct(s.FirstTryRate()), pct(ratio(s.PassRetry, s.Fights)), pct(s.FallbackRate()),
		pct(s.PassRate()), thousands(s.MeanIn), thousands(s.P90In), thousands(s.MeanOut), thousands(s.P90Out),
		usd(s.TotalCostUSD, s.CostEstimated), usd(s.MeanCostUSD, s.CostEstimated),
		s.MeanLatencyMS/msPerSecond, s.P90LatencyMS/msPerSecond)
}

const msPerSecond = 1000.0

func pct(x float64) string { return fmt.Sprintf("%.0f%%", x*percent) }

func thousands(x float64) string { return fmt.Sprintf("%.0f", x) }

func usd(x float64, estimated bool) string {
	mark := ""
	if estimated {
		mark = "~"
	}
	return fmt.Sprintf("%s$%.4f", mark, x)
}

func writeProblemKinds(b *strings.Builder, stats []ModelStats) {
	b.WriteString("\n## Final-answer problems by kind\n\n")
	for _, s := range stats {
		kinds := make([]string, 0, len(s.ProblemKinds))
		for k, n := range s.ProblemKinds {
			kinds = append(kinds, fmt.Sprintf("%s %d", k, n))
		}
		sort.Strings(kinds)
		if len(kinds) == 0 {
			kinds = []string{"none"}
		}
		fmt.Fprintf(b, "- %s: %s\n", s.Model, strings.Join(kinds, ", "))
	}
}

func writeFightLines(b *strings.Builder, results []Result) {
	b.WriteString("\n## Per fight\n\n| Fight | Spec | Size | Model | Outcome | Tokens in/out | Problems (last answer) |\n|---|---|---|---|---|---|---|\n")
	for _, r := range results {
		fmt.Fprintf(b, "| %s | %s | %d | %s | %s | %d / %d | %s |\n", r.Label, r.Spec, r.Size, r.Model, outcomeCell(r),
			r.TokensIn, r.TokensOut, problemCell(r.Problems))
	}
}

func outcomeCell(r Result) string {
	if r.Note != "" {
		return string(r.Outcome) + " (" + r.Note + ")"
	}
	return string(r.Outcome)
}

func problemCell(ps []Problem) string {
	if len(ps) == 0 {
		return ""
	}
	parts := make([]string, len(ps))
	for i, p := range ps {
		parts[i] = p.Kind + ": " + strings.ReplaceAll(p.Value, "|", "/")
	}
	return strings.Join(parts, "; ")
}
