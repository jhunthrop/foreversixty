package auth

import (
	"html"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

// M1: the start answers a confirm page that needs a click. It never redirects on its own, so
// a link alone cannot finish a sign-in while Battle.net consents silently.
func TestAppStartIsAConfirmPageWithNoAutoRedirect(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	res := h.do(t, http.MethodGet, appStartQuery("ios", challengeOf(testVerifier), "S256", testAppState), "")
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK || res.Header.Get("Location") != "" {
		t.Fatalf("start = %d to %q, want 200 with no redirect", res.StatusCode, res.Header.Get("Location"))
	}
	b, _ := io.ReadAll(res.Body)
	body := string(b)
	for _, want := range []string{appConfirmCopy.Heading, appConfirmCopy.Button, appConfirmCopy.NotYou} {
		if !strings.Contains(html.UnescapeString(body), want) {
			t.Errorf("confirm page is missing %q", want)
		}
	}
	for _, want := range []string{`href="/v1/auth/app/continue?`} {
		if !strings.Contains(body, want) {
			t.Errorf("confirm page is missing %q", want)
		}
	}
	for _, banned := range []string{"<script", "http-equiv", "<form", "<iframe", "src="} {
		if strings.Contains(strings.ToLower(body), banned) {
			t.Errorf("confirm page carries %q; it must need a click and load nothing", banned)
		}
	}
	for header, want := range map[string]string{
		"Content-Type":    "text/html; charset=utf-8",
		"Cache-Control":   "no-store",
		"Referrer-Policy": "no-referrer",
		"X-Frame-Options": "DENY",
	} {
		if got := res.Header.Get(header); got != want {
			t.Errorf("%s = %q, want %q", header, got, want)
		}
	}
	if csp := res.Header.Get("Content-Security-Policy"); !strings.Contains(csp, "default-src 'none'") ||
		!strings.Contains(csp, "frame-ancestors 'none'") {
		t.Errorf("CSP = %q, want default-src and frame-ancestors 'none'", csp)
	}
}

// The continue link works only in the browser whose start set the state cookie: an attacker
// who opens a start and sends someone its continue link gets nothing.
func TestAppContinueNeedsTheBrowserThatStarted(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	cont := continueLink(t, h.do(t, http.MethodGet, appStartQuery("ios", challengeOf(testVerifier), "S256", testAppState), ""))

	victim := &http.Client{CheckRedirect: h.client.CheckRedirect}
	res, err := victim.Get(h.server.URL + cont)
	if err != nil {
		t.Fatal(err)
	}
	wantStatus(t, res, http.StatusBadRequest, "continue from a browser that never started")

	wantStatus(t, h.do(t, http.MethodGet, "/v1/auth/app/continue?state=not-the-state", ""),
		http.StatusBadRequest, "continue with a forged state")
	wantStatus(t, h.do(t, http.MethodGet, "/v1/auth/app/continue", ""), http.StatusBadRequest, "continue with no state")
}

// A web sign-in's state cookie cannot be continued as an app flow.
func TestAppContinueRefusesAWebSignInState(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/start", "")
	res.Body.Close()
	target, _ := url.Parse(res.Header.Get("Location"))
	wantStatus(t, h.do(t, http.MethodGet, "/v1/auth/app/continue?state="+url.QueryEscape(target.Query().Get("state")), ""),
		http.StatusBadRequest, "continue on a web sign-in's state")
}

// The redirect is ours alone: a redirect_uri the client sends is ignored, Battle.net is sent
// the configured callback, and the code goes to /app/auth.
func TestAppSignInIgnoresAClientRedirectURI(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	q := appStartQuery("android", challengeOf(testVerifier), "S256", testAppState) +
		"&redirect_uri=" + url.QueryEscape("https://evil.example/app/auth/callback")
	cont := continueLink(t, h.do(t, http.MethodGet, q, ""))
	res := h.do(t, http.MethodGet, cont, "")
	res.Body.Close()
	target, _ := url.Parse(res.Header.Get("Location"))
	if got := target.Query().Get("redirect_uri"); got != h.svc.BNet.Config.RedirectURL {
		t.Fatalf("Battle.net redirect_uri = %q, want the configured %q", got, h.svc.BNet.Config.RedirectURL)
	}
	h.appCallback(t, target.Query().Get("state")) // asserts https://foreversixty.gg/app/auth
}

// A Battle.net callback replayed against an app flow that already has its code mints nothing.
func TestAppCallbackReplayMintsNoSecondCode(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	state := h.appStart(t, "ios")
	h.appCallback(t, state)
	u, _ := url.Parse(h.server.URL)
	h.client.Jar.SetCookies(u, []*http.Cookie{{Name: bnetStateCookie, Value: encodeState(state, appFlowNext), Path: "/"}})
	wantStatus(t, h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state="+state, ""),
		http.StatusBadRequest, "replayed callback")
	var n int
	if err := h.svc.Store.Pool.QueryRow(t.Context(), `select count(*) from app_auth_flows where code_hash is not null`).Scan(&n); err != nil || n != 1 {
		t.Fatalf("codes = %d, %v; want the one", n, err)
	}
}

// M2: an app token lapses after its idle limit or its maximum age, answering session_expired.
func TestAppTokenLapsesWhenIdleOrOld(t *testing.T) {
	for name, age := range map[string]string{
		"idle":    `update devices set last_seen_at = now() - interval '31 days'`,
		"too old": `update devices set created_at = now() - interval '181 days', last_seen_at = now()`,
	} {
		t.Run(name, func(t *testing.T) {
			h := newHarness(t)
			h.withAppSignIn(t)
			got := h.signInApp(t, "ios")
			if _, err := h.svc.Store.Pool.Exec(t.Context(), age); err != nil {
				t.Fatal(err)
			}
			wantErrorCode(t, bearer(t, h, http.MethodGet, "/v1/me", got.Token), http.StatusUnauthorized, "session_expired")
			// A lapsed app is revoked, so it leaves the account's device list.
			var revoked bool
			if err := h.svc.Store.Pool.QueryRow(t.Context(),
				`select revoked_at is not null from devices where id = $1`, got.Device.ID).Scan(&revoked); err != nil || !revoked {
				t.Fatalf("revoked = %v, %v", revoked, err)
			}
			wantErrorCode(t, bearer(t, h, http.MethodGet, "/v1/me", got.Token), http.StatusUnauthorized, "unauthorized")
		})
	}
}

func TestAppTokenLifetimesComeFromTheStore(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	got := h.signInApp(t, "android")
	h.svc.Store.AppTokenIdle = time.Hour
	if _, err := h.svc.Store.Pool.Exec(t.Context(), `update devices set last_seen_at = now() - interval '2 hours'`); err != nil {
		t.Fatal(err)
	}
	wantErrorCode(t, bearer(t, h, http.MethodGet, "/v1/me", got.Token), http.StatusUnauthorized, "session_expired")
}

func TestARecentlyUsedAppTokenKeepsWorking(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	got := h.signInApp(t, "ios")
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`update devices set created_at = now() - interval '179 days', last_seen_at = now() - interval '29 days'`); err != nil {
		t.Fatal(err)
	}
	wantStatus(t, bearer(t, h, http.MethodGet, "/v1/me", got.Token), http.StatusOK, "/v1/me inside both limits")
}

