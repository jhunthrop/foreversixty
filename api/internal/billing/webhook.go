// api/internal/billing/webhook.go
package billing

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/webhook"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// maxWebhookBody matches Stripe's own Go example (spec §2.6): a
// subscription event's JSON is a few KB; this is generous headroom.
const maxWebhookBody = 64 << 10

func (s *Service) webhook(w http.ResponseWriter, r *http.Request) {
	if s.Gateway == nil || s.WebhookSecret == "" {
		writeUnavailable(w, r)
		return
	}
	payload, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxWebhookBody))
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "the request body could not be read", nil)
		return
	}
	// Signature verified on the raw body before anything else runs
	// against it (spec §3). The library's default 5-minute replay
	// tolerance is used unmodified.
	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), s.WebhookSecret)
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "the webhook signature could not be verified", nil)
		return
	}
	proceed, err := s.Store.RecordEventOnce(r.Context(), event.ID, string(event.Type), payload)
	if err != nil {
		s.fail(w, r, "webhook", err, "could not record that event just now")
		return
	}
	if !proceed {
		// Already fully processed (a genuine redelivery of a completed
		// event, or Stripe's own retry after a slow-but-eventually-200
		// first attempt): 200, no reprocessing.
		httpx.WriteOK(w, r, http.StatusOK, nil)
		return
	}
	if err := s.handleEvent(r.Context(), event); err != nil {
		s.fail(w, r, "webhook", err, "could not process that event just now")
		return
	}
	if err := s.Store.MarkEventProcessed(r.Context(), event.ID); err != nil {
		s.fail(w, r, "webhook", err, "could not process that event just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, nil)
}

func (s *Service) handleEvent(ctx context.Context, event stripe.Event) error {
	actor := "stripe_webhook:" + event.ID
	switch event.Type {
	case "checkout.session.completed":
		return s.onCheckoutCompleted(ctx, event, actor)
	case "customer.subscription.created", "customer.subscription.updated":
		return s.onSubscriptionEvent(ctx, event, actor)
	case "customer.subscription.deleted":
		return s.onSubscriptionDeleted(ctx, event, actor)
	case "invoice.paid", "invoice.payment_failed":
		return s.onInvoiceEvent(ctx, event, actor)
	default:
		// Not in this endpoint's enabled_events list in practice; kept
		// as a safety net if the Dashboard config ever drifts.
		return nil
	}
}

// stripeRefID reads a Stripe reference field the API may encode as a
// bare object id string or as an expanded object with its own "id" key.
// This webhook never asks for expansion, but a defensive read costs
// nothing.
func stripeRefID(v interface{}) string {
	switch t := v.(type) {
	case string:
		return t
	case map[string]interface{}:
		if id, ok := t["id"].(string); ok {
			return id
		}
	}
	return ""
}

func (s *Service) onCheckoutCompleted(ctx context.Context, event stripe.Event, actor string) error {
	subID := stripeRefID(event.Data.Object["subscription"])
	if subID == "" {
		return nil
	}
	sub, err := s.Gateway.GetSubscription(ctx, subID)
	if err != nil {
		return fmt.Errorf("billing: checkout completed: get subscription %s: %w", subID, err)
	}
	if custID := stripeRefID(event.Data.Object["customer"]); custID != "" {
		if uid, ok := metadataUserID(sub); ok {
			if err := s.Store.SaveCustomerID(ctx, uid, custID); err != nil {
				return err
			}
		}
	}
	return s.upsertFromSubscription(ctx, sub, actor)
}

func (s *Service) onSubscriptionEvent(ctx context.Context, event stripe.Event, actor string) error {
	id := stripeRefID(event.Data.Object["id"])
	sub, err := s.Gateway.GetSubscription(ctx, id)
	if err != nil {
		return fmt.Errorf("billing: subscription event: get %s: %w", id, err)
	}
	return s.upsertFromSubscription(ctx, sub, actor)
}

func (s *Service) onSubscriptionDeleted(ctx context.Context, event stripe.Event, actor string) error {
	id := stripeRefID(event.Data.Object["id"])
	sub, err := s.Gateway.GetSubscription(ctx, id)
	if err != nil {
		return fmt.Errorf("billing: subscription deleted: get %s: %w", id, err)
	}
	var grace *time.Time
	if end, ok := subscriptionPeriodEnd(sub); ok {
		g := end.Add(retentionGraceDays)
		grace = &g
	}
	return s.Entitlements.SetCanceled(ctx, id, grace, actor)
}

func (s *Service) onInvoiceEvent(ctx context.Context, event stripe.Event, actor string) error {
	subID := stripeRefID(event.Data.Object["subscription"])
	if subID == "" {
		return nil
	}
	sub, err := s.Gateway.GetSubscription(ctx, subID)
	if err != nil {
		return fmt.Errorf("billing: invoice event: get subscription %s: %w", subID, err)
	}
	return s.upsertFromSubscription(ctx, sub, actor)
}

// metadataUserID reads sub.Metadata["user_id"] — always the buying
// account, premium or guild plan alike (spec §2.2).
func metadataUserID(sub *stripe.Subscription) (int64, bool) {
	uid, err := strconv.ParseInt(sub.Metadata["user_id"], 10, 64)
	if err != nil {
		return 0, false
	}
	return uid, true
}

// subscriptionPeriodEnd reads a subscription's current billing period
// end from its first line item (Ruling E): stripe-go v82 / Stripe API
// version 2025-08-27 moved current_period_end off the Subscription
// object onto each SubscriptionItem. Every subscription this integration
// creates has exactly one item (one price, quantity 1, spec §2.2).
func subscriptionPeriodEnd(sub *stripe.Subscription) (time.Time, bool) {
	if sub.Items == nil || len(sub.Items.Data) == 0 {
		return time.Time{}, false
	}
	return time.Unix(sub.Items.Data[0].CurrentPeriodEnd, 0).UTC(), true
}

// upsertFromSubscription is the only writer of entitlements from Stripe
// state (spec RULING 8): plan and subject come from sub.Metadata, set at
// Checkout via subscription_data.metadata (spec §2.2) and therefore
// present on the Subscription object for its whole lifetime, independent
// of which event triggered this call — never from the event's own
// embedded snapshot.
func (s *Service) upsertFromSubscription(ctx context.Context, sub *stripe.Subscription, actor string) error {
	plan := sub.Metadata["plan"]
	if plan != entitlements.PlanPremium && plan != entitlements.PlanGuild {
		return fmt.Errorf("billing: subscription %s carries no valid plan metadata", sub.ID)
	}
	billingUserID, ok := metadataUserID(sub)
	if !ok {
		return fmt.Errorf("billing: subscription %s carries no valid user_id metadata", sub.ID)
	}
	p := entitlements.StripeUpsert{
		Plan: plan, Status: string(sub.Status), CancelAtPeriodEnd: sub.CancelAtPeriodEnd,
		StripeSubscriptionID: sub.ID, BillingUserID: billingUserID, Actor: actor,
	}
	if end, ok := subscriptionPeriodEnd(sub); ok {
		p.CurrentPeriodEnd = &end
	}
	if plan == entitlements.PlanGuild {
		guildID, err := strconv.ParseInt(sub.Metadata["guild_id"], 10, 64)
		if err != nil {
			return fmt.Errorf("billing: guild subscription %s carries no valid guild_id metadata: %w", sub.ID, err)
		}
		p.GuildID = &guildID
	} else {
		p.UserID = &billingUserID
	}
	return s.Entitlements.UpsertStripe(ctx, p)
}
