package main

import (
	"strings"
	"testing"
	"time"
)

// reportFixture is a spec in tree One (index 0) whose guide takes A, B and
// E. B is a talent the engine neither models nor credits, so a variant that
// drops it sims as free.
func reportFixture() report {
	guide := build{1: 3, 2: 2, 5: 3}
	drops := build{1: 5, 5: 3} // moves B's points into A
	return report{
		opts: options{level: 60, faction: "alliance", preset: "bare", screenIters: 300, finalIters: 1000, probeIters: 200, seed: 7, keep: "B, E"},
		in: inputs{
			setup:       simSetup{spec: specInfo{Name: "Fire", ClassSlug: "mage", TreeIndex: 0}, band: bisBand{Race: "Gnome"}, trees: testTrees()},
			clientBuild: "1.2.3",
			guide:       guide,
			guideCode:   "3/0/0 (client 1.0.0)",
			modeled:     map[int]bool{1: true},
		},
		credits: map[int]credit{
			1: {PerPoint: 4, Diff: 40, Err: 2, Damage: true},
			2: {Diff: 0, Err: 2},
			5: {Diff: 0, Err: 2},
		},
		eval: evaluation{
			Guide: result{candidate: candidate{Label: "guide", Build: guide}, DPS: estimate{Mean: 100, Err: 1}},
			Final: []result{
				{candidate: candidate{Label: "no B", Build: drops}, DPS: estimate{Mean: 110, Err: 1}, Delta: 10, DeltaErr: 1.4, Significant: true},
				{candidate: candidate{Label: "guide", Build: guide}, DPS: estimate{Mean: 100, Err: 1}},
			},
			Screened: []result{
				{candidate: candidate{Label: "no B", Build: drops, Estimate: 9}, DPS: estimate{Mean: 110, Err: 1}, Delta: 10, DeltaErr: 1.4},
				{candidate: candidate{Label: "guide", Build: guide}, DPS: estimate{Mean: 100, Err: 1}},
			},
		},
		poolSize: 12,
		elapsed:  3 * time.Second,
	}
}

func TestRemovedUnmodeledNamesTheDroppedBlindSpots(t *testing.T) {
	r := reportFixture()
	got := removedUnmodeled(r.in, r.credits, build{1: 5, 5: 3})
	if len(got) != 1 || got[0] != "B (-2)" {
		t.Fatalf("got %v", got)
	}
	if got := removedUnmodeled(r.in, r.credits, r.in.guide); len(got) != 0 {
		t.Fatalf("the guide drops nothing, got %v", got)
	}
	// a damage-credited talent is not a blind spot even when dropped
	if got := removedUnmodeled(r.in, r.credits, build{2: 2, 5: 3}); len(got) != 0 {
		t.Fatalf("A has damage credit, got %v", got)
	}
}

func TestBestCleanSkipsWinnersThatDropUnmodeledTalents(t *testing.T) {
	r := reportFixture()
	res, ok := r.bestClean()
	if !ok || res.Label != "guide" {
		t.Fatalf("got %+v %v", res, ok)
	}
	r.eval.Final = r.eval.Final[:1]
	if _, ok := r.bestClean(); ok {
		t.Fatal("no clean finalist exists")
	}
}

func TestWriteVerdictThreeOutcomes(t *testing.T) {
	r := reportFixture()
	var b strings.Builder
	r.writeVerdict(&b, r.eval.Final[0])
	if !strings.Contains(b.String(), "no B beats the guide by +10.0 DPS (+10.0%), beyond the combined error ±1.4") {
		t.Errorf("significant: %s", b.String())
	}
	b.Reset()
	weak := r.eval.Final[0]
	weak.Significant = false
	r.writeVerdict(&b, weak)
	if !strings.Contains(b.String(), "no variant beats the guide beyond error") {
		t.Errorf("insignificant: %s", b.String())
	}
	b.Reset()
	r.writeVerdict(&b, r.eval.Guide)
	if !strings.Contains(b.String(), "the guide build is the best build found") {
		t.Errorf("guide: %s", b.String())
	}
}

func TestWriteCleanVerdictOutcomes(t *testing.T) {
	r := reportFixture()
	var b strings.Builder
	r.writeCleanVerdict(&b)
	if !strings.Contains(b.String(), "the guide is the best") {
		t.Errorf("guide clean: %s", b.String())
	}

	clean := result{candidate: candidate{Label: "more A", Build: build{1: 4, 2: 2, 5: 3}}, Delta: 6, DeltaErr: 1, Significant: true}
	r.eval.Final = []result{clean, r.eval.Guide}
	b.Reset()
	r.writeCleanVerdict(&b)
	if !strings.Contains(b.String(), "more A, +6.0 DPS (±1.0 combined), beyond error") {
		t.Errorf("significant clean: %s", b.String())
	}

	clean.Significant = false
	r.eval.Final = []result{clean, r.eval.Guide}
	b.Reset()
	r.writeCleanVerdict(&b)
	if !strings.Contains(b.String(), "best is more A at +6.0 DPS (±1.0 combined), within error") {
		t.Errorf("within error clean: %s", b.String())
	}

	r.eval.Final = r.eval.Final[:1]
	r.eval.Final[0].Build = build{1: 5, 5: 3} // drops B
	b.Reset()
	r.writeCleanVerdict(&b)
	if b.Len() != 0 {
		t.Errorf("no clean finalist must print nothing: %s", b.String())
	}
}

