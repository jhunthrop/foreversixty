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

// A raid Arms warrior Charges on the pull, never before it: at -2 s the
// Charge reached the boss ahead of the tank (round 13, wow-player review).
// The pull is time 0 in every band, so one 0 s Charge serves the raid,
// where the tank has just engaged, and the bare bands, where the warrior
// pulls with it. The pre-pull Battle Shout follows the Charge, whose rage
// pays for it.
func TestArmsChargesOnThePullNotBefore(t *testing.T) {
	const battleShout, charge = 11551, 11578
	curated, err := loadLadderCurated(repoRoot, "warrior-arms")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	chargeIndex, shoutIndex := -1, -1
	for i, action := range rot.PrepullActions {
		at, err := time.ParseDuration(action.GetDoAtValue().GetConst().GetVal())
		if err != nil {
			t.Fatalf("pre-pull time %q: %v", action.GetDoAtValue().GetConst().GetVal(), err)
		}
		switch action.GetAction().GetCastSpell().GetSpellId().GetSpellId() {
		case charge:
			chargeIndex = i
			if at != 0 {
				t.Errorf("Charge at %v, want 0s (on the pull, not before the tank)", at)
			}
		case battleShout:
			shoutIndex = i
			if at != 0 {
				t.Errorf("pre-pull Battle Shout at %v, want 0s, after the Charge that pays for it", at)
			}
		}
	}
	if chargeIndex < 0 {
		t.Fatal("no Charge on the pull")
	}
	if shoutIndex >= 0 && shoutIndex < chargeIndex {
		t.Errorf("pre-pull Battle Shout (action %d) comes before the Charge (action %d)", shoutIndex, chargeIndex)
	}
}
