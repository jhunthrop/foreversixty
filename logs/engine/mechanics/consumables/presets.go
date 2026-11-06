package consumables

import (
	"embed"
	"encoding/json"
	"fmt"
	"maps"
)

// PresetBand is one level band of a spec's Raid-buffed preset: from FromLevel up to the
// next band, the categories the preset carries, each with the preset ids that carry it.
type PresetBand struct {
	FromLevel int64               `json:"from_level"`
	Expects   map[string][]string `json:"expects"`
}

// PresetTable is presets.json: each spec's bands, lowest first, derived from Source
// (sim/request/raid-consumes.json) by TestPresetsJSONMatchesRaidConsumes, which fails when
// the two disagree.
type PresetTable struct {
	Source string                  `json:"source"`
	Specs  map[string][]PresetBand `json:"specs"`
}

//go:embed presets.json
var presetsFile embed.FS

var defaultPresets = mustParsePresets()

func mustParsePresets() PresetTable {
	data, err := presetsFile.ReadFile("presets.json")
	if err != nil {
		panic(fmt.Errorf("consumables: reading presets.json: %w", err))
	}
	var p PresetTable
	if err := json.Unmarshal(data, &p); err != nil {
		panic(fmt.Errorf("consumables: presets.json: %w", err))
	}
	for spec, bands := range p.Specs {
		for _, b := range bands {
			for cat := range b.Expects {
				if !isCategory(cat) {
					panic(fmt.Errorf("consumables: presets.json: %s names unknown category %q", spec, cat))
				}
			}
		}
	}
	return p
}

// PresetExpects is the categories spec's raid preset carries at level, read from the band
// whose from_level is the highest at or below level. ok is false when the level is unknown
// (0), the spec has no preset, or no band starts at or below level: the caller keeps the
// role's expectation.
func PresetExpects(spec string, level int64) (map[string]bool, bool) {
	if level <= 0 {
		return nil, false
	}
	var band *PresetBand
	for i, b := range defaultPresets.Specs[spec] {
		if b.FromLevel <= level && (band == nil || b.FromLevel > band.FromLevel) {
			band = &defaultPresets.Specs[spec][i]
		}
	}
	if band == nil {
		return nil, false
	}
	out := make(map[string]bool, len(band.Expects))
	for cat := range band.Expects {
		out[cat] = true
	}
	return out, true
}

// ForPlayer is role's catalogue narrowed to the categories spec's raid preset carries at
// level (PresetExpects): the one expectation the coach's pull check and the Forever Score's
// Preparation share. With no level or no preset for spec the role alone decides. false
// when the role has no catalogue.
func (c Catalogue) ForPlayer(role, spec string, level int64) (RoleCatalogue, bool) {
	rc, ok := c.For(role)
	if !ok {
		return rc, false
	}
	if expects, ok := PresetExpects(spec, level); ok {
		return rc.Expect(expects), true
	}
	return rc, true
}

// Expect is rc narrowed to the given categories: Expects is false for every other preset
// category, so MissedAtPull never calls one missing and IsComplete never lets one prove
// absence. world_buffs, which no preset carries, stays the role's. rc itself is unchanged.
func (rc RoleCatalogue) Expect(categories map[string]bool) RoleCatalogue {
	rc.expected = maps.Clone(categories)
	if rc.expected == nil {
		rc.expected = map[string]bool{}
	}
	return rc
}

// Expects reports whether the player is expected to carry the category: every category
// when rc was never narrowed by Expect, else world_buffs and the ones Expect named.
func (rc RoleCatalogue) Expects(category string) bool {
	return rc.expected == nil || category == CategoryWorldBuffs || rc.expected[category]
}
