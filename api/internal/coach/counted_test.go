package coach

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// countedFury is furyInput over a 60 s fight with one Bloodthirst at 4 s as the warrior's
// only cast, and no downtime table.
func countedFury(t *testing.T) Input {
	t.Helper()
	in := furyInput(t)
	in.Fight = withCasts(withoutCasts(withDuration(in.Fight, 60_000), furyGUID), furyGUID, 23881, "Bloodthirst", 4_000)
	in.Tables.Mechanics = nil
	return in
}

func TestCountedTime_ExcludesDowntimeDeathAndPreFirstCast(t *testing.T) {
	in := countedFury(t)
	in.Fight = withDeath(in.Fight, furyGUID, 50_000, 0, "Melee")
	got := CountedTime(in)
	want := Counted{AliveMS: 50_000, CountedMS: 46_000, FirstCastMS: 4_000}
	if got.AliveMS != want.AliveMS || got.CountedMS != want.CountedMS || got.FirstCastMS != want.FirstCastMS {
		t.Fatalf("CountedTime = %+v, want %+v", got, want)
	}
	if len(got.Excluded) != 1 || got.Excluded[0] != (ratingengine.Window{FromMS: 0, ToMS: 4_000}) {
		t.Fatalf("Excluded = %+v, want the time before the first cast", got.Excluded)
	}
}

func TestCountedTime_LeavesOutCuratedDowntime(t *testing.T) {
	in := countedFury(t)
	// A downtime window [20 s, 30 s) triggered by an aura applied at 20 s, and one
	// [2 s, 6 s) that overlaps the time before the first cast.
	in.Fight = withAura(in.Fight, "Creature-0-1", 999, "Lift Off", nil, span(20_000, 21_000))
	in.Fight = withAura(in.Fight, "Creature-0-1", 998, "Deep Breath", nil, span(2_000, 3_000))
	in.Tables.Mechanics = &mechanics.Table{Downtime: []mechanics.Downtime{
		{Trigger: mechanics.PhaseStart{SpellID: 999, On: mechanics.OnAuraApplied}, DurationMS: 10_000},
		{Trigger: mechanics.PhaseStart{SpellID: 998, On: mechanics.OnAuraApplied}, DurationMS: 4_000},
	}}
	got := CountedTime(in)
	if got.CountedMS != 60_000-6_000-10_000 {
		t.Fatalf("CountedMS = %d, want 44000 (60 s less [0, 6 s) and [20 s, 30 s))", got.CountedMS)
	}
	want := []ratingengine.Window{{FromMS: 0, ToMS: 6_000}, {FromMS: 20_000, ToMS: 30_000}}
	if len(got.Excluded) != 2 || got.Excluded[0] != want[0] || got.Excluded[1] != want[1] {
		t.Fatalf("Excluded = %+v, want %+v merged", got.Excluded, want)
	}
}

// §6.1: the windows ActivityIdleMS subtracts include the officer's assignment windows; a
// mark on another player sets nothing aside for this one.
func TestCountedTime_LeavesOutAssignmentWindows(t *testing.T) {
	in := countedFury(t)
	in.Tables.Assignments = []ratingengine.Assignment{
		{PlayerKey: furyGUID, Job: "decurse", FromMS: 30_000, ToMS: 40_000},
		{PlayerKey: priestGUID, Job: "dispel", FromMS: 10_000, ToMS: 20_000},
	}
	got := CountedTime(in)
	if got.CountedMS != 60_000-4_000-10_000 {
		t.Fatalf("CountedMS = %d, want 46000 (60 s less [0, 4 s) and the mark's [30 s, 40 s))", got.CountedMS)
	}
	want := []ratingengine.Window{{FromMS: 0, ToMS: 4_000}, {FromMS: 30_000, ToMS: 40_000}}
	if len(got.Excluded) != 2 || got.Excluded[0] != want[0] || got.Excluded[1] != want[1] {
		t.Fatalf("Excluded = %+v, want %+v", got.Excluded, want)
	}
}

