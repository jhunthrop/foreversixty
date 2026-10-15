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

// cell is one talent in the order LadderPointOrder searches: the spec's
// own tree first, then the other two in tree order, each top row down
// and left to right (LoadTalentTrees' own order).
type cell struct {
	tree int
	node TalentNode
}

// LadderPointOrder is the guide build's learning order: one point per
// level from 10 to 60, never removed, so the addon can spend a fresh
// character's points the way the site's own leveling bands do.
//
// Each level adds the first cell (own tree first, then tier, then
// column) where that level's LadderTalentString build has more ranks
// than the order so far and the point is legal - the site planner's own
// rules (web/src/lib/planner/rules.ts validateOrder): pointsPerTier*tier
// points already in the tree, the prerequisite at its rank, below max
// rank. Where the ladder is monotonic that is exactly its per-level
// diff. Where it is not (the own-tree spine reservation moves a rank
// that a character could only move with a respec), no such cell may
// exist, and the point goes instead to the first legal cell where the
// guide's level-60 targets have more ranks than the order so far.
//
// The order must end on the level-60 ladder build itself, with 51
// points; a level with no legal point, or an order that ends anywhere
// else, is an error naming the level - a data defect, not something to
// paper over.
func LadderPointOrder(activeTrees []TalentTree, targets map[int]int, ownTreeIndex int) ([]Point, error) {
	cells := searchOrder(activeTrees, ownTreeIndex)
	ranks := make([]map[int]int, len(activeTrees))
	inTree := make([]int, len(activeTrees))
	for i := range ranks {
		ranks[i] = make(map[int]int)
	}
	legal := func(c cell) bool {
		if inTree[c.tree] < pointsPerTier*c.node.Tier || ranks[c.tree][c.node.ID] >= c.node.MaxRank {
			return false
		}
		return c.node.PrereqTalentID == 0 || ranks[c.tree][c.node.PrereqTalentID] >= c.node.PrereqRank
	}
	order := make([]Point, 0, 51)
	for level := 10; level <= 60; level++ {
		ladder, err := TalentRanksFromString(activeTrees, LadderTalentString(activeTrees, targets, ownTreeIndex, level))
		if err != nil {
			return nil, err
		}
		pick, found := firstCell(cells, func(c cell) bool {
			return ladder[c.node.ID] > ranks[c.tree][c.node.ID] && legal(c)
		})
		if !found {
			pick, found = firstCell(cells, func(c cell) bool {
				return min(targets[c.node.ID], c.node.MaxRank) > ranks[c.tree][c.node.ID] && legal(c)
			})
		}
		if !found {
			return nil, fmt.Errorf("leveling: level %d has no legal point toward the guide build", level)
		}
		ranks[pick.tree][pick.node.ID]++
		inTree[pick.tree]++
		order = append(order, Point{Tab: pick.tree + 1, Tier: pick.node.Tier + 1, Column: pick.node.Column + 1})
	}
	final := LadderTalentString(activeTrees, targets, ownTreeIndex, 60)
	if got := rankString(activeTrees, ranks); got != final {
		return nil, fmt.Errorf("leveling: level 60 order ends on %s, not the guide build %s", got, final)
	}
	return order, nil
}

func searchOrder(trees []TalentTree, ownTreeIndex int) []cell {
	treeOrder := []int{ownTreeIndex}
	for i := range trees {
		if i != ownTreeIndex {
			treeOrder = append(treeOrder, i)
		}
	}
	var cells []cell
	for _, ti := range treeOrder {
		if ti < 0 || ti >= len(trees) {
			continue
		}
		for _, node := range trees[ti].Talents {
			cells = append(cells, cell{tree: ti, node: node})
		}
	}
	return cells
}

func firstCell(cells []cell, ok func(cell) bool) (cell, bool) {
	for _, c := range cells {
		if ok(c) {
			return c, true
		}
	}
	return cell{}, false
}

// rankString writes ranks in LadderTalentString's own positional form.
func rankString(trees []TalentTree, ranks []map[int]int) string {
	parts := make([]string, len(trees))
	for i, tree := range trees {
		var sb strings.Builder
		for _, node := range tree.Talents {
			sb.WriteByte(byte('0' + ranks[i][node.ID]))
		}
		parts[i] = sb.String()
	}
	return strings.Join(parts, "-")
}
