package main

import (
	"strings"
	"testing"
)

// TestWriteProbeRendersEveryRowAndSurfacesWrongAndMissingLines covers
// the probe table rendering: a removal row that hurts (a wrong line),
// an insertion row that helps (a missing line), a no-effect row, and
// a sim-error row (which must render without a delta and must not be
// counted as either wrong or missing).
func TestWriteProbeRendersEveryRowAndSurfacesWrongAndMissingLines(t *testing.T) {
	r := report{
		opts: options{confirmIterations: 800},
		probe: []probeRow{
			{Kind: "remove", Label: "spell 100", Delta: -12.3, Err: 2.0, Verdict: verdictHurts},
			{Kind: "remove", Label: "spell 200", Delta: 15.0, Err: 2.0, Verdict: verdictHelps},
			{Kind: "remove", Label: "spell 300", Delta: 0.4, Err: 2.0, Verdict: verdictNoEffect},
			{Kind: "insert", Label: "Missing Spell", Delta: 9.0, Err: 1.0, Verdict: verdictHelps},
			{Kind: "insert", Label: "Useless Spell", Delta: -1.0, Err: 2.0, Verdict: verdictNoEffect},
			{Kind: "insert", Label: "Broken Spell", Verdict: verdictError, Error: "the engine reported an error: boom"},
		},
	}
	var b strings.Builder
	r.writeProbe(&b)
	out := b.String()

	for _, want := range []string{
		"| remove | spell 100 | -12.3 | 2.0 | hurts |",
		"| remove | spell 200 | +15.0 | 2.0 | helps |",
		"| remove | spell 300 | +0.4 | 2.0 | no effect |",
		"| insert | Missing Spell | +9.0 | 1.0 | helps |",
		"| insert | Useless Spell | -1.0 | 2.0 | no effect |",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("probe table missing row %q in:\n%s", want, out)
		}
	}
	if !strings.Contains(out, "Broken Spell") || !strings.Contains(out, "sim error") || !strings.Contains(out, "boom") {
		t.Errorf("probe table did not render the sim-error row with its message:\n%s", out)
	}

	if !strings.Contains(out, "**Wrong lines**") {
		t.Errorf("report did not render a wrong-lines line:\n%s", out)
	}
	if !strings.Contains(out, "**Missing lines**") {
		t.Errorf("report did not render a missing-lines line:\n%s", out)
	}

	wrong, missing := r.wrongAndMissing()
	if len(wrong) != 1 || wrong[0] != "spell 100" {
		t.Errorf("wrongAndMissing() wrong = %v, want [\"spell 100\"]", wrong)
	}
	if len(missing) != 1 || missing[0] != "Missing Spell" {
		t.Errorf("wrongAndMissing() missing = %v, want [\"Missing Spell\"]", missing)
	}
}

func TestFirstLine(t *testing.T) {
	if got := firstLine("one\ntwo\nthree"); got != "one" {
		t.Errorf("firstLine = %q, want %q", got, "one")
	}
	if got := firstLine("just one line"); got != "just one line" {
		t.Errorf("firstLine = %q, want %q", got, "just one line")
	}
}

func TestWriteProbeWithNoFindingsSaysSo(t *testing.T) {
	r := report{probe: []probeRow{
		{Kind: "remove", Label: "spell 1", Delta: 0.1, Err: 2.0, Verdict: verdictNoEffect},
	}}
	var b strings.Builder
	r.writeProbe(&b)
	out := b.String()
	if !strings.Contains(out, "No wrong lines and no missing lines") {
		t.Errorf("expected the no-findings message, got:\n%s", out)
	}
}
