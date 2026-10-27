package entitlements

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var standardFeatures = []Feature{
	FeatureServerSims, FeatureRetention, FeatureMultiCompare,
	FeatureHistory, FeatureNotifications, FeatureRosterCheck,
}

// seedTieredGuildPlan inserts an active whole-plan Stripe row of tier.
func seedTieredGuildPlan(t *testing.T, pool *pgxpool.Pool, gid int64, tier, status string) {
	t.Helper()
	if _, err := pool.Exec(context.Background(),
		`insert into entitlements (guild_id, plan, source, status, tier, stripe_subscription_id, current_period_end)
		 values ($1, 'guild', 'stripe', $2, $3, $4, now() + interval '30 days')`,
		gid, status, tier, "sub_t4_"+time.Now().Format("150405.000000000")); err != nil {
		t.Fatal(err)
	}
}

// seedSeat gives holder a seat in state on guild gid; startsAt may be nil
// for a held seat.
func seedSeat(t *testing.T, pool *pgxpool.Pool, gid, holder int64, slot int, state string, startsAt *time.Time) {
	t.Helper()
	if _, err := pool.Exec(context.Background(),
		`insert into premium_seats (guild_id, slot, state, holder_user_id, given_at, movable_at, starts_at)
		 values ($1, $2, $3, $4, now(), now() + interval '30 days', $5)`,
		gid, slot, state, holder, startsAt); err != nil {
		t.Fatal(err)
	}
}

func assertStandard(t *testing.T, s *Store, uid int64, want bool) {
	t.Helper()
	for _, f := range standardFeatures {
		ok, _, err := s.Can(context.Background(), uid, f)
		if err != nil {
			t.Fatal(err)
		}
		if ok != want {
			t.Fatalf("Can(%s) = %v, want %v", f, ok, want)
		}
	}
	sup, err := s.IsSupporter(context.Background(), uid)
	if err != nil {
		t.Fatal(err)
	}
	if sup != want {
		t.Fatalf("IsSupporter = %v, want %v", sup, want)
	}
}

// seatFixture is a guild on tier with one verified member holding slot.
func seatFixture(t *testing.T, tier string, slot int, state string, startsAt *time.Time) (*Store, *pgxpool.Pool, int64, int64) {
	t.Helper()
	pool := testPool(t)
	gid := seedGuild(t, pool)
	seedTieredGuildPlan(t, pool, gid, tier, "active")
	uid := seedUser(t, pool)
	seedGuildMember(t, pool, gid, uid, "member", true)
	seedSeat(t, pool, gid, uid, slot, state, startsAt)
	return &Store{Pool: pool}, pool, gid, uid
}

func TestCan_MemberWithoutSeatHasNoPremium(t *testing.T) {
	pool := testPool(t)
	gid := seedGuild(t, pool)
	seedTieredGuildPlan(t, pool, gid, TierGuild, "active")
	uid := seedUser(t, pool)
	seedGuildMember(t, pool, gid, uid, "leader", true)
	assertStandard(t, &Store{Pool: pool}, uid, false)
}

func TestCan_HeldSeatGrantsPremium(t *testing.T) {
	s, _, _, uid := seatFixture(t, TierRaidGuild, 3, "held", nil)
	assertStandard(t, s, uid, true)
}

