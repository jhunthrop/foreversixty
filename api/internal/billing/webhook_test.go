// api/internal/billing/webhook_test.go
package billing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
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
		ID: id, Status: status, Metadata: meta,
		Items: &stripe.SubscriptionItemList{Data: []*stripe.SubscriptionItem{
			{ID: "si_1", CurrentPeriodEnd: periodEnd.Unix()},
		}},
	}
}

// entitlementsStripeUpsertFor turns a fixture *stripe.Subscription into
// the same entitlements.StripeUpsert upsertFromSubscription would build,
// so a test can seed a row identical to what a real webhook delivery
// would have written, without going through the handler.
func entitlementsStripeUpsertFor(sub *stripe.Subscription) entitlements.StripeUpsert {
	uid, _ := metadataUserID(sub)
	end, _ := subscriptionPeriodEnd(sub)
	return entitlements.StripeUpsert{
		UserID: &uid, Plan: sub.Metadata["plan"], Status: string(sub.Status),
		CurrentPeriodEnd: &end, StripeSubscriptionID: sub.ID, BillingUserID: uid,
		Actor: "test_fixture",
	}
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
	if err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlementsStripeUpsertFor(sub)); err != nil {
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
