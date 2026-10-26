package notify

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// GuildStore is every guild_notification_webhooks and notify_connect_states read and write:
// a guild's one raid channel and the guild connect's OAuth states.
type GuildStore struct{ Pool *pgxpool.Pool }

// GuildWebhook is a guild's saved channel as the settings page and the raid posts see it.
// The token is deliberately absent; GuildStore.TokenBox is the one read that returns it.
type GuildWebhook struct {
	GuildID   int64
	ChannelID string
	// DiscordGuildID is the Discord server of the channel; nil when unknown.
	DiscordGuildID *string
	Failures       int
	// DisabledReason is set once five counted failures switched posting off.
	DisabledReason *string
	LastError      *string
	ConnectedBy    *int64
	// PostDefault is the raid form's remembered "Post to Discord" box.
	PostDefault bool
	UpdatedAt   time.Time
}

// GuildSaveInput is everything Save writes.
type GuildSaveInput struct {
	GuildID        int64
	ChannelID      string
	DiscordGuildID *string
	TokenBox       []byte
	ConnectedBy    int64
}

// Save upserts the guild's channel with a new token and resets the failure counter, last
// error and disabled reason (a reconnect is a re-enable). post_default is kept.
func (s *GuildStore) Save(ctx context.Context, in GuildSaveInput) error {
	_, err := s.Pool.Exec(ctx, `
		insert into guild_notification_webhooks (guild_id, token_box, channel_id, discord_guild_id, connected_by)
		values ($1, $2, $3, $4, $5)
		on conflict (guild_id) do update set
		  token_box = excluded.token_box, channel_id = excluded.channel_id,
		  discord_guild_id = excluded.discord_guild_id, connected_by = excluded.connected_by,
		  failures = 0, last_error = null, disabled_reason = null, updated_at = now()`,
		in.GuildID, in.TokenBox, in.ChannelID, in.DiscordGuildID, in.ConnectedBy)
	if err != nil {
		return fmt.Errorf("notify: save guild %d: %w", in.GuildID, err)
	}
	return nil
}

// Get returns the guild's channel without its token; ok is false when none is saved.
func (s *GuildStore) Get(ctx context.Context, guildID int64) (GuildWebhook, bool, error) {
	w := GuildWebhook{GuildID: guildID}
	err := s.Pool.QueryRow(ctx, `
		select channel_id, discord_guild_id, failures, disabled_reason, last_error, connected_by, post_default, updated_at
		from guild_notification_webhooks where guild_id = $1`, guildID).
		Scan(&w.ChannelID, &w.DiscordGuildID, &w.Failures, &w.DisabledReason, &w.LastError,
			&w.ConnectedBy, &w.PostDefault, &w.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return GuildWebhook{}, false, nil
	}
	if err != nil {
		return GuildWebhook{}, false, fmt.Errorf("notify: get guild %d: %w", guildID, err)
	}
	return w, true, nil
}

// TokenBox returns the sealed webhook URL; ErrNotFound when nothing is saved.
func (s *GuildStore) TokenBox(ctx context.Context, guildID int64) ([]byte, error) {
	var box []byte
	err := s.Pool.QueryRow(ctx, `select token_box from guild_notification_webhooks where guild_id = $1`, guildID).Scan(&box)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("notify: guild token box %d: %w", guildID, err)
	}
	return box, nil
}

// Delete removes the channel; deleting an absent one is not an error.
func (s *GuildStore) Delete(ctx context.Context, guildID int64) error {
	if _, err := s.Pool.Exec(ctx, `delete from guild_notification_webhooks where guild_id = $1`, guildID); err != nil {
		return fmt.Errorf("notify: delete guild %d: %w", guildID, err)
	}
	return nil
}

// RecordSuccess clears the failure counter, last error and disabled reason: a post that
// lands switches posting back on.
func (s *GuildStore) RecordSuccess(ctx context.Context, guildID int64) error {
	_, err := s.Pool.Exec(ctx, `update guild_notification_webhooks
		set failures = 0, last_error = null, disabled_reason = null where guild_id = $1`, guildID)
	if err != nil {
		return fmt.Errorf("notify: record guild success %d: %w", guildID, err)
	}
	return nil
}

// RecordFailure is the account's five-failure rule over the guild's channel (failureRule);
// a 429 or a manual test passes counts false.
func (s *GuildStore) RecordFailure(ctx context.Context, guildID int64, lastErr string, counts bool) (bool, error) {
	return guildFailures.record(ctx, s.Pool, lastErr, counts, guildID)
}

// ConnectState is a guild connect in flight.
type ConnectState struct {
	UserID, GuildID int64
	// Expired is true once the state's ten minutes have passed (by the database clock).
	Expired bool
}

// CreateConnectState stores a new state for userID connecting guildID, valid for ttl. It
// first sweeps the states that have already expired, so the table stays small.
func (s *GuildStore) CreateConnectState(ctx context.Context, state string, userID, guildID int64, ttl time.Duration) error {
	if _, err := s.Pool.Exec(ctx, `delete from notify_connect_states where expires_at < now()`); err != nil {
		return fmt.Errorf("notify: sweep connect states: %w", err)
	}
	_, err := s.Pool.Exec(ctx, `insert into notify_connect_states (state, user_id, guild_id, expires_at)
		values ($1, $2, $3, now() + make_interval(secs => $4))`, state, userID, guildID, ttl.Seconds())
	if err != nil {
		return fmt.Errorf("notify: create connect state: %w", err)
	}
	return nil
}

// TakeConnectState deletes and returns the state's row (each state is used once); ok is
// false when no row matches.
func (s *GuildStore) TakeConnectState(ctx context.Context, state string) (ConnectState, bool, error) {
	if state == "" {
		return ConnectState{}, false, nil
	}
	var cs ConnectState
	err := s.Pool.QueryRow(ctx, `delete from notify_connect_states where state = $1
		returning user_id, guild_id, expires_at <= now()`, state).Scan(&cs.UserID, &cs.GuildID, &cs.Expired)
	if errors.Is(err, pgx.ErrNoRows) {
		return ConnectState{}, false, nil
	}
	if err != nil {
		return ConnectState{}, false, fmt.Errorf("notify: take connect state: %w", err)
	}
	return cs, true, nil
}
