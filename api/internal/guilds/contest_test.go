// api/internal/guilds/contest_test.go
package guilds

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestContestClaimRequiresAnActiveClaim(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	contester := seedUser(t, pool, "contester1@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/contester1", "officer", false)

	if err := s.ContestClaim(ctx, gid, contester); !errors.Is(err, ErrNoActiveClaim) {
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

	if err := s.ContestClaim(ctx, gid, claimant); !errors.Is(err, ErrSameAccount) {
		t.Fatalf("the claimant contesting their own claim = %v, want ErrSameAccount", err)
	}

	noStanding := seedUser(t, pool, "nostanding@example.com")
	seedCharacter(t, pool, gid, noStanding, "us/hardcore/nostanding", "member", false)
	if err := s.ContestClaim(ctx, gid, noStanding); !errors.Is(err, ErrNotEligible) {
		t.Fatalf("a plain member contesting = %v, want ErrNotEligible", err)
	}

	contester := seedUser(t, pool, "realcontester@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/realcontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester); err != nil {
		t.Fatal(err)
	}
	var contestedBy *int64
	pool.QueryRow(ctx, `select claim_contested_by from guilds where id = $1`, gid).Scan(&contestedBy)
	if contestedBy == nil || *contestedBy != contester {
		t.Fatalf("claim_contested_by = %v, want %d", contestedBy, contester)
	}

	second := seedUser(t, pool, "secondcontester@example.com")
	seedCharacter(t, pool, gid, second, "us/hardcore/secondcontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, second); !errors.Is(err, ErrAlreadyContested) {
		t.Fatalf("contesting an already-contested claim = %v, want ErrAlreadyContested", err)
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
	if err := s.ContestClaim(ctx, gid, contester); err != nil {
		t.Fatal(err)
	}

	if err := s.ResolveClaim(ctx, gid, "uphold"); err != nil {
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
	if err := s.ContestClaim(ctx, gid, contester); err != nil {
		t.Fatal(err)
	}

	if err := s.ResolveClaim(ctx, gid, "release"); err != nil {
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
	if err := s.ContestClaim(ctx, gid, contester); err != nil {
		t.Fatal(err)
	}

	if err := s.ResolveClaim(ctx, gid, "transfer"); err != nil {
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
