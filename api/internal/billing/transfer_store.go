// api/internal/billing/transfer_store.go
package billing

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// transferDeclined is the state a withdrawn or declined offer ends in.
const transferDeclined = "declined"

// planTransfer is a guild's plan_transfers row; a guild has at most one.
type planTransfer struct {
	fromUserID         int64
	toUserID           int64
	fromSubscriptionID string
	state              string
	offeredAt          time.Time
	decidedAt          *time.Time
}

// GuildTransfer reads the guild's transfer row, nil when it has none.
func (s *Store) GuildTransfer(ctx context.Context, guildID int64) (*planTransfer, error) {
	var t planTransfer
	err := s.Pool.QueryRow(ctx, `
		select from_user_id, to_user_id, from_subscription_id, state, offered_at, decided_at
		from plan_transfers where guild_id = $1`, guildID).
		Scan(&t.fromUserID, &t.toUserID, &t.fromSubscriptionID, &t.state, &t.offeredAt, &t.decidedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("billing: read transfer of guild %d: %w", guildID, err)
	}
	return &t, nil
}

// OfferTransfer writes the guild's one transfer row as a fresh offer away
// from fromSubscriptionID, replacing whatever row was there.
func (s *Store) OfferTransfer(ctx context.Context, guildID, fromUserID, toUserID int64, fromSubscriptionID string) error {
	if _, err := s.Pool.Exec(ctx, `
		insert into plan_transfers (guild_id, from_user_id, to_user_id, from_subscription_id, state)
		values ($1, $2, $3, $4, 'offered')
		on conflict (guild_id) do update set
		  from_user_id = excluded.from_user_id, to_user_id = excluded.to_user_id,
		  from_subscription_id = excluded.from_subscription_id, state = 'offered',
		  offered_at = now(), decided_at = null`,
		guildID, fromUserID, toUserID, fromSubscriptionID); err != nil {
		return fmt.Errorf("billing: offer transfer of guild %d: %w", guildID, err)
	}
	return nil
}

// WithdrawTransfer closes the guild's open offer; false when none is open.
func (s *Store) WithdrawTransfer(ctx context.Context, guildID int64) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `
		update plan_transfers set state = 'declined', decided_at = now()
		where guild_id = $1 and state = 'offered'`, guildID)
	if err != nil {
		return false, fmt.Errorf("billing: withdraw transfer of guild %d: %w", guildID, err)
	}
	return tag.RowsAffected() == 1, nil
}

// DeclineTransfer closes the guild's open offer to toUserID; false when
// there is no such offer.
func (s *Store) DeclineTransfer(ctx context.Context, guildID, toUserID int64) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `
		update plan_transfers set state = 'declined', decided_at = now()
		where guild_id = $1 and to_user_id = $2 and state = 'offered'`, guildID, toUserID)
	if err != nil {
		return false, fmt.Errorf("billing: decline transfer of guild %d: %w", guildID, err)
	}
	return tag.RowsAffected() == 1, nil
}

// VerifiedCount is the live count of distinct verified accounts in the
// guild, its leader included.
func (s *Store) VerifiedCount(ctx context.Context, guildID int64) (int, error) {
	var n int
	if err := s.Pool.QueryRow(ctx, `
		select count(distinct user_id) from guild_members
		where guild_id = $1 and verified_at is not null`, guildID).Scan(&n); err != nil {
		return 0, fmt.Errorf("billing: verified count of guild %d: %w", guildID, err)
	}
	return n, nil
}
