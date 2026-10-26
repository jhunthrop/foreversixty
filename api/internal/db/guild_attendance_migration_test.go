package db

import (
	"context"
	"testing"
)

// guildAttendanceSchema counts what 0052 adds: its columns, its two tables and its three
// new indexes. Zero means none of it is there.
func guildAttendanceSchema(t *testing.T, url string) (columns, tables, indexes int) {
	t.Helper()
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	if err := pool.QueryRow(ctx,
		`select count(*) from information_schema.columns
		  where (table_name = 'reports' and column_name = 'log_zone')
		     or (table_name = 'guild_raids' and column_name in
		         ('pre_done_status', 'matched_by', 'matched_at', 'report_posted_at'))
		     or (table_name = 'loot_awards' and column_name = 'raid_id')`).Scan(&columns); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx,
		`select count(*) from information_schema.tables
		  where table_name in ('guild_raid_rejected_reports', 'guild_raid_attendance')`).Scan(&tables); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx,
		`select count(*) from pg_indexes
		  where indexname in ('guild_raids_report_idx', 'loot_awards_drop_unmatched_idx')
		     or (indexname = 'loot_awards_drop_idx' and indexdef like '%raid_id%')`).Scan(&indexes); err != nil {
		t.Fatal(err)
	}
	return columns, tables, indexes
}

// 0052_guild_attendance carries Lane G2's whole schema; its down removes all of it and puts
// back 0030's one-award-per-drop index; up again restores it.
func TestMigration0052AddsTheAttendanceSchemaAndDownReversesIt(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	if c, tb, ix := guildAttendanceSchema(t, url); c != 6 || tb != 2 || ix != 3 {
		t.Fatalf("after up: columns %d tables %d indexes %d, want 6 2 3", c, tb, ix)
	}

	migrateTo(t, url, 51)
	if c, tb, ix := guildAttendanceSchema(t, url); c != 0 || tb != 0 || ix != 0 {
		t.Fatalf("after down: columns %d tables %d indexes %d, want 0 0 0", c, tb, ix)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	var restored int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from pg_indexes where indexname = 'loot_awards_drop_idx'
		    and indexdef like '%(guild_id, encounter_id, item_id)%'`).Scan(&restored); err != nil {
		t.Fatal(err)
	}
	pool.Close()
	if restored != 1 {
		t.Fatal("the down migration did not restore 0030's loot_awards_drop_idx")
	}

	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	if c, tb, ix := guildAttendanceSchema(t, url); c != 6 || tb != 2 || ix != 3 {
		t.Fatalf("after up again: columns %d tables %d indexes %d, want 6 2 3", c, tb, ix)
	}
}
