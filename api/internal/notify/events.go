package notify

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/textfmt"
	"github.com/jhunthrop/foreversixty/api/internal/zoneicons"
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

// pushPairBody joins two parts of a push body: a log's zone and its fights line, a run's DPS
// and its name.
const pushPairBody = "%s · %s"

// SimFinishedEvent is the run as an event: plain text for push, the fact for Discord. Push and
// the webhook embed word the run with the same keys (embedCopy).
func SimFinishedEvent(publicBaseURL string, e SimFinished) events.Event {
	title := embedCopy.SimTitleNoSpec
	if spec, ok := specs.ByKey[e.Spec]; ok {
		label := spec.FullName
		title = copygen.Fill(embedCopy.SimTitle, "{spec}", label)
		if name := collapse(e.Character); name != "" {
			title = copygen.Fill(embedCopy.SimTitleNamed, "{name}", name, "{spec}", label)
		}
	}
	body := copygen.Fill(embedCopy.SimResultDPS, "{dps}", textfmt.Grouped(e.DPS))
	if run := collapse(e.Title); run != "" {
		body = fmt.Sprintf(pushPairBody, body, run)
	}
	if headline := collapse(e.Headline); headline != "" {
		body = headline
	}
	out := events.Event{
		Kind: events.SimFinished, Audience: events.Users(e.UserID), DedupeKey: "sim.finished:" + e.SimID,
		Title: title, Body: body, Link: absoluteURL(publicBaseURL, simPathPrefix+e.SimID), Icon: events.IconSim,
		Data: map[string]string{"sim_id": e.SimID}, Source: e,
	}
	if e.DPS > 0 {
		dps := e.DPS
		out.Figure = &dps
		// The Activity row draws the DPS itself: it keeps the headline alone.
		headline := collapse(e.Headline)
		out.ActivityBody = &headline
	}
	return out
}

// UploadParsedEvent is the parsed upload as an event to its owner.
func UploadParsedEvent(publicBaseURL string, e UploadParsed) events.Event {
	title := embedCopy.LogTitleNoTitle
	if name := collapse(e.Title); name != "" {
		title = copygen.Fill(embedCopy.LogTitle, "{title}", name)
	}
	body := copygen.Fill(embedCopy.LogPushFights, "{fights}", embedCopy.LogBossFights.For(e.Fights), "{kills}", killsPhrase(e.Kills))
	icon := ""
	if zone := collapse(e.Zone); zone != "" {
		icon = zoneicons.ForName(zone)
		// The zone leads the body unless the title already names it ("Tuesday Molten Core").
		if !strings.Contains(strings.ToLower(title), strings.ToLower(zone)) {
			body = fmt.Sprintf(pushPairBody, zone, body)
		}
	}
	return events.Event{
		Kind: events.ReportReady, Audience: events.Users(e.UserID), DedupeKey: "report.ready:" + e.ReportID,
		Title: title, Body: body, Link: absoluteURL(publicBaseURL, reportPathPrefix+e.ReportID),
		Data: map[string]string{"report_id": e.ReportID}, Source: e, Icon: icon,
	}
}

// killsPhrase is the kills as the embeds word them.
func killsPhrase(kills int) string {
	if kills == 0 {
		return embedCopy.LogKillsNone
	}
	return embedCopy.LogKills.For(kills)
}
