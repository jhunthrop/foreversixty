package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// majorManaPotionRow is the breakdown's row for Major Mana Potion (13444).
const majorManaPotionRow = "item 13444"

// A raid warlock drinks a Major Mana Potion on each of its two-minute
// cooldowns: the raid preset carries it ("used on cooldown by the engine
// when mana allows"), and every published warlock set Life Taps a dozen
// times or more a fight, so mana is short all fight. The curated rotations
// used to name the potion in a line that also required the execute phase,
// which held it to the last 35% of the fight: one potion in 180 s, not two,
// and Life Tap's globals paid for the second.
func TestWarlockRaidRotationsDrinkAManaPotionOnEachCooldown(t *testing.T) {
	for _, spec := range []string{"warlock-affliction", "warlock-demonology", "warlock-destruction"} {
		t.Run(spec, func(t *testing.T) {
			o := options{
				repoRoot: "../../..", spec: spec, preset: "raid", band: 60,
				faction: "horde", iterations: 200, seed: 1,
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
			casts := 0.0
			for _, a := range r.Actions {
				if a.Name == majorManaPotionRow {
					casts = a.CastsPerIter
				}
			}
			if casts < 1.9 {
				t.Errorf("Major Mana Potion drunk %.2f times a %.0f s fight, want about 2 (one per cooldown)", casts, r.FightSeconds)
			}
		})
	}
}
