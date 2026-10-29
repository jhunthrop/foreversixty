package notify

import (
	"context"
	"errors"
	"log/slog"
	"time"
)

// Raid event names, as the logs read them.
const (
	eventRaidPosted   = "raid_posted"
	eventRaidChanged  = "raid_changed"
	eventRaidReminder = "raid_reminder"
	eventReportReady  = "report_ready"
)

// GuildDeliverer posts the raid embeds to a guild's channel. A nil *GuildDeliverer posts
// nothing, so the raid handlers need no wiring check.
type GuildDeliverer struct {
	Store  *GuildStore
	Sealer *Sealer
	Poster Poster
	Log    *slog.Logger
	// now stamps the embed; nil means time.Now.
	now func() time.Time
	// postTimeout bounds one post; zero means posterTimeout.
	postTimeout time.Duration
}

// GuildDeliverer returns a GuildDeliverer sharing the service's wiring; nil when the
// deployment has no guild channel store.
func (s *Service) GuildDeliverer() *GuildDeliverer {
	if s == nil || s.GuildStore == nil {
		return nil
	}
	return &GuildDeliverer{Store: s.GuildStore, Sealer: s.Sealer, Poster: s.Poster, Log: s.Log}
}

// SetGuildDiscordDefaults remembers the raid form's "Post to Discord" and "Remind the day
// before" boxes on the guild's channel.
func (s *Service) SetGuildDiscordDefaults(ctx context.Context, guildID int64, post, remind bool) error {
	return s.GuildStore.SetDefaults(ctx, guildID, post, remind)
}

// RaidPosted posts the raid_posted embed; a failure is logged, never returned, because
// the raid itself was saved.
func (d *GuildDeliverer) RaidPosted(ctx context.Context, guildID int64, e RaidPosted) {
	if d == nil {
		return
	}
	d.deliverQuietly(ctx, guildID, eventRaidPosted, RaidPostedEmbed(e, d.clock()))
}

// SendRaidPost posts the raid_posted embed for the raid-posts job, which retries: a post the
// channel was found for but that did not land returns a *NotSentError. False with no error
// means no channel, or it is switched off. A failure is the job's to count (failureUntallied).
func (d *GuildDeliverer) SendRaidPost(ctx context.Context, guildID int64, e RaidPosted) (bool, error) {
	if d == nil {
		return false, nil
	}
	embed := RaidPostedEmbed(e, d.clock())
	return d.deliverWith(ctx, guildID, eventRaidPosted, failureUntallied, func(ctx context.Context, url string) error {
		return d.Poster.Post(ctx, url, embed)
	})
}

// NotSentError is a send the guild's channel was found for that did not land: Discord
// refused it, or its context ended first. It never holds the URL or token.
type NotSentError struct{ Err error }

func (e *NotSentError) Error() string {
	return "notify: the guild's channel did not take the post: " + e.Err.Error()
}

func (e *NotSentError) Unwrap() error { return e.Err }

// quietly is the save-time contract: a send that did not land is already logged and
// counted on the channel, so it is reported as not posted, never as an error.
func quietly(sent bool, err error) (bool, error) {
	var notSent *NotSentError
	if errors.As(err, &notSent) {
		return false, nil
	}
	return sent, err
}

// RaidChanged posts the raid_changed embed (moved or cancelled); see RaidPosted.
func (d *GuildDeliverer) RaidChanged(ctx context.Context, guildID int64, e RaidChanged) {
	if d == nil {
		return
	}
	d.deliverQuietly(ctx, guildID, eventRaidChanged, RaidChangedEmbed(e, d.clock()))
}

// ReportReady posts the report_ready embed and says whether it reached the channel: false
// with no error when the guild has no channel (or it is switched off), so the caller can
// release its one-post claim.
func (d *GuildDeliverer) ReportReady(ctx context.Context, guildID int64, e ReportReady) (bool, error) {
	if d == nil {
		return false, nil
	}
	return d.deliver(ctx, guildID, eventReportReady, ReportReadyEmbed(e, d.clock()))
}

func (d *GuildDeliverer) deliverQuietly(ctx context.Context, guildID int64, event string, embed Embed) {
	if err := d.Deliver(ctx, guildID, event, embed); err != nil {
		d.logger().Error("notify", "guild", guildID, "event", event, "outcome", "error", "err", err)
	}
}

// Deliver posts embed when the guild has a channel that is switched on. A guild with no
// channel, or one switched off, is a silent skip. The error never holds the URL or token.
func (d *GuildDeliverer) Deliver(ctx context.Context, guildID int64, event string, embed Embed) error {
	_, err := d.deliver(ctx, guildID, event, embed)
	return err
}

// RaidReminder posts the day-before reminder through the guild's webhook and returns the
// reference to withdraw it by. False with no error means no channel, or it is switched off;
// a reminder that did not land returns a *NotSentError, so the job retries it. A failure is
// the job's to count (failureUntallied).
func (d *GuildDeliverer) RaidReminder(ctx context.Context, guildID int64, e RaidPosted) (ReminderRef, bool, error) {
	if d == nil {
		return ReminderRef{}, false, nil
	}
	var ref ReminderRef
	sent, err := d.deliverWith(ctx, guildID, eventRaidReminder, failureUntallied, func(ctx context.Context, url string) error {
		id, err := d.Poster.PostContent(ctx, url, WebhookReminderContent(e), WebhookReminderEmbed(e, d.clock()))
		ref = ReminderRef{Via: ReminderViaWebhook, MessageID: id}
		return err
	})
	return ref, sent, err
}

