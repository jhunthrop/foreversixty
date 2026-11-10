package card

import (
	"fmt"
	"image/color"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"golang.org/x/image/font"
)

// Boss is one pull of an encounter on a report card: the boss, whether it died, and how long
// the pull ran.
type Boss struct {
	Name       string
	Kill       bool
	DurationMS int64
}

// ReportInput is everything a report card shows: the pulls in order, the night's length and
// its deaths.
type ReportInput struct {
	Title      string
	Zone       string
	Bosses     []Boss
	DurationMS int64
	Deaths     int
}

// The report card's words.
const (
	copyKilled    = "Killed"
	copyWiped     = "Wiped"
	copyNoBosses  = "No boss fights yet"
	copyLogZone   = "Combat log"
	copyMoreBoss  = " and %d more"
	copyWipesOne  = "1 wipe"
	copyWipesMany = "%d wipes"
	copyKillTime  = "%s %s"
	copyKillAfter = "%s %s after %s"
	copyDeathsOne = "1 death"
	copyDeaths    = "%d deaths"
	reportLabelGp = 14 // between "Killed" and the names
)

// RenderReport draws the 1200x630 card an unfurled report link shows: under the seal and name,
// the report's title, the zone with the kill count and the night's length, the bosses killed
// and wiped on, and the fan-made line.
func RenderReport(in ReportInput) ([]byte, error) {
	f, err := newFaces()
	if err != nil {
		return nil, err
	}
	img, err := canvas(f)
	if err != nil {
		return nil, err
	}
	maxW := Width - 2*marginX
	y := Height - marginY - 34
	drawText(img, f.body, inkText, marginX, y, fit(f.body, copyReport+lineJoin+copyFanLine, maxW))

	rows := bossRows(f, in.Bosses, maxW)
	for i := len(rows) - 1; i >= 0; i-- {
		y -= 40
		x := marginX + rows[i].indent
		if rows[i].label != "" {
			x = drawText(img, f.label, rows[i].ink, x, y, rows[i].label) + reportLabelGp
		}
		drawText(img, f.body, inkCream2, x, y, rows[i].names)
	}

	y -= 56
	drawText(img, f.lead, inkCream2, marginX, y, fit(f.lead, summaryLine(in), maxW))
	y -= 70
	drawTitle(img, f, marginX, y, reportTitle(in.Title), maxW)
	return encode(img)
}

func reportTitle(title string) string {
	if t := strings.TrimSpace(title); t != "" {
		return t
	}
	return copyReport
}

// summaryLine is the zone, the kill count and the length of the night: "Molten Core · 8 of 10
// killed · 2h 05m".
func summaryLine(in ReportInput) string {
	zone := strings.TrimSpace(in.Zone)
	if zone == "" {
		zone = copyLogZone
	}
	killed, wiped := outcomes(in.Bosses)
	parts := []string{zone, fmt.Sprintf("%d of %d killed", len(killed), len(killed)+len(wiped))}
	if in.DurationMS > 0 {
		parts = append(parts, notify.NightLength(time.Duration(in.DurationMS)*time.Millisecond))
	}
	switch {
	case in.Deaths == 1:
		parts = append(parts, copyDeathsOne)
	case in.Deaths > 1:
		parts = append(parts, fmt.Sprintf(copyDeaths, in.Deaths))
	}
	return strings.Join(parts, lineJoin)
}

// bossRow is one line of bosses: the label in its outcome colour, then the names.
type bossRow struct {
	label  string
	ink    color.NRGBA
	names  string
	indent int // a continuation line sits under the names, past the label
}

// reportBossLines is how many lines the bosses take, kills and wipes together.
const reportBossLines = 3

