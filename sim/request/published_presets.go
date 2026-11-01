package request

// Which preset each level band's published number is simmed under, and what the bare preset
// sends, for the /sim page (owner ruling, round 6: below 60 the page's default buffs match
// the published number at the character's band).
//
// The BiS ranker (sim/cmd/leveling-bis) ranks every band bare and adds a raid entry only at
// RaidPresetBand; the BiS page, the tier list and the home card headline the raid entry
// where a band has one. PublishedPreset states that rule once, the ranker reads it, and
// `go run ./internal/genraidconsumes` writes it with the bare sets to
// sim/request/published-presets.json for the page (held current by
// TestPublishedPresetsJSONIsCommitted).

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// RaidPresetBand is the one band the ranker also ranks under the raid preset: level 60 is
// the headline number and the only band a raid context describes.
const RaidPresetBand = 60

// PublishedPreset is the preset the published number at level is simmed under: the raid
// preset from RaidPresetBand, the bare character below it.
func PublishedPreset(level int) string {
	if level >= RaidPresetBand {
		return RaidPreset
	}
	return BarePreset
}

// PresetFrom is the preset published from FromLevel until the next row's level.
type PresetFrom struct {
	FromLevel int    `json:"from_level"`
	Preset    string `json:"preset"`
}

// LevelSet is what a preset sends for one spec from FromLevel until the next row's level.
type LevelSet struct {
	FromLevel int      `json:"from_level"`
	Buffs     []string `json:"buffs"`
	Consumes  []string `json:"consumes"`
}

// publishedPresetsFile is the generated file's shape.
type publishedPresetsFile struct {
	Published []PresetFrom          `json:"published"`
	Bare      map[string][]LevelSet `json:"bare"`
}

// PublishedPresets is PublishedPreset as rows over levels 1 to max, each starting where the
// preset changes.
func PublishedPresets() []PresetFrom {
	var rows []PresetFrom
	for level := 1; level <= api.MaxLevel; level++ {
		preset := PublishedPreset(level)
		if len(rows) == 0 || rows[len(rows)-1].Preset != preset {
			rows = append(rows, PresetFrom{FromLevel: level, Preset: preset})
		}
	}
	return rows
}

// BareSets is, per spec key, what the bare preset sends at each level: the class kit
// (sim/leveling KitBuffs and KitConsumes), as the ranker's bare pass and spec-breakdown
// -preset bare build it. Rows start where the set changes.
func BareSets() map[string][]LevelSet {
	out := make(map[string][]LevelSet, len(specs.All))
	for _, spec := range specs.All {
		var rows []LevelSet
		for level := 1; level <= api.MaxLevel; level++ {
			buffs := nonNil(leveling.KitBuffs(spec.Spec, level))
			consumes := nonNil(leveling.KitConsumes(spec.Spec, level))
			if len(rows) > 0 {
				last := rows[len(rows)-1]
				if slices.Equal(last.Buffs, buffs) && slices.Equal(last.Consumes, consumes) {
					continue
				}
			}
			rows = append(rows, LevelSet{FromLevel: level, Buffs: buffs, Consumes: consumes})
		}
		out[spec.Spec] = rows
	}
	return out
}

// SetAtLevel is the row in force at level; rows start at level 1.
func SetAtLevel(rows []LevelSet, level int) LevelSet {
	var at LevelSet
	for _, row := range rows {
		if row.FromLevel > level {
			break
		}
		at = row
	}
	return at
}

// PublishedPresetsJSON is PublishedPresets and BareSets as the committed file states them.
func PublishedPresetsJSON() ([]byte, error) {
	b, err := json.MarshalIndent(publishedPresetsFile{Published: PublishedPresets(), Bare: BareSets()}, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("encoding the published presets: %w", err)
	}
	return append(b, '\n'), nil
}

// nonNil keeps an empty set an empty JSON list rather than null.
func nonNil(ids []string) []string {
	if ids == nil {
		return []string{}
	}
	return ids
}
