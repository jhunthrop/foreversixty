package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/bulk"
	"github.com/jhunthrop/foreversixty/sim/enginever"
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
