package auth

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// The cookie names the contract fixes.
const (
	SessionCookie = "fs_session"
	CSRFCookie    = "fs_csrf"
	CSRFHeader    = "X-CSRF-Token"
)

// Actor is who made a request: a signed-in person, a paired device, or
// nobody. A device carries its owner's user id too, so a companion's
// uploads land on the right account.
type Actor struct {
	UserID   int64
	Role     string
	Method   string
	DeviceID string
}

// Signed reports whether the request carries any identity at all.
func (a Actor) Signed() bool { return a.UserID != 0 }

// IsDevice reports whether the identity came from a device token.
func (a Actor) IsDevice() bool { return a.DeviceID != "" }

// IsModerator reports whether the actor may act on other people's
// reports and rankings.
func (a Actor) IsModerator() bool { return a.Role == "moderator" || a.Role == "admin" }

type ctxKey int

const actorKey ctxKey = 1

// WithActor puts an actor on a context, for tests and for the middleware.
func WithActor(ctx context.Context, a Actor) context.Context {
	return context.WithValue(ctx, actorKey, a)
}

// ActorFrom reads the actor the middleware resolved. The zero Actor means
// an anonymous request.
func ActorFrom(ctx context.Context) Actor {
	a, _ := ctx.Value(actorKey).(Actor)
	return a
}

// Authenticator resolves identity and enforces CSRF. It never rejects an
// anonymous request: routes that need identity say so with Require.
type Authenticator struct {
	Store        *Store
	CookieDomain string
	// Secure marks the cookies Secure. It is true in production and
	// false for a plain-HTTP local run, where a Secure cookie would
	// never be sent back.
	Secure bool
	Log    *slog.Logger
}

func (a *Authenticator) logger() *slog.Logger {
	if a.Log != nil {
		return a.Log
	}
	return slog.Default()
}

// Middleware resolves the request's actor and enforces double-submit
// CSRF on state-changing requests that carry a session cookie. Requests
// authenticated by a device token are exempt, as the contract says: a
// companion is not a browser and cannot be made to submit a form.
func (a *Authenticator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token, ok := bearerToken(r); ok {
			d, err := a.Store.DeviceByToken(r.Context(), TokenHash(token))
			if err == nil {
				u, uerr := a.Store.User(r.Context(), d.UserID)
				if uerr == nil {
					next.ServeHTTP(w, r.WithContext(WithActor(r.Context(),
						Actor{UserID: u.ID, Role: u.Role, Method: "device", DeviceID: d.ID})))
					return
				}
				err = uerr
			}
			if err != ErrNotFound {
				a.logger().Error("auth", "id", httpx.RequestIDFrom(r.Context()), "op", "device", "err", err)
			}
			httpx.WriteError(w, r, http.StatusUnauthorized, "unauthorized", "that device token is not valid", nil)
			return
		}

		c, err := r.Cookie(SessionCookie)
		if err != nil || c.Value == "" {
			next.ServeHTTP(w, r)
			return
		}
		u, err := a.Store.SessionUser(r.Context(), c.Value)
		if err == ErrNotFound {
			// An expired or unknown session is simply anonymous, and the
			// stale cookie is cleared so the browser stops sending it.
			a.ClearSession(w)
			next.ServeHTTP(w, r)
			return
		}
		if err != nil {
			a.logger().Error("auth", "id", httpx.RequestIDFrom(r.Context()), "op", "session", "err", err)
			httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not read that session", nil)
			return
		}
		// fs_csrf is readable by script and outlives nothing that
		// fs_session does not, so a browser can end up holding the
		// session and not the token - and then every state-changing
		// request 403s with no way out, DELETE /v1/sessions included.
		// Re-issuing it here is the recovery: this request still fails
		// the check below, because the caller could not have sent a
		// token it did not have, but the retry carries one. Handing a
		// fresh random token to a request that is missing it gives a
		// cross-site caller nothing: it still cannot read the cookie to
		// echo it in the header.
		if c, err := r.Cookie(CSRFCookie); err != nil || c.Value == "" {
			a.setCookie(w, CSRFCookie, NewSessionID(), SessionTTL, false)
		}
		if stateChanging(r.Method) && !a.csrfOK(r) {
			httpx.WriteError(w, r, http.StatusForbidden, "csrf",
				"this request needs the X-CSRF-Token header to match the fs_csrf cookie", nil)
			return
		}
		next.ServeHTTP(w, r.WithContext(WithActor(r.Context(),
			Actor{UserID: u.ID, Role: u.Role, Method: "session"})))
	})
}

