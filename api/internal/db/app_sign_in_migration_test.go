package db

import (
	"context"
	"testing"
)

const appSignInVersion = 60

// 0060 adds app_auth_flows and devices.kind; every existing device reads as a companion.
func TestMigrationAddsAppSignIn(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, appSignInVersion-1)
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
	if err := pool.QueryRow(ctx, `insert into users (email) values ('o@example.com') returning id`).Scan(&uid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into devices (id, user_id, name, platform, token_hash) values ('old', $1, 'PC', 'windows', 'h')`, uid); err != nil {
		t.Fatal(err)
	}
	migrateTo(t, url, appSignInVersion)

	var kind string
	if err := pool.QueryRow(ctx, `select kind from devices where id = 'old'`).Scan(&kind); err != nil || kind != "companion" {
		t.Fatalf("existing device kind = %q, %v; want companion", kind, err)
	}
	if _, err := pool.Exec(ctx, `update devices set kind = 'tablet' where id = 'old'`); err == nil {
		t.Fatal("an unknown device kind was stored")
	}
	const challenge = "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM"
	if _, err := pool.Exec(ctx,
		`insert into app_auth_flows (state_hash, platform, code_challenge, app_state, expires_at)
		 values ('s1', 'ios', $1, 'st', now())`, challenge); err != nil {
		t.Fatalf("a valid flow is refused: %v", err)
	}
	for name, q := range map[string]string{
		"desktop platform": `insert into app_auth_flows (state_hash, platform, code_challenge, app_state, expires_at) values ('s2', 'windows', '` + challenge + `', 'st', now())`,
		"short challenge":  `insert into app_auth_flows (state_hash, platform, code_challenge, app_state, expires_at) values ('s3', 'ios', 'abc', 'st', now())`,
	} {
		if _, err := pool.Exec(ctx, q); err == nil {
			t.Errorf("%s was stored", name)
		}
	}
	if _, err := pool.Exec(ctx, `truncate app_auth_flows`); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into devices (id, user_id, name, platform, token_hash, kind) values ('phone', $1, 'iPhone', 'ios', 'p', 'app')`, uid); err != nil {
		t.Fatal(err)
	}
	migrateTo(t, url, appSignInVersion-1)
	// Without kind the phone would read as a companion, so going down revokes it.
	var phoneRevoked, oldRevoked bool
	if err := pool.QueryRow(ctx,
		`select (select revoked_at is not null from devices where id = 'phone'),
		        (select revoked_at is not null from devices where id = 'old')`).Scan(&phoneRevoked, &oldRevoked); err != nil {
		t.Fatal(err)
	}
	if !phoneRevoked || oldRevoked {
		t.Fatalf("after down: app revoked = %v, companion revoked = %v; want true, false", phoneRevoked, oldRevoked)
	}
	var n int
	if err := pool.QueryRow(ctx,
		`select count(*) from information_schema.columns where table_name = 'devices' and column_name = 'kind'`).Scan(&n); err != nil || n != 0 {
		t.Fatalf("down migration left devices.kind (%d, %v)", n, err)
	}
}