// bossRows are the kills, then the wipes, a line each, the names in pull order and cut with
// "and N more" where the line runs out. A report with no fights says so.
func bossRows(f *faces, bosses []Boss, maxWidth int) []bossRow {
	if len(bosses) == 0 {
		return []bossRow{{names: copyNoBosses}}
	}
	killed, wiped := outcomes(bosses)
	var rows []bossRow
	killedWords, wipedWords := make([]string, len(killed)), make([]string, len(wiped))
	for i, b := range killed {
		killedWords[i] = fmt.Sprintf(copyKillTime, b.name, notify.KillClock(b.kill))
		if b.wipes > 0 {
			killedWords[i] = fmt.Sprintf(copyKillAfter, b.name, notify.KillClock(b.kill), wipeWords(b.wipes))
		}
	}
	for i, b := range wiped {
		wipedWords[i] = fmt.Sprintf("%s (%s)", b.name, wipeWords(b.wipes))
	}
	for _, g := range []struct {
		label string
		ink   color.NRGBA
		names []string
	}{{copyKilled, inkKill, killedWords}, {copyWiped, inkWipe, wipedWords}} {
		if len(g.names) == 0 {
			continue
		}
		indent := measure(f.label, g.label) + reportLabelGp
		most := 1
		if g.label == copyKilled {
			most = reportBossLines - min(len(wipedWords), 1)
		}
		for i, line := range nameLines(f.body, g.names, maxWidth-indent, most) {
			row := bossRow{label: g.label, ink: g.ink, names: line}
			if i > 0 {
				row = bossRow{names: line, indent: indent}
			}
			rows = append(rows, row)
		}
	}
	return rows
}

// nameLines lays names out over at most most lines of maxWidth, joined with " · ", the last
// line ending in " and N more" when some do not fit.
func nameLines(face font.Face, names []string, maxWidth, most int) []string {
	var lines []string
	for len(names) > 0 {
		if len(lines) == most-1 {
			return append(lines, nameList(face, names, maxWidth))
		}
		line, used := names[0], 1
		for used < len(names) && measure(face, line+lineJoin+names[used]) <= maxWidth {
			line += lineJoin + names[used]
			used++
		}
		lines = append(lines, line)
		names = names[used:]
	}
	return lines
}

// wipeWords is "1 wipe" or "3 wipes".
func wipeWords(n int) string {
	if n == 1 {
		return copyWipesOne
	}
	return fmt.Sprintf(copyWipesMany, n)
}

// bossNight is one boss over the night: its kill time, zero when it never died, and the wipes.
type bossNight struct {
	name  string
	kill  time.Duration
	wipes int
}

// outcomes are the night's bosses by name in pull order: those killed with the killing pull's
// time, and those wiped on and never killed with their wipe count. A boss wiped on and then
// killed is a kill.
func outcomes(bosses []Boss) (killed, wiped []bossNight) {
	var order []string
	nights := map[string]*bossNight{}
	for _, b := range bosses {
		n, seen := nights[b.Name]
		if !seen {
			n = &bossNight{name: b.Name}
			nights[b.Name] = n
			order = append(order, b.Name)
		}
		if b.Kill {
			n.kill = max(time.Duration(b.DurationMS)*time.Millisecond, time.Second)
		} else {
			n.wipes++
		}
	}
	for _, name := range order {
		if n := nights[name]; n.kill > 0 {
			killed = append(killed, *n)
		} else {
			wiped = append(wiped, *n)
		}
	}
	return killed, wiped
}

// nameList joins names with " · " within maxWidth, ending in " and N more" when some do not fit.
func nameList(face font.Face, names []string, maxWidth int) string {
	line := ""
	for i, n := range names {
		next := n
		if line != "" {
			next = line + lineJoin + n
		}
		rest := len(names) - i - 1
		tail := ""
		if rest > 0 {
			tail = fmt.Sprintf(copyMoreBoss, rest)
		}
		if measure(face, next+tail) <= maxWidth || (line == "" && measure(face, next) <= maxWidth) {
			line = next
			continue
		}
		if line == "" {
			return fit(face, n, maxWidth)
		}
		return line + fmt.Sprintf(copyMoreBoss, len(names)-i)
	}
	return line
}
