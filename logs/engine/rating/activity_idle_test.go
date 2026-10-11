// logs/engine/rating/activity_idle_test.go
package rating

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// ActivityIdleMS is scoreActivity's own measure as a duration: the scored seconds the
// player did nothing, windows excluded.
func TestActivityIdleMSIsTheActivityComponentsOwnMeasure(t *testing.T) {
	const player = "Player-Idle"
	// Active 0-59 and 65-169; the downtime window is [60000, 65000); idle 170-179.
	actives := rangeSet(0, 60)
	for k := range rangeSet(65, 170) {
		actives[k] = true
	}
	fight := summary.Summary{
		DurationMS: 180000,
		Roster:     []summary.RosterRow{{GUID: player, ActiveMS: 1000}}, // the fallback, unused with windows
		DamageDone: []summary.Actor{{GUID: player, Series: activitySeries(actives)}},
		Auras: []summary.AuraTrack{
			{SpellID: 999, Name: "Detonation", Segments: []summary.Segment{{StartMS: 60000, EndMS: 60000}}},
		},
	}
	table := &mechanics.Table{EncounterID: 1, Name: "Test", Downtime: []mechanics.Downtime{
		{Trigger: mechanics.PhaseStart{SpellID: 999, On: mechanics.OnAuraApplied}, DurationMS: 5000},
	}}
	if got := ActivityIdleMS(fight, player, table, nil); got != 10000 {
		t.Fatalf("idle = %d ms, want 10000 (175 scored seconds, 165 active)", got)
	}
	// An assignment window over the idle tail takes it out of the count.
	tail := []Assignment{{PlayerKey: player, FromMS: 170000, ToMS: 180000}}
	if got := ActivityIdleMS(fight, player, table, tail); got != 0 {
		t.Fatalf("idle with the tail assigned = %d ms, want 0", got)
	}
	// With no window at all: the time alive less ActiveMS, here to the death at 150s.
	fight.Deaths = []summary.Death{{GUID: player, AtMS: 150000}}
	fight.Roster[0].ActiveMS = 120000
	if got := ActivityIdleMS(fight, player, nil, nil); got != 30000 {
		t.Fatalf("idle with no windows = %d ms, want 30000", got)
	}
	if got := ActivityIdleMS(fight, "someone-else", table, nil); got != 0 {
		t.Fatalf("a player off the roster = %d, want 0", got)
	}
}
