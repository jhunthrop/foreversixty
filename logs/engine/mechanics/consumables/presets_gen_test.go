package consumables

import (
	"bytes"
	"encoding/json"
	"flag"
	"os"
	"slices"
	"strings"
	"testing"
)

var updatePresets = flag.Bool("update-presets", false, "rewrite presets.json from sim/request/raid-consumes.json")

// TestPresetsJSONMatchesRaidConsumes regenerates presets.json from raid-consumes.json (the
// /sim page's Raid-buffed preset, which this package cannot import: sim/request pulls the
// engine), presets.json's labels and the catalogue, and fails when the committed file
// differs. Run with -update-presets after a preset or catalogue change.
func TestPresetsJSONMatchesRaidConsumes(t *testing.T) {
	var src struct {
		Specs map[string][]struct {
			FromLevel int64    `json:"from_level"`
			Consumes  []string `json:"consumes"`
		} `json:"specs"`
	}
	readJSON(t, raidConsumesPath, &src)
	classify := presetClassifier(Load(), presetLabels(t))
	out := PresetTable{Source: "sim/request/raid-consumes.json", Specs: map[string][]PresetBand{}}
	for spec, bands := range src.Specs {
		for _, b := range bands {
			band := PresetBand{FromLevel: b.FromLevel, Expects: map[string][]string{}}
			for _, id := range b.Consumes {
				if cat := classify(id); cat != "" {
					band.Expects[cat] = append(band.Expects[cat], id)
				}
			}
			for _, ids := range band.Expects {
				slices.Sort(ids)
			}
			out.Specs[spec] = append(out.Specs[spec], band)
		}
	}
	want, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	want = append(want, '\n')
	if *updatePresets {
		if err := os.WriteFile("presets.json", want, 0o600); err != nil {
			t.Fatal(err)
		}
		return
	}
	got, err := presetsFile.ReadFile("presets.json")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Fatal("presets.json is stale against sim/request/raid-consumes.json or the catalogue: go test -run TestPresetsJSONMatchesRaidConsumes -update-presets")
	}
}

// The preset's labels are the catalogue's names: each id below resolves to its entry.
func TestPresetClassifier_MapsByTheCataloguesNames(t *testing.T) {
	labels := presetLabels(t)
	classify := presetClassifier(Load(), labels)
	all := Load().AllRoles()
	for id, want := range map[string]string{
		"flask_of_supreme_power":                     CategoryFlask,
		"food_smoked_desert_dumpling":                CategoryFood,
		"food_sages_tea":                             CategoryFood,
		"main_hand_imbue:elemental_sharpening_stone": CategoryWeaponEnchant,
		"main_hand_imbue:brilliant_wizard_oil":       CategoryWeaponEnchant,
		"major_frenzy_potion":                        CategoryCombatPotion,
		"frenzy_potion":                              CategoryCombatPotion,
		"mighty_rage_potion":                         CategoryCombatPotion,
	} {
		if !slices.ContainsFunc(all.Entries(want), func(e Entry) bool { return e.Name == labels[id] }) {
			t.Errorf("%s: label %q is no %s entry in the catalogue", id, labels[id], want)
		}
		if got := classify(id); got != want {
			t.Errorf("%s: %q, want %q", id, got, want)
		}
	}
	// Outside the catalogue the id's shape decides: an oil or stone on a weapon is a weapon
	// enchant; a class imbue, an elixir, a juju or a rune is none of the four.
	for id, want := range map[string]string{
		"main_hand_imbue:brilliant_mana_oil":         CategoryWeaponEnchant,
		"main_hand_stone:elemental_sharpening_stone": CategoryWeaponEnchant,
		"off_hand_stone:dense_sharpening_stone":      CategoryWeaponEnchant,
		"main_hand_imbue:instant_poison":             "",
		"main_hand_imbue:rockbiter_weapon":           "",
		"elixir_of_the_mongoose":                     "",
		"conjured_demonic_rune":                      "",
		"juju_might":                                 "",
	} {
		if got := classify(id); got != want {
			t.Errorf("%s: %q, want %q", id, got, want)
		}
	}
}

// weaponSlots are the preset id prefixes that put an item on a weapon.
var weaponSlots = []string{"main_hand_imbue", "off_hand_imbue", "main_hand_stone", "off_hand_stone"}

// presetCatalogueCategories are the categories a raid preset can carry; world buffs never
// sit in one.
var presetCatalogueCategories = []string{CategoryFlask, CategoryFood, CategoryWeaponEnchant, CategoryCombatPotion}

// presetClassifier answers the catalogue category a preset id belongs to: the category of
// the catalogue entry named by the id's label, else by the id's shape (an oil or stone on a
// weapon slot, flask_, food_, _potion), else "" (elixirs, jujus, runes, class imbues).
func presetClassifier(c Catalogue, labels map[string]string) func(id string) string {
	all := c.AllRoles()
	byName := map[string]string{}
	for _, cat := range presetCatalogueCategories {
		for _, e := range all.Entries(cat) {
			byName[e.Name] = cat
		}
	}
	return func(id string) string {
		if cat, ok := byName[labels[id]]; ok {
			return cat
		}
		return presetCategoryByShape(id)
	}
}

func presetCategoryByShape(id string) string {
	slot, item, onWeapon := strings.Cut(id, ":")
	switch {
	case onWeapon && slices.Contains(weaponSlots, slot) && (strings.HasSuffix(item, "_oil") || strings.HasSuffix(item, "_stone")):
		return CategoryWeaponEnchant
	case strings.HasPrefix(id, "flask_"):
		return CategoryFlask
	case strings.HasPrefix(id, "food_"):
		return CategoryFood
	case strings.HasSuffix(id, "_potion"):
		return CategoryCombatPotion
	}
	return ""
}
