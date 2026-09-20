package sims

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// saveBrowserResult posts one finished browser run and returns its id.
func saveBrowserResult(h *harness, spec string, mean float64, title string) string {
	h.t.Helper()
	res := browserResult(spec, mean)
	body, err := json.Marshal(struct {
		simapi.SimResult
		Title string `json:"title"`
	}{res, title})
	if err != nil {
		h.t.Fatal(err)
	}
	out := struct {
		SimID string `json:"sim_id"`
	}{}
	h.data(h.json(http.MethodPost, "/v1/sims", string(body)), &out)
	if out.SimID == "" {
		h.t.Fatal("no sim id came back")
	}
	return out.SimID
}

func TestABrowserResultIsSavedAndReadBack(t *testing.T) {
	h := newHarness(t)
	id := saveBrowserResult(h, "warrior-fury", 1042.5, "Tuesday")
	if len(id) != 12 {
		t.Fatalf("sim id %q is not twelve characters", id)
	}
	if strings.ToLower(id) != id {
		t.Fatalf("sim id %q is not lowercase", id)
	}

	h.anonymous()
	res := h.do(http.MethodGet, "/v1/sims/"+id, "", nil)
	if cc := res.Header.Get("Cache-Control"); !strings.Contains(cc, "public") {
		t.Errorf("Cache-Control %q: a saved sim never changes", cc)
	}
	var got simapi.SimResult
	h.data(res, &got)
	if got.SimID != id || got.DPS.Mean != 1042.5 || got.Lane != simapi.LaneBrowser {
		t.Fatalf("read back: %+v", got)
	}
	if got.Request.Character.Class != "warrior" {
		t.Errorf("the character did not survive: %+v", got.Request.Character)
	}
}

// TestASaveAcceptsAResultFromAnOlderEngineAndReadsBackStale pins
// HIGH-4: saving is not running, so a browser result naming an engine
// build behind the deployment's pin is stored, not refused - the
// member does not lose the sim they just ran to a stale WASM bundle.
// It reads back through SimResult.Stale, the contract's own answer.
func TestASaveAcceptsAResultFromAnOlderEngineAndReadsBackStale(t *testing.T) {
	h := newHarness(t)
	res := browserResult("warrior-fury", 1000)
	res.EngineVersion, res.Request.EngineVersion = "anoldbuild", "anoldbuild"
	body, err := json.Marshal(struct {
		simapi.SimResult
		Title string `json:"title"`
	}{res, "an old build"})
	if err != nil {
		t.Fatal(err)
	}
	out := struct {
		SimID string `json:"sim_id"`
	}{}
	res2 := h.json(http.MethodPost, "/v1/sims", string(body))
	if res2.StatusCode != http.StatusCreated {
		t.Fatalf("status %d, want 201: a result from an older engine build must be saved, not refused", res2.StatusCode)
	}
	h.data(res2, &out)
	if out.SimID == "" {
		t.Fatal("no sim id came back")
	}
	var got simapi.SimResult
	h.data(h.do(http.MethodGet, "/v1/sims/"+out.SimID, "", nil), &got)
	if got.EngineVersion != "anoldbuild" || !got.Stale(testEngine) {
		t.Errorf("engine_version %q did not read back stale against the pin %q", got.EngineVersion, testEngine)
	}
}

// TestASaveRefusesAnAbortedResult pins the other half of HIGH-4's
// ruling: ValidateSaved relaxes the engine-version check but still
// refuses a partial run. The browser only ever posts a finished
// result, so an aborted one reaching this route is malformed.
func TestASaveRefusesAnAbortedResult(t *testing.T) {
	h := newHarness(t)
	res := browserResult("warrior-fury", 1000)
	res.Aborted = true
	body, err := json.Marshal(struct {
		simapi.SimResult
		Title string `json:"title"`
	}{res, ""})
	if err != nil {
		t.Fatal(err)
	}
	got := h.json(http.MethodPost, "/v1/sims", string(body))
	if got.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", got.StatusCode)
	}
	if code := h.errorCode(got); code != "invalid" {
		t.Fatalf("code %q", code)
	}
}

