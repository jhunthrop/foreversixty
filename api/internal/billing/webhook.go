// api/internal/billing/webhook.go
package billing

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/webhook"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
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
		// The library also refuses an event whose API version is not the
		// one it was built for (stripe.APIVersion), so the endpoint is
		// created pinned to it (api/README.md, "Stripe: going live"); the
		// reason is logged because the client sees only a 400.
		s.Log.Warn("billing: webhook rejected", "err", err)
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "the webhook signature could not be verified", nil)
		return
	}
	// Serialized per event id: Stripe's own retry can outrun a slow
	// first attempt that has not yet responded, and without this lock
	// both deliveries could observe RecordEventOnce's processed_at as
	// still null and both call handleEvent concurrently.
	err = s.Store.WithEventLock(r.Context(), event.ID, func(ctx context.Context) error {
		proceed, err := s.Store.RecordEventOnce(ctx, event.ID, string(event.Type), payload)
		if err != nil {
			return err
		}
		if !proceed {
			// Already fully processed (a genuine redelivery of a
			// completed event, or Stripe's own retry after a
			// slow-but-eventually-200 first attempt): no reprocessing.
			return nil
		}
		if err := s.handleEvent(ctx, event); err != nil {
			return err
		}
		return s.Store.MarkEventProcessed(ctx, event.ID)
	})
	if err != nil {
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
	if err := s.upsertSubscription(ctx, subID, actor); err != nil {
		return err
	}
	// A completed Checkout Session no longer needs to block a later,
	// genuinely new checkout for the same guild — cleared whether this
	// upsert wrote the row or, on a caught duplicate, left it untouched
	// (security review fix, 2026-09-21): either way the session itself
	// did complete, which is what pending_checkouts tracks.
	if sessionID := stripeRefID(event.Data.Object["id"]); sessionID != "" {
		if err := s.Store.DeletePendingCheckout(ctx, sessionID); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) onSubscriptionEvent(ctx context.Context, event stripe.Event, actor string) error {
	return s.upsertSubscription(ctx, stripeRefID(event.Data.Object["id"]), actor)
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
	subID := invoiceSubscriptionID(event.Data.Object)
	if subID == "" {
		return nil
	}
	subject, ok, err := s.upsertSubscriptionOf(ctx, subID, actor)
	if err != nil || !ok {
		return err
	}
	if event.Type == "invoice.payment_failed" {
		s.publishPaymentFailed(ctx, subject.billingUserID, stripeRefID(event.Data.Object["id"]), event.ID)
	}
	return nil
}

// metadataUserID reads sub.Metadata["user_id"] — always the buying
// account, premium or guild plan alike (spec §2.2).
// invoiceSubscriptionID is the subscription an invoice bills. Since API
// version 2025-03-31.basil an invoice carries it at
// parent.subscription_details.subscription; older shapes carried a
// top-level subscription, which is still read so a replayed older event
// resolves the same way.
func invoiceSubscriptionID(invoice map[string]any) string {
	if parent, ok := invoice["parent"].(map[string]any); ok {
		if details, ok := parent["subscription_details"].(map[string]any); ok {
			if id := stripeRefID(details["subscription"]); id != "" {
				return id
			}
		}
	}
	return stripeRefID(invoice["subscription"])
}

// errNotOurSubscription marks a subscription this service did not sell:
// one with no plan or subject in its metadata (made in the Dashboard, or
// edited by support). Its events are acknowledged and logged rather than
// failed, because a 500 makes Stripe retry for three days and can disable
// the endpoint, and no retry would ever make the metadata appear.
var errNotOurSubscription = errors.New("billing: subscription carries no plan metadata")

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

// upsertSubscription is the only writer of entitlements from Stripe
// state (spec RULING 8): plan and subject come from the subscription's
// metadata, set at Checkout via subscription_data.metadata (spec §2.2)
// and therefore present for its whole lifetime, independent of which
// event triggered this call — never from the event's own snapshot.
//
// The subscription is fetched twice (issue #3 M3). The first fetch only
// learns the plan and subject, which pick the lock; the write uses the
// second fetch, taken inside the lock, so a slower delivery cannot write
// a status it read before a faster one committed.
func (s *Service) upsertSubscription(ctx context.Context, subscriptionID, actor string) error {
	_, _, err := s.upsertSubscriptionOf(ctx, subscriptionID, actor)
	return err
}

// upsertSubscriptionOf is upsertSubscription answering the subject it upserted; ok is false
// for a subscription this service did not sell.
func (s *Service) upsertSubscriptionOf(ctx context.Context, subscriptionID, actor string) (subscriptionSubject, bool, error) {
	first, err := s.Gateway.GetSubscription(ctx, subscriptionID)
	if err != nil {
		return subscriptionSubject{}, false, fmt.Errorf("billing: get subscription %s: %w", subscriptionID, err)
	}
	subject, ok := s.subjectOf(first)
	if !ok {
		return subscriptionSubject{}, false, nil
	}
	return subject, true, subject.withLock(ctx, s.Store, func(ctx context.Context) error {
		sub, err := s.Gateway.GetSubscription(ctx, subscriptionID)
		if err != nil {
			return fmt.Errorf("billing: get subscription %s inside lock: %w", subscriptionID, err)
		}
		if again, err := parseSubject(sub); err != nil || again != subject {
			s.logger().Warn("billing: subscription metadata changed between fetches", "subscription", sub.ID)
			return nil
		}
		p, err := s.upsertFor(ctx, sub, subject, actor)
		if err != nil {
			return err
		}
		return s.applyUpsert(ctx, p, sub)
	})
}

// subscriptionSubject is the row owner a subscription bills: a guild for
// the guild plan, the buying user for premium.
type subscriptionSubject struct {
	plan          string
	billingUserID int64
	guildID       int64 // set only for the guild plan
}

func (sj subscriptionSubject) withLock(ctx context.Context, store *Store, fn func(context.Context) error) error {
	if sj.plan == entitlements.PlanGuild {
		return store.WithGuildLock(ctx, sj.guildID, fn)
	}
	return store.WithUserLock(ctx, sj.billingUserID, fn)
}

// subjectOf is parseSubject for the first fetch: a subscription without
// our metadata is acknowledged and logged rather than failed, because no
// retry would ever make the metadata appear.
func (s *Service) subjectOf(sub *stripe.Subscription) (subscriptionSubject, bool) {
	subject, err := parseSubject(sub)
	if err != nil {
		s.Log.Warn("billing: subscription is not ours; event acknowledged", "subscription", sub.ID, "reason", err)
		return subscriptionSubject{}, false
	}
	return subject, true
}

// parseSubject reads plan, buyer and guild from sub.Metadata.
func parseSubject(sub *stripe.Subscription) (subscriptionSubject, error) {
	plan := sub.Metadata["plan"]
	if plan != entitlements.PlanPremium && plan != entitlements.PlanGuild {
		return subscriptionSubject{}, errNotOurSubscription
	}
	billingUserID, ok := metadataUserID(sub)
	if !ok {
		return subscriptionSubject{}, errors.New("no valid user_id metadata")
	}
	subject := subscriptionSubject{plan: plan, billingUserID: billingUserID}
	if plan == entitlements.PlanGuild {
		guildID, err := strconv.ParseInt(sub.Metadata["guild_id"], 10, 64)
		if err != nil {
			return subscriptionSubject{}, errors.New("no valid guild_id metadata")
		}
		subject.guildID = guildID
	}
	return subject, nil
}

// buildUpsert is the upsert a subscription calls for, and whether its
// price is one the catalog sells for the subscription's plan. A guild row
// of an unrecognised price takes the plain guild tier.
func buildUpsert(sub *stripe.Subscription, subject subscriptionSubject, actor string) (entitlements.StripeUpsert, bool) {
	p := entitlements.StripeUpsert{
		Plan: subject.plan, Status: string(sub.Status), CancelAtPeriodEnd: sub.CancelAtPeriodEnd,
		StripeSubscriptionID: sub.ID, BillingUserID: subject.billingUserID, Actor: actor,
		// spec §2.9 RULING 10's "Take over billing" handoff: the new
		// subscription is meant to overwrite the guild's row, not be
		// flagged as a duplicate (StripeUpsert.Transfer's own doc).
		Transfer: sub.Metadata["intent"] == "transfer",
	}
	if sub.Created != 0 { // time.Unix(0, 0) is 1970, not the zero time "unknown" needs
		p.SubscriptionCreated = time.Unix(sub.Created, 0).UTC()
	}
	if end, ok := subscriptionPeriodEnd(sub); ok {
		p.CurrentPeriodEnd = &end
	}
	applyPriceFacts(&p, sub)
	offer, known := catalog.Offer(""), false
	if price, ok := catalog.Lookup(p.PriceLookupKey); ok {
		offer, known = price.Offer, true
	}
	if subject.plan == entitlements.PlanGuild {
		p.GuildID = &subject.guildID
		p.Tier = catalog.TierGuild
		if known && offer.Tier() != "" {
			p.Tier = offer.Tier()
		} else {
			known = false
		}
	} else {
		p.UserID = &subject.billingUserID
		known = known && offer == catalog.OfferPremium
	}
	return p, known
}

// applyPriceFacts copies the first item's price key and amount, and the
// trial end, onto the upsert.
func applyPriceFacts(p *entitlements.StripeUpsert, sub *stripe.Subscription) {
	if price := itemPrice(sub); price != nil {
		p.PriceLookupKey = price.LookupKey
		p.UnitAmountCents = price.UnitAmount
	}
	if sub.TrialEnd != 0 {
		end := time.Unix(sub.TrialEnd, 0).UTC()
		p.TrialEnd = &end
	}
}

// itemPrice is the price of sub's first line item, nil when it has none.
func itemPrice(sub *stripe.Subscription) *stripe.Price {
	if sub.Items == nil || len(sub.Items.Data) == 0 {
		return nil
	}
	return sub.Items.Data[0].Price
}

// upsertFor is buildUpsert plus the anomaly a priced subscription outside
// the catalog for its plan deserves. A subscription carrying no price at
// all has no price facts to judge and is left alone.
func (s *Service) upsertFor(ctx context.Context, sub *stripe.Subscription, subject subscriptionSubject, actor string) (entitlements.StripeUpsert, error) {
	p, known := buildUpsert(sub, subject, actor)
	price := itemPrice(sub)
	if known || price == nil {
		return p, nil
	}
	s.logger().Error("billing", "op", "unknown_price", "subscription_id", sub.ID,
		"plan", subject.plan, "price_id", price.ID, "lookup_key", p.PriceLookupKey)
	if err := s.Entitlements.RecordAnomaly(ctx, p.UserID, p.GuildID, subject.plan,
		entitlements.AnomalyUnknownPrice, sub.ID, actor); err != nil {
		return p, err
	}
	return p, nil
}

// applyUpsert is upsertSubscription's locked half: the entitlements
// write, then the action its outcome calls for. Always called from
// inside WithGuildLock/WithUserLock.
func (s *Service) applyUpsert(ctx context.Context, p entitlements.StripeUpsert, sub *stripe.Subscription) error {
	if retired, err := s.retiredByTransfer(ctx, p); err != nil || retired {
		return err
	}
	if p.Transfer {
		state, err := s.transferStateFor(ctx, *p.GuildID, sub.Metadata["from_subscription_id"])
		if err != nil {
			return err
		}
		p.OfferGone = state != transferOffered && state != transferAccepted
	}
	result, err := s.Entitlements.UpsertStripe(ctx, p)
	if err != nil {
		return err
	}
	switch result.Outcome {
	case entitlements.OutcomeDuplicate:
		return s.cancelDuplicate(ctx, sub, result.ExistingSubscriptionID)
	case entitlements.OutcomeStaleTransfer:
		s.logger().Info("billing", "op", "stale_transfer_ignored",
			"subscription_id", sub.ID, "kept_subscription_id", result.ExistingSubscriptionID)
	case entitlements.OutcomeStrayTransfer:
		return s.cancelStrayTransfer(ctx, p, sub)
	case entitlements.OutcomeDeadNewcomer:
		s.logger().Info("billing", "op", "dead_newcomer_ignored",
			"subscription_id", sub.ID, "kept_subscription_id", result.ExistingSubscriptionID)
	case entitlements.OutcomeWritten:
		return s.afterWritten(ctx, p, sub)
	}
	return nil
}

// retiredByTransfer reports whether the subscription is the one an
// accepted transfer retired. It keeps sending events until its period
// ends; they are acknowledged and written nowhere, so they neither record
// a false duplicate nor cancel it a second time.
func (s *Service) retiredByTransfer(ctx context.Context, p entitlements.StripeUpsert) (bool, error) {
	if p.GuildID == nil {
		return false, nil
	}
	state, err := s.Store.TransferState(ctx, *p.GuildID, p.StripeSubscriptionID)
	if err != nil || state != transferAccepted {
		return false, err
	}
	s.logger().Info("billing", "op", "outgoing_transfer_event_ignored", "subscription_id", p.StripeSubscriptionID)
	return true, nil
}

// afterWritten runs the follow-ups a written upsert owes: the guild's
// trial record, a transfer's retirement of the outgoing subscription,
// and a premium comeback's redemption. Each is idempotent.
func (s *Service) afterWritten(ctx context.Context, p entitlements.StripeUpsert, sub *stripe.Subscription) error {
	if p.Plan == entitlements.PlanGuild {
		if p.Transfer {
			return s.completeTransfer(ctx, p, sub.Metadata["from_subscription_id"])
		}
		if p.TrialEnd != nil {
			return s.Store.RecordGuildTrial(ctx, *p.GuildID, sub.ID)
		}
		return nil
	}
	raw, ok := sub.Metadata["comeback_id"]
	if !ok {
		return nil
	}
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		s.logger().Warn("billing: ignoring malformed comeback_id", "subscription", sub.ID)
		return nil
	}
	return s.Store.MarkComebackUsed(ctx, id, p.BillingUserID)
}

