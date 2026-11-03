package talents

import (
	"fmt"
	"strings"
)

// pointsPerTier is the game's row gate: a talent in row t (0-based) opens
// once its tree holds pointsPerTier*t points in the rows above it
// (web/src/lib/planner/types.ts POINTS_PER_TIER, sim/leveling's own copy).
const pointsPerTier = 5

// firstTalentLevel is the level of a character's first talent point;
// every level from there gives one more, up to maxTalentPoints.
const (
	firstTalentLevel = 10
	maxTalentPoints  = 51
)

// PointsAt is how many talent points a character of level has.
func PointsAt(level int) int {
	return min(max(level-firstTalentLevel+1, 0), maxTalentPoints)
}

// Check refuses a positional talent string the game would not let a
// character of level hold: a malformed string, a rank past a talent's
// last, more points than the level gives, a row whose gate is not met
// or a prerequisite unmet. The check is on the finished tree; a tree
// that passes has a legal spending order (every talent's row gate and
// prerequisite hold once the points above and before it are in).
//
// Every refusal wraps ErrIllegalBuild, and says which talent and rule,
// in words a player can act on.
func (l *Layouts) Check(classSlug, talents string, level int) error {
	lay, ok := l.byClass[classSlug]
	if !ok {
		return fmt.Errorf("%w: %q", ErrUnknownClass, classSlug)
	}
	ranks, err := lay.ranks(talents)
	if err != nil {
		return err
	}
	if spent, have := total(ranks), PointsAt(level); spent > have {
		return fmt.Errorf("%w: the build spends %d points; a level %d character has %d", ErrIllegalBuild, spent, level, have)
	}
	for ti, t := range lay.trees {
		if err := t.gatesHold(ranks[ti], lay, ranks); err != nil {
			return err
		}
	}
	return nil
}

// ranks reads the string into one rank per talent slot, refusing a
// shape the layout cannot hold or a rank past a talent's last.
func (lay layout) ranks(s string) ([][]int, error) {
	parts := strings.Split(s, "-")
	if len(parts) > len(lay.trees) {
		return nil, fmt.Errorf("%w: %q has %d trees; the class has %d", ErrIllegalBuild, s, len(parts), len(lay.trees))
	}
	out := make([][]int, len(lay.trees))
	for ti, t := range lay.trees {
		out[ti] = make([]int, len(t.nodes))
		if ti >= len(parts) {
			continue
		}
		if len(parts[ti]) > len(t.nodes) {
			return nil, fmt.Errorf("%w: %q gives %s %d digits; %s has %d talents", ErrIllegalBuild, s, t.name, len(parts[ti]), t.name, len(t.nodes))
		}
		for si, c := range parts[ti] {
			if c < '0' || c > '9' {
				return nil, fmt.Errorf("%w: %q has %q, not a digit", ErrIllegalBuild, s, c)
			}
			n := t.nodes[si]
			if r := int(c - '0'); r > n.maxRank {
				return nil, fmt.Errorf("%w: %s has %d of %d points", ErrIllegalBuild, n.name, r, n.maxRank)
			}
			out[ti][si] = int(c - '0')
		}
	}
	return out, nil
}

// gatesHold checks every ranked talent of one tree against its row gate
// and its prerequisite (looked up across the class, as the client does).
func (t tree) gatesHold(treeRanks []int, lay layout, all [][]int) error {
	for si, n := range t.nodes {
		if treeRanks[si] == 0 {
			continue
		}
		if need, above := pointsPerTier*n.tier, pointsAbove(t, treeRanks, n.tier); above < need {
			return fmt.Errorf("%w: %s needs %d points in %s above its row; the build has %d", ErrIllegalBuild, n.name, need, t.name, above)
		}
		if n.prereqID == 0 {
			continue
		}
		at, ok := lay.index[n.prereqID]
		if !ok {
			return fmt.Errorf("%w: %d (prerequisite of %s) in this class", ErrUnknownTalent, n.prereqID, n.name)
		}
		if got := all[at.tree][at.slot]; got < n.prereqRank {
			pre := lay.trees[at.tree].nodes[at.slot]
			return fmt.Errorf("%w: %s needs %d points in %s; the build has %d", ErrIllegalBuild, n.name, n.prereqRank, pre.name, got)
		}
	}
	return nil
}

// pointsAbove is the tree's points in rows above tier.
func pointsAbove(t tree, treeRanks []int, tier int) int {
	n := 0
	for si, node := range t.nodes {
		if node.tier < tier {
			n += treeRanks[si]
		}
	}
	return n
}

func total(ranks [][]int) int {
	n := 0
	for _, tree := range ranks {
		for _, r := range tree {
			n += r
		}
	}
	return n
}
