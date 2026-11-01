package main

import (
	"bytes"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// Task L1.7 fix round 3: the ranker is run-to-run deterministic. score()
// summed over a Go map, whose iteration order is random, so two items with
// identical stats could score an ulp apart and the tie-break went to the
// noise instead of the item id.

// TestScoreIsBitIdenticalWhateverTheStatMapOrder scores one candidate 1,000
// times, each from a fresh copy of its stat map built in a shuffled insertion
// order, and wants one bit-identical result.
func TestScoreIsBitIdenticalWhateverTheStatMapOrder(t *testing.T) {
	stats := map[string]float64{
		"agility": 7, "strength": 3, "stamina": 11, "intellect": 5, "spirit": 9,
		"crit": 1, "hit": 2, "attack_power": 13, "ranged_attack_power": 17,
		"spell_power": 19, "healing_power": 23, "mp5": 4, "defense": 6,
	}
	weights := map[string]float64{
		"agility": 0.1, "strength": 0.7, "stamina": 0.3, "intellect": 0.2, "spirit": 0.6,
		"crit": 13.3, "hit": 11.1, "attack_power": 0.07, "ranged_attack_power": 1,
		"spell_power": 0.9, "healing_power": 0.11, "mp5": 1.3, "defense": 0.17,
		armorWeightStat: 0.013,
	}
	keys := make([]string, 0, len(stats))
	for k := range stats {
		keys = append(keys, k)
	}
	rng := rand.New(rand.NewSource(1))
	seen := map[uint64]int{}
	for i := 0; i < 1000; i++ {
		fresh := make(map[string]float64, len(stats))
		for _, j := range rng.Perm(len(keys)) {
			fresh[keys[j]] = stats[keys[j]]
		}
		c := candidate{Stats: fresh, Armor: 37, DPS: 12.3}
		seen[math.Float64bits(score(c, "main_hand", weights, 0))]++
	}
	if len(seen) != 1 {
		got := map[float64]int{}
		for bits, n := range seen {
			got[math.Float64frombits(bits)] = n
		}
		t.Fatalf("1,000 scores of one candidate gave %d distinct values: %v", len(seen), got)
	}
}

// TestStoredScoresTieWhenTheSumsAreEqualUpToRounding: two stat blocks whose
// weighted sums are equal in exact arithmetic but not in floating point
// (0.1+0.2 against 0.3) store one score through buildBandPool, so the id
// tie-break decides between them, and an alternative's dps_delta is 0.
func TestStoredScoresTieWhenTheSumsAreEqualUpToRounding(t *testing.T) {
	weights := map[string]float64{"agility": 0.1, "crit": 0.2, "hit": 0.3}
	a := candidate{ID: 2, Name: "A", Slots: []string{"wrist"}, Stats: map[string]float64{"agility": 1, "crit": 1}}
	b := candidate{ID: 1, Name: "B", Slots: []string{"wrist"}, Stats: map[string]float64{"hit": 1}}
	if score(a, "wrist", weights, 0) == score(b, "wrist", weights, 0) {
		t.Fatal("the raw sums agree; the fixture no longer shows the float gap")
	}
	if roundScore(score(a, "wrist", weights, 0)) != roundScore(score(b, "wrist", weights, 0)) {
		t.Fatalf("rounded scores differ: %v vs %v", roundScore(score(a, "wrist", weights, 0)), roundScore(score(b, "wrist", weights, 0)))
	}
	if got := roundScore(-1e-15); math.Signbit(got) || got != 0 {
		t.Fatalf("roundScore(-1e-15) = %v, want a positive 0", got)
	}

	idx := lootIndex{
		1: {{Kind: "quest", Label: "A Quest"}},
		2: {{Kind: "quest", Label: "A Quest"}},
	}
	pool := buildBandPool([]candidate{a, b}, idx, "hunter", 20, "alliance", weights, 0)
	if len(pool.Scored) != 2 {
		t.Fatalf("scored %d candidates, want 2", len(pool.Scored))
	}
	if pool.Scored[0].Score != pool.Scored[1].Score {
		t.Fatalf("stored scores differ: %v vs %v", pool.Scored[0].Score, pool.Scored[1].Score)
	}
	list := append([]scored(nil), pool.Scored...)
	sortCandidates(list)
	if list[0].ID != 1 {
		t.Fatalf("wrist order = [%d %d], want the lower id first on a tie", list[0].ID, list[1].ID)
	}
}

// determinismRepo is the shared reporoot fixture with two wrist items whose
// stat blocks are identical (the shaman-restoration band 50 wrist shape,
// 252543/252573) and a head with several weighted stats, so a score summed
// in map order differs by an ulp from run to run.
func determinismRepo(t *testing.T) (repoRoot, buildDir string) {
	t.Helper()
	repoRoot = t.TempDir()
	if err := os.CopyFS(repoRoot, os.DirFS(repoRootFixture)); err != nil {
		t.Fatalf("copying %s into %s: %v", repoRootFixture, repoRoot, err)
	}
	buildDir = filepath.Join(repoRoot, "data", "builds", "testbuild")
	if err := writeFile(t, filepath.Join(buildDir, "items.json"), `[
		{"id": 1001, "name": "Test Helm", "quality": 3, "item_level": 20, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 1, "suffixes": [], "faction_restriction": ""},
		{"id": 1002, "name": "Test Bow", "quality": 3, "item_level": 25, "required_level": 15, "class_id": 2, "subclass_id": 2, "inventory_type": 15, "suffixes": [], "faction_restriction": "horde"},
		{"id": 1031, "name": "Twin Bracers A", "quality": 3, "item_level": 20, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 9, "suffixes": [], "faction_restriction": ""},
		{"id": 1032, "name": "Twin Bracers B", "quality": 3, "item_level": 20, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 9, "suffixes": [], "faction_restriction": ""},
		{"id": 1033, "name": "Busy Helm", "quality": 3, "item_level": 20, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 1, "suffixes": [], "faction_restriction": ""}
	]`); err != nil {
		t.Fatal(err)
	}
	const twinStats = `{"agility": 7, "attack_power": 13, "ranged_attack_power": 11, "crit": 3}`
	if err := writeFile(t, filepath.Join(buildDir, "items", "hunter.json"), `{
		"build": "testbuild", "class_slug": "hunter",
		"items": [
			{"id": 1001, "name": "Test Helm", "slot": "head", "quality": 3, "required_level": 10, "item_level": 20, "armor": 40, "stats": {"agility": 5}, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": false},
			{"id": 1002, "name": "Test Bow", "slot": "ranged", "quality": 3, "required_level": 15, "item_level": 25, "armor": 0, "stats": {"agility": 2}, "damage_min": 10, "damage_max": 20, "speed": 2.8, "dps": 5.3, "two_hand": false, "effect_text": "", "set_id": null, "unique": false, "weapon_type": "bow"},
			{"id": 1031, "name": "Twin Bracers A", "slot": "wrist", "quality": 3, "required_level": 10, "item_level": 20, "armor": 20, "stats": `+twinStats+`, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": false},
			{"id": 1032, "name": "Twin Bracers B", "slot": "wrist", "quality": 3, "required_level": 10, "item_level": 20, "armor": 20, "stats": `+twinStats+`, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": false},
			{"id": 1033, "name": "Busy Helm", "slot": "head", "quality": 3, "required_level": 10, "item_level": 20, "armor": 40, "stats": {"agility": 9, "attack_power": 17, "ranged_attack_power": 19, "crit": 5}, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": false}
		]
	}`); err != nil {
		t.Fatal(err)
	}
	if err := writeFile(t, filepath.Join(buildDir, "loot.json"), `{
		"sources": [
			{"id": "quest-1", "kind": "quest", "name": "A Test Quest", "items": [1001]},
			{"id": "dungeon-1", "kind": "dungeon", "name": "A Test Dungeon", "bosses": [{"name": "Test Boss", "items": [1002, 1031, 1032, 1033]}]}
		],
		"quests": {
			"1001": [{"quest_id": 1, "name": "A Test Quest", "faction": "both", "min_level": 5, "level": 7, "level_source": "wowhead"}]
		}
	}`); err != nil {
		t.Fatal(err)
	}
	return repoRoot, buildDir
}

// volatileReportFields are the fields a rerun may change: the timestamp and
// the wall-clock run times.
var volatileReportFields = regexp.MustCompile(`"(generated_at|weights_run_seconds|verify_run_seconds)":\s*("[^"]*"|[-0-9.eE+]+)`)

// TestTwoRunSpecPassesPublishByteIdenticalJSON runs the whole spec over the
// determinism fixture and wants the same JSON from every pass, apart from
// the timestamp and run times, and the twin wrists' tie going to the lower
// id. Two passes are the contract; more make a map-order flip near certain
// to show.
func TestTwoRunSpecPassesPublishByteIdenticalJSON(t *testing.T) {
	repoRoot, buildDir := determinismRepo(t)
	run := func() (published []byte, path string) {
		fake := &fakeEngine{
			DefaultDPS: 500,
			WeightsResult: map[string]api.StatWeight{
				"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0},
				"agility":             {Stat: "agility", Weight: 0.1},
				"crit":                {Stat: "crit", Weight: 0.7},
			},
			ReferenceDPSPerPoint: 0.3,
		}
		outDir := t.TempDir()
		if err := runSpec(fake, repoRoot, buildDir, "testbuild", outDir, "hunter-marksmanship", []int{20}, 5, identityTalentLayout); err != nil {
			t.Fatalf("runSpec: %v", err)
		}
		path = filepath.Join(outDir, "hunter-marksmanship.json")
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		return volatileReportFields.ReplaceAll(raw, []byte(`"$1": null`)), path
	}
	first, path := run()
	for pass := 2; pass <= 12; pass++ {
		if again, _ := run(); !bytes.Equal(first, again) {
			t.Fatalf("pass %d published different JSON than pass 1:\n--- pass 1\n%s\n--- pass %d\n%s", pass, first, pass, again)
		}
	}
	wrists := 0
	for _, b := range readSpecReportForTest(t, path).Bands {
		for _, row := range b.Slots {
			if row.Slot != "wrist" {
				continue
			}
			wrists++
			if row.ItemID != 1031 {
				t.Errorf("band %d %s wrist = %d, want 1031 (the lower id of two identical items)", b.Band, b.Faction, row.ItemID)
			}
		}
	}
	if wrists == 0 {
		t.Fatal("no wrist row published")
	}
}
