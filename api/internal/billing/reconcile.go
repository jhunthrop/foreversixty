// api/internal/billing/reconcile.go
package billing

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// reconcileActor tags every entitlement_audit/entitlement_anomalies row
// this job writes.
const reconcileActor = "stripe_reconcile"

// reconcileProducts is every Stripe Product id the orphan check (below)
// scans — both plans' products, spec §2.8's "our product ids" scope.
var reconcileProducts = []string{catalog.ProductIDPremium, catalog.ProductIDGuild, catalog.ProductIDRaidGuild}

// Reconcile is the stripe-reconcile job's whole job (spec §2.8), now
// two-directional (security review fix, 2026-09-21; the original design
// deferred the Stripe→database half — see the spec's own "RULING" text
// for why — but never actually implemented even the logging it promised):
//
//  1. Database → Stripe (unchanged): every entitlements row still
//     believed live is re-fetched fresh from Stripe and upserted exactly
//     as a webhook event would, healing anything Stripe's own three-day
//     retry window never successfully delivered.
//  2. Expired pending_checkouts rows are swept — the same backstop this
//     job already provides for webhooks, extended to an abandoned
//     Checkout Session nobody ever completed (security review fix).
//  3. Stripe → database: every subscription Stripe currently considers
//     live against one of our own products, with no matching
//     entitlements.stripe_subscription_id, is recorded as an
//     orphan_subscription anomaly for a human to investigate — this is
//     what would have caught the double-billing finding's orphaned first
//     subscription, which no database-→-Stripe pass can ever see (it has
//     no row to start from).
//  4. stripe_events rows past stripeEventRetention are deleted (issue #3
//     M8): nothing reads a raw payload once Stripe can no longer resend
//     its event, and each one carries the buyer's name, email and address.
//
// Every step runs whatever an earlier one hit, and a bad row never stops
// the rows after it (issue #3 M6): each failure is logged with its
// subscription id and the returned error joins them all. Returns how many
// entitlements rows step 1 healed.
func (s *Service) Reconcile(ctx context.Context) (int, error) {
	ids, err := s.Entitlements.StripeSubscriptionIDs(ctx)
	if err != nil {
		return 0, fmt.Errorf("billing: reconcile: list: %w", err)
	}
	known := make(map[string]bool, len(ids))
	var errs []error
	healed := 0
	for _, id := range ids {
		known[id] = true
		if err := s.upsertSubscription(ctx, id, reconcileActor); err != nil {
			errs = append(errs, s.reconcileFailure("upsert", id, err))
			continue
		}
		healed++
	}
	if outgoing, err := s.Store.OutgoingTransferSubscriptionIDs(ctx); err != nil {
		errs = append(errs, s.reconcileFailure("outgoing transfers", "", err))
	} else {
		errs = append(errs, s.reportOrphanSubscriptions(ctx, known, outgoingSet(outgoing))...)
	}
	if _, err := s.Store.SweepExpiredPendingCheckouts(ctx); err != nil {
		errs = append(errs, s.reconcileFailure("sweep pending checkouts", "", err))
	}
	if _, err := s.Store.DeleteStripeEventsBefore(ctx, s.now().Add(-stripeEventRetention)); err != nil {
		errs = append(errs, s.reconcileFailure("delete old stripe events", "", err))
	}
	return healed, errors.Join(errs...)
}

// stripeEventRetention is how long a stripe_events row is kept: Stripe
// keeps an event, and so can resend it, for 30 days. The webhook's
// idempotency gate needs the row only while a resend can arrive, and
// reconcile heals from Stripe itself, never from a stored payload. A
// resend after deletion is processed again, which is safe only because
// every handler re-fetches the subscription from Stripe; keep it so.
const stripeEventRetention = 30 * 24 * time.Hour

// reconcileFailure logs one reconcile failure and wraps it for the job's
// joined error. subscriptionID is "" for a step that is not about one row.
func (s *Service) reconcileFailure(step, subscriptionID string, err error) error {
	s.logger().Error("stripe-reconcile", "step", step, "subscription_id", subscriptionID, "err", err)
	if subscriptionID == "" {
		return fmt.Errorf("billing: reconcile: %s: %w", step, err)
	}
	return fmt.Errorf("billing: reconcile: %s %s: %w", step, subscriptionID, err)
}

// reportOrphanSubscriptions is Reconcile's Stripe→database half (spec
// §2.8's last paragraph, security review fix 2026-09-21): known is every
// stripe_subscription_id Reconcile already believes live. A product whose
// listing fails, or an anomaly that cannot be written, is reported and the
// rest still run.
func (s *Service) reportOrphanSubscriptions(ctx context.Context, known, outgoing map[string]bool) []error {
	var errs []error
	for _, productID := range reconcileProducts {
		subs, err := s.Gateway.ListSubscriptions(ctx, productID)
		if err != nil {
			errs = append(errs, s.reconcileFailure("list subscriptions for "+productID, "", err))
			continue
		}
		plan := entitlements.PlanPremium
		if productID != catalog.ProductIDPremium {
			plan = entitlements.PlanGuild
		}
		for _, sub := range subs {
			if known[sub.ID] || (outgoing[sub.ID] && sub.CancelAtPeriodEnd) {
				continue
			}
			if err := s.Entitlements.RecordAnomaly(ctx, nil, nil, plan,
				entitlements.AnomalyOrphanSubscription, sub.ID, reconcileActor); err != nil {
				errs = append(errs, s.reconcileFailure("record orphan", sub.ID, err))
			}
		}
	}
	return errs
}

// outgoingSet indexes the subscriptions accepted transfers retire. One
// counts as known only while Stripe shows it canceling at period end; a
// re-enabled one is an orphan again.
func outgoingSet(ids []string) map[string]bool {
	set := make(map[string]bool, len(ids))
	for _, id := range ids {
		set[id] = true
	}
	return set
}
