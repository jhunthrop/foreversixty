package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/wowsims/classic/sim/core/simsignals"
)

// abortRunID predicts the id runID() will hand out to the nth run this
// test starts from here (1 = the very next one), the same trick
// TestExecuteBulkAbortedMidStageKeepsEarlierStages uses: runID's
// counter is a package-level atomic.Int64, so its next value is knowable
// before the run that will register it even starts.
func abortRunID(n int64) string {
	return fmt.Sprintf("forever-sim-%d-%d", os.Getpid(), runs.Load()+n)
}

// abortAfterRegistration polls simsignals.AbortById(id) until it
// succeeds (meaning the target run has registered and is now signalled
// to stop) or the timeout elapses. Call sites should <-done before
// asserting, so the goroutine can't outlive its test.
func abortAfterRegistration(id string) (done chan struct{}) {
	done = make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 20000; i++ {
			if simsignals.AbortById(id) {
				return
			}
			time.Sleep(time.Millisecond)
		}
	}()
	return done
}

// A fixed-count run is one sim: the loop must not turn today's
// requests into several.
func TestExecuteToTargetIsOneRunWithoutATarget(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	req := fixtureRequest(t, "warrior-fury")
	req.EngineVersion = enginever.Version
	req.Iterations = 500
	req.Encounter.DurationSec = 60
	res, err := ExecuteToTarget(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.IterationsRun != 500 {
		t.Errorf("iterations_run = %d, want exactly the 500 asked for", res.IterationsRun)
	}
}

// A target-error run steps until the error is inside the target or the
// ceiling is reached, and reports the total.
func TestExecuteToTargetSteps(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	req := fixtureRequest(t, "warrior-fury")
	req.EngineVersion = enginever.Version
	req.Encounter.DurationSec = 60
	// A target nothing reaches, so the ceiling ends it and the count
	// is exactly the ceiling rather than a step past it.
	req.TargetError = 0.00001
	req.Iterations = 3 * api.StepIterations
	res, err := ExecuteToTarget(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.IterationsRun != 3*api.StepIterations {
		t.Errorf("iterations_run = %d, want the ceiling %d", res.IterationsRun, 3*api.StepIterations)
	}
	if res.DPS.Mean <= 0 || res.DPS.Error <= 0 {
		t.Errorf("dps = %+v", res.DPS)
	}
	if res.Summary.EngineVersion == "" {
		t.Error("a stepped run lost its summary")
	}

	// A target anything reaches stops early, so the two runs differ.
	req.TargetError = 0.5
	early, err := ExecuteToTarget(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if early.IterationsRun != api.StepIterations {
		t.Errorf("a slack target ran %d iterations, want one step", early.IterationsRun)
	}
}

func TestExecuteWeights(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	req := fixtureRequest(t, "warrior-fury")
	req.EngineVersion = enginever.Version
	req.Iterations = 500
	req.Encounter.DurationSec = 60
	req.Weights = &api.WeightsSpec{
		Stats:     []string{"agility", "attack_power", "crit"},
		Reference: "attack_power",
	}
	res, err := executeWeights(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Weights) != 3 {
		t.Fatalf("got %d weights: %+v", len(res.Weights), res.Weights)
	}
	if res.Weights[1].Stat != "attack_power" || res.Weights[1].Weight != 1 {
		t.Errorf("the reference stat is %+v, want a weight of exactly 1", res.Weights[1])
	}
	if res.Lane != api.LaneServer || res.EngineVersion == "" {
		t.Errorf("provenance: lane=%q engine=%q", res.Lane, res.EngineVersion)
	}
	if res.Summary.DamageDone == nil {
		t.Error("a weights result carries a null summary; it must carry the empty one")
	}
	// iterations_run is the engine's cumulative total across the whole
	// sweep (a baseline plus one low/high pair per stat), never the
	// per-sim req.Iterations a caller asked for - so it must be more
	// than that, not equal to it.
	if res.IterationsRun <= req.Iterations {
		t.Errorf("iterations_run = %d, want the sweep's cumulative total, which is more than the per-sim count %d", res.IterationsRun, req.Iterations)
	}
}

// An abort on the very FIRST step has no completed step to fall back
// on, so the step's own aborted result - not a zeroed one - is what
// "how far it got" means. main()'s exit code depends on the error
// being adapter.ErrAborted, not nil, the same way Execute's and
// executeBulk's own abort paths do.
func TestExecuteToTargetAbortOnFirstStepReturnsThePartial(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	req := fixtureRequest(t, "warrior-fury")
	req.EngineVersion = enginever.Version
	req.Encounter.DurationSec = 60
	req.TargetError = 0.00001
	req.Iterations = 3 * api.StepIterations

	done := abortAfterRegistration(abortRunID(1))
	res, err := ExecuteToTarget(req, nil)
	<-done

	if !errors.Is(err, adapter.ErrAborted) {
		t.Fatalf("ExecuteToTarget = %v, want adapter.ErrAborted", err)
	}
	if !res.Aborted {
		t.Error("the result does not say it was stopped")
	}
	if res.EngineVersion == "" || res.Lane != api.LaneServer {
		t.Errorf("an aborted first step still needs its provenance: engine=%q lane=%q", res.EngineVersion, res.Lane)
	}
	if res.Summary.DamageDone == nil {
		t.Error("an aborted step's result carries a null summary; it must carry the empty one")
	}
}

// An abort AFTER at least one step has already finished has real data
// to report: the loop keeps the pooled result of the steps that
// completed and drops the interrupted one rather than merging a
// result Execute never computed a DPS for. That data is still written
// out AND the result says it was stopped - adapter.ErrAborted and
// Aborted: true - the same as every other abort path, rather than
// leaving a reader to infer the stop from TargetError and
// IterationsRun alone.
func TestExecuteToTargetAbortAfterAStepKeepsWhatCompleted(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	req := fixtureRequest(t, "warrior-fury")
	req.EngineVersion = enginever.Version
	req.Encounter.DurationSec = 60
	req.TargetError = 0.00001
	req.Iterations = 3 * api.StepIterations

	// Step 1 is runID n=1; aborting step 2 (n=2) lets step 1 finish and
	// pool before the signal lands.
	done := abortAfterRegistration(abortRunID(2))
	res, err := ExecuteToTarget(req, nil)
	<-done

	if !errors.Is(err, adapter.ErrAborted) {
		t.Fatalf("ExecuteToTarget = %v, want adapter.ErrAborted", err)
	}
	if !res.Aborted {
		t.Error("a stop after a full step completed must still say it was stopped")
	}
	if res.IterationsRun != api.StepIterations {
		t.Errorf("iterations_run = %d, want exactly the one completed step (%d); the interrupted step must not be merged in", res.IterationsRun, api.StepIterations)
	}
	if res.EngineVersion == "" || res.Lane != api.LaneServer {
		t.Errorf("the completed step's provenance must survive the abort: engine=%q lane=%q", res.EngineVersion, res.Lane)
	}
}

// An aborted weights run's ErrorOutcome carries an empty message, so a
// message check would misreport the stop as ErrNoWeights. It must come
// back as {Aborted: true} with adapter.ErrAborted, the same as every
// other lane's abort - not the nil error a message-blind check on the
// happy path would otherwise fall through to.
func TestExecuteWeightsAborted(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	req := fixtureRequest(t, "warrior-fury")
	req.EngineVersion = enginever.Version
	req.Iterations = 3000
	req.Encounter.DurationSec = 60
	req.Weights = &api.WeightsSpec{
		Stats:     []string{"agility", "attack_power", "crit", "hit", "melee_haste"},
		Reference: "attack_power",
	}

	done := abortAfterRegistration(abortRunID(1))
	res, err := executeWeights(req, nil)
	<-done

	if !errors.Is(err, adapter.ErrAborted) {
		t.Fatalf("executeWeights = %v, want adapter.ErrAborted", err)
	}
	if !res.Aborted {
		t.Error("the result does not say it was stopped")
	}
	if res.Weights != nil {
		t.Errorf("an aborted weights run reports no weights, got %+v", res.Weights)
	}
	if res.Summary.DamageDone == nil {
		t.Error("an aborted weights result must still carry the empty summary shape")
	}
}
