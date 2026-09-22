// api/internal/bnetapi/client_test.go
package bnetapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// fixtureServer serves canned JSON per exact path+query, recording every
// request it sees so a test can assert on call counts (AppToken caching,
// the retry-once behaviour).
type fixtureServer struct {
	t        *testing.T
	server   *httptest.Server
	handlers map[string]func(w http.ResponseWriter, r *http.Request)
	calls    map[string]*int32
}

func newFixtureServer(t *testing.T) *fixtureServer {
	t.Helper()
	fs := &fixtureServer{t: t, handlers: map[string]func(http.ResponseWriter, *http.Request){}, calls: map[string]*int32{}}
	fs.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Fixtures are keyed without the locale the client always adds
		// (see withLocale); the locale itself is asserted once, below.
		key := r.Method + " " + strings.Replace(r.URL.RequestURI(), "&locale="+apiLocale, "", 1)
		if r.URL.Query().Has("namespace") && r.URL.Query().Get("locale") != apiLocale {
			t.Errorf("fixtureServer: %s carried no locale=%s", r.URL.RequestURI(), apiLocale)
		}
		if n, ok := fs.calls[key]; ok {
			atomic.AddInt32(n, 1)
		} else {
			n := int32(1)
			fs.calls[key] = &n
		}
		h, ok := fs.handlers[key]
		if !ok {
			t.Fatalf("fixtureServer: unexpected request %s", key)
			return
		}
		h(w, r)
	}))
	t.Cleanup(fs.server.Close)
	return fs
}

func (fs *fixtureServer) json(method, path string, status int, body any) {
	fs.handlers[method+" "+path] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if body != nil {
			_ = json.NewEncoder(w).Encode(body)
		}
	}
}

func (fs *fixtureServer) count(method, path string) int32 {
	n, ok := fs.calls[method+" "+path]
	if !ok {
		return 0
	}
	return atomic.LoadInt32(n)
}

// newTestClient wires a Client whose token endpoint and every regional API
// host point at fs, with a near-zero retry backoff.
func newTestClient(fs *fixtureServer) *Client {
	c := New(Config{
		TokenURL: fs.server.URL + "/token",
		APIHost:  func(string) string { return fs.server.URL },
		ClientID: "id", ClientSecret: "secret", Game: "classic1x", Regions: []string{"us", "eu"},
	})
	c.RetryBackoff = time.Millisecond
	return c
}

func TestNamespaces(t *testing.T) {
	c := New(Config{Game: "classic1x"})
	if got := c.ProfileNamespace("US"); got != "profile-classic1x-us" {
		t.Errorf("ProfileNamespace = %q", got)
	}
	if got := c.DynamicNamespace("EU"); got != "dynamic-classic1x-eu" {
		t.Errorf("DynamicNamespace = %q", got)
	}
}

func TestAppTokenIsCachedUntilNearExpiry(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok-1", "expires_in": 3600})
	c := newTestClient(fs)
	ctx := context.Background()

	tok, err := c.AppToken(ctx)
	if err != nil || tok != "tok-1" {
		t.Fatalf("AppToken = %q, %v", tok, err)
	}
	if _, err := c.AppToken(ctx); err != nil {
		t.Fatal(err)
	}
	if n := fs.count(http.MethodPost, "/token"); n != 1 {
		t.Fatalf("token endpoint called %d times, want 1 (cached)", n)
	}
}

func TestAppTokenRefreshesNearExpiry(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok-1", "expires_in": 60})
	c := newTestClient(fs)
	fixedNow := time.Now()
	c.now = func() time.Time { return fixedNow }
	ctx := context.Background()

	if _, err := c.AppToken(ctx); err != nil {
		t.Fatal(err)
	}
	// Past the 60s expiry minus the 60s margin: already stale.
	fixedNow = fixedNow.Add(2 * time.Second)
	if _, err := c.AppToken(ctx); err != nil {
		t.Fatal(err)
	}
	if n := fs.count(http.MethodPost, "/token"); n != 2 {
		t.Fatalf("token endpoint called %d times, want 2 (refreshed near expiry)", n)
	}
}

func TestDoRetriesOnceOn5xxAndNeverOn4xx(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})

	var calls int32
	fs.handlers[http.MethodGet+" /flaky"] = func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&calls, 1)
		if n == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"ok": "yes"})
	}
	c := newTestClient(fs)
	var out map[string]string
	if err := c.getJSON(context.Background(), "flaky", fs.server.URL+"/flaky", &out); err != nil {
		t.Fatal(err)
	}
	if calls != 2 {
		t.Fatalf("flaky endpoint called %d times, want 2 (one retry)", calls)
	}

	fs.json(http.MethodGet, "/notfound", http.StatusNotFound, nil)
	err := c.getJSON(context.Background(), "notfound", fs.server.URL+"/notfound", &out)
	if err == nil {
		t.Fatal("want an error for 404")
	}
	if n := fs.count(http.MethodGet, "/notfound"); n != 1 {
		t.Fatalf("404 endpoint called %d times, want 1 (no retry on 4xx)", n)
	}
}

func TestStatusErrorsClassify(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/nf", http.StatusNotFound, nil)
	fs.json(http.MethodGet, "/fb", http.StatusForbidden, nil)
	fs.json(http.MethodGet, "/rl", http.StatusTooManyRequests, nil)
	c := newTestClient(fs)
	c.RetryBackoff = time.Millisecond

	var out any
	if err := c.getJSON(context.Background(), "op", fs.server.URL+"/nf", &out); !errors.Is(err, ErrNotFound) {
		t.Errorf("nf err = %v, want ErrNotFound", err)
	}
	if err := c.getJSON(context.Background(), "op", fs.server.URL+"/fb", &out); !errors.Is(err, ErrForbidden) {
		t.Errorf("fb err = %v, want ErrForbidden", err)
	}
	if err := c.getJSON(context.Background(), "op", fs.server.URL+"/rl", &out); !errors.Is(err, ErrRateLimited) {
		t.Errorf("rl err = %v, want ErrRateLimited", err)
	}
}
