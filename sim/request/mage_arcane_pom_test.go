package request

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// Round 18 kept the Presence of Mind Frostbolt off the raid card (the raid
// build takes the gating talent, the leveling build never does). Sim-accuracy
// round 21 (design/reviews/2026-11-05-sim-audit-round21/mage.md A5) moved the
// gate from Heating Up to Pyroblast: the Arcane raid order takes Pyroblast at 54 and Heating Up at 57. A
// Presence of Mind Pyroblast beats a Presence of Mind Frostbolt at 54 to 56
// too, so the two Presence of Mind lines are gated on Pyroblast being known,
// not on Heating Up.
func TestArcanePresenceOfMindWaitsForPyroblastNotHeatingUp(t *testing.T) {
	const (
		pyroblast      = 18809
		frostbolt      = 10181
		presenceOfMind = 12043
		arcaneSpecSlug = "mage-arcane"
	)
	curated, err := loadLadderCurated(repoRoot, arcaneSpecSlug)
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	found := map[int32]bool{}
	for i, item := range rot.PriorityList {
		action := item.GetAction()
		id := action.GetCastSpell().GetSpellId().GetSpellId()
		if id != pyroblast && id != frostbolt {
			continue
		}
		vals := action.GetCondition().GetAnd().GetVals()
		if !anyValue(vals, func(v *proto.APLValue) bool { return v.GetAuraIsActive().GetAuraId().GetSpellId() == presenceOfMind }) {
			continue // not a Presence of Mind line
		}
		found[id] = true
		knowsPyro := func(v *proto.APLValue) bool { return v.GetSpellIsKnown().GetSpellId().GetSpellId() == pyroblast }
		switch id {
		case pyroblast:
			if !anyValue(vals, knowsPyro) {
				t.Errorf("line %d: the Presence of Mind Pyroblast must be gated on spellIsKnown(%d), got %v", i, pyroblast, vals)
			}
		case frostbolt:
			if !anyValue(vals, func(v *proto.APLValue) bool { return knowsPyro(v.GetNot().GetVal()) }) {
				t.Errorf("line %d: the Presence of Mind Frostbolt must be gated on not spellIsKnown(%d), got %v", i, pyroblast, vals)
			}
		}
	}
	if !found[pyroblast] || !found[frostbolt] {
		t.Fatalf("want both Presence of Mind lines, got %v", found)
	}
}

func anyValue(vals []*proto.APLValue, pred func(*proto.APLValue) bool) bool {
	for _, v := range vals {
		if pred(v) {
			return true
		}
	}
	return false
}

// Sim-accuracy round 23 (design/reviews/2026-11-05-sim-audit-round23/mage.md A1, guide-accuracy
// round 24 sims requested 3): before Pyroblast at 54 the raid order holds Arcane Wrath (31),
// a 3.5 second cast on a 6 second cooldown, longer than Frostbolt. Presence of Mind goes to it
// when it is ready (+0.4 to +0.9% at 32 to 53), and to Frostbolt only after it.
func TestArcanePresenceOfMindTakesArcaneWrathBeforeFrostbolt(t *testing.T) {
	const (
		pyroblast      = 18809
		frostbolt      = 10181
		arcaneWrath    = 1326082
		presenceOfMind = 12043
	)
	curated, err := loadLadderCurated(repoRoot, "mage-arcane")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	wrathLine, frostLine := -1, -1
	for i, item := range rot.PriorityList {
		action := item.GetAction()
		vals := action.GetCondition().GetAnd().GetVals()
		if !anyValue(vals, func(v *proto.APLValue) bool { return v.GetAuraIsActive().GetAuraId().GetSpellId() == presenceOfMind }) {
			continue
		}
		switch action.GetCastSpell().GetSpellId().GetSpellId() {
		case arcaneWrath:
			wrathLine = i
			if !anyValue(vals, func(v *proto.APLValue) bool { return v.GetSpellIsReady().GetSpellId().GetSpellId() == arcaneWrath }) {
				t.Errorf("line %d: the Presence of Mind Arcane Wrath must wait for spellIsReady(%d)", i, arcaneWrath)
			}
			if !mentionsSpellKnown(action.GetCondition(), pyroblast, false) {
				t.Errorf("line %d: the Presence of Mind Arcane Wrath must give way to Pyroblast once it is known", i)
			}
		case frostbolt:
			frostLine = i
		}
	}
	if wrathLine < 0 || frostLine < 0 || wrathLine > frostLine {
		t.Fatalf("want a Presence of Mind Arcane Wrath line ahead of the Frostbolt one, got Arcane Wrath %d, Frostbolt %d", wrathLine, frostLine)
	}
}
