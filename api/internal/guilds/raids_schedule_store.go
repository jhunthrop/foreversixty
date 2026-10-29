package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	// postHorizon is how far ahead of the start a raid posts when saved earlier.
	postHorizon = 7 * 24 * time.Hour
	// reminderLead is how long before the start the reminder goes out.
	reminderLead = 24 * time.Hour
	// reminderFresh is the 12-hour rule: a post newer than this before the reminder is
	// already the reminder, so none is sent.
	reminderFresh = 12 * time.Hour
	// scheduleLease keeps two runners off one row.
	scheduleLease = 2 * time.Minute
	// scheduleBatch is how many rows one claim takes.
	scheduleBatch = 50
)

// postAtFor is when a raid starting at startsAt posts: now when it is within the horizon,
// otherwise the horizon before the start.
func postAtFor(startsAt, now time.Time) time.Time {
	if at := startsAt.Add(-postHorizon); at.After(now) {
		return at
	}
	return now
}

// RaidSchedule is one raid's Discord posting plan, whichever destination sends it.
type RaidSchedule struct {
	RaidID, GuildID int64
	PostAt          time.Time
	Remind          bool
	PostedAt        *time.Time
	RemindedAt      *time.Time
	Reminder        *notify.ReminderRef
}

type dueKind int

const (
	duePost dueKind = iota
	dueReminder
)

// DueSchedule is a claimed row and what is due on it.
type DueSchedule struct {
	RaidSchedule
	Kind     dueKind
	Attempts int
}

const scheduleColumns = `s.raid_id, s.guild_id, s.post_at, s.remind, s.posted_at, s.reminded_at,
	s.reminder_via, s.reminder_message_id`

type rowScanner interface{ Scan(dest ...any) error }

func scanSchedule(row rowScanner, extra ...any) (RaidSchedule, error) {
	var rs RaidSchedule
	var via, msg *string
	dest := append([]any{&rs.RaidID, &rs.GuildID, &rs.PostAt, &rs.Remind, &rs.PostedAt, &rs.RemindedAt, &via, &msg}, extra...)
	if err := row.Scan(dest...); err != nil {
		return RaidSchedule{}, err
	}
	if via != nil && msg != nil {
		rs.Reminder = &notify.ReminderRef{Via: *via, MessageID: *msg}
	}
	return rs, nil
}

func (s *Store) scheduleExec(ctx context.Context, what, query string, args ...any) error {
	if _, err := s.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("guilds: %s: %w", what, err)
	}
	return nil
}

// UpsertSchedule plans a raid's post. An unposted row takes the new post_at and remind; a
// posted row keeps post_at and posted_at and takes only remind.
func (s *Store) UpsertSchedule(ctx context.Context, raidID, guildID int64, postAt time.Time, remind bool) error {
	return s.scheduleExec(ctx, fmt.Sprintf("upsert schedule of raid %d", raidID), `
		insert into guild_raid_discord_schedule (raid_id, guild_id, post_at, remind)
		values ($1, $2, $3, $4)
		on conflict (raid_id) do update set
		  post_at = case when guild_raid_discord_schedule.posted_at is null then excluded.post_at else guild_raid_discord_schedule.post_at end,
		  remind = excluded.remind,
		  attempts = case when guild_raid_discord_schedule.posted_at is null then 0 else guild_raid_discord_schedule.attempts end,
		  next_attempt_at = case when guild_raid_discord_schedule.posted_at is null then now() else guild_raid_discord_schedule.next_attempt_at end`,
		raidID, guildID, postAt, remind)
}

// SetScheduleRemind turns the reminder on or off for a raid's row.
func (s *Store) SetScheduleRemind(ctx context.Context, raidID int64, remind bool) error {
	return s.scheduleExec(ctx, fmt.Sprintf("set remind of raid %d", raidID),
		`update guild_raid_discord_schedule set remind = $2 where raid_id = $1`, raidID, remind)
}

// DeleteUnpostedSchedule drops a row whose post has not gone out; a posted row stays.
func (s *Store) DeleteUnpostedSchedule(ctx context.Context, raidID int64) error {
	return s.scheduleExec(ctx, fmt.Sprintf("delete schedule of raid %d", raidID),
		`delete from guild_raid_discord_schedule where raid_id = $1 and posted_at is null`, raidID)
}

// MarkSchedulePosted records the post as sent and clears its lease and retry state.
func (s *Store) MarkSchedulePosted(ctx context.Context, raidID int64, at time.Time) error {
	return s.scheduleExec(ctx, fmt.Sprintf("mark raid %d posted", raidID), `
		update guild_raid_discord_schedule set posted_at = $2, lease_until = null, attempts = 0,
		  next_attempt_at = now(), last_error = null where raid_id = $1`, raidID, at)
}

