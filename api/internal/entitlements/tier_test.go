package entitlements

import (
	"context"
	"testing"
)

func assertPlanTier(t *testing.T, s *Store, uid int64, want string) {
	t.Helper()
	got, err := s.PlanTier(context.Background(), uid)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("PlanTier = %q, want %q", got, want)
	}
}

func TestPlanTier_NoPlanIsEmpty(t *testing.T) {
	pool := testPool(t)
	assertPlanTier(t, &Store{Pool: pool}, seedUser(t, pool), "")
}

func TestPlanTier_PersonalPremium(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	seedEntitlement(t, pool, Subject{UserID: i64(uid)}, PlanPremium, "active")
	assertPlanTier(t, &Store{Pool: pool}, uid, PlanPremium)
}

func TestPlanTier_SeatHolderIsPremium(t *testing.T) {
	s, _, _, uid := seatFixture(t, TierRaidGuild, 1, "held", nil)
	assertPlanTier(t, s, uid, PlanPremium)
}

func TestPlanTier_OfficerUsesTheGuildPlansTier(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	guild, raid := seedGuild(t, pool), seedGuild(t, pool)
	seedTieredGuildPlan(t, pool, guild, TierGuild, "active")
	seedTieredGuildPlan(t, pool, raid, TierRaidGuild, "active")
	officer := seedUser(t, pool)
	seedEntitlement(t, pool, Subject{UserID: i64(officer)}, PlanPremium, "active")
	seedGuildMember(t, pool, guild, officer, "officer", true)
	assertPlanTier(t, s, officer, TierGuild)
	seedGuildMember(t, pool, raid, officer, "leader", true)
	assertPlanTier(t, s, officer, TierRaidGuild)
}

func TestPlanTier_MemberOrUnverifiedOfficerGetsNoGuildTier(t *testing.T) {
	pool := testPool(t)
	gid := seedGuild(t, pool)
	seedTieredGuildPlan(t, pool, gid, TierRaidGuild, "active")
	member, unverified := seedUser(t, pool), seedUser(t, pool)
	seedGuildMember(t, pool, gid, member, "member", true)
	seedGuildMember(t, pool, gid, unverified, "officer", false)
	s := &Store{Pool: pool}
	assertPlanTier(t, s, member, "")
	assertPlanTier(t, s, unverified, "")
}

func TestPlanTier_CanceledGuildPlanFallsBack(t *testing.T) {
	pool := testPool(t)
	gid := seedGuild(t, pool)
	seedTieredGuildPlan(t, pool, gid, TierGuild, "canceled")
	uid := seedUser(t, pool)
	seedGuildMember(t, pool, gid, uid, "officer", true)
	seedEntitlement(t, pool, Subject{UserID: i64(uid)}, PlanPremium, "active")
	assertPlanTier(t, &Store{Pool: pool}, uid, PlanPremium)
}

// grantOfficer is a verified officer of a guild whose only guild row is a live grant with
// features (nil for a whole-plan comp grant), and whether they also hold personal Premium.
func grantOfficer(t *testing.T, features []string, premium bool) (*Store, int64) {
	t.Helper()
	pool := testPool(t)
	gid := seedGuild(t, pool)
	if _, err := pool.Exec(context.Background(),
		`insert into entitlements (guild_id, plan, source, status, tier, features)
		 values ($1, 'guild', 'grant', 'active', 'raid_guild', $2)`, gid, features); err != nil {
		t.Fatal(err)
	}
	uid := seedUser(t, pool)
	seedGuildMember(t, pool, gid, uid, "officer", true)
	if premium {
		seedEntitlement(t, pool, Subject{UserID: i64(uid)}, PlanPremium, "active")
	}
	return &Store{Pool: pool}, uid
}

func TestPlanTier_ScopedOfficerViewsGrantRaisesNoTier(t *testing.T) {
	s, uid := grantOfficer(t, []string{FeatureScopeOfficerViews}, false)
	assertPlanTier(t, s, uid, "") // no plan: the caps use Premium's, the lowest
	s, uid = grantOfficer(t, []string{FeatureScopeOfficerViews}, true)
	assertPlanTier(t, s, uid, PlanPremium)
}

func TestPlanTier_CompGuildGrantRaisesNoTier(t *testing.T) {
	s, uid := grantOfficer(t, nil, false)
	assertPlanTier(t, s, uid, "")
	s, uid = grantOfficer(t, nil, true)
	assertPlanTier(t, s, uid, PlanPremium)
}

func TestPlanTier_PaidRaidGuildSubscription(t *testing.T) {
	pool := testPool(t)
	gid := seedGuild(t, pool)
	seedTieredGuildPlan(t, pool, gid, TierRaidGuild, "active")
	uid := seedUser(t, pool)
	seedGuildMember(t, pool, gid, uid, "leader", true)
	assertPlanTier(t, &Store{Pool: pool}, uid, TierRaidGuild)
}
