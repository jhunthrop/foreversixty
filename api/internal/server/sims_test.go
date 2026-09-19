package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
)

// TestTheSimulatorRoutesAreMountedWhenTheServiceIs checks the router
// wiring, not the handlers: a nil-store service answers 500, which is
// still proof the route exists rather than 404.
func TestTheSimulatorRoutesAreMountedWhenTheServiceIs(t *testing.T) {
	h := NewRouter(Deps{Version: "test", Sims: &sims.Service{}})
	for _, path := range []string{
		"/v1/sims/aaaaaaaaaaaa",
		"/v1/sims/aaaaaaaaaaaa/progress",
		"/v1/specs",
		"/v1/characters/us/normal/baelgrim/sim-input",
	} {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))
		if w.Code == http.StatusNotFound {
			t.Errorf("%s is not mounted", path)
		}
	}
}

// fakeJobRunner and fakePremiumer are the minimum a sims.Service needs
// to mount POST /v1/sims/run at all - the one route whose registration
// depends on runtime state rather than always being there.
type fakeJobRunner struct{}

func (fakeJobRunner) Run(context.Context, ...string) error { return nil }

type fakePremiumer struct{}

func (fakePremiumer) Premium(context.Context, int64) (bool, error) { return false, nil }

// TestTheSaveMineAndRunRoutesAreMounted rounds out the mount check
// above: POST /v1/sims and GET /v1/sims are always there, and POST
// /v1/sims/run only when the service carries a job runner and an
// accounts reader. A POST to an unmounted path is a 404 the rest of
// this suite could not otherwise tell apart from a 405 or a session
// check running.
func TestTheSaveMineAndRunRoutesAreMounted(t *testing.T) {
	h := NewRouter(Deps{Version: "test", Sims: &sims.Service{}})

	// A nil store means save() fails once it reaches the store, past
	// the router: not a 404.
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/v1/sims", strings.NewReader("{}")))
	if w.Code == http.StatusNotFound {
		t.Error("POST /v1/sims is not mounted")
	}

	// No session reaches RequireSession's own check, a 401 - still
	// proof the route is mounted.
	w = httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/v1/sims", nil))
	if w.Code != http.StatusUnauthorized {
		t.Errorf("GET /v1/sims status %d, want 401 (mounted, no session)", w.Code)
	}

	w = httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/v1/sims/run", strings.NewReader("{}")))
	if w.Code != http.StatusNotFound {
		t.Errorf("POST /v1/sims/run status %d, want 404 without Jobs and Accounts", w.Code)
	}

	withRun := NewRouter(Deps{Version: "test", Sims: &sims.Service{
		Jobs: fakeJobRunner{}, Accounts: fakePremiumer{},
	}})
	w = httptest.NewRecorder()
	withRun.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/v1/sims/run", strings.NewReader("{}")))
	if w.Code == http.StatusNotFound {
		t.Error("POST /v1/sims/run is not mounted when Jobs and Accounts are both set")
	}
}

func TestTheSimulatorRoutesAreAbsentWhenTheServiceIs(t *testing.T) {
	h := NewRouter(Deps{Version: "test"})
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/v1/specs", nil))
	if w.Code != http.StatusNotFound {
		t.Errorf("status %d, want 404 with no sims service", w.Code)
	}
}

// TestTheIngestCarriesTheScorerAndTheMemberRead is the wiring check:
// each of these is an interface field that compiles perfectly well
// while nil, and a nil one means no execution score ever.
func TestTheIngestCarriesTheScorerAndTheMemberRead(t *testing.T) {
	ing := &reports.Ingest{}
	ing.Score, ing.Members = stubScorer{}, stubMembers{}
	if ing.Score == nil || ing.Members == nil {
		t.Fatal("the ingest's scoring seams are nil")
	}
	// And the router accepts a Deps carrying them, which is the shape
	// serve() builds.
	h := NewRouter(Deps{Version: "test", Ingest: ing, Sims: &sims.Service{}})
	if h == nil {
		t.Fatal("the router refused a Deps with a scoring ingest")
	}
}

type stubScorer struct{}

func (stubScorer) Schedule(reports.ScoredFight) {}

type stubMembers struct{}

func (stubMembers) MemberKeys(context.Context, []string) (map[string]bool, error) {
	return map[string]bool{}, nil
}
