package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
	"golang.org/x/oauth2"
)

// harness is a mounted auth service over the test database, with a fake
// mailer and a client that keeps cookies, so a test can follow a whole
// sign-in the way a browser would.
type harness struct {
	svc    *Service
	server *httptest.Server
	client *http.Client
	mailer *mail.Fake
}

func newHarness(t *testing.T) *harness {
	t.Helper()
	store := testStore(t)
	mailer := &mail.Fake{}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	a := &Authenticator{Store: store, Log: quiet}
	svc := &Service{
		Store: store, Auth: a, Mailer: mailer, Log: quiet,
		Entitlements:  &entitlements.Store{Pool: store.Pool},
		PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg",
	}
	mux := http.NewServeMux()
	Mount(mux, svc, 0)
	srv := httptest.NewServer(a.Middleware(mux))
	t.Cleanup(srv.Close)

	jar, err := newJar()
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{
		Jar: jar,
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return &harness{svc: svc, server: srv, client: client, mailer: mailer}
}

// seedUser creates an account directly through the store, skipping the
// email magic-link flow, for tests that only care what a signed-in
// account sees.
func (h *harness) seedUser(t *testing.T, email string) int64 {
	t.Helper()
	u, err := h.svc.Store.UpsertEmailUser(t.Context(), email)
	if err != nil {
		t.Fatal(err)
	}
	return u.ID
}

// seedGuild inserts a bare guild row and returns its id.
func (h *harness) seedGuild(t *testing.T, name string) int64 {
	t.Helper()
	var id int64
	if err := h.svc.Store.Pool.QueryRow(t.Context(),
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', $1) returning id`, name).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// seedGuildMember links userID into guildID at rank, verified or not.
func (h *harness) seedGuildMember(t *testing.T, guildID, userID int64, rank string, verified bool) {
	t.Helper()
	var verifiedAt any
	if verified {
		verifiedAt = time.Now()
	}
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`insert into guild_members (guild_id, user_id, rank, verified_at, refreshed_at) values ($1, $2, $3, $4, now())`,
		guildID, userID, rank, verifiedAt); err != nil {
		t.Fatal(err)
	}
}

// sessionRequest signs in as uid directly through the store, skipping
// the email flow, and sends method/path with that session's cookie.
func (h *harness) sessionRequest(t *testing.T, method, path string, uid int64) *http.Response {
	t.Helper()
	sess, err := h.svc.Store.CreateSession(t.Context(), uid, "email", SessionTTL)
	if err != nil {
		t.Fatal(err)
	}
	r, err := http.NewRequest(method, h.server.URL+path, nil)
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(&http.Cookie{Name: SessionCookie, Value: sess.ID})
	res, err := h.client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

// do sends a request with the CSRF header filled in from the cookie jar,
// which is what the site's fetch wrapper does.
func (h *harness) do(t *testing.T, method, path, body string) *http.Response {
	t.Helper()
	var r *http.Request
	var err error
	if body == "" {
		r, err = http.NewRequest(method, h.server.URL+path, nil)
	} else {
		r, err = http.NewRequest(method, h.server.URL+path, strings.NewReader(body))
		r.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		t.Fatal(err)
	}
	u, _ := url.Parse(h.server.URL)
	for _, c := range h.client.Jar.Cookies(u) {
		if c.Name == CSRFCookie {
			r.Header.Set(CSRFHeader, c.Value)
		}
	}
	res, err := h.client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func (h *harness) decode(t *testing.T, res *http.Response, into any) {
	t.Helper()
	defer res.Body.Close()
	var env struct {
		OK   bool            `json:"ok"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	if !env.OK {
		t.Fatalf("envelope reports failure for %s", res.Request.URL)
	}
	if into != nil {
		if err := json.Unmarshal(env.Data, into); err != nil {
			t.Fatal(err)
		}
	}
}

// signIn takes the harness through the whole email magic link.
func (h *harness) signIn(t *testing.T, address string) {
	t.Helper()
	res := h.do(t, http.MethodPost, "/v1/auth/email", `{"email":"`+address+`"}`)
	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("POST /v1/auth/email = %d", res.StatusCode)
	}
	res.Body.Close()
	if len(h.mailer.Sent) == 0 {
		t.Fatal("no link was mailed")
	}
	link := h.mailer.Sent[len(h.mailer.Sent)-1].Text
	i := strings.Index(link, "token=")
	if i < 0 {
		t.Fatalf("the mail carries no token: %q", link)
	}
	token := strings.TrimSpace(link[i+len("token="):])
	token, _, _ = strings.Cut(token, "\n")
	res = h.do(t, http.MethodGet, "/v1/auth/email/callback?token="+token, "")
	defer res.Body.Close()
	if res.StatusCode != http.StatusFound {
		t.Fatalf("callback = %d, want a redirect", res.StatusCode)
	}
	if got := res.Header.Get("Location"); got != "https://foreversixty.gg/logs" {
		t.Fatalf("redirected to %q", got)
	}
}

func TestEmailMagicLinkSignsInAndMeAnswers(t *testing.T) {
	h := newHarness(t)
	res := h.do(t, http.MethodGet, "/v1/me", "")
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("/v1/me before signing in = %d, want 401", res.StatusCode)
	}
	res.Body.Close()

	h.signIn(t, "raider@example.com")

	res = h.do(t, http.MethodGet, "/v1/me", "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("/v1/me = %d", res.StatusCode)
	}
	var me Me
	h.decode(t, res, &me)
	if me.User.Email == nil || *me.User.Email != "raider@example.com" || me.User.Role != "user" {
		t.Fatalf("me = %+v", me)
	}
	if me.Characters == nil || me.Guilds == nil {
		t.Fatalf("me must carry empty lists rather than null: %+v", me)
	}
}

func TestSecondClickOnAMagicLinkFails(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")
	link := h.mailer.Sent[0].Text
	token := strings.TrimSpace(strings.SplitN(link[strings.Index(link, "token=")+6:], "\n", 2)[0])
	res := h.do(t, http.MethodGet, "/v1/auth/email/callback?token="+token, "")
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("second click = %d, want 400", res.StatusCode)
	}
}

func TestMagicLinksAreCappedAtFivePerHour(t *testing.T) {
	h := newHarness(t)
	for i := range LoginsPerHour {
		res := h.do(t, http.MethodPost, "/v1/auth/email", `{"email":"raider@example.com"}`)
		res.Body.Close()
		if res.StatusCode != http.StatusAccepted {
			t.Fatalf("request %d = %d", i, res.StatusCode)
		}
	}
	res := h.do(t, http.MethodPost, "/v1/auth/email", `{"email":"raider@example.com"}`)
	defer res.Body.Close()
	if res.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("sixth request = %d, want 429", res.StatusCode)
	}
}

// TestEmailIsRateLimitedPerIPOnTopOfThePerAddressCap reaches the per-IP
// budget (emailsPerHour) by spreading requests across distinct addresses,
// each well under the five-per-address cap, so it is the IP-wide limiter
// - not the address one - that answers 429.
func TestEmailIsRateLimitedPerIPOnTopOfThePerAddressCap(t *testing.T) {
	h := newHarness(t)
	for i := range emailsPerHour {
		body := fmt.Sprintf(`{"email":"raider%d@example.com"}`, i)
		res := h.do(t, http.MethodPost, "/v1/auth/email", body)
		res.Body.Close()
		if res.StatusCode != http.StatusAccepted {
			t.Fatalf("request %d = %d", i, res.StatusCode)
		}
	}
	res := h.do(t, http.MethodPost, "/v1/auth/email", `{"email":"one-more@example.com"}`)
	defer res.Body.Close()
	if res.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("request %d = %d, want 429 from the per-IP cap", emailsPerHour+1, res.StatusCode)
	}
}

func TestAnInvalidAddressIsRejectedWithAField(t *testing.T) {
	h := newHarness(t)
	res := h.do(t, http.MethodPost, "/v1/auth/email", `{"email":"not-an-address"}`)
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", res.StatusCode)
	}
	var env struct {
		Error struct {
			Fields map[string]string `json:"fields"`
		} `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	if env.Error.Fields["email"] == "" {
		t.Fatalf("the 400 should name the email field: %+v", env)
	}
}

func TestStateChangingRequestsNeedTheCSRFHeader(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")

	// The same request without the header, which is what a cross-site
	// form post looks like.
	r, err := http.NewRequest(http.MethodPost, h.server.URL+"/v1/devices/pair", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := h.client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("status = %d, want 403 without the CSRF header", res.StatusCode)
	}
}

func TestPairClaimListAndRevokeADevice(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")

	res := h.do(t, http.MethodPost, "/v1/devices/pair", "")
	var pairing struct {
		Code string `json:"code"`
	}
	h.decode(t, res, &pairing)
	if pairing.Code == "" {
		t.Fatal("no pairing code")
	}

	// The companion has no cookies at all.
	claimed, err := http.Post(h.server.URL+"/v1/devices/claim", "application/json",
		strings.NewReader(`{"code":"`+pairing.Code+`","name":"Raid PC","platform":"windows/amd64"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer claimed.Body.Close()
	if claimed.StatusCode != http.StatusCreated {
		t.Fatalf("claim = %d", claimed.StatusCode)
	}
	var device struct {
		DeviceID string `json:"device_id"`
		Token    string `json:"token"`
	}
	var env struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.NewDecoder(claimed.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(env.Data, &device); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(device.Token, DeviceTokenPrefix) {
		t.Fatalf("token = %q", device.Token)
	}

	// That token authenticates as the owner, with no CSRF header.
	r, _ := http.NewRequest(http.MethodGet, h.server.URL+"/v1/me", nil)
	r.Header.Set("Authorization", "Bearer "+device.Token)
	asDevice, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer asDevice.Body.Close()
	// A device token uploads; it does not read the account behind it.
	if asDevice.StatusCode != http.StatusUnauthorized {
		t.Fatalf("/v1/me as the device = %d, want 401", asDevice.StatusCode)
	}

	// But it cannot pair another device: that needs a browser session.
	r, _ = http.NewRequest(http.MethodPost, h.server.URL+"/v1/devices/pair", nil)
	r.Header.Set("Authorization", "Bearer "+device.Token)
	res2, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer res2.Body.Close()
	if res2.StatusCode != http.StatusUnauthorized {
		t.Fatalf("pairing as a device = %d, want 401", res2.StatusCode)
	}

	res = h.do(t, http.MethodGet, "/v1/devices", "")
	var list []Device
	h.decode(t, res, &list)
	if len(list) != 1 || list[0].ID != device.DeviceID {
		t.Fatalf("devices = %+v", list)
	}

	res = h.do(t, http.MethodDelete, "/v1/devices/"+device.DeviceID, "")
	res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("revoke = %d", res.StatusCode)
	}
	r, _ = http.NewRequest(http.MethodGet, h.server.URL+"/v1/me", nil)
	r.Header.Set("Authorization", "Bearer "+device.Token)
	after, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer after.Body.Close()
	if after.StatusCode != http.StatusUnauthorized {
		t.Fatalf("a revoked token = %d, want 401", after.StatusCode)
	}
}

func TestAnUnknownPairingCodeIs404(t *testing.T) {
	h := newHarness(t)
	res, err := http.Post(h.server.URL+"/v1/devices/claim", "application/json",
		strings.NewReader(`{"code":"aaaaaaaa"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", res.StatusCode)
	}
}

// TestClaimIsRateLimitedPerIP covers the one control standing between a
// guesser and the live pairing-code pool: pairing_codes.code is eight
// base32 characters (40 bits) in plaintext, so an unthrottled claim route
// would make brute-forcing it online feasible. Each attempt here uses an
// unknown code and answers 404 on its own, but the budget is still spent
// - the limiter runs before the handler ever looks the code up.
func TestClaimIsRateLimitedPerIP(t *testing.T) {
	h := newHarness(t)
	for i := range claimsPerHour {
		res, err := http.Post(h.server.URL+"/v1/devices/claim", "application/json",
			strings.NewReader(`{"code":"aaaaaaaa"}`))
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()
		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("request %d = %d, want 404 from an unknown code, not the limiter", i, res.StatusCode)
		}
	}
	res, err := http.Post(h.server.URL+"/v1/devices/claim", "application/json",
		strings.NewReader(`{"code":"aaaaaaaa"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("request %d = %d, want 429", claimsPerHour+1, res.StatusCode)
	}
}

func TestLogoutEndsTheSession(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")
	res := h.do(t, http.MethodPost, "/v1/auth/logout", "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("logout = %d", res.StatusCode)
	}
	res = h.do(t, http.MethodGet, "/v1/me", "")
	defer res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("/v1/me after logout = %d, want 401", res.StatusCode)
	}
}

// fakeBattleNet is an OAuth provider: the token endpoint and the
// userinfo endpoint, enough for the callback to complete.
func fakeBattleNet(t *testing.T, sub, battletag string) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /token", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Error(err)
		}
		if r.Form.Get("code") != "the-code" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"at","token_type":"bearer","expires_in":3600}`))
	})
	mux.HandleFunc("GET /oauth/userinfo", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer at" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"sub":"` + sub + `","battletag":"` + battletag + `"}`))
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

func (h *harness) withBattleNet(t *testing.T, sub, battletag string) {
	t.Helper()
	bnet := fakeBattleNet(t, sub, battletag)
	h.svc.BNet = &BattleNet{
		Config: &oauth2.Config{
			ClientID: "id", ClientSecret: "secret",
			RedirectURL: h.server.URL + "/v1/auth/battlenet/callback",
			Scopes:      []string{"openid", "wow.profile"},
			Endpoint:    oauth2.Endpoint{AuthURL: bnet.URL + "/authorize", TokenURL: bnet.URL + "/token"},
		},
		UserInfoURL: bnet.URL + "/oauth/userinfo",
		HTTP:        bnet.Client(),
	}
	// Remount so the Battle.net routes exist now that it is configured.
	mux := http.NewServeMux()
	Mount(mux, h.svc, 0)
	h.server.Config.Handler = h.svc.Auth.Middleware(mux)
}

// TestBattleNetRoutesAreNotMountedWhenUnconfigured covers honest
// degradation: newHarness leaves Service.BNet nil (no withBattleNet
// call), so an unconfigured deployment must answer 404 here rather than
// nil-dereferencing s.BNet.AuthURL, and the email magic link stays the
// only way in.
func TestBattleNetRoutesAreNotMountedWhenUnconfigured(t *testing.T) {
	h := newHarness(t)
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/start", "")
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("status = %d, want 404 when Battle.net is not configured", res.StatusCode)
	}
}

func TestBattleNetSignInCreatesTheAccountAndSession(t *testing.T) {
	h := newHarness(t)
	h.withBattleNet(t, "12345", "Baelgrim#1234")

	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/start?next=/account", "")
	res.Body.Close()
	if res.StatusCode != http.StatusFound {
		t.Fatalf("start = %d, want a redirect", res.StatusCode)
	}
	target, err := url.Parse(res.Header.Get("Location"))
	if err != nil {
		t.Fatal(err)
	}
	if got := target.Query().Get("scope"); got != "openid wow.profile" {
		t.Fatalf("scope = %q", got)
	}
	state := target.Query().Get("state")
	if state == "" {
		t.Fatal("no state was sent to Battle.net")
	}

	res = h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state="+state, "")
	res.Body.Close()
	if res.StatusCode != http.StatusFound {
		t.Fatalf("callback = %d", res.StatusCode)
	}
	if got := res.Header.Get("Location"); got != "https://foreversixty.gg/account" {
		t.Fatalf("redirected to %q, want the next page", got)
	}

	res = h.do(t, http.MethodGet, "/v1/me", "")
	var me Me
	h.decode(t, res, &me)
	if me.User.Battletag == nil || *me.User.Battletag != "Baelgrim#1234" {
		t.Fatalf("me = %+v", me)
	}
}

func TestBattleNetCallbackRejectsAForgedState(t *testing.T) {
	h := newHarness(t)
	h.withBattleNet(t, "12345", "Baelgrim#1234")
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/start", "")
	res.Body.Close()
	res = h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state=not-the-state", "")
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", res.StatusCode)
	}
}

func TestStartRefusesToRedirectOffSite(t *testing.T) {
	h := newHarness(t)
	h.withBattleNet(t, "12345", "Baelgrim#1234")
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/start?next=https://evil.example/steal", "")
	res.Body.Close()
	target, _ := url.Parse(res.Header.Get("Location"))
	state := target.Query().Get("state")
	res = h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state="+state, "")
	defer res.Body.Close()
	if got := res.Header.Get("Location"); got != "https://foreversixty.gg/logs" {
		t.Fatalf("redirected to %q, want the safe default", got)
	}
}

// newJar is a cookie jar for the test client. net/http/cookiejar needs a
// public suffix list to be useful across hosts; every request here is to
// one host, so the default nil list is right.
func newJar() (http.CookieJar, error) { return cookiejar.New(nil) }

func TestTheAccountPageSignsOutAndSetsThePseudonymFlag(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")

	res := h.do(t, http.MethodPatch, "/v1/me", `{"anonymize":true}`)
	var me Me
	h.decode(t, res, &me)
	if !me.User.Anonymize {
		t.Fatalf("me = %+v, want the flag set", me.User)
	}
	if me.User.Battletag != nil {
		t.Fatalf("battletag = %v, want null for an email account", *me.User.Battletag)
	}

	res = h.do(t, http.MethodPatch, "/v1/me", `{}`)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("an empty patch = %d, want 400", res.StatusCode)
	}
	res = h.do(t, http.MethodPatch, "/v1/me", `{`)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a broken patch = %d, want 400", res.StatusCode)
	}

	res = h.do(t, http.MethodDelete, "/v1/sessions", "")
	res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("sign out = %d, want 204", res.StatusCode)
	}
	res = h.do(t, http.MethodGet, "/v1/me", "")
	res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("/v1/me after signing out = %d, want 401", res.StatusCode)
	}
}

