package server

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/addon"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// testPool connects a migrated pool for a router fixture, skipping the
// test when TEST_DATABASE_URL is unset.
func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

// logsRouter is the whole Phase 3 surface mounted over the test
// database, which is what a deployment with every credential looks
// like.
func logsRouter(t *testing.T) http.Handler {
	t.Helper()
	pool := testPool(t)
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	authStore := &auth.Store{Pool: pool}
	authenticator := &auth.Authenticator{Store: authStore, Log: quiet}
	reportStore := &reports.Store{Pool: pool}
	rankStore := &rankings.Store{Pool: pool}
	return NewRouter(Deps{
		Version: "test", Log: quiet, AllowedOrigin: "https://foreversixty.gg",
		Auth: authenticator,
		Accounts: &auth.Service{
			Store: authStore, Auth: authenticator, PublicBaseURL: "https://foreversixty.gg",
			APIBaseURL: "https://api.foreversixty.gg", Log: quiet,
			BNet: auth.NewBattleNet("id", "secret", "https://api.foreversixty.gg/v1/auth/battlenet/callback"),
		},
		Reports:  &reports.Service{Store: reportStore, Accounts: authStore, Log: quiet},
		Ingest:   &reports.Ingest{Store: reportStore, Put: store.NewDir(t.TempDir()), Log: quiet},
		Uploads:  &reports.Uploads{Store: reportStore, Jobs: &jobs.Fake{}, Log: quiet},
		Rankings: &rankings.Service{Store: rankStore, Log: quiet},
		Addon:    &addon.Service{Store: &addon.Store{Pool: pool}, Log: quiet},
	})
}

// degradedRouter is the shape a deployment with no R2 and no Battle.net
// credentials actually runs in - the Phase 0 environment carried
// requirements 1 and 2 exist to protect. It returns the router and the
// report store behind it, so a test can seed a report to exercise the
// always-mounted file route.
func degradedRouter(t *testing.T) (http.Handler, *reports.Store) {
	t.Helper()
	pool := testPool(t)
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	authStore := &auth.Store{Pool: pool}
	authenticator := &auth.Authenticator{Store: authStore, Log: quiet}
	reportStore := &reports.Store{Pool: pool}
	rankStore := &rankings.Store{Pool: pool}
	h := NewRouter(Deps{
		Version: "test", Log: quiet, AllowedOrigin: "https://foreversixty.gg",
		Auth: authenticator,
		Accounts: &auth.Service{
			Store: authStore, Auth: authenticator, PublicBaseURL: "https://foreversixty.gg",
			APIBaseURL: "https://api.foreversixty.gg", Log: quiet,
			// No BNet: the email-only shape.
		},
		// Reports is mounted (it does not need R2), but its Signer is
		// left nil: the no-R2 shape. Ingest and Uploads are left nil
		// entirely, so their routes are not registered at all.
		Reports:  &reports.Service{Store: reportStore, Accounts: authStore, Log: quiet},
		Rankings: &rankings.Service{Store: rankStore, Log: quiet},
		Addon:    &addon.Service{Store: &addon.Store{Pool: pool}, Log: quiet},
	})
	return h, reportStore
}

func TestEveryPhase3RouteIsMounted(t *testing.T) {
	h := logsRouter(t)
	for _, tc := range []struct{ method, path string }{
		{http.MethodGet, "/v1/auth/battlenet/start"},
		{http.MethodPost, "/v1/auth/email"},
		{http.MethodGet, "/v1/me"},
		{http.MethodDelete, "/v1/sessions"},
		{http.MethodPost, "/v1/devices/pair"},
		{http.MethodPost, "/v1/devices/claim"},
		{http.MethodGet, "/v1/devices"},
		{http.MethodPost, "/v1/reports"},
		{http.MethodGet, "/v1/reports?mine=1"},
		{http.MethodGet, "/v1/reports/abc"},
		{http.MethodGet, "/v1/reports/abc/visibility"},
		{http.MethodPut, "/v1/reports/abc/fights/1"},
		{http.MethodPut, "/v1/reports/abc/raw?offset=0"},
		{http.MethodPost, "/v1/reports/abc/complete"},
		{http.MethodPost, "/v1/uploads"},
		{http.MethodGet, "/v1/rankings?encounter=1"},
		{http.MethodGet, "/v1/rankings/percentile?encounter=1&difficulty=1&value=1"},
		{http.MethodGet, "/v1/rankings/guilds?encounter=1"},
		{http.MethodGet, "/v1/characters/us/hardcore/baelgrim"},
		{http.MethodGet, "/v1/guilds/us/hardcore/forever"},
		{http.MethodGet, "/v1/addon/inbox"},
		{http.MethodPost, "/v1/addon/exports"},
		{http.MethodGet, "/reports/abc/card.png"},
	} {
		r := httptest.NewRequest(tc.method, tc.path, strings.NewReader("{}"))
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		// A mounted route answers about the request; only the router's
		// own catch-all says there is no such route.
		if strings.Contains(w.Body.String(), "no such route") {
			t.Errorf("%s %s is not mounted: %d %s", tc.method, tc.path, w.Code, w.Body.String())
		}
	}
}

