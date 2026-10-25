package main

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
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
		Store: &notify.Store{Pool: pool}, Sealer: sealer, Entitlements: ents,
		Poster: notify.NewHTTPPoster(notify.AvatarURL(cfg.PublicBaseURL)), PublicBaseURL: cfg.PublicBaseURL, Log: log,
	}
}

// newJobDeliverer is the poster the sim and parse jobs use, built from the
// same wiring as the routes. It is nil when notifications are off, so the
// jobs' optional Notify stays a true nil interface.
func newJobDeliverer(cfg config.Config, pool *pgxpool.Pool, log *slog.Logger) *notify.Deliverer {
	service := newNotifyService(cfg, pool, &entitlements.Store{Pool: pool}, log)
	if service == nil {
		return nil
	}
	return service.Deliverer()
}
