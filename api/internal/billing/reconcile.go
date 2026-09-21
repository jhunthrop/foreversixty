// api/internal/billing/reconcile.go
package billing

import (
	"context"
	"fmt"
)

// Reconcile is the stripe-reconcile job's whole job (spec §2.8): every
// entitlements row still believed live is re-fetched fresh from Stripe
// and upserted exactly as a webhook event would, healing anything
// Stripe's own three-day retry window never successfully delivered.
// One-directional (database → Stripe), not the reverse. Returns how many
// rows it processed.
func (s *Service) Reconcile(ctx context.Context) (int, error) {
	ids, err := s.Entitlements.StripeSubscriptionIDs(ctx)
	if err != nil {
		return 0, fmt.Errorf("billing: reconcile: list: %w", err)
	}
	for _, id := range ids {
		sub, err := s.Gateway.GetSubscription(ctx, id)
		if err != nil {
			return 0, fmt.Errorf("billing: reconcile: get %s: %w", id, err)
		}
		if err := s.upsertFromSubscription(ctx, sub, "stripe_reconcile"); err != nil {
			return 0, fmt.Errorf("billing: reconcile: upsert %s: %w", id, err)
		}
	}
	return len(ids), nil
}
