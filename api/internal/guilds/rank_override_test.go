// api/internal/guilds/rank_override_test.go
package guilds

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// memberRank reads guild_members for the account, "" when it has no row.
func memberRank(t *testing.T, pool *pgxpool.Pool, guildID, userID int64) (string, bool) {
	t.Helper()
	var rank string
	var verified bool
	err := pool.QueryRow(context.Background(),
		`select rank, verified_at is not null from guild_members where guild_id = $1 and user_id = $2`,
		guildID, userID).Scan(&rank, &verified)
	if err != nil {
		return "", false
	}
	return rank, verified
}

func grantOverride(t *testing.T, s *Store, guildID, userID int64, expires time.Time) RankOverrideResult {
	t.Helper()
	res, err := s.GrantRankOverride(context.Background(), RankOverride{
		GuildID: guildID, UserID: userID, Rank: RankOfficer,
		ExpiresAt: expires, Note: "testing guild tools", CreatedBy: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func mustExec(t *testing.T, pool *pgxpool.Pool, sql string, args ...any) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), sql, args...); err != nil {
		t.Fatal(err)
	}
}

func TestRankOverrideSurvivesResync(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool, "Censored")
	uid := seedUser(t, pool, "override-resync@example.com")
	seedCharacter(t, pool, gid, uid, "us-hc-alt", "member", false)
	mustExec(t, pool, `update guild_characters set rank_index = 5 where guild_id = $1 and user_id = $2`, gid, uid)

	res := grantOverride(t, s, gid, uid, time.Now().Add(48*time.Hour))
	if res.Rank != RankOfficer || !res.Verified {
		t.Fatalf("after grant = %+v, want officer and verified", res)
	}

	// A sync rewrites the character from the in-game rank and recomputes.
	mustExec(t, pool, `update guild_characters set rank = 'member', refreshed_at = now() where guild_id = $1 and user_id = $2`, gid, uid)
	recomputeMembership(t, pool, gid, uid)
	if rank, verified := memberRank(t, pool, gid, uid); rank != RankOfficer || !verified {
		t.Fatalf("after re-sync = %q verified=%v, want officer verified", rank, verified)
	}
	var rankIndex int
	if err := pool.QueryRow(context.Background(),
		`select rank_index from guild_characters where guild_id = $1 and user_id = $2`, gid, uid).Scan(&rankIndex); err != nil {
		t.Fatal(err)
	}
	if rankIndex != 5 {
		t.Fatalf("rank_index = %d, want 5 untouched", rankIndex)
	}
}

func TestRankOverrideNeverLowersARealRank(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool, "Censored")
	uid := seedUser(t, pool, "override-leader@example.com")
	seedCharacter(t, pool, gid, uid, "us-hc-gm", "leader", true)
	if res := grantOverride(t, s, gid, uid, time.Now().Add(time.Hour)); res.Rank != RankLeader {
		t.Fatalf("an officer override on a leader = %q, want leader", res.Rank)
	}
}

func TestRankOverrideExpiryStopsIt(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool, "Censored")
	uid := seedUser(t, pool, "override-expiry@example.com")
	seedCharacter(t, pool, gid, uid, "us-hc-main", "member", true)
	grantOverride(t, s, gid, uid, time.Now().Add(time.Hour))

	// A recompute ignores an expired row even before the sweep removes it.
	mustExec(t, pool, `update guild_rank_overrides set expires_at = now() - interval '1 minute' where guild_id = $1`, gid)
	recomputeMembership(t, pool, gid, uid)
	if rank, _ := memberRank(t, pool, gid, uid); rank != "member" {
		t.Fatalf("after expiry, recompute = %q, want member", rank)
	}

	// The sweep removes an expired row and recomputes without any sync.
	grantOverride(t, s, gid, uid, time.Now().Add(time.Hour))
	expired, err := s.ExpireRankOverrides(context.Background(), time.Now().Add(2*time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	if len(expired) != 1 || expired[0] != [2]int64{gid, uid} {
		t.Fatalf("expired = %v, want [[%d %d]]", expired, gid, uid)
	}
	if rank, _ := memberRank(t, pool, gid, uid); rank != "member" {
		t.Fatalf("after the sweep = %q, want member", rank)
	}
	var left int
	if err := pool.QueryRow(context.Background(), `select count(*) from guild_rank_overrides`).Scan(&left); err != nil {
		t.Fatal(err)
	}
	if left != 0 {
		t.Fatalf("overrides left after the sweep = %d, want 0", left)
	}
}

func TestRankOverrideRevokeRemovesIt(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool, "Censored")
	uid := seedUser(t, pool, "override-revoke@example.com")
	seedCharacter(t, pool, gid, uid, "us-hc-main", "member", false)
	grantOverride(t, s, gid, uid, time.Now().Add(time.Hour))

	res, removed, err := s.RevokeRankOverride(context.Background(), gid, uid)
	if err != nil {
		t.Fatal(err)
	}
	if !removed || res.Rank != "member" || res.Verified {
		t.Fatalf("after revoke = %+v removed=%v, want an unverified member", res, removed)
	}
	if _, removed, err := s.RevokeRankOverride(context.Background(), gid, uid); err != nil || removed {
		t.Fatalf("a second revoke = removed %v err %v, want false nil", removed, err)
	}
}

func TestRankOverrideNeedsACharacterInTheGuild(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool, "Censored")
	other := seedGuild(t, pool, "Elsewhere")
	uid := seedUser(t, pool, "override-stranger@example.com")
	seedCharacter(t, pool, other, uid, "us-hc-elsewhere", "member", true)

	if res := grantOverride(t, s, gid, uid, time.Now().Add(time.Hour)); res.Rank != "" {
		t.Fatalf("an override with no character in the guild = %+v, want no membership", res)
	}
	recomputeMembership(t, pool, gid, uid)
	if rank, _ := memberRank(t, pool, gid, uid); rank != "" {
		t.Fatalf("after recompute = %q, want no guild_members row", rank)
	}
}
