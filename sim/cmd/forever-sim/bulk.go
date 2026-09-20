package main

// The native plan-rank loop.
//
// This is the same three functions the browser calls - bulk.Plan, a
// run per request, bulk.Rank - in a loop instead of across a worker
// pool. Nothing about WHICH sims run or HOW they are ranked lives
// here; that is sim/bulk's, compiled into both artifacts, which is
// what stops the premium lane and the free lane disagreeing about who
// won.
//
// Concurrency is inside each run: core.RunRaidSimConcurrentAsync
// already splits one sim across every CPU. Running the stage's sims
// in parallel on top of that would oversubscribe the 4-CPU job and
// make the progress line meaningless. The browser does the mirror
// image - whole, unsplit runs across the worker pool - for the same
// reason (contract 10.2).

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/bulk"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// executeBulk runs every stage of a bulk request and returns the
// ranked result.
func executeBulk(req api.SimRequest, opt bulk.Options, progress io.Writer) (api.SimResult, error) {
	start := time.Now()
	stage, err := bulk.PlanWith(req, opt)
	if err != nil {
		return api.SimResult{}, fmt.Errorf("%w: %v", errBadInput, err)
	}
	for {
		results := make([]api.SimResult, 0, len(stage.Requests))
		for i, r := range stage.Requests {
			// The stage's own progress, written before each sim rather
			// than after, so a long stage says what it is doing rather
			// than going quiet and then jumping.
			writeStageTick(progress, stage, i)
			// The per-sim ticks are suppressed: one line per iteration
			// batch times ninety-six combinations is a stderr stream
			// nobody can read, and the api lane's reader wants the
			// stage counts.
			res, err := Execute(r, nil)
			if err != nil {
				if errors.Is(err, adapter.ErrAborted) {
					// A stopped bulk run returns what it has: the
					// stage it reached, marked partial, the way a
					// stopped single run does.
					return abortedBulk(req, stage, start), err
				}
				return api.SimResult{}, err
			}
			results = append(results, res)
		}
		writeStageTick(progress, stage, len(stage.Requests))

		next, final, err := bulk.Rank(req, stage, results)
		if err != nil {
			return api.SimResult{}, err
		}
		if final != nil {
			final.EngineVersion = enginever.Version
			final.Lane = api.LaneServer
			final.DurationMS = time.Since(start).Milliseconds()
			return *final, nil
		}
		stage = *next
	}
}

// writeStageTick writes one bulk progress line. done counts the
// requests of this stage that have finished, and the equipped set is
// one of them - the page says "31 of 96" about the whole stage, not
// about the combinations alone, because that is what the bar fills to.
func writeStageTick(progress io.Writer, stage bulk.StageRequests, done int) {
	if progress == nil {
		return
	}
	_ = json.NewEncoder(progress).Encode(struct {
		Completed   int     `json:"completed"`
		Total       int     `json:"total"`
		DPS         float64 `json:"dps"`
		Stage       int     `json:"stage"`
		CombosDone  int     `json:"combos_done"`
		CombosTotal int     `json:"combos_total"`
	}{
		Completed:   done * stage.Iterations,
		Total:       len(stage.Requests) * stage.Iterations,
		Stage:       stage.Stage,
		CombosDone:  done,
		CombosTotal: len(stage.Requests),
	})
}

// abortedBulk is what a stopped bulk run writes: the request, the
// stage it reached, and nothing pretending to be a ranking.
func abortedBulk(req api.SimRequest, stage bulk.StageRequests, start time.Time) api.SimResult {
	return api.SimResult{
		EngineVersion: enginever.Version,
		Request:       req,
		Lane:          api.LaneServer,
		Aborted:       true,
		DurationMS:    time.Since(start).Milliseconds(),
		Summary:       adapter.EmptySummary(),
		Stages:        []api.Stage{{Iterations: stage.Iterations, Combos: len(stage.Combos)}},
	}
}
