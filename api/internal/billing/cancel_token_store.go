// api/internal/billing/cancel_token_store.go
package billing

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// cancelTokenBytes is the one-click cancel link's randomness (spec 4.3).
const cancelTokenBytes = 32

// purposeCancelTrial is the only thing a cancel link can do.
const purposeCancelTrial = "cancel_trial"

// CancelTarget is the subscription a cancel link names, as the link's page shows it. Guild is
// empty for Premium.
type CancelTarget struct {
	EntitlementID        int64
	StripeSubscriptionID string
	Plan                 string
	Tier                 string
	Guild                string
	TrialEnd             time.Time
	PriceLookupKey       string
	UnitAmountCents      int64
	CancelAtPeriodEnd    bool
}

// hashToken is the stored form of a token: its SHA-256, so the table never holds a working link.
func hashToken(token string) []byte {
	sum := sha256.Sum256([]byte(token))
	return sum[:]
}

// decodeToken is a token's bytes, for checking its shape before any read.
func decodeToken(token string) ([]byte, error) { return base64.RawURLEncoding.DecodeString(token) }

// CreateCancelToken mints a cancel link's token for one trialing subscription, good until the
// trial ends. The token is answered once and kept only as its hash.
func (s *Store) CreateCancelToken(ctx context.Context, entitlementID int64, subscriptionID string, expires time.Time) (string, error) {
	raw := make([]byte, cancelTokenBytes)
	if _, err := rand.Read(raw); err != nil {
		return "", fmt.Errorf("billing: cancel token: %w", err)
	}
	token := base64.RawURLEncoding.EncodeToString(raw)
	if _, err := s.Pool.Exec(ctx, `insert into billing_cancel_tokens
		(token_hash, entitlement_id, stripe_subscription_id, purpose, expires_at) values ($1, $2, $3, $4, $5)`,
		hashToken(token), entitlementID, subscriptionID, purposeCancelTrial, expires); err != nil {
		return "", fmt.Errorf("billing: store cancel token for %d: %w", entitlementID, err)
	}
	return token, nil
}

// DeleteCancelToken drops a token whose email was never sent.
func (s *Store) DeleteCancelToken(ctx context.Context, token string) error {
	if _, err := s.Pool.Exec(ctx, `delete from billing_cancel_tokens where token_hash = $1`, hashToken(token)); err != nil {
		return fmt.Errorf("billing: delete cancel token: %w", err)
	}
	return nil
}

// cancelTargetQuery reads a token's subscription while the token is unused and unexpired and the
// entitlement still trials on the same subscription; $1 is the hash, $2 now.
const cancelTargetQuery = `
	select e.id, t.stripe_subscription_id, e.plan, coalesce(e.tier, ''), coalesce(g.name, ''), e.trial_end,
	       coalesce(e.price_lookup_key, ''), coalesce(e.unit_amount_cents, 0), e.cancel_at_period_end
	from billing_cancel_tokens t
	join entitlements e on e.id = t.entitlement_id and e.stripe_subscription_id = t.stripe_subscription_id
	left join guilds g on g.id = e.guild_id
	where t.token_hash = $1 and t.used_at is null and t.expires_at > $2
	  and e.source = 'stripe' and e.status = 'trialing' and e.trial_end is not null`

// CancelTokenTarget is what a token names, ok false for a used, expired or unknown token or a
// subscription that no longer trials.
func (s *Store) CancelTokenTarget(ctx context.Context, token string, now time.Time) (CancelTarget, bool, error) {
	var t CancelTarget
	err := s.Pool.QueryRow(ctx, cancelTargetQuery, hashToken(token), now).Scan(&t.EntitlementID, &t.StripeSubscriptionID,
		&t.Plan, &t.Tier, &t.Guild, &t.TrialEnd, &t.PriceLookupKey, &t.UnitAmountCents, &t.CancelAtPeriodEnd)
	if errors.Is(err, pgx.ErrNoRows) {
		return CancelTarget{}, false, nil
	}
	if err != nil {
		return CancelTarget{}, false, fmt.Errorf("billing: cancel token target: %w", err)
	}
	return t, true, nil
}

// UseCancelToken marks the token used; false when another request used it first.
func (s *Store) UseCancelToken(ctx context.Context, token string, now time.Time) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `update billing_cancel_tokens set used_at = $2
		where token_hash = $1 and used_at is null and expires_at > $2`, hashToken(token), now)
	if err != nil {
		return false, fmt.Errorf("billing: use cancel token: %w", err)
	}
	return tag.RowsAffected() == 1, nil
}

// ReleaseCancelToken undoes UseCancelToken when Stripe refused the cancel, so the link still works.
func (s *Store) ReleaseCancelToken(ctx context.Context, token string) error {
	if _, err := s.Pool.Exec(ctx, `update billing_cancel_tokens set used_at = null where token_hash = $1`, hashToken(token)); err != nil {
		return fmt.Errorf("billing: release cancel token: %w", err)
	}
	return nil
}
