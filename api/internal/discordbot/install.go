package discordbot

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	botPath         = "/v1/guilds/{id}/discord/bot"
	botInstallPath  = botPath + "/install"
	botChannelsPath = botPath + "/channels"
	botChannelPath  = botPath + "/channel"
	botTestPath     = botPath + "/test"

	installsPerMinute = 20
	// layoutReadsPerMinute caps an account's channel list and channel save together: each
	// one reads the server's layout through the shared bot token.
	layoutReadsPerMinute = 20
	botTestsPerMinute    = 5
	// refusalTTL is how long an officer's "server taken" refusal waits for the next page load.
	refusalTTL = 10 * time.Minute

	botOutcomeParam = "/settings?discord_bot="
	discordAnchor   = "#discord"

	installInstalled = "installed"
	installDenied    = "denied"
	installFailed    = "failed"
	installTaken     = "taken"

	noInstallMessage = "the bot is not installed"
)

// mountInstall registers the bot's guild routes; Mount calls it.
func (s *Service) mountInstall(mux *http.ServeMux) {
	install := httpx.RateLimitByKey(installsPerMinute, time.Minute, auth.AccountRateLimitKey)
	test := httpx.RateLimitByKey(botTestsPerMinute, time.Minute, auth.AccountRateLimitKey)
	mux.Handle("GET "+botInstallPath, install(auth.RequireSession(s.installStart)))
	layout := httpx.RateLimitByKey(layoutReadsPerMinute, time.Minute, auth.AccountRateLimitKey)
	mux.Handle("GET "+botChannelsPath, layout(auth.RequireSession(s.channelsList)))
	mux.Handle("PUT "+botChannelPath, layout(auth.RequireSession(s.channelSave)))
	mux.Handle("POST "+botTestPath, test(auth.RequireSession(s.botTest)))
	// Remove is never behind the plan: stopping posts must always work.
	mux.HandleFunc("DELETE "+botPath, auth.RequireSession(s.botRemove))
}

// officer is notify's gate for the bot's routes.
func (s *Service) officer(w http.ResponseWriter, r *http.Request, needPlan bool) (int64, notify.GuildStanding, bool) {
	return notify.GuildOfficerGate(w, r, s.d.Guilds, needPlan, func(op string, err error) { s.logFailure(r, op, err) })
}

