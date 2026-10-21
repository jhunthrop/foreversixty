package db

import (
	"context"
	"testing"
	"time"
)

// 0043_rekey_two_word_units moves fight_metrics and rating_scores rows keyed by the first word
// of a two-word name onto the First-Last key, leaves a row whose corrected key already exists
// for the same fight, and leaves every row that is not half-keyed. A second run changes nothing.
func TestMigration0043RekeysHalfNameRows(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	ctx := context.Background()
	pool, err := Connect(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	at := time.Date(2026, 10, 5, 20, 0, 0, 0, time.UTC)
	if err := EnsureMetricsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}
	if err := EnsureRatingsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}

	migrateTo(t, url, 42)

	// fight 1: half-keyed only, so it is re-keyed. fight 2: half-keyed and already correctly
	// keyed, so the half-keyed row stays. fight 3: a one-word character (two-segment unit) and
	// a one-word character whose unit carries only a realm word, neither of which matches.
	// fight 4: a key in a different region and ruleset keeps them.
	type seed struct {
		fight     int
		key, name string
	}
	seeds := []seed{
		{1, "us/pvp/bow", "Bow-Jackzon-ClassicBetaPvP-"},
		{2, "us/pvp/bow", "Bow-Jackzon-ClassicBetaPvP-"},
		{2, "us/pvp/bow-jackzon", "Bow-Jackzon-ClassicBetaPvP-"},
		{3, "us/hardcore/baelgrim", "Baelgrim-Nightslayer"},
		{3, "us/pvp/bow", "Bow-ClassicBetaPvP-"},
		{4, "eu/normal/ana", "Ana-Maria-Realm"},
	}
	if _, err := pool.Exec(ctx, `delete from fight_metrics where report_id = 'rk-0043'`); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `delete from rating_scores where report_id = 'rk-0043'`); err != nil {
		t.Fatal(err)
	}
	for _, s := range seeds {
		if _, err := pool.Exec(ctx,
			`insert into fight_metrics (report_id, fight_index, player_key, player_name, fought_at)
			 values ('rk-0043', $1, $2, $3, $4)`, s.fight, s.key, s.name, at); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx,
			`insert into rating_scores (report_id, fight_index, player_key, player_name, kill, components, model_version, fought_at)
			 values ('rk-0043', $1, $2, $3, true, '{}', 'test', $4)`, s.fight, s.key, s.name, at); err != nil {
			t.Fatal(err)
		}
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from fight_metrics where report_id = 'rk-0043'`)
		pool.Exec(ctx, `delete from rating_scores where report_id = 'rk-0043'`)
	})

	keysOf := func(table string) map[string]int {
		t.Helper()
		rows, err := pool.Query(ctx,
			`select fight_index::text || ' ' || player_key, 1 from `+table+` where report_id = 'rk-0043'`)
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()
		got := map[string]int{}
		for rows.Next() {
			var k string
			var n int
			if err := rows.Scan(&k, &n); err != nil {
				t.Fatal(err)
			}
			got[k] += n
		}
		return got
	}
	want := map[string]int{
		"1 us/pvp/bow-jackzon":   1,
		"2 us/pvp/bow":           1,
		"2 us/pvp/bow-jackzon":   1,
		"3 us/hardcore/baelgrim": 1,
		"3 us/pvp/bow":           1,
		"4 eu/normal/ana-maria":  1,
	}
	check := func(stage string) {
		t.Helper()
		for _, table := range []string{"fight_metrics", "rating_scores"} {
			got := keysOf(table)
			if len(got) != len(want) {
				t.Errorf("%s %s: keys = %v, want %v", stage, table, got, want)
				continue
			}
			for k, n := range want {
				if got[k] != n {
					t.Errorf("%s %s: keys = %v, want %v", stage, table, got, want)
					break
				}
			}
		}
	}

	migrateTo(t, url, 43)
	check("after 0043")
	// A second run is a no-op: down is a documented no-op and up applies the same statements.
	migrateTo(t, url, 42)
	check("after the 0043 down")
	migrateTo(t, url, 43)
	check("after 0043 again")
}
