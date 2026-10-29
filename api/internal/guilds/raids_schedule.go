// api/internal/guilds/raids_schedule.go
//
// When a raid posts to Discord and reminds the day before (design/specs/
// 2026-10-28-discord-guided-cards.md §P.4): a save schedules the post (now within 7 days,
// else 7 days before the start), an edit follows a move, a cancel or lock withdraws a sent
// webhook reminder, and RunRaidSchedule (the raid-posts job) sends what has come due through
// RaidEvents, the bot's router or the webhook. A bot reminder is never tracked here: the
// drainer withdraws it when the raid's own move, cancel or lock marks the post.
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
	// scheduleRetryBase is the first wait after a failed send; it doubles up to the cap.
	scheduleRetryBase = 2 * time.Minute
	scheduleRetryCap  = time.Hour
	// scheduleMaxAttempts is how many failed sends a row takes before it is given up.
	scheduleMaxAttempts = 5
	// scheduleErrorMax bounds the failure text kept on the row.
	scheduleErrorMax = 500
)

// ScheduleRun counts what one RunRaidSchedule sent and what failed.
type ScheduleRun struct{ Posted, Reminded, Failed int }

type dueOutcome int

const (
	outcomeSkipped dueOutcome = iota
	outcomePosted
	outcomeReminded
	outcomeFailed
)

func raidIDsOf(views []RaidView) []int64 {
	ids := make([]int64, len(views))
	for i, v := range views {
		ids[i] = v.ID
	}
	return ids
}

// discordPostOf is raid id's discord_post: nil without a schedule row.
func discordPostOf(schedules map[int64]RaidSchedule, id int64) *RaidDiscordPost {
	rs, ok := schedules[id]
	if !ok {
		return nil
	}
	postAt := rs.PostAt
	return &RaidDiscordPost{PostedAt: rs.PostedAt, PostAt: &postAt, RemindedAt: rs.RemindedAt, Remind: rs.Remind}
}

func (s *Service) logSchedule(op string, raidID int64, err error) {
	s.logger().Error("guilds", "op", op, "raid_id", raidID, "err", err)
}

// withSchedules returns views with each discord_post read again, after a save scheduled them.
func (s *Service) withSchedules(ctx context.Context, views []RaidView) []RaidView {
	schedules, err := s.Store.SchedulesOf(ctx, raidIDsOf(views))
	if err != nil {
		s.logger().Error("guilds", "op", "raid_schedule_read", "err", err)
		return views
	}
	out := make([]RaidView, len(views))
	for i, v := range views {
		v.DiscordPost = discordPostOf(schedules, v.ID)
		out[i] = v
	}
	return out
}

// scheduleCreated plans each created raid's post and posts the ones due now, so a repeat
// posts only the raid within the week. A failure is logged: the raids were saved.
func (s *Service) scheduleCreated(ctx context.Context, guildID int64, views []RaidView, remind bool, now time.Time) {
	if s.RaidEvents == nil {
		return
	}
	ctx = context.WithoutCancel(ctx)
	var due []RaidView
	for _, v := range views {
		postAt := postAtFor(v.StartsAt, now)
		if err := s.Store.UpsertSchedule(ctx, v.ID, guildID, postAt, remind); err != nil {
			s.logSchedule("raid_schedule", v.ID, err)
			continue
		}
		if !postAt.After(now) {
			due = append(due, v)
		}
	}
	s.postDue(ctx, guildID, due, now)
}

// postDue posts the views whose rows it can lease, so the schedule run cannot post them too,
// and marks them posted. A row the budget left out is released for the run.
func (s *Service) postDue(ctx context.Context, guildID int64, views []RaidView, now time.Time) {
	leased := make([]RaidView, 0, len(views))
	leases := map[int64]time.Time{}
	for _, v := range views {
		lease, ok, err := s.Store.leaseUnposted(ctx, v.ID, now)
		if err != nil {
			s.logSchedule("raid_schedule_lease", v.ID, err)
			continue
		}
		if ok {
			leased, leases[v.ID] = append(leased, v), lease
		}
	}
	posted := map[int64]bool{}
	for _, id := range s.postCreated(ctx, guildID, leased) {
		posted[id] = true
	}
	for _, v := range leased {
		if posted[v.ID] {
			s.markPosted(ctx, v.ID, leases[v.ID], v.StartsAt, now)
		} else {
			s.releaseDue(ctx, v.ID, leases[v.ID])
		}
	}
}

