package discordbot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
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
	// rowBudget bounds one row's facts read, Discord calls and bookkeeping. It runs detached
	// from the drain's cancel, so a stop lets the row in flight land, and it stays inside the
	// API's 9 s shutdown budget. A row reached later than its lease less this is released.
	rowBudget = 7 * time.Second
	// releaseBudget bounds handing unreached rows back.
	releaseBudget = 2 * time.Second
	// backoffBase and backoffCap shape the retry delay: 2 s doubling, at most 5 minutes.
	backoffBase = 2 * time.Second
	backoffCap  = 5 * time.Minute
	// maxAttempts is when a failing row is parked instead of retried again.
	maxAttempts = 8
)

// PostFacts reads what a raid's post shows; *guilds.Service implements it.
type PostFacts interface {
	RaidPostFacts(ctx context.Context, raidID int64) (guilds.RaidPostFacts, bool, error)
}

// Drainer creates and edits the raid posts that are due. Any instance may run one: rows are
// leased, so two drainers never hold one row, and every write of a claim is guarded on its
// lease.
type Drainer struct {
	Store                 *Store
	REST                  REST
	Facts                 PostFacts
	Now                   func() time.Time
	Tick, Debounce, Lease time.Duration
	Log                   *slog.Logger

	// The rate-limit holds Discord answered, kept across ticks: a channel whose bucket is
	// spent, and the bot's global limit. 429s count against Discord's invalid-request limit.
	mu         sync.Mutex
	held       map[string]time.Time
	globalTill time.Time
}

// Run drains every Tick until ctx ends. The row in flight when ctx ends still lands; the
// rows claimed but not reached are handed back.
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
	if d.globalHeld(now) {
		return nil
	}
	posts, err := d.Store.ClaimDue(ctx, now, orDefault(d.Debounce, defaultDrainDebounce),
		orDefault(d.Lease, defaultDrainLease), drainBatch)
	if err != nil {
		return err
	}
	for i, p := range posts {
		if ctx.Err() != nil || d.now().After(p.Lease.Add(-rowBudget)) {
			d.release(ctx, posts[i:])
			return nil
		}
		rowCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), rowBudget)
		d.drainRow(rowCtx, p, now)
		cancel()
	}
	return nil
}

// release hands rows back unsent: a stop, or a lease too near its end to finish inside.
func (d *Drainer) release(ctx context.Context, posts []Post) {
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), releaseBudget)
	defer cancel()
	if err := d.Store.ReleaseClaim(ctx, posts); err != nil {
		d.logger().Error("discordbot", "op", "post_release", "rows", len(posts), "err", err)
	}
}

