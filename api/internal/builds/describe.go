package builds

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// defaultClassColor is the site's gold, used when the class has no color -
// or when the build's tree version has no data at all.
const defaultClassColor = "#e5b955"

// Description is everything the page and the card need beyond the record
// itself. The spec calls all of it derived, never stored.
type Description struct {
	ClassName  string
	RaceName   string
	ClassColor string
	// Heading is "{Race} {Class}".
	Heading string
	// Title is the record's title, or Heading when it has none.
	Title string
	// Split is points per tree in client order; SplitText joins it with "/"
	// ("31/0/20" for a three-tree class).
	Split     []int
	SplitText string
	// Level is 10 + len(point_order) - 1, per the contract: the first point
	// is spent at level 10, so a build with no points reads as level 9.
	Level int
	// DPSLine is the build's simmed DPS as card text - "<mean> DPS
	// ±<margin>" - once WithSimDPS has found a done sim for this
	// build. Empty until then, and the card simply omits the line.
	DPSLine string
}

func Describe(data *trees.Data, b Build) Description {
	d := Description{
		ClassName:  fmt.Sprintf("Class %d", b.ClassID),
		RaceName:   fmt.Sprintf("Race %d", b.RaceID),
		ClassColor: defaultClassColor,
		Level:      10 + len(b.PointOrder) - 1,
	}
	if tb, ok := buildData(data, b.TreeVersion); ok {
		if c, ok := tb.Class(b.ClassID); ok {
			d.ClassName = c.Name
			if c.Color != "" {
				d.ClassColor = c.Color
			}
		}
		if r, ok := tb.Race(b.RaceID); ok {
			d.RaceName = r.Name
		}
		treeList := tb.Trees(b.ClassID)
		counts := make([]int, len(treeList))
		position := make(map[int]int, len(treeList))
		for i, tree := range treeList {
			position[tree.ID] = i
		}
		for _, id := range b.PointOrder {
			t, ok := tb.Talent(b.ClassID, id)
			if !ok {
				continue
			}
			if i, ok := position[t.TreeID]; ok {
				counts[i]++
			}
		}
		d.Split = counts
	}

	parts := make([]string, len(d.Split))
	for i, n := range d.Split {
		parts[i] = strconv.Itoa(n)
	}
	d.SplitText = strings.Join(parts, "/")
	if d.SplitText == "" {
		// No tree data for this record: one zero still reads as a split.
		d.SplitText = "0"
	}
	d.Heading = d.RaceName + " " + d.ClassName
	d.Title = b.Title
	if d.Title == "" {
		d.Title = d.Heading
	}
	return d
}

func buildData(data *trees.Data, version string) (*trees.Build, bool) {
	if data == nil {
		return nil, false
	}
	return data.Build(version)
}

// SimLookup is the part of sims.Store the build card needs: the newest
// simmed result for a build, if the planner ever ran and shared one.
// It is defined here rather than imported from sims, so that builds
// never imports sims - sims already imports builds - and is satisfied
// by *sims.Store without either package naming the other.
type SimLookup interface {
	ForBuild(ctx context.Context, buildID string) (simapi.SimResult, bool, error)
}

// WithSimDPS returns d with DPSLine set from the build's newest done
// sim, when sims has one. sims may be nil, for a caller with no lookup
// wired up; d is then returned unchanged. A lookup error is returned
// for the caller to log - the card must still render, without a DPS
// line, rather than fail.
func (d Description) WithSimDPS(ctx context.Context, buildID string, sims SimLookup) (Description, error) {
	if sims == nil {
		return d, nil
	}
	res, ok, err := sims.ForBuild(ctx, buildID)
	if err != nil {
		return d, fmt.Errorf("builds: sim lookup for %s: %w", buildID, err)
	}
	if !ok {
		return d, nil
	}
	d.DPSLine = formatDPSLine(res.DPS.Mean, res.DPS.Error)
	return d, nil
}

// formatDPSLine renders a build's simmed DPS as the card's one line of
// text: the mean, rounded, with a 95% margin (1.96 standard errors)
// after it. Thousands are comma-separated, the way the rest of the
// site formats a number.
func formatDPSLine(mean, stdErr float64) string {
	margin := math.Round(1.96 * stdErr)
	return fmt.Sprintf("%s DPS ±%s", withCommas(math.Round(mean)), withCommas(margin))
}

// withCommas renders a rounded number with a comma every three digits.
func withCommas(n float64) string {
	s := strconv.FormatInt(int64(n), 10)
	neg := strings.HasPrefix(s, "-")
	s = strings.TrimPrefix(s, "-")
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	if neg {
		s = "-" + s
	}
	return s
}
