package rating

import (
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The discipline and consumables-missed readers read what the player brought through
// consumables.AtPull: the seeded aura tracks (log coach spec §13). A category is reported
// missing only when the pull is known and the catalogue's list of it is complete.

// pullFight is a fight with row's character sheet (preparedCombatant) and auras on it.
func pullFight(row summary.RosterRow, auras ...summary.AuraTrack) summary.Summary {
	return summary.Summary{DurationMS: 180000, Roster: []summary.RosterRow{row},
		Combatants: []summary.CombatantRow{preparedCombatant(row.GUID)}, Auras: auras}
}

func TestConsumablesMissed_ReadsSeededAuras(t *testing.T) {
	dps, _ := consumables.Load().For("dps")
	row := summary.RosterRow{GUID: "g", Name: "Fed"}
	fight := pullFight(row, seededTrack("g", dps.Food[0].SpellID))
	if got := consumablesMissed(fight, row); !reflect.DeepEqual(got, []string{consumableFlask}) {
		t.Fatalf("missed %v, want [flask]: the food is seeded, no flask is", got)
	}
	flasked := pullFight(row, seededTrack("g", dps.Flask[0].SpellID))
	if got := consumablesMissed(flasked, row); got == nil || len(got) != 0 {
		t.Fatalf("missed %#v, want an empty list: the flask is seeded", got)
	}
}

// offListFood is a food buff the catalogue does not list (Smoked Desert Dumplings' Well
// Fed): the food list is incomplete, so it must never read as "food missed".
const offListFood = 24799

func requireIncompleteFood(t *testing.T) {
	t.Helper()
	if disciplineCatalogue.IsComplete(consumables.CategoryFood) {
		t.Fatal("the catalogue's food list is now complete: these cases assume an incomplete one")
	}
	for _, e := range disciplineCatalogue.Food {
		if e.SpellID == offListFood {
			t.Fatalf("%d is now listed: pick another off-list food", offListFood)
		}
	}
}

func TestConsumablesMissed_OffListFoodIsNotMissed(t *testing.T) {
	requireIncompleteFood(t)
	dps, _ := consumables.Load().For("dps")
	row := summary.RosterRow{GUID: "g", Name: "Dumplings"}
	fight := pullFight(row, seededTrack("g", dps.Flask[0].SpellID), seededTrack("g", offListFood))
	if got := consumablesMissed(fight, row); len(got) != 0 {
		t.Fatalf("missed %v, want nothing: an off-list food is not a missed food", got)
	}
}

func TestFightDiscipline_OffListFoodIsPrepared(t *testing.T) {
	requireIncompleteFood(t)
	dps, _ := consumables.Load().For("dps")
	row := summary.RosterRow{GUID: "g", Name: "Dumplings", Role: "dps"}
	fight := pullFight(row, seededTrack("g", dps.Flask[0].SpellID), seededTrack("g", offListFood))
	if got := fightDiscipline(fight, row); got == nil || !*got {
		t.Fatalf("discipline %v, want true: flask, an off-list food and every enchant", got)
	}
}

func TestConsumablesMissed_FlaskAppliedInTheGraceIsBrought(t *testing.T) {
	dps, _ := consumables.Load().For("dps")
	row := summary.RosterRow{GUID: "g", Name: "Late"}
	onPull := summary.AuraTrack{TargetGUID: "g", SpellID: dps.Flask[0].SpellID,
		Segments: []summary.Segment{{StartMS: consumables.AtPullGraceMS, EndMS: 180000}}}
	fight := pullFight(row, onPull, seededTrack("g", dps.Food[0].SpellID))
	if got := consumablesMissed(fight, row); len(got) != 0 {
		t.Fatalf("missed %v, want nothing: a flask drunk on the pull counts", got)
	}
}

func TestFightDiscipline_SeededFlaskAndFoodAndEnchants(t *testing.T) {
	dps, _ := consumables.Load().For("dps")
	row := summary.RosterRow{GUID: "g", Name: "Prepared", Role: "dps"}
	fight := pullFight(row, seededTrack("g", dps.Flask[0].SpellID), seededTrack("g", dps.Food[0].SpellID))
	if got := fightDiscipline(fight, row); got == nil || !*got {
		t.Fatalf("discipline %v, want true: flask and food seeded, every enchant slot enchanted", got)
	}
}

func TestFightDiscipline_NoSheetIsUnknown(t *testing.T) {
	dps, _ := consumables.Load().For("dps")
	row := summary.RosterRow{GUID: "g", Name: "Unseen"}
	fight := summary.Summary{Roster: []summary.RosterRow{row},
		Auras: []summary.AuraTrack{seededTrack("g", dps.Flask[0].SpellID), seededTrack("g", dps.Food[0].SpellID)}}
	if got := fightDiscipline(fight, row); got != nil {
		t.Fatalf("discipline %v, want nil: no combatant row, the log cannot say", *got)
	}
	if got := consumablesMissed(fight, row); got == nil || len(got) != 0 {
		t.Fatalf("missed %#v, want an empty list: unknown is never missing", got)
	}
}
