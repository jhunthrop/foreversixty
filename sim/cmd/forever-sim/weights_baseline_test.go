package main

import (
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/request"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
)

// A weights run answers with DPS too (task L6.11 fix 1): the result's DPS and error are the
// sweep's own baseline sim, the same figures the baseline request the engine builds for that
// sweep gives when it is run on its own, and the summary is that baseline's, not the empty one.
// Through the real engine; the seed is fixed so the baseline can be run again to compare.
func TestAWeightsRunReportsItsBaselineDPS(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	req := fixtureRequest(t, "warrior-fury")
	req.EngineVersion = enginever.Version
	req.Iterations = 500
	req.RandomSeed = 4242
	req.Encounter.DurationSec = 60
	req.Weights = &api.WeightsSpec{Stats: []string{"attack_power", "crit"}, Reference: "attack_power"}

	res, err := executeWeights(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Weights) != 2 {
		t.Fatalf("weights = %+v, want two", res.Weights)
	}
	if res.DPS.Mean <= 0 || res.DPS.Error <= 0 {
		t.Fatalf("a weights run's DPS = %+v, want the baseline's mean and error", res.DPS)
	}
	if len(res.Summary.DamageDone) == 0 {
		t.Error("a weights result carries the empty summary; it must carry the baseline's")
	}

	// The baseline on its own: the request the engine builds for the same sweep.
	registerOnce.Do(engine.RegisterAll)
	engineReq, err := request.BuildWeights(req, request.Options{OpenIterations: true})
	if err != nil {
		t.Fatal(err)
	}
	if err := simdb.AttachWeights(engineReq); err != nil {
		t.Fatal(err)
	}
	base := core.RunRaidSimConcurrent(core.StatWeightRequests(engineReq).BaseRequest)
	if err := adapter.ResultError(base); err != nil {
		t.Fatal(err)
	}
	want := adapter.DPS(base)
	if math.Abs(res.DPS.Mean-want.Mean) > 1e-6 || math.Abs(res.DPS.Error-want.Error) > 1e-6 {
		t.Errorf("weights run DPS %+v, want the baseline's %+v", res.DPS, want)
	}
}
