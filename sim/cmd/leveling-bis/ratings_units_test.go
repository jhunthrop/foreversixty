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
	items, _, err := loadCandidates(buildDir, "warrior")
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
	for _, want := range []struct {
		id        int
		hit, crit float64
	}{
		{lionheartHelmID, 2, 2},
		{furyVisorID, 1, 1},
	} {
		c, ok := byID[want.id]
		if !ok {
			t.Fatalf("item %d missing from the warrior candidates", want.id)
		}
		if c.Stats["hit"] != want.hit || c.Stats["crit"] != want.crit {
			t.Errorf("item %d scores hit %v / crit %v, want %v / %v", want.id, c.Stats["hit"], c.Stats["crit"], want.hit, want.crit)
		}
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
	a, r := score(converted[0], "head", weights, 0, false), score(converted[1], "head", weights, 0, false)
	if a != 1.5 || r != 1.5 {
		t.Fatalf("1%% aura hit scored %v and 10-rating hit scored %v, both want 1.5", a, r)
	}
}
