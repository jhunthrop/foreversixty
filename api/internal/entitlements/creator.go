package entitlements

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// SourceCreator marks the Premium the creator program gives a creator's own account (the creators
// migration). It is written only by GrantCreatorPremium and ended only by EndCreatorPremium, so it never
// replaces a live paid subscription or another grant, and ending it never touches one either.
// The account page reads it as a grant (PersonalBilling).
const SourceCreator = "creator"

// GrantCreatorPremium gives userID open-ended Premium as a creator. A live row from any other
// source (a paid subscription, a hand grant) is left exactly as it is, and granted is false.
func (s *Store) GrantCreatorPremium(ctx context.Context, userID int64, actor string) (granted bool, err error) {
	subj := Subject{UserID: &userID}
	err = s.withTx(ctx, func(tx pgx.Tx) error {
		existing, isNew, err := existingEntitlement(ctx, tx, subj, PlanPremium)
		if err != nil {
			return err
		}
		if !isNew && existing.active && existing.source != SourceCreator {
			return nil
		}
		var id int64
		if err := tx.QueryRow(ctx, `
			insert into entitlements (user_id, plan, source, status, grant_note)
			values ($1, 'premium', 'creator', 'active', 'creator program')
			on conflict (user_id, plan) do update set
				source = 'creator', status = 'active', current_period_end = null, grant_note = 'creator program',
				features = null, updated_at = now(), billing_user_id = null, cancel_at_period_end = false,
				stripe_subscription_id = null, grace_until = null, tier = null, price_lookup_key = null,
				unit_amount_cents = null, trial_end = null
			returning id`, userID).Scan(&id); err != nil {
			return fmt.Errorf("entitlements: creator grant: write: %w", err)
		}
		var old *string
		if !isNew {
			old = &existing.status
		}
		granted = true
		return recordAudit(ctx, tx, id, subj.UserID, nil, PlanPremium, old, "active", actor)
	})
	s.changed(ctx, subj, err)
	return granted, err
}

// EndCreatorPremium cancels userID's creator Premium. Any other row is left alone; ended is false
// when there was no live creator row.
func (s *Store) EndCreatorPremium(ctx context.Context, userID int64, actor string) (ended bool, err error) {
	subj := Subject{UserID: &userID}
	err = s.withTx(ctx, func(tx pgx.Tx) error {
		var id int64
		err := tx.QueryRow(ctx, `update entitlements set status = 'canceled', updated_at = now()
			where user_id = $1 and plan = 'premium' and source = 'creator' and status = 'active' returning id`, userID).Scan(&id)
		if err == pgx.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf("entitlements: end creator grant: %w", err)
		}
		ended = true
		old := "active"
		return recordAudit(ctx, tx, id, subj.UserID, nil, PlanPremium, &old, "canceled", actor)
	})
	s.changed(ctx, subj, err)
	return ended, err
}
