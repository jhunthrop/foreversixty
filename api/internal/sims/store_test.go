package sims

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

func TestASavedBrowserResultComesBackWhole(t *testing.T) {
	h := newHarness(t)
	res := browserResult("warrior-fury", 1042.5)
	if err := h.store.Save(t.Context(), "aaaaaaaaaaaa", &h.owner, "Tuesday", res); err != nil {
		t.Fatal(err)
	}
	got, title, err := h.store.Get(t.Context(), "aaaaaaaaaaaa")
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
	// Defect fix: the name a member gave this sim ("Tuesday", saved above) used
	// to live only in the `sims.title` column -- Get() never selected it, and
	// the saved page, its <title> and its og tags all had nothing to read.
	if title != "Tuesday" {
		t.Errorf("title = %q, want %q", title, "Tuesday")
	}
}

// A sim saved with no title (the common case: naming a run is optional) reads
// back as "", not a Postgres NULL a caller has to special-case.
func TestASavedResultWithNoTitleReadsBackEmpty(t *testing.T) {
	h := newHarness(t)
	res := browserResult("warrior-fury", 1042.5)
	if err := h.store.Save(t.Context(), "notitleaaaaa", &h.owner, "", res); err != nil {
		t.Fatal(err)
	}
	_, title, err := h.store.Get(t.Context(), "notitleaaaaa")
	if err != nil {
		t.Fatal(err)
	}
	if title != "" {
		t.Errorf("title = %q, want empty", title)
	}
}

