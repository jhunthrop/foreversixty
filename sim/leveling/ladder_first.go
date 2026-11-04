package leveling

// Per-spec leveling reservations (sim-accuracy round 10, healprofile lane).
//
// Below 60 a band's build is the guide's level-60 build cut down by the top-row-down walk
// (LadderTalentString). For most specs that is the build; for a spec whose band builds lose
// beyond error on their own fight, data/curated/ladder-first.json names ranks of its own
// tree to take first, in order, each once the build can legally hold it, and never past the
// guide's target for that talent. Because every level spends the same list the same way,
// the ladder still only ever adds points, and at 60 every target is met, so the level-60
// build is the guide's. A spec with no row walks exactly as before.

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"
)

// LadderFirstRank is one reserved rank: the talent, by its client name, and the rank to
// hold before the walk.
type LadderFirstRank struct {
	Talent string `json:"talent"`
	Rank   int    `json:"rank"`
}

type ladderFirstSpec struct {
	Reason string            `json:"reason"`
	First  []LadderFirstRank `json:"first"`
}

type ladderFirstFile struct {
	Notes string                     `json:"notes"`
	Specs map[string]ladderFirstSpec `json:"specs"`
}

// ladderFirstJSON is data/curated/ladder-first.json (`make ladder-first-sync`);
// TestEmbeddedLadderFirstMatchesCurated holds it to the curated file.
//
//go:embed ladder-first.json
var ladderFirstJSON []byte

func parseLadderFirst(raw []byte) (ladderFirstFile, error) {
	var f ladderFirstFile
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&f); err != nil {
		return ladderFirstFile{}, fmt.Errorf("leveling: decoding ladder-first.json: %w", err)
	}
	for spec, row := range f.Specs {
		if row.Reason == "" || len(row.First) == 0 {
			return ladderFirstFile{}, fmt.Errorf("leveling: ladder-first %s states its reason and at least one rank", spec)
		}
		for _, r := range row.First {
			if r.Talent == "" || r.Rank <= 0 {
				return ladderFirstFile{}, fmt.Errorf("leveling: ladder-first %s names a talent and a positive rank", spec)
			}
		}
	}
	return f, nil
}

var embeddedLadderFirst = sync.OnceValue(func() ladderFirstFile {
	f, err := parseLadderFirst(ladderFirstJSON)
	if err != nil {
		panic(err) // embedded data; TestEmbeddedLadderFirstMatchesCurated fails first
	}
	return f
})

// LadderOption changes how LadderTalentString spends a level's points.
type LadderOption func(*ladderConfig)

type ladderConfig struct {
	first []LadderFirstRank
}

// ForSpec spends spec's curated reservation first, if it has one, then its hand-set
// leveling order (ladder_order.go), so a spec can carry both.
func ForSpec(spec string) LadderOption {
	first := embeddedLadderFirst().Specs[spec].First
	order := embeddedLadderOrder().Specs[spec].First
	return withFirst(append(append([]LadderFirstRank{}, first...), order...))
}

func withFirst(first []LadderFirstRank) LadderOption {
	return func(c *ladderConfig) { c.first = first }
}

// reservedRanks resolves first against tree by talent name into cumulative steps of one
// rank each, in list order, each talent capped at the guide's target and its max rank. Names not in tree are
// skipped (TestCuratedLadderFirstIsSound names them).
func reservedRanks(tree TalentTree, targets map[int]int, first []LadderFirstRank) []map[int]int {
	byName := make(map[string]TalentNode, len(tree.Talents))
	for _, node := range tree.Talents {
		byName[node.Name] = node
	}
	steps := make([]map[int]int, 0, len(first))
	acc := map[int]int{}
	for _, r := range first {
		node, ok := byName[r.Talent]
		if !ok {
			continue
		}
		// One step a rank, so a level that can hold only part of an entry takes that part.
		for rank := acc[node.ID] + 1; rank <= min(r.Rank, targets[node.ID], node.MaxRank); rank++ {
			next := make(map[int]int, len(acc)+1)
			for id, held := range acc {
				next[id] = held
			}
			next[node.ID] = rank
			acc = next
			steps = append(steps, next)
		}
	}
	return steps
}

// spendReserved is spendTree with the longest prefix of the reservation the budget can
// legally hold: the whole list first, then dropping from its end.
func spendReserved(digits []int, tree TalentTree, targets map[int]int, steps []map[int]int, budget int) int {
	for k := len(steps); k > 0; k-- {
		reserve := steps[k-1]
		if spent(reserve) > budget {
			continue
		}
		trial := make([]int, len(digits))
		left := spendTree(trial, tree, targets, reserve, budget)
		if treeIsLegal(tree, trial) {
			copy(digits, trial)
			return left
		}
	}
	return spendTree(digits, tree, targets, nil, budget)
}

// treeIsLegal holds a tree's ranks to the planner's rules: pointsPerTier*tier points in the
// rows above a talent with a rank, and its prerequisite at the rank it lists.
func treeIsLegal(tree TalentTree, digits []int) bool {
	rankOf := make(map[int]int, len(tree.Talents))
	for j, node := range tree.Talents {
		rankOf[node.ID] = digits[j]
	}
	for j, node := range tree.Talents {
		if digits[j] == 0 {
			continue
		}
		above := 0
		for k, other := range tree.Talents {
			if other.Tier < node.Tier {
				above += digits[k]
			}
		}
		if above < pointsPerTier*node.Tier {
			return false
		}
		if node.PrereqTalentID != 0 && rankOf[node.PrereqTalentID] < max(node.PrereqRank, 1) {
			return false
		}
	}
	return true
}
