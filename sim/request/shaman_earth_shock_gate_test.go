package request

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// maelstromWeaponTalentSpellID is Maelstrom Weapon's talent spell. The
// engine's trigger aura carries the talent rank as its ActionID tag, and the
// stacking buff keeps the bare id.
const maelstromWeaponTalentSpellID = 408498

// Sim-accuracy round 18, shaman lane: Enhancement held Earth Shock back below
// half its mana. With raid buffs it never runs dry, so the reserve only threw
// shocks away, and without them a quarter is the reserve that keeps the most
// damage (design/reviews/2026-11-04-sim-audit-round18/shaman.md).
// Round 21: Maelstrom Weapon's chance doubled (Blizzard, 3 November), and a
// five-stack Lightning Bolt below 5/5 now loses to the swings it delays
// (+1.9% to +3.7% at 38 to 46 without it), so the bolt waits for all five
// points and the half reserve that paid for those bolts is gone
// (design/reviews/2026-11-05-sim-audit-round21/shaman.md).
func TestEnhancementShocksDownToAQuarterOfItsMana(t *testing.T) {
	const earthShock = 10414
	for _, item := range enhancementRotation(t).PriorityList {
		action := item.GetAction()
		if action.GetCastSpell().GetSpellId().GetSpellId() != earthShock {
			continue
		}
		assertManaAbove(t, action.GetCondition(), "25%")
		return
	}
	t.Fatal("the Enhancement rotation casts no Earth Shock rank 7")
}

func TestEnhancementBoltsMaelstromOnlyAtFiveOfFive(t *testing.T) {
	const lightningBolt = 15208
	for _, item := range enhancementRotation(t).PriorityList {
		action := item.GetAction()
		if action.GetCastSpell().GetSpellId().GetSpellId() != lightningBolt {
			continue
		}
		for _, v := range action.GetCondition().GetAnd().GetVals() {
			id := v.GetAuraIsKnown().GetAuraId()
			if id.GetSpellId() == maelstromWeaponTalentSpellID && id.GetTag() == 5 {
				return
			}
		}
		t.Fatalf("Lightning Bolt's condition does not ask for Maelstrom Weapon 5/5: %v", action.GetCondition())
	}
	t.Fatal("the Enhancement rotation casts no Lightning Bolt rank 10")
}

func enhancementRotation(t *testing.T) *proto.APLRotation {
	t.Helper()
	curated, err := loadLadderCurated(repoRoot, "shaman-enhancement")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	return rot
}

func assertManaAbove(t *testing.T, value *proto.APLValue, want string) {
	t.Helper()
	cmp := value.GetCmp()
	if cmp.GetLhs().GetCurrentManaPercent() == nil || cmp.GetOp() != proto.APLValueCompare_OpGt {
		t.Fatalf("condition is not 'mana percent above': %v", value)
	}
	if got := cmp.GetRhs().GetConst().GetVal(); got != want {
		t.Errorf("Earth Shock holds back below %s mana, want %s", got, want)
	}
}