// bearerToken reads a device token from the Authorization header.
func bearerToken(r *http.Request) (string, bool) {
	h := r.Header.Get("Authorization")
	if !strings.HasPrefix(h, "Bearer ") {
		return "", false
	}
	token := strings.TrimSpace(strings.TrimPrefix(h, "Bearer "))
	if !strings.HasPrefix(token, DeviceTokenPrefix) {
		return "", false
	}
	return token, true
}

func stateChanging(method string) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		return false
	default:
		return true
	}
}

func (a *Authenticator) csrfOK(r *http.Request) bool {
	c, err := r.Cookie(CSRFCookie)
	if err != nil || c.Value == "" {
		return false
	}
	return SameToken(c.Value, r.Header.Get(CSRFHeader))
}

// StartSession creates a session for a user and sets both cookies.
func (a *Authenticator) StartSession(ctx context.Context, w http.ResponseWriter, userID int64, method string) (Session, error) {
	sess, err := a.Store.CreateSession(ctx, userID, method, SessionTTL)
	if err != nil {
		return Session{}, err
	}
	a.setCookie(w, SessionCookie, sess.ID, SessionTTL, true)
	a.setCookie(w, CSRFCookie, NewSessionID(), SessionTTL, false)
	return sess, nil
}

// EndSession deletes the session the request carries and clears both
// cookies. An anonymous request is a no-op, so signing out twice is fine.
func (a *Authenticator) EndSession(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c, err := r.Cookie(SessionCookie)
	if err == nil && c.Value != "" {
		if err := a.Store.DeleteSession(ctx, c.Value); err != nil {
			return err
		}
	}
	a.ClearSession(w)
	return nil
}

// ClearSession expires both cookies in the browser.
func (a *Authenticator) ClearSession(w http.ResponseWriter) {
	a.setCookie(w, SessionCookie, "", -time.Hour, true)
	a.setCookie(w, CSRFCookie, "", -time.Hour, false)
}

func (a *Authenticator) setCookie(w http.ResponseWriter, name, value string, ttl time.Duration, httpOnly bool) {
	c := &http.Cookie{
		Name: name, Value: value, Path: "/", Domain: a.CookieDomain,
		HttpOnly: httpOnly, Secure: a.Secure, SameSite: http.SameSiteLaxMode,
		Expires: time.Now().UTC().Add(ttl), MaxAge: int(ttl.Seconds()),
	}
	http.SetCookie(w, c)
}

// Require wraps a handler so only a signed-in person or device reaches
// it, answering 401 in the envelope otherwise.
func Require(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !ActorFrom(r.Context()).Signed() {
			httpx.WriteError(w, r, http.StatusUnauthorized, "unauthorized", "sign in first", nil)
			return
		}
		h(w, r)
	}
}

// RequireSession wraps a handler so only a browser session reaches it: a
// device token cannot pair another device or sign anything out.
func RequireSession(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a := ActorFrom(r.Context())
		if !a.Signed() || a.IsDevice() {
			httpx.WriteError(w, r, http.StatusUnauthorized, "unauthorized", "sign in on the site first", nil)
			return
		}
		h(w, r)
	}
}

// RequireDevice wraps a handler so only a companion reaches it.
func RequireDevice(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !ActorFrom(r.Context()).IsDevice() {
			httpx.WriteError(w, r, http.StatusUnauthorized, "unauthorized",
				"this route needs a device token", nil)
			return
		}
		h(w, r)
	}
}
