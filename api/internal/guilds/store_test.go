// api/internal/guilds/store_test.go
package guilds

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestIsMemberReadsGuildCharactersRegardlessOfVerification(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "member@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/baelgrim", "member", false)

	ok, err := s.IsMember(ctx, gid, uid)
	if err != nil || !ok {
		t.Fatalf("IsMember = %v, %v, want true", ok, err)
	}
	other := seedUser(t, pool, "stranger@example.com")
	ok, err = s.IsMember(ctx, gid, other)
	if err != nil || ok {
		t.Fatalf("IsMember for a non-member = %v, %v, want false", ok, err)
	}
}

func TestRecomputeMembershipDerivesTheHighestVerifiedRank(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "alt-haver@example.com")
	gid := seedGuild(t, pool, "Forever")
	// An unverified officer character must not make the account read as an officer.
	seedCharacter(t, pool, gid, uid, "us/hardcore/officeralt", "officer", false)
	seedCharacter(t, pool, gid, uid, "us/hardcore/mainchar", "member", true)

	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := RecomputeMembership(ctx, tx, gid, &uid); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}

	var rank string
	var verified bool
	if err := pool.QueryRow(ctx,
		`select rank, verified_at is not null from guild_members where guild_id = $1 and user_id = $2`,
		gid, uid).Scan(&rank, &verified); err != nil {
		t.Fatal(err)
	}
	if rank != "member" || !verified {
		t.Fatalf("rank = %q verified = %v, want member/true (the unverified officer row must not count)", rank, verified)
	}
	_ = s
}

func TestRecomputeMembershipRemovesTheRowWhenNoCharacterRemains(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "leaving@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/leaver", "member", true)

	tx, _ := pool.Begin(ctx)
	if err := RecomputeMembership(ctx, tx, gid, &uid); err != nil {
		t.Fatal(err)
	}
	_ = tx.Commit(ctx)
	var n int
	pool.QueryRow(ctx, `select count(*) from guild_members where guild_id = $1 and user_id = $2`, gid, uid).Scan(&n)
	if n != 1 {
		t.Fatalf("guild_members rows = %d, want 1 after the character joined", n)
	}

	if _, err := pool.Exec(ctx, `delete from guild_characters where guild_id = $1 and user_id = $2`, gid, uid); err != nil {
		t.Fatal(err)
	}
	tx, _ = pool.Begin(ctx)
	if err := RecomputeMembership(ctx, tx, gid, &uid); err != nil {
		t.Fatal(err)
	}
	_ = tx.Commit(ctx)
	pool.QueryRow(ctx, `select count(*) from guild_members where guild_id = $1 and user_id = $2`, gid, uid).Scan(&n)
	if n != 0 {
		t.Fatalf("guild_members rows = %d, want 0 once no guild_characters row remains", n)
	}
	_ = s
}

func TestReleaseClaimIfLostClearsClaimedByOnlyWhenVerificationIsGone(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	uid := seedUser(t, pool, "leader@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/gm", "leader", true)
	if _, err := pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, uid, gid); err != nil {
		t.Fatal(err)
	}
	tx, _ := pool.Begin(ctx)
	if err := RecomputeMembership(ctx, tx, gid, &uid); err != nil {
		t.Fatal(err)
	}
	if err := ReleaseClaimIfLost(ctx, tx, gid, uid); err != nil {
		t.Fatal(err)
	}
	_ = tx.Commit(ctx)
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy == nil || *claimedBy != uid {
		t.Fatalf("claimed_by = %v, want still %d (the row is still verified)", claimedBy, uid)
	}

	if _, err := pool.Exec(ctx, `delete from guild_characters where guild_id = $1 and user_id = $2`, gid, uid); err != nil {
		t.Fatal(err)
	}
	tx, _ = pool.Begin(ctx)
	if err := RecomputeMembership(ctx, tx, gid, &uid); err != nil {
		t.Fatal(err)
	}
	if err := ReleaseClaimIfLost(ctx, tx, gid, uid); err != nil {
		t.Fatal(err)
	}
	_ = tx.Commit(ctx)
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy != nil {
		t.Fatalf("claimed_by = %v, want nil once the claimant has no character left", claimedBy)
	}
}

