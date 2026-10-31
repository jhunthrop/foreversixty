package db

import (
	"context"
	"testing"
)

const liveActivityTokensVersion = 61

// 0061 adds live_activity_tokens; its down drops it, and a revoked device takes its rows.
func TestMigrationAddsLiveActivityTokens(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, liveActivityTokensVersion)
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `truncate users cascade`); err != nil {
		t.Fatal(err)
	}
	var uid int64
	if err := pool.QueryRow(ctx, `insert into users (email) values ('l@example.com') returning id`).Scan(&uid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into devices (id, user_id, name, platform, token_hash, kind) values ('ph', $1, 'iPhone', 'ios', 'h', 'app')`, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into live_activity_tokens
		(user_id, device_id, kind, platform, object_id, push_token, expires_at)
		values ($1, 'ph', 'raid', 'ios', '7', 'abcd', now() + interval '1 hour')`, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into live_activity_tokens
		(user_id, kind, platform, object_id, push_token, expires_at)
		values ($1, 'party', 'ios', '7', 'abcd', now())`, uid); err == nil {
		t.Fatal("an unknown kind was accepted")
	}
	if _, err := pool.Exec(ctx, `delete from devices where id = 'ph'`); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := pool.QueryRow(ctx, `select count(*) from live_activity_tokens`).Scan(&n); err != nil || n != 0 {
		t.Fatalf("rows after the device went = %d, %v", n, err)
	}
	migrateTo(t, url, liveActivityTokensVersion-1)
	var exists bool
	if err := pool.QueryRow(ctx, `select to_regclass('live_activity_tokens') is not null`).Scan(&exists); err != nil || exists {
		t.Fatalf("table after down = %v, %v", exists, err)
	}
}