func TestUnknownDeviceTokenIs401(t *testing.T) {
	h := newHarness(t)
	r, _ := http.NewRequest(http.MethodGet, h.server.URL+"/v1/me", nil)
	r.Header.Set("Authorization", "Bearer fsd_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", res.StatusCode)
	}
}

func TestANonDeviceBearerIsIgnoredRatherThanRejected(t *testing.T) {
	h := newHarness(t)
	r, _ := http.NewRequest(http.MethodGet, h.server.URL+"/v1/me", nil)
	r.Header.Set("Authorization", "Bearer some-other-scheme")
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	// Not a device token at all, so the request is simply anonymous and
	// /v1/me answers its own 401 rather than the middleware's.
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", res.StatusCode)
	}
}

func TestLogoutWithoutASessionIsHarmless(t *testing.T) {
	h := newHarness(t)
	res := h.do(t, http.MethodPost, "/v1/auth/logout", "")
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d", res.StatusCode)
	}
}

func TestMalformedBodiesAreRejected(t *testing.T) {
	h := newHarness(t)
	res := h.do(t, http.MethodPost, "/v1/auth/email", `{`)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("email = %d, want 400", res.StatusCode)
	}
	res = h.do(t, http.MethodPost, "/v1/devices/claim", `{`)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("claim = %d, want 400", res.StatusCode)
	}
	res = h.do(t, http.MethodGet, "/v1/auth/email/callback", "")
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("callback with no token = %d, want 400", res.StatusCode)
	}
	res = h.do(t, http.MethodGet, "/v1/auth/email/callback?token=nope", "")
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("callback with a bad token = %d, want 400", res.StatusCode)
	}
}