// transferStateFor is the state of the transfer away from the
// subscription a transfer upsert names; transferNone when it names none.
func (s *Service) transferStateFor(ctx context.Context, guildID int64, from string) (string, error) {
	if from == "" {
		return transferNone, nil
	}
	return s.Store.TransferState(ctx, guildID, from)
}

// cancelStrayTransfer is the other half of OutcomeStrayTransfer: a
// transfer subscription whose offer was withdrawn, declined or replaced
// while its Checkout was open is recorded and canceled at period end (it
// is still in its trial, so nothing was charged). One already set to end,
// or no longer live, needs no cancel.
func (s *Service) cancelStrayTransfer(ctx context.Context, p entitlements.StripeUpsert, sub *stripe.Subscription) error {
	settled := sub.CancelAtPeriodEnd || !entitlements.StatusGrantsAccess(string(sub.Status))
	// A stray already set to end is a later event of one handled before: Warn.
	level := slog.LevelError
	if settled {
		level = slog.LevelWarn
	}
	s.logger().Log(ctx, level, "billing", "op", "stray_transfer_canceled", "subscription_id", sub.ID,
		"from_subscription_id", sub.Metadata["from_subscription_id"], "guild_id", *p.GuildID)
	if err := s.Entitlements.RecordAnomaly(ctx, nil, p.GuildID, entitlements.PlanGuild,
		entitlements.AnomalyUnexpectedTransfer, p.StripeSubscriptionID, p.Actor); err != nil {
		return err
	}
	if settled {
		return nil
	}
	if err := s.Gateway.CancelSubscriptionAtPeriodEnd(ctx, sub.ID); err != nil {
		return fmt.Errorf("billing: cancel stray transfer subscription %s: %w", sub.ID, err)
	}
	return nil
}

