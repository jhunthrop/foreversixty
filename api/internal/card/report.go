package card

import (
	"fmt"
	"image/color"
	"strings"

	"golang.org/x/image/font"
)

// Boss is one encounter on a report card.
type Boss struct {
	Name string
	Kill bool
}

// ReportInput is everything a report card shows.
type ReportInput struct {
	Title      string
	Zone       string
	Bosses     []Boss
	DurationMS int64
}

// The report card's words.
const (
	copyKilled    = "Killed"
	copyWiped     = "Wiped"
	copyNoBosses  = "No boss fights yet"
	copyLogZone   = "Combat log"
	copyMoreBoss  = " and %d more"
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
		y -= 42
		x := marginX
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
		parts = append(parts, duration(in.DurationMS))
	}
	return strings.Join(parts, lineJoin)
}

// duration renders milliseconds as hours and minutes.
func duration(ms int64) string {
	minutes := ms / 60000
	if minutes < 60 {
		return fmt.Sprintf("%d min", minutes)
	}
	return fmt.Sprintf("%dh %02dm", minutes/60, minutes%60)
}

// bossRow is one line of bosses: the label in its outcome colour, then the names.
type bossRow struct {
	label string
	ink   color.NRGBA
	names string
}

// bossRows are the kills, then the wipes, a line each, the names in pull order and cut with
// "and N more" where the line runs out. A report with no fights says so.
func bossRows(f *faces, bosses []Boss, maxWidth int) []bossRow {
	if len(bosses) == 0 {
		return []bossRow{{names: copyNoBosses}}
	}
	killed, wiped := outcomes(bosses)
	var rows []bossRow
	for _, g := range []struct {
		label string
		ink   color.NRGBA
		names []string
	}{{copyKilled, inkKill, killed}, {copyWiped, inkWipe, wiped}} {
		if len(g.names) == 0 {
			continue
		}
		room := maxWidth - measure(f.label, g.label) - reportLabelGp
		rows = append(rows, bossRow{label: g.label, ink: g.ink, names: nameList(f.body, g.names, room)})
	}
	return rows
}

// outcomes are the night's bosses by name in pull order: those killed, and those wiped on and
// never killed. A boss wiped on and then killed is a kill.
func outcomes(bosses []Boss) (killed, wiped []string) {
	var order []string
	kill := map[string]bool{}
	for _, b := range bosses {
		if _, seen := kill[b.Name]; !seen {
			order = append(order, b.Name)
		}
		kill[b.Name] = kill[b.Name] || b.Kill
	}
	for _, name := range order {
		if kill[name] {
			killed = append(killed, name)
		} else {
			wiped = append(wiped, name)
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