func TestAnUnknownSimIs404(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodGet, "/v1/sims/zzzzzzzzzzzz", "", nil)
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("status %d", res.StatusCode)
	}
	if code := h.errorCode(res); code != "not_found" {
		t.Fatalf("code %q", code)
	}
}

func TestOnlyABrowserResultIsSavedThisWay(t *testing.T) {
	h := newHarness(t)
	res := browserResult("warrior-fury", 1000)
	res.Lane = simapi.LaneServer
	body, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	got := h.json(http.MethodPost, "/v1/sims", string(body))
	if got.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", got.StatusCode)
	}
	if code := h.errorCode(got); code != "invalid" {
		t.Fatalf("code %q", code)
	}
}

func TestARequestTheEnvelopeRejectsIsRefused(t *testing.T) {
	h := newHarness(t)
	for _, c := range []struct {
		name string
		mut  func(*simapi.SimResult)
	}{
		{"no spec", func(r *simapi.SimResult) { r.Request.Spec = "" }},
		{"no race", func(r *simapi.SimResult) { r.Request.Character.Race = "" }},
		{"a level the engine cannot sim", func(r *simapi.SimResult) { r.Request.Character.Level = simapi.SimLevel - 1 }},
		{"an iteration count outside the set", func(r *simapi.SimResult) { r.Request.Iterations = 7 }},
		{"a fight of five seconds", func(r *simapi.SimResult) { r.Request.Encounter.DurationSec = 5 }},
		{"ninety-nine targets", func(r *simapi.SimResult) { r.Request.Encounter.Targets = 99 }},
	} {
		t.Run(c.name, func(t *testing.T) {
			res := browserResult("warrior-fury", 1000)
			c.mut(&res)
			body, err := json.Marshal(res)
			if err != nil {
				t.Fatal(err)
			}
			got := h.json(http.MethodPost, "/v1/sims", string(body))
			if got.StatusCode != http.StatusBadRequest {
				t.Fatalf("status %d, want 400", got.StatusCode)
			}
		})
	}
}

func TestGarbageIsRefused(t *testing.T) {
	h := newHarness(t)
	got := h.json(http.MethodPost, "/v1/sims", `{"lane":`)
	if got.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", got.StatusCode)
	}
}

func TestMyOwnSimsNeedASessionAndMineEqualsOne(t *testing.T) {
	h := newHarness(t)
	saveBrowserResult(h, "warrior-fury", 1000, "one")
	saveBrowserResult(h, "mage-frost", 1100, "two")

	var page Page
	h.data(h.do(http.MethodGet, "/v1/sims?mine=1", "", nil), &page)
	if page.Total != 2 || len(page.Rows) != 2 || page.PerPage != PerPage {
		t.Fatalf("page: %+v", page)
	}

	// Without mine=1 the route refuses rather than guessing.
	if got := h.do(http.MethodGet, "/v1/sims", "", nil); got.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", got.StatusCode)
	}
	// A bad page number is refused rather than silently corrected.
	if got := h.do(http.MethodGet, "/v1/sims?mine=1&page=0", "", nil); got.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", got.StatusCode)
	}

	h.anonymous()
	if got := h.do(http.MethodGet, "/v1/sims?mine=1", "", nil); got.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status %d, want 401", got.StatusCode)
	}
}

