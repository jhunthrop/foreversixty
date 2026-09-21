// api/internal/guilds/jobs_test.go
package guilds

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestAgeOutRemovesStaleCharactersAndRecomputesTheAccount(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "stale@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/stale", "member", true)
	if _, err := pool.Exec(ctx,
		`update guild_characters set refreshed_at = now() - interval '46 days' where character_key = 'us/hardcore/stale'`); err != nil {
		t.Fatal(err)
	}
	tx, _ := pool.Begin(ctx)
	if err := RecomputeMembership(ctx, tx, gid, &uid); err != nil {
		t.Fatal(err)
	}
	_ = tx.Commit(ctx)

	if err := s.AgeOut(ctx, time.Now()); err != nil {
		t.Fatal(err)
	}
	var n int
	pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/hardcore/stale'`).Scan(&n)
	if n != 0 {
		t.Fatal("a character stale for more than 45 days should be aged out")
	}
	pool.QueryRow(ctx, `select count(*) from guild_members where guild_id = $1 and user_id = $2`, gid, uid).Scan(&n)
	if n != 0 {
		t.Fatal("guild_members should follow the aged-out character")
	}
}

func TestAgeOutLeavesFreshCharactersAlone(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "fresh@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/fresh", "member", true)

	if err := s.AgeOut(ctx, time.Now()); err != nil {
		t.Fatal(err)
	}
	var n int
	pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/hardcore/fresh'`).Scan(&n)
	if n != 1 {
		t.Fatal("a character refreshed today must not be aged out")
	}
}

func TestAgeOutReleasesAClaimTheAccountNoLongerQualifiesFor(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "aging-leader@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/agingleader", "leader", true)
	if _, err := pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, uid, gid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`update guild_characters set refreshed_at = now() - interval '46 days' where character_key = 'us/hardcore/agingleader'`); err != nil {
		t.Fatal(err)
	}

	if err := s.AgeOut(ctx, time.Now()); err != nil {
		t.Fatal(err)
	}
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy != nil {
		t.Fatal("ageing out a claimant's last character should release the claim, exactly as leaving would")
	}
}

func TestVerifyByLogsRequiresTwoDistinctNightsWithinThirtyDays(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "logger@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/logger", "member", false)
	// Owned by a different account than the character being verified -
	// item 3's fixed independence rule (fourth security review
	// response) never counts a report toward its own account's
	// verification.
	uploader := seedUser(t, pool, "logger-uploader@example.com")

	seedReportWithFight := func(id string, createdAt time.Time, players []string) {
		t.Helper()
		if _, err := pool.Exec(ctx,
			`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
			 values ($1, $2, $3, 'guild', 'complete', $4)`, id, uploader, gid, createdAt); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx,
			`insert into fights (report_id, fight_index, players) values ($1, 0, $2)`, id, players); err != nil {
			t.Fatal(err)
		}
	}
	seedReportWithFight("r1spoofnight1", time.Now().Add(-20*24*time.Hour), []string{"us/hardcore/logger"})

	if err := s.VerifyByLogs(ctx); err != nil {
		t.Fatal(err)
	}
	var verified bool
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/logger'`).Scan(&verified)
	if verified {
		t.Fatal("one appearance must not verify the character")
	}

	seedReportWithFight("r2spoofnight2", time.Now().Add(-5*24*time.Hour), []string{"us/hardcore/logger"})
	if err := s.VerifyByLogs(ctx); err != nil {
		t.Fatal(err)
	}
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/logger'`).Scan(&verified)
	if !verified {
		t.Fatal("two distinct report dates within 30 days should verify the character")
	}
	var memberVerified bool
	pool.QueryRow(ctx, `select verified_at is not null from guild_members where guild_id = $1 and user_id = $2`, gid, uid).Scan(&memberVerified)
	if !memberVerified {
		t.Fatal("guild_members should be recomputed as verified once the character is")
	}
}

