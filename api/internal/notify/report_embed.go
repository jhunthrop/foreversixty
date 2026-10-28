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
}

// ReportKill is one boss killed on the night.
type ReportKill struct {
	Boss     string
	Duration time.Duration
}

// ReportReady is what an officer's "Post to Discord" on a done raid holds: counts only, never
// the names of who came or who was missing (R-A3). RaidPosted carries the guild and the raid;
// its RaidURL should be the raid page, the embed adds the report anchor.
type ReportReady struct {
	RaidPosted
	Kills        []ReportKill
	Wipes        int
	Came, Signed int
	LootItems    int
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
		raidField(reportCopy.CameName, fmt.Sprintf(reportCopy.CameOf, in.Came, in.Signed)),
		raidField(reportCopy.LootName, lootValue(in.LootItems)),
	}
	return e
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

func lootValue(items int) string {
	if items == 1 {
		return reportCopy.LootOne
	}
	return fmt.Sprintf(reportCopy.LootMany, items)
}

// killClock is a kill time as m:ss, or h:mm:ss from an hour up.
func killClock(d time.Duration) string {
	total := int(d.Round(time.Second) / time.Second)
	if total >= 3600 {
		return fmt.Sprintf("%d:%02d:%02d", total/3600, total%3600/60, total%60)
	}
	return fmt.Sprintf("%d:%02d", total/60, total%60)
}
