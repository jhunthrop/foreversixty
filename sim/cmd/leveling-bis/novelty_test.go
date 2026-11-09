package main

import (
	"os"
	"path/filepath"
	"testing"
)

// Feedback 2026-11-09: a fishing catch (Rockhide Strongfish) was a
// level-40 Protection main hand.
func TestAFishingCatchIsNeverASource(t *testing.T) {
	dir := t.TempDir()
	loot := `{"sources": [
		{"id": "world:fishing", "kind": "world", "name": "Fishing", "items": [19808]},
		{"id": "world:aged-kodo", "kind": "world", "name": "Aged Kodo", "items": [1]}
	], "quests": {}}`
	if err := os.WriteFile(filepath.Join(dir, "loot.json"), []byte(loot), 0o644); err != nil {
		t.Fatal(err)
	}
	idx, _, err := loadLootIndex(dir, nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	if srcs := idx[19808]; len(srcs) != 0 {
		t.Errorf("Rockhide Strongfish has sources %v, want none from fishing", srcs)
	}
	if srcs := idx[1]; len(srcs) != 1 {
		t.Errorf("an ordinary world drop has %d sources, want 1", len(srcs))
	}
}

func TestANoveltyItemIsNeverEligible(t *testing.T) {
	hat := candidate{ID: 19972, Name: "Lucky Fishing Hat", ClassID: armorClassID, SubclassID: armorSubclothID}
	if eligible(hat, "paladin", 20, "alliance") {
		t.Error("Lucky Fishing Hat is eligible, want it denied as a fishing novelty")
	}
	other := candidate{ID: 1, ClassID: armorClassID, SubclassID: armorSubclothID}
	if !eligible(other, "paladin", 20, "alliance") {
		t.Error("an ordinary cloth hat is not eligible")
	}
}
