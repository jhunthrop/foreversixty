package auth

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// wantSignInFailed checks a failed web Battle.net callback went back to the site, to want,
// and started no session.
func wantSignInFailed(t *testing.T, res *http.Response, want string) {
	t.Helper()
	res.Body.Close()
	if res.StatusCode != http.StatusFound {
		t.Fatalf("callback = %d, want a redirect back to the site", res.StatusCode)
	}
	if got := res.Header.Get("Location"); got != want {
		t.Fatalf("redirected to %q, want %q", got, want)
	}
	for _, c := range res.Cookies() {
		if c.Name == SessionCookie && c.Value != "" {
			t.Fatal("a failed sign-in must not start a session")
		}
	}
}

func TestSignInFailedPathKeepsTheQueryAndRefusesAnUnsafeNext(t *testing.T) {
	for next, want := range map[string]string{
		"":                             "/setup?signin=failed",
		"/account":                     "/account?signin=failed",
		"/connect?code=kq47tm":         "/connect?code=kq47tm&signin=failed",
		"/setup?signin=ok&refreshed=1": "/setup?refreshed=1&signin=failed",
		"https://evil.example/steal":   "/setup?signin=failed",
		"//evil.example":               "/setup?signin=failed",
		"/\\evil.example":              "/setup?signin=failed",
		"app":                          "/setup?signin=failed",
	} {
		if got := signInFailedPath(next); got != want {
			t.Errorf("signInFailedPath(%q) = %q, want %q", next, got, want)
		}
	}
}

// A next with its own query keeps it on the way back, with signin=failed added.
func TestBattleNetFailureKeepsTheNextQuery(t *testing.T) {
	h := newHarness(t)
	h.withBattleNet(t, "1", "A#1")
	start := h.do(t, http.MethodGet, "/v1/auth/battlenet/start?next="+url.QueryEscape("/connect?code=kq47tm"), "")
	start.Body.Close()
	target, _ := url.Parse(start.Header.Get("Location"))
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=wrong&state="+target.Query().Get("state"), "")
	wantSignInFailed(t, res, "https://foreversixty.gg/connect?code=kq47tm&signin=failed")
}

// An unsafe next was already refused at the start, so the failure returns to the start's
// own fallback, never off the site.
func TestBattleNetFailureNeverRedirectsOffSite(t *testing.T) {
	h := newHarness(t)
	h.withBattleNet(t, "1", "A#1")
	start := h.do(t, http.MethodGet, "/v1/auth/battlenet/start?next=https://evil.example/steal", "")
	start.Body.Close()
	target, _ := url.Parse(start.Header.Get("Location"))
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=wrong&state="+target.Query().Get("state"), "")
	wantSignInFailed(t, res, "https://foreversixty.gg/logs?signin=failed")
}

// The account write failing (the database gone) returns to the site too.
func TestBattleNetAccountWriteFailureReturnsToTheSite(t *testing.T) {
	h := newHarness(t)
	h.withBattleNet(t, "1", "A#1")
	start := h.do(t, http.MethodGet, "/v1/auth/battlenet/start?next=/setup", "")
	start.Body.Close()
	target, _ := url.Parse(start.Header.Get("Location"))
	closed, err := db.Connect(context.Background(), os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	closed.Close()
	h.svc.Store = &Store{Pool: closed}
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state="+target.Query().Get("state"), "")
	wantSignInFailed(t, res, "https://foreversixty.gg/setup?signin=failed")
}