func TestEngineGapsListOwnTreeOrGuideTalentsOnly(t *testing.T) {
	r := reportFixture()
	r.credits = map[int]credit{} // nothing credited: every unmodeled talent is a candidate gap
	r.in.modeled = map[int]bool{1: true, 3: true}
	got := strings.Join(r.engineGaps(), "|")
	// tree One (spec's own): B and D are unmodeled; tree Two: only E (the guide takes it), not F.
	for _, want := range []string{"B (One, tier 0)", "D (One, tier 2)", "E (Two, tier 0)"} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in %s", want, got)
		}
	}
	for _, not := range []string{"A (", "C (", "F ("} {
		if strings.Contains(got, not) {
			t.Errorf("%q must not be listed in %s", not, got)
		}
	}
	var b strings.Builder
	r.writeEngineGaps(&b)
	if !strings.Contains(b.String(), "- B (One, tier 0)") {
		t.Errorf("section: %s", b.String())
	}

	r.in.modeled = map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true}
	b.Reset()
	r.writeEngineGaps(&b)
	if !strings.Contains(b.String(), "None:") {
		t.Errorf("empty section: %s", b.String())
	}
}

func TestUnmodeledInAndEngineStatus(t *testing.T) {
	r := reportFixture()
	got := r.unmodeledIn(build{2: 2, 5: 1, 1: 1})
	if strings.Join(got, "|") != "B (2)|E (1)" {
		t.Fatalf("got %v", got)
	}
	if r.engineStatus(1, r.credits[1]) != "damage" || r.engineStatus(2, r.credits[2]) != "unmodeled" {
		t.Fatal("engineStatus mislabelled")
	}
}

func TestMarkdownRendersEverySection(t *testing.T) {
	r := reportFixture()
	r.in.guideIssue = "spends 8 points"
	out := r.markdown()
	for _, want := range []string{
		"# Talent search: Fire, level 60",
		"12 candidates generated, 1 screened at 300 iterations, the best 1 plus the guide re-simmed at 1000 iterations",
		"Run time 3s",
		"**Verdict: no B beats the guide",
		"- The guide is not a full legal build in 1.2.3: spends 8 points",
		"- Kept at the guide's rank or more in every candidate (-keep): B, E",
		"## Finalists",
		"| 1 | no B | 5/3 | 110.0 | 1.0 | +10.0 ± 1.4 | yes |",
		"| B (-2) |", // drops unmodeled guide talents column
		"## Engine gaps",
		"## What the engine credits each talent with",
		"| One | A | 0 | 3 | 5 | +4.00 | +40.0 ± 2.0 | damage |",
		"## Every screened candidate (300 iterations)",
		"| no B | 5/3 | 110.0 | +10.0 ± 1.4 | +9.0 |",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("report missing %q", want)
		}
	}
	if strings.Contains(out, "Winner keeping every unmodeled guide talent") {
		t.Error("the guide is the clean winner, so no separate clean-winner line is due")
	}

	r.eval.Final = []result{{candidate: candidate{Label: "more A", Build: build{1: 4, 2: 2, 5: 3}}, DPS: estimate{Mean: 106, Err: 1}, Delta: 6, DeltaErr: 1, Significant: true}, r.eval.Guide}
	r.opts.keep = ""
	out = r.markdown()
	if !strings.Contains(out, "Winner keeping every unmodeled guide talent: `FS1:1.2.3:mage:Gnome:") {
		t.Error("a clean non-guide winner must be named")
	}
	if strings.Contains(out, "(-keep)") {
		t.Error("no keep list, no keep line")
	}
}

func TestYesNoAndJoinOrDash(t *testing.T) {
	if yesNo(true) != "yes" || yesNo(false) != "no" {
		t.Error("yesNo")
	}
	if joinOrDash(nil) != "-" || joinOrDash([]string{"a", "b"}) != "a, b" {
		t.Error("joinOrDash")
	}
}

func TestSummaryAndAbs64(t *testing.T) {
	if got := testTrees().summary(build{1: 5, 3: 1, 5: 2}); got != "6/2" {
		t.Errorf("summary = %q", got)
	}
	if abs64(-2.5) != 2.5 || abs64(2.5) != 2.5 || abs64(0) != 0 {
		t.Error("abs64")
	}
}

func TestHoldsKeptNeedsEveryKeptTalentAtGuideRank(t *testing.T) {
	guide := build{1: 3, 2: 2}
	keep := map[int]bool{1: true, 2: true}
	if !holdsKept(build{1: 3, 2: 4}, guide, keep) {
		t.Error("equal-or-more rejected")
	}
	if holdsKept(build{1: 3, 2: 1}, guide, keep) {
		t.Error("a kept talent below the guide was accepted")
	}
	if !holdsKept(build{}, guide, nil) {
		t.Error("an empty keep list protects nothing")
	}
	pool := []candidate{{Label: "ok", Build: build{1: 3, 2: 2}}, {Label: "bad", Build: build{1: 2, 2: 2}}}
	if got := withoutBelowGuide(pool, guide, keep); len(got) != 1 || got[0].Label != "ok" {
		t.Errorf("got %+v", got)
	}
}
