package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

// estimate is one sim's mean DPS and its standard error.
type estimate struct {
	Mean, Err float64
}

// combinedErr is the error of a difference of two independent
// estimates - the bar a delta has to clear to count. Paired seeds
// (every sim in one search shares the same -seed) make the true error
// of a difference smaller than this; it is the conservative side.
func combinedErr(a, b estimate) float64 {
	return math.Hypot(a.Err, b.Err)
}

// simSetup is everything held fixed across this whole search: one
// spec, the guide's own talents (converted to the engine's own
// layout), and the committed BiS band's gear and race. Only the
// rotation varies from here on.
type simSetup struct {
	spec          specInfo
	band          bisBand
	level         int
	seed          int64
	engineTalents string
}

func (s simSetup) character() api.CharacterSpec {
	ch := api.CharacterSpec{
		Name:     "rotation-search",
		Race:     s.band.Race,
		Class:    s.spec.ClassSlug,
		Level:    s.level,
		Talents:  s.engineTalents,
		Gear:     s.band.gear(),
		Consumes: leveling.KitConsumes(s.spec.Spec, s.level),
		Buffs:    leveling.KitBuffs(s.spec.Spec, s.level),
	}
	if leveling.NoMeleeAutoAttackSpecs[s.spec.Spec] {
		ch.DistanceFromTarget = leveling.CasterDistanceFromTarget
	}
	return ch
}

// dpsFunc sims one rotation at a given iteration count. Every call in
// one search uses the same seed, so two rotations' difference is a
// paired comparison (common random numbers).
type dpsFunc func(r rotation, iterations int) (estimate, error)

// engineRun builds a dpsFunc backed by simRotation,
// memoised by rotation JSON and iteration count so a rotation this
// search has already simmed at this iteration count (the baseline,
// re-screened inside a later round) is never re-run. The search
// screens many mutations concurrently (goroutines - see search.go's
// screenAll), so the cache is guarded by a mutex; inproc's own
// Register() is a sync.Once and every call builds its own request and
// run id, so concurrent calls into the engine are otherwise safe.
func engineRun(s simSetup) (dpsFunc, error) {
	var mu sync.Mutex
	cache := map[string]estimate{}
	return func(r rotation, iterations int) (estimate, error) {
		b, err := r.marshalIndent()
		if err != nil {
			return estimate{}, fmt.Errorf("encoding the rotation to sim: %w", err)
		}
		key := fmt.Sprintf("%s@%d", string(b), iterations)
		mu.Lock()
		est, cached := cache[key]
		mu.Unlock()
		if cached {
			return est, nil
		}
		apiEst, _, err := simRotation(s, r, iterations)
		if err != nil {
			return estimate{}, err
		}
		out := estimate{Mean: apiEst.Mean, Err: apiEst.Error}
		mu.Lock()
		cache[key] = out
		mu.Unlock()
		return out, nil
	}, nil
}

// simRotation sims one rotation in this process and returns the
// estimate with the player's own metrics (the cast tally reads them).
func simRotation(s simSetup, r rotation, iterations int) (api.Estimate, *proto.UnitMetrics, error) {
	b, err := r.marshalIndent()
	if err != nil {
		return api.Estimate{}, nil, fmt.Errorf("encoding the rotation to sim: %w", err)
	}
	apl := &proto.APLRotation{}
	if err := protojson.Unmarshal(b, apl); err != nil {
		return api.Estimate{}, nil, fmt.Errorf("the mutated rotation does not parse as an engine APL: %w", err)
	}
	return inproc.PlainRunWithRotation(api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          s.spec.Spec,
		Source:        api.CharacterSource{Kind: api.SourceBuild},
		Character:     s.character(),
		Encounter:     api.DefaultEncounter(),
		Iterations:    iterations,
		RandomSeed:    s.seed,
	}, apl)
}
