package runner

import (
	"context"
	"fmt"
	"sync"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
)

// Fixture is a Runner that answers with a checked-in result instead
// of running anything. Every caller of a Runner is tested against it,
// and a deployment with no engine binary falls back to it rather than
// failing, so an environment without the artifact still serves.
//
// The result is built in Go, not read from a JSON file, so the
// compiler checks it against the logs engine's summary types: a field
// renamed in logs/ breaks this build rather than silently decoding to
// a zero.
type Fixture struct {
	mu sync.Mutex
	// Runs records every request it was handed, so a test can check
	// what a job asked for.
	Runs []api.SimRequest
	// Err, when set, is returned instead of a result.
	Err error
	// Mean, when non-zero, overrides the fixture's DPS, so a test can
	// put a scorer on a known ratio.
	Mean float64
	// Aborted, when true, makes Run answer the way Native does on
	// forever-sim's exit 130: a partial SimResult with Aborted set,
	// half the requested iterations run, and ErrAborted wrapped in
	// the returned error — so a caller like sims.Run is tested
	// against the real abort contract without invoking a binary.
	// Err takes precedence when both are set.
	Aborted bool
	// PlanCombinations is the combination count Plan answers with on
	// success, and what its IterationsTotal is computed from. Zero is
	// a valid answer - an empty bulk expansion - not "unset": unlike
	// Mean, a combination count has no obvious checked-in baseline to
	// default to instead.
	PlanCombinations int
	// CapBreach, when set, makes Plan answer this error instead of a
	// summary - the fixture's way of exercising the cap_exceeded path
	// without a bulk request large enough to hit one for real. Err
	// takes precedence over it, the same as Err does over Aborted.
	CapBreach *api.ErrCapExceeded
}

// Run answers from the fixture, reporting progress once at the
// halfway mark and once at the end, the way a real run streams. It is
// RunStaged with the three bulk fields dropped.
func (f *Fixture) Run(ctx context.Context, req api.SimRequest, onProgress Progress) (api.SimResult, error) {
	var staged StageProgress
	if onProgress != nil {
		staged = func(p api.Progress) { onProgress(p.IterationsRun, p.DPS.Mean) }
	}
	return f.RunStaged(ctx, req, staged)
}

// RunStaged answers from the fixture, reporting the wider stage
// progress. The fixture never runs stages, so the three bulk fields
// stay zero, the same as a plain run's ticks.
func (f *Fixture) RunStaged(_ context.Context, req api.SimRequest, onProgress StageProgress) (api.SimResult, error) {
	f.mu.Lock()
	f.Runs = append(f.Runs, req)
	err, mean, aborted := f.Err, f.Mean, f.Aborted
	f.mu.Unlock()
	if err != nil {
		return api.SimResult{}, err
	}
	if aborted {
		return f.abortedResult(req, onProgress), fmt.Errorf("%w: fixture", ErrAborted)
	}
	out := api.SimResult{
		EngineVersion: req.EngineVersion,
		Request:       req,
		Lane:          api.LaneServer,
		DPS:           api.Estimate{Mean: 1042.5, StdDev: 88.1, Error: 1.61, Min: 812, Max: 1290.3},
		IterationsRun: req.Iterations,
		DurationMS:    2470,
		Summary:       fixtureSummary(req),
	}
	if mean != 0 {
		out.DPS.Mean = mean
	}
	if onProgress != nil {
		onProgress(api.Progress{IterationsRun: req.Iterations / 2, DPS: api.Estimate{Mean: out.DPS.Mean}})
		onProgress(api.Progress{IterationsRun: req.Iterations, DPS: api.Estimate{Mean: out.DPS.Mean}})
	}
	return out, nil
}

// abortedResult builds the partial SimResult RunStaged returns
// alongside ErrAborted when Aborted is set: half the run, one
// progress tick, Aborted true — the same shape forever-sim writes on
// exit 130.
func (f *Fixture) abortedResult(req api.SimRequest, onProgress StageProgress) api.SimResult {
	done := req.Iterations / 2
	if onProgress != nil {
		onProgress(api.Progress{IterationsRun: done})
	}
	return api.SimResult{
		EngineVersion: req.EngineVersion,
		Request:       req,
		Lane:          api.LaneServer,
		IterationsRun: done,
		Aborted:       true,
	}
}

