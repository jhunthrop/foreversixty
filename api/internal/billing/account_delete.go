// api/internal/billing/account_delete.go
package billing

import (
	"context"
	"errors"
	"fmt"

	stripe "github.com/stripe/stripe-go/v82"
)

// ErrBillingUnavailable is CancelAccountSubscriptions' answer when the account has a live
// subscription and this deployment has no Stripe gateway to cancel it with: the deletion
// stops before anything is removed.
var ErrBillingUnavailable = errors.New("billing: stripe is not configured; cannot cancel a live subscription")

// LiveStripeSQL is the predicate for an entitlements row (unaliased, or the only table in
// scope) whose Stripe subscription can still charge: every status but Stripe's two terminal
// ones, canceled and incomplete_expired. The account deletion's guild check reads it too.
const LiveStripeSQL = `source = 'stripe' and stripe_subscription_id is not null
	and status not in ('canceled', 'incomplete_expired')`

// liveAccountSubscriptionsSQL is every live Stripe subscription the account pays for: its own
// Premium row and any guild plan it is the billing user of.
const liveAccountSubscriptionsSQL = `
	select stripe_subscription_id from entitlements
	where ` + LiveStripeSQL + ` and (user_id = $1 or billing_user_id = $1)
	order by id`

// CancelAccountSubscriptions cancels, at once, every live Stripe subscription userID pays
// for: the account deleting itself (DELETE /v1/me). There is no automatic refund (the refund
// policy is a hand step), and the Stripe Customer stays at Stripe, where billing records are
// retained for tax and legal reasons; the deletion removes only our link to it. A
// subscription Stripe already holds as ended is skipped, so a retried deletion is safe.
func (s *Service) CancelAccountSubscriptions(ctx context.Context, userID int64) error {
	ids, err := s.Store.liveAccountSubscriptions(ctx, userID)
	if err != nil {
		return err
	}
	if len(ids) == 0 {
		return nil
	}
	if s.Gateway == nil {
		return ErrBillingUnavailable
	}
	for _, id := range ids {
		if err := s.cancelNowIfLive(ctx, id); err != nil {
			return err
		}
	}
	return nil
}

// cancelNowIfLive re-reads the subscription from Stripe (never trusting our copy of its
// status) and cancels it unless Stripe already ended it or has no such subscription.
func (s *Service) cancelNowIfLive(ctx context.Context, id string) error {
	sub, err := s.Gateway.GetSubscription(ctx, id)
	if errors.Is(err, ErrSubscriptionNotFound) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("billing: account deletion: read %s: %w", id, err)
	}
	if sub.Status == stripe.SubscriptionStatusCanceled || sub.Status == stripe.SubscriptionStatusIncompleteExpired {
		return nil
	}
	if err := s.Gateway.CancelSubscriptionNow(ctx, id); err != nil {
		return fmt.Errorf("billing: account deletion: %w", err)
	}
	s.logger().Info("billing", "op", "account_delete_cancel", "subscription_id", id)
	return nil
}

// liveAccountSubscriptions runs liveAccountSubscriptionsSQL.
func (s *Store) liveAccountSubscriptions(ctx context.Context, userID int64) ([]string, error) {
	rows, err := s.Pool.Query(ctx, liveAccountSubscriptionsSQL, userID)
	if err != nil {
		return nil, fmt.Errorf("billing: live subscriptions of %d: %w", userID, err)
	}
	defer rows.Close()
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("billing: live subscriptions of %d: %w", userID, err)
		}
		ids = append(ids, id)
	}
	return ids, rows.Err()
}

// userExists reports whether the account a subscription's metadata names still exists. A
// deleted account's subscription keeps sending events (the cancellation itself, a late
// invoice); they are acknowledged rather than written, because the entitlement row would
// reference an account that is gone, and a 500 would have Stripe retry for days.
func (s *Store) userExists(ctx context.Context, userID int64) (bool, error) {
	var ok bool
	if err := s.Pool.QueryRow(ctx, `select exists (select 1 from users where id = $1)`, userID).Scan(&ok); err != nil {
		return false, fmt.Errorf("billing: account %d exists: %w", userID, err)
	}
	return ok, nil
}