// completeTransfer cancels the outgoing subscription at period end and
// marks the transfer accepted. An outgoing subscription that has already
// ended at Stripe needs no cancel (Stripe would refuse one). An already
// accepted transfer is left alone, so a later event for the same
// subscription cancels nothing twice; any other state means the row was
// already this subscription's (cancelStrayTransfer keeps strays off it)
// and a later offer has replaced the transfer row, so there is nothing
// to complete.
func (s *Service) completeTransfer(ctx context.Context, p entitlements.StripeUpsert, from string) error {
	guildID := *p.GuildID
	state, err := s.transferStateFor(ctx, guildID, from)
	if err != nil || state != transferOffered {
		return err
	}
	ended, err := s.outgoingEnded(ctx, from)
	if err != nil {
		return err
	}
	if !ended {
		if err := s.Gateway.CancelSubscriptionAtPeriodEnd(ctx, from); err != nil {
			return fmt.Errorf("billing: cancel outgoing subscription %s: %w", from, err)
		}
	}
	return s.Store.MarkTransferAccepted(ctx, guildID, from)
}

// cancelDuplicate is the other half of a caught duplicate (security
// review fix, 2026-09-21): UpsertStripe kept the existing row and
// recorded the anomaly; the newcomer must not keep billing with nothing
// in our database honoring it, so it is canceled at Stripe (at period
// end: whoever paid keeps what they paid for). A failed cancel is
// returned so Stripe retries; the anomaly is recorded once per
// subscription, so the retry adds no second row. No secret or metadata
// dump in the log line — subscription ids only.
func (s *Service) cancelDuplicate(ctx context.Context, sub *stripe.Subscription, keptID string) error {
	s.logger().Error("billing", "op", "duplicate_subscription",
		"newcomer_subscription_id", sub.ID, "kept_subscription_id", keptID,
		"plan", sub.Metadata["plan"])
	if err := s.Gateway.CancelSubscriptionAtPeriodEnd(ctx, sub.ID); err != nil {
		return fmt.Errorf("billing: cancel duplicate subscription %s: %w", sub.ID, err)
	}
	return nil
}
