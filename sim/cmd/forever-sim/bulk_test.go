package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/bulk"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/wowsims/classic/sim/core/simsignals"
)

// bulkRequest is the checked-in warrior fixture with two candidates,
// at the fewest iterations that still answers: the binary's job here
// is the LOOP, and the engine's numbers are tested in sim/adapter.
func bulkRequest(t *testing.T) api.SimRequest {
	t.Helper()
	req := fixtureRequest(t, "warrior-fury")
	req.EngineVersion = enginever.Version
	req.Iterations = 3000
	req.Encounter.DurationSec = 60
	req.Bulk = &api.BulkSpec{
		Mode:      api.KindGear,
		Precision: api.PrecisionNormal,
		Cap:       api.Caps[api.LaneServer],
		Candidates: []api.Candidate{
			{Slot: "head", ItemID: req.Character.Gear[0].ItemID, Origin: api.OriginBag},
			{Slot: "neck", ItemID: req.Character.Gear[1].ItemID, Origin: api.OriginBag},
		},
	}
	return req
}

// The whole loop, end to end, at the smallest useful size. It is slow
// (two stages of three sims), so it is one test rather than a table.
func TestExecuteBulkRunsEveryStage(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	var progress bytes.Buffer
	req := bulkRequest(t)
	res, err := executeBulk(req, bulk.Options{}, &progress)
	if err != nil {
		t.Fatal(err)
	}
	if res.Equipped == nil || res.Equipped.Mean <= 0 {
		t.Fatalf("equipped = %+v", res.Equipped)
	}
	if len(res.Combos) != 3 {
		t.Errorf("got %d combos, want the head, the neck and both", len(res.Combos))
	}
	if len(res.Stages) != 2 || res.Stages[0].Iterations != 1000 || res.Stages[1].Iterations != 3000 {
		t.Errorf("stages = %+v", res.Stages)
	}
	if res.Request.Bulk == nil {
		t.Error("the result lost the request that produced it")
	}
	if res.Lane != api.LaneServer {
		t.Errorf("lane = %q", res.Lane)
	}
	if res.EngineVersion != enginever.Version {
		t.Errorf("engine_version = %q", res.EngineVersion)
	}
	// A within-error group is assigned for every row, and the leader
	// is group 0.
	if res.Combos[0].Group != 0 {
		t.Errorf("the leader is group %d", res.Combos[0].Group)
	}

	// Every stderr line is JSON, and the bulk ones carry the stage.
	var stages, combos int
	for _, line := range strings.Split(strings.TrimSpace(progress.String()), "\n") {
		if line == "" {
			continue
		}
		var tick struct {
			Stage       int `json:"stage"`
			CombosDone  int `json:"combos_done"`
			CombosTotal int `json:"combos_total"`
			Log         string
		}
		if err := json.Unmarshal([]byte(line), &tick); err != nil {
			t.Fatalf("a progress line is not JSON: %s", line)
		}
		if tick.Stage > stages {
			stages = tick.Stage
		}
		if tick.CombosTotal > combos {
			combos = tick.CombosTotal
		}
	}
	if stages != 2 {
		t.Errorf("progress reported %d stages, want 2", stages)
	}
	if combos == 0 {
		t.Error("progress never reported a combination count")
	}
}

func TestExecuteBulkRefusesAPlainRun(t *testing.T) {
	if _, err := executeBulk(fixtureRequest(t, "warrior-fury"), bulk.Options{}, nil); err == nil {
		t.Error("executeBulk accepted a request with no bulk block")
	}
}

// The invariant that lets -plan answer "how many combinations" from one
// bulk.PlanWith call instead of a second bulk.CountWith call:
// PlanWith's stage.Combos IS ExpandWith's own slice, untrimmed, so its
// length already answers what CountWith would. Pinned here so a future
// PlanWith that filtered or trimmed Combos would fail this test rather
// than silently making -plan disagree with simCount (contract 10.2's
// one rule: the two must never disagree).
func TestPlanCombosMatchCount(t *testing.T) {
	req := bulkRequest(t)
	stage, err := bulk.PlanWith(req, bulk.Options{})
	if err != nil {
		t.Fatal(err)
	}
	count, err := bulk.CountWith(req, bulk.Options{})
	if err != nil {
		t.Fatal(err)
	}
	if len(stage.Combos) != count {
		t.Errorf("stage.Combos has %d, bulk.CountWith says %d", len(stage.Combos), count)
	}
}

