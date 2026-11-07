// api/internal/bnetimport/token.go
package bnetimport

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// TokenSealer seals the sign-in's access token at rest; *notify.Sealer satisfies it.
type TokenSealer interface {
	Seal(plain string) ([]byte, error)
	Open(box []byte) (string, error)
}

// tokenExpiryMargin is how long before its expiry a stored token stops being used, so a
// refresh never starts on a token about to lapse mid-run.
const tokenExpiryMargin = 5 * time.Minute

// tokenUsable reports whether a token expiring at expiresAt may still start a refresh at now.
func tokenUsable(expiresAt, now time.Time) bool {
	return expiresAt.After(now.Add(tokenExpiryMargin))
}

// KeepToken stores the sign-in's access token, sealed, until it expires (auth.TokenKeeper).
// Battle.net issues no refresh token, so this is all a later visit refresh can use to read
// the account's character list. Without a Sealer, or for a token with no usable expiry,
// nothing is stored. The token is never logged; a failed write is logged without it.
func (s *Service) KeepToken(ctx context.Context, userID int64, token string, expiresAt time.Time) {
	if s.Sealer == nil || token == "" || !tokenUsable(expiresAt, time.Now()) {
		return
	}
	box, err := s.Sealer.Seal(token)
	if err != nil {
		s.logger().Warn("bnetimport", "op", "keep_token", "user_id", userID, "err", err)
		return
	}
	if _, err := s.Pool.Exec(ctx,
		`update users set bnet_token_box = $2, bnet_token_expires_at = $3 where id = $1`,
		userID, box, expiresAt); err != nil {
		s.logger().Warn("bnetimport", "op", "keep_token", "user_id", userID, "err", err)
	}
}

// liveToken is the account's stored access token while it is usable, else "". A token past
// its margin, or one that no longer opens (a rotated key), is forgotten.
func (s *Service) liveToken(ctx context.Context, userID int64) (string, error) {
	if s.Sealer == nil {
		return "", nil
	}
	var box []byte
	var expiresAt *time.Time
	err := s.Pool.QueryRow(ctx,
		`select bnet_token_box, bnet_token_expires_at from users where id = $1`, userID).Scan(&box, &expiresAt)
	if errors.Is(err, pgx.ErrNoRows) || (err == nil && box == nil) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("bnetimport: read token: %w", err)
	}
	if expiresAt == nil || !tokenUsable(*expiresAt, time.Now()) {
		return "", s.forgetToken(ctx, userID)
	}
	token, err := s.Sealer.Open(box)
	if err != nil {
		s.logger().Warn("bnetimport", "op", "open_token", "user_id", userID, "err", err)
		return "", s.forgetToken(ctx, userID)
	}
	return token, nil
}

// forgetToken drops the account's stored token.
func (s *Service) forgetToken(ctx context.Context, userID int64) error {
	if _, err := s.Pool.Exec(ctx,
		`update users set bnet_token_box = null, bnet_token_expires_at = null where id = $1`, userID); err != nil {
		return fmt.Errorf("bnetimport: forget token: %w", err)
	}
	return nil
}
