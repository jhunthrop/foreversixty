// api/internal/guilds/contest_test.go
package guilds

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestContestClaimRequiresAnActiveClaim(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	contester := seedUser(t, pool, "contester1@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/contester1", "officer", false)

	if err := s.ContestClaim(ctx, gid, contester, true); !errors.Is(err, ErrNoActiveClaim) {
		t.Fatalf("contesting an unclaimed guild = %v, want ErrNoActiveClaim", err)
	}
}

func TestContestClaimNeedsARawOfficerOrLeaderCharacterOnADifferentAccount(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/claimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}

	if err := s.ContestClaim(ctx, gid, claimant, true); !errors.Is(err, ErrSameAccount) {
		t.Fatalf("the claimant contesting their own claim = %v, want ErrSameAccount", err)
	}

	noStanding := seedUser(t, pool, "nostanding@example.com")
	seedCharacter(t, pool, gid, noStanding, "us/hardcore/nostanding", "member", false)
	if err := s.ContestClaim(ctx, gid, noStanding, true); !errors.Is(err, ErrNotEligible) {
		t.Fatalf("a plain member contesting = %v, want ErrNotEligible", err)
	}

	contester := seedUser(t, pool, "realcontester@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/realcontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester, true); err != nil {
		t.Fatal(err)
	}
	var contestedBy *int64
	pool.QueryRow(ctx, `select claim_contested_by from guilds where id = $1`, gid).Scan(&contestedBy)
	if contestedBy == nil || *contestedBy != contester {
		t.Fatalf("claim_contested_by = %v, want %d", contestedBy, contester)
	}

	second := seedUser(t, pool, "secondcontester@example.com")
	seedCharacter(t, pool, gid, second, "us/hardcore/secondcontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, second, true); !errors.Is(err, ErrAlreadyContested) {
		t.Fatalf("contesting an already-contested claim = %v, want ErrAlreadyContested", err)
	}
}

// TestContestClaimRefusesAnAccountWithNoBattleNetIdentity is A5's first
// exploit step: an email-only account with a single forged officer
// export must never be able to open a contest, let alone freeze a real
// guild's officer tools (A1).
func TestContestClaimRefusesAnAccountWithNoBattleNetIdentity(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "young-claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/youngclaimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}

	attacker := seedUser(t, pool, "email-only-attacker@example.com")
	seedCharacter(t, pool, gid, attacker, "us/hardcore/attacker", "officer", false)
	if err := s.ContestClaim(ctx, gid, attacker, false); !errors.Is(err, ErrNoBattleNetIdentity) {
		t.Fatalf("contest with no Battle.net identity = %v, want ErrNoBattleNetIdentity", err)
	}
	var contestedAt any
	pool.QueryRow(ctx, `select claim_contested_at from guilds where id = $1`, gid).Scan(&contestedAt)
	if contestedAt != nil {
		t.Fatal("a refused contest must never mark the claim contested")
	}
}

