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
	r := newReport(req, est, player, nil)
	for _, a := range r.Actions {
		if strings.HasSuffix(a.Name, "(9908)") && a.CastsPerIter > 0 {
			t.Errorf("one target: Swipe cast %.1f times a fight, want 0", a.CastsPerIter)
		}
	}
}

// From 40, when Primal Bite arrives, to 53 a bear is short of rage, and a
// Maul queued at 15 rage leaves Primal Bite waiting on its cooldown: Maul at
// 30 rage there is +0.2 to +0.8% threat (bare, both seeds, both factions),
// while from 54 and at raid 60 the 15 holds. At 43 Primal Bite then fires on
// nearly every cooldown (29.5 of a possible 30 a fight, against 28.4).
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
	r := newReport(req, est, player, nil)
	bites := 0.0
	for _, a := range r.Actions {
		if strings.HasSuffix(a.Name, "(407995)") {
			bites += a.CastsPerIter
		}
	}
	if bites < 29.2 {
		t.Errorf("level 43: Primal Bite cast %.2f times a fight, want at least 29.2 (Maul is spending its rage)", bites)
	}
}
