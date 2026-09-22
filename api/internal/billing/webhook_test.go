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
// would have written, without going through the handler. Mirrors
// upsertFromSubscription's own plan-based subject choice: a premium
// fixture sets UserID, a guild fixture sets GuildID (never both — the
// entitlements_one_subject constraint refuses that combination).
func entitlementsStripeUpsertFor(sub *stripe.Subscription) entitlements.StripeUpsert {
	uid, _ := metadataUserID(sub)
	end, _ := subscriptionPeriodEnd(sub)
	up := entitlements.StripeUpsert{
		Plan: sub.Metadata["plan"], Status: string(sub.Status),
		CurrentPeriodEnd: &end, StripeSubscriptionID: sub.ID, BillingUserID: uid,
		Actor: "test_fixture",
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
