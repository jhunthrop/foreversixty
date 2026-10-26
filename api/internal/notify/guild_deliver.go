package notify

import (
	"context"
	"errors"
	"log/slog"
	"time"
)

// Raid event names, as the logs read them.
const (
	eventRaidPosted  = "raid_posted"
	eventRaidChanged = "raid_changed"
	eventReportReady = "report_ready"
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

// SetGuildPostDefault remembers the raid form's "Post to Discord" box on the guild's channel.
func (s *Service) SetGuildPostDefault(ctx context.Context, guildID int64, on bool) error {
	return s.GuildStore.SetPostDefault(ctx, guildID, on)
}

// RaidPosted posts the raid_posted embed; a failure is logged, never returned, because
// the raid itself was saved.
func (d *GuildDeliverer) RaidPosted(ctx context.Context, guildID int64, e RaidPosted) {
	if d == nil {
		return
	}
	d.deliverQuietly(ctx, guildID, eventRaidPosted, RaidPostedEmbed(e, d.clock()))
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

// deliver is Deliver that also says whether the post reached the channel.
func (d *GuildDeliverer) deliver(ctx context.Context, guildID int64, event string, embed Embed) (bool, error) {
	hook, found, err := d.Store.Get(ctx, guildID)
	if err != nil {
		return false, err
	}
	if !found || hook.DisabledReason != nil {
		return false, nil
	}
	box, err := d.Store.TokenBox(ctx, guildID)
	if errors.Is(err, ErrNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	url, err := d.Sealer.Open(box)
	if err != nil {
		return false, errors.New("notify: the guild's saved webhook could not be opened")
	}
	postCtx, cancel := context.WithTimeout(ctx, d.effectiveTimeout())
	defer cancel()
	postErr := d.Poster.Post(postCtx, url, embed)
	if postErr != nil && ctx.Err() != nil {
		// The request's shared budget ended the post; that says nothing about the channel.
		d.logger().Warn("notify", "guild", guildID, "event", event, "outcome", "raid_post_skipped")
		return false, nil
	}
	if err := d.record(ctx, guildID, event, postErr); err != nil {
		if postErr == nil {
			// The post landed; failing the caller over a bookkeeping error would have it retry
			// and post twice.
			d.logger().Error("notify", "guild", guildID, "event", event, "outcome", "record_failed", "err", err)
			return true, nil
		}
		return false, err
	}
	return postErr == nil, nil
}

// record applies the five-failure rule to the guild's channel on a context that outlives
// a cancelled request, so a post that happened is never left unrecorded.
func (d *GuildDeliverer) record(ctx context.Context, guildID int64, event string, postErr error) error {
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), posterTimeout)
	defer cancel()
	if postErr == nil {
		d.logger().Info("notify", "guild", guildID, "event", event, "outcome", "posted")
		return d.Store.RecordSuccess(ctx, guildID)
	}
	f := failureOf(postErr)
	disabled, err := d.Store.RecordFailure(ctx, guildID, f.text, f.counts)
	d.logger().Warn("notify", "guild", guildID, "event", event, "outcome", "failed",
		"status", f.status, "disabled", disabled)
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
