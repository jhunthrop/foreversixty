package billing

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v82"
)

// The event sequences issue #3 listed as untested (coverage gaps 2, 3, 6 and 12).

// premiumState is a premium row's status and cancel flag.
func (h *harness) premiumState(t *testing.T, uid int64) (status string, cancelAtEnd bool) {
	t.Helper()
	if err := h.pool.QueryRow(t.Context(),
		`select status, cancel_at_period_end from entitlements where user_id = $1 and plan = 'premium'`,
		uid).Scan(&status, &cancelAtEnd); err != nil {
		t.Fatal(err)
	}
	return status, cancelAtEnd
}

// deliver sends one event for sub after Stripe moved it to status and cancelAtEnd.
func (h *harness) deliver(t *testing.T, sub *stripe.Subscription, eventType string, status stripe.SubscriptionStatus, cancelAtEnd bool) {
	t.Helper()
	sub.Status, sub.CancelAtPeriodEnd = status, cancelAtEnd
	object := map[string]any{"id": sub.ID}
	if eventType == "invoice.paid" {
		object = map[string]any{"parent": map[string]any{"subscription_details": map[string]any{"subscription": sub.ID}}}
	}
	if w := h.do(t, signedWebhookRequest(t, eventType, object)); w.Code != http.StatusOK {
		t.Fatalf("%s: status = %d, body = %s", eventType, w.Code, w.Body.String())
	}
}

// Gap 2: subscription.updated follows Stripe through a cancel and its undo, dunning and its
// recovery, and unpaid, which ends access.
func TestWebhookUpdatedFollowsCancelDunningAndUnpaid(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "seq-updated@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_seq_upd", "premium", uid, 0, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}
	for i, step := range []struct {
		status      stripe.SubscriptionStatus
		cancelAtEnd bool
		wantLive    bool
	}{
		{stripe.SubscriptionStatusActive, true, true},
		{stripe.SubscriptionStatusActive, false, true},
		{stripe.SubscriptionStatusPastDue, false, true},
		{stripe.SubscriptionStatusActive, false, true},
		{stripe.SubscriptionStatusUnpaid, false, false},
	} {
		h.deliver(t, sub, "customer.subscription.updated", step.status, step.cancelAtEnd)
		status, cancelAtEnd := h.premiumState(t, uid)
		if status != string(step.status) || cancelAtEnd != step.cancelAtEnd {
			t.Fatalf("step %d: row = %s/%v, want %s/%v", i, status, cancelAtEnd, step.status, step.cancelAtEnd)
		}
		billing, err := h.svc.Entitlements.PersonalBilling(t.Context(), uid)
		if err != nil {
			t.Fatal(err)
		}
		if billing.Active != step.wantLive {
			t.Fatalf("step %d: active = %v, want %v", i, billing.Active, step.wantLive)
		}
	}
}

// Gap 3: an invoice arriving after the subscription was deleted does not bring it back, and a
// created event (incomplete) that loses the race to invoice.paid does not undo the payment.
func TestWebhookLateEventsNeverReviveOrUndo(t *testing.T) {
	h := newHarness(t)
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)

	gone := h.seedUser(t, "seq-invoice-after-deleted@example.com")
	deleted := fixtureSubscription("sub_seq_gone", "premium", gone, 0, stripe.SubscriptionStatusActive, end)
	paid := h.seedUser(t, "seq-created-late@example.com")
	racing := fixtureSubscription("sub_seq_race", "premium", paid, 0, stripe.SubscriptionStatusIncomplete, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{deleted.ID: deleted, racing.ID: racing}

	h.deliver(t, deleted, "customer.subscription.created", stripe.SubscriptionStatusActive, false)
	h.deliver(t, deleted, "customer.subscription.deleted", stripe.SubscriptionStatusCanceled, false)
	h.deliver(t, deleted, "invoice.paid", stripe.SubscriptionStatusCanceled, false)
	if status, _ := h.premiumState(t, gone); status != "canceled" {
		t.Fatalf("invoice after deleted: status = %q, want canceled", status)
	}

	h.deliver(t, racing, "invoice.paid", stripe.SubscriptionStatusActive, false)
	h.deliver(t, racing, "customer.subscription.created", stripe.SubscriptionStatusActive, false)
	if status, _ := h.premiumState(t, paid); status != "active" {
		t.Fatalf("created after invoice.paid: status = %q, want active", status)
	}
}

// Gap 6: an officer who abandons a guild checkout is refused a second one while the first
// Checkout page is still open, and may start again once it expired.
func TestGuildCheckoutAbandonedThenRetriedBySameOfficer(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "seq-abandoned", true)
	officer := h.seedUser(t, "seq-abandoned-officer@example.com")
	h.seedGuildMember(t, gid, officer, "leader")
	checkout := func() int {
		return h.do(t, h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", officer,
			map[string]any{"plan": "guild", "interval": "monthly", "guild_id": gid})).Code
	}
	if code := checkout(); code != http.StatusOK {
		t.Fatalf("first checkout = %d", code)
	}
	if code := checkout(); code != http.StatusConflict {
		t.Fatalf("retry while the first page is open = %d, want 409", code)
	}
	if _, err := h.pool.Exec(t.Context(),
		`update pending_checkouts set expires_at = now() - interval '1 minute' where guild_id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	if code := checkout(); code != http.StatusOK {
		t.Fatalf("retry after the first page expired = %d, want 200", code)
	}
	if len(h.gateway.Customers) != 1 {
		t.Fatalf("customers = %v, want one Customer reused", h.gateway.Customers)
	}
}

// Gap 12: the subject lock is a transaction-scoped lock held for the whole of fn, so it holds
// through a transaction-mode pooler; another session cannot take it until fn returns.
func TestSubjectLockIsHeldByATransactionForTheWholeOfFn(t *testing.T) {
	h := newHarness(t)
	const guildID = 987654
	key := fmt.Sprintf("guild:%d", guildID)
	tryLock := func() bool {
		t.Helper()
		tx, err := h.pool.Begin(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = tx.Rollback(t.Context()) }()
		var got bool
		if err := tx.QueryRow(t.Context(), `select pg_try_advisory_xact_lock(hashtext($1))`, key).Scan(&got); err != nil {
			t.Fatal(err)
		}
		return got
	}
	if err := h.svc.Store.WithGuildLock(t.Context(), guildID, func(context.Context) error {
		var xactHeld int
		if err := h.pool.QueryRow(t.Context(), `
			select count(*) from pg_locks l join pg_stat_activity a on a.pid = l.pid
			where l.locktype = 'advisory' and l.granted and a.xact_start is not null`).Scan(&xactHeld); err != nil {
			t.Fatal(err)
		}
		if xactHeld == 0 {
			t.Fatal("no advisory lock held inside an open transaction")
		}
		if tryLock() {
			t.Fatal("another session took the lock while fn was running")
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if !tryLock() {
		t.Fatal("the lock was still held after fn returned")
	}
}
