package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// demonicSacrificeBuffs are the four buffs Demonic Sacrifice (18788) grants,
// one per sacrificed demon: Burning Shadow (Imp), Fel Stamina (Voidwalker),
// Touch of Fire (Succubus) and Fel Energy (Felhunter).
var demonicSacrificeBuffs = []int32{18789, 18790, 18791, 18792}

// minSacrificeUptimePct is how much of the fight a sacrifice buff must cover
// once it is applied. Nothing in the fight removes it except a summon, so a
// sacrifice the rotation keeps is up from the pull to the end.
const minSacrificeUptimePct = 95.0

// Demonic Sacrifice is cancelled by summoning any demon unless the build takes
// Demonic Pact, which keeps it across a DIFFERENT demon. A rotation that
// sacrifices a demon and then summons another without Pact throws the
// sacrifice away before the pull (owner ruling, sim-accuracy round 9). At
// every published Demonology band, a sacrifice buff that was applied must
// still be up for the fight.
func TestDemonologyNeverSummonsAwayItsSacrifice(t *testing.T) {
	cases := []struct {
		band   int
		preset string
	}{
		{20, "bare"}, {30, "bare"}, {40, "bare"}, {50, "bare"}, {60, "bare"}, {60, "raid"},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", tc.band, tc.preset), func(t *testing.T) {
			o := options{
				repoRoot: "../../..", spec: "warlock-demonology", preset: tc.preset, band: tc.band,
				faction: "horde", iterations: 50, seed: 1,
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
			for _, a := range r.Auras {
				for _, id := range demonicSacrificeBuffs {
					if !strings.HasSuffix(a.Name, fmt.Sprintf("(%d)", id)) || a.ProcsPerIter == 0 {
						continue
					}
					if a.UptimePct < minSacrificeUptimePct {
						t.Errorf("%s applied %.2f times a fight but up %.1f%% of it: the rotation summoned a demon that cancelled the sacrifice", a.Name, a.ProcsPerIter, a.UptimePct)
					}
				}
			}
		})
	}
}
