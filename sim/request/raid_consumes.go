package request

// The consumables every spec's published level-60 raid number carries, for the /sim page.
//
// The BiS ranker, the tier list and spec-breakdown resolve the raid preset for a spec and
// layer the class kit on top (a rogue's poisons, Thistle Tea; a Survival hunter's ranged
// group). The page cannot run that resolution itself, so `go run ./internal/genraidconsumes`
// writes its output to sim/request/raid-consumes.json, web/scripts/sync-raid-preset.mjs
// copies it into the page's build, and TestRaidConsumesJSONIsCommitted fails when the
// committed file and this code disagree. One source of truth: data/curated/presets.json
// plus sim/leveling's kit.

import (
	"encoding/json"
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// raidConsumesFile is the generated file's shape.
type raidConsumesFile struct {
	Preset string              `json:"preset"`
	Level  int                 `json:"level"`
	Specs  map[string][]string `json:"specs"`
}

// RaidConsumes is, per spec key, the consumable ids the published level-60 raid numbers
// send: the raid preset resolved for the spec with the class kit layered on top.
func RaidConsumes(presets Presets) (map[string][]string, error) {
	out := make(map[string][]string, len(specs.All))
	for _, spec := range specs.All {
		resolved, err := presets.Resolve(RaidPreset, spec)
		if err != nil {
			return nil, fmt.Errorf("raid consumes for %s: %w", spec.Spec, err)
		}
		_, consumeIDs := resolved.Layer(
			leveling.KitBuffs(spec.Spec, api.MaxLevel),
			leveling.KitConsumes(spec.Spec, api.MaxLevel),
		)
		if len(consumeIDs) == 0 {
			return nil, fmt.Errorf("raid consumes for %s: the published set is empty", spec.Spec)
		}
		out[spec.Spec] = consumeIDs
	}
	return out, nil
}

// RaidConsumesJSON is RaidConsumes as the committed file states it.
func RaidConsumesJSON(presets Presets) ([]byte, error) {
	bySpec, err := RaidConsumes(presets)
	if err != nil {
		return nil, err
	}
	b, err := json.MarshalIndent(raidConsumesFile{Preset: RaidPreset, Level: api.MaxLevel, Specs: bySpec}, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("encoding the raid consumes: %w", err)
	}
	return append(b, '\n'), nil
}
