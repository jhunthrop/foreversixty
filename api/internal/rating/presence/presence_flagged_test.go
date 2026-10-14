// api/internal/rating/presence/presence_flagged_test.go
package presence

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// Task L6.10: a report flagged tampered or unverified is not evidence, so In finds nobody in
// it even when it is asked about it; the same rows in an unflagged report are presence.
func TestInFindsNobodyInAFlaggedReport(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	t.Cleanup(func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id like 'l610p-%'`)
		pool.Exec(bg, `delete from fight_metrics where report_id like 'l610p-%'`)
		pool.Exec(bg, `delete from reports where id like 'l610p-%'`)
	})
	at := time.Date(2026, 10, 9, 21, 0, 0, 0, time.UTC)
	if err := db.EnsureMetricsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}
	if err := db.EnsureRatingsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}
	const key = "us/normal/l610present"
	for _, r := range []struct{ id, flagged string }{{"l610p-clean", ""}, {"l610p-flagged", "tampered"}} {
		mustExec(t, pool, `insert into reports (id, visibility, status, created_at, flagged)
			values ($1, 'public', 'complete', $2, nullif($3, ''))`, r.id, at, r.flagged)
		mustExec(t, pool, `insert into fight_metrics (report_id, fight_index, player_key, fought_at) values ($1, 0, $2, $3)`, r.id, key, at)
		mustExec(t, pool, `insert into rating_scores (report_id, fight_index, player_key, overall, components, model_version, fought_at, kill)
			values ($1, 0, $2, 50, '[]', 'l610', $3, true)`, r.id, key, at)
	}
	got, err := In(ctx, pool, []string{"l610p-clean", "l610p-flagged"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].ReportID != "l610p-clean" || got[0].Key != key {
		t.Fatalf("In(clean, flagged) = %+v, want only %s in l610p-clean", got, key)
	}
}
