package discordbot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

// ErrNoLivePost is a reminder for a raid without a post the bot can still edit.
var ErrNoLivePost = errors.New("discordbot: the raid has no live post")

// ErrLeaseLost is a write by a drainer whose lease on the row has lapsed or been taken: the
// row is someone else's now, so the write does not land.
var ErrLeaseLost = errors.New("discordbot: the post's lease was lost")

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
	// Lease is the claim's lease_until; every write of the claim is guarded on it.
	Lease time.Time
	// CreatedAt is when this generation of the post was queued (it names the nonce).
	CreatedAt time.Time
	// RemindDue asks the drain for the day-before reminder.
	RemindDue bool
	// ReminderMessageID is the reminder's reply, nil when none is up (or the post took it).
	ReminderMessageID *string
	// RemindedAt is when the reminder went out; nil before it and after a move withdrew it.
	RemindedAt *time.Time
}

// enqueueSQL queues the post in the install's channel, read in the same statement: for share
// waits on a channel save in flight and reads the channel it saved.
const enqueueSQL = `
	with i as (
	  select channel_id from guild_discord_installs
	  where guild_id = $2 and channel_id is not null and not removed_in_discord
	    and channel_unusable_reason is null
	  for share),
	ins as (
	  insert into guild_raid_discord_posts (raid_id, guild_id, channel_id)
	  select $1, $2, channel_id from i
	  on conflict (raid_id) do nothing)
	select exists (select 1 from i)`

// EnqueuePost records a raid's post in the guild's channel, due now (the row is born dirty).
// queued is false when the guild has no channel the bot can post in, so nothing was queued. A
// raid that already has a row, even a gone one, keeps it: one post per raid, never re-posted.
func (s *Store) EnqueuePost(ctx context.Context, raidID, guildID int64) (queued bool, err error) {
	if err := s.Pool.QueryRow(ctx, enqueueSQL, raidID, guildID).Scan(&queued); err != nil {
		return false, fmt.Errorf("discordbot: enqueue post: %w", err)
	}
	return queued, nil
}

// unparkSet, in a mark, gives a parked row another go: the raid changed, so the message did.
const unparkSet = `next_attempt_at = case when next_attempt_at = 'infinity' then now() else next_attempt_at end,
	attempts = case when next_attempt_at = 'infinity' then 0 else attempts end`

// MarkRaidDirty implements guilds.RaidPostMarker: the raid's post is out of date. A raid
// with no post row is a no-op. A failure is logged: the raid change itself was saved.
func (s *Store) MarkRaidDirty(ctx context.Context, raidID int64) {
	err := s.exec(ctx, `update guild_raid_discord_posts
		set dirty_seq = dirty_seq + 1, dirty_at = coalesce(dirty_at, now()), `+unparkSet+`
		where raid_id = $1 and not gone`, raidID)
	if err != nil {
		s.logger().Error("discordbot", "op", "mark_raid_dirty", "raid_id", raidID, "err", err)
	}
}

// MarkRaidMoved implements guilds.RaidPostMarker: the raid's start moved from was, ticked or
// not, so the post shows Moved and Was. A failure is logged.
func (s *Store) MarkRaidMoved(ctx context.Context, raidID int64, was time.Time) {
	err := s.exec(ctx, `update guild_raid_discord_posts
		set moved_from = $2, reminded_at = null, dirty_seq = dirty_seq + 1, dirty_at = coalesce(dirty_at, now()), `+unparkSet+`
		where raid_id = $1 and not gone`, raidID, was)
	if err != nil {
		s.logger().Error("discordbot", "op", "mark_raid_moved", "raid_id", raidID, "err", err)
	}
}

// SetMovedFrom records the start the raid was moved from, which the post shows as Was. A
// reminder sent for the old start no longer stands (Ruling 4): reminded_at is cleared, so the
// post loses its Tomorrow look and the drain withdraws a reminder reply still up.
func (s *Store) SetMovedFrom(ctx context.Context, raidID int64, was time.Time) error {
	return s.exec(ctx, `update guild_raid_discord_posts set moved_from = $2, reminded_at = null
		where raid_id = $1`, raidID, was)
}

