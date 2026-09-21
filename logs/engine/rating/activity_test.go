// logs/engine/rating/activity_test.go
package rating

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// activitySeries builds a 180-entry per-second series: value 100 for
// every index in actives, 0 elsewhere -- the shape Actor.Series already
// has (one bucket per second).
func activitySeries(actives map[int]bool) []int64 {
	out := make([]int64, 180)
	for i := range out {
		if actives[i] {
			out[i] = 100
		}
	}
	return out
}

func rangeSet(from, to int) map[int]bool {
	m := map[int]bool{}
	for i := from; i < to; i++ {
		m[i] = true
	}
	return m
}

func TestScoreActivityExcludesACuratedDowntimeWindowFromBothNumeratorAndDenominator(t *testing.T) {
	const player = "Player-Active"
	// Active seconds 0-59 and 65-169 (165 total); idle 60-64 (inside the
	// downtime window, so it wouldn't matter either way) and idle 170-179.
	actives := rangeSet(0, 60)
	for k := range rangeSet(65, 170) {
		actives[k] = true
	}
	fight := summary.Summary{
		DurationMS: 180000,
		Roster:     []summary.RosterRow{{GUID: player, ActivityPct: 50}}, // fallback value, must not be used
		DamageDone: []summary.Actor{{GUID: player, Series: activitySeries(actives)}},
		// The downtime trigger: an enemy's own aura (spell 999) opens at
		// 60,000ms -- the same aura_applied shape spec §3.4's Garr example
		// uses -- giving a 5-second downtime window [60000, 65000).
		Auras: []summary.AuraTrack{
			{SpellID: 999, Name: "Detonation", Segments: []summary.Segment{{StartMS: 60000, EndMS: 60000}}},
		},
	}
	table := &mechanics.Table{EncounterID: 1, Name: "Test", Downtime: []mechanics.Downtime{
		{Trigger: mechanics.PhaseStart{SpellID: 999, On: mechanics.OnAuraApplied}, DurationMS: 5000},
	}}
	src := &fakePercentiles{} // ok=false: forces the absolute (self) standard
	c := scoreActivity(fight, player, Bracket{}, table, nil, src)

	if c.Excluded {
		t.Fatalf("Activity excluded (%s), want scored (Activity is always computable)", c.Reason)
	}
	want := roundTo(165.0/175.0*100, 2)
	if c.Score != want {
		t.Fatalf("Score = %v, want %v (165 active seconds of 175 non-downtime seconds)", c.Score, want)
	}
	if len(src.recorded[ComponentNameActivity]) != 1 || roundTo(src.recorded[ComponentNameActivity][0], 2) != want {
		t.Fatalf("recorded activity placements = %v, want [%v]", src.recorded[ComponentNameActivity], want)
	}
}

func TestScoreActivityExcludesAnAssignmentWindowForThisPlayerOnly(t *testing.T) {
	const player = "Player-Active"
	const other = "Player-Other"
	actives := rangeSet(0, 180) // fully active every second
	fight := summary.Summary{
		DurationMS: 180000,
		Roster:     []summary.RosterRow{{GUID: player}, {GUID: other}},
		DamageDone: []summary.Actor{{GUID: player, Series: activitySeries(actives)}},
	}
	// An officer-marked assignment (spec §2) takes this player out of the
	// active pool for 10 of the 180 seconds -- their "free time" denominator
	// shrinks by exactly that window, per §2's Output/Activity exemption.
	assignments := []Assignment{{PlayerKey: player, Job: "kiting", FromMS: 100000, ToMS: 110000}}
	src := &fakePercentiles{}
	c := scoreActivity(fight, player, Bracket{}, nil, assignments, src)

	want := roundTo(170.0/170.0*100, 2) // 170s active out of 170s free time: still 100%
	if c.Score != want {
		t.Fatalf("Score = %v, want %v", c.Score, want)
	}

	// The OTHER player has no assignment at all, so their own Activity
	// still falls back to whole-fight ActivityPct rather than going
	// through the windowed path -- confirming assignments are filtered per
	// player, not applied to the whole roster.
	otherFight := fight
	otherFight.Roster = []summary.RosterRow{{GUID: other, ActivityPct: 42}}
	otherC := scoreActivity(otherFight, other, Bracket{}, nil, assignments, src)
	if otherC.Score != 42 {
		t.Fatalf("other player's Score = %v, want 42 (the fallback ActivityPct, unaffected by player's own assignment)", otherC.Score)
	}
}