// MarkScheduleReminded records the reminder as sent, with the reference to withdraw it by.
func (s *Store) MarkScheduleReminded(ctx context.Context, raidID int64, ref notify.ReminderRef, at time.Time) error {
	return s.scheduleExec(ctx, fmt.Sprintf("mark raid %d reminded", raidID), `
		update guild_raid_discord_schedule set reminded_at = $2, reminder_via = $3, reminder_message_id = $4,
		  lease_until = null, attempts = 0, next_attempt_at = now(), last_error = null where raid_id = $1`,
		raidID, at, ref.Via, ref.MessageID)
}

// ResetScheduleReminder clears the sent reminder and returns its old reference (nil when
// none was sent), so the caller can delete the message and the row is armed again.
func (s *Store) ResetScheduleReminder(ctx context.Context, raidID int64) (*notify.ReminderRef, error) {
	var via, msg *string
	err := s.Pool.QueryRow(ctx, `
		update guild_raid_discord_schedule n set reminded_at = null, reminder_via = null, reminder_message_id = null
		from guild_raid_discord_schedule o
		where n.raid_id = $1 and o.raid_id = n.raid_id
		returning o.reminder_via, o.reminder_message_id`, raidID).Scan(&via, &msg)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("guilds: reset reminder of raid %d: %w", raidID, err)
	}
	if via == nil || msg == nil {
		return nil, nil
	}
	return &notify.ReminderRef{Via: *via, MessageID: *msg}, nil
}

// SchedulesOf returns the rows of the given raids, keyed by raid.
func (s *Store) SchedulesOf(ctx context.Context, raidIDs []int64) (map[int64]RaidSchedule, error) {
	out := map[int64]RaidSchedule{}
	if len(raidIDs) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx, `select `+scheduleColumns+` from guild_raid_discord_schedule s where s.raid_id = any($1)`, raidIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: read schedules: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		rs, err := scanSchedule(rows)
		if err != nil {
			return nil, fmt.Errorf("guilds: scan schedule: %w", err)
		}
		out[rs.RaidID] = rs
	}
	return out, rows.Err()
}

// ClaimDueSchedules leases up to limit rows that are due at now: an unposted row whose
// post_at has come, or a posted row whose reminder is on, unsent, within 24 h of the start
// and posted at least 36 h before it. Only planned raids that have not started qualify.
func (s *Store) ClaimDueSchedules(ctx context.Context, now time.Time, limit int) ([]DueSchedule, error) {
	rows, err := s.Pool.Query(ctx, `
		update guild_raid_discord_schedule s set lease_until = $1::timestamptz + make_interval(secs => $3)
		where s.raid_id in (
		  select d.raid_id from guild_raid_discord_schedule d
		  join guild_raids r on r.id = d.raid_id
		  where (d.lease_until is null or d.lease_until <= $1)
		    and d.next_attempt_at <= $1
		    and r.status = 'planned' and r.starts_at > $1
		    and ((d.posted_at is null and d.post_at <= $1)
		      or (d.remind and d.posted_at is not null and d.reminded_at is null
		          and r.starts_at - make_interval(secs => $4) <= $1
		          and d.posted_at <= r.starts_at - make_interval(secs => $5)))
		  order by d.next_attempt_at
		  limit $2
		  for update of d skip locked)
		returning `+scheduleColumns+`, s.attempts`,
		now, limit, scheduleLease.Seconds(), reminderLead.Seconds(), (reminderLead + reminderFresh).Seconds())
	if err != nil {
		return nil, fmt.Errorf("guilds: claim due schedules: %w", err)
	}
	defer rows.Close()
	var out []DueSchedule
	for rows.Next() {
		var attempts int
		rs, err := scanSchedule(rows, &attempts)
		if err != nil {
			return nil, fmt.Errorf("guilds: scan due schedule: %w", err)
		}
		kind := dueReminder
		if rs.PostedAt == nil {
			kind = duePost
		}
		out = append(out, DueSchedule{RaidSchedule: rs, Kind: kind, Attempts: attempts})
	}
	return out, rows.Err()
}

// FailSchedule counts a failed send, keeps its message and holds the row until retryAt.
func (s *Store) FailSchedule(ctx context.Context, raidID int64, retryAt time.Time, lastErr string) error {
	return s.scheduleExec(ctx, fmt.Sprintf("fail schedule of raid %d", raidID), `
		update guild_raid_discord_schedule set attempts = attempts + 1, next_attempt_at = $2,
		  last_error = $3, lease_until = null where raid_id = $1`, raidID, retryAt, lastErr)
}

// ReleaseSchedule drops the lease without counting an attempt.
func (s *Store) ReleaseSchedule(ctx context.Context, raidID int64) error {
	return s.scheduleExec(ctx, fmt.Sprintf("release schedule of raid %d", raidID),
		`update guild_raid_discord_schedule set lease_until = null where raid_id = $1`, raidID)
}
