// api/cmd/seedguild/integration_test.go
//
// The one integration test: apply against a fresh guild with a single unverified owner
// (the real production shape), assert the guild home and rankings pages return the
// expected shapes and counts, then remove and assert the guild is back to exactly that
// single unverified member with no seed rows left anywhere.
//
// Needs TEST_DATABASE_URL (api/docker-compose.test.yml); skips otherwise, the same as
// every other package's own harness_test.go.
package main

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start api/docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	if _, err := pool.Exec(context.Background(),
		`truncate users, guilds, reports, fight_metrics, seed_rows cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

// seedOwnerGuild recreates the exact production shape this tool is built against: one
// guild, one account, one character already in it as an unverified member - returns the
// guild id, the owner's user id and their character key.
func seedOwnerGuild(t *testing.T, ctx context.Context, pool *pgxpool.Pool) (guildID, ownerID int64, charKey string) {
	t.Helper()
	if err := pool.QueryRow(ctx,
		`insert into users (battletag) values ('hunthrop#1894') returning id`).Scan(&ownerID); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'pvp', 'OLYMPUS XXVII') returning id`,
	).Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	charKey = "us/pvp/obnoxious-yell"
	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, class, user_id, level)
		 values ($1, 'us', 'pvp', 'Obnoxious Yell', 'warrior', $2, 23)`, charKey, ownerID); err != nil {
		t.Fatal(err)
	}
	// A real guild_characters row only ever exists alongside an addon_exports row for
	// the same character - the export that put it there (addon.Store.PutExports writes
	// both in one transaction) - so HomeRoster's own inner join on addon_exports finds
	// it. level=23 matches the owner's real character (the task brief); no gear.
	if _, err := pool.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at)
		 values ($1, $2, 'us', 'pvp', 'Obnoxious Yell', 'FS1:1.60.1.70009:warrior:human:0/0/0:|level=23', 'addon', now())`,
		charKey, ownerID); err != nil {
		t.Fatal(err)
	}
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if err := guilds.UpsertCharacterMembership(ctx, tx, guilds.MembershipRow{
		GuildID: guildID, CharacterKey: charKey, UserID: ownerID, RankIndex: 5, Rank: "member", Source: "export",
	}); err != nil {
		t.Fatal(err)
	}
	if err := guilds.AfterGuildChange(ctx, tx, guildID, ownerID); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	return guildID, ownerID, charKey
}

