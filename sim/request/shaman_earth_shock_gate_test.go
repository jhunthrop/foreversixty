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
// Round 19: while Maelstrom Weapon is below 5/5 a five-stack Lightning Bolt
// still costs mana, and half is the better reserve there (+0.3% to +1.1% at
// 40 to 49 on Maelstrom Weapon 3/5); a quarter holds without the talent and
// at 5/5 (design/reviews/2026-11-04-sim-audit-round19/shaman.md).
func TestEnhancementShocksDownToAQuarterOfItsManaOrHalfBelowFullMaelstrom(t *testing.T) {
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
		vals := action.GetCondition().GetAnd().GetVals()
		if len(vals) != 2 {
			t.Fatalf("Earth Shock's condition is not 'a quarter, and half below 5/5 Maelstrom': %v", action.GetCondition())
		}
		assertManaAbove(t, vals[0], "25%")
		alts := vals[1].GetOr().GetVals()
		if len(alts) != 3 {
			t.Fatalf("Earth Shock's half-mana clause has %d alternatives, want 3: %v", len(alts), vals[1])
		}
		assertManaAbove(t, alts[0], "50%")
		untalented := alts[1].GetNot().GetVal().GetAuraIsKnown().GetAuraId()
		if untalented.GetSpellId() != maelstromWeaponTalentSpellID || untalented.GetTag() != 0 {
			t.Errorf("the quarter reserve without the talent reads %v, want aura %d known", alts[1], maelstromWeaponTalentSpellID)
		}
		full := alts[2].GetAuraIsKnown().GetAuraId()
		if full.GetSpellId() != maelstromWeaponTalentSpellID || full.GetTag() != 5 {
			t.Errorf("the quarter reserve at 5/5 reads %v, want aura %d tag 5 known", alts[2], maelstromWeaponTalentSpellID)
		}
		return
	}
	t.Fatal("the Enhancement rotation casts no Earth Shock rank 7")
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
