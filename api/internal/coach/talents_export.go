package coach

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
)

// TalentSource reads a character's talents from its stored build. nil talents with no error
// is "not on file": no stored build, one that does not decode, or one of another class.
type TalentSource interface {
	TalentsOf(ctx context.Context, key, class string) (PlayerTalents, error)
}

// RowQuerier is the one pgx method ExportTalents reads with (*pgxpool.Pool satisfies it).
type RowQuerier interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// ExportTalents is the production TalentSource: the character's addon_exports row (the
// addon's export or the Battle.net import, whichever came last), decoded against the
// talent trees of the data build the export names. The talents are today's, not the
// fight's: like the level, they are a current-character fact.
type ExportTalents struct {
	DB        RowQuerier
	BuildsDir string

	mu    sync.Mutex
	trees map[classKey][][]string
}

// dataBuildPattern is a client build version; anything else in an export is never a path.
var dataBuildPattern = regexp.MustCompile(`^[0-9]+(\.[0-9]+)+$`)

// TalentsOf is the talents of the character key, when its export is class's.
func (e *ExportTalents) TalentsOf(ctx context.Context, key, class string) (PlayerTalents, error) {
	var code string
	err := e.DB.QueryRow(ctx, `select export from addon_exports where character_key = $1`, key).Scan(&code)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("coach: export of %s: %w", key, err)
	}
	d, ok := fs1.Decode(code)
	if !ok || d.ClassSlug != classSlug(class) || !dataBuildPattern.MatchString(d.DataBuild) {
		return nil, nil
	}
	names, err := e.treeNames(d.DataBuild, d.ClassSlug)
	if errors.Is(err, ErrNoTables) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	talents, ok := DecodeTalents(d.Talents.Trees, names)
	if !ok {
		return nil, nil
	}
	return talents, nil
}

// treeNames loads and caches one class's talent names at a build.
func (e *ExportTalents) treeNames(build, class string) ([][]string, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	key := classKey{build, class}
	if names, ok := e.trees[key]; ok {
		return names, nil
	}
	names, err := readBuildFile(e.BuildsDir, build, "talents", class, ParseTalentNames)
	if err != nil {
		return nil, err
	}
	if e.trees == nil {
		e.trees = map[classKey][][]string{}
	}
	e.trees[key] = names
	return names, nil
}

// talentNode and talentTree are the fields of talents/<class>.json ParseTalentNames reads.
type talentNode struct {
	Name   string `json:"name"`
	Tier   int    `json:"tier"`
	Column int    `json:"column"`
}

type talentTree struct {
	Position int          `json:"position"`
	Talents  []talentNode `json:"talents"`
}

// ParseTalentNames reads data/builds/<build>/talents/<class>.json into each tree's talent
// names, trees by position and talents in tier then column order: the order an FS1
// export's one base-36 digit per talent follows.
func ParseTalentNames(data []byte) ([][]string, error) {
	var f struct {
		Trees []talentTree `json:"trees"`
	}
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("coach talents: %w", err)
	}
	trees := slices.SortedStableFunc(slices.Values(f.Trees), func(a, b talentTree) int {
		return cmp.Compare(a.Position, b.Position)
	})
	out := make([][]string, len(trees))
	for i, tree := range trees {
		talents := slices.SortedStableFunc(slices.Values(tree.Talents), func(a, b talentNode) int {
			return cmp.Or(cmp.Compare(a.Tier, b.Tier), cmp.Compare(a.Column, b.Column))
		})
		for _, t := range talents {
			out[i] = append(out[i], t.Name)
		}
	}
	return out, nil
}

// DecodeTalents is the talent ranks an export's tree strings spend, by name. false when the
// strings do not fit the trees (a digit that is not base 36, or more digits than talents):
// an export of another tree layout is not read on a guess.
func DecodeTalents(trees []string, names [][]string) (PlayerTalents, bool) {
	if len(trees) != len(names) {
		return nil, false
	}
	out := PlayerTalents{}
	for i, tree := range trees {
		if len(tree) > len(names[i]) {
			return nil, false
		}
		for j, digit := range tree {
			rank, err := strconv.ParseInt(string(digit), 36, 0)
			if err != nil {
				return nil, false
			}
			if rank > 0 {
				out[names[i][j]] = int(rank)
			}
		}
	}
	return out, true
}
