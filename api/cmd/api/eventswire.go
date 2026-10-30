// api/cmd/api/eventswire.go
//
// The events router's wiring, shared by serve and the jobs: the account webhook is the
// Discord channel whenever notifications are on (as before the router), and push joins
// only with APP_NOTIFICATIONS_ENABLED=1 and a sender whose keys are set. Until then the
// new routes stay unmounted and no producer beyond the sim and parse jobs publishes.
package main

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/push"
	"github.com/jhunthrop/foreversixty/api/internal/server"
)

// newEventsRouter builds the router over n's webhook (n nil: no Discord channel) and, when
// the flag is on, the push senders. It is nil when no channel exists at all.
func newEventsRouter(ctx context.Context, cfg config.Config, pool *pgxpool.Pool, n *notify.Service, log *slog.Logger) *events.Router {
	sinks := map[events.Channel]events.Sink{}
	if n != nil {
		sinks[events.ChannelDiscord] = &notify.DiscordSink{Deliverer: n.Deliverer()}
	}
	if cfg.Push.Enabled {
		if senders := pushSenders(ctx, cfg.Push, log); len(senders) > 0 {
			sinks[events.ChannelPush] = &push.Sink{Store: &push.Store{Pool: pool}, Senders: senders, Log: log}
		}
	}
	if len(sinks) == 0 {
		return nil
	}
	return &events.Router{Store: &events.Store{Pool: pool}, Sinks: sinks, Log: log}
}

// pushSenders is a sender per platform whose keys are set. A key that does not parse is
// logged by variable name and leaves that platform off.
func pushSenders(ctx context.Context, cfg config.PushConfig, log *slog.Logger) map[push.Platform]push.Sender {
	senders := map[push.Platform]push.Sender{}
	if cfg.APNsConfigured() {
		host := push.APNsProductionURL
		if cfg.APNsSandbox {
			host = push.APNsSandboxURL
		}
		apns, err := push.NewAPNs(push.APNsConfig{KeyP8: cfg.APNsKeyP8, KeyID: cfg.APNsKeyID, TeamID: cfg.APNsTeamID,
			Topic: cfg.APNsTopic, BaseURL: host}, nil)
		if err != nil {
			log.Error("push", "state", "APNs off", "variable", "APNS_KEY_P8", "err", err)
		} else {
			senders[push.IOS] = apns
		}
	}
	if cfg.FCMConfigured() {
		fcm, err := push.NewFCM(ctx, cfg.FCMServiceAccountJSON, push.FCMBaseURL, nil)
		if err != nil {
			log.Error("push", "state", "FCM off", "variable", "FCM_SERVICE_ACCOUNT_JSON", "err", err)
		} else {
			senders[push.Android] = fcm
		}
	}
	if len(senders) == 0 {
		log.Warn("push", "state", "no push sender", "effect", "events reach Discord only")
	}
	return senders
}

// newJobNotifier is the sim and parse jobs' notify port: their events through the router.
// It is nil when the router has no channel, so the jobs' optional Notify stays a true nil.
func newJobNotifier(ctx context.Context, cfg config.Config, pool *pgxpool.Pool, log *slog.Logger) *notify.EventProducer {
	n := newNotifyService(cfg, pool, &entitlements.Store{Pool: pool}, log)
	router := newEventsRouter(ctx, cfg, pool, n, log)
	if router == nil {
		return nil
	}
	return &notify.EventProducer{Publisher: router, PublicBaseURL: cfg.PublicBaseURL, Log: log}
}

// wireEvents connects serve's producers and, with the flag on, mounts the preference and
// push token routes and lets a device token answer a raid sign-up. deps.Notify must already
// be wired.
func wireEvents(ctx context.Context, cfg config.Config, pool *pgxpool.Pool, deps *server.Deps, log *slog.Logger) {
	if !cfg.Push.Enabled {
		return
	}
	deps.Notifications = &events.Handler{Store: &events.Store{Pool: pool}, Log: log}
	deps.PushTokens = &push.Handler{Store: &push.Store{Pool: pool}, Log: log}
	if deps.Guilds != nil {
		deps.Guilds.AppSignups = true
	}
	router := newEventsRouter(ctx, cfg, pool, deps.Notify, log)
	if router == nil {
		return
	}
	connectProducers(deps, router, cfg.PublicBaseURL, log)
}

// connectProducers gives every producer serve holds the router.
func connectProducers(deps *server.Deps, router events.Publisher, publicBaseURL string, log *slog.Logger) {
	if deps.Guilds != nil {
		deps.Guilds.Events = router
	}
}
