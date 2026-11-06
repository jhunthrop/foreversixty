package rating

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const presetPrepEncounter = 667

// preparedButFlask is a fight in which the row's player brought one of every food and
// weapon enchant on the role's list, every listed world buff and the role's full count of
// combat potions, but no flask.
func preparedButFlask(t *testing.T, row summary.RosterRow) summary.Summary {
	t.Helper()
	rc, ok := consumables.Load().For(row.Role)
	if !ok {
		t.Fatalf("no %s catalogue", row.Role)
	}
	sum := summary.Summary{EncounterID: presetPrepEncounter, Kill: true, DurationMS: 120_000,
		Roster:     []summary.RosterRow{row},
		Combatants: []summary.CombatantRow{{GUID: row.GUID, Gear: []event.Item{{ID: 1}}}}}
	seeded := append([]consumables.Entry{rc.Food[0], rc.WeaponEnchant[0]}, rc.WorldBuffs...)
	for _, e := range seeded {
		sum.Auras = append(sum.Auras, summary.AuraTrack{TargetGUID: row.GUID, SpellID: e.SpellID, Seeded: true,
			Segments: []summary.Segment{{StartMS: 0, EndMS: sum.DurationMS}}})
	}
	potion := rc.CombatPotion.Entries[0]
	sum.Casts = []summary.CastRow{{GUID: row.GUID, OwnerGUID: row.GUID, SpellID: potion.SpellID,
		Succeeded: rc.CombatPotion.MaxUses}}
	return sum
}

// preparationOf is the Preparation value the Forever Score places for row's player, read
// through the curated tables the rating hands the engine.
func preparationOf(t *testing.T, row summary.RosterRow) float64 {
	t.Helper()
	cat := curatedTablesFor(row, presetPrepEncounter).Consumables
	v, ok := ratingengine.PreparationValue(preparedButFlask(t, row), row.GUID, cat)
	if !ok {
		t.Fatalf("%s %s at %d: no preparation value", row.Class, row.Spec, row.Level)
	}
	return v
}

func presetRow(class, spec string, level int64) summary.RosterRow {
	return summary.RosterRow{GUID: "g1", Name: "Prepared", Class: class, Spec: spec, Role: "dps", Level: level}
}

// A spec whose raid preset carries no flask at the player's level loses no points for it.
func TestPreparation_PresetWithoutAFlaskCostsNoFlaskPoints(t *testing.T) {
	for _, row := range []summary.RosterRow{
		presetRow("Druid", "Feral", 60),
		presetRow("Hunter", "Marksmanship", 60),
		presetRow("Druid", "Balance", 30),
	} {
		if got := preparationOf(t, row); got != 100 {
			t.Errorf("%s %s at %d: Preparation %v, want 100 with everything but the flask its preset leaves out",
				row.Class, row.Spec, row.Level, got)
		}
	}
}

// A spec whose preset carries a flask still pays for going without.
func TestPreparation_FireMageAt60StillPaysForNoFlask(t *testing.T) {
	if got := preparationOf(t, presetRow("Mage", "Fire", 60)); got >= 100 {
		t.Errorf("Fire mage at 60 without a flask: Preparation %v, want under 100", got)
	}
}

// A log with no level keeps the role's expectation.
func TestPreparation_NoLevelKeepsTheRoleExpectation(t *testing.T) {
	if got := preparationOf(t, presetRow("Druid", "Feral", 0)); got >= 100 {
		t.Errorf("feral with no level: Preparation %v, want the role's flask cost", got)
	}
}
