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
	var raid struct {
		Specs map[string][]struct {
			Consumes []string `json:"consumes"`
		} `json:"specs"`
	}
	readJSON(t, raidConsumesPath, &raid)
	var presets struct {
		Presets struct {
			Raid struct {
				Consumes struct {
					Groups map[string]struct {
						IDs []struct {
							ID    string `json:"id"`
							Label string `json:"label"`
						} `json:"ids"`
					} `json:"groups"`
				} `json:"consumes"`
			} `json:"raid"`
		} `json:"presets"`
	}
	readJSON(t, presetsPath, &presets)
	labels := map[string]string{}
	for _, g := range presets.Presets.Raid.Consumes.Groups {
		for _, id := range g.IDs {
			labels[id.ID] = id.Label
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
			if !strings.HasPrefix(id, "food_") {
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
		t.Fatal("the preset names no food: the tables moved")
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
