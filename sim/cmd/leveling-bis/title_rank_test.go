package main

import (
	"path/filepath"
	"testing"
)

// Round 8, owner ruling 2026-11-01: Forever's PvP rewards need no rank, so
// the pipeline relabels the Era-titled ones as unranked "PvP reward"
// sources and keeps the old rank in title_rank. Default eligibility must not
// move until the owner rules: a source carrying title_rank above pvpRankCap
// stays out of the default picks, its vendor row inherits the same gate, and
// the label names a PvP reward rather than a rank.
func TestAnUnrankedPvpSourceCarryingATitleRankKeepsTheRankCap(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "pvp:reward:horde", "kind": "pvp", "name": "PvP reward", "faction": "horde", "title_rank": 12, "items": [3001]},
			{"id": "vendor:12792", "kind": "vendor", "name": "Lady Palanseer", "npc_id": 12792, "faction": "horde", "items": [3001]},
			{"id": "pvp:reward:horde-low", "kind": "pvp", "name": "PvP reward", "faction": "horde", "title_rank": 8, "items": [3002]}
		],
		"quests": {}
	}`
	if err := writeFile(t, filepath.Join(dir, "loot.json"), lootJSON); err != nil {
		t.Fatal(err)
	}
	idx, _, err := loadLootIndex(dir, nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	if len(idx[3001]) != 2 {
		t.Fatalf("idx[3001] = %+v, want the pvp and the vendor source", idx[3001])
	}
	for _, s := range idx[3001] {
		if !pvpRankExceedsCap(s) {
			t.Errorf("%s source %+v: pvpRankExceedsCap = false, want true (title_rank 12 is above the cap)", s.Kind, s)
		}
	}
	pool := excludeAbovePvpRankCap([]scored{
		{candidate: candidate{ID: 3001}, Source: idx[3001][0], HasSource: true},
		{candidate: candidate{ID: 3003}, Source: itemSource{Kind: "quest"}, HasSource: true},
	})
	if len(pool) != 1 || pool[0].ID != 3003 {
		t.Errorf("excludeAbovePvpRankCap kept %+v, want only the quest item", pool)
	}
	for _, s := range idx[3002] {
		if pvpRankExceedsCap(s) {
			t.Errorf("idx[3002] %+v: title_rank 8 is within the cap", s)
		}
	}
	if got, want := idx[3002][0].Label, "PvP reward · Horde"; got != want {
		t.Errorf("label = %q, want %q", got, want)
	}
}
