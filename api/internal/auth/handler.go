package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	netmail "net/mail"
	"net/url"
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

// GuildCapabilities is the part of guilds.Store the account read needs. It is an interface
// because guilds imports this package.
type GuildCapabilities interface {
	CapabilitiesFor(ctx context.Context, guildIDs []int64, userID int64, moderator bool) (map[int64][]string, error)
}

// GuildAttendanceReader is the part of guilds.Store GET /v1/me needs for a character's raid
// attendance: given the character keys of each guild, an entry per key of a guild that counts
// raids and none for a guild still counting reports. It is an interface because guilds imports
// this package.
type GuildAttendanceReader interface {
	CharacterAttendance(ctx context.Context, keysByGuild map[int64][]string) (map[string]Attendance, error)
}

// FinishedSim is one server run that finished while its owner was away, as GET /v1/me
// reports it. FinishedAt is RFC3339 UTC.
type FinishedSim struct {
	SimID      string  `json:"sim_id"`
	Spec       string  `json:"spec"`
	Title      string  `json:"title"`
	DPS        float64 `json:"dps"`
	FinishedAt string  `json:"finished_at"`
}

// SimsView is the part of the sims store the account read needs. It is an interface
// because sims imports this package.
type SimsView interface {
	FinishedUnseen(ctx context.Context, userID int64, limit int) ([]FinishedSim, error)
	RunsThisMonth(ctx context.Context, userID int64) (int, error)
}

// DiscordBlock is the caller's saved Discord channel as GET /v1/me and the notifications
// routes answer it. It never carries the webhook URL, its token or its id.
type DiscordBlock struct {
	ChannelID      string        `json:"channel_id"`
	Name           *string       `json:"name"`
	Events         DiscordEvents `json:"events"`
	DisabledReason *string       `json:"disabled_reason"`
	LastError      *string       `json:"last_error"`
	Failures       int           `json:"failures"`
	// Source is "pasted" or "connected" (made by the one-click connect).
	Source string `json:"source"`
	// GuildID is the Discord server of the channel; null when unknown.
	GuildID *string `json:"guild_id"`
}

// DiscordEvents are the two things a Discord channel can be told about, and
// the modifier that lets private logs post.
type DiscordEvents struct {
	SimFinished   bool `json:"sim_finished"`
	UploadParsed  bool `json:"upload_parsed"`
	UploadPrivate bool `json:"upload_private"`
}

// finishedSimsLimit caps the "done while away" list on GET /v1/me.
const finishedSimsLimit = 3

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
	Entitlements *entitlements.Store
	// GuildCapabilities answers each guild entry's Can list on GET /v1/me. Nil is safe (every
	// list reads empty) for a harness that does not mount guilds; guilds.Store satisfies it.
	GuildCapabilities GuildCapabilities
	// GuildAttendance fills each guilded character's attendance on GET /v1/me. Nil is safe (no
	// attendance is shown); guilds.Store satisfies it.
	GuildAttendance GuildAttendanceReader
	// Sims answers GET /v1/me's finished_sims and server_sims. Nil is safe (the list reads
	// empty and server_sims is omitted) for a harness that does not mount sims.
	Sims SimsView
	// Notifications answers GET /v1/me's notifications block. Nil means the feature is off
	// (no encryption key) and the block is omitted.
	Notifications NotificationsView
	// DiscordLinks answers GET /v1/me's Discord link fields. Nil means the bot is off.
	DiscordLinks DiscordLinkView
	// Seats answers GET /v1/me's seat, seat_offer and ended_seat. Nil means they are left out.
	Seats SeatFactsSource
	// Standing answers each guilds[] entry's identity and step facts on GET /v1/me. Nil means
	// they are left out (the older shape).
	Standing GuildStandingSource
	// Importer runs the Battle.net character/guild import right after a
	// successful Battle.net sign-in. Nil means the feature is off (no
	// Battle.net credentials configured, or a test harness that does
	// not exercise it): login behaves exactly as before.
	Importer      Importer
	PublicBaseURL string
	APIBaseURL    string
	Log           *slog.Logger
	// DevLogin mounts POST /v1/dev/login (devlogin.go). Only the local walkthrough stack
	// sets it, and config.Load refuses FS_DEV_LOGIN off a loopback listener.
	DevLogin bool
	// AppSignIn mounts the native app sign-in (app_auth.go). Only config.AppSignIn sets it,
	// and it needs BNet as well.
	AppSignIn bool
}