func TestAnUnknownSimIsNotFound(t *testing.T) {
	h := newHarness(t)
	if _, _, err := h.store.Get(t.Context(), "zzzzzzzzzzzz"); err != ErrNotFound {
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
	page, err := h.store.Mine(t.Context(), h.owner, 1, "", "")
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
	queued, _, err := h.store.Get(t.Context(), "dddddddddddd")
	if err != nil {
		t.Fatal(err)
	}
	if queued.Request.Spec != "warrior-fury" || queued.Request.Character.Class != "warrior" {
		t.Fatalf("queued request: %+v", queued.Request)
	}

	if err := h.store.Advance(t.Context(), "dddddddddddd", Tick{
		IterationsDone: 1500, Mean: 1000,
	}); err != nil {
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
	stored, _, err := h.store.Get(t.Context(), "dddddddddddd")
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

func TestAdvanceCannotReviveAFailedRun(t *testing.T) {
	h := newHarness(t)
	if err := h.store.Queue(t.Context(), "hhhhhhhhhhhh", h.owner,
		browserResult("warrior-fury", 0).Request); err != nil {
		t.Fatal(err)
	}
	if err := h.store.Fail(t.Context(), "hhhhhhhhhhhh", "the job could not be started"); err != nil {
		t.Fatal(err)
	}
	// A late or duplicate progress tick arriving after the failure must
	// not walk the row back to running: error is terminal.
	if err := h.store.Advance(t.Context(), "hhhhhhhhhhhh", Tick{
		IterationsDone: 1500, Mean: 1000,
	}); err != nil {
		t.Fatal(err)
	}
	p, err := h.store.Progress(t.Context(), "hhhhhhhhhhhh")
	if err != nil {
		t.Fatal(err)
	}
	if p.State != StateError {
		t.Fatalf("state %q, want %q; Advance revived a failed run", p.State, StateError)
	}
	if p.DPS != nil || p.IterationsDone != 0 {
		t.Fatalf("progress: %+v; a terminal row must not pick up Advance's figures", p)
	}
}

func TestAdvanceCannotReopenAFinishedRun(t *testing.T) {
	h := newHarness(t)
	if err := h.store.Queue(t.Context(), "iiiiiiiiiiii", h.owner,
		browserResult("warrior-fury", 0).Request); err != nil {
		t.Fatal(err)
	}
	done := browserResult("warrior-fury", 1042.5)
	done.Lane = simapi.LaneServer
	if err := h.store.Finish(t.Context(), "iiiiiiiiiiii", done); err != nil {
		t.Fatal(err)
	}
	// A late or duplicate progress tick arriving after the result must
	// not overwrite it: done is terminal.
	if err := h.store.Advance(t.Context(), "iiiiiiiiiiii", Tick{
		IterationsDone: 1500, Mean: 1000,
	}); err != nil {
		t.Fatal(err)
	}
	p, err := h.store.Progress(t.Context(), "iiiiiiiiiiii")
	if err != nil {
		t.Fatal(err)
	}
	if p.State != StateDone {
		t.Fatalf("state %q, want %q; Advance reopened a finished run", p.State, StateDone)
	}
	if p.DPS == nil || *p.DPS != 1042.5 {
		t.Fatalf("progress: %+v; Advance overwrote the finished result", p)
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
	// And it composes no headline. Headline would read the failed
	// result's zero fields as an answer — "0 DPS" — which is a sentence
	// claiming the run finished and found nothing. The row carries its
	// state instead, which is what the list has to show.
	page, err := h.store.Mine(t.Context(), h.owner, 1, "", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Rows) != 1 {
		t.Fatalf("rows = %+v, want the one failed run", page.Rows)
	}
	if got := page.Rows[0]; got.Headline != "" || got.State != StateError {
		t.Fatalf("failed row: %+v; want an empty headline and state %q", got, StateError)
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
	got, _, err := h.store.Get(t.Context(), "gggggggggggg")
	if err != nil {
		t.Fatal(err)
	}
	if got.DPS.Mean != 900 {
		t.Fatalf("mean %v, want the first save's 900", got.DPS.Mean)
	}
}

func TestEveryRowRecordsWhichToolProducedIt(t *testing.T) {
	h := newHarness(t)
	// A browser save of a plain run.
	if err := h.store.Save(t.Context(), "aaaaaaaaaaaa", &h.owner, "",
		browserResult("warrior-fury", 1000)); err != nil {
		t.Fatal(err)
	}
	// A queued server run of a Top Gear request.
	gear := browserResult("warrior-fury", 0).Request
	gear.Bulk = &simapi.BulkSpec{
		Mode:      simapi.KindGear,
		Precision: simapi.PrecisionNormal,
		Cap:       simapi.Caps[simapi.LaneServer],
		Candidates: []simapi.Candidate{
			{Slot: "main_hand", ItemID: 19019, Origin: "bag"},
		},
	}
	if err := h.store.Queue(t.Context(), "bbbbbbbbbbbb", h.owner, gear); err != nil {
		t.Fatal(err)
	}
	// A queued weights run.
	weights := browserResult("warrior-fury", 0).Request
	weights.Weights = &simapi.WeightsSpec{
		Stats: []string{"strength", "crit"}, Reference: "crit",
	}
	if err := h.store.Queue(t.Context(), "cccccccccccc", h.owner, weights); err != nil {
		t.Fatal(err)
	}

	for _, c := range []struct{ id, want string }{
		{"aaaaaaaaaaaa", simapi.KindRun},
		{"bbbbbbbbbbbb", simapi.KindGear},
		{"cccccccccccc", simapi.KindWeights},
	} {
		var got string
		if err := h.store.Pool.QueryRow(t.Context(),
			`select kind from sims where id = $1`, c.id).Scan(&got); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: kind %q, want %q", c.id, got, c.want)
		}
	}
}

func TestMyHistoryCarriesTheKindAndHeadlineAndFiltersByKind(t *testing.T) {
	h := newHarness(t)
	plain := browserResult("warrior-fury", 1204.4)
	if err := h.store.Save(t.Context(), "aaaaaaaaaaaa", &h.owner, "", plain); err != nil {
		t.Fatal(err)
	}
	gear := browserResult("warrior-fury", 1204.4)
	gear.Request.Bulk = &simapi.BulkSpec{Mode: simapi.KindGear, Precision: simapi.PrecisionNormal}
	gear.Combos = []simapi.Combo{{
		Substitutions: []simapi.Substitution{
			{Kind: "item", ItemID: 17182, Name: "Vis'kag the Bloodletter", Origin: "bag"},
		},
		Delta: simapi.Estimate{Mean: 41.2},
	}}
	if err := h.store.Save(t.Context(), "bbbbbbbbbbbb", &h.owner, "", gear); err != nil {
		t.Fatal(err)
	}

	all, err := h.store.Mine(t.Context(), h.owner, 1, "", "")
	if err != nil {
		t.Fatal(err)
	}
	if all.Total != 2 {
		t.Fatalf("total %d, want both kinds", all.Total)
	}
	byID := map[string]Row{}
	for _, r := range all.Rows {
		byID[r.SimID] = r
	}
	if got := byID["aaaaaaaaaaaa"]; got.Kind != simapi.KindRun || got.Headline != "1,204 DPS" ||
		got.State != StateDone {
		t.Errorf("plain row: %+v", got)
	}
	if got := byID["bbbbbbbbbbbb"]; got.Kind != simapi.KindGear ||
		got.Headline != "+41 DPS from Vis'kag the Bloodletter" {
		t.Errorf("gear row: %+v", got)
	}

	only, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindGear, "")
	if err != nil {
		t.Fatal(err)
	}
	if only.Total != 1 || len(only.Rows) != 1 || only.Rows[0].SimID != "bbbbbbbbbbbb" {
		t.Fatalf("filtered: total %d rows %+v", only.Total, only.Rows)
	}

	none, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindDrops, "")
	if err != nil {
		t.Fatal(err)
	}
	if none.Total != 0 || len(none.Rows) != 0 {
		t.Fatalf("a kind with no rows: total %d rows %+v", none.Total, none.Rows)
	}
}

func TestAServerRunsHeadlineIsWrittenWhenItFinishes(t *testing.T) {
	h := newHarness(t)
	req := browserResult("warrior-fury", 0).Request
	req.Weights = &simapi.WeightsSpec{
		// crit, not melee_crit: contract 10.8 carries one hit and one
		// crit, not the lane-split melee/spell pairs.
		Stats: []string{"crit", "agility"}, Reference: "crit",
	}
	if err := h.store.Queue(t.Context(), "cccccccccccc", h.owner, req); err != nil {
		t.Fatal(err)
	}
	// A queued run has nothing to say yet.
	queued, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindWeights, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(queued.Rows) != 1 || queued.Rows[0].Headline != "" ||
		queued.Rows[0].State != StateQueued {
		t.Fatalf("queued row: %+v", queued.Rows)
	}

	done := browserResult("warrior-fury", 1000)
	done.Lane, done.Request = simapi.LaneServer, req
	done.Weights = []simapi.StatWeight{
		{Stat: "crit", Weight: 1}, {Stat: "agility", Weight: 0.874},
	}
	if err := h.store.Finish(t.Context(), "cccccccccccc", done); err != nil {
		t.Fatal(err)
	}
	finished, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindWeights, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(finished.Rows) != 1 || finished.Rows[0].Headline != "Crit 1.00 · Agility 0.87" ||
		finished.Rows[0].State != StateDone {
		t.Fatalf("finished row: %+v", finished.Rows)
	}
}

// TestAnAbsurdPageIsClampedRatherThanFailing pins the overflow:
// (page-1)*PerPage past the int range wraps negative, and Postgres
// refuses a negative OFFSET — a 500 where an empty page is the answer.
func TestAnAbsurdPageIsClampedRatherThanFailing(t *testing.T) {
	h := newHarness(t)
	if err := h.store.Save(t.Context(), "aaaaaaaaaaaa", &h.owner, "",
		browserResult("warrior-fury", 900)); err != nil {
		t.Fatal(err)
	}
	page, err := h.store.Mine(t.Context(), h.owner, math.MaxInt, "", "")
	if err != nil {
		t.Fatalf("an absurd page must be empty, not an error: %v", err)
	}
	if page.Page != MaxPage {
		t.Errorf("page = %d, want it clamped to %d", page.Page, MaxPage)
	}
	if len(page.Rows) != 0 {
		t.Errorf("rows = %+v, want none that far in", page.Rows)
	}
	if page.Total != 1 {
		t.Errorf("total = %d; the count is of the whole list, not the page", page.Total)
	}
}

// buildResult is a plausible finished sim whose source names a build,
// for ForBuild's tests.
func buildResult(buildID string, mean float64) simapi.SimResult {
	r := browserResult("warrior-arms", mean)
	r.Request.Source = simapi.CharacterSource{Kind: simapi.SourceBuild, Ref: buildID}
	return r
}

func TestForBuildFindsTheNewestDoneSimForThatBuild(t *testing.T) {
	h := newHarness(t)
	if err := h.store.Save(t.Context(), "aaaaaaaaaaaa", &h.owner, "",
		buildResult("znorjmts", 900)); err != nil {
		t.Fatal(err)
	}
	if err := h.store.Save(t.Context(), "bbbbbbbbbbbb", &h.owner, "",
		buildResult("znorjmts", 1500)); err != nil {
		t.Fatal(err)
	}
	got, ok, err := h.store.ForBuild(t.Context(), "znorjmts")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("ok = false, want a sim for this build")
	}
	if got.SimID != "bbbbbbbbbbbb" || got.DPS.Mean != 1500 {
		t.Fatalf("got = %+v, want the newest save (1500)", got)
	}
}

func TestForBuildIgnoresAnotherBuildsRefAndANonBuildSource(t *testing.T) {
	h := newHarness(t)
	if err := h.store.Save(t.Context(), "aaaaaaaaaaaa", &h.owner, "",
		buildResult("otherbuild", 900)); err != nil {
		t.Fatal(err)
	}
	if err := h.store.Save(t.Context(), "bbbbbbbbbbbb", &h.owner, "",
		browserResult("warrior-fury", 1000)); err != nil {
		t.Fatal(err)
	}
	_, ok, err := h.store.ForBuild(t.Context(), "znorjmts")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("ForBuild found a sim for a build it was never run against")
	}
}

