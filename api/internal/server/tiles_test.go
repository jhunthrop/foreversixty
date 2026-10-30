package server

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// tilesRouter mounts the tiles over the trees fixture, with build test-1's icon folder holding
// the one icon the route needs for an item it does not know: the question mark.
func tilesRouter(t *testing.T) http.Handler {
	t.Helper()
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	icons := filepath.Join(root, "test-1", "icons")
	if err := os.MkdirAll(icons, 0o755); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(filepath.Join("..", "sheet", "testdata", "icons", sheet.QuestionMark+".webp"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(icons, sheet.QuestionMark+".webp"), raw, 0o600); err != nil {
		t.Fatal(err)
	}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	return NewRouter(Deps{Version: "test-1", Log: quiet, Tiles: sheet.NewTiles(root, sheet.TreesItems{Data: data}, quiet)})
}

func TestTheTilesRouteIsMountedAndSkipsTheRouterLimit(t *testing.T) {
	h := tilesRouter(t)
	frozen := time.Date(2026, time.October, 29, 0, 0, 0, 0, time.UTC)
	httpx.Now = func() time.Time { return frozen }
	t.Cleanup(func() { httpx.Now = time.Now })
	get := func(path string) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.RemoteAddr = "198.51.100.9:4321"
		h.ServeHTTP(rec, req)
		return rec
	}
	// Item 1 is not in test-1: the question mark, in the empty ring.
	rec := get(sheet.TilePath + "test-1/1.png")
	if rec.Code != http.StatusOK || rec.Header().Get("Content-Type") != "image/png" {
		t.Fatalf("code %d type %q cache %q", rec.Code, rec.Header().Get("Content-Type"), rec.Header().Get("Cache-Control"))
	}
	if rec := get(sheet.TilePath + "no-such-build/1.png"); rec.Code != http.StatusNotFound {
		t.Fatalf("unknown build: %d", rec.Code)
	}
	if rec := get(sheet.TilePath + "test-1/x.png"); rec.Code != http.StatusNotFound {
		t.Fatalf("non-digit item: %d", rec.Code)
	}
	if cc := rec.Header().Get("Cache-Control"); cc != "public, max-age=3600" {
		t.Fatalf("an unknown item's question mark is cached %q, want an hour", cc)
	}
	// Discord's media proxy fetches from a few IPs: past the router budget, tiles still answer
	// while everything else is limited.
	for i := 0; i < routerRequestsPerMinute; i++ {
		get("/version")
	}
	if rec := get(sheet.TilePath + "test-1/12640.png"); rec.Code != http.StatusOK {
		t.Fatalf("a tile past the budget: %d, want 200", rec.Code)
	}
	if rec := get("/version"); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("/version past the budget: %d, want 429", rec.Code)
	}
}
