package main

import (
	"encoding/json"
	"path/filepath"
	"testing"
)

func TestCollectSpellIDsFindsBareAndRankedShapes(t *testing.T) {
	var doc any = map[string]any{
		"prepullActions": []any{
			map[string]any{"action": map[string]any{"castSpell": map[string]any{
				"spellId": map[string]any{"spellId": float64(11269), "rank": float64(6)},
			}}},
		},
		"priorityList": []any{
			map[string]any{"action": map[string]any{"castSpell": map[string]any{
				"spellId": map[string]any{"spellId": float64(14278)},
			}}},
			map[string]any{"condition": map[string]any{"cmp": map[string]any{
				"lhs": map[string]any{"currentComboPoints": map[string]any{}},
			}}},
		},
	}
	got := map[int]bool{}
	collectSpellIDs(doc, got)
	want := map[int]bool{11269: true, 14278: true}
	if len(got) != len(want) {
		t.Fatalf("collectSpellIDs = %v, want %v", got, want)
	}
	for id := range want {
		if !got[id] {
			t.Errorf("collectSpellIDs missing id %d: %v", id, got)
		}
	}
}

func writeAPLFixture(t *testing.T, dir, spec string, spellIDs ...int) {
	t.Helper()
	actions := make([]any, len(spellIDs))
	for i, id := range spellIDs {
		actions[i] = map[string]any{"action": map[string]any{"castSpell": map[string]any{"spellId": map[string]any{"spellId": id}}}}
	}
	doc := map[string]any{
		"spec":  spec,
		"state": "written",
		"rotation": map[string]any{
			"type":         "TypeAPL",
			"priorityList": actions,
		},
	}
	b, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "data", "curated", "apl", spec+".json")
	if err := writeFile(t, path, string(b)); err != nil {
		t.Fatal(err)
	}
}

// This lane's brief, item 5: rogue-subtlety's own rotation casts Ambush
// (spell id 11269 at rank 6, matching data/curated/apl/rogue-
// subtlety.json's real prepull action) - a dagger-only Classic ability
// - so its rotation must be found to require a dagger.
func TestAPLRotationRequiresDaggerTrueForAmbush(t *testing.T) {
	dir := t.TempDir()
	writeAPLFixture(t, dir, "rogue-subtlety", 11269, 14278)
	got, err := aplRotationRequiresDagger(dir, "rogue-subtlety")
	if err != nil {
		t.Fatalf("aplRotationRequiresDagger: %v", err)
	}
	if !got {
		t.Error("aplRotationRequiresDagger = false, want true: this rotation casts Ambush (id 11269)")
	}
}

// A rotation that never names Backstab or Ambush at any rank (Ghostly
// Strike, id 14278, is neither) must not be flagged - this is the
// negative control every other rogue spec, and every other class
// entirely, relies on to stay unrestricted.
func TestAPLRotationRequiresDaggerFalseWithNeitherAbility(t *testing.T) {
	dir := t.TempDir()
	writeAPLFixture(t, dir, "rogue-combat", 14278, 1752)
	got, err := aplRotationRequiresDagger(dir, "rogue-combat")
	if err != nil {
		t.Fatalf("aplRotationRequiresDagger: %v", err)
	}
	if got {
		t.Error("aplRotationRequiresDagger = true, want false: this rotation names neither Backstab nor Ambush")
	}
}

// A spec with no apl file yet (not curated - writtenSpecs' own "skipped,
// not an error" reading) must not fail runSpec over a requirement that
// does not apply to it yet.
func TestAPLRotationRequiresDaggerMissingFileIsFalseNotError(t *testing.T) {
	dir := t.TempDir()
	got, err := aplRotationRequiresDagger(dir, "nonexistent-spec")
	if err != nil {
		t.Fatalf("aplRotationRequiresDagger for a missing apl file: want no error, got %v", err)
	}
	if got {
		t.Error("aplRotationRequiresDagger for a missing apl file: want false")
	}
}

// restrictToDaggers's own contract, this lane's brief item 5: Ardent
// Custodian (a real Mace1H, subclass 4) is dropped; a real dagger and a
// non-weapon (should never appear in main_hand/off_hand, but left alone
// rather than assumed impossible) both pass through.
func TestRestrictToDaggersDropsNonDaggerWeaponsOnly(t *testing.T) {
	mace := scored{candidate: candidate{ID: 868, Name: "Ardent Custodian", ClassID: itemClassWeapon, SubclassID: 4}}
	dagger := scored{candidate: candidate{ID: 1, Name: "A Dagger", ClassID: itemClassWeapon, SubclassID: daggerSubclassID}}
	nonWeapon := scored{candidate: candidate{ID: 2, Name: "Not A Weapon", ClassID: armorClassID}}
	got := restrictToDaggers([]scored{mace, dagger, nonWeapon})
	if len(got) != 2 {
		t.Fatalf("restrictToDaggers = %+v, want 2 entries (the mace dropped)", got)
	}
	for _, c := range got {
		if c.ID == 868 {
			t.Fatalf("restrictToDaggers kept the mace (id 868): %+v", got)
		}
	}
}
