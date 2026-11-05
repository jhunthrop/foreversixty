package request

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

// Round 22 platform (druid round 21 C2): Rend and Tear reads "on Bleeding
// targets", and any bleed counts. In a raid another player's bleed (a
// warrior's Deep Wounds, a rogue's Rupture) is on the boss from the opening
// seconds, so the raid fight's feral and bear assume a bleeding target; the
// bare fight is solo and counts only the druid's own bleeds.
func TestTheRaidFightAssumesAnotherPlayersBleed(t *testing.T) {
	for _, tc := range []struct {
		spec, preset string
		want         bool
	}{
		{"druid-feral", RaidPreset, true},
		{"druid-feral", BarePreset, false},
		{"druid-feral-bear", RaidPreset, true},
		{"druid-feral-bear", BarePreset, false},
	} {
		player := &proto.Player{Class: proto.Class_ClassDruid, Race: proto.Race_RaceTauren, Level: 60}
		if err := applySpec(player, tc.spec, "druid", tc.preset, "tauren", 60); err != nil {
			t.Fatal(err)
		}
		var got bool
		switch spec := player.Spec.(type) {
		case *proto.Player_FeralDruid:
			got = spec.FeralDruid.GetOptions().GetAssumeBleedActive()
		case *proto.Player_FeralTankDruid:
			got = spec.FeralTankDruid.GetOptions().GetAssumeBleedActive()
		default:
			t.Fatalf("%s: unexpected spec %T", tc.spec, spec)
		}
		if got != tc.want {
			t.Errorf("%s, %s preset: assume_bleed_active = %v, want %v", tc.spec, tc.preset, got, tc.want)
		}
	}
}
