package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// The bear's raid preset at 60 carries Mighty Rage Potion (item 13442; the
// deep dive makes rage potions usable by druids), with Greater Stoneshield
// behind it below 46, and Forever lets a druid drink a potion in Bear Form
// without leaving it (the druid deep dive, news.blizzard.com article
// 24301515). The rotation
// must reach it: every other tank's list opens with the line that uses its
// potions, trinkets and racials, and the bear's had none, so it never drank.
func TestTheBearDrinksItsPotion(t *testing.T) {
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
	for _, a := range newReport(req, fightSecondsOf(req), est, player, nil).Actions {
		if strings.Contains(a.Name, "13442") {
			casts += a.CastsPerIter
		}
	}
	if casts < 1 {
		t.Errorf("the bear drank %.2f Mighty Rage Potions a fight, want at least one", casts)
	}
}
