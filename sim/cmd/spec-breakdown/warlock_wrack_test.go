package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// wrackRowSuffix ends the breakdown's row for Wrack, named "Wrack (1316697)"
// with the spell names loaded and "spell (1316697)" without them.
const wrackRowSuffix = "(1316697)"

// Sim-accuracy round 15 (warlock lane): the 70755 hotfix push raised Wrack
// from 36 to 86 a tick at 60 and doubled its coefficient (0.143 to 0.286).
// With Soul Siphon's 36% it now out-damages Shadow Bolt per second of
// casting and costs a quarter of the mana, so Affliction's filler is Wrack.
// The rotation never cast it, so every band whose build takes Wrack carried a
// dead capstone, and the raid build left it out.
func TestAfflictionCastsWrackWhenItsBuildTakesIt(t *testing.T) {
	_, trees, err := leveling.GuideRaidBuildTalents("../../..", "warlock", "affliction")
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name    string
		preset  string
		band    int
		talents string
	}{
		{"bare 50", "bare", 50, ""},
		{"bare 60", "bare", 60, ""},
		{"raid 60 guide raid build", "raid", 60, strings.Join(trees[:], "-")},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			o := options{
				repoRoot: "../../..", spec: "warlock-affliction", preset: c.preset, band: c.band,
				faction: "horde", iterations: 100, seed: 1, talents: c.talents,
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
			casts := 0.0
			for _, a := range r.Actions {
				if strings.HasSuffix(a.Name, wrackRowSuffix) {
					casts = a.CastsPerIter
				}
			}
			if casts < 10 {
				t.Errorf("Wrack cast %.2f times a %.0f s fight, want it as the filler (10 or more)", casts, r.FightSeconds)
			}
		})
	}
}
