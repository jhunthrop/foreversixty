// api/internal/billing/webhook_store.go
package billing

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// RecordGuildTrial notes that a guild has used its one trial; a second
// record for the same guild is ignored.
func (s *Store) RecordGuildTrial(ctx context.Context, guildID int64, subscriptionID string) error {
	if _, err := s.Pool.Exec(ctx, `
		insert into guild_trials (guild_id, stripe_subscription_id) values ($1, $2)
		on conflict (guild_id) do nothing`, guildID, subscriptionID); err != nil {
		return fmt.Errorf("billing: record guild trial: %w", err)
	}
	return nil
}

// Transfer states a guild's plan_transfers row can be in, as far as the
// webhook cares; transferNone means no row away from that subscription.
const (
	transferNone     = ""
	transferOffered  = "offered"
	transferAccepted = "accepted"
)

// TransferState is the state of the guild's transfer away from
// fromSubscriptionID, or transferNone when there is no such row.
func (s *Store) TransferState(ctx context.Context, guildID int64, fromSubscriptionID string) (string, error) {
	var state string
	err := s.Pool.QueryRow(ctx, `
		select state from plan_transfers where guild_id = $1 and from_subscription_id = $2`,
		guildID, fromSubscriptionID).Scan(&state)
	if errors.Is(err, pgx.ErrNoRows) {
		return transferNone, nil
	}
	if err != nil {
		return transferNone, fmt.Errorf("billing: read transfer: %w", err)
	}
	return state, nil
}

// MarkTransferAccepted completes an offered transfer; accepting an
// accepted or declined transfer changes nothing.
func (s *Store) MarkTransferAccepted(ctx context.Context, guildID int64, fromSubscriptionID string) error {
	if _, err := s.Pool.Exec(ctx, `
		update plan_transfers set state = 'accepted', decided_at = now()
		where guild_id = $1 and from_subscription_id = $2 and state = 'offered'`,
		guildID, fromSubscriptionID); err != nil {
		return fmt.Errorf("billing: mark transfer accepted: %w", err)
	}
	return nil
}

// MarkComebackUsed stamps the user's own unexpired comeback offer as
// redeemed, once; any other id, user or state changes nothing.
func (s *Store) MarkComebackUsed(ctx context.Context, id, userID int64) error {
	if _, err := s.Pool.Exec(ctx, `
		update seat_comebacks set used_at = now()
		where id = $1 and user_id = $2 and used_at is null and expires_at > now()`,
		id, userID); err != nil {
		return fmt.Errorf("billing: mark comeback used: %w", err)
	}
	return nil
}

// OutgoingTransferSubscriptionIDs lists the subscriptions that accepted
// transfers are retiring; they stay live until period end and are not orphans.
func (s *Store) OutgoingTransferSubscriptionIDs(ctx context.Context) ([]string, error) {
	rows, err := s.Pool.Query(ctx,
		`select from_subscription_id from plan_transfers where state = 'accepted'`)
	if err != nil {
		return nil, fmt.Errorf("billing: list outgoing transfers: %w", err)
	}
	defer rows.Close()
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("billing: scan outgoing transfer: %w", err)
		}
		ids = append(ids, id)
	}
	return ids, rows.Err()
}
