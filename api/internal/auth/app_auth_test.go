package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"html"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"
)

// The verifier an app holds and the challenge it sends (RFC 7636 appendix B's example pair).
const (
	testVerifier  = "dBjftJeZ4CVP-mB92K0uEtqmtSx1hzXuHg7vyOQGxZ8wjE"
	testAppState  = "app-state-0123456789abcdef"
	wrongVerifier = "this-is-not-the-verifier-the-app-started-with-at-all"
)

func challengeOf(verifier string) string {
	sum := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

// withAppSignIn configures Battle.net and turns the app routes on.
func (h *harness) withAppSignIn(t *testing.T) {
	t.Helper()
	h.withBattleNet(t, "12345", "Baelgrim#1234")
	h.svc.AppSignIn = true
	mux := http.NewServeMux()
	Mount(mux, h.svc, 0)
	h.server.Config.Handler = h.svc.Auth.Middleware(mux)
}

func appStartQuery(platform, challenge, method, state string) string {
	q := url.Values{}
	for k, v := range map[string]string{
		"platform": platform, "code_challenge": challenge, "code_challenge_method": method, "state": state,
	} {
		if v != "" {
			q.Set(k, v)
		}
	}
	return "/v1/auth/app/start?" + q.Encode()
}

// appStart begins an app sign-in: the confirm page, then the click through to Battle.net.
// It returns the state Battle.net was sent.
func (h *harness) appStart(t *testing.T, platform string) string {
	t.Helper()
	page := h.do(t, http.MethodGet, appStartQuery(platform, challengeOf(testVerifier), "S256", testAppState), "")
	cont := continueLink(t, page)
	res := h.do(t, http.MethodGet, cont, "")
	res.Body.Close()
	if res.StatusCode != http.StatusFound {
		t.Fatalf("continue = %d, want a redirect to Battle.net", res.StatusCode)
	}
	target, err := url.Parse(res.Header.Get("Location"))
	if err != nil {
		t.Fatal(err)
	}
	state := target.Query().Get("state")
	if state == "" || state == testAppState {
		t.Fatalf("Battle.net state = %q, want a fresh server state", state)
	}
	return state
}

// continueLink reads the confirm page and returns the path its button goes to.
func continueLink(t *testing.T, res *http.Response) string {
	t.Helper()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("app start = %d, want the confirm page", res.StatusCode)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	m := regexp.MustCompile(`href="(/v1/auth/app/continue\?[^"]+)"`).FindStringSubmatch(string(b))
	if m == nil {
		t.Fatalf("the confirm page has no continue link:\n%s", b)
	}
	return html.UnescapeString(m[1])
}

// appCallback finishes Battle.net and returns the one-time code from the app link.
func (h *harness) appCallback(t *testing.T, state string) string {
	t.Helper()
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state="+state, "")
	res.Body.Close()
	if res.StatusCode != http.StatusFound {
		t.Fatalf("callback = %d, want a redirect to the app link", res.StatusCode)
	}
	loc, err := url.Parse(res.Header.Get("Location"))
	if err != nil {
		t.Fatal(err)
	}
	if loc.Scheme != "https" || loc.Host != "foreversixty.gg" || loc.Path != "/app/auth" {
		t.Fatalf("redirected to %q, want https://foreversixty.gg/app/auth", loc)
	}
	if got := loc.Query().Get("state"); got != testAppState {
		t.Fatalf("app link state = %q, want the app's own %q", got, testAppState)
	}
	for _, c := range res.Cookies() {
		if c.Name == SessionCookie && c.Value != "" {
			t.Fatal("an app sign-in must not set a web session")
		}
	}
	code := loc.Query().Get("code")
	if code == "" {
		t.Fatal("the app link carries no code")
	}
	return code
}

type appExchangeBody struct {
	Token  string `json:"token"`
	Device struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Platform string `json:"platform"`
	} `json:"device"`
	Me Me `json:"me"`
}

