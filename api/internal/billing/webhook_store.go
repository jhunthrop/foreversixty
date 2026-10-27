// api/internal/billing/webhook_store.go
package billing

import (
	"context"
	"fmt"
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

// TransferOffered reports whether the guild has an offered, not yet
// accepted, transfer away from fromSubscriptionID.
func (s *Store) TransferOffered(ctx context.Context, guildID int64, fromSubscriptionID string) (bool, error) {
	var offered bool
	if err := s.Pool.QueryRow(ctx, `
		select exists (select 1 from plan_transfers
		where guild_id = $1 and from_subscription_id = $2 and state = 'offered')`,
		guildID, fromSubscriptionID).Scan(&offered); err != nil {
		return false, fmt.Errorf("billing: read transfer: %w", err)
	}
	return offered, nil
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

// MarkComebackUsed stamps a comeback offer as redeemed, once.
func (s *Store) MarkComebackUsed(ctx context.Context, id int64) error {
	if _, err := s.Pool.Exec(ctx,
		`update seat_comebacks set used_at = now() where id = $1 and used_at is null`, id); err != nil {
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
