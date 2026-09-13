package subscribe

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct{ Pool *pgxpool.Pool }

func (s *Store) Upsert(ctx context.Context, email, normalized, token string) (bool, error) {
	tag, err := s.Pool.Exec(ctx,
		`insert into subscribers (email, email_normalized, token) values ($1, $2, $3)
		 on conflict (email_normalized) do nothing`, email, normalized, token)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() == 0, nil
}

func (s *Store) Confirm(ctx context.Context, token string) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `update subscribers set confirmed_at = coalesce(confirmed_at, now()), unsubscribed_at = null where token = $1`, token)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() == 1, nil
}

func (s *Store) Unsubscribe(ctx context.Context, token string) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `update subscribers set unsubscribed_at = now() where token = $1`, token)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() == 1, nil
}