func TestBattleNetCallbackNeedsItsStateCookieAndCode(t *testing.T) {
	h := newHarness(t)
	h.withBattleNet(t, "1", "A#1")
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state=x", "")
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("no cookie = %d, want 400", res.StatusCode)
	}
	start := h.do(t, http.MethodGet, "/v1/auth/battlenet/start", "")
	start.Body.Close()
	target, _ := url.Parse(start.Header.Get("Location"))
	res = h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?state="+target.Query().Get("state"), "")
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("no code = %d, want 400", res.StatusCode)
	}
}

func TestBattleNetUpstreamFailureIs502(t *testing.T) {
	h := newHarness(t)
	h.withBattleNet(t, "1", "A#1")
	start := h.do(t, http.MethodGet, "/v1/auth/battlenet/start", "")
	start.Body.Close()
	target, _ := url.Parse(start.Header.Get("Location"))
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=wrong&state="+target.Query().Get("state"), "")
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadGateway {
		t.Fatalf("status = %d, want 502", res.StatusCode)
	}
}

// brokenService is an auth service whose database has gone away, so
// every handler takes its failure path.
func brokenService(t *testing.T) *Service {
	t.Helper()
	store := testStore(t)
	store.Pool.Close()
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	return &Service{
		Store: store, Auth: &Authenticator{Store: store, Log: quiet},
		Mailer: &mail.Fake{}, PublicBaseURL: "https://foreversixty.gg",
		APIBaseURL: "https://api.foreversixty.gg", Log: quiet,
	}
}

