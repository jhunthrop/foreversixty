// api/internal/billing/reconcile.go
package billing

import (
	"context"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// reconcileActor tags every entitlement_audit/entitlement_anomalies row
// this job writes.
const reconcileActor = "stripe_reconcile"

// reconcileProducts is every Stripe Product id the orphan check (below)
// scans — both plans' products, spec §2.8's "our product ids" scope.
var reconcileProducts = []string{ProductIDPremium, ProductIDGuild}

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
//
// Returns how many entitlements rows step 1 processed.
func (s *Service) Reconcile(ctx context.Context) (int, error) {
	ids, err := s.Entitlements.StripeSubscriptionIDs(ctx)
	if err != nil {
		return 0, fmt.Errorf("billing: reconcile: list: %w", err)
	}
	known := make(map[string]bool, len(ids))
	for _, id := range ids {
		known[id] = true
		sub, err := s.Gateway.GetSubscription(ctx, id)
		if err != nil {
			return 0, fmt.Errorf("billing: reconcile: get %s: %w", id, err)
		}
		if err := s.upsertFromSubscription(ctx, sub, reconcileActor); err != nil {
			return 0, fmt.Errorf("billing: reconcile: upsert %s: %w", id, err)
		}
	}
	if _, err := s.Store.SweepExpiredPendingCheckouts(ctx); err != nil {
		return len(ids), fmt.Errorf("billing: reconcile: sweep pending checkouts: %w", err)
	}
	if err := s.reportOrphanSubscriptions(ctx, known); err != nil {
		return len(ids), err
	}
	return len(ids), nil
}

// reportOrphanSubscriptions is Reconcile's Stripe→database half (spec
// §2.8's last paragraph, security review fix 2026-09-21): known is every
// stripe_subscription_id Reconcile already believes live.
func (s *Service) reportOrphanSubscriptions(ctx context.Context, known map[string]bool) error {
	for _, productID := range reconcileProducts {
		subs, err := s.Gateway.ListSubscriptions(ctx, productID)
		if err != nil {
			return fmt.Errorf("billing: reconcile: list subscriptions for %s: %w", productID, err)
		}
		for _, sub := range subs {
			if known[sub.ID] {
				continue
			}
			plan := entitlements.PlanPremium
			if productID == ProductIDGuild {
				plan = entitlements.PlanGuild
			}
			if err := s.Entitlements.RecordAnomaly(ctx, nil, nil, plan,
				entitlements.AnomalyOrphanSubscription, sub.ID, reconcileActor); err != nil {
				return fmt.Errorf("billing: reconcile: record orphan %s: %w", sub.ID, err)
			}
		}
	}
	return nil
}
