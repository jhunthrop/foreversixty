package notify

import (
	"fmt"
	"strings"
	"time"
)

const (
	// reportColor is #7fd48a, the site's Done mark (spec 2026-10-27-guild-attendance §6.2).
	reportColor = 8377482
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
	Description string
	// CoachSentence follows Description after one space when a raider has the log coach: it
	// names and counts nobody (R-A3; experience spec 2026-10-28-log-coach §8.3 raidNight).
	CoachSentence string
	KillsName     string
	KillLine      string // %s the escaped boss, %s the kill time
	MoreKills     string // %d the kills left out
	NoKills       string
	WipesName     string
	CameName      string
	CameOf        string // %d came, %d signed
	LootName      string
	LootOne       string
	LootMany      string // %d items
	LootLine      string // %s the item, %s the raider
	MoreLoot      string // %d awards left out
	TimeName      string
	DeathsName    string
	ParsesName    string
	ParseLine     string // %s the raider, %d the figure, %s the metric, %s the boss
}{
	TitlePrefix:   "Report: ",
	Description:   "Kill times, loot and the biggest changes on Forever Sixty.",
	CoachSentence: "Coach reviews for tonight's pulls are ready; open the report or use /report.",
	KillsName:     "Kills",
	KillLine:      "%s %s",
	MoreKills:     "+%d more",
	NoKills:       "No kills",
	WipesName:     "Wipes",
	CameName:      "Came",
	CameOf:        "%d of %d",
	LootName:      "Loot",
	LootOne:       "1 item",
	LootMany:      "%d items",
	LootLine:      "%s to %s",
	MoreLoot:      "+%d more",
	TimeName:      "Time",
	DeathsName:    "Deaths",
	ParsesName:    "Top parses",
	ParseLine:     "%s %d %s on %s",
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
		e.Fields = append(e.Fields, raidField(reportCopy.TimeName, killClock(in.Time)))
	}
	e.Fields = append(e.Fields,
		raidField(reportCopy.CameName, fmt.Sprintf(reportCopy.CameOf, in.Came, in.Signed)),
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
		lines[i] = fmt.Sprintf(reportCopy.ParseLine, escapedWithin(collapse(p.Name), maxNameRunes, maxNameRunes*2),
			p.Value, strings.ToUpper(p.Metric), escapedWithin(collapse(p.Boss), maxBossRunes, maxBossRunes*2))
	}
	return strings.Join(lines, "\n")
}

func reportDescription(coachReady bool) string {
	if coachReady {
		return reportCopy.Description + " " + reportCopy.CoachSentence
	}
	return reportCopy.Description
}

func killsValue(kills []ReportKill) string {
	if len(kills) == 0 {
		return reportCopy.NoKills
	}
	lines := make([]string, 0, reportMaxKills+1)
	for i, k := range kills {
		if i == reportMaxKills {
			lines = append(lines, fmt.Sprintf(reportCopy.MoreKills, len(kills)-reportMaxKills))
			break
		}
		boss := escapedWithin(collapse(k.Boss), maxBossRunes, maxBossRunes*2)
		lines = append(lines, fmt.Sprintf(reportCopy.KillLine, boss, killClock(k.Duration)))
	}
	return strings.Join(lines, "\n")
}

func lootValue(in ReportReady) string {
	if len(in.Loot) == 0 {
		if in.LootItems == 1 {
			return reportCopy.LootOne
		}
		return fmt.Sprintf(reportCopy.LootMany, in.LootItems)
	}
	lines := make([]string, 0, reportMaxLoot+1)
	for i, l := range in.Loot {
		if i == reportMaxLoot {
			lines = append(lines, fmt.Sprintf(reportCopy.MoreLoot, len(in.Loot)-reportMaxLoot))
			break
		}
		lines = append(lines, fmt.Sprintf(reportCopy.LootLine, escapedWithin(collapse(l.Item), maxNameRunes, maxNameRunes*2),
			escapedWithin(collapse(l.Winner), maxNameRunes, maxNameRunes*2)))
	}
	return strings.Join(lines, "\n")
}

// killClock is a kill time as m:ss, or h:mm:ss from an hour up.
func killClock(d time.Duration) string {
	total := int(d.Round(time.Second) / time.Second)
	if total >= 3600 {
		return fmt.Sprintf("%d:%02d:%02d", total/3600, total%3600/60, total%60)
	}
	return fmt.Sprintf("%d:%02d", total/60, total%60)
}
