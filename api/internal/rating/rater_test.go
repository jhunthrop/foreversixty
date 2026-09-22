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

// TestRaterCloseDrainsFightsAlreadyQueued proves the drain guarantee Close's doc comment
// promises: every fight already buffered when Close is called is still rated before Close
// returns, the same as sims.Scorer.Close's "waits for the score in flight ... to finish."
// Close is called immediately after scheduling, with no wait, so a Close that stopped the
// loop early (the bug this fix corrects) would leave some of these fights unrated.
func TestRaterCloseDrainsFightsAlreadyQueued(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	rater := NewRater(RateDeps{Store: store})
	go rater.Run(t.Context())

	ids := []string{"rater-drain-1", "rater-drain-2", "rater-drain-3"}
	for _, id := range ids {
		rater.Schedule(reports.RatedFight{
			ReportID: id, FightIndex: 1, Region: "us", Ruleset: "normal",
			FoughtAt: time.Date(2026, 12, 9, 2, 0, 0, 0, time.UTC), EncounterID: 667,
			Summary: summary.Summary{
				EncounterID: 667, Difficulty: 3, DurationMS: 180000, Kill: true,
				Roster: []summary.RosterRow{{GUID: "g1", Name: "Drained", Class: "Hunter", Spec: "Survival", Role: "dps"}},
			},
		})
	}
	t.Cleanup(func() {
		pool.Exec(context.Background(), `delete from rating_scores where report_id = any($1)`, ids)
	})

	rater.Close() // must block until every already-queued fight is rated, not just stop

	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(distinct report_id) from rating_scores where report_id = any($1)`, ids).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != len(ids) {
		t.Fatalf("Close returned with %d/%d already-queued fights rated; the drain guarantee is broken", n, len(ids))
	}
}

func TestRaterScheduleNeverBlocksAfterClose(t *testing.T) {
	rater := NewRater(RateDeps{Store: &Store{Pool: nil}})
	go rater.Run(t.Context())
	rater.Close()
	// Scheduling after Close must not panic (send on a closed channel) - mirrors
	// sims.Scorer's own Close-then-Schedule safety.
	rater.Schedule(reports.RatedFight{ReportID: "closed-test"})
}
