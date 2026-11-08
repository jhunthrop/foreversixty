package apptesters

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store is the app_testers table.
type Store struct{ Pool *pgxpool.Pool }

// Add records email (already normalized) for platform, and reports whether
// the row is new: false when that address already asked for that platform.
func (s *Store) Add(ctx context.Context, email, platform string) (bool, error) {
	tag, err := s.Pool.Exec(ctx,
		`insert into app_testers (email, platform) values ($1, $2)
		 on conflict (email, platform) do nothing`, email, platform)
	if err != nil {
		return false, fmt.Errorf("apptesters: add: %w", err)
	}
	return tag.RowsAffected() == 1, nil
}

// List is every signup, newest first.
func (s *Store) List(ctx context.Context) ([]Tester, error) {
	rows, err := s.Pool.Query(ctx,
		`select email, platform, created_at from app_testers order by created_at desc, id desc`)
	if err != nil {
		return nil, fmt.Errorf("apptesters: list: %w", err)
	}
	defer rows.Close()
	out := []Tester{}
	for rows.Next() {
		var t Tester
		if err := rows.Scan(&t.Email, &t.Platform, &t.CreatedAt); err != nil {
			return nil, fmt.Errorf("apptesters: list: %w", err)
		}
		out = append(out, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("apptesters: list: %w", err)
	}
	return out, nil
}
