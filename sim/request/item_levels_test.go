package request

import (
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// loadRealItemLevels reads the active build's items.json, the client's ItemSparse rows.
func loadRealItemLevels(t *testing.T) ItemLevels {
	t.Helper()
	build, err := leveling.ReadActiveBuild("../..")
	if err != nil {
		t.Fatal(err)
	}
	levels, err := LoadItemLevelsFile(filepath.Join("../../data/builds", build, "items.json"))
	if err != nil {
		t.Fatal(err)
	}
	return levels
}

func TestItemLevelsReadTheClientsRequiredLevel(t *testing.T) {
	levels := loadRealItemLevels(t)
	cases := []struct {
		id    string
		level int
		known bool
	}{
		{"elixir_of_the_mongoose", 46, true},
		{"flask_of_the_titans", 50, true},
		{"major_mana_potion", 49, true},
		{"food_grilled_squid", 35, true},
		{"main_hand_imbue:elemental_sharpening_stone", 50, true},
		{"off_hand_stone:dense_sharpening_stone", 35, true},
		{"main_hand_imbue:deadly_poison", 30, true},
		{"conjured_rogue_thistle_tea", 5, true},
		{"elixir_of_firepower", 18, true},
		{"juju_power", 0, true},
		// A class's own weapon spell is no item: nothing to gate.
		{"main_hand_imbue:windfury_weapon", 0, false},
	}
	for _, c := range cases {
		got, ok := levels.RequiredLevel(c.id)
		if ok != c.known || got != c.level {
			t.Errorf("RequiredLevel(%q) = %d, %v; want %d, %v", c.id, got, ok, c.level, c.known)
		}
	}
}

// A character never carries a consumable its level cannot use (ItemSparse RequiredLevel).
func TestRaidConsumesCarryOnlyUsableItems(t *testing.T) {
	levels := loadRealItemLevels(t)
	all, err := RaidConsumes(loadRealPresets(t), levels)
	if err != nil {
		t.Fatal(err)
	}
	for spec, rows := range all {
		for level := 1; level <= 60; level++ {
			for _, id := range ConsumesAtLevel(rows, level) {
				if need, ok := levels.RequiredLevel(id); ok && need > level {
					t.Errorf("%s at %d carries %s, which needs level %d", spec, level, id, need)
				}
			}
		}
	}
	fury40 := ConsumesAtLevel(all["warrior-fury"], 40)
	for _, id := range []string{"elixir_of_the_mongoose", "mighty_rage_potion", "main_hand_imbue:elemental_sharpening_stone"} {
		if slices.Contains(fury40, id) {
			t.Errorf("warrior-fury at 40 carries %s: %v", id, fury40)
		}
	}
	if !slices.Contains(fury40, "off_hand_imbue:dense_sharpening_stone") || !slices.Contains(fury40, "juju_power") {
		t.Errorf("warrior-fury at 40 lost a usable item: %v", fury40)
	}
	combat40 := ConsumesAtLevel(all["rogue-combat"], 40)
	if !slices.ContainsFunc(combat40, func(id string) bool { return strings.HasSuffix(id, "deadly_poison") }) {
		t.Errorf("rogue-combat at 40 lost its poison: %v", combat40)
	}
}
