package consumables

import (
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// The published raid preset: each spec's consumables by band (the last row is level 60's)
// and the labels of the ids it names.
const (
	raidConsumesPath = "../../../../sim/request/raid-consumes.json"
	presetsPath      = "../../../../data/curated/presets.json"
	specsPath        = "../../../../data/curated/specs.json"
)

// presetFoodsByRole is, per role, the name of every food a spec of that role eats in the
// published raid preset at its top band (round 8: Nightfin Soup for damage casters, Smoked
// Desert Dumplings for Strength users and tanks, Grilled Squid for rogues and hunters, Sage's
// Tea for healers).
func presetFoodsByRole(t *testing.T) map[string][]string {
	t.Helper()
	return presetConsumablesByRole(t, func(id string) bool { return strings.HasPrefix(id, "food_") })
}

// presetConsumablesByRole is, per role, the label of every consumable keep accepts that a
// spec of that role carries in the published raid preset at its top band.
func presetConsumablesByRole(t *testing.T, keep func(id string) bool) map[string][]string {
	t.Helper()
	var raid struct {
		Specs map[string][]struct {
			Consumes []string `json:"consumes"`
		} `json:"specs"`
	}
	readJSON(t, raidConsumesPath, &raid)
	type labelled struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	}
	var presets struct {
		Presets struct {
			Raid struct {
				Consumes struct {
					Groups map[string]struct {
						IDs []labelled `json:"ids"`
					} `json:"groups"`
					ByClass map[string][]labelled `json:"by_class"`
					BySpec  map[string][]labelled `json:"by_spec"`
				} `json:"consumes"`
			} `json:"raid"`
		} `json:"presets"`
	}
	readJSON(t, presetsPath, &presets)
	// Round 9 moved foods into the class and spec lists too (a level-gated pair per spec),
	// so a food's label can sit in any of the three.
	labels := map[string]string{}
	consumes := presets.Presets.Raid.Consumes
	for _, g := range consumes.Groups {
		for _, id := range g.IDs {
			labels[id.ID] = id.Label
		}
	}
	for _, lists := range []map[string][]labelled{consumes.ByClass, consumes.BySpec} {
		for _, ids := range lists {
			for _, id := range ids {
				labels[id.ID] = id.Label
			}
		}
	}
	var specs []struct {
		Spec string `json:"spec"`
		Role string `json:"role"`
	}
	readJSON(t, specsPath, &specs)
	out := map[string][]string{}
	for _, s := range specs {
		rows := raid.Specs[s.Spec]
		if len(rows) == 0 {
			t.Fatalf("%s: no raid-consumes rows", s.Spec)
		}
		for _, id := range rows[len(rows)-1].Consumes {
			if !keep(id) {
				continue
			}
			label, ok := labels[id]
			if !ok {
				t.Fatalf("%s: %s has no label in presets.json", s.Spec, id)
			}
			if !slices.Contains(out[s.Role], label) {
				out[s.Role] = append(out[s.Role], label)
			}
		}
	}
	if len(out) == 0 {
		t.Fatal("the preset names none of these consumables: the tables moved")
	}
	return out
}

// Each role's food list names every food its specs eat in the published preset, so the
// coach credits the food the site tells them to bring (round 8 moved casters to Nightfin
// Soup, Strength users and tanks to Smoked Desert Dumplings, healers to Sage's Tea).
func TestCatalogue_FoodsAreThePublishedPresetsFoods(t *testing.T) {
	c := Load()
	for role, want := range presetFoodsByRole(t) {
		rc, ok := c.Roles[role]
		if !ok {
			t.Fatalf("the catalogue has no %s role", role)
		}
		var got []string
		for _, e := range rc.Food {
			got = append(got, e.Name)
		}
		for _, name := range want {
			if !slices.Contains(got, name) {
				t.Errorf("%s food %v lacks the preset's %s", role, got, name)
			}
		}
	}
}

// A Forever food's use spell ("Nutritious Food", "Nutritious Drink") is a 10 second aura
// that triggers the 15 minute Well Fed (SpellMisc duration index 1 against 347), so only
// Well Fed is up at the pull. Each listed food carries its item's use spell, the cast, and
// a Well Fed id, the buff.
func TestCatalogue_EachFoodCarriesItsUseSpellAndItsWellFed(t *testing.T) {
	var active struct {
		Build string `json:"build"`
	}
	readJSON(t, activeBuildPath, &active)
	var items []struct {
		Name     string  `json:"name"`
		SpellIDs []int64 `json:"spell_ids"`
	}
	readJSON(t, filepath.Join(buildsDir, active.Build, "simconsumes.json"), &items)
	var spells []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	readJSON(t, filepath.Join(buildsDir, active.Build, "spells.json"), &spells)
	spellName := make(map[int64]string, len(spells))
	for _, s := range spells {
		spellName[s.ID] = s.Name
	}
	useSpells := map[string][]int64{}
	for _, it := range items {
		useSpells[it.Name] = append(useSpells[it.Name], it.SpellIDs...)
	}
	for role, rc := range Load().Roles {
		byName := map[string][]int64{}
		for _, e := range rc.Food {
			byName[e.Name] = append(byName[e.Name], e.SpellID)
		}
		for name, ids := range byName {
			hasUse, hasWellFed := false, false
			for _, id := range ids {
				hasUse = hasUse || slices.Contains(useSpells[name], id)
				hasWellFed = hasWellFed || spellName[id] == "Well Fed"
			}
			if !hasUse {
				t.Errorf("%s %s: none of %v is the item's use spell %v", role, name, ids, useSpells[name])
			}
			if !hasWellFed {
				t.Errorf("%s %s: none of %v is a Well Fed buff, the aura up at the pull", role, name, ids)
			}
		}
	}
}

// Each role's combat potion list names every potion its specs drink in the published preset
// at the top band, so the coach credits the potion the site tells them to drink. Round 10:
// the damage list held only Mighty Rage Potion, so a mage or a Shadow priest drinking the
// preset's Major Mana Potion earned no credit, and the auras lane's 30-second Frenzy and
// Spellblasting potions joined the preset.
func TestCatalogue_PotionsAreThePublishedPresetsPotions(t *testing.T) {
	c := Load()
	for role, want := range presetConsumablesByRole(t, isPotion) {
		var got []string
		for _, e := range c.Roles[role].CombatPotion.Entries {
			got = append(got, e.Name)
		}
		for _, name := range want {
			if !slices.Contains(got, name) {
				t.Errorf("%s combat potions %v lack the preset's %s", role, got, name)
			}
		}
	}
}

func isPotion(id string) bool { return strings.HasSuffix(id, "_potion") }