// asUser calls one handler with an actor already resolved, which is what
// the middleware would have done.
func asUser(h http.HandlerFunc, method, target string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, nil)
	r = r.WithContext(WithActor(r.Context(), Actor{UserID: 1, Role: "user", Method: "session"}))
	w := httptest.NewRecorder()
	h(w, r)
	return w
}

func TestHandlersAnswer500WhenTheDatabaseIsGone(t *testing.T) {
	s := brokenService(t)
	for name, w := range map[string]*httptest.ResponseRecorder{
		"me":      asUser(s.me, http.MethodGet, "/v1/me"),
		"pair":    asUser(s.pair, http.MethodPost, "/v1/devices/pair"),
		"devices": asUser(s.listDevices, http.MethodGet, "/v1/devices"),
		"revoke":  asUser(s.revokeDevice, http.MethodDelete, "/v1/devices/x"),
		"email":   emailPost(s, `{"email":"raider@example.com"}`),
		"claim":   claimPost(s, `{"code":"aaaaaaaa"}`),
	} {
		if w.Code != http.StatusInternalServerError {
			t.Errorf("%s = %d, want 500", name, w.Code)
		}
		if !strings.Contains(w.Body.String(), `"ok":false`) {
			t.Errorf("%s did not answer in the envelope: %s", name, w.Body.String())
		}
	}
}

