package entitlements

import (
	"context"
	"errors"
	"testing"
	"time"
)

// premiumRow is the part of a personal premium row a Grant or a Stripe write may change.
type premiumRow struct {
	Source         string
	Status         string
	SubscriptionID *string
	GraceUntil     *time.Time
}

func readPremiumRow(t *testing.T, s *Store, uid int64) premiumRow {
	t.Helper()
	var r premiumRow
	if err := s.Pool.QueryRow(context.Background(),
		`select source, status, stripe_subscription_id, grace_until from entitlements where user_id = $1 and plan = 'premium'`,
		uid).Scan(&r.Source, &r.Status, &r.SubscriptionID, &r.GraceUntil); err != nil {
		t.Fatal(err)
	}
	return r
}

func countAnomaliesOfKind(t *testing.T, s *Store, kind AnomalyKind, subID string) int {
	t.Helper()
	var n int
	if err := s.Pool.QueryRow(context.Background(),
		`select count(*) from entitlement_anomalies where kind = $1 and stripe_subscription_id = $2`,
		string(kind), subID).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func premiumUpsert(uid int64, subID, status string, cancelAtEnd bool) StripeUpsert {
	end := tomorrow().UTC().Truncate(time.Second)
	return StripeUpsert{
		UserID: i64(uid), Plan: PlanPremium, Status: status, CurrentPeriodEnd: &end, CancelAtPeriodEnd: cancelAtEnd,
		StripeSubscriptionID: subID, BillingUserID: uid, Actor: "stripe_webhook:evt_m5",
	}
}

// A live whole-plan grant is never replaced by a Stripe write for a subscription the row does
// not point at (issue #3 M5): a dead one is ignored, a live one is kept off the row and recorded.
func TestUpsertStripeNeverReplacesALiveWholePlanGrant(t *testing.T) {
	for _, tc := range []struct {
		name        string
		status      string
		cancelAtEnd bool
		wantAnomaly int
	}{
		{"canceled subscription", "canceled", false, 0},
		{"live subscription", "active", false, 1},
		{"live subscription ending at period end", "active", true, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := &Store{Pool: testPool(t)}
			uid := seedUser(t, s.Pool)
			seedEntitlementUntil(t, s.Pool, Subject{UserID: i64(uid)}, PlanPremium, "active", nil, nil)
			before := readPremiumRow(t, s, uid)
			result, err := s.UpsertStripe(context.Background(), premiumUpsert(uid, "sub_stale", tc.status, tc.cancelAtEnd))
			if err != nil {
				t.Fatal(err)
			}
			if result.Outcome == OutcomeWritten {
				t.Fatalf("outcome = %q, want the grant kept", result.Outcome)
			}
			if after := readPremiumRow(t, s, uid); after.Source != before.Source || after.Status != before.Status || after.SubscriptionID != nil {
				t.Fatalf("row = %+v, want the grant untouched (%+v)", after, before)
			}
			if n := countAnomaliesOfKind(t, s, AnomalyGrantConflict, "sub_stale"); n != tc.wantAnomaly {
				t.Fatalf("grant_conflict anomalies = %d, want %d", n, tc.wantAnomaly)
			}
		})
	}
}

// A dead subscription never overwrites a live scoped grant either; a live one still upgrades it
// (TestUpsertStripeReplacesScopedGrant).
func TestUpsertStripeDeadSubscriptionLeavesAScopedGrant(t *testing.T) {
	s := &Store{Pool: testPool(t)}
	gid := seedGuild(t, s.Pool)
	uid := seedUser(t, s.Pool)
	seedEntitlementUntil(t, s.Pool, Subject{GuildID: i64(gid)}, PlanGuild, "active",
		[]string{FeatureScopeOfficerViews}, timePtr(tomorrow()))
	end := tomorrow().UTC().Truncate(time.Second)
	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Tier: TierGuild, Status: "canceled", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_dead_scoped", BillingUserID: uid, Actor: "stripe_webhook:evt_d",
	}); err != nil {
		t.Fatal(err)
	}
	if got := readGuildRowState(t, s.Pool, gid); got.Source != "grant" || len(got.Features) != 1 {
		t.Fatalf("row = %+v, want the scoped grant untouched", got)
	}
}

