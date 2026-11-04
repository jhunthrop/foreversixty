package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/linkedrole"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/server"
)

// linkedRoleSetupWithin bounds the startup metadata and verification-URL calls.
const linkedRoleSetupWithin = 30 * time.Second

// linkedRoleSyncer keeps the Discord Premium role's metadata in step with Premium. It is nil,
// with one warn line, unless the Discord application (DISCORD_CLIENT_ID/SECRET, whose id is the
// application id) and NOTIFY_ENCRYPTION_KEY are set; the daily reconcile then catches up from
// wherever they are.
func linkedRoleSyncer(cfg config.Config, pool *pgxpool.Pool, log *slog.Logger) *linkedrole.Syncer {
	if cfg.DiscordClientID == "" || len(cfg.NotifyEncryptionKey) == 0 {
		log.Warn("linkedrole", "state", "off", "effect", "the Discord Premium role is not re-pushed from this process")
		return nil
	}
	sealer, err := notify.NewSealer(cfg.NotifyEncryptionKey)
	if err != nil {
		log.Error("linkedrole", "state", "off", "err", err)
		return nil
	}
	return &linkedrole.Syncer{
		Store: &linkedrole.PGStore{Pool: pool}, Sealer: sealer, Log: log,
		Discord: linkedrole.NewHTTPDiscord(cfg.DiscordClientID, cfg.DiscordClientSecret, notify.CallbackURL(cfg.APIBaseURL)),
	}
}

// entitlementsStore is the entitlements store whose writes re-push the linked role when the
// syncer is on.
func entitlementsStore(pool *pgxpool.Pool, syncer *linkedrole.Syncer) *entitlements.Store {
	store := &entitlements.Store{Pool: pool}
	if syncer != nil {
		store.OnChange = syncer.OnEntitlementChange
	}
	return store
}

// wireLinkedRole mounts the verification start (it finishes at notify's shared callback, so the
// guild connect must be on) and, with the bot token, registers the metadata and the
// verification URL off the serving path.
func wireLinkedRole(stop context.Context, cfg config.Config, syncer *linkedrole.Syncer, deps *server.Deps, log *slog.Logger) {
	if syncer == nil {
		return
	}
	if deps.Notify == nil || !deps.Notify.GuildConnectEnabled() {
		log.Warn("linkedrole", "state", "no verification route", "effect", "the shared Discord callback is off")
		return
	}
	deps.LinkedRole = linkedrole.NewHandler(syncer, deps.Notify, cfg.PublicBaseURL, log)
	if cfg.DiscordBotToken == "" {
		return
	}
	go setupLinkedRole(stop, linkedrole.NewApp(cfg.DiscordClientID, cfg.DiscordBotToken), cfg.APIBaseURL+linkedrole.StartPath, log)
}

// setupLinkedRole runs the idempotent application setup; a failure is logged (the error carries
// no token) and the next start tries again.
func setupLinkedRole(stop context.Context, app *linkedrole.App, verifyURL string, log *slog.Logger) {
	ctx, cancel := context.WithTimeout(stop, linkedRoleSetupWithin)
	defer cancel()
	if err := app.Setup(ctx, verifyURL); err != nil {
		log.Error("linkedrole", "op", "setup", "outcome", "failed", "err", err)
		return
	}
	log.Info("linkedrole", "op", "setup", "outcome", "ok")
}
