package character

import "github.com/jhunthrop/foreversixty/sim/specs"

// The sources SpecFor names a spec by.
const (
	SpecFromFight   = "fight"
	SpecFromTalents = "talents"
	SpecFromDefault = "default"
)

// SpecFor is the spec the site treats a character as, and where it came from.
//
//  1. fightSpec, the site's slug ("warrior-fury") when a logged fight named one: "fight".
//  2. Otherwise the spec of the tree with the most points in treeRanks (the export's points
//     per tree in the client's tree order; a tie goes to the first tree, as spec.Inferrer.Name
//     does): "talents".
//  3. Otherwise, with no points spent, the class's first spec in the data's order (the
//     planner's default for the class): "default".
//
// class is the class slug ("warrior"). A class the data does not know has no spec: ("", "").
// Where a tree has more than one spec (a feral druid's cat and bear), the first listed wins.
func SpecFor(class, fightSpec string, treeRanks []int) (slug, source string) {
	if fightSpec != "" {
		return fightSpec, SpecFromFight
	}
	best, most := -1, 0
	for n, points := range treeRanks {
		if points > most {
			best, most = n, points
		}
	}
	var first string
	for _, s := range specs.All {
		if s.ClassSlug != class {
			continue
		}
		if first == "" {
			first = s.Spec
		}
		if best >= 0 && s.TreeIndex == best {
			return s.Spec, SpecFromTalents
		}
	}
	if first == "" {
		return "", ""
	}
	return first, SpecFromDefault
}
