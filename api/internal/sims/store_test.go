package sims

import (
	"testing"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

func TestASavedBrowserResultComesBackWhole(t *testing.T) {
	h := newHarness(t)
	res := browserResult("warrior-fury", 1042.5)
	if err := h.store.Save(t.Context(), "aaaaaaaaaaaa", &h.owner, "Tuesday", res); err != nil {
		t.Fatal(err)
	}
	got, err := h.store.Get(t.Context(), "aaaaaaaaaaaa")
	if err != nil {
		t.Fatal(err)
	}
	if got.SimID != "aaaaaaaaaaaa" || got.DPS.Mean != 1042.5 {
		t.Fatalf("round trip: %+v", got)
	}
	if got.Request.Spec != "warrior-fury" {
		t.Errorf("spec %q", got.Request.Spec)
	}
	// The whole request is stored, so a saved sim can be re-run by
	// handing it straight back to the engine.
	if got.Request.Character.Race != "orc" || len(got.Request.Character.Gear) != 2 {
		t.Errorf("character did not survive the round trip: %+v", got.Request.Character)
	}
}

func TestAnUnknownSimIsNotFound(t *testing.T) {
	h := newHarness(t)
	if _, err := h.store.Get(t.Context(), "zzzzzzzzzzzz"); err != ErrNotFound {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
	if _, err := h.store.Progress(t.Context(), "zzzzzzzzzzzz"); err != ErrNotFound {
		t.Fatalf("progress err = %v, want ErrNotFound", err)
	}
}

func TestMyHistoryIsMineAndNewestFirst(t *testing.T) {
	h := newHarness(t)
	for _, c := range []struct {
		id    string
		owner *int64
		mean  float64
	}{
		{"aaaaaaaaaaaa", &h.owner, 900},
		{"bbbbbbbbbbbb", &h.owner, 1000},
		{"cccccccccccc", nil, 1100}, // an anonymous save belongs to nobody
	} {
		if err := h.store.Save(t.Context(), c.id, c.owner, "", browserResult("mage-frost", c.mean)); err != nil {
			t.Fatal(err)
		}
	}
	page, err := h.store.Mine(t.Context(), h.owner, 1)
	if err != nil {
		t.Fatal(err)
	}
	if page.Total != 2 || len(page.Rows) != 2 {
		t.Fatalf("total %d rows %d, want 2 and 2", page.Total, len(page.Rows))
	}
	if page.PerPage != PerPage || page.Page != 1 {
		t.Fatalf("paging: %+v", page)
	}
	for _, r := range page.Rows {
		if r.SimID == "cccccccccccc" {
			t.Error("an anonymous sim appeared in someone's history")
		}
		if r.Spec != "mage-frost" || r.EngineVersion != testEngine {
			t.Errorf("row: %+v", r)
		}
	}
}

func TestAServerRunWalksQueuedThenRunningThenDone(t *testing.T) {
	h := newHarness(t)
	req := browserResult("warrior-fury", 0).Request
	if err := h.store.Queue(t.Context(), "dddddddddddd", h.owner, req); err != nil {
		t.Fatal(err)
	}
	p, err := h.store.Progress(t.Context(), "dddddddddddd")
	if err != nil {
		t.Fatal(err)
	}
	if p.State != StateQueued || p.DPS != nil || p.IterationsDone != 0 {
		t.Fatalf("queued: %+v", p)
	}
	// The queued row carries the whole request: it is what the job
	// reads back, and there is no second copy anywhere.
	queued, err := h.store.Get(t.Context(), "dddddddddddd")
	if err != nil {
		t.Fatal(err)
	}
	if queued.Request.Spec != "warrior-fury" || queued.Request.Character.Class != "warrior" {
		t.Fatalf("queued request: %+v", queued.Request)
	}

	if err := h.store.Advance(t.Context(), "dddddddddddd", 1500, 1000); err != nil {
		t.Fatal(err)
	}
	p, err = h.store.Progress(t.Context(), "dddddddddddd")
	if err != nil {
		t.Fatal(err)
	}
	if p.State != StateRunning || p.IterationsDone != 1500 || p.DPS == nil || *p.DPS != 1000 {
		t.Fatalf("running: %+v", p)
	}

	done := browserResult("warrior-fury", 1042.5)
	done.Lane = simapi.LaneServer
	if err := h.store.Finish(t.Context(), "dddddddddddd", done); err != nil {
		t.Fatal(err)
	}
	p, err = h.store.Progress(t.Context(), "dddddddddddd")
	if err != nil {
		t.Fatal(err)
	}
	if p.State != StateDone || p.DPS == nil || *p.DPS != 1042.5 {
		t.Fatalf("done: %+v", p)
	}
	stored, err := h.store.Get(t.Context(), "dddddddddddd")
	if err != nil {
		t.Fatal(err)
	}
	if stored.Lane != simapi.LaneServer {
		t.Fatalf("finished result: lane %q", stored.Lane)
	}
}

func TestAQueuedRunThatCannotStartIsMarkedFailed(t *testing.T) {
	h := newHarness(t)
	if err := h.store.Queue(t.Context(), "eeeeeeeeeeee", h.owner,
		browserResult("warrior-fury", 0).Request); err != nil {
		t.Fatal(err)
	}
	if err := h.store.Fail(t.Context(), "eeeeeeeeeeee", "the job could not be started"); err != nil {
		t.Fatal(err)
	}
	p, err := h.store.Progress(t.Context(), "eeeeeeeeeeee")
	if err != nil {
		t.Fatal(err)
	}
	if p.State != StateError {
		t.Fatalf("state %q, want %q", p.State, StateError)
	}
}

func TestAFinishedResultCarryingAnEngineErrorIsStoredAsAnError(t *testing.T) {
	h := newHarness(t)
	if err := h.store.Queue(t.Context(), "ffffffffffff", h.owner,
		browserResult("warrior-fury", 0).Request); err != nil {
		t.Fatal(err)
	}
	bad := browserResult("warrior-fury", 0)
	bad.Error = "the engine has no model for that spec"
	if err := h.store.Finish(t.Context(), "ffffffffffff", bad); err != nil {
		t.Fatal(err)
	}
	p, err := h.store.Progress(t.Context(), "ffffffffffff")
	if err != nil {
		t.Fatal(err)
	}
	if p.State != StateError {
		t.Fatalf("state %q, want %q", p.State, StateError)
	}
}

func TestSavingTheSameIdTwiceKeepsTheFirst(t *testing.T) {
	h := newHarness(t)
	if err := h.store.Save(t.Context(), "gggggggggggg", &h.owner, "first",
		browserResult("warrior-fury", 900)); err != nil {
		t.Fatal(err)
	}
	if err := h.store.Save(t.Context(), "gggggggggggg", &h.owner, "second",
		browserResult("warrior-fury", 1900)); err != nil {
		t.Fatal(err)
	}
	got, err := h.store.Get(t.Context(), "gggggggggggg")
	if err != nil {
		t.Fatal(err)
	}
	if got.DPS.Mean != 900 {
		t.Fatalf("mean %v, want the first save's 900", got.DPS.Mean)
	}
}

func TestTheResultKeyIsLaidOutLikeAReports(t *testing.T) {
	if got := (Keys{SimID: "aaaaaaaaaaaa"}).Result(); got != "sims/aaaaaaaaaaaa/result.json" {
		t.Errorf("result key %q", got)
	}
}
