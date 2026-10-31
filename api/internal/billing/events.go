package billing

import (
	"context"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// paymentFailedCopy is billing.failed's words. It names the failure and where to fix it,
// never a price or an upsell.
var paymentFailedCopy = struct{ Title, Body string }{
	Title: "A payment failed",
	Body:  "Your Forever Sixty payment did not go through. Update your card on the website.",
}

// accountBillingPath is the account page, where the plan and its card are managed.
const accountBillingPath = "/account"

// publishPaymentFailed tells the payer a payment failed. The invoice id, else the event id,
// holds it to once per failed invoice.
func (s *Service) publishPaymentFailed(ctx context.Context, userID int64, invoiceID, eventID string) {
	ref := invoiceID
	if ref == "" {
		ref = eventID
	}
	events.Quietly(ctx, s.Events, s.logger(), events.Event{
		Kind: events.BillingFailed, Audience: events.Users(userID), DedupeKey: "billing.failed:" + ref,
		Title: paymentFailedCopy.Title, Body: paymentFailedCopy.Body,
		Link: strings.TrimRight(s.PublicBaseURL, "/") + accountBillingPath,
		Data: map[string]string{"user_id": strconv.FormatInt(userID, 10)},
	})
}
