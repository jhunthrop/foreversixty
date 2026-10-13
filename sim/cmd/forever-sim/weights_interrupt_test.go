package main

import (
	"errors"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// A SIGINT that lands between two sub-sims of a weights sweep, when no engine sim is running
// to take the abort, still stops the sweep (task L6.12, the L6.11 review's D1): the next
// sub-sim never starts and the run answers aborted, as an interrupt during a sub-sim does.
// The signal is real: it is sent to this process at the gap after the baseline.
func TestExecuteWeightsStopsOnASignalBetweenSubSims(t *testing.T) {
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

	// The test's own handler, so the process is never killed by the signal it sends itself.
	own := make(chan os.Signal, 1)
	signal.Notify(own, os.Interrupt)
	defer signal.Stop(own)
	gaps := 0
	afterSubSim = func() {
		gaps++
		if gaps != 1 {
			return
		}
		self, err := os.FindProcess(os.Getpid())
		if err == nil {
			err = self.Signal(os.Interrupt)
		}
		if err != nil {
			t.Fatalf("raising SIGINT: %v", err)
		}
		<-own
		// executeWeights' own handler has the same signal; give it the time it takes to act
		// on it before the next sub-sim would start.
		time.Sleep(200 * time.Millisecond)
	}
	defer func() { afterSubSim = nil }()

	res, err := executeWeights(req, nil)
	if !errors.Is(err, adapter.ErrAborted) {
		t.Fatalf("executeWeights = %v (aborted %v, %d weights), want adapter.ErrAborted: the signal between sub-sims was lost",
			err, res.Aborted, len(res.Weights))
	}
	if !res.Aborted || res.Weights != nil {
		t.Errorf("result aborted %v with weights %+v, want aborted and none", res.Aborted, res.Weights)
	}
	if gaps != 1 {
		t.Errorf("%d sub-sims ended, want only the baseline: a sub-sim started after the signal", gaps)
	}
}