// scheduleEdited follows an edit. A posted raid takes the reminder box, and a move withdraws
// and re-arms its reminder. An unposted raid (or one with no row) is planned again when the
// box is on or it already had a row, posting at once when due, and dropped when the box is
// off. wasPosted reports whether the raid had posted before this edit, which is when a
// ticked move still announces the change.
func (s *Service) scheduleEdited(ctx context.Context, guildID int64, before, after Raid, req raidRequest, now time.Time) (wasPosted bool) {
	ctx = context.WithoutCancel(ctx)
	rows, err := s.Store.SchedulesOf(ctx, []int64{after.ID})
	if err != nil {
		s.logSchedule("raid_schedule_read", after.ID, err)
		return false
	}
	row, has := rows[after.ID]
	if has && row.PostedAt != nil {
		s.editPosted(ctx, guildID, after.ID, req.RemindDayBefore, !after.StartsAt.Equal(before.StartsAt))
		return true
	}
	if req.PostToDiscord != nil && !*req.PostToDiscord {
		if err := s.Store.DeleteUnpostedSchedule(ctx, after.ID); err != nil {
			s.logSchedule("raid_schedule_delete", after.ID, err)
		}
		return false
	}
	if (req.PostToDiscord == nil && !has) || s.RaidEvents == nil {
		return false
	}
	remind := row.Remind
	if !has || req.RemindDayBefore != nil {
		remind = s.remindChoice(ctx, guildID, req.RemindDayBefore)
	}
	s.planUnposted(ctx, guildID, after, remind, now)
	return false
}

// editPosted applies an edit to a posted raid's row.
func (s *Service) editPosted(ctx context.Context, guildID, raidID int64, remind *bool, moved bool) {
	if remind != nil {
		if err := s.Store.SetScheduleRemind(ctx, raidID, *remind); err != nil {
			s.logSchedule("raid_schedule_remind", raidID, err)
		}
	}
	if moved {
		s.withdrawReminder(ctx, guildID, raidID)
	}
}

// planUnposted upserts an unposted raid's row and posts it at once when due and planned.
func (s *Service) planUnposted(ctx context.Context, guildID int64, raid Raid, remind bool, now time.Time) {
	postAt := postAtFor(raid.StartsAt, now)
	if err := s.Store.UpsertSchedule(ctx, raid.ID, guildID, postAt, remind); err != nil {
		s.logSchedule("raid_schedule", raid.ID, err)
		return
	}
	if postAt.After(now) || raid.Status != RaidPlanned {
		return
	}
	view, err := s.raidViewOf(ctx, guildID, noViewer, raid)
	if err != nil {
		s.logSchedule("raid_schedule_view", raid.ID, err)
		return
	}
	s.postDue(ctx, guildID, []RaidView{view}, now)
}

// remindChoice is the save's reminder box, else the guild's remembered default (false when
// it cannot be read, logged).
func (s *Service) remindChoice(ctx context.Context, guildID int64, box *bool) bool {
	if box != nil {
		return *box
	}
	d, err := s.raidDiscord(ctx, guildID)
	if err != nil {
		s.logger().Error("guilds", "op", "raid_remind_default", "guild_id", guildID, "err", err)
	}
	return d.RemindDefault
}

// withdrawOnStatus withdraws a sent reminder when a raid is cancelled or locked.
func (s *Service) withdrawOnStatus(ctx context.Context, guildID, raidID int64) {
	s.withdrawReminder(context.WithoutCancel(ctx), guildID, raidID)
}

// withdrawReminder clears the raid's sent reminder, arming it again, and deletes a webhook
// reminder's message. A bot reminder has no reference here; the drainer withdraws it.
func (s *Service) withdrawReminder(ctx context.Context, guildID, raidID int64) {
	ref, err := s.Store.ResetScheduleReminder(ctx, raidID)
	if err != nil {
		s.logSchedule("raid_reminder_reset", raidID, err)
		return
	}
	if ref == nil || ref.Via != notify.ReminderViaWebhook || s.RaidEvents == nil {
		return
	}
	s.RaidEvents.RaidReminderWithdrawn(ctx, guildID, *ref)
}

