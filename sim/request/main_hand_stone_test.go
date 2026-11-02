package request

import (
	"slices"
	"strings"
	"testing"
)

// A melee main hand carries the Dense stone from level 35, the level the
// off hand already gets it, until the Elemental Sharpening Stone at 50
// takes the slot. Before, raid rows at 35 to 49 left the main hand bare
// (round 11, found by the regression lane). A rogue's Deadly Poison keeps
// the imbue slot, so its stone goes beside the poison, and only one: at 50
// the Elemental stone, not both.
func TestMeleeMainHandsCarryTheDenseStoneUntilTheElementalStone(t *testing.T) {
	all, err := RaidConsumes(loadRealPresets(t), loadRealItemLevels(t))
	if err != nil {
		t.Fatal(err)
	}
	const (
		dense     = "dense_sharpening_stone"
		elemental = "elemental_sharpening_stone"
	)
	for _, c := range []struct {
		spec  string
		level int
		want  string
		slot  string
	}{
		{"warrior-fury", 34, "", "main_hand_imbue"},
		{"warrior-fury", 35, dense, "main_hand_imbue"},
		{"warrior-fury", 49, dense, "main_hand_imbue"},
		{"warrior-fury", 50, elemental, "main_hand_imbue"},
		{"warrior-arms", 40, dense, "main_hand_imbue"},
		{"warrior-protection", 40, dense, "main_hand_imbue"},
		{"warrior-protection", 60, elemental, "main_hand_imbue"},
		{"paladin-retribution", 40, dense, "main_hand_imbue"},
		{"druid-feral", 40, dense, "main_hand_imbue"},
		{"druid-feral-bear", 40, dense, "main_hand_imbue"},
		{"druid-feral-bear", 60, elemental, "main_hand_imbue"},
		{"hunter-survival", 40, dense, "main_hand_imbue"},
		{"hunter-survival", 60, elemental, "main_hand_imbue"},
		{"rogue-combat", 40, dense, "main_hand_stone"},
		{"rogue-combat", 60, elemental, "main_hand_stone"},
		{"shaman-enhancement", 40, "windfury_weapon", "main_hand_imbue"},
	} {
		var got []string
		for _, id := range ConsumesAtLevel(all[c.spec], c.level) {
			if strings.HasPrefix(id, c.slot+":") {
				got = append(got, strings.TrimPrefix(id, c.slot+":"))
			}
		}
		want := []string{}
		if c.want != "" {
			want = []string{c.want}
		}
		if !slices.Equal(got, want) && !(len(got) == 0 && len(want) == 0) {
			t.Errorf("%s at %d: %s carries %v, want %v", c.spec, c.level, c.slot, got, want)
		}
	}
}
