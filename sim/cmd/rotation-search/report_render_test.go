package main

import (
	"math"
	"strings"
	"testing"
	"time"
)

func TestClassifyVerdictUsesTheCombinedErrorBand(t *testing.T) {
	cases := []struct {
		delta, err float64
		want       verdict
	}{
		{5, 2, verdictHelps},
		{-5, 2, verdictHurts},
		{2, 2, verdictNoEffect},
		{-2, 2, verdictNoEffect},
		{0, 0, verdictNoEffect},
	}
	for _, c := range cases {
		if got := classifyVerdict(c.delta, c.err); got != c.want {
			t.Errorf("classifyVerdict(%v, %v) = %q, want %q", c.delta, c.err, got, c.want)
		}
	}
}

func TestCombinedErrIsTheHypotenuse(t *testing.T) {
	if got := combinedErr(estimate{Err: 3}, estimate{Err: 4}); math.Abs(got-5) > 1e-9 {
		t.Fatalf("got %v", got)
	}
}

func TestRemoveEntryDropsOnlyTheFirstIdenticalEntry(t *testing.T) {
	a := buildCastEntry("a", 1, 0, nil)
	b := buildCastEntry("b", 2, 0, nil)
	got := removeEntry([]entry{a, b, a}, a)
	if len(got) != 2 || got[0].Notes != "b" || got[1].Notes != "a" {
		t.Fatalf("got %+v", got)
	}
	if !sameEntry(a, buildCastEntry("a", 1, 0, nil)) || sameEntry(a, b) {
		t.Fatal("sameEntry misjudged identity")
	}
	if got := removeEntry([]entry{a}, b); len(got) != 1 {
		t.Fatal("removing an absent entry changed the list")
	}
}

func TestHuntBestAcceptsGainsBeyondErrorAndStopsOtherwise(t *testing.T) {
	// Fewer priority lines is "better": 10 DPS per removed line.
	score := func(r rotation, _ int) (estimate, error) {
		return estimate{Mean: 200 - 10*float64(len(r.PriorityList)), Err: 1}, nil
	}
	base := testRotation()
	baseEst, _ := score(base, 1)

	best, est, accepted := huntBest(base, nil, nil, score, 10, 2, baseEst, nil)
	if len(accepted) != 2 || len(best.PriorityList) != len(base.PriorityList)-2 {
		t.Fatalf("accepted %d, list %d", len(accepted), len(best.PriorityList))
	}
	if accepted[0].Round != 1 || accepted[1].Round != 2 || accepted[0].Gain <= accepted[0].GainErr {
		t.Fatalf("accepted = %+v", accepted)
	}
	if est.Mean != baseEst.Mean+20 {
		t.Fatalf("final estimate %v", est)
	}

	flat := func(rotation, int) (estimate, error) { return estimate{Mean: 100, Err: 1}, nil }
	_, _, none := huntBest(base, nil, nil, flat, 10, 5, estimate{Mean: 100, Err: 1}, nil)
	if len(none) != 0 {
		t.Fatalf("a flat landscape accepted %+v", none)
	}

	failing := func(rotation, int) (estimate, error) { return estimate{}, errSim }
	got, _, none := huntBest(base, nil, nil, failing, 10, 5, estimate{Mean: 100, Err: 1}, nil)
	if len(none) != 0 || len(got.PriorityList) != len(base.PriorityList) {
		t.Fatal("unsimmable mutations must never be adopted")
	}
}

type simError string

func (e simError) Error() string { return string(e) }

const errSim = simError("boom")

func TestRunProbeLabelsSortsAndClassifiesRows(t *testing.T) {
	base := testRotation()
	baseline := estimate{Mean: 100, Err: 1}
	run := func(r rotation, _ int) (estimate, error) {
		switch {
		case len(r.PriorityList) == len(base.PriorityList)+1: // an insertion
			return estimate{Mean: 110, Err: 1}, nil
		case len(r.PriorityList) == len(base.PriorityList)-1 && r.PriorityList[len(r.PriorityList)-1].Action["castSpell"] == nil:
			return estimate{}, errSim
		}
		return estimate{Mean: 90, Err: 1}, nil
	}
	rows := runProbe(base, testCandidates()[1:], baseline, run, 10, nil)
	if len(rows) != len(base.PriorityList)+1 {
		t.Fatalf("rows = %d", len(rows))
	}
	last := rows[len(rows)-1]
	if last.Kind != "insert" || last.Label != "PlainSpell" || last.Verdict != verdictHelps || last.Delta != 10 {
		t.Fatalf("insert row = %+v", last)
	}
	for _, r := range rows[:len(rows)-1] {
		if r.Kind != "remove" {
			t.Fatalf("removals must sort first: %+v", r)
		}
		if r.Verdict == verdictHelps && r.Delta != 10 {
			t.Fatalf("removal delta must be baseline minus removed: %+v", r)
		}
	}
}