func TestMyOwnSimsFilterByKind(t *testing.T) {
	h := newHarness(t)
	saveBrowserResult(h, "warrior-fury", 1204.4, "a plain run")

	gear := browserResult("mage-frost", 1100)
	// Cap and Candidates are here only so ValidateSaved accepts the
	// request on its way through POST /v1/sims; the headline this
	// test checks is read from Combos, not from either field.
	gear.Request.Bulk = &simapi.BulkSpec{
		Mode: simapi.KindGear, Precision: simapi.PrecisionNormal,
		Cap: simapi.Caps[simapi.LaneBrowser],
		Candidates: []simapi.Candidate{
			{Slot: "main_hand", ItemID: 19019, Origin: "bag"},
		},
	}
	gear.Combos = []simapi.Combo{{
		Substitutions: []simapi.Substitution{
			{Kind: "item", ItemID: 19019, Name: "Thunderfury", Origin: "search"},
		},
		Delta: simapi.Estimate{Mean: 41.2},
	}}
	body, err := json.Marshal(struct {
		simapi.SimResult
		Title string `json:"title"`
	}{gear, "top gear"})
	if err != nil {
		t.Fatal(err)
	}
	h.data(h.json(http.MethodPost, "/v1/sims", string(body)), nil)

	var all Page
	h.data(h.do(http.MethodGet, "/v1/sims?mine=1", "", nil), &all)
	if all.Total != 2 {
		t.Fatalf("unfiltered total %d, want 2", all.Total)
	}

	var only Page
	h.data(h.do(http.MethodGet, "/v1/sims?mine=1&kind=gear", "", nil), &only)
	if only.Total != 1 || len(only.Rows) != 1 {
		t.Fatalf("filtered: %+v", only)
	}
	if only.Rows[0].Kind != simapi.KindGear ||
		only.Rows[0].Headline != "+41 DPS from Thunderfury" {
		t.Fatalf("row: %+v", only.Rows[0])
	}

	// An empty kind is "every kind", not "a kind called empty".
	var empty Page
	h.data(h.do(http.MethodGet, "/v1/sims?mine=1&kind=", "", nil), &empty)
	if empty.Total != 2 {
		t.Fatalf("kind= total %d, want 2", empty.Total)
	}
}

func TestAnUnknownKindIsRefusedRatherThanAnsweredEmpty(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodGet, "/v1/sims?mine=1&kind=topgear", "", nil)
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", res.StatusCode)
	}
	if code := h.errorCode(res); code != "invalid" {
		t.Fatalf("code %q, want invalid", code)
	}
}

func TestAnAnonymousSaveIsStillShareable(t *testing.T) {
	h := newHarness(t)
	h.anonymous()
	id := saveBrowserResult(h, "warrior-fury", 1000, "")
	var got simapi.SimResult
	h.data(h.do(http.MethodGet, "/v1/sims/"+id, "", nil), &got)
	if got.DPS.Mean != 1000 {
		t.Fatalf("read back: %+v", got)
	}
}

func TestALongTitleIsTrimmedRatherThanRefused(t *testing.T) {
	long := strings.Repeat("a", maxTitle+40)
	if got := trimTitle(long); len([]rune(got)) != maxTitle {
		t.Fatalf("trimmed to %d runes, want %d", len([]rune(got)), maxTitle)
	}
	if got := trimTitle("Tuesday"); got != "Tuesday" {
		t.Fatalf("a short title was changed: %q", got)
	}
	// Cut by rune, not by byte: a byte cut through a multi-byte rune
	// would store invalid UTF-8, which every reader down the line
	// would have to cope with.
	wide := strings.Repeat("辿", maxTitle+10)
	got := trimTitle(wide)
	if len([]rune(got)) != maxTitle {
		t.Fatalf("trimmed to %d runes, want %d", len([]rune(got)), maxTitle)
	}
	if !utf8.ValidString(got) {
		t.Fatalf("the trim produced invalid UTF-8: %q", got)
	}
	// An emoji is several runes' worth of bytes and still counts as
	// its runes, not its bytes.
	if got := trimTitle("Tuesday 🐻"); got != "Tuesday 🐻" {
		t.Fatalf("a short title with an emoji was changed: %q", got)
	}
}

// A browser save writes its row done, so its progress is answered
// straight away rather than left to a poller.
func TestASavedSimsProgressIsAnsweredDone(t *testing.T) {
	h := newHarness(t)
	id := saveBrowserResult(h, "warrior-fury", 1042.5, "Tuesday")

	res := h.do(http.MethodGet, "/v1/sims/"+id+"/progress", "", nil)
	if cc := res.Header.Get("Cache-Control"); cc != "no-store" {
		t.Errorf("Cache-Control %q, want no-store: a run in flight is never cached", cc)
	}
	var p Progress
	h.data(res, &p)
	if p.State != StateDone || p.DPS == nil || *p.DPS != 1042.5 {
		t.Fatalf("progress: %+v", p)
	}
}

