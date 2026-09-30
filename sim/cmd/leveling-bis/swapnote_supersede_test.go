package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// TestRunSpecNeverCitesASupersededItemInSwapNote is this lane's brief
// (bis-ranker-integrity-14), item 2: the caster-spec player review
// (day3/player-review-33/casters.md, 20 instances across
// warlock-affliction/demonology/destruction and priest-shadow at band
// 60) found a published swap_note naming a superseded vanilla PvP item
// id/name ("Field Marshal's Coronal (id 17578)") while the row's own
// alternatives correctly named the current, re-itemised successor
// (id 231584, same display name, different stats).
//
// Investigation (this lane): loadCandidates (data.go) already excludes
// any row either the flat or per-class file marks superseded_by
// (TestLoadCandidatesExcludesSupersededItems*, above) - confirmed by
// hand against the live build (data/builds/1.60.1.70009), id 17578 does
// carry superseded_by: 231584 on both its flat and per-class rows, and
// re-running sim/cmd/leveling-bis fresh against that exact data, at
// both the current HEAD and at the exact commit the stale swap_note
// was published from (72fc081e), never reproduces it - every fresh run
// omits the swap_note for that row entirely rather than naming either
// id. That means the published artifact is stale relative to its own
// commit's source and data (a nightly-pipeline staleness question, not
// a bug this lane's code can fix), but nothing end-to-end already
// proved the INVARIANT the bug report actually cares about: once an
// item is excluded as superseded, nothing downstream (pick()'s
// score-based runner-up, verifyBand's swap comparison, buildReport's
// SwapNote text) can ever resurrect its id/name, even when a same-named
// successor is its own runner-up. This test builds a synthetic
// superseded pair (a legacy head item 1013, superseded_by 1012, the
// same shape as the live Field Marshal's Coronal pair) through the
// real file-loading path and the real pick/verify/report pipeline (a
// fake engine only), so a future regression that reintroduces the
// legacy id anywhere in this path (not just at loadCandidates) fails
// loudly.
func TestRunSpecNeverCitesASupersededItemInSwapNote(t *testing.T) {
	repoRoot := t.TempDir()
	if err := os.CopyFS(repoRoot, os.DirFS(repoRootFixture)); err != nil {
		t.Fatalf("copying %s into %s: %v", repoRootFixture, repoRoot, err)
	}
	buildDir := filepath.Join(repoRoot, "data", "builds", "testbuild")

	// items.json: keep 1001/1002 (the shared fixture's existing items,
	// other tests' expectations depend on their shape unchanged
	// elsewhere), add 1012 (the current, re-itemised head item) and
	// 1013 (the legacy item 1012 superseded - same display name, a
	// touch less agility, exactly the "same name, different id/stats"
	// shape the live Field Marshal's Coronal pair has).
	if err := writeFile(t, filepath.Join(buildDir, "items.json"), `[
		{"id": 1001, "name": "Test Helm", "quality": 3, "item_level": 20, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 1, "suffixes": [], "faction_restriction": ""},
		{"id": 1002, "name": "Test Bow", "quality": 3, "item_level": 25, "required_level": 15, "class_id": 2, "subclass_id": 2, "inventory_type": 15, "suffixes": [], "faction_restriction": "horde"},
		{"id": 1012, "name": "Marshal's Test Helm", "quality": 4, "item_level": 74, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 1, "suffixes": [], "faction_restriction": ""},
		{"id": 1013, "name": "Marshal's Test Helm", "quality": 4, "item_level": 74, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 1, "suffixes": [], "faction_restriction": "", "superseded_by": 1012}
	]`); err != nil {
		t.Fatal(err)
	}
	if err := writeFile(t, filepath.Join(buildDir, "items", "hunter.json"), `{
		"build": "testbuild", "class_slug": "hunter",
		"items": [
			{"id": 1001, "name": "Test Helm", "slot": "head", "quality": 3, "required_level": 10, "item_level": 20, "armor": 40, "stats": {"agility": 5}, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": false},
			{"id": 1002, "name": "Test Bow", "slot": "ranged", "quality": 3, "required_level": 15, "item_level": 25, "armor": 0, "stats": {"agility": 2}, "damage_min": 10, "damage_max": 20, "speed": 2.8, "dps": 5.3, "two_hand": false, "effect_text": "", "set_id": null, "unique": false, "weapon_type": "bow"},
			{"id": 1003, "name": "Ghost Item", "slot": "waist", "quality": 1, "required_level": 5, "item_level": 5, "armor": 1, "stats": {}, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": false},
			{"id": 1012, "name": "Marshal's Test Helm", "slot": "head", "quality": 4, "required_level": 10, "item_level": 74, "armor": 60, "stats": {"agility": 4}, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": true},
			{"id": 1013, "name": "Marshal's Test Helm", "slot": "head", "quality": 4, "required_level": 10, "item_level": 74, "armor": 60, "stats": {"agility": 4}, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": true, "superseded_by": 1012}
		]
	}`); err != nil {
		t.Fatal(err)
	}
	// loot.json: source every head item (1001, 1012 - 1013 never needs
	// one, loadCandidates excludes it before sourcing is ever asked)
	// plus keep the existing fixture's other sources so nothing else
	// this reporoot's other tests exercise regresses.
	if err := writeFile(t, filepath.Join(buildDir, "loot.json"), `{
		"sources": [
			{"id": "quest-1", "kind": "quest", "name": "A Test Quest", "items": [1001]},
			{"id": "dungeon-1", "kind": "dungeon", "name": "A Test Dungeon", "bosses": [{"name": "Test Boss", "items": [1002]}]},
			{"id": "dungeon:ragefire-chasm", "kind": "dungeon", "name": "Ragefire Chasm", "bosses": [{"name": "Taragaman the Hungerer", "items": [1004]}]},
			{"id": "rep-1", "kind": "rep", "name": "A Test Faction", "faction_id": 1, "standing": "revered", "items": [1005]},
			{"id": "vendor-1", "kind": "vendor", "name": "A Test Quartermaster", "items": [1005]},
			{"id": "dungeon-marshal", "kind": "dungeon", "name": "A Test Battleground", "bosses": [{"name": "Test Field Marshal", "items": [1012]}]}
		],
		"quests": {
			"1001": [{"quest_id": 1, "name": "A Test Quest", "faction": "both", "min_level": 5, "level": 7, "level_source": "wowhead"}]
		}
	}`); err != nil {
		t.Fatal(err)
	}

	// fakeEngine: every gear fingerprint measures the same DPS, so
	// verify.go's own real-sim comparison between the score-based pick
	// (1001) and its runner-up (1012, the current/successor item) is a
	// clean "kept the pick" - the same shape the live bug's own
	// swap_note text used, the one most likely to carry a stale
	// id/name into published text.
	fake := &fakeEngine{
		DefaultDPS: 500,
		WeightsResult: map[string]api.StatWeight{
			"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0},
			"agility":             {Stat: "agility", Weight: 1.8},
		},
		ReferenceDPSPerPoint: 0.5,
	}

	outDir := t.TempDir()
	if err := runSpec(fake, repoRoot, buildDir, "testbuild", outDir, "hunter-marksmanship", []int{20}, 5); err != nil {
		t.Fatalf("runSpec: %v", err)
	}

	raw, err := os.ReadFile(filepath.Join(outDir, "hunter-marksmanship.json"))
	if err != nil {
		t.Fatalf("reading report: %v", err)
	}
	var report specReport
	if err := json.Unmarshal(raw, &report); err != nil {
		t.Fatalf("decoding report: %v", err)
	}

	// The legacy id/name must never surface anywhere in the published
	// report - not in a pick, not in an alternative, not in a swap_note
	// - the exact invariant the live bug broke (a swap_note is the one
	// place that failed; this test checks the whole document so a
	// regression anywhere in the pipeline, not only in the exact
	// reproduction shape, is caught).
	if strings.Contains(string(raw), "1013") {
		t.Fatalf("published report contains the superseded item's id (1013) somewhere: %s", string(raw))
	}

	var headRow *slotRow
	for _, b := range report.Bands {
		if b.Band != 20 {
			continue
		}
		for i := range b.Slots {
			if b.Slots[i].Slot == "head" {
				headRow = &b.Slots[i]
			}
		}
	}
	if headRow == nil {
		t.Fatal("no head row published at band 20")
	}
	// The pick must be the surviving, higher-agility current item
	// (1001 outscores 1012 under these weights; what matters is that
	// 1013 is never the pick and never named anywhere on the row).
	if headRow.ItemID != 1001 {
		t.Fatalf("head pick = %+v, want 1001 (the real, higher-scored head candidate)", headRow)
	}
	if headRow.SwapNote != "" && strings.Contains(headRow.SwapNote, "1013") {
		t.Fatalf("head row.SwapNote = %q, must never cite the superseded id 1013", headRow.SwapNote)
	}
	// The runner-up (1012, the current/successor item) must be the one
	// named in alternatives, with its own real id/name - not silently
	// dropped, and never replaced by its superseded predecessor's id.
	foundRunnerUp := false
	for _, alt := range headRow.Alternatives {
		if alt.ItemID == 1013 {
			t.Fatalf("head row.Alternatives = %+v, must never list the superseded id 1013", headRow.Alternatives)
		}
		if alt.ItemID == 1012 {
			foundRunnerUp = true
			if alt.ItemName != "Marshal's Test Helm" {
				t.Errorf("alternative for id 1012 has name %q, want the current row's own name", alt.ItemName)
			}
		}
	}
	if !foundRunnerUp {
		t.Fatalf("head row.Alternatives = %+v, want the current item (1012) named as the runner-up", headRow.Alternatives)
	}
}
