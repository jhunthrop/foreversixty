// Package leveling holds the talent-truncation rule shared by the
// rotation ladder (sim/request/ladder.go, Phase 1a of
// docs/superpowers/specs/2026-09-28-rotation-accuracy-program-design.md)
// and the leveling BiS pipeline (sim/cmd/leveling-bis,
// docs/superpowers/specs/2026-09-28-leveling-bis-design.md): a guide's
// level-60 build, truncated to the points a leveling character of a
// given level would actually have spent.
//
// Moved verbatim out of sim/request/ladder.go (lane bis-all's brief)
// so a second caller does not have to re-derive or duplicate it;
// sim/request/ladder.go now imports this package instead of defining
// these symbols itself.
package leveling

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// TalentNode is the fields LadderTalentString needs from one talent of
// data/builds/<build>/talents/<class>.json.
type TalentNode struct {
	ID      int `json:"id"`
	Tier    int `json:"tier"`
	Column  int `json:"column"`
	MaxRank int `json:"max_rank"`
}

// TalentTree is one of a class's three trees.
type TalentTree struct {
	Position int          `json:"position"`
	Talents  []TalentNode `json:"talents"`
}

// TalentsFile is data/builds/<build>/talents/<class>.json's shape.
type TalentsFile struct {
	Trees []TalentTree `json:"trees"`
}

// LoadTalentTrees reads a build's talent trees for a class, sorted by
// tree position and then by (tier, column) within each tree - the same
// top-row-first, left-to-right order the client's own talent string
// digits walk (TestFillTalentsProto's WarriorTalents proto, generated
// from the client, carries one field per talent in exactly this order).
func LoadTalentTrees(repoRoot, build, class string) ([]TalentTree, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "talents", class+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("leveling: reading %s: %w", path, err)
	}
	var f TalentsFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("leveling: parsing %s: %w", path, err)
	}
	trees := append([]TalentTree(nil), f.Trees...)
	sort.SliceStable(trees, func(i, j int) bool { return trees[i].Position < trees[j].Position })
	for i := range trees {
		talents := append([]TalentNode(nil), trees[i].Talents...)
		sort.SliceStable(talents, func(a, b int) bool {
			if talents[a].Tier != talents[b].Tier {
				return talents[a].Tier < talents[b].Tier
			}
			return talents[a].Column < talents[b].Column
		})
		trees[i].Talents = talents
	}
	return trees, nil
}

// guideBuildRe pulls the FS1 build code out of a guide's frontmatter:
// build: 'FS1:<build>:<class>:<race>:<tree1>/<tree2>/<tree3>:'
var guideBuildRe = regexp.MustCompile(`build:\s*'FS1:([^:]+):([^:]+):([^:]+):([^/]+)/([^/]+)/([^:]+):'`)

// GuideBuildTalents reads a spec's guide and returns the client build its
// FS1 code names and the three trees' digit strings, in the code's own
// order.
func GuideBuildTalents(repoRoot, class, specSlug string) (build string, trees [3]string, err error) {
	path := filepath.Join(repoRoot, "web", "src", "content", "guides", class, specSlug+".md")
	b, err := os.ReadFile(path)
	if err != nil {
		return "", trees, fmt.Errorf("leveling: reading %s: %w", path, err)
	}
	m := guideBuildRe.FindSubmatch(b)
	if m == nil {
		return "", trees, fmt.Errorf("leveling: %s has no FS1 build line", path)
	}
	return string(m[1]), [3]string{string(m[4]), string(m[5]), string(m[6])}, nil
}

// GuideTalentTargets is the level-60 guide build, addressed by each
// talent's stable id rather than its position: the guide's own build
// code (GuideBuildTalents) is pinned to whatever client build the FS1
// tool exported it from, which is not always the site's active build -
// paladin lost two talents (Improved Holy Strike, Crusade) between
// 1.60.1.69893 and 1.60.1.70009, which would silently misalign every
// digit after them if the guide's string were read positionally against
// the active build's tree instead. Reading by id and re-resolving each
// talent's row against the active build (LadderTalentString) sidesteps
// that: a talent the active build no longer carries simply has nowhere
// to receive its points and is dropped, rather than shifting every
// talent after it.
func GuideTalentTargets(guideTrees []TalentTree, treeDigits [3]string) map[int]int {
	targets := make(map[int]int)
	for i, tree := range guideTrees {
		if i >= len(treeDigits) {
			break
		}
		digits := treeDigits[i]
		for j, node := range tree.Talents {
			rank := 0
			if j < len(digits) {
				if v, err := strconv.Atoi(string(digits[j])); err == nil {
					rank = v
				}
			}
			targets[node.ID] = rank
		}
	}
	return targets
}

// LadderTalentString is the truncated talent string for level: level-9
// points (0 below level 10), spent against the guide's level-60
// targets (GuideTalentTargets), walking the ACTIVE build's trees from
// the top row down - the spec's own tree first (ownTreeIndex, from
// sim/specs' TreeIndex), then the other two in the build code's own
// order (0, 1, 2, skipping the spec's own).
//
// Within a tree the walk is simply "top row down": each talent receives
// min(remaining budget, its guide target), in (tier, column) order. A
// row whose points were never reached because the budget ran out first
// is left at 0 without any special-casing - that is what "skipping a
// row not yet reachable" reduces to once the walk stops spending. If the
// guide build itself spends fewer than level-9 points (a guide is not
// always minmaxed to the last point), the remainder is left unspent
// rather than invented: this ladder approximates a leveling build, it
// does not design one.
func LadderTalentString(activeTrees []TalentTree, targets map[int]int, ownTreeIndex, level int) string {
	budget := level - 9
	if budget < 0 {
		budget = 0
	}
	order := make([]int, 0, len(activeTrees))
	order = append(order, ownTreeIndex)
	for i := range activeTrees {
		if i != ownTreeIndex {
			order = append(order, i)
		}
	}

	digits := make([][]int, len(activeTrees))
	for i, tree := range activeTrees {
		digits[i] = make([]int, len(tree.Talents))
	}
	for _, ti := range order {
		if ti < 0 || ti >= len(activeTrees) {
			continue
		}
		for j, node := range activeTrees[ti].Talents {
			if budget <= 0 {
				break
			}
			target := targets[node.ID]
			if target > node.MaxRank {
				target = node.MaxRank
			}
			give := target
			if give > budget {
				give = budget
			}
			digits[ti][j] = give
			budget -= give
		}
	}

	parts := make([]string, len(digits))
	for i, row := range digits {
		var sb strings.Builder
		for _, d := range row {
			sb.WriteByte(byte('0' + d))
		}
		parts[i] = sb.String()
	}
	return strings.Join(parts, "-")
}
