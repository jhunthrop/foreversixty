package request

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// Owner ruling 4 (sim-accuracy round 14): the tank boss carries a Classic
// creature's attack power for its level, so Demoralizing Shout and
// Demoralizing Roar take a share of its swing. The boss is three levels
// above the character: 268 at 63 and 80 at 23 (cmangos classic-db
// creature_template_classlevelstats, warrior class). The mean swing with
// that attack power is still the curated one: the weapon roll under it is
// min_base_damage over (1 + attack power x the engine's coefficient).
func TestTankBossCarriesClassicAttackPower(t *testing.T) {
	for _, c := range []struct {
		level       int
		attackPower float64
	}{{60, 268}, {20, 80}} {
		req := tankRequest("warrior-protection", "warrior", c.level)
		enc := encounter(req.Encounter, c.level)
		if err := applyTankFight(&proto.Player{}, &proto.Raid{}, enc, req); err != nil {
			t.Fatal(err)
		}
		profile, err := TankProfileForLevel(c.level)
		if err != nil {
			t.Fatal(err)
		}
		boss := enc.Targets[0]
		if got := boss.Stats[proto.Stat_StatAttackPower]; got != c.attackPower {
			t.Errorf("level %d: boss attack power %v, want %v", c.level, got, c.attackPower)
		}
		full := boss.MinBaseDamage * (1 + c.attackPower*core.EnemyAutoAttackAPCoefficient)
		if math.Abs(full-profile.Boss.MinBaseDamage) > 1e-6 {
			t.Errorf("level %d: weapon roll %v with %v attack power swings for %v at least, want the curated %v",
				c.level, boss.MinBaseDamage, c.attackPower, full, profile.Boss.MinBaseDamage)
		}
	}
}
