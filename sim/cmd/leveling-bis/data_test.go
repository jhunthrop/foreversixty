package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// repoRootFixture is testdata/reporoot: a minimal, hand-built stand-in
// for the site repository's own layout (data/curated/specs.json, its
// apl/ state files, web/src/data/active-build.json, one guide's
// frontmatter, and one build's items.json/items/hunter.json/loot.json)
// - just enough for every data.go loader to exercise its real file, no
// stub.
const repoRootFixture = "testdata/reporoot"

func buildDirFixture() string {
	return filepath.Join(repoRootFixture, "data", "builds", "testbuild")
}

func TestPlannerSlotsExpandsAliases(t *testing.T) {
	if got := plannerSlots("finger"); len(got) != 2 || got[0] != "finger1" || got[1] != "finger2" {
		t.Errorf("plannerSlots(finger) = %v, want [finger1 finger2]", got)
	}
	if got := plannerSlots("trinket"); len(got) != 2 || got[0] != "trinket1" || got[1] != "trinket2" {
		t.Errorf("plannerSlots(trinket) = %v, want [trinket1 trinket2]", got)
	}
	if got := plannerSlots("head"); len(got) != 1 || got[0] != "head" {
		t.Errorf("plannerSlots(head) = %v, want [head] (no alias)", got)
	}
}

