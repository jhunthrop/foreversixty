package db

import (
	"context"
	"testing"
)

// adherenceColumns counts rating_scores.adherence (0 or 1).
func adherenceColumns(t *testing.T, url string) int {
	t.Helper()
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from information_schema.columns
		  where table_name = 'rating_scores' and column_name = 'adherence' and data_type = 'numeric'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

// 0037_adherence adds the nullable numeric rating_scores.adherence; its down takes the
// column away and leaves 0036's last-raid facts and the table itself; up again restores it.
func TestMigration0037AddsAdherenceAndDownReversesIt(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	if n := adherenceColumns(t, url); n != 1 {
		t.Fatalf("rating_scores.adherence count = %d after migrating up, want 1", n)
	}

	migrateTo(t, url, 36)
	if n := adherenceColumns(t, url); n != 0 {
		t.Fatal("rating_scores.adherence survived the 0037 down migration")
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	var facts int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from information_schema.columns where table_name = 'rating_scores'
		    and column_name in ('deaths', 'biggest_loss', 'loss_seconds', 'consumables_missed')`).Scan(&facts); err != nil {
		t.Fatal(err)
	}
	if facts != 4 {
		t.Errorf("the 0037 down migration took 0036's last-raid facts: %d of 4 left", facts)
	}

	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	if n := adherenceColumns(t, url); n != 1 {
		t.Fatal("rating_scores.adherence did not come back after migrating up again")
	}
}