// Plan answers Plan from configuration, never from a real bulk
// expansion: Fixture imports neither sim/bulk nor sim/internal, so it
// could not compute one even if it wanted to. PlanCombinations is
// what a success answers with; CapBreach, when set, is returned
// instead of a summary, and Err takes precedence over both - the same
// order Err takes over Aborted in RunStaged.
func (f *Fixture) Plan(_ context.Context, req api.SimRequest) (api.PlanSummary, error) {
	f.mu.Lock()
	f.Runs = append(f.Runs, req)
	err, combos, breach := f.Err, f.PlanCombinations, f.CapBreach
	f.mu.Unlock()
	if err != nil {
		return api.PlanSummary{}, err
	}
	if breach != nil {
		return api.PlanSummary{}, *breach
	}
	out := api.PlanSummary{Kind: req.Kind(), Combinations: combos}
	if req.Bulk != nil {
		out.Cap = req.Bulk.Cap
		if ladder, ok := api.Ladders[req.Bulk.Precision]; ok {
			out.IterationsTotal = api.LadderIterations(ladder, combos)
		}
	}
	return out, nil
}

// Asked reports the requests handed to the fixture so far.
func (f *Fixture) Asked() []api.SimRequest {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]api.SimRequest{}, f.Runs...)
}

// fixtureSummary is a plausible three-minute Fury Warrior fight in the
// logs engine's own shape: one damage actor with three abilities, one
// aura, one cast row, one resource track and one roster row. It is
// built fresh per call so two runs never share a slice.
func fixtureSummary(req api.SimRequest) summary.Summary {
	const guid = "sim-player"
	durationMS := int64(req.Encounter.DurationSec) * 1000
	if durationMS == 0 {
		durationMS = 180_000
	}
	return summary.Summary{
		EngineVersion: "sim:" + req.EngineVersion,
		FightIndex:    1,
		DurationMS:    durationMS,
		DamageDone: []summary.Actor{{
			GUID: guid, Name: "Simulated", Class: req.Character.Class,
			Total: 187650, Effective: 187650, ActiveMS: durationMS,
			Abilities: []summary.Ability{
				{SpellID: 0, Name: "Melee", Total: 91200, Effective: 91200,
					Hits: 148, Crits: 61, Min: 210, Max: 940},
				{SpellID: 23881, Name: "Bloodthirst", Total: 62300, Effective: 62300,
					Hits: 30, Crits: 22, Min: 560, Max: 1480},
				{SpellID: 1680, Name: "Whirlwind", Total: 34150, Effective: 34150,
					Hits: 14, Crits: 6, Min: 480, Max: 1310},
			},
			Targets: []summary.Pair{{GUID: "sim-target-1", Name: "Training Dummy", Total: 187650}},
			Series:  []int64{},
		}},
		DamageTaken:  []summary.Actor{},
		Healing:      []summary.Actor{},
		HealingTaken: []summary.Actor{},
		Deaths:       []summary.Death{},
		Auras: []summary.AuraTrack{{
			TargetGUID: guid, TargetName: "Simulated", SpellID: 12974, Name: "Flurry",
			Type: "BUFF", Applications: 61, MaxStacks: 3, UptimeMS: durationMS * 78 / 100,
			Segments: []summary.Segment{}, Appliers: []string{guid},
		}},
		Casts: []summary.CastRow{{
			GUID: guid, Name: "Simulated", OwnerGUID: guid, SpellID: 23881,
			SpellName: "Bloodthirst", Started: 52, Succeeded: 52, Sequence: []int64{},
		}},
		Interrupts: []summary.ExchangeRow{},
		Dispels:    []summary.ExchangeRow{},
		Resources: []summary.ResourceTrack{{
			GUID: guid, Name: "Simulated", PowerType: 1, Series: []int64{},
			Gained: 3480, Spent: 3410, Max: 100,
		}},
		Threat:         []summary.ThreatRow{},
		ThreatByTarget: []summary.ThreatPair{},
		Taunts:         []summary.Taunt{},
		Combatants:     []summary.CombatantRow{},
		Roster: []summary.RosterRow{{
			GUID: guid, Name: "Simulated", Class: req.Character.Class, Spec: req.Spec,
			Role: "dps", ItemLevel: 66, ActiveMS: durationMS, ActivityPct: 100,
			DamageDone: 187650, DPS: 1042.5,
		}},
		Mechanics: summary.MechanicsBlock{TableFound: false, Rows: []summary.MechanicRow{}},
		Phases:    []summary.Phase{},
	}
}
