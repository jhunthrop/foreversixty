// api/internal/billing/checkout_store.go
package billing

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// GuildTrialUsed reports whether guildID has already had its one trial.
func (s *Store) GuildTrialUsed(ctx context.Context, guildID int64) (bool, error) {
	var used bool
	if err := s.Pool.QueryRow(ctx,
		`select exists (select 1 from guild_trials where guild_id = $1)`, guildID).Scan(&used); err != nil {
		return false, fmt.Errorf("billing: guild trial %d: %w", guildID, err)
	}
	return used, nil
}

// Comeback reads the price a seat holder had, ok=false unless comeback id
// belongs to userID and is unexpired and unused at now.
func (s *Store) Comeback(ctx context.Context, id, userID int64, now time.Time) (string, bool, error) {
	var key string
	err := s.Pool.QueryRow(ctx,
		`select price_lookup_key from seat_comebacks
		 where id = $1 and user_id = $2 and used_at is null and expires_at > $3`,
		id, userID, now).Scan(&key)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("billing: comeback %d: %w", id, err)
	}
	return key, true, nil
}

// OpenTransferTo reads the subscription a guild's open transfer offer to
// toUserID would replace, ok=false when no such offer exists.
func (s *Store) OpenTransferTo(ctx context.Context, guildID, toUserID int64) (string, bool, error) {
	var from string
	err := s.Pool.QueryRow(ctx,
		`select from_subscription_id from plan_transfers
		 where guild_id = $1 and to_user_id = $2 and state = 'offered'`,
		guildID, toUserID).Scan(&from)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("billing: open transfer guild %d: %w", guildID, err)
	}
	return from, true, nil
}
