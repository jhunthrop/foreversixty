package billing

import (
	"errors"
	"net/http"
	"slices"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v82"
)

func TestADeletedAccountsSubscriptionEventsAreAcknowledgedNotWritten(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "deleted-account@example.com")
	end := time.Now().Add(30 * 24 * time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_deleted", "premium", uid, 0, stripe.SubscriptionStatusCanceled, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}
	if _, err := h.pool.Exec(t.Context(), `delete from users where id = $1`, uid); err != nil {
		t.Fatal(err)
	}
	for _, event := range []string{"customer.subscription.updated", "customer.subscription.deleted"} {
		w := h.do(t, signedWebhookRequest(t, event, map[string]any{"id": sub.ID}))
		if w.Code != http.StatusOK {
			t.Fatalf("%s: status = %d, body = %s; want acknowledged", event, w.Code, w.Body.String())
		}
	}
	var rows int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from entitlements where stripe_subscription_id = $1`, sub.ID).Scan(&rows); err != nil {
		t.Fatal(err)
	}
	if rows != 0 {
		t.Fatalf("entitlement rows = %d, want none written for a deleted account", rows)
	}
}

func TestCancelAccountSubscriptionsCancelsEveryLiveOneAtOnce(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "leaving@example.com")
	guild := h.seedGuild(t, "Night Watch", true)
	end := time.Now().Add(30 * 24 * time.Hour).UTC().Truncate(time.Second)
	premium := fixtureSubscription("sub_own", "premium", uid, 0, stripe.SubscriptionStatusActive, end)
	plan := fixtureSubscription("sub_plan", "guild", uid, guild, stripe.SubscriptionStatusPastDue, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{premium.ID: premium, plan.ID: plan}
	for _, sub := range []*stripe.Subscription{premium, plan} {
		if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(sub)); err != nil {
			t.Fatal(err)
		}
	}
	if err := h.svc.CancelAccountSubscriptions(t.Context(), uid); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(h.gateway.CanceledNow, []string{"sub_own", "sub_plan"}) {
		t.Fatalf("cancelled now = %v, want both, at once", h.gateway.CanceledNow)
	}
	// A retry finds both ended at Stripe and makes no call.
	if err := h.svc.CancelAccountSubscriptions(t.Context(), uid); err != nil || len(h.gateway.CanceledNow) != 2 {
		t.Fatalf("retry: err %v, cancelled now %v", err, h.gateway.CanceledNow)
	}
}

func TestCancelAccountSubscriptionsWithNoGatewayRefusesOnlyALiveSubscription(t *testing.T) {
	h := newHarness(t)
	h.svc.Gateway = nil
	free := h.seedUser(t, "free@example.com")
	if err := h.svc.CancelAccountSubscriptions(t.Context(), free); err != nil {
		t.Fatalf("an account with nothing to cancel: %v", err)
	}
	paying := h.seedUser(t, "paying@example.com")
	end := time.Now().Add(30 * 24 * time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_live", "premium", paying, 0, stripe.SubscriptionStatusActive, end)
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(sub)); err != nil {
		t.Fatal(err)
	}
	if err := h.svc.CancelAccountSubscriptions(t.Context(), paying); !errors.Is(err, ErrBillingUnavailable) {
		t.Fatalf("err = %v, want ErrBillingUnavailable", err)
	}
}
