package consumables

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// The active build's tables, relative to this package.
const (
	activeBuildPath = "../../../../web/src/data/active-build.json"
	buildsDir       = "../../../../data/builds"
)

// buildFlaskIDs is every flask spell id the active build ships: each "Flask of" item's use
// spell in simconsumes.json (a level-0 drink named "Flask of ..." is not a flask), and
// every spells.json id that carries one of those flasks' names.
func buildFlaskIDs(t *testing.T) map[int64]string {
	t.Helper()
	var active struct {
		Build string `json:"build"`
	}
	readJSON(t, activeBuildPath, &active)
	if active.Build == "" {
		t.Fatal("active-build.json names no build")
	}
	var items []struct {
		Name          string  `json:"name"`
		RequiredLevel int     `json:"required_level"`
		SpellIDs      []int64 `json:"spell_ids"`
	}
	readJSON(t, filepath.Join(buildsDir, active.Build, "simconsumes.json"), &items)
	var spells []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	readJSON(t, filepath.Join(buildsDir, active.Build, "spells.json"), &spells)
	out := map[int64]string{}
	names := map[string]bool{}
	for _, it := range items {
		if !strings.HasPrefix(it.Name, "Flask of") || it.RequiredLevel <= 0 {
			continue
		}
		names[it.Name] = true
		for _, id := range it.SpellIDs {
			out[id] = it.Name
		}
	}
	for _, s := range spells {
		if _, ok := out[s.ID]; !ok && names[s.Name] {
			out[s.ID] = s.Name
		}
	}
	if len(out) == 0 {
		t.Fatal("the build ships no flask: the tables moved")
	}
	return out
}

func readJSON(t *testing.T, path string, v any) {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, v); err != nil {
		t.Fatalf("%s: %v", path, err)
	}
}

// A role's flask list marked complete lists every flask the build ships (final review A2):
// a player pulling with an unlisted flask would otherwise read as "No flask at the pull".
func TestCatalogue_ACompleteFlaskListCoversEveryFlaskTheBuildShips(t *testing.T) {
	want := buildFlaskIDs(t)
	c := Load()
	for role, rc := range c.Roles {
		if !rc.IsComplete(CategoryFlask) {
			continue
		}
		have := make([]int64, 0, len(rc.Flask))
		for _, e := range rc.Flask {
			have = append(have, e.SpellID)
		}
		for id, name := range want {
			if !slices.Contains(have, id) {
				t.Errorf("%s: the flask list is complete but lacks %s (%d)", role, name, id)
			}
		}
	}
}
