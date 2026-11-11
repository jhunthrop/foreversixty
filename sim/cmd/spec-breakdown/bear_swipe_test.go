package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// On a single boss Swipe's rage does more threat in Maul, Lacerate and
// Primal Bite: the bear card's "Swipe above 40 rage" line cost about 0.8% of
// raid 60 threat (1.6% before Swipe gained its attack power share) and 1% at
// bare 30 to 50. Swipe is the pack button, so a one-target fight casts none.
func TestBearSwipesOnlyOnAPack(t *testing.T) {
	o := options{
		repoRoot: "../../..", spec: "druid-feral-bear", preset: "bare", band: 60,
		faction: "alliance", iterations: 50, seed: 1,
	}
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	est, player, err := inproc.PlainRunWithRotation(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	r := newReport(req, fightSecondsOf(req), est, player, nil)
	for _, a := range r.Actions {
		if strings.HasSuffix(a.Name, "(9908)") && a.CastsPerIter > 0 {
			t.Errorf("one target: Swipe cast %.1f times a fight, want 0", a.CastsPerIter)
		}
	}
}

// From 40, when Primal Bite arrives, a bear is short of rage, and a Maul queued at 15 rage
// leaves Primal Bite waiting on its cooldown (28.4 casts of a possible 30 a fight at 43, against
// 29.5 with round 16's 30-rage gate). Round 30 retuned the gate on round 29's order and the
// dungeon-table boss: at 45 rage from 40 to 59 (sim/request TestBearMaulGateFollowsTheLevel)
// Lacerate takes more of the rage, Primal Bite still fires about 28 times, and threat rises
// 1.6 to 2.1% at 40 to 45.
func TestBearMaulKeepsRageForPrimalBiteFromFortyToFiftyThree(t *testing.T) {
	o := options{
		repoRoot: "../../..", spec: "druid-feral-bear", preset: "bare", band: 40, level: 43,
		faction: "alliance", iterations: 300, seed: 1,
	}
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	est, player, err := inproc.PlainRunWithRotation(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	r := newReport(req, fightSecondsOf(req), est, player, nil)
	bites := 0.0
	for _, a := range r.Actions {
		if strings.HasSuffix(a.Name, "(407995)") {
			bites += a.CastsPerIter
		}
	}
	if bites < 27.5 {
		t.Errorf("level 43: Primal Bite cast %.2f times a fight, want at least 27.5 (Maul is spending its rage)", bites)
	}
}
