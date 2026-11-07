package auth

import (
	"net/http"
	"net/url"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// signInFailedParam and signInFailedValue are what a failed web Battle.net sign-in appends
// to the page it returns to, so /setup and /connect can say it did not finish
// (design/specs/2026-11-07-setup.md section 9, G2). The reason stays in the log, never the URL.
const (
	signInFailedParam = "signin"
	signInFailedValue = "failed"
)

// signInFailedFallback is where a failed web sign-in returns when the flow carries no next.
const signInFailedFallback = "/setup"

// callbackFailure is one way a Battle.net callback can end without a sign-in: the JSON
// answer an app sign-in still gets, and the detail to log (nil when there is nothing to log).
type callbackFailure struct {
	status  int
	code    string
	message string
	err     error
}

// bnetCallback finishes every Battle.net sign-in. A web sign-in starts a session and goes
// to next; an app sign-in (next is appFlowNext) hands a one-time code to the app link.
func (s *Service) bnetCallback(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivate(w)
	state, next, f := s.checkBnetState(w, r)
	if f != nil {
		s.answerCallbackFailure(w, r, next, f)
		return
	}
	bu, u, f := s.identifyBnet(r)
	if f != nil {
		s.answerCallbackFailure(w, r, next, f)
		return
	}
	if next == appFlowNext {
		s.finishAppSignIn(w, r, state, u.ID, bu.AccessToken)
		return
	}
	if _, err := s.Auth.StartSession(r.Context(), w, u.ID, "battlenet"); err != nil {
		s.answerCallbackFailure(w, r, next, internalCallbackFailure(err))
		return
	}
	s.runImport(r, u.ID, bu.AccessToken)
	http.Redirect(w, r, s.PublicBaseURL+safeNext(next), http.StatusFound)
}

func internalCallbackFailure(err error) *callbackFailure {
	return &callbackFailure{http.StatusInternalServerError, "internal", "could not sign you in just now", err}
}

// answerCallbackFailure logs a failed callback as before. An app sign-in keeps its JSON
// answer (the app reads it); a web sign-in goes back to the site with signin=failed.
func (s *Service) answerCallbackFailure(w http.ResponseWriter, r *http.Request, next string, f *callbackFailure) {
	if f.err != nil {
		s.logger().Error("auth", "id", httpx.RequestIDFrom(r.Context()), "op", "battlenet", "err", f.err)
	}
	if next == appFlowNext {
		httpx.WriteError(w, r, f.status, f.code, f.message, nil)
		return
	}
	http.Redirect(w, r, s.PublicBaseURL+signInFailedPath(next), http.StatusFound)
}

// signInFailedPath is next (or /setup when next is missing or unsafe) with signin=failed
// added to whatever query it already carries.
func signInFailedPath(next string) string {
	u, err := url.Parse(safeNextOr(next, signInFailedFallback))
	if err != nil {
		u = &url.URL{Path: signInFailedFallback}
	}
	q := u.Query()
	q.Set(signInFailedParam, signInFailedValue)
	u.RawQuery = q.Encode()
	return u.RequestURI()
}

// checkBnetState reads and clears the state cookie and checks Battle.net echoed its state.
// next is whatever the cookie carried, even when the check fails, so the failure can return
// the browser where it was going.
func (s *Service) checkBnetState(w http.ResponseWriter, r *http.Request) (state, next string, f *callbackFailure) {
	c, err := r.Cookie(bnetStateCookie)
	if err != nil || c.Value == "" {
		return "", "", &callbackFailure{status: http.StatusBadRequest, code: "invalid",
			message: "this sign-in has expired; start again"}
	}
	state, next = decodeState(c.Value)
	http.SetCookie(w, &http.Cookie{Name: bnetStateCookie, Value: "", Path: "/",
		Domain: s.Auth.CookieDomain, HttpOnly: true, Secure: s.Auth.Secure, MaxAge: -1})
	if state == "" || !SameToken(state, r.URL.Query().Get("state")) {
		return "", next, &callbackFailure{status: http.StatusBadRequest, code: "invalid",
			message: "this sign-in could not be verified; start again"}
	}
	return state, next, nil
}

// identifyBnet spends Battle.net's authorization code and upserts the account behind it.
func (s *Service) identifyBnet(r *http.Request) (BnetUser, User, *callbackFailure) {
	code := r.URL.Query().Get("code")
	if code == "" {
		return BnetUser{}, User{}, &callbackFailure{status: http.StatusBadRequest, code: "invalid",
			message: "Battle.net sent no authorization code"}
	}
	bu, err := s.BNet.Identify(r.Context(), code)
	if err != nil {
		return BnetUser{}, User{}, &callbackFailure{http.StatusBadGateway, "upstream",
			"Battle.net could not confirm that sign-in", err}
	}
	u, err := s.Store.UpsertBnetUser(r.Context(), bu.Sub, bu.Battletag)
	if err != nil {
		return BnetUser{}, User{}, internalCallbackFailure(err)
	}
	return bu, u, nil
}
