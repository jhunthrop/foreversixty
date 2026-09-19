package bulk

// Staging: how many sims at what precision.
//
// The ladder is api.Ladders, so the page can show it before the run
// starts. This file turns one rung into runnable requests.
//
// Three things about the requests matter and are easy to get wrong.
//
// The equipped set is Requests[0] of EVERY stage. It is not an
// optimisation to run it once at the end: a delta is only honest
// against a baseline measured the same way, and a baseline from a
// 100-iteration stage compared against a 3,000-iteration finalist
// would put the whole error of the cheap run into every delta.
//
// Every request in one stage shares a seed. The engine advances its
// seed once per iteration, so two requests with one seed see the same
// rolls, and the difference between them is a PAIRED comparison -
// which is what makes a 41-DPS gain detectable at 1,000 iterations
// instead of 10,000. Stages get different seeds so that a combination
// that got a lucky stream once does not keep it.
//
// Every request a stage builds - the equipped baseline included - sets
// NoSample: a stage sim's cast log is never read (contract 10.3), so
// paying the engine's median-iteration replay to build one would be
// pure cost for a report nothing looks at.
import (
	"errors"
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// StageRequests is one rung: what to run, and what each run means.
type StageRequests struct {
	Stage      int `json:"stage"`
	Iterations int `json:"iterations"`
	// Requests[0] is ALWAYS the equipped set; Requests[1:] are
	// parallel to Combos.
	Requests []api.SimRequest `json:"requests"`
	Combos   []Combination    `json:"combos"`
}

// Plan is the first stage, with no build tables. See PlanWith.
func Plan(req api.SimRequest) (StageRequests, error) {
	return PlanWith(req, Options{})
}

// PlanWith is the first stage: every combination plus the equipped set,
// each as a request at the first rung's iteration count.
func PlanWith(req api.SimRequest, opt Options) (StageRequests, error) {
	combos, err := ExpandWith(req, opt)
	if err != nil {
		return StageRequests{}, err
	}
	if len(combos) == 0 {
		return StageRequests{}, errors.New("bulk: nothing to rank; every candidate was for another class, another faction, a locked slot, or a slot it does not fit")
	}
	ladder, ok := api.Ladders[req.Bulk.Precision]
	if !ok || len(ladder.Iterations) == 0 {
		return StageRequests{}, fmt.Errorf("bulk: no ladder for precision %q", req.Bulk.Precision)
	}
	return stageRequests(req, 1, ladder.Iterations[0], combos), nil
}

// stageRequests builds one stage's runnable requests. Rank calls it for
// every rung after the first.
func stageRequests(req api.SimRequest, stage, iterations int, combos []Combination) StageRequests {
	// A stage's seed is the request's plus the stage number, so a
	// re-run of a saved request reproduces the whole ladder and no two
	// stages share a stream.
	seed := req.RandomSeed + int64(stage)

	equipped := req
	equipped.Bulk = nil
	equipped.Weights = nil
	equipped.TargetError = 0
	equipped.Iterations = iterations
	equipped.RandomSeed = seed
	// A stage request is a plain run by shape (contract 10.3): its
	// cast log is never read, so it must say so or every stage sim -
	// including this baseline - pays the engine's sample replay for
	// nothing.
	equipped.NoSample = true

	out := StageRequests{
		Stage:      stage,
		Iterations: iterations,
		Requests:   make([]api.SimRequest, 0, len(combos)+1),
		Combos:     combos,
	}
	out.Requests = append(out.Requests, equipped)
	for _, c := range combos {
		r := c.Request
		r.Iterations = iterations
		r.RandomSeed = seed
		r.NoSample = true
		out.Requests = append(out.Requests, r)
	}
	return out
}
