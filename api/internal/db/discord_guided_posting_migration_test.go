package db

import (
	"context"
	"testing"
)

const guidedPostingVersion = 57

// 0057 backfills every existing bot post into the schedule as posted with no reminder, and
// turns the remembered boxes on.
func TestMigrationBackfillsBotPosts(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, guidedPostingVersion-1)

	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `truncate users, guilds cascade`); err != nil {
		t.Fatal(err)
	}
	var uid, gid, raid int64
	if err := pool.QueryRow(ctx, `insert into users (email) values ('o@example.com') returning id`).Scan(&uid); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into guilds (region, ruleset, name) values ('us','normal','G') returning id`).Scan(&gid); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, created_by)
		values ($1, 'raid:onyxias-lair', 'Ony', now() + interval '2 days', 180, $2) returning id`, gid, uid).Scan(&raid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into guild_discord_installs (guild_id, server_id, server_name) values ($1, 's1', 'S')`, gid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into guild_raid_discord_posts (raid_id, guild_id, channel_id) values ($1, $2, 'c1')`, raid, gid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into guild_notification_webhooks (guild_id, token_box, channel_id, post_default)
		values ($1, '\x00', 'c1', false)`, gid); err != nil {
		t.Fatal(err)
	}
	migrateTo(t, url, guidedPostingVersion)

	var hookPost, hookRemind bool
	if err := pool.QueryRow(ctx, `select post_default, remind_default from guild_notification_webhooks where guild_id = $1`, gid).
		Scan(&hookPost, &hookRemind); err != nil || !hookPost || !hookRemind {
		t.Fatalf("webhook defaults %v %v err %v", hookPost, hookRemind, err)
	}

	var remind, posted, samePostAt, postDefault, remindDefault bool
	if err := pool.QueryRow(ctx, `select remind, posted_at is not null, post_at = posted_at
		from guild_raid_discord_schedule where raid_id = $1`, raid).Scan(&remind, &posted, &samePostAt); err != nil {
		t.Fatal(err)
	}
	if remind || !posted || !samePostAt {
		t.Fatalf("backfill remind=%v posted=%v post_at=posted_at %v", remind, posted, samePostAt)
	}
	if err := pool.QueryRow(ctx, `select post_default, remind_default from guild_discord_installs where guild_id = $1`, gid).
		Scan(&postDefault, &remindDefault); err != nil || !postDefault || !remindDefault {
		t.Fatalf("install defaults %v %v err %v", postDefault, remindDefault, err)
	}
}
