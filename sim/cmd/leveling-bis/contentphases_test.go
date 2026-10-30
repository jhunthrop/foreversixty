package main

import (
	"io"
	"path/filepath"
	"strings"
	"testing"
)

// testPhases is a phases.json with the three boundaries the gate needs.
const testPhases = `[
	{"name": "pre-beta", "start": "0001-01-01T00:00:00Z"},
	{"name": "launch", "start": "2026-11-04T23:00:00Z"},
	{"name": "raids-1", "start": "2026-12-09T00:00:00Z"}
]`

func testGates(t *testing.T, gates string) contentGates {
	t.Helper()
	g, err := parseContentGates([]byte(`{"sources": [], "notes": "n", "gates": `+gates+`}`), []byte(testPhases))
	if err != nil {
		t.Fatalf("parseContentGates: %v", err)
	}
	return g
}

func gatedIndex(t *testing.T, lootJSON, gates string) (lootIndex, []string) {
	t.Helper()
	dir := t.TempDir()
	if err := writeFile(t, filepath.Join(dir, "loot.json"), lootJSON); err != nil {
		t.Fatal(err)
	}
	idx, _, err := loadLootIndex(dir, nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	return testGates(t, gates).apply(idx)
}

func opensOf(t *testing.T, idx lootIndex, item int) []string {
	t.Helper()
	var out []string
	for _, s := range idx[item] {
		out = append(out, s.Opens)
	}
	return out
}

func TestTheCuratedContentPhasesLoad(t *testing.T) {
	g, err := loadContentGates(publishedRepoRoot)
	if err != nil {
		t.Fatalf("loadContentGates: %v", err)
	}
	// The hand lists data.go used to carry now live in the curated file.
	for _, q := range []int{8756, 8789, 8790} {
		if _, ok := g.quests[q]; !ok {
			t.Errorf("quest %d (an AQ War Effort reward) is not gated", q)
		}
	}
	for _, f := range []int{609, 910, 270} {
		if _, ok := g.factions[f]; !ok {
			t.Errorf("faction %d is not gated", f)
		}
	}
	for _, name := range []string{"Lord Kazzak", "Azuregos", "Emeriss", "Lethon", "Taerar", "Ysondre"} {
		if ref, ok := g.npcNames[name]; !ok || ref.opens != "raids-1" {
			t.Errorf("world boss %s = %+v, want gated to raids-1", name, ref)
		}
	}
	if ref := g.npcs[15192]; ref.opens != "later" {
		t.Errorf("Anachronos (npc 15192) = %+v, want gated later", ref)
	}
}

func TestParseContentGatesFailsFast(t *testing.T) {
	cases := map[string]string{
		"unknown phase":  `[{"key": "a", "phase": "phase-9", "npcs": [{"id": 1, "name": "x"}]}]`,
		"unknown field":  `[{"key": "a", "phase": "later", "npc": [1]}]`,
		"no id":          `[{"key": "a", "phase": "later"}]`,
		"duplicate npc":  `[{"key": "a", "phase": "later", "npcs": [{"id": 1, "name": "x"}]}, {"key": "b", "phase": "later", "npcs": [{"id": 1, "name": "x"}]}]`,
		"bad standing":   `[{"key": "a", "phase": "later", "factions": [{"id": 1, "name": "x", "from_standing": "best"}]}]`,
		"duplicate gate": `[{"key": "a", "phase": "later", "kinds": ["pvp"]}, {"key": "a", "phase": "later", "kinds": ["rep"]}]`,
	}
	for name, gates := range cases {
		if _, err := parseContentGates([]byte(`{"sources": [], "notes": "n", "gates": `+gates+`}`), []byte(testPhases)); err == nil {
			t.Errorf("%s: parseContentGates succeeded, want an error", name)
		}
	}
}

func TestContentGatesCloseEveryKeyType(t *testing.T) {
	loot := `{
		"sources": [
			{"id": "vendor:15192", "kind": "vendor", "name": "Anachronos", "npc_id": 15192, "items": [1]},
			{"id": "world:prince-skaldrenox", "kind": "world", "name": "Prince Skaldrenox", "items": [2]},
			{"id": "world:aged-kodo", "kind": "world", "name": "Aged Kodo", "items": [3]},
			{"id": "dungeon:stratholme", "kind": "dungeon", "name": "Stratholme", "zone_id": 2017, "bosses": [
				{"name": "Balzaphon", "npc_id": 14684, "items": [4]},
				{"name": "Baron Rivendare", "npc_id": 10440, "items": [5]}
			]},
			{"id": "dungeon:dire-maul", "kind": "dungeon", "name": "Dire Maul", "zone_id": 2557, "bosses": [{"name": "Immol'thar", "npc_id": 11496, "items": [6]}]},
			{"id": "pvp:rank-12:horde", "kind": "pvp", "name": "Rank 12", "rank": 12, "faction": "horde", "items": [7]},
			{"id": "rep:silverwing-sentinels:exalted", "kind": "rep", "name": "Silverwing Sentinels", "faction_id": 889, "standing": "exalted", "items": [8]},
			{"id": "rep:timbermaw-hold:honored", "kind": "rep", "name": "Timbermaw Hold", "faction_id": 576, "standing": "honored", "items": [9]},
			{"id": "crafted:leatherworking", "kind": "crafted", "name": "Leatherworking", "items": [10, 11]}
		],
		"quests": {
			"12": [{"quest_id": 9222, "name": "Epic Armaments of Battle - Friend of the Dawn", "faction": "both"}],
			"13": [{"quest_id": 8104, "name": "The Jewel of Kajaro", "faction": "both", "required_rep_faction": 270, "required_rep_standing": "revered"}],
			"14": [{"quest_id": 100, "name": "An Ordinary Quest", "faction": "both"}]
		}
	}`
	gates := `[
		{"key": "npc", "phase": "later", "npcs": [
			{"id": 15192, "name": "Anachronos"}, {"id": 15203, "name": "Prince Skaldrenox"}, {"id": 14684, "name": "Balzaphon"}
		]},
		{"key": "zone", "phase": "later", "zones": [{"id": 2557, "name": "Dire Maul"}]},
		{"key": "honor", "phase": "later", "kinds": ["pvp"]},
		{"key": "rep", "phase": "later", "factions": [{"id": 889, "name": "Silverwing Sentinels"}, {"id": 270, "name": "Zandalar Tribe"}]},
		{"key": "quest", "phase": "later", "quests": [{"id": 9222, "name": "Epic Armaments"}]},
		{"key": "item", "phase": "later", "items": [{"id": 10, "name": "Stormshroud Gloves"}]}
	]`
	idx, missing := gatedIndex(t, loot, gates)
	want := map[int]string{1: "later", 2: "later", 3: "", 4: "later", 5: "", 6: "later", 7: "later", 8: "later", 9: "", 10: "later", 11: "", 12: "later", 13: "later", 14: ""}
	for item, opens := range want {
		if got := opensOf(t, idx, item); len(got) != 1 || got[0] != opens {
			t.Errorf("item %d opens %v, want [%q]", item, got, opens)
		}
	}
	if len(missing) != 0 {
		t.Errorf("missing = %v, want none", missing)
	}
	if _, ok := sourceFor(4, 60, "alliance", "", idx); ok {
		t.Error("Balzaphon's drop is still obtainable at 60; it must not inherit Stratholme's open status")
	}
	if _, ok := sourceFor(5, 60, "alliance", "", idx); !ok {
		t.Error("Baron Rivendare's drop is gated; only the invasion boss should be")
	}
}

func TestContentGatesFromAStandingUp(t *testing.T) {
	loot := `{
		"sources": [
			{"id": "rep:x:friendly", "kind": "rep", "name": "X", "faction_id": 5, "standing": "friendly", "items": [1]},
			{"id": "rep:x:exalted", "kind": "rep", "name": "X", "faction_id": 5, "standing": "exalted", "items": [2]}
		],
		"quests": {}
	}`
	idx, _ := gatedIndex(t, loot, `[{"key": "x", "phase": "later", "factions": [{"id": 5, "name": "X", "from_standing": "revered"}]}]`)
	if got := opensOf(t, idx, 1); got[0] != "" {
		t.Errorf("friendly reward opens %v, want open", got)
	}
	if got := opensOf(t, idx, 2); got[0] != "later" {
		t.Errorf("exalted reward opens %v, want later", got)
	}
}

// One ungated way to an item is enough for the ranker, so a bare copy of
// a quest used to defeat the gate its twin carried (Treasure of the
// Timeless One, The Eye of Zuldazar).
func TestContentGatesCloseEveryCopyOfADuplicateQuest(t *testing.T) {
	loot := `{
		"sources": [],
		"quests": {
			"21520": [
				{"quest_id": 8745, "name": "Treasure of the Timeless One", "faction": "both"},
				{"quest_id": 8745, "name": "Treasure of the Timeless One", "faction": "both", "opens": "later", "required_rep_faction": 910}
			],
			"19593": [
				{"quest_id": 8051, "name": "The Eye of Zuldazar", "faction": "both"},
				{"quest_id": 8051, "name": "The Eye of Zuldazar", "faction": "both", "required_rep_faction": 270, "required_rep_standing": "revered"}
			]
		}
	}`
	idx, _ := gatedIndex(t, loot, `[{"key": "zg", "phase": "later", "factions": [{"id": 270, "name": "Zandalar Tribe"}]}]`)
	for _, item := range []int{21520, 19593} {
		if _, ok := sourceFor(item, 60, "alliance", "", idx); ok {
			t.Errorf("item %d is obtainable through a bare duplicate quest entry", item)
		}
	}
}

// Earthstrike: quest 8573 states no reputation, but the same item is a
// Cenarion Circle exalted reward.
func TestContentGatesCloseAQuestThroughItsItemsReputation(t *testing.T) {
	loot := `{
		"sources": [{"id": "rep:cenarion-circle:exalted", "kind": "rep", "name": "Cenarion Circle", "faction_id": 609, "standing": "exalted", "items": [21180]}],
		"quests": {"21180": [{"quest_id": 8573, "name": "Champion's Battlegear", "faction": "both"}]}
	}`
	idx, _ := gatedIndex(t, loot, `[{"key": "aq", "phase": "later", "factions": [{"id": 609, "name": "Cenarion Circle"}]}]`)
	if got := opensOf(t, idx, 21180); len(got) != 2 || got[0] != "later" || got[1] != "later" {
		t.Errorf("Earthstrike opens %v, want both ways gated later", got)
	}
}

func TestAContentGateDatedAtLaunchIsOpen(t *testing.T) {
	loot := `{"sources": [{"id": "vendor:15192", "kind": "vendor", "name": "Anachronos", "npc_id": 15192, "items": [1]}], "quests": {}}`
	idx, _ := gatedIndex(t, loot, `[{"key": "aq", "phase": "launch", "npcs": [{"id": 15192, "name": "Anachronos"}]}]`)
	if got := opensOf(t, idx, 1); got[0] != "" {
		t.Errorf("a gate dated at launch closed the vendor: %v", got)
	}
}

func TestAContentGateKeepsAStricterComputedOpens(t *testing.T) {
	loot := `{"sources": [{"id": "world:lord-kazzak", "kind": "world", "name": "Lord Kazzak", "items": [1], "opens": "later"}], "quests": {}}`
	idx, _ := gatedIndex(t, loot, `[{"key": "wb", "phase": "raids-1", "npcs": [{"id": 12397, "name": "Lord Kazzak"}]}]`)
	if got := opensOf(t, idx, 1); got[0] != "later" {
		t.Errorf("Lord Kazzak opens %v, want the stricter later kept", got)
	}
}

// Commit 3592e40a5 deleted the Anachronos gate when the vendor dropped
// out of one fetch, and it came back ungated. A missing id is now only
// reported; the gate stays and closes the source the night it returns.
func TestAContentGateWhoseSourceIsMissingIsKept(t *testing.T) {
	gates := `[{"key": "aq", "phase": "later", "npcs": [{"id": 15192, "name": "Anachronos"}]}]`
	_, missing := gatedIndex(t, `{"sources": [{"id": "world:aged-kodo", "kind": "world", "name": "Aged Kodo", "items": [3]}], "quests": {}}`, gates)
	if len(missing) != 1 || !strings.Contains(missing[0], "npc 15192 Anachronos") {
		t.Fatalf("missing = %v, want the Anachronos gate reported", missing)
	}
	idx, missing := gatedIndex(t, `{"sources": [{"id": "vendor:15192", "kind": "vendor", "name": "Anachronos", "npc_id": 15192, "items": [234032]}], "quests": {}}`, gates)
	if len(missing) != 0 {
		t.Errorf("missing = %v after the vendor returned, want none", missing)
	}
	if _, ok := sourceFor(234032, 60, "alliance", "", idx); ok {
		t.Error("the returned Anachronos vendor is ungated")
	}
}

// TestNoBannedItemIsObtainableAtLaunch is the launch guard on the
// ranker's own input: over the active build's committed loot.json with
// the curated gates applied, no item launch-banned.json bans while its
// gate is closed has a source a level-60 character of either faction can
// use. The ranker picks only from such sources, so no banned item can be
// a pick or alternative in any spec.
func TestNoBannedItemIsObtainableAtLaunch(t *testing.T) {
	idx := activeGatedLootIndex(t)
	for id, gate := range bannedWhileClosed(t) {
		for _, faction := range []string{"alliance", "horde"} {
			if s, ok := sourceFor(id, 60, faction, faction, idx); ok {
				t.Errorf("item %d (gate %s) is obtainable at 60 for %s via %s %q", id, gate, faction, s.Kind, s.Label)
			}
		}
	}
}

func activeGatedLootIndex(t *testing.T) lootIndex {
	t.Helper()
	buildDir := activeBuildDir(t)
	idx, _, err := loadGatedLootIndex(publishedRepoRoot, buildDir, nil, io.Discard)
	if err != nil {
		t.Fatalf("loadGatedLootIndex(%s): %v", buildDir, err)
	}
	return idx
}
