// api/internal/bnetapi/realms_test.go
package bnetapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
)

func TestRulesetOf(t *testing.T) {
	cases := []struct {
		name string
		r    Realm
		want string
	}{
		{"normal", Realm{Type: "NORMAL"}, character.RulesetNormal},
		{"pvp", Realm{Type: "PVP"}, character.RulesetPvP},
		{"rp", Realm{Type: "RP"}, character.RulesetRP},
		{"rp_pvp", Realm{Type: "RP_PVP"}, character.RulesetRP},
		{"forever pvp_rp", Realm{Type: "PVP_RP"}, character.RulesetRP},
		{"hardcore by category", Realm{Type: "PVP", Category: "Hardcore"}, character.RulesetHardcore},
		{"hardcore by name", Realm{Type: "NORMAL", Name: "Doomhowl Hardcore"}, character.RulesetHardcore},
		{"unknown type resolves nothing", Realm{Type: "SEASONAL"}, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got, ok := RulesetOf(c.r); got != c.want || ok != (c.want != "") {
				t.Errorf("RulesetOf(%+v) = %q, %v; want %q", c.r, got, ok, c.want)
			}
		})
	}
}

func TestRealmsFetchesIndexAndPerRealmDetailAndCaches(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us", http.StatusOK, map[string]any{
		"realms": []map[string]any{
			{"id": 1, "name": "Whitemane", "slug": "whitemane"},
			{"id": 2, "name": "Faerlina", "slug": "faerlina"},
		},
	})
	fs.json(http.MethodGet, "/data/wow/realm/whitemane?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"type": map[string]string{"type": "PVP"}, "category": "PvP"})
	fs.json(http.MethodGet, "/data/wow/realm/faerlina?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"type": map[string]string{"type": "NORMAL"}, "category": "PvE"})

	c := newTestClient(fs)
	realms, err := c.Realms(context.Background(), "us")
	if err != nil {
		t.Fatal(err)
	}
	if len(realms) != 2 || realms[0].Slug != "whitemane" || realms[0].Type != "PVP" {
		t.Fatalf("realms = %+v", realms)
	}

	// A second call must hit the cache, not the index endpoint again.
	if _, err := c.Realms(context.Background(), "us"); err != nil {
		t.Fatal(err)
	}
	if n := fs.count(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us"); n != 1 {
		t.Fatalf("realm index called %d times, want 1 (cached)", n)
	}
}

// TestRealmsNeverCachesAPartialListOnDetailFailure is the regression test
// for spec A1's production defect: a partial realm list (one realm's
// detail failed) must never be cached, or the next lookup for that realm
// silently reuses a wrong (default "normal") ruleset for up to 24h.
func TestRealmsNeverCachesAPartialListOnDetailFailure(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"realms": []map[string]any{
			{"id": 1, "name": "Whitemane", "slug": "whitemane"},
			{"id": 2, "name": "Broken", "slug": "broken"},
		}})
	fs.json(http.MethodGet, "/data/wow/realm/whitemane?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"type": map[string]string{"type": "PVP"}, "category": "PvP"})
	fs.json(http.MethodGet, "/data/wow/realm/broken?namespace=dynamic-classic1x-us", http.StatusInternalServerError, nil)

	c := newTestClient(fs)
	if _, err := c.Realms(context.Background(), "us"); err == nil {
		t.Fatal("want an error when a realm detail fails, not a partial list")
	}
	c.realmMu.Lock()
	_, cached := c.realmCache["us"]
	c.realmMu.Unlock()
	if cached {
		t.Fatal("a failed batch must not be cached")
	}
}

// TestRealmsFetchesDetailsConcurrently is the regression test for the
// other half of A1: 66 realms fetched serially took ~7s on Era; this
// asserts the worker pool actually overlaps requests rather than merely
// compiling one.
func TestRealmsFetchesDetailsConcurrently(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})

	const realmCount = 16
	var realmsList []map[string]any
	for i := 1; i <= realmCount; i++ {
		slug := fmt.Sprintf("realm%d", i)
		realmsList = append(realmsList, map[string]any{"id": i, "name": slug, "slug": slug})
	}
	fs.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"realms": realmsList})

	var mu sync.Mutex
	inFlight, maxInFlight := 0, 0
	for i := 1; i <= realmCount; i++ {
		slug := fmt.Sprintf("realm%d", i)
		fs.handlers[http.MethodGet+" /data/wow/realm/"+slug+"?namespace=dynamic-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
			mu.Lock()
			inFlight++
			if inFlight > maxInFlight {
				maxInFlight = inFlight
			}
			mu.Unlock()
			time.Sleep(20 * time.Millisecond)
			mu.Lock()
			inFlight--
			mu.Unlock()
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"type": map[string]string{"type": "NORMAL"}, "category": "PvE"})
		}
	}

	c := newTestClient(fs)
	start := time.Now()
	realms, err := c.Realms(context.Background(), "us")
	elapsed := time.Since(start)
	if err != nil {
		t.Fatal(err)
	}
	if len(realms) != realmCount {
		t.Fatalf("realms = %d, want %d", len(realms), realmCount)
	}
	// Serial would be realmCount*20ms = 320ms; realmWorkerCount=8 in
	// flight should finish in about two batches (~40ms) plus overhead.
	if elapsed > 200*time.Millisecond {
		t.Fatalf("elapsed = %s, want well under the serial time thanks to concurrency", elapsed)
	}
	mu.Lock()
	got := maxInFlight
	mu.Unlock()
	if got < 2 {
		t.Fatalf("maxInFlight = %d, want concurrent requests (>1)", got)
	}
}