// RunRaidSchedule sends every post and reminder due at now, claiming rows in batches until
// none are left or ctx ends. It returns an error only when the claim fails or every claimed
// row failed.
func (s *Service) RunRaidSchedule(ctx context.Context, now time.Time) (ScheduleRun, error) {
	var run ScheduleRun
	if s.RaidEvents == nil {
		return run, nil
	}
	if n, err := s.Store.RecoverSentSchedules(ctx, now); err != nil {
		s.logger().Error("guilds", "op", "raid_schedule_recover", "err", err)
	} else if n > 0 {
		s.logger().Info("guilds", "op", "raid_schedule_recover", "marked", n)
	}
	// Each batch is claimed at the run's clock as it has advanced, so a long run never claims
	// with a now its own earlier leases have outlived.
	began := time.Now()
	claimed := 0
	for ctx.Err() == nil {
		at := now.Add(time.Since(began))
		due, err := s.Store.ClaimDueSchedules(ctx, at, scheduleBatch)
		if err != nil {
			return run, err
		}
		if len(due) == 0 {
			break
		}
		claimed += len(due)
		for _, d := range due {
			run = run.count(s.runDue(ctx, d, at))
		}
	}
	if claimed > 0 && run.Failed == claimed {
		return run, fmt.Errorf("guilds: raid schedule: all %d due sends failed", claimed)
	}
	return run, nil
}

func (r ScheduleRun) count(o dueOutcome) ScheduleRun {
	switch o {
	case outcomePosted:
		r.Posted++
	case outcomeReminded:
		r.Reminded++
	case outcomeFailed:
		r.Failed++
	}
	return r
}

// RaidPostSender is a RaidEvents that can also post for the raid-posts job and say whether
// the post landed; a send that did not land is an error, so the row is retried. A RaidEvents
// without it posts through RaidPosted, whose failures are only logged.
type RaidPostSender interface {
	SendRaidPost(ctx context.Context, guildID int64, e notify.RaidPosted) (sent bool, err error)
}

var _ RaidPostSender = (*notify.GuildDeliverer)(nil)

// runDue sends one claimed row's post or reminder: only while the raid is still planned and
// the row still holds this claim's lease, renewed just before the send.
func (s *Service) runDue(ctx context.Context, d DueSchedule, now time.Time) dueOutcome {
	facts, found, err := s.RaidPostFacts(ctx, d.RaidID)
	if err != nil {
		return s.failDue(ctx, d.RaidID, d.Attempts, d.LeaseUntil, now, err)
	}
	if !found || facts.Status != RaidPlanned {
		s.releaseDue(ctx, d.RaidID, d.LeaseUntil)
		return outcomeSkipped
	}
	lease, held, err := s.Store.RenewScheduleLease(ctx, d.RaidID, d.LeaseUntil)
	if err != nil {
		s.logSchedule("raid_schedule_renew", d.RaidID, err)
		return outcomeSkipped
	}
	if !held {
		return outcomeSkipped
	}
	if d.Kind == duePost {
		return s.sendPost(ctx, d, lease, facts.Post.RaidPosted, now)
	}
	return s.sendReminder(ctx, d, lease, facts.Post.RaidPosted, now)
}

// releaseDue drops the lease held; a lease another runner took since is left alone.
func (s *Service) releaseDue(ctx context.Context, raidID int64, held time.Time) {
	if _, err := s.Store.ReleaseSchedule(ctx, raidID, held); err != nil {
		s.logSchedule("raid_schedule_release", raidID, err)
	}
}

// errPostReachedNoChannel is a due post that found no channel to post to (none saved, or
// switched off): it is retried like a refusal, never marked posted.
var errPostReachedNoChannel = errors.New("guilds: the raid post reached no channel")

// sendPost posts a due raid and marks it posted; a post that did not land, or reached no
// channel, is retried.
func (s *Service) sendPost(ctx context.Context, d DueSchedule, lease time.Time, e notify.RaidPosted, now time.Time) dueOutcome {
	if sender, ok := s.RaidEvents.(RaidPostSender); ok {
		sent, err := sender.SendRaidPost(ctx, d.GuildID, e)
		if err == nil && !sent {
			err = errPostReachedNoChannel
		}
		if err != nil {
			return s.failDue(ctx, d.RaidID, d.Attempts, lease, now, err)
		}
	} else {
		s.RaidEvents.RaidPosted(ctx, d.GuildID, e)
	}
	s.markPosted(ctx, d.RaidID, lease, e.StartsAt, now)
	return outcomePosted
}

// Marking a sent post is retried this many times, the wait doubling from scheduleMarkWait.
const (
	scheduleMarkTries = 3
	scheduleMarkWait  = 100 * time.Millisecond
)

