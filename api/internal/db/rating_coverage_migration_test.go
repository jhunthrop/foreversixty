// api/internal/db/rating_coverage_migration_test.go
package db

import (
	"context"
	"testing"
)

func TestMigrateAddsRatingCoverageColumns(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	for _, col := range []string{"coverage", "insufficient", "insufficient_reason"} {
		var n int
		if err := pool.QueryRow(context.Background(),
			`select count(*) from information_schema.columns where table_name = 'rating_scores' and column_name = $1`,
			col).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Errorf("rating_scores.%s: count = %d, want 1", col, n)
		}
	}
}

func TestMigration0023DownReversesUpWithoutDroppingRatingScores(t *testing.T) {
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

	migrateTo(t, url, 22)
	for _, col := range []string{"coverage", "insufficient", "insufficient_reason"} {
		var n int
		if err := pool.QueryRow(context.Background(),
			`select count(*) from information_schema.columns where table_name = 'rating_scores' and column_name = $1`,
			col).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 0 {
			t.Errorf("rating_scores.%s survived the 0023 down migration", col)
		}
	}
	if n := tableCount(t, pool, "rating_scores"); n != 1 {
		t.Error("the 0023 down migration took the rating_scores table itself, which 0022 owns")
	}

	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	for _, col := range []string{"coverage", "insufficient", "insufficient_reason"} {
		var n int
		if err := pool.QueryRow(context.Background(),
			`select count(*) from information_schema.columns where table_name = 'rating_scores' and column_name = $1`,
			col).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Errorf("rating_scores.%s did not come back after migrating up again", col)
		}
	}
}
