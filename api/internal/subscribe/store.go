package subscribe

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct{ Pool *pgxpool.Pool }

// UpsertResult describes the row Upsert acted on. Existing is false only for
// a brand-new row, in which case Token/Confirmed/Unsubscribed/
// ConfirmationSentAt are the row's zero values (the caller has no use for
// them: it already knows the token it just inserted, and a fresh row is
// never confirmed, unsubscribed, or previously sent-to).
type UpsertResult struct {
	Existing           bool
	Token              string
	Confirmed          bool
	Unsubscribed       bool
	ConfirmationSentAt *time.Time
}

// Upsert inserts a new subscriber row when normalized is not already present.
// When the address already exists, nothing is inserted; the existing row's
// token, confirmation state, unsubscribed state, and last confirmation-send
// time are returned so the caller can decide whether to resend the
// confirmation (never inventing a new token for a row that already has one,
// never re-mailing an already-confirmed-and-subscribed address, and never
// resending faster than its own cooldown/cap policy allows).
func (s *Store) Upsert(ctx context.Context, email, normalized, token, unsubscribeToken string) (UpsertResult, error) {
	var returnedToken string
	err := s.Pool.QueryRow(ctx,
		`insert into subscribers (email, email_normalized, token, unsubscribe_token) values ($1, $2, $3, $4)
		 on conflict (email_normalized) do nothing
		 returning token`, email, normalized, token, unsubscribeToken).Scan(&returnedToken)
	if err == nil {
		return UpsertResult{Existing: false}, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return UpsertResult{}, err
	}
	var res UpsertResult
	res.Existing = true
	err = s.Pool.QueryRow(ctx,
		`select token, confirmed_at is not null, unsubscribed_at is not null, confirmation_sent_at
		 from subscribers where email_normalized = $1`, normalized).
		Scan(&res.Token, &res.Confirmed, &res.Unsubscribed, &res.ConfirmationSentAt)
	if err != nil {
		return UpsertResult{}, err
	}
	return res, nil
}

// MarkConfirmationSent records that a confirmation email was just sent for
// normalized, so future sends can be rate-limited against it.
func (s *Store) MarkConfirmationSent(ctx context.Context, normalized string) error {
	_, err := s.Pool.Exec(ctx, `update subscribers set confirmation_sent_at = now() where email_normalized = $1`, normalized)
	return err
}

func (s *Store) Confirm(ctx context.Context, token string) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `update subscribers set confirmed_at = coalesce(confirmed_at, now()), unsubscribed_at = null where token = $1`, token)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() == 1, nil
}

func (s *Store) Unsubscribe(ctx context.Context, unsubscribeToken string) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `update subscribers set unsubscribed_at = now() where unsubscribe_token = $1`, unsubscribeToken)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() == 1, nil
}
