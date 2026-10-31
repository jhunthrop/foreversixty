package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/textx"
)

// The native app sign-in: Battle.net OAuth in the system browser, with PKCE (RFC 7636)
// binding the browser leg to the app that started it.
//
//  1. The app opens GET /v1/auth/app/start with its platform, a S256 code_challenge and its
//     own state. The flow is recorded, the state is bound to this browser in the same cookie
//     the site's sign-in uses, and a confirm page (app_confirm.go) asks for a click. The
//     click, GET /v1/auth/app/continue, goes to Battle.net. Nothing redirects on its own, so
//     a link someone else sends cannot finish a sign-in while Battle.net consents silently.
//     Any redirect_uri the client sends is ignored: Battle.net always gets the configured
//     callback and the code always goes to appLinkPath.
//  2. Battle.net returns to the shared callback. For an app flow it mints a one-time code
//     (stored hashed, single use, two minutes) and redirects to the app link
//     https://foreversixty.gg/app/auth?code=…&state=<the app's state>. No web session is set.
//  3. The app posts the code and its code_verifier to POST /v1/auth/app/exchange and gets a
//     device token for a devices row of kind app.
//  4. POST /v1/auth/app/logout revokes the calling device.

// appFlowNext is the next value the state cookie carries for an app sign-in. safeNext only
// ever yields a path starting with "/", so a web sign-in can never carry it.
const appFlowNext = "app"

// appLinkPath is the universal link / App Link path the callback hands the code to. Only
// /app/* opens the apps (web/public/.well-known).
const appLinkPath = "/app/auth"

// Per-IP budgets for the two unauthenticated app routes. They are looser than the
// companion's claim budget because phones on one carrier often share an address (CGNAT);
// the code and the verifier carry 256 bits each, so the budget is about abuse, not guessing.
const (
	appStartsPerHour    = 60
	appExchangesPerHour = 60
)

// appPlatforms are the platforms an app device can carry, with the name a device gets when
// the app sends none.
var appPlatforms = map[string]string{"ios": "iOS", "android": "Android"}

var (
	// s256Challenge is a base64url SHA-256 with no padding.
	s256Challenge = regexp.MustCompile(`^[A-Za-z0-9_-]{43}$`)
	// pkceVerifier is RFC 7636 §4.1's code_verifier.
	pkceVerifier = regexp.MustCompile(`^[A-Za-z0-9._~-]{43,128}$`)
	// appState is the app's own CSRF state, echoed back on the app link.
	appState = regexp.MustCompile(`^[A-Za-z0-9._~-]{16,128}$`)
)

func mountApp(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	startLimit := httpx.RateLimitPer(appStartsPerHour, time.Hour, trustedProxyHops)
	exchangeLimit := httpx.RateLimitPer(appExchangesPerHour, time.Hour, trustedProxyHops)
	mux.Handle("GET /v1/auth/app/start", startLimit(http.HandlerFunc(s.appStart)))
	mux.Handle("GET /v1/auth/app/continue", startLimit(http.HandlerFunc(s.appContinue)))
	mux.Handle("POST /v1/auth/app/exchange", exchangeLimit(http.HandlerFunc(s.appExchange)))
	mux.HandleFunc("POST /v1/auth/app/logout", RequireApp(s.appLogout))
}

// parseAppStart validates the start query, naming each field that is wrong.
func parseAppStart(q url.Values) (AppFlow, map[string]string) {
	f := AppFlow{Platform: q.Get("platform"), CodeChallenge: q.Get("code_challenge"), AppState: q.Get("state")}
	bad := map[string]string{}
	if _, ok := appPlatforms[f.Platform]; !ok {
		bad["platform"] = "ios or android"
	}
	if q.Get("code_challenge_method") != "S256" {
		bad["code_challenge_method"] = "S256"
	}
	if !s256Challenge.MatchString(f.CodeChallenge) {
		bad["code_challenge"] = "the base64url SHA-256 of the code verifier, 43 characters"
	}
	if !appState.MatchString(f.AppState) {
		bad["state"] = "16 to 128 URL-safe characters"
	}
	if len(bad) > 0 {
		return AppFlow{}, bad
	}
	return f, nil
}

func (s *Service) appStart(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivate(w)
	f, bad := parseAppStart(r.URL.Query())
	if bad != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that app sign-in request is not valid", bad)
		return
	}
	state := NewSessionID()
	if err := s.Store.CreateAppFlow(r.Context(), TokenHash(state), f); err != nil {
		s.fail(w, r, "app_start", err, "could not start that sign-in just now")
		return
	}
	s.setBnetState(w, state, appFlowNext)
	s.writeAppConfirm(w, r, "/v1/auth/app/continue?"+url.Values{"state": {state}}.Encode())
}

