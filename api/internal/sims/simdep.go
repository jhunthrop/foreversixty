// Package sims is the simulator's server side: saved results, the
// premium lane's dispatch and progress, the spec fidelity figures the
// support page reads, and the character model the sim page starts
// from. It sits beside reports and rankings and restructures neither.
//
// It imports only the engine-free half of the sim module - sim/api,
// sim/enginever, sim/specs and sim/runner. sim/adapter and sim/request
// import the engine itself, and a dependency module's replace does not
// apply to its consumer, so importing either here would break the
// image's GOWORK=off build. The native engine is reached as a
// subprocess, through sim/runner.
package sims

import (
	"slices"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
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

// roleDPS is the only role the simulator models at launch: tanks and
// healers are research problems and stay out of the first cut
// (design, "Scope at launch").
const roleDPS = "dps"

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