// recordingHandler keeps every slog record for assertions on what a client logged.
type recordingHandler struct {
	mu      sync.Mutex
	records []slog.Record
}

func (h *recordingHandler) Enabled(context.Context, slog.Level) bool { return true }
func (h *recordingHandler) Handle(_ context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.records = append(h.records, r)
	return nil
}
func (h *recordingHandler) WithAttrs([]slog.Attr) slog.Handler { return h }
func (h *recordingHandler) WithGroup(string) slog.Handler      { return h }

// warnings counts the WARN records whose "op" attribute is op.
func (h *recordingHandler) warnings(op string) int {
	h.mu.Lock()
	defer h.mu.Unlock()
	n := 0
	for _, r := range h.records {
		if r.Level != slog.LevelWarn {
			continue
		}
		r.Attrs(func(a slog.Attr) bool {
			if a.Key == "op" && a.Value.String() == op {
				n++
			}
			return true
		})
	}
	return n
}

// Forever's realm index has answered 404 since 2026-11-05 (and still did on
// 2026-11-12, as did the connected-realm index and every realm by id or slug).
// That is a fact about the namespace, not a failure of the run: Realms answers
// an empty list with no error, remembers the absence for an hour, and says so
// once, so an import a minute is not a WARN a minute.
func TestRealmsRemembersAMissingIndexForAnHourAndWarnsOnce(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	const index = "/data/wow/realm/index?namespace=dynamic-forever-us"
	fs.json(http.MethodGet, index, http.StatusNotFound, map[string]any{"code": 404, "detail": "Not Found"})

	logs := &recordingHandler{}
	c := newTestClient(fs)
	c.Game = "forever"
	c.Log = slog.New(logs)
	now := time.Date(2026, 11, 12, 16, 0, 0, 0, time.UTC)
	c.now = func() time.Time { return now }

	for i := 0; i < 3; i++ {
		realms, err := c.Realms(context.Background(), "us")
		if err != nil || len(realms) != 0 {
			t.Fatalf("call %d: realms = %v, err = %v; want an empty list and no error", i, realms, err)
		}
	}
	if n := fs.count(http.MethodGet, index); n != 1 {
		t.Fatalf("realm index fetched %d times in an hour, want 1", n)
	}
	if n := logs.warnings("realm_index"); n != 1 {
		t.Fatalf("realm_index warned %d times, want once", n)
	}

	now = now.Add(realmIndexMissingTTL + time.Second)
	if _, err := c.Realms(context.Background(), "us"); err != nil {
		t.Fatal(err)
	}
	if n := fs.count(http.MethodGet, index); n != 2 {
		t.Fatalf("realm index fetched %d times after the hour, want 2", n)
	}
	if n := logs.warnings("realm_index"); n != 2 {
		t.Fatalf("realm_index warned %d times after the hour, want 2", n)
	}
}

// Any other failure (a 403 namespace, a 5xx) stays an error and is never
// remembered: the next call tries again.
func TestRealmsDoesNotRememberOtherIndexFailures(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	const index = "/data/wow/realm/index?namespace=dynamic-classic1x-us"
	fs.json(http.MethodGet, index, http.StatusForbidden, nil)
	c := newTestClient(fs)
	for i := 0; i < 2; i++ {
		if _, err := c.Realms(context.Background(), "us"); !errors.Is(err, ErrForbidden) {
			t.Fatalf("call %d: err = %v, want ErrForbidden", i, err)
		}
	}
	if n := fs.count(http.MethodGet, index); n != 2 {
		t.Fatalf("realm index fetched %d times, want 2 (a 403 is not cached)", n)
	}
}
