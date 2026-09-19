package sims

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"os"
	"testing"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/runner"
)

// queued puts a server-lane run in the queued state, exactly as
// POST /v1/sims/run leaves it.
func (h *harness) queued(t *testing.T, id, spec string) simapi.SimRequest {
	t.Helper()
	req := browserResult(spec, 0).Request
	req.Iterations = preciseIterations
	if err := h.store.Queue(t.Context(), id, h.owner, req); err != nil {
		t.Fatal(err)
	}
	return req
}

func (h *harness) jobDeps(engine runner.Runner) JobDeps {
	return JobDeps{
		Store: h.store, Put: h.files, Engine: engine,
		Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

func TestTheRunJobStreamsProgressAndWritesTheResultToBothPlaces(t *testing.T) {
	h := newHarness(t)
	h.queued(t, "aaaaaaaaaaaa", "warrior-fury")

	if err := Run(t.Context(), h.jobDeps(&runner.Fixture{}), "aaaaaaaaaaaa"); err != nil {
		t.Fatal(err)
	}

	p, err := h.store.Progress(t.Context(), "aaaaaaaaaaaa")
	if err != nil {
		t.Fatal(err)
	}
	if p.State != StateDone || p.DPS == nil || *p.DPS == 0 {
		t.Fatalf("progress: %+v", p)
	}
	if p.IterationsDone != preciseIterations {
		t.Errorf("iterations %d, want the whole run", p.IterationsDone)
	}

	stored, err := h.store.Get(t.Context(), "aaaaaaaaaaaa")
	if err != nil {
		t.Fatal(err)
	}
	if stored.Lane != simapi.LaneServer || stored.SimID != "aaaaaaaaaaaa" {
		t.Fatalf("stored: lane %q id %q", stored.Lane, stored.SimID)
	}
	if len(stored.Summary.DamageDone) == 0 {
		t.Error("the stored result has no summary; the report components have nothing to render")
	}

	// And the whole thing is in the bucket too.
	b, err := os.ReadFile(h.dir + "/" + Keys{SimID: "aaaaaaaaaaaa"}.Result())
	if err != nil {
		t.Fatal(err)
	}
	var fromBucket simapi.SimResult
	if err := json.Unmarshal(b, &fromBucket); err != nil {
		t.Fatal(err)
	}
	if fromBucket.DPS.Mean != stored.DPS.Mean {
		t.Errorf("bucket %v, row %v", fromBucket.DPS.Mean, stored.DPS.Mean)
	}
}

func TestTheRunJobHandsTheEngineTheQueuedRequest(t *testing.T) {
	h := newHarness(t)
	want := h.queued(t, "bbbbbbbbbbbb", "mage-frost")
	engine := &runner.Fixture{}
	if err := Run(t.Context(), h.jobDeps(engine), "bbbbbbbbbbbb"); err != nil {
		t.Fatal(err)
	}
	asked := engine.Asked()
	if len(asked) != 1 {
		t.Fatalf("%d runs", len(asked))
	}
	got := asked[0]
	if got.Spec != want.Spec || got.Iterations != want.Iterations {
		t.Errorf("request %+v, want %+v", got, want)
	}
	if got.Character.Race != want.Character.Race || len(got.Character.Gear) != len(want.Character.Gear) {
		t.Errorf("the character did not reach the engine: %+v", got.Character)
	}
}

func TestARunTheEngineRefusesIsRecordedAsFailed(t *testing.T) {
	h := newHarness(t)
	h.queued(t, "cccccccccccc", "warrior-fury")
	err := Run(t.Context(), h.jobDeps(&runner.Fixture{Err: errAnyway}), "cccccccccccc")
	if err == nil {
		t.Fatal("the job should exit non-zero when the engine fails")
	}
	p, perr := h.store.Progress(t.Context(), "cccccccccccc")
	if perr != nil {
		t.Fatal(perr)
	}
	if p.State != StateError {
		t.Fatalf("state %q, want %q: the page must stop polling", p.State, StateError)
	}
}

func TestARequestTheEngineRefusesIsRecordedAndNotRetried(t *testing.T) {
	h := newHarness(t)
	h.queued(t, "eeeeeeeeeeee", "warrior-fury")
	// Exit 2 from forever-sim: an unknown consumable id, a character
	// sim/request cannot build. runner.Native turns it into this.
	err := Run(t.Context(), h.jobDeps(&runner.Fixture{Err: runner.ErrBadInput}), "eeeeeeeeeeee")
	if !errors.Is(err, runner.ErrBadInput) {
		t.Fatalf("err = %v, want the engine's refusal to reach the caller", err)
	}
	p, perr := h.store.Progress(t.Context(), "eeeeeeeeeeee")
	if perr != nil {
		t.Fatal(perr)
	}
	if p.State != StateError {
		t.Fatalf("state %q, want %q", p.State, StateError)
	}
}

func TestAnUnknownSimIsAnErrorAndWritesNothing(t *testing.T) {
	h := newHarness(t)
	if err := Run(t.Context(), h.jobDeps(&runner.Fixture{}), "zzzzzzzzzzzz"); err != ErrNotFound {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}
