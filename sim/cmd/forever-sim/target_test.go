package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

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
}
