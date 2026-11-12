// api/internal/billing/trial_cancel_test.go
package billing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

// linkRequest is a cancel link's anonymous POST.
func linkRequest(t *testing.T, path, token string) *http.Request {
	t.Helper()
	body, err := json.Marshal(map[string]string{"token": token})
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRequest(http.MethodPost, path, bytes.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	return r
}

// reminded seeds a 7-day Premium trial, sends its reminder and answers the link's token. The
// gateway knows the subscription, as Stripe would, for the re-read after a cancel.
func reminded(t *testing.T, h *trialHarness, email string) (seededTrial, string) {
	t.Helper()
	uid := h.seedUser(t, email)
	h.seedCustomer(t, uid, "stripe-"+email)
	tr := h.seedTrial(t, 0, uid, catalog.LookupKeyPremiumFounderMonthly, trialStart, trialStart.Add(days(7)))
	sub := fixtureSubscription(tr.subscriptionID, "premium", uid, 0, stripe.SubscriptionStatusTrialing, tr.end)
	sub.TrialEnd = tr.end.Unix()
	if h.gateway.Subscriptions == nil {
		h.gateway.Subscriptions = map[string]*stripe.Subscription{}
	}
	h.gateway.Subscriptions[sub.ID] = sub
	h.clockAt(trialStart.Add(days(4.5)))
	sent, _ := h.sendAt(t, trialStart.Add(days(4.5)))
	if len(sent) != 1 {
		t.Fatalf("no reminder for %s", email)
	}
	return tr, tokenOf(t, sent[0])
}

// Opening the link changes nothing; the POST cancels its own subscription at the trial end, once.
func TestCancelLinkCancelsOnlyAfterThePostAndOnlyOnce(t *testing.T) {
	h := newTrialHarness(t)
	tr, token := reminded(t, h, "link@example.com")
	other, _ := reminded(t, h, "other@example.com")

	w := h.do(t, linkRequest(t, "/v1/billing/cancel-link/preview", token))
	if w.Code != http.StatusOK {
		t.Fatalf("preview = %d %s", w.Code, w.Body.String())
	}
	var preview CancelLinkView
	raw, _ := json.Marshal(decodeEnvelope(t, w).Data)
	if err := json.Unmarshal(raw, &preview); err != nil {
		t.Fatal(err)
	}
	if preview.Plan != "premium" || preview.Guild != nil || preview.EndsAt != tr.end.Format(time.RFC3339) ||
		preview.AmountCents == nil || *preview.AmountCents != 500 || preview.Canceled {
		t.Fatalf("preview %+v", preview)
	}
	if len(h.gateway.CanceledAtPeriodEnd) != 0 {
		t.Fatalf("the preview cancelled %v", h.gateway.CanceledAtPeriodEnd)
	}
	if w := h.do(t, linkRequest(t, "/v1/billing/cancel-link", token)); w.Code != http.StatusOK {
		t.Fatalf("cancel = %d %s", w.Code, w.Body.String())
	}
	if !slices.Equal(h.gateway.CanceledAtPeriodEnd, []string{tr.subscriptionID}) || slices.Contains(h.gateway.CanceledAtPeriodEnd, other.subscriptionID) {
		t.Fatalf("cancelled %v, want only %s", h.gateway.CanceledAtPeriodEnd, tr.subscriptionID)
	}
	if w := h.do(t, linkRequest(t, "/v1/billing/cancel-link", token)); w.Code != http.StatusGone {
		t.Fatalf("a second use = %d, want 410", w.Code)
	}
	if len(h.gateway.CanceledAtPeriodEnd) != 1 {
		t.Fatalf("a second use cancelled again")
	}
}

// A link stops working at the trial end, and a token of any other shape is just an expired link.
func TestCancelLinkExpiresAtTheTrialEnd(t *testing.T) {
	h := newTrialHarness(t)
	tr, token := reminded(t, h, "late@example.com")
	h.clockAt(tr.end.Add(time.Minute))
	for _, path := range []string{"/v1/billing/cancel-link/preview", "/v1/billing/cancel-link"} {
		if w := h.do(t, linkRequest(t, path, token)); w.Code != http.StatusGone {
			t.Fatalf("%s after the end = %d, want 410", path, w.Code)
		}
		if w := h.do(t, linkRequest(t, path, "not-a-token")); w.Code != http.StatusGone {
			t.Fatalf("%s with a bad token = %d, want 410", path, w.Code)
		}
	}
	if len(h.gateway.CanceledAtPeriodEnd) != 0 {
		t.Fatalf("an expired link cancelled %v", h.gateway.CanceledAtPeriodEnd)
	}
}

// Stripe refusing the cancel leaves the link working.
func TestCancelLinkSurvivesAStripeFailure(t *testing.T) {
	h := newTrialHarness(t)
	_, token := reminded(t, h, "flaky@example.com")
	h.gateway.CancelErr = errMailDown
	if w := h.do(t, linkRequest(t, "/v1/billing/cancel-link", token)); w.Code != http.StatusInternalServerError {
		t.Fatalf("a refused cancel = %d", w.Code)
	}
	h.gateway.CancelErr = nil
	if w := h.do(t, linkRequest(t, "/v1/billing/cancel-link", token)); w.Code != http.StatusOK {
		t.Fatalf("the retry = %d", w.Code)
	}
}

// The account page cancels and keeps its own trial; a plan that is not trialing is 409, and only
// the guild's payer can cancel a guild's trial. The entitlement stays live to the trial end.
func TestAccountCancelsAndKeepsItsTrial(t *testing.T) {
	h := newTrialHarness(t)
	tr, _ := reminded(t, h, "account-cancel@example.com")
	var uid int64
	if err := h.pool.QueryRow(t.Context(), `select billing_user_id from entitlements where id = $1`, tr.entitlementID).Scan(&uid); err != nil {
		t.Fatal(err)
	}
	h.gateway.Subscriptions[tr.subscriptionID].CancelAtPeriodEnd = true
	if w := h.do(t, h.sessionRequest(t, http.MethodPost, "/v1/billing/trial/cancel", uid, map[string]any{})); w.Code != http.StatusOK {
		t.Fatalf("cancel = %d %s", w.Code, w.Body.String())
	}
	var canceling bool
	var status string
	if err := h.pool.QueryRow(t.Context(), `select cancel_at_period_end, status from entitlements where id = $1`, tr.entitlementID).
		Scan(&canceling, &status); err != nil || !canceling || status != "trialing" {
		t.Fatalf("after cancel: canceling %v, status %s (%v)", canceling, status, err)
	}
	h.gateway.Subscriptions[tr.subscriptionID].CancelAtPeriodEnd = false
	if w := h.do(t, h.sessionRequest(t, http.MethodPost, "/v1/billing/trial/keep", uid, map[string]any{})); w.Code != http.StatusOK {
		t.Fatalf("keep = %d", w.Code)
	}
	if !slices.Equal(h.gateway.CanceledAtPeriodEnd, []string{tr.subscriptionID}) || !slices.Equal(h.gateway.Resumed, []string{tr.subscriptionID}) {
		t.Fatalf("cancelled %v, resumed %v", h.gateway.CanceledAtPeriodEnd, h.gateway.Resumed)
	}
	none := h.seedUser(t, "no-trial@example.com")
	if w := h.do(t, h.sessionRequest(t, http.MethodPost, "/v1/billing/trial/cancel", none, map[string]any{})); w.Code != http.StatusConflict {
		t.Fatalf("no trial = %d, want 409", w.Code)
	}
	payer := h.seedUser(t, "guild-payer@example.com")
	g := h.seedGuild(t, "Cancel guild", true)
	h.seedTrial(t, g, payer, catalog.LookupKeyGuildFounderYearly, trialStart, trialStart.Add(days(7)))
	officer := h.seedUser(t, "not-payer@example.com")
	if w := h.do(t, h.sessionRequest(t, http.MethodPost, "/v1/billing/trial/cancel", officer, map[string]any{"guild_id": g})); w.Code != http.StatusForbidden {
		t.Fatalf("a non-payer = %d, want 403", w.Code)
	}
}