// Ruling B1: F3's one measure is IdleWindows. The total it states is IdleTotalMS, the sum
// of the windows, and the title's longest window is the first. With no downtime or
// assignment window the rating's ActivityIdleMS (alive less the roster's ActiveMS) is a
// different figure, which F3 never states.
func TestIdleTotalMS_IsTheSumOfIdleWindows(t *testing.T) {
	in := countedFury(t)
	series := repeat(100, 60)
	for _, sec := range []int{10, 11, 12, 30, 31, 32, 33, 34} { // idle 0:10–0:13 and 0:30–0:35
		series[sec] = 0
	}
	in.Fight.DamageDone = []summary.Actor{{GUID: furyGUID, Series: series}}
	in.Fight.Healing = nil // the golden's own heals by the warrior would mark seconds active
	in.Fight = mapRosterRow(in.Fight, furyGUID, func(r summary.RosterRow) summary.RosterRow { r.ActiveMS = 58_000; return r })
	windows := ratingengine.IdleWindows(in.Fight, furyGUID, nil, nil)
	if len(windows) != 2 || windows[0] != (ratingengine.Window{FromMS: 30_000, ToMS: 35_000}) {
		t.Fatalf("IdleWindows = %+v, want the 5 s window first", windows)
	}
	if got := IdleTotalMS(windows); got != 8_000 {
		t.Fatalf("IdleTotalMS = %d, want 8000 (5 s + 3 s)", got)
	}
	if rating := ratingengine.ActivityIdleMS(in.Fight, furyGUID, nil, nil); rating != 2_000 {
		t.Fatalf("ActivityIdleMS = %d, want the rating's own 2000 (60 s less ActiveMS 58 s)", rating)
	}
	// With an assignment window in force the two measures agree.
	marks := []ratingengine.Assignment{{PlayerKey: furyGUID, FromMS: 50_000, ToMS: 52_000}}
	w := ratingengine.IdleWindows(in.Fight, furyGUID, nil, marks)
	if got, rating := IdleTotalMS(w), ratingengine.ActivityIdleMS(in.Fight, furyGUID, nil, marks); got != rating || got != 8_000 {
		t.Fatalf("with a window in force IdleTotalMS %d and ActivityIdleMS %d, want both 8000", got, rating)
	}
}

func TestCountedTime_PetCastsCountAsTheOwners(t *testing.T) {
	in := countedFury(t)
	in.Fight.Casts = append(in.Fight.Casts, summary.CastRow{GUID: "Pet-0-1", OwnerGUID: furyGUID, SpellID: 1, Succeeded: 1, Sequence: []int64{1_000}})
	if got := CountedTime(in); got.FirstCastMS != 1_000 {
		t.Fatalf("FirstCastMS = %d, want the pet's 1000", got.FirstCastMS)
	}
}

// §6.1 (pre-flight F14): on a 60 s wipe, deaths in the last WipeTailMS (from 45 s) are part
// of the wipe; earlier ones are reviewed. On a kill every death is reviewed.
func TestReviewableDeath_WipeTail(t *testing.T) {
	in := countedFury(t)
	in.Fight.Kill = false
	cases := []struct {
		atMS int64
		want bool
	}{{30_000, true}, {44_999, true}, {45_000, false}, {50_000, false}}
	for _, c := range cases {
		if got := ReviewableDeath(in, summary.Death{GUID: furyGUID, AtMS: c.atMS}); got != c.want {
			t.Errorf("wipe, death at %d ms: reviewable = %v, want %v", c.atMS, got, c.want)
		}
	}
	in.Fight.Kill = true
	if !ReviewableDeath(in, summary.Death{GUID: furyGUID, AtMS: 50_000}) {
		t.Error("on a kill the death at 50 s is reviewable")
	}
}
