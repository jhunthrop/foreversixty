// api/internal/billing/reconcile_test.go
package billing

import (
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v82"
)

func TestReconcileUpsertsEveryLiveStripeSubscriptionFresh(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "reconcile@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(
		fixtureSubscription("sub_reconcile_1", "premium", uid, 0, stripe.SubscriptionStatusPastDue, end))); err != nil {
		t.Fatal(err)
	}
	// The reconcile job re-fetches by id and finds the subscription has
	// actually recovered to active — this is what "heals a missed
	// webhook" means concretely.
	h.gateway.Subscriptions = map[string]*stripe.Subscription{
		"sub_reconcile_1": fixtureSubscription("sub_reconcile_1", "premium", uid, 0, stripe.SubscriptionStatusActive, end),
	}
	n, err := h.svc.Reconcile(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("reconciled = %d, want 1", n)
	}
	var status string
	if err := h.pool.QueryRow(t.Context(),
		`select status from entitlements where user_id = $1`, uid).Scan(&status); err != nil {
		t.Fatal(err)
	}
	if status != "active" {
		t.Fatalf("status = %q, want active", status)
	}
}

// TestReconcileReportsAStripeSubscriptionWithNoLocalRowAsAnOrphan is the
// reconcile-layer regression test for the security review's finding that
// the database→Stripe-only reconcile job "never finds" an orphaned
// subscription (2026-09-21) — a live Stripe subscription for one of our
// products with nothing in entitlements pointing at it must now surface
// as an anomaly for a human to investigate.
func TestReconcileReportsAStripeSubscriptionWithNoLocalRowAsAnOrphan(t *testing.T) {
	h := newHarness(t)
	h.gateway.ListSubscriptionsResult = map[string][]StripeSubscriptionSummary{
		ProductIDGuild: {{ID: "sub_orphan_guild", Status: "active"}},
	}
	if _, err := h.svc.Reconcile(t.Context()); err != nil {
		t.Fatal(err)
	}
	var kind, plan, subID, actor string
	if err := h.pool.QueryRow(t.Context(),
		`select kind, plan, stripe_subscription_id, actor from entitlement_anomalies
		 where stripe_subscription_id = 'sub_orphan_guild'`).Scan(&kind, &plan, &subID, &actor); err != nil {
		t.Fatal(err)
	}
	if kind != "orphan_subscription" || plan != "guild" || subID != "sub_orphan_guild" || actor != "stripe_reconcile" {
		t.Fatalf("anomaly = kind=%q plan=%q sub=%q actor=%q", kind, plan, subID, actor)
	}
}

// TestReconcileDoesNotReportAKnownSubscriptionAsAnOrphan confirms the
// orphan check only reports subscriptions with no matching local row —
// every subscription entitlements already knows about must be silent.
func TestReconcileDoesNotReportAKnownSubscriptionAsAnOrphan(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "reconcile-known@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_known", "premium", uid, 0, stripe.SubscriptionStatusActive, end)
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(sub)); err != nil {
		t.Fatal(err)
	}
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}
	h.gateway.ListSubscriptionsResult = map[string][]StripeSubscriptionSummary{
		ProductIDPremium: {{ID: "sub_known", Status: "active"}},
	}
	if _, err := h.svc.Reconcile(t.Context()); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from entitlement_anomalies`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("anomaly rows = %d, want 0 for a subscription entitlements already knows about", n)
	}
}

// TestReconcileSweepsExpiredPendingCheckouts is the reconcile-layer half
// of the guild-checkout pending-row bookkeeping (security review fix,
// 2026-09-21): an abandoned Checkout Session's pending_checkouts row
// must eventually be swept even with no webhook ever landing for it.
func TestReconcileSweepsExpiredPendingCheckouts(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "reconcile-sweep-guild", true)
	uid := h.seedUser(t, "reconcile-sweep@example.com")
	past := time.Now().Add(-time.Hour).UTC().Truncate(time.Second)
	future := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	if err := h.svc.Store.RecordPendingCheckout(t.Context(), gid, uid, "cs_reconcile_expired", past); err != nil {
		t.Fatal(err)
	}
	if err := h.svc.Store.RecordPendingCheckout(t.Context(), gid, uid, "cs_reconcile_live", future); err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.Reconcile(t.Context()); err != nil {
		t.Fatal(err)
	}
	var remaining int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from pending_checkouts where guild_id = $1`, gid).Scan(&remaining); err != nil {
		t.Fatal(err)
	}
	if remaining != 1 {
		t.Fatalf("remaining pending_checkouts rows = %d, want 1 (the live one)", remaining)
	}
}
