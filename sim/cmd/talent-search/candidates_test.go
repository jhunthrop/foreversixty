package main

import (
	"errors"
	"strings"
	"testing"
)

// testCredits: A and C deal damage, D a lot; B, E and F nothing.
func testCredits() map[int]credit {
	return map[int]credit{
		1: {PerPoint: 1, Damage: true},
		2: {},
		3: {PerPoint: 2, Damage: true},
		4: {PerPoint: 10, Damage: true},
		5: {},
		6: {},
	}
}

func TestSwapsAreLegalAtTheSameTotalAndMoveOnlyNonDamagePoints(t *testing.T) {
	tr := testTrees()
	guide := build{1: 3, 2: 2, 5: 3}
	got := swaps(tr, guide, testCredits(), "")
	if len(got) == 0 {
		t.Fatal("no swaps generated")
	}
	for _, c := range got {
		if err := tr.legal(c.Build, guide.points()); err != nil {
			t.Errorf("%s: illegal swap %v: %v", c.Label, c.Build, err)
		}
		if c.Build[1] < guide[1] {
			t.Errorf("%s: moved points out of a damage talent", c.Label)
		}
	}
	// Moving all of E into C would need 5 points in tree 0's tier 0
	// first - the guide has exactly 5 there, so 2..3 E points into C
	// is legal, but C is capped at 3.
	for _, c := range got {
		if c.Build[3] > 3 {
			t.Errorf("%s: %d points in a 3-rank talent", c.Label, c.Build[3])
		}
	}
}

func TestSwapsRespectTierGatesWhenPointsLeave(t *testing.T) {
	tr := testTrees()
	// B's 2 points hold C's tier-1 gate (A 3 + B 2 = 5): taking one out
	// of B would strand C, so no swap may do it.
	guide := build{1: 3, 2: 2, 3: 1, 5: 1}
	for _, c := range swaps(tr, guide, testCredits(), "") {
		if c.Build[2] < 2 && c.Build[1]+c.Build[2] < 5 {
			t.Errorf("%s: stranded C: %v", c.Label, c.Build)
		}
	}
}

func TestFillSpendsTheWholeBudgetLegallyAndReachesTheBigTalent(t *testing.T) {
	tr := testTrees()
	s := scorer{credits: testCredits()}
	b := s.fill(tr, build{}, 11, -1)
	if err := tr.legal(b, 11); err != nil {
		t.Fatalf("fill = %v: %v", b, err)
	}
	if b[4] != 1 {
		t.Fatalf("fill = %v: did not bundle its way to D (10 DPS)", b)
	}
}

func TestFillDeepTreeFirst(t *testing.T) {
	tr := testTrees()
	s := scorer{credits: testCredits()}
	b := s.fill(tr, build{}, 12, 1)
	if tr.treePoints(b)[1] != 7 {
		t.Fatalf("deep tree Two: %v puts %d points there, want all 7 it holds", b, tr.treePoints(b)[1])
	}
	if err := tr.legal(b, 12); err != nil {
		t.Fatal(err)
	}
}

func TestScorerPrefersGuideTalentsAmongEqualCredits(t *testing.T) {
	tr := testTrees()
	s := scorer{credits: testCredits(), prefer: build{5: 1}}
	id, ok := s.bestPoint(tr, build{1: 5}, func(id int) bool { return id == 2 || id == 5 })
	if !ok || id != 5 {
		t.Fatalf("bestPoint = %d, want the guide's own E over B", id)
	}
}

func TestStripKeepsGatesAndDamage(t *testing.T) {
	tr := testTrees()
	credits := testCredits()
	b := strip(tr, build{1: 3, 2: 2, 3: 1, 5: 3}, func(id int) bool { return !credits[id].Damage })
	if b[5] != 0 {
		t.Fatalf("strip left E: %v", b)
	}
	if b[2] != 2 || b[3] != 1 || b[1] != 3 {
		t.Fatalf("strip = %v, want A and C kept and B kept as C's gate", b)
	}
	if err := tr.legal(b, b.points()); err != nil {
		t.Fatal(err)
	}
}

func TestGenerateIsLegalDedupedAndKeepsUnmodeledInTheModeledRespend(t *testing.T) {
	tr := testTrees()
	guide := build{1: 5, 2: 2, 5: 3}
	modeled := map[int]bool{1: true, 2: true, 3: true, 4: true, 6: true} // E is unmodeled
	pool := generate(tr, guide, testCredits(), modeled, guide.points(), 2)
	seen := map[string]bool{tr.key(guide): true}
	var modeledRespend *candidate
	for i, c := range pool {
		if err := tr.legal(c.Build, guide.points()); err != nil {
			t.Errorf("%s: %v", c.Label, err)
		}
		if seen[tr.key(c.Build)] {
			t.Errorf("%s: duplicate", c.Label)
		}
		seen[tr.key(c.Build)] = true
		if c.Label == "guide, modeled non-damage points re-spent" {
			modeledRespend = &pool[i]
		}
	}
	if modeledRespend == nil || modeledRespend.Build[5] != 3 || modeledRespend.Build[3] != 2 {
		t.Fatalf("the modeled-only re-spend must keep unmodeled E: %+v", modeledRespend)
	}
}

