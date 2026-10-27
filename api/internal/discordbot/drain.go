package discordbot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// Drainer defaults (product §8.9: at most one edit per raid every 2 s).
const (
	defaultDrainTick     = time.Second
	defaultDrainDebounce = 2 * time.Second
	defaultDrainLease    = 30 * time.Second
	// drainBatch is how many rows one DrainOnce claims.
	drainBatch = 10
	// rowBudget bounds one row's facts read, Discord call and bookkeeping. It runs detached
	// from the drain's cancel, so a stop lets the row in flight land.
	rowBudget = 15 * time.Second
	// backoffBase and backoffCap shape the retry delay: 2 s doubling, at most 5 minutes.
	backoffBase = 2 * time.Second
	backoffCap  = 5 * time.Minute
)

// PostFacts reads what a raid's post shows; *guilds.Service implements it.
type PostFacts interface {
	RaidPostFacts(ctx context.Context, raidID int64) (guilds.RaidPostFacts, bool, error)
}

// Drainer creates and edits the raid posts that are due. Any instance may run one: rows are
// leased, so two drainers never hold one row.
type Drainer struct {
	Store *Store
	REST  REST
	Facts PostFacts
	// OnRemoved is told when Discord says the bot may not reach the guild's channel.
	OnRemoved             func(ctx context.Context, guildID int64, lastErr string)
	Now                   func() time.Time
	Tick, Debounce, Lease time.Duration
	Log                   *slog.Logger
}

