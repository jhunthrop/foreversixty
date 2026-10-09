package request

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

// The published Marksmanship build (no Lone Wolf) and the same build with the
// Lone Wolf point set and the pet points moved, in the engine's own string form.
const (
	marksmanshipWithPet  = "5320000501000000-0051550001503050-500000000000000000"
	marksmanshipLoneWolf = "5000000000000000-0053550011503050-500050030000000000"
)

func TestAHunterBringsTheCatUnlessTheBuildTakesLoneWolf(t *testing.T) {
	for _, tc := range []struct {
		talents string
		pet     proto.Hunter_Options_PetType
		uptime  float64
	}{
		{marksmanshipWithPet, proto.Hunter_Options_Cat, fullPetUptime},
		{marksmanshipLoneWolf, proto.Hunter_Options_PetNone, 0},
		{"", proto.Hunter_Options_Cat, fullPetUptime},
	} {
		p := &proto.Player{TalentsString: tc.talents}
		hunterOptions(p)
		got := p.GetHunter().GetOptions()
		if got.PetType != tc.pet || got.PetUptime != tc.uptime {
			t.Errorf("talents %q: pet %v uptime %v, want %v %v", tc.talents, got.PetType, got.PetUptime, tc.pet, tc.uptime)
		}
	}
}
