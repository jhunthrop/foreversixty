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
	page, err := h.store.Mine(t.Context(), h.owner, 1, "")
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
	if err := h.store.Advance(t.Context(), "hhhhhhhhhhhh", 1500, 1000); err != nil {
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
	if err := h.store.Advance(t.Context(), "iiiiiiiiiiii", 1500, 1000); err != nil {
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

	all, err := h.store.Mine(t.Context(), h.owner, 1, "")
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
	if got := byID["aaaaaaaaaaaa"]; got.Kind != simapi.KindRun || got.Headline != "1,204 DPS" {
		t.Errorf("plain row: %+v", got)
	}
	if got := byID["bbbbbbbbbbbb"]; got.Kind != simapi.KindGear ||
		got.Headline != "+41 DPS from Vis'kag the Bloodletter" {
		t.Errorf("gear row: %+v", got)
	}

	only, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindGear)
	if err != nil {
		t.Fatal(err)
	}
	if only.Total != 1 || len(only.Rows) != 1 || only.Rows[0].SimID != "bbbbbbbbbbbb" {
		t.Fatalf("filtered: total %d rows %+v", only.Total, only.Rows)
	}

	none, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindDrops)
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
	queued, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindWeights)
	if err != nil {
		t.Fatal(err)
	}
	if len(queued.Rows) != 1 || queued.Rows[0].Headline != "" {
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
	finished, err := h.store.Mine(t.Context(), h.owner, 1, simapi.KindWeights)
	if err != nil {
		t.Fatal(err)
	}
	if len(finished.Rows) != 1 || finished.Rows[0].Headline != "Crit 1.00 · Agility 0.87" {
		t.Fatalf("finished row: %+v", finished.Rows)
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
