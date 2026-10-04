package main

import (
	"math/rand"
	"testing"
)

func TestBuildFightMetricsOneRowPerCharacter(t *testing.T) {
	roster := mockRoster()
	keys := make([]string, len(roster))
	for i, c := range roster {
		keys[i] = c.Name
	}
	plan := fightPlan{Name: "Onyxia", EncounterID: encOnyxia, Kill: true, DurationMS: 240_000}
	rows := buildFightMetrics(roster, keys, plan, rand.New(rand.NewSource(1)))
	if len(rows) != len(roster) {
		t.Fatalf("len(rows) = %d, want %d", len(rows), len(roster))
	}
	for i, row := range rows {
		if row.Role != roster[i].Role {
			t.Errorf("row %d role = %s, want %s", i, row.Role, roster[i].Role)
		}
		if row.Role == roleHealer && row.MetricHPS == nil {
			t.Errorf("healer %s has no hps", row.PlayerName)
		}
		if row.Role != roleHealer && row.MetricHPS != nil {
			t.Errorf("non-healer %s has an hps", row.PlayerName)
		}
		if row.MetricDPS <= 0 {
			t.Errorf("%s: non-positive dps %v", row.PlayerName, row.MetricDPS)
		}
	}
}

func TestBuildFightMetricsTrashNeverHasDeaths(t *testing.T) {
	roster := mockRoster()
	keys := make([]string, len(roster))
	plan := fightPlan{Name: "Trash", Kill: true, Trash: true, DurationMS: 60_000}
	for seed := int64(0); seed < 20; seed++ {
		rows := buildFightMetrics(roster, keys, plan, rand.New(rand.NewSource(seed)))
		for _, row := range rows {
			if row.Deaths != 0 {
				t.Fatalf("seed %d: trash fight has a death (%s)", seed, row.PlayerName)
			}
		}
	}
}

func TestBuildFightMetricsWipesCanHaveDeaths(t *testing.T) {
	roster := mockRoster()
	keys := make([]string, len(roster))
	plan := fightPlan{Name: "Onyxia", EncounterID: encOnyxia, Kill: false, DurationMS: 300_000}
	anyDeath := false
	for seed := int64(0); seed < 20; seed++ {
		rows := buildFightMetrics(roster, keys, plan, rand.New(rand.NewSource(seed)))
		for _, row := range rows {
			if row.Deaths > 0 {
				anyDeath = true
			}
		}
	}
	if !anyDeath {
		t.Fatal("no deaths across 20 wipe-fight samples")
	}
}
