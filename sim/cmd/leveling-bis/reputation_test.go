package main

import (
	"path/filepath"
	"testing"
)

// Round 9: opening the Darkspear Islands gate made Creeg Bothunk's stock
// eligible, and his loot.json vendor row states no standing and no side.
// The client does: ItemSparse.MinFactionID 2798 (Darkspear Raiders, a
// Horde Forces faction) and MinReputation 4 to 7 per item. Raider Gauntlets
// (272096, level 50) need Exalted, so a band-50 list may not name them, and
// an Alliance list may never name them.
func TestAnItemsClientReputationGatesEverySource(t *testing.T) {
	idx := lootIndex{
		272096: {{Kind: "vendor", Label: "Creeg Bothunk"}},
		272067: {{Kind: "vendor", Label: "Creeg Bothunk"}},
		15000:  {{Kind: "vendor", Label: "Timbermaw quartermaster"}},
		16000:  {{Kind: "vendor", Label: "Timbermaw quartermaster"}},
	}
	items := []candidate{
		{ID: 272096, Reputation: &itemReputation{FactionID: 2798, Standing: "exalted"}},
		{ID: 272067, Reputation: &itemReputation{FactionID: 2798, Standing: "honored"}},
		{ID: 15000, Reputation: &itemReputation{FactionID: 576, Standing: "revered"}},
		{ID: 1},
		{ID: 16000, Reputation: &itemReputation{FactionID: 576, Standing: "neutral"}},
	}
	gated := applyItemReputations(idx, items)
	cases := []struct {
		item, level int
		faction     string
		want        bool
	}{
		{272096, 50, "horde", false},    // Exalted while leveling
		{272096, 60, "horde", true},     // Exalted at 60
		{272096, 60, "alliance", false}, // a Horde Forces reputation
		{272067, 30, "horde", true},     // Honored while leveling
		{272067, 30, "alliance", false},
		{15000, 50, "alliance", false}, // Revered while leveling, any side
		{15000, 60, "alliance", true},
		{16000, 20, "horde", true}, // neutral gates nothing
	}
	for _, c := range cases {
		if _, got := sourceFor(c.item, c.level, c.faction, "", gated); got != c.want {
			t.Errorf("item %d at %d for %s: obtainable %v, want %v", c.item, c.level, c.faction, got, c.want)
		}
	}
	if idx[272096][0].Standing != "" {
		t.Error("applyItemReputations changed the index it was given")
	}
}

// A source that already states a stricter standing keeps it.
func TestAnItemsReputationNeverLoosensASourcesStanding(t *testing.T) {
	idx := lootIndex{7: {{Kind: "rep", Standing: "exalted", Side: "alliance"}}}
	gated := applyItemReputations(idx, []candidate{{ID: 7, Reputation: &itemReputation{FactionID: 2799, Standing: "honored"}}})
	if s := gated[7][0]; s.Standing != "exalted" || s.Side != "alliance" {
		t.Errorf("source = %+v, want exalted and alliance kept", s)
	}
}

func TestTheClassFileCarriesTheItemsReputation(t *testing.T) {
	dir := t.TempDir()
	if err := writeFile(t, filepath.Join(dir, "items.json"), `[{"id": 272096, "name": "Raider Gauntlets", "inventory_type": 10}]`); err != nil {
		t.Fatal(err)
	}
	row := `{"build": "b", "class_slug": "warrior", "items": [{"id": 272096, "name": "Raider Gauntlets", "slot": "hands", "quality": 4, "required_level": 50, "reputation": {"faction_id": 2798, "standing": "exalted"}}]}`
	if err := writeFile(t, filepath.Join(dir, "items", "warrior.json"), row); err != nil {
		t.Fatal(err)
	}
	got, _, err := loadClassCandidates(dir, "warrior")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].Reputation == nil || *got[0].Reputation != (itemReputation{FactionID: 2798, Standing: "exalted"}) {
		t.Fatalf("candidates = %+v, want the reputation carried", got)
	}
}
