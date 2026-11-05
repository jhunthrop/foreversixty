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
