// logs/engine/mechanics/mechanics_test.go
package mechanics

import "testing"

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