func (s *Service) logFailure(r *http.Request, op string, err error) {
	s.log.Error("discordbot", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
}

func (s *Service) internalError(w http.ResponseWriter, r *http.Request, op string, err error) {
	s.logFailure(r, op, err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not save that just now", nil)
}

// BotAvailable implements guilds.DiscordBot: the deployment runs the bot.
func (s *Service) BotAvailable() bool {
	return s != nil && s.d.OAuth != nil && s.d.Notify != nil && s.d.Guilds != nil
}

// ActiveInstall is the guild's install when the bot can post: it has a channel, Discord has
// not said it was removed, and the channel is not unusable.
func (s *Service) ActiveInstall(ctx context.Context, guildID int64) (Install, bool, error) {
	in, found, err := s.d.Store.InstallFor(ctx, guildID)
	if err != nil || !found || !in.canPost() {
		return Install{}, false, err
	}
	return in, true, nil
}

// canPost is ActiveInstall's rule for an install that exists.
func (in Install) canPost() bool {
	return in.ChannelID != nil && !in.RemovedInDiscord && in.ChannelUnusable == nil
}

// Destination implements guilds.DiscordBot: whether the bot posts the guild's raids
// (ActiveInstall's rule), the install's remembered boxes whenever there is an install (so a
// save that carries one box keeps the other), and the channel's name when active.
func (s *Service) Destination(ctx context.Context, guildID int64) (guilds.BotDestination, error) {
	in, found, err := s.d.Store.InstallFor(ctx, guildID)
	if err != nil || !found {
		return guilds.BotDestination{}, err
	}
	dest := guilds.BotDestination{Active: in.canPost(), PostDefault: in.PostDefault, RemindDefault: in.RemindDefault}
	if dest.Active && in.ChannelName != nil {
		dest.ChannelName = *in.ChannelName
	}
	return dest, nil
}

// SetDiscordDefaults implements guilds.DiscordBot: it remembers the raid form's two boxes
// on the guild's install; a guild with no install is left as it is.
func (s *Service) SetDiscordDefaults(ctx context.Context, guildID int64, post, remind bool) error {
	return s.d.Store.SetDiscordDefaults(ctx, guildID, post, remind)
}

// SettingsBot implements guilds.DiscordBot.
func (s *Service) SettingsBot(ctx context.Context, guildID, viewerID int64) (*guilds.SettingsBot, *guilds.BotRefusal, error) {
	bot, err := s.installView(ctx, guildID)
	if err != nil {
		return nil, nil, err
	}
	// The refusal is read last, so a failure above never consumes it.
	held, err := s.d.Store.TakeRefusal(ctx, viewerID, guildID)
	if err != nil {
		return nil, nil, err
	}
	if held == nil {
		return bot, nil, nil
	}
	return bot, &guilds.BotRefusal{GuildName: *held}, nil
}

// installView is the guild's install as settings shows it; nil when there is none.
func (s *Service) installView(ctx context.Context, guildID int64) (*guilds.SettingsBot, error) {
	in, found, err := s.d.Store.InstallFor(ctx, guildID)
	if err != nil || !found {
		return nil, err
	}
	linked, members, err := s.d.Store.LinkedMembers(ctx, guildID)
	if err != nil {
		return nil, err
	}
	return &guilds.SettingsBot{
		ServerID: in.ServerID, ServerName: in.ServerName, ChannelID: in.ChannelID, ChannelName: in.ChannelName,
		LinkedMembers: linked, Members: members, Failures: in.Failures, LastError: in.LastError,
		RemovedInDiscord: in.RemovedInDiscord, ChannelUnusable: unusableView(in.ChannelUnusable),
	}, nil
}

// installStart sends a plan-holding officer to Discord's install page.
func (s *Service) installStart(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	guildID, _, ok := s.officer(w, r, true)
	if !ok {
		return
	}
	state, err := s.d.Notify.StartState(r.Context(), notify.StateKindBotInstall, auth.ActorFrom(r.Context()).UserID, &guildID)
	if err != nil {
		s.internalError(w, r, "install state", err)
		return
	}
	http.Redirect(w, r, s.d.OAuth.InstallURL(state), http.StatusFound)
}

// installCallback finishes an install at the shared callback. The code never leaves the
// exchange; only the server's id and name are kept.
func (s *Service) installCallback(w http.ResponseWriter, r *http.Request, cs notify.ConnectState) {
	httpx.CachePrivateNoStore(w)
	standing, found, err := s.d.Guilds.GuildStanding(r.Context(), cs.GuildID, cs.UserID)
	if err != nil {
		s.logFailure(r, "install standing", err)
	}
	if err != nil || !found {
		http.Redirect(w, r, s.d.PublicBaseURL+"/", http.StatusFound)
		return
	}
	outcome := s.installOutcome(r, cs, standing)
	http.Redirect(w, r, s.d.PublicBaseURL+standing.Path+botOutcomeParam+outcome+discordAnchor, http.StatusFound)
}

func (s *Service) installOutcome(r *http.Request, cs notify.ConnectState, standing notify.GuildStanding) string {
	actor := auth.ActorFrom(r.Context())
	if !actor.Signed() || actor.IsDevice() || actor.UserID != cs.UserID || cs.Expired ||
		!standing.Officer || !standing.Plan || standing.Frozen {
		return installFailed
	}
	query := r.URL.Query()
	if discordError := query.Get("error"); discordError != "" && discordError != oauthAccessDenied {
		return installFailed
	}
	code := query.Get("code")
	if code == "" || query.Get("error") == oauthAccessDenied {
		return installDenied
	}
	server, err := s.d.OAuth.ExchangeInstall(r.Context(), code)
	if err != nil || server.ID == "" {
		s.log.Warn("discordbot", "id", httpx.RequestIDFrom(r.Context()), "op", "install exchange", "err", err)
		return installFailed
	}
	return s.bindOutcome(r, cs, standing, server)
}

func (s *Service) bindOutcome(r *http.Request, cs notify.ConnectState, standing notify.GuildStanding, server InstalledServer) string {
	ctx := r.Context()
	res, err := s.d.Store.BindServer(ctx, cs.GuildID, server.ID, server.Name, cs.UserID)
	if err != nil {
		s.logFailure(r, "install bind", err)
		return installFailed
	}
	if res.Outcome == BindInstalled {
		s.log.Info("discordbot", "op", "install", "guild_id", cs.GuildID, "user_id", cs.UserID)
		return installInstalled
	}
	holder, err := s.holderName(ctx, res.HeldBy, cs, standing)
	if err == nil {
		err = s.d.Store.SaveRefusal(ctx, cs.UserID, cs.GuildID, holder, refusalTTL)
	}
	if err != nil {
		s.logFailure(r, "install refusal", err)
		return installFailed
	}
	return installTaken
}

// holderName is the name of the guild that blocks the install: the officer's own guild when
// it already has another server.
func (s *Service) holderName(ctx context.Context, heldBy int64, cs notify.ConnectState, own notify.GuildStanding) (string, error) {
	if heldBy == cs.GuildID {
		return own.Name, nil
	}
	standing, found, err := s.d.Guilds.GuildStanding(ctx, heldBy, cs.UserID)
	if err != nil {
		return "", fmt.Errorf("discordbot: holder name: %w", err)
	}
	if !found {
		return "", fmt.Errorf("discordbot: holder guild %d not found", heldBy)
	}
	return standing.Name, nil
}

// unusableView is the settings shape of an unusable channel; nil while the channel is fine.
func unusableView(u *ChannelUnusable) *guilds.ChannelUnusable {
	if u == nil {
		return nil
	}
	return &guilds.ChannelUnusable{Reason: u.Reason, Missing: append([]string{}, u.Missing...)}
}
