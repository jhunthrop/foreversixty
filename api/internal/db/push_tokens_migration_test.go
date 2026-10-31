package db

import (
	"context"
	"testing"
)

const pushTokensVersion = 62

// 0062's push tokens go with their device: deleted with it, and dropped on revoke. 0061 and
// 0062 go down and up again cleanly.
func TestMigrationDropsPushTokensOnRevoke(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, raidSpecVersion)
	migrateTo(t, url, pushTokensVersion)

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
	if err := pool.QueryRow(ctx, `insert into users (email) values ('p@example.com') returning id`).Scan(&uid); err != nil {
		t.Fatal(err)
	}
	for _, dev := range []string{"dev-a", "dev-b"} {
		if _, err := pool.Exec(ctx, `insert into devices (id, user_id, token_hash) values ($1, $2, $3)`, dev, uid, []byte(dev)); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx, `insert into push_tokens (token, platform, user_id, device_id) values ($1, 'ios', $2, $3)`,
			dev+"-0123456789abcdef", uid, dev); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := pool.Exec(ctx, `insert into push_tokens (token, platform, user_id, device_id) values ('0123456789abcdefX', 'web', $1, 'dev-a')`, uid); err == nil {
		t.Fatal("an unknown platform was accepted")
	}
	if _, err := pool.Exec(ctx, `update devices set revoked_at = now() where id = 'dev-a'`); err != nil {
		t.Fatal(err)
	}
	var left []string
	rows, err := pool.Query(ctx, `select device_id from push_tokens order by device_id`)
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		var d string
		if err := rows.Scan(&d); err != nil {
			t.Fatal(err)
		}
		left = append(left, d)
	}
	if len(left) != 1 || left[0] != "dev-b" {
		t.Fatalf("tokens left after revoke: %v", left)
	}
	if _, err := pool.Exec(ctx, `delete from devices where id = 'dev-b'`); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := pool.QueryRow(ctx, `select count(*) from push_tokens`).Scan(&n); err != nil || n != 0 {
		t.Fatalf("tokens after delete: %d %v", n, err)
	}
	if _, err := pool.Exec(ctx, `insert into notification_prefs (user_id, kind, channel, enabled) values ($1, 'Raid', 'push', true)`, uid); err == nil {
		t.Fatal("a malformed kind was accepted")
	}
}
