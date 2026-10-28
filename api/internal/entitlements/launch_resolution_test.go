package entitlements

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

var standardFeatures = []Feature{
	FeatureServerSims, FeatureRetention, FeatureMultiCompare,
	FeatureHistory, FeatureNotifications, FeatureRosterCheck, FeatureLogCoach,
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

func TestCan_SeatFollowsPlanLifecycle(t *testing.T) {
	for _, tc := range []struct {
		name string
		seed func(t *testing.T, pool *pgxpool.Pool, gid int64)
		want bool
	}{
		{"active", func(t *testing.T, p *pgxpool.Pool, g int64) { seedTieredGuildPlan(t, p, g, TierRaidGuild, "active") }, true},
		{"past_due", func(t *testing.T, p *pgxpool.Pool, g int64) { seedTieredGuildPlan(t, p, g, TierRaidGuild, "past_due") }, true},
		{"trialing", func(t *testing.T, p *pgxpool.Pool, g int64) { seedTieredGuildPlan(t, p, g, TierRaidGuild, "trialing") }, true},
		{"canceled", func(t *testing.T, p *pgxpool.Pool, g int64) { seedTieredGuildPlan(t, p, g, TierRaidGuild, "canceled") }, false},
		{"unpaid", func(t *testing.T, p *pgxpool.Pool, g int64) { seedTieredGuildPlan(t, p, g, TierRaidGuild, "unpaid") }, false},
		{"expired whole-plan grant", func(t *testing.T, p *pgxpool.Pool, g int64) {
			seedEntitlementUntil(t, p, Subject{GuildID: i64(g)}, PlanGuild, "active", nil, timePtr(yesterday()))
		}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pool := testPool(t)
			gid := seedGuild(t, pool)
			tc.seed(t, pool, gid)
			uid := seedUser(t, pool)
			seedGuildMember(t, pool, gid, uid, "member", true)
			seedSeat(t, pool, gid, uid, 1, "held", nil)
			assertStandard(t, &Store{Pool: pool}, uid, tc.want)
		})
	}
}

func TestCan_SeatHolderVerifiedElsewhereOnly(t *testing.T) {
	pool := testPool(t)
	seatGuild := seedGuild(t, pool)
	seedTieredGuildPlan(t, pool, seatGuild, TierRaidGuild, "active")
	elsewhere := seedGuild(t, pool)
	uid := seedUser(t, pool)
	seedGuildMember(t, pool, elsewhere, uid, "member", true)
	seedGuildMember(t, pool, seatGuild, uid, "member", false)
	seedSeat(t, pool, seatGuild, uid, 1, "held", nil)
	assertStandard(t, &Store{Pool: pool}, uid, false)
}

func TestHeldSeatClauseMatchesCatalogSeatCounts(t *testing.T) {
	for _, tier := range []string{TierGuild, TierRaidGuild} {
		t.Run(tier, func(t *testing.T) {
			pool := testPool(t)
			s := &Store{Pool: pool}
			gid := seedGuild(t, pool)
			seedTieredGuildPlan(t, pool, gid, tier, "active")
			resolving := 0
			for slot := 1; slot <= 3; slot++ {
				uid := seedUser(t, pool)
				seedGuildMember(t, pool, gid, uid, "member", true)
				seedSeat(t, pool, gid, uid, slot, "held", nil)
				ok, err := s.IsSupporter(context.Background(), uid)
				if err != nil {
					t.Fatal(err)
				}
				if ok {
					resolving++
				}
			}
			if want := catalog.SeatsForTier(tier); resolving != want {
				t.Fatalf("%d seats resolve Premium on %s, catalog says %d", resolving, tier, want)
			}
		})
	}
}

func TestCan_RaidGuildToolsIsPerGuild(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	seedEntitlement(t, pool, Subject{UserID: i64(uid)}, PlanPremium, "active")
	if ok, _, err := (&Store{Pool: pool}).Can(context.Background(), uid, FeatureRaidGuildTools); err == nil || ok {
		t.Fatalf("Can(raid_guild_tools) = %v, %v, want an error and false", ok, err)
	}
}

func TestSupportersAmong(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	seedTieredGuildPlan(t, pool, gid, TierGuild, "active")
	seated, personal, none := seedUser(t, pool), seedUser(t, pool), seedUser(t, pool)
	seedGuildMember(t, pool, gid, seated, "member", true)
	seedSeat(t, pool, gid, seated, 1, "held", nil)
	seedGuildMember(t, pool, gid, none, "member", true)
	seedEntitlement(t, pool, Subject{UserID: i64(personal)}, PlanPremium, "active")

	got, err := s.SupportersAmong(context.Background(), []int64{seated, personal, none})
	if err != nil {
		t.Fatal(err)
	}
	if !got[seated] || !got[personal] || got[none] || len(got) != 3 {
		t.Fatalf("SupportersAmong = %v", got)
	}
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
	// The seat's plan lapses, so only the grant can resolve.
	if _, err := pool.Exec(context.Background(),
		`update entitlements set status = 'canceled' where guild_id = $1`, seated); err != nil {
		t.Fatal(err)
	}

	other := seedGuild(t, pool)
	seedTieredGuildPlan(t, pool, other, TierGuild, "active")
	seedGuildMember(t, pool, other, owner, "member", true)

	assertStandard(t, s, owner, true)
}

func TestCanForGuild_RaidGuildTools(t *testing.T) {
	raid := func(status string) func(*testing.T, *pgxpool.Pool, int64) {
		return func(t *testing.T, p *pgxpool.Pool, g int64) { seedTieredGuildPlan(t, p, g, TierRaidGuild, status) }
	}
	for _, tc := range []struct {
		name     string
		setup    func(t *testing.T, pool *pgxpool.Pool, gid int64)
		rank     string
		verified bool
		wantOK   bool
		reason   Reason
	}{
		{"guild tier", func(t *testing.T, p *pgxpool.Pool, g int64) { seedTieredGuildPlan(t, p, g, TierGuild, "active") }, "officer", true, false, ReasonNoPlan},
		{"raid guild", raid("active"), "officer", true, true, ReasonEntitled},
		{"raid guild past_due", raid("past_due"), "officer", true, true, ReasonEntitled},
		{"raid guild member", raid("active"), "member", true, false, ReasonNotOfficer},
		{"raid guild unverified officer", raid("active"), "officer", false, false, ReasonNoPlan},
		{"scoped grant", func(t *testing.T, p *pgxpool.Pool, g int64) {
			seedEntitlementUntil(t, p, Subject{GuildID: i64(g)}, PlanGuild, "active", []string{FeatureScopeOfficerViews}, timePtr(tomorrow()))
		}, "officer", true, true, ReasonEntitled},
		{"lapsed", raid("canceled"), "officer", true, false, ReasonNoPlan},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pool := testPool(t)
			gid := seedGuild(t, pool)
			tc.setup(t, pool, gid)
			uid := seedUser(t, pool)
			seedGuildMember(t, pool, gid, uid, tc.rank, tc.verified)
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
