// api/internal/dataaddon/guild_aggregate_test.go
package dataaddon

import "testing"

func fraction(v float64) *float64 { return &v }

func TestMeanAttendanceIsTheMeanOfTheMembersFractions(t *testing.T) {
	if got := meanAttendance([]float64{0.75, 0.25}); got == nil || *got != 0.5 {
		t.Errorf("mean = %v, want 0.5", got)
	}
	if got := meanAttendance(nil); got != nil {
		t.Errorf("no scored members: %v, want none", *got)
	}
}

func TestStandingsRankByKillsThenAttendanceAndShareTies(t *testing.T) {
	got := standings([]guildStanding{
		{ID: 1, Kills: 2, Attendance: fraction(0.5)},
		{ID: 2, Kills: 5, Attendance: nil},
		{ID: 3, Kills: 2, Attendance: fraction(0.9)},
		{ID: 4, Kills: 2, Attendance: fraction(0.5)},
		{ID: 5, Kills: 0, Attendance: nil},
	})
	want := map[int64]int{2: 1, 3: 2, 1: 3, 4: 3, 5: 5}
	for id, rank := range want {
		if got[id] != rank {
			t.Errorf("guild %d: standing %d, want %d (all: %v)", id, got[id], rank, got)
		}
	}
}