// An expired grant no longer holds: Stripe takes the row back.
func TestUpsertStripeTakesBackAnExpiredGrant(t *testing.T) {
	s := &Store{Pool: testPool(t)}
	uid := seedUser(t, s.Pool)
	seedEntitlementUntil(t, s.Pool, Subject{UserID: i64(uid)}, PlanPremium, "active", nil, timePtr(yesterday()))
	result, err := s.UpsertStripe(context.Background(), premiumUpsert(uid, "sub_after_grant", "active", false))
	if err != nil {
		t.Fatal(err)
	}
	if result.Outcome != OutcomeWritten {
		t.Fatalf("outcome = %q, want written", result.Outcome)
	}
	if got := readPremiumRow(t, s, uid); got.Source != "stripe" || got.SubscriptionID == nil || *got.SubscriptionID != "sub_after_grant" {
		t.Fatalf("row = %+v, want the Stripe subscription", got)
	}
}

// A grant never replaces a live Stripe subscription that keeps billing (issue #3 M5): the
// subscription would bill on with nothing in the database pointing at it.
func TestGrantRefusesALiveStripeSubscription(t *testing.T) {
	s := &Store{Pool: testPool(t)}
	uid := seedUser(t, s.Pool)
	if _, err := s.UpsertStripe(context.Background(), premiumUpsert(uid, "sub_paying", "active", false)); err != nil {
		t.Fatal(err)
	}
	before := readPremiumRow(t, s, uid)
	err := s.Grant(context.Background(), Subject{UserID: i64(uid)}, PlanPremium, nil, nil, "n", "a", nil)
	if !errors.Is(err, ErrLiveSubscription) {
		t.Fatalf("err = %v, want ErrLiveSubscription", err)
	}
	if after := readPremiumRow(t, s, uid); after.Source != before.Source || after.SubscriptionID == nil || *after.SubscriptionID != "sub_paying" {
		t.Fatalf("row = %+v, want the subscription kept (%+v)", after, before)
	}
}

// A subscription already set to end at its period end may be replaced by a grant: nothing
// bills after it, and its later events no longer match the row.
func TestGrantReplacesAStripeSubscriptionEndingAtPeriodEnd(t *testing.T) {
	s := &Store{Pool: testPool(t)}
	uid := seedUser(t, s.Pool)
	if _, err := s.UpsertStripe(context.Background(), premiumUpsert(uid, "sub_ending", "active", true)); err != nil {
		t.Fatal(err)
	}
	if err := s.Grant(context.Background(), Subject{UserID: i64(uid)}, PlanPremium, nil, nil, "n", "a", nil); err != nil {
		t.Fatal(err)
	}
	if got := readPremiumRow(t, s, uid); got.Source != "grant" || got.SubscriptionID != nil {
		t.Fatalf("row = %+v, want the grant", got)
	}
}

// A comeback subscription clears the retention grace the canceled one left (issue #3 L2): a
// paying account is not in a grace.
func TestUpsertStripeLiveWriteClearsTheRetentionGrace(t *testing.T) {
	s := &Store{Pool: testPool(t)}
	uid := seedUser(t, s.Pool)
	if _, err := s.UpsertStripe(context.Background(), premiumUpsert(uid, "sub_old", "active", false)); err != nil {
		t.Fatal(err)
	}
	grace := tomorrow().Add(30 * 24 * time.Hour)
	if err := s.SetCanceled(context.Background(), "sub_old", &grace, "a"); err != nil {
		t.Fatal(err)
	}
	if _, err := s.UpsertStripe(context.Background(), premiumUpsert(uid, "sub_old", "canceled", false)); err != nil {
		t.Fatal(err)
	}
	if got := readPremiumRow(t, s, uid); got.GraceUntil == nil {
		t.Fatalf("row = %+v, want the grace kept while canceled", got)
	}
	if _, err := s.UpsertStripe(context.Background(), premiumUpsert(uid, "sub_new", "active", false)); err != nil {
		t.Fatal(err)
	}
	if got := readPremiumRow(t, s, uid); got.GraceUntil != nil || got.Status != "active" {
		t.Fatalf("row = %+v, want active with no grace", got)
	}
}
