package main

import "strings"

// truncateTalents drops points from a level-60 talent build, expressed
// as the engine's own three per-tree digit strings (one digit per
// talent slot, in the tree's on-screen tier/column order - see
// sim/talents' package doc), until at most budget points remain.
//
// The rule is the one this lane's brief states, deliberately simpler
// than the rotation-accuracy ladder's own "spend points bottom-up,
// this spec's tree first" approximation (that rule builds a leveling
// build UP from nothing; the ladder lane owns it and it is still
// unwritten on main as of this lane's run - see the brief's Inputs
// section): dropping talents from the LAST tree first, and within a
// tree from its LAST slot first, until level-9 points remain. This is
// a stated approximation of a leveling build, not a claim about what
// any real leveling hunter would actually choose - it exists so the
// prototype's bands 20/30/40 have SOME talents to weigh gear against,
// and a reader can see exactly which points were kept.
//
// It returns new strings; the input is never mutated (trees' own
// backing arrays are copied before decrementing).
func truncateTalents(trees []string, budget int) []string {
	points := make([][]int, len(trees))
	total := 0
	for i, s := range trees {
		points[i] = make([]int, len(s))
		for j, c := range s {
			n := int(c - '0')
			points[i][j] = n
			total += n
		}
	}
	for ti := len(points) - 1; ti >= 0 && total > budget; ti-- {
		for pi := len(points[ti]) - 1; pi >= 0 && total > budget; pi-- {
			for points[ti][pi] > 0 && total > budget {
				points[ti][pi]--
				total--
			}
		}
	}
	out := make([]string, len(points))
	for i, tree := range points {
		var b strings.Builder
		for _, n := range tree {
			b.WriteByte(byte('0' + n))
		}
		out[i] = strings.TrimRight(b.String(), "0")
	}
	return out
}

// talentString joins truncated (or full) per-tree digit strings into
// the engine's positional talent string: trees separated by "-".
func talentString(trees []string) string {
	return strings.Join(trees, "-")
}

// levelBudget is the ladder's own rule (Phase 1a of the rotation
// accuracy program design, which this lane's brief points at): a
// leveling character has spent level-9 talent points by level, one
// point per level from the first point at 10. Below 10 that is
// negative, which truncateTalents treats as "drop everything" (total
// > budget stays true until every digit is 0); no band this prototype
// runs is below 20, but the formula is written the same way the
// ladder states it rather than clamped, so a caller that does pass a
// low level gets the honest zero-point build instead of a silently
// different rule.
func levelBudget(level int) int { return level - 9 }
