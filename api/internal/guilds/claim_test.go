// api/internal/guilds/claim_test.go
package guilds

import (
	"context"
	"errors"
	"sync"
	"testing"
)

func TestClaimByTheGuildMasterIsImmediate(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "gm@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/gm", "leader", false)

	result, err := s.Claim(ctx, gid, uid, true)
	if err != nil || result.Status != "confirmed" {
		t.Fatalf("Claim = %+v, %v, want confirmed", result, err)
	}
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy == nil || *claimedBy != uid {
		t.Fatalf("claimed_by = %v, want %d", claimedBy, uid)
	}
	var verified bool
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/gm'`).Scan(&verified)
	if !verified {
		t.Fatal("the GM's own claiming character should be verified immediately")
	}
}

func TestClaimByAnOfficerGoesPendingAndRequiresEligibility(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")

	stranger := seedUser(t, pool, "stranger@example.com")
	if _, err := s.Claim(ctx, gid, stranger, true); !errors.Is(err, ErrNotEligible) {
		t.Fatalf("Claim by a non-member = %v, want ErrNotEligible", err)
	}

	officer := seedUser(t, pool, "officer@example.com")
	seedCharacter(t, pool, gid, officer, "us/hardcore/officer", "officer", false)
	result, err := s.Claim(ctx, gid, officer, true)
	if err != nil || result.Status != "pending" || result.ExpiresAt == nil {
		t.Fatalf("Claim by an officer = %+v, %v, want pending with an expiry", result, err)
	}
	var pendingBy *int64
	pool.QueryRow(ctx, `select claim_pending_by from guilds where id = $1`, gid).Scan(&pendingBy)
	if pendingBy == nil || *pendingBy != officer {
		t.Fatalf("claim_pending_by = %v, want %d", pendingBy, officer)
	}

	if _, err := s.Claim(ctx, gid, officer, true); !errors.Is(err, ErrClaimPending) {
		t.Fatalf("claiming again while pending = %v, want ErrClaimPending", err)
	}
}

func TestConfirmClaimNeedsASecondDistinctEligibleAccount(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	officer := seedUser(t, pool, "officer1@example.com")
	seedCharacter(t, pool, gid, officer, "us/hardcore/officer1", "officer", false)
	if _, err := s.Claim(ctx, gid, officer, true); err != nil {
		t.Fatal(err)
	}

	if _, err := s.ConfirmClaim(ctx, gid, officer); !errors.Is(err, ErrSameAccount) {
		t.Fatalf("confirming your own pending claim = %v, want ErrSameAccount", err)
	}

	nonOfficer := seedUser(t, pool, "plain@example.com")
	seedCharacter(t, pool, gid, nonOfficer, "us/hardcore/plain", "member", false)
	if _, err := s.ConfirmClaim(ctx, gid, nonOfficer); !errors.Is(err, ErrNotEligible) {
		t.Fatalf("confirming with no officer/leader character = %v, want ErrNotEligible", err)
	}

	officer2 := seedUser(t, pool, "officer2@example.com")
	seedCharacter(t, pool, gid, officer2, "us/hardcore/officer2", "officer", false)
	claimant, err := s.ConfirmClaim(ctx, gid, officer2)
	if err != nil || claimant != officer {
		t.Fatalf("ConfirmClaim = %d, %v, want %d, nil", claimant, err, officer)
	}
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy == nil || *claimedBy != officer {
		t.Fatalf("claimed_by = %v, want %d", claimedBy, officer)
	}
}

