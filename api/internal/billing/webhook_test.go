// api/internal/billing/webhook_test.go
package billing

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/webhook"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// signedWebhookRequest builds a raw JSON payload and a valid
// Stripe-Signature header against testWebhookSecret, using stripe-go's
// own signing helper (spec §7: "the Stripe library's own signature
// helper") — no network call, pure local HMAC.
func signedWebhookRequest(t *testing.T, eventType string, object map[string]any) *http.Request {
	t.Helper()
	body := map[string]any{
		"id":          "evt_" + eventType + "_" + time.Now().Format("150405.000000000"),
		"type":        eventType,
		"api_version": stripe.APIVersion, // ConstructEvent refuses a mismatched/absent version
		"data":        map[string]any{"object": object},
	}
	payload, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	signed := webhook.GenerateTestSignedPayload(&webhook.UnsignedPayload{
		Payload: payload, Secret: testWebhookSecret,
	})
	r := httptest.NewRequest(http.MethodPost, "/v1/billing/webhook", bytes.NewReader(signed.Payload))
	r.Header.Set("Stripe-Signature", signed.Header)
	return r
}

func fixtureSubscription(id, plan string, userID, guildID int64, status stripe.SubscriptionStatus, periodEnd time.Time) *stripe.Subscription {
	meta := map[string]string{"plan": plan, "user_id": strconv.FormatInt(userID, 10)}
	if guildID != 0 {
		meta["guild_id"] = strconv.FormatInt(guildID, 10)
	}
	return &stripe.Subscription{
		ID: id, Status: status, Metadata: meta, Created: time.Now().Add(-time.Hour).Unix(),
		Items: &stripe.SubscriptionItemList{Data: []*stripe.SubscriptionItem{
			{ID: "si_1", CurrentPeriodEnd: periodEnd.Unix()},
		}},
	}
}

// entitlementsStripeUpsertFor turns a fixture *stripe.Subscription into
// the same entitlements.StripeUpsert upsertSubscription would build,
// so a test can seed a row identical to what a real webhook delivery
// would have written, without going through the handler. Mirrors
// upsertSubscription's own plan-based subject choice: a premium
// fixture sets UserID, a guild fixture sets GuildID (never both — the
// entitlements_one_subject constraint refuses that combination).
func entitlementsStripeUpsertFor(sub *stripe.Subscription) entitlements.StripeUpsert {
	uid, _ := metadataUserID(sub)
	end, _ := subscriptionPeriodEnd(sub)
	up := entitlements.StripeUpsert{
		Plan: sub.Metadata["plan"], Status: string(sub.Status),
		CurrentPeriodEnd: &end, StripeSubscriptionID: sub.ID, BillingUserID: uid,
		Actor: "test_fixture", Transfer: sub.Metadata["intent"] == "transfer",
	}
	if sub.Created != 0 {
		up.SubscriptionCreated = time.Unix(sub.Created, 0).UTC()
	}
	if up.Plan == entitlements.PlanGuild {
		gid, _ := strconv.ParseInt(sub.Metadata["guild_id"], 10, 64)
		up.GuildID = &gid
	} else {
		up.UserID = &uid
	}
	return up
}

