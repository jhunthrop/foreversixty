package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
