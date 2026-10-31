package notify

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// The account webhook under the events router: the sim and parse jobs publish typed events
// (EventProducer), and DiscordSink renders the ones the webhook has always carried with the
// same embeds and the same gates (Deliver), so its output does not change.

// DiscordSink is the events router's Discord channel: the account's own webhook. It
// renders sim.finished and report.ready from their notify facts; every other kind has no
// Discord rendering per account yet (no DM sender), so it sends nothing.
type DiscordSink struct {
	Deliverer *Deliverer
}

var (
	_ events.Sink       = (*DiscordSink)(nil)
	_ events.KindFilter = (*DiscordSink)(nil)
)

// Renders reports whether the webhook carries kind.
func (s *DiscordSink) Renders(kind events.Kind) bool {
	return kind == events.SimFinished || kind == events.ReportReady
}

// Deliver posts e's embed when e carries a fact the webhook renders for userID.
func (s *DiscordSink) Deliver(ctx context.Context, userID int64, e events.Event) error {
	d := s.Deliverer
	switch src := e.Source.(type) {
	case SimFinished:
		if src.UserID != userID {
			return nil
		}
		return d.Deliver(ctx, userID, SimFinishedEmbed(d.PublicBaseURL, src), EventSimFinished)
	case UploadParsed:
		if src.UserID != userID {
			return nil
		}
		return d.Deliver(ctx, userID, UploadParsedEmbed(d.PublicBaseURL, src), uploadNeeds(src)...)
	}
	return nil
}

// EventProducer is the sim and parse jobs' notify port as events: each finished run or
// parsed upload is published, and the router sends it to the webhook and to push.
type EventProducer struct {
	Publisher     events.Publisher
	PublicBaseURL string
	Log           *slog.Logger
}

var (
	_ SimEvents    = (*EventProducer)(nil)
	_ UploadEvents = (*EventProducer)(nil)
)

// SimFinished publishes sim.finished to the run's owner.
func (p *EventProducer) SimFinished(ctx context.Context, e SimFinished) {
	events.Quietly(ctx, p.Publisher, p.Log, SimFinishedEvent(p.PublicBaseURL, e))
}

// UploadParsed publishes report.ready to the upload's owner.
func (p *EventProducer) UploadParsed(ctx context.Context, e UploadParsed) {
	events.Quietly(ctx, p.Publisher, p.Log, UploadParsedEvent(p.PublicBaseURL, e))
}

// pushCopy is the words of the push notifications built from notify facts.
var pushCopy = struct {
	SimTitle, SimTitleNamed, SimTitleNoSpec, SimBody             string
	LogTitle, LogTitleNoTitle, LogBody, LogBodyNoZone, LogFights string
}{
	SimTitle:        "Your %s sim finished",
	SimTitleNamed:   "%s's %s sim finished",
	SimTitleNoSpec:  "Your sim finished",
	SimBody:         "%s DPS",
	LogTitle:        "%s finished parsing",
	LogTitleNoTitle: "Your log finished parsing",
	LogBody:         "%s · %s",
	LogBodyNoZone:   "%s",
	LogFights:       "%d boss fights · %s",
}

// SimFinishedEvent is the run as an event: plain text for push, the fact for Discord.
func SimFinishedEvent(publicBaseURL string, e SimFinished) events.Event {
	title := pushCopy.SimTitleNoSpec
	if spec, ok := specs.ByKey[e.Spec]; ok {
		label := spec.Name + " " + classLabel(spec.ClassSlug)
		title = fmt.Sprintf(pushCopy.SimTitle, label)
		if name := collapse(e.Character); name != "" {
			title = fmt.Sprintf(pushCopy.SimTitleNamed, name, label)
		}
	}
	body := fmt.Sprintf(pushCopy.SimBody, thousands(e.DPS))
	if headline := collapse(e.Headline); headline != "" {
		body = headline
	}
	return events.Event{
		Kind: events.SimFinished, Audience: events.Users(e.UserID), DedupeKey: "sim.finished:" + e.SimID,
		Title: title, Body: body, Link: absoluteURL(publicBaseURL, simPathPrefix+e.SimID),
		Data: map[string]string{"sim_id": e.SimID}, Source: e,
	}
}

// UploadParsedEvent is the parsed upload as an event to its owner.
func UploadParsedEvent(publicBaseURL string, e UploadParsed) events.Event {
	title := pushCopy.LogTitleNoTitle
	if name := collapse(e.Title); name != "" {
		title = fmt.Sprintf(pushCopy.LogTitle, name)
	}
	fights := fmt.Sprintf(pushCopy.LogFights, e.Fights, killsPhrase(e.Kills))
	body := fmt.Sprintf(pushCopy.LogBodyNoZone, fights)
	if zone := collapse(e.Zone); zone != "" {
		body = fmt.Sprintf(pushCopy.LogBody, zone, fights)
	}
	return events.Event{
		Kind: events.ReportReady, Audience: events.Users(e.UserID), DedupeKey: "report.ready:" + e.ReportID,
		Title: title, Body: body, Link: absoluteURL(publicBaseURL, reportPathPrefix+e.ReportID),
		Data: map[string]string{"report_id": e.ReportID}, Source: e,
	}
}

// killsPhrase is the kills as the embeds word them.
func killsPhrase(kills int) string {
	switch kills {
	case 0:
		return embedCopy.LogKillsNone
	case 1:
		return embedCopy.LogOneKill
	}
	return fmt.Sprintf(embedCopy.LogManyKills, kills)
}
