package guilds

import (
	"encoding/json"
	"os"
	"testing"
)

// raidSuppliesPath is the one list of counted consumables the site and the addon share.
const raidSuppliesPath = "../../../data/curated/raid-supplies.json"

// raidSupplies is a copy of data/curated/raid-supplies.json's supplies (the API image carries no
// data directory); the two must be the same item for item, name and kind.
func TestRaidSuppliesMatchTheCuratedCatalogue(t *testing.T) {
	raw, err := os.ReadFile(raidSuppliesPath)
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Supplies []struct {
			ItemID int    `json:"item_id"`
			Name   string `json:"name"`
			Kind   string `json:"kind"`
		} `json:"supplies"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	if len(file.Supplies) != len(raidSupplies) {
		t.Errorf("%s lists %d supplies, raidSupplies %d", raidSuppliesPath, len(file.Supplies), len(raidSupplies))
	}
	for _, s := range file.Supplies {
		got, ok := raidSupplies[s.ItemID]
		if !ok || got.Name != s.Name || got.Kind != s.Kind {
			t.Errorf("item %d: curated %q (%s), raidSupplies %+v", s.ItemID, s.Name, s.Kind, got)
		}
	}
}