// exchange posts the code with no cookies at all, the way the app does.
func (h *harness) exchange(t *testing.T, code, verifier string) *http.Response {
	t.Helper()
	body, _ := json.Marshal(map[string]string{"code": code, "code_verifier": verifier, "device_name": "Baelgrim's iPhone"})
	res, err := http.Post(h.server.URL+"/v1/auth/app/exchange", "application/json", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func bearer(t *testing.T, h *harness, method, path, token string) *http.Response {
	t.Helper()
	r, err := http.NewRequest(method, h.server.URL+path, nil)
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func wantStatus(t *testing.T, res *http.Response, want int, what string) {
	t.Helper()
	res.Body.Close()
	if res.StatusCode != want {
		t.Fatalf("%s = %d, want %d", what, res.StatusCode, want)
	}
}

func (h *harness) signInApp(t *testing.T, platform string) appExchangeBody {
	t.Helper()
	code := h.appCallback(t, h.appStart(t, platform))
	res := h.exchange(t, code, testVerifier)
	if res.StatusCode != http.StatusCreated {
		res.Body.Close()
		t.Fatalf("exchange = %d, want 201", res.StatusCode)
	}
	var out appExchangeBody
	h.decode(t, res, &out)
	return out
}

func TestAppSignInRoutesAreDarkByDefault(t *testing.T) {
	h := newHarness(t)
	h.withBattleNet(t, "12345", "Baelgrim#1234")
	wantStatus(t, h.do(t, http.MethodGet, appStartQuery("ios", challengeOf(testVerifier), "S256", testAppState), ""),
		http.StatusNotFound, "app start with the flag off")
	wantStatus(t, h.exchange(t, "x", testVerifier), http.StatusNotFound, "exchange with the flag off")
	wantStatus(t, h.do(t, http.MethodPost, "/v1/auth/app/logout", ""), http.StatusNotFound, "logout with the flag off")
}

func TestAppSignInHappyPath(t *testing.T) {
	for _, platform := range []string{"ios", "android"} {
		t.Run(platform, func(t *testing.T) {
			h := newHarness(t)
			h.withAppSignIn(t)
			got := h.signInApp(t, platform)
			if !strings.HasPrefix(got.Token, DeviceTokenPrefix) {
				t.Fatalf("token = %q", got.Token)
			}
			if got.Device.ID == "" || got.Device.Platform != platform || got.Device.Name != "Baelgrim's iPhone" {
				t.Fatalf("device = %+v", got.Device)
			}
			if got.Me.User.Battletag == nil || *got.Me.User.Battletag != "Baelgrim#1234" {
				t.Fatalf("me = %+v", got.Me.User)
			}

			// The token reads the account, with no cookie and no CSRF header.
			res := bearer(t, h, http.MethodGet, "/v1/me", got.Token)
			var me Me
			h.decode(t, res, &me)
			if me.User.ID != got.Me.User.ID {
				t.Fatalf("/v1/me as the app = user %d, want %d", me.User.ID, got.Me.User.ID)
			}

			// The site's account page lists the app beside any companion, marked as an app.
			res = bearer(t, h, http.MethodGet, "/v1/devices", got.Token)
			var list []Device
			h.decode(t, res, &list)
			if len(list) != 1 || list[0].ID != got.Device.ID || list[0].Kind != DeviceKindApp || list[0].Platform != platform {
				t.Fatalf("devices = %+v", list)
			}

			// Only the code's hash is at rest.
			var n int
			if err := h.svc.Store.Pool.QueryRow(t.Context(),
				`select count(*) from app_auth_flows where code_hash is not null and used_at is not null`).Scan(&n); err != nil || n != 1 {
				t.Fatalf("spent flows = %d, %v", n, err)
			}
		})
	}
}

func TestAppExchangeRefusesAWrongVerifierAndBurnsTheCode(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	code := h.appCallback(t, h.appStart(t, "ios"))
	res := h.exchange(t, code, wrongVerifier)
	wantErrorCode(t, res, http.StatusBadRequest, "invalid_grant")
	// A stolen code tried without its verifier is spent: the real app must start again.
	wantErrorCode(t, h.exchange(t, code, testVerifier), http.StatusBadRequest, "invalid_grant")
	assertNoDevices(t, h)
}

func TestAppExchangeRefusesAReplayedCode(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	code := h.appCallback(t, h.appStart(t, "android"))
	wantStatus(t, h.exchange(t, code, testVerifier), http.StatusCreated, "first exchange")
	wantErrorCode(t, h.exchange(t, code, testVerifier), http.StatusBadRequest, "invalid_grant")
}

func TestAppExchangeRefusesAnExpiredCode(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	code := h.appCallback(t, h.appStart(t, "ios"))
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`update app_auth_flows set code_expires_at = now() - interval '1 second'`); err != nil {
		t.Fatal(err)
	}
	wantErrorCode(t, h.exchange(t, code, testVerifier), http.StatusBadRequest, "invalid_grant")
	assertNoDevices(t, h)
}

func TestAppCodeExpiresInTwoMinutes(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	h.appCallback(t, h.appStart(t, "ios"))
	var secs float64
	if err := h.svc.Store.Pool.QueryRow(t.Context(),
		`select extract(epoch from code_expires_at - now()) from app_auth_flows`).Scan(&secs); err != nil {
		t.Fatal(err)
	}
	if secs <= 100 || secs > 120 {
		t.Fatalf("code lives %.0fs, want about 120", secs)
	}
}

func TestAppCallbackRefusesAStateMismatch(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	h.appStart(t, "ios")
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state=not-the-state", "")
	wantStatus(t, res, http.StatusBadRequest, "callback with a forged state")
	var n int
	if err := h.svc.Store.Pool.QueryRow(t.Context(),
		`select count(*) from app_auth_flows where code_hash is not null`).Scan(&n); err != nil || n != 0 {
		t.Fatalf("codes minted = %d, %v; a forged state must mint none", n, err)
	}
}

// A state that passes the cookie check but whose app flow has expired gets no code and no
// web session either.
func TestAppCallbackRefusesAnExpiredFlow(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	state := h.appStart(t, "ios")
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`update app_auth_flows set expires_at = now() - interval '1 second'`); err != nil {
		t.Fatal(err)
	}
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state="+state, "")
	wantStatus(t, res, http.StatusBadRequest, "callback for an expired app flow")
	for _, c := range res.Cookies() {
		if c.Name == SessionCookie && c.Value != "" {
			t.Fatal("an expired app flow must not fall through to a web session")
		}
	}
}

func TestAppStartValidatesItsQuery(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	good := challengeOf(testVerifier)
	for name, q := range map[string]string{
		"no platform":      appStartQuery("", good, "S256", testAppState),
		"desktop platform": appStartQuery("windows", good, "S256", testAppState),
		"no challenge":     appStartQuery("ios", "", "S256", testAppState),
		"short challenge":  appStartQuery("ios", "abc", "S256", testAppState),
		"plain method":     appStartQuery("ios", good, "plain", testAppState),
		"no method":        appStartQuery("ios", good, "", testAppState),
		"no state":         appStartQuery("ios", good, "S256", ""),
		"short state":      appStartQuery("ios", good, "S256", "abc"),
		"state with space": appStartQuery("ios", good, "S256", "a state with spaces in it"),
	} {
		res := h.do(t, http.MethodGet, q, "")
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s: start = %d, want 400", name, res.StatusCode)
		}
	}
}

