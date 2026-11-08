package livesurface

// The raid Live Activity's pushes (spec 2026-11-08-app-flows F2): on raid.starting a
// push-to-start for every account In or Late that keeps a start token and has no countdown
// running for the raid yet, and when the raid is logged (marked done, or its report matched) an
// end with the night's kills, the loot the raider got and whether the report is ready. iOS goes
// through APNs on the liveactivity push type; Android has no Live Activity, so its ongoing
// notification gets the same facts as an FCM data message. A token the service refuses goes.

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/push"
)

// LiveAPNs sends an ActivityKit push (push.APNsSender).
type LiveAPNs interface {
	SendLive(ctx context.Context, token string, payload map[string]any) error
}

// LiveFCM sends a whole FCM send body (push.FCMSender).
type LiveFCM interface {
	SendBody(ctx context.Context, body map[string]any) error
}

// LiveRaidSource is what the raid pushes read (guilds.Service).
type LiveRaidSource interface {
	RaidAttendees(ctx context.Context, raidID int64) ([]guilds.RaidAttendee, error)
	WidgetRaidFor(ctx context.Context, userID int64, characterKey string, now time.Time) (guilds.WidgetRaid, error)
	RaidNightOf(ctx context.Context, viewerID, raidID int64, now time.Time) (guilds.RaidNightSummary, bool, error)
}

// liveEndDismissal is how long an ended raid's activity stays on the lock screen.
const liveEndDismissal = 4 * time.Hour

// liveStartCopy is the push-to-start's alert.
var liveStartCopy = struct{ Body string }{Body: "Starts in %d minutes."}

// LivePusher sends the raid Live Activity's start and end (guilds.LiveRaids). A nil APNs or FCM
// skips that platform.
type LivePusher struct {
	APNs   LiveAPNs
	FCM    LiveFCM
	Tokens *Store
	Raids  LiveRaidSource
	Log    *slog.Logger
	// Now is the clock; nil reads time.Now.
	Now func() time.Time
}

var _ guilds.LiveRaids = (*LivePusher)(nil)

func (p *LivePusher) now() time.Time {
	if p.Now != nil {
		return p.Now()
	}
	return time.Now()
}

func (p *LivePusher) logger() *slog.Logger {
	if p.Log != nil {
		return p.Log
	}
	return slog.Default()
}

// RaidStarting starts the raid's countdown for every In or Late account with a start token and
// none running. Failures are logged; the event has already gone.
func (p *LivePusher) RaidStarting(ctx context.Context, raidID int64) {
	attendees, err := p.Raids.RaidAttendees(ctx, raidID)
	if err != nil {
		p.logger().Error("livesurface", "op", "live_raid_start", "raid_id", raidID, "err", err)
		return
	}
	object := strconv.FormatInt(raidID, 10)
	for _, a := range attendees {
		if err := p.startFor(ctx, raidID, object, a); err != nil {
			p.logger().Error("livesurface", "op", "live_raid_start", "raid_id", raidID, "user", a.UserID, "err", err)
		}
	}
}

// startFor starts a's countdown on each of their start tokens, unless one is running.
func (p *LivePusher) startFor(ctx context.Context, raidID int64, object string, a guilds.RaidAttendee) error {
	tokens, err := p.Tokens.StartTokens(ctx, a.UserID, KindRaid)
	if err != nil || len(tokens) == 0 {
		return err
	}
	running, err := p.Tokens.Running(ctx, a.UserID, KindRaid, object)
	if err != nil || running {
		return err
	}
	now := p.now()
	w, err := p.Raids.WidgetRaidFor(ctx, a.UserID, a.CharacterKey, now)
	if err != nil || w.NextRaid == nil || w.NextRaid.RaidID != raidID {
		return err
	}
	attrs, state := RaidLive(*w.NextRaid, a.Name)
	minutes := int(time.Until(w.NextRaid.StartsAt).Round(time.Minute).Minutes())
	alert := &Alert{Title: attrs.ZoneName, Body: fmt.Sprintf(liveStartCopy.Body, max(minutes, 0))}
	stale := w.NextRaid.EndsAt
	for _, t := range tokens {
		err := p.send(ctx, t.Platform, t.Token,
			func() map[string]any {
				return APNsLive("start", now, RaidActivityType, attrs, state, alert, &stale, nil)
			},
			func() map[string]any { return FCMData(t.Token, EventRaidStart, RaidData(attrs, state), true) })
		if errors.Is(err, push.ErrInvalidToken) {
			_ = p.Tokens.DropStartToken(ctx, a.UserID, t.DeviceID, KindRaid)
			continue
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// RaidEnded ends every countdown kept for the raid with the raider's night. Failures are logged.
func (p *LivePusher) RaidEnded(ctx context.Context, raidID int64) {
	now := p.now()
	subs, err := p.Tokens.Subscribers(ctx, KindRaid, strconv.FormatInt(raidID, 10), now)
	if err != nil {
		p.logger().Error("livesurface", "op", "live_raid_end", "raid_id", raidID, "err", err)
		return
	}
	for _, sub := range subs {
		if err := p.endFor(ctx, raidID, sub, now); err != nil {
			p.logger().Error("livesurface", "op", "live_raid_end", "raid_id", raidID, "user", sub.UserID, "err", err)
		}
	}
}

// endFor sends sub the raid's end from their raid night.
func (p *LivePusher) endFor(ctx context.Context, raidID int64, sub Subscriber, now time.Time) error {
	night, found, err := p.Raids.RaidNightOf(ctx, sub.UserID, raidID, now)
	if err != nil || !found {
		return err
	}
	state := RaidEnd(night)
	attrs := RaidAttributes{RaidID: strconv.FormatInt(raidID, 10), GuildID: strconv.FormatInt(night.GuildID, 10),
		GuildName: night.GuildName, ZoneName: night.ZoneName, CharacterName: night.Character.Name, ZoneIcon: night.ZoneIcon}
	dismissal := now.Add(liveEndDismissal)
	err = p.send(ctx, sub.Platform, sub.PushToken,
		func() map[string]any { return APNsLive("end", now, RaidActivityType, nil, state, nil, nil, &dismissal) },
		func() map[string]any { return FCMData(sub.PushToken, EventRaidEnd, RaidData(attrs, state), true) })
	if errors.Is(err, push.ErrInvalidToken) {
		return p.Tokens.Drop(ctx, sub.ID)
	}
	return err
}

// send posts the platform's payload; a platform with no sender is skipped.
func (p *LivePusher) send(ctx context.Context, platform, token string, apns, fcm func() map[string]any) error {
	switch platform {
	case string(push.IOS):
		if p.APNs != nil {
			return p.APNs.SendLive(ctx, token, apns())
		}
	case string(push.Android):
		if p.FCM != nil {
			return p.FCM.SendBody(ctx, fcm())
		}
	}
	return nil
}
