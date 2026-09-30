package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

// TestLoadCandidatesExcludesSupersededItems is this lane's brief
// (bis-ranker-integrity-10), item 2: data-followups-4's own
// superseded_by key means "never a candidate" - the pinned example
// being "Grand Marshal's Stave" 18873 (superseded_by 234571), a real
// client row Forever kept beside its own re-itemised copy. Read
// defensively off either row (flatItem.SupersededBy's own doc): this
// case sets it only on the flat items.json row, the shape data-
// followups-4's own brief describes first ("mark the legacy row...
// on the flat and per-class rows"), so this test also stands as the
// minimal case a build that has only regenerated one of the two rows
// still excludes correctly.
func TestLoadCandidatesExcludesSupersededItems(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "items.json"), `[
		{"id": 18873, "name": "Grand Marshal's Stave", "quality": 4, "item_level": 78, "required_level": 60, "class_id": 2, "subclass_id": 10, "inventory_type": 17, "suffixes": [], "faction_restriction": "", "superseded_by": 234571},
		{"id": 234571, "name": "Grand Marshal's Runed Stave", "quality": 4, "item_level": 80, "required_level": 60, "class_id": 2, "subclass_id": 10, "inventory_type": 17, "suffixes": [], "faction_restriction": ""}
	]`)
	writeFile(t, filepath.Join(dir, "items", "priest.json"), `{
		"build": "testbuild", "class_slug": "priest",
		"items": [
			{"id": 18873, "name": "Grand Marshal's Stave", "slot": "main_hand", "quality": 4, "required_level": 60, "item_level": 78, "armor": 0, "stats": {}, "damage_min": 100, "damage_max": 150, "speed": 2.8, "dps": 44.6, "two_hand": true, "effect_text": "", "set_id": null, "unique": true},
			{"id": 234571, "name": "Grand Marshal's Runed Stave", "slot": "main_hand", "quality": 4, "required_level": 60, "item_level": 80, "armor": 0, "stats": {"spirit": 10}, "damage_min": 105, "damage_max": 155, "speed": 2.8, "dps": 46.4, "two_hand": true, "effect_text": "", "set_id": null, "unique": true}
		]
	}`)
	candidates, missing, err := loadCandidates(dir, "priest")
	if err != nil {
		t.Fatalf("loadCandidates: %v", err)
	}
	if len(candidates) != 1 || candidates[0].ID != 234571 {
		t.Fatalf("loadCandidates = %+v, want only the re-itemised copy (234571), the legacy row (18873, superseded_by 234571) excluded", candidates)
	}
	found := false
	for _, m := range missing {
		if strings.Contains(m, "18873") && strings.Contains(m, "superseded_by 234571") {
			found = true
		}
	}
	if !found {
		t.Errorf("missing = %v, want a note naming 18873's superseded_by 234571", missing)
	}
}

