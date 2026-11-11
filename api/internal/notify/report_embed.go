package notify

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/textfmt"
)

// reportColor is the site's Done mark, the kill green (spec 2026-10-27-guild-attendance §6.2).
var reportColor = brand.RGB(brand.ColorOutcomeKill)

const (
	// reportMaxKills is how many kill lines the embed lists before "+n more".
	reportMaxKills = 6
	// maxBossRunes bounds a boss name a kill line carries.
	maxBossRunes = 60
	// reportMaxLoot is how many awards the Loot field names before "+n more".
	reportMaxLoot = 5
	// maxNameRunes bounds a raider or item name a line carries.
	maxNameRunes = 60
	// reportAnchor is the raid page's report section.
	reportAnchor = "#report"
)

// reportCopy is the words of the report_ready post.
var reportCopy = struct {
	TitlePrefix string
	// CoachSentence is the description when a raider has the log coach: it
	// names and counts nobody (R-A3; experience spec 2026-10-28-log-coach §8.3 raidNight).
	CoachSentence string
	KillsName     string
	KillLine      string // %s the escaped boss, %s the kill time (markup, no words)
	MoreKills     string // {n} the kills left out
	NoKills       string
	WipesName     string
	CameName      string
	CameOf        string // {came} came, {of} signed
	LootName      string
	LootOne       string
	LootMany      string // {n} items
	LootLine      string // {item} the item, {winner} the raider
	MoreLoot      string // {n} awards left out
	TimeName      string
	DeathsName    string
	ParsesName    string
	ParseLine     string // {name} the raider, {value} the figure, {metric} the metric, {boss} the boss
}{
	TitlePrefix:   copygen.DiscordOfficerReportPrefix,
	CoachSentence: copygen.DiscordCoachRaidNight,
	KillsName:     copygen.DiscordRaidKillsName,
	KillLine:      "%s %s",
	MoreKills:     copygen.DiscordReportMoreKills,
	NoKills:       copygen.DiscordReportNoKills,
	WipesName:     copygen.DiscordReportWipesName,
	CameName:      copygen.DiscordRaidCameName,
	CameOf:        copygen.DiscordRaidCameOf,
	LootName:      copygen.DiscordOfficerLoot,
	LootOne:       copygen.DiscordReportLootOne,
	LootMany:      copygen.DiscordOfficerLootAll,
	LootLine:      copygen.DiscordReportLootLine,
	MoreLoot:      copygen.DiscordReportMoreKills,
	TimeName:      copygen.DiscordReportTimeName,
	DeathsName:    copygen.DiscordReportDeathsName,
	ParsesName:    copygen.DiscordReportParsesName,
	ParseLine:     copygen.DiscordReportParseLine,
}

// ReportKill is one boss killed on the night.
type ReportKill struct {
	Boss     string
	Duration time.Duration
}

// ReportParse is one of the night's best parses: the raider, the figure, its metric, the boss.
type ReportParse struct {
	Name, Metric, Boss string
	Value              int
}

// ReportLoot is one award of the night.
type ReportLoot struct {
	Item, Winner string
}

// ReportReady is what an officer's "Post to Discord" on a done raid holds. It names raiders only
// for good news (the best parses and the loot winners), never who came, who was missing or who
// died most (R-A3; a channel is no place for a worst-first list). RaidPosted carries the guild and
// the raid; its RaidURL should be the raid page, the embed adds the report anchor.
type ReportReady struct {
	RaidPosted
	Kills        []ReportKill
	Wipes        int
	Came, Signed int
	LootItems    int
	// Loot names the awards; with none listed the field counts LootItems.
	Loot []ReportLoot
	// Time is the night from the first pull to the last; zero leaves the field out.
	Time time.Duration
	// Deaths counts the night's deaths, every pull.
	Deaths int
	// TopParses are the night's best damage parses, best first; none leaves the field out.
	TopParses []ReportParse
	// CoachReady adds the coach sentence: some raider in the log holds log_coach.
	CoachReady bool
}

// ReportReadyEmbed is the report_ready message.
func ReportReadyEmbed(in ReportReady, now time.Time) Embed {
	post := in.RaidPosted
	if isAbsolute(post.RaidURL) {
		post.RaidURL += reportAnchor
	}
	e := raidBaseEmbed(post, reportCopy.TitlePrefix, now)
	e.Color = reportColor
	e.Description = reportDescription(in.CoachReady)
	e.Fields = []EmbedField{
		raidField(reportCopy.KillsName, killsValue(in.Kills)),
		raidField(reportCopy.WipesName, fmt.Sprint(in.Wipes)),
	}
	if in.Time > 0 {
		e.Fields = append(e.Fields, raidField(reportCopy.TimeName, textfmt.Duration(in.Time)))
	}
	e.Fields = append(e.Fields,
		raidField(reportCopy.CameName, copygen.Fill(reportCopy.CameOf, "{came}", strconv.Itoa(in.Came), "{of}", strconv.Itoa(in.Signed))),
		raidField(reportCopy.DeathsName, fmt.Sprint(in.Deaths)))
	if len(in.TopParses) > 0 {
		e.Fields = append(e.Fields, EmbedField{Name: reportCopy.ParsesName, Value: parsesValue(in.TopParses)})
	}
	e.Fields = append(e.Fields, EmbedField{Name: reportCopy.LootName, Value: lootValue(in), Inline: len(in.Loot) == 0})
	return e
}

func parsesValue(parses []ReportParse) string {
	lines := make([]string, len(parses))
	for i, p := range parses {
		lines[i] = copygen.Fill(reportCopy.ParseLine,
			"{name}", escapedWithin(collapse(p.Name), maxNameRunes, maxNameRunes*2),
			"{value}", strconv.Itoa(p.Value), "{metric}", strings.ToUpper(p.Metric),
			"{boss}", escapedWithin(collapse(p.Boss), maxBossRunes, maxBossRunes*2))
	}
	return strings.Join(lines, "\n")
}

// reportDescription is the coach sentence when a raider has the log coach, else nothing: the
// fields are the night (round 5 dropped the line that described them).
func reportDescription(coachReady bool) string {
	if coachReady {
		return reportCopy.CoachSentence
	}
	return ""
}

func killsValue(kills []ReportKill) string {
	if len(kills) == 0 {
		return reportCopy.NoKills
	}
	lines := make([]string, 0, reportMaxKills+1)
	for i, k := range kills {
		if i == reportMaxKills {
			lines = append(lines, counted(reportCopy.MoreKills, len(kills)-reportMaxKills))
			break
		}
		boss := escapedWithin(collapse(k.Boss), maxBossRunes, maxBossRunes*2)
		lines = append(lines, fmt.Sprintf(reportCopy.KillLine, boss, textfmt.FightClock(k.Duration)))
	}
	return strings.Join(lines, "\n")
}

func lootValue(in ReportReady) string {
	if len(in.Loot) == 0 {
		if in.LootItems == 1 {
			return reportCopy.LootOne
		}
		return counted(reportCopy.LootMany, in.LootItems)
	}
	lines := make([]string, 0, reportMaxLoot+1)
	for i, l := range in.Loot {
		if i == reportMaxLoot {
			lines = append(lines, counted(reportCopy.MoreLoot, len(in.Loot)-reportMaxLoot))
			break
		}
		lines = append(lines, copygen.Fill(reportCopy.LootLine,
			"{item}", escapedWithin(collapse(l.Item), maxNameRunes, maxNameRunes*2),
			"{winner}", escapedWithin(collapse(l.Winner), maxNameRunes, maxNameRunes*2)))
	}
	return strings.Join(lines, "\n")
}
