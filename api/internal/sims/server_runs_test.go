package sims

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// serverRun queues one server-lane run for the harness owner and leaves it
// queued; the caller walks it on.
func serverRun(h *harness, id, spec string) simapi.SimRequest {
	h.t.Helper()
	req := browserResult(spec, 0).Request
	if err := h.store.Queue(h.t.Context(), id, h.owner, req); err != nil {
		h.t.Fatal(err)
	}
	return req
}

// finishedServerRun queues and finishes one server-lane run.
func finishedServerRun(h *harness, id, spec string) {
	h.t.Helper()
	serverRun(h, id, spec)
	done := browserResult(spec, 1000)
	done.Lane = simapi.LaneServer
	if err := h.store.Finish(h.t.Context(), id, done); err != nil {
		h.t.Fatal(err)
	}
}

func failedServerRun(h *harness, id string) simapi.SimRequest {
	h.t.Helper()
	req := serverRun(h, id, "warrior-fury")
	if err := h.store.Fail(h.t.Context(), id, "the engine stopped"); err != nil {
		h.t.Fatal(err)
	}
	return req
}

func TestMineLaneServerListsOnlyServerRunsWithEstimates(t *testing.T) {
	h := newHarness(t)
	for _, id := range []string{"aaaaaaaaaaa1", "aaaaaaaaaaa2", "aaaaaaaaaaa3"} {
		finishedServerRun(h, id, "warrior-fury")
	}
	serverRun(h, "aaaaaaaaaaa4", "warrior-fury")
	if err := h.store.Advance(t.Context(), "aaaaaaaaaaa4", Tick{IterationsDone: 100, Mean: 900}); err != nil {
		t.Fatal(err)
	}
	browserID := saveBrowserResult(h, "warrior-fury", 1000, "browser")

	var page Page
	h.data(h.do(http.MethodGet, "/v1/sims?mine=1&lane=server", "", nil), &page)
	if page.Total != 4 || len(page.Rows) != 4 {
		t.Fatalf("page: %+v", page)
	}
	for _, r := range page.Rows {
		switch {
		case r.SimID == browserID:
			t.Fatalf("a browser save is in the server list: %+v", r)
		case r.State == StateRunning && r.EstimateSeconds == nil:
			t.Errorf("running row has no estimate: %+v", r)
		case r.State == StateDone && (r.EstimateSeconds != nil || r.Seen):
			t.Errorf("done row: %+v", r)
		}
	}
}

func TestMineRejectsAnUnknownLane(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodGet, "/v1/sims?mine=1&lane=moon", "", nil)
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", res.StatusCode)
	}
	if code := h.errorCode(res); code != "invalid" {
		t.Fatalf("code %q, want invalid", code)
	}
}

func TestSeenMarksTheOwnersDoneRunAndRefusesOthers(t *testing.T) {
	h := newHarness(t)
	finishedServerRun(h, "bbbbbbbbbbb1", "warrior-fury")

	if res := h.do(http.MethodPost, "/v1/sims/bbbbbbbbbbb1/seen", "", nil); res.StatusCode != http.StatusNoContent {
		t.Fatalf("status %d, want 204", res.StatusCode)
	}
	var seen bool
	if err := h.store.Pool.QueryRow(t.Context(),
		`select seen_at is not null from sims where id = $1`, "bbbbbbbbbbb1").Scan(&seen); err != nil || !seen {
		t.Fatalf("seen_at set = %v, err %v", seen, err)
	}

	h.actor.UserID = h.otherUser()
	if res := h.do(http.MethodPost, "/v1/sims/bbbbbbbbbbb1/seen", "", nil); res.StatusCode != http.StatusNotFound {
		t.Fatalf("other user: status %d, want 404", res.StatusCode)
	}
	h.anonymous()
	if res := h.do(http.MethodPost, "/v1/sims/bbbbbbbbbbb1/seen", "", nil); res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("anonymous: status %d, want 401", res.StatusCode)
	}
}

func TestRerunStartsTheFailedRunsStoredRequest(t *testing.T) {
	h := newHarness(t)
	h.entitlements.allowed = true
	old := failedServerRun(h, "ccccccccccc1")

	res := h.do(http.MethodPost, "/v1/sims/ccccccccccc1/rerun", "", nil)
	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("status %d, want 202", res.StatusCode)
	}
	var out struct {
		SimID string `json:"sim_id"`
	}
	h.data(res, &out)
	if len(out.SimID) != 12 || out.SimID == "ccccccccccc1" {
		t.Fatalf("sim id %q", out.SimID)
	}
	ran := h.jobs.Ran()
	if len(ran) != 1 || ran[0][0] != SimRunJobCommand || ran[0][1] != out.SimID {
		t.Fatalf("dispatched %v", ran)
	}
	got, _, err := h.store.Get(t.Context(), out.SimID)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got.Request, old) {
		t.Fatalf("stored request %+v, want %+v", got.Request, old)
	}
	var state string
	var reason *string
	var body []byte
	if err := h.store.Pool.QueryRow(t.Context(),
		`select state, failure_reason, result from sims where id = $1`, "ccccccccccc1").
		Scan(&state, &reason, &body); err != nil || state != StateError {
		t.Fatalf("old row state %q, err %v", state, err)
	}
	if reason == nil || *reason != "the engine stopped" {
		t.Fatalf("old row failure_reason %v, want the original", reason)
	}
	var stored simapi.SimResult
	if err := json.Unmarshal(body, &stored); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(stored.Request, old) {
		t.Fatalf("old row request %+v, want %+v", stored.Request, old)
	}
}

func TestRerunRefusesADoneRowAndAnotherUsersRow(t *testing.T) {
	h := newHarness(t)
	h.entitlements.allowed = true
	finishedServerRun(h, "ddddddddddd1", "warrior-fury")
	failedServerRun(h, "ddddddddddd2")

	if res := h.do(http.MethodPost, "/v1/sims/ddddddddddd1/rerun", "", nil); res.StatusCode != http.StatusNotFound {
		t.Fatalf("done row: status %d, want 404", res.StatusCode)
	}
	h.actor.UserID = h.otherUser()
	if res := h.do(http.MethodPost, "/v1/sims/ddddddddddd2/rerun", "", nil); res.StatusCode != http.StatusNotFound {
		t.Fatalf("other user: status %d, want 404", res.StatusCode)
	}
	if ran := h.jobs.Ran(); len(ran) != 0 {
		t.Fatalf("dispatched %v", ran)
	}
}

func TestRerunIsGatedByTheEntitlement(t *testing.T) {
	h := newHarness(t)
	failedServerRun(h, "eeeeeeeeeee1")
	res := h.do(http.MethodPost, "/v1/sims/eeeeeeeeeee1/rerun", "", nil)
	if res.StatusCode != http.StatusPaymentRequired {
		t.Fatalf("status %d, want 402", res.StatusCode)
	}
	if code := h.errorCode(res); code != "premium_required" {
		t.Fatalf("code %q", code)
	}
}
