package auth

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	netmail "net/mail"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// LoginTokenTTL and PairingCodeTTL are Task 5's store.go constants
// (store.go), reused rather than redeclared here so the handlers and the
// store agree on exactly one number for each.
const (
	// LoginsPerHour is the contract's five-links-per-address-per-hour cap.
	LoginsPerHour = 5
	// maxAuthBody is the ceiling on a sign-in or pairing body.
	maxAuthBody = 4 << 10
	// claimsPerHour caps device claims per IP, so pairing codes cannot be
	// guessed by brute force.
	claimsPerHour = 30
	// emailsPerHour caps magic links per IP on top of the per-address cap.
	emailsPerHour = 20
)

// Service serves every sign-in route.
type Service struct {
	Store         *Store
	Auth          *Authenticator
	BNet          *BattleNet
	Mailer        mail.Mailer
	PublicBaseURL string
	APIBaseURL    string
	Log           *slog.Logger
}

// Mount registers the sign-in, device, and account routes. The Battle.net
// pair is registered only when the service is configured for it, so an
// unconfigured deployment answers 404 there rather than redirecting into
// a broken flow.
func Mount(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	if s.BNet != nil {
		mux.HandleFunc("GET /v1/auth/battlenet/start", s.bnetStart)
		mux.HandleFunc("GET /v1/auth/battlenet/callback", s.bnetCallback)
	}
	emailLimit := httpx.RateLimitPer(emailsPerHour, time.Hour, trustedProxyHops)
	mux.Handle("POST /v1/auth/email", emailLimit(http.HandlerFunc(s.emailStart)))
	mux.HandleFunc("GET /v1/auth/email/callback", s.emailCallback)
	mux.HandleFunc("POST /v1/auth/logout", s.logout)
	// The site's own sign-out. It is the same act as the contract's
	// logout route, in the shape the account page uses: a DELETE of
	// the session resource, answering 204.
	mux.HandleFunc("DELETE /v1/sessions", RequireSession(s.deleteSession))
	mux.HandleFunc("GET /v1/me", Require(s.me))
	mux.HandleFunc("PATCH /v1/me", RequireSession(s.patchMe))
	mux.HandleFunc("POST /v1/devices/pair", RequireSession(s.pair))
	claimLimit := httpx.RateLimitPer(claimsPerHour, time.Hour, trustedProxyHops)
	mux.Handle("POST /v1/devices/claim", claimLimit(http.HandlerFunc(s.claim)))
	mux.HandleFunc("GET /v1/devices", RequireSession(s.listDevices))
	mux.HandleFunc("DELETE /v1/devices/{id}", RequireSession(s.revokeDevice))
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

// fail logs the detail and answers with the envelope's generic message,
// so an internal error never reaches the client as prose.
func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	s.logger().Error("auth", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}

func (s *Service) bnetStart(w http.ResponseWriter, r *http.Request) {
	state := NewSessionID()
	next := safeNext(r.URL.Query().Get("next"))
	http.SetCookie(w, &http.Cookie{
		Name: bnetStateCookie, Value: encodeState(state, next), Path: "/",
		Domain: s.Auth.CookieDomain, HttpOnly: true, Secure: s.Auth.Secure,
		SameSite: http.SameSiteLaxMode, MaxAge: int(bnetStateTTL.Seconds()),
	})
	http.Redirect(w, r, s.BNet.AuthURL(state), http.StatusFound)
}

func (s *Service) bnetCallback(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(bnetStateCookie)
	if err != nil || c.Value == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "this sign-in has expired; start again", nil)
		return
	}
	state, next := decodeState(c.Value)
	http.SetCookie(w, &http.Cookie{Name: bnetStateCookie, Value: "", Path: "/",
		Domain: s.Auth.CookieDomain, HttpOnly: true, Secure: s.Auth.Secure, MaxAge: -1})
	if state == "" || !SameToken(state, r.URL.Query().Get("state")) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "this sign-in could not be verified; start again", nil)
		return
	}
	code := r.URL.Query().Get("code")
	if code == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "Battle.net sent no authorization code", nil)
		return
	}
	bu, err := s.BNet.Identify(r.Context(), code)
	if err != nil {
		s.logger().Error("auth", "id", httpx.RequestIDFrom(r.Context()), "op", "battlenet", "err", err)
		httpx.WriteError(w, r, http.StatusBadGateway, "upstream", "Battle.net could not confirm that sign-in", nil)
		return
	}
	u, err := s.Store.UpsertBnetUser(r.Context(), bu.Sub, bu.Battletag)
	if err != nil {
		s.fail(w, r, "battlenet", err, "could not sign you in just now")
		return
	}
	if _, err := s.Auth.StartSession(r.Context(), w, u.ID, "battlenet"); err != nil {
		s.fail(w, r, "battlenet", err, "could not sign you in just now")
		return
	}
	http.Redirect(w, r, s.PublicBaseURL+safeNext(next), http.StatusFound)
}

