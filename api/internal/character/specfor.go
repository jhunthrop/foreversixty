package character

import "github.com/jhunthrop/foreversixty/sim/specs"

// The sources SpecFor names a spec by.
const (
	SpecFromFight   = "fight"
	SpecFromTalents = "talents"
	// SpecFromAsk is points spent that do not name a tree (DecideTree): the addon asks the
	// player. The spec SpecFor returns with it is the class's default, a placeholder for
	// callers that need one; a surface that names the spec to a person names none.
	SpecFromAsk     = "ask"
	SpecFromDefault = "default"
)

// SureLead is how far the leading tree must lead every other before the points name it, and
// SureMinPoints the fewest spent points that can name a tree at all: the addon's
// Talents.SURE_LEAD and SURE_MIN_POINTS (owner 2026-11-11, "a wrong silent guess is worse than
// one question"). web/src/lib/character-spec.ts holds the TypeScript copy; the three tests read
// the same cases.
const (
	SureLead      = 5
	SureMinPoints = 2
)

// DecideTree is the tree index the points name beyond doubt, as the addon's Talents.decideTree
// reads them: at least SureMinPoints spent and all in one tree, or the leading tree holding
// more than half of the points and leading every other by SureLead. ok is false otherwise (a
// tie, a close split, nothing spent), where the addon asks. spent is the points' total.
func DecideTree(points []int) (tree int, spent int, ok bool) {
	lead, top, second := -1, 0, 0
	for n, p := range points {
		if p <= 0 {
			continue
		}
		spent += p
		switch {
		case p > top:
			lead, top, second = n, p, top
		case p > second:
			second = p
		}
	}
	if spent < SureMinPoints {
		return -1, spent, false
	}
	if second == 0 || (top*2 > spent && top-second >= SureLead) {
		return lead, spent, true
	}
	return -1, spent, false
}

// SpecFor is the spec the site treats a character as, and where it came from.
//
//  1. fightSpec, the site's slug ("warrior-fury") when a logged fight named one: "fight".
//  2. Otherwise the spec of the tree treeRanks names (DecideTree; treeRanks is the export's
//     points per tree in the client's tree order): "talents".
//  3. Otherwise, with points spent that name no tree, the class's first spec: "ask".
//  4. Otherwise, with no points spent, the class's first spec in the data's order (the
//     planner's default for the class): "default".
//
// class is the class slug ("warrior"). A class the data does not know has no spec: ("", "").
// Where a tree has more than one spec (a feral druid's cat and bear), the first listed wins.
func SpecFor(class, fightSpec string, treeRanks []int) (slug, source string) {
	if fightSpec != "" {
		return fightSpec, SpecFromFight
	}
	best, spent, ok := DecideTree(treeRanks)
	var first string
	for _, s := range specs.All {
		if s.ClassSlug != class {
			continue
		}
		if first == "" {
			first = s.Spec
		}
		if ok && s.TreeIndex == best {
			return s.Spec, SpecFromTalents
		}
	}
	switch {
	case first == "":
		return "", ""
	case spent > 0:
		return first, SpecFromAsk
	default:
		return first, SpecFromDefault
	}
}
