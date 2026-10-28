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