// TestRouterDegradesHonestlyWithoutR2OrBattleNet is carried requirements
// 1 and 2: an unconfigured deployment must not offer the routes it has
// no credentials for, and the object-storage routes are the ones this
// task's own wiring is responsible for gating. The report-file route is
// the accepted exception - it stays mounted and answers 503, which the
// Global Constraints' "degrading honestly" language calls for and which
// is a better answer than the router's generic 404 for a route that
// does exist, just not right now.
func TestRouterDegradesHonestlyWithoutR2OrBattleNet(t *testing.T) {
	h, reportStore := degradedRouter(t)

	rep, err := reportStore.Create(context.Background(), reports.Report{
		ID: auth.NewReportID(), Title: "degraded shape", Visibility: reports.Public,
		Status: reports.StatusComplete,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name       string
		method     string
		path       string
		wantStatus int
		wantBody   string
	}{
		{"ingest not mounted", http.MethodPut, "/v1/reports/abc/fights/1", http.StatusNotFound, "no such route"},
		{"uploads not mounted", http.MethodPost, "/v1/uploads", http.StatusNotFound, "no such route"},
		{"battlenet not mounted", http.MethodGet, "/v1/auth/battlenet/start", http.StatusNotFound, "no such route"},
		{
			"report files answer 503, not unmounted", http.MethodGet,
			"/v1/reports/" + rep.ID + "/files/report.json", http.StatusServiceUnavailable,
			"report files are not configured on this deployment",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, tc.path, strings.NewReader("{}"))
			w := httptest.NewRecorder()
			h.ServeHTTP(w, r)
			if w.Code != tc.wantStatus {
				t.Errorf("%s %s = %d, want %d (body %s)", tc.method, tc.path, w.Code, tc.wantStatus, w.Body.String())
			}
			if !strings.Contains(w.Body.String(), tc.wantBody) {
				t.Errorf("%s %s body = %q, want it to contain %q", tc.method, tc.path, w.Body.String(), tc.wantBody)
			}
		})
	}

	// Routes that need neither R2 nor Battle.net still serve normally:
	// this deployment is missing two credentials, not broken.
	for _, tc := range []struct{ method, path string }{
		{http.MethodGet, "/health"},
		{http.MethodPost, "/v1/auth/email"},
		{http.MethodGet, "/v1/rankings?encounter=1"},
		{http.MethodGet, "/reports/" + rep.ID + "/card.png"},
	} {
		r := httptest.NewRequest(tc.method, tc.path, strings.NewReader("{}"))
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		if strings.Contains(w.Body.String(), "no such route") {
			t.Errorf("%s %s should still be mounted: %d %s", tc.method, tc.path, w.Code, w.Body.String())
		}
	}
}

func TestTheCardAndBuildSurfacesSkipTheRouterLimiter(t *testing.T) {
	for path, want := range map[string]bool{
		"/b/abc": true, "/b/abc/card.png": true, "/reports/abc/card.png": true,
		"/v1/reports/abc": false, "/health": false,
	} {
		r := httptest.NewRequest(http.MethodGet, path, nil)
		if got := sharedRead(r); got != want {
			t.Errorf("sharedRead(%s) = %v, want %v", path, got, want)
		}
	}
	post := httptest.NewRequest(http.MethodPost, "/reports/abc/card.png", nil)
	if sharedRead(post) {
		t.Error("only reads skip the limiter")
	}
}

func TestCORSAllowsTheSitesIslandsToCallWithCookies(t *testing.T) {
	h := logsRouter(t)
	r := httptest.NewRequest(http.MethodOptions, "/v1/me", nil)
	r.Header.Set("Origin", "https://foreversixty.gg")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if got := w.Header().Get("Access-Control-Allow-Credentials"); got != "true" {
		t.Fatalf("credentials = %q", got)
	}
	for _, method := range []string{"PUT", "PATCH", "DELETE"} {
		if !strings.Contains(w.Header().Get("Access-Control-Allow-Methods"), method) {
			t.Errorf("methods = %q, want %s", w.Header().Get("Access-Control-Allow-Methods"), method)
		}
	}
	if !strings.Contains(w.Header().Get("Access-Control-Allow-Headers"), "X-CSRF-Token") {
		t.Fatalf("headers = %q", w.Header().Get("Access-Control-Allow-Headers"))
	}
}
