// api/internal/billing/webhook_launch_test.go
package billing

import (
	"net/http"
	"slices"
	"strconv"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

func (h *harness) deliverSubscription(t *testing.T, sub *stripe.Subscription) {
	t.Helper()
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}
	w := h.do(t, signedWebhookRequest(t, "customer.subscription.updated", map[string]any{"id": sub.ID}))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
}

func guildFixture(h *harness, t *testing.T, subID, lookupKey string, amount int64) (*stripe.Subscription, int64, int64) {
	t.Helper()
	uid := h.seedUser(t, subID+"@example.com")
	gid := h.seedGuild(t, subID, false)
	sub := fixtureSubscription(subID, "guild", uid, gid, stripe.SubscriptionStatusActive, time.Now().Add(time.Hour))
	sub.Items.Data[0].Price = &stripe.Price{ID: "price_" + subID, LookupKey: lookupKey, UnitAmount: amount}
	return sub, uid, gid
}

func TestWebhook_RaidGuildTierFromPrice(t *testing.T) {
	h := newHarness(t)
	sub, _, gid := guildFixture(h, t, "sub_raid", catalog.LookupKeyRaidGuildMonthly, 2500)
	h.deliverSubscription(t, sub)
	var tier string
	if err := h.pool.QueryRow(t.Context(), `select tier from entitlements where guild_id = $1`, gid).Scan(&tier); err != nil {
		t.Fatal(err)
	}
	if tier != catalog.TierRaidGuild {
		t.Fatalf("tier = %q", tier)
	}
}

func TestWebhook_UnknownGuildKeyFallsBackToGuildTier(t *testing.T) {
	h := newHarness(t)
	sub, _, gid := guildFixture(h, t, "sub_unknown", "mystery_key", 100)
	h.deliverSubscription(t, sub)
	var tier string
	if err := h.pool.QueryRow(t.Context(), `select tier from entitlements where guild_id = $1`, gid).Scan(&tier); err != nil {
		t.Fatal(err)
	}
	if tier != catalog.TierGuild {
		t.Fatalf("tier = %q", tier)
	}
}

func TestWebhook_FounderFactsStored(t *testing.T) {
	h := newHarness(t)
	sub, _, gid := guildFixture(h, t, "sub_founder", catalog.LookupKeyGuildFounderYearly, 12000)
	h.deliverSubscription(t, sub)
	var key string
	var cents int64
	if err := h.pool.QueryRow(t.Context(),
		`select price_lookup_key, unit_amount_cents from entitlements where guild_id = $1`, gid).Scan(&key, &cents); err != nil {
		t.Fatal(err)
	}
	if key != catalog.LookupKeyGuildFounderYearly || cents != 12000 {
		t.Fatalf("key=%q cents=%d", key, cents)
	}
}

func TestWebhook_LookupKeyKeptWhenLaterUpsertHasNone(t *testing.T) {
	h := newHarness(t)
	sub, _, gid := guildFixture(h, t, "sub_keep", catalog.LookupKeyGuildMonthly, 1500)
	h.deliverSubscription(t, sub)
	sub.Items.Data[0].Price = nil
	h.deliverSubscription(t, sub)
	var key string
	if err := h.pool.QueryRow(t.Context(),
		`select coalesce(price_lookup_key, '') from entitlements where guild_id = $1`, gid).Scan(&key); err != nil {
		t.Fatal(err)
	}
	if key != catalog.LookupKeyGuildMonthly {
		t.Fatalf("key = %q", key)
	}
}

