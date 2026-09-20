package sims

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// runBody is a well-formed POST /v1/sims/run body: the same envelope
// the browser lane runs, at the precision setting.
func runBody(t *testing.T) string {
	t.Helper()
	b, err := json.Marshal(simapi.SimRequest{
		EngineVersion: "an old one the page was holding", Spec: "warrior-fury",
		Iterations: preciseIterations,
		Source:     simapi.CharacterSource{Kind: simapi.SourceAddon, Ref: "us/normal/baelgrim"},
		Character:  aCharacter("warrior", "orc"),
	})
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

// bulkBody is a well-formed Top Gear submit: the same envelope with a
// bulk block on it. Iterations is the precision's final-stage count,
// which is what Validate expects of a bulk request (contract A3).
func bulkBody(t *testing.T) string {
	t.Helper()
	b, err := json.Marshal(simapi.SimRequest{
		EngineVersion: "an old one the page was holding", Spec: "warrior-fury",
		Iterations: defaultIterations,
		Source:     simapi.CharacterSource{Kind: simapi.SourceAddon, Ref: "us/normal/baelgrim"},
		Character:  aCharacter("warrior", "orc"),
		Bulk: &simapi.BulkSpec{
			Mode: simapi.KindGear, Precision: simapi.PrecisionNormal,
			// A client-chosen cap the server must overwrite.
			Cap: 1_000_000,
			Candidates: []simapi.Candidate{
				{Slot: "main_hand", ItemID: 19019, Origin: "bag"},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

// refusal reads a failing envelope's code and fields together.
func refusal(t *testing.T, res *http.Response) (string, map[string]string) {
	t.Helper()
	defer res.Body.Close()
	var env struct {
		Error struct {
			Code   string            `json:"code"`
			Fields map[string]string `json:"fields"`
		} `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	return env.Error.Code, env.Error.Fields
}

// assertCapExceededBody checks the one shape both cap-breach paths must
// produce: 400 cap_exceeded with cap and combinations as decimal
// strings (contract 10.6). Both TestABulkRunPastTheLanesCap tests below
// call this with the same numbers, so the two paths cannot drift from
// each other without one of them failing.
func assertCapExceededBody(t *testing.T, res *http.Response, cap, combinations int) {
	t.Helper()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", res.StatusCode)
	}
	code, fields := refusal(t, res)
	if code != "cap_exceeded" {
		t.Fatalf("code %q, want cap_exceeded", code)
	}
	if fields["combinations"] != strconv.Itoa(combinations) || fields["cap"] != strconv.Itoa(cap) {
		t.Fatalf("fields %+v, want cap=%d combinations=%d", fields, cap, combinations)
	}
}

// TestABulkRunPastTheLanesCapIsRefusedWithBothNumbers pins the
// reachable shape: both runner.Native and runner.Fixture (with
// CapBreach set) fail Plan's call itself with a typed
// api.ErrCapExceeded rather than answering a successful over-cap
// summary - the binary refuses the request outright, the same way
// bulk.Count does. checkSize must recover that typed error with
// errors.As and answer the same cap_exceeded body a successful
// over-cap summary would.
func TestABulkRunPastTheLanesCapIsRefusedWithBothNumbers(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	h.planner.err = simapi.ErrCapExceeded{Cap: simapi.Caps[simapi.LaneServer], Combinations: 31200}

	res := h.json(http.MethodPost, "/v1/sims/run", bulkBody(t))
	assertCapExceededBody(t, res, simapi.Caps[simapi.LaneServer], 31200)
	if ran := h.jobs.Ran(); len(ran) != 0 {
		t.Fatalf("a refused run was dispatched anyway: %v", ran)
	}
}

// TestABulkRunPastTheLanesCapViaAnOverCapSummaryIsAlsoRefused covers
// the other shape a Planner can answer with: a plain, successful
// PlanSummary whose Combinations already exceeds its Cap (what a
// Fixture driven by PlanCombinations alone, with no CapBreach,
// returns). checkSize must answer the identical body either way.
func TestABulkRunPastTheLanesCapViaAnOverCapSummaryIsAlsoRefused(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	h.planner.summary = simapi.PlanSummary{
		Kind: simapi.KindGear, Combinations: 31200,
		Cap: simapi.Caps[simapi.LaneServer], IterationsTotal: 100,
	}

	res := h.json(http.MethodPost, "/v1/sims/run", bulkBody(t))
	assertCapExceededBody(t, res, simapi.Caps[simapi.LaneServer], 31200)
	if ran := h.jobs.Ran(); len(ran) != 0 {
		t.Fatalf("a refused run was dispatched anyway: %v", ran)
	}
}

// TestAPlannerFailureThatIsNotACapBreachFailsTheSubmit is the third
// case checkSize must tell apart from a cap breach: a genuine failure
// to size the request (the binary crashed, timed out, or refused for
// a reason other than the cap) must still fail the submit with a
// generic 500 - it must never fall through and queue a run nobody
// sized.
func TestAPlannerFailureThatIsNotACapBreachFailsTheSubmit(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	h.planner.err = errAnyway

	res := h.json(http.MethodPost, "/v1/sims/run", bulkBody(t))
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status %d, want 500", res.StatusCode)
	}
	if code := h.errorCode(res); code != "internal" {
		t.Fatalf("code %q, want internal", code)
	}
	if ran := h.jobs.Ran(); len(ran) != 0 {
		t.Fatalf("an unsized run was dispatched anyway: %v", ran)
	}
}

func TestTheServerSetsTheCapItself(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	if res := h.json(http.MethodPost, "/v1/sims/run", bulkBody(t)); res.StatusCode != http.StatusAccepted {
		t.Fatalf("status %d, want 202", res.StatusCode)
	}
	if len(h.planner.asked) != 1 {
		t.Fatalf("%d plans", len(h.planner.asked))
	}
	if got := h.planner.asked[0].Bulk.Cap; got != simapi.Caps[simapi.LaneServer] {
		t.Errorf("cap %d, want the server lane's %d; a client may not raise its own bound",
			got, simapi.Caps[simapi.LaneServer])
	}
}

func TestAPlainRunIsNeverPlanned(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	if res := h.json(http.MethodPost, "/v1/sims/run", runBody(t)); res.StatusCode != http.StatusAccepted {
		t.Fatalf("status %d, want 202", res.StatusCode)
	}
	if len(h.planner.asked) != 0 {
		t.Errorf("a plain run has nothing to expand, but the planner was asked: %+v", h.planner.asked)
	}
}

func TestAnAccountWithoutPremiumIsAnswered402(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = false
	res := h.json(http.MethodPost, "/v1/sims/run", runBody(t))
	if res.StatusCode != http.StatusPaymentRequired {
		t.Fatalf("status %d, want 402", res.StatusCode)
	}
	if code := h.errorCode(res); code != "premium_required" {
		t.Fatalf("code %q, want premium_required", code)
	}
	if ran := h.jobs.Ran(); len(ran) != 0 {
		t.Fatalf("a job was started for a non-premium account: %v", ran)
	}
}

func TestAPremiumRunIsQueuedDispatchedAndPollable(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	res := h.json(http.MethodPost, "/v1/sims/run", runBody(t))
	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("status %d, want 202", res.StatusCode)
	}
	var out struct {
		SimID string `json:"sim_id"`
	}
	h.data(res, &out)
	if len(out.SimID) != 12 {
		t.Fatalf("sim id %q", out.SimID)
	}

	ran := h.jobs.Ran()
	if len(ran) != 1 || ran[0][0] != SimRunJobCommand || ran[0][1] != out.SimID {
		t.Fatalf("dispatched %v, want [%s %s]", ran, SimRunJobCommand, out.SimID)
	}

	var p Progress
	h.data(h.do(http.MethodGet, "/v1/sims/"+out.SimID+"/progress", "", nil), &p)
	if p.State != StateQueued {
		t.Fatalf("state %q, want %q", p.State, StateQueued)
	}

	stored, err := h.store.Get(t.Context(), out.SimID)
	if err != nil {
		t.Fatal(err)
	}
	// The deployment's own pin wins over whatever the page sent.
	if stored.Request.EngineVersion != testEngine {
		t.Fatalf("engine version %q, want the deployment's pin", stored.Request.EngineVersion)
	}
	// The encounter the page left empty is the design's default.
	// EncounterSpec now carries TargetsOverTime ([]TargetCount), so it is
	// no longer comparable with ==; DeepEqual is the correct successor.
	if !reflect.DeepEqual(stored.Request.Encounter, simapi.DefaultEncounter()) {
		t.Fatalf("encounter defaults were not applied: %+v", stored.Request.Encounter)
	}
	// And the whole character is in the row, which is what the job
	// reads back: there is no second copy in the bucket.
	if stored.Request.Character.Race != "orc" || len(stored.Request.Character.Gear) == 0 {
		t.Fatalf("the queued row lost the character: %+v", stored.Request.Character)
	}
}

func TestARunTheEnvelopeRejectsIsRefused(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	b, err := json.Marshal(simapi.SimRequest{
		Spec: "warrior-fury", Iterations: defaultIterations,
		Source: simapi.CharacterSource{Kind: simapi.SourceManual},
		// No character at all: no class, no race, no level.
	})
	if err != nil {
		t.Fatal(err)
	}
	res := h.json(http.MethodPost, "/v1/sims/run", string(b))
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", res.StatusCode)
	}
}

func TestARunNobodyCanStartIsRecordedAsFailed(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	h.jobs.Err = errAnyway
	res := h.json(http.MethodPost, "/v1/sims/run", runBody(t))
	if res.StatusCode != http.StatusBadGateway {
		t.Fatalf("status %d, want 502", res.StatusCode)
	}
	// The row exists and says it failed, so the page stops polling.
	var n int
	if err := h.store.Pool.QueryRow(t.Context(),
		`select count(*) from sims where state = $1`, StateError).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("%d failed rows, want 1", n)
	}
}

// cancelingJobs fails a dispatch the way a client disconnect does:
// the same event that makes Jobs.Run fail also cancels the request
// context, so the compensating write right after it cannot rely on
// that context still being alive.
type cancelingJobs struct{ cancel context.CancelFunc }

func (j cancelingJobs) Run(ctx context.Context, args ...string) error {
	j.cancel()
	return errAnyway
}

func TestARunNobodyCanStartIsRecordedAsFailedEvenWhenTheRequestContextIsDone(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true

	ctx, cancel := context.WithCancel(auth.WithActor(context.Background(), h.actor))
	h.service.Jobs = cancelingJobs{cancel: cancel}

	r := httptest.NewRequest(http.MethodPost, "/v1/sims/run", strings.NewReader(runBody(t))).WithContext(ctx)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.service.run(w, r)

	if w.Code != http.StatusBadGateway {
		t.Fatalf("status %d, want 502", w.Code)
	}
	var n int
	if err := h.store.Pool.QueryRow(t.Context(),
		`select count(*) from sims where state = $1`, StateError).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("%d failed rows, want 1: the compensating write must not ride the now-cancelled request context", n)
	}
}

func TestABulkRunPastTheBudgetIsRefusedWithItsEstimate(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	// Exactly four thousand seconds of engine time, whatever the
	// benchmark's current figure is.
	h.planner.summary = simapi.PlanSummary{
		Kind: simapi.KindGear, Combinations: 4000,
		Cap: simapi.Caps[simapi.LaneServer], IterationsTotal: nativeRate * 4000,
	}

	res := h.json(http.MethodPost, "/v1/sims/run", bulkBody(t))
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400", res.StatusCode)
	}
	code, fields := refusal(t, res)
	if code != "too_large" {
		t.Fatalf("code %q, want too_large", code)
	}
	if fields["estimate_sec"] != "4000" ||
		fields["budget_sec"] != strconv.Itoa(int(BulkBudget.Seconds())) {
		t.Fatalf("fields %+v", fields)
	}
	if ran := h.jobs.Ran(); len(ran) != 0 {
		t.Fatalf("a refused run was dispatched anyway: %v", ran)
	}
}

// TestAWeightsRunIsAcceptedUnsizedAndBoundOnlyByTheJobTimeout pins the
// honest current behaviour, not the brief's original assertion that a
// weights request is planned: runner.Native.Plan and runner.Fixture.Plan
// both refuse req.Bulk == nil with ErrBadInput (sim/runner/native.go,
// sim/runner/fixture.go), and a weights request carries req.Weights
// instead of req.Bulk. checkSize's req.Bulk != nil guard (Task 5) is
// therefore load-bearing, not incidental: without it, every weights
// submit would 500 rather than skip a check it has no combinations to
// answer.
//
// That means a weights run today has no submit-time size estimate at
// all - it is bounded only by timeoutFor's BulkBudget once it is
// running, the same as an unusually large bulk run that slipped under
// its cap. sim/api exports WeightsSpec, StatWeight and an unexported
// validate, and nothing that costs a weights run in iterations; adding
// one would mean guessing whether the engine runs one sim per stat or
// two (plus/minus delta), and that guess would become a number shown
// to a user in a refusal. For a weights run to get an estimate, the
// module needs to publish its own iteration-cost function the way
// sim/api.LadderIterations does for bulk, and this package would then
// call it here exactly as checkSize already calls the planner for bulk.
func TestAWeightsRunIsAcceptedUnsizedAndBoundOnlyByTheJobTimeout(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	// Cap 0: weights expand to no combinations, so the cap test must not
	// fire on a zero and refuse every one of them.
	h.planner.summary = simapi.PlanSummary{
		Kind: simapi.KindWeights, Combinations: 0, Cap: 0, IterationsTotal: 60_000,
	}
	b, err := json.Marshal(simapi.SimRequest{
		EngineVersion: testEngine, Spec: "warrior-fury", Iterations: defaultIterations,
		Source:    simapi.CharacterSource{Kind: simapi.SourceAddon, Ref: "us/normal/baelgrim"},
		Character: aCharacter("warrior", "orc"),
		Weights: &simapi.WeightsSpec{
			Stats: []string{"strength", "melee_crit"}, Reference: "melee_crit",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if res := h.json(http.MethodPost, "/v1/sims/run", string(b)); res.StatusCode != http.StatusAccepted {
		t.Fatalf("status %d, want 202", res.StatusCode)
	}
	// Not planned: a weights request has no Bulk block, so checkSize's
	// req.Bulk != nil guard skips the planner call entirely.
	if len(h.planner.asked) != 0 {
		t.Fatalf("a weights request has nothing for the planner to size, but it was asked: %d plans",
			len(h.planner.asked))
	}
	if ran := h.jobs.Ran(); len(ran) != 1 {
		t.Fatalf("%d jobs dispatched, want 1: an accepted weights run must still be queued", len(ran))
	}
}

func TestTheRunRouteNeedsASession(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	h.anonymous()
	if res := h.json(http.MethodPost, "/v1/sims/run", runBody(t)); res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status %d, want 401", res.StatusCode)
	}
}