// TestLoadCandidatesExcludesSupersededItemsMarkedOnlyOnClassRow
// confirms the defensive read the other way: a class regeneration
// that has landed superseded_by on items/<class>.json before the flat
// items.json regeneration has (or ever will, for a build the flat
// pipeline stage never revisits) still excludes the item - the flat
// row here carries no superseded_by at all.
func TestLoadCandidatesExcludesSupersededItemsMarkedOnlyOnClassRow(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "items.json"), `[
		{"id": 18874, "name": "High Warlord's War Staff", "quality": 4, "item_level": 78, "required_level": 60, "class_id": 2, "subclass_id": 10, "inventory_type": 17, "suffixes": [], "faction_restriction": ""}
	]`)
	writeFile(t, filepath.Join(dir, "items", "priest.json"), `{
		"build": "testbuild", "class_slug": "priest",
		"items": [
			{"id": 18874, "name": "High Warlord's War Staff", "slot": "main_hand", "quality": 4, "required_level": 60, "item_level": 78, "armor": 0, "stats": {}, "damage_min": 100, "damage_max": 150, "speed": 2.8, "dps": 44.6, "two_hand": true, "effect_text": "", "set_id": null, "unique": true, "superseded_by": 234549}
		]
	}`)
	candidates, missing, err := loadCandidates(dir, "priest")
	if err != nil {
		t.Fatalf("loadCandidates: %v", err)
	}
	if len(candidates) != 0 {
		t.Fatalf("loadCandidates = %+v, want zero candidates: the class row's own superseded_by must exclude it even with no flat-row match", candidates)
	}
	if len(missing) != 1 || !strings.Contains(missing[0], "superseded_by 234549") {
		t.Errorf("missing = %v, want a note naming superseded_by 234549", missing)
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
	// This lane's brief (bis-ranker-integrity-6), item 9, corrected per
	// the controller's own direct note: weapon_type lives on the
	// PER-CLASS row (items/<class>.json, classItem), not the flat
	// items.json row (flatItem) - loadCandidates must read it from
	// there.
	if bow.WeaponType != "bow" {
		t.Errorf("bow WeaponType = %q, want %q (from the class file, not the flat file)", bow.WeaponType, "bow")
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

// TestLoadLootIndexCraftedSourceWithPhaseSuffixedIDCarriesOpensThrough
// is this lane's brief (bis-ranker-integrity-10), item 2: data-
// followups-3's own `crafted:<profession>:<phase>` source id (the
// crafted-source builder's own gate, splitting a phase-restricted
// crafted item like Sulfuron Hammer out of a launch-open
// `crafted:blacksmithing` bucket) must carry `opens` "like a raid" -
// this test confirms the loader already reads it that way with no
// code change: nothing in loadLootIndex's add() closure special-cases
// "crafted" by id shape at all (unlike "world"/"vendor"/"rep"/"pvp",
// each of which DOES get a kind-specific override below) - every
// source's own Opens (src.Opens) is carried onto its itemSource in the
// one generic assignment every kind shares, so a longer, phase-
// suffixed crafted id changes nothing about how Opens reaches
// band.go's sourceObtainable, which itself gates on Opens being
// non-empty with no kind check either (see its own doc: "no raid is
// open on launch day at all... Opens non-empty... is refused at EVERY
// band, not only below 60"). Named explicitly per this lane's brief:
// wiring is confirmed by this test, not added.
func TestLoadLootIndexCraftedSourceWithPhaseSuffixedIDCarriesOpensThrough(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "crafted:blacksmithing:phase2", "kind": "crafted", "name": "Blacksmithing", "profession": "blacksmithing", "opens": "raids-1", "items": [12583]},
			{"id": "crafted:blacksmithing", "kind": "crafted", "name": "Blacksmithing", "profession": "blacksmithing", "items": [12719]}
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
	gated, ok := idx[12583]
	if !ok || len(gated) != 1 || gated[0].Kind != "crafted" || gated[0].Opens != "raids-1" {
		t.Fatalf("idx[12583] = %+v, want one crafted source with Opens \"raids-1\"", gated)
	}
	if sourceObtainable(gated[0], 60, "alliance", "") {
		t.Error("sourceObtainable(the phase-gated crafted source) = true at level 60, want false - Opens must gate a crafted source exactly like a raid, at every band")
	}
	open, ok := idx[12719]
	if !ok || len(open) != 1 || open[0].Kind != "crafted" || open[0].Opens != "" {
		t.Fatalf("idx[12719] = %+v, want one crafted source with Opens empty (launch-open)", open)
	}
	if !sourceObtainable(open[0], 20, "alliance", "") {
		t.Error("sourceObtainable(the ungated crafted source) = false at level 20, want true")
	}
}

// pvp-faction lane, 2026-09-29: loadLootIndex must copy a "pvp" kind
// source's own "faction" JSON field onto the itemSource's Side, the
// same fact factionExclusiveDungeons already supplies for a dungeon --
// see band_test.go's TestSourceForGatesAPvpRankItemToItsOwnFaction for
// the gating behaviour this makes possible.
func TestLoadLootIndexPvpSourceSideComesFromItsOwnFaction(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "pvp:rank-11:alliance", "kind": "pvp", "name": "Rank 11 (Alliance)",
			 "rank": 11, "faction": "alliance", "items": [16338]},
			{"id": "pvp:rank-11:horde", "kind": "pvp", "name": "Rank 11 (Horde)",
			 "rank": 11, "faction": "horde", "items": [16391]}
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
	alliance := idx[16338]
	if len(alliance) != 1 || alliance[0].Side != "alliance" || alliance[0].Rank != 11 {
		t.Errorf("idx[16338] = %+v, want one pvp source with Side alliance, Rank 11", alliance)
	}
	horde := idx[16391]
	if len(horde) != 1 || horde[0].Side != "horde" || horde[0].Rank != 11 {
		t.Errorf("idx[16391] = %+v, want one pvp source with Side horde, Rank 11", horde)
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

// TestLoadLootIndexVendorInheritsPvpFactionSideAndLabel is the fourth
// wow-player sweep's own item 1 (real production shape, post pvp-
// faction lane's rename): loot.json's own pvp:rank-18:alliance carries
// the split faction and this rank's title, and Captain O'Neal's vendor
// row shares the same item id - the vendor row must inherit BOTH Rank
// and Side (not just Rank, which is all the earlier fix checked), and
// its own Label must become the "PvP rank N · Title · Faction" line
// (item 2), never the bare quartermaster name.
func TestLoadLootIndexVendorInheritsPvpFactionSideAndLabel(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "vendor:12782", "kind": "vendor", "name": "Captain O'Neal", "items": [18873]},
			{"id": "pvp:rank-18:alliance", "kind": "pvp", "name": "Rank 18 (Alliance)",
			 "rank": 18, "faction": "alliance", "title": "Grand Marshal", "items": [18873]}
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
	if vendorRow.Rank != 18 || vendorRow.Side != "alliance" {
		t.Errorf("Captain O'Neal's vendor row = %+v, want Rank 18, Side alliance", vendorRow)
	}
	wantLabel := "PvP rank 18 · Grand Marshal · Alliance"
	if vendorRow.Label != wantLabel {
		t.Errorf("Captain O'Neal's vendor row Label = %q, want %q", vendorRow.Label, wantLabel)
	}
	// This lane's brief, item 1's own test ask: "an alliance rank-14
	// weapon under a Horde quartermaster row must be unobtainable for
	// Horde and capped out of default picks for Alliance." The inherited
	// Side (now "alliance", not "" as the pre-fix bug left it) is what
	// makes the first half true; pvpRankExceedsCap(pvpRankCap == 10) on
	// the inherited Rank (18) is what makes the second half true.
	if sourceObtainable(*vendorRow, 60, "horde", "") {
		t.Errorf("vendor row with inherited Side %q must be UNobtainable for horde", vendorRow.Side)
	}
	if !sourceObtainable(*vendorRow, 60, "alliance", "") {
		t.Errorf("vendor row with inherited Side %q must be obtainable for its own faction", vendorRow.Side)
	}
	if !pvpRankExceedsCap(*vendorRow) {
		t.Errorf("vendor row with inherited Rank %d must exceed pvpRankCap (%d), capping it out of the default pick", vendorRow.Rank, pvpRankCap)
	}
}

// TestLoadLootIndexVendorInheritsPvpRankAmbiguousFactionUsesVendorsOwnSide
// is the pathological case this lane's brief calls out: an alliance
// rank-14 (ladder position) weapon somehow shares its classic item id
// with a Horde vendor row too (a data anomaly, not something today's
// build actually produces - verified against build 1.60.1.70009's own
// loot.json directly). The vendor row's own known side (pipeline.loot.
// pvp_faction's npc-faction resolution, carried as lootSource.Faction)
// must decide which of the two same-id pvp sources it inherits from,
// rather than the first one found.
func TestLoadLootIndexVendorInheritsPvpRankAmbiguousFactionUsesVendorsOwnSide(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "vendor:horde-qm", "kind": "vendor", "name": "Horde Quartermaster",
			 "faction": "horde", "items": [99001]},
			{"id": "pvp:rank-18:alliance", "kind": "pvp", "name": "Rank 18 (Alliance)",
			 "rank": 18, "faction": "alliance", "title": "Grand Marshal", "items": [99001]},
			{"id": "pvp:rank-18:horde", "kind": "pvp", "name": "Rank 18 (Horde)",
			 "rank": 18, "faction": "horde", "title": "High Warlord", "items": [99001]}
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
	for i, s := range idx[99001] {
		if s.Kind == "vendor" {
			vendorRow = &idx[99001][i]
		}
	}
	if vendorRow == nil {
		t.Fatalf("idx[99001] = %+v, want a vendor row", idx[99001])
	}
	if vendorRow.Side != "horde" || vendorRow.Title != "High Warlord" {
		t.Errorf("Horde Quartermaster's vendor row = %+v, want Side horde, Title High Warlord (matched by the vendor's own known side, not the first pvp source found)", vendorRow)
	}
}