// Mount registers the sign-in, device, and account routes. The Battle.net
// pair is registered only when the service is configured for it, so an
// unconfigured deployment answers 404 there rather than redirecting into
// a broken flow.
func Mount(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	if s.BNet != nil {
		mux.HandleFunc("GET /v1/auth/battlenet/start", s.bnetStart)
		mux.HandleFunc("GET /v1/auth/battlenet/callback", s.bnetCallback)
		if s.AppSignIn {
			mountApp(mux, s, trustedProxyHops)
		}
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
	mux.Handle("POST /v1/devices/connect", claimLimit(http.HandlerFunc(s.connectStart)))
	// The poll is the companion's, with the unguessable id as its secret, so
	// it is not rate limited; an unknown id answers 404 in one shape.
	mux.HandleFunc("GET /v1/devices/connect/lookup", RequireSession(s.connectLookup))
	mux.HandleFunc("GET /v1/devices/connect/{connect_id}", s.connectPoll)
	mux.HandleFunc("POST /v1/devices/connect/approve", RequireSession(s.connectApprove))
	mux.HandleFunc("POST /v1/devices/connect/deny", RequireSession(s.connectDeny))
	mux.HandleFunc("GET /v1/devices", RequireSession(s.listDevices))
	// The literal wins over {id}: a companion forgets itself with its own token.
	mux.HandleFunc("DELETE /v1/devices/self", RequireDevice(s.revokeSelf))
	mux.HandleFunc("DELETE /v1/devices/{id}", RequireSession(s.revokeDevice))
	if s.DevLogin {
		mux.HandleFunc("POST /v1/dev/login", s.devLogin)
	}
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
	httpx.CachePrivate(w)
	s.sendToBattleNet(w, r, NewSessionID(), safeNext(r.URL.Query().Get("next")))
}

// sendToBattleNet is the start of a web Battle.net sign-in: it binds state to this browser
// and sends the browser to Battle.net.
func (s *Service) sendToBattleNet(w http.ResponseWriter, r *http.Request, state, next string) {
	s.setBnetState(w, state, next)
	http.Redirect(w, r, s.BNet.AuthURL(state), http.StatusFound)
}

// setBnetState binds every Battle.net sign-in, the site's and the app's, to this browser in
// the state cookie, with next riding beside it.
func (s *Service) setBnetState(w http.ResponseWriter, state, next string) {
	http.SetCookie(w, &http.Cookie{
		Name: bnetStateCookie, Value: encodeState(state, next), Path: "/",
		Domain: s.Auth.CookieDomain, HttpOnly: true, Secure: s.Auth.Secure,
		SameSite: http.SameSiteLaxMode, MaxAge: int(bnetStateTTL.Seconds()),
	})
}

// bnetCallback finishes every Battle.net sign-in. A web sign-in starts a session and goes
// to next; an app sign-in (next is appFlowNext) hands a one-time code to the app link.
func (s *Service) bnetCallback(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivate(w)
	state, next, ok := s.checkBnetState(w, r)
	if !ok {
		return
	}
	bu, u, ok := s.identifyBnet(w, r)
	if !ok {
		return
	}
	if next == appFlowNext {
		s.finishAppSignIn(w, r, state, u.ID, bu.AccessToken)
		return
	}
	if _, err := s.Auth.StartSession(r.Context(), w, u.ID, "battlenet"); err != nil {
		s.fail(w, r, "battlenet", err, "could not sign you in just now")
		return
	}
	s.runImport(r, u.ID, bu.AccessToken)
	http.Redirect(w, r, s.PublicBaseURL+safeNext(next), http.StatusFound)
}

// checkBnetState reads and clears the state cookie and checks Battle.net echoed its state.
// It answers the request itself when it reports false.
func (s *Service) checkBnetState(w http.ResponseWriter, r *http.Request) (state, next string, ok bool) {
	c, err := r.Cookie(bnetStateCookie)
	if err != nil || c.Value == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "this sign-in has expired; start again", nil)
		return "", "", false
	}
	state, next = decodeState(c.Value)
	http.SetCookie(w, &http.Cookie{Name: bnetStateCookie, Value: "", Path: "/",
		Domain: s.Auth.CookieDomain, HttpOnly: true, Secure: s.Auth.Secure, MaxAge: -1})
	if state == "" || !SameToken(state, r.URL.Query().Get("state")) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "this sign-in could not be verified; start again", nil)
		return "", "", false
	}
	return state, next, true
}