func TestForBuildIgnoresASimThatIsNotDoneYet(t *testing.T) {
	h := newHarness(t)
	req := buildResult("znorjmts", 0).Request
	if err := h.store.Queue(t.Context(), "cccccccccccc", h.owner, req); err != nil {
		t.Fatal(err)
	}
	_, ok, err := h.store.ForBuild(t.Context(), "znorjmts")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("ForBuild found a queued sim, which has not finished")
	}
}

func TestTheResultKeyIsLaidOutLikeAReports(t *testing.T) {
	if got := (Keys{SimID: "aaaaaaaaaaaa"}).Result(); got != "sims/aaaaaaaaaaaa/result.json" {
		t.Errorf("result key %q", got)
	}
}

// serverRun queues a server run for the owner and returns its id.
func (h *harness) serverRun(id, spec string) string {
	h.t.Helper()
	if err := h.store.Queue(h.t.Context(), id, h.owner, browserResult(spec, 0).Request); err != nil {
		h.t.Fatal(err)
	}
	return id
}

// finishedRun queues and finishes a server run whose engine took durationMS.
func (h *harness) finishedRun(id, spec string, durationMS int64) {
	h.t.Helper()
	h.serverRun(id, spec)
	res := browserResult(spec, 1000)
	res.Lane, res.DurationMS = simapi.LaneServer, durationMS
	if err := h.store.Finish(h.t.Context(), id, res); err != nil {
		h.t.Fatal(err)
	}
}

