package request

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// Sim-accuracy round 18, shaman lane: Enhancement held Earth Shock back below
// half its mana. With raid buffs it never runs dry, so the reserve only threw
// shocks away, and without them a quarter is the reserve that keeps the most
// damage (design/reviews/2026-11-04-sim-audit-round18/shaman.md).
func TestEnhancementShocksDownToAQuarterOfItsMana(t *testing.T) {
	const earthShock = 10414
	curated, err := loadLadderCurated(repoRoot, "shaman-enhancement")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	for _, item := range rot.PriorityList {
		action := item.GetAction()
		if action.GetCastSpell().GetSpellId().GetSpellId() != earthShock {
			continue
		}
		cmp := action.GetCondition().GetCmp()
		if cmp.GetLhs().GetCurrentManaPercent() == nil || cmp.GetOp() != proto.APLValueCompare_OpGt {
			t.Fatalf("Earth Shock's condition is not 'mana percent above': %v", action.GetCondition())
		}
		if got := cmp.GetRhs().GetConst().GetVal(); got != "25%" {
			t.Errorf("Earth Shock holds back below %s mana, want 25%%", got)
		}
		return
	}
	t.Fatal("the Enhancement rotation casts no Earth Shock rank 7")
}
