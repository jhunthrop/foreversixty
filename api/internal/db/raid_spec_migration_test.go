package db

import (
	"context"
	"testing"
)

const raidSpecVersion = 58

// 0058 adds guild_characters.raid_spec, null by default, holding only a specs.json key.
func TestMigrationAddsRaidSpec(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, raidSpecVersion-1)
	migrateTo(t, url, raidSpecVersion)

	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `truncate users, guilds cascade`); err != nil {
		t.Fatal(err)
	}
	var uid, gid int64
	if err := pool.QueryRow(ctx, `insert into users (email) values ('o@example.com') returning id`).Scan(&uid); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into guilds (region, ruleset, name) values ('us','normal','G') returning id`).Scan(&gid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into guild_characters (guild_id, user_id, character_key) values ($1, $2, 'us/normal/aurelith')`, gid, uid); err != nil {
		t.Fatal(err)
	}
	var spec *string
	if err := pool.QueryRow(ctx, `select raid_spec from guild_characters where guild_id = $1`, gid).Scan(&spec); err != nil || spec != nil {
		t.Fatalf("new row raid_spec = %v, err %v", spec, err)
	}
	if _, err := pool.Exec(ctx, `update guild_characters set raid_spec = 'druid-feral-bear' where guild_id = $1`, gid); err != nil {
		t.Fatalf("a specs.json key is refused: %v", err)
	}
	if _, err := pool.Exec(ctx, `update guild_characters set raid_spec = 'Holy' where guild_id = $1`, gid); err == nil {
		t.Fatal("a catalogue name was stored as a spec key")
	}
}