func TestWebhookFirstDeliveryIsProcessedAndAudited(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "webhook-first@example.com")
	end := time.Now().Add(30 * 24 * time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_webhook_1", "premium", uid, 0, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}

	r := signedWebhookRequest(t, "customer.subscription.created", map[string]any{"id": sub.ID})
	w := h.do(t, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}

	var status string
	if err := h.pool.QueryRow(t.Context(),
		`select status from entitlements where user_id = $1 and plan = 'premium'`, uid).Scan(&status); err != nil {
		t.Fatal(err)
	}
	if status != "active" {
		t.Fatalf("status = %q, want active", status)
	}
	var audits int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from entitlement_audit where user_id = $1`, uid).Scan(&audits); err != nil {
		t.Fatal(err)
	}
	if audits != 1 {
		t.Fatalf("audit rows = %d, want 1", audits)
	}
}

func TestWebhookRedeliveryOfTheSameEventIsANoOp(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "webhook-redeliver@example.com")
	end := time.Now().Add(time.Hour).UTC()
	sub := fixtureSubscription("sub_webhook_2", "premium", uid, 0, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}

	body, err := json.Marshal(map[string]any{
		"id": "evt_fixed_1", "type": "customer.subscription.created", "api_version": stripe.APIVersion,
		"data": map[string]any{"object": map[string]any{"id": sub.ID}},
	})
	if err != nil {
		t.Fatal(err)
	}
	signed := webhook.GenerateTestSignedPayload(&webhook.UnsignedPayload{Payload: body, Secret: testWebhookSecret})
	send := func() *httptest.ResponseRecorder {
		r := httptest.NewRequest(http.MethodPost, "/v1/billing/webhook", bytes.NewReader(signed.Payload))
		r.Header.Set("Stripe-Signature", signed.Header)
		return h.do(t, r)
	}
	if w := send(); w.Code != http.StatusOK {
		t.Fatalf("first delivery: %d, %s", w.Code, w.Body.String())
	}
	if w := send(); w.Code != http.StatusOK {
		t.Fatalf("redelivery: %d, %s", w.Code, w.Body.String())
	}
	var audits int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from entitlement_audit where user_id = $1`, uid).Scan(&audits); err != nil {
		t.Fatal(err)
	}
	if audits != 1 {
		t.Fatalf("audit rows = %d, want 1 (redelivery must not reprocess)", audits)
	}
	var events int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from stripe_events where id = 'evt_fixed_1'`).Scan(&events); err != nil {
		t.Fatal(err)
	}
	if events != 1 {
		t.Fatalf("stripe_events rows = %d, want 1", events)
	}
}

func TestWebhookBadSignatureIs400AndWritesNothing(t *testing.T) {
	h := newHarness(t)
	body, err := json.Marshal(map[string]any{"id": "evt_bad_sig", "type": "customer.subscription.created",
		"data": map[string]any{"object": map[string]any{"id": "sub_x"}}})
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRequest(http.MethodPost, "/v1/billing/webhook", bytes.NewReader(body))
	r.Header.Set("Stripe-Signature", "t=1,v1=deadbeef")
	w := h.do(t, r)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var n int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from stripe_events where id = 'evt_bad_sig'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatal("a bad signature must never write a stripe_events row")
	}
}

func TestWebhookBodyOverTheLimitIsRefusedBeforeSignatureVerification(t *testing.T) {
	h := newHarness(t)
	oversized := bytes.Repeat([]byte("a"), maxWebhookBody+1)
	r := httptest.NewRequest(http.MethodPost, "/v1/billing/webhook", bytes.NewReader(oversized))
	r.Header.Set("Stripe-Signature", "t=1,v1=deadbeef")
	w := h.do(t, r)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d", w.Code)
	}
}

func TestWebhookCheckoutSessionCompletedSavesTheCustomerAndUpserts(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "webhook-checkout@example.com")
	end := time.Now().Add(time.Hour).UTC()
	sub := fixtureSubscription("sub_webhook_checkout", "premium", uid, 0, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}

	r := signedWebhookRequest(t, "checkout.session.completed", map[string]any{
		"id": "cs_test_1", "subscription": sub.ID, "customer": "cus_from_webhook",
	})
	if w := h.do(t, r); w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	id, ok, err := h.svc.Store.CustomerID(t.Context(), uid)
	if err != nil || !ok || id != "cus_from_webhook" {
		t.Fatalf("customer id = %q, %v, %v", id, ok, err)
	}
}

func TestWebhookSubscriptionDeletedSetsCanceledAndGrace(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "webhook-deleted@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_webhook_deleted", "premium", uid, 0, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(sub)); err != nil {
		t.Fatal(err)
	}
	r := signedWebhookRequest(t, "customer.subscription.deleted", map[string]any{"id": sub.ID})
	if w := h.do(t, r); w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var status string
	var grace time.Time
	if err := h.pool.QueryRow(t.Context(),
		`select status, grace_until from entitlements where user_id = $1`, uid).Scan(&status, &grace); err != nil {
		t.Fatal(err)
	}
	if status != "canceled" {
		t.Fatalf("status = %q", status)
	}
	if !grace.Equal(end.Add(retentionGraceDays)) {
		t.Fatalf("grace_until = %v, want %v", grace, end.Add(retentionGraceDays))
	}
}

// TestWebhookDuplicateSubscriptionForTheSameGuildIsKeptAndTheNewcomerCanceled
// is the webhook-layer regression test for the security review's
// double-billing finding (2026-09-21): a second, live subscription for a
// guild that already has an active one must never overwrite it — the
// first officer's subscription (sub_a) stays entitled, an anomaly is
// recorded, and the newcomer (sub_b) is canceled at Stripe instead of
// silently taking over.
func TestWebhookDuplicateSubscriptionForTheSameGuildIsKeptAndTheNewcomerCanceled(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "webhook-dupe-guild", true)
	officerA := h.seedUser(t, "webhook-dupe-a@example.com")
	officerB := h.seedUser(t, "webhook-dupe-b@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)

	subA := fixtureSubscription("sub_dupe_a", "guild", officerA, gid, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{subA.ID: subA}
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(subA)); err != nil {
		t.Fatal(err)
	}

	subB := fixtureSubscription("sub_dupe_b", "guild", officerB, gid, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions["sub_dupe_b"] = subB

	r := signedWebhookRequest(t, "customer.subscription.created", map[string]any{"id": subB.ID})
	w := h.do(t, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}

	var subID string
	if err := h.pool.QueryRow(t.Context(),
		`select stripe_subscription_id from entitlements where guild_id = $1 and plan = 'guild'`, gid).Scan(&subID); err != nil {
		t.Fatal(err)
	}
	if subID != "sub_dupe_a" {
		t.Fatalf("kept subscription id = %q, want sub_dupe_a (untouched)", subID)
	}

	var anomalies int
	var kind, gotSubID string
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from entitlement_anomalies where guild_id = $1`, gid).Scan(&anomalies); err != nil {
		t.Fatal(err)
	}
	if anomalies != 1 {
		t.Fatalf("anomaly rows = %d, want 1", anomalies)
	}
	if err := h.pool.QueryRow(t.Context(),
		`select kind, stripe_subscription_id from entitlement_anomalies where guild_id = $1`, gid).
		Scan(&kind, &gotSubID); err != nil {
		t.Fatal(err)
	}
	if kind != "duplicate_subscription" || gotSubID != "sub_dupe_b" {
		t.Fatalf("anomaly = kind=%q sub=%q, want duplicate_subscription/sub_dupe_b", kind, gotSubID)
	}

	if len(h.gateway.CanceledAtPeriodEnd) != 1 || h.gateway.CanceledAtPeriodEnd[0] != "sub_dupe_b" {
		t.Fatalf("CanceledAtPeriodEnd = %v, want [sub_dupe_b]", h.gateway.CanceledAtPeriodEnd)
	}
}