func emailPost(s *Service, body string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(http.MethodPost, "/v1/auth/email", strings.NewReader(body))
	w := httptest.NewRecorder()
	s.emailStart(w, r)
	return w
}

func claimPost(s *Service, body string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(http.MethodPost, "/v1/devices/claim", strings.NewReader(body))
	w := httptest.NewRecorder()
	s.claim(w, r)
	return w
}

func TestLogoutAnswers500WhenTheSessionCannotBeDeleted(t *testing.T) {
	s := brokenService(t)
	r := httptest.NewRequest(http.MethodPost, "/v1/auth/logout", nil)
	r.AddCookie(&http.Cookie{Name: SessionCookie, Value: "whatever"})
	w := httptest.NewRecorder()
	s.logout(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", w.Code)
	}
}

func TestAMailerFailureIs500AndTheLinkIsNotClaimedSent(t *testing.T) {
	store := testStore(t)
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	s := &Service{
		Store: store, Auth: &Authenticator{Store: store, Log: quiet},
		Mailer: &mail.Fake{Err: errors.New("resend is down")},
		Log:    quiet, APIBaseURL: "https://api.foreversixty.gg",
	}
	w := emailPost(s, `{"email":"raider@example.com"}`)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", w.Code)
	}
}

func TestMeIncludesEntitlementsAndDefaultsToAllFalseWithNoRows(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "me-entitlements@example.com")
	res := h.sessionRequest(t, http.MethodGet, "/v1/me", uid)
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		t.Fatalf("status = %d", res.StatusCode)
	}
	var body struct {
		Entitlements struct {
			ServerSims    bool `json:"server_sims"`
			SupporterMark bool `json:"supporter_mark"`
			Billing       *struct {
				Status string `json:"status"`
			} `json:"billing"`
		} `json:"entitlements"`
	}
	h.decode(t, res, &body)
	if body.Entitlements.ServerSims || body.Entitlements.SupporterMark {
		t.Fatalf("a fresh account must show every feature false: %+v", body.Entitlements)
	}
	if body.Entitlements.Billing != nil {
		t.Fatalf("billing should be nil with no entitlement row, got %+v", body.Entitlements.Billing)
	}
}

