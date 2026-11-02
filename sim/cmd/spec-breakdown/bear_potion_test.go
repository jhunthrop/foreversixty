package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// The bear's raid preset carries Greater Stoneshield Potion (item 13455),
// and Forever lets a druid drink a potion in Bear Form without leaving it
// (the druid deep dive, news.blizzard.com article 24301515). The rotation
// must reach it: every other tank's list opens with the line that uses its
// potions, trinkets and racials, and the bear's had none, so it never drank.
func TestTheBearDrinksItsStoneshieldPotion(t *testing.T) {
	o := options{
		repoRoot: "../../..", spec: "druid-feral-bear", preset: "raid", band: 60,
		faction: "horde", iterations: 100, seed: 1,
	}
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	est, player, err := inproc.PlainRunWithRotation(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	casts := 0.0
	for _, a := range newReport(req, est, player, nil).Actions {
		if strings.Contains(a.Name, "13455") {
			casts += a.CastsPerIter
		}
	}
	if casts < 1 {
		t.Errorf("the bear drank %.2f Greater Stoneshield Potions a fight, want at least one", casts)
	}
}
