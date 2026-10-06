// Package leveling holds rules shared by the rotation ladder
// (sim/request/ladder.go, Phase 1a of
// docs/superpowers/specs/2026-09-28-rotation-accuracy-program-design.md)
// and the leveling BiS pipeline (sim/cmd/leveling-bis,
// docs/superpowers/specs/2026-09-28-leveling-bis-design.md):
//
//   - Talent truncation (this file): a guide's level-60 build, truncated
//     to the points a leveling character of a given level would
//     actually have spent. Moved verbatim out of sim/request/ladder.go
//     (lane bis-all's brief) so a second caller does not have to
//     re-derive or duplicate it; sim/request/ladder.go now imports this
//     package instead of defining these symbols itself.
//   - Effective required level (required_level.go; 2026-09-28
//     quest-levels lane): a candidate item's REAL level gate, which a
//     quest reward's or crafted item's own client required_level (very
//     often 0) does not state.
package leveling

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// TalentNode is the fields LadderTalentString and sim/cmd/talent-search
// need from one talent of data/builds/<build>/talents/<class>.json.
// PrereqTalentID is 0 when the talent has no prerequisite (the file's
// null).
type TalentNode struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Tier           int    `json:"tier"`
	Column         int    `json:"column"`
	MaxRank        int    `json:"max_rank"`
	PrereqTalentID int    `json:"prereq_talent_id"`
	PrereqRank     int    `json:"prereq_rank"`
}

// TalentTree is one of a class's three trees.
type TalentTree struct {
	Name     string       `json:"name"`
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

// TalentRanksFromString is the inverse of the positional encoding
// LadderTalentString writes ("a-b-c", one digit per talent in trees'
// own (tier, column) order - the format the published band `talents`
// field carries): ranks by stable talent id, for whichever digit each
// one actually holds. It is the one read every site that needs to
// re-express an already-positional talent string against a DIFFERENT
// layout (sim/internal/enginetalents.Layout.Reposition, and
// sim/cmd/talent-search's own decodeActive) starts from - decode by
// position once, by id from there on, rather than each caller
// re-deriving this loop.
func TalentRanksFromString(trees []TalentTree, s string) (map[int]int, error) {
	parts := strings.Split(s, "-")
	if len(parts) != len(trees) {
		return nil, fmt.Errorf("leveling: talent string %q has %d trees, want %d", s, len(parts), len(trees))
	}
	ranks := make(map[int]int)
	for ti, part := range parts {
		if len(part) > len(trees[ti].Talents) {
			return nil, fmt.Errorf("leveling: talent string %q tree %d has %d digits for %d talents", s, ti, len(part), len(trees[ti].Talents))
		}
		for j, c := range part {
			if c < '0' || c > '9' {
				return nil, fmt.Errorf("leveling: talent string %q has a non-digit %q", s, c)
			}
			if r := int(c - '0'); r > 0 {
				ranks[trees[ti].Talents[j].ID] = r
			}
		}
	}
	return ranks, nil
}

// ErrGuideBuildMismatch is RequireGuideBuildMatchesActive's error.
var ErrGuideBuildMismatch = errors.New("leveling: guide build does not match the active build")

// RequireGuideBuildMatchesActive fails loudly when a guide's own FS1
// stamp (GuideBuildTalents' first return value) is not the site's
// active build.
//
// GuideTalentTargets maps a guide's digits onto talent ids by walking
// the STAMPED build's own trees POSITIONALLY; LadderTalentString then
// re-resolves those ids against the active build's trees, which only
// recovers the right talent if the stamped build's own tree ordering
// put each digit on the id GuideTalentTargets assumed. A guide whose
// stamp drifts from the active build (every guide's did, briefly,
// before lane guide-codes-70009 restamped them from 1.60.1.69893 to
// 1.60.1.70009 - the digits were always authored in 70009's order, the
// stamp just said otherwise) silently misaligns every digit after a
// tree whose shape changed between the two builds, exactly the way an
// engine whose compiled proto predates the active build misaligns a
// positional string handed to it unconverted
// (sim/internal/enginetalents' own doc). This is that same invariant,
// checked at the other end of the pipeline: a guide's stamp must name
// the build its own digits are already positioned for.
func RequireGuideBuildMatchesActive(guideBuild, activeBuild string) error {
	if guideBuild != activeBuild {
		return fmt.Errorf("%w: guide stamp %q, active build %q", ErrGuideBuildMismatch, guideBuild, activeBuild)
	}
	return nil
}
