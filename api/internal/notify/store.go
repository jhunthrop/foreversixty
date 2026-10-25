package notify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store is every notification_webhooks read and write.
type Store struct{ Pool *pgxpool.Pool }

// Save upserts the member's channel with a new token and events, and resets
// the failure counter, last error and disabled flag (a re-save is a re-enable).
func (s *Store) Save(ctx context.Context, userID int64, kind, channelID string, tokenBox []byte, events map[EventName]bool) error {
	raw, err := json.Marshal(events)
	if err != nil {
		return fmt.Errorf("notify: encode events: %w", err)
	}
	_, err = s.Pool.Exec(ctx, `
		insert into notification_webhooks (user_id, kind, channel_id, token_box, events)
		values ($1, $2, $3, $4, $5)
		on conflict (user_id, kind) do update set
		  channel_id = excluded.channel_id, token_box = excluded.token_box, events = excluded.events,
		  failures = 0, last_error = '', disabled_at = null, updated_at = now()`,
		userID, kind, channelID, tokenBox, raw)
	if err != nil {
		return fmt.Errorf("notify: save %d/%s: %w", userID, kind, err)
	}
	return nil
}

// SaveEvents changes only the events of a saved channel, leaving the token
// and the counters alone. ErrNotFound when nothing is saved.
func (s *Store) SaveEvents(ctx context.Context, userID int64, kind string, events map[EventName]bool) error {
	raw, err := json.Marshal(events)
	if err != nil {
		return fmt.Errorf("notify: encode events: %w", err)
	}
	return s.mustAffect(ctx, `update notification_webhooks set events = $3, updated_at = now()
		where user_id = $1 and kind = $2`, userID, kind, raw)
}

// Get returns the saved channel without its token; ok is false when none is saved.
func (s *Store) Get(ctx context.Context, userID int64, kind string) (Webhook, bool, error) {
	w := Webhook{UserID: userID, Kind: kind}
	var raw []byte
	err := s.Pool.QueryRow(ctx, `
		select channel_id, events, failures, last_error, disabled_at, updated_at
		from notification_webhooks where user_id = $1 and kind = $2`, userID, kind).
		Scan(&w.ChannelID, &raw, &w.Failures, &w.LastError, &w.DisabledAt, &w.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Webhook{}, false, nil
	}
	if err != nil {
		return Webhook{}, false, fmt.Errorf("notify: get %d/%s: %w", userID, kind, err)
	}
	if err := json.Unmarshal(raw, &w.Events); err != nil {
		return Webhook{}, false, fmt.Errorf("notify: decode events %d/%s: %w", userID, kind, err)
	}
	return w, true, nil
}

// TokenBox returns the sealed token. It is the delivery job's one read of
// the credential; ErrNotFound when nothing is saved.
func (s *Store) TokenBox(ctx context.Context, userID int64, kind string) ([]byte, error) {
	var box []byte
	err := s.Pool.QueryRow(ctx, `select token_box from notification_webhooks
		where user_id = $1 and kind = $2`, userID, kind).Scan(&box)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("notify: token box %d/%s: %w", userID, kind, err)
	}
	return box, nil
}

// Delete removes the channel; deleting an absent one is not an error.
func (s *Store) Delete(ctx context.Context, userID int64, kind string) error {
	if _, err := s.Pool.Exec(ctx, `delete from notification_webhooks
		where user_id = $1 and kind = $2`, userID, kind); err != nil {
		return fmt.Errorf("notify: delete %d/%s: %w", userID, kind, err)
	}
	return nil
}

// RecordSuccess clears the failure counter, last error and disabled flag.
func (s *Store) RecordSuccess(ctx context.Context, userID int64, kind string) error {
	_, err := s.Pool.Exec(ctx, `update notification_webhooks
		set failures = 0, last_error = '', disabled_at = null
		where user_id = $1 and kind = $2`, userID, kind)
	if err != nil {
		return fmt.Errorf("notify: record success %d/%s: %w", userID, kind, err)
	}
	return nil
}

// RecordFailure stores lastErr. Only when counts is true does it also add
// to the consecutive-failure counter, and at MaxFailures set disabled_at
// (reported as disabled). A 429 or a manual test passes counts false.
func (s *Store) RecordFailure(ctx context.Context, userID int64, kind, lastErr string, counts bool) (bool, error) {
	step := 0
	if counts {
		step = 1
	}
	var disabled bool
	err := s.Pool.QueryRow(ctx, `
		update notification_webhooks set
		  failures = failures + $3,
		  last_error = $4,
		  disabled_at = case when failures + $3 >= $5 then coalesce(disabled_at, now()) else disabled_at end
		where user_id = $1 and kind = $2
		returning disabled_at is not null`,
		userID, kind, step, lastErr, MaxFailures).Scan(&disabled)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, ErrNotFound
	}
	if err != nil {
		return false, fmt.Errorf("notify: record failure %d/%s: %w", userID, kind, err)
	}
	return disabled, nil
}

func (s *Store) mustAffect(ctx context.Context, sql string, args ...any) error {
	tag, err := s.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("notify: update: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
