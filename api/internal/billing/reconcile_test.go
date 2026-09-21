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
	if err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(
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