// MarkReminderDue asks the drain for the raid's day-before reminder and marks the post dirty.
// ErrNoLivePost when the raid has no post row or its post is gone.
func (s *Store) MarkReminderDue(ctx context.Context, raidID int64) error {
	tag, err := s.Pool.Exec(ctx, `update guild_raid_discord_posts
		set remind_due = true, dirty_seq = dirty_seq + 1, dirty_at = coalesce(dirty_at, now()), `+unparkSet+`
		where raid_id = $1 and not gone`, raidID)
	if err != nil {
		return fmt.Errorf("discordbot: mark reminder due: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNoLivePost
	}
	return nil
}

// FinishReminder records a reminder sent (reminderID its reply, nil when the post took it) or
// dropped unsent. It keeps the lease for the post's own FinishEdit. A move that landed while
// the reminder was in flight leaves reminded_at null, so the next drain withdraws the reply.
func (s *Store) FinishReminder(ctx context.Context, p Post, reminderID *string, now time.Time) error {
	return s.execLeased(ctx, `update guild_raid_discord_posts set remind_due = false, reminder_message_id = $3,
		reminded_at = case when moved_from is not distinct from $5 then $4::timestamptz end
		where raid_id = $1 and lease_until = $2`, p.RaidID, p.Lease, reminderID, now, p.MovedFrom)
}

// ClearReminder records the reminder reply withdrawn; rearm also forgets that it was sent, so
// the post loses its Tomorrow look. It keeps the lease for the post's own FinishEdit.
func (s *Store) ClearReminder(ctx context.Context, p Post, rearm bool) error {
	return s.execLeased(ctx, `update guild_raid_discord_posts set reminder_message_id = null,
		reminded_at = case when $3 then null else reminded_at end
		where raid_id = $1 and lease_until = $2`, p.RaidID, p.Lease, rearm)
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
// not leased, past its retry time, its guild's bot is present with a usable channel, and it
// is either dirty and last edited at least the debounce ago, or its wake time has come. skip
// locked keeps two drainers off one row; the lease keeps the next claim off it until the
// claim's guarded write.
const claimDueSQL = `
	update guild_raid_discord_posts p set lease_until = $1::timestamptz + $3::interval
	where p.raid_id in (
	  select d.raid_id from guild_raid_discord_posts d
	  join guild_discord_installs i on i.guild_id = d.guild_id
	    and not i.removed_in_discord and i.channel_unusable_reason is null
	  where not d.gone
	    and (d.lease_until is null or d.lease_until < $1)
	    and d.next_attempt_at <= $1
	    and ((d.dirty_at is not null and (d.edited_at is null or d.edited_at <= $1::timestamptz - $2::interval))
	         or d.wake_at <= $1)
	  order by d.next_attempt_at
	  limit $4
	  for update of d skip locked)
	returning p.raid_id, p.guild_id, p.channel_id, p.message_id, p.moved_from, p.dirty_seq, p.attempts,
	  p.lease_until, p.created_at, p.remind_due, p.reminder_message_id, p.reminded_at`

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
		if err := rows.Scan(&p.RaidID, &p.GuildID, &p.ChannelID, &p.MessageID, &p.MovedFrom, &p.DirtySeq,
			&p.Attempts, &p.Lease, &p.CreatedAt, &p.RemindDue, &p.ReminderMessageID, &p.RemindedAt); err != nil {
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
// cleared only when no newer one landed since the claim (dirty_seq still the claim's).
func (s *Store) FinishEdit(ctx context.Context, p Post, messageID string, wakeAt *time.Time, now time.Time) error {
	return s.execLeased(ctx, `update guild_raid_discord_posts set message_id = $4, edited_at = $6, wake_at = $5,
		dirty_at = case when dirty_seq = $3 then null else dirty_at end,
		attempts = 0, next_attempt_at = $6, lease_until = null, last_error = null
		where raid_id = $1 and lease_until = $2`, p.RaidID, p.Lease, p.DirtySeq, messageID, wakeAt, now)
}

// FailEdit releases the lease and holds the row until retryAt, keeping the mark.
func (s *Store) FailEdit(ctx context.Context, p Post, retryAt time.Time, lastErr string) error {
	return s.execLeased(ctx, `update guild_raid_discord_posts set attempts = attempts + 1, next_attempt_at = $3,
		lease_until = null, last_error = $4 where raid_id = $1 and lease_until = $2`, p.RaidID, p.Lease, retryAt, lastErr)
}

// DeferPost releases the lease and holds the row until at without counting an attempt: the
// drainer never called Discord for it (a rate-limit hold).
func (s *Store) DeferPost(ctx context.Context, p Post, at time.Time) error {
	return s.execLeased(ctx, `update guild_raid_discord_posts set next_attempt_at = $3, lease_until = null
		where raid_id = $1 and lease_until = $2`, p.RaidID, p.Lease, at)
}

// ParkPost stops retrying a row Discord will never take as it is; its next mark, or a channel
// save, unparks it.
func (s *Store) ParkPost(ctx context.Context, p Post, lastErr string) error {
	return s.execLeased(ctx, `update guild_raid_discord_posts set attempts = attempts + 1,
		next_attempt_at = 'infinity', lease_until = null, last_error = $3
		where raid_id = $1 and lease_until = $2`, p.RaidID, p.Lease, lastErr)
}

// ReleaseClaim hands rows back unsent, as they were before the claim, for any drainer.
func (s *Store) ReleaseClaim(ctx context.Context, posts []Post) error {
	if len(posts) == 0 {
		return nil
	}
	ids := make([]int64, len(posts))
	for i, p := range posts {
		ids[i] = p.RaidID
	}
	// One claim leases every row to the same instant.
	return s.exec(ctx, `update guild_raid_discord_posts set lease_until = null
		where raid_id = any($1) and lease_until = $2`, ids, posts[0].Lease)
}

// MarkGone records that the post was deleted in Discord: it is never edited or re-posted.
// It is guarded on the claim's lease, so a row re-pointed at a new channel meanwhile stays.
func (s *Store) MarkGone(ctx context.Context, p Post) error {
	return s.execLeased(ctx, `update guild_raid_discord_posts set gone = true, dirty_at = null, wake_at = null,
		lease_until = null where raid_id = $1 and lease_until = $2`, p.RaidID, p.Lease)
}

// DeletePost drops a raid's row (the raid itself is gone).
func (s *Store) DeletePost(ctx context.Context, raidID int64) error {
	return s.exec(ctx, `delete from guild_raid_discord_posts where raid_id = $1`, raidID)
}

// execLeased is exec for a claim's write: ErrLeaseLost when the row is no longer leased to it.
func (s *Store) execLeased(ctx context.Context, query string, args ...any) error {
	tag, err := s.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("discordbot: store write: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrLeaseLost
	}
	return nil
}

func (s *Store) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}