// TestWebhookTransferSubscriptionOverwritesInsteadOfBeingCanceled is the
// webhook-layer half of the transfer-flow regression fix: a subscription
// carrying intent=transfer metadata (spec §2.9 RULING 10's "Take over
// billing" handoff) must overwrite the guild's existing row and must
// never be canceled at Stripe — the opposite of an ordinary duplicate.
func TestWebhookTransferSubscriptionOverwritesInsteadOfBeingCanceled(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "webhook-transfer-guild", true)
	oldOfficer := h.seedUser(t, "webhook-transfer-old@example.com")
	newOfficer := h.seedUser(t, "webhook-transfer-new@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)

	oldSub := fixtureSubscription("sub_transfer_old", "guild", oldOfficer, gid, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{oldSub.ID: oldSub}
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(oldSub)); err != nil {
		t.Fatal(err)
	}

	newSub := fixtureSubscription("sub_transfer_new", "guild", newOfficer, gid, stripe.SubscriptionStatusActive, end)
	newSub.Metadata["intent"] = "transfer"
	newSub.Created = oldSub.Created + 60
	h.gateway.Subscriptions["sub_transfer_new"] = newSub

	r := signedWebhookRequest(t, "customer.subscription.created", map[string]any{"id": newSub.ID})
	if w := h.do(t, r); w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}

	var subID string
	if err := h.pool.QueryRow(t.Context(),
		`select stripe_subscription_id from entitlements where guild_id = $1 and plan = 'guild'`, gid).Scan(&subID); err != nil {
		t.Fatal(err)
	}
	if subID != "sub_transfer_new" {
		t.Fatalf("stripe_subscription_id = %q, want sub_transfer_new (the transfer must overwrite)", subID)
	}
	if len(h.gateway.CanceledAtPeriodEnd) != 0 {
		t.Fatalf("CanceledAtPeriodEnd = %v, want none — a transfer must never be canceled", h.gateway.CanceledAtPeriodEnd)
	}
	var anomalies int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from entitlement_anomalies where guild_id = $1`, gid).Scan(&anomalies); err != nil {
		t.Fatal(err)
	}
	if anomalies != 0 {
		t.Fatalf("anomaly rows = %d, want 0 for a deliberate transfer", anomalies)
	}
}

// TestWebhookConcurrentDeliveriesForDifferentSubscriptionsOnTheSameGuildAreSerialized
// is a regression test for a gap an independent review found in an
// earlier version of this fix (2026-09-21): checkGuildCheckout's
// advisory lock only ever serializes *creating* a Checkout Session, not
// the webhook deliveries that later land for whatever subscriptions got
// created — two deliveries for two different subscriptions on the same
// guild, with no entitlements row yet for either, could previously both
// read "no existing row" before either had written one, letting the
// second's unconditional upsert silently overwrite the first with no
// anomaly recorded. upsertSubscription now holds WithGuildLock
// around the entire write, closing that window: exactly one subscription
// must end up kept, and the other recorded as a duplicate and canceled.
func TestWebhookConcurrentDeliveriesForDifferentSubscriptionsOnTheSameGuildAreSerialized(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "webhook-race-guild", true)
	officer1 := h.seedUser(t, "webhook-race-1@example.com")
	officer2 := h.seedUser(t, "webhook-race-2@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)

	sub1 := fixtureSubscription("sub_race_1", "guild", officer1, gid, stripe.SubscriptionStatusActive, end)
	sub2 := fixtureSubscription("sub_race_2", "guild", officer2, gid, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub1.ID: sub1, sub2.ID: sub2}

	r1 := signedWebhookRequest(t, "customer.subscription.created", map[string]any{"id": sub1.ID})
	r2 := signedWebhookRequest(t, "customer.subscription.created", map[string]any{"id": sub2.ID})

	codes := make([]int, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); codes[0] = h.do(t, r1).Code }()
	go func() { defer wg.Done(); codes[1] = h.do(t, r2).Code }()
	wg.Wait()

	for i, c := range codes {
		if c != http.StatusOK {
			t.Fatalf("delivery %d: status = %d, want 200 (a webhook always answers 200 once processed, duplicate or not)", i, c)
		}
	}

	var rows int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from entitlements where guild_id = $1 and plan = 'guild'`, gid).Scan(&rows); err != nil {
		t.Fatal(err)
	}
	if rows != 1 {
		t.Fatalf("entitlements rows for the guild = %d, want exactly 1", rows)
	}
	var anomalies int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from entitlement_anomalies where guild_id = $1 and kind = 'duplicate_subscription'`, gid).
		Scan(&anomalies); err != nil {
		t.Fatal(err)
	}
	if anomalies != 1 {
		t.Fatalf("duplicate_subscription anomalies = %d, want exactly 1", anomalies)
	}
	if len(h.gateway.CanceledAtPeriodEnd) != 1 {
		t.Fatalf("CanceledAtPeriodEnd = %v, want exactly one cancellation", h.gateway.CanceledAtPeriodEnd)
	}
	if len(h.gateway.GetSubscriptionCalls) != 4 {
		t.Fatalf("GetSubscriptionCalls = %v, want 4 (two per delivery: lock key, then inside the lock)", h.gateway.GetSubscriptionCalls)
	}
}

func TestWebhookRetriesProcessingWhenAPriorAttemptFailedBeforeCompleting(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "webhook-retry@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_webhook_retry", "premium", uid, 0, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}

	// Simulate a prior attempt that recorded the event row but crashed
	// before processing completed.
	if _, err := h.pool.Exec(t.Context(),
		`insert into stripe_events (id, type, payload) values ($1, $2, $3)`,
		"evt_retry_test", "customer.subscription.created", []byte(`{}`)); err != nil {
		t.Fatal(err)
	}

	body, err := json.Marshal(map[string]any{
		"id": "evt_retry_test", "type": "customer.subscription.created", "api_version": stripe.APIVersion,
		"data": map[string]any{"object": map[string]any{"id": sub.ID}},
	})
	if err != nil {
		t.Fatal(err)
	}
	signed := webhook.GenerateTestSignedPayload(&webhook.UnsignedPayload{Payload: body, Secret: testWebhookSecret})
	r := httptest.NewRequest(http.MethodPost, "/v1/billing/webhook", bytes.NewReader(signed.Payload))
	r.Header.Set("Stripe-Signature", signed.Header)
	w := h.do(t, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}

	var status string
	if err := h.pool.QueryRow(t.Context(),
		`select status from entitlements where user_id = $1 and plan = 'premium'`, uid).Scan(&status); err != nil {
		t.Fatal(err)
	}
	if status != "active" {
		t.Fatalf("status = %q, want active — the retry must have actually processed the event", status)
	}
}

// Since API version 2025-03-31.basil an invoice names its subscription at
// parent.subscription_details.subscription; the older top-level field is
// still read so a replayed older event resolves the same way.
func TestInvoiceSubscriptionIDReadsTheBasilShapeAndTheOldOne(t *testing.T) {
	basil := map[string]any{"id": "in_1", "parent": map[string]any{
		"subscription_details": map[string]any{"subscription": "sub_basil"},
	}}
	if got := invoiceSubscriptionID(basil); got != "sub_basil" {
		t.Errorf("basil shape: %q, want sub_basil", got)
	}
	expanded := map[string]any{"id": "in_2", "parent": map[string]any{
		"subscription_details": map[string]any{"subscription": map[string]any{"id": "sub_expanded"}},
	}}
	if got := invoiceSubscriptionID(expanded); got != "sub_expanded" {
		t.Errorf("expanded basil shape: %q, want sub_expanded", got)
	}
	if got := invoiceSubscriptionID(map[string]any{"id": "in_3", "subscription": "sub_old"}); got != "sub_old" {
		t.Errorf("old shape: %q, want sub_old", got)
	}
	if got := invoiceSubscriptionID(map[string]any{"id": "in_4"}); got != "" {
		t.Errorf("one-off invoice: %q, want empty", got)
	}
}

// invoice.paid is the renewal signal (spec §2.6): a basil-shaped invoice
// for a known subscription re-fetches it and writes the entitlement.
func TestWebhookInvoicePaidUpsertsFromTheBasilInvoiceShape(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "webhook-invoice@example.com")
	end := time.Now().Add(30 * 24 * time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_webhook_invoice", "premium", uid, 0, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}

	r := signedWebhookRequest(t, "invoice.paid", map[string]any{
		"id": "in_webhook_1", "parent": map[string]any{
			"subscription_details": map[string]any{"subscription": sub.ID},
		},
	})
	if w := h.do(t, r); w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var status string
	if err := h.pool.QueryRow(t.Context(),
		`select status from entitlements where user_id = $1 and stripe_subscription_id = $2`, uid, sub.ID).Scan(&status); err != nil {
		t.Fatalf("no entitlement row written from invoice.paid: %v", err)
	}
	if status != "active" {
		t.Fatalf("status = %q, want active", status)
	}
}

// A subscription this service did not sell (no plan metadata: one made in
// the Dashboard, or edited by support) is acknowledged and marked
// processed, never failed: a 500 would make Stripe retry for three days
// and no retry could make the metadata appear.
func TestWebhookSubscriptionWithoutOurMetadataIsAcknowledged(t *testing.T) {
	h := newHarness(t)
	sub := &stripe.Subscription{ID: "sub_not_ours", Status: stripe.SubscriptionStatusActive, Metadata: map[string]string{}}
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}

	r := signedWebhookRequest(t, "customer.subscription.updated", map[string]any{"id": sub.ID})
	w := h.do(t, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var processed bool
	if err := h.pool.QueryRow(t.Context(),
		`select processed_at is not null from stripe_events order by received_at desc limit 1`).Scan(&processed); err != nil {
		t.Fatal(err)
	}
	if !processed {
		t.Fatal("the event was answered 200 but not marked processed")
	}
	var rows int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from entitlements where stripe_subscription_id = $1`, sub.ID).Scan(&rows); err != nil {
		t.Fatal(err)
	}
	if rows != 0 {
		t.Fatalf("an entitlement row was written for a subscription that is not ours")
	}
}

