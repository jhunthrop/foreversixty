package discordbot

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	// optReportRaid is /report's raid option: a raid id from the autocomplete.
	optReportRaid = "raid"
	// maxLootEmbeds is how many award embeds follow the report (spec §3.3.3: nine in all).
	maxLootEmbeds = 8
	// maxMostDeaths is how many raiders the Most deaths field names.
	maxMostDeaths = 3
	// reportChoices is how many done raids the raid option suggests.
	reportChoices = 10
	// doneGreen is #7fd48a, a done raid and its report.
	doneGreen = 8377482
	// fieldValueMax is Discord's field value limit.
	fieldValueMax = 1024
	// reportAnchor is the raid page's report section.
	reportAnchor = "#report"
	// raidTitleDate is "Thursday 10 December", in the guild's raid zone.
	raidTitleDate = "Monday 2 January"
)

// cmdReport answers /report: the last done raid's report (or the named one), ephemeral, with
// Share in channel only when the matched log is public.
func (s *Service) cmdReport(ctx context.Context, in Interaction) (Reply, error) {
	guildID, viewer, bound, err := s.lookupCaller(ctx, in)
	if err != nil || !bound {
		return pressReply(commonCopy.Unbound, false), err
	}
	if viewer == nil {
		return pressReply(officerCopy.ReportUnlinked, false, s.linkButton()), nil
	}
	raw, _ := in.Option(optReportRaid)
	return s.reportReply(ctx, in, guildID, *viewer, raw, false)
}

// reportReply is the one answer behind the command and a pick of the night select: the named
// raid's report (the last one for ""), with Share when its log is public, the coach lines and
// the select for another night. update edits the pressed card in place.
func (s *Service) reportReply(ctx context.Context, in Interaction, guildID, viewer int64, raw string, update bool) (Reply, error) {
	msg, ans, found, err := s.reportMessage(ctx, guildID, &viewer, map[string]string{optReportRaid: raw})
	if err != nil {
		return Reply{}, err
	}
	if !found {
		return pressReply(officerCopy.ReportNone, update), nil
	}
	if ans.Public {
		opts := map[string]string{optReportRaid: strconv.FormatInt(ans.Raid.ID, 10)}
		button, err := s.offerShare(ctx, in, guildID, "report", opts)
		if err != nil {
			return Reply{}, err
		}
		msg = withShare(msg, button)
	}
	msg = s.withCoachLines(ctx, msg, viewer, ans.Raid.Match.ReportID)
	nights, err := s.reportNightRow(ctx, guildID, viewer, ans.Raid.ID)
	if err != nil {
		return Reply{}, err
	}
	msg.Components = append(msg.Components, nights...)
	msg.Ephemeral = true
	return Reply{Message: msg, Update: update}, nil
}

// answerReport is /report's answer for viewer (nil: a visitor, who sees only a public log's
// night line and kill order, linking the public report).
func (s *Service) answerReport(ctx context.Context, guildID int64, viewer *int64, opts map[string]string) (Message, bool, error) {
	msg, _, found, err := s.reportMessage(ctx, guildID, viewer, opts)
	return msg, found, err
}

