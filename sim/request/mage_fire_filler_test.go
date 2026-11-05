package request

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// Sim-accuracy round 21 (design/reviews/2026-11-05-sim-audit-round21/mage.md
// F-S1): below Heating Up (both Fire orders take it at 30) a rank 1 or 2
// Scorch does less damage for its mana than Fireball and the band-20 mage
// runs dry, so Fireball is the filler at 22 to 29 and Scorch takes over only
// once Heating Up is known. The Improved Scorch stack line still casts Scorch
// from 27.
func TestFireScorchFillerWaitsForHeatingUp(t *testing.T) {
	const (
		scorch     = 10207
		fireball   = 10151
		heatingUp  = 400625
		impScorch  = 12873
		fireSpecID = "mage-fire"
	)
	curated, err := loadLadderCurated(repoRoot, fireSpecID)
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	fillerAt, fireballAt := -1, -1
	for i, item := range rot.PriorityList {
		action := item.GetAction()
		switch action.GetCastSpell().GetSpellId().GetSpellId() {
		case scorch:
			cond := action.GetCondition()
			if mentionsAura(cond, impScorch) {
				continue // the Improved Scorch stack line
			}
			fillerAt = i
			if !knowsAura(cond, heatingUp) {
				t.Errorf("the Scorch filler (line %d) must be gated on auraIsKnown(%d), got %v", i, heatingUp, cond)
			}
		case fireball:
			fireballAt = i
			if action.GetCondition() != nil {
				t.Errorf("the Fireball filler (line %d) must be unconditional, got %v", i, action.GetCondition())
			}
		}
	}
	if fillerAt < 0 || fireballAt < 0 || fillerAt > fireballAt {
		t.Fatalf("want the gated Scorch filler above the Fireball filler, got Scorch %d, Fireball %d", fillerAt, fireballAt)
	}
}

func knowsAura(v *proto.APLValue, id int32) bool {
	return v.GetAuraIsKnown().GetAuraId().GetSpellId() == id
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
