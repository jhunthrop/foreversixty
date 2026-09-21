// api/internal/db/ratings_migration_test.go
package db

import (
	"context"
	"testing"
	"time"
)

func TestMigrateCreatesRatingsTables(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	for _, name := range []string{"rating_scores", "rating_percentile_digests", "kill_duration_digests"} {
		if n := tableCount(t, pool, name); n != 1 {
			t.Errorf("table %s: count = %d, want 1", name, n)
		}
	}
	var strategy string
	if err := pool.QueryRow(context.Background(),
		`select partstrat from pg_partitioned_table p
		 join pg_class c on c.oid = p.partrelid where c.relname = 'rating_scores'`).Scan(&strategy); err != nil {
		t.Fatalf("rating_scores is not partitioned: %v", err)
	}
	if strategy != "r" {
		t.Fatalf("partition strategy = %q, want range", strategy)
	}
}

func TestMigration0022DownReversesUp(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	migrateTo(t, url, mergeBaseMax)
	for _, gone := range []string{"rating_scores", "rating_percentile_digests", "kill_duration_digests"} {
		if n := tableCount(t, pool, gone); n != 0 {
			t.Errorf("table %s survived the down migration", gone)
		}
	}
	if n := tableCount(t, pool, "guild_members"); n != 1 {
		t.Error("the down migration took a table from an earlier migration")
	}
	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	if n := tableCount(t, pool, "rating_scores"); n != 1 {
		t.Error("rating_scores did not come back")
	}
}

func TestRatingsTableIsPartitionedByMonthViaEnsurePartition(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	at := time.Date(2026, 12, 9, 0, 0, 0, 0, time.UTC)
	if err := EnsureRatingsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := pool.QueryRow(ctx,
		`select count(*) from information_schema.tables where table_name = 'rating_scores_2026_12'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("rating_scores_2026_12 missing")
	}
}

func TestPartitionJobSweepsEveryConfiguredTable(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	job := &PartitionJob{
		Pool: pool, Ahead: 0, Every: time.Hour, Tables: []string{"fight_metrics", RatingsTable},
		Now: func() time.Time { return time.Date(2027, 3, 1, 0, 0, 0, 0, time.UTC) },
	}
	if err := job.Run(ctx); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"fight_metrics_2027_03", "rating_scores_2027_03"} {
		var n int
		if err := pool.QueryRow(ctx,
			`select count(*) from information_schema.tables where table_name = $1`, want).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Errorf("partition %s missing", want)
		}
	}
}
