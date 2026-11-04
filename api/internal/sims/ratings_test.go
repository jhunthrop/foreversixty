package sims

import (
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// The API's rating table is docs/contracts/fixtures/rating-per-percent.json, which the
// site's vitest pins to the client's combatratings.txt.
func TestRatingPerPercentIsTheContractFixture(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("..", "..", "..", "docs", "contracts", "fixtures", "rating-per-percent.json"))
	if err != nil {
		t.Fatal(err)
	}
	var fixture struct {
		RatingPerPercent map[string]float64 `json:"rating_per_percent"`
	}
	if err := json.Unmarshal(b, &fixture); err != nil {
		t.Fatal(err)
	}
	// haste is the addon's alone: the API's item lines carry no haste key.
	delete(fixture.RatingPerPercent, "haste")
	if len(fixture.RatingPerPercent) != len(ratingPerPercent) {
		t.Fatalf("ratingPerPercent = %v, fixture %v", ratingPerPercent, fixture.RatingPerPercent)
	}
	for stat, factor := range fixture.RatingPerPercent {
		if ratingPerPercent[stat] != factor {
			t.Errorf("ratingPerPercent[%q] = %v, fixture %v", stat, ratingPerPercent[stat], factor)
		}
	}
}

// Sim audit round 17: every weight is per 1%, an item's line rating points, so a hit- and
// crit-rated glove scored 10x and 14x too high. The Fury hands of 1.60.1.70755
// (items/warrior.json) under stat-weights.json warrior-fury.
func TestStatScoreReadsRatingAsPercent(t *testing.T) {
	fury := map[string]float64{"attack_power": 1, "strength": 1.9, "agility": 0.7, "hit": 32, "crit": 18, "expertise": 18}
	hands := map[string]trees.Item{
		"Gauntlets of Annihilation": {Stats: map[string]int{"strength": 35, "stamina": 15, "crit": 14, "hit": 10}},
		"Fists of the Unrelenting": {
			Stats:        map[string]int{"stamina": 26, "strength": 26, "crit": 1, "hit": 1},
			PercentStats: map[string]float64{"crit": 1, "hit": 1},
		},
		"Gloves of Enforcement": {Stats: map[string]int{"agility": 20, "strength": 28, "stamina": 6, "hit": 10}},
		"Primal Batskin Gloves": {Stats: map[string]int{"agility": 10, "stamina": 9, "hit": 20}},
	}
	want := map[string]float64{
		"Gauntlets of Annihilation": 116.5, // 35*1.9 + 14/14*18 + 10/10*32
		"Fists of the Unrelenting":  99.4,  // 26*1.9 + 1*18 + 1*32, literal percents
		"Gloves of Enforcement":     99.2,  // 28*1.9 + 20*0.7 + 10/10*32
		"Primal Batskin Gloves":     71,    // 10*0.7 + 20/10*32 (was 647)
	}
	names := make([]string, 0, len(hands))
	for name, item := range hands {
		if got := statScore(item, fury); math.Abs(got-want[name]) > 1e-9 {
			t.Errorf("%s scores %v, want %v", name, got, want[name])
		}
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool { return statScore(hands[names[i]], fury) > statScore(hands[names[j]], fury) })
	order := []string{"Gauntlets of Annihilation", "Fists of the Unrelenting", "Gloves of Enforcement", "Primal Batskin Gloves"}
	for i := range order {
		if names[i] != order[i] {
			t.Fatalf("order %v, want %v", names, order)
		}
	}
}

func TestStatScoreReadsEveryRatingStatAtItsOwnFactor(t *testing.T) {
	weights := map[string]float64{"expertise": 18, "dodge": 12, "parry": 15, "block": 5, "defense": 2}
	item := trees.Item{Stats: map[string]int{"expertise": 20, "dodge": 12, "parry": 15, "block": 5, "defense": 3}}
	if got, want := statScore(item, weights), 36.0+12+15+5+6; math.Abs(got-want) > 1e-9 {
		t.Errorf("statScore = %v, want %v", got, want)
	}
}
