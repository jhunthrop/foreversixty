package server

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
	"github.com/jhunthrop/foreversixty/api/internal/site"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func TestHealthAndVersion(t *testing.T) {
	h := NewRouter(Deps{Version: "test-1"})
	for _, tc := range []struct{ path, want string }{
		{"/health", `"status":"ok"`},
		{"/version", `"version":"test-1"`},
	} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, tc.path, nil))
		if rec.Code != 200 || !strings.Contains(rec.Body.String(), tc.want) {
			t.Fatalf("%s: code=%d body=%s", tc.path, rec.Code, rec.Body.String())
		}
	}
}

func TestThePhasesRouteServesTheBoundaries(t *testing.T) {
	h := NewRouter(Deps{Version: "test"})
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/v1/phases", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("status %d, want 200", w.Code)
	}
	var env struct {
		OK   bool `json:"ok"`
		Data struct {
			Phases []struct {
				Name  string    `json:"name"`
				Start time.Time `json:"start"`
			} `json:"phases"`
		} `json:"data"`
	}
	if err := json.NewDecoder(w.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	if !env.OK || len(env.Data.Phases) != len(phase.Boundaries) {
		t.Fatalf("envelope: %+v", env)
	}
	if env.Data.Phases[0].Name != phase.Boundaries[0].Name {
		t.Errorf("first phase %q, want %q", env.Data.Phases[0].Name, phase.Boundaries[0].Name)
	}
	want := "public, max-age=3600"
	if cc := w.Header().Get("Cache-Control"); cc != want {
		t.Errorf("Cache-Control = %q, want %q", cc, want)
	}
}

func TestUnknownRouteReturnsEnvelope(t *testing.T) {
	h := NewRouter(Deps{Version: "test-1"})
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/nope", nil))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("code = %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.HasPrefix(ct, "application/json") {
		t.Fatalf("content-type = %q", ct)
	}
	body := rec.Body.String()
	if !strings.Contains(body, `"ok":false`) || !strings.Contains(body, `"code":"not_found"`) {
		t.Fatalf("unexpected body: %s", body)
	}
}

func TestRouterWithoutBuildsOrSiteStillServesHealth(t *testing.T) {
	h := NewRouter(Deps{Version: "test-1"})
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/b/znorjmts", nil))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("code = %d, want the catch-all 404 when the page is not mounted", rec.Code)
	}
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/health", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("health code = %d", rec.Code)
	}
}

// oneBuildGetter is a site.BuildGetter that always returns the same record,
// so the shared build page renders without a database.
type oneBuildGetter struct{ b builds.Build }

func (g oneBuildGetter) Get(context.Context, string) (builds.Build, error) { return g.b, nil }

func TestSharedBuildPageStaysServedAfterTheRouterWideBudgetIsSpent(t *testing.T) {
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	b, err := builds.New(builds.Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", PointOrder: []int{101}})
	if err != nil {
		t.Fatal(err)
	}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	h := NewRouter(Deps{
		Version: "test-1",
		Log:     quiet,
		Site: &site.Deps{
			Store:         oneBuildGetter{b: b},
			Data:          data,
			PublicBaseURL: "https://foreversixty.gg",
			Log:           quiet,
		},
		TrustedProxyHops: 1,
	})
	// Freeze the limiter's clock: with real time, a slow runner refills part of
	// the bucket during the loop below and the request past the budget is let
	// through, which CI observed. The budget itself, not the refill, is under test.
	frozen := time.Date(2026, time.September, 14, 0, 0, 0, 0, time.UTC)
	httpx.Now = func() time.Time { return frozen }
	t.Cleanup(func() { httpx.Now = time.Now })

	get := func(path string) int {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.RemoteAddr = "198.51.100.7:4321"
		h.ServeHTTP(rec, req)
		return rec.Code
	}

	// Spend the whole router-wide per-minute budget from one address.
	for i := 0; i < routerRequestsPerMinute; i++ {
		if code := get("/version"); code != http.StatusOK {
			t.Fatalf("/version request %d code = %d, want 200", i+1, code)
		}
	}
	if code := get("/b/" + b.ID); code != http.StatusOK {
		t.Fatalf("/b/{id} past the budget code = %d, want 200 (the page must be exempt)", code)
	}
	if code := get("/b/" + b.ID + "/card.png"); code != http.StatusOK {
		t.Fatalf("/b/{id}/card.png past the budget code = %d, want 200 (the card must be exempt)", code)
	}
	// ...and the limiter is still armed for everything else.
	if code := get("/version"); code != http.StatusTooManyRequests {
		t.Fatalf("/version past the budget code = %d, want 429", code)
	}
}