// Companion tokens never lapse: a raid PC can sit unused for a season.
func TestACompanionTokenDoesNotLapse(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "raider@example.com")
	token, hash := NewDeviceToken()
	if _, err := h.svc.Store.CreateDevice(t.Context(), uid, "Raid PC", "windows", hash); err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`update devices set created_at = now() - interval '400 days', last_seen_at = now() - interval '200 days'`); err != nil {
		t.Fatal(err)
	}
	wantStatus(t, bearer(t, h, http.MethodDelete, "/v1/devices/self", token), http.StatusNoContent, "an old companion's own route")
}

// M2: billing writes and moderator powers stay with a browser session.
func TestRequireBrowserSessionRefusesEveryDevice(t *testing.T) {
	h := RequireBrowserSession(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) })
	for name, c := range map[string]struct {
		actor  Actor
		status int
		code   string
	}{
		"browser":   {Actor{UserID: 1, Method: "session"}, http.StatusNoContent, ""},
		"app":       {Actor{UserID: 1, Method: "device", DeviceID: "d", DeviceKind: DeviceKindApp}, http.StatusForbidden, "browser_session_required"},
		"companion": {Actor{UserID: 1, Method: "device", DeviceID: "d", DeviceKind: DeviceKindCompanion}, http.StatusUnauthorized, "unauthorized"},
		"anonymous": {Actor{}, http.StatusUnauthorized, "unauthorized"},
	} {
		rec := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/v1/billing/checkout", nil)
		h(rec, r.WithContext(WithActor(r.Context(), c.actor)))
		if rec.Code != c.status || (c.code != "" && !strings.Contains(rec.Body.String(), `"code":"`+c.code+`"`)) {
			t.Errorf("%s: %d %s, want %d %s", name, rec.Code, rec.Body.String(), c.status, c.code)
		}
	}
}

func TestAnAppTokenCarriesNoModeratorPowers(t *testing.T) {
	app := Actor{UserID: 1, Role: "admin", Method: "device", DeviceID: "d", DeviceKind: DeviceKindApp}
	if app.IsModerator() {
		t.Fatal("an admin's app token must not act as a moderator")
	}
	if !(Actor{UserID: 1, Role: "moderator", Method: "session"}).IsModerator() {
		t.Fatal("a moderator's browser session keeps its powers")
	}
}

// LOW: the site's generic sign-outs revoke an app's device rather than answering success
// and doing nothing.
func TestGenericSignOutsRevokeAnAppDevice(t *testing.T) {
	for _, route := range []struct{ method, path string }{
		{http.MethodDelete, "/v1/sessions"},
		{http.MethodPost, "/v1/auth/logout"},
	} {
		t.Run(route.path, func(t *testing.T) {
			h := newHarness(t)
			h.withAppSignIn(t)
			got := h.signInApp(t, "ios")
			res := bearer(t, h, route.method, route.path, got.Token)
			res.Body.Close()
			if res.StatusCode >= 300 {
				t.Fatalf("%s %s = %d", route.method, route.path, res.StatusCode)
			}
			wantStatus(t, bearer(t, h, http.MethodGet, "/v1/me", got.Token), http.StatusUnauthorized, "/v1/me after the sign-out")
		})
	}
}