// identifyBnet spends Battle.net's authorization code and upserts the account behind it.
// It answers the request itself when it reports false.
func (s *Service) identifyBnet(w http.ResponseWriter, r *http.Request) (BnetUser, User, bool) {
	code := r.URL.Query().Get("code")
	if code == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "Battle.net sent no authorization code", nil)
		return BnetUser{}, User{}, false
	}
	bu, err := s.BNet.Identify(r.Context(), code)
	if err != nil {
		s.logger().Error("auth", "id", httpx.RequestIDFrom(r.Context()), "op", "battlenet", "err", err)
		httpx.WriteError(w, r, http.StatusBadGateway, "upstream", "Battle.net could not confirm that sign-in", nil)
		return BnetUser{}, User{}, false
	}
	u, err := s.Store.UpsertBnetUser(r.Context(), bu.Sub, bu.Battletag)
	if err != nil {
		s.fail(w, r, "battlenet", err, "could not sign you in just now")
		return BnetUser{}, User{}, false
	}
	return bu, u, true
}

// runImport calls the Battle.net importer, bounded by importBudget, and
// logs the outcome. It never fails the login (spec §4.1): an error is
// logged at WARN and the caller redirects exactly as it would have with
// no importer at all.
func (s *Service) runImport(r *http.Request, userID int64, accessToken string) {
	if s.Importer == nil || accessToken == "" {
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), importBudget)
	defer cancel()
	summary, err := s.Importer.ImportAccount(ctx, userID, accessToken)
	id := httpx.RequestIDFrom(r.Context())
	if err != nil {
		s.logger().Warn("auth", "id", id, "op", "battlenet_import", "user_id", userID, "err", err)
		return
	}
	s.logger().Info("auth", "id", id, "op", "battlenet_import", "user_id", userID,
		"regions", summary.Regions, "characters", summary.Characters,
		"guilds", summary.Guilds, "skipped", summary.Skipped, "unavailable", summary.Unavailable)
}

type emailRequest struct {
	Next  string `json:"next"`
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
	// next rides in the link, ahead of the token, and is checked again by the
	// callback; only a given one is carried so plain sign-ins read as before.
	link := s.APIBaseURL + "/v1/auth/email/callback?"
	if in.Next != "" {
		link += "next=" + url.QueryEscape(safeNext(in.Next)) + "&"
	}
	link += "token=" + token
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
	httpx.CachePrivate(w)
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
	if err := s.endSignIn(w, r); err != nil {
		s.fail(w, r, "logout", err, "could not sign you out just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]bool{"signed_out": true})
}

// endSignIn is every generic sign-out: a browser's session ends, and an app calling it
// revokes its own device rather than hearing success while its token stays live.
func (s *Service) endSignIn(w http.ResponseWriter, r *http.Request) error {
	if a := ActorFrom(r.Context()); a.IsApp() {
		if _, err := s.Store.RevokeDeviceByID(r.Context(), a.DeviceID); err != nil {
			return err
		}
	}
	return s.Auth.EndSession(r.Context(), w, r)
}

func (s *Service) deleteSession(w http.ResponseWriter, r *http.Request) {
	if err := s.endSignIn(w, r); err != nil {
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
	// MainCharacterKey chooses the account's main; must be one of its own
	// characters. Everything else the account owns is an alt.
	MainCharacterKey *string `json:"main_character_key"`
}

// patchMe sets the read-time pseudonym flag the privacy section
// promises to anyone who asks.
func (s *Service) patchMe(w http.ResponseWriter, r *http.Request) {
	var in MeInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxAuthBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return
	}
	if in.Anonymize == nil && in.MainCharacterKey == nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "nothing to change",
			map[string]string{"anonymize": "true or false", "main_character_key": "one of your characters"})
		return
	}
	userID := ActorFrom(r.Context()).UserID
	if in.Anonymize != nil {
		if _, err := s.Store.SetAnonymize(r.Context(), userID, *in.Anonymize); err != nil {
			s.fail(w, r, "me", err, "could not change your account just now")
			return
		}
	}
	if in.MainCharacterKey != nil {
		if _, err := s.Store.SetMainCharacter(r.Context(), userID, *in.MainCharacterKey); err != nil {
			if errors.Is(err, ErrNotYourCharacter) {
				httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not one of your characters",
					map[string]string{"main_character_key": "one of your characters"})
				return
			}
			s.fail(w, r, "me", err, "could not change your account just now")
			return
		}
	}
	s.me(w, r)
}

