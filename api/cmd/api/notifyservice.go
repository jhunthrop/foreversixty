package main

import (
	"context"
	"log/slog"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/server"
)

// newNotifyService builds the Discord notification routes. Without the
// encryption key there is nothing to seal webhook tokens with, so the
// feature is off: no routes and no /v1/me block. notify.Service itself is
// the auth.NotificationsView, so no adapter sits between them.
func newNotifyService(cfg config.Config, pool *pgxpool.Pool, ents notify.Entitlementer, log *slog.Logger) *notify.Service {
	if len(cfg.NotifyEncryptionKey) == 0 {
		log.Warn("notify", "state", "notifications off", "effect", "no Discord routes, no /v1/me block")
		return nil
	}
	sealer, err := notify.NewSealer(cfg.NotifyEncryptionKey)
	if err != nil {
		log.Error("notify", "state", "notifications off", "err", err)
		return nil
	}
	return &notify.Service{
		Store: &notify.Store{Pool: pool}, GuildStore: &notify.GuildStore{Pool: pool}, Sealer: sealer, Entitlements: ents,
		Poster: notify.NewHTTPPoster(notify.AvatarURL(cfg.PublicBaseURL)), PublicBaseURL: cfg.PublicBaseURL, Log: log,
		OAuth: newDiscordOAuth(cfg), CookieSecure: strings.HasPrefix(cfg.APIBaseURL, "https://"),
	}
}

// newDiscordOAuth builds the one-click connect's client; nil (connect off)
// unless the Discord application is configured. config.Load guarantees the
// id and secret come together.
func newDiscordOAuth(cfg config.Config) notify.OAuthClient {
	if cfg.DiscordClientID == "" {
		return nil
	}
	return notify.NewHTTPOAuthClient(cfg.DiscordClientID, cfg.DiscordClientSecret, notify.CallbackURL(cfg.APIBaseURL))
}

// newJobGuilds is the guilds service a job tells of a completed report
// (reports.Completed), on the job's own pool. It carries no Discord
// poster; with the app notifications on it publishes the night's
// report.ready to the raiders.
func newJobGuilds(cfg config.Config, pool *pgxpool.Pool, log *slog.Logger) *guilds.Service {
	g := &guilds.Service{Store: &guilds.Store{Pool: pool}, PublicBaseURL: cfg.PublicBaseURL, Log: log}
	if !cfg.Push.Enabled {
		return g
	}
	if router := newEventsRouter(context.Background(), cfg, pool, nil, log); router != nil {
		// The night's raid.summary names the coach notes to an account that holds the coach.
		g.Events, g.Entitlements = router, &entitlements.Store{Pool: pool}
	}
	return g
}

// wireNotify builds the notify service into deps and connects it to the
// accounts and guilds services it talks to. It leaves deps.Notify nil when
// notifications are off.
func wireNotify(cfg config.Config, pool *pgxpool.Pool, ents notify.Entitlementer, accounts *auth.Service, deps *server.Deps, log *slog.Logger) {
	deps.Notify = newNotifyService(cfg, pool, ents, log)
	if deps.Notify == nil {
		return
	}
	accounts.Notifications = deps.Notify
	wireGuildRaidChannel(deps.Notify, deps.Guilds)
}

// wireGuildRaidChannel connects notify and guilds: notify asks guilds for an officer's
// standing, guilds reads the channel back for the settings payload and the raid view, and the
// raid posts follow the channel (none while the deployment has connect off). serve and the
// raid-posts job wire it the same way.
func wireGuildRaidChannel(n *notify.Service, g *guilds.Service) {
	n.Guilds = g
	g.Discord = n
	if poster := n.GuildDeliverer(); poster != nil && n.GuildConnectEnabled() {
		g.RaidEvents = poster
	}
}