func (h *harness) countAnomalies(t *testing.T, guildID int64) int {
	t.Helper()
	var n int
	if err := h.pool.QueryRow(t.Context(),
		`select count(*) from entitlement_anomalies where guild_id = $1`, guildID).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func (h *harness) guildSubscriptionID(t *testing.T, guildID int64) (id string, cancelAtPeriodEnd bool) {
	t.Helper()
	if err := h.pool.QueryRow(t.Context(),
		`select stripe_subscription_id, cancel_at_period_end from entitlements where guild_id = $1 and plan = 'guild'`,
		guildID).Scan(&id, &cancelAtPeriodEnd); err != nil {
		t.Fatal(err)
	}
	return id, cancelAtPeriodEnd
}

func TestWebhookStaleTransferEventDoesNotOverwriteTheNewerTakeover(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "webhook-stale-transfer-guild", true)
	officerB := h.seedUser(t, "webhook-stale-b@example.com")
	officerC := h.seedUser(t, "webhook-stale-c@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)

	subC := fixtureSubscription("sub_stale_c", "guild", officerC, gid, stripe.SubscriptionStatusActive, end)
	subC.Metadata["intent"] = "transfer"
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(subC)); err != nil {
		t.Fatal(err)
	}
	subB := fixtureSubscription("sub_stale_b", "guild", officerB, gid, stripe.SubscriptionStatusActive, end)
	subB.Metadata["intent"] = "transfer"
	subB.Created = subC.Created - 3600
	subB.CancelAtPeriodEnd = true
	h.gateway.Subscriptions = map[string]*stripe.Subscription{subB.ID: subB}

	r := signedWebhookRequest(t, "customer.subscription.updated", map[string]any{"id": subB.ID})
	if w := h.do(t, r); w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	if id, cancel := h.guildSubscriptionID(t, gid); id != "sub_stale_c" || cancel {
		t.Fatalf("row = %q cancel_at_period_end=%v, want sub_stale_c untouched", id, cancel)
	}
	if len(h.gateway.CanceledAtPeriodEnd) != 0 {
		t.Fatalf("CanceledAtPeriodEnd = %v, want none", h.gateway.CanceledAtPeriodEnd)
	}
	if n := h.countAnomalies(t, gid); n != 0 {
		t.Fatalf("anomaly rows = %d, want 0", n)
	}
}

