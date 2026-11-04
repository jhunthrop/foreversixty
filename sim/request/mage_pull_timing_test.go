package request

import (
	"slices"
	"testing"
	"time"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// Round 15 mage (guide-accuracy round 1 SIM-2): Arcane Power and Presence
// of Mind go out after the tank's first Sunders, not on the pull. The
// rotation's own note says "On a fresh pull, give the tank a Sunder or two
// first"; the guide's opener blasts twice first. Measured at a tie.
func TestArcaneHoldsArcanePowerForTheTanksFirstSunders(t *testing.T) {
	const (
		arcanePower    = 12042
		presenceOfMind = 12043
		minDelay       = 3 * time.Second
	)
	curated, err := loadLadderCurated(repoRoot, "mage-arcane")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	seen := map[int32]bool{}
	for _, item := range rot.PriorityList {
		id := item.GetAction().GetCastSpell().GetSpellId().GetSpellId()
		if id != arcanePower && id != presenceOfMind {
			continue
		}
		seen[id] = true
		cmp := item.GetAction().GetCondition().GetCmp()
		if cmp == nil || cmp.GetOp() != proto.APLValueCompare_OpGe || cmp.GetLhs().GetCurrentTime() == nil {
			t.Errorf("spell %d: want a currentTime >= condition, got %v", id, item.GetAction().GetCondition())
			continue
		}
		d, err := time.ParseDuration(cmp.GetRhs().GetConst().GetVal())
		if err != nil || d < minDelay {
			t.Errorf("spell %d waits %q, want at least %v", id, cmp.GetRhs().GetConst().GetVal(), minDelay)
		}
	}
	if !seen[arcanePower] || !seen[presenceOfMind] {
		t.Errorf("Arcane must cast Arcane Power and Presence of Mind, saw %v", seen)
	}
}

// Round 15 wave 4 (guide-accuracy round 2 SIM-R2-1, SIM-R2-2): the Arcane raid build plays a
// Fire second tree (Heating Up, Pyroblast, Improved Fire Blast) and Arcane Wrath. The rotation
// casts Pyroblast on three Heating Up stacks and with Presence of Mind, Fire Blast on cooldown,
// Arcane Wrath at three Blast stacks and Fireball as the filler, each gated on the talent, so
// the leveling build keeps its Frostbolt lines. +6.4 to +7.1% at raid 60.
func TestArcaneRaidPlaysItsFireSecondTree(t *testing.T) {
	const (
		pyroblast   = 18809
		fireBlast   = 10199
		arcaneWrath = 1326082
		fireball    = 10151
		frostbolt   = 10181
	)
	ids := castOrder(t, "mage-arcane")
	for _, id := range []int32{pyroblast, fireBlast, arcaneWrath, fireball, frostbolt} {
		if !slices.Contains(ids, id) {
			t.Errorf("Arcane never casts %d: %v", id, ids)
		}
	}
	lastFrostbolt := -1
	for i, id := range ids {
		if id == frostbolt {
			lastFrostbolt = i
		}
	}
	if slices.Index(ids, fireball) > lastFrostbolt {
		t.Errorf("the Fireball filler must sit above the Frostbolt filler: %v", ids)
	}
}