// TestContestClaimRateLimitsToOnePerAccountPerThirtyDays is A2/A5: a
// Battle.net account's second contest attempt inside 30 days is
// refused, even against a different guild.
func TestContestClaimRateLimitsToOnePerAccountPerThirtyDays(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()

	g1 := seedGuild(t, pool, "RateOne")
	claimant1 := seedUser(t, pool, "rate-claimant1@example.com")
	seedCharacter(t, pool, g1, claimant1, "us/hardcore/rateclaimant1", "leader", false)
	if _, err := s.Claim(ctx, g1, claimant1, true); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "rate-contester@example.com")
	seedCharacter(t, pool, g1, contester, "us/hardcore/ratecontester", "officer", false)
	if err := s.ContestClaim(ctx, g1, contester, true); err != nil {
		t.Fatal(err)
	}
	if err := s.ResolveClaim(ctx, g1, seedUser(t, pool, "rate-mod1@example.com"), "uphold"); err != nil {
		t.Fatal(err)
	}

	g2 := seedGuild(t, pool, "RateTwo")
	claimant2 := seedUser(t, pool, "rate-claimant2@example.com")
	seedCharacter(t, pool, g2, claimant2, "us/hardcore/rateclaimant2", "leader", false)
	if _, err := s.Claim(ctx, g2, claimant2, true); err != nil {
		t.Fatal(err)
	}
	// The contester's own character in g2, so eligibility is not what
	// blocks the second attempt.
	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank) values ($1, 'us/hardcore/ratecontesteralt', $2, 'officer')`,
		g2, contester); err != nil {
		t.Fatal(err)
	}
	if err := s.ContestClaim(ctx, g2, contester, true); !errors.Is(err, ErrContestRateLimited) {
		t.Fatalf("a second contest within 30 days on a different guild = %v, want ErrContestRateLimited", err)
	}
}

// TestContestClaimRefusesASecondOpenContestOnAnotherGuild is A2: at
// most one OPEN contest per account across every guild.
func TestContestClaimRefusesASecondOpenContestOnAnotherGuild(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()

	g1 := seedGuild(t, pool, "OpenOne")
	claimant1 := seedUser(t, pool, "open-claimant1@example.com")
	seedCharacter(t, pool, g1, claimant1, "us/hardcore/openclaimant1", "leader", false)
	if _, err := s.Claim(ctx, g1, claimant1, true); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "open-contester@example.com")
	seedCharacter(t, pool, g1, contester, "us/hardcore/opencontester", "officer", false)
	if err := s.ContestClaim(ctx, g1, contester, true); err != nil {
		t.Fatal(err)
	}

	g2 := seedGuild(t, pool, "OpenTwo")
	claimant2 := seedUser(t, pool, "open-claimant2@example.com")
	seedCharacter(t, pool, g2, claimant2, "us/hardcore/openclaimant2", "leader", false)
	if _, err := s.Claim(ctx, g2, claimant2, true); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank) values ($1, 'us/hardcore/opencontesteralt', $2, 'officer')`,
		g2, contester); err != nil {
		t.Fatal(err)
	}
	if err := s.ContestClaim(ctx, g2, contester, true); !errors.Is(err, ErrAlreadyContestingAnotherGuild) {
		t.Fatalf("a second open contest while one is already open = %v, want ErrAlreadyContestingAnotherGuild", err)
	}
}

// TestContestClaimRefusesARepeatAfterAnUphold is A3: once a moderator
// has upheld a claim against a particular contester, that same account
// may not re-contest the same guild, and the forged row that lost the
// dispute is gone (its unverified guild_characters row was deleted on
// uphold).
func TestContestClaimRefusesARepeatAfterAnUphold(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "upheld-again-claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/upheldagainclaimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "upheld-again-contester@example.com")
	// The forged, unverified row the contest itself relied on...
	seedCharacter(t, pool, gid, contester, "us/hardcore/upheldagaincontester", "officer", false)
	// ...plus a second, VERIFIED officer row on the same account, so the
	// contester still has eligibility standing after the uphold deletes
	// the unverified one - otherwise a bare re-contest attempt would
	// already be refused by ErrNotEligible before ever reaching the
	// already-upheld check this test targets.
	seedCharacter(t, pool, gid, contester, "us/hardcore/upheldagaincontestersurvivor", "officer", true)
	if err := s.ContestClaim(ctx, gid, contester, true); err != nil {
		t.Fatal(err)
	}
	moderator := seedUser(t, pool, "upheld-again-mod@example.com")
	if err := s.ResolveClaim(ctx, gid, moderator, "uphold"); err != nil {
		t.Fatal(err)
	}
	var forgedRows int
	pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/hardcore/upheldagaincontester'`).Scan(&forgedRows)
	if forgedRows != 0 {
		t.Fatal("the contester's unverified row should be gone after the contest is upheld against them")
	}
	var survivorRows int
	pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/hardcore/upheldagaincontestersurvivor'`).Scan(&survivorRows)
	if survivorRows != 1 {
		t.Fatal("the contester's VERIFIED row must survive the uphold")
	}

	if err := s.ContestClaim(ctx, gid, contester, true); !errors.Is(err, ErrContestAlreadyUpheld) {
		t.Fatalf("re-contesting after an uphold = %v, want ErrContestAlreadyUpheld", err)
	}
}

