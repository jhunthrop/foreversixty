package db

import (
	"context"
	"testing"
)

// simRequestSimIDIndexes counts the indexes on addon_sim_requests that lead with sim_id.
func simRequestSimIDIndexes(t *testing.T, url string) int {
	t.Helper()
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from pg_indexes
		  where tablename = 'addon_sim_requests' and indexname = 'addon_sim_requests_sim_id_idx'
		    and indexdef like '%(sim_id)%'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

// 0039_sim_request_sim_id_idx indexes addon_sim_requests.sim_id, the column sims.Store.Mine's
// FromAddon (exists ... where sim_id = sims.id) and the inbox's run joins look rows up by; its
// down drops the index and leaves 0038's column; up again restores it.
func TestMigration0039IndexesTheRequestsSimIDAndDownReversesIt(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	if n := simRequestSimIDIndexes(t, url); n != 1 {
		t.Fatalf("addon_sim_requests_sim_id_idx count = %d after migrating up, want 1", n)
	}

	migrateTo(t, url, 38)
	if n := simRequestSimIDIndexes(t, url); n != 0 {
		t.Fatal("addon_sim_requests_sim_id_idx survived the 0039 down migration")
	}
	if n := simRequestExportColumns(t, url); n != 1 {
		t.Error("the 0039 down migration took 0038's export_seen_at with it")
	}

	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	if n := simRequestSimIDIndexes(t, url); n != 1 {
		t.Fatal("addon_sim_requests_sim_id_idx did not come back after migrating up again")
	}
}
