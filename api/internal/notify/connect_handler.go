package notify

import (
	"context"
	"errors"
	"net/http"
	"time"

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
	mux.HandleFunc("GET "+connectPath, s.browserSession(s.connect))
	mux.HandleFunc("GET "+callbackPath, s.callbackEntry)
}

// accountRedirect sends the browser back to the account page with the outcome.
func (s *Service) accountRedirect(w http.ResponseWriter, r *http.Request, outcome string) {
	http.Redirect(w, r, s.PublicBaseURL+"/account?discord="+outcome+"#discord", http.StatusFound)
}

// browserSession wraps a browser-navigation handler: without a browser
// session it redirects to the account page's failed outcome instead of
// answering 401 JSON the browser would show as raw text.
func (s *Service) browserSession(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if actor := auth.ActorFrom(r.Context()); !actor.Signed() || actor.IsDevice() {
			httpx.CachePrivateNoStore(w)
			s.accountRedirect(w, r, outcomeFailed)
			return
		}
		h(w, r)
	}
}

// entitledToConnect reports whether the member may connect; an entitlement
// error is logged and reads as no.
func (s *Service) entitledToConnect(r *http.Request, userID int64) bool {
	allowed, _, err := s.Entitlements.Can(r.Context(), userID, entitlements.FeatureNotifications)
	if err != nil {
		s.logger().Error("notify", "id", httpx.RequestIDFrom(r.Context()), "op", "connect premium", "err", err)
		return false
	}
	return allowed
}

func (s *Service) stateCookie(value string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name: discordStateCookie, Value: value, Path: callbackPath,
		HttpOnly: true, Secure: s.CookieSecure, SameSite: http.SameSiteLaxMode, MaxAge: maxAge,
	}
}

func (s *Service) connect(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	if !s.entitledToConnect(r, auth.ActorFrom(r.Context()).UserID) {
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
	userID := auth.ActorFrom(r.Context()).UserID
	if !s.entitledToConnect(r, userID) {
		s.accountRedirect(w, r, outcomeFailed)
		return
	}
	save := func(ctx context.Context, hook OAuthWebhook) error { return s.saveConnected(ctx, userID, hook) }
	if err := s.exchangeAndSave(r.Context(), code, save); err != nil {
		s.logger().Warn("notify", "id", httpx.RequestIDFrom(r.Context()), "op", "connect", "err", err)
		s.accountRedirect(w, r, outcomeFailed)
		return
	}
	s.accountRedirect(w, r, outcomeConnected)
}

// exchangeAndSave exchanges the code and hands the new webhook to save. When the save
// fails after Discord created the webhook, the orphan is deleted. The returned error never
// carries the code, the token or the URL.
func (s *Service) exchangeAndSave(ctx context.Context, code string, save func(context.Context, OAuthWebhook) error) error {
	hook, err := s.OAuth.Exchange(ctx, code)
	if err != nil {
		return err
	}
	if err := save(ctx, hook); err != nil {
		s.discardOrphan(ctx, hook)
		return err
	}
	return nil
}

// sealHook checks the exchanged webhook (a canonical URL and a channel) and seals its URL.
func (s *Service) sealHook(hook OAuthWebhook) (ParsedURL, []byte, error) {
	parsed, err := ParseWebhookURL(hook.URL)
	if err != nil {
		return ParsedURL{}, nil, err
	}
	if hook.ChannelID == "" {
		return ParsedURL{}, nil, errNoChannel
	}
	box, err := s.Sealer.Seal(parsed.URL())
	if err != nil {
		return ParsedURL{}, nil, err
	}
	return parsed, box, nil
}

// saveConnected seals and saves hook, then removes the member's previous
// connected webhook from Discord.
func (s *Service) saveConnected(ctx context.Context, userID int64, hook OAuthWebhook) error {
	parsed, box, err := s.sealHook(hook)
	if err != nil {
		return err
	}
	events := defaultEvents()
	previous, had, err := s.Store.Get(ctx, userID, KindDiscord)
	if err != nil {
		return err
	}
	if had {
		events = previous.Events
	}
	oldURL := s.connectedURLOf(ctx, previous, had)
	in := SaveInput{
		UserID: userID, Kind: KindDiscord, ChannelID: hook.ChannelID, GuildID: optionalText(hook.GuildID),
		TokenBox: box, Events: events, Source: SourceConnected,
	}
	if err := s.Store.Save(ctx, in); err != nil {
		return err
	}
	if oldURL != "" && oldURL != parsed.URL() {
		s.deleteAtDiscord(ctx, oldURL, "replace")
	}
	return nil
}

// orphanDeleteTimeout bounds the cleanup of a webhook Discord created but we
// could not keep.
const orphanDeleteTimeout = 5 * time.Second

// discardOrphan deletes the webhook Discord just created, best effort, on a
// fresh context so a cancelled request does not strand it. With no usable URL
// it only logs that it skipped.
func (s *Service) discardOrphan(ctx context.Context, hook OAuthWebhook) {
	url := orphanURL(hook)
	if url == "" {
		s.logger().Warn("notify", "op", "discord delete orphan", "outcome", "skipped")
		return
	}
	fresh, cancel := context.WithTimeout(context.WithoutCancel(ctx), orphanDeleteTimeout)
	defer cancel()
	s.deleteAtDiscord(fresh, url, "orphan")
}

// orphanURL is the webhook URL to delete: hook.URL when it parses, else the
// canonical URL built from a well-formed id and token, else "".
func orphanURL(hook OAuthWebhook) string {
	if parsed, err := ParseWebhookURL(hook.URL); err == nil {
		return parsed.URL()
	}
	built := ParsedURL{WebhookID: hook.ID, Token: hook.Token}
	if parsed, err := ParseWebhookURL(built.URL()); err == nil {
		return parsed.URL()
	}
	return ""
}

func optionalText(text string) *string {
	if text == "" {
		return nil
	}
	return &text
}
