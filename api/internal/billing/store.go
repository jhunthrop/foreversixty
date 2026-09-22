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

// withAdvisoryLock takes a session-level Postgres advisory lock scoped to
// key for the duration of fn, so at most one caller is ever inside fn for
// that key at a time. Uses a dedicated connection acquired from the pool
// for the lock/unlock pair, not tied to any of fn's own transactions.
// Shared by WithEventLock and WithGuildLock (DRY) — their key spaces
// never collide because a raw Stripe event id (WithEventLock's key) is
// never spelled like WithGuildLock's "guild-checkout:<id>" prefix.
func (s *Store) withAdvisoryLock(ctx context.Context, key string, fn func(context.Context) error) error {
	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("billing: lock %s: acquire: %w", key, err)
	}
	defer conn.Release()
	if _, err := conn.Exec(ctx, `select pg_advisory_lock(hashtext($1))`, key); err != nil {
		return fmt.Errorf("billing: lock %s: lock: %w", key, err)
	}
	defer func() {
		_, _ = conn.Exec(context.WithoutCancel(ctx), `select pg_advisory_unlock(hashtext($1))`, key)
	}()
	return fn(ctx)
}

// WithEventLock takes a session-level Postgres advisory lock scoped to
// event id for the duration of fn, so at most one webhook delivery for
// that event id is ever inside fn at a time. This is what makes
// RecordEventOnce's existing-row check race-free against a second,
// overlapping delivery of the same event id (Stripe's own retry can
// outrun a slow first attempt that has not yet responded): the second
// caller blocks here until the first's fn has returned (success or
// failure) and the lock is released, then sees the now-correct
// processed_at state. handleEvent's writes go through entitlements.Store,
// which manages its own transactions independently of this lock.
func (s *Store) WithEventLock(ctx context.Context, id string, fn func(context.Context) error) error {
	return s.withAdvisoryLock(ctx, id, fn)
}

// WithGuildLock takes a session-level Postgres advisory lock scoped to
// guildID for the duration of fn, so at most one caller is ever inside fn
// for that guild at a time — the fix for the security review's
// double-billing finding (2026-09-21): without it, two officers checking
// out for the same guild concurrently could each read "no active plan"
// from checkGuildCheckout before either had created a Checkout Session,
// each pay, and the second webhook would silently overwrite the first
// subscription id in entitlements (UpsertStripe's duplicate guard,
// entitlements/store.go, is the second, independent layer against the
// same race). Held by handler.go's guildCheckout from the entitlement
// read through Checkout Session creation and the pending_checkouts write.
func (s *Store) WithGuildLock(ctx context.Context, guildID int64, fn func(context.Context) error) error {
	return s.withAdvisoryLock(ctx, fmt.Sprintf("guild-checkout:%d", guildID), fn)
}

// RecordPendingCheckout writes one row marking a guild-plan Checkout
// Session in flight (security review fix, 2026-09-21) — called inside
// WithGuildLock, after the Checkout Session is created, so a second
// caller for the same guild inside a later lock window sees it via
// PendingGuildCheckout and is refused exactly as an already-active plan
// would refuse it.
func (s *Store) RecordPendingCheckout(ctx context.Context, guildID, userID int64, stripeSessionID string, expiresAt time.Time) error {
	if _, err := s.Pool.Exec(ctx,
		`insert into pending_checkouts (guild_id, user_id, stripe_session_id, expires_at) values ($1, $2, $3, $4)`,
		guildID, userID, stripeSessionID, expiresAt); err != nil {
		return fmt.Errorf("billing: record pending checkout guild %d: %w", guildID, err)
	}
	return nil
}

// PendingGuildCheckout reports whether guildID has an unexpired pending
// Checkout Session in flight. Read inside checkGuildCheckout's own
// WithGuildLock window, so it only ever race-frees against a genuinely
// concurrent caller — never a stale row left by a session nobody
// completed, which expires_at (and the reconcile sweep below) bounds.
func (s *Store) PendingGuildCheckout(ctx context.Context, guildID int64) (bool, error) {
	var exists bool
	if err := s.Pool.QueryRow(ctx,
		`select exists (select 1 from pending_checkouts where guild_id = $1 and expires_at > now())`,
		guildID).Scan(&exists); err != nil {
		return false, fmt.Errorf("billing: pending checkout guild %d: %w", guildID, err)
	}
	return exists, nil
}

// DeletePendingCheckout removes a guild's pending-checkout row once its
// Checkout Session's checkout.session.completed webhook lands — a
// completed session (successful or a caught duplicate) no longer needs
// to block a later, genuinely new checkout for the same guild.
func (s *Store) DeletePendingCheckout(ctx context.Context, stripeSessionID string) error {
	if _, err := s.Pool.Exec(ctx, `delete from pending_checkouts where stripe_session_id = $1`, stripeSessionID); err != nil {
		return fmt.Errorf("billing: delete pending checkout %s: %w", stripeSessionID, err)
	}
	return nil
}

// SweepExpiredPendingCheckouts deletes every pending_checkouts row whose
// Checkout Session expired with no webhook ever landing for it (an
// abandoned checkout) — called from the nightly stripe-reconcile job
// (spec §2.8), the existing sweep this fix piggybacks on rather than
// standing up a new background job. Returns how many rows it removed.
func (s *Store) SweepExpiredPendingCheckouts(ctx context.Context) (int64, error) {
	tag, err := s.Pool.Exec(ctx, `delete from pending_checkouts where expires_at <= now()`)
	if err != nil {
		return 0, fmt.Errorf("billing: sweep pending checkouts: %w", err)
	}
	return tag.RowsAffected(), nil
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
