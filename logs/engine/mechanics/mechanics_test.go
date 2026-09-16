// logs/engine/mechanics/mechanics_test.go
package mechanics

import (
	"strconv"
	"testing"
)

func TestParseAcceptsTheSpecFormatAndRejectsAnUnknownKind(t *testing.T) {
	good := []byte(`{"encounter_id": 2363, "name": "General Kaal", "mechanics": [
		{"spell_id": 331415, "name": "Wicked Gash", "kind": "avoidable", "note": "Frontal."}]}`)
	table, err := Parse(good)
	if err != nil {
		t.Fatal(err)
	}
	if table.EncounterID != 2363 || len(table.Mechanics) != 1 || table.Mechanics[0].Kind != Avoidable {
		t.Fatalf("table = %+v", table)
	}
	if _, ok := table.Lookup(331415); !ok {
		t.Fatal("Lookup by spell id must find the mechanic")
	}
	bad := []byte(`{"encounter_id": 1, "name": "X", "mechanics": [{"spell_id": 5, "name": "Y", "kind": "soak"}]}`)
	if _, err := Parse(bad); err == nil {
		t.Fatal("an unknown kind must be refused")
	}
	zero := []byte(`{"encounter_id": 1, "name": "X", "mechanics": [{"spell_id": 0, "name": "Y", "kind": "avoidable"}]}`)
	if _, err := Parse(zero); err == nil {
		t.Fatal("a spell id of zero must be refused")
	}
}

func TestLoadFindsAnEmbeddedTableAndSaysWhenThereIsNone(t *testing.T) {
	if table, ok := Load(2363); !ok || table.Name != "General Kaal" {
		t.Fatalf("Load(2363) = %+v, %v", table, ok)
	}
	if _, ok := Load(424242); ok {
		t.Fatal("an encounter without a table must report false")
	}
}

// TestEveryEmbeddedTableParses walks the embed FS rather than naming the
// tables, so a new one is covered the day it is added and a malformed one
// fails here rather than the first time that boss is pulled.
func TestEveryEmbeddedTableParses(t *testing.T) {
	entries, err := tables.ReadDir("tables")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) == 0 {
		t.Fatal("no embedded tables: the go:embed pattern matched nothing")
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
			if want := strconv.FormatInt(table.EncounterID, 10) + ".json"; entry.Name() != want {
				t.Fatalf("%s holds encounter %d, so Load would never find it; want the file named %s",
					entry.Name(), table.EncounterID, want)
			}
			if len(table.Mechanics) == 0 {
				t.Fatalf("%s lists no mechanics", entry.Name())
			}
			if table.Name == "" {
				t.Fatalf("%s has no encounter name", entry.Name())
			}
		})
	}
}

func TestParseRejectsADuplicateSpellID(t *testing.T) {
	dup := []byte(`{"encounter_id": 1, "name": "X", "mechanics": [
		{"spell_id": 5, "name": "Y", "kind": "avoidable"},
		{"spell_id": 5, "name": "Y again", "kind": "interrupt"}]}`)
	if _, err := Parse(dup); err == nil {
		t.Fatal("a spell id listed twice must be refused: Lookup would only ever find the first")
	}
}
