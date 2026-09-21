// api/internal/guilds/store_test.go
package guilds

import (
	"context"
	"testing"
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