func TestCan_AcceptedSeatStartsAtStartsAt(t *testing.T) {
	future := time.Now().Add(time.Hour)
	s, pool, gid, uid := seatFixture(t, TierGuild, 1, "accepted", &future)
	assertStandard(t, s, uid, false)
	if _, err := pool.Exec(context.Background(),
		`update premium_seats set starts_at = now() - interval '1 minute' where guild_id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	assertStandard(t, s, uid, true)
}

func TestCan_SeatBeyondTierCountIgnored(t *testing.T) {
	s, _, _, uid := seatFixture(t, TierGuild, 2, "held", nil)
	assertStandard(t, s, uid, false)
}

func TestCan_OfferedSeatGrantsNothing(t *testing.T) {
	s, pool, gid, uid := seatFixture(t, TierGuild, 1, "held", nil)
	if _, err := pool.Exec(context.Background(),
		`update premium_seats set state = 'offered', offered_at = now() where guild_id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	assertStandard(t, s, uid, false)
}

func TestCan_SeatEndsWithPlan(t *testing.T) {
	s, pool, gid, uid := seatFixture(t, TierRaidGuild, 1, "held", nil)
	assertStandard(t, s, uid, true)
	if _, err := pool.Exec(context.Background(),
		`update entitlements set status = 'canceled' where guild_id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	assertStandard(t, s, uid, false)
}

func TestCan_SeatHolderNoLongerVerified(t *testing.T) {
	s, pool, gid, uid := seatFixture(t, TierRaidGuild, 1, "held", nil)
	if _, err := pool.Exec(context.Background(),
		`update guild_members set verified_at = null where guild_id = $1 and user_id = $2`, gid, uid); err != nil {
		t.Fatal(err)
	}
	assertStandard(t, s, uid, false)
}

func TestCan_ScopedGrantGivesNoSeats(t *testing.T) {
	pool := testPool(t)
	gid := seedGuild(t, pool)
	seedEntitlementUntil(t, pool, Subject{GuildID: i64(gid)}, PlanGuild, "active",
		[]string{FeatureScopeOfficerViews}, timePtr(tomorrow()))
	uid := seedUser(t, pool)
	seedGuildMember(t, pool, gid, uid, "member", true)
	seedSeat(t, pool, gid, uid, 1, "held", nil)
	assertStandard(t, &Store{Pool: pool}, uid, false)
}

func TestCan_OpenEndedPersonalGrantSurvives(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	owner := seedUser(t, pool)
	seedEntitlementUntil(t, pool, Subject{UserID: i64(owner)}, PlanPremium, "active", nil, nil)

	seated := seedGuild(t, pool)
	seedTieredGuildPlan(t, pool, seated, TierGuild, "active")
	seedGuildMember(t, pool, seated, owner, "member", true)
	seedSeat(t, pool, seated, owner, 1, "held", nil)

	other := seedGuild(t, pool)
	seedTieredGuildPlan(t, pool, other, TierGuild, "active")
	seedGuildMember(t, pool, other, owner, "member", true)

	assertStandard(t, s, owner, true)
}

func TestCanForGuild_RaidGuildTools(t *testing.T) {
	for _, tc := range []struct {
		name   string
		setup  func(t *testing.T, pool *pgxpool.Pool, gid int64)
		wantOK bool
		reason Reason
	}{
		{"guild tier", func(t *testing.T, p *pgxpool.Pool, g int64) { seedTieredGuildPlan(t, p, g, TierGuild, "active") }, false, ReasonNoPlan},
		{"raid guild", func(t *testing.T, p *pgxpool.Pool, g int64) { seedTieredGuildPlan(t, p, g, TierRaidGuild, "active") }, true, ReasonEntitled},
		{"scoped grant", func(t *testing.T, p *pgxpool.Pool, g int64) {
			seedEntitlementUntil(t, p, Subject{GuildID: i64(g)}, PlanGuild, "active", []string{FeatureScopeOfficerViews}, timePtr(tomorrow()))
		}, true, ReasonEntitled},
		{"lapsed", func(t *testing.T, p *pgxpool.Pool, g int64) { seedTieredGuildPlan(t, p, g, TierRaidGuild, "canceled") }, false, ReasonNoPlan},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pool := testPool(t)
			gid := seedGuild(t, pool)
			tc.setup(t, pool, gid)
			uid := seedUser(t, pool)
			seedGuildMember(t, pool, gid, uid, "officer", true)
			ok, reason, err := (&Store{Pool: pool}).CanForGuild(context.Background(), gid, uid, FeatureRaidGuildTools)
			if err != nil {
				t.Fatal(err)
			}
			if ok != tc.wantOK || reason != tc.reason {
				t.Fatalf("got %v, %q want %v, %q", ok, reason, tc.wantOK, tc.reason)
			}
		})
	}
}