func (s *Service) reportMessage(ctx context.Context, guildID int64, viewer *int64, opts map[string]string) (Message, guilds.ReportAnswer, bool, error) {
	var raidID *int64
	if raw := strings.TrimSpace(opts[optReportRaid]); raw != "" {
		id, err := strconv.ParseInt(raw, 10, 64)
		if err != nil || id <= 0 {
			return Message{}, guilds.ReportAnswer{}, false, nil
		}
		raidID = &id
	}
	who := int64(0)
	if viewer != nil {
		who = *viewer
	}
	ans, found, err := s.d.Guilds.ReportFor(ctx, guildID, who, raidID)
	if err != nil {
		return Message{}, ans, false, fmt.Errorf("discordbot: report: %w", err)
	}
	if !found {
		return Message{}, ans, false, nil
	}
	g, err := s.d.Guilds.BotGuildFacts(ctx, guildID)
	if err != nil {
		return Message{}, ans, false, fmt.Errorf("discordbot: report guild: %w", err)
	}
	link := s.d.PublicBaseURL + "/reports/" + ans.Raid.Match.ReportID
	// Anyone who is not a verified member (a visitor, a share, a linked non-member) gets the
	// public answer: the night line and kill order, linking the public report. Loot and the raid
	// page are the members'.
	if ans.Member {
		path, err := s.d.Guilds.GuildPagePathOf(ctx, guildID)
		if err != nil {
			return Message{}, ans, false, fmt.Errorf("discordbot: report path: %w", err)
		}
		link = s.d.PublicBaseURL + guilds.RaidPagePath(path, ans.Raid.ID) + reportAnchor
	}
	embeds := []notify.Embed{reportHead(ans, link)}
	if ans.Member {
		embeds = append(embeds, s.lootEmbeds(ans.Report, link)...)
	}
	closeEmbeds(embeds, fill(guildFooter, "{Guild}", notify.Collapse(g.Name)), s.now())
	button := Component{Kind: KindButton, Style: StyleLink, Label: officerCopy.OpenReport, URL: link}
	return Message{Embeds: embeds, Components: []ActionRow{{button}}}, ans, true, nil
}

// guildFooter closes a guild answer (spec §3.3.0).
const guildFooter = "{Guild} · Forever Sixty"

// raidTitle is "{Title} · Thursday 10 December" in the guild's raid zone.
func raidTitle(title string, startsAt time.Time, zone *string) string {
	loc := time.UTC
	if zone != nil {
		loc = raidLocation(*zone)
	}
	return memberText(title) + lineJoin + startsAt.In(loc).Format(raidTitleDate)
}

// reportHead is the report's main embed: the night line and kill order, and for a member the
// Loot count, with the officer rows before it for an officer.
func reportHead(ans guilds.ReportAnswer, link string) notify.Embed {
	v := ans.Raid
	e := notify.Embed{
		Title: officerCopy.ReportPrefix + raidTitle(v.Title, v.StartsAt, v.GuildTimezone), URL: link, Color: doneGreen,
		Description: nightLine(*v.Night),
	}
	if kills := boundedLines(killLines(ans.Report.KillOrder)); kills != "" {
		e.Fields = append(e.Fields, notify.EmbedField{Name: officerCopy.KillOrder, Value: kills})
	}
	if !ans.Member {
		return e
	}
	if ans.Officer {
		if d := mostDeaths(ans.Report.Deaths); d != "" {
			e.Fields = append(e.Fields, notify.EmbedField{Name: officerCopy.MostDeaths, Value: d})
		}
		e.Fields = append(e.Fields, notify.EmbedField{Name: officerCopy.Consumables, Value: consumablesLine(ans.Report.Consumables)})
	}
	e.Fields = append(e.Fields, notify.EmbedField{Name: officerCopy.Loot, Value: lootLine(lootCount(ans.Report))})
	return e
}

// nightLine is G2's night line: "22 of 23 came · 1 kill · 3 wipes · 14 deaths · 2 h 41 min",
// a zero clause left out.
func nightLine(n guilds.NightView) string {
	parts := []string{fill(nightCopy.Came, "{came}", strconv.Itoa(n.Came), "{signed}", strconv.Itoa(n.Signed))}
	if n.Joined > 0 {
		parts = append(parts, fmt.Sprintf("%d joined", n.Joined))
	}
	for _, c := range []struct {
		n    int
		unit string
	}{{n.Kills, "kill"}, {n.Wipes, "wipe"}, {n.Deaths, "death"}} {
		if c.n > 0 {
			parts = append(parts, units(c.n, c.unit))
		}
	}
	if n.DurationMS > 0 {
		parts = append(parts, nightDuration(n.DurationMS))
	}
	return strings.Join(parts, lineJoin)
}

