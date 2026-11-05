package request

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
)

// Fire Nova bursts from the shaman's active Fire totem. Below 26 an
// Elemental shaman's only Fire totem is Searing Totem, and the card dropped
// it on a single target only, so on a pack Fire Nova (trained at 12) never
// had a totem to burst from (guide-accuracy round 19 sim finding, sim round
// 21 shaman S1). Until Magma Totem is trained the pack gets Searing Totem.
// The first 30 seconds of the pull: a naked smoke character runs out of
// mana later on, which is not what this pins.
func TestElementalDropsSearingTotemForFireNovaOnPacksBelow26(t *testing.T) {
	registerEngine.Do(engine.RegisterAll)
	for _, level := range []int{14, 20, 25} {
		req := smokeRequest(t, "shaman-elemental")
		req.Character.Level = level
		req.Encounter.TargetLevel = 0
		req.Encounter.Targets = 3
		req.Encounter.DurationSec = 30
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
		var searing, nova float64
		for _, c := range castSet(player) {
			switch c.name {
			case "spell:3599", "spell:6363", "spell:6364", "spell:6365", "spell:10437", "spell:10438":
				searing += c.casts
			case "spell:408341", "spell:408342", "spell:408343", "spell:408344", "spell:408345":
				nova += c.casts
			}
		}
		t.Logf("level %d, 3 targets: Searing Totem %.1f, Fire Nova %.1f casts in 30 s", level, searing, nova)
		if searing < 1 || nova < 1 {
			t.Errorf("level %d on 3 targets: Searing Totem %.1f, Fire Nova %.1f casts in 30 s; want the totem down and Fire Nova bursting from it", level, searing, nova)
		}
	}
}
