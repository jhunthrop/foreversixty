// logs/engine/mechanics/encounters_test.go
package mechanics

import (
	"encoding/json"
	"os"
	"testing"
)

// encounterList is encounters.json: the ids read off the Forever client's
// DungeonEncounter table, with the build they came from. It is the answer to
// "where did this number come from", and the test below makes a table whose
// id is not in it fail rather than sit there addressing nothing.
type encounterList struct {
	Build     string `json:"build"`
	Instances []struct {
		Name       string `json:"name"`
		Encounters []struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"encounters"`
	} `json:"instances"`
}

func loadEncounterList(t *testing.T) encounterList {
	t.Helper()
	data, err := os.ReadFile("encounters.json")
	if err != nil {
		t.Fatal(err)
	}
	var list encounterList
	if err := json.Unmarshal(data, &list); err != nil {
		t.Fatal(err)
	}
	if list.Build == "" {
		t.Fatal("encounters.json must say which client build the ids were read from")
	}
	return list
}

// TestEveryTableIsAnEncounterTheClientKnows walks the embedded tables and
// insists each one's id and name appear in encounters.json. A table named for
// an id no build carries -- a typo, or an id that moved between builds -- would
// otherwise just never load, and the page would be quietly empty for that boss.
// The name has to match too, because a health_pct phase trigger is matched
// against the boss unit's own name, which is the name the client's table gives.
func TestEveryTableIsAnEncounterTheClientKnows(t *testing.T) {
	list := loadEncounterList(t)
	known := map[int64]string{}
	for _, inst := range list.Instances {
		for _, e := range inst.Encounters {
			if prior, dup := known[e.ID]; dup {
				t.Fatalf("encounters.json lists %d twice (%q and %q)", e.ID, prior, e.Name)
			}
			known[e.ID] = e.Name
		}
	}
	entries, err := tables.ReadDir("tables")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		t.Run(entry.Name(), func(t *testing.T) {
			data, err := tables.ReadFile("tables/" + entry.Name())
			if err != nil {
				t.Fatal(err)
			}
			table, err := Parse(data)
			if err != nil {
				t.Fatal(err)
			}
			name, ok := known[table.EncounterID]
			if !ok {
				t.Fatalf("encounter %d (%q) is in no instance in encounters.json, read from build %s;"+
					" either the id is wrong or the list is stale", table.EncounterID, table.Name, list.Build)
			}
			if name != table.Name {
				t.Errorf("encounter %d is %q in encounters.json and %q in the table;"+
					" a health_pct phase matches the client's name, so they must agree", table.EncounterID, name, table.Name)
			}
		})
	}
}

// TestEveryPhaseTriggerIsUsable holds the phase triggers to what the engine can
// actually distinguish. Parse enforces the shape of a trigger; this enforces
// that a table's triggers are different from each other and in the order the
// fight reaches them, which Parse has no way to know and which a reader of the
// rendered phases would be misled by.
func TestEveryPhaseTriggerIsUsable(t *testing.T) {
	entries, err := tables.ReadDir("tables")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		data, err := tables.ReadFile("tables/" + entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		table, err := Parse(data)
		if err != nil {
			t.Fatal(err)
		}
		seen := map[string]bool{}
		// Health thresholds have to fall as the list goes on: the phases are in the
		// order they happen, and a boss only loses health, so a later phase written
		// at a higher percentage would fire before the one above it.
		lastHealth := 100.0
		for _, p := range table.Phases {
			if seen[p.Name] {
				t.Errorf("%s: two phases are both called %q", entry.Name(), p.Name)
			}
			seen[p.Name] = true
			if p.Starts.HealthPct == 0 {
				continue
			}
			if p.Starts.HealthPct >= lastHealth {
				t.Errorf("%s: phase %q starts at %.0f%% health, at or above the %.0f%% before it;"+
					" a boss passes thresholds on the way down, so this one fires first or never",
					entry.Name(), p.Name, p.Starts.HealthPct, lastHealth)
			}
			lastHealth = p.Starts.HealthPct
		}
	}
}
