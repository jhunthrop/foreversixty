package auth

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// devLoginServer mounts the auth routes with the dev sign-in on or off, over the test
// database, on a loopback httptest server.
func devLoginServer(t *testing.T, on bool) (*httptest.Server, *Store) {
	t.Helper()
	h := newHarness(t)
	svc := *h.svc
	svc.DevLogin = on
	mux := http.NewServeMux()
	Mount(mux, &svc, 0)
	srv := httptest.NewServer(svc.Auth.Middleware(mux))
	t.Cleanup(srv.Close)
	return srv, svc.Store
}

func postDevLogin(t *testing.T, srv *httptest.Server, battletag string) (*http.Response, string) {
	t.Helper()
	res, err := http.Post(srv.URL+"/v1/dev/login", "application/json",
		strings.NewReader(`{"battletag":"`+battletag+`"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return res, string(body)
}

func TestDevLoginSignsInASeededAccountOnALoopbackPeer(t *testing.T) {
	srv, store := devLoginServer(t, true)
	if _, err := store.UpsertBnetUser(context.Background(), "dev-sub-1", "Seeded#0001"); err != nil {
		t.Fatal(err)
	}
	res, body := postDevLogin(t, srv, "Seeded#0001")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, body %s", res.StatusCode, body)
	}
	var session *http.Cookie
	for _, c := range res.Cookies() {
		if c.Name == SessionCookie {
			session = c
		}
	}
	if session == nil || session.Value == "" {
		t.Fatalf("no %s cookie in %v", SessionCookie, res.Cookies())
	}
	req, _ := http.NewRequest(http.MethodGet, srv.URL+"/v1/me", nil)
	req.AddCookie(session)
	me, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer me.Body.Close()
	if me.StatusCode != http.StatusOK {
		t.Fatalf("GET /v1/me = %d", me.StatusCode)
	}
}

// An unknown battletag is a 404: the route signs in a seeded account, it never creates one.
func TestDevLoginNeverCreatesAnAccount(t *testing.T) {
	srv, store := devLoginServer(t, true)
	res, _ := postDevLogin(t, srv, "Nobody#0000")
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", res.StatusCode)
	}
	if _, err := store.UserByBattletag(context.Background(), "Nobody#0000"); err == nil {
		t.Fatal("the dev sign-in created an account")
	}
}

// With the flag off the route is not mounted at all, which is what every deployment runs.
func TestDevLoginRouteIsAbsentWhenTheFlagIsOff(t *testing.T) {
	srv, store := devLoginServer(t, false)
	if _, err := store.UpsertBnetUser(context.Background(), "dev-sub-2", "Seeded#0002"); err != nil {
		t.Fatal(err)
	}
	res, _ := postDevLogin(t, srv, "Seeded#0002")
	if res.StatusCode == http.StatusOK {
		t.Fatal("POST /v1/dev/login answered 200 with FS_DEV_LOGIN off")
	}
	for _, c := range res.Cookies() {
		if c.Name == SessionCookie {
			t.Fatal("a session was minted with FS_DEV_LOGIN off")
		}
	}
}

// A peer that is not on loopback is refused even with the flag on, so a listener bound too
// wide, or a proxy in front of it, still cannot hand out sessions.
func TestDevLoginRefusesAPeerOffLoopback(t *testing.T) {
	h := newHarness(t)
	svc := *h.svc
	svc.DevLogin = true
	if _, err := svc.Store.UpsertBnetUser(context.Background(), "dev-sub-3", "Seeded#0003"); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/v1/dev/login", strings.NewReader(`{"battletag":"Seeded#0003"}`))
	req.RemoteAddr = "203.0.113.9:51234"
	rec := httptest.NewRecorder()
	svc.devLogin(rec, req)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want 403", rec.Code)
	}
	if len(rec.Result().Cookies()) != 0 {
		t.Fatal("a session cookie was set for an off-loopback peer")
	}
}
