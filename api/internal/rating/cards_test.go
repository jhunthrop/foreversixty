package rating

import (
	"encoding/json"
	"testing"
	"time"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func sampleCard() ratingengine.Card {
	pct := 71.2
	return ratingengine.Card{
		Overall: 70, OverallUncapped: 70, OverallCapped: false,
		Basis: "percentile", ModelVersion: ratingengine.DefaultModelVersion, KillTimeBand: "typical",
		Components: [6]ratingengine.Component{
			{Name: ratingengine.ComponentNameOutput, Score: 71, Weight: 35, Basis: "percentile",
				Percentile: &pct, BracketN: 142, Excluded: false},
			{Name: ratingengine.ComponentNameSurvival, Score: 76, Weight: 15, Basis: "percentile",
				Excluded: false, Moments: []ratingengine.Moment{
					{Kind: "death", AtMS: 140000, SpellID: 19712, SpellName: "Arcane Explosion",
						Avoidable: true, Anchor: "death-g1-140000"},
				}},
			{Name: ratingengine.ComponentNameMechanics, Excluded: true, Reason: ratingengine.ReasonNoMechanicsTable},
			{Name: ratingengine.ComponentNameUtility, Score: 62, Weight: 15, Basis: "percentile"},
			{Name: ratingengine.ComponentNamePreparation, Score: 95, Weight: 10, Basis: "absolute"},
			{Name: ratingengine.ComponentNameActivity, Score: 80, Weight: 5, Basis: "percentile"},
		},
	}
}

func TestToComponentDTOsMatchesTheSpecFieldNames(t *testing.T) {
	dtos := toComponentDTOs(sampleCard().Components)
	b, err := json.Marshal(dtos[0])
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]any
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{"name", "score", "weight", "basis", "percentile", "bracket_n", "excluded", "moments"} {
		if _, ok := got[field]; !ok {
			t.Errorf("component JSON missing field %q: %s", field, b)
		}
	}
}

func TestToComponentDTOsNullsScoreWhenExcluded(t *testing.T) {
	dtos := toComponentDTOs(sampleCard().Components)
	mechanics := dtos[2]
	b, _ := json.Marshal(mechanics)
	var got map[string]any
	_ = json.Unmarshal(b, &got)
	if got["score"] != nil {
		t.Errorf("excluded component's score = %v, want null", got["score"])
	}
	if got["reason"] != ratingengine.ReasonNoMechanicsTable {
		t.Errorf("reason = %v, want %q", got["reason"], ratingengine.ReasonNoMechanicsTable)
	}
}

func TestNewCardRowRoundTripsThroughJSON(t *testing.T) {
	meta := fightMeta{
		ReportID: "r1", FightIndex: 3, EncounterID: 667, Difficulty: 3, Size: 40,
		DurationMS: 180000, Kill: true, KillTimeBand: "typical", FoughtAt: time.Date(2026, 12, 9, 0, 0, 0, 0, time.UTC),
	}
	row := summary.RosterRow{GUID: "g1", Name: "Simfury", Class: "Warrior", Spec: "Fury", Role: "dps"}
	cr := newCardRow(meta, row, sampleCard())
	if cr.PlayerKey != "" {
		t.Error("newCardRow must not invent a player key - the caller (Store) computes it from region/ruleset")
	}
	var decoded []componentDTO
	if err := json.Unmarshal(cr.Components, &decoded); err != nil {
		t.Fatalf("Components is not valid JSON: %v", err)
	}
	if len(decoded) != 6 {
		t.Fatalf("decoded %d components, want 6", len(decoded))
	}
}