func TestWebhookDeadNewcomerIsAcknowledgedWithoutCancelOrAnomaly(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "webhook-dead-newcomer-guild", true)
	officerA := h.seedUser(t, "webhook-dead-a@example.com")
	officerX := h.seedUser(t, "webhook-dead-x@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)

	subA := fixtureSubscription("sub_dead_a", "guild", officerA, gid, stripe.SubscriptionStatusActive, end)
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(subA)); err != nil {
		t.Fatal(err)
	}
	subX := fixtureSubscription("sub_dead_x", "guild", officerX, gid, stripe.SubscriptionStatusCanceled, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{subX.ID: subX}

	r := signedWebhookRequest(t, "customer.subscription.updated", map[string]any{"id": subX.ID})
	if w := h.do(t, r); w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	if id, _ := h.guildSubscriptionID(t, gid); id != "sub_dead_a" {
		t.Fatalf("row = %q, want sub_dead_a", id)
	}
	if len(h.gateway.CanceledAtPeriodEnd) != 0 {
		t.Fatalf("CanceledAtPeriodEnd = %v, want none", h.gateway.CanceledAtPeriodEnd)
	}
	if n := h.countAnomalies(t, gid); n != 0 {
		t.Fatalf("anomaly rows = %d, want 0", n)
	}
}

