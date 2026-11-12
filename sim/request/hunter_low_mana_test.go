package request

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
)

// At 18 and 19 a Beast Mastery or Marksmanship hunter has Multi-Shot but no
// Aimed Shot. Below a quarter of his mana Multi-Shot is held, so its
// cooldown never runs, and Serpent Sting and Arcane Shot, which wait on
// Multi-Shot's cooldown, never fire either: the card cast nothing but
// Hunter's Mark until the mana was back (guide-accuracy round 20, sim
// round 21 G1). Under a quarter mana they now fire on their own.
func TestLowManaHunterKeepsShootingAt18(t *testing.T) {
	registerEngine.Do(engine.RegisterAll)
	const level = 18
	for _, spec := range []string{"hunter-beast-mastery", "hunter-marksmanship"} {
		t.Run(spec, func(t *testing.T) {
			req := smokeRequest(t, spec)
			req.Character.Level = level
			req.Encounter.TargetLevel = 0
			req.Encounter.DurationSec = 300
			engineReq, err := BuildWith(req, Options{OpenIterations: true})
			if err != nil {
				t.Fatal(err)
			}
			res := core.RunRaidSim(engineReq)
			if err := adapter.ResultError(res); err != nil {
				t.Fatal(err)
			}
			player, err := adapter.PlayerMetrics(res)
			if err != nil {
				t.Fatal(err)
			}
			var arcane, multi float64
			for _, c := range castSet(player) {
				switch c.name {
				case "spell:3044", "spell:14281", "spell:14282":
					arcane += c.casts
				case "spell:2643":
					multi += c.casts
				}
			}
			t.Logf("level %d: Arcane Shot %.1f, Multi-Shot %.1f casts in 300 s", level, arcane, multi)
			// About one Arcane Shot every 15 s at least: its cooldown is 6 s and
			// a hunter out of Multi-Shot mana still has mana for it. (The stall
			// this guards cast none. Since 2026-11-12 a Beast Mastery hunter at
			// 18 keeps the newest Serpent Sting, which costs more than rank 1.)
			if arcane < 20 {
				t.Errorf("Arcane Shot cast %.1f times in 300 s at level %d, want 20 or more", arcane, level)
			}
		})
	}
}
