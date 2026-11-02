package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func memoTestRequest(name string, item, seed int) api.SimRequest {
	return api.SimRequest{
		Spec:       "warrior-fury",
		Iterations: 300,
		RandomSeed: int64(seed),
		Character:  api.CharacterSpec{Name: name, Level: 60, Gear: []api.GearSlot{{Slot: "head", ItemID: item}}},
	}
}

func TestTheMemoSimulatesARepeatedRequestOnce(t *testing.T) {
	fake := &fakeEngine{DPSByGear: map[string]float64{"head=1;": 100, "head=2;": 120}, DefaultStdErr: 0.5}
	memo := newMemoRunner(fake)
	for _, name := range []string{"verify", "trinket-rank", "verify"} {
		mean, stdErr, err := memo.RunPlainDPSWithError(memoTestRequest(name, 1, 7))
		if err != nil || mean != 100 || stdErr != 0.5 {
			t.Fatalf("%s: got %v ± %v (%v), want 100 ± 0.5", name, mean, stdErr, err)
		}
	}
	if len(fake.Calls) != 1 {
		t.Errorf("the engine ran %d times for one request under three labels, want 1", len(fake.Calls))
	}
	if mean, _ := memo.RunPlainDPS(memoTestRequest("verify", 2, 7)); mean != 120 {
		t.Errorf("other gear = %v, want its own 120", mean)
	}
	if _, err := memo.RunPlainDPS(memoTestRequest("verify", 1, 8)); err != nil || len(fake.Calls) != 3 {
		t.Errorf("another seed ran %d sims in all, want 3 (a new seed is a new sim)", len(fake.Calls))
	}
	if hits, misses := memo.stats(); hits != 2 || misses != 3 {
		t.Errorf("hits %d misses %d, want 2 and 3", hits, misses)
	}
}

func TestTheMemoNeverReusesAnUnseededRequest(t *testing.T) {
	fake := &fakeEngine{DefaultDPS: 50}
	memo := newMemoRunner(fake)
	for range 2 {
		if _, err := memo.RunPlainDPS(memoTestRequest("verify", 1, 0)); err != nil {
			t.Fatal(err)
		}
	}
	if len(fake.Calls) != 2 {
		t.Errorf("an unseeded request ran %d times in two calls, want 2: it draws a fresh seed each run", len(fake.Calls))
	}
}

func TestTheMemoReusesAWeightsSweep(t *testing.T) {
	fake := &fakeEngine{WeightsResult: map[string]api.StatWeight{"stamina": {Stat: "stamina", Weight: 1}}, ReferenceDPSPerPoint: 0.25}
	memo := newMemoRunner(fake)
	for range 2 {
		w, ref, err := memo.RunWeights(memoTestRequest("ladder", 1, 3))
		if err != nil || ref != 0.25 || w["stamina"].Weight != 1 {
			t.Fatalf("got %v %v %v", w, ref, err)
		}
	}
	if len(fake.WeightsIterationsSeen) != 1 {
		t.Errorf("the sweep ran %d times, want 1", len(fake.WeightsIterationsSeen))
	}
}

// A tank band's figures and a healer's fields come from interfaces only
// the role runner has; the memo must not hide them.
func TestTheMemoKeepsTheRoleRunnerReachable(t *testing.T) {
	tank := roleRunner(&fakeEngine{}, specInfo{Role: roleTank})
	if _, ok := unwrapRunner(newMemoRunner(tank)).(tankRunner); !ok {
		t.Error("a memoized tank runner is no longer a tankRunner")
	}
	if got := unwrapRunner(tank); got != tank {
		t.Error("unwrapRunner changed a runner with no memo")
	}
}

// slowEngine counts its plain sims and takes a moment over each, so
// concurrent callers overlap.
type slowEngine struct {
	fakeEngine
	runs atomic.Int64
}

func (s *slowEngine) RunPlainDPSWithError(req api.SimRequest) (float64, float64, error) {
	s.runs.Add(1)
	time.Sleep(20 * time.Millisecond)
	return float64(req.Character.Gear[0].ItemID), 0, nil
}

// Both weight sets of a band rank at once (band_pipeline.go); a set both
// reach at the same moment is simulated once and both get its answer.
func TestTheMemoRunsAnInFlightRequestOnceForConcurrentCallers(t *testing.T) {
	engine := &slowEngine{}
	memo := newMemoRunner(engine)
	var wg sync.WaitGroup
	got := make([]float64, 8)
	for i := range got {
		wg.Add(1)
		go func() {
			defer wg.Done()
			got[i], _ = memo.RunPlainDPS(memoTestRequest(fmt.Sprintf("caller-%d", i), 5, 7))
		}()
	}
	wg.Wait()
	if n := engine.runs.Load(); n != 1 {
		t.Errorf("eight concurrent callers ran %d sims, want 1", n)
	}
	for i, v := range got {
		if v != 5 {
			t.Errorf("caller %d got %v, want 5", i, v)
		}
	}
}