func TestWebhookDuplicateCancelFailureRetriesWithoutASecondAnomaly(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "webhook-cancel-retry-guild", true)
	officerA := h.seedUser(t, "webhook-retry-a@example.com")
	officerB := h.seedUser(t, "webhook-retry-b@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)

	subA := fixtureSubscription("sub_cretry_a", "guild", officerA, gid, stripe.SubscriptionStatusActive, end)
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(subA)); err != nil {
		t.Fatal(err)
	}
	subB := fixtureSubscription("sub_cretry_b", "guild", officerB, gid, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{subB.ID: subB}
	h.gateway.CancelErr = errors.New("stripe is down")

	first := signedWebhookRequest(t, "customer.subscription.created", map[string]any{"id": subB.ID})
	body, err := io.ReadAll(first.Body)
	if err != nil {
		t.Fatal(err)
	}
	send := func() *httptest.ResponseRecorder {
		r := httptest.NewRequest(http.MethodPost, "/v1/billing/webhook", bytes.NewReader(body))
		r.Header.Set("Stripe-Signature", first.Header.Get("Stripe-Signature"))
		return h.do(t, r)
	}
	if w := send(); w.Code != http.StatusInternalServerError {
		t.Fatalf("first delivery: status = %d, want 500", w.Code)
	}
	h.gateway.CancelErr = nil
	if w := send(); w.Code != http.StatusOK {
		t.Fatalf("redelivery: status = %d, body = %s", w.Code, w.Body.String())
	}
	if n := h.countAnomalies(t, gid); n != 1 {
		t.Fatalf("anomaly rows = %d, want 1", n)
	}
	if got := h.gateway.CanceledAtPeriodEnd; len(got) != 1 || got[0] != "sub_cretry_b" {
		t.Fatalf("CanceledAtPeriodEnd = %v, want [sub_cretry_b]", got)
	}
}

