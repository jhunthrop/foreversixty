// api/internal/billing/webhook_launch_test.go
package billing

import (
	"errors"
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
		catalog.ProductIDGuild: {{ID: "sub_outgoing", Status: "active", CancelAtPeriodEnd: true}},
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

func (h *harness) anomalyCount(t *testing.T, kind, subID string) int {
	t.Helper()
	var n int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from entitlement_anomalies where kind = $1 and stripe_subscription_id = $2`,
		kind, subID).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func (h *harness) guildRow(t *testing.T, gid int64) (subID, tier, key string, cents int64) {
	t.Helper()
	if err := h.pool.QueryRow(t.Context(), `
		select stripe_subscription_id, tier, coalesce(price_lookup_key, ''), coalesce(unit_amount_cents, 0)
		from entitlements where guild_id = $1`, gid).Scan(&subID, &tier, &key, &cents); err != nil {
		t.Fatal(err)
	}
	return subID, tier, key, cents
}

func (h *harness) offerTransfer(t *testing.T, gid, fromUID, toUID int64, from, state string) {
	t.Helper()
	if _, err := h.pool.Exec(t.Context(), `
		insert into plan_transfers (guild_id, from_user_id, to_user_id, from_subscription_id, state)
		values ($1, $2, $3, $4, $5)`, gid, fromUID, toUID, from, state); err != nil {
		t.Fatal(err)
	}
}

func TestWebhook_EveryGuildKeyStoresItsCatalogTier(t *testing.T) {
	for _, price := range catalog.Prices() {
		if price.Offer.Tier() == "" {
			continue
		}
		t.Run(price.LookupKey, func(t *testing.T) {
			h := newHarness(t)
			sub, _, gid := guildFixture(h, t, "sub_tier_"+price.LookupKey, price.LookupKey, price.AmountCents)
			h.deliverSubscription(t, sub)
			_, tier, key, cents := h.guildRow(t, gid)
			if tier != price.Offer.Tier() || key != price.LookupKey || cents != price.AmountCents {
				t.Fatalf("tier=%q key=%q cents=%d", tier, key, cents)
			}
			if n := h.anomalyCount(t, "unknown_price", sub.ID); n != 0 {
				t.Fatalf("unknown_price anomalies = %d", n)
			}
		})
	}
}

func TestWebhook_RaidTierSurvivesALaterEventWithNoPrice(t *testing.T) {
	h := newHarness(t)
	sub, _, gid := guildFixture(h, t, "sub_raid_keep", catalog.LookupKeyRaidGuildYearly, 25000)
	h.deliverSubscription(t, sub)
	sub.Items.Data[0].Price = nil
	h.deliverSubscription(t, sub)
	_, tier, key, cents := h.guildRow(t, gid)
	if tier != catalog.TierRaidGuild || key != catalog.LookupKeyRaidGuildYearly || cents != 25000 {
		t.Fatalf("tier=%q key=%q cents=%d", tier, key, cents)
	}
}

func TestWebhook_CheckoutAndUpdateEventsInAnyOrderAndRepeated(t *testing.T) {
	h := newHarness(t)
	sub, _, gid := guildFixture(h, t, "sub_events", catalog.LookupKeyRaidGuildMonthly, 2500)
	sub.TrialEnd = time.Now().Add(24 * time.Hour).Unix()
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}
	checkout := map[string]any{"id": "cs_events", "subscription": sub.ID}
	update := map[string]any{"id": sub.ID}
	for _, ev := range []struct {
		kind string
		obj  map[string]any
	}{
		{"customer.subscription.updated", update}, {"checkout.session.completed", checkout},
		{"checkout.session.completed", checkout}, {"customer.subscription.updated", update},
	} {
		if w := h.do(t, signedWebhookRequest(t, ev.kind, ev.obj)); w.Code != http.StatusOK {
			t.Fatalf("%s: status = %d, body = %s", ev.kind, w.Code, w.Body.String())
		}
	}
	gotSub, tier, _, _ := h.guildRow(t, gid)
	var trials int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from guild_trials where guild_id = $1`, gid).Scan(&trials); err != nil {
		t.Fatal(err)
	}
	if gotSub != sub.ID || tier != catalog.TierRaidGuild || trials != 1 || len(h.gateway.CanceledAtPeriodEnd) != 0 {
		t.Fatalf("sub=%q tier=%q trials=%d canceled=%v", gotSub, tier, trials, h.gateway.CanceledAtPeriodEnd)
	}
}

