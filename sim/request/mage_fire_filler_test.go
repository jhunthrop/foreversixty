package request

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

const (
	fireScorch      = 10207
	fireFireball    = 10151
	fireFireBlast   = 10199
	fireMissiles    = 10212
	fireArcaneBlast = 1239700
	fireBlastBuff   = 400573
	fireBarrage     = 400589
	fireHeatingUp   = 400625
	fireImpScorch   = 12873
)

// Sim-accuracy round 21 (design/reviews/2026-11-05-sim-audit-round21/mage.md
// F-S1): below Heating Up (both Fire orders take it at 30) a rank 1 or 2
// Scorch does less damage for its mana than Fireball and the band-20 mage
// runs dry, so Fireball is the filler at 22 to 29 and Scorch takes over only
// once Heating Up is known. The Improved Scorch stack line still casts Scorch
// from 27.
//
// Round 22 platform: the below-30 Fireball filler carries the complement of
// the Scorch filler's gate, not auraIsKnown(Heating Up), so the raid card no
// longer prints a "Below 30" Fireball row with the raid rank.
//
// Round 22 mage (F1): the raid build takes Accelerant 5/5, Missile Barrage
// and Arcane Blast, and on it Fireball is the filler. The Scorch filler is
// the leveling path from 30 to 56 (Heating Up known, Arcane Blast not).
func TestFireScorchFillerWaitsForHeatingUp(t *testing.T) {
	rot := fireRotation(t)
	scorchFiller, belowThirty := -1, -1
	for i, item := range rot.PriorityList {
		action := item.GetAction()
		cond := action.GetCondition()
		switch action.GetCastSpell().GetSpellId().GetSpellId() {
		case fireScorch:
			if mentionsAura(cond, fireImpScorch) || mentionsSpellKnown(cond, fireArcaneBlast, true) {
				continue // the Improved Scorch stack line, or the raid build's gap Scorch
			}
			scorchFiller = i
			if !hasConjunct(cond, func(v *proto.APLValue) bool { return knowsAura(v, fireHeatingUp) }) {
				t.Errorf("the Scorch filler (line %d) must be gated on auraIsKnown(%d), got %v", i, fireHeatingUp, cond)
			}
			if !hasConjunct(cond, func(v *proto.APLValue) bool { return knowsSpell(v.GetNot().GetVal(), fireArcaneBlast) }) {
				t.Errorf("the Scorch filler (line %d) must be gated on not spellIsKnown(%d), got %v", i, fireArcaneBlast, cond)
			}
		case fireFireball:
			if knowsAura(cond.GetNot().GetVal(), fireHeatingUp) {
				belowThirty = i
			}
		}
	}
	if scorchFiller < 0 || belowThirty < 0 || scorchFiller > belowThirty {
		t.Fatalf("want the gated Scorch filler above the below-30 Fireball filler, got Scorch %d, Fireball %d", scorchFiller, belowThirty)
	}
}

// Round 22 mage (F1): the raid build's filler is Fireball, cast only when Fire
// Blast is not about to come back (a Scorch fills that gap); Arcane Missiles
// spends every Missile Barrage, and Arcane Blast goes out whenever its buff
// has dropped.
func TestFireRaidBuildFillsWithFireballAndKeepsArcaneBlastUp(t *testing.T) {
	rot := fireRotation(t)
	var fireball, gapScorch, missiles, blast bool
	for _, item := range rot.PriorityList {
		action := item.GetAction()
		cond := action.GetCondition()
		switch action.GetCastSpell().GetSpellId().GetSpellId() {
		case fireFireball:
			fireball = fireball || hasConjunct(cond, func(v *proto.APLValue) bool { return knowsSpell(v, fireArcaneBlast) })
		case fireScorch:
			gapScorch = gapScorch || (hasConjunct(cond, func(v *proto.APLValue) bool { return knowsSpell(v, fireArcaneBlast) }) &&
				hasConjunct(cond, func(v *proto.APLValue) bool {
					return v.GetCmp().GetLhs().GetSpellTimeToReady().GetSpellId().GetSpellId() == fireFireBlast
				}))
		case fireMissiles:
			missiles = missiles || hasConjunct(cond, func(v *proto.APLValue) bool {
				return v.GetAuraIsActive().GetAuraId().GetSpellId() == fireBarrage
			})
		case fireArcaneBlast:
			blast = blast || hasConjunct(cond, func(v *proto.APLValue) bool {
				return v.GetNot().GetVal().GetAuraIsActive().GetAuraId().GetSpellId() == fireBlastBuff
			})
		}
	}
	if !fireball || !gapScorch || !missiles || !blast {
		t.Errorf("want Fireball on spellIsKnown(Arcane Blast) %v, a gap Scorch before Fire Blast %v, Missiles on Missile Barrage %v, Arcane Blast when its buff is down %v",
			fireball, gapScorch, missiles, blast)
	}
}

func fireRotation(t *testing.T) *proto.APLRotation {
	t.Helper()
	curated, err := loadLadderCurated(repoRoot, "mage-fire")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	return rot
}

// hasConjunct reports whether v, or one of its top-level and-terms, satisfies ok.
func hasConjunct(v *proto.APLValue, ok func(*proto.APLValue) bool) bool {
	if v == nil {
		return false
	}
	if ok(v) {
		return true
	}
	for _, sub := range v.GetAnd().GetVals() {
		if ok(sub) {
			return true
		}
	}
	return false
}

func knowsAura(v *proto.APLValue, id int32) bool {
	return v.GetAuraIsKnown().GetAuraId().GetSpellId() == id
}

func knowsSpell(v *proto.APLValue, id int32) bool {
	return v.GetSpellIsKnown().GetSpellId().GetSpellId() == id
}

// mentionsSpellKnown reports whether a top-level term asks spellIsKnown(id)
// with the given polarity.
func mentionsSpellKnown(v *proto.APLValue, id int32, positive bool) bool {
	return hasConjunct(v, func(term *proto.APLValue) bool {
		if positive {
			return knowsSpell(term, id)
		}
		return knowsSpell(term.GetNot().GetVal(), id)
	})
}

func mentionsAura(v *proto.APLValue, id int32) bool {
	if v == nil {
		return false
	}
	if knowsAura(v, id) {
		return true
	}
	for _, sub := range v.GetAnd().GetVals() {
		if mentionsAura(sub, id) {
			return true
		}
	}
	for _, sub := range v.GetOr().GetVals() {
		if mentionsAura(sub, id) {
			return true
		}
	}
	return false
}