// markPosted records a sent post. A post that went out must never go out again, so a failed
// mark is retried a few times; if it still fails, the row's lease is held to the raid's start
// and flagged, so no runner re-posts it and the next run's recovery is the only retry.
func (s *Service) markPosted(ctx context.Context, raidID int64, lease, startsAt, now time.Time) {
	// The post is out: the bookkeeping outlives a run or request that ends now.
	ctx = context.WithoutCancel(ctx)
	err := retrySchedule(ctx, func() error { return s.Store.MarkSchedulePosted(ctx, raidID, now) })
	if err == nil {
		return
	}
	s.logger().Error("guilds", "op", "raid_schedule_posted", "raid_id", raidID, "err", err)
	held, holdErr := s.Store.HoldSentSchedule(ctx, raidID, lease, startsAt)
	if holdErr != nil || !held {
		s.logger().Error("guilds", "op", "raid_schedule_hold", "raid_id", raidID, "held", held, "err", holdErr)
	}
}

// retrySchedule runs write up to scheduleMarkTries times with a short doubling wait, and
// returns the last error.
func retrySchedule(ctx context.Context, write func() error) error {
	wait := scheduleMarkWait
	var err error
	for try := 0; try < scheduleMarkTries; try++ {
		if err = write(); err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return err
		case <-time.After(wait):
		}
		wait *= 2
	}
	return err
}

// sendReminder sends a due reminder and marks it reminded; a reminder that did not land is
// retried, and one that reached no channel is marked so it is not retried every run.
func (s *Service) sendReminder(ctx context.Context, d DueSchedule, lease time.Time, e notify.RaidPosted, now time.Time) dueOutcome {
	ref, sent, err := s.RaidEvents.RaidReminder(ctx, d.GuildID, e)
	if err != nil {
		return s.failDue(ctx, d.RaidID, d.Attempts, lease, now, err)
	}
	// Reached no channel: marked anyway, so it is not retried on every run.
	if err := s.Store.MarkScheduleReminded(ctx, d.RaidID, trackedRef(ref, sent), now); err != nil {
		s.logSchedule("raid_schedule_reminded", d.RaidID, err)
	}
	if !sent {
		return outcomeSkipped
	}
	return outcomeReminded
}

// trackedRef is the reference the row keeps: only a sent webhook reminder's, the one the
// schedule withdraws itself; zero otherwise.
func trackedRef(ref notify.ReminderRef, sent bool) notify.ReminderRef {
	if sent && ref.Via == notify.ReminderViaWebhook && ref.MessageID != "" {
		return ref
	}
	return notify.ReminderRef{}
}

// failDue holds a failed row back for scheduleBackoff and counts it. The last of
// scheduleMaxAttempts, or a refusal Discord will repeat (a non-429 4xx), gives the row up
// instead: it keeps the error and is not retried until an edit plans it again. A lease
// another runner took since is left alone.
func (s *Service) failDue(ctx context.Context, raidID int64, attempts int, held, now time.Time, cause error) dueOutcome {
	msg := cause.Error()
	if len(msg) > scheduleErrorMax {
		msg = msg[:scheduleErrorMax]
	}
	giveUp := attempts+1 >= scheduleMaxAttempts || notify.Refused(cause)
	s.logger().Warn("guilds", "op", "raid_schedule_send", "raid_id", raidID, "attempts", attempts+1,
		"gave_up", giveUp, "err", msg)
	var err error
	if giveUp {
		_, err = s.Store.GiveUpSchedule(ctx, raidID, held, msg)
	} else {
		_, err = s.Store.FailSchedule(ctx, raidID, held, now.Add(scheduleBackoff(attempts)), msg)
	}
	if err != nil {
		s.logSchedule("raid_schedule_fail", raidID, err)
	}
	return outcomeFailed
}

// scheduleBackoff is the wait after attempts earlier failures: 2 minutes, doubling, at most
// an hour.
func scheduleBackoff(attempts int) time.Duration {
	wait := scheduleRetryBase
	for i := 0; i < attempts && wait < scheduleRetryCap; i++ {
		wait *= 2
	}
	return min(wait, scheduleRetryCap)
}

// leaseUnposted takes raidID's lease for a post sent from a save and returns it; false when
// the row is posted already or another runner holds it.
func (s *Store) leaseUnposted(ctx context.Context, raidID int64, now time.Time) (time.Time, bool, error) {
	var lease time.Time
	err := s.Pool.QueryRow(ctx, `
		update guild_raid_discord_schedule set lease_until = $2::timestamptz + make_interval(secs => $3)
		where raid_id = $1 and posted_at is null and (lease_until is null or lease_until <= $2)
		returning lease_until`, raidID, now, scheduleLease.Seconds()).Scan(&lease)
	if errors.Is(err, pgx.ErrNoRows) {
		return time.Time{}, false, nil
	}
	if err != nil {
		return time.Time{}, false, fmt.Errorf("guilds: lease schedule of raid %d: %w", raidID, err)
	}
	return lease, true, nil
}