// abort mid-stage-2 must report BOTH stages: stage 1, which finished,
// and stage 2, which was cut short. abortedBulk used to report only
// the stage it was standing in when the signal arrived, so a run
// stopped in stage 2 claimed the whole ladder was "1 stage, 3000
// iterations" - dropping the 1000-iteration stage 1 that already
// completed. SimResult.Stages is documented as "what the planner
// actually ran" and the api lane's progress reader keys off it.
func TestExecuteBulkAbortedMidStageKeepsEarlierStages(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	req := bulkRequest(t)
	stage1, err := bulk.PlanWith(req, bulk.Options{})
	if err != nil {
		t.Fatal(err)
	}
	// Stage 1 runs len(stage1.Requests) sims (the equipped set plus
	// every combination); the next one runID hands out is stage 2's
	// first. Aborting exactly that id stops the run partway into stage
	// 2, which is the case that exposes a truncated Ran history.
	target := fmt.Sprintf("forever-sim-%d-%d", os.Getpid(), runs.Load()+int64(len(stage1.Requests))+1)
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 20000; i++ {
			if simsignals.AbortById(target) {
				return
			}
			time.Sleep(time.Millisecond)
		}
	}()

	res, err := executeBulk(req, bulk.Options{}, nil)
	<-done
	if !errors.Is(err, adapter.ErrAborted) {
		t.Fatalf("executeBulk = %v, want adapter.ErrAborted", err)
	}
	if !res.Aborted {
		t.Error("the result does not say it was stopped")
	}
	if len(res.Stages) != 2 {
		t.Fatalf("Stages = %+v, want stage 1 (completed) and stage 2 (partial)", res.Stages)
	}
	if res.Stages[0].Iterations != stage1.Iterations || res.Stages[0].Combos != len(stage1.Combos) {
		t.Errorf("Stages[0] = %+v, want the completed stage 1 (%d iterations, %d combos)", res.Stages[0], stage1.Iterations, len(stage1.Combos))
	}
}

// -plan is how the API counts a bulk request without importing
// sim/internal. It runs nothing, so it is fast enough for a submit
// handler.
func TestPlanOnlyRunsNothing(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "request.json")
	out := filepath.Join(dir, "plan.json")
	body, err := json.Marshal(bulkRequest(t))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(in, body, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := run(in, out, overrides{}, true, nil); err != nil {
		t.Fatal(err)
	}
	var got struct {
		Combinations int `json:"combinations"`
		Stage        struct {
			Stage      int `json:"stage"`
			Iterations int `json:"iterations"`
		} `json:"stage"`
	}
	b, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("-plan wrote %s", b)
	}
	if got.Combinations != 3 || got.Stage.Stage != 1 || got.Stage.Iterations != 1000 {
		t.Errorf("-plan wrote %+v", got)
	}
}

// A cap breach is the same structured answer the wasm gives, and a
// non-zero exit, so the API can answer 400 cap_exceeded from either
// lane without parsing prose.
func TestPlanOnlyReportsACapBreach(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "request.json")
	out := filepath.Join(dir, "plan.json")
	req := bulkRequest(t)
	req.Bulk.Cap = 1
	body, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(in, body, 0o644); err != nil {
		t.Fatal(err)
	}
	err = run(in, out, overrides{}, true, nil)
	if !errors.Is(err, errBadInput) {
		t.Fatalf("run = %v, want errBadInput", err)
	}
	var got struct {
		Error        string `json:"error"`
		Cap          int    `json:"cap"`
		Combinations int    `json:"combinations"`
	}
	b, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("-plan wrote %s", b)
	}
	if got.Error != "cap_exceeded" || got.Cap != 1 || got.Combinations <= 1 {
		t.Errorf("-plan wrote %+v", got)
	}
}
