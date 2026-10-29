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
	// rowBudget bounds one row's facts read and Discord calls. It runs detached from the
	// drain's cancel, so a stop lets the row in flight land. A stop's worst case is this, the
	// row's two writes and the release: 5 + 2 × 1 + 2 = 9 s, the API's shutdown budget, and
	// the drain's landing context (drain_stop.go) caps all of it at StopWithin regardless. A
	// row reached later than its lease less this is released.
	rowBudget = 5 * time.Second
	// releaseBudget bounds handing unreached rows back.
	releaseBudget = 2 * time.Second
	// backoffBase and backoffCap shape the retry delay: 2 s doubling, at most 5 minutes.
	backoffBase = 2 * time.Second
	backoffCap  = 5 * time.Minute
	// writeBudget bounds one store write after Discord answered, on its own context, so a
	// slow Discord call never leaves the row unrecorded (waiting out its lease).
	writeBudget = time.Second
	// probeNeeds is the row budget a channel refusal's server read must have left; with less,
	// the row is retried instead.
	probeNeeds = 3 * time.Second
	// authHold is how long a 401 (a bad or rotated bot token) holds every post.
	authHold = 10 * time.Minute
	// shareSweepEvery is how often the drainer deletes Share tokens past their keep.
	shareSweepEvery = time.Hour
)

// PostFacts reads what a raid's post shows; *guilds.Service implements it.
type PostFacts interface {
	RaidPostFacts(ctx context.Context, raidID int64) (guilds.RaidPostFacts, bool, error)
}

// Drainer creates and edits the raid posts that are due. Any instance may run one: rows are
// leased, so two drainers never hold one row, and every write of a claim is guarded on its
// lease.
type Drainer struct {
	Store *Store
	REST  REST
	Facts PostFacts
	// Marks are the spec emoji the roster draws (spec 2026-10-28 §P.9.1); nil draws text lines.
	Marks                 MarkSource
	Now                   func() time.Time
	Tick, Debounce, Lease time.Duration
	// StopWithin is how long a stop lets the drain's last work run once its context ends;
	// zero means defaultStopWithin.
	StopWithin time.Duration
	Log        *slog.Logger

	// The rate-limit holds Discord answered, kept across ticks: a channel whose bucket is
	// spent, and the bot's global limit. 429s count against Discord's invalid-request limit.
	mu         sync.Mutex
	held       map[string]time.Time
	globalTill time.Time
	// authTill is the 401 hold's end; one error line is logged per hold.
	authTill time.Time
}

// MarkSource is where the drainer reads the spec emoji ids: the bot's EmojiRegistry.
type MarkSource interface {
	Marks(ctx context.Context) notify.Marks
}

