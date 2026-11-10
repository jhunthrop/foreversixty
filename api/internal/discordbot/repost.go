package discordbot

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

var _ guilds.RaidReposter = (*RaidRouter)(nil)

// oldPost is the message a repost replaces, and the reminder reply under it.
type oldPost struct {
	ChannelID string
	MessageID *string
	Reminder  *string
}

// resetPostSQL turns the raid's row into a fresh post in the install's channel (repointSQL's
// reset for one raid, whichever channel it was in) and returns the message it replaces. A row
// a drainer holds loses its lease, so that drainer's guarded writes no longer land.
const resetPostSQL = `
	with i as (
	  select channel_id from guild_discord_installs
	  where guild_id = $2 and channel_id is not null and not removed_in_discord
	    and channel_unusable_reason is null
	  for share),
	old as (
	  select channel_id, message_id, reminder_message_id from guild_raid_discord_posts
	  where raid_id = $1 and guild_id = $2
	  for update)
	update guild_raid_discord_posts p set channel_id = i.channel_id, message_id = null,
	  reminder_message_id = null, gone = false, dirty_seq = p.dirty_seq + 1, dirty_at = now(),
	  edited_at = null, wake_at = null, attempts = 0, next_attempt_at = now(), lease_until = null,
	  last_error = null, created_at = now()
	from i, old
	where p.raid_id = $1 and p.guild_id = $2
	returning old.channel_id, old.message_id, old.reminder_message_id`

// ResetPost makes the raid's post a fresh one; found is false when the raid has no row or the
// guild has no channel the bot can post in.
func (s *Store) ResetPost(ctx context.Context, raidID, guildID int64) (oldPost, bool, error) {
	var old oldPost
	err := s.Pool.QueryRow(ctx, resetPostSQL, raidID, guildID).Scan(&old.ChannelID, &old.MessageID, &old.Reminder)
	if errors.Is(err, pgx.ErrNoRows) {
		return oldPost{}, false, nil
	}
	if err != nil {
		return oldPost{}, false, fmt.Errorf("discordbot: reset post: %w", err)
	}
	return old, true, nil
}

// RepostRaid implements guilds.RaidReposter: with the bot active the raid's post goes out
// again as a new message (the drainer creates it) and the old one and its reminder reply are
// deleted; a raid with no row yet is queued. Without the bot the webhook posts it again.
func (r *RaidRouter) RepostRaid(ctx context.Context, guildID int64, e notify.RaidPosted) (bool, error) {
	if _, ok := r.activeChannel(ctx, guildID); ok {
		old, found, err := r.Bot.d.Store.ResetPost(ctx, e.RaidID, guildID)
		if err != nil {
			return false, err
		}
		if found {
			r.Bot.deleteOldPost(ctx, e.RaidID, old)
			return true, nil
		}
		if r.enqueue(ctx, guildID, e.RaidID) {
			return true, nil
		}
	}
	return r.Webhook.SendRaidPost(ctx, guildID, e)
}

// deleteOldPost deletes the replaced message and its reminder reply; a failure is logged,
// never returned: the fresh post is already queued.
func (s *Service) deleteOldPost(ctx context.Context, raidID int64, old oldPost) {
	for _, id := range []*string{old.Reminder, old.MessageID} {
		if id == nil {
			continue
		}
		if err := s.d.REST.DeleteMessage(ctx, old.ChannelID, *id); err != nil {
			s.log.Warn("discordbot", "op", "repost_delete", "raid_id", raidID, "err", describe(err))
		}
	}
}