func TestLoadFlatItems(t *testing.T) {
	items, err := loadFlatItems(buildDirFixture())
	if err != nil {
		t.Fatalf("loadFlatItems: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("loadFlatItems = %d items, want 2", len(items))
	}
	if items[0].ID != 1001 || items[0].FactionRestriction != "" {
		t.Errorf("items[0] = %+v, want id 1001 no faction restriction", items[0])
	}
	if items[1].FactionRestriction != "horde" {
		t.Errorf("items[1].FactionRestriction = %q, want horde", items[1].FactionRestriction)
	}
}

func TestLoadFlatItemsMissingFile(t *testing.T) {
	if _, err := loadFlatItems(t.TempDir()); err == nil {
		t.Fatal("loadFlatItems on an empty dir: want an error, got nil")
	}
}

func TestLoadClassItems(t *testing.T) {
	f, err := loadClassItems(buildDirFixture(), "hunter")
	if err != nil {
		t.Fatalf("loadClassItems: %v", err)
	}
	if f.ClassSlug != "hunter" || len(f.Items) != 3 {
		t.Fatalf("loadClassItems = %+v, want class hunter with 3 items", f)
	}
}

func TestLoadClassItemsMissingClass(t *testing.T) {
	if _, err := loadClassItems(buildDirFixture(), "nonexistent-class"); err == nil {
		t.Fatal("loadClassItems for a class with no file: want an error, got nil")
	}
}

func TestLoadCandidatesMergesFlatAndClassFiles(t *testing.T) {
	candidates, missing, err := loadCandidates(buildDirFixture(), "hunter")
	if err != nil {
		t.Fatalf("loadCandidates: %v", err)
	}
	// Item 1003 ("Ghost Item") is in items/hunter.json but not
	// items.json - it must be reported as missing, not merged in with
	// zero-value ClassID/SubclassID/FactionRestriction (those decide
	// eligibility, so a silent zero value would be a wrong answer, not
	// a missing one).
	if len(candidates) != 2 {
		t.Fatalf("loadCandidates = %d candidates, want 2 (the ghost item excluded)", len(candidates))
	}
	if len(missing) != 1 || missing[0] == "" {
		t.Fatalf("missing = %v, want exactly one note about item 1003", missing)
	}
	byID := map[int]candidate{}
	for _, c := range candidates {
		byID[c.ID] = c
	}
	helm, ok := byID[1001]
	if !ok {
		t.Fatal("candidate 1001 missing")
	}
	if helm.ClassID != 4 || helm.SubclassID != 2 {
		t.Errorf("helm ClassID/SubclassID = %d/%d, want 4/2 (from the flat file)", helm.ClassID, helm.SubclassID)
	}
	if len(helm.Slots) != 1 || helm.Slots[0] != "head" {
		t.Errorf("helm Slots = %v, want [head]", helm.Slots)
	}
	bow, ok := byID[1002]
	if !ok {
		t.Fatal("candidate 1002 missing")
	}
	if bow.FactionRestriction != "horde" {
		t.Errorf("bow FactionRestriction = %q, want horde (from the flat file)", bow.FactionRestriction)
	}
	if bow.DPS != 5.3 {
		t.Errorf("bow DPS = %v, want 5.3", bow.DPS)
	}
}

func TestLoadLootIndex(t *testing.T) {
	idx, questFloors, err := loadLootIndex(buildDirFixture(), nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	quest, ok := idx[1001]
	if !ok || len(quest) != 1 || quest[0].Kind != "quest" || quest[0].Label != "A Test Quest" {
		t.Errorf("idx[1001] = %+v, want one quest source named %q", quest, "A Test Quest")
	}
	boss, ok := idx[1002]
	if !ok || len(boss) != 1 || boss[0].Kind != "dungeon" || boss[0].Label != "A Test Dungeon: Test Boss" {
		t.Errorf("idx[1002] = %+v, want one dungeon source labelled %q", boss, "A Test Dungeon: Test Boss")
	}
	if _, ok := idx[9999]; ok {
		t.Error("idx[9999] present, want absent (no source names it)")
	}
	// Item 1004's only source is the fixture's "dungeon:ragefire-chasm"
	// row - factionExclusiveDungeons marks that id horde-only, so its
	// itemSource must carry Faction "horde", the same as the real
	// build's Subterranean Cape (14149) does.
	ragefire, ok := idx[1004]
	if !ok || len(ragefire) != 1 || ragefire[0].Side != "horde" {
		t.Errorf("idx[1004] = %+v, want one source with Faction \"horde\"", ragefire)
	}
	// An ordinary dungeon source (not in factionExclusiveDungeons)
	// carries no Faction at all - it must not inherit one by accident.
	if boss[0].Side != "" {
		t.Errorf("idx[1002][0].Side = %q, want empty (not a faction-exclusive source)", boss[0].Side)
	}
	// The fixture's item 1001 (required_level 10) is a quest reward
	// whose quest ("A Test Quest") states min_level 25 in loot.json's
	// quests map - the Polar Leggings shape this lane's brief names.
	if questFloors[1001] != 25 {
		t.Errorf("questFloors[1001] = %d, want 25 (from loot.json's quests map)", questFloors[1001])
	}
	if questFloors[1002] != 0 {
		t.Errorf("questFloors[1002] = %d, want 0 (not a quest reward)", questFloors[1002])
	}
	// This lane's brief, defect 2: item 1005's fixture rows are a rep
	// source (revered) and a quartermaster vendor row for the same
	// item - loadLootIndex's own vendorInheritsRepStandingGate must
	// give the vendor row that rep source's own Standing (and Side),
	// so sourceObtainable's revered gate reaches it too, not just the
	// rep row.
	gated, ok := idx[1005]
	if !ok || len(gated) != 2 {
		t.Fatalf("idx[1005] = %+v, want a rep source and a vendor source", gated)
	}
	for _, s := range gated {
		if s.Kind == "vendor" && s.Standing != "revered" {
			t.Errorf("idx[1005] vendor source = %+v, want Standing \"revered\" inherited from the matching rep source", s)
		}
	}
}

// This lane's brief (bis-ranker-integrity, 2026-09-29), item 3:
// loadLootIndex must carry a source's own "opens" field through to its
// itemSource unchanged, so band.go's sourceObtainable can gate a
// leveling list on it. A self-contained fixture (not the shared
// testdata/reporoot one, which no test may add a new item id to
// without touching every other test's exact counts) with one raid
// source curated "later" and one dungeon source with no opens tag at
// all - the shape data/curated/loot/forever-raid-phases.json's own
// real rows take.
func TestLoadLootIndexCarriesOpensThrough(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "raid:naxxramas", "kind": "raid", "name": "Naxxramas", "opens": "later", "bosses": [{"name": "Kel'Thuzad", "items": [2001]}]},
			{"id": "dungeon-1", "kind": "dungeon", "name": "A Test Dungeon", "bosses": [{"name": "Test Boss", "items": [2002]}]}
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
	raidSrc, ok := idx[2001]
	if !ok || len(raidSrc) != 1 || raidSrc[0].Opens != "later" {
		t.Fatalf("idx[2001] = %+v, want one source with Opens \"later\"", raidSrc)
	}
	dungeonSrc, ok := idx[2002]
	if !ok || len(dungeonSrc) != 1 || dungeonSrc[0].Opens != "" {
		t.Fatalf("idx[2002] = %+v, want one source with Opens empty (not phase-gated)", dungeonSrc)
	}
}

// TestLoadLootIndexVendorInheritsPvpRank is the real shape this lane's
// own dogfood run found regenerating band-60 weapons for item 1: loot.json
// lists Grand Marshal's Stave (18873) under BOTH pvp:rank-18 (Rank 18)
// AND vendor:12782 "Captain O'Neal" (no rank field of its own) - and
// "vendor" outranks "pvp" in band.go's own sourceKindPriority, so
// sourceFor always picks Captain O'Neal's rank-less row, silently
// defeating pvpRankCap for every dual-listed item unless the vendor row
// inherits Rank the same way it already inherits a rep source's
// Standing (vendorInheritsRepStandingGate). An ordinary vendor with no
// matching pvp source is untouched.
func TestLoadLootIndexVendorInheritsPvpRank(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "vendor:12782", "kind": "vendor", "name": "Captain O'Neal", "items": [18873]},
			{"id": "pvp:rank-18", "kind": "pvp", "name": "Rank 18", "rank": 18, "items": [18873]},
			{"id": "vendor:1", "kind": "vendor", "name": "An Ordinary Vendor", "items": [999]}
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
	var vendorRow *itemSource
	for i, s := range idx[18873] {
		if s.Kind == "vendor" {
			vendorRow = &idx[18873][i]
		}
	}
	if vendorRow == nil {
		t.Fatalf("idx[18873] = %+v, want a vendor row", idx[18873])
	}
	if vendorRow.Rank != 18 {
		t.Errorf("Captain O'Neal's vendor row Rank = %d, want 18 (inherited from pvp:rank-18)", vendorRow.Rank)
	}
	ordinary := idx[999]
	if len(ordinary) != 1 || ordinary[0].Rank != 0 {
		t.Errorf("idx[999] = %+v, want the ordinary vendor's Rank untouched at 0", ordinary)
	}
}

