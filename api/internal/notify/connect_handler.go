package notify

import (
	"context"
	"errors"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const (
	connectPath  = "/v1/notifications/discord/connect"
	callbackPath = "/v1/notifications/discord/callback"

	// discordStateCookie carries the OAuth state from connect to callback.
	discordStateCookie = "fs_discord_state"
	discordStateTTL    = 600 // seconds: ten minutes

	// The three outcomes the account page reads from ?discord=.
	outcomeConnected = "connected"
	outcomeDenied    = "denied"
	outcomeFailed    = "failed"

	oauthAccessDenied = "access_denied"
)

var errNoChannel = errors.New("notify: discord's webhook carried no channel id")

// CallbackURL is the redirect URI registered with the Discord application.
func CallbackURL(apiBaseURL string) string { return apiBaseURL + callbackPath }

// defaultEvents are the events a first connect starts with (sim on, log on,
// private off).
func defaultEvents() map[EventName]bool {
	return map[EventName]bool{EventSimFinished: true, EventUploadParsed: true, EventUploadPrivate: false}
}

// ConnectEnabled implements auth.NotificationsView.
func (s *Service) ConnectEnabled() bool { return s.OAuth != nil }

func (s *Service) mountConnect(mux *http.ServeMux) {
	if s.OAuth == nil {
		return
	}
	mux.HandleFunc("GET "+connectPath, auth.RequireSession(s.connect))
	mux.HandleFunc("GET "+callbackPath, auth.RequireSession(s.callback))
}

// accountRedirect sends the browser back to the account page with the outcome.
func (s *Service) accountRedirect(w http.ResponseWriter, r *http.Request, outcome string) {
	http.Redirect(w, r, s.PublicBaseURL+"/account?discord="+outcome+"#discord", http.StatusFound)
}

func (s *Service) stateCookie(value string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name: discordStateCookie, Value: value, Path: callbackPath,
		HttpOnly: true, Secure: s.CookieSecure, SameSite: http.SameSiteLaxMode, MaxAge: maxAge,
	}
}

func (s *Service) connect(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	allowed, _, err := s.Entitlements.Can(r.Context(), auth.ActorFrom(r.Context()).UserID, entitlements.FeatureNotifications)
	if err != nil {
		s.logger().Error("notify", "id", httpx.RequestIDFrom(r.Context()), "op", "connect premium", "err", err)
	}
	if err != nil || !allowed {
		s.accountRedirect(w, r, outcomeFailed)
		return
	}
	state := auth.NewSessionID()
	http.SetCookie(w, s.stateCookie(state, discordStateTTL))
	http.Redirect(w, r, s.OAuth.AuthURL(state), http.StatusFound)
}

func (s *Service) callback(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	cookie, cookieErr := r.Cookie(discordStateCookie)
	http.SetCookie(w, s.stateCookie("", -1))
	query := r.URL.Query()
	code := query.Get("code")
	if code == "" || query.Get("error") == oauthAccessDenied {
		s.accountRedirect(w, r, outcomeDenied)
		return
	}
	if cookieErr != nil || cookie.Value == "" || !auth.SameToken(cookie.Value, query.Get("state")) {
		s.accountRedirect(w, r, outcomeFailed)
		return
	}
	if err := s.storeConnected(r.Context(), auth.ActorFrom(r.Context()).UserID, code); err != nil {
		s.logger().Warn("notify", "id", httpx.RequestIDFrom(r.Context()), "op", "connect", "err", err)
		s.accountRedirect(w, r, outcomeFailed)
		return
	}
	s.accountRedirect(w, r, outcomeConnected)
}

// storeConnected exchanges the code, seals and saves the new webhook, then
// removes the previous connected webhook from Discord. The returned error
// never carries the code, the token or the URL.
func (s *Service) storeConnected(ctx context.Context, userID int64, code string) error {
	hook, err := s.OAuth.Exchange(ctx, code)
	if err != nil {
		return err
	}
	parsed, err := ParseWebhookURL(hook.URL)
	if err != nil {
		return err
	}
	if hook.ChannelID == "" {
		return errNoChannel
	}
	events := defaultEvents()
	if previous, ok, err := s.Store.Get(ctx, userID, KindDiscord); err != nil {
		return err
	} else if ok {
		events = previous.Events
	}
	oldURL := s.connectedURL(ctx, userID)
	box, err := s.Sealer.Seal(parsed.URL())
	if err != nil {
		return err
	}
	if err := s.Store.Save(ctx, userID, KindDiscord, hook.ChannelID, optionalText(hook.GuildID), nil, box, events, SourceConnected); err != nil {
		return err
	}
	if oldURL != "" && oldURL != parsed.URL() {
		s.deleteAtDiscord(ctx, oldURL, "replace")
	}
	return nil
}

func optionalText(text string) *string {
	if text == "" {
		return nil
	}
	return &text
}
