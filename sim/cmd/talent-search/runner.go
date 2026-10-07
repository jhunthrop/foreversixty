package main

import (
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// simSetup is everything but the talents: one spec, one band's gear
// and race, one seed.
type simSetup struct {
	spec   specInfo
	band   bisBand
	level  int
	seed   int64
	layout enginetalents.Layout
	trees  talentTrees
}

// character is the band's character wearing the band's BiS set with
// talents - built the way the ranker builds its verify character
// (sim/cmd/leveling-bis bandCharacter + plainRequest): the class kit,
// and casters standing out of melee range.
func (s simSetup) character(talents string) api.CharacterSpec {
	ch := api.CharacterSpec{
		Name:     "talent-search",
		Race:     s.band.Race,
		Class:    s.spec.ClassSlug,
		Level:    s.level,
		Talents:  talents,
		Gear:     s.band.gear(),
		Consumes: leveling.KitConsumes(s.spec.Spec, s.level),
		Buffs:    leveling.KitBuffs(s.spec.Spec, s.level),
	}
	if leveling.NoMeleeAutoAttackSpecs[s.spec.Spec] {
		ch.DistanceFromTarget = leveling.CasterDistanceFromTarget
	}
	return ch
}

// engineRun is the real dpsFunc: the build written in the compiled
// engine's own layout (enginetalents), simmed in process, memoised by
// build and iteration count.
func engineRun(s simSetup) dpsFunc {
	cache := map[string]estimate{}
	return func(b build, iterations int) (estimate, error) {
		key := fmt.Sprintf("%s@%d", s.trees.key(b), iterations)
		if est, ok := cache[key]; ok {
			return est, nil
		}
		talents, err := s.layout.Encode(s.trees.trees, b)
		if err != nil {
			return estimate{}, err
		}
		est, err := inproc.PlainDPS(api.SimRequest{
			EngineVersion: enginever.Version,
			Spec:          s.spec.Spec,
			Source:        api.CharacterSource{Kind: api.SourceBuild},
			Character:     s.character(talents),
			Encounter:     api.DefaultEncounter(),
			Iterations:    iterations,
			RandomSeed:    s.seed,
		})
		if err != nil {
			return estimate{}, err
		}
		out := estimate{Mean: est.Mean, Err: est.Error}
		cache[key] = out
		return out, nil
	}
}
