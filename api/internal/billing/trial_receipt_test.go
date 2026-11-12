// api/internal/billing/trial_receipt_test.go
package billing

import (
	"net/http"
	"strings"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

// paidAfterTrial is the invoice.paid of a trial's first charge.
func paidAfterTrial(subscriptionID, invoiceID string, amount int64, reason string) map[string]any {
	return map[string]any{
		"id": invoiceID, "amount_paid": float64(amount), "billing_reason": reason,
		"hosted_invoice_url": "https://invoice.stripe.com/i/" + invoiceID,
		"parent":             map[string]any{"subscription_details": map[string]any{"subscription": subscriptionID}},
	}
}

// The first charge after a free trial sends one receipt with Stripe's invoice and the next date;
// a redelivery, a $0 invoice and a subscription that never trialed send none.
func TestTrialReceiptAfterTheFirstCharge(t *testing.T) {
	h := newTrialHarness(t)
	uid := h.seedUser(t, "receipt@example.com")
	h.seedCustomer(t, uid, "stripe-receipt@example.com")
	tr := h.seedTrial(t, 0, uid, catalog.LookupKeyPremiumFounderMonthly, trialStart, trialStart.Add(days(7)))
	next := tr.end.AddDate(0, 1, 0)
	sub := fixtureSubscription(tr.subscriptionID, "premium", uid, 0, stripe.SubscriptionStatusActive, next)
	sub.TrialEnd = tr.end.Unix()
	sub.Items.Data[0].Price = &stripe.Price{LookupKey: catalog.LookupKeyPremiumFounderMonthly, UnitAmount: 500}
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}

	if w := h.do(t, signedWebhookRequest(t, "invoice.paid", paidAfterTrial(sub.ID, "in_zero", 0, "subscription_create"))); w.Code != http.StatusOK {
		t.Fatalf("the trial's $0 invoice = %d", w.Code)
	}
	if len(h.mailer.Sent) != 0 {
		t.Fatalf("a $0 invoice sent a receipt")
	}
	for i := 0; i < 2; i++ {
		if w := h.do(t, signedWebhookRequest(t, "invoice.paid", paidAfterTrial(sub.ID, "in_first", 500, billingReasonCycle))); w.Code != http.StatusOK {
			t.Fatalf("the first charge = %d %s", w.Code, w.Body.String())
		}
	}
	if len(h.mailer.Sent) != 1 {
		t.Fatalf("sent %d receipts, want 1", len(h.mailer.Sent))
	}
	m := h.mailer.Sent[0]
	if m.To != "stripe-receipt@example.com" || m.Subject != "Receipt: Premium, $5 a month" ||
		!strings.Contains(m.Text, "https://invoice.stripe.com/i/in_first") ||
		!strings.Contains(m.Text, "The next charge is $5 on "+next.Format("2 Jan 2006")+".") {
		t.Fatalf("%q to %q:\n%s", m.Subject, m.To, m.Text)
	}
	if strings.Contains(h.logs.String(), "@example.com") {
		t.Fatalf("the logs carry an email:\n%s", h.logs.String())
	}

	plain := h.seedUser(t, "never-trialed@example.com")
	h.seedCustomer(t, plain, "stripe-never@example.com")
	paid := fixtureSubscription("sub_never_trialed", "premium", plain, 0, stripe.SubscriptionStatusActive, time.Now().Add(30*24*time.Hour))
	h.gateway.Subscriptions[paid.ID] = paid
	if w := h.do(t, signedWebhookRequest(t, "invoice.paid", paidAfterTrial(paid.ID, "in_plain", 500, billingReasonCycle))); w.Code != http.StatusOK {
		t.Fatalf("a renewal = %d", w.Code)
	}
	if len(h.mailer.Sent) != 1 {
		t.Fatalf("a subscription that never trialed got a receipt")
	}
}
