package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store reads the stored addon exports.
type Store struct {
	Pool *pgxpool.Pool
}

// Export is key's addon export; found is false when it has none.
func (s *Store) Export(ctx context.Context, key string) (Export, bool, error) {
	var e Export
	err := s.Pool.QueryRow(ctx,
		`select export, source, captured_at from addon_exports where character_key = $1`, key).
		Scan(&e.Code, &e.Source, &e.CapturedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Export{}, false, nil
	}
	if err != nil {
		return Export{}, false, fmt.Errorf("app: export of %s: %w", key, err)
	}
	return e, true, nil
}
