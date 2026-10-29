package sims

// The rotation a spec's players are measured against for rotation adherence
// (api/internal/rating/adherence.go; docs/contracts/addon-v2.md contract 3, last_raid's
// adherence).
//
// The rotation is the one the addon ships, read from the data build's addon-data.json
// (`rotations`: spec -> one band per pipeline.addonrotation.LEVEL_BANDS rung, each band's
// lines resolved to the rank learned by that rung). The API image carries data/builds but
// not data/curated, and reading the pipeline's own output shares its band rule and its rank
// resolving exactly instead of re-implementing them here. A band line names one rank; the
// same ability cast at any other rank still counts, through spellranks.json's family
// grouping (the grouping pipeline.addonrotation._family_ranks resolves lines with).

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	addonDataFile  = "addon-data.json"
	spellRanksFile = "spellranks.json"
)

// Rotation is one spec's shipped rotation at one level band.
type Rotation struct {
	// Band is the band's level: the highest pipeline.addonrotation.LEVEL_BANDS rung at or
	// below the character's level, or the lowest rung below it.
	Band int
	// listed is every spell id, at every rank, of every ability the band lists.
	listed map[int64]bool
	// family is the class's spellranks.json family name of each of its spell ids.
	family map[int64]string
}

// Lists reports whether spellID is a rank of an ability the band lists.
func (r Rotation) Lists(spellID int64) bool { return r.listed[spellID] }

// Family is the spellranks.json family name (the data's own ability name, whatever the log
// calls it) of spellID among the spec's class spells; ok is false for a spell that is no rank
// of any of the class's spells: an item's, a trinket's, a racial's or a consumable's effect.
func (r Rotation) Family(spellID int64) (name string, ok bool) {
	name, ok = r.family[spellID]
	return name, ok
}

// rotationBand is one band of a spec's shipped rotation.
type rotationBand struct {
	level  int
	listed map[int64]bool
}

// buildRotations is one build's shipped rotations: spec -> its bands, lowest level first,
// and spec -> its class's spell families by id.
type buildRotations struct {
	bands    map[string][]rotationBand
	familyOf map[string]map[int64]string
}

// RotationFor is spec's shipped rotation in build at level's band; build "" is the active
// build (trees.Data.Active). The band is the addon's own rule (Rotation.lua's bandFor): the highest band
// whose level is at or below level, else the lowest band. ok is false when no gear data is
// wired, the build's files cannot be read, or the build ships no rotation for the spec.
func RotationFor(build, spec string, level int) (Rotation, bool) {
	src := currentGearData()
	if src == nil {
		return Rotation{}, false
	}
	if build == "" {
		b, ok := src.trees.Active()
		if !ok {
			return Rotation{}, false
		}
		build = b.Version
	}
	src.mu.Lock()
	rots, ok := src.rotations[build]
	if !ok {
		var err error
		if rots, err = loadRotations(src.dir, build); err != nil {
			if !src.rotationsLogged[build] {
				src.rotationsLogged[build] = true
				slog.Default().Warn("sims", "op", "rotations", "build", build, "err", err)
			}
			src.mu.Unlock()
			return Rotation{}, false
		}
		src.rotations[build] = rots
		if len(rots.bands) == 0 && !src.rotationsLogged[build] {
			// Read once and kept, so this is said once per build, not once per rated player.
			src.rotationsLogged[build] = true
			slog.Default().Warn("sims", "op", "rotations", "build", build,
				"err", addonDataFile+" has no rotations; adherence is absent for every spec")
		}
	}
	src.mu.Unlock()

	bands := rots.bands[spec]
	if len(bands) == 0 {
		return Rotation{}, false
	}
	chosen := bands[0]
	for _, b := range bands {
		if b.level <= level {
			chosen = b
		}
	}
	return Rotation{Band: chosen.level, listed: chosen.listed, family: rots.familyOf[spec]}, true
}

// loadRotations reads <dir>/<build>/addon-data.json's rotations and spellranks.json's
// families, and resolves each band's lines to every rank of their abilities.
func loadRotations(dir, build string) (*buildRotations, error) {
	body, err := os.ReadFile(filepath.Join(dir, build, addonDataFile))
	if err != nil {
		return nil, fmt.Errorf("sims: read rotations for %s: %w", build, err)
	}
	var data struct {
		Rotations map[string][]struct {
			Level int `json:"level"`
			Lines []struct {
				SpellID int64  `json:"spell_id"`
				Name    string `json:"name"`
			} `json:"lines"`
		} `json:"rotations"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("sims: decode rotations for %s: %w", build, err)
	}
	body, err = os.ReadFile(filepath.Join(dir, build, spellRanksFile))
	if err != nil {
		return nil, fmt.Errorf("sims: read spell ranks for %s: %w", build, err)
	}
	var ranks struct {
		Classes map[string]map[string][]struct {
			ID int64 `json:"id"`
		} `json:"classes"`
	}
	if err := json.Unmarshal(body, &ranks); err != nil {
		return nil, fmt.Errorf("sims: decode spell ranks for %s: %w", build, err)
	}

	out := &buildRotations{bands: map[string][]rotationBand{}, familyOf: map[string]map[int64]string{}}
	byClass := map[string]map[int64]string{}
	for spec, bands := range data.Rotations {
		// The class is the spec slug's first word ("druid-feral-bear" is a druid), the
		// pipeline's own _class_slug.
		class, _, _ := strings.Cut(spec, "-")
		families := ranks.Classes[class]
		if _, done := byClass[class]; !done {
			// An id two families share (rare: an alternate id) names the first family by
			// name, so the answer is the same on every load.
			familyOf := map[int64]string{}
			for name, ids := range families {
				for _, rank := range ids {
					if prev, seen := familyOf[rank.ID]; !seen || name < prev {
						familyOf[rank.ID] = name
					}
				}
			}
			byClass[class] = familyOf
		}
		out.familyOf[spec] = byClass[class]
		for _, b := range bands {
			listed := map[int64]bool{}
			for _, line := range b.Lines {
				listed[line.SpellID] = true
				for _, rank := range families[line.Name] {
					listed[rank.ID] = true
				}
			}
			out.bands[spec] = append(out.bands[spec], rotationBand{level: b.Level, listed: listed})
		}
		sort.SliceStable(out.bands[spec], func(i, j int) bool { return out.bands[spec][i].level < out.bands[spec][j].level })
	}
	return out, nil
}
