// logs/engine/rating/preparation_test.go
package rating

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// preparedFight is fixtureSummary with a character sheet for the fixture
// player and the given aura tracks on them.
func preparedFight(auras ...summary.AuraTrack) summary.Summary {
	fight := fixtureSummary()
	fight.Combatants = []summary.CombatantRow{{GUID: fixturePlayer, Gear: []event.Item{{ID: 1}}}}
	fight.Auras = auras
	return fight
}

// seededOn is an aura the player's own COMBATANT_INFO listed, up the whole
// 180-second fixture fight.
func seededOn(guid string, spellID int64) summary.AuraTrack {
	return summary.AuraTrack{TargetGUID: guid, SpellID: spellID, Seeded: true,
		Segments: []summary.Segment{{StartMS: 0, EndMS: 180000}}}
}

func seededAura(spellID int64) summary.AuraTrack { return seededOn(fixturePlayer, spellID) }

// checkPlaced asserts the one raw value component placed through src.
func checkPlaced(t *testing.T, src *fakePercentiles, component string, want float64) {
	t.Helper()
	if got := src.recorded[component]; len(got) != 1 || roundTo1(got[0]) != want {
		t.Errorf("%s placed %v, want [%v]", component, got, want)
	}
}

// placedPreparation is the one value scorePreparation placed for the dps
// catalogue against fight.
func placedPreparation(t *testing.T, fight summary.Summary) []float64 {
	t.Helper()
	rc, ok := consumables.Load().For(RoleDPS)
	if !ok {
		t.Fatal("no dps catalogue")
	}
	src := &fakePercentiles{}
	_ = scorePreparation(fight, fixturePlayer, Bracket{}, &rc, src)
	return src.recorded[ComponentNamePreparation]
}

func TestPreparation_SeededFlaskEarnsFlaskWeight(t *testing.T) {
	rc, _ := consumables.Load().For(RoleDPS)
	got := placedPreparation(t, preparedFight(seededAura(17626)))
	if len(got) != 1 || got[0] != float64(rc.Weights["flask"]) {
		t.Fatalf("placed %v, want the flask weight alone (%d)", got, rc.Weights["flask"])
	}
}

func TestPreparation_FlaskDrunkAfterTheGraceEarnsNothing(t *testing.T) {
	late := summary.AuraTrack{TargetGUID: fixturePlayer, SpellID: 17626,
		Segments: []summary.Segment{{StartMS: consumables.AtPullGraceMS + 1, EndMS: 180000}}}
	if got := placedPreparation(t, preparedFight(late)); len(got) != 1 || got[0] != 0 {
		t.Fatalf("placed %v, want 0: a flask drunk mid-fight is not preparation", got)
	}
}

func TestPreparation_SeededFlaskWithoutASheetIsNotRead(t *testing.T) {
	fight := preparedFight(seededAura(17626))
	fight.Combatants = nil
	if got := placedPreparation(t, fight); len(got) != 1 || got[0] != 0 {
		t.Fatalf("placed %v, want 0: no character sheet, nothing is known present", got)
	}
}

func TestPreparation_PrePullPotionCountsAsOneUse(t *testing.T) {
	rc, _ := consumables.Load().For(RoleDPS)
	if len(rc.CombatPotion.Entries) == 0 || rc.CombatPotion.MaxUses == 0 {
		t.Fatal("the dps catalogue needs a combat potion for this case to mean anything")
	}
	potion := rc.CombatPotion.Entries[0].SpellID
	want := float64(rc.Weights["combat_potion"]) / float64(rc.CombatPotion.MaxUses)
	if got := placedPreparation(t, preparedFight(seededAura(potion))); len(got) != 1 || roundTo1(got[0]) != roundTo1(want) {
		t.Fatalf("placed %v, want %v: one use of %d", got, want, rc.CombatPotion.MaxUses)
	}
}

func TestModelVersionMovedForConsumables(t *testing.T) {
	if DefaultModelVersion != "rating-2026-10-28-consumables-at-pull" {
		t.Fatal(DefaultModelVersion)
	}
}