func (d *Drainer) drainRow(ctx context.Context, p Post, now time.Time) {
	if at, held := d.holdFor(p.ChannelID, now); held {
		d.logIfErr(d.Store.DeferPost(ctx, p, at), "post_defer", p.RaidID)
		return
	}
	facts, found, err := d.Facts.RaidPostFacts(ctx, p.RaidID)
	if err != nil {
		d.logger().Error("discordbot", "op", "post_facts", "raid_id", p.RaidID, "err", err)
		d.retry(ctx, p, now, err)
		return
	}
	if !found {
		d.logIfErr(d.Store.DeletePost(ctx, p.RaidID), "post_delete", p.RaidID)
		return
	}
	state := PostState(facts.Status, facts.Post.StartsAt, facts.EndsAt, p.MovedFrom, now)
	messageID, err := d.send(ctx, p, postFor(facts, state, p.MovedFrom), facts, now)
	if err != nil {
		d.refused(ctx, p, now, err)
		return
	}
	d.logIfErr(d.Store.FinishEdit(ctx, p, messageID, wakeAt(state, facts), now), "post_finish", p.RaidID)
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

// send creates the post the first time and edits it after; it answers the message id. The
// create carries the row's nonce, so a create whose answer was lost is not made twice.
func (d *Drainer) send(ctx context.Context, p Post, post notify.RaidPost, facts guilds.RaidPostFacts, now time.Time) (string, error) {
	msg := PostMessage(post, facts.RaidURL, facts.NightURL, now)
	if p.MessageID == nil {
		msg.Nonce = postNonce(p)
		return d.REST.CreateMessage(ctx, p.ChannelID, msg)
	}
	return *p.MessageID, d.REST.EditMessage(ctx, p.ChannelID, *p.MessageID, msg)
}

// postNonce names one generation of a raid's post: r{raid}.{queued at, ms}, base 36, so it
// stays inside Discord's 25 characters for any id.
func postNonce(p Post) string {
	return "r" + strconv.FormatInt(p.RaidID, 36) + "." + strconv.FormatInt(p.CreatedAt.UnixMilli(), 36)
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

// refused handles Discord's no:
//   - 10008, the message was deleted: gone for good (Ruling 7);
//   - a channel or server refusal: one read of the server decides removed_in_discord or the
//     unusable channel, and either pauses the guild's rows and sends the router to the webhook;
//   - a 429: wait retry_after, holding the channel (or everything, on the global limit);
//   - a 400 or 401: Discord will never take it as it is, so the row is parked;
//   - anything else: back off, parking after maxAttempts.
func (d *Drainer) refused(ctx context.Context, p Post, now time.Time, err error) {
	var api *APIError
	errors.As(err, &api)
	switch {
	case api != nil && api.Code == codeUnknownMessage:
		d.logger().Info("discordbot", "op", "post_gone", "raid_id", p.RaidID)
		d.logIfErr(d.Store.MarkGone(ctx, p), "post_gone", p.RaidID)
	case api != nil && channelRefused(api):
		d.channelRefused(ctx, p, now, err)
	case api != nil && api.RateLimited():
		retry := now.Add(backoff(p.Attempts, api.RetryAfter))
		d.hold(p.ChannelID, retry, api.Global)
		d.logger().Warn("discordbot", "op", "post_rate_limited", "raid_id", p.RaidID,
			"retry_after", api.RetryAfter, "global", api.Global)
		d.logIfErr(d.Store.FailEdit(ctx, p, retry, describe(err)), "post_fail", p.RaidID)
	case api != nil && (api.Status == http.StatusBadRequest || api.Status == http.StatusUnauthorized):
		d.park(ctx, p, err)
	default:
		d.retry(ctx, p, now, err)
	}
}

// channelRefused judges a channel-level refusal; a decided one waits, uncounted, for the
// officer's next channel save (the claim skips the guild's rows until then).
func (d *Drainer) channelRefused(ctx context.Context, p Post, now time.Time, err error) {
	verdict, judgeErr := judgeChannel(ctx, d.Store, d.REST, p.GuildID, p.ChannelID, err)
	if judgeErr != nil {
		d.logger().Error("discordbot", "op", "post_judge", "raid_id", p.RaidID, "err", judgeErr)
	}
	if verdict == verdictRetry {
		d.retry(ctx, p, now, err)
		return
	}
	d.logger().Warn("discordbot", "op", "post_refused", "raid_id", p.RaidID, "guild_id", p.GuildID,
		"removed", verdict == verdictRemoved, "err", describe(err))
	d.logIfErr(d.Store.DeferPost(ctx, p, now), "post_defer", p.RaidID)
}

// retry backs the row off and counts the failure on the install; the last attempt parks it.
func (d *Drainer) retry(ctx context.Context, p Post, now time.Time, err error) {
	if p.Attempts+1 >= maxAttempts {
		d.park(ctx, p, err)
		return
	}
	d.logger().Warn("discordbot", "op", "post_failed", "raid_id", p.RaidID, "attempts", p.Attempts+1, "err", describe(err))
	d.logIfErr(d.Store.FailEdit(ctx, p, now.Add(backoff(p.Attempts, 0)), describe(err)), "post_fail", p.RaidID)
	d.logIfErr(d.Store.RecordPostFailure(ctx, p.GuildID, describe(err)), "post_failure", p.RaidID)
}

func (d *Drainer) park(ctx context.Context, p Post, err error) {
	d.logger().Error("discordbot", "op", "post_parked", "raid_id", p.RaidID, "attempts", p.Attempts+1, "err", describe(err))
	d.logIfErr(d.Store.ParkPost(ctx, p, describe(err)), "post_park", p.RaidID)
	d.logIfErr(d.Store.RecordPostFailure(ctx, p.GuildID, describe(err)), "post_failure", p.RaidID)
}

func (d *Drainer) hold(channelID string, until time.Time, global bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if global {
		d.globalTill = until
	}
	if d.held == nil {
		d.held = map[string]time.Time{}
	}
	d.held[channelID] = until
}

// holdFor is the time a held channel (or the global hold) frees up; expired holds drop.
func (d *Drainer) holdFor(channelID string, now time.Time) (time.Time, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.globalTill.After(now) {
		return d.globalTill, true
	}
	at, ok := d.held[channelID]
	if ok && !at.After(now) {
		delete(d.held, channelID)
		return time.Time{}, false
	}
	return at, ok
}

func (d *Drainer) globalHeld(now time.Time) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.globalTill.After(now)
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

func (d *Drainer) logIfErr(err error, op string, raidID int64) {
	switch {
	case errors.Is(err, ErrLeaseLost):
		d.logger().Warn("discordbot", "op", op, "raid_id", raidID, "outcome", "lease_lost")
	case err != nil:
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
