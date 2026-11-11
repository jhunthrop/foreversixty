package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// Evocation (12051) is every mage's mana cooldown from level 20
// (SkillLineAbility; SpellLevels 20): an 8 s channel at +1500% regeneration
// on an 8 minute cooldown. Only the Fire rotation cast it, so a bare-band
// Arcane or Frost mage ran dry for a third to a half of the fight and
// wanded with Evocation ready.
func TestBareMagesEvocateWhenTheyRunDry(t *testing.T) {
	for _, spec := range []string{"mage-arcane", "mage-fire", "mage-frost"} {
		for _, band := range []int{20, 30, 40, 50} {
			o := options{
				repoRoot: "../../..", spec: spec, preset: "bare", band: band,
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
				if strings.Contains(a.Name, "(12051)") {
					casts += a.CastsPerIter
				}
			}
			if casts < 0.5 {
				t.Errorf("%s band %d: %.2f Evocations a fight, want at least one in two", spec, band, casts)
			}
		}
	}
}