// ServerSimsView is the caller's server-run account row.
type ServerSimsView struct {
	RunsThisMonth int `json:"runs_this_month"`
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
	LogCoach      bool         `json:"log_coach"`
	SupporterMark bool         `json:"supporter_mark"`
	Billing       *BillingView `json:"billing"`
	// PersonalPortal is true when the account's personal row carries a Stripe subscription
	// id, whatever its source or status: a person who once paid and was then granted still
	// has a Stripe history for the billing portal to show.
	PersonalPortal bool `json:"personal_portal"`
}

func (s *Service) me(w http.ResponseWriter, r *http.Request) {
	a := ActorFrom(r.Context())
	me, err := s.loadMe(r.Context(), a.UserID, a.IsModerator())
	if err != nil {
		s.fail(w, r, "me", err, "could not load your account just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, me)
}

// loadMe assembles the account read GET /v1/me answers and the app exchange returns.
// moderator is whether this caller may use the account's moderator powers, which an app
// token never may, so each guild's Can list only offers what the caller can do.
func (s *Service) loadMe(ctx context.Context, userID int64, moderator bool) (Me, error) {
	u, err := s.Store.User(ctx, userID)
	if err != nil {
		return Me{}, err
	}
	chars, err := s.Store.Characters(ctx, u.ID)
	if err != nil {
		return Me{}, err
	}
	if chars, err = s.attachCharacterAttendance(ctx, chars); err != nil {
		return Me{}, err
	}
	guilds, err := s.Store.Guilds(ctx, u.ID)
	if err != nil {
		return Me{}, err
	}
	facts := s.seatFacts(ctx, u.ID)
	if guilds, err = s.attachGuildPlans(ctx, guilds, u.ID, facts.HeldGuildIDs); err != nil {
		return Me{}, err
	}
	if guilds, err = s.attachGuildCapabilities(ctx, guilds, u.ID, moderator); err != nil {
		return Me{}, err
	}
	guilds = s.attachStanding(ctx, u.ID, guilds)
	ev, err := s.buildEntitlementsView(ctx, u.ID)
	if err != nil {
		return Me{}, err
	}
	finished, serverSims, err := s.buildSimsView(ctx, u.ID, ev.ServerSims)
	if err != nil {
		return Me{}, err
	}
	notifications, err := s.buildNotifications(ctx, u.ID, ev.Notifications)
	if err != nil {
		return Me{}, err
	}
	link, err := s.buildDiscordLink(ctx, u.ID)
	if err != nil {
		return Me{}, err
	}
	return Me{
		User: u, Characters: chars, Guilds: guilds, Entitlements: ev,
		BnetImportedAt: rfc3339Ptr(u.BnetImportedAt), MainCharacterKey: u.MainCharacterKey,
		FinishedSims: finished, ServerSims: serverSims, Notifications: notifications,
		DiscordLinkAvailable: link.available, DiscordLink: link.block, DiscordLinkSuggested: link.suggested,
		SeatViews: seatViewsOf(facts, ev.Billing),
	}, nil
}

// buildSimsView resolves the sims half of /v1/me. The finished list is only read for a
// caller who holds server_sims (a lapsed payer's unseen rows stay unseen, unannounced); the
// monthly count is a true fact for everyone. Without a sims store it is an empty list and nil.
func (s *Service) buildSimsView(ctx context.Context, userID int64, entitled bool) ([]FinishedSim, *ServerSimsView, error) {
	finished := []FinishedSim{}
	if s.Sims == nil {
		return finished, nil, nil
	}
	if entitled {
		rows, err := s.Sims.FinishedUnseen(ctx, userID, finishedSimsLimit)
		if err != nil {
			return nil, nil, fmt.Errorf("auth: me: finished sims: %w", err)
		}
		finished = append(finished, rows...)
	}
	n, err := s.Sims.RunsThisMonth(ctx, userID)
	if err != nil {
		return nil, nil, fmt.Errorf("auth: me: runs this month: %w", err)
	}
	return finished, &ServerSimsView{RunsThisMonth: n}, nil
}

// entitlementFeatures pairs every gated Feature with the EntitlementsView
// field it sets, so buildEntitlementsView is one loop; Premium resolves once
// and fans out to every standard feature.
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
	{entitlements.FeatureLogCoach, func(v *EntitlementsView, ok bool) { v.LogCoach = ok }},
}

func (s *Service) buildEntitlementsView(ctx context.Context, userID int64) (EntitlementsView, error) {
	var ev EntitlementsView
	if s.Entitlements == nil {
		return ev, nil
	}
	supporter, err := s.Entitlements.IsSupporter(ctx, userID)
	if err != nil {
		return EntitlementsView{}, fmt.Errorf("auth: me: supporter: %w", err)
	}
	ev.SupporterMark = supporter
	for _, f := range entitlementFeatures {
		ok := supporter
		if !f.feature.Standard() {
			ok, _, err = s.Entitlements.Can(ctx, userID, f.feature)
			if err != nil {
				return EntitlementsView{}, fmt.Errorf("auth: me: entitlements: %w", err)
			}
		}
		f.set(&ev, ok)
	}
	b, err := s.Entitlements.PersonalBilling(ctx, userID)
	if err != nil {
		return EntitlementsView{}, fmt.Errorf("auth: me: billing: %w", err)
	}
	if b != nil {
		ev.Billing = personalBillingView(b)
		ev.PersonalPortal = b.StripeSubscriptionID != ""
	}
	return ev, nil
}

// attachGuildCapabilities fills Guild.Can for every guild the account belongs to, from the
// one rule guilds.Capabilities holds. Every list is non-nil, so the JSON is [] and not null.
func (s *Service) attachGuildCapabilities(ctx context.Context, guilds []Guild, userID int64, moderator bool) ([]Guild, error) {
	out := make([]Guild, len(guilds))
	copy(out, guilds)
	var lists map[int64][]string
	if s.GuildCapabilities != nil && len(guilds) > 0 {
		ids := make([]int64, len(guilds))
		for i, g := range guilds {
			ids[i] = g.ID
		}
		var err error
		lists, err = s.GuildCapabilities.CapabilitiesFor(ctx, ids, userID, moderator)
		if err != nil {
			return nil, fmt.Errorf("auth: me: guild capabilities: %w", err)
		}
	}
	for i, g := range out {
		out[i].Can = []string{}
		if can := lists[g.ID]; can != nil {
			out[i].Can = can
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
	// Who it connected as, so the companion can say so; a failed read only
	// loses the label, never the device that was just made.
	label, err := s.Store.AccountLabel(r.Context(), userID)
	if err != nil {
		s.logger().Warn("auth", "id", httpx.RequestIDFrom(r.Context()), "op", "claim label", "err", err)
	}
	httpx.WriteOK(w, r, http.StatusCreated, map[string]string{"device_id": d.ID, "token": token, "account_label": label})
}

func (s *Service) listDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := s.Store.Devices(r.Context(), ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, "devices", err, "could not list your devices just now")
		return
	}
	httpx.CachePrivate(w)
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

// revokeSelf is the companion's Forget: the device the bearer token
// belongs to revokes itself. A token that is already dead never gets here;
// the middleware answers it 401, which the companion counts as done.
func (s *Service) revokeSelf(w http.ResponseWriter, r *http.Request) {
	if _, err := s.Store.RevokeDeviceByID(r.Context(), ActorFrom(r.Context()).DeviceID); err != nil {
		s.fail(w, r, "devices", err, "could not remove this device just now")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// attachCharacterAttendance fills CharacterGuild.Attendance for every verified guilded
// character whose guild counts raids.
func (s *Service) attachCharacterAttendance(ctx context.Context, chars []Character) ([]Character, error) {
	if s.GuildAttendance == nil {
		return chars, nil
	}
	keysByGuild := map[int64][]string{}
	for _, c := range chars {
		if c.Guild != nil && c.Guild.Verified {
			keysByGuild[c.Guild.ID] = append(keysByGuild[c.Guild.ID], c.Key)
		}
	}
	if len(keysByGuild) == 0 {
		return chars, nil
	}
	found, err := s.GuildAttendance.CharacterAttendance(ctx, keysByGuild)
	if err != nil {
		return nil, fmt.Errorf("auth: character attendance: %w", err)
	}
	for i := range chars {
		a, ok := found[chars[i].Key]
		if ok && chars[i].Guild != nil && chars[i].Guild.Verified {
			att := a
			chars[i].Guild.Attendance = &att
		}
	}
	return chars, nil
}