// setColumn writes one sims column directly, for the timestamps a test must place.
func (h *harness) setColumn(id, column string, value any) {
	h.t.Helper()
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`update sims set `+column+` = $2 where id = $1`, id, value); err != nil {
		h.t.Fatal(err)
	}
}

func (h *harness) otherUser() int64 {
	h.t.Helper()
	id, err := (&auth.Store{Pool: h.store.Pool}).UpsertEmailUser(h.t.Context(), "other@example.com")
	if err != nil {
		h.t.Fatal(err)
	}
	return id.ID
}

func (h *harness) finishedAt(id string) *time.Time {
	h.t.Helper()
	var at *time.Time
	if err := h.store.Pool.QueryRow(h.t.Context(),
		`select finished_at from sims where id = $1`, id).Scan(&at); err != nil {
		h.t.Fatal(err)
	}
	return at
}

func TestFinishSetsFinishedAt(t *testing.T) {
	h := newHarness(t)
	h.serverRun("fin000000001", "warrior-fury")
	if h.finishedAt("fin000000001") != nil {
		t.Fatal("a queued run has no finish time")
	}
	h.finishedRun("fin000000002", "warrior-fury", 1000)
	at := h.finishedAt("fin000000002")
	if at == nil || time.Since(*at) > time.Minute {
		t.Fatalf("finished_at = %v, want about now", at)
	}
}

func TestFailStoresThePlainReasonAndFinishedAt(t *testing.T) {
	h := newHarness(t)
	h.serverRun("fail00000001", "warrior-fury")
	if err := h.store.Fail(t.Context(), "fail00000001", "the engine was busy"); err != nil {
		t.Fatal(err)
	}
	page, err := h.store.Mine(t.Context(), h.owner, 1, "", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Rows) != 1 || page.Rows[0].State != StateError ||
		page.Rows[0].FailureReason != "the engine was busy" || page.Rows[0].FinishedAt == nil {
		t.Fatalf("failed row: %+v", page.Rows)
	}
}

