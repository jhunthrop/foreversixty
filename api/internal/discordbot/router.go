package discordbot

import (
	"context"
	"errors"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// The bot's halves of the guilds ports, checked here so cmd/api cannot drift.
var (
	_ guilds.RaidEvents     = (*RaidRouter)(nil)
	_ guilds.RaidPostMarker = (*Store)(nil)
	_ PostFacts             = (*guilds.Service)(nil)
	_ GuildService          = (*guilds.Service)(nil)
)

// RaidRouter implements guilds.RaidEvents: a guild whose bot can post gets the one edited
// post per raid; any other guild gets the webhook, as before the bot. It adds no plan
// check: every caller already passed the officer route's gate (Ruling 10).
type RaidRouter struct {
	Bot     *Service
	Webhook *notify.GuildDeliverer
}

// RaidPosted queues the raid's post when the bot is active, else posts through the webhook.
func (r *RaidRouter) RaidPosted(ctx context.Context, guildID int64, e notify.RaidPosted) {
	if channel, ok := r.activeChannel(ctx, guildID); ok {
		err := r.Bot.d.Store.EnqueuePost(ctx, e.RaidID, guildID, channel)
		if err == nil {
			return
		}
		r.Bot.log.Error("discordbot", "op", "enqueue_post", "raid_id", e.RaidID, "err", err)
	}
	r.Webhook.RaidPosted(ctx, guildID, e)
}

// RaidChanged edits the raid's one post: a raid with a row is marked (and its old start
// kept when moved); a webhook-era raid gets its first bot post. Without an active bot the
// webhook posts the change, and a row the bot left behind is still brought up to date for
// its return. A gone row is never re-posted.
func (r *RaidRouter) RaidChanged(ctx context.Context, guildID int64, e notify.RaidChanged) {
	channel, active := r.activeChannel(ctx, guildID)
	exists, err := r.postExists(ctx, e.RaidID)
	switch {
	case err != nil:
	case exists:
		r.setMovedFrom(ctx, e)
		r.Bot.d.Store.MarkRaidDirty(ctx, e.RaidID)
		if active {
			return
		}
	case active:
		if err := r.Bot.d.Store.EnqueuePost(ctx, e.RaidID, guildID, channel); err != nil {
			r.Bot.log.Error("discordbot", "op", "enqueue_post", "raid_id", e.RaidID, "err", err)
			break
		}
		r.setMovedFrom(ctx, e)
		return
	}
	r.Webhook.RaidChanged(ctx, guildID, e)
}

// ReportReady posts the night's report to the bot's channel as its own message (Ruling 6),
// else through the webhook. posted is false, with no error, when Discord refused it.
func (r *RaidRouter) ReportReady(ctx context.Context, guildID int64, e notify.ReportReady) (bool, error) {
	channel, ok := r.activeChannel(ctx, guildID)
	if !ok {
		return r.Webhook.ReportReady(ctx, guildID, e)
	}
	msg := Message{Embeds: []notify.Embed{notify.ReportReadyEmbed(e, r.Bot.now())}}
	if _, err := r.Bot.d.REST.CreateMessage(ctx, channel, msg); err != nil {
		r.Bot.reportRefused(ctx, guildID, err)
		return false, nil
	}
	if err := r.Bot.d.Store.RecordPostSuccess(ctx, guildID); err != nil {
		// The report landed; failing the caller would have it post twice.
		r.Bot.log.Error("discordbot", "op", "report_record", "guild_id", guildID, "err", err)
	}
	return true, nil
}

// activeChannel is the bot's channel when it can post for the guild. A failed read is
// logged and answers false, so the webhook still posts.
func (r *RaidRouter) activeChannel(ctx context.Context, guildID int64) (string, bool) {
	if r.Bot == nil {
		return "", false
	}
	in, active, err := r.Bot.Active(ctx, guildID)
	if err != nil {
		r.Bot.log.Error("discordbot", "op", "router_install", "guild_id", guildID, "err", err)
		return "", false
	}
	if !active {
		return "", false
	}
	return *in.ChannelID, true
}

func (r *RaidRouter) postExists(ctx context.Context, raidID int64) (bool, error) {
	if r.Bot == nil {
		return false, errors.New("discordbot: no bot")
	}
	exists, err := r.Bot.d.Store.PostExists(ctx, raidID)
	if err != nil {
		r.Bot.log.Error("discordbot", "op", "post_exists", "raid_id", raidID, "err", err)
	}
	return exists, err
}

func (r *RaidRouter) setMovedFrom(ctx context.Context, e notify.RaidChanged) {
	if e.Was == nil {
		return
	}
	if err := r.Bot.d.Store.SetMovedFrom(ctx, e.RaidID, *e.Was); err != nil {
		r.Bot.log.Error("discordbot", "op", "set_moved_from", "raid_id", e.RaidID, "err", err)
	}
}

// reportRefused records a report Discord did not take: a bot that lost access marks the
// install removed; any other refusal counts a failure.
func (s *Service) reportRefused(ctx context.Context, guildID int64, err error) {
	s.log.Warn("discordbot", "op", "report_post", "guild_id", guildID, "err", describe(err))
	var api *APIError
	if errors.As(err, &api) && api.MissingAccess() {
		s.MarkRemoved(ctx, guildID, removedInDiscordNote)
		return
	}
	if recErr := s.d.Store.RecordPostFailure(ctx, guildID, describe(err)); recErr != nil {
		s.log.Error("discordbot", "op", "report_record", "guild_id", guildID, "err", recErr)
	}
}

// MarkRemoved records that Discord says the bot cannot reach the guild's channel: the
// router then posts through the webhook, and the guild's rows wait for a new channel.
func (s *Service) MarkRemoved(ctx context.Context, guildID int64, lastErr string) {
	if err := s.d.Store.MarkRemovedInDiscord(ctx, guildID, lastErr); err != nil {
		s.log.Error("discordbot", "op", "mark_removed", "guild_id", guildID, "err", err)
	}
}
