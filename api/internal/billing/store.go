// api/internal/billing/store.go
package billing

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store is stripe_customers and stripe_events reads/writes. It never
// touches the entitlements table — see gateway.go's package doc.
type Store struct{ Pool *pgxpool.Pool }

// CustomerID reads userID's Stripe Customer id, ok=false when none exists
// yet (spec §2.4: created lazily on first checkout).
func (s *Store) CustomerID(ctx context.Context, userID int64) (string, bool, error) {
	var id string
	err := s.Pool.QueryRow(ctx,
		`select stripe_customer_id from stripe_customers where user_id = $1`, userID).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("billing: read customer %d: %w", userID, err)
	}
	return id, true, nil
}

// SaveCustomerID records userID's Stripe Customer id, upserting so a
// retried write after a genuine failure never creates a duplicate row.
func (s *Store) SaveCustomerID(ctx context.Context, userID int64, customerID string) error {
	if _, err := s.Pool.Exec(ctx,
		`insert into stripe_customers (user_id, stripe_customer_id) values ($1, $2)
		 on conflict (user_id) do update set stripe_customer_id = excluded.stripe_customer_id`,
		userID, customerID); err != nil {
		return fmt.Errorf("billing: save customer %d: %w", userID, err)
	}
	return nil
}

// RecordEventOnce is the webhook's idempotency gate (spec §2.6). It
// returns proceed=true when handleEvent must (re)run: either this is
// the first time event id has ever been seen, or a prior attempt
// recorded the row but failed before MarkEventProcessed ever ran (a
// transient failure mid-processing — Stripe's own retry then correctly
// causes a real second attempt, not a silent drop). proceed=false means
// this event id was already fully processed — a genuine redelivery,
// correctly a no-op.
func (s *Store) RecordEventOnce(ctx context.Context, id, eventType string, payload []byte) (bool, error) {
	tag, err := s.Pool.Exec(ctx,
		`insert into stripe_events (id, type, payload) values ($1, $2, $3) on conflict (id) do nothing`,
		id, eventType, payload)
	if err != nil {
		return false, fmt.Errorf("billing: record event %s: %w", id, err)
	}
	if tag.RowsAffected() == 1 {
		return true, nil
	}
	var processedAt *time.Time
	if err := s.Pool.QueryRow(ctx,
		`select processed_at from stripe_events where id = $1`, id).Scan(&processedAt); err != nil {
		return false, fmt.Errorf("billing: record event %s: check processed: %w", id, err)
	}
	return processedAt == nil, nil
}

// MarkEventProcessed stamps processed_at once the event's entitlement
// write has committed.
func (s *Store) MarkEventProcessed(ctx context.Context, id string) error {
	if _, err := s.Pool.Exec(ctx, `update stripe_events set processed_at = now() where id = $1`, id); err != nil {
		return fmt.Errorf("billing: mark event %s processed: %w", id, err)
	}
	return nil
}

// GuildClaimed reports whether guildID has been claimed (spec §2.3
// precondition 1: an unclaimed guild has no accountable officer and
// cannot be sold the plan). Read directly against the guilds table —
// api/internal/guilds/ is not imported (Ruling A).
func (s *Store) GuildClaimed(ctx context.Context, guildID int64) (bool, error) {
	var claimed bool
	if err := s.Pool.QueryRow(ctx,
		`select claimed_by is not null from guilds where id = $1`, guildID).Scan(&claimed); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("billing: guild claimed %d: %w", guildID, err)
	}
	return claimed, nil
}
