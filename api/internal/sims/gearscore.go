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
	full    map[string]map[bandKey]float64           // build -> (spec, band) -> FullGearScore
	// dpsPerPoint and raids are lootvalues.go's per-build reads of dps-per-point.json
	// and loot.json.
	dpsPerPoint       map[string]map[string]DPSPoint        // build -> spec -> DPS per reference point
	dpsPerPointLogged map[string]bool                       // builds whose file failed to read, logged once
	raids             map[string]map[string]raidLoot        // build -> raidKey -> loot
	bis               map[string]map[string]map[int][][]int // build -> spec -> band -> sets (bestInSlot)
	bisRows           map[string]map[string][]bisEntry      // build -> spec -> entries (bandpicks.go)
	// rotations is rotation.go's per-build read of addon-data.json and spellranks.json.
	rotations       map[string]*buildRotations // build -> the shipped rotations
	rotationsLogged map[string]bool            // builds whose files failed to read, logged once
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
	gearSrc = &gearData{
		trees: t, dir: dir,
		weights:           map[string]map[string]map[string]float64{},
		full:              map[string]map[bandKey]float64{},
		dpsPerPoint:       map[string]map[string]DPSPoint{},
		dpsPerPointLogged: map[string]bool{},
		raids:             map[string]map[string]raidLoot{},
		bis:               map[string]map[string]map[int][][]int{},
		bisRows:           map[string]map[string][]bisEntry{},
		rotations:         map[string]*buildRotations{},
		rotationsLogged:   map[string]bool{},
	}
}

// currentGearData is what SetGearData last wired, or nil.
func currentGearData() *gearData {
	gearMu.RLock()
	defer gearMu.RUnlock()
	return gearSrc
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

// ErrNoBestInSlot is returned by FullGearScore for a spec the build has no
// best-in-slot set for in the band asked for (no bis/<spec>.json, or no set of
// that band in it).
var ErrNoBestInSlot = errors.New("sims: no best-in-slot set")

// The best-in-slot bands: pipeline.addonbis.BIS_LEVEL_BANDS, one every ten
// levels from 20 to 60, the same rungs the addon's Gear.bandFor climbs.
const (
	// FirstGearBand is the lowest band: below it no best-in-slot set exists.
	FirstGearBand = 20
	// FullGearBand is the level-60 band, whose set is "full gear".
	FullGearBand = 60
	gearBandStep = 10
)

// bandKey is one FullGearScore cache entry: a spec in one band.
type bandKey struct {
	spec string
	band int
}

// GearBand is the band a character of level is measured in, by the addon's
// own rule (Gear.bandFor, addon/ForeverSixty/Gear.lua): the highest band at or
// below the level, rounded down to the band step and capped at 60, that
// spec's bis/<spec>.json in build has a set for (a spec in progress may skip
// a band, so the rule steps down to the next one). ok is false below
// FirstGearBand, and when no band at or below the level has a set: there is
// nothing to measure against.
func GearBand(ctx context.Context, build, spec string, level int) (band int, ok bool, err error) {
	if err := ctx.Err(); err != nil {
		return 0, false, err
	}
	if level < FirstGearBand {
		return 0, false, nil
	}
	gearMu.RLock()
	src := gearSrc
	gearMu.RUnlock()
	if src == nil {
		return 0, false, fmt.Errorf("%w: no gear data is loaded", ErrGearUnscorable)
	}
	sets, err := src.bestInSlot(build, spec)
	if err != nil {
		return 0, false, err
	}
	for candidate := min(FullGearBand, level/gearBandStep*gearBandStep); candidate >= FirstGearBand; candidate -= gearBandStep {
		if len(sets[candidate]) > 0 {
			return candidate, true, nil
		}
	}
	return 0, false, nil
}

// FullGearScore is the mark a gear score is measured against: GearScore of
// spec's own best-in-slot set for band in build (the highest-scoring of the
// band's sets <TreeDataDir>/<build>/bis/<spec>.json lists, one per preset and
// faction), under the same build's weights GearScore uses - so a score and its
// mark always come from one build. A spec's weights are in its own units (a
// healer's full set scores about 400, a rogue's about 5,700), which is why
// the mark is per spec; band is the character's own (GearBand), so a
// character below 60 is measured against the best set for its level.
// Computed once per (build, spec, band) and cached.
func FullGearScore(ctx context.Context, build, spec string, band int) (float64, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	gearMu.RLock()
	src := gearSrc
	gearMu.RUnlock()
	if src == nil {
		return 0, fmt.Errorf("%w: no gear data is loaded", ErrGearUnscorable)
	}
	key := bandKey{spec: spec, band: band}
	src.mu.Lock()
	if v, ok := src.full[build][key]; ok {
		src.mu.Unlock()
		return v, nil
	}
	src.mu.Unlock()

	b, classID, weights, err := src.specTables(build, spec)
	if err != nil {
		return 0, err
	}
	sets, err := src.bestInSlot(build, spec)
	if err != nil {
		return 0, err
	}
	best, found := 0.0, false
	for _, items := range sets[band] {
		total := 0.0
		for _, itemID := range items {
			if item, ok := b.Item(classID, itemID); ok {
				total += statScore(item.Stats, weights)
			}
		}
		if !found || total > best {
			best, found = total, true
		}
	}
	if !found {
		return 0, fmt.Errorf("%w: %s band %d in build %s", ErrNoBestInSlot, spec, band, build)
	}
	src.mu.Lock()
	if src.full[build] == nil {
		src.full[build] = map[bandKey]float64{}
	}
	src.full[build][key] = best
	src.mu.Unlock()
	return best, nil
}

// bestInSlot is loadBestInSlot of build and spec, kept once read; an error is not kept.
func (g *gearData) bestInSlot(build, spec string) (map[int][][]int, error) {
	g.mu.Lock()
	sets, ok := g.bis[build][spec]
	g.mu.Unlock()
	if ok {
		return sets, nil
	}
	sets, err := loadBestInSlot(g.dir, build, spec)
	if err != nil {
		return nil, err
	}
	g.mu.Lock()
	if g.bis[build] == nil {
		g.bis[build] = map[string]map[int][][]int{}
	}
	g.bis[build][spec] = sets
	g.mu.Unlock()
	return sets, nil
}

// loadBestInSlot reads <dir>/<build>/bis/<spec>.json's sets by band, each as
// its slots' item ids (an empty slot has none). A missing file is
// ErrNoBestInSlot.
func loadBestInSlot(dir, build, spec string) (map[int][][]int, error) {
	body, err := os.ReadFile(filepath.Join(dir, build, "bis", spec+".json"))
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("%w: %s in build %s", ErrNoBestInSlot, spec, build)
	}
	if err != nil {
		return nil, fmt.Errorf("sims: read best-in-slot for %s/%s: %w", build, spec, err)
	}
	var file struct {
		Bands []struct {
			Band  int `json:"band"`
			Slots []struct {
				ItemID int `json:"item_id"`
			} `json:"slots"`
		} `json:"bands"`
	}
	if err := json.Unmarshal(body, &file); err != nil {
		return nil, fmt.Errorf("sims: decode best-in-slot for %s/%s: %w", build, spec, err)
	}
	sets := map[int][][]int{}
	for _, band := range file.Bands {
		items := make([]int, 0, len(band.Slots))
		for _, slot := range band.Slots {
			if slot.ItemID != 0 {
				items = append(items, slot.ItemID)
			}
		}
		sets[band.Band] = append(sets[band.Band], items)
	}
	return sets, nil
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
