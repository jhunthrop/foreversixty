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

// Round 16 mage: the Arcane raid build has no Improved Frostbolt, so its Frostbolt is a 3 s cast
// and a pre-pull Frostbolt at -2.5 s lands half a second into the fight. With Heating Up known
// (the raid build's Fire tree) the opener is a Pyroblast 6 s out, its whole cast, as Fire opens;
// without it (the leveling build, Improved Frostbolt 5/5) it stays a 2.5 s Frostbolt. Each pre-pull
// line carries the Heating Up gate. +0.4 to +0.6% at raid 60, bare bands unchanged.
func TestArcaneOpensWithPyroblastOnTheRaidBuildAndFrostboltLeveling(t *testing.T) {
	const (
		heatingUp = 400625
		pyroblast = 18809
		frostbolt = 10181
	)
	curated, err := loadLadderCurated(repoRoot, "mage-arcane")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	want := map[int32]struct {
		at    string
		known bool
	}{pyroblast: {"-6s", true}, frostbolt: {"-2.5s", false}}
	if len(rot.PrepullActions) != len(want) {
		t.Fatalf("Arcane prepull has %d lines, want %d (Pyroblast with Heating Up, Frostbolt without)", len(rot.PrepullActions), len(want))
	}
	for _, p := range rot.PrepullActions {
		id := p.GetAction().GetCastSpell().GetSpellId().GetSpellId()
		w, ok := want[id]
		if !ok {
			t.Errorf("unexpected pre-pull cast %d", id)
			continue
		}
		if got := p.GetDoAtValue().GetConst().GetVal(); got != w.at {
			t.Errorf("pre-pull %d at %q, want %q", id, got, w.at)
		}
		cond := p.GetAction().GetCondition()
		known := cond.GetAuraIsKnown()
		if !w.known {
			known = cond.GetNot().GetVal().GetAuraIsKnown()
		}
		if known.GetAuraId().GetSpellId() != heatingUp {
			t.Errorf("pre-pull %d: want gate %s Heating Up (%d) known, got %v", id, map[bool]string{true: "", false: "not"}[w.known], heatingUp, cond)
		}
	}
}

// Round 16 wave 2 (shaman lane's phantom lines): a line that reads Presence of Mind's aura must
// also check the talent is known. The engine drops an auraIsActive on an aura it never
// registered, so without the talent the rest of the line's condition stands alone: the Pyroblast
// line would fire on every global for any build with Heating Up and no Presence of Mind.
func TestArcanePresenceOfMindLinesCheckTheTalent(t *testing.T) {
	const presenceOfMind = 12043
	curated, err := loadLadderCurated(repoRoot, "mage-arcane")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	for i, item := range rot.PriorityList {
		vals := item.GetAction().GetCondition().GetAnd().GetVals()
		readsAura, checksTalent := false, false
		for _, v := range vals {
			if v.GetAuraIsActive().GetAuraId().GetSpellId() == presenceOfMind {
				readsAura = true
			}
			if v.GetSpellIsKnown().GetSpellId().GetSpellId() == presenceOfMind || v.GetAuraIsKnown().GetAuraId().GetSpellId() == presenceOfMind {
				checksTalent = true
			}
		}
		if readsAura && !checksTalent {
			t.Errorf("line %d reads Presence of Mind's aura without spellIsKnown(%d)", i, presenceOfMind)
		}
	}
}
