package db

import (
	"context"
	"testing"
)

const activeDaysVersion = 104

// 0104 adds user_active_days and backfills it from what history keeps: signups, sessions, the
// player's own sims, addon sims and syncs, the latest addon export and the devices, each on its
// surface, every row marked backfilled. Import sims stay out: the nightly refresh starts them.
func TestMigrationBackfillsTheActiveDays(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, activeDaysVersion-1)
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	for _, sql := range []string{
		`truncate users cascade`,
		`insert into users (id, battletag, created_at) values (5, 'a#1', '2026-11-01T10:00:00Z')`,
		`insert into sessions (id, user_id, method, created_at, expires_at) values
		   ('s1', 5, 'bnet', '2026-11-03T10:00:00Z', '2026-12-03T10:00:00Z')`,
		`insert into characters (key, region, ruleset, name, user_id) values ('us/pvp/a', 'us', 'pvp', 'A', 5)`,
		`insert into sims (id, user_id, spec, engine_version, lane, dps_mean, dps_error, iterations, result, created_at) values
		   ('own', 5, 'mage-fire', 'v', 'server', 1, 1, 1, '{}', '2026-11-04T10:00:00Z'),
		   ('imp', 5, 'mage-fire', 'v', 'server', 1, 1, 1, '{}', '2026-11-05T10:00:00Z'),
		   ('add', 5, 'mage-fire', 'v', 'server', 1, 1, 1, '{}', '2026-11-06T10:00:00Z')`,
		`insert into character_sims (sim_id, character_key, user_id, kind, spec, role, band, preset, created_at) values
		   ('imp', 'us/pvp/a', 5, 'import', 'mage-fire', 'dps', 60, 'bare', '2026-11-05T10:00:00Z'),
		   ('add', 'us/pvp/a', 5, 'addon', 'mage-fire', 'dps', 60, 'bare', '2026-11-06T10:00:00Z')`,
		`insert into character_syncs (character_key, source, outcome, created_at) values
		   ('us/pvp/a', 'addon', 'ok', '2026-11-07T10:00:00Z'), ('us/pvp/a', 'blizzard', 'ok', '2026-11-12T10:00:00Z')`,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, seen_at) values
		   ('us/pvp/a', 5, 'us', 'pvp', 'A', '', '2026-11-08T10:00:00Z', '2026-11-08T10:00:00Z')`,
		`insert into devices (id, user_id, token_hash, kind, created_at, last_seen_at) values
		   ('pc', 5, '\x01', 'companion', '2026-11-09T10:00:00Z', '2026-11-10T10:00:00Z'),
		   ('ph', 5, '\x02', 'app', '2026-11-11T10:00:00Z', null)`,
	} {
		if _, err := pool.Exec(ctx, sql); err != nil {
			t.Fatalf("%s: %v", sql, err)
		}
	}
	migrateTo(t, url, activeDaysVersion)

	rows, err := pool.Query(ctx, `select day::text, surfaces, backfilled from user_active_days where user_id = 5 order by day`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	got := map[string]int16{}
	for rows.Next() {
		var day string
		var surfaces int16
		var backfilled bool
		if err := rows.Scan(&day, &surfaces, &backfilled); err != nil {
			t.Fatal(err)
		}
		if !backfilled {
			t.Errorf("%s is not marked backfilled", day)
		}
		got[day] = surfaces
	}
	want := map[string]int16{
		"2026-11-01": 1, "2026-11-03": 1, "2026-11-04": 1, "2026-11-06": 2, "2026-11-07": 2,
		"2026-11-08": 2, "2026-11-09": 2, "2026-11-10": 2, "2026-11-11": 4,
	}
	if len(got) != len(want) {
		t.Fatalf("days = %v, want %v", got, want)
	}
	for day, s := range want {
		if got[day] != s {
			t.Errorf("%s: surfaces %d, want %d", day, got[day], s)
		}
	}
	migrateTo(t, url, activeDaysVersion-1)
	var exists bool
	if err := pool.QueryRow(ctx, `select to_regclass('user_active_days') is not null`).Scan(&exists); err != nil || exists {
		t.Fatalf("down left the table: %v %v", exists, err)
	}
}
