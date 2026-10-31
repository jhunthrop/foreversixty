package billing

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

type eventRecorder struct{ got []events.Event }

func (r *eventRecorder) Publish(_ context.Context, e events.Event) error {
	if err := e.Validate(); err != nil {
		return err
	}
	r.got = append(r.got, e)
	return nil
}

func TestAFailedPaymentTellsThePayer(t *testing.T) {
	h := newHarness(t)
	rec := &eventRecorder{}
	h.svc.Events = rec
	uid := h.seedUser(t, "webhook-failed@example.com")
	sub := fixtureSubscription("sub_failed", "premium", uid, 0, stripe.SubscriptionStatusPastDue,
		time.Now().Add(30*24*time.Hour).UTC().Truncate(time.Second))
	h.gateway.Subscriptions = map[string]*stripe.Subscription{sub.ID: sub}
	invoice := map[string]any{"id": "in_failed_1", "parent": map[string]any{"subscription_details": map[string]any{"subscription": sub.ID}}}
	if w := h.do(t, signedWebhookRequest(t, "invoice.payment_failed", invoice)); w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	if len(rec.got) != 1 || rec.got[0].Kind != events.BillingFailed || rec.got[0].Audience.UserIDs[0] != uid ||
		rec.got[0].DedupeKey != "billing.failed:in_failed_1" {
		t.Fatalf("events %+v", rec.got)
	}
	// A paid invoice tells nothing.
	paid := map[string]any{"id": "in_paid_1", "parent": invoice["parent"]}
	if w := h.do(t, signedWebhookRequest(t, "invoice.paid", paid)); w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if len(rec.got) != 1 {
		t.Fatalf("a paid invoice told %+v", rec.got)
	}
}
