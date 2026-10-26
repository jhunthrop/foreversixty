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

// SaveInput is everything Save writes. GuildID and Name are nil when unknown.
type SaveInput struct {
	UserID    int64
	Kind      string
	ChannelID string
	GuildID   *string
	Name      *string
	TokenBox  []byte
	Events    map[EventName]bool
	Source    Source
}

// Save upserts the member's channel with its guild id, source and name, a new
// token and events, and resets the failure counter, last error and disabled
// flag (a re-save is a re-enable). All three event keys are written explicitly.
func (s *Store) Save(ctx context.Context, in SaveInput) error {
	raw, err := json.Marshal(explicitEvents(in.Events))
	if err != nil {
		return fmt.Errorf("notify: encode events: %w", err)
	}
	_, err = s.Pool.Exec(ctx, `
		insert into notification_webhooks (user_id, kind, channel_id, guild_id, webhook_name, token_box, events, source)
		values ($1, $2, $3, $4, $5, $6, $7, $8)
		on conflict (user_id, kind) do update set
		  channel_id = excluded.channel_id, guild_id = excluded.guild_id, source = excluded.source, webhook_name = excluded.webhook_name, token_box = excluded.token_box, events = excluded.events,
		  failures = 0, last_error = '', disabled_at = null, updated_at = now()`,
		in.UserID, in.Kind, in.ChannelID, in.GuildID, in.Name, in.TokenBox, raw, string(in.Source))
	if err != nil {
		return fmt.Errorf("notify: save %d/%s: %w", in.UserID, in.Kind, err)
	}
	return nil
}

// SaveEvents changes only the events of a saved channel, leaving the token
// and the counters alone. ErrNotFound when nothing is saved.
func (s *Store) SaveEvents(ctx context.Context, userID int64, kind string, events map[EventName]bool) error {
	raw, err := json.Marshal(explicitEvents(events))
	if err != nil {
		return fmt.Errorf("notify: encode events: %w", err)
	}
	return s.mustAffect(ctx, `update notification_webhooks set events = $3, updated_at = now()
		where user_id = $1 and kind = $2`, userID, kind, raw)
}

// storedEvents are the keys every stored events object carries.
var storedEvents = []EventName{EventSimFinished, EventUploadParsed, EventUploadPrivate}

// explicitEvents returns events with every stored key present (false when absent).
func explicitEvents(events map[EventName]bool) map[EventName]bool {
	out := make(map[EventName]bool, len(storedEvents))
	for _, name := range storedEvents {
		out[name] = events[name]
	}
	return out
}

// Get returns the saved channel without its token; ok is false when none is saved.
func (s *Store) Get(ctx context.Context, userID int64, kind string) (Webhook, bool, error) {
	w := Webhook{UserID: userID, Kind: kind}
	var raw []byte
	err := s.Pool.QueryRow(ctx, `
		select channel_id, guild_id, source, webhook_name, events, failures, last_error, disabled_at, updated_at
		from notification_webhooks where user_id = $1 and kind = $2`, userID, kind).
		Scan(&w.ChannelID, &w.GuildID, &w.Source, &w.Name, &raw, &w.Failures, &w.LastError, &w.DisabledAt, &w.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Webhook{}, false, nil
	}
	if err != nil {
		return Webhook{}, false, fmt.Errorf("notify: get %d/%s: %w", userID, kind, err)
	}
	if err := json.Unmarshal(raw, &w.Events); err != nil {
		return Webhook{}, false, fmt.Errorf("notify: decode events %d/%s: %w", userID, kind, err)
	}
	if _, present := w.Events[EventUploadPrivate]; !present {
		// A row written by an older API instance: the legacy behaviour posted private logs.
		w.Events[EventUploadPrivate] = true
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
	return accountFailures.record(ctx, s.Pool, lastErr, counts, userID, kind)
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
