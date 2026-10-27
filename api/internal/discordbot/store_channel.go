package discordbot

import (
	"context"
	"fmt"
)

// Why a channel the bot is in the server for cannot take its posts.
const (
	UnusableDeleted     = "deleted"
	UnusablePermissions = "permissions"
)

// ChannelUnusable is the install's channel state when the bot cannot post there. Missing is
// the lacking subset of View Channel, Send Messages and Embed Links, in that order; empty
// when the channel was deleted.
type ChannelUnusable struct {
	Reason  string
	Missing []string
}

// channelSaveSQL saves the channel and clears every failure state, the unusable one too, so
// an officer who fixed the permissions just saves the same channel again.
const channelSaveSQL = `update guild_discord_installs set channel_id = $2, channel_name = $3,
	failures = 0, last_error = null, removed_in_discord = false,
	channel_unusable_reason = null, channel_unusable_missing = '{}', updated_at = now()
	where guild_id = $1`

// repointSQL moves the posts of raids that have not ended to the new channel as fresh posts:
// their old message ids belong to the old channel. moved_from stays: it belongs to the raid,
// not the message. A row a drainer holds loses its lease, so that drainer's guarded writes no
// longer land. "Not ended" is guilds.Raid.Signable's rule (planned or locked, ending after
// now) restated in SQL; TestRepointFollowsRaidSignable pins the two together.
const repointSQL = `update guild_raid_discord_posts p set channel_id = $2, message_id = null, gone = false,
	dirty_seq = p.dirty_seq + 1, dirty_at = now(), edited_at = null, wake_at = null,
	attempts = 0, next_attempt_at = now(), lease_until = null, last_error = null, created_at = now()
	from guild_raids r
	where r.id = p.raid_id and p.guild_id = $1 and p.channel_id <> $2
	  and r.status in ('planned', 'locked') and r.starts_at + make_interval(mins => r.duration_minutes) > now()`

// dropOldChannelSQL drops what is left in other channels: the ended, done and cancelled raids.
const dropOldChannelSQL = `delete from guild_raid_discord_posts where guild_id = $1 and channel_id <> $2`

// unparkSQL gives parked rows another go once an officer has saved a channel.
const unparkSQL = `update guild_raid_discord_posts set next_attempt_at = now(), attempts = 0
	where guild_id = $1 and next_attempt_at = 'infinity'`

// SetChannel saves the raid channel, clears the failure and unusable states, and in the same
// transaction re-points the open raids' posts at it and drops the rest of the old channel's.
func (s *Store) SetChannel(ctx context.Context, guildID int64, channelID, channelName string) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("discordbot: set channel begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	steps := []struct {
		sql  string
		args []any
	}{
		{channelSaveSQL, []any{guildID, channelID, channelName}},
		{repointSQL, []any{guildID, channelID}},
		{dropOldChannelSQL, []any{guildID, channelID}},
		{unparkSQL, []any{guildID}},
	}
	for _, step := range steps {
		if _, err := tx.Exec(ctx, step.sql, step.args...); err != nil {
			return fmt.Errorf("discordbot: set channel: %w", err)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("discordbot: set channel commit: %w", err)
	}
	return nil
}

// MarkChannelUnusable records that the bot cannot post in the guild's channel. The router
// then posts through the webhook and the drainer leaves the guild's rows until a save.
func (s *Store) MarkChannelUnusable(ctx context.Context, guildID int64, u ChannelUnusable, lastErr string) error {
	missing := u.Missing
	if missing == nil {
		missing = []string{}
	}
	return s.exec(ctx, `update guild_discord_installs set channel_unusable_reason = $2,
		channel_unusable_missing = $3, last_error = $4, updated_at = now() where guild_id = $1`,
		guildID, u.Reason, missing, lastErr)
}
