package request

import (
	"testing"
	"time"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// Charge "Cannot be used in combat", and the client still ships Bloodrage's
// "considered in combat" effect (29131 effect 1, aura 94, which
// cmangos-classic turns into the in-combat flag). So no warrior opener
// Charges after a pre-pull Bloodrage: Charge comes first, and Bloodrage
// waits for the pull (round 12 wow-player review of the regression-melee
// lane; the landing's Arms opener was Bloodrage, Battle Shout, Charge).
func TestNoWarriorChargesAfterAPrepullBloodrage(t *testing.T) {
	const bloodrage, charge = 2687, 11578
	for _, spec := range []string{"warrior-arms", "warrior-fury", "warrior-protection"} {
		curated, err := loadLadderCurated(repoRoot, spec)
		if err != nil {
			t.Fatal(err)
		}
		rot := &proto.APLRotation{}
		if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
			t.Fatal(err)
		}
		var bloodrageAt, chargeAt *time.Duration
		for _, action := range rot.PrepullActions {
			at, err := time.ParseDuration(action.GetDoAtValue().GetConst().GetVal())
			if err != nil {
				t.Fatalf("%s: pre-pull time %q: %v", spec, action.GetDoAtValue().GetConst().GetVal(), err)
			}
			switch action.GetAction().GetCastSpell().GetSpellId().GetSpellId() {
			case bloodrage:
				bloodrageAt = &at
			case charge:
				chargeAt = &at
			}
		}
		if bloodrageAt != nil && chargeAt != nil && *chargeAt >= *bloodrageAt {
			t.Errorf("%s: Charge at %v comes after a pre-pull Bloodrage at %v", spec, *chargeAt, *bloodrageAt)
		}
	}
}
