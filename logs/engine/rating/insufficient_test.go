// logs/engine/rating/insufficient_test.go
package rating

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/utility"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// TestScoreProducesAnInsufficientCardOnAWipeWithNoMechanicsTable reproduces the
// production case the coordinator found: a wipe, no curated mechanics table for
// the encounter, no validated simulator spec (no ExecutionScore). Output,
// Survival and Mechanics are correctly excluded, leaving Utility, Preparation
// and Activity -- DPS weight 15+10+5 = 30 of 100, coverage 0.3, below
// MinCoverage. The overall must not read as a judgement built on almost
// nothing.
func TestScoreProducesAnInsufficientCardOnAWipeWithNoMechanicsTable(t *testing.T) {
	const player = "Player-Shadow"
	fight := summary.Summary{
		DurationMS: 60000, EncounterID: 9999, Kill: false, // a wipe
		Roster: []summary.RosterRow{
			{GUID: player, Class: "Priest", Spec: "Shadow", Role: RoleDPS, DPS: 300, ActivityPct: 40, ActiveMS: 24000},
		},
		// No Mechanics.TableFound: the fixture never sets it, so tables.Mechanics is nil below.
	}
	utilTable := &utility.Table{Spec: "priest-shadow", Owned: []utility.Entry{
		{SpellID: 12345, Name: "Shadow Weaving", Kind: utility.Debuff, Target: "enemy", Verified: "test"},
	}}
	cat := &consumables.RoleCatalogue{
		Weights: map[string]int{"flask": 30, "food": 15, "weapon_enchant": 15, "world_buffs": 20, "combat_potion": 20},
	}
	src := &fakePercentiles{} // ok=false everywhere: every scored component uses its own absolute standard
	card := Score(fight, player, CuratedTables{Mechanics: nil, Utility: utilTable, Consumables: cat}, src, nil, DefaultModelInfo())

	outputC := componentByName(card, ComponentNameOutput)
	survivalC := componentByName(card, ComponentNameSurvival)
	mechanicsC := componentByName(card, ComponentNameMechanics)
	if !outputC.Excluded || !survivalC.Excluded || !mechanicsC.Excluded {
		t.Fatalf("output/survival/mechanics must all be excluded: output=%v survival=%v mechanics=%v",
			outputC.Excluded, survivalC.Excluded, mechanicsC.Excluded)
	}
	utilityC := componentByName(card, ComponentNameUtility)
	prepC := componentByName(card, ComponentNamePreparation)
	activityC := componentByName(card, ComponentNameActivity)
	if utilityC.Excluded || prepC.Excluded || activityC.Excluded {
		t.Fatalf("utility/preparation/activity must all be scored (not excluded): utility=%+v prep=%+v activity=%+v",
			utilityC, prepC, activityC)
	}

	if want := 0.3; card.Coverage != want {
		t.Fatalf("Coverage = %v, want %v (30 of 100 DPS weight points scored: utility 15 + preparation 10 + activity 5)",
			card.Coverage, want)
	}
	if !card.Insufficient {
		t.Fatal("Insufficient = false, want true: coverage 0.3 is below MinCoverage 0.5")
	}
	if card.Overall != 0 || card.OverallUncapped != 0 || card.OverallCapped != false {
		t.Fatalf("an insufficient card must zero Overall/OverallUncapped/OverallCapped, got %v/%v/%v",
			card.Overall, card.OverallUncapped, card.OverallCapped)
	}
	if card.InsufficientReason == "" {
		t.Fatal("InsufficientReason must name what was missing, got empty string")
	}
	for _, want := range []string{"wipe", "mechanics table"} {
		if !strings.Contains(card.InsufficientReason, want) {
			t.Errorf("InsufficientReason = %q, want it to mention %q", card.InsufficientReason, want)
		}
	}
	// Components themselves are still returned in full, scores and reasons
	// included, so the page can show what WAS measured even on an
	// insufficient card.
	if utilityC.Score < 0 || prepC.Score < 0 {
		t.Errorf("scored components must still carry real scores: utility=%v prep=%v", utilityC.Score, prepC.Score)
	}
}
