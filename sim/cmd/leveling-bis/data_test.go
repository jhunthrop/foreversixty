package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
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
	idx, questFloors, err := loadLootIndex(buildDirFixture())
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
	if !ok || len(ragefire) != 1 || ragefire[0].Faction != "horde" {
		t.Errorf("idx[1004] = %+v, want one source with Faction \"horde\"", ragefire)
	}
	// An ordinary dungeon source (not in factionExclusiveDungeons)
	// carries no Faction at all - it must not inherit one by accident.
	if boss[0].Faction != "" {
		t.Errorf("idx[1002][0].Faction = %q, want empty (not a faction-exclusive source)", boss[0].Faction)
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
}

func TestLoadLootIndexMissingFile(t *testing.T) {
	if _, _, err := loadLootIndex(t.TempDir()); err == nil {
		t.Fatal("loadLootIndex on an empty dir: want an error, got nil")
	}
}

func TestApplyEffectiveRequiredLevelsRaisesAQuestRewardsGate(t *testing.T) {
	idx, questFloors, err := loadLootIndex(buildDirFixture())
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
	if len(specs) != 3 {
		t.Fatalf("loadAllSpecs = %d specs, want 3", len(specs))
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
	races, err := loadGuideRaces(repoRootFixture, "hunter", "marksmanship")
	if err != nil {
		t.Fatalf("loadGuideRaces: %v", err)
	}
	if races.AllianceRace != "dwarf" || races.HordeRace != "troll" {
		t.Errorf("races = %+v, want dwarf/troll", races)
	}
}

func TestLoadGuideRacesMissingFile(t *testing.T) {
	if _, err := loadGuideRaces(repoRootFixture, "hunter", "nonexistent-spec"); err == nil {
		t.Fatal("loadGuideRaces for a missing guide: want an error, got nil")
	}
}

func TestLoadGuideRacesNoFrontmatterLine(t *testing.T) {
	dir := t.TempDir()
	guideDir := filepath.Join(dir, "web", "src", "content", "guides", "hunter")
	if err := writeFile(t, filepath.Join(guideDir, "broken.md"), "---\ntitle: no races here\n---\n"); err != nil {
		t.Fatal(err)
	}
	if _, err := loadGuideRaces(dir, "hunter", "broken"); err == nil {
		t.Fatal("loadGuideRaces with no recommendedRaces line: want an error, got nil")
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