func TestWebhook_TrialRecordedOnce(t *testing.T) {
	h := newHarness(t)
	sub, _, gid := guildFixture(h, t, "sub_trial", catalog.LookupKeyGuildMonthly, 1500)
	sub.TrialEnd = time.Now().Add(7 * 24 * time.Hour).Unix()
	h.deliverSubscription(t, sub)
	h.deliverSubscription(t, sub)
	var n int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from guild_trials where guild_id = $1`, gid).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("guild_trials rows = %d", n)
	}
}

func TestWebhook_TransferCancelsOutgoing(t *testing.T) {
	h := newHarness(t)
	oldUID := h.seedUser(t, "old-payer@example.com")
	sub, newUID, gid := guildFixture(h, t, "sub_new", catalog.LookupKeyGuildMonthly, 1500)
	if _, err := h.pool.Exec(t.Context(), `
		insert into plan_transfers (guild_id, from_user_id, to_user_id, from_subscription_id, state)
		values ($1, $2, $3, 'sub_old', 'offered')`, gid, oldUID, newUID); err != nil {
		t.Fatal(err)
	}
	sub.Metadata["intent"] = "transfer"
	sub.Metadata["from_subscription_id"] = "sub_old"
	h.deliverSubscription(t, sub)
	h.deliverSubscription(t, sub)
	if !slices.Equal(h.gateway.CanceledAtPeriodEnd, []string{"sub_old"}) {
		t.Fatalf("canceled = %v", h.gateway.CanceledAtPeriodEnd)
	}
	var state string
	if err := h.pool.QueryRow(t.Context(), `select state from plan_transfers where guild_id = $1`, gid).Scan(&state); err != nil {
		t.Fatal(err)
	}
	if state != "accepted" {
		t.Fatalf("state = %q", state)
	}
}

func TestWebhook_ComebackMarkedUsed(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "comeback@example.com")
	var id int64
	if err := h.pool.QueryRow(t.Context(), `
		insert into seat_comebacks (user_id, guild_name, ended_at, price_lookup_key, expires_at)
		values ($1, 'G', now() - interval '1 day', $2, now() + interval '1 day') returning id`,
		uid, catalog.LookupKeyPremiumMonthly).Scan(&id); err != nil {
		t.Fatal(err)
	}
	sub := fixtureSubscription("sub_cb", "premium", uid, 0, stripe.SubscriptionStatusActive, time.Now().Add(time.Hour))
	sub.Metadata["comeback_id"] = strconv.FormatInt(id, 10)
	h.deliverSubscription(t, sub)
	var used bool
	if err := h.pool.QueryRow(t.Context(), `select used_at is not null from seat_comebacks where id = $1`, id).Scan(&used); err != nil {
		t.Fatal(err)
	}
	if !used {
		t.Fatal("comeback not marked used")
	}
}

func TestReconcile_ScansRaidGuild(t *testing.T) {
	h := newHarness(t)
	h.gateway.ListSubscriptionsResult = map[string][]StripeSubscriptionSummary{
		catalog.ProductIDRaidGuild: {{ID: "sub_orphan_raid", Status: "active"}},
	}
	if _, err := h.svc.Reconcile(t.Context()); err != nil {
		t.Fatal(err)
	}
	var plan string
	if err := h.pool.QueryRow(t.Context(),
		`select plan from entitlement_anomalies where stripe_subscription_id = 'sub_orphan_raid'`).Scan(&plan); err != nil {
		t.Fatal(err)
	}
	if plan != "guild" {
		t.Fatalf("plan = %q", plan)
	}
}

func TestReconcile_OutgoingTransferIsNotOrphan(t *testing.T) {
	h := newHarness(t)
	oldUID := h.seedUser(t, "out-old@example.com")
	newUID := h.seedUser(t, "out-new@example.com")
	gid := h.seedGuild(t, "outgoing", false)
	if _, err := h.pool.Exec(t.Context(), `
		insert into plan_transfers (guild_id, from_user_id, to_user_id, from_subscription_id, state)
		values ($1, $2, $3, 'sub_outgoing', 'accepted')`, gid, oldUID, newUID); err != nil {
		t.Fatal(err)
	}
	h.gateway.ListSubscriptionsResult = map[string][]StripeSubscriptionSummary{
		catalog.ProductIDGuild: {{ID: "sub_outgoing", Status: "active"}},
	}
	if _, err := h.svc.Reconcile(t.Context()); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from entitlement_anomalies where stripe_subscription_id = 'sub_outgoing'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("anomalies = %d", n)
	}
}
