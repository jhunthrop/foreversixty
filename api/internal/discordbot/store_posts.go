package discordbot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

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

// MarkRaidMoved implements guilds.RaidPostMarker: the raid's start moved from was, ticked or
// not, so the post shows Moved and Was. A failure is logged.
func (s *Store) MarkRaidMoved(ctx context.Context, raidID int64, was time.Time) {
	err := s.exec(ctx, `update guild_raid_discord_posts
		set moved_from = $2, dirty_seq = dirty_seq + 1, dirty_at = coalesce(dirty_at, now())
		where raid_id = $1 and not gone`, raidID, was)
	if err != nil {
		s.logger().Error("discordbot", "op", "mark_raid_moved", "raid_id", raidID, "err", err)
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
	  p.lease_until, p.created_at`

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
			&p.Attempts, &p.Lease, &p.CreatedAt); err != nil {
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

// ParkPost stops retrying a row Discord will never take as it is; a channel save unparks it.
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
