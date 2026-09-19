package sims

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestTheRunRouteNeedsASession(t *testing.T) {
	h := newHarness(t)
	h.premium.premium = true
	h.anonymous()
	if res := h.json(http.MethodPost, "/v1/sims/run", runBody(t)); res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status %d, want 401", res.StatusCode)
	}
}
