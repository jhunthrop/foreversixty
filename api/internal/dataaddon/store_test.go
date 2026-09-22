// api/internal/dataaddon/store_test.go
package dataaddon

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func TestCharacterFightsAppliesTheWindowVisibilityAndAnonymizeRules(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)

	var normalUser, anonUser int64
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('Normal') returning id`).Scan(&normalUser); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into users (battletag, anonymize) values ('Hidden', true) returning id`).Scan(&anonUser); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from users where id in ($1, $2)`, normalUser, anonUser)
	})

	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id) values
		 ('us/normal/thoradin', 'us', 'normal', 'Thoradin', $1),
		 ('us/normal/hiddenhero', 'us', 'normal', 'Hiddenhero', $2)`, normalUser, anonUser); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from characters where key in ('us/normal/thoradin', 'us/normal/hiddenhero')`)
	})

	insertReport := func(id, visibility string, createdAt time.Time) {
		if _, err := pool.Exec(ctx,
			`insert into reports (id, visibility, status, created_at) values ($1, $2, 'complete', $3)`,
			id, visibility, createdAt); err != nil {
			t.Fatal(err)
		}
	}
	insertReport("dataaddon-store-public", "public", now.Add(-2*24*time.Hour))
	insertReport("dataaddon-store-private", "private", now.Add(-2*24*time.Hour))
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from reports where id in ('dataaddon-store-public', 'dataaddon-store-private')`)
	})

	insertScore := func(reportID string, index int, playerKey string, overall float64, foughtAt time.Time) {
		// rating_scores is partitioned by month on fought_at (migration
		// 0022): a plain insert with no matching partition fails outright,
		// so every write through this test helper ensures one first, the
		// same way rating.Store.RateFight does before its own insert.
		if err := db.EnsureRatingsPartition(ctx, pool, foughtAt); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx,
			`insert into rating_scores (report_id, fight_index, player_key, player_name, encounter_id,
			   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at)
			 values ($1, $2, $3, $3, 1, false, $4, $4, false, '[]', 'test', $5)`,
			reportID, index, playerKey, overall, foughtAt); err != nil {
			t.Fatal(err)
		}
	}
	insertScore("dataaddon-store-public", 1, "us/normal/thoradin", 80, now.Add(-2*24*time.Hour))    // in window, public: counts
	insertScore("dataaddon-store-public", 2, "us/normal/thoradin", 999, now.Add(-100*24*time.Hour)) // out of window: excluded
	insertScore("dataaddon-store-private", 1, "us/normal/thoradin", 999, now.Add(-1*24*time.Hour))  // private report: excluded
	insertScore("dataaddon-store-public", 3, "us/normal/hiddenhero", 999, now.Add(-2*24*time.Hour)) // anonymized owner: excluded
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from rating_scores where report_id in ('dataaddon-store-public', 'dataaddon-store-private')`)
	})

	store := &Store{Pool: pool}
	rows, err := store.characterFights(ctx, now.Add(-ratingWindow))
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("rows = %d, want exactly 1 (thoradin's in-window public fight); got %+v", len(rows), rows)
	}
	if rows[0].PlayerKey != "us/normal/thoradin" || rows[0].Overall != 80 {
		t.Errorf("row = %+v", rows[0])
	}
}

func TestGuildQueriesReadVerifiedMembersNightsAndProgression(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)

	var u1, u2 int64
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('One') returning id`).Scan(&u1); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('Two') returning id`).Scan(&u2); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from users where id in ($1, $2)`, u1, u2) })

	var guildID int64
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'Iron Vanguard') returning id`,
	).Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from guilds where id = $1`, guildID) })

	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, verified_at) values
		 ($1, 'us/normal/thoradin', $2, now()),
		 ($1, 'us/normal/notyet', $2, null)`, guildID, u1); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from guild_characters where guild_id = $1`, guildID) })

	insertReport := func(id string, createdAt time.Time) {
		if _, err := pool.Exec(ctx,
			`insert into reports (id, guild_id, visibility, status, created_at) values ($1, $2, 'public', 'complete', $3)`,
			id, guildID, createdAt); err != nil {
			t.Fatal(err)
		}
	}
	insertReport("dataaddon-guild-1", now.Add(-2*24*time.Hour))
	insertReport("dataaddon-guild-2", now.Add(-10*24*time.Hour))
	t.Cleanup(func() { pool.Exec(ctx, `delete from reports where guild_id = $1`, guildID) })

	insertFight := func(reportID string, index int, encounterID int64, kill bool) {
		if _, err := pool.Exec(ctx,
			`insert into fights (report_id, fight_index, encounter_id, kill) values ($1, $2, $3, $4)`,
			reportID, index, encounterID, kill); err != nil {
			t.Fatal(err)
		}
	}
	insertFight("dataaddon-guild-1", 1, 100, true)
	insertFight("dataaddon-guild-1", 2, 101, false)
	insertFight("dataaddon-guild-2", 1, 100, true)
	insertFight("dataaddon-guild-2", 2, 102, true)
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from fights where report_id in ('dataaddon-guild-1', 'dataaddon-guild-2')`)
	})

	store := &Store{Pool: pool}
	identities, err := store.guildsWithVerifiedMembers(ctx)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, g := range identities {
		if g.ID == guildID {
			found = true
		}
	}
	if !found {
		t.Fatal("the guild with a verified member should be listed")
	}

	members, err := store.verifiedMembers(ctx, []int64{guildID})
	if err != nil {
		t.Fatal(err)
	}
	if len(members[guildID]) != 1 || members[guildID][0] != "us/normal/thoradin" {
		t.Errorf("verified members = %v, want exactly [us/normal/thoradin]", members[guildID])
	}

	nights, err := store.nightsByGuild(ctx, []int64{guildID}, now.Add(-ratingWindow))
	if err != nil {
		t.Fatal(err)
	}
	if nights[guildID] != 2 {
		t.Errorf("nights = %d, want 2", nights[guildID])
	}

	killed, total, err := store.progressionByGuild(ctx, []int64{guildID})
	if err != nil {
		t.Fatal(err)
	}
	if killed[guildID] != 2 || total[guildID] != 3 {
		t.Errorf("progression = %d/%d, want 2/3", killed[guildID], total[guildID])
	}
}