func TestMineFiltersByLaneAndCarriesLifecycleFields(t *testing.T) {
	h := newHarness(t)
	if err := h.store.Save(t.Context(), "lane00000001", &h.owner, "",
		browserResult("warrior-fury", 900)); err != nil {
		t.Fatal(err)
	}
	h.finishedRun("lane00000002", "warrior-fury", 5000)
	server, err := h.store.Mine(t.Context(), h.owner, 1, "", simapi.LaneServer)
	if err != nil {
		t.Fatal(err)
	}
	if server.Total != 1 || len(server.Rows) != 1 {
		t.Fatalf("server lane: total %d, rows %+v", server.Total, server.Rows)
	}
	row := server.Rows[0]
	if row.SimID != "lane00000002" || row.Lane != simapi.LaneServer ||
		row.FinishedAt == nil || row.FailureReason != "" || row.Seen {
		t.Fatalf("server row: %+v", row)
	}
	browser, err := h.store.Mine(t.Context(), h.owner, 1, "", simapi.LaneBrowser)
	if err != nil {
		t.Fatal(err)
	}
	if browser.Total != 1 || browser.Rows[0].SimID != "lane00000001" {
		t.Fatalf("browser lane: %+v", browser)
	}
	all, err := h.store.Mine(t.Context(), h.owner, 1, "", "")
	if err != nil {
		t.Fatal(err)
	}
	if all.Total != 2 {
		t.Fatalf("every lane: total %d, want 2", all.Total)
	}
}

func TestMarkSeenOnlyForTheOwnersDoneRun(t *testing.T) {
	h := newHarness(t)
	h.finishedRun("seen00000001", "warrior-fury", 1000)
	h.serverRun("seen00000002", "warrior-fury")
	if ok, err := h.store.MarkSeen(t.Context(), "seen00000001", h.otherUser()); err != nil || ok {
		t.Fatalf("another user: ok=%v err=%v, want false", ok, err)
	}
	if ok, err := h.store.MarkSeen(t.Context(), "seen00000002", h.owner); err != nil || ok {
		t.Fatalf("a queued run: ok=%v err=%v, want false", ok, err)
	}
	if ok, err := h.store.MarkSeen(t.Context(), "seen00000001", h.owner); err != nil || !ok {
		t.Fatalf("the owner's done run: ok=%v err=%v, want true", ok, err)
	}
	page, err := h.store.Mine(t.Context(), h.owner, 1, "", simapi.LaneServer)
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, r := range page.Rows {
		seen[r.SimID] = r.Seen
	}
	if !seen["seen00000001"] || seen["seen00000002"] {
		t.Fatalf("seen flags: %v", seen)
	}
}

func TestFailedRequestReturnsTheOwnersErrorRowOnly(t *testing.T) {
	h := newHarness(t)
	h.serverRun("req000000001", "warrior-fury")
	if err := h.store.Fail(t.Context(), "req000000001", "boom"); err != nil {
		t.Fatal(err)
	}
	h.finishedRun("req000000002", "warrior-fury", 1000)
	got, err := h.store.FailedRequest(t.Context(), "req000000001", h.owner)
	if err != nil {
		t.Fatal(err)
	}
	if got.Spec != "warrior-fury" || got.Character.Class != "warrior" {
		t.Fatalf("request: %+v", got)
	}
	other := h.otherUser()
	cases := map[string]struct {
		id   string
		user int64
	}{
		"a done row":     {"req000000002", h.owner},
		"another user's": {"req000000001", other},
		"an unknown id":  {"zzzzzzzzzzzz", h.owner},
	}
	for name, c := range cases {
		if _, err := h.store.FailedRequest(t.Context(), c.id, c.user); !errors.Is(err, ErrNotFound) {
			t.Errorf("%s: err = %v, want ErrNotFound", name, err)
		}
	}
}

func TestFinishedUnseenNewestFirstAndCapped(t *testing.T) {
	h := newHarness(t)
	base := time.Now().Add(-time.Hour)
	ids := []string{"unseen000001", "unseen000002", "unseen000003", "unseen000004", "unseen000005"}
	for i, id := range ids {
		h.finishedRun(id, "warrior-fury", 1000)
		h.setColumn(id, "finished_at", base.Add(time.Duration(i)*time.Minute))
	}
	h.setColumn("unseen000005", "seen_at", time.Now())
	got, err := h.store.FinishedUnseen(t.Context(), h.owner, 3)
	if err != nil {
		t.Fatal(err)
	}
	gotIDs := make([]string, len(got))
	for i, r := range got {
		gotIDs[i] = r.SimID
	}
	want := []string{"unseen000004", "unseen000003", "unseen000002"}
	if !slices.Equal(gotIDs, want) {
		t.Fatalf("ids = %v, want %v", gotIDs, want)
	}
	if got[0].Spec != "warrior-fury" || got[0].DPS != 1000 || got[0].FinishedAt.IsZero() {
		t.Errorf("fields: %+v", got[0])
	}
}