func seedRowCount(t *testing.T, ctx context.Context, pool *pgxpool.Pool, tag string) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(ctx, `select count(*) from seed_rows where tag = $1`, tag).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func TestApplyThenRemoveRoundTrips(t *testing.T) {
	ctx := context.Background()
	pool := testPool(t)
	guildID, ownerID, charKey := seedOwnerGuild(t, ctx, pool)
	tag := seedTag(guildID)

	// --- apply ---
	if err := ApplySeed(ctx, pool, guildID, "hunthrop#1894"); err != nil {
		t.Fatalf("ApplySeed: %v", err)
	}

	gstore := &guilds.Store{Pool: pool}
	home, err := gstore.Home(ctx, guildID, ownerID, true, false, true, nil)
	if err != nil {
		t.Fatalf("Home: %v", err)
	}
	if len(home.Roster) != 25 {
		t.Errorf("roster len = %d, want 25 (owner + 24 mock)", len(home.Roster))
	}
	var ownerRow *guilds.RosterRow
	officers, unverified, loggedRecently := 0, 0, 0
	for i := range home.Roster {
		r := &home.Roster[i]
		if r.CharacterKey == charKey {
			ownerRow = r
		}
		if r.Rank == "officer" {
			officers++
		}
		if !r.Verified {
			unverified++
		}
		if r.LoggedRecently {
			loggedRecently++
		}
	}
	if ownerRow == nil {
		t.Fatal("owner's own character missing from the roster")
	}
	if ownerRow.Rank != "leader" || !ownerRow.Verified {
		t.Errorf("owner row = rank=%s verified=%v, want leader/true", ownerRow.Rank, ownerRow.Verified)
	}
	if officers != 2 {
		t.Errorf("officers = %d, want 2", officers)
	}
	if unverified != 3 {
		t.Errorf("unverified = %d, want 3", unverified)
	}
	if loggedRecently == 0 || loggedRecently == 24 {
		t.Errorf("logged_recently = %d of 24 mock rows, want some but not all (addon_exports.updated_at is spread over 10 days)", loggedRecently)
	}
	if len(home.Reports) == 0 {
		t.Error("home.Reports is empty, want this week's reports to have rows")
	}
	for _, r := range home.Reports {
		if r.FightCount == 0 {
			t.Errorf("report %s has no fights", r.ID)
		}
	}
	if home.Claim.State != "claimed" {
		t.Errorf("claim state = %s, want claimed", home.Claim.State)
	}

	rstore := &rankings.Store{Pool: pool}
	page, ok, err := rstore.Guild(ctx, "us", "pvp", "OLYMPUS XXVII")
	if err != nil {
		t.Fatalf("rankings.Guild: %v", err)
	}
	if !ok {
		t.Fatal("rankings.Guild: ok = false")
	}
	if len(page.Progression) == 0 {
		t.Error("progression is empty, want Onyxia's encounter")
	}
	for _, p := range page.Progression {
		if p.PullCount == 0 {
			t.Errorf("encounter %d has a zero pull count", p.EncounterID)
		}
	}
	if len(page.RosterBest) == 0 {
		t.Error("roster_best is empty, want at least one kill-fight parse per raider")
	}

	accountRowCount := seedRowCount(t, ctx, pool, tag)
	if accountRowCount == 0 {
		t.Error("seed_rows has nothing recorded for this tag after apply")
	}

	// --- idempotency: a second apply is a no-op ---
	if err := ApplySeed(ctx, pool, guildID, "hunthrop#1894"); err != nil {
		t.Fatalf("second ApplySeed: %v", err)
	}
	if got := seedRowCount(t, ctx, pool, tag); got != accountRowCount {
		t.Errorf("seed_rows count after a second apply = %d, want unchanged %d", got, accountRowCount)
	}
	var userCount int
	if err := pool.QueryRow(ctx, `select count(*) from users`).Scan(&userCount); err != nil {
		t.Fatal(err)
	}
	if userCount != 23 { // 1 owner + 22 mock accounts (two accounts carry a main and an alt)
		t.Errorf("users count after a second apply = %d, want 23 (not duplicated)", userCount)
	}

	// --- remove ---
	if err := RemoveSeed(ctx, pool, guildID, "hunthrop#1894"); err != nil {
		t.Fatalf("RemoveSeed: %v", err)
	}
	if got := seedRowCount(t, ctx, pool, tag); got != 0 {
		t.Errorf("seed_rows count after remove = %d, want 0", got)
	}
	if err := pool.QueryRow(ctx, `select count(*) from users`).Scan(&userCount); err != nil {
		t.Fatal(err)
	}
	if userCount != 1 {
		t.Errorf("users count after remove = %d, want 1 (only the owner)", userCount)
	}

	var gcCount int
	var rank, source string
	var rankIndex *int16
	var verifiedBy *string
	var verifiedAt *string
	if err := pool.QueryRow(ctx, `select count(*) from guild_characters where guild_id = $1`, guildID).Scan(&gcCount); err != nil {
		t.Fatal(err)
	}
	if gcCount != 1 {
		t.Fatalf("guild_characters rows for the guild after remove = %d, want 1", gcCount)
	}
	if err := pool.QueryRow(ctx,
		`select rank, source, rank_index, verified_by, verified_at::text
		 from guild_characters where guild_id = $1 and character_key = $2`,
		guildID, charKey).Scan(&rank, &source, &rankIndex, &verifiedBy, &verifiedAt); err != nil {
		t.Fatal(err)
	}
	if rank != "member" || source != "export" || verifiedBy != nil || verifiedAt != nil {
		t.Errorf("owner's guild_characters row after remove = rank=%s source=%s verified_by=%v verified_at=%v, "+
			"want member/export/nil/nil", rank, source, verifiedBy, verifiedAt)
	}
	if rankIndex == nil || *rankIndex != 5 {
		t.Errorf("owner's rank_index after remove = %v, want 5 (its original value)", rankIndex)
	}

	var gmCount int
	if err := pool.QueryRow(ctx, `select count(*) from guild_members where guild_id = $1`, guildID).Scan(&gmCount); err != nil {
		t.Fatal(err)
	}
	if gmCount != 1 {
		t.Errorf("guild_members rows for the guild after remove = %d, want 1", gmCount)
	}

	var claimedBy *int64
	var officerMax int
	if err := pool.QueryRow(ctx,
		`select claimed_by, officer_max_rank_index from guilds where id = $1`, guildID,
	).Scan(&claimedBy, &officerMax); err != nil {
		t.Fatal(err)
	}
	if claimedBy != nil {
		t.Errorf("guilds.claimed_by after remove = %v, want nil (unclaimed)", *claimedBy)
	}
	if officerMax != 1 {
		t.Errorf("guilds.officer_max_rank_index after remove = %d, want 1 (its original default)", officerMax)
	}

	var reportCount, fmCount int
	if err := pool.QueryRow(ctx, `select count(*) from reports where guild_id = $1`, guildID).Scan(&reportCount); err != nil {
		t.Fatal(err)
	}
	if reportCount != 0 {
		t.Errorf("reports for the guild after remove = %d, want 0", reportCount)
	}
	if err := pool.QueryRow(ctx,
		`select count(*) from fight_metrics m where exists (
		   select 1 from reports r where r.id = m.report_id and r.guild_id = $1)`, guildID,
	).Scan(&fmCount); err != nil {
		t.Fatal(err)
	}
	if fmCount != 0 {
		t.Errorf("fight_metrics for the guild's reports after remove = %d, want 0", fmCount)
	}
}
