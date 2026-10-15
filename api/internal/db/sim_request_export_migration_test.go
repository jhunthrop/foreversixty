package db

import (
	"context"
	"testing"
)

// simRequestExportColumns counts addon_sim_requests.export_seen_at as a timestamptz (0 or 1).
func simRequestExportColumns(t *testing.T, url string) int {
	t.Helper()
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from information_schema.columns
		  where table_name = 'addon_sim_requests' and column_name = 'export_seen_at'
		    and data_type = 'timestamp with time zone' and is_nullable = 'YES'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

// 0038_sim_request_export adds the nullable addon_sim_requests.export_seen_at; its down takes
// the column away and leaves the table and 0036's state column; up again restores it.
func TestMigration0038AddsTheRequestsExportAndDownReversesIt(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	if n := simRequestExportColumns(t, url); n != 1 {
		t.Fatalf("addon_sim_requests.export_seen_at count = %d after migrating up, want 1", n)
	}

	migrateTo(t, url, 37)
	if n := simRequestExportColumns(t, url); n != 0 {
		t.Fatal("addon_sim_requests.export_seen_at survived the 0038 down migration")
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	var state int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from information_schema.columns
		  where table_name = 'addon_sim_requests' and column_name = 'state'`).Scan(&state); err != nil {
		t.Fatal(err)
	}
	if state != 1 {
		t.Error("the 0038 down migration took addon_sim_requests.state with it")
	}

	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	if n := simRequestExportColumns(t, url); n != 1 {
		t.Fatal("addon_sim_requests.export_seen_at did not come back after migrating up again")
	}
}
