package request

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// maelstromWeaponTalentSpellID is Maelstrom Weapon's talent spell, the
// stacking buff's id.
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
// Sim-accuracy 2026-11-12, shamans: on the 10 November Stormstrike (100% on
// 8 s, less mana spent) the quarter costs shocks: shocking down to 15%
// gains 3.1% at bare 60 and 1.3% at bare 40 over it and never runs dry,
// where shocking to empty gains 2.2% and 0.8% more but spends 9 s of a
// bare-60 fight and 55 s of a bare-40 one with no mana for Stormstrike or
// a totem (design/reviews/2026-11-12-sim-accuracy/shamans.md).
func TestEnhancementShocksDownToFifteenPercentOfItsMana(t *testing.T) {
	const earthShock = 10414
	for _, item := range enhancementRotation(t).PriorityList {
		action := item.GetAction()
		if action.GetCastSpell().GetSpellId().GetSpellId() != earthShock {
			continue
		}
		assertManaAbove(t, action.GetCondition(), "15%")
		return
	}
	t.Fatal("the Enhancement rotation casts no Earth Shock rank 7")
}

// Blizzard's 10 November hotfix: every Maelstrom Weapon stack takes 20% at
// every rank, so five stacks make the bolt instant and free from the first
// point, and the bolt no longer waits for 5/5 (+4.8 to +9.1% at 35 to 38,
// design/reviews/2026-11-11-sim-accuracy/hotfixes-other-classes.md).
func TestEnhancementBoltsAtFiveMaelstromStacksAtAnyRank(t *testing.T) {
	const lightningBolt = 15208
	for _, item := range enhancementRotation(t).PriorityList {
		action := item.GetAction()
		if action.GetCastSpell().GetSpellId().GetSpellId() != lightningBolt {
			continue
		}
		known, stacks := false, false
		for _, v := range action.GetCondition().GetAnd().GetVals() {
			if id := v.GetAuraIsKnown().GetAuraId(); id.GetSpellId() == maelstromWeaponTalentSpellID {
				if id.GetTag() != 0 {
					t.Fatalf("Lightning Bolt still waits for Maelstrom Weapon rank %d: %v", id.GetTag(), action.GetCondition())
				}
				known = true
			}
			cmp := v.GetCmp()
			if cmp.GetLhs().GetAuraNumStacks().GetAuraId().GetSpellId() == maelstromWeaponTalentSpellID && cmp.GetRhs().GetConst().GetVal() == "5" {
				stacks = true
			}
		}
		if !known || !stacks {
			t.Fatalf("Lightning Bolt's condition is not the talent known and five stacks: %v", action.GetCondition())
		}
		return
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