// Run drains every Tick until ctx ends. A row in flight when ctx ends still lands; rows
// claimed but not reached keep their lease until it lapses, then any drainer takes them.
func (d *Drainer) Run(ctx context.Context) {
	ticker := time.NewTicker(orDefault(d.Tick, defaultDrainTick))
	defer ticker.Stop()
	for {
		if err := d.DrainOnce(ctx); err != nil && ctx.Err() == nil {
			d.logger().Error("discordbot", "op", "drain", "err", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

// DrainOnce claims the due rows and brings each post up to date.
func (d *Drainer) DrainOnce(ctx context.Context) error {
	now := d.now()
	posts, err := d.Store.ClaimDue(ctx, now, orDefault(d.Debounce, defaultDrainDebounce),
		orDefault(d.Lease, defaultDrainLease), drainBatch)
	if err != nil {
		return err
	}
	b := batch{held: map[string]time.Time{}}
	for _, p := range posts {
		if ctx.Err() != nil {
			return nil
		}
		rowCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), rowBudget)
		d.drainRow(rowCtx, p, now, &b)
		cancel()
	}
	return nil
}

// batch is what one DrainOnce learned from Discord's rate limits: a channel whose bucket is
// spent, or the bot's global limit, holds the rest of the batch without another request
// (429s count against Discord's invalid-request limit).
type batch struct {
	held       map[string]time.Time
	globalTill *time.Time
}

func (b *batch) holdFor(channelID string) (time.Time, bool) {
	if b.globalTill != nil {
		return *b.globalTill, true
	}
	at, ok := b.held[channelID]
	return at, ok
}

func (d *Drainer) drainRow(ctx context.Context, p Post, now time.Time, b *batch) {
	if at, held := b.holdFor(p.ChannelID); held {
		d.fail(ctx, p, at, "rate limited")
		return
	}
	facts, found, err := d.Facts.RaidPostFacts(ctx, p.RaidID)
	if err != nil {
		d.logger().Error("discordbot", "op", "post_facts", "raid_id", p.RaidID, "err", err)
		d.fail(ctx, p, now.Add(backoff(p.Attempts, 0)), "facts unreadable")
		return
	}
	if !found {
		d.logIfErr(d.Store.DeletePost(ctx, p.RaidID), "post_delete", p.RaidID)
		return
	}
	state := PostState(facts.Status, facts.Post.StartsAt, facts.EndsAt, p.MovedFrom, now)
	messageID, err := d.send(ctx, p, postFor(facts, state, p.MovedFrom), facts, now)
	if err != nil {
		d.refused(ctx, p, now, err, b)
		return
	}
	d.logIfErr(d.Store.FinishEdit(ctx, p.RaidID, p.DirtySeq, messageID, wakeAt(state, facts), now), "post_finish", p.RaidID)
	d.logIfErr(d.Store.RecordPostSuccess(ctx, p.GuildID), "post_success", p.RaidID)
}

// postFor is the raid's post in state; a moved post shows where it was.
func postFor(facts guilds.RaidPostFacts, state notify.RaidPostState, movedFrom *time.Time) notify.RaidPost {
	post := facts.Post
	post.State = state
	if state == notify.PostMoved {
		post.Was = movedFrom
	}
	return post
}

// send creates the post the first time and edits it after; it answers the message id.
func (d *Drainer) send(ctx context.Context, p Post, post notify.RaidPost, facts guilds.RaidPostFacts, now time.Time) (string, error) {
	msg := PostMessage(post, facts.RaidURL, facts.NightURL, now)
	if p.MessageID == nil {
		return d.REST.CreateMessage(ctx, p.ChannelID, msg)
	}
	return *p.MessageID, d.REST.EditMessage(ctx, p.ChannelID, *p.MessageID, msg)
}

// wakeAt is the next time the post changes on its own: the start while moved, else the end
// while the raid has not ended; nil once nothing is left to change by the clock.
func wakeAt(state notify.RaidPostState, facts guilds.RaidPostFacts) *time.Time {
	switch state {
	case notify.PostMoved:
		at := facts.Post.StartsAt
		return &at
	case notify.PostOpen, notify.PostLocked:
		at := facts.EndsAt
		return &at
	}
	return nil
}

// refused handles Discord's no: a deleted message or channel is gone for good (Ruling 7); a
// bot without access marks the install removed, which pauses the guild's rows and sends the
// router to the webhook; anything else retries after a backoff.
func (d *Drainer) refused(ctx context.Context, p Post, now time.Time, err error, b *batch) {
	var api *APIError
	errors.As(err, &api)
	switch {
	case api != nil && postDeleted(api):
		d.logger().Info("discordbot", "op", "post_gone", "raid_id", p.RaidID, "code", api.Code)
		d.logIfErr(d.Store.MarkGone(ctx, p.RaidID), "post_gone", p.RaidID)
	case api != nil && api.MissingAccess():
		d.logger().Warn("discordbot", "op", "post_refused", "raid_id", p.RaidID, "guild_id", p.GuildID,
			"status", api.Status, "code", api.Code)
		if d.OnRemoved != nil {
			d.OnRemoved(ctx, p.GuildID, removedInDiscordNote)
		}
		d.fail(ctx, p, now.Add(backoff(p.Attempts, 0)), describe(err))
	case api != nil && api.RateLimited():
		retry := now.Add(backoff(p.Attempts, api.RetryAfter))
		if api.Global {
			b.globalTill = &retry
		}
		b.held[p.ChannelID] = retry
		d.logger().Warn("discordbot", "op", "post_rate_limited", "raid_id", p.RaidID,
			"retry_after", api.RetryAfter, "global", api.Global)
		d.fail(ctx, p, retry, describe(err))
	default:
		d.logger().Warn("discordbot", "op", "post_failed", "raid_id", p.RaidID, "err", describe(err))
		d.fail(ctx, p, now.Add(backoff(p.Attempts, 0)), describe(err))
		d.logIfErr(d.Store.RecordPostFailure(ctx, p.GuildID, describe(err)), "post_failure", p.RaidID)
	}
}

// postDeleted is a message or channel that no longer exists. An expired interaction token
// (also Gone) never comes from a channel message call, so it is not counted.
func postDeleted(api *APIError) bool {
	return api.Code == codeUnknownMessage || api.Code == codeUnknownChannel
}

// describe is an error as a row keeps it: status and code for a Discord answer, the
// transport's own words otherwise. Neither carries a credential.
func describe(err error) string {
	var api *APIError
	if errors.As(err, &api) {
		return fmt.Sprintf("discord status %d code %d", api.Status, api.Code)
	}
	return err.Error()
}

func (d *Drainer) fail(ctx context.Context, p Post, retryAt time.Time, lastErr string) {
	d.logIfErr(d.Store.FailEdit(ctx, p.RaidID, retryAt, lastErr), "post_fail", p.RaidID)
}

func (d *Drainer) logIfErr(err error, op string, raidID int64) {
	if err != nil {
		d.logger().Error("discordbot", "op", op, "raid_id", raidID, "err", err)
	}
}

// backoff is the wait before a retry: 2 s doubling per attempt, at most 5 minutes, and never
// shorter than Discord's retry_after.
func backoff(attempts int, retryAfter time.Duration) time.Duration {
	wait := backoffCap
	if attempts >= 0 && attempts < 8 {
		wait = min(backoffBase<<attempts, backoffCap)
	}
	return max(wait, retryAfter)
}

func (d *Drainer) now() time.Time {
	if d.Now != nil {
		return d.Now()
	}
	return time.Now()
}

func (d *Drainer) logger() *slog.Logger {
	if d.Log != nil {
		return d.Log
	}
	return slog.Default()
}

func orDefault(v, def time.Duration) time.Duration {
	if v > 0 {
		return v
	}
	return def
}
