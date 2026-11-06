package raidbosses

import (
	"strings"
	"testing"
)

const sample = `{
  "build": "1.60.1.1",
  "raids": [
    {"key": "raid:barrow-deeps", "name": "Barrow Deeps", "status": "not_in_client",
     "note": "not in client 1.60.1.1", "max_players": null, "bosses": []},
    {"key": "raid:onyxias-lair", "name": "Onyxia's Lair", "status": "in_client", "note": null,
     "max_players": 40, "bosses": [
      {"encounter_id": 1084, "journal_encounter_id": null, "name": "Onyxia", "order": 1,
       "loot_source": "loot.json raid:onyxias-lair:10184 (npc 10184, classic-db)",
       "loot": [{"item_id": 17064, "name": "Shard of the Scale", "quality": 4, "item_level": 71,
                 "required_level": 60, "class_id": 4, "subclass_id": 0, "inventory_type": 12}]}
    ]}
  ]
}`

func TestParseRefusesAnEmptyOrInconsistentFile(t *testing.T) {
	for _, body := range []string{
		`{}`, `{"build": "x"}`, `{"raids": [{"key": "a"}]}`, `not json`,
		`{"build": "x", "raids": [{"key": "raid:a", "status": "in_client", "bosses": []}]}`,
	} {
		if _, err := Parse([]byte(body)); err == nil {
			t.Errorf("Parse(%q) took it, want an error", body)
		}
	}
}

func TestGoSourceRendersOnlyTheRaidsTheClientNames(t *testing.T) {
	doc, err := Parse([]byte(sample))
	if err != nil {
		t.Fatal(err)
	}
	src, err := GoSource(doc, SourcePath(doc.Build))
	if err != nil {
		t.Fatal(err)
	}
	got := string(src)
	for _, want := range []string{
		"from data/builds/1.60.1.1/raid-bosses.json. DO NOT EDIT.",
		`const raidBossesBuild = "1.60.1.1"`,
		"// Not in this client, so absent: Barrow Deeps (not_in_client).",
		`{Key: "raid:onyxias-lair", Name: "Onyxia's Lair", MaxPlayers: 40,`,
		`{EncounterID: 1084, JournalEncounterID: 0, Name: "Onyxia", Order: 1,`,
		`{ItemID: 17064, Name: "Shard of the Scale", Quality: 4, ItemLevel: 71, RequiredLevel: 60, ItemClass: 4, Subclass: 0, InventoryType: 12},`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("source lacks %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, `Key: "raid:barrow-deeps"`) {
		t.Errorf("a raid not in the client is in the table:\n%s", got)
	}
}