func TestWebhookWritesFromTheFetchTakenInsideTheLock(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "webhook-inlock@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_p", "premium", uid, 0, stripe.SubscriptionStatusPastDue, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}
	h.gateway.OnGetSubscription = func(_ string, call int) {
		if call == 1 {
			sub.Status = stripe.SubscriptionStatusActive
		}
	}

	r := signedWebhookRequest(t, "invoice.paid", map[string]any{"parent": map[string]any{
		"subscription_details": map[string]any{"subscription": sub.ID},
	}})
	if w := h.do(t, r); w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var status string
	if err := h.pool.QueryRow(t.Context(),
		`select status from entitlements where user_id = $1 and plan = 'premium'`, uid).Scan(&status); err != nil {
		t.Fatal(err)
	}
	if status != "active" {
		t.Fatalf("status = %q, want active (the write must use the second fetch)", status)
	}
	if got := h.gateway.GetSubscriptionCalls; len(got) != 2 || got[0] != "sub_p" || got[1] != "sub_p" {
		t.Fatalf("GetSubscriptionCalls = %v, want [sub_p sub_p]", got)
	}
}

func TestWebhookUpdatedAfterDeletedStaysCanceled(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "webhook-updated-after-deleted@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_d", "premium", uid, 0, stripe.SubscriptionStatusActive, end)
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(sub)); err != nil {
		t.Fatal(err)
	}
	sub.Status = stripe.SubscriptionStatusCanceled
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}

	read := func() (status string, grace *time.Time) {
		if err := h.pool.QueryRow(t.Context(),
			`select status, grace_until from entitlements where user_id = $1 and plan = 'premium'`, uid).Scan(&status, &grace); err != nil {
			t.Fatal(err)
		}
		return status, grace
	}
	if w := h.do(t, signedWebhookRequest(t, "customer.subscription.deleted", map[string]any{"id": sub.ID})); w.Code != http.StatusOK {
		t.Fatalf("deleted: status = %d, body = %s", w.Code, w.Body.String())
	}
	status, grace := read()
	if status != "canceled" || grace == nil {
		t.Fatalf("after deleted: status = %q grace = %v, want canceled with grace", status, grace)
	}
	if w := h.do(t, signedWebhookRequest(t, "customer.subscription.updated", map[string]any{"id": sub.ID})); w.Code != http.StatusOK {
		t.Fatalf("updated: status = %d, body = %s", w.Code, w.Body.String())
	}
	status2, grace2 := read()
	if status2 != "canceled" || grace2 == nil || !grace2.Equal(*grace) {
		t.Fatalf("after updated: status = %q grace = %v, want canceled with grace %v unchanged", status2, grace2, grace)
	}
}

func TestWebhookMetadataChangedBetweenFetchesWritesNothing(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "webhook-meta-changed-guild", true)
	officer := h.seedUser(t, "webhook-meta-changed@example.com")
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	sub := fixtureSubscription("sub_meta", "guild", officer, gid, stripe.SubscriptionStatusActive, end)
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}
	h.gateway.OnGetSubscription = func(_ string, call int) {
		if call == 2 {
			sub.Metadata["guild_id"] = strconv.FormatInt(gid+1, 10)
		}
	}

	r := signedWebhookRequest(t, "customer.subscription.updated", map[string]any{"id": sub.ID})
	if w := h.do(t, r); w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var rows int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from entitlements where plan = 'guild' and stripe_subscription_id = $1`, sub.ID).Scan(&rows); err != nil {
		t.Fatal(err)
	}
	if rows != 0 {
		t.Fatalf("entitlements rows = %d, want 0", rows)
	}
	if len(h.gateway.CanceledAtPeriodEnd) != 0 {
		t.Fatalf("CanceledAtPeriodEnd = %v, want none", h.gateway.CanceledAtPeriodEnd)
	}
	if n := h.countAnomalies(t, gid) + h.countAnomalies(t, gid+1); n != 0 {
		t.Fatalf("anomaly rows = %d, want 0", n)
	}
}