func TestAppExchangeValidatesItsBody(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	for name, body := range map[string]string{
		"not json":    `{`,
		"no code":     `{"code_verifier":"` + testVerifier + `"}`,
		"no verifier": `{"code":"abc"}`,
	} {
		res, err := http.Post(h.server.URL+"/v1/auth/app/exchange", "application/json", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s: exchange = %d, want 400", name, res.StatusCode)
		}
	}
}

func TestAppDeviceNameDefaultsToThePlatform(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	code := h.appCallback(t, h.appStart(t, "android"))
	res, err := http.Post(h.server.URL+"/v1/auth/app/exchange", "application/json",
		strings.NewReader(`{"code":"`+code+`","code_verifier":"`+testVerifier+`"}`))
	if err != nil {
		t.Fatal(err)
	}
	var out appExchangeBody
	h.decode(t, res, &out)
	if out.Device.Name != "Android" {
		t.Fatalf("name = %q, want Android", out.Device.Name)
	}
}

func TestAppLogoutRevokesTheCallingDevice(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	got := h.signInApp(t, "ios")
	wantStatus(t, bearer(t, h, http.MethodPost, "/v1/auth/app/logout", got.Token), http.StatusOK, "app logout")
	wantStatus(t, bearer(t, h, http.MethodGet, "/v1/me", got.Token), http.StatusUnauthorized, "/v1/me after logout")
	// Logout needs a device token: a cookie-less anonymous call is refused.
	res, err := http.Post(h.server.URL+"/v1/auth/app/logout", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	wantStatus(t, res, http.StatusUnauthorized, "anonymous app logout")
}

func TestAnAppRevokedFromTheSiteIsRefused(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	got := h.signInApp(t, "android")
	// The same person on the site, through the web flow in the same browser.
	h.webBattleNetSignIn(t)
	wantStatus(t, h.do(t, http.MethodDelete, "/v1/devices/"+got.Device.ID, ""), http.StatusNoContent, "revoke from the site")
	wantStatus(t, bearer(t, h, http.MethodGet, "/v1/me", got.Token), http.StatusUnauthorized, "/v1/me after revoke")
}

// webBattleNetSignIn runs the unchanged web flow and checks it lands on the site with a session.
func (h *harness) webBattleNetSignIn(t *testing.T) {
	t.Helper()
	res := h.do(t, http.MethodGet, "/v1/auth/battlenet/start?next=/account", "")
	res.Body.Close()
	target, _ := url.Parse(res.Header.Get("Location"))
	res = h.do(t, http.MethodGet, "/v1/auth/battlenet/callback?code=the-code&state="+target.Query().Get("state"), "")
	res.Body.Close()
	if got := res.Header.Get("Location"); got != "https://foreversixty.gg/account" {
		t.Fatalf("web callback redirected to %q, want the next page", got)
	}
	res = h.do(t, http.MethodGet, "/v1/me", "")
	var me Me
	h.decode(t, res, &me)
	if me.User.Battletag == nil || *me.User.Battletag != "Baelgrim#1234" {
		t.Fatalf("web session me = %+v", me.User)
	}
}

func TestWebBattleNetFlowIsUnchangedWithAppSignInOn(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	h.webBattleNetSignIn(t)
	var n int
	if err := h.svc.Store.Pool.QueryRow(t.Context(), `select count(*) from app_auth_flows`).Scan(&n); err != nil || n != 0 {
		t.Fatalf("app flows = %d, %v; the web flow must write none", n, err)
	}
}

// A companion's token still cannot read the account: only an app's can.
func TestACompanionTokenStillCannotReadTheAccount(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	uid := h.seedUser(t, "raider@example.com")
	token, hash := NewDeviceToken()
	if _, err := h.svc.Store.CreateDevice(t.Context(), uid, "Raid PC", "ios", hash); err != nil {
		t.Fatal(err)
	}
	wantStatus(t, bearer(t, h, http.MethodGet, "/v1/me", token), http.StatusUnauthorized, "/v1/me as a companion named ios")
}

func TestAppStartAndExchangeAreRateLimited(t *testing.T) {
	h := newHarness(t)
	h.withAppSignIn(t)
	q := appStartQuery("ios", challengeOf(testVerifier), "S256", testAppState)
	for i := range appStartsPerHour {
		res := h.do(t, http.MethodGet, q, "")
		res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Fatalf("start %d = %d", i, res.StatusCode)
		}
	}
	wantStatus(t, h.do(t, http.MethodGet, q, ""), http.StatusTooManyRequests, "start past the budget")
	for i := range appExchangesPerHour {
		res := h.exchange(t, "nope", testVerifier)
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("exchange %d = %d", i, res.StatusCode)
		}
	}
	wantStatus(t, h.exchange(t, "nope", testVerifier), http.StatusTooManyRequests, "exchange past the budget")
}

func wantErrorCode(t *testing.T, res *http.Response, status int, code string) {
	t.Helper()
	defer res.Body.Close()
	var env struct {
		OK    bool `json:"ok"`
		Error struct {
			Code string `json:"code"`
		} `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != status || env.OK || env.Error.Code != code {
		t.Fatalf("got %d %+v, want %d %s", res.StatusCode, env, status, code)
	}
}

func assertNoDevices(t *testing.T, h *harness) {
	t.Helper()
	var n int
	if err := h.svc.Store.Pool.QueryRow(t.Context(), `select count(*) from devices`).Scan(&n); err != nil || n != 0 {
		t.Fatalf("devices = %d, %v; want none", n, err)
	}
}
