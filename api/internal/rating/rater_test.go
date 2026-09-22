// api/internal/rating/rater_test.go
package rating

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func TestRaterRunProcessesScheduledFights(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	rater := NewRater(RateDeps{Store: store})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go rater.Run(ctx)
	defer rater.Close()

	rf := reports.RatedFight{
		ReportID: "rater-test-1", FightIndex: 1, Region: "us", Ruleset: "normal",
		FoughtAt: time.Date(2026, 12, 9, 2, 0, 0, 0, time.UTC), EncounterID: 667,
		Summary: summary.Summary{
			EncounterID: 667, Difficulty: 3, DurationMS: 180000, Kill: true,
			Roster: []summary.RosterRow{{GUID: "g1", Name: "Ratered", Class: "Hunter", Spec: "Survival", Role: "dps"}},
		},
	}
	rater.Schedule(rf)
	t.Cleanup(func() { pool.Exec(context.Background(), `delete from rating_scores where report_id = $1`, rf.ReportID) })

	deadline := time.After(5 * time.Second)
	for {
		var n int
		pool.QueryRow(context.Background(), `select count(*) from rating_scores where report_id = $1`, rf.ReportID).Scan(&n)
		if n == 1 {
			return
		}
		select {
		case <-deadline:
			t.Fatal("the scheduled fight was never rated within 5s")
		case <-time.After(50 * time.Millisecond):
		}
	}
}

func TestRaterScheduleNeverBlocksAfterClose(t *testing.T) {
	rater := NewRater(RateDeps{Store: &Store{Pool: nil}})
	rater.Close()
	// Scheduling after Close must not panic (send on a closed channel) - mirrors
	// sims.Scorer's own Close-then-Schedule safety.
	rater.Schedule(reports.RatedFight{ReportID: "closed-test"})
}
