// Package talents turns a set of chosen talent nodes into the engine's
// positional talent string.
//
// The engine takes talents as "0550000505021051-05-": one digit per
// talent, trees separated by "-", trailing zeros trimmed per tree, and
// the digit's position is the talent's position in its tree. Everything
// on this side - a planner build, a combat log's COMBATANT_INFO - names
// talents by node id instead. This package is the one conversion, and
// it reads the order from the build's own layout rather than a table
// written out by hand, so a patch that adds a talent is a pipeline run
// and not a code change.
//
// It imports nothing but the standard library, which is what lets the
// api module use it: sim/request, which would otherwise be the natural
// home, imports the engine.
package talents

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	// ErrUnknownClass is a class slug with no layout in this build.
	ErrUnknownClass = errors.New("talents: unknown class")
	// ErrUnknownTalent is a node id that is not in the class's trees.
	ErrUnknownTalent = errors.New("talents: unknown talent")
	// ErrTooManyRanks is more points in a talent than it has ranks,
	// which is a corrupt input rather than a digit the engine could
	// read.
	ErrTooManyRanks = errors.New("talents: too many ranks")
)

// Layouts is one client build's talent layout, by class slug.
type Layouts struct {
	byClass map[string]layout
}

// layout is one class: its trees, each already in tree order, each
// tree's talents already in (tier, column) order.
type layout struct {
	trees []tree
	// index maps a node id onto where it sits and how far it goes.
	index map[int64]position
}

type tree struct {
	position int
	nodes    []node
}

type node struct {
	id      int64
	maxRank int
}

type position struct {
	tree, slot, maxRank int
}

// The JSON the pipeline writes to data/builds/<build>/talents/<class>.json.
type layoutFile struct {
	ClassSlug string `json:"class_slug"`
	Trees     []struct {
		Position int `json:"position"`
		Talents  []struct {
			ID      int64 `json:"id"`
			MaxRank int   `json:"max_rank"`
			Tier    int   `json:"tier"`
			Column  int   `json:"column"`
		} `json:"talents"`
	} `json:"trees"`
}

// Load reads every <class>.json in one build's talents directory.
func Load(dir string) (*Layouts, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		return nil, fmt.Errorf("talents: reading %s: %w", dir, err)
	}
	out := &Layouts{byClass: make(map[string]layout, len(paths))}
	for _, path := range paths {
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("talents: reading %s: %w", path, err)
		}
		var f layoutFile
		if err := json.Unmarshal(b, &f); err != nil {
			return nil, fmt.Errorf("talents: %s is not a talent layout: %w", path, err)
		}
		if f.ClassSlug == "" || len(f.Trees) == 0 {
			return nil, fmt.Errorf("talents: %s carries no class or no trees", path)
		}
		out.byClass[f.ClassSlug] = build(f)
	}
	if len(out.byClass) == 0 {
		return nil, fmt.Errorf("talents: no layouts in %s", dir)
	}
	return out, nil
}

// build turns one file into the ordered form String walks.
func build(f layoutFile) layout {
	l := layout{index: map[int64]position{}}
	sorted := f.Trees
	sort.SliceStable(sorted, func(i, j int) bool { return sorted[i].Position < sorted[j].Position })
	for ti, t := range sorted {
		talents := t.Talents
		// The engine's talents proto is generated from this layout in
		// tier-then-column order, which is also how the tree reads on
		// screen, top-left first.
		sort.SliceStable(talents, func(i, j int) bool {
			if talents[i].Tier != talents[j].Tier {
				return talents[i].Tier < talents[j].Tier
			}
			return talents[i].Column < talents[j].Column
		})
		nodes := make([]node, 0, len(talents))
		for si, ta := range talents {
			max := ta.MaxRank
			if max <= 0 {
				max = 1
			}
			nodes = append(nodes, node{id: ta.ID, maxRank: max})
			l.index[ta.ID] = position{tree: ti, slot: si, maxRank: max}
		}
		l.trees = append(l.trees, tree{position: t.Position, nodes: nodes})
	}
	return l
}

// Classes is every class slug this build has a layout for, sorted.
func (l *Layouts) Classes() []string {
	out := make([]string, 0, len(l.byClass))
	for slug := range l.byClass {
		out = append(out, slug)
	}
	sort.Strings(out)
	return out
}

// nodes is every node id in a class's layout, for the tests.
func (l *Layouts) nodes(class string) []int64 {
	out := []int64{}
	for id := range l.byClass[class].index {
		out = append(out, id)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

// String turns a list of chosen nodes - one entry per point spent, so
// a node three times is three ranks in it - into the engine's
// positional talent string.
func (l *Layouts) String(classSlug string, points []int64) (string, error) {
	lay, ok := l.byClass[classSlug]
	if !ok {
		return "", fmt.Errorf("%w: %q", ErrUnknownClass, classSlug)
	}
	ranks := make([][]int, len(lay.trees))
	for i, t := range lay.trees {
		ranks[i] = make([]int, len(t.nodes))
	}
	for _, id := range points {
		at, ok := lay.index[id]
		if !ok {
			return "", fmt.Errorf("%w: %d in %s", ErrUnknownTalent, id, classSlug)
		}
		ranks[at.tree][at.slot]++
		if ranks[at.tree][at.slot] > at.maxRank {
			return "", fmt.Errorf("%w: %d points in talent %d, which has %d",
				ErrTooManyRanks, ranks[at.tree][at.slot], id, at.maxRank)
		}
	}
	parts := make([]string, len(ranks))
	for i, tree := range ranks {
		var b strings.Builder
		for _, r := range tree {
			b.WriteByte(byte('0' + r))
		}
		// Trailing zeros are trimmed per tree: the engine infers the
		// rest from the tree's size, which is why FillTalentsProto
		// takes the sizes separately.
		parts[i] = strings.TrimRight(b.String(), "0")
	}
	return strings.Join(parts, "-"), nil
}