func TestGetGuildReadsContestFields(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	contester := seedUser(t, pool, "contester@example.com")
	if _, err := pool.Exec(ctx,
		`update guilds set claim_contested_at = now(), claim_contested_by = $2 where id = $1`, gid, contester); err != nil {
		t.Fatal(err)
	}
	g, err := s.getGuild(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	if g.ClaimContestedAt == nil || g.ClaimContestedBy == nil || *g.ClaimContestedBy != contester {
		t.Fatalf("g.ClaimContestedAt/By = %v, %v, want set and %d", g.ClaimContestedAt, g.ClaimContestedBy, contester)
	}
}

func TestContestedReportsWhetherAGuildsClaimIsDisputed(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")

	yes, err := s.contested(ctx, gid)
	if err != nil || yes {
		t.Fatalf("contested = %v, %v, want false on a fresh guild", yes, err)
	}
	if _, err := pool.Exec(ctx, `update guilds set claim_contested_at = now() where id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	yes, err = s.contested(ctx, gid)
	if err != nil || !yes {
		t.Fatalf("contested = %v, %v, want true once claim_contested_at is set", yes, err)
	}
}

func TestSetVerifiedForAccountVerifiesEveryRowAndRecordsSource(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	uid := seedUser(t, pool, "multichar@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/main", "leader", false)
	seedCharacter(t, pool, gid, uid, "us/hardcore/alt", "member", false)

	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := setVerifiedForAccount(ctx, tx, gid, uid, "claim"); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}

	rows, err := pool.Query(ctx,
		`select verified_at is not null, verified_by from guild_characters where guild_id = $1 and user_id = $2`, gid, uid)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	n := 0
	for rows.Next() {
		var verified bool
		var by *string
		if err := rows.Scan(&verified, &by); err != nil {
			t.Fatal(err)
		}
		if !verified || by == nil || *by != "claim" {
			t.Fatalf("row: verified = %v, verified_by = %v, want true/claim", verified, by)
		}
		n++
	}
	if n != 2 {
		t.Fatalf("verified %d rows, want 2 (both of the account's characters)", n)
	}

	// A second call with a different source must not reclassify an
	// already-verified row - the FIRST path that verified it is what a
	// later release/transfer un-verifies.
	tx, err = pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := setVerifiedForAccount(ctx, tx, gid, uid, "officer"); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	var by string
	if err := pool.QueryRow(ctx,
		`select verified_by from guild_characters where character_key = 'us/hardcore/main'`).Scan(&by); err != nil {
		t.Fatal(err)
	}
	if by != "claim" {
		t.Fatalf("verified_by = %q after a second call, want it to stay claim (first writer wins)", by)
	}
}

func TestRecomputeMembershipsAdvisoryLockSerialisesConcurrentTransactions(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	uid := seedUser(t, pool, "lock-hold@example.com")
	seedCharacter(t, pool, gid, uid, "us/hardcore/lockhold", "member", true)

	txA, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = txA.Rollback(ctx) }()
	if err := RecomputeMembership(ctx, txA, gid, &uid); err != nil {
		t.Fatal(err)
	}
	// txA now holds the advisory lock for this guild, uncommitted.

	started := make(chan struct{})
	done := make(chan error, 1)
	go func() {
		txB, err := pool.Begin(ctx)
		if err != nil {
			done <- err
			return
		}
		defer func() { _ = txB.Rollback(ctx) }()
		close(started)
		done <- RecomputeMembership(ctx, txB, gid, &uid)
	}()
	<-started
	// Give txB every chance to race ahead if the lock did not block it.
	select {
	case err := <-done:
		t.Fatalf("a concurrent RecomputeMembership on the same guild returned (err=%v) before the lock-holding transaction committed - the advisory lock did not serialise it", err)
	case <-time.After(200 * time.Millisecond):
		// Expected: txB is still blocked waiting for the advisory lock.
	}
	if err := txA.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("txB never completed after txA committed and released the lock")
	}
}

func TestContestedReturnsNotFoundForAnUnknownGuild(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	if _, err := s.contested(ctx, 999999999); !errors.Is(err, ErrNotFound) {
		t.Fatalf("contested for an unknown guild id = %v, want ErrNotFound", err)
	}
}
