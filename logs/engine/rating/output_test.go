// logs/engine/rating/output_test.go
package rating

import "testing"

func TestScoreOutputUsesTheAbsoluteCapWhenTheBracketLacksSamples(t *testing.T) {
	fight := fixtureSummary()
	fight.Roster[0].ExecutionScore = f64(0.97)
	row, ok := rosterRow(fight, fixturePlayer)
	if !ok {
		t.Fatal("fixture roster row not found")
	}
	src := &fakePercentiles{} // no placements configured: every Placement call answers ok=false
	c := scoreOutput(fight, row, Bracket{}, src)

	if c.Excluded {
		t.Fatalf("Output excluded (%s), want the absolute-cap path (an execution score always has an absolute standard)", c.Reason)
	}
	if c.Basis != BasisAbsolute {
		t.Fatalf("Basis = %q, want %q", c.Basis, BasisAbsolute)
	}
	if want := 97.0; c.Score != want {
		t.Fatalf("Score = %v, want %v (0.97 * 100)", c.Score, want)
	}
	if len(src.recorded[ComponentNameOutput]) != 1 || src.recorded[ComponentNameOutput][0] != 0.97 {
		t.Fatalf("recorded output placements = %v, want [0.97]", src.recorded[ComponentNameOutput])
	}
}

func TestScoreOutputCapsAnExecutionScoreAboveOneAtOneHundred(t *testing.T) {
	fight := fixtureSummary()
	fight.Roster[0].ExecutionScore = f64(1.30) // a lucky pull: 130% of the sim's mean
	row, _ := rosterRow(fight, fixturePlayer)
	src := &fakePercentiles{}
	c := scoreOutput(fight, row, Bracket{}, src)

	if c.Basis != BasisAbsolute {
		t.Fatalf("Basis = %q, want %q", c.Basis, BasisAbsolute)
	}
	if c.Score != 100 {
		t.Fatalf("Score = %v, want 100 (spec §1.3: capped at 100, not scaled past it -- variance belongs to "+
			"the percentile path, not the absolute one)", c.Score)
	}
}

func TestScoreOutputExcludesTheRawFallbackWhenTheBracketLacksSamples(t *testing.T) {
	fight := fixtureSummary()
	fight.Roster[0].ExecutionScore = nil // no validated sim for this spec: falls back to raw DPS/HPS
	fight.Roster[0].DPS = 750
	row, _ := rosterRow(fight, fixturePlayer)
	src := &fakePercentiles{} // ok=false for every bracket
	c := scoreOutput(fight, row, Bracket{}, src)

	if !c.Excluded {
		t.Fatalf("Excluded = false, want true: raw DPS has no absolute standard (spec §1.3: unbounded and gear-dependent)")
	}
	if c.Reason != ReasonNotEnoughSamples {
		t.Fatalf("Reason = %q, want %q", c.Reason, ReasonNotEnoughSamples)
	}
	if len(src.recorded[ComponentNameOutput]) != 1 || src.recorded[ComponentNameOutput][0] != 750 {
		t.Fatalf("recorded output placements = %v, want [750] (the raw DPS value, not the execution score)", src.recorded[ComponentNameOutput])
	}
}

func TestScoreOutputExcludesOnAWipeRegardlessOfExecutionScore(t *testing.T) {
	fight := fixtureSummary()
	fight.Kill = false
	fight.Roster[0].ExecutionScore = f64(0.97)
	row, _ := rosterRow(fight, fixturePlayer)
	src := &fakePercentiles{placements: map[string]fakePlacement{ComponentNameOutput: {pct: 0.9, n: 100, ok: true}}}
	c := scoreOutput(fight, row, Bracket{}, src)

	if !c.Excluded || c.Reason != ReasonWipe {
		t.Fatalf("Excluded=%v Reason=%q, want excluded with %q", c.Excluded, c.Reason, ReasonWipe)
	}
	if len(src.recorded) != 0 {
		t.Fatalf("a wipe must return before ever calling Placement, recorded = %v", src.recorded)
	}
}