// TestLoadLootIndexVendorInheritsPvpRankLeavesAmbiguousUnknownSideAlone
// pins the "left out, never shown as fact" half of the same rule: when
// two same-id pvp sources exist and the vendor row's OWN side is
// unknown (no Faction pipeline.loot.pvp_faction could resolve), nothing
// is guessed - the vendor row keeps whatever Rank/Side it already had
// (none, in this fixture).
func TestLoadLootIndexVendorInheritsPvpRankLeavesAmbiguousUnknownSideAlone(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "vendor:unknown", "kind": "vendor", "name": "Unresolved Vendor", "items": [99002]},
			{"id": "pvp:rank-18:alliance", "kind": "pvp", "name": "Rank 18 (Alliance)",
			 "rank": 18, "faction": "alliance", "items": [99002]},
			{"id": "pvp:rank-18:horde", "kind": "pvp", "name": "Rank 18 (Horde)",
			 "rank": 18, "faction": "horde", "items": [99002]}
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
	for i, s := range idx[99002] {
		if s.Kind == "vendor" {
			vendorRow = &idx[99002][i]
		}
	}
	if vendorRow == nil {
		t.Fatalf("idx[99002] = %+v, want a vendor row", idx[99002])
	}
	if vendorRow.Rank != 0 || vendorRow.Side != "" {
		t.Errorf("Unresolved Vendor's row = %+v, want Rank 0, Side \"\" (left alone, not guessed)", vendorRow)
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

// This lane's brief, item 4: Earthstrike (21180,
// rep:cenarion-circle:exalted) published with no gate at all - loot.json
// never states an "opens" value for any rep-kind source, but Cenarion
// Circle (faction 609) is Gates of Ahn'Qiraj content, the same patch
// this build already gates raid:ahnqiraj "later" for.
// repFactionRaidPhaseOpens (data.go) supplies that gate on top of
// loot.json's own (always-empty) rep opens; an ordinary reputation
// faction not in that map stays open at launch, exactly as a previous
// lane already confirmed for reputation sources in general.
func TestLoadLootIndexGatesCenarionCircleRepToLaterPhase(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "rep:cenarion-circle:exalted", "kind": "rep", "name": "Cenarion Circle", "faction_id": 609, "standing": "exalted", "items": [21180]},
			{"id": "rep:timbermaw-hold:friendly", "kind": "rep", "name": "Timbermaw Hold", "faction_id": 576, "standing": "friendly", "items": [9999]}
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
	earthstrike, ok := idx[21180]
	if !ok || len(earthstrike) != 1 || earthstrike[0].Opens != "later" {
		t.Fatalf("idx[21180] (Earthstrike, Cenarion Circle exalted) = %+v, want one source with Opens \"later\"", earthstrike)
	}
	ordinary, ok := idx[9999]
	if !ok || len(ordinary) != 1 || ordinary[0].Opens != "" {
		t.Fatalf("idx[9999] (Timbermaw Hold, an ordinary launch-day reputation) = %+v, want Opens empty", ordinary)
	}
}

