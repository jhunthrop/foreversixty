package main

import (
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// The page's trust line ("simulated against a level-63 target ... over
// 180 seconds") reads these from the file instead of hard-coding them
// (coordinator scope, 2026-11-05): runConditions is top level, constant
// for every band; bandReport.TargetLevel is the one per-band number. The
// buffs and consumes each preset applies are already the file's
// top-level presets map, named by bandReport.Preset ("bare" applies none).

// runConditions is the encounter and harness every published number in
// the file was simmed under.
type runConditions struct {
	// TargetType is the creature type ("unknown": slaying bonuses count
	// for nothing).
	TargetType string `json:"target_type"`
	// TargetLevelOffset is the target's level above the band's character
	// level (api.DefaultTargetLevel); bandReport.TargetLevel applies it.
	TargetLevelOffset int     `json:"target_level_offset"`
	DurationSec       int     `json:"duration_sec"`
	Variation         float64 `json:"variation"`
	Targets           int     `json:"targets"`
	ExecuteRatio      float64 `json:"execute_ratio"`
	// Iterations is what the ranker asks the engine for, per run.
	Iterations runIterations `json:"iterations"`
	// HealIterationScale multiplies every healer run's iterations (capped
	// at the request layer's maximum); absent for other roles.
	HealIterationScale int   `json:"heal_iteration_scale,omitempty"`
	Seed               int64 `json:"seed"`
}

type runIterations struct {
	// Set is the run set_dps, set_dps_error and filtered_set_dps
	// publish (final_set.go).
	Set int `json:"set"`
	// Verify is one swap or set check (verify.go).
	Verify int `json:"verify"`
	// Guard is the guard that compares the final set with the sets the
	// ranker considered (incumbent.go).
	Guard int `json:"guard"`
	// TrinketRank is one trinket tournament entry (trinkets.go).
	TrinketRank int `json:"trinket_rank"`
	// WeightsPerDirection is one stat's weights run, each direction.
	WeightsPerDirection int `json:"weights_per_direction"`
}

// rankerRunConditions is the run block for one spec's file.
func rankerRunConditions(spec specInfo, weightsIterations int) runConditions {
	e := leveling.RankerEncounter()
	rc := runConditions{
		TargetType:        e.TargetType,
		TargetLevelOffset: api.DefaultTargetLevel(0),
		DurationSec:       e.DurationSec,
		Variation:         e.Variation,
		Targets:           e.Targets,
		ExecuteRatio:      e.ExecuteRatio,
		Iterations: runIterations{
			Set:                 finalSetEffectiveIterations(spec),
			Verify:              verifyIterations,
			Guard:               guardIterations(spec),
			TrinketRank:         trinketRankIterations,
			WeightsPerDirection: weightsIterations * api.WeightsIterationsFactor,
		},
		Seed: verifySeed,
	}
	if isHealer(spec) {
		rc.HealIterationScale = healIterationScale
	}
	return rc
}

// bandTargetLevel is the target level a band's sims fight.
func bandTargetLevel(band int) int {
	if t := leveling.RankerEncounter().TargetLevel; t != 0 {
		return t
	}
	return api.DefaultTargetLevel(band)
}
