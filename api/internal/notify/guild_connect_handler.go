package notify

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const (
	guildConnectPath = "/v1/guilds/{id}/notifications/discord/connect"
	// guildConnectStateTTL is how long a guild connect's state stays good.
	guildConnectStateTTL = 10 * time.Minute
	// guildSettingsReturn follows the guild page's path: the settings page and its anchor.
	guildSettingsReturn = "/settings?discord="
	discordAnchor       = "#discord"
)

// guildConnect starts a guild connect: a state row bound to this officer and guild, then
// Discord's authorize page with the account flow's redirect URI.
func (s *Service) guildConnect(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	guildID, _, ok := s.guildOfficer(w, r, true)
	if !ok {
		return
	}
	state := auth.NewSessionID()
	userID := auth.ActorFrom(r.Context()).UserID
	if err := s.GuildStore.CreateConnectState(r.Context(), state, userID, guildID, guildConnectStateTTL); err != nil {
		s.fail(w, r, "guild connect state", err)
		return
	}
	http.Redirect(w, r, s.OAuth.AuthURL(state), http.StatusFound)
}

// callbackEntry is the one redirect URI both connects share. A state matching a guild
// connect row (taken, so used once) is the guild's; anything else is the account's flow.
func (s *Service) callbackEntry(w http.ResponseWriter, r *http.Request) {
	if s.GuildConnectEnabled() {
		cs, ok, err := s.GuildStore.TakeConnectState(r.Context(), r.URL.Query().Get("state"))
		if err != nil {
			httpx.CachePrivateNoStore(w)
			s.logger().Error("notify", "id", httpx.RequestIDFrom(r.Context()), "op", "guild connect state", "err", err)
			s.accountRedirect(w, r, outcomeFailed)
			return
		}
		if ok {
			s.guildCallback(w, r, cs)
			return
		}
	}
	s.browserSession(s.callback)(w, r)
}

// guildCallback finishes a guild connect. It answers failed unless the session is the
// officer who started it, inside ten minutes, and still a verified officer of a guild with
// the plan and no contested claim; denied when the officer refused at Discord.
func (s *Service) guildCallback(w http.ResponseWriter, r *http.Request, cs ConnectState) {
	httpx.CachePrivateNoStore(w)
	standing, found, err := s.Guilds.GuildStanding(r.Context(), cs.GuildID, cs.UserID)
	if err != nil || !found {
		s.logger().Warn("notify", "id", httpx.RequestIDFrom(r.Context()), "op", "guild connect standing", "found", found, "err", err)
		s.accountRedirect(w, r, outcomeFailed)
		return
	}
	back := func(outcome string) { s.guildRedirect(w, r, standing.Path, outcome) }
	actor := auth.ActorFrom(r.Context())
	if !actor.Signed() || actor.IsDevice() || actor.UserID != cs.UserID || cs.Expired {
		back(outcomeFailed)
		return
	}
	query := r.URL.Query()
	code := query.Get("code")
	if code == "" || query.Get("error") == oauthAccessDenied {
		back(outcomeDenied)
		return
	}
	if !standing.Officer || !standing.Plan || standing.Frozen {
		back(outcomeFailed)
		return
	}
	save := func(ctx context.Context, hook OAuthWebhook) error {
		return s.saveGuildConnected(ctx, cs.GuildID, cs.UserID, hook)
	}
	if err := s.exchangeAndSave(r.Context(), code, save); err != nil {
		s.logger().Warn("notify", "id", httpx.RequestIDFrom(r.Context()), "op", "guild connect", "err", err)
		back(outcomeFailed)
		return
	}
	s.logger().Info("notify", "op", "guild connect", "guild_id", cs.GuildID, "user_id", cs.UserID)
	back(outcomeConnected)
}

// guildRedirect sends the browser back to the guild's settings page with the outcome.
func (s *Service) guildRedirect(w http.ResponseWriter, r *http.Request, guildPath, outcome string) {
	http.Redirect(w, r, s.PublicBaseURL+guildPath+guildSettingsReturn+outcome+discordAnchor, http.StatusFound)
}

// saveGuildConnected seals and saves hook as the guild's channel, then removes the
// guild's previous webhook from Discord.
func (s *Service) saveGuildConnected(ctx context.Context, guildID, userID int64, hook OAuthWebhook) error {
	parsed, box, err := s.sealHook(hook)
	if err != nil {
		return fmt.Errorf("notify: seal guild webhook %d: %w", guildID, err)
	}
	oldURL := s.guildURL(ctx, guildID)
	in := GuildSaveInput{
		GuildID: guildID, ChannelID: hook.ChannelID, DiscordGuildID: optionalText(hook.GuildID),
		TokenBox: box, ConnectedBy: userID,
	}
	if err := s.GuildStore.Save(ctx, in); err != nil {
		return fmt.Errorf("notify: save guild webhook %d: %w", guildID, err)
	}
	if oldURL != "" && oldURL != parsed.URL() {
		s.deleteAtDiscord(ctx, oldURL, "guild replace")
	}
	return nil
}
