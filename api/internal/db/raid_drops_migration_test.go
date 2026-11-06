package db

import (
	"context"
	"testing"
)

// raidDropsTables counts 0076's table. Zero means it is not there.
func raidDropsTables(t *testing.T, url string) int {
	t.Helper()
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from information_schema.tables where table_name = 'raid_drops'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

// 0076 adds the recorded drops; down removes them; up restores them.
func TestMigration0076AddsRaidDropsAndDownReversesIt(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, 76)
	if n := raidDropsTables(t, url); n != 1 {
		t.Fatalf("after up: %d tables, want 1", n)
	}
	migrateTo(t, url, 75)
	if n := raidDropsTables(t, url); n != 0 {
		t.Fatalf("after down: %d tables, want 0", n)
	}
	migrateTo(t, url, 76)
	if n := raidDropsTables(t, url); n != 1 {
		t.Fatalf("after up again: %d tables, want 1", n)
	}
}
