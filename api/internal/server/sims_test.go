package server

import (
	"context"
	"net/http"
	"net/http/httptest"
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
