// api/internal/guilds/claim_test.go
package guilds

import (
	"context"
	"errors"
	"testing"
)

func TestClaimByTheGuildMasterIsImmediate(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "gm@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/gm", "leader", false)

	result, err := s.Claim(ctx, gid, uid)
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
	if _, err := s.Claim(ctx, gid, stranger); !errors.Is(err, ErrNotEligible) {
		t.Fatalf("Claim by a non-member = %v, want ErrNotEligible", err)
	}

	officer := seedUser(t, pool, "officer@example.com")
	seedCharacter(t, pool, gid, officer, "us/hardcore/officer", "officer", false)
	result, err := s.Claim(ctx, gid, officer)
	if err != nil || result.Status != "pending" || result.ExpiresAt == nil {
		t.Fatalf("Claim by an officer = %+v, %v, want pending with an expiry", result, err)
	}
	var pendingBy *int64
	pool.QueryRow(ctx, `select claim_pending_by from guilds where id = $1`, gid).Scan(&pendingBy)
	if pendingBy == nil || *pendingBy != officer {
		t.Fatalf("claim_pending_by = %v, want %d", pendingBy, officer)
	}

	if _, err := s.Claim(ctx, gid, officer); !errors.Is(err, ErrClaimPending) {
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
	if _, err := s.Claim(ctx, gid, officer); err != nil {
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
	if _, err := s.Claim(ctx, gid, officer); err != nil {
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
	if _, err := s.Claim(ctx, gid, officer2); err != nil {
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
	if _, err := s.Claim(ctx, gid, officer); err != nil {
		t.Fatal(err)
	}

	gm := seedUser(t, pool, "auto-gm@example.com")
	seedCharacter(t, pool, gid, gm, "us/hardcore/autogm", "leader", false)
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := AutoConfirmClaimIfPending(ctx, tx, gid); err != nil {
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
	if _, err := s.Claim(ctx, gid, uid); err != nil {
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