func TestMeReflectsAnActivePersonalPremiumEntitlement(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "me-active-premium@example.com")
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`insert into entitlements (user_id, plan, source, status) values ($1, 'premium', 'grant', 'active')`,
		uid); err != nil {
		t.Fatal(err)
	}
	res := h.sessionRequest(t, http.MethodGet, "/v1/me", uid)
	var body struct {
		Entitlements struct {
			ServerSims bool `json:"server_sims"`
			Billing    *struct {
				Status string `json:"status"`
			} `json:"billing"`
		} `json:"entitlements"`
	}
	h.decode(t, res, &body)
	if !body.Entitlements.ServerSims {
		t.Fatal("server_sims should be true")
	}
	if body.Entitlements.Billing == nil || body.Entitlements.Billing.Status != "active" {
		t.Fatalf("billing = %+v", body.Entitlements.Billing)
	}
}

func TestMeShowsGuildPlanOnlyToAVerifiedOfficer(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "me-guild-plan")
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`insert into entitlements (guild_id, plan, source, status) values ($1, 'guild', 'grant', 'active')`,
		gid); err != nil {
		t.Fatal(err)
	}
	officer := h.seedUser(t, "me-guild-officer@example.com")
	h.seedGuildMember(t, gid, officer, "officer", true)
	member := h.seedUser(t, "me-guild-member@example.com")
	h.seedGuildMember(t, gid, member, "member", true)

	type meBody struct {
		Guilds []struct {
			ID   int64 `json:"id"`
			Plan *struct {
				Status string `json:"status"`
			} `json:"plan,omitempty"`
		} `json:"guilds"`
	}
	decode := func(uid int64) meBody {
		res := h.sessionRequest(t, http.MethodGet, "/v1/me", uid)
		var b meBody
		h.decode(t, res, &b)
		return b
	}
	oBody := decode(officer)
	if len(oBody.Guilds) != 1 || oBody.Guilds[0].Plan == nil || oBody.Guilds[0].Plan.Status != "active" {
		t.Fatalf("officer's guilds = %+v", oBody.Guilds)
	}
	mBody := decode(member)
	if len(mBody.Guilds) != 1 || mBody.Guilds[0].Plan != nil {
		t.Fatalf("non-officer must not see Plan: %+v", mBody.Guilds)
	}
}

