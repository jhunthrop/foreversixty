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
// actually fire on: a health threshold needs the boss's own name, which the
// test above pins to the client's, and a spell trigger needs an id and a known
// "on". Parse enforces the shape; this says the shape is enough to fire.
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
		for _, p := range table.Phases {
			if seen[p.Name] {
				t.Errorf("%s: two phases are both called %q", entry.Name(), p.Name)
			}
			seen[p.Name] = true
			if p.Starts.HealthPct > 0 && table.Name == "" {
				t.Errorf("%s: a health trigger is matched against the boss's name, which this table does not give",
					entry.Name())
			}
		}
	}
}
