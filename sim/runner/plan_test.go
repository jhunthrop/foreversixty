package runner

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// planFixtureRequest loads the checked-in warrior-fury fixture and
// turns it into a drops-mode bulk request with exactly n candidates,
// all the same real, already-equipped item in the same slot.
//
// singleCombinations (drops and talents mode) keeps one Combination
// per placement regardless of duplicates, and a candidate that names
// the item already worn there passes every validity check trivially
// (its enchant is already on record for that slot, and replacing one
// slot with what is already in it cannot create a unique-item or
// two-hand conflict elsewhere) - so this is a request whose
// combination count is exactly n without needing n distinct items
// from the build's table.
func planFixtureRequest(t *testing.T, n int) api.SimRequest {
	t.Helper()
	path := filepath.Join("..", "adapter", "testdata", "warrior-fury.request.json")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var req api.SimRequest
	if err := json.Unmarshal(b, &req); err != nil {
		t.Fatalf("%s is not a SimRequest: %v", path, err)
	}
	req.EngineVersion = enginever.Version
	req.Iterations = 3000 // PrecisionNormal's final stage (api.Ladders)

	candidates := make([]api.Candidate, n)
	for i := range candidates {
		candidates[i] = api.Candidate{Slot: "head", ItemID: 12640, Origin: "drop:1"}
	}
	req.Bulk = &api.BulkSpec{
		Mode:       api.KindDrops,
		Precision:  api.PrecisionNormal,
		Cap:        api.Caps[api.LaneServer],
		Candidates: candidates,
	}
	return req
}

// planBinary builds forever-sim to a temp path outside the worktree,
// the same way every other -plan test in this branch does, and
// returns it.
func planBinary(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "forever-sim")
	cmd := exec.Command("go", "build", "-o", bin, "../cmd/forever-sim")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go build ../cmd/forever-sim: %v\n%s", err, out)
	}
	return bin
}

// At the browser cap exactly, -plan answers a summary, not a breach:
// combinations == cap is inside the lane, not past it.
func TestNativePlanAtTheBrowserCapSucceeds(t *testing.T) {
	bin := planBinary(t)
	req := planFixtureRequest(t, api.Caps[api.LaneBrowser])
	req.Bulk.Cap = api.Caps[api.LaneBrowser]

	summary, err := (&Native{Binary: bin}).Plan(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if summary.Kind != api.KindDrops {
		t.Errorf("kind = %q, want %q", summary.Kind, api.KindDrops)
	}
	if summary.Combinations != api.Caps[api.LaneBrowser] {
		t.Errorf("combinations = %d, want %d", summary.Combinations, api.Caps[api.LaneBrowser])
	}
	if summary.Cap != api.Caps[api.LaneBrowser] {
		t.Errorf("cap = %d, want %d", summary.Cap, api.Caps[api.LaneBrowser])
	}
	want := api.LadderIterations(api.Ladders[api.PrecisionNormal], api.Caps[api.LaneBrowser])
	if summary.IterationsTotal != want {
		t.Errorf("iterations_total = %d, want %d", summary.IterationsTotal, want)
	}
}

// One combination past the browser cap, -plan answers cap_exceeded,
// and Plan turns that body into api.ErrCapExceeded rather than a
// summary - the same type bulk.Count itself returns, provable with
// errors.As.
func TestNativePlanOnePastTheBrowserCapReportsCapExceeded(t *testing.T) {
	bin := planBinary(t)
	req := planFixtureRequest(t, api.Caps[api.LaneBrowser]+1)
	req.Bulk.Cap = api.Caps[api.LaneBrowser]

	_, err := (&Native{Binary: bin}).Plan(context.Background(), req)
	var capped api.ErrCapExceeded
	if !errors.As(err, &capped) {
		t.Fatalf("err = %v, want an api.ErrCapExceeded", err)
	}
	if capped.Cap != api.Caps[api.LaneBrowser] || capped.Combinations != api.Caps[api.LaneBrowser]+1 {
		t.Errorf("capped = %+v", capped)
	}
}

func TestFixturePlanAnswersFromConfiguration(t *testing.T) {
	req := planFixtureRequest(t, 3) // the candidate count is irrelevant to the fixture
	f := &Fixture{PlanCombinations: 96}
	summary, err := f.Plan(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if summary.Kind != api.KindDrops || summary.Combinations != 96 || summary.Cap != req.Bulk.Cap {
		t.Errorf("summary = %+v", summary)
	}
	want := api.LadderIterations(api.Ladders[api.PrecisionNormal], 96)
	if summary.IterationsTotal != want {
		t.Errorf("iterations_total = %d, want %d", summary.IterationsTotal, want)
	}
	if len(f.Asked()) != 1 || f.Asked()[0].Bulk.Mode != api.KindDrops {
		t.Fatalf("Plan did not record the request: %+v", f.Asked())
	}
}

func TestFixturePlanCanBeMadeToBreachTheCap(t *testing.T) {
	req := planFixtureRequest(t, 3)
	breach := api.ErrCapExceeded{Cap: 400, Combinations: 401}
	f := &Fixture{CapBreach: &breach}
	_, err := f.Plan(context.Background(), req)
	var capped api.ErrCapExceeded
	if !errors.As(err, &capped) || capped != breach {
		t.Fatalf("err = %v, want %v", err, breach)
	}
}

// Err, when set, wins over CapBreach - the same precedence RunStaged
// gives Err over Aborted - so a test can force a plain failure without
// also having to leave CapBreach nil to prove it took effect.
func TestFixturePlanErrTakesPrecedenceOverCapBreach(t *testing.T) {
	req := planFixtureRequest(t, 3)
	breach := api.ErrCapExceeded{Cap: 400, Combinations: 401}
	f := &Fixture{Err: context.DeadlineExceeded, CapBreach: &breach}
	_, err := f.Plan(context.Background(), req)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v, want context.DeadlineExceeded", err)
	}
}