// TestMeWorksWithNoEntitlementsStoreConfigured covers the nil-Entitlements
// success path: a harness that never sets Service.Entitlements (mirrors
// the doc comment on Service.Entitlements and reports.Service.Guilds's
// own nil-safety pattern) must still answer /v1/me with 200 and a
// zero-valued EntitlementsView, and attachGuildPlans must leave every
// Guild.Plan nil even for a verified officer of a guild that does have
// an active plan in the database — the nil check has to short-circuit
// before ever asking, not merely happen to find nothing.
func TestMeWorksWithNoEntitlementsStoreConfigured(t *testing.T) {
	h := newHarness(t)
	h.svc.Entitlements = nil

	gid := h.seedGuild(t, "me-guild-no-entitlements-store")
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`insert into entitlements (guild_id, plan, source, status) values ($1, 'guild', 'grant', 'active')`,
		gid); err != nil {
		t.Fatal(err)
	}
	officer := h.seedUser(t, "me-no-entitlements-store@example.com")
	h.seedGuildMember(t, gid, officer, "officer", true)

	res := h.sessionRequest(t, http.MethodGet, "/v1/me", officer)
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		t.Fatalf("status = %d, want 200 with no panic or error even though Entitlements is nil", res.StatusCode)
	}
	var body struct {
		Entitlements struct {
			ServerSims    bool `json:"server_sims"`
			Retention     bool `json:"retention"`
			MultiCompare  bool `json:"multi_compare"`
			History       bool `json:"history"`
			Notifications bool `json:"notifications"`
			OfficerViews  bool `json:"officer_views"`
			RosterCheck   bool `json:"roster_check"`
			SupporterMark bool `json:"supporter_mark"`
			Billing       *struct {
				Status string `json:"status"`
			} `json:"billing"`
		} `json:"entitlements"`
		Guilds []struct {
			ID   int64 `json:"id"`
			Plan *struct {
				Status string `json:"status"`
			} `json:"plan,omitempty"`
		} `json:"guilds"`
	}
	h.decode(t, res, &body)
	ev := body.Entitlements
	if ev.ServerSims || ev.Retention || ev.MultiCompare || ev.History || ev.Notifications ||
		ev.OfficerViews || ev.RosterCheck || ev.SupporterMark {
		t.Fatalf("a nil Entitlements store must answer every feature false: %+v", ev)
	}
	if ev.Billing != nil {
		t.Fatalf("billing should be nil with no Entitlements store, got %+v", ev.Billing)
	}
	if len(body.Guilds) != 1 || body.Guilds[0].Plan != nil {
		t.Fatalf("attachGuildPlans must leave Plan nil with no Entitlements store, even for a "+
			"verified officer of a guild with an active plan row: %+v", body.Guilds)
	}
}