// This lane's brief, item 3's own named case: Atiesh's own quest (id
// 9270) is a "kind": "quest" loot.json row, which the general
// raid-source Opens gate never sees at all (that loop skips "quest"
// rows entirely) - raidLockedQuestOpens is the separate table that
// catches it, and this test is the exact regression a future edit to
// either table could otherwise reintroduce silently.
// TestLoadLootIndexNoLongerHandGatesAtieshsQuest pins that
// raidLockedQuestOpens (data.go) no longer carries a hand-picked entry
// for Atiesh's own quest ids - this lane's brief, item 2: "replace the
// hand list with a rule" (band.go's legendaryGatedLater, gated on the
// candidate's own quality field) replaces it, one layer above this
// function: loadLootIndex/sourceFor have no item quality to gate on at
// all, so a quest reward's own itemSource.Opens is correctly empty
// here now - see TestBuildBandPoolGatesEveryLegendaryRegardlessOfSourceKind
// (band_test.go) for the actual gate, which buildBandPool applies
// before sourceFor ever runs.
func TestLoadLootIndexNoLongerHandGatesAtieshsQuest(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [],
		"quests": {
			"22589": [{"quest_id": 9270, "name": "Atiesh, Greatstaff of the Guardian", "faction": "both", "min_level": 60, "level": 60}]
		}
	}`
	if err := writeFile(t, filepath.Join(dir, "loot.json"), lootJSON); err != nil {
		t.Fatal(err)
	}
	idx, _, err := loadLootIndex(dir, nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	src, ok := idx[22589]
	if !ok || len(src) != 1 || src[0].Opens != "" {
		t.Fatalf("idx[22589] = %+v, want one source with Opens \"\" (unset - this layer never gates on quality)", src)
	}
	if _, ok := sourceFor(22589, 60, "alliance", "", idx); !ok {
		t.Fatal("sourceFor(22589, level 60) = not ok, want ok=true: sourceFor alone no longer gates Atiesh; buildBandPool's legendaryGatedLater does, from the candidate's own Quality")
	}
}

// TestLoadLootIndexCarriesPerQuestOpensThrough is the quest-gates lane's
// own regression, 2026-09-29: loot.json's `quests` map now states its
// OWN `opens` per entry (data/pipeline/loot/sources.py's
// apply_quest_opens_gate), computed for a quest whose classic-db
// turn-in item is itself a raid boss drop - "For All To See"/
// "Celebrating Good Times" (Onyxia Tooth Pendant/Blood Talisman, gated
// via the "Victory for the Horde/Alliance" predecessor's own Head of
// Onyxia turn-in) is the real shape this pins. An entry with no
// `opens` at all (ordinary, non-raid-gated quest) stays open.
func TestLoadLootIndexCarriesPerQuestOpensThrough(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [],
		"quests": {
			"18404": [
				{"quest_id": 7491, "name": "For All To See", "faction": "horde", "min_level": 60, "level": 60, "opens": "raids-1"},
				{"quest_id": 7496, "name": "Celebrating Good Times", "faction": "alliance", "min_level": 60, "level": 60, "opens": "raids-1"}
			],
			"9001": [
				{"quest_id": 100, "name": "An Ordinary Quest", "faction": "both", "min_level": 10, "level": 10}
			]
		}
	}`
	if err := writeFile(t, filepath.Join(dir, "loot.json"), lootJSON); err != nil {
		t.Fatal(err)
	}
	idx, _, err := loadLootIndex(dir, nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	gated, ok := idx[18404]
	if !ok || len(gated) != 2 || gated[0].Opens != "raids-1" || gated[1].Opens != "raids-1" {
		t.Fatalf("idx[18404] = %+v, want two sources both with Opens \"raids-1\"", gated)
	}
	if _, ok := sourceFor(18404, 60, "horde", "", idx); ok {
		t.Fatal("sourceFor(18404, level 60, horde) = ok, want ok=false: every source is Opens-gated")
	}
	open, ok := idx[9001]
	if !ok || len(open) != 1 || open[0].Opens != "" {
		t.Fatalf("idx[9001] = %+v, want one source with Opens empty (an ordinary quest)", open)
	}
}

