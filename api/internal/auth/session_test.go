package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// newCSRFRequest builds a state-changing request carrying sess as the
// fs_session cookie, an fs_csrf cookie set to csrfCookie (omitted when
// empty), and the X-CSRF-Token header set to header when setHeader is
// true (including an empty value, to distinguish a missing header from
// an empty one).
func newCSRFRequest(sessID, csrfCookie string, setHeader bool, header string) *http.Request {
	r := httptest.NewRequest(http.MethodPost, "/v1/me", nil)
	r.AddCookie(&http.Cookie{Name: SessionCookie, Value: sessID})
	if csrfCookie != "" {
		r.AddCookie(&http.Cookie{Name: CSRFCookie, Value: csrfCookie})
	}
	if setHeader {
		r.Header.Set(CSRFHeader, header)
	}
	return r
}

func TestCSRFIsEnforcedOnStateChangingSessionRequests(t *testing.T) {
	store := testStore(t)
	u, err := store.UpsertEmailUser(context.Background(), "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	sess, err := store.CreateSession(context.Background(), u.ID, "email", SessionTTL)
	if err != nil {
		t.Fatal(err)
	}
	a := &Authenticator{Store: store}
	h := a.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	for _, tc := range []struct {
		name       string
		csrfCookie string
		setHeader  bool
		header     string
		want       int
	}{
		{"matching token succeeds", "csrf-token", true, "csrf-token", http.StatusOK},
		{"missing header is refused", "csrf-token", false, "", http.StatusForbidden},
		{"empty header is refused", "csrf-token", true, "", http.StatusForbidden},
		{"mismatched header is refused", "csrf-token", true, "wrong-token", http.StatusForbidden},
		{"missing cookie is refused", "", true, "csrf-token", http.StatusForbidden},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			h.ServeHTTP(w, newCSRFRequest(sess.ID, tc.csrfCookie, tc.setHeader, tc.header))
			if w.Code != tc.want {
				t.Fatalf("status = %d, want %d", w.Code, tc.want)
			}
		})
	}
}

func TestDeviceTokenRequestsAreExemptFromCSRF(t *testing.T) {
	store := testStore(t)
	u, err := store.UpsertEmailUser(context.Background(), "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	token, hash := NewDeviceToken()
	d, err := store.CreateDevice(context.Background(), u.ID, "Raid PC", "windows/amd64", hash)
	if err != nil {
		t.Fatal(err)
	}
	a := &Authenticator{Store: store}
	var reached Actor
	h := a.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reached = ActorFrom(r.Context())
		w.WriteHeader(http.StatusOK)
	}))

	// A state-changing request with a device token and no CSRF cookie or
	// header at all must still succeed: a companion is not a browser and
	// cannot be made to submit a form.
	r := httptest.NewRequest(http.MethodPost, "/v1/fights", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200 (device tokens are exempt from CSRF)", w.Code)
	}
	if !reached.IsDevice() || reached.DeviceID != d.ID || reached.UserID != u.ID {
		t.Fatalf("actor = %+v, want the device %q owned by user %d", reached, d.ID, u.ID)
	}
}

func TestDeviceTokenThatDoesNotResolveIsRejected(t *testing.T) {
	store := testStore(t)
	a := &Authenticator{Store: store}
	h := a.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("the handler should not be reached for an unknown device token")
	}))
	r := httptest.NewRequest(http.MethodPost, "/v1/fights", nil)
	r.Header.Set("Authorization", "Bearer "+DeviceTokenPrefix+Base32ID(DeviceTokenChars))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", w.Code)
	}
}

func TestStartSessionSetsBothCookiesWithTheContractAttributes(t *testing.T) {
	store := testStore(t)
	u, err := store.UpsertEmailUser(context.Background(), "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name       string
		domain     string
		wantDomain string
	}{
		{"empty CookieDomain is a host-only cookie", "", ""},
		// net/http's Set-Cookie writer sanitizes away a leading "." on the
		// Domain attribute (RFC 6265 treats a leading dot as equivalent to
		// none, so modern cookie writers drop it): CookieDomain
		// ".foreversixty.gg" is what SESSION_COOKIE_DOMAIN is documented
		// as, and it still comes out matching the whole *.foreversixty.gg
		// tree — just without the redundant dot on the wire.
		{"a set CookieDomain comes out as Domain", ".foreversixty.gg", "foreversixty.gg"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a := &Authenticator{Store: store, CookieDomain: tc.domain, Secure: true}
			w := httptest.NewRecorder()
			if _, err := a.StartSession(context.Background(), w, u.ID, "email"); err != nil {
				t.Fatal(err)
			}

			var session, csrf *http.Cookie
			for _, c := range w.Result().Cookies() {
				switch c.Name {
				case SessionCookie:
					session = c
				case CSRFCookie:
					csrf = c
				}
			}
			if session == nil || csrf == nil {
				t.Fatalf("cookies = %v, want both %s and %s", w.Header().Values("Set-Cookie"), SessionCookie, CSRFCookie)
			}
			if !session.HttpOnly {
				t.Error("fs_session should be HttpOnly")
			}
			if csrf.HttpOnly {
				t.Error("fs_csrf should be readable by the page, not HttpOnly")
			}
			for _, c := range []*http.Cookie{session, csrf} {
				if c.Path != "/" {
					t.Errorf("%s path = %q, want /", c.Name, c.Path)
				}
				if c.SameSite != http.SameSiteLaxMode {
					t.Errorf("%s SameSite = %v, want Lax", c.Name, c.SameSite)
				}
				if !c.Secure {
					t.Errorf("%s should be Secure when Authenticator.Secure is true", c.Name)
				}
				if c.Domain != tc.wantDomain {
					t.Errorf("%s domain = %q, want %q", c.Name, c.Domain, tc.wantDomain)
				}
			}
			wantExpiry := time.Now().Add(SessionTTL)
			if d := session.Expires.Sub(wantExpiry); d < -time.Minute || d > time.Minute {
				t.Errorf("fs_session expiry = %v, want ~30 days from now (%v)", session.Expires, wantExpiry)
			}
		})
	}
}