type emailRequest struct {
	Email string `json:"email"`
}

func (s *Service) emailStart(w http.ResponseWriter, r *http.Request) {
	var in emailRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxAuthBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with an email", nil)
		return
	}
	address, err := normalizeEmail(in.Email)
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not an email address",
			map[string]string{"email": "enter an address like you@example.com"})
		return
	}
	n, err := s.Store.RecentLoginTokens(r.Context(), address, time.Hour)
	if err != nil {
		s.fail(w, r, "email", err, "could not send that link just now")
		return
	}
	if n >= LoginsPerHour {
		httpx.WriteError(w, r, http.StatusTooManyRequests, "rate_limited",
			"that address has been sent five links in the last hour", nil)
		return
	}
	token := NewSessionID()
	if err := s.Store.CreateLoginToken(r.Context(), address, TokenHash(token), LoginTokenTTL); err != nil {
		s.fail(w, r, "email", err, "could not send that link just now")
		return
	}
	link := s.APIBaseURL + "/v1/auth/email/callback?token=" + token
	if err := s.Mailer.Send(r.Context(), mail.Message{
		To:      address,
		Subject: "Your Forever Sixty sign-in link",
		Text: "Sign in to Forever Sixty:\n\n" + link +
			"\n\nThe link works once and expires in 20 minutes. If you did not ask for it, ignore this mail.\n",
	}); err != nil {
		s.fail(w, r, "email", err, "could not send that link just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusAccepted, map[string]bool{"sent": true})
}

func (s *Service) emailCallback(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that link is missing its token", nil)
		return
	}
	address, err := s.Store.ConsumeLoginToken(r.Context(), TokenHash(token))
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"that link has already been used or has expired; ask for another", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "email", err, "could not sign you in just now")
		return
	}
	u, err := s.Store.UpsertEmailUser(r.Context(), address)
	if err != nil {
		s.fail(w, r, "email", err, "could not sign you in just now")
		return
	}
	if _, err := s.Auth.StartSession(r.Context(), w, u.ID, "email"); err != nil {
		s.fail(w, r, "email", err, "could not sign you in just now")
		return
	}
	http.Redirect(w, r, s.PublicBaseURL+safeNext(r.URL.Query().Get("next")), http.StatusFound)
}