func units(n int, singular string) string {
	if n == 1 {
		return fmt.Sprintf("%d %s", n, singular)
	}
	return fmt.Sprintf("%d %ss", n, singular)
}

// nightDuration is "2 h 41 min", "2 h" or "41 min"; under a minute reads "1 min".
func nightDuration(ms int64) string {
	total := max(1, int((time.Duration(ms)*time.Millisecond+30*time.Second)/time.Minute))
	hours, minutes := total/60, total%60
	switch {
	case hours == 0:
		return fmt.Sprintf("%d min", minutes)
	case minutes == 0:
		return fmt.Sprintf("%d h", hours)
	default:
		return fmt.Sprintf("%d h %d min", hours, minutes)
	}
}

// killLines are "1. Onyxia 1:42 after 3 wipes", "2. Onyxia 1:42 first pull" or "3. Nefarian
// Not killed · 4 wipes", one per named boss in kill order.
func killLines(kills []guilds.KillOrderView) []string {
	out := make([]string, 0, len(kills))
	for i, k := range kills {
		boss := memberText(k.Name)
		if !k.Killed || k.KillMS == nil {
			out = append(out, fmt.Sprintf("%d. %s %s%s%s", i+1, boss, nightCopy.NotKilled, lineJoin, units(k.WipesBefore, "wipe")))
			continue
		}
		note := nightCopy.FirstPull
		if k.WipesBefore > 0 {
			note = fill(nightCopy.After, "{wipes}", units(k.WipesBefore, "wipe"))
		}
		out = append(out, fmt.Sprintf("%d. %s %s %s", i+1, boss, killClock(*k.KillMS), note))
	}
	return out
}

// killClock is a kill time as m:ss, or h:mm:ss from an hour up.
func killClock(ms int64) string {
	total := int((time.Duration(ms)*time.Millisecond + 500*time.Millisecond) / time.Second)
	if total >= 3600 {
		return fmt.Sprintf("%d:%02d:%02d", total/3600, total%3600/60, total%60)
	}
	return fmt.Sprintf("%d:%02d", total/60, total%60)
}

// boundedLines joins lines inside a field value, dropping the tail that does not fit; "" for none.
func boundedLines(lines []string) string {
	out := ""
	for _, l := range lines {
		next := l
		if out != "" {
			next = out + "\n" + l
		}
		if len(next) > fieldValueMax {
			break
		}
		out = next
	}
	return out
}

// mostDeaths is "Pyrewisp 4 · Windtalon 3 · Vexlash 2", the top three; "" when nobody died.
func mostDeaths(deaths []guilds.ReportDeathView) string {
	rows := append([]guilds.ReportDeathView(nil), deaths...)
	sort.SliceStable(rows, func(i, j int) bool {
		return rows[i].OnWipes+rows[i].OnKills > rows[j].OnWipes+rows[j].OnKills
	})
	parts := []string{}
	for _, d := range rows {
		if n := d.OnWipes + d.OnKills; n > 0 && len(parts) < maxMostDeaths {
			parts = append(parts, fmt.Sprintf("%s %d", memberText(d.Name), n))
		}
	}
	return strings.Join(parts, lineJoin)
}

// consumablesLine is "3 short before the raid", then " · 2 with no flask in the log" once the
// log carries buff data.
func consumablesLine(c guilds.ReportConsumablesView) string {
	short, noFlask := 0, 0
	for _, r := range c.Rows {
		if r.Before == "short" {
			short++
		}
		if r.Flask != nil && !*r.Flask {
			noFlask++
		}
	}
	line := fill(nightCopy.ShortBefore, "{n}", strconv.Itoa(short))
	if c.LogHasBuffs {
		line += lineJoin + fill(nightCopy.NoFlask, "{n}", strconv.Itoa(noFlask))
	}
	return line
}