func TestAnUnknownSimsProgressIs404(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodGet, "/v1/sims/zzzzzzzzzzzz/progress", "", nil)
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("status %d", res.StatusCode)
	}
}

// The sim-input route Task 7 fills in is mounted, even though this
// task leaves it as a stub: a request reaching it answers a plain 404
// rather than a routing failure. GET /v1/specs left this map when
// Task 6 gave it a real handler; its own tests are in specs_test.go.
func TestTheStubbedRoutesAreMountedAndAnswer404(t *testing.T) {
	h := newHarness(t)
	for name, res := range map[string]*http.Response{
		"sim-input": h.do(http.MethodGet, "/v1/characters/us/normal/baelgrim/sim-input", "", nil),
	} {
		res.Body.Close()
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("%s = %d, want 404", name, res.StatusCode)
		}
	}
}

// POST /v1/sims/run is not mounted at all when this deployment has
// neither a job runner nor an accounts store behind it. newHarness
// wires up both (Task 5 needs a harness that can exercise the run
// route), so this builds its own bare Service to exercise Mount's
// guard directly. The path still matches GET /v1/sims/{id}'s pattern
// (id "run"), so the mux answers 405 rather than 404 - that pattern
// is registered, just not for POST - which is still proof the run
// route itself was never mounted.
func TestTheRunRouteIsNotMountedWithoutJobsAndAccounts(t *testing.T) {
	svc := &Service{Store: &Store{}, EngineVersion: testEngine}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodPost, "/v1/sims/run", strings.NewReader(`{}`))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status %d, want 405", rec.Code)
	}
}

// TestTheRunRouteNeedsAPlannerToo mirrors
// TestTheRunRouteIsNotMountedWithoutJobsAndAccounts above: on a bare
// mux with no catch-all "/", a POST to /v1/sims/run still matches GET
// /v1/sims/{id}'s path pattern (id "run"), so the mux answers 405
// rather than 404 - that pattern is registered, just not for POST -
// which is still proof the run route was never mounted without a
// planner. (The real router in api/internal/server does register a
// catch-all "/", so there the same gap answers a plain 404 -
// TestTheSaveMineAndRunRoutesAreMounted covers that.)
func TestTheRunRouteNeedsAPlannerToo(t *testing.T) {
	mux := http.NewServeMux()
	Mount(mux, &Service{Jobs: &jobs.Fake{}, Accounts: &fakePremium{}})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/v1/sims/run", strings.NewReader("{}")))
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("status %d, want 405: a lane that cannot size a bulk request does not offer one", w.Code)
	}
}

// Every handler that reads or writes through the store answers the
// envelope's 500 rather than crashing or hanging when the database is
// gone.
func TestEveryHandlerAnswers500WhenTheDatabaseIsGone(t *testing.T) {
	h := newHarness(t)
	id := saveBrowserResult(h, "warrior-fury", 1000, "Tuesday")
	h.store.Pool.Close()

	body, err := json.Marshal(struct {
		simapi.SimResult
		Title string `json:"title"`
	}{browserResult("warrior-fury", 1000), "x"})
	if err != nil {
		t.Fatal(err)
	}
	for name, res := range map[string]*http.Response{
		"save":     h.json(http.MethodPost, "/v1/sims", string(body)),
		"get":      h.do(http.MethodGet, "/v1/sims/"+id, "", nil),
		"mine":     h.do(http.MethodGet, "/v1/sims?mine=1", "", nil),
		"progress": h.do(http.MethodGet, "/v1/sims/"+id+"/progress", "", nil),
	} {
		res.Body.Close()
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("%s = %d, want 500", name, res.StatusCode)
		}
	}
}
