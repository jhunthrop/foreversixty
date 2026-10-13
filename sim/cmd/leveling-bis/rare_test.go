package main

import (
	"path/filepath"
	"testing"
)

// Task L1.7 (the persona passes' pick audit): loot.json marks a rare
// spawn's own rows `rare_spawn` (pipeline.loot.npc_rank: a rare-ranked
// creature with a pooled or open-world spawn, from the pinned classic-db
// dump). The ranker names such a source "(rare)", prefers any other source
// the item has, and never makes a rare spawn's drop the pick of a slot that
// has a candidate from another source at the band. A trash drop several
// creatures of one instance share, none with a rate, is named by its
// instance alone ("Gnomeregan"), never by whichever creature loot.json
// happens to list first.

func loadTestLootIndex(t *testing.T, lootJSON string) lootIndex {
	t.Helper()
	dir := t.TempDir()
	if err := writeFile(t, filepath.Join(dir, "loot.json"), lootJSON); err != nil {
		t.Fatal(err)
	}
	idx, _, err := loadLootIndex(dir, nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	return idx
}

func TestLoadLootIndexNamesARareSpawnsSourceRare(t *testing.T) {
	idx := loadTestLootIndex(t, `{
		"sources": [
			{"id": "dungeon:scarlet-monastery", "kind": "dungeon", "name": "Scarlet Monastery", "bosses": [
				{"name": "Ironspine", "npc_id": 6489, "rank": 2, "rare_spawn": true, "items": [7686]},
				{"name": "Herod", "npc_id": 3975, "rank": 1, "items": [7718]}
			]},
			{"id": "world:swiftmane", "kind": "world", "name": "Swiftmane", "rare_spawn": true, "items": [285330]},
			{"id": "world:bluegill-muckdweller", "kind": "world", "name": "Bluegill Muckdweller", "items": [3022]}
		],
		"quests": {}
	}`)
	cases := []struct {
		id    int
		label string
		rare  bool
	}{
		{7686, "Scarlet Monastery: Ironspine (rare)", true},
		{7718, "Scarlet Monastery: Herod", false},
		{285330, "Swiftmane (rare)", true},
		{3022, "Bluegill Muckdweller", false},
	}
	for _, c := range cases {
		srcs := idx[c.id]
		if len(srcs) != 1 || srcs[0].Label != c.label || srcs[0].RareSpawn != c.rare {
			t.Errorf("idx[%d] = %+v, want one source %q with RareSpawn %v", c.id, srcs, c.label, c.rare)
		}
	}
}

func TestLoadLootIndexNamesAMultiCreatureTrashDropWithNoRateByItsInstance(t *testing.T) {
	idx := loadTestLootIndex(t, `{
		"sources": [
			{"id": "dungeon:gnomeregan", "kind": "dungeon", "name": "Gnomeregan", "trash": [9509, 9510, 9511], "bosses": [
				{"name": "Caverndeep Burrower", "npc_id": 6206, "items": [9509, 9510, 9511], "item_chances": {"9510": 1.5}},
				{"name": "Caverndeep Ambusher", "npc_id": 6207, "items": [9509, 9510]},
				{"name": "Mekgineer Thermaplugg", "npc_id": 7800, "items": [9509, 9999]}
			]},
			{"id": "dungeon:wailing-caverns", "kind": "dungeon", "name": "Wailing Caverns", "trash": [10413], "bosses": [
				{"name": "Druid of the Fang", "npc_id": 3840, "items": [10413], "item_chances": {"10413": 2.34}}
			]}
		],
		"quests": {}
	}`)
	// Petrolspill Leggings: trash, three creatures, no rate anywhere.
	if srcs := idx[9509]; len(srcs) != 1 || srcs[0].Label != "Gnomeregan" || srcs[0].Kind != "dungeon" || srcs[0].Chance != 0 {
		t.Errorf("idx[9509] = %+v, want the one instance-wide source \"Gnomeregan\"", srcs)
	}
	// A trash drop one creature states a rate for keeps its creatures.
	if src, ok := sourceFor(9510, 30, "alliance", "", idx); !ok || src.Label != "Gnomeregan: Caverndeep Burrower" || src.Chance != 1.5 {
		t.Errorf("sourceFor(9510) = %+v, %v, want the Burrower at 1.5%%", src, ok)
	}
	// A trash drop only one creature names keeps that creature.
	if srcs := idx[9511]; len(srcs) != 1 || srcs[0].Label != "Gnomeregan: Caverndeep Burrower" {
		t.Errorf("idx[9511] = %+v, want the one creature that drops it", srcs)
	}
	// A boss's own item that is not trash keeps its boss.
	if srcs := idx[9999]; len(srcs) != 1 || srcs[0].Label != "Gnomeregan: Mekgineer Thermaplugg" {
		t.Errorf("idx[9999] = %+v, want Thermaplugg", srcs)
	}
	// Gloves of the Fang: trash, one creature, with its rate.
	if srcs := idx[10413]; len(srcs) != 1 || srcs[0].Label != "Wailing Caverns: Druid of the Fang" || srcs[0].Chance != 2.34 {
		t.Errorf("idx[10413] = %+v, want the Druid of the Fang at 2.34%%", srcs)
	}
}

func TestSourceForPrefersAnySourceThatIsNotARareSpawn(t *testing.T) {
	idx := lootIndex{
		1: {
			{Kind: "dungeon", Label: "Scarlet Monastery: Ironspine (rare)", Chance: 20, RareSpawn: true},
			{Kind: "crafted", Label: "Leatherworking"},
		},
		2: {{Kind: "dungeon", Label: "Scarlet Monastery: Ironspine (rare)", RareSpawn: true}},
	}
	if src, ok := sourceFor(1, 30, "alliance", "", idx); !ok || src.Label != "Leatherworking" {
		t.Errorf("sourceFor(1) = %+v, %v, want the crafted source over the rare spawn", src, ok)
	}
	if src, ok := sourceFor(2, 30, "alliance", "", idx); !ok || !src.RareSpawn {
		t.Errorf("sourceFor(2) = %+v, %v, want the rare spawn, the item's only source", src, ok)
	}
}

func TestBuildBandPoolKeepsARareSpawnsDropOnlyWhereNothingElseFillsTheSlot(t *testing.T) {
	finger := []string{"finger1", "finger2"}
	items := []candidate{
		{ID: 7686, Name: "Ironspine's Eye", RequiredLevel: 30, EffectiveRequiredLevel: 30, Stats: map[string]float64{"agility": 9}, Slots: finger},
		{ID: 1491, Name: "Ring of Precision", RequiredLevel: 25, EffectiveRequiredLevel: 25, Stats: map[string]float64{"agility": 4}, Slots: finger},
		{ID: 9001, Name: "Rare Helm", RequiredLevel: 25, EffectiveRequiredLevel: 25, Stats: map[string]float64{"agility": 5}, Slots: []string{"head"}},
	}
	idx := lootIndex{
		7686: {{Kind: "dungeon", Label: "Scarlet Monastery: Ironspine (rare)", RareSpawn: true}},
		1491: {{Kind: "dungeon", Label: "Blackfathom Deeps: Fallenroot Shadowstalker"}},
		9001: {{Kind: "world", Label: "Lupos (rare)", RareSpawn: true}},
	}
	pool := buildBandPool(items, idx, "hunter", 30, "alliance", map[string]float64{"agility": 1}, 0, false)
	got := map[int]bool{}
	for _, s := range pool.Scored {
		got[s.ID] = true
	}
	if got[7686] || !got[1491] || !got[9001] || len(pool.Scored) != 2 {
		t.Errorf("pool.Scored ids = %v, want 1491 and 9001: the rare ring gives way to the other ring, the rare helm has no rival", got)
	}
	if pool.RareSpawnSkipped != 1 {
		t.Errorf("pool.RareSpawnSkipped = %d, want 1", pool.RareSpawnSkipped)
	}
}

// The nightly's published-file check (published_test.go) compares every
// published source with sourceFor's own. 1.60.1.70009's files were ranked
// before task L1.7 and are never re-ranked, so for that build alone the
// check accepts exactly L1.7's two relabels - a creature's drop now named
// by its instance alone, a rare spawn's now ending "(rare)" - and nothing
// else. The active build must match sourceFor exactly.
func TestThePublishedSourceCheckAcceptsOnlyTaskL17sRelabelsInAnOlderBuild(t *testing.T) {
	cases := []struct {
		published, want string
		active          bool
		ok              bool
	}{
		{"Gnomeregan", "Gnomeregan", true, true},
		{"Gnomeregan: Caverndeep Burrower", "Gnomeregan", true, false},
		{"Gnomeregan: Caverndeep Burrower", "Gnomeregan", false, true},
		{"Scarlet Monastery: Ironspine", "Scarlet Monastery: Ironspine (rare)", false, true},
		{"Scarlet Monastery: Ironspine", "Scarlet Monastery: Ironspine (rare)", true, false},
		{"Gnomeregan: Caverndeep Burrower", "Uldaman", false, false},
		{"Gnomeregan", "Gnomeregan: Caverndeep Burrower", false, false},
		{"Leatherworking", "Tailoring", false, false},
	}
	for _, c := range cases {
		msg := sourceLabelError("item", 1, "Thing", c.published, c.want, c.active)
		if (msg == "") != c.ok {
			t.Errorf("sourceLabelError(%q, %q, active=%v) = %q, want ok=%v", c.published, c.want, c.active, msg, c.ok)
		}
	}
}