func TestConfirmClaimRefusesOnceExpired(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	officer := seedUser(t, pool, "expiring@example.com")
	seedCharacter(t, pool, gid, officer, "us/hardcore/expiring", "officer", false)
	if _, err := s.Claim(ctx, gid, officer, true); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`update guilds set claim_requested_at = now() - interval '15 days' where id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	officer2 := seedUser(t, pool, "officer2exp@example.com")
	seedCharacter(t, pool, gid, officer2, "us/hardcore/officer2exp", "officer", false)
	if _, err := s.ConfirmClaim(ctx, gid, officer2); !errors.Is(err, ErrNoPendingClaim) {
		t.Fatalf("confirming an expired claim = %v, want ErrNoPendingClaim", err)
	}
	// A fresh Claim call is what actually clears the stale pending fields.
	if _, err := s.Claim(ctx, gid, officer2, true); err != nil {
		t.Fatalf("a fresh claim after expiry should be allowed: %v", err)
	}
}

func TestAutoConfirmClaimIfPendingFiresOnTheGuildMastersOwnExport(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	officer := seedUser(t, pool, "pending-officer@example.com")
	seedCharacter(t, pool, gid, officer, "us/hardcore/pendingofficer", "officer", false)
	if _, err := s.Claim(ctx, gid, officer, true); err != nil {
		t.Fatal(err)
	}

	gm := seedUser(t, pool, "auto-gm@example.com")
	seedCharacter(t, pool, gid, gm, "us/hardcore/autogm", "leader", false)
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := AutoConfirmClaimIfPending(ctx, tx, gid, gm); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy == nil || *claimedBy != officer {
		t.Fatalf("claimed_by = %v, want the originally pending officer %d, confirmed by the GM's own export", claimedBy, officer)
	}
}

func TestReleaseClaimIsTheClaimantOrAModerator(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	uid := seedUser(t, pool, "claimant@example.com")
	seedCharacter(t, pool, gid, uid, "us/hardcore/claimant", "leader", false)
	if _, err := s.Claim(ctx, gid, uid, true); err != nil {
		t.Fatal(err)
	}

	stranger := seedUser(t, pool, "not-claimant@example.com")
	if err := s.ReleaseClaim(ctx, gid, stranger, false); !errors.Is(err, ErrNotClaimant) {
		t.Fatalf("release by a stranger = %v, want ErrNotClaimant", err)
	}
	if err := s.ReleaseClaim(ctx, gid, stranger, true); err != nil {
		t.Fatalf("a moderator should be able to release: %v", err)
	}
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy != nil {
		t.Fatal("claimed_by should be nil after release")
	}
}

func TestClaimRefusesAnAccountWithNoBattleNetIdentity(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "email-only-gm@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/emailgm", "leader", false)

	if _, err := s.Claim(ctx, gid, uid, false); !errors.Is(err, ErrNoBattleNetIdentity) {
		t.Fatalf("Claim with no Battle.net identity = %v, want ErrNoBattleNetIdentity", err)
	}
	result, err := s.Claim(ctx, gid, uid, true)
	if err != nil || result.Status != "confirmed" {
		t.Fatalf("Claim with a Battle.net identity = %+v, %v, want confirmed", result, err)
	}
}

func TestClaimRateLimitsToOnePerAccountPerThirtyDays(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "repeat-claimant@example.com")
	g1 := seedGuild(t, pool, "First")
	g2 := seedGuild(t, pool, "Second")
	seedCharacter(t, pool, g1, uid, "us/hardcore/first", "leader", false)
	if _, err := s.Claim(ctx, g1, uid, true); err != nil {
		t.Fatal(err)
	}
	if err := s.ReleaseClaim(ctx, g1, uid, false); err != nil {
		t.Fatal(err)
	}
	seedCharacter(t, pool, g2, uid, "us/hardcore/second", "leader", false)
	if _, err := s.Claim(ctx, g2, uid, true); !errors.Is(err, ErrClaimRateLimited) {
		t.Fatalf("a second claim within 30 days = %v, want ErrClaimRateLimited (even after releasing the first)", err)
	}
}

func TestClaimRefusesASecondGuildWhileAnotherIsStillHeld(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "two-guild-claimant@example.com")
	g1 := seedGuild(t, pool, "First")
	g2 := seedGuild(t, pool, "Second")
	seedCharacter(t, pool, g1, uid, "us/hardcore/heldfirst", "leader", false)
	if _, err := s.Claim(ctx, g1, uid, true); err != nil {
		t.Fatal(err)
	}
	seedCharacter(t, pool, g2, uid, "us/hardcore/heldsecond", "leader", false)
	if _, err := s.Claim(ctx, g2, uid, true); !errors.Is(err, ErrAlreadyClaimsAnotherGuild) {
		t.Fatalf("claiming a second guild while still holding the first = %v, want ErrAlreadyClaimsAnotherGuild", err)
	}
}

func TestAutoConfirmClaimIfPendingRefusesTheSameAccountsSecondForgedCharacter(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	attacker := seedUser(t, pool, "self-confirm@example.com")
	seedCharacter(t, pool, gid, attacker, "us/hardcore/alt1", "officer", false)
	if _, err := s.Claim(ctx, gid, attacker, true); err != nil {
		t.Fatal(err)
	}
	var pendingBy *int64
	pool.QueryRow(ctx, `select claim_pending_by from guilds where id = $1`, gid).Scan(&pendingBy)
	if pendingBy == nil || *pendingBy != attacker {
		t.Fatal("the attacker's claim should be pending")
	}

	// Same account, a second forged character at rank 0 - must NOT
	// auto-confirm its own pending claim.
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := AutoConfirmClaimIfPending(ctx, tx, gid, attacker); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy != nil {
		t.Fatal("a same-account rank-0 signal must not auto-confirm the account's own pending claim")
	}

	// A genuinely distinct account's rank-0 export still auto-confirms.
	gm := seedUser(t, pool, "real-gm@example.com")
	seedCharacter(t, pool, gid, gm, "us/hardcore/realgm", "leader", false)
	tx, err = pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := AutoConfirmClaimIfPending(ctx, tx, gid, gm); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy == nil || *claimedBy != attacker {
		t.Fatalf("claimed_by = %v, want the originally pending %d, confirmed by a genuinely distinct account", claimedBy, attacker)
	}
}

func TestClaimEnforcesOneClaimedGuildPerAccountUnderConcurrency(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "concurrent-claimant@example.com")
	g1 := seedGuild(t, pool, "ConcurrentFirst")
	g2 := seedGuild(t, pool, "ConcurrentSecond")
	seedCharacter(t, pool, g1, uid, "us/hardcore/concurrentfirst", "leader", false)
	seedCharacter(t, pool, g2, uid, "us/hardcore/concurrentsecond", "leader", false)

	var wg sync.WaitGroup
	results := make(chan error, 2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, err := s.Claim(ctx, g1, uid, true)
		results <- err
	}()
	go func() {
		defer wg.Done()
		_, err := s.Claim(ctx, g2, uid, true)
		results <- err
	}()
	wg.Wait()
	close(results)

	var succeeded, rejected int
	for err := range results {
		switch {
		case err == nil:
			succeeded++
		case errors.Is(err, ErrAlreadyClaimsAnotherGuild), errors.Is(err, ErrClaimRateLimited):
			rejected++
		default:
			t.Fatalf("unexpected error from a concurrent claim: %v", err)
		}
	}
	if succeeded != 1 {
		t.Fatalf("succeeded = %d, want exactly 1 (the unique index must stop a second concurrent claim from the same account)", succeeded)
	}
	if rejected != 1 {
		t.Fatalf("rejected = %d, want exactly 1", rejected)
	}
	var claimedCount int
	if err := pool.QueryRow(ctx, `select count(*) from guilds where claimed_by = $1`, uid).Scan(&claimedCount); err != nil {
		t.Fatal(err)
	}
	if claimedCount != 1 {
		t.Fatalf("guilds claimed by the account = %d, want exactly 1", claimedCount)
	}
	// lockAccountForClaimActivity (HIGH, third security review response)
	// serialises the count-then-insert per account: the losing
	// transaction's own attempt-row insert rolls back with the rest of
	// its transaction, so exactly one 'claim'-kind attempt row should
	// exist, never two (which a lost-update race could otherwise leave
	// behind even where the guilds-table unique index alone would not
	// catch it).
	var attemptCount int
	if err := pool.QueryRow(ctx,
		`select count(*) from guild_claim_attempts where user_id = $1 and kind = 'claim'`, uid).
		Scan(&attemptCount); err != nil {
		t.Fatal(err)
	}
	if attemptCount != 1 {
		t.Fatalf("recorded 'claim' attempts for the account = %d, want exactly 1", attemptCount)
	}
}

// TestContestClaimEnforcesOneOpenContestPerAccountUnderConcurrency is
// item 2 of the third security review response: N concurrent contests
// from one account must not all pass the count-then-act rate-limit
// checks. Mirrors TestClaimEnforcesOneClaimedGuildPerAccountUnderConcurrency.
func TestContestClaimEnforcesOneOpenContestPerAccountUnderConcurrency(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()

	g1 := seedGuild(t, pool, "ConcurrentContestFirst")
	claimant1 := seedUser(t, pool, "concurrent-contest-claimant1@example.com")
	seedCharacter(t, pool, g1, claimant1, "us/hardcore/concurrentcontestclaimant1", "leader", false)
	if _, err := s.Claim(ctx, g1, claimant1, true); err != nil {
		t.Fatal(err)
	}
	g2 := seedGuild(t, pool, "ConcurrentContestSecond")
	claimant2 := seedUser(t, pool, "concurrent-contest-claimant2@example.com")
	seedCharacter(t, pool, g2, claimant2, "us/hardcore/concurrentcontestclaimant2", "leader", false)
	if _, err := s.Claim(ctx, g2, claimant2, true); err != nil {
		t.Fatal(err)
	}

	contester := seedUser(t, pool, "concurrent-contester@example.com")
	seedCharacter(t, pool, g1, contester, "us/hardcore/concurrentcontester1", "officer", false)
	seedCharacter(t, pool, g2, contester, "us/hardcore/concurrentcontester2", "officer", false)

	var wg sync.WaitGroup
	results := make(chan error, 2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		results <- s.ContestClaim(ctx, g1, contester, true)
	}()
	go func() {
		defer wg.Done()
		results <- s.ContestClaim(ctx, g2, contester, true)
	}()
	wg.Wait()
	close(results)

	var succeeded, rejected int
	for err := range results {
		switch {
		case err == nil:
			succeeded++
		case errors.Is(err, ErrAlreadyContestingAnotherGuild), errors.Is(err, ErrContestRateLimited):
			rejected++
		default:
			t.Fatalf("unexpected error from a concurrent contest: %v", err)
		}
	}
	if succeeded != 1 {
		t.Fatalf("succeeded = %d, want exactly 1 (guilds_claim_contested_by_idx must stop a second concurrent contest from the same account)", succeeded)
	}
	if rejected != 1 {
		t.Fatalf("rejected = %d, want exactly 1", rejected)
	}
	var contestedCount int
	if err := pool.QueryRow(ctx, `select count(*) from guilds where claim_contested_by = $1`, contester).Scan(&contestedCount); err != nil {
		t.Fatal(err)
	}
	if contestedCount != 1 {
		t.Fatalf("guilds contested by the account = %d, want exactly 1", contestedCount)
	}
	// Same rationale as the claim-side assertion above: the per-account
	// advisory lock must serialise the attempts count-then-insert, so
	// the losing transaction's attempt row never survives its rollback.
	var attemptCount int
	if err := pool.QueryRow(ctx,
		`select count(*) from guild_claim_attempts where user_id = $1 and kind = 'contest'`, contester).
		Scan(&attemptCount); err != nil {
		t.Fatal(err)
	}
	if attemptCount != 1 {
		t.Fatalf("recorded 'contest' attempts for the account = %d, want exactly 1", attemptCount)
	}
}

// TestContestClaimSameGuildConcurrentContestersRaceCleanly is the
// missing concurrency case the third re-review noted (fourth security
// review response, item 5): two DIFFERENT accounts contesting the SAME
// guild at the same time must not both silently overwrite
// claim_contested_by - the "where claim_contested_at is null" guard on
// the UPDATE, backstopped by RowsAffected, must leave exactly one
// winner and answer the loser ErrAlreadyContested.
func TestContestClaimSameGuildConcurrentContestersRaceCleanly(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()

	gid := seedGuild(t, pool, "SameGuildRace")
	claimant := seedUser(t, pool, "samerace-claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/sameraceclaimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	contesterA := seedUser(t, pool, "samerace-contester-a@example.com")
	seedCharacter(t, pool, gid, contesterA, "us/hardcore/sameracecontestera", "officer", false)
	contesterB := seedUser(t, pool, "samerace-contester-b@example.com")
	seedCharacter(t, pool, gid, contesterB, "us/hardcore/sameracecontesterb", "officer", false)

	var wg sync.WaitGroup
	results := make(chan error, 2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		results <- s.ContestClaim(ctx, gid, contesterA, true)
	}()
	go func() {
		defer wg.Done()
		results <- s.ContestClaim(ctx, gid, contesterB, true)
	}()
	wg.Wait()
	close(results)

	var succeeded, rejected int
	for err := range results {
		switch {
		case err == nil:
			succeeded++
		case errors.Is(err, ErrAlreadyContested):
			rejected++
		default:
			t.Fatalf("unexpected error from a same-guild concurrent contest: %v", err)
		}
	}
	if succeeded != 1 {
		t.Fatalf("succeeded = %d, want exactly 1 (the claim_contested_at is null guard must let only one contester win)", succeeded)
	}
	if rejected != 1 {
		t.Fatalf("rejected = %d, want exactly 1 (ErrAlreadyContested)", rejected)
	}
	var contestedBy *int64
	if err := pool.QueryRow(ctx, `select claim_contested_by from guilds where id = $1`, gid).Scan(&contestedBy); err != nil {
		t.Fatal(err)
	}
	if contestedBy == nil || (*contestedBy != contesterA && *contestedBy != contesterB) {
		t.Fatalf("claim_contested_by = %v, want one of the two contesters, not overwritten or lost", contestedBy)
	}
}

// TestLockAccountForClaimActivityHandlesAUserIDAboveInt32Max is item 4
// (fifth security review response): the reviewer's own minor note - an
// earlier version of lockAccountForClaimActivity cast the user id
// straight to a single int4 (pg_advisory_xact_lock(0, $1::int)), which
// Postgres errors "integer out of range" for above 2^31-1. The fix
// splits the id into its high and low 32-bit halves; this seeds an
// explicit id above that boundary and proves Claim (which takes the
// lock via checkClaimRateLimit) no longer errors on it.
func TestLockAccountForClaimActivityHandlesAUserIDAboveInt32Max(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()

	const bigUserID int64 = 5_000_000_000 // well above 2^31-1 (2147483647)
	if _, err := pool.Exec(ctx,
		`insert into users (id, email) values ($1, 'big-user-id@example.com')`, bigUserID); err != nil {
		t.Fatal(err)
	}
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, bigUserID, "us/hardcore/biguserid", "leader", false)

	result, err := s.Claim(ctx, gid, bigUserID, true)
	if err != nil {
		t.Fatalf("Claim with a user id above 2^31-1 = %v, want no error (the advisory lock must not overflow)", err)
	}
	if result.Status != "confirmed" {
		t.Fatalf("Claim status = %q, want confirmed", result.Status)
	}
}