// Player-review sweep 15/16, 2026-09-30: the same gap as Earthstrike
// above, a second AQ War Effort faction. "Signet Ring of the Bronze
// Dragonflight" (21200/21205/21210) is reachable directly by reputation
// (rep:brood-of-nozdormu:exalted) with no gate at all in loot.json - the
// quest-reward path is fixed at the pipeline level
// (pipeline.loot.sources.REP_FACTION_RAID_PHASE_OPENS bakes the gate
// into loot.json's own quest entries), but this rep-kind source is a
// second, independent way to reach the same item, and an item counts as
// launch-day reachable the moment any ONE of its sources is ungated -
// so this fallback has to cover it too, exactly the way Cenarion
// Circle's own entry already does for Earthstrike.
func TestLoadLootIndexGatesBroodOfNozdormuRepToLaterPhase(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "rep:brood-of-nozdormu:exalted", "kind": "rep", "name": "Brood of Nozdormu", "faction_id": 910, "standing": "exalted", "items": [21210]}
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
	ring, ok := idx[21210]
	if !ok || len(ring) != 1 || ring[0].Opens != "later" {
		t.Fatalf(
			"idx[21210] (Signet Ring of the Bronze Dragonflight, Brood of Nozdormu exalted) = %+v, want one source with Opens \"later\"",
			ring,
		)
	}
}

