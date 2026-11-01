package request

import (
	"slices"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// The 2026-10-31 caster theorycraft audit
// (design/reviews/2026-10-31-theorycraft-audit/casters.md) found rotation
// lines that disagreed with the build or with Forever theorycraft. These
// pin the corrected lines so a later rotation search cannot quietly undo
// them without a decision.

// castOrder lists the spell ids of a curated rotation's top-level castSpell
// lines, in priority order.
func castOrder(t *testing.T, spec string) []int32 {
	t.Helper()
	curated, err := loadLadderCurated(repoRoot, spec)
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatalf("%s: %v", spec, err)
	}
	var ids []int32
	for _, item := range rot.PriorityList {
		if cast := item.GetAction().GetCastSpell(); cast != nil {
			ids = append(ids, cast.GetSpellId().GetSpellId())
		}
		if channel := item.GetAction().GetChannelSpell(); channel != nil {
			ids = append(ids, channel.GetSpellId().GetSpellId())
		}
	}
	return ids
}

func requireBefore(t *testing.T, spec string, ids []int32, first, second int32) {
	t.Helper()
	i, j := slices.Index(ids, first), slices.Index(ids, second)
	if i < 0 || j < 0 {
		t.Fatalf("%s: want both %d and %d cast, got %v", spec, first, second, ids)
	}
	if i > j {
		t.Errorf("%s: %d must sit above %d, got %v", spec, first, second, ids)
	}
}

// W1, W2: the Affliction build takes Siphon Life, so the rotation keeps it
// up; Bane of Agony, not Bane of Doom, is the maintained curse.
func TestAfflictionKeepsSiphonLifeAndAgony(t *testing.T) {
	const (
		siphonLife = 18881
		agony      = 11713
		doom       = 603
	)
	ids := castOrder(t, "warlock-affliction")
	if !slices.Contains(ids, siphonLife) {
		t.Errorf("Affliction never casts Siphon Life (%d): %v", siphonLife, ids)
	}
	if !slices.Contains(ids, agony) || slices.Contains(ids, doom) {
		t.Errorf("Affliction must keep Bane of Agony (%d), not Bane of Doom (%d): %v", agony, doom, ids)
	}
}

// D2: Demonic Pact is real and modeled. The Demonology raid build takes it,
// and the prepull summons the Imp, sacrifices it, then summons the Succubus
// before Soul Link; Shadow Bolt is the filler under both Shadow bonuses.
func TestDemonologySacrificesTheImpUnderDemonicPact(t *testing.T) {
	const (
		summonImp         = 688
		demonicSacrifice  = 18788
		summonSuccubus    = 712
		soulLink          = 19028
		shadowBolt        = 11661
		searingPain       = 17923
		demonologySpecKey = "warlock-demonology"
	)
	curated, err := loadLadderCurated(repoRoot, demonologySpecKey)
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	var prepull []int32
	for _, p := range rot.PrepullActions {
		prepull = append(prepull, p.GetAction().GetCastSpell().GetSpellId().GetSpellId())
	}
	want := []int32{summonImp, demonicSacrifice, summonSuccubus, soulLink}
	if !slices.Equal(prepull, want) {
		t.Errorf("Demonology prepull = %v, want %v", prepull, want)
	}
	ids := castOrder(t, demonologySpecKey)
	if !slices.Contains(ids, shadowBolt) || slices.Contains(ids, searingPain) {
		t.Errorf("Demonology's filler must be Shadow Bolt (%d), not Searing Pain (%d): %v", shadowBolt, searingPain, ids)
	}
}

// P1, P2: Mind Blast sits above Mind Flay, and Devouring Plague is kept up
// by its own line rather than only behind Inner Focus.
func TestShadowCastsMindBlastBeforeMindFlayAndKeepsDevouringPlague(t *testing.T) {
	const (
		mindBlast       = 10947
		mindFlay        = 18807
		devouringPlague = 19280
		shadowSpecKey   = "priest-shadow"
	)
	ids := castOrder(t, shadowSpecKey)
	requireBefore(t, shadowSpecKey, ids, mindBlast, mindFlay)
	if !slices.Contains(ids, devouringPlague) {
		t.Errorf("Shadow has no standalone Devouring Plague line: %v", ids)
	}
}

// R1: Frost casts Arcane Missiles on Missile Barrage procs; A4: Arcane
// stacks Arcane Blast to 3 before its Frostbolt filler, and since client
// 1.60.1.70707 (+100% cost a stack, 18 s) keeps blasting past 3 while the
// mana bar is ahead of the fight.
func TestMageMissileBarrageAndArcaneBlastStacks(t *testing.T) {
	const arcaneMissiles = 25345
	if ids := castOrder(t, "mage-frost"); !slices.Contains(ids, arcaneMissiles) {
		t.Errorf("Frost never casts Arcane Missiles on Missile Barrage: %v", ids)
	}
	curated, err := loadLadderCurated(repoRoot, "mage-arcane")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	for _, item := range rot.PriorityList {
		either := item.GetAction().GetCondition().GetOr().GetVals()
		if len(either) != 2 || either[0].GetCmp().GetLhs().GetAuraNumStacks() == nil {
			continue
		}
		if got := either[0].GetCmp().GetRhs().GetConst().GetVal(); got != "3" {
			t.Errorf("Arcane Blast stacks to %s before Frostbolt, want 3", got)
		}
		if !conditionPacesMana(either[1]) {
			t.Error("Arcane Blast's past-3 branch does not compare mana against the fight's remaining time")
		}
		return
	}
	t.Error("the Arcane rotation has no Arcane Blast stack condition")
}

// conditionPacesMana reports whether v (or one of the values it ands)
// compares currentManaPercent against remainingTimePercent.
func conditionPacesMana(v *proto.APLValue) bool {
	if cmp := v.GetCmp(); cmp != nil {
		return cmp.GetLhs().GetCurrentManaPercent() != nil && cmp.GetRhs().GetRemainingTimePercent() != nil
	}
	for _, inner := range v.GetAnd().GetVals() {
		if conditionPacesMana(inner) {
			return true
		}
	}
	return false
}