func TestDedupeDropsTheGuideAndRepeats(t *testing.T) {
	tr := testTrees()
	guide := build{1: 5}
	got := dedupe(tr, guide, []candidate{
		{Label: "same as guide", Build: build{1: 5}},
		{Label: "x", Build: build{1: 4, 2: 1}},
		{Label: "x again", Build: build{2: 1, 1: 4}},
		{Label: "y", Build: build{1: 4, 5: 1}},
	})
	if len(got) != 2 || got[0].Label != "x" || got[1].Label != "y" {
		t.Fatalf("dedupe = %+v", got)
	}
}

func TestCapPrefersStructuralThenDistinctThenFills(t *testing.T) {
	pool := []candidate{
		{Label: "swap best", Build: build{1: 4, 2: 1}, Estimate: 5},
		{Label: "swap near-duplicate", Build: build{1: 4, 5: 1}, Estimate: 4},
		{Label: "swap distinct", Build: build{2: 5}, Estimate: 3},
		{Label: "archetype", Build: build{5: 5}, Structural: true},
	}
	got := capCandidates(pool, 3)
	want := []string{"archetype", "swap best", "swap distinct"}
	for i, w := range want {
		if got[i].Label != w {
			t.Fatalf("cap = %v, want %v", labels(got), want)
		}
	}
	if all := capCandidates(pool, 10); len(all) != 4 || all[3].Label != "swap near-duplicate" {
		t.Fatalf("with room, the near-duplicate fills in last: %v", labels(all))
	}
}

func labels(cs []candidate) []string {
	out := make([]string, len(cs))
	for i, c := range cs {
		out[i] = c.Label
	}
	return out
}

// fakeRun scores a build as the sum of a fixed per-talent value, with
// a fixed error, so probe and evaluate can be checked without the engine.
func fakeRun(values map[int]float64) dpsFunc {
	return func(b build, _ int) (estimate, error) {
		dps := 100.0
		for id, r := range b {
			dps += float64(r) * values[id]
		}
		return estimate{Mean: dps, Err: 0.5}, nil
	}
}

func TestProbeCreditsMeasuresTakenAndUntakenTalents(t *testing.T) {
	tr := testTrees()
	credits, err := probeCredits(tr, build{1: 2, 5: 1}, fakeRun(map[int]float64{1: 3, 3: 2}), 10)
	if err != nil {
		t.Fatal(err)
	}
	if c := credits[1]; c.PerPoint != 3 || !c.Damage {
		t.Errorf("taken A: %+v", c)
	}
	if c := credits[3]; c.PerPoint != 2 || !c.Damage {
		t.Errorf("untaken C: %+v", c)
	}
	if c := credits[5]; c.PerPoint != 0 || c.Damage {
		t.Errorf("taken E, worth nothing: %+v", c)
	}
}

func TestEvaluateReportsDeltaAndSignificance(t *testing.T) {
	run := fakeRun(map[int]float64{1: 1, 3: 2})
	guide := build{1: 5}
	ev, err := evaluate(guide, []candidate{
		{Label: "big", Build: build{1: 2, 3: 3}},
		{Label: "tiny", Build: build{1: 5, 2: 0}},
	}, run, 10, 20, 5, func(build) bool { return true })
	if err != nil {
		t.Fatal(err)
	}
	best := ev.Final[0]
	if best.Label != "big" || best.Delta != 3 || !best.Significant {
		t.Fatalf("best = %+v", best)
	}
	if ev.Guide.Label != "guide" || ev.Guide.Delta != 0 {
		t.Fatalf("guide = %+v", ev.Guide)
	}
}

func TestEvaluateAlwaysFinalsTheBestCleanCandidate(t *testing.T) {
	run := fakeRun(map[int]float64{1: 1, 2: 2, 3: 3})
	ev, err := evaluate(build{}, []candidate{
		{Label: "dirty best", Build: build{3: 3}},
		{Label: "dirty second", Build: build{2: 3}},
		{Label: "clean", Build: build{1: 1}},
	}, run, 10, 20, 1, func(b build) bool { return b[3] == 0 && b[2] == 0 })
	if err != nil {
		t.Fatal(err)
	}
	got := strings.Join(labelsOf(ev.Final), ",")
	if got != "dirty best,clean,guide" {
		t.Fatalf("finals = %s, want the top 1, the best clean one and the guide", got)
	}
}

func labelsOf(rs []result) []string {
	out := make([]string, len(rs))
	for i, r := range rs {
		out[i] = r.Label
	}
	return out
}

func TestEvaluatePropagatesSimErrors(t *testing.T) {
	boom := errors.New("boom")
	_, err := evaluate(build{}, nil, func(build, int) (estimate, error) { return estimate{}, boom }, 1, 1, 1, func(build) bool { return true })
	if !errors.Is(err, boom) {
		t.Fatalf("err = %v, want boom", err)
	}
}
