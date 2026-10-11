package sims

// The gear score: one number for how good a set of gear is for a spec,
// defined exactly as the addon's own Gear.score (addon/ForeverSixty/Gear.lua)
// so the site and the game never disagree about it. It is a linear sum: for
// every item, every stat on its stat line times the spec's weight for that
// stat, a stat the spec does not weight counting nothing. Hit and crit are
// one stat each, as they are in the weights table.
//
// The stat line is the item's own stats map from the build's per-class item
// table (trees.Build.Item), which is the same line builds.AddonCode writes
// into the addon code the addon scores planned items from. The weights are
// the curated table (data/curated/stat-weights.json), read through the
// build's own copy of it, <TreeDataDir>/<build>/stat-weights.json, which the
// pipeline's weights gate keeps identical to the curated file and which is
// the table Data.lua ships as weights. docs/contracts/fixtures/gearscore.json
// pins one worked example for both sides.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// ErrGearUnscorable is returned when a gear code cannot be scored: it is not
// an FS1 code, it names a client build or class the loaded data does not
// have, the spec has no weights, or no gear data is wired at all.
var ErrGearUnscorable = errors.New("sims: gear cannot be scored")

// weightsFile is the build's copy of data/curated/stat-weights.json.
const weightsFile = "stat-weights.json"

// gearData is what GearScore reads: the item tables and the directory the
// per-build weights live in. GearScore is a plain function, so it is set
// once at startup (SetGearData) rather than passed in.
type gearData struct {
	trees *trees.Data
	dir   string

	mu      sync.Mutex
	weights map[string]map[string]map[string]float64 // build -> spec -> stat -> weight
}

var (
	gearMu  sync.RWMutex
	gearSrc *gearData
)

// SetGearData wires GearScore to the loaded item tables and the TreeDataDir
// they came from. A nil t unwires it, after which every score is
// ErrGearUnscorable.
func SetGearData(t *trees.Data, dir string) {
	gearMu.Lock()
	defer gearMu.Unlock()
	if t == nil {
		gearSrc = nil
		return
	}
	gearSrc = &gearData{trees: t, dir: dir, weights: map[string]map[string]map[string]float64{}}
}

// GearScore is Gear.score summed over every item gearCode (an FS1 export
// string) carries, under spec's weights, against the client build the code
// names. An item the build's tables do not have adds nothing, as a client
// stat key Gear.lua does not know adds nothing: a conservative score rather
// than a guessed one.
func GearScore(ctx context.Context, gearCode string, spec string) (float64, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	gearMu.RLock()
	src := gearSrc
	gearMu.RUnlock()
	if src == nil {
		return 0, fmt.Errorf("%w: no gear data is loaded", ErrGearUnscorable)
	}
	decoded, ok := fs1.Decode(gearCode)
	if !ok {
		return 0, fmt.Errorf("%w: not an FS1 code", ErrGearUnscorable)
	}
	b, ok := src.trees.Build(decoded.DataBuild)
	if !ok {
		return 0, fmt.Errorf("%w: no data for client build %q", ErrGearUnscorable, decoded.DataBuild)
	}
	classID, ok := classIDBySlug(b, decoded.ClassSlug)
	if !ok {
		return 0, fmt.Errorf("%w: no class %q in build %s", ErrGearUnscorable, decoded.ClassSlug, decoded.DataBuild)
	}
	weights, err := src.specWeights(decoded.DataBuild, spec)
	if err != nil {
		return 0, err
	}
	total := 0.0
	for _, itemID := range decoded.Gear {
		item, ok := b.Item(classID, itemID)
		if !ok {
			continue
		}
		total += statScore(item.Stats, weights)
	}
	return total, nil
}

// statScore is Gear.score for one stat line.
func statScore(stats map[string]int, weights map[string]float64) float64 {
	total := 0.0
	for name, value := range stats {
		total += float64(value) * weights[name]
	}
	return total
}

// classIDBySlug finds the build's class id for an FS1 class slug.
func classIDBySlug(b *trees.Build, slug string) (int, bool) {
	for _, c := range b.Classes() {
		if c.Slug == slug {
			return c.ID, true
		}
	}
	return 0, false
}

// specWeights reads one spec's weights for build, loading the build's file
// the first time it is asked for.
func (g *gearData) specWeights(build, spec string) (map[string]float64, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	bySpec, ok := g.weights[build]
	if !ok {
		var err error
		if bySpec, err = loadSpecWeights(g.dir, build); err != nil {
			return nil, err
		}
		g.weights[build] = bySpec
	}
	w, ok := bySpec[spec]
	if !ok {
		return nil, fmt.Errorf("%w: no weights for spec %q", ErrGearUnscorable, spec)
	}
	return w, nil
}

// loadSpecWeights reads <dir>/<build>/stat-weights.json into spec -> stat ->
// weight.
func loadSpecWeights(dir, build string) (map[string]map[string]float64, error) {
	body, err := os.ReadFile(filepath.Join(dir, build, weightsFile))
	if err != nil {
		return nil, fmt.Errorf("%w: read weights for %s: %v", ErrGearUnscorable, build, err)
	}
	var rows []struct {
		Spec    string             `json:"spec"`
		Weights map[string]float64 `json:"weights"`
	}
	if err := json.Unmarshal(body, &rows); err != nil {
		return nil, fmt.Errorf("sims: decode weights for %s: %w", build, err)
	}
	out := make(map[string]map[string]float64, len(rows))
	for _, r := range rows {
		out[r.Spec] = r.Weights
	}
	return out, nil
}
