package discordbot

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

// Post is one raid's channel post as the drainer claims it.
type Post struct {
	RaidID, GuildID int64
	ChannelID       string
	// MessageID is nil until the post is created.
	MessageID *string
	MovedFrom *time.Time
	// DirtySeq is the mark count at the claim; FinishEdit clears the mark only if no newer
	// one landed while the edit was in flight.
	DirtySeq int64
	Attempts int
}

// EnqueuePost records a raid's post, due now (the row is born dirty). A raid that already
// has a row, even a gone one, keeps it: one post per raid, never re-posted.
func (s *Store) EnqueuePost(ctx context.Context, raidID, guildID int64, channelID string) error {
	return s.exec(ctx, `insert into guild_raid_discord_posts (raid_id, guild_id, channel_id)
		values ($1, $2, $3) on conflict (raid_id) do nothing`, raidID, guildID, channelID)
}

// MarkRaidDirty implements guilds.RaidPostMarker: the raid's post is out of date. A raid
// with no post row is a no-op. A failure is logged: the raid change itself was saved.
func (s *Store) MarkRaidDirty(ctx context.Context, raidID int64) {
	err := s.exec(ctx, `update guild_raid_discord_posts
		set dirty_seq = dirty_seq + 1, dirty_at = coalesce(dirty_at, now())
		where raid_id = $1 and not gone`, raidID)
	if err != nil {
		s.logger().Error("discordbot", "op", "mark_raid_dirty", "raid_id", raidID, "err", err)
	}
}

// SetMovedFrom records the start the raid was moved from, which the post shows as Was.
func (s *Store) SetMovedFrom(ctx context.Context, raidID int64, was time.Time) error {
	return s.exec(ctx, `update guild_raid_discord_posts set moved_from = $2 where raid_id = $1`, raidID, was)
}

// PostExists reports whether the raid has a post row, gone or not.
func (s *Store) PostExists(ctx context.Context, raidID int64) (bool, error) {
	var exists bool
	err := s.Pool.QueryRow(ctx, `select exists (select 1 from guild_raid_discord_posts where raid_id = $1)`,
		raidID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("discordbot: post exists: %w", err)
	}
	return exists, nil
}

// claimDueSQL leases up to $4 due rows in one statement. A row is due when it is not gone,
// not leased, past its retry time, its guild's bot is present, and it is either dirty and
// last edited at least the debounce ago, or its wake time has come. skip locked keeps two
// drainers off one row; the lease keeps the next claim off it until FinishEdit or FailEdit.
const claimDueSQL = `
	update guild_raid_discord_posts p set lease_until = $1::timestamptz + $3::interval
	where p.raid_id in (
	  select d.raid_id from guild_raid_discord_posts d
	  join guild_discord_installs i on i.guild_id = d.guild_id and not i.removed_in_discord
	  where not d.gone
	    and (d.lease_until is null or d.lease_until < $1)
	    and d.next_attempt_at <= $1
	    and ((d.dirty_at is not null and (d.edited_at is null or d.edited_at <= $1::timestamptz - $2::interval))
	         or d.wake_at <= $1)
	  order by d.next_attempt_at
	  limit $4
	  for update of d skip locked)
	returning p.raid_id, p.guild_id, p.channel_id, p.message_id, p.moved_from, p.dirty_seq, p.attempts`

// ClaimDue leases the rows due at now for lease.
func (s *Store) ClaimDue(ctx context.Context, now time.Time, debounce, lease time.Duration, limit int) ([]Post, error) {
	rows, err := s.Pool.Query(ctx, claimDueSQL, now, debounce, lease, limit)
	if err != nil {
		return nil, fmt.Errorf("discordbot: claim posts: %w", err)
	}
	defer rows.Close()
	var out []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.RaidID, &p.GuildID, &p.ChannelID, &p.MessageID, &p.MovedFrom, &p.DirtySeq, &p.Attempts); err != nil {
			return nil, fmt.Errorf("discordbot: claim posts: %w", err)
		}
		out = append(out, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("discordbot: claim posts: %w", err)
	}
	return out, nil
}

// FinishEdit records a landed create or edit: the message id, the edit time, the next time
// the post changes on its own (wakeAt, nil for never), and the lease released. The mark is
// cleared only when no newer one landed since the claim (dirty_seq still seq).
func (s *Store) FinishEdit(ctx context.Context, raidID, seq int64, messageID string, wakeAt *time.Time, now time.Time) error {
	return s.exec(ctx, `update guild_raid_discord_posts set message_id = $3, edited_at = $5, wake_at = $4,
		dirty_at = case when dirty_seq = $2 then null else dirty_at end,
		attempts = 0, next_attempt_at = $5, lease_until = null, last_error = null
		where raid_id = $1`, raidID, seq, messageID, wakeAt, now)
}

// FailEdit releases the lease and holds the row until retryAt, keeping the mark.
func (s *Store) FailEdit(ctx context.Context, raidID int64, retryAt time.Time, lastErr string) error {
	return s.exec(ctx, `update guild_raid_discord_posts set attempts = attempts + 1, next_attempt_at = $2,
		lease_until = null, last_error = $3 where raid_id = $1`, raidID, retryAt, lastErr)
}

// MarkGone records that the post was deleted in Discord: it is never edited or re-posted.
func (s *Store) MarkGone(ctx context.Context, raidID int64) error {
	return s.exec(ctx, `update guild_raid_discord_posts set gone = true, dirty_at = null, wake_at = null,
		lease_until = null where raid_id = $1`, raidID)
}

// DeletePost drops a raid's row (the raid itself is gone).
func (s *Store) DeletePost(ctx context.Context, raidID int64) error {
	return s.exec(ctx, `delete from guild_raid_discord_posts where raid_id = $1`, raidID)
}

// DeletePostsForGuild drops every post row of the guild (its posts belong to a channel the
// bot no longer posts in).
func (s *Store) DeletePostsForGuild(ctx context.Context, guildID int64) error {
	return s.exec(ctx, `delete from guild_raid_discord_posts where guild_id = $1`, guildID)
}

func (s *Store) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}
