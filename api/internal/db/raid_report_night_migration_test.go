package db

import (
	"context"
	"testing"
)

// raidReportNightTables counts 0074's two tables. Zero means neither is there.
func raidReportNightTables(t *testing.T, url string) int {
	t.Helper()
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from information_schema.tables
		  where table_name in ('guild_raid_attendance_overrides', 'guild_report_pulls')`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

// 0074 adds the attendance fixes and the per-pull digests; down removes both; up restores them.
func TestMigration0074AddsTheReportNightTablesAndDownReversesIt(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, 74)
	if n := raidReportNightTables(t, url); n != 2 {
		t.Fatalf("after up: %d tables, want 2", n)
	}
	migrateTo(t, url, 71) // 0072 and 0073 belong to other lanes; 71 is below all three
	if n := raidReportNightTables(t, url); n != 0 {
		t.Fatalf("after down: %d tables, want 0", n)
	}
	migrateTo(t, url, 74)
	if n := raidReportNightTables(t, url); n != 2 {
		t.Fatalf("after up again: %d tables, want 2", n)
	}
}