func TestWebhook_OutgoingSubscriptionEventsAreIgnoredAfterTransfer(t *testing.T) {
	h := newHarness(t)
	oldUID := h.seedUser(t, "old-event@example.com")
	sub, newUID, gid := guildFixture(h, t, "sub_new_ev", catalog.LookupKeyGuildMonthly, 1500)
	h.offerTransfer(t, gid, oldUID, newUID, "sub_old_ev", "offered")
	sub.Metadata["intent"] = "transfer"
	sub.Metadata["from_subscription_id"] = "sub_old_ev"
	h.deliverSubscription(t, sub)

	old := fixtureSubscription("sub_old_ev", "guild", oldUID, gid, stripe.SubscriptionStatusActive, time.Now().Add(time.Hour))
	old.CancelAtPeriodEnd = true
	old.Items.Data[0].Price = &stripe.Price{ID: "price_old", LookupKey: catalog.LookupKeyGuildMonthly, UnitAmount: 1500}
	h.gateway.Subscriptions["sub_old_ev"] = old
	w := h.do(t, signedWebhookRequest(t, "customer.subscription.updated", map[string]any{"id": "sub_old_ev"}))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if !slices.Equal(h.gateway.CanceledAtPeriodEnd, []string{"sub_old_ev"}) {
		t.Fatalf("canceled = %v", h.gateway.CanceledAtPeriodEnd)
	}
	if n := h.anomalyCount(t, "duplicate_subscription", "sub_old_ev"); n != 0 {
		t.Fatalf("duplicate anomalies = %d", n)
	}
	if gotSub, _, _, _ := h.guildRow(t, gid); gotSub != "sub_new_ev" {
		t.Fatalf("row points at %q", gotSub)
	}
}

func TestWebhook_TransferWithoutAnOfferIsAnAnomaly(t *testing.T) {
	cases := map[string]struct{ state, from string }{
		"withdrawn": {"declined", "sub_gone"},
		"no row":    {"", "sub_gone"},
		"no from":   {"", ""},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			h := newHarness(t)
			oldUID := h.seedUser(t, "x-old@example.com")
			sub, newUID, gid := guildFixture(h, t, "sub_unexpected", catalog.LookupKeyGuildMonthly, 1500)
			if c.state != "" {
				h.offerTransfer(t, gid, oldUID, newUID, c.from, c.state)
			}
			sub.Metadata["intent"] = "transfer"
			if c.from != "" {
				sub.Metadata["from_subscription_id"] = c.from
			}
			h.deliverSubscription(t, sub)
			if n := h.anomalyCount(t, "unexpected_transfer", sub.ID); n != 1 {
				t.Fatalf("unexpected_transfer anomalies = %d", n)
			}
			if len(h.gateway.CanceledAtPeriodEnd) != 0 {
				t.Fatalf("canceled = %v", h.gateway.CanceledAtPeriodEnd)
			}
		})
	}
}