// A rep source loot.json itself already gives an explicit opens value
// for must keep that computed value, not the hand-maintained
// repFactionRaidPhaseOpens fallback - the same firstNonEmpty priority
// TestLoadLootIndexComputedOpensWinsOverTheHandList pins for quests.
func TestLoadLootIndexComputedRepOpensWinsOverTheHandList(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "rep:cenarion-circle:exalted", "kind": "rep", "name": "Cenarion Circle", "faction_id": 609, "standing": "exalted", "items": [21180], "opens": "raids-1"}
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
	src, ok := idx[21180]
	if !ok || len(src) != 1 || src[0].Opens != "raids-1" {
		t.Fatalf("idx[21180] = %+v, want Opens \"raids-1\" (the computed value, not repFactionRaidPhaseOpens's \"later\")", src)
	}
}

// bis-ranker-integrity-7 lane, item 3: worldBossSources (data.go)
// gates the six named world bosses to "raids-1" even though loot.json's
// own "world" kind (shared with every ordinary named-mob world drop)
// never carries an opens value at all - a fresh level-60 character does
// not solo Lord Kazzak or a Dragon of Nightmare. An ordinary world
// source (a plain mob drop, not a world boss) stays open at launch.
func TestLoadLootIndexGatesWorldBossesToRaidsOnePhase(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "world:lord-kazzak", "kind": "world", "name": "Lord Kazzak", "items": [18543]},
			{"id": "world:azuregos", "kind": "world", "name": "Azuregos", "items": [18202]},
			{"id": "world:emeriss", "kind": "world", "name": "Emeriss", "items": [20579]},
			{"id": "world:lethon", "kind": "world", "name": "Lethon", "items": [20625]},
			{"id": "world:taerar", "kind": "world", "name": "Taerar", "items": [20631]},
			{"id": "world:ysondre", "kind": "world", "name": "Ysondre", "items": [20635]},
			{"id": "world:aged-kodo", "kind": "world", "name": "Aged Kodo", "items": [6249]}
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
	for id, boss := range map[int]string{
		18543: "Lord Kazzak", 18202: "Azuregos", 20579: "Emeriss",
		20625: "Lethon", 20631: "Taerar", 20635: "Ysondre",
	} {
		src, ok := idx[id]
		if !ok || len(src) != 1 || src[0].Opens != "raids-1" {
			t.Fatalf("idx[%d] (%s) = %+v, want one source with Opens \"raids-1\"", id, boss, src)
		}
	}
	ordinary, ok := idx[6249]
	if !ok || len(ordinary) != 1 || ordinary[0].Opens != "" {
		t.Fatalf("idx[6249] (Aged Kodo, an ordinary world mob) = %+v, want Opens empty", ordinary)
	}
}