// TestResolveClaimUpholdKeepsAVerifiedContesterRow is A3's parenthetical:
// a VERIFIED row of the losing contester survives an uphold - they are
// a real member who lost a dispute, not an impostor.
func TestResolveClaimUpholdKeepsAVerifiedContesterRow(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "keep-verified-claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/keepverifiedclaimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "keep-verified-contester@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/keepverifiedofficer", "officer", true)
	if err := s.ContestClaim(ctx, gid, contester, true); err != nil {
		t.Fatal(err)
	}
	moderator := seedUser(t, pool, "keep-verified-mod@example.com")
	if err := s.ResolveClaim(ctx, gid, moderator, "uphold"); err != nil {
		t.Fatal(err)
	}
	var stillThere bool
	pool.QueryRow(ctx,
		`select exists(select 1 from guild_characters where character_key = 'us/hardcore/keepverifiedofficer')`).
		Scan(&stillThere)
	if !stillThere {
		t.Fatal("the contester's own VERIFIED row must survive an uphold - they lost a dispute, not their membership")
	}
}

func TestResolveClaimUphold(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "upheld-claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/upheldclaimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "upheld-contester@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/upheldcontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester, true); err != nil {
		t.Fatal(err)
	}

	moderator := seedUser(t, pool, "uphold-mod@example.com")
	if err := s.ResolveClaim(ctx, gid, moderator, "uphold"); err != nil {
		t.Fatal(err)
	}
	var claimedBy *int64
	var contestedAt any
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	pool.QueryRow(ctx, `select claim_contested_at from guilds where id = $1`, gid).Scan(&contestedAt)
	if claimedBy == nil || *claimedBy != claimant {
		t.Fatalf("claimed_by = %v, want unchanged (%d)", claimedBy, claimant)
	}
	if contestedAt != nil {
		t.Fatal("claim_contested_at should be cleared after an uphold")
	}
	var recordedModerator int64
	if err := pool.QueryRow(ctx,
		`select moderator_id from guild_claim_resolutions where guild_id = $1 and contester_id = $2`,
		gid, contester).Scan(&recordedModerator); err != nil {
		t.Fatal(err)
	}
	if recordedModerator != moderator {
		t.Fatalf("recorded moderator_id = %d, want %d", recordedModerator, moderator)
	}
}

func TestResolveClaimReleaseUnverifiesOnlyClaimSourcedRows(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "released-claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/releasedclaimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	// A second, officer-verified row on a DIFFERENT account, verified via
	// approval rather than the claim - must survive the release.
	other := seedUser(t, pool, "officer-verified@example.com")
	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank, verified_at, verified_by)
		 values ($1, 'us/hardcore/officerverified', $2, 'member', now(), 'officer')`, gid, other); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "release-contester@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/releasecontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester, true); err != nil {
		t.Fatal(err)
	}

	moderator := seedUser(t, pool, "release-mod@example.com")
	if err := s.ResolveClaim(ctx, gid, moderator, "release"); err != nil {
		t.Fatal(err)
	}
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy != nil {
		t.Fatal("claimed_by should be cleared after a release")
	}
	var claimantVerified, otherVerified bool
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/releasedclaimant'`).Scan(&claimantVerified)
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/officerverified'`).Scan(&otherVerified)
	if claimantVerified {
		t.Fatal("the claim-sourced row should be un-verified after a release")
	}
	if !otherVerified {
		t.Fatal("the officer-sourced row must survive the release untouched")
	}
}

func TestResolveClaimTransferMovesTheClaimAndVerification(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "transferred-from@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/transferredfrom", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "transferred-to@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/transferredto", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester, true); err != nil {
		t.Fatal(err)
	}

	moderator := seedUser(t, pool, "transfer-mod@example.com")
	if err := s.ResolveClaim(ctx, gid, moderator, "transfer"); err != nil {
		t.Fatal(err)
	}
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy == nil || *claimedBy != contester {
		t.Fatalf("claimed_by = %v, want the contesting account %d", claimedBy, contester)
	}
	var contesterVerified, originalVerified bool
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/transferredto'`).Scan(&contesterVerified)
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/transferredfrom'`).Scan(&originalVerified)
	if !contesterVerified {
		t.Fatal("the new claimant's character should be verified")
	}
	if originalVerified {
		t.Fatal("the original claimant's claim-sourced verification should not survive the transfer")
	}
}

