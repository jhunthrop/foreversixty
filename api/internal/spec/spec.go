// Package spec names a build from its talents: the three-number point
// split a ranking row shows, and the spec name it is bracketed under.
//
// The Phase 3 contract wants the names from the per-class table in
// data/curated/specs.json, which the data plan has yet to produce.
// Until it does, the spec is the tree with the most points, named by
// the tree's own name - "Fury", "Holy", "Subtlety" - which is what a
// reader expects for all but the hybrid builds the curated table will
// name properly.
//
// FOLLOW-UP (data): when data/curated/specs.json ships, load it beside
// the talent data and replace Name below. Nothing else in the API
// looks at trees to decide a spec.
package spec

import (
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// MaxPoints is the most points a level-60 build can spend, and the
// ceiling that tells a points-per-tree split from a list of talent ids.
const MaxPoints = 51

// Inferrer names specs from one client build's talent data.
type Inferrer struct {
	build   *trees.Build
	byClass map[string]int
}

// New returns an inferrer over a build's trees, or nil when there is no
// tree data: every method on a nil inferrer answers empty, so a
// deployment with no talent data still ranks, just without spec names.
func New(b *trees.Build) *Inferrer {
	if b == nil {
		return nil
	}
	byClass := map[string]int{}
	for _, c := range b.Classes() {
		byClass[strings.ToLower(c.Name)] = c.ID
		byClass[strings.ToLower(c.Slug)] = c.ID
	}
	return &Inferrer{build: b, byClass: byClass}
}

// ClassID resolves a class name as a combat log spells it ("Warrior")
// to the id the talent data is keyed by.
func (i *Inferrer) ClassID(class string) (int, bool) {
	if i == nil {
		return 0, false
	}
	id, ok := i.byClass[strings.ToLower(strings.TrimSpace(class))]
	return id, ok
}

// Split is the points per tree in client order, from a COMBATANT_INFO
// talents field.
//
// The field's shape is settled by the first beta log: retail writes a
// list of talent ids, and a Classic-style client may write the three
// tree totals instead. Both are read here, and this is the only place
// that decides which is which.
func (i *Inferrer) Split(class string, talents []int64) []int {
	classID, ok := i.ClassID(class)
	if !ok {
		return nil
	}
	trees := i.build.Trees(classID)
	if len(trees) == 0 {
		return nil
	}
	if split, ok := asTotals(talents, len(trees)); ok {
		return split
	}
	split := make([]int, len(trees))
	position := map[int]int{}
	for n, t := range trees {
		position[t.ID] = n
	}
	for _, id := range talents {
		ref, ok := i.build.Talent(classID, int(id))
		if !ok {
			// The 1.60 client writes the talent's spell id, not the trait
			// node id the emitted data is keyed by. Both are tried, in that
			// order, so a log from either client reads correctly.
			ref, ok = i.build.TalentBySpellID(classID, int(id))
		}
		if !ok {
			continue
		}
		if n, ok := position[ref.TreeID]; ok {
			split[n]++
		}
	}
	return split
}

// asTotals reads a talents field that is already the per-tree totals:
// one entry per tree, each within the point budget, summing to no more
// than it. A list of talent ids never looks like that, because talent
// ids are in the thousands.
func asTotals(talents []int64, trees int) ([]int, bool) {
	if len(talents) != trees {
		return nil, false
	}
	out := make([]int, 0, trees)
	sum := 0
	for _, v := range talents {
		if v < 0 || v > MaxPoints {
			return nil, false
		}
		sum += int(v)
		out = append(out, int(v))
	}
	if sum == 0 || sum > MaxPoints {
		return nil, false
	}
	return out, true
}

// SplitString renders a split the way a ranking row shows it: "31/20/0".
func SplitString(split []int) string {
	if len(split) == 0 {
		return ""
	}
	parts := make([]string, len(split))
	for i, n := range split {
		parts[i] = strconv.Itoa(n)
	}
	return strings.Join(parts, "/")
}

// Name is the spec a split is bracketed under: the tree with the most
// points. A tie goes to the earlier tree, which is the client's own
// order, so the answer never depends on map iteration.
func (i *Inferrer) Name(class string, split []int) string {
	classID, ok := i.ClassID(class)
	if !ok || len(split) == 0 {
		return ""
	}
	trees := i.build.Trees(classID)
	if len(trees) != len(split) {
		return ""
	}
	best, most := -1, 0
	for n, points := range split {
		if points > most {
			best, most = n, points
		}
	}
	if best < 0 {
		return ""
	}
	return trees[best].Name
}

// Of is Split and Name together, which is what a ranking row needs.
func (i *Inferrer) Of(class string, talents []int64) (name, split string) {
	points := i.Split(class, talents)
	return i.Name(class, points), SplitString(points)
}