// RaidReminderWithdrawn deletes a webhook reminder again (the raid moved, was cancelled or
// locked). A bot reference is the bot's to delete; a failure is logged, never returned.
func (d *GuildDeliverer) RaidReminderWithdrawn(ctx context.Context, guildID int64, ref ReminderRef) {
	if d == nil || ref.Via != ReminderViaWebhook {
		return
	}
	url, found, err := d.openURL(ctx, guildID)
	if err == nil && found {
		postCtx, cancel := context.WithTimeout(ctx, d.effectiveTimeout())
		defer cancel()
		err = d.Poster.DeleteMessage(postCtx, url, ref.MessageID)
	}
	if err != nil {
		d.logger().Warn("notify", "guild", guildID, "event", eventRaidReminder, "outcome", "withdraw_failed", "err", err)
	}
}

// openURL is the guild's webhook URL, false when it has no channel or the channel is off.
func (d *GuildDeliverer) openURL(ctx context.Context, guildID int64) (string, bool, error) {
	hook, found, err := d.Store.Get(ctx, guildID)
	if err != nil || !found || hook.DisabledReason != nil {
		return "", false, err
	}
	box, err := d.Store.TokenBox(ctx, guildID)
	if errors.Is(err, ErrNotFound) {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	url, err := d.Sealer.Open(box)
	if err != nil {
		return "", false, errors.New("notify: the guild's saved webhook could not be opened")
	}
	return url, true, nil
}

// deliver is Deliver that also says whether the post reached the channel; a post that did
// not land is false with no error (quietly).
func (d *GuildDeliverer) deliver(ctx context.Context, guildID int64, event string, embed Embed) (bool, error) {
	return quietly(d.deliverWith(ctx, guildID, event, failureTallied, func(ctx context.Context, url string) error {
		return d.Poster.Post(ctx, url, embed)
	}))
}

// failureTally says whether a send that did not land counts toward the channel's
// five-failure rule.
type failureTally bool

const (
	// failureTallied is the save path's: a refused post counts toward switching the channel off.
	failureTallied failureTally = true
	// failureUntallied is the raid-posts job's: the job caps its own retries on the schedule
	// row, so one refused raid retried never switches the guild's channel off. The failure
	// is still kept as the channel's last error.
	failureUntallied failureTally = false
)

// deliverWith sends through send once the channel is found, then applies the five-failure
// rule as tally says. A send that did not land returns a *NotSentError.
func (d *GuildDeliverer) deliverWith(ctx context.Context, guildID int64, event string, tally failureTally, send func(ctx context.Context, url string) error) (bool, error) {
	url, found, err := d.openURL(ctx, guildID)
	if err != nil || !found {
		return false, err
	}
	postCtx, cancel := context.WithTimeout(ctx, d.effectiveTimeout())
	defer cancel()
	postErr := send(postCtx, url)
	if postErr != nil && ctx.Err() != nil {
		// The request's shared budget ended the post; that says nothing about the channel.
		d.logger().Warn("notify", "guild", guildID, "event", event, "outcome", "raid_post_skipped")
		return false, &NotSentError{Err: postErr}
	}
	if err := d.record(ctx, guildID, event, tally, postErr); err != nil {
		if postErr == nil {
			// The post landed; failing the caller over a bookkeeping error would have it retry
			// and post twice.
			d.logger().Error("notify", "guild", guildID, "event", event, "outcome", "record_failed", "err", err)
			return true, nil
		}
		return false, err
	}
	if postErr != nil {
		return false, &NotSentError{Err: postErr}
	}
	return true, nil
}

// record applies the five-failure rule to the guild's channel on a context that outlives
// a cancelled request, so a post that happened is never left unrecorded; an untallied
// failure keeps its text without counting.
func (d *GuildDeliverer) record(ctx context.Context, guildID int64, event string, tally failureTally, postErr error) error {
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), posterTimeout)
	defer cancel()
	if postErr == nil {
		d.logger().Info("notify", "guild", guildID, "event", event, "outcome", "posted")
		return d.Store.RecordSuccess(ctx, guildID)
	}
	f := failureOf(postErr)
	counts := f.counts && bool(tally)
	disabled, err := d.Store.RecordFailure(ctx, guildID, f.text, counts)
	d.logger().Warn("notify", "guild", guildID, "event", event, "outcome", "failed",
		"status", f.status, "counted", counts, "disabled", disabled)
	return err
}

func (d *GuildDeliverer) clock() time.Time {
	if d.now != nil {
		return d.now()
	}
	return time.Now()
}

func (d *GuildDeliverer) effectiveTimeout() time.Duration {
	if d.postTimeout == 0 {
		return posterTimeout
	}
	return d.postTimeout
}

func (d *GuildDeliverer) logger() *slog.Logger {
	if d.Log != nil {
		return d.Log
	}
	return slog.Default()
}
