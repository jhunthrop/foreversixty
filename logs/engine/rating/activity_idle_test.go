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

// IdleWindows is ActivityIdleMS's spans: with a downtime window in force, the idle
// seconds outside it, grouped into runs, longest first, summing to ActivityIdleMS.
func TestIdleWindowsLongestFirst(t *testing.T) {
	const player = "Player-Idle"
	// Active 0-19, 25-59, 65-169; idle 20-24 (5 s) and 170-179 (10 s); 60-64 is downtime.
	actives := rangeSet(0, 20)
	for k := range rangeSet(25, 60) {
		actives[k] = true
	}
	for k := range rangeSet(65, 170) {
		actives[k] = true
	}
	fight := summary.Summary{
		DurationMS: 180000,
		Roster:     []summary.RosterRow{{GUID: player, ActiveMS: 1000}},
		DamageDone: []summary.Actor{{GUID: player, Series: activitySeries(actives)}},
		Auras: []summary.AuraTrack{
			{SpellID: 999, Name: "Detonation", Segments: []summary.Segment{{StartMS: 60000, EndMS: 60000}}},
		},
	}
	table := &mechanics.Table{EncounterID: 1, Name: "Test", Downtime: []mechanics.Downtime{
		{Trigger: mechanics.PhaseStart{SpellID: 999, On: mechanics.OnAuraApplied}, DurationMS: 5000},
	}}
	got := IdleWindows(fight, player, table, nil)
	want := []Window{{FromMS: 170000, ToMS: 180000}, {FromMS: 20000, ToMS: 25000}}
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Fatalf("IdleWindows = %+v, want %+v", got, want)
	}
	var sum int64
	for _, w := range got {
		sum += w.ToMS - w.FromMS
	}
	if idle := ActivityIdleMS(fight, player, table, nil); sum != idle {
		t.Fatalf("the windows sum to %d ms, ActivityIdleMS is %d", sum, idle)
	}
	if w := IdleWindows(fight, "someone-else", table, nil); w != nil {
		t.Fatalf("a player off the roster has no idle windows, got %+v", w)
	}
}

func TestDowntimeWindowsAndTimeAliveAreExported(t *testing.T) {
	fight := summary.Summary{
		DurationMS: 180000,
		Deaths:     []summary.Death{{GUID: "p", AtMS: 62000}},
		Auras:      []summary.AuraTrack{{SpellID: 999, Segments: []summary.Segment{{StartMS: 60000, EndMS: 60000}}}},
	}
	table := &mechanics.Table{Downtime: []mechanics.Downtime{
		{Trigger: mechanics.PhaseStart{SpellID: 999, On: mechanics.OnAuraApplied}, DurationMS: 5000},
	}}
	alive := TimeAliveMS(fight, "p")
	if alive != 62000 {
		t.Fatalf("TimeAliveMS = %d, want 62000", alive)
	}
	got := DowntimeWindows(fight, table, alive)
	if len(got) != 1 || got[0] != (Window{FromMS: 60000, ToMS: 62000}) {
		t.Fatalf("DowntimeWindows = %+v, want one window clipped to the death", got)
	}
	marks := []Assignment{{PlayerKey: "p", FromMS: 50000, ToMS: 70000}, {PlayerKey: "q", FromMS: 0, ToMS: 1000}}
	if got := AssignmentWindows(marks, "p", alive); len(got) != 1 || got[0] != (Window{FromMS: 50000, ToMS: 62000}) {
		t.Fatalf("AssignmentWindows = %+v, want p's one window clipped to the death", got)
	}
}