func TestVerifyByLogsRefusesTwoAppearancesMoreThanThirtyDaysApart(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "stale-logger@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/staleoverlap", "member", false)
	uploader := seedUser(t, pool, "stale-logger-uploader@example.com")

	for i, days := range []int{40, 5} {
		id := "far" + string(rune('a'+i))
		if _, err := pool.Exec(ctx,
			`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
			 values ($1, $2, $3, 'guild', 'complete', now() - ($4::int * interval '1 day'))`,
			id, uploader, gid, days); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx,
			`insert into fights (report_id, fight_index, players) values ($1, 0, $2)`,
			id, []string{"us/hardcore/staleoverlap"}); err != nil {
			t.Fatal(err)
		}
	}
	if err := s.VerifyByLogs(ctx); err != nil {
		t.Fatal(err)
	}
	var verified bool
	pool.QueryRow(ctx,
		`select verified_at is not null from guild_characters where character_key = 'us/hardcore/staleoverlap'`).Scan(&verified)
	if verified {
		t.Fatal("only the trailing-30-days appearance counts, so the 40-day-old one falls outside the window")
	}
}

func TestMembershipJobRunsBothSweepsOnStartup(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	uid := seedUser(t, pool, "job@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/job", "member", true)
	pool.Exec(ctx, `update guild_characters set refreshed_at = now() - interval '46 days'`)

	job := &MembershipJob{Store: s, Every: time.Hour}
	if err := job.Run(ctx); err != nil {
		t.Fatal(err)
	}
	var n int
	pool.QueryRow(ctx, `select count(*) from guild_characters`).Scan(&n)
	if n != 0 {
		t.Fatal("Run should perform the first sweep pass synchronously before returning")
	}
}

func TestVerifyByLogsRecordsVerifiedByLogs(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "logs-source@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/logssource", "member", false)
	uploader := seedUser(t, pool, "logs-source-uploader@example.com")

	first := time.Now().Add(-20 * 24 * time.Hour)
	for i, days := range []float64{20, 5} {
		id := fmt.Sprintf("logsource%d", i)
		if _, err := pool.Exec(ctx,
			`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
			 values ($1, $2, $3, 'guild', 'complete', $4)`,
			id, uploader, gid, first.Add(time.Duration(20-days)*24*time.Hour)); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx,
			`insert into fights (report_id, fight_index, players) values ($1, 0, $2)`,
			id, []string{"us/hardcore/logssource"}); err != nil {
			t.Fatal(err)
		}
	}
	if err := s.VerifyByLogs(ctx); err != nil {
		t.Fatal(err)
	}
	var by string
	if err := pool.QueryRow(ctx,
		`select verified_by from guild_characters where character_key = 'us/hardcore/logssource'`).Scan(&by); err != nil {
		t.Fatal(err)
	}
	if by != "logs" {
		t.Fatalf("verified_by = %q, want logs", by)
	}
}

// TestVerifyByLogsVerifiesNobodyInAContestedGuild is item 3's second
// half (fourth security review response): while a guild's claim is
// contested, VerifyByLogs verifies nobody in that guild at all - since
// approve is already frozen for the same reason, membership of a
// disputed guild cannot change under a moderator's feet while they are
// looking at the dispute.
func TestVerifyByLogsVerifiesNobodyInAContestedGuild(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	uid := seedUser(t, pool, "contested-logger@example.com")
	seedCharacter(t, pool, gid, uid, "us/hardcore/contestedlogger", "member", false)
	uploader := seedUser(t, pool, "contested-logger-uploader@example.com")

	first := time.Now().Add(-20 * 24 * time.Hour)
	for i, offset := range []time.Duration{0, 15 * 24 * time.Hour} {
		id := fmt.Sprintf("contestlog%d", i)
		if _, err := pool.Exec(ctx,
			`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
			 values ($1, $2, $3, 'guild', 'complete', $4)`,
			id, uploader, gid, first.Add(offset)); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx,
			`insert into fights (report_id, fight_index, players) values ($1, 0, $2)`,
			id, []string{"us/hardcore/contestedlogger"}); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := pool.Exec(ctx, `update guilds set claim_contested_at = now() where id = $1`, gid); err != nil {
		t.Fatal(err)
	}

	if err := s.VerifyByLogs(ctx); err != nil {
		t.Fatal(err)
	}
	var verified bool
	pool.QueryRow(ctx,
		`select verified_at is not null from guild_characters where character_key = 'us/hardcore/contestedlogger'`).
		Scan(&verified)
	if verified {
		t.Fatal("VerifyByLogs must verify nobody in a contested guild, even with two genuinely independent qualifying reports")
	}
}
