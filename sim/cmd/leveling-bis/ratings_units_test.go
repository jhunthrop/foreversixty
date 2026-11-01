package main

import (
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

const (
	lionheartHelmID = 12640 // 20 hit / 28 crit rating points
	furyVisorID     = 20521 // on-equip aura: a literal 1% hit / 1% crit
	arcanoweaveID   = 272411
)

func TestConvertRatingStatsLeavesAuraPercentUntouched(t *testing.T) {
	stats := map[string]float64{"hit": 1, "crit": 1, "strength": 18}
	percent := map[string]float64{"hit": 1, "crit": 1}
	got := convertRatingStats(stats, percent, wantRatingFactors)
	if got["hit"] != 1 || got["crit"] != 1 || got["strength"] != 18 {
		t.Fatalf("aura-only item converted to %+v, want hit 1, crit 1, strength 18", got)
	}
}

func TestConvertRatingStatsSplitsAMixedStat(t *testing.T) {
	// 20 rating points (2%) plus an aura's literal 1% share one stats entry.
	stats := map[string]float64{"hit": 21}
	percent := map[string]float64{"hit": 1}
	got := convertRatingStats(stats, percent, wantRatingFactors)
	if got["hit"] != 3 {
		t.Fatalf("mixed hit = %v, want 3 (20/10 + 1)", got["hit"])
	}
}

func TestConvertRatingStatsConvertsEveryRatingFamilyStat(t *testing.T) {
	stats := map[string]float64{"hit": 10, "crit": 14, "dodge": 12, "parry": 15, "block": 5, "defense": 4}
	got := convertRatingStats(stats, nil, wantRatingFactors)
	for stat, want := range map[string]float64{"hit": 1, "crit": 1, "dodge": 1, "parry": 1, "block": 1, "defense": 4} {
		if got[stat] != want {
			t.Errorf("%s = %v, want %v", stat, got[stat], want)
		}
	}
}

func TestRankerScoresRealItemsInEnginePercent(t *testing.T) {
	activeBuild, err := leveling.ReadActiveBuild(publishedRepoRoot)
	if err != nil {
		t.Fatalf("ReadActiveBuild: %v", err)
	}
	buildDir := filepath.Join(publishedRepoRoot, "data", "builds", activeBuild)
	items, _, _, err := loadCandidates(buildDir, "warrior")
	if err != nil {
		t.Fatalf("loadCandidates: %v", err)
	}
	factors, err := loadRatingFactors(buildDir)
	if err != nil {
		t.Fatalf("loadRatingFactors: %v", err)
	}
	byID := map[int]candidate{}
	for _, c := range convertCandidateRatings(items, factors) {
		byID[c.ID] = c
	}
	c, ok := byID[lionheartHelmID]
	if !ok {
		t.Fatalf("item %d missing from the warrior candidates", lionheartHelmID)
	}
	if c.Stats["hit"] != 2 || c.Stats["crit"] != 2 {
		t.Errorf("item %d scores hit %v / crit %v, want 2 / 2", lionheartHelmID, c.Stats["hit"], c.Stats["crit"])
	}
	// Every aura-percent item in this build is client_unconfirmed, so none
	// is a candidate any more (task L1.6, loadCandidates); the conversion
	// is still checked on the real row.
	if _, ok := byID[furyVisorID]; ok {
		t.Errorf("item %d (client_unconfirmed) is a warrior candidate", furyVisorID)
	}
	classFile, err := loadClassItems(buildDir, "warrior")
	if err != nil {
		t.Fatalf("loadClassItems: %v", err)
	}
	found := false
	for _, ci := range classFile.Items {
		if ci.ID != furyVisorID {
			continue
		}
		found = true
		got := convertRatingStats(ci.Stats, ci.PercentStats, factors)
		if got["hit"] != 1 || got["crit"] != 1 {
			t.Errorf("item %d converts to hit %v / crit %v, want 1 / 1", furyVisorID, got["hit"], got["crit"])
		}
	}
	if !found {
		t.Fatalf("item %d missing from items/warrior.json", furyVisorID)
	}
	if got := byID[arcanoweaveID].Stats["hit"]; got != 1 {
		t.Errorf("Arcanoweave Cloak hit = %v, want 1", got)
	}
}

func TestPercentSharedScoreMatchesAStatUnitItem(t *testing.T) {
	aura := candidate{Stats: map[string]float64{"hit": 1}, PercentStats: map[string]float64{"hit": 1}}
	rating := candidate{Stats: map[string]float64{"hit": 10}}
	weights := map[string]float64{"hit": 1.5}
	converted := convertCandidateRatings([]candidate{aura, rating}, wantRatingFactors)
	a, r := score(converted[0], "head", weights, 0), score(converted[1], "head", weights, 0)
	if a != 1.5 || r != 1.5 {
		t.Fatalf("1%% aura hit scored %v and 10-rating hit scored %v, both want 1.5", a, r)
	}
}

// A shield's block value is a flat amount taken off a blocked hit, not a
// rating: dividing it by the block factor would score it as chance.
func TestConvertRatingStatsLeavesBlockValueFlat(t *testing.T) {
	got := convertRatingStats(map[string]float64{"block": 5, "block_value": 27}, nil, wantRatingFactors)
	if got["block"] != 1 || got["block_value"] != 27 {
		t.Fatalf("block %v, block_value %v; want 1 and 27", got["block"], got["block_value"])
	}
}

const immovableObjectID = 19321

// The shield's flat block value in the active client: 25 since 1.60.1.70425
// moved it to item level 60 (27 before).
const immovableObjectBlockValue = 25.0

func TestTheImmovableObjectScoresBlockValueAndNoBlockChance(t *testing.T) {
	activeBuild, err := leveling.ReadActiveBuild(publishedRepoRoot)
	if err != nil {
		t.Fatalf("ReadActiveBuild: %v", err)
	}
	buildDir := filepath.Join(publishedRepoRoot, "data", "builds", activeBuild)
	items, _, _, err := loadCandidates(buildDir, "warrior")
	if err != nil {
		t.Fatalf("loadCandidates: %v", err)
	}
	factors, err := loadRatingFactors(buildDir)
	if err != nil {
		t.Fatalf("loadRatingFactors: %v", err)
	}
	for _, c := range convertCandidateRatings(items, factors) {
		if c.ID != immovableObjectID {
			continue
		}
		if c.Stats["block_value"] != immovableObjectBlockValue || c.Stats["block"] != 0 {
			t.Fatalf("The Immovable Object scores block_value %v, block %v; want %v and 0", c.Stats["block_value"], c.Stats["block"], immovableObjectBlockValue)
		}
		return
	}
	t.Fatalf("item %d missing from the warrior candidates", immovableObjectID)
}
