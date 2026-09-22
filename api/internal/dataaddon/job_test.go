// api/internal/dataaddon/job_test.go
package dataaddon

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

type recordingLogger struct {
	warnings []string
}

func (l *recordingLogger) Info(string, ...any) {}
func (l *recordingLogger) Warn(msg string, args ...any) {
	l.warnings = append(l.warnings, msg)
}
func (l *recordingLogger) Error(string, ...any) {}

func TestRunSkipsAnUnparseablePlayerKeyAndCountsIt(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)
	foughtAt := now.Add(-2 * 24 * time.Hour)

	if _, err := pool.Exec(ctx,
		`insert into reports (id, visibility, status, created_at) values ('dataaddon-job-bad', 'public', 'complete', $1)`,
		foughtAt); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from reports where id = 'dataaddon-job-bad'`) })

	// rating_scores is partitioned by month on fought_at (migration 0022);
	// ensure the partition exists before inserting into it, the same way
	// rating.Store.RateFight does before its own insert.
	if err := db.EnsureRatingsPartition(ctx, pool, foughtAt); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into rating_scores (report_id, fight_index, player_key, player_name, encounter_id,
		   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at)
		 values ('dataaddon-job-bad', 1, 'not-a-valid-key', 'x', 1, false, 50, 50, false, '[]', 'test', $1)`,
		foughtAt); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from rating_scores where report_id = 'dataaddon-job-bad'`) })

	log := &recordingLogger{}
	result, err := Run(ctx, Deps{
		Store: &Store{Pool: pool}, Build: "1.60.1.69893", Now: now, Log: log,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.SkippedRows != 1 {
		t.Errorf("skipped = %d, want 1", result.SkippedRows)
	}
	if len(log.warnings) == 0 {
		t.Error("the skip should be logged")
	}
}

func TestRunSkipsTheUploadAndLogsWhyWhenNoBucketIsConfigured(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	log := &recordingLogger{}
	result, err := Run(ctx, Deps{
		Store: &Store{Pool: pool}, Build: "1.60.1.69893", Now: time.Now().UTC(), Log: log,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Files["Data.lua"] == 0 {
		t.Error("Data.lua should still be rendered (and its size logged) even with no bucket")
	}
	found := false
	for _, w := range log.warnings {
		if w == "dataaddon" {
			found = true
		}
	}
	if !found {
		t.Error("the skipped upload should be logged")
	}
}

func TestRunUploadsEveryRenderedFileAndAManifest(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	up := &FakeUploader{}
	result, err := Run(ctx, Deps{
		Store: &Store{Pool: pool}, Upload: up, Bucket: "my-bucket",
		Build: "1.60.1.69893", Now: time.Now().UTC(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Files["Data.lua"] == 0 {
		t.Fatal("Data.lua should be rendered")
	}
	var sawData, sawManifest bool
	for _, c := range up.Calls {
		if c.Bucket != "my-bucket" {
			t.Errorf("bucket = %q, want my-bucket", c.Bucket)
		}
		if c.Object == "data-addon/Data.lua" {
			sawData = true
		}
		if c.Object == "data-addon/manifest.txt" {
			sawManifest = true
			if string(c.Data) != "Data.lua\n" {
				t.Errorf("manifest = %q, want \"Data.lua\\n\"", c.Data)
			}
		}
	}
	if !sawData || !sawManifest {
		t.Errorf("calls = %+v, want Data.lua and manifest.txt uploaded", up.Calls)
	}
}
