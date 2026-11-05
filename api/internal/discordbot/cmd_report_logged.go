package discordbot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// loggedNightReply is a bare /report's answer when no planned raid has a matched log: the
// newest night on the guild's Raids tab, read from its log alone (bot spec §3.3.3, amended
// 2026-11-05). It links the report page; there is no raid page, loot or Share to offer.
func (s *Service) loggedNightReply(ctx context.Context, guildID, viewer int64, update bool) (Reply, error) {
	ans, found, err := s.d.Guilds.LatestLoggedNight(ctx, guildID, viewer)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: logged night: %w", err)
	}
	if !found {
		return pressReply(officerCopy.ReportNone, update), nil
	}
	g, err := s.d.Guilds.BotGuildFacts(ctx, guildID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: logged night guild: %w", err)
	}
	link := s.d.PublicBaseURL + "/reports/" + ans.ReportID
	embeds := []notify.Embed{loggedNightEmbed(ans, link)}
	closeEmbeds(embeds, fill(guildFooter, "{Guild}", notify.Collapse(g.Name)), s.now())
	button := Component{Kind: KindButton, Style: StyleLink, Label: officerCopy.OpenReport, URL: link}
	msg := Message{Embeds: embeds, Components: []ActionRow{{button}}}
	msg = s.withCoachLines(ctx, msg, viewer, ans.ReportID)
	msg.Ephemeral = true
	return Reply{Message: msg, Update: update}, nil
}

// loggedNightEmbed is the logged night's head: "Report: Onyxia's Lair · Thursday 10 December",
// "5 present · 1 kill · 3 wipes · 14 deaths · 2 h 41 min", the note, then the kill order.
func loggedNightEmbed(ans guilds.LoggedNightAnswer, link string) notify.Embed {
	name := ans.Zone
	if name == "" {
		name = ans.Title
	}
	e := notify.Embed{
		Title: officerCopy.ReportPrefix + raidTitle(name, ans.At, ans.GuildTimezone), URL: link, Color: doneGreen,
		Description: loggedNightLine(ans) + "\n" + officerCopy.LoggedNote,
	}
	if kills := boundedLines(killLines(ans.KillOrder)); kills != "" {
		e.Fields = append(e.Fields, notify.EmbedField{Name: officerCopy.KillOrder, Value: kills})
	}
	return e
}

// loggedNightLine is the night line with "{n} present" where a planned raid's says who came.
func loggedNightLine(ans guilds.LoggedNightAnswer) string {
	parts := []string{fill(officerCopy.LoggedPresent, "{n}", strconv.Itoa(ans.Present))}
	return strings.Join(append(parts, tallyClauses(ans.Kills, ans.Wipes, ans.Deaths, ans.DurationMS)...), lineJoin)
}
