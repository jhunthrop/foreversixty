// api/internal/billing/reconcile_test.go
package billing

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
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
		catalog.ProductIDGuild: {{ID: "sub_orphan_guild", Status: "active"}},
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
		catalog.ProductIDPremium: {{ID: "sub_known", Status: "active"}},
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

// TestReconcileContinuesPastABadRowAndReportsEveryOne is issue #3 M6: one subscription Stripe
// cannot return no longer stops the job; every later row is still healed, the orphan pass and
// the sweeps still run, and the error names every row that failed.
func TestReconcileContinuesPastABadRowAndReportsEveryOne(t *testing.T) {
	h := newHarness(t)
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	subs := map[string]*stripe.Subscription{}
	var users []int64
	for _, id := range []string{"sub_bad_a", "sub_good", "sub_bad_b"} {
		uid := h.seedUser(t, id+"@example.com")
		users = append(users, uid)
		sub := fixtureSubscription(id, "premium", uid, 0, stripe.SubscriptionStatusPastDue, end)
		if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(sub)); err != nil {
			t.Fatal(err)
		}
		healed := *sub
		healed.Status = stripe.SubscriptionStatusActive
		subs[id] = &healed
	}
	delete(subs, "sub_bad_a")
	delete(subs, "sub_bad_b")
	h.gateway.Subscriptions = subs
	h.gateway.ListSubscriptionsResult = map[string][]StripeSubscriptionSummary{
		catalog.ProductIDGuild: {{ID: "sub_orphan_after_bad", Status: "active"}},
	}

	_, err := h.svc.Reconcile(t.Context())
	if err == nil {
		t.Fatal("Reconcile succeeded, want an error naming the bad rows")
	}
	for _, id := range []string{"sub_bad_a", "sub_bad_b"} {
		if !strings.Contains(err.Error(), id) {
			t.Fatalf("err = %v, want it to name %s", err, id)
		}
	}
	var status string
	if err := h.pool.QueryRow(t.Context(),
		`select status from entitlements where user_id = $1`, users[1]).Scan(&status); err != nil {
		t.Fatal(err)
	}
	if status != "active" {
		t.Fatalf("good row status = %q, want active (healed past the bad row)", status)
	}
	var orphans int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from entitlement_anomalies where stripe_subscription_id = 'sub_orphan_after_bad'`).Scan(&orphans); err != nil {
		t.Fatal(err)
	}
	if orphans != 1 {
		t.Fatalf("orphan anomalies = %d, want 1 (the orphan pass still ran)", orphans)
	}
}

// TestReconcileDeletesStripeEventsPastTheRetentionWindow is issue #3 M8: a raw event payload
// carries the buyer's name, email and address, and nothing reads it once Stripe can no longer
// resend the event, so the nightly job deletes rows past the window and keeps the rest.
func TestReconcileDeletesStripeEventsPastTheRetentionWindow(t *testing.T) {
	h := newHarness(t)
	if _, err := h.pool.Exec(t.Context(), `
		insert into stripe_events (id, type, payload, received_at, processed_at) values
		  ('evt_old', 'invoice.paid', '{"customer_email":"old@example.com"}', now() - interval '31 days', now() - interval '31 days'),
		  ('evt_old_unprocessed', 'invoice.paid', '{"customer_email":"lost@example.com"}', now() - interval '31 days', null),
		  ('evt_recent', 'invoice.paid', '{"customer_email":"new@example.com"}', now() - interval '29 days', now() - interval '29 days')`); err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.Reconcile(t.Context()); err != nil {
		t.Fatal(err)
	}
	rows, err := h.pool.Query(t.Context(), `select id from stripe_events order by id`)
	if err != nil {
		t.Fatal(err)
	}
	ids, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) != 1 || ids[0] != "evt_recent" {
		t.Fatalf("stripe_events = %v, want only evt_recent", ids)
	}
}

// TestReconcileHandsTheRowToAConflictingSubscriptionOnceTheGrantLapses: a subscription kept off
// a grant's row (issue #3 M5) is re-read every night, so it takes the row the night the grant
// lapses instead of billing on with no access until its next invoice.
func TestReconcileHandsTheRowToAConflictingSubscriptionOnceTheGrantLapses(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "reconcile-grant-lapse@example.com")
	until := time.Now().Add(time.Hour)
	if err := h.svc.Entitlements.(*entitlements.Store).Grant(t.Context(), entitlements.Subject{UserID: &uid},
		entitlements.PlanPremium, &until, nil, "n", "a", nil); err != nil {
		t.Fatal(err)
	}
	end := time.Now().Add(30 * 24 * time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_beside_grant", "premium", uid, 0, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}
	if w := h.do(t, signedWebhookRequest(t, "customer.subscription.created", map[string]any{"id": sub.ID})); w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	if _, err := h.pool.Exec(t.Context(),
		`update entitlements set current_period_end = now() - interval '1 minute' where user_id = $1`, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.Reconcile(t.Context()); err != nil {
		t.Fatal(err)
	}
	var source, subID string
	if err := h.pool.QueryRow(t.Context(),
		`select source, coalesce(stripe_subscription_id, '') from entitlements where user_id = $1 and plan = 'premium'`,
		uid).Scan(&source, &subID); err != nil {
		t.Fatal(err)
	}
	if source != "stripe" || subID != sub.ID {
		t.Fatalf("row = %s/%q, want the subscription once the grant lapsed", source, subID)
	}
}