func TestScoreActivityFallsBackToWholeFightActivityPctWithNoWindowsAtAll(t *testing.T) {
	const player = "Player-Active"
	fight := summary.Summary{
		DurationMS: 180000,
		Roster:     []summary.RosterRow{{GUID: player, ActivityPct: 73.5}},
	}
	src := &fakePercentiles{}
	c := scoreActivity(fight, player, Bracket{}, nil, nil, src)
	if c.Score != 73.5 {
		t.Fatalf("Score = %v, want 73.5 (whole-fight ActivityPct fallback, no curated downtime or assignments)", c.Score)
	}
}

func TestScoreActivityWithNoRosterRowFallsBackToZero(t *testing.T) {
	fight := summary.Summary{DurationMS: 180000}
	src := &fakePercentiles{}
	c := scoreActivity(fight, "does-not-exist", Bracket{}, nil, nil, src)
	if c.Score != 0 {
		t.Fatalf("Score = %v, want 0 for a player with no roster row", c.Score)
	}
}

func TestTriggerInstantsMSSupportsAuraAndCastTriggersOnly(t *testing.T) {
	fight := summary.Summary{
		Auras: []summary.AuraTrack{
			{SpellID: 100, Segments: []summary.Segment{{StartMS: 1000, EndMS: 4000}, {StartMS: 9000, EndMS: 9500}}},
		},
		Casts: []summary.CastRow{
			{SpellID: 200, Sequence: []int64{500, 6000}},
		},
	}

	applied := triggerInstantsMS(fight, mechanics.PhaseStart{SpellID: 100, On: mechanics.OnAuraApplied})
	if len(applied) != 2 || applied[0] != 1000 || applied[1] != 9000 {
		t.Fatalf("aura_applied instants = %v, want [1000 9000]", applied)
	}

	removed := triggerInstantsMS(fight, mechanics.PhaseStart{SpellID: 100, On: mechanics.OnAuraRemoved})
	if len(removed) != 2 || removed[0] != 4000 || removed[1] != 9500 {
		t.Fatalf("aura_removed instants = %v, want [4000 9500]", removed)
	}

	casts := triggerInstantsMS(fight, mechanics.PhaseStart{SpellID: 200, On: mechanics.OnCastSuccess})
	if len(casts) != 2 || casts[0] != 500 || casts[1] != 6000 {
		t.Fatalf("cast_success instants = %v, want [500 6000]", casts)
	}

	// RULING R7: cast_start and a health_pct trigger have no equivalent
	// data surviving in Summary and must produce no instants, not a guess.
	if got := triggerInstantsMS(fight, mechanics.PhaseStart{SpellID: 200, On: mechanics.OnCastStart}); got != nil {
		t.Errorf("cast_start instants = %v, want nil (unsupported, documented)", got)
	}
	if got := triggerInstantsMS(fight, mechanics.PhaseStart{HealthPct: 50}); got != nil {
		t.Errorf("health_pct instants = %v, want nil (unsupported, documented)", got)
	}
	if got := triggerInstantsMS(fight, mechanics.PhaseStart{}); got != nil {
		t.Errorf("a trigger with no spell id must produce no instants, got %v", got)
	}
}

func TestClipWindowsDropsAndClampsOutOfRangeWindows(t *testing.T) {
	windows := []window{
		{fromMS: -5000, toMS: 3000},    // clamped to [0, 3000)
		{fromMS: 200000, toMS: 210000}, // entirely past aliveMS: dropped
		{fromMS: 170000, toMS: 190000}, // clamped to [170000, 180000)
		{fromMS: 50000, toMS: 50000},   // zero-width: dropped
		{fromMS: -10000, toMS: -5000},  // entirely negative: dropped
	}
	got := clipWindows(windows, 180000)
	if len(got) != 2 {
		t.Fatalf("clipWindows = %+v, want 2 surviving windows", got)
	}
	if got[0].fromMS != 0 || got[0].toMS != 3000 {
		t.Errorf("window 0 = %+v, want [0,3000)", got[0])
	}
	if got[1].fromMS != 170000 || got[1].toMS != 180000 {
		t.Errorf("window 1 = %+v, want [170000,180000)", got[1])
	}
}
