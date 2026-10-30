package discordbot

import (
	"encoding/json"
	"testing"
)

// The loot this reads is a cut of data/builds/1.60.1.70630/loot.json's own shapes.
const sourceLoot = `{
 "sources": [
  {"kind": "vendor", "name": "Pix Xizzix", "items": [272408]},
  {"kind": "vendor", "name": "Captain Dirgehammer", "faction": "alliance", "items": [16466]},
  {"kind": "pvp", "name": "Rank 17 (Alliance)", "rank": 17, "faction": "alliance", "items": [16466]},
  {"kind": "rep", "name": "Stormpike Guard", "standing": "exalted", "items": [19325]},
  {"kind": "crafted", "name": "Leatherworking", "profession": "leatherworking", "items": [252483]},
  {"kind": "dungeon", "name": "Stratholme", "bosses": [{"name": "Baron Rivendare", "items": [13340], "item_chances": {"13340": 11.5}}]},
  {"kind": "dungeon", "name": "Scholomance", "bosses": [{"name": "Darkmaster Gandling", "items": [13398]}]},
  {"kind": "world_drop", "name": "World drops", "level_min": 55, "level_max": 63, "items": [7000]}
 ],
 "quests": {"9533": [{"name": "Divino-matic Rod", "faction": "both", "level": 47, "min_level": 40}],
  "9534": [{"name": "No Requirement", "faction": "both", "level": 47}]}
}`

func TestSourceLinesAreTheBisPagesOwn(t *testing.T) {
	var loot lootFile
	if err := json.Unmarshal([]byte(sourceLoot), &loot); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		id         int
		kind, want string
	}{
		{272408, "vendor", "Vendor: Pix Xizzix"},
		{16466, "vendor", "PvP rank 17 · Field Marshal · Alliance"},
		{19325, "rep", "Stormpike Guard (exalted)"},
		{9533, "quest", "Quest: Divino-matic Rod · requires level 40"},
		{9534, "quest", "Quest: No Requirement"},
		{252483, "crafted", "Crafted: Leatherworking"},
		{13340, "dungeon", "12% from Baron Rivendare"},
		{13398, "dungeon", "Scholomance · Darkmaster Gandling"},
		{7000, "world_drop", "World drop (BoE) · levels 55-63"},
		{1, "vendor", "Vendors"},
		{1, "quest", "Alliance quest"},
	}
	for _, c := range cases {
		if got := describeSource(&loot, c.id, c.kind, "alliance"); got != c.want {
			t.Errorf("%d %s = %q, want %q", c.id, c.kind, got, c.want)
		}
	}
	if got := describeSource(&loot, 1, "quest", "horde"); got != "Horde quest" {
		t.Errorf("horde quest badge = %q", got)
	}
}