// TestMeCharacterShapeForAGuildedBnetCharacter pins spec §6's frozen
// GET /v1/me interface: a single Battle.net-imported, guilded character,
// with the exact fields the design's own example names.
func TestMeCharacterShapeForAGuildedBnetCharacter(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "guilded-bnet@example.com")
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`update users set bnet_imported_at = '2026-09-22T03:30:00Z' where id = $1`, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`insert into characters (key, region, ruleset, name, class, user_id, realm_name, level, faction, source, imported_at, refreshed_at)
		 values ('us/pvp/thoradin', 'us', 'pvp', 'Thoradin', 'warrior', $1, 'Whitemane', 60, 'alliance', 'bnet', now(), now())`,
		uid); err != nil {
		t.Fatal(err)
	}
	gid := h.seedGuild(t, "Iron Vanguard")
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`insert into guild_characters (guild_id, character_key, user_id, rank, rank_index, source, verified_by, verified_at)
		 values ($1, 'us/pvp/thoradin', $2, 'officer', 1, 'bnet', 'bnet', now())`,
		gid, uid); err != nil {
		t.Fatal(err)
	}

	res := h.sessionRequest(t, http.MethodGet, "/v1/me", uid)
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		t.Fatalf("status = %d", res.StatusCode)
	}
	var body struct {
		BnetImportedAt string      `json:"bnet_imported_at"`
		Characters     []Character `json:"characters"`
	}
	h.decode(t, res, &body)

	if body.BnetImportedAt == "" {
		t.Fatal("bnet_imported_at was omitted, want it present")
	}
	if len(body.Characters) != 1 {
		t.Fatalf("characters = %+v, want exactly one", body.Characters)
	}
	c := body.Characters[0]
	if c.Key != "us/pvp/thoradin" || c.Region != "us" || c.Ruleset != "pvp" || c.Name != "Thoradin" ||
		c.Class != "warrior" || c.Realm != "Whitemane" || c.Level == nil || *c.Level != 60 ||
		c.Faction != "alliance" || c.Source != "bnet" {
		t.Fatalf("character = %+v, does not match spec §6's shape", c)
	}
	if c.Guild == nil || c.Guild.ID != gid || c.Guild.Name != "Iron Vanguard" ||
		c.Guild.Rank != "officer" || c.Guild.RankIndex == nil || *c.Guild.RankIndex != 1 || !c.Guild.Verified {
		t.Fatalf("guild = %+v, does not match spec §6's shape", c.Guild)
	}
}
