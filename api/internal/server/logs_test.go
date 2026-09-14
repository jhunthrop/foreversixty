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

	"github.com/jhunthrop/foreversixty/api/internal/addon"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// logsRouter is the whole Phase 3 surface mounted over the test
// database, which is what a deployment with every credential looks
// like.
func logsRouter(t *testing.T) http.Handler {
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
