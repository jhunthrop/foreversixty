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
	_ guilds.RaidPostSender = (*RaidRouter)(nil)
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
	if _, ok := r.activeChannel(ctx, guildID); ok && r.enqueue(ctx, guildID, e.RaidID) {
		return
	}
	r.Webhook.RaidPosted(ctx, guildID, e)
}

// SendRaidPost is RaidPosted for the raid-posts job: a post queued for the bot is sent (the
// drainer owns it from there); otherwise the webhook posts, and a refusal is returned so the
// schedule retries it with backoff.
func (r *RaidRouter) SendRaidPost(ctx context.Context, guildID int64, e notify.RaidPosted) (bool, error) {
	if _, ok := r.activeChannel(ctx, guildID); ok && r.enqueue(ctx, guildID, e.RaidID) {
		return true, nil
	}
	return r.Webhook.SendRaidPost(ctx, guildID, e)
}

// RaidChanged edits the raid's one post: a raid with a row is marked (and its old start
// kept when moved); a webhook-era raid gets its first bot post. Without an active bot the
// webhook posts the change, and a row the bot left behind is still brought up to date for
// its return. A gone row is never re-posted. With the bot active the webhook never posts a
// raid change, even when the row could not be read: the raid is marked (a no-op without a
// row) and the failure logged.
func (r *RaidRouter) RaidChanged(ctx context.Context, guildID int64, e notify.RaidChanged) {
	_, active := r.activeChannel(ctx, guildID)
	exists, err := r.postExists(ctx, e.RaidID)
	switch {
	case err != nil && active:
		r.markChanged(ctx, e)
		return
	case err != nil:
	case exists:
		r.markChanged(ctx, e)
		if active {
			return
		}
	case active && r.enqueue(ctx, guildID, e.RaidID):
		r.setMovedFrom(ctx, e)
		return
	}
	r.Webhook.RaidChanged(ctx, guildID, e)
}

// RaidReminder asks the drainer for the reminder when the bot is active and the raid has a
// live post: the drainer edits the post or replies under it (drain_reminder.go). Otherwise
// the webhook sends it. A failed mark with the bot active is returned, so the caller retries
// the reminder instead of sending it a second way.
func (r *RaidRouter) RaidReminder(ctx context.Context, guildID int64, e notify.RaidPosted) (notify.ReminderRef, bool, error) {
	if _, ok := r.activeChannel(ctx, guildID); ok {
		err := r.Bot.d.Store.MarkReminderDue(ctx, e.RaidID)
		switch {
		case err == nil:
			return notify.ReminderRef{Via: notify.ReminderViaBot}, true, nil
		case !errors.Is(err, ErrNoLivePost):
			return notify.ReminderRef{}, false, err
		}
	}
	return r.Webhook.RaidReminder(ctx, guildID, e)
}

// RaidReminderWithdrawn deletes a webhook reminder. A bot reminder is the drainer's to
// withdraw: the raid's own mark (a move, cancel or lock) brings it there.
func (r *RaidRouter) RaidReminderWithdrawn(ctx context.Context, guildID int64, ref notify.ReminderRef) {
	if ref.Via == notify.ReminderViaBot {
		return
	}
	r.Webhook.RaidReminderWithdrawn(ctx, guildID, ref)
}

// enqueue queues the raid's post; false (logged on a failure) when nothing was queued
// because the channel went away since ActiveInstall read it, so the caller uses the webhook.
func (r *RaidRouter) enqueue(ctx context.Context, guildID, raidID int64) bool {
	queued, err := r.Bot.d.Store.EnqueuePost(ctx, raidID, guildID)
	if err != nil {
		r.Bot.log.Error("discordbot", "op", "enqueue_post", "raid_id", raidID, "err", err)
		return false
	}
	return queued
}

func (r *RaidRouter) markChanged(ctx context.Context, e notify.RaidChanged) {
	r.setMovedFrom(ctx, e)
	r.Bot.d.Store.MarkRaidDirty(ctx, e.RaidID)
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
		r.Bot.reportRefused(ctx, guildID, channel, err)
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
	in, active, err := r.Bot.ActiveInstall(ctx, guildID)
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

// reportRefused records a report Discord did not take: a channel or server refusal is
// judged as the drainer judges it; any other refusal counts a failure.
func (s *Service) reportRefused(ctx context.Context, guildID int64, channelID string, err error) {
	s.log.Warn("discordbot", "op", "report_post", "guild_id", guildID, "err", describe(err))
	var api *APIError
	if errors.As(err, &api) && channelRefused(api) {
		if _, judgeErr := judgeChannel(ctx, s.d.Store, s.d.REST, guildID, channelID, err); judgeErr != nil {
			s.log.Error("discordbot", "op", "report_judge", "guild_id", guildID, "err", judgeErr)
		}
		return
	}
	if recErr := s.d.Store.RecordPostFailure(ctx, guildID, describe(err)); recErr != nil {
		s.log.Error("discordbot", "op", "report_record", "guild_id", guildID, "err", recErr)
	}
}
