package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	netmail "net/mail"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
	"github.com/jhunthrop/foreversixty/api/internal/textx"
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
	// maxDeviceName and maxDevicePlatform bound the two labels a
	// companion names itself with.
	maxDeviceName     = 60
	maxDevicePlatform = 40
)

// Service serves every sign-in route.
type Service struct {
	Store  *Store
	Auth   *Authenticator
	BNet   *BattleNet
	Mailer mail.Mailer
	// Entitlements answers GET /v1/me's Entitlements block and each
	// guild's Plan field. Nil is safe (every field reads as its zero
	// value, every Guild.Plan stays nil) for a test harness that does
	// not exercise billing.
	Entitlements  *entitlements.Store
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
	// A paired device holds a token for uploading, not for reading the
	// account behind it; the email and guild list stay with the browser.
	mux.HandleFunc("GET /v1/me", RequireSession(s.me))
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

// Me is the /v1/me body: who you are, what you own, and what you're
// entitled to (spec §1.4).
type Me struct {
	User         User             `json:"user"`
	Characters   []Character      `json:"characters"`
	Guilds       []Guild          `json:"guilds"`
	Entitlements EntitlementsView `json:"entitlements"`
}

// EntitlementsView is every Can() feature pre-resolved for the caller,
// plus their own personal billing state.
type EntitlementsView struct {
	ServerSims    bool         `json:"server_sims"`
	Retention     bool         `json:"retention"`
	MultiCompare  bool         `json:"multi_compare"`
	History       bool         `json:"history"`
	Notifications bool         `json:"notifications"`
	OfficerViews  bool         `json:"officer_views"`
	RosterCheck   bool         `json:"roster_check"`
	SupporterMark bool         `json:"supporter_mark"`
	Billing       *BillingView `json:"billing"`
}

// BillingView is the caller's own personal premium subscription state,
// nil when they have none (spec §1.4).
type BillingView struct {
	Plan              string  `json:"plan"`
	Status            string  `json:"status"`
	CurrentPeriodEnd  *string `json:"current_period_end"`
	CancelAtPeriodEnd bool    `json:"cancel_at_period_end"`
}

// GuildBillingView is a guild's billing state, shown only to a verified
// officer/leader of that guild (spec §1.4; Ruling C).
type GuildBillingView struct {
	Status               string  `json:"status"`
	CurrentPeriodEnd     *string `json:"current_period_end"`
	CancelAtPeriodEnd    bool    `json:"cancel_at_period_end"`
	BilledBy             string  `json:"billed_by"`
	YouAreBillingContact bool    `json:"you_are_billing_contact"`
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
	guilds, err = s.attachGuildPlans(r.Context(), guilds, u.ID)
	if err != nil {
		s.fail(w, r, "me", err, "could not load your account just now")
		return
	}
	ev, err := s.buildEntitlementsView(r.Context(), u.ID)
	if err != nil {
		s.fail(w, r, "me", err, "could not load your account just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, Me{User: u, Characters: chars, Guilds: guilds, Entitlements: ev})
}

// entitlementFeatures pairs every gated Feature with the EntitlementsView
// field it sets, so buildEntitlementsView is one loop rather than seven
// near-identical Can calls.
var entitlementFeatures = []struct {
	feature entitlements.Feature
	set     func(*EntitlementsView, bool)
}{
	{entitlements.FeatureServerSims, func(v *EntitlementsView, ok bool) { v.ServerSims = ok }},
	{entitlements.FeatureRetention, func(v *EntitlementsView, ok bool) { v.Retention = ok }},
	{entitlements.FeatureMultiCompare, func(v *EntitlementsView, ok bool) { v.MultiCompare = ok }},
	{entitlements.FeatureHistory, func(v *EntitlementsView, ok bool) { v.History = ok }},
	{entitlements.FeatureNotifications, func(v *EntitlementsView, ok bool) { v.Notifications = ok }},
	{entitlements.FeatureOfficerViews, func(v *EntitlementsView, ok bool) { v.OfficerViews = ok }},
	{entitlements.FeatureRosterCheck, func(v *EntitlementsView, ok bool) { v.RosterCheck = ok }},
}

func (s *Service) buildEntitlementsView(ctx context.Context, userID int64) (EntitlementsView, error) {
	var ev EntitlementsView
	if s.Entitlements == nil {
		return ev, nil
	}
	for _, f := range entitlementFeatures {
		ok, _, err := s.Entitlements.Can(ctx, userID, f.feature)
		if err != nil {
			return EntitlementsView{}, fmt.Errorf("auth: me: entitlements: %w", err)
		}
		f.set(&ev, ok)
	}
	supporter, err := s.Entitlements.IsSupporter(ctx, userID)
	if err != nil {
		return EntitlementsView{}, fmt.Errorf("auth: me: supporter: %w", err)
	}
	ev.SupporterMark = supporter
	b, err := s.Entitlements.PersonalBilling(ctx, userID)
	if err != nil {
		return EntitlementsView{}, fmt.Errorf("auth: me: billing: %w", err)
	}
	if b != nil {
		ev.Billing = &BillingView{
			Plan: b.Plan, Status: b.Status, CancelAtPeriodEnd: b.CancelAtPeriodEnd,
			CurrentPeriodEnd: rfc3339Ptr(b.CurrentPeriodEnd),
		}
	}
	return ev, nil
}

// isActiveEntitlementStatus mirrors entitlements.Store's own
// activeStatusClause (unexported there): the three statuses that keep a
// guild's plan visible/active on this response (spec §1.3 rule 4).
func isActiveEntitlementStatus(status string) bool {
	return status == "active" || status == "trialing" || status == "past_due"
}

// attachGuildPlans fills Guild.Plan for every guild in guilds where
// userID is a verified officer/leader and that guild currently has an
// active plan (spec §1.4; Ruling C).
func (s *Service) attachGuildPlans(ctx context.Context, guilds []Guild, userID int64) ([]Guild, error) {
	if s.Entitlements == nil {
		return guilds, nil
	}
	out := make([]Guild, len(guilds))
	copy(out, guilds)
	for i, g := range out {
		if !g.Verified || (g.Rank != "officer" && g.Rank != "leader") {
			continue
		}
		gb, err := s.Entitlements.GuildBilling(ctx, g.ID)
		if err != nil {
			return nil, fmt.Errorf("auth: me: guild billing %d: %w", g.ID, err)
		}
		if gb == nil || !isActiveEntitlementStatus(gb.Status) {
			continue
		}
		billedBy, youAre := "", false
		switch {
		case gb.BillingUserID == nil:
			// A CLI grant, never billed through Stripe — no billing
			// contact to name.
		case *gb.BillingUserID == userID:
			billedBy, youAre = "you", true
		default:
			if bu, err := s.Store.User(ctx, *gb.BillingUserID); err == nil {
				billedBy = bu.PublicName()
			} else {
				billedBy = "a former member"
			}
		}
		out[i].Plan = &GuildBillingView{
			Status: gb.Status, CancelAtPeriodEnd: gb.CancelAtPeriodEnd,
			CurrentPeriodEnd: rfc3339Ptr(gb.CurrentPeriodEnd), BilledBy: billedBy, YouAreBillingContact: youAre,
		}
	}
	return out, nil
}

// rfc3339Ptr formats t as an RFC 3339 string pointer, nil in, nil out —
// EntitlementsView/GuildBillingView's CurrentPeriodEnd shape (spec §1.4:
// "null for a grant with no expiry").
func rfc3339Ptr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.UTC().Format(time.RFC3339)
	return &s
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
	d, err := s.Store.CreateDevice(r.Context(), userID,
		textx.Trim(in.Name, maxDeviceName), textx.Trim(in.Platform, maxDevicePlatform), hash)
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