func TestResolveClaimNeverUnverifiesAnUnrelatedAccountsStaleClaimSourcedRow(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")

	// An unrelated account with a stale claim-sourced row from some
	// earlier, already-released claim - no live claim references it.
	stale := seedUser(t, pool, "stale-claim-source@example.com")
	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank, verified_at, verified_by)
		 values ($1, 'us/hardcore/staleclaimsource', $2, 'member', now(), 'claim')`, gid, stale); err != nil {
		t.Fatal(err)
	}

	claimant := seedUser(t, pool, "unrelated-resolve-claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/unrelatedresolveclaimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "unrelated-resolve-contester@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/unrelatedresolvecontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester, true); err != nil {
		t.Fatal(err)
	}

	moderator := seedUser(t, pool, "unrelated-resolve-mod@example.com")
	if err := s.ResolveClaim(ctx, gid, moderator, "release"); err != nil {
		t.Fatal(err)
	}
	var staleVerified bool
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/staleclaimsource'`).Scan(&staleVerified)
	if !staleVerified {
		t.Fatal("resolving a different account's claim must never touch an unrelated account's stale claim-sourced verification")
	}
}

func TestClaimStateReflectsEachPhase(t *testing.T) {
	now := time.Now()
	unclaimed := Guild{}
	if got := claimState(unclaimed, now); got.State != "unclaimed" {
		t.Fatalf("state = %q, want unclaimed", got.State)
	}
	pendingBy := int64(1)
	requested := now
	pending := Guild{ClaimPendingBy: &pendingBy, ClaimRequestedAt: &requested}
	if got := claimState(pending, now); got.State != "pending" {
		t.Fatalf("state = %q, want pending", got.State)
	}
	claimedByID := int64(2)
	claimed := Guild{ClaimedBy: &claimedByID}
	if got := claimState(claimed, now); got.State != "claimed" {
		t.Fatalf("state = %q, want claimed", got.State)
	}
	contested := Guild{ClaimedBy: &claimedByID, ClaimContestedAt: &requested}
	if got := claimState(contested, now); got.State != "contested" {
		t.Fatalf("state = %q, want contested", got.State)
	}
}

// TestAContestAlwaysFreezes is the fourth security review response's
// simplification: a contest freezes the disputed claimant's officer
// tools unconditionally, with no exception for a young, an established,
// or an otherwise "corroborated" claim - two narrower rules (a 14-day
// young-claim test, then an independence-checked corroboration test)
// were each found gameable by a squatter in turn, so this response
// deletes the escape rather than patch it a third time.
func TestAContestAlwaysFreezes(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "always-frozen-claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/alwaysfrozenclaimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	// Establish the claim well past the old 14-day threshold - under
	// the deleted rule this alone would have been enough to avoid a
	// freeze; under the current rule it changes nothing.
	if _, err := pool.Exec(ctx,
		`update guilds set claimed_at = now() - interval '90 days' where id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "always-frozen-contester@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/alwaysfrozencontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester, true); err != nil {
		t.Fatal(err)
	}

	g, err := s.getGuild(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	view := claimState(g, time.Now())
	if view.State != "contested" || !view.Frozen {
		t.Fatalf("claim view = %+v, want contested and frozen (a 90-day-old claim included)", view)
	}
}

// TestAnUpheldContestUnfreezes: resolving a contest with "uphold"
// clears claim_contested_at, so claimState no longer reports contested
// or frozen - officer tools return to the standing claimant.
func TestAnUpheldContestUnfreezes(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "unfreeze-claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/unfreezeclaimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "unfreeze-contester@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/unfreezecontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester, true); err != nil {
		t.Fatal(err)
	}
	g, err := s.getGuild(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	if view := claimState(g, time.Now()); !view.Frozen {
		t.Fatalf("claim view = %+v, want frozen before the uphold", view)
	}

	moderator := seedUser(t, pool, "unfreeze-mod@example.com")
	if err := s.ResolveClaim(ctx, gid, moderator, "uphold"); err != nil {
		t.Fatal(err)
	}
	g, err = s.getGuild(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	view := claimState(g, time.Now())
	if view.State != "claimed" || view.Frozen {
		t.Fatalf("claim view after uphold = %+v, want claimed and not frozen", view)
	}
}

// seedIndependentLogReport inserts one report+fight row naming players
// so VerifyByLogs's real job can verify them - tests go through the
// real verification path, not a raw insert of verified_by = 'logs'.
func seedIndependentLogReport(t *testing.T, pool *pgxpool.Pool, id string, ownerID, guildID int64, createdAt time.Time, players []string) {
	t.Helper()
	ctx := context.Background()
	if _, err := pool.Exec(ctx,
		`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
		 values ($1, $2, $3, 'guild', 'complete', $4)`, id, ownerID, guildID, createdAt); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into fights (report_id, fight_index, players) values ($1, 0, $2)`, id, players); err != nil {
		t.Fatal(err)
	}
}