// Losing fs_csrf while fs_session survives used to be a lockout: every
// state-changing request 403s, and DELETE /v1/sessions is itself state
// changing, so even signing out on the account page failed. The
// middleware now hands the token back, so the retry works.
func TestALostCSRFCookieIsReissuedWhileTheSessionStands(t *testing.T) {
	store := testStore(t)
	u, err := store.UpsertEmailUser(context.Background(), "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	sess, err := store.CreateSession(context.Background(), u.ID, "email", SessionTTL)
	if err != nil {
		t.Fatal(err)
	}
	a := &Authenticator{Store: store}
	h := a.Middleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	w := httptest.NewRecorder()
	h.ServeHTTP(w, newCSRFRequest(sess.ID, "", true, "anything"))
	if w.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want the request itself still refused", w.Code)
	}
	reissued := setCookie(w.Result(), CSRFCookie)
	if reissued == nil || reissued.Value == "" {
		t.Fatalf("no fs_csrf was re-issued: %v", w.Result().Cookies())
	}

	// The browser now has the token, so the retry is accepted.
	w = httptest.NewRecorder()
	h.ServeHTTP(w, newCSRFRequest(sess.ID, reissued.Value, true, reissued.Value))
	if w.Code != http.StatusOK {
		t.Fatalf("the retry = %d, want 200", w.Code)
	}

	// A GET is not state changing and never had a token to lose, but
	// it is the cheapest place to get one back.
	w = httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/me", nil)
	r.AddCookie(&http.Cookie{Name: SessionCookie, Value: sess.ID})
	h.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("a read = %d, want 200", w.Code)
	}
	if got := setCookie(w.Result(), CSRFCookie); got == nil || got.Value == "" {
		t.Fatalf("a read with a session and no token must re-issue one: %v", w.Result().Cookies())
	}
	if got := w.Header().Get("Cache-Control"); got != "no-store" {
		t.Fatalf("a response that re-issues fs_csrf must be no-store, got %q", got)
	}

	// A public read is shared through caches, so it never carries the
	// re-issue: one viewer's cookie must not be served to the next.
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/v1/rankings?encounter=1", nil)
	r.AddCookie(&http.Cookie{Name: SessionCookie, Value: sess.ID})
	h.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("a public read = %d, want 200", w.Code)
	}
	if got := setCookie(w.Result(), CSRFCookie); got != nil {
		t.Fatalf("a public read re-issued fs_csrf into a cacheable response: %+v", got)
	}
	if got := w.Header().Get("Cache-Control"); got != "" {
		t.Fatalf("a public read had its cache policy overridden: %q", got)
	}
}

// A request that already carries fs_csrf is left alone: re-issuing on
// every request would race a second tab into a 403.
func TestACSRFCookieThatIsPresentIsNotReissued(t *testing.T) {
	store := testStore(t)
	u, err := store.UpsertEmailUser(context.Background(), "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	sess, err := store.CreateSession(context.Background(), u.ID, "email", SessionTTL)
	if err != nil {
		t.Fatal(err)
	}
	a := &Authenticator{Store: store}
	h := a.Middleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, newCSRFRequest(sess.ID, "csrf-token", true, "csrf-token"))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if got := setCookie(w.Result(), CSRFCookie); got != nil {
		t.Fatalf("fs_csrf was re-issued over a good one: %+v", got)
	}
}

// setCookie reads one Set-Cookie off a response, or nil when it is not
// there.
func setCookie(res *http.Response, name string) *http.Cookie {
	for _, c := range res.Cookies() {
		if c.Name == name {
			return c
		}
	}
	return nil
}
