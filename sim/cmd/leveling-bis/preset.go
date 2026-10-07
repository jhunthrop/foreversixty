package main

import (
	"path/filepath"

	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// The sim presets a band entry is measured under. A bare entry is the
// character as the ranker has always built it (gear, talents and the
// class kit); a raid entry adds data/curated/presets.json's Phase 1 raid
// context on top of the same character.
const (
	presetBare = "bare"
	presetRaid = request.RaidPreset
)

// raidPresetBand is the one band that also gets a raid entry: level 60
// is the headline number and the only band a raid context describes.
const raidPresetBand = 60

// presetsPath is where the curated presets live under the repository root.
func presetsPath(repoRoot string) string {
	return filepath.Join(repoRoot, "data", "curated", "presets.json")
}

// resolveRaidPreset reads the curated presets and resolves the raid one
// for spec. It fails when the file is missing or names an id the request
// layer cannot map, so a typo stops the run rather than ranking without
// the buff.
func resolveRaidPreset(repoRoot string, spec specInfo) (request.ResolvedPreset, error) {
	presets, err := request.LoadPresets(presetsPath(repoRoot))
	if err != nil {
		return request.ResolvedPreset{}, err
	}
	return presets.Resolve(presetRaid, specs.Spec{
		Spec:          spec.Spec,
		ClassSlug:     spec.ClassSlug,
		Role:          spec.Role,
		ReferenceStat: spec.ReferenceStat,
	})
}

// rankPass is one full ranking of a band: the spec as that pass sees it.
type rankPass struct {
	name string
	spec specInfo
}

// passesFor lists the rankings a band gets: always the bare one, first,
// and the raid one as well at raidPresetBand.
func passesFor(base specInfo, band int, raid request.ResolvedPreset) []rankPass {
	bare := base
	bare.Applied = nil
	passes := []rankPass{{name: presetBare, spec: bare}}
	if band != raidPresetBand {
		return passes
	}
	raided := base
	raided.Applied = &raid
	return append(passes, rankPass{name: presetRaid, spec: raided})
}