// Run drains every Tick until ctx ends. The row in flight when ctx ends still lands; the
// rows claimed but not reached are handed back.
func (d *Drainer) Run(ctx context.Context) {
	ticker := time.NewTicker(orDefault(d.Tick, defaultDrainTick))
	defer ticker.Stop()
	var swept time.Time
	for {
		if err := d.DrainOnce(ctx); err != nil && ctx.Err() == nil {
			d.logger().Error("discordbot", "op", "drain", "err", err)
		}
		if now := d.now(); now.Sub(swept) >= shareSweepEvery {
			swept = now
			if err := d.Store.SweepShares(ctx, now); err != nil && ctx.Err() == nil {
				d.logger().Error("discordbot", "op", "sweep_shares", "err", err)
			}
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
	land, end := landing(ctx, orDefault(d.StopWithin, defaultStopWithin))
	defer end()
	for i, p := range posts {
		if ctx.Err() != nil || d.now().After(p.Lease.Add(-rowBudget)) {
			d.release(land, posts[i:])
			return nil
		}
		rowCtx, cancel := context.WithTimeout(land, rowBudget)
		d.drainRow(rowCtx, p, now)
		cancel()
	}
	return nil
}

// release hands rows back unsent: a stop, or a lease too near its end to finish inside.
func (d *Drainer) release(ctx context.Context, posts []Post) {
	ctx, cancel := detached(ctx, releaseBudget)
	defer cancel()
	if err := d.Store.ReleaseClaim(ctx, posts); err != nil {
		d.logger().Error("discordbot", "op", "post_release", "rows", len(posts), "err", err)
	}
}

// write runs one store write on its own short context, detached from the row's budget but
// capped by the drain's landing context, so a stop still ends it in time.
func (d *Drainer) write(ctx context.Context, op string, raidID int64, do func(context.Context) error) {
	ctx, cancel := detached(ctx, writeBudget)
	defer cancel()
	d.logIfErr(do(ctx), op, raidID)
}

func (d *Drainer) drainRow(ctx context.Context, p Post, now time.Time) {
	if at, held := d.holdFor(p.ChannelID, now); held {
		d.write(ctx, "post_defer", p.RaidID, func(ctx context.Context) error { return d.Store.DeferPost(ctx, p, at) })
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
	messageID, look, err := d.sendPost(ctx, p, facts, state, now)
	if errors.Is(err, ErrLeaseLost) {
		// Another drainer holds the row now; it redoes the work, so nothing is counted here.
		d.logIfErr(err, "post_finish", p.RaidID)
		return
	}
	if err != nil {
		d.refused(ctx, p, now, err)
		return
	}
	d.write(ctx, "post_finish", p.RaidID, func(ctx context.Context) error {
		return d.Store.FinishEdit(ctx, p, messageID, wakeAt(look, facts), now)
	})
	d.write(ctx, "post_success", p.RaidID, func(ctx context.Context) error { return d.Store.RecordPostSuccess(ctx, p.GuildID) })
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
	if d.Marks != nil {
		post.Marks = d.Marks.Marks(ctx)
	}
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

// wakeAt is the next time the post changes on its own: the start while moved or Tomorrow,
// else the end while the raid has not ended; nil once nothing is left to change by the clock.
func wakeAt(state notify.RaidPostState, facts guilds.RaidPostFacts) *time.Time {
	switch state {
	case notify.PostMoved, notify.PostTomorrow:
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
//   - a 429: wait retry_after, holding the channel (or everything, on the global limit),
//     without counting an attempt;
//   - a 401: the bot token is bad, so every post waits authHold (one error line per hold);
//   - any other 4xx: Discord will never take the message as it is, so the row is parked
//     until its next mark;
//   - anything else (no answer, a 5xx, an unexplained refusal): back off, at most 5 minutes,
//     and keep trying.
func (d *Drainer) refused(ctx context.Context, p Post, now time.Time, err error) {
	var api *APIError
	errors.As(err, &api)
	switch {
	case api != nil && api.Code == codeUnknownMessage:
		d.logger().Info("discordbot", "op", "post_gone", "raid_id", p.RaidID)
		d.write(ctx, "post_gone", p.RaidID, func(ctx context.Context) error { return d.Store.MarkGone(ctx, p) })
	case api != nil && channelRefused(api):
		d.channelRefused(ctx, p, now, err)
	case api != nil && api.RateLimited():
		retry := now.Add(backoff(0, api.RetryAfter))
		d.hold(p.ChannelID, retry, api.Global)
		d.logger().Warn("discordbot", "op", "post_rate_limited", "raid_id", p.RaidID,
			"retry_after", api.RetryAfter, "global", api.Global)
		d.write(ctx, "post_defer", p.RaidID, func(ctx context.Context) error { return d.Store.DeferPost(ctx, p, retry) })
	case api != nil && api.Status == http.StatusUnauthorized:
		until := d.holdAuth(now)
		d.write(ctx, "post_defer", p.RaidID, func(ctx context.Context) error { return d.Store.DeferPost(ctx, p, until) })
	case api != nil && nonRetryable(api):
		d.park(ctx, p, err)
	default:
		d.retry(ctx, p, now, err)
	}
}

// nonRetryable is a 4xx that no wait will change: not 401 (the token), 403 or 404 (judged
// as channel refusals or retried) or 429 (a wait).
func nonRetryable(api *APIError) bool {
	switch api.Status {
	case http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusTooManyRequests:
		return false
	}
	return api.Status >= 400 && api.Status < 500
}

// channelRefused judges a channel-level refusal; a decided one waits, uncounted, for the
// officer's next channel save (the claim skips the guild's rows until then). Without time
// left in the row budget for the server read, the row is retried instead.
func (d *Drainer) channelRefused(ctx context.Context, p Post, now time.Time, err error) {
	if deadline, ok := ctx.Deadline(); ok && time.Until(deadline) < probeNeeds {
		d.retry(ctx, p, now, err)
		return
	}
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
	d.write(ctx, "post_defer", p.RaidID, func(ctx context.Context) error { return d.Store.DeferPost(ctx, p, now) })
}

// retry backs the row off (at most 5 minutes) and counts the failure on the install. It never
// parks: a transient failure outlasts no outage.
func (d *Drainer) retry(ctx context.Context, p Post, now time.Time, err error) {
	d.logger().Warn("discordbot", "op", "post_failed", "raid_id", p.RaidID, "attempts", p.Attempts+1, "err", describe(err))
	d.write(ctx, "post_fail", p.RaidID, func(ctx context.Context) error {
		return d.Store.FailEdit(ctx, p, now.Add(backoff(p.Attempts, 0)), describe(err))
	})
	d.write(ctx, "post_failure", p.RaidID, func(ctx context.Context) error {
		return d.Store.RecordPostFailure(ctx, p.GuildID, describe(err))
	})
}

func (d *Drainer) park(ctx context.Context, p Post, err error) {
	d.logger().Error("discordbot", "op", "post_parked", "raid_id", p.RaidID, "err", describe(err))
	d.write(ctx, "post_park", p.RaidID, func(ctx context.Context) error { return d.Store.ParkPost(ctx, p, describe(err)) })
	d.write(ctx, "post_failure", p.RaidID, func(ctx context.Context) error {
		return d.Store.RecordPostFailure(ctx, p.GuildID, describe(err))
	})
}

// holdAuth starts, or keeps, the 401 hold and answers its end. The first 401 of a hold logs.
func (d *Drainer) holdAuth(now time.Time) time.Time {
	d.mu.Lock()
	defer d.mu.Unlock()
	if !d.authTill.After(now) {
		d.authTill = now.Add(authHold)
		d.logger().Error("discordbot", "op", "post_unauthorized", "effect", "every post waits", "hold", authHold)
	}
	if d.authTill.After(d.globalTill) {
		d.globalTill = d.authTill
	}
	return d.authTill
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
