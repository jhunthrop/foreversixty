// api/cmd/api/rerate_flagged_test.go
package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
)

// emptyBucket is a bucket that holds no fight summary: every read misses, so every picked
// fight fails to re-rate.
type emptyBucket struct{}

func (emptyBucket) Get(context.Context, string) (io.ReadCloser, error) {
	return nil, errors.New("no such object")
}

// The rating-rerate-flagged job exits non-zero when a fight could not be re-rated (task
// L6.10's exit branch), after trying every fight and logging the counts: on the real
// database, one character in a flagged report whose newest clean fight's summary is not in
// the bucket.
func TestTheRerateJobFailsWhenAFightCouldNotBeRerated(t *testing.T) {
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	pool, err := db.Connect(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	const key = "us/normal/l612rerate"
	t.Cleanup(func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id like 'l612-rerate-%'`)
		pool.Exec(bg, `delete from reports where id like 'l612-rerate-%'`)
	})
	at := time.Now().Add(-48 * time.Hour)
	if err := db.EnsureRatingsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}
	for _, r := range []struct{ id, flagged string }{{"l612-rerate-clean", ""}, {"l612-rerate-flagged", "tampered"}} {
		if _, err := pool.Exec(ctx, `insert into reports (id, title, visibility, zone, status, flagged)
			values ($1, 'L6.12', 'public', 'Molten Core', 'complete', nullif($2, ''))`, r.id, r.flagged); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx, `insert into rating_scores (report_id, fight_index, player_key, overall, components,
			model_version, fought_at, kill) values ($1, 1, $2, 50, '[]', '', $3, true)`, r.id, key, at); err != nil {
			t.Fatal(err)
		}
	}

	var logged bytes.Buffer
	log := slog.New(slog.NewTextHandler(&logged, nil))
	err = rerateFlagged(ctx, rating.BackfillDeps{Store: &rating.Store{Pool: pool, Log: log}, Summaries: emptyBucket{}, Log: log}, log)
	if err == nil {
		t.Fatalf("the job succeeded with a fight it could not re-rate; it logged:\n%s", logged.String())
	}
	if !strings.Contains(err.Error(), "could not be re-rated") {
		t.Errorf("error = %v, want the failed-fight count", err)
	}
	if !strings.Contains(logged.String(), "msg="+rating.RerateFlaggedJobCommand) || !strings.Contains(logged.String(), "failed=") {
		t.Errorf("the counts were not logged before the job failed:\n%s", logged.String())
	}
}
