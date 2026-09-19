package sims

import (
	"context"
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

func (h *harness) jobDeps(engine runner.StageRunner) JobDeps {
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

func TestAnAbortedRunIsFailedNotStoredAsDone(t *testing.T) {
	h := newHarness(t)
	h.queued(t, "dddddddddddd", "warrior-fury")
	// runner.Fixture{Aborted: true} answers the way Native does on
	// forever-sim's exit 130: a partial result plus a wrapped
	// runner.ErrAborted. The job must never store that partial result
	// as a finished run.
	err := Run(t.Context(), h.jobDeps(&runner.Fixture{Aborted: true}), "dddddddddddd")
	if !errors.Is(err, runner.ErrAborted) {
		t.Fatalf("err = %v, want the abort to reach the caller", err)
	}

	p, perr := h.store.Progress(t.Context(), "dddddddddddd")
	if perr != nil {
		t.Fatal(perr)
	}
	if p.State != StateError {
		t.Fatalf("state %q, want %q: a partial result must never read as done", p.State, StateError)
	}

	stored, gerr := h.store.Get(t.Context(), "dddddddddddd")
	if gerr != nil {
		t.Fatal(gerr)
	}
	if stored.DPS.Mean != 0 || stored.IterationsRun != 0 {
		t.Fatalf("the aborted run's partial result must not be written: %+v", stored)
	}
}

func TestAnUnknownSimIsAnErrorAndWritesNothing(t *testing.T) {
	h := newHarness(t)
	if err := Run(t.Context(), h.jobDeps(&runner.Fixture{}), "zzzzzzzzzzzz"); err != ErrNotFound {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}

// bulkRunner answers the way forever-sim does for a bulk request: it
// streams a tick per stage, then returns a ranked result. The binary
// detects the kind and runs sim/bulk's plan-rank loop itself, so the
// job hands it the request whole and reads what comes back — which is
// exactly what this stands in for.
//
// It implements runner.StageRunner: Run answers plainly (no test here
// exercises it), and RunStaged is what the job actually calls, since
// the sim lane's widening (contract 10.8) put the stage fields on
// simapi.Progress behind StageRunner.RunStaged rather than on
// runner.Progress itself.
type bulkRunner struct{ got simapi.SimRequest }

func (b *bulkRunner) Run(ctx context.Context, req simapi.SimRequest,
	onProgress runner.Progress) (simapi.SimResult, error) {
	return b.RunStaged(ctx, req, nil)
}

func (b *bulkRunner) RunStaged(_ context.Context, req simapi.SimRequest,
	onProgress runner.StageProgress) (simapi.SimResult, error) {
	b.got = req
	if onProgress != nil {
		onProgress(simapi.Progress{
			IterationsRun: 1000, DPS: simapi.Estimate{Mean: 1000},
			Stage: 1, CombosDone: 40, CombosTotal: 96,
		})
		onProgress(simapi.Progress{
			IterationsRun: 4000, DPS: simapi.Estimate{Mean: 1040},
			Stage: 2, CombosDone: 10, CombosTotal: 10,
		})
	}
	return simapi.SimResult{
		EngineVersion: req.EngineVersion, Request: req, Lane: simapi.LaneServer,
		DPS:           simapi.Estimate{Mean: 1042.5, Error: 1.6},
		IterationsRun: 4000, DurationMS: 90_000,
		Equipped: &simapi.Estimate{Mean: 1001.3},
		Stages:   []simapi.Stage{{Iterations: 1000, Combos: 96}, {Iterations: 3000, Combos: 10}},
		Combos: []simapi.Combo{{
			Substitutions: []simapi.Substitution{
				{Kind: "item", Slot: "main_hand", ItemID: 17182,
					Name: "Vis'kag the Bloodletter", Origin: "bag"},
			},
			DPS:   simapi.Estimate{Mean: 1042.5},
			Delta: simapi.Estimate{Mean: 41.2, Error: 2.2},
		}},
	}, nil
}

// queuedBulk puts a Top Gear run in the queued state.
func (h *harness) queuedBulk(t *testing.T, id string) simapi.SimRequest {
	t.Helper()
	req := browserResult("warrior-fury", 0).Request
	req.Bulk = &simapi.BulkSpec{
		Mode: simapi.KindGear, Precision: simapi.PrecisionNormal,
		Cap:        simapi.Caps[simapi.LaneServer],
		Candidates: []simapi.Candidate{{Slot: "main_hand", ItemID: 17182, Origin: "bag"}},
	}
	if err := h.store.Queue(t.Context(), id, h.owner, req); err != nil {
		t.Fatal(err)
	}
	return req
}

func TestABulkJobStreamsItsStagesAndStoresTheRanking(t *testing.T) {
	h := newHarness(t)
	want := h.queuedBulk(t, "ffffffffffff")
	engine := &bulkRunner{}

	if err := Run(t.Context(), h.jobDeps(engine), "ffffffffffff"); err != nil {
		t.Fatal(err)
	}

	// The whole request reached the binary, bulk block and all: the
	// binary detects the kind and runs the plan-rank loop itself.
	if engine.got.Bulk == nil || engine.got.Bulk.Mode != want.Bulk.Mode ||
		len(engine.got.Bulk.Candidates) != 1 {
		t.Fatalf("the bulk block did not reach the engine: %+v", engine.got.Bulk)
	}

	p, err := h.store.Progress(t.Context(), "ffffffffffff")
	if err != nil {
		t.Fatal(err)
	}
	if p.State != StateDone {
		t.Fatalf("state %q", p.State)
	}
	// The last tick is what the row remembers.
	if p.Stage != 2 || p.CombosDone != 10 || p.CombosTotal != 10 {
		t.Errorf("progress %+v, want the final stage's counts", p)
	}

	stored, err := h.store.Get(t.Context(), "ffffffffffff")
	if err != nil {
		t.Fatal(err)
	}
	if len(stored.Combos) != 1 || stored.Equipped == nil || len(stored.Stages) != 2 {
		t.Fatalf("the ranking did not survive storage: %+v", stored)
	}

	// And the history row says what it found.
	page, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindGear)
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Rows) != 1 || page.Rows[0].Headline != "+41 DPS from Vis'kag the Bloodletter" {
		t.Fatalf("history row: %+v", page.Rows)
	}
}

func TestAPlainRunReportsNoStages(t *testing.T) {
	h := newHarness(t)
	h.queued(t, "gggggggggggg", "warrior-fury")
	if err := Run(t.Context(), h.jobDeps(&runner.Fixture{}), "gggggggggggg"); err != nil {
		t.Fatal(err)
	}
	p, err := h.store.Progress(t.Context(), "gggggggggggg")
	if err != nil {
		t.Fatal(err)
	}
	if p.Stage != 0 || p.CombosDone != 0 || p.CombosTotal != 0 {
		t.Errorf("a plain run reported stages: %+v", p)
	}
}
