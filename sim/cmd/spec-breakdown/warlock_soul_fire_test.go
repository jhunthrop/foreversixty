package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// Sim-accuracy round 15 wave 4 (warlock lane, guide-accuracy round 2 S1):
// with Demonic Pact the raid build keeps the Imp sacrifice (+15% Shadow),
// Master Demonologist's Shadow 10% and Improved Shadow Bolt up, and Soul Fire
// (Fire) gets none of them, so even Decimation's free 2.4 s Soul Fire does less
// than the Shadow Bolt it replaces (raid 60 +1.3% without it). The leveling
// builds hold no Pact, and there the free Soul Fire still wins (bare 60 +1.5%).
func TestDemonologySoulFireOnlyWithoutDemonicPact(t *testing.T) {
	cases := []struct {
		name   string
		preset string
		want   bool
	}{
		{"raid 60, Pact build", "raid", false},
		{"bare 60, leveling build", "bare", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			o := options{
				repoRoot: "../../..", spec: "warlock-demonology", preset: c.preset, band: 60,
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
			r := newReport(req, fightSecondsOf(req), est, player, nil)
			casts := 0.0
			for _, a := range r.Actions {
				if strings.HasSuffix(a.Name, "(17924)") {
					casts = a.CastsPerIter
				}
			}
			if got := casts > 0.5; got != c.want {
				t.Errorf("Soul Fire cast %.2f times a fight; want casting it = %v", casts, c.want)
			}
		})
	}
}