// appContinue is the confirm page's button. It goes on to Battle.net only for the browser
// whose start set the state cookie, and only for an app flow; the cookie stays for the
// callback.
func (s *Service) appContinue(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivate(w)
	c, err := r.Cookie(bnetStateCookie)
	if err != nil || c.Value == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "this sign-in has expired; start again in the app", nil)
		return
	}
	state, next := decodeState(c.Value)
	if next != appFlowNext || state == "" || !SameToken(state, r.URL.Query().Get("state")) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "this sign-in could not be verified; start again in the app", nil)
		return
	}
	http.Redirect(w, r, s.BNet.AuthURL(state), http.StatusFound)
}

// finishAppSignIn is the app half of the Battle.net callback: the account is confirmed, so
// the flow gets its one-time code and the browser goes to the app link.
func (s *Service) finishAppSignIn(w http.ResponseWriter, r *http.Request, state string, userID int64, accessToken string) {
	code := NewSessionID()
	f, err := s.Store.IssueAppCode(r.Context(), TokenHash(state), userID, TokenHash(code))
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "this sign-in has expired; start again in the app", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "app_callback", err, "could not sign you in just now")
		return
	}
	s.runImport(r, userID, accessToken)
	q := url.Values{"code": {code}, "state": {f.AppState}}
	http.Redirect(w, r, s.PublicBaseURL+appLinkPath+"?"+q.Encode(), http.StatusFound)
}

type appExchangeRequest struct {
	Code         string `json:"code"`
	CodeVerifier string `json:"code_verifier"`
	DeviceName   string `json:"device_name"`
}

// AppDevice is the device block the exchange answers.
type AppDevice struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

// AppExchangeResult is POST /v1/auth/app/exchange's data: the token, shown only here.
type AppExchangeResult struct {
	Token  string    `json:"token"`
	Device AppDevice `json:"device"`
	Me     Me        `json:"me"`
}

// pkceMatches reports whether verifier hashes to the S256 challenge.
func pkceMatches(verifier, challenge string) bool {
	sum := sha256.Sum256([]byte(verifier))
	return SameToken(base64.RawURLEncoding.EncodeToString(sum[:]), challenge)
}

func (s *Service) appExchange(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivate(w)
	var in appExchangeRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxAuthBody)).Decode(&in); err != nil ||
		in.Code == "" || !pkceVerifier.MatchString(in.CodeVerifier) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with a code and a code_verifier", nil)
		return
	}
	userID, f, err := s.Store.SpendAppCode(r.Context(), TokenHash(in.Code))
	if err == nil && !pkceMatches(in.CodeVerifier, f.CodeChallenge) {
		err = ErrNotFound
	}
	if errors.Is(err, ErrNotFound) {
		// One answer for unknown, spent, expired and wrong-verifier, so none can be told apart.
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid_grant",
			"that sign-in code is not valid; sign in again", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "app_exchange", err, "could not sign you in just now")
		return
	}
	s.issueAppDevice(w, r, userID, f.Platform, in.DeviceName)
}

// issueAppDevice makes the app's device and answers with its token and the account read.
func (s *Service) issueAppDevice(w http.ResponseWriter, r *http.Request, userID int64, platform, name string) {
	name = textx.Trim(name, maxDeviceName)
	if name == "" {
		name = appPlatforms[platform]
	}
	token, hash := NewDeviceToken()
	d, err := s.Store.CreateAppDevice(r.Context(), userID, name, platform, hash)
	if err != nil {
		s.fail(w, r, "app_exchange", err, "could not sign you in just now")
		return
	}
	me, err := s.loadMe(r.Context(), userID, false)
	if err != nil {
		s.fail(w, r, "app_exchange", err, "could not sign you in just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusCreated, AppExchangeResult{
		Token: token, Device: AppDevice{ID: d.ID, Name: d.Name, Platform: d.Platform}, Me: me,
	})
}

// appLogout is the app's Sign out: the device behind the bearer token revokes itself. A
// token already dead never gets here; the middleware answers it 401.
func (s *Service) appLogout(w http.ResponseWriter, r *http.Request) {
	if _, err := s.Store.RevokeDeviceByID(r.Context(), ActorFrom(r.Context()).DeviceID); err != nil {
		s.fail(w, r, "app_logout", err, "could not sign you out just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]bool{"signed_out": true})
}
