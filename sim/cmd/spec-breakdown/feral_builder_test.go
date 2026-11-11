package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// The cat's builder switches from Claw to Shred at 21, the second point of
// Shredding Attacks (Shred 48 energy against Claw's 40 with Ferocity full):
// Shred is +1% at 21 and +7.5% at 22 and 23 behind the target, Claw +4% at 20
// (bare band 20, guide leveling order, seeds 1 and 50001, both factions). The
// card switched at Ferocious Bite's trainer level, 24.
func TestFeralBuilderIsShredFromTwentyOne(t *testing.T) {
	for _, c := range []struct {
		level     int
		talents   string
		wantShred bool
	}{
		{20, "0000000000000000-50212010000000000000-0000000000000000", false},
		{21, "0000000000000000-50212020000000000000-0000000000000000", true},
		{23, "0000000000000000-50212031000000000000-0000000000000000", true},
	} {
		o := options{
			repoRoot: "../../..", spec: "druid-feral", preset: "bare", band: 20, level: c.level,
			faction: "horde", iterations: 30, seed: 1, talents: c.talents,
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
		shred, claw := 0.0, 0.0
		for _, a := range r.Actions {
			switch {
			case strings.HasSuffix(a.Name, "(5221)"):
				shred += a.CastsPerIter
			case strings.HasSuffix(a.Name, "(1082)"):
				claw += a.CastsPerIter
			}
		}
		if c.wantShred && (shred == 0 || claw > 0) {
			t.Errorf("level %d: Shred %.1f, Claw %.1f casts; want Shred as the builder", c.level, shred, claw)
		}
		if !c.wantShred && (claw == 0 || shred > 0) {
			t.Errorf("level %d: Shred %.1f, Claw %.1f casts; want Claw as the builder", c.level, shred, claw)
		}
	}
}
