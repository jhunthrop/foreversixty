package card

import (
	"fmt"
	"strings"

	"golang.org/x/image/font"
)

// Report card layout. The build card's baselines are tuned for four
// short lines; a report card carries a zone, a title, a row of bosses,
// and a duration, so it has its own.
const (
	zoneBaseline     = 128
	reportTitleLine1 = 208
	reportTitleLine2 = 268
	bossesBaseline   = 380
	bossesBaseline2  = 432
	durationBaseline = 510
)

// maxBossLines is how many lines of boss names a card shows before it
// gives up and counts the rest.
const maxBossLines = 2

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
	// Accent colours the top band; the site's gold when empty.
	Accent string
}

// RenderReport draws the 1200x630 card an unfurled report link shows:
// the zone, the report's title, the bosses with a kill or wipe mark,
// and how long the night ran.
func RenderReport(in ReportInput) ([]byte, error) {
	f, err := newFaces()
	if err != nil {
		return nil, err
	}
	accent := parseHexColor(in.Accent)
	img := canvas(accent)

	zone := in.Zone
	if zone == "" {
		zone = "Combat log"
	}
	drawText(img, f.body, mutedColor, marginX, zoneBaseline, fit(f.body, strings.ToUpper(zone), Width-2*marginX))

	title := strings.TrimSpace(in.Title)
	if title == "" {
		title = "Raid report"
	}
	first, second := wrap(f.title, title, Width-2*marginX)
	drawText(img, f.title, goldColor, marginX, reportTitleLine1, first)
	if second != "" {
		drawText(img, f.title, goldColor, marginX, reportTitleLine2, second)
	}

	for i, line := range bossLines(f, in.Bosses, Width-2*marginX) {
		baseline := bossesBaseline
		if i == 1 {
			baseline = bossesBaseline2
		}
		drawText(img, f.body, textColor, marginX, baseline, line)
	}

	drawText(img, f.body, mutedColor, marginX, durationBaseline, summaryLine(in))
	drawText(img, f.wordmark, mutedColor, marginX, wordmarkBaseline, "foreversixty.gg")
	return encode(img)
}

// summaryLine is the count of kills and the length of the night.
func summaryLine(in ReportInput) string {
	kills := 0
	for _, b := range in.Bosses {
		if b.Kill {
			kills++
		}
	}
	parts := []string{fmt.Sprintf("%d of %d killed", kills, len(in.Bosses))}
	if in.DurationMS > 0 {
		parts = append(parts, duration(in.DurationMS))
	}
	return strings.Join(parts, "  -  ")
}

// duration renders milliseconds as hours and minutes.
func duration(ms int64) string {
	minutes := ms / 60000
	if minutes < 60 {
		return fmt.Sprintf("%d min", minutes)
	}
	return fmt.Sprintf("%dh %02dm", minutes/60, minutes%60)
}

// bossLines lays the bosses out over at most maxBossLines lines, each
// name marked with a kill or a wipe. The marks are plain ASCII: the
// vendored faces are the latin subsets and carry no check mark.
func bossLines(f *faces, bosses []Boss, maxWidth int) []string {
	if len(bosses) == 0 {
		return []string{"No boss fights yet"}
	}
	var (
		lines   []string
		current string
	)
	for i, b := range bosses {
		mark := "x"
		if b.Kill {
			mark = "+"
		}
		piece := b.Name + " " + mark
		next := piece
		if current != "" {
			next = current + "   " + piece
		}
		if measure(f.body, next) <= maxWidth {
			current = next
			continue
		}
		lines = append(lines, current)
		current = piece
		if len(lines) == maxBossLines {
			remaining := len(bosses) - i
			lines[maxBossLines-1] = fit(f.body,
				lines[maxBossLines-1]+fmt.Sprintf("   and %d more", remaining), maxWidth)
			return lines
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	if len(lines) > maxBossLines {
		lines = lines[:maxBossLines]
	}
	return lines
}

// wrap splits s over two lines of the given face, ellipsising whatever
// will not fit.
func wrap(face font.Face, s string, maxWidth int) (string, string) {
	if measure(face, s) <= maxWidth {
		return s, ""
	}
	words := strings.Fields(s)
	first := ""
	i := 0
	for ; i < len(words); i++ {
		candidate := strings.TrimSpace(first + " " + words[i])
		if measure(face, candidate) > maxWidth {
			break
		}
		first = candidate
	}
	if first == "" {
		return fit(face, s, maxWidth), ""
	}
	return first, fit(face, strings.Join(words[i:], " "), maxWidth)
}