// A world boss source loot.json itself already gives an explicit opens
// value for must keep that computed value, not the hand-maintained
// worldBossSources fallback - the same firstNonEmpty priority
// TestLoadLootIndexComputedOpensWinsOverTheHandList pins for quests and
// TestLoadLootIndexComputedRepOpensWinsOverTheHandList pins for
// reputation.
func TestLoadLootIndexComputedWorldBossOpensWinsOverTheHandList(t *testing.T) {
	dir := t.TempDir()
	lootJSON := `{
		"sources": [
			{"id": "world:lord-kazzak", "kind": "world", "name": "Lord Kazzak", "items": [18543], "opens": "later"}
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
	src, ok := idx[18543]
	if !ok || len(src) != 1 || src[0].Opens != "later" {
		t.Fatalf("idx[18543] = %+v, want Opens \"later\" (the computed value, not worldBossSources's \"raids-1\")", src)
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

// TestLoadLootIndexCarriesQuestClassesThrough is this lane's brief
// (bis-ranker-integrity-12), item 3: loot.json's own quests map does
// not carry a "classes" field yet (checked directly against this
// build's own data/builds/1.60.1.70009/loot.json before writing this
// fix), so this test builds the field synthetically - a data lane
// running in parallel is expected to add it, the pinned example being
// quest 8253 "Destroy Morphaz" (a MAGE class quest whose reward, Fire
// Ruby, can never be a hunter's or shaman's trinket). loadLootIndex
// must carry it through onto the quest's own itemSource unchanged, and
// a quest with no "classes" at all (every quest in this build today)
// must carry a nil Classes, not an empty-but-non-nil one that would
// read differently to a future consumer.
func TestLoadLootIndexCarriesQuestClassesThrough(t *testing.T) {
	dir := t.TempDir()
	loot := `{
  "sources": [],
  "quests": {
    "20036": [
      {"quest_id": 8253, "name": "Destroy Morphaz", "faction": "both", "min_level": 50, "level": 52, "classes": ["mage"]}
    ],
    "9001": [
      {"quest_id": 100, "name": "An Ordinary Quest", "faction": "both", "min_level": 10, "level": 10}
    ]
  }
}`
	if err := writeFile(t, filepath.Join(dir, "loot.json"), loot); err != nil {
		t.Fatal(err)
	}
	idx, _, err := loadLootIndex(dir, nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	restricted, ok := idx[20036]
	if !ok || len(restricted) != 1 || len(restricted[0].Classes) != 1 || restricted[0].Classes[0] != "mage" {
		t.Fatalf("idx[20036] = %+v, want one source with Classes [\"mage\"]", restricted)
	}
	unrestricted, ok := idx[9001]
	if !ok || len(unrestricted) != 1 || unrestricted[0].Classes != nil {
		t.Fatalf("idx[9001] = %+v, want one source with Classes nil (no restriction stated)", unrestricted)
	}
}

// TestClassAllowed pins classAllowed's own nil-safe contract (band.go):
// no restriction at all is always allowed; a restriction only ever
// allows the classes it names, case-insensitively (loot.json's own
// data lane doc gives no guarantee on casing).
func TestClassAllowed(t *testing.T) {
	if !classAllowed(nil, "rogue") {
		t.Error("classAllowed(nil, rogue) = false, want true: no restriction stated")
	}
	if !classAllowed([]string{}, "rogue") {
		t.Error("classAllowed([], rogue) = false, want true: an empty list is the same as no restriction")
	}
	if !classAllowed([]string{"mage"}, "mage") {
		t.Error("classAllowed([mage], mage) = false, want true")
	}
	if !classAllowed([]string{"Mage"}, "mage") {
		t.Error("classAllowed([Mage], mage) = false, want true: case-insensitive")
	}
	if classAllowed([]string{"mage"}, "hunter") {
		t.Error("classAllowed([mage], hunter) = true, want false: hunter is not in the list")
	}
}

// TestBuildBandPoolExcludesAClassRestrictedQuestRewardForTheWrongClass
// is the real call site, end to end, with a synthetic quests map
// (this lane's brief, item 3's own instruction: "add the reader
// behind a nil-safe check with a test that uses a synthetic quests
// map, so the nightly picks it up the moment the data lane lands").
// Fire Ruby (quest 8253) must never reach a hunter's candidate pool,
// but must still reach a mage's - the exact fix for hunter-beast-
// mastery/shaman-elemental band 50's own repro (this lane's item 1).
func TestBuildBandPoolExcludesAClassRestrictedQuestRewardForTheWrongClass(t *testing.T) {
	items := []candidate{
		{ID: 20036, Name: "Fire Ruby", RequiredLevel: 50, EffectiveRequiredLevel: 50, Slots: []string{"trinket1", "trinket2"}},
	}
	idx := lootIndex{
		20036: {{Kind: "quest", Label: "Destroy Morphaz", Classes: []string{"mage"}}},
	}
	weights := map[string]float64{}
	hunterPool := buildBandPool(items, idx, "hunter", 50, "alliance", weights, 0, false)
	if len(hunterPool.Scored) != 0 {
		t.Fatalf("hunter pool.Scored = %+v, want empty: Fire Ruby's own quest is mage-only", hunterPool.Scored)
	}
	if len(hunterPool.NoSource) != 1 || hunterPool.NoSource[0].ID != 20036 {
		t.Fatalf("hunter pool.NoSource = %+v, want Fire Ruby (class-excluded, same bucket as no source at all)", hunterPool.NoSource)
	}
	magePool := buildBandPool(items, idx, "mage", 50, "alliance", weights, 0, false)
	if len(magePool.Scored) != 1 || magePool.Scored[0].ID != 20036 {
		t.Fatalf("mage pool.Scored = %+v, want Fire Ruby (the quest's own class)", magePool.Scored)
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

// TestRogueSpecsWeighStrength is this lane's brief
// (bis-ranker-integrity-12), item 5's second half: rogue band 20's own
// tie repro (Serpent Gloves vs Gloves of the Fang) traced to
// data/curated/specs.json never listing "strength" among rogue's
// weight_stats at all - not a noisy near-zero measurement (item 1's
// retry/carry shape), a stat this build's own weights sweep never
// measured for a rogue in the first place, even though Strength
// converts to Attack Power for a rogue in Classic. Reads the REAL
// curated file (publishedRepoRoot, published_test.go's own constant),
// not a test fixture, since the fixture pre-dates this fix and a
// regression here would only ever show up against the real file.
func TestRogueSpecsWeighStrength(t *testing.T) {
	for _, spec := range []string{"rogue-assassination", "rogue-combat", "rogue-subtlety"} {
		s, err := loadSpec(publishedRepoRoot, spec)
		if err != nil {
			t.Fatalf("loadSpec(%s): %v", spec, err)
		}
		found := false
		for _, stat := range s.WeightStats {
			if stat == "strength" {
				found = true
			}
		}
		if !found {
			t.Errorf("%s weight_stats = %v, want \"strength\" included (a rogue converts it to attack power 1:1)", spec, s.WeightStats)
		}
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

// wantRatingFactors is the fixture build's own gametables/
// combatratings.txt level 60 row (testdata/reporoot/data/builds/
// testbuild/gametables/combatratings.txt), copied from
// data/tests/fixtures/sim/combatratings.txt, itself the real client's
// 1.60.1.70009 level 60 row (data/builds/1.60.1.70009/gametables/
// combatratings.txt): 10 hit rating = 1% hit, 14 crit = 1% crit,
// defense 1, dodge 12, parry 15, block 5.
var wantRatingFactors = ratingFactors{
	"hit": 10, "crit": 14, "dodge": 12, "parry": 15, "block": 5, "defense": 1,
}

func TestLoadRatingFactorsReadsTheBuildsLevel60Row(t *testing.T) {
	got, err := loadRatingFactors(buildDirFixture())
	if err != nil {
		t.Fatalf("loadRatingFactors: %v", err)
	}
	for stat, want := range wantRatingFactors {
		if got[stat] != want {
			t.Errorf("factors[%q] = %v, want %v", stat, got[stat], want)
		}
	}
}

func TestLoadRatingFactorsMissingFile(t *testing.T) {
	if _, err := loadRatingFactors(t.TempDir()); err == nil {
		t.Fatal("loadRatingFactors on a dir with no gametables/combatratings.txt: want an error, got nil")
	}
}

func TestLoadRatingFactorsMissingColumn(t *testing.T) {
	dir := t.TempDir()
	writeGameTable(t, dir, "Level\tDodge\tParry\tBlock\tHit - Melee\tHit - Ranged\tHit - Spell\tCrit - Melee\tCrit - Ranged\tCrit - Spell\n60\t12\t15\t5\t10\t10\t10\t14\t14\t14\n")
	if _, err := loadRatingFactors(dir); err == nil {
		t.Fatal("loadRatingFactors with no Defense Skill column: want an error, got nil")
	}
}

func TestLoadRatingFactorsNoLevel60Row(t *testing.T) {
	dir := t.TempDir()
	writeGameTable(t, dir, "Level\tDefense Skill\tDodge\tParry\tBlock\tHit - Melee\tHit - Ranged\tHit - Spell\tCrit - Melee\tCrit - Ranged\tCrit - Spell\n59\t1\t12\t15\t5\t10\t10\t10\t14\t14\t14\n")
	if _, err := loadRatingFactors(dir); err == nil {
		t.Fatal("loadRatingFactors with no level 60 row: want an error, got nil")
	}
}

func TestLoadRatingFactorsDisagreeingColumnsErrors(t *testing.T) {
	// Hit - Melee and Hit - Ranged disagree (10 vs 11): loadRatingFactors
	// must fail loudly rather than silently pick one, mirroring
	// data/pipeline/simdb/ratings.py's own load_rating_factors.
	dir := t.TempDir()
	writeGameTable(t, dir, "Level\tDefense Skill\tDodge\tParry\tBlock\tHit - Melee\tHit - Ranged\tHit - Spell\tCrit - Melee\tCrit - Ranged\tCrit - Spell\n60\t1\t12\t15\t5\t10\t11\t10\t14\t14\t14\n")
	if _, err := loadRatingFactors(dir); err == nil {
		t.Fatal("loadRatingFactors with disagreeing hit columns: want an error, got nil")
	}
}

func writeGameTable(t *testing.T, buildDir, contents string) {
	t.Helper()
	dir := filepath.Join(buildDir, "gametables")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "combatratings.txt"), []byte(contents), 0o644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
}

func TestConvertRatingStatsDividesRatingFamilyStatsOnly(t *testing.T) {
	stats := map[string]float64{"crit": 14, "hit": 20, "agility": 10, "spell_power": 5}
	got := convertRatingStats(stats, wantRatingFactors)
	want := map[string]float64{"crit": 1, "hit": 2, "agility": 10, "spell_power": 5}
	for stat, w := range want {
		if got[stat] != w {
			t.Errorf("convertRatingStats[%q] = %v, want %v", stat, got[stat], w)
		}
	}
	// Immutability: the input map must be untouched.
	if stats["crit"] != 14 || stats["hit"] != 20 {
		t.Errorf("convertRatingStats mutated its input: %+v", stats)
	}
}

// TestScoreAgreesForARatingItemOnceConvertedAtFactor is this lane's
// brief, item 2's exact synthetic-item test: a +14 crit RATING item
// (Medallion of the Dawn-shaped: 14 crit rating at this build's own
// factor of 14) must, once run through convertCandidateRatings, score
// identically to a hand-built +1 crit STAT UNIT item against the same
// weight - proving the ranker now scores in the same sim units
// (percent) the weights sweep measured a weight in, not raw rating
// points.
func TestScoreAgreesForARatingItemOnceConvertedAtFactor(t *testing.T) {
	ratingItem := []candidate{{ID: 1, Stats: map[string]float64{"crit": 14}}}
	converted := convertCandidateRatings(ratingItem, ratingFactors{"crit": 14})
	weights := map[string]float64{"crit": 1.0}

	got := score(converted[0], "neck", weights, 0, false)
	statUnitItem := candidate{Stats: map[string]float64{"crit": 1}}
	want := score(statUnitItem, "neck", weights, 0, false)

	if got != want || got != 1.0 {
		t.Fatalf("score(converted +14 crit rating) = %v, want %v (== 1.0, matching +1 crit stat unit at weight 1.0)", got, want)
	}
	// The original, unconverted candidate slice is untouched
	// (convertCandidateRatings returns a new slice/new Stats maps).
	if ratingItem[0].Stats["crit"] != 14 {
		t.Errorf("convertCandidateRatings mutated its input: %+v", ratingItem[0].Stats)
	}
}
