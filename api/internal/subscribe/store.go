package subscribe

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct{ Pool *pgxpool.Pool }

// Upsert inserts a new subscriber row when normalized is not already present.
// When the address already exists, nothing is inserted; the existing row's
// token and confirmation state are returned so the caller can decide whether
// to resend the confirmation (never inventing a new token for a row that
// already has one, and never re-mailing an already-confirmed address).
func (s *Store) Upsert(ctx context.Context, email, normalized, token, unsubscribeToken string) (bool, string, bool, error) {
	var returnedToken string
	err := s.Pool.QueryRow(ctx,
		`insert into subscribers (email, email_normalized, token, unsubscribe_token) values ($1, $2, $3, $4)
		 on conflict (email_normalized) do nothing
		 returning token`, email, normalized, token, unsubscribeToken).Scan(&returnedToken)
	if err == nil {
		return false, "", false, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return false, "", false, err
	}
	var existingToken string
	var confirmed bool
	err = s.Pool.QueryRow(ctx,
		`select token, confirmed_at is not null from subscribers where email_normalized = $1`, normalized).
		Scan(&existingToken, &confirmed)
	if err != nil {
		return false, "", false, err
	}
	return true, existingToken, confirmed, nil
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