// TestLoadLootIndexComputedOpensWinsOverTheHandList pins
// firstNonEmpty's own priority in loadLootIndex: a quest id
// raidLockedQuestOpens ALSO names must still show whatever the
// pipeline's own computed `opens` states, not silently be overridden by
// the hand list - the hand list is a fallback for ids the computation
// cannot see (the Ahn'Qiraj war-effort three), never a ceiling on one it
// can.
func TestLoadLootIndexComputedOpensWinsOverTheHandList(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [],
		"quests": {
			"21504": [
				{"quest_id": 8756, "name": "The Qiraji Conqueror", "faction": "both", "min_level": 60, "level": 60, "opens": "raids-1"}
			]
		}
	}`
	if err := writeFile(t, filepath.Join(dir, "loot.json"), lootJSON); err != nil {
		t.Fatal(err)
	}
	idx, _, err := loadLootIndex(dir, nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	src, ok := idx[21504]
	if !ok || len(src) != 1 || src[0].Opens != "raids-1" {
		t.Fatalf(
			"idx[21504] = %+v, want Opens \"raids-1\" (the computed value, not raidLockedQuestOpens's \"later\")",
			src,
		)
	}
}

func TestLoadLootIndexMissingFile(t *testing.T) {
	if _, _, err := loadLootIndex(t.TempDir(), nil); err == nil {
		t.Fatal("loadLootIndex on an empty dir: want an error, got nil")
	}
}

// TestLoadLootIndexPerQuestFactionSide is the quest-faction lane's own
// regression for the owner defect this lane fixes: item 270018
// (Hammerbone) showed as Alliance best-in-slot from quest 914 Leaders
// of the Fang, a horde-only quest, because loot.json's flat "quest"
// LootSource named the item with no per-quest faction at all and the
// item itself carries no factionRestriction. Two fixtures, a temp
// loot.json rather than testdata/reporoot's shared one, since this is
// specifically about the `quests` map's per-quest `faction`, not the
// rest of loadLootIndex's file shape:
//
//   - item 2001 mirrors Hammerbone exactly: one horde-only quest, no
//     item-level restriction - an alliance character must find no
//     obtainable source at all, a horde character must.
//   - item 2002 is reachable through two faction-mirrored quests, one
//     per side (the design's own "both" case) - it must carry TWO
//     itemSource records, never merged into one, so each faction finds
//     its own.
func TestLoadLootIndexPerQuestFactionSide(t *testing.T) {
	dir := t.TempDir()
	loot := `{
  "sources": [
    {"id": "quest", "kind": "quest", "name": "Quests", "items": [2001, 2002]}
  ],
  "quests": {
    "2001": [
      {"quest_id": 914, "name": "Leaders of the Fang", "faction": "horde", "faction_source": "classic-db", "min_level": 20, "level": 20, "level_source": "classic-db"}
    ],
    "2002": [
      {"quest_id": 100, "name": "A Horde Quest", "faction": "horde", "faction_source": "classic-db", "min_level": 20, "level": 20, "level_source": "classic-db"},
      {"quest_id": 200, "name": "An Alliance Mirror", "faction": "alliance", "faction_source": "classic-db", "min_level": 20, "level": 20, "level_source": "classic-db"}
    ]
  }
}`
	if err := os.WriteFile(filepath.Join(dir, "loot.json"), []byte(loot), 0o644); err != nil {
		t.Fatalf("writing loot.json: %v", err)
	}
	idx, _, err := loadLootIndex(dir, nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}

	if _, ok := sourceFor(2001, 20, "alliance", "", idx); ok {
		t.Error("sourceFor(2001, alliance) = ok, want false: quest 914 is horde-only")
	}
	if src, ok := sourceFor(2001, 20, "horde", "", idx); !ok || src.Label != "Leaders of the Fang" {
		t.Errorf("sourceFor(2001, horde) = %+v, %v, want ok=true, %q", src, ok, "Leaders of the Fang")
	}

	if got := len(idx[2002]); got != 2 {
		t.Fatalf("idx[2002] = %+v, want 2 records (one per quest, never merged)", idx[2002])
	}
	if src, ok := sourceFor(2002, 20, "horde", "", idx); !ok || src.Label != "A Horde Quest" {
		t.Errorf("sourceFor(2002, horde) = %+v, %v, want ok=true, %q", src, ok, "A Horde Quest")
	}
	if src, ok := sourceFor(2002, 20, "alliance", "", idx); !ok || src.Label != "An Alliance Mirror" {
		t.Errorf("sourceFor(2002, alliance) = %+v, %v, want ok=true, %q", src, ok, "An Alliance Mirror")
	}
}

func TestApplyEffectiveRequiredLevelsRaisesAQuestRewardsGate(t *testing.T) {
	idx, questFloors, err := loadLootIndex(buildDirFixture(), nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	items := []candidate{
		{ID: 1001, RequiredLevel: 10, ItemLevel: 20, EffectiveRequiredLevel: 10},
		{ID: 1002, RequiredLevel: 15, ItemLevel: 25, EffectiveRequiredLevel: 15},
	}
	got := applyEffectiveRequiredLevels(items, idx, questFloors)
	byID := map[int]candidate{}
	for _, c := range got {
		byID[c.ID] = c
	}
	// 1001's own required_level (10) is well under its quest's min_level
	// (25): the quest floor wins, the same shape as this lane's Polar
	// Leggings finding (item level 80, required_level 0, but level-60
	// quest-gated in the client).
	if byID[1001].EffectiveRequiredLevel != 25 {
		t.Errorf("candidate 1001 EffectiveRequiredLevel = %d, want 25", byID[1001].EffectiveRequiredLevel)
	}
	// 1002 is a dungeon drop, not a quest reward: no floor, unchanged.
	if byID[1002].EffectiveRequiredLevel != 15 {
		t.Errorf("candidate 1002 EffectiveRequiredLevel = %d, want 15 (unchanged)", byID[1002].EffectiveRequiredLevel)
	}
	// applyEffectiveRequiredLevels must not mutate its input slice
	// in place (immutability: this lane returns a new slice).
	if items[0].EffectiveRequiredLevel != 10 {
		t.Errorf("input slice was mutated: items[0].EffectiveRequiredLevel = %d, want unchanged 10", items[0].EffectiveRequiredLevel)
	}
}

func TestApplyEffectiveRequiredLevelsFallsBackToItemLevelProxyForACraftedItem(t *testing.T) {
	idx := lootIndex{2001: {{Kind: "crafted", Label: "Blacksmithing"}}}
	items := []candidate{{ID: 2001, RequiredLevel: 0, ItemLevel: 44, EffectiveRequiredLevel: 0}}
	got := applyEffectiveRequiredLevels(items, idx, map[int]int{})
	// leveling.ItemLevelProxyRequiredLevel(44) = min(60, 39) = 39.
	if got[0].EffectiveRequiredLevel != 39 {
		t.Errorf("crafted candidate EffectiveRequiredLevel = %d, want 39 (item_level_proxy)", got[0].EffectiveRequiredLevel)
	}
}

func TestLoadAllSpecs(t *testing.T) {
	specs, err := loadAllSpecs(repoRootFixture)
	if err != nil {
		t.Fatalf("loadAllSpecs: %v", err)
	}
	if len(specs) != 4 {
		t.Fatalf("loadAllSpecs = %d specs, want 4", len(specs))
	}
}

func TestLoadAllSpecsMissingFile(t *testing.T) {
	if _, err := loadAllSpecs(t.TempDir()); err == nil {
		t.Fatal("loadAllSpecs on an empty dir: want an error, got nil")
	}
}

func TestLoadSpec(t *testing.T) {
	s, err := loadSpec(repoRootFixture, "hunter-marksmanship")
	if err != nil {
		t.Fatalf("loadSpec: %v", err)
	}
	if s.ClassSlug != "hunter" || s.SpecSlug != "marksmanship" || s.TreeIndex != 1 {
		t.Errorf("loadSpec = %+v, want hunter/marksmanship tree_index 1", s)
	}
}

func TestLoadSpecUnknown(t *testing.T) {
	if _, err := loadSpec(repoRootFixture, "nonexistent-spec"); err == nil {
		t.Fatal("loadSpec(nonexistent-spec): want an error, got nil")
	}
}

func TestLoadAPLState(t *testing.T) {
	state, err := loadAPLState(repoRootFixture, "hunter-marksmanship")
	if err != nil {
		t.Fatalf("loadAPLState: %v", err)
	}
	if state != "written" {
		t.Errorf("state = %q, want written", state)
	}
}

func TestLoadAPLStateMissing(t *testing.T) {
	_, err := loadAPLState(repoRootFixture, "no-such-spec")
	if err == nil || !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("loadAPLState(no-such-spec) = %v, want a not-exist error", err)
	}
}

func TestWrittenSpecs(t *testing.T) {
	specs, err := writtenSpecs(repoRootFixture)
	if err != nil {
		t.Fatalf("writtenSpecs: %v", err)
	}
	// hunter-marksmanship: state "written" -> included.
	// mage-fire: state "draft" -> excluded.
	// priest-shadow: no apl file at all -> skipped, not an error.
	if len(specs) != 1 || specs[0] != "hunter-marksmanship" {
		t.Fatalf("writtenSpecs = %v, want [hunter-marksmanship]", specs)
	}
}

func TestLoadGuideRaces(t *testing.T) {
	races, err := loadGuideRaces(repoRootFixture, buildDirFixture(), "hunter", "marksmanship")
	if err != nil {
		t.Fatalf("loadGuideRaces: %v", err)
	}
	if races.AllianceRace != "dwarf" || races.HordeRace != "troll" {
		t.Errorf("races = %+v, want dwarf/troll", races)
	}
}

func TestLoadGuideRacesMissingFile(t *testing.T) {
	if _, err := loadGuideRaces(repoRootFixture, buildDirFixture(), "hunter", "nonexistent-spec"); err == nil {
		t.Fatal("loadGuideRaces for a missing guide: want an error, got nil")
	}
}

func TestLoadGuideRacesNoFrontmatterLine(t *testing.T) {
	dir := t.TempDir()
	guideDir := filepath.Join(dir, "web", "src", "content", "guides", "hunter")
	if err := writeFile(t, filepath.Join(guideDir, "broken.md"), "---\ntitle: no races here\n---\n"); err != nil {
		t.Fatal(err)
	}
	if _, err := loadGuideRaces(dir, buildDirFixture(), "hunter", "broken"); err == nil {
		t.Fatal("loadGuideRaces with no recommendedRaces line: want an error, got nil")
	}
}

// This lane's brief (bis-ranker-integrity, 2026-09-29), item 4: the
// exact shape of the paladin-retribution defect - a third race in the
// list silently taking HordeRace's slot instead of the real Horde race
// - must fail to load, not silently publish the wrong faction's race.
func TestLoadGuideRacesRejectsMoreThanTwoRaces(t *testing.T) {
	dir := t.TempDir()
	guideDir := filepath.Join(dir, "web", "src", "content", "guides", "paladin")
	if err := writeFile(t, filepath.Join(guideDir, "retribution.md"), "---\ntitle: broken\nrecommendedRaces: [human, dwarf, undead]\n---\n"); err != nil {
		t.Fatal(err)
	}
	_, err := loadGuideRaces(dir, buildDirFixture(), "paladin", "retribution")
	if err == nil {
		t.Fatal("loadGuideRaces with 3 races: want an error, got nil")
	}
}

// The other half of item 4: a race that IS a real race, but not on the
// faction the slot claims (a copy-paste of an Alliance-only race into
// HordeRace, or vice versa), must fail against races.json rather than
// publish an unplayable faction/race pairing.
func TestLoadGuideRacesRejectsAllianceRaceInHordeSlot(t *testing.T) {
	dir := t.TempDir()
	guideDir := filepath.Join(dir, "web", "src", "content", "guides", "paladin")
	if err := writeFile(t, filepath.Join(guideDir, "retribution.md"), "---\ntitle: broken\nrecommendedRaces: [human, dwarf]\n---\n"); err != nil {
		t.Fatal(err)
	}
	_, err := loadGuideRaces(dir, buildDirFixture(), "paladin", "retribution")
	if err == nil {
		t.Fatal("loadGuideRaces with an Alliance-only race (dwarf) in the Horde slot: want an error, got nil")
	}
}

// The battleground reputations' sides, keyed by the faction ids this
// build's loot.json carries (Silverwing Sentinels is 889, Warsong
// Outriders 890 -- easy to swap, and swapping them hands every Horde
// bow to Alliance lists). Read off the real build when it is checked
// out beside the module; skipped otherwise.
func TestRepSideMatchesTheBuildsFactionIDs(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("..", "..", "..", "data", "builds", "1.60.1.70009", "loot.json"))
	if err != nil {
		t.Skip("no real build beside the module: " + err.Error())
	}
	var f lootFile
	if err := json.Unmarshal(b, &f); err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"Warsong Outriders": "horde", "The Defilers": "horde", "Frostwolf Clan": "horde", "Silverwing Sentinels": "alliance", "The League of Arathor": "alliance", "Stormpike Guard": "alliance"}
	checked := 0
	for _, src := range f.Sources {
		side, named := want[src.Name]
		if src.Kind != "rep" || !named {
			continue
		}
		if got := repSide[src.FactionID]; got != side {
			t.Fatalf("%s (faction %d) maps to %q, want %q", src.Name, src.FactionID, got, side)
		}
		checked++
	}
	if checked == 0 {
		t.Fatal("the build's loot.json names none of the six battleground reputations")
	}
}

func TestCorrectedRepSourceSwapsWheneverTheItemsOwnRestrictionDisagreesWithTheMinedSide(t *testing.T) {
	names := map[int]string{889: "Silverwing Sentinels", 890: "Warsong Outriders"}
	// Scout's Medallion (20442, horde_only) mined under Silverwing
	// Sentinels/889 (alliance): corrected to Warsong Outriders/890.
	if id, label, swapped := correctedRepSource("horde", 889, names); !swapped || id != 890 || label != "Warsong Outriders" {
		t.Errorf("correctedRepSource(horde, 889) = %d, %q, %v, want 890, Warsong Outriders, true", id, label, swapped)
	}
	// Sentinel's Medallion (20444 AND, separately, 19541 - two
	// different items, same name and restriction) mined under Warsong
	// Outriders/890 (horde): corrected to Silverwing Sentinels/889. The
	// rule is driven by the item's own restriction, not its id, so it
	// catches every item this shape without needing to enumerate them.
	if id, label, swapped := correctedRepSource("alliance", 890, names); !swapped || id != 889 || label != "Silverwing Sentinels" {
		t.Errorf("correctedRepSource(alliance, 890) = %d, %q, %v, want 889, Silverwing Sentinels, true", id, label, swapped)
	}
	// An unrestricted item (most items, e.g. Rune of Perfection) is left
	// alone: there is no restriction to disagree with the mined Side.
	if id, _, swapped := correctedRepSource("", 889, names); swapped || id != 889 {
		t.Errorf("correctedRepSource(\"\", 889) swapped an unrestricted item: id=%d swapped=%v", id, swapped)
	}
	// A restriction that already agrees with the mined Side is left
	// alone (nothing to correct).
	if id, _, swapped := correctedRepSource("alliance", 889, names); swapped || id != 889 {
		t.Errorf("correctedRepSource(alliance, 889) swapped an already-consistent item: id=%d swapped=%v", id, swapped)
	}
}

func TestAnItemWithNoRequiredLevelIsGatedByItsItemLevel(t *testing.T) {
	items := []candidate{{ID: 270052, Name: "Swamp Ring", ItemLevel: 35, RequiredLevel: 0}}
	out := applyEffectiveRequiredLevels(items, lootIndex{}, map[int]int{})
	if got := out[0].EffectiveRequiredLevel; got != leveling.ItemLevelProxyRequiredLevel(35) {
		t.Fatalf("EffectiveRequiredLevel = %d, want the item-level proxy %d", got, leveling.ItemLevelProxyRequiredLevel(35))
	}
	stated := []candidate{{ID: 1, Name: "Stated", ItemLevel: 35, RequiredLevel: 12}}
	if got := applyEffectiveRequiredLevels(stated, lootIndex{}, map[int]int{})[0].EffectiveRequiredLevel; got != 12 {
		t.Fatalf("a stated required level must stand: got %d", got)
	}
}