// TestAContestFreezesRegardlessOfWhatASquatterManufactures is the
// CRITICAL finding's own exploit scenario, reproduced through the real
// Claim/VerifyByLogs/ContestClaim path: a squatter claims an unclaimed
// guild, is now its only officer, and attaches their OWN uploaded
// reports naming a throwaway account's character on two distinct
// nights. Under item 3's fixed, always-on rule (a report only fails to
// count when it is owned by the character's OWN account), this
// self-attachment actually succeeds in verifying the throwaway account
// - an accepted, explicitly out-of-scope risk (it grants the throwaway
// nothing but access to the squatter's own reports, and a moderator's
// release un-verifies every claim-derived row regardless). What matters
// here is that this manufactured "verified" appearance changes nothing
// about freezing: the contest still freezes, unconditionally.
func TestAContestFreezesRegardlessOfWhatASquatterManufactures(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")

	squatter := seedUser(t, pool, "squatter@example.com")
	seedCharacter(t, pool, gid, squatter, "us/hardcore/squatter", "leader", false)
	if _, err := s.Claim(ctx, gid, squatter, true); err != nil {
		t.Fatal(err)
	}

	throwaway := seedUser(t, pool, "throwaway@example.com")
	seedCharacter(t, pool, gid, throwaway, "us/hardcore/throwaway", "member", false)
	first := time.Now().Add(-10 * 24 * time.Hour)
	seedIndependentLogReport(t, pool, "squat-night-1", squatter, gid, first, []string{"us/hardcore/throwaway"})
	seedIndependentLogReport(t, pool, "squat-night-2", squatter, gid, first.Add(3*24*time.Hour), []string{"us/hardcore/throwaway"})
	if err := s.VerifyByLogs(ctx); err != nil {
		t.Fatal(err)
	}
	var manufacturedVerified bool
	if err := pool.QueryRow(ctx,
		`select verified_at is not null from guild_characters where character_key = 'us/hardcore/throwaway'`).
		Scan(&manufacturedVerified); err != nil {
		t.Fatal(err)
	}
	if !manufacturedVerified {
		t.Fatal("item 3's accepted rule lets a squatter verify a sockpuppet from their own reports - it should have succeeded here")
	}

	realGM := seedUser(t, pool, "real-gm@example.com")
	seedCharacter(t, pool, gid, realGM, "us/hardcore/realgm", "officer", false)
	if err := s.ContestClaim(ctx, gid, realGM, true); err != nil {
		t.Fatal(err)
	}

	g, err := s.getGuild(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	view := claimState(g, time.Now())
	if view.State != "contested" {
		t.Fatalf("state = %q, want contested", view.State)
	}
	if !view.Frozen {
		t.Fatal("a contest must freeze regardless of what the disputed claimant has manufactured in their own guild shell")
	}
}