func TestWebhook_FailedCancelThenRetryCompletesTheTransfer(t *testing.T) {
	h := newHarness(t)
	oldUID := h.seedUser(t, "retry-old@example.com")
	sub, newUID, gid := guildFixture(h, t, "sub_retry", catalog.LookupKeyGuildMonthly, 1500)
	h.offerTransfer(t, gid, oldUID, newUID, "sub_retry_old", "offered")
	sub.Metadata["intent"] = "transfer"
	sub.Metadata["from_subscription_id"] = "sub_retry_old"
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}
	h.gateway.CancelErr = errors.New("stripe is down")
	w := h.do(t, signedWebhookRequest(t, "customer.subscription.updated", map[string]any{"id": sub.ID}))
	if w.Code == http.StatusOK {
		t.Fatal("a failed cancel must not acknowledge the event")
	}
	if state, err := h.svc.Store.TransferState(t.Context(), gid, "sub_retry_old"); err != nil || state != "offered" {
		t.Fatalf("state = %q, err = %v", state, err)
	}
	h.gateway.CancelErr = nil
	h.deliverSubscription(t, sub)
	if state, _ := h.svc.Store.TransferState(t.Context(), gid, "sub_retry_old"); state != "accepted" {
		t.Fatalf("state = %q", state)
	}
	if !slices.Equal(h.gateway.CanceledAtPeriodEnd, []string{"sub_retry_old"}) {
		t.Fatalf("canceled = %v", h.gateway.CanceledAtPeriodEnd)
	}
}

func TestWebhook_UnknownPriceIsAnAnomaly(t *testing.T) {
	h := newHarness(t)
	sub, _, _ := guildFixture(h, t, "sub_mystery", "mystery_key", 100)
	h.deliverSubscription(t, sub)
	if n := h.anomalyCount(t, "unknown_price", sub.ID); n != 1 {
		t.Fatalf("guild unknown_price anomalies = %d", n)
	}
	uid := h.seedUser(t, "premium-wrong@example.com")
	prem := fixtureSubscription("sub_prem_wrong", "premium", uid, 0, stripe.SubscriptionStatusActive, time.Now().Add(time.Hour))
	prem.Items.Data[0].Price = &stripe.Price{ID: "price_x", LookupKey: catalog.LookupKeyGuildMonthly, UnitAmount: 1500}
	h.deliverSubscription(t, prem)
	if n := h.anomalyCount(t, "unknown_price", prem.ID); n != 1 {
		t.Fatalf("premium unknown_price anomalies = %d", n)
	}
}

func TestWebhook_ComebackOfAnotherUserIsNotMarkedUsed(t *testing.T) {
	h := newHarness(t)
	owner := h.seedUser(t, "cb-owner@example.com")
	buyer := h.seedUser(t, "cb-buyer@example.com")
	var id int64
	if err := h.pool.QueryRow(t.Context(), `
		insert into seat_comebacks (user_id, guild_name, ended_at, price_lookup_key, expires_at)
		values ($1, 'G', now() - interval '1 day', $2, now() + interval '1 day') returning id`,
		owner, catalog.LookupKeyPremiumMonthly).Scan(&id); err != nil {
		t.Fatal(err)
	}
	sub := fixtureSubscription("sub_cb_other", "premium", buyer, 0, stripe.SubscriptionStatusActive, time.Now().Add(time.Hour))
	sub.Metadata["comeback_id"] = strconv.FormatInt(id, 10)
	h.deliverSubscription(t, sub)
	var used bool
	if err := h.pool.QueryRow(t.Context(), `select used_at is not null from seat_comebacks where id = $1`, id).Scan(&used); err != nil {
		t.Fatal(err)
	}
	if used {
		t.Fatal("another user's comeback was marked used")
	}
}

func TestReconcile_ReenabledOutgoingTransferIsAnOrphan(t *testing.T) {
	h := newHarness(t)
	oldUID := h.seedUser(t, "re-old@example.com")
	newUID := h.seedUser(t, "re-new@example.com")
	gid := h.seedGuild(t, "reenabled", false)
	h.offerTransfer(t, gid, oldUID, newUID, "sub_reenabled", "accepted")
	h.gateway.ListSubscriptionsResult = map[string][]StripeSubscriptionSummary{
		catalog.ProductIDGuild: {{ID: "sub_reenabled", Status: "active"}},
	}
	if _, err := h.svc.Reconcile(t.Context()); err != nil {
		t.Fatal(err)
	}
	if n := h.anomalyCount(t, "orphan_subscription", "sub_reenabled"); n != 1 {
		t.Fatalf("orphan anomalies = %d", n)
	}
}
