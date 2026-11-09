package db

import (
	"context"
	"testing"
)

// creatorsVersion is the creator program's migration (design/specs/2026-11-08-creator-program.md).
const creatorsVersion = 94

// beforeCreators is the newest migration on main before ours (going down to it also takes back
// any lane's migration numbered between, which the cleanup restores).
const beforeCreators = 93

// The creator program's migration adds its tables and the entitlements source 'creator', holds
// one open application per account and the code and channel rules, and its down removes all of
// it, turning a creator grant back into a plain grant.
func TestMigrationAddsTheCreatorProgram(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, creatorsVersion)
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	mustExec := func(sql string, args ...any) {
		t.Helper()
		if _, err := pool.Exec(ctx, sql, args...); err != nil {
			t.Fatalf("%s: %v", sql, err)
		}
	}
	mustExec(`truncate users cascade`)
	mustExec(`delete from creators`)
	mustExec(`delete from go_links where slug = 'mig-creator'`)
	mustExec(`insert into users (id) values (9101), (9102)`)
	mustExec(`insert into go_links (slug, target, label) values ('mig-creator', '/', 'Creator')`)
	mustExec(`insert into creators (name, channel_url, slug, code, starts_on) values ('Mig', 'https://twitch.tv/mig', 'mig-creator', 'MIGCODE', current_date)`)
	for _, bad := range []string{`insert into creators (name, channel_url, slug, code, starts_on) values ('X', 'https://twitch.tv/x', 'reddit', 'lower', current_date)`,
		`insert into creators (name, channel_url, slug, code, starts_on) values ('X', 'http://twitch.tv/x', 'reddit', 'GOODCODE', current_date)`} {
		if _, err := pool.Exec(ctx, bad); err == nil {
			t.Errorf("accepted: %s", bad)
		}
	}
	mustExec(`insert into creator_applications (user_id, channel_url, platform, about) values (9101, 'https://twitch.tv/a', 'twitch', 'Forever runs')`)
	if _, err := pool.Exec(ctx, `insert into creator_applications (user_id, channel_url, platform, about) values (9101, 'https://twitch.tv/a', 'twitch', 'again')`); err == nil {
		t.Error("a second open application was accepted")
	}
	mustExec(`insert into entitlements (user_id, plan, source, status) values (9102, 'premium', 'creator', 'active')`)

	migrateTo(t, url, beforeCreators)
	var tables int
	if err := pool.QueryRow(ctx, `select count(*) from information_schema.tables where table_name like 'creator%'`).Scan(&tables); err != nil {
		t.Fatal(err)
	}
	if tables != 0 {
		t.Fatalf("%d creator tables left after down", tables)
	}
	var source string
	if err := pool.QueryRow(ctx, `select source from entitlements where user_id = 9102`).Scan(&source); err != nil || source != "grant" {
		t.Fatalf("creator grant after down = %q (%v), want grant", source, err)
	}
}
