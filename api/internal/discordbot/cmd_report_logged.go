package discordbot

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

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
	note := officerCopy.LoggedNote
	if ans.PlannedRaid != nil {
		note = plannedNote(ans.PlannedRaid.Title, s.raidPageURL(ctx, ans.PlannedRaid.ID))
	}
	embeds := []notify.Embed{loggedNightEmbed(ans, link, note, s.now())}
	s.closeEmbeds(embeds, fill(guildFooter, "{Guild}", notify.Collapse(g.Name)), s.now())
	button := Component{Kind: KindButton, Style: StyleLink, Label: officerCopy.OpenReport, URL: link}
	msg := Message{Embeds: embeds, Components: []ActionRow{{button}}}
	msg = s.withCoachLines(ctx, msg, viewer, ans.ReportID)
	msg.Ephemeral = true
	return Reply{Message: msg, Update: update}, nil
}

// plannedNote is the note for a log a planned raid waits for that night: "{Raid} was planned
// that night: match the log on the raid page.", the raid page a link when there is one.
func plannedNote(title, raidURL string) string {
	page := officerCopy.LoggedPlannedPage
	if raidURL != "" {
		page = "[" + page + "](" + raidURL + ")"
	}
	return fill(officerCopy.LoggedPlanned, "{Raid}", notify.Collapse(title), "{page}", page)
}

// loggedNightEmbed is the logged night's head: "Report: Onyxia's Lair · Thursday 10 December",
// "5 present · 1 kill · 3 wipes · 14 deaths · 2 h 41 min", the note, then the kill order.
func loggedNightEmbed(ans guilds.LoggedNightAnswer, link, note string, now time.Time) notify.Embed {
	name := ans.Zone
	if name == "" {
		name = ans.Title
	}
	e := notify.Embed{
		Title: officerCopy.ReportPrefix + raidTitle(name, ans.At, ans.GuildTimezone, now), URL: link, Color: doneGreen,
		Description: loggedNightLine(ans) + "\n" + note,
	}
	if kills := boundedLines(killLines(ans.KillOrder)); kills != "" {
		e.Fields = append(e.Fields, notify.EmbedField{Name: officerCopy.KillOrder, Value: kills})
	}
	return e
}

// loggedNightLine is the night line with "{n} present" where a planned raid's says who came,
// then the boss kills, the trash packs (pulls with no named encounter, never counted as kills:
// the Overview's words) and the rest of the tally.
func loggedNightLine(ans guilds.LoggedNightAnswer) string {
	parts := []string{fill(officerCopy.LoggedPresent, "{n}", strconv.Itoa(ans.Present))}
	tally := tallyClauses(ans.Kills, ans.Wipes, ans.Deaths, ans.DurationMS)
	if ans.TrashPacks > 0 {
		at := 0
		if ans.Kills > 0 {
			at = 1
		}
		tally = append(tally[:at], append([]string{units(ans.TrashPacks, "trash pack")}, tally[at:]...)...)
	}
	return strings.Join(append(parts, tally...), lineJoin)
}