func (s *Service) logout(w http.ResponseWriter, r *http.Request) {
	if err := s.Auth.EndSession(r.Context(), w, r); err != nil {
		s.fail(w, r, "logout", err, "could not sign you out just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]bool{"signed_out": true})
}

func (s *Service) deleteSession(w http.ResponseWriter, r *http.Request) {
	if err := s.Auth.EndSession(r.Context(), w, r); err != nil {
		s.fail(w, r, "logout", err, "could not sign you out just now")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// MeInput is the body of PATCH /v1/me. Only the pseudonym flag can be
// set: everything else about an account comes from Battle.net or from
// the sign-in itself.
type MeInput struct {
	Anonymize *bool `json:"anonymize"`
}

// patchMe sets the read-time pseudonym flag the privacy section
// promises to anyone who asks.
func (s *Service) patchMe(w http.ResponseWriter, r *http.Request) {
	var in MeInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxAuthBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return
	}
	if in.Anonymize == nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "nothing to change",
			map[string]string{"anonymize": "true or false"})
		return
	}
	if _, err := s.Store.SetAnonymize(r.Context(), ActorFrom(r.Context()).UserID, *in.Anonymize); err != nil {
		s.fail(w, r, "me", err, "could not change your account just now")
		return
	}
	s.me(w, r)
}

// Me is the /v1/me body: who you are, and what you own.
type Me struct {
	User       User        `json:"user"`
	Characters []Character `json:"characters"`
	Guilds     []Guild     `json:"guilds"`
}

func (s *Service) me(w http.ResponseWriter, r *http.Request) {
	a := ActorFrom(r.Context())
	u, err := s.Store.User(r.Context(), a.UserID)
	if err != nil {
		s.fail(w, r, "me", err, "could not load your account just now")
		return
	}
	chars, err := s.Store.Characters(r.Context(), u.ID)
	if err != nil {
		s.fail(w, r, "me", err, "could not load your account just now")
		return
	}
	guilds, err := s.Store.Guilds(r.Context(), u.ID)
	if err != nil {
		s.fail(w, r, "me", err, "could not load your account just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, Me{User: u, Characters: chars, Guilds: guilds})
}

func (s *Service) pair(w http.ResponseWriter, r *http.Request) {
	code, err := s.Store.CreatePairingCode(r.Context(), ActorFrom(r.Context()).UserID, PairingCodeTTL)
	if err != nil {
		s.fail(w, r, "pair", err, "could not make a pairing code just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{
		"code": code, "expires_in": int(PairingCodeTTL.Seconds()),
	})
}

type claimRequest struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

func (s *Service) claim(w http.ResponseWriter, r *http.Request) {
	var in claimRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxAuthBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with a code", nil)
		return
	}
	userID, err := s.Store.ClaimPairingCode(r.Context(), strings.ToLower(strings.TrimSpace(in.Code)))
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found",
			"that pairing code is not valid; make a new one on the site", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "claim", err, "could not pair that device just now")
		return
	}
	token, hash := NewDeviceToken()
	d, err := s.Store.CreateDevice(r.Context(), userID, trim(in.Name, 60), trim(in.Platform, 40), hash)
	if err != nil {
		s.fail(w, r, "claim", err, "could not pair that device just now")
		return
	}
	// The only time the token is ever readable.
	httpx.WriteOK(w, r, http.StatusCreated, map[string]string{"device_id": d.ID, "token": token})
}

func (s *Service) listDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := s.Store.Devices(r.Context(), ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, "devices", err, "could not list your devices just now")
		return
	}
	// The array is the body: the account page reads data as the list,
	// with no wrapper object around it.
	httpx.WriteOK(w, r, http.StatusOK, devices)
}

func (s *Service) revokeDevice(w http.ResponseWriter, r *http.Request) {
	ok, err := s.Store.RevokeDevice(r.Context(), ActorFrom(r.Context()).UserID, r.PathValue("id"))
	if err != nil {
		s.fail(w, r, "devices", err, "could not revoke that device just now")
		return
	}
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such device", nil)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// normalizeEmail parses and lowercases an address, rejecting anything
// net/mail will not read as a bare address.
func normalizeEmail(s string) (string, error) {
	addr, err := netmail.ParseAddress(strings.TrimSpace(s))
	if err != nil || addr.Name != "" || !strings.Contains(addr.Address, "@") {
		return "", errors.New("auth: invalid email")
	}
	return strings.ToLower(addr.Address), nil
}

// trim bounds a client-supplied label.
func trim(s string, max int) string {
	s = strings.TrimSpace(s)
	if len(s) > max {
		return s[:max]
	}
	return s
}
