package request

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

// A hunter carries the best ammo and quiver its level can use (Wowhead
// Forever item tooltips, read 2026-11-02): Sharp Arrow 3.5 DPS (level 10),
// Razor Arrow 7.5 (25), Jagged Arrow 13 (40), Mithril Gyro-Shot 15 (44),
// Ice Threaded Arrow 16.5 (51), Thorium Headed Arrow 17.5 (52); Light
// Quiver 10%, Heavy Quiver 12% (30), Quickdraw Quiver 13% (40), Gnoll Skin
// Bandolier 15% (55). Every band was simmed with level-52 arrows and a
// level-55 bandolier.
func TestAHuntersAmmoAndQuiverFollowItsLevel(t *testing.T) {
	for _, tc := range []struct {
		level  int32
		ammo   proto.Hunter_Options_Ammo
		quiver proto.Hunter_Options_QuiverBonus
	}{
		{9, proto.Hunter_Options_AmmoNone, proto.Hunter_Options_Speed10},
		{10, proto.Hunter_Options_SharpArrow, proto.Hunter_Options_Speed10},
		{20, proto.Hunter_Options_SharpArrow, proto.Hunter_Options_Speed10},
		{25, proto.Hunter_Options_RazorArrow, proto.Hunter_Options_Speed10},
		{30, proto.Hunter_Options_RazorArrow, proto.Hunter_Options_Speed12},
		{40, proto.Hunter_Options_JaggedArrow, proto.Hunter_Options_Speed13},
		{44, proto.Hunter_Options_MithrilGyroShot, proto.Hunter_Options_Speed13},
		{50, proto.Hunter_Options_MithrilGyroShot, proto.Hunter_Options_Speed13},
		{51, proto.Hunter_Options_IceThreadedArrow, proto.Hunter_Options_Speed13},
		{52, proto.Hunter_Options_ThoriumHeadedArrow, proto.Hunter_Options_Speed13},
		{55, proto.Hunter_Options_ThoriumHeadedArrow, proto.Hunter_Options_Speed15},
		{60, proto.Hunter_Options_ThoriumHeadedArrow, proto.Hunter_Options_Speed15},
	} {
		p := &proto.Player{Level: tc.level}
		hunterOptions(p)
		got := p.GetHunter().GetOptions()
		if got.Ammo != tc.ammo || got.QuiverBonus != tc.quiver {
			t.Errorf("level %d: ammo %v quiver %v, want %v %v", tc.level, got.Ammo, got.QuiverBonus, tc.ammo, tc.quiver)
		}
	}
}