func TestVerdictWordsFollowTheBestEstimate(t *testing.T) {
	r := report{baseline: estimate{Mean: 100, Err: 1}, bestEst: estimate{Mean: 110, Err: 1}, opts: options{confirmIterations: 800}}
	var b strings.Builder
	r.writeVerdict(&b)
	if !strings.Contains(b.String(), "found a better rotation, +10.0 DPS (+10.0%)") {
		t.Fatalf("got %s", b.String())
	}
	r.bestEst = estimate{Mean: 100.5, Err: 1}
	b.Reset()
	r.writeVerdict(&b)
	if !strings.Contains(b.String(), "already the best found") {
		t.Fatalf("got %s", b.String())
	}
}

func TestWriteAcceptedListsRoundsOrSaysNone(t *testing.T) {
	var b strings.Builder
	report{}.writeAccepted(&b)
	if !strings.Contains(b.String(), "No mutation beat the incumbent") {
		t.Fatal(b.String())
	}
	b.Reset()
	report{accepted: []acceptedMutation{{Round: 2, Label: "swap #1", Gain: 4.25, GainErr: 1.5}}}.writeAccepted(&b)
	if !strings.Contains(b.String(), "| 2 | swap #1 | +4.2 | 1.5 |") && !strings.Contains(b.String(), "| 2 | swap #1 | +4.3 | 1.5 |") {
		t.Fatal(b.String())
	}
}

func TestWriteRotationsPadsTheShorterSide(t *testing.T) {
	base := testRotation()
	r := report{in: inputs{base: base}, best: base.withPriorityList(base.PriorityList[:2])}
	var b strings.Builder
	r.writeRotations(&b)
	out := b.String()
	if !strings.Contains(out, "| 3 | spell 150 (rank 2) -- condition: spell 100 (rank 1) dot is active | - |") {
		t.Fatalf("got:\n%s", out)
	}
	if strings.Count(out, "\n| ") != len(base.PriorityList)+1 { // header row + one per line
		t.Fatalf("row count wrong:\n%s", out)
	}
}

func TestDanglingGatesFlagsGatesNothingApplies(t *testing.T) {
	rot := testRotation()
	if got := danglingGates(rot, nil); len(got) != 0 {
		t.Fatalf("a rotation that casts its gate's dot is not dangling: %v", got)
	}
	rot = rot.withPriorityList(append(rot.PriorityList[:1:1], rot.PriorityList[2:]...))
	got := danglingGates(rot, map[int]string{150: "Bonus", 100: "Dot"})
	if len(got) != 1 || !strings.Contains(got[0], "Bonus (rank 2) requires Dot (rank 1)") {
		t.Fatalf("got %v", got)
	}
	var b strings.Builder
	report{best: rot, names: map[int]string{}}.writeDanglingGates(&b)
	if !strings.Contains(b.String(), "## Dangling gates in the winning rotation") {
		t.Fatal(b.String())
	}
	b.Reset()
	report{best: testRotation()}.writeDanglingGates(&b)
	if b.Len() != 0 {
		t.Fatal("clean rotation produced a section")
	}
}

func TestPassiveProcAurasAreNeverDangling(t *testing.T) {
	if !isPassiveProcAura(fingersOfFrostAuraID) || isPassiveProcAura(1) {
		t.Fatal("isPassiveProcAura misjudged")
	}
}

func TestWinnerJSONAndBuildDescription(t *testing.T) {
	r := report{best: testRotation(), in: inputs{setup: simSetup{spec: specInfo{Spec: "mage-fire", Name: "Fire"}}}}
	var b strings.Builder
	r.writeWinnerJSON(&b)
	out := b.String()
	if !strings.Contains(out, "data/curated/apl/mage-fire.json") || !strings.Contains(out, "```json\n{") || !strings.HasSuffix(out, "```\n") {
		t.Fatalf("got %s", out)
	}
	if r.buildDescription() != "the guide's own FS1 build" {
		t.Fatal(r.buildDescription())
	}
	r.opts.buildCode = "FS1:x"
	if r.buildDescription() != "the build FS1:x" {
		t.Fatal(r.buildDescription())
	}
}

func TestMarkdownProbeOnlyStopsAfterTheProbe(t *testing.T) {
	r := report{
		opts:      options{level: 60, faction: "alliance", preset: "bare", seed: 7, confirmIterations: 800},
		in:        inputs{clientBuild: "1.2.3", setup: simSetup{spec: specInfo{Spec: "mage-fire", Name: "Fire"}}},
		baseline:  estimate{Mean: 100, Err: 1},
		probeOnly: true,
		elapsed:   1500 * time.Millisecond,
	}
	out := r.markdown()
	if !strings.Contains(out, "# Rotation search: Fire, level 60") || !strings.Contains(out, "## Action probe") || strings.Contains(out, "**Verdict:") {
		t.Fatalf("got %s", out)
	}
	r.probeOnly = false
	r.best = testRotation()
	r.bestEst = estimate{Mean: 100, Err: 1}
	r.in.base = testRotation()
	r.finishers = &finisherCasts{abilities: []string{"Rupture"}, baseline: castTable{"Rupture": 1}, best: castTable{"Rupture": 1}}
	full := r.markdown()
	for _, want := range []string{"**Verdict:", "## Search: accepted mutations", "## Finisher casts", "## Rotation, baseline vs winner", "## Winning rotation, engine APL JSON"} {
		if !strings.Contains(full, want) {
			t.Errorf("full report missing %q", want)
		}
	}
}
