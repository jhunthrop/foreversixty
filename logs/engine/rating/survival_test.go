// logs/engine/rating/survival_test.go
package rating

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// avoidableRoleFixture is a 180-second fight with one Role-tagged avoidable
// mechanic ("Cleave", role tank) that hit the player under test for 4000
// total damage between 100,000ms and 104,000ms. otherTank, when non-empty,
// is added to the roster as a second tank so an assignment can name them.
func avoidableRoleFixture(otherTank string) summary.Summary {
	const player = "Player-Tank"
	roster := []summary.RosterRow{{GUID: player, Role: RoleTank}}
	if otherTank != "" {
		roster = append(roster, summary.RosterRow{GUID: otherTank, Role: RoleTank})
	}
	return summary.Summary{
		DurationMS: 180000,
		Roster:     roster,
		Mechanics: summary.MechanicsBlock{TableFound: true, Rows: []summary.MechanicRow{
			{SpellID: 19998, Name: "Cleave", Kind: mechanics.Avoidable, Role: RoleTank,
				Players: []summary.MechanicHit{{GUID: player, Hits: 2, Damage: 4000, FirstMS: 100000, LastMS: 104000}}},
		}},
	}
}

const avoidableFixturePlayer = "Player-Tank"
const avoidableFixtureOtherTank = "Player-OtherTank"

// TestDeathScoreForNamesAnUnclassifiedDeathWithNoKillingBlow proves a death
// the log never recorded a killing blow for still produces a reader-facing
// Moment, not an empty SpellName paired with SpellID 0 (whole-branch
// review LOW finding).
func TestDeathScoreForNamesAnUnclassifiedDeathWithNoKillingBlow(t *testing.T) {
	const player = "Player-Unknown"
	fight := summary.Summary{
		DurationMS: 100000,
		Deaths:     []summary.Death{{GUID: player, AtMS: 50000, KillingBlow: nil}},
	}
	table := &mechanics.Table{EncounterID: 1, Name: "Test"}
	_, moments := deathScoreFor(fight, player, table)
	if len(moments) != 1 {
		t.Fatalf("moments = %+v, want exactly one", moments)
	}
	m := moments[0]
	if m.SpellName == "" {
		t.Fatalf("Moment.SpellName is empty, want a player-facing name such as %q", unknownDeathCause)
	}
	if m.SpellName != unknownDeathCause {
		t.Errorf("Moment.SpellName = %q, want %q", m.SpellName, unknownDeathCause)
	}
	if m.SpellID != 0 {
		t.Errorf("Moment.SpellID = %d, want 0 (no killing blow spell id exists to report)", m.SpellID)
	}
}

func TestAvoidableDamagePerSecondExcusesWithNoAssignmentAtAll(t *testing.T) {
	fight := avoidableRoleFixture("")
	got := avoidableDamagePerSecond(fight, avoidableFixturePlayer, RoleTank, nil)
	if got != 0 {
		t.Fatalf("avoidableDamagePerSecond = %v, want 0 (no assignment at all for this role: excused, "+
			"we cannot tell the main tank doing their job from the off tank standing in it)", got)
	}
}

func TestAvoidableDamagePerSecondExcusesWhenAssignedToThisPlayer(t *testing.T) {
	fight := avoidableRoleFixture(avoidableFixtureOtherTank)
	assignments := []Assignment{
		{PlayerKey: avoidableFixturePlayer, Job: "cleave duty", FromMS: 90000, ToMS: 110000},
	}
	got := avoidableDamagePerSecond(fight, avoidableFixturePlayer, RoleTank, assignments)
	if got != 0 {
		t.Fatalf("avoidableDamagePerSecond = %v, want 0 (assigned to this player at the time of the hit: excused)", got)
	}
}

func TestAvoidableDamagePerSecondCountsInFullWhenAssignedToTheOtherTank(t *testing.T) {
	fight := avoidableRoleFixture(avoidableFixtureOtherTank)
	assignments := []Assignment{
		{PlayerKey: avoidableFixtureOtherTank, Job: "cleave duty", FromMS: 90000, ToMS: 110000},
	}
	got := avoidableDamagePerSecond(fight, avoidableFixturePlayer, RoleTank, assignments)
	want := 4000.0 / 180.0
	if got != want {
		t.Fatalf("avoidableDamagePerSecond = %v, want %v (assignment names the other tank for this window: "+
			"the hit counts in full against this player, spec §1.6's \"stood in a cleave meant for the other tank\")",
			got, want)
	}
}

func TestAvoidableDamagePerSecondExcusesWhenTheOnlyAssignmentDoesNotOverlapTheHit(t *testing.T) {
	fight := avoidableRoleFixture(avoidableFixtureOtherTank)
	// An assignment for the other tank exists, but for a completely
	// different, unrelated stretch of the fight (0-50000ms) that does not
	// overlap this hit's own window (100000-104000ms): irrelevant evidence,
	// the same as no assignment existing at all for this mechanic.
	assignments := []Assignment{
		{PlayerKey: avoidableFixtureOtherTank, Job: "unrelated duty", FromMS: 0, ToMS: 50000},
	}
	got := avoidableDamagePerSecond(fight, avoidableFixturePlayer, RoleTank, assignments)
	if got != 0 {
		t.Fatalf("avoidableDamagePerSecond = %v, want 0 (the only assignment for this role does not cover "+
			"the hit's own time window, which is the same as no assignment existing for this mechanic at all)", got)
	}
}
