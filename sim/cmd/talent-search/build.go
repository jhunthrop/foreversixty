package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// pointsPerTier is the planner's own rule (web/src/lib/planner/types.ts
// POINTS_PER_TIER): tier t of a tree unlocks once 5*t points sit in it.
const pointsPerTier = 5

// build is a talent build by stable talent id: id -> points. A build is
// never changed in place; with returns a new one.
type build map[int]int

// with is a copy of b with delta points added to id (a talent left at
// zero is dropped, so two equal builds always compare equal by key).
func (b build) with(id, delta int) build {
	out := make(build, len(b)+1)
	for k, v := range b {
		out[k] = v
	}
	out[id] += delta
	if out[id] == 0 {
		delete(out, id)
	}
	return out
}

func (b build) points() int {
	n := 0
	for _, v := range b {
		n += v
	}
	return n
}

// talentTrees is one class's active-build trees plus an id index.
type talentTrees struct {
	trees  []leveling.TalentTree
	treeOf map[int]int
	byID   map[int]leveling.TalentNode
}

func newTalentTrees(trees []leveling.TalentTree) talentTrees {
	t := talentTrees{trees: trees, treeOf: map[int]int{}, byID: map[int]leveling.TalentNode{}}
	for ti, tree := range trees {
		for _, n := range tree.Talents {
			t.treeOf[n.ID] = ti
			t.byID[n.ID] = n
		}
	}
	return t
}

// treePoints is the points b spends in each tree.
func (t talentTrees) treePoints(b build) []int {
	out := make([]int, len(t.trees))
	for id, r := range b {
		out[t.treeOf[id]] += r
	}
	return out
}

// pointsBelow is the points b spends in tree ti on tiers below tier.
func (t talentTrees) pointsBelow(b build, ti, tier int) int {
	n := 0
	for _, node := range t.trees[ti].Talents {
		if node.Tier < tier {
			n += b[node.ID]
		}
	}
	return n
}

// legal reports why b is not a build a player could actually spend:
// an unknown talent, ranks above max, a tier not yet unlocked by the
// points beneath it, an unmet prerequisite, or a total other than
// total. nil means legal. Order-free: a final rank map is reachable
// exactly when every taken talent's gate is met by the points in lower
// tiers, because points can always be spent tier by tier.
func (t talentTrees) legal(b build, total int) error {
	ids := make([]int, 0, len(b))
	for id := range b {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	for _, id := range ids {
		r := b[id]
		node, ok := t.byID[id]
		if !ok {
			return fmt.Errorf("talent %d is not in this class", id)
		}
		if r < 0 || r > node.MaxRank {
			return fmt.Errorf("%s has %d of %d points", node.Name, r, node.MaxRank)
		}
		if r == 0 {
			continue
		}
		if below := t.pointsBelow(b, t.treeOf[id], node.Tier); below < pointsPerTier*node.Tier {
			return fmt.Errorf("%s (tier %d) needs %d points above it, has %d", node.Name, node.Tier, pointsPerTier*node.Tier, below)
		}
		if node.PrereqTalentID != 0 && b[node.PrereqTalentID] < node.PrereqRank {
			return fmt.Errorf("%s needs %d points in %s", node.Name, node.PrereqRank, t.byID[node.PrereqTalentID].Name)
		}
	}
	if p := b.points(); p != total {
		return fmt.Errorf("spends %d points, want %d", p, total)
	}
	return nil
}

// canAdd reports whether one more point in id keeps b's gates met
// (the total is the caller's concern).
func (t talentTrees) canAdd(b build, id int) bool {
	node, ok := t.byID[id]
	if !ok || b[id] >= node.MaxRank {
		return false
	}
	if t.pointsBelow(b, t.treeOf[id], node.Tier) < pointsPerTier*node.Tier {
		return false
	}
	return node.PrereqTalentID == 0 || b[node.PrereqTalentID] >= node.PrereqRank
}

// activeDigits is b in the active build's own positional order, one
// string per tree, trailing zeros kept.
func (t talentTrees) activeDigits(b build) []string {
	out := make([]string, len(t.trees))
	for ti, tree := range t.trees {
		var sb strings.Builder
		for _, n := range tree.Talents {
			sb.WriteByte(byte('0' + b[n.ID]))
		}
		out[ti] = sb.String()
	}
	return out
}

// key is a build's identity for dedupe.
func (t talentTrees) key(b build) string { return strings.Join(t.activeDigits(b), "-") }

// fs1 is the build's FS1 code against the active client build (the
// format the guides' frontmatter carries: trailing zeros trimmed per
// tree, an empty tree written "0").
func (t talentTrees) fs1(clientBuild, class, race string, b build) string {
	digits := t.activeDigits(b)
	for i, d := range digits {
		d = strings.TrimRight(d, "0")
		if d == "" {
			d = "0"
		}
		digits[i] = d
	}
	return fmt.Sprintf("FS1:%s:%s:%s:%s:", clientBuild, class, race, strings.Join(digits, "/"))
}

// decodeActive reads an active-layout positional string ("a-b-c", as
// leveling.LadderTalentString writes it) back into a build by id.
func (t talentTrees) decodeActive(s string) (build, error) {
	parts := strings.Split(s, "-")
	if len(parts) != len(t.trees) {
		return nil, fmt.Errorf("talent string %q has %d trees, want %d", s, len(parts), len(t.trees))
	}
	b := build{}
	for ti, part := range parts {
		if len(part) > len(t.trees[ti].Talents) {
			return nil, fmt.Errorf("talent string %q tree %d has %d digits for %d talents", s, ti, len(part), len(t.trees[ti].Talents))
		}
		for j, c := range part {
			if c < '0' || c > '9' {
				return nil, fmt.Errorf("talent string %q has a non-digit %q", s, c)
			}
			if r := int(c - '0'); r > 0 {
				b[t.trees[ti].Talents[j].ID] = r
			}
		}
	}
	return b, nil
}

// distance is the number of points two builds place differently
// (half the L1 distance between their rank vectors, so one point moved
// is distance 1).
func distance(a, b build) int {
	d := 0
	for id, r := range a {
		d += abs(r - b[id])
	}
	for id, r := range b {
		if _, ok := a[id]; !ok {
			d += r
		}
	}
	return d / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// diffNames is the talents added and removed going from base to b,
// by name ("+3 Vengeance", "-5 Toughness"), in tree then tier order.
func (t talentTrees) diffNames(base, b build) (added, removed []string) {
	for _, tree := range t.trees {
		for _, n := range tree.Talents {
			switch d := b[n.ID] - base[n.ID]; {
			case d > 0:
				added = append(added, fmt.Sprintf("+%d %s", d, n.Name))
			case d < 0:
				removed = append(removed, fmt.Sprintf("%d %s", d, n.Name))
			}
		}
	}
	return added, removed
}

// summary is a build's points per tree, "a/b/c".
func (t talentTrees) summary(b build) string {
	tp := t.treePoints(b)
	parts := make([]string, len(tp))
	for i, p := range tp {
		parts[i] = fmt.Sprint(p)
	}
	return strings.Join(parts, "/")
}