func lootCount(r *guilds.ReportView) int {
	n := 0
	for _, b := range r.Loot {
		n += len(b.Awards)
	}
	return n
}

// lootLine is "6 items." and, when items follow, " The first 6:".
func lootLine(n int) string {
	template := officerCopy.LootCount
	if n > 0 && n <= maxLootEmbeds { // every item follows: no "The first" repeating the count
		template = officerCopy.LootAll
	}
	line := fill(template, "{n}", strconv.Itoa(n))
	if n == 1 {
		line = strings.Replace(line, "items", "item", 1)
	}
	if n > maxLootEmbeds {
		line += fill(officerCopy.LootFirst, "{k}", strconv.Itoa(maxLootEmbeds))
	}
	return line
}

type lootAward struct {
	guilds.ReportLootAward
	boss string
}

// lootEmbeds are the night's awards, most valuable quality first, at most eight, each the item
// embed with "to {Name} · {Boss}".
func (s *Service) lootEmbeds(r *guilds.ReportView, link string) []notify.Embed {
	var awards []lootAward
	for _, b := range r.Loot {
		for _, a := range b.Awards {
			awards = append(awards, lootAward{a, b.Boss})
		}
	}
	sort.SliceStable(awards, func(i, j int) bool { return awards[i].Quality > awards[j].Quality })
	if len(awards) > maxLootEmbeds {
		awards = awards[:maxLootEmbeds]
	}
	build, haveBuild := s.d.Guilds.ActiveBuild()
	out := make([]notify.Embed, 0, len(awards))
	for i, a := range awards {
		source := fill(officerCopy.LootTo, "{Name}", notify.Collapse(a.Name), "{Boss}", notify.Collapse(a.boss))
		// The award's index joins the id: the same item can drop twice in one night.
		key := strconv.Itoa(a.ItemID) + "-" + strconv.Itoa(i)
		facts := ItemFacts{ID: a.ItemID, Name: a.ItemName, Quality: a.Quality, Source: source, URL: itemLink(link, key)}
		if haveBuild { // without a build there is no item table and no icon path: name and quality only
			if it, ok := s.d.Guilds.BuildItem(build, a.Class, a.ItemID); ok {
				facts.ItemLevel = it.ItemLevel
			}
			if a.Icon != "" {
				facts.Icon = s.d.PublicBaseURL + "/data/" + build + "/icons/" + a.Icon + ".webp"
			}
		}
		out = append(out, ItemEmbed(facts, "", "", ""))
	}
	return out
}

// autocompleteReport suggests /report's raid: the guild's last ten logged raids a member may
// open (a linked non-member, the public ones), newest first.
func (s *Service) autocompleteReport(ctx context.Context, in Interaction) ([]Choice, error) {
	focused, typed := in.Focused()
	if focused != optReportRaid {
		return []Choice{}, nil
	}
	guildID, viewer, bound, err := s.lookupCaller(ctx, in)
	if err != nil || !bound || viewer == nil {
		return []Choice{}, err
	}
	member, err := s.verifiedMember(ctx, guildID, viewer)
	if err != nil {
		return nil, err
	}
	raids, err := s.d.Guilds.DoneRaids(ctx, guildID, !member, reportChoices)
	if err != nil {
		return nil, fmt.Errorf("discordbot: report autocomplete: %w", err)
	}
	zone, err := s.d.Guilds.GuildRaidTimezone(ctx, guildID)
	if err != nil {
		return nil, fmt.Errorf("discordbot: report autocomplete zone: %w", err)
	}
	loc := raidLocation(zone)
	out := []Choice{}
	for _, r := range raids {
		date := choiceJoin + r.StartsAt.In(loc).Format(raidChoiceDate)
		title := truncateRunes(r.Title, maxChoiceRunes-len([]rune(date)))
		if matches(title, typed) {
			out = append(out, Choice{Name: title + date, Value: strconv.FormatInt(r.ID, 10)})
		}
	}
	return out, nil
}
