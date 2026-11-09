package rating

// A player may hide their own Forever Score (owner approval 2026-11-08; addon-v2 contract 4's
// export.score_visible, stored as characters.score_visible by migration 0092). A hidden score is
// hidden from everyone but the character's own account, on every surface that reads it here:
// the character page and its card, the applicant sheet, the nightly Data addon. Hidden reads
// exactly as no score, with no label of its own.

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// ScoreHidden reports whether playerKey's owner hid its Forever Score. A key with no
// characters row was never hidden.
func (s *Store) ScoreHidden(ctx context.Context, playerKey string) (bool, error) {
	var hidden bool
	err := s.Pool.QueryRow(ctx, `select not score_visible from characters where key = $1`, playerKey).Scan(&hidden)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("rating: score visibility %s: %w", playerKey, err)
	}
	return hidden, nil
}

// ScoreHiddenFrom reports whether playerKey's score is hidden from viewerID: its owner hid it
// and viewerID is not that owner (0, no one signed in, is never the owner).
func (s *Store) ScoreHiddenFrom(ctx context.Context, playerKey string, viewerID int64) (bool, error) {
	var hidden bool
	err := s.Pool.QueryRow(ctx,
		`select not score_visible and (user_id is null or user_id <> $2) from characters where key = $1`,
		playerKey, viewerID).Scan(&hidden)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("rating: score visibility %s: %w", playerKey, err)
	}
	return hidden, nil
}

// hiddenFromViewer is ScoreHiddenFrom for the request's own actor.
func (s *Store) hiddenFromViewer(ctx context.Context, playerKey string) (bool, error) {
	return s.ScoreHiddenFrom(ctx, playerKey, auth.ActorFrom(ctx).UserID)
}
