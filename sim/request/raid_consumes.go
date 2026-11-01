package request

// The consumables every spec carries under the raid preset at each level, for the /sim page.
//
// The BiS ranker, the tier list and spec-breakdown resolve the raid preset for a spec and
// layer the class kit on top (a rogue's poisons, Thistle Tea; a Survival hunter's ranged
// group, and the kit follows the level: a rogue's poisons from 20, Deadly from 30). The page cannot run that resolution itself, so `go run ./internal/genraidconsumes`
// writes its output to sim/request/raid-consumes.json, web/scripts/sync-raid-preset.mjs
// copies it into the page's build, and TestRaidConsumesJSONIsCommitted fails when the
// committed file and this code disagree. One source of truth: data/curated/presets.json
// plus sim/leveling's kit.

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// LevelConsumes is the set a spec carries from FromLevel until the next row's level.
type LevelConsumes struct {
	FromLevel int      `json:"from_level"`
	Consumes  []string `json:"consumes"`
}

// raidConsumesFile is the generated file's shape.
type raidConsumesFile struct {
	Preset string                     `json:"preset"`
	Specs  map[string][]LevelConsumes `json:"specs"`
}

// RaidConsumes is, per spec key, the consumable ids the raid preset sends at each level,
// as rows that start where the set changes: the raid preset resolved for the spec, cut to
// the items that level can use (levels, the client's RequiredLevel), with the class kit of
// that level layered on top, as spec-breakdown and the BiS ranker build it at 60.
func RaidConsumes(presets Presets, levels ItemLevels) (map[string][]LevelConsumes, error) {
	out := make(map[string][]LevelConsumes, len(specs.All))
	for _, spec := range specs.All {
		resolved, err := presets.Resolve(RaidPreset, spec)
		if err != nil {
			return nil, fmt.Errorf("raid consumes for %s: %w", spec.Spec, err)
		}
		rows, err := levelRows(resolved, levels, spec.Spec)
		if err != nil {
			return nil, err
		}
		out[spec.Spec] = rows
	}
	return out, nil
}

// levelRows walks levels 1 to max and starts a row wherever the layered set changes.
func levelRows(resolved ResolvedPreset, levels ItemLevels, spec string) ([]LevelConsumes, error) {
	var rows []LevelConsumes
	for level := 1; level <= api.MaxLevel; level++ {
		_, consumeIDs := resolved.UsableAt(level, levels).Layer(leveling.KitBuffs(spec, level), leveling.KitConsumes(spec, level))
		if len(rows) == 0 || !slices.Equal(rows[len(rows)-1].Consumes, consumeIDs) {
			rows = append(rows, LevelConsumes{FromLevel: level, Consumes: consumeIDs})
		}
	}
	return rows, nil
}

// ConsumesAtLevel is the row in force at level; rows start at level 1.
func ConsumesAtLevel(rows []LevelConsumes, level int) []string {
	var at []string
	for _, row := range rows {
		if row.FromLevel > level {
			break
		}
		at = row.Consumes
	}
	return at
}

// RaidConsumesJSON is RaidConsumes as the committed file states it.
func RaidConsumesJSON(presets Presets, levels ItemLevels) ([]byte, error) {
	bySpec, err := RaidConsumes(presets, levels)
	if err != nil {
		return nil, err
	}
	b, err := json.MarshalIndent(raidConsumesFile{Preset: RaidPreset, Specs: bySpec}, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("encoding the raid consumes: %w", err)
	}
	return append(b, '\n'), nil
}
