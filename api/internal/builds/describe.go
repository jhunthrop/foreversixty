package builds

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
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
