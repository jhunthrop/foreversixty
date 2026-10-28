package rating

import (
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The discipline and consumables-missed readers read what the player brought through
// consumables.AtPull: the seeded aura tracks (log coach spec §13), never the
// CombatantRow.Consumables list production never fills.

// pullFight is a fight with row's character sheet (preparedCombatant) and auras on it.
func pullFight(row summary.RosterRow, auras ...summary.AuraTrack) summary.Summary {
	return summary.Summary{DurationMS: 180000, Roster: []summary.RosterRow{row},
		Combatants: []summary.CombatantRow{preparedCombatant(row.GUID)}, Auras: auras}
}

func TestConsumablesMissed_ReadsSeededAuras(t *testing.T) {
	dps, _ := consumables.Load().For("dps")
	row := summary.RosterRow{GUID: "g", Name: "Flasked"}
	fight := pullFight(row, seededTrack("g", dps.Flask[0].SpellID))
	if got := consumablesMissed(fight, row); !reflect.DeepEqual(got, []string{consumableFood}) {
		t.Fatalf("missed %v, want [food]: the flask is seeded, no food is", got)
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
