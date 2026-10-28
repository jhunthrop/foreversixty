package mechanics

import "testing"

func TestAll_EveryEmbeddedTableLoads(t *testing.T) {
	all, err := All()
	if err != nil || len(all) == 0 {
		t.Fatalf("%d tables, err %v", len(all), err)
	}
	for _, table := range all {
		if got, ok := Load(table.EncounterID); !ok || got.Name != table.Name {
			t.Fatalf("%s does not load by its encounter id", table.Name)
		}
	}
}
