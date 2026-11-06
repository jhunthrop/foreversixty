package db

import (
	"context"
	"testing"
)

// raidAssignmentTables counts 0077's two tables. Zero means neither is there.
func raidAssignmentTables(t *testing.T, url string) int {
	t.Helper()
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from information_schema.tables
		  where table_name in ('raid_assignments', 'raid_assignment_posts')`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

// 0077 adds the raid assignments and their Discord post throttle; down removes both; up
// restores them. The down step goes to 75, below 0076 (another lane's) as well.
func TestMigration0077AddsTheRaidAssignmentTablesAndDownReversesIt(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, 77)
	if n := raidAssignmentTables(t, url); n != 2 {
		t.Fatalf("after up: %d tables, want 2", n)
	}
	migrateTo(t, url, 75)
	if n := raidAssignmentTables(t, url); n != 0 {
		t.Fatalf("after down: %d tables, want 0", n)
	}
	migrateTo(t, url, 77)
	if n := raidAssignmentTables(t, url); n != 2 {
		t.Fatalf("after up again: %d tables, want 2", n)
	}
}
