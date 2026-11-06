// api/internal/db/guild_invite_loot_migration_test.go
package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func connectTest(t *testing.T, url string) *pgxpool.Pool {
	t.Helper()
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

func countOf(t *testing.T, pool *pgxpool.Pool, query string, args ...any) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(), query, args...).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

const inviteColumnsQuery = `select count(*) from information_schema.columns
	where table_name = 'guilds' and column_name in ('invite_token_enc', 'invite_rotated_by')`

const lootIndexQuery = `select count(*) from pg_indexes where indexname = $1`

// 0070 adds the sealed invite token and its maker to guilds; its down removes both.
func TestMigration0070AddsTheKeptInviteAndDownRemovesIt(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	pool := connectTest(t, url)
	if n := countOf(t, pool, inviteColumnsQuery); n != 2 {
		t.Fatalf("after up: %d invite columns, want 2", n)
	}
	migrateTo(t, url, 69)
	if n := countOf(t, pool, inviteColumnsQuery); n != 0 {
		t.Fatalf("after down: %d invite columns, want 0", n)
	}
	migrateTo(t, url, 70)
	if n := countOf(t, pool, inviteColumnsQuery); n != 2 {
		t.Fatalf("after up again: %d invite columns, want 2", n)
	}
}

// 0071 swaps the one-raid-less-award-per-drop index for one per drop per logged night, so two
// raid-less awards of one item with no log both stand. Its down keeps the newest of each drop
// (what 0052's upsert kept) and restores the old index.
func TestMigration0071LetsTwoUnmatchedAwardsStandAndDownKeepsTheNewest(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	pool := connectTest(t, url)
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `truncate guilds cascade`); err != nil {
		t.Fatal(err)
	}
	if countOf(t, pool, lootIndexQuery, "loot_awards_drop_unmatched_idx") != 0 ||
		countOf(t, pool, lootIndexQuery, "loot_awards_drop_night_idx") != 1 {
		t.Fatal("after up: want the night index and no unmatched index")
	}
	var gid int64
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Mig71') returning id`).Scan(&gid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `
		insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_at) values
		  ($1, 1084, 16900, 'us/hardcore/older', now() - interval '7 days'),
		  ($1, 1084, 16900, 'us/hardcore/newer', now())`, gid); err != nil {
		t.Fatalf("two raid-less awards of one item: %v", err)
	}

	migrateTo(t, url, 70)
	if countOf(t, pool, lootIndexQuery, "loot_awards_drop_unmatched_idx") != 1 ||
		countOf(t, pool, lootIndexQuery, "loot_awards_drop_night_idx") != 0 {
		t.Fatal("after down: want the unmatched index back and no night index")
	}
	var kept string
	if err := pool.QueryRow(ctx, `select string_agg(character_key, ',') from loot_awards where guild_id = $1`, gid).
		Scan(&kept); err != nil {
		t.Fatal(err)
	}
	if kept != "us/hardcore/newer" {
		t.Fatalf("after down the awards are %q, want only the newer", kept)
	}

	migrateTo(t, url, 71)
	if countOf(t, pool, lootIndexQuery, "loot_awards_drop_night_idx") != 1 {
		t.Fatal("after up again: want the night index")
	}
}
