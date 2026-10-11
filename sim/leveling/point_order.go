package leveling

import (
	"fmt"
	"strings"
)

// Point is one talent point as the addon's codec writes it: the tab
// (tree position + 1), the talent's tier and its column, all 1-based.
type Point struct {
	Tab, Tier, Column int
}

// LadderPointOrder is the guide build's learning order: the one point
// each level from 10 to 60 adds to LadderTalentString's truncation, in
// level order, so the addon can spend a fresh character's points the way
// the site's own leveling bands do. A level whose truncation does not
// differ from the level below by exactly one point (a guide build that
// spends fewer than 51 points, or a spine reservation that moves points
// rather than adding one) is an error naming that level: such a ladder
// has no single-point order to give.
func LadderPointOrder(activeTrees []TalentTree, targets map[int]int, ownTreeIndex int) ([]Point, error) {
	order := make([]Point, 0, 51)
	prev := strings.Split(LadderTalentString(activeTrees, targets, ownTreeIndex, 9), "-")
	for level := 10; level <= 60; level++ {
		cur := strings.Split(LadderTalentString(activeTrees, targets, ownTreeIndex, level), "-")
		var added []Point
		valid := true
		for ti := range cur {
			for j := range cur[ti] {
				switch d := int(cur[ti][j]) - int(prev[ti][j]); {
				case d == 1:
					node := activeTrees[ti].Talents[j]
					added = append(added, Point{Tab: ti + 1, Tier: node.Tier + 1, Column: node.Column + 1})
				case d != 0:
					valid = false
				}
			}
		}
		if !valid || len(added) != 1 {
			return nil, fmt.Errorf("leveling: level %d is not exactly one point more than level %d (%s -> %s)",
				level, level-1, strings.Join(prev, "-"), strings.Join(cur, "-"))
		}
		order = append(order, added[0])
		prev = cur
	}
	return order, nil
}
