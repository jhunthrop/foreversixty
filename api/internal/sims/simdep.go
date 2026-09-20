// Package sims is the simulator's server side: saved results, the
// premium lane's dispatch and progress, the spec fidelity figures the
// support page reads, and the character model the sim page starts
// from. It sits beside reports and rankings and restructures neither.
//
// It imports only the engine-free half of the sim module - sim/api,
// sim/enginever, sim/specs, sim/runner and sim/measure. sim/adapter and
// sim/request import the engine itself, and a dependency module's
// replace does not apply to its consumer, so importing either here
// would break the image's GOWORK=off build. The native engine is
// reached as a subprocess, through sim/runner.
package sims

import (
	"slices"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/measure"
	"github.com/jhunthrop/foreversixty/sim/runner"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// The iteration counts the product offers, named: the planner's live
// estimate, the default run, and the precision toggle. They are the
// closed set simapi.ValidIterations accepts, and simdep_test.go
// asserts they stay in step with it.
const (
	liveIterations    = 500
	defaultIterations = 3000
	preciseIterations = 10000
)

// The shape of the job the estimate is made against. simJobCPUs and
// nativeRate both read their numbers from sim/measure rather than
// restating them: simJobCPUs is measure.NativeJobCPUs, and the
// engine's own rate is measure.NativeIterationsPerCPUSecond, which
// sim/measure publishes from its benchmark (contract A2). A second
// copy of a benchmark figure is a copy that goes stale, so neither is
// written here as a literal. api/README.md creates the sim-run job
// with `--cpu 4`, matching measure.NativeJobCPUs; the two have to
// change together, and the README says so.
const (
	simJobCPUs = measure.NativeJobCPUs
	nativeRate = measure.NativeIterationsPerCPUSecond * simJobCPUs
)

// estimateSec is how long the sim-run job needs to complete iterations
// iterations, rounded up: a run that would take a fraction of a second
// still takes some, and rounding up rather than down is what keeps
// "over budget" a certainty rather than an optimistic guess.
// Multiplied out, nativeRate x BulkBudget is the iteration ceiling
// contract A2 names.
func estimateSec(iterations int) int {
	if iterations <= 0 {
		return 0
	}
	return (iterations + nativeRate - 1) / nativeRate
}

// roleDPS is the only role the simulator models at launch: tanks and
// healers are research problems and stay out of the first cut
// (design, "Scope at launch").
const roleDPS = "dps"

// Engine is a StageRunner that can also plan: what cmd/api hands the
// service and the jobs alike, so one binary-presence check serves
// both. sim/runner already publishes Planner — counting a bulk or
// weights request without running it, by invoking `forever-sim -plan`
// (contract 10.2) — so it is used directly here rather than
// redeclared; sim/bulk cannot be imported by this package (it reads
// sim/internal/simdb, which imports the engine), so the count crosses
// the boundary as a subprocess's JSON, exactly the way a run does.
// Engine embeds StageRunner rather than Runner because the job
// already needs stage progress for a bulk run; a Runner-only Engine
// would silently drop it.
type Engine interface {
	runner.StageRunner
	runner.Planner
}

// withEncounterDefaults fills a partially specified encounter from the
// envelope's own defaults. Zero is not a legal value for any of these
// fields, so "unset" and "zero" do not have to be told apart, and a
// request that arrives with an empty encounter is runnable rather than
// rejected by SimRequest.Validate for a sixty-second floor it never
// meant to break.
func withEncounterDefaults(e simapi.EncounterSpec) simapi.EncounterSpec {
	d := simapi.DefaultEncounter()
	if e.DurationSec == 0 {
		e.DurationSec = d.DurationSec
	}
	if e.Variation == 0 {
		e.Variation = d.Variation
	}
	if e.Targets == 0 {
		e.Targets = d.Targets
	}
	if e.ExecuteRatio == 0 {
		e.ExecuteRatio = d.ExecuteRatio
	}
	return e
}

// DPSSpecs is every spec the simulator models, sorted. The list is the
// data lane's generated one; nothing in this package names a spec. It
// is exported because cmd/api/main.go builds the nightly job's spec
// list from it (Task 13).
func DPSSpecs() []string {
	out := make([]string, 0, len(specs.All))
	for _, s := range specs.All {
		if s.Role == roleDPS {
			out = append(out, s.Spec)
		}
	}
	slices.Sort(out)
	return out
}
