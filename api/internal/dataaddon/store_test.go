// api/internal/dataaddon/store_test.go
package dataaddon

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// windowParses keeps the 90-day window, public reports and sufficient cards, and names each
// boss by its fight (else the encounter's newest named fight), keeping the best overall.
func TestWindowParsesKeepsTheWindowVisibilityAndSufficiencyRules(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)
	t.Cleanup(func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id like 'dataaddon-store-%'`)
		pool.Exec(bg, `delete from fights where report_id like 'dataaddon-store-%'`)
		pool.Exec(bg, `delete from reports where id like 'dataaddon-store-%'`)
	})
	for id, visibility := range map[string]string{"dataaddon-store-public": "public", "dataaddon-store-private": "private"} {
		if _, err := pool.Exec(ctx, `insert into reports (id, visibility, status, created_at) values ($1, $2, 'complete', $3)`,
			id, visibility, now.Add(-2*24*time.Hour)); err != nil {
			t.Fatal(err)
		}
	}
	fight := func(index, encounter int, name string, at time.Time) {
		if _, err := pool.Exec(ctx, `insert into fights (report_id, fight_index, encounter_id, name, kill, start_ms)
			values ('dataaddon-store-public', $1, $2, $3, true, $4)`, index, encounter, name, at.UnixMilli()); err != nil {
			t.Fatal(err)
		}
	}
	fight(1, 990801, "Lucifron", now.Add(-3*24*time.Hour))
	fight(2, 990801, "", now.Add(-2*24*time.Hour)) // unnamed: the encounter's newest name
	fight(3, 990802, "Magmadar", now.Add(-2*24*time.Hour))

	score := func(report string, index int, key string, encounter int, overall float64, insufficient bool, at time.Time) {
		if err := db.EnsureRatingsPartition(ctx, pool, at); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx,
			`insert into rating_scores (report_id, fight_index, player_key, player_name, encounter_id,
			   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at, insufficient)
			 values ($1, $2, $3, $3, $4, true, $5, $5, false, '[]', 'test', $6, $7)`,
			report, index, key, encounter, overall, at, insufficient); err != nil {
			t.Fatal(err)
		}
	}
	const thoradin = "us/normal/dataaddon-thoradin"
	at := now.Add(-2 * 24 * time.Hour)
	score("dataaddon-store-public", 1, thoradin, 990801, 64.4, false, at)
	score("dataaddon-store-public", 2, thoradin, 990801, 71.2, false, at)                       // the better Lucifron, named by fight 1
	score("dataaddon-store-public", 3, thoradin, 990802, 99, true, at)                          // insufficient: not a score
	score("dataaddon-store-public", 4, thoradin, 990802, 99, false, now.Add(-100*24*time.Hour)) // out of window
	score("dataaddon-store-private", 1, thoradin, 990802, 99, false, at)                        // private report
	score("dataaddon-store-public", 3, "us/normal/dataaddon-onlyinsufficient", 990802, 0, true, at)

	got, err := (&Store{Pool: pool}).windowParses(ctx, now.Add(-ratingWindow))
	if err != nil {
		t.Fatal(err)
	}
	if bosses := got[thoradin]; len(bosses) != 1 || bosses["Lucifron"] != 71.2 {
		t.Errorf("thoradin = %v, want only Lucifron at 71.2", bosses)
	}
	if _, ok := got["us/normal/dataaddon-onlyinsufficient"]; ok {
		t.Error("a character whose only card is insufficient has no rated fight in the window")
	}
}

func TestStoredNamesPreferTheCharacterRowThenTheRatedRow(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	at := time.Date(2026, 9, 20, 4, 0, 0, 0, time.UTC)
	t.Cleanup(func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id = 'dataaddon-store-names'`)
		pool.Exec(bg, `delete from characters where key = 'us/normal/dataaddon-lady-sunwick'`)
		pool.Exec(bg, `delete from reports where id = 'dataaddon-store-names'`)
	})
	if _, err := pool.Exec(ctx, `insert into characters (key, region, ruleset, name) values
		('us/normal/dataaddon-lady-sunwick', 'us', 'normal', 'Lady Sunwick')`); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into reports (id, visibility, status, created_at) values
		('dataaddon-store-names', 'public', 'complete', $1)`, at); err != nil {
		t.Fatal(err)
	}
	if err := db.EnsureRatingsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into rating_scores (report_id, fight_index, player_key, player_name, encounter_id,
		   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at)
		 values ('dataaddon-store-names', 1, 'eu/pvp/dataaddon-mörk', 'Mörk-Pvp', 1, false, 50, 50, false, '[]', 'test', $1)`,
		at); err != nil {
		t.Fatal(err)
	}
	got, err := (&Store{Pool: pool}).storedNames(ctx,
		[]string{"us/normal/dataaddon-lady-sunwick", "eu/pvp/dataaddon-mörk", "us/normal/dataaddon-nobody"})
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"us/normal/dataaddon-lady-sunwick": "Lady Sunwick",
		"eu/pvp/dataaddon-mörk":            "Mörk-Pvp",
		"us/normal/dataaddon-nobody":       "",
	}
	for key, name := range want {
		if got[key] != name {
			t.Errorf("%s: %q, want %q", key, got[key], name)
		}
	}
}

func TestGuildQueriesReadVerifiedMembersNightsAndKills(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)

	var u1 int64
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('One') returning id`).Scan(&u1); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from users where id = $1`, u1) })

	var guildID int64
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'Dataaddon Vanguard') returning id`,
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
	insertReport("dataaddon-guild-old", now.Add(-100*24*time.Hour))
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
	insertFight("dataaddon-guild-old", 1, 103, true) // out of window
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from fights where report_id like 'dataaddon-guild-%'`)
	})

	store := &Store{Pool: pool}
	identities, err := store.guildsWithVerifiedMembers(ctx)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, g := range identities {
		found = found || g.ID == guildID
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

	since := now.Add(-ratingWindow)
	nights, err := store.nightsByGuild(ctx, []int64{guildID}, since)
	if err != nil {
		t.Fatal(err)
	}
	if nights[guildID] != 2 {
		t.Errorf("nights = %d, want 2", nights[guildID])
	}

	kills, err := store.bossesKilledByGuild(ctx, []int64{guildID}, since)
	if err != nil {
		t.Fatal(err)
	}
	if kills[guildID] != 2 {
		t.Errorf("bosses killed in the window = %d, want 2 (100 and 102; 101 was a wipe, 103 is too old)", kills[guildID])
	}
}