func TestRunsThisMonthCountsServerRunsOnly(t *testing.T) {
	h := newHarness(t)
	if err := h.store.Save(t.Context(), "month0000001", &h.owner, "",
		browserResult("warrior-fury", 900)); err != nil {
		t.Fatal(err)
	}
	h.serverRun("month0000002", "warrior-fury")
	h.finishedRun("month0000003", "warrior-fury", 1000)
	h.finishedRun("month0000004", "warrior-fury", 1000)
	// A day before this month began: AddDate(0, -1, 0) on the 31st normalizes back into this month.
	now := time.Now().UTC()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	h.setColumn("month0000004", "created_at", monthStart.AddDate(0, 0, -1))
	n, err := h.store.RunsThisMonth(t.Context(), h.owner)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("runs this month = %d, want 2", n)
	}
}

func TestEstimateSecondsIsTheMedianOfTheLastTenDoneRuns(t *testing.T) {
	h := newHarness(t)
	h.finishedRun("est000000001", "warrior-fury", 60000)
	h.finishedRun("est000000002", "warrior-fury", 120000)
	got, err := h.store.EstimateSeconds(t.Context(), "warrior-fury")
	if err != nil || got != nil {
		t.Fatalf("two runs: %v, %v, want nil", got, err)
	}
	h.finishedRun("est000000003", "warrior-fury", 600000)
	if got, err = h.store.EstimateSeconds(t.Context(), "warrior-fury"); err != nil ||
		got == nil || *got != 120 {
		t.Fatalf("three runs: %v, %v, want 120", got, err)
	}
	// Ten newer runs of 10s push the three old ones out of the window.
	for i := range 10 {
		id := fmt.Sprintf("est0000001%02d", i)
		h.finishedRun(id, "warrior-fury", 10000)
		h.setColumn(id, "finished_at", time.Now().Add(time.Hour))
	}
	if got, err = h.store.EstimateSeconds(t.Context(), "warrior-fury"); err != nil ||
		got == nil || *got != 10 {
		t.Fatalf("eleven-plus runs: %v, %v, want 10", got, err)
	}
	if got, err = h.store.EstimateSeconds(t.Context(), "mage-fire"); err != nil || got != nil {
		t.Fatalf("another spec: %v, %v, want nil", got, err)
	}
	// A done row whose result lacks the duration is not a sample: the
	// median is unchanged.
	h.setColumn("est000000100", "result", []byte(`{}`))
	if got, err = h.store.EstimateSeconds(t.Context(), "warrior-fury"); err != nil ||
		got == nil || *got != 10 {
		t.Fatalf("a result without duration_ms: %v, %v, want 10", got, err)
	}
}

func TestEstimateSecondsDoesNotCountARunWithoutADuration(t *testing.T) {
	h := newHarness(t)
	h.finishedRun("estblank0001", "warrior-fury", 60000)
	h.finishedRun("estblank0002", "warrior-fury", 120000)
	h.finishedRun("estblank0003", "warrior-fury", 0)
	h.setColumn("estblank0003", "result", []byte(`{}`))
	if got, err := h.store.EstimateSeconds(t.Context(), "warrior-fury"); err != nil || got != nil {
		t.Fatalf("two real runs and a blank one: %v, %v, want nil", got, err)
	}
}

func TestFinishWithAnEngineErrorStoresItAsTheReason(t *testing.T) {
	h := newHarness(t)
	h.serverRun("refuse000001", "warrior-fury")
	res := browserResult("warrior-fury", 0)
	res.Lane, res.Error = simapi.LaneServer, "no such talent"
	if err := h.store.Finish(t.Context(), "refuse000001", res); err != nil {
		t.Fatal(err)
	}
	page, err := h.store.Mine(t.Context(), h.owner, 1, "", simapi.LaneServer)
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Rows) != 1 || page.Rows[0].State != StateError ||
		page.Rows[0].FailureReason != "no such talent" || page.Rows[0].FinishedAt == nil {
		t.Fatalf("refused row: %+v", page.Rows)
	}
}
