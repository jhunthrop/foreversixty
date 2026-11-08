package main

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/livesurface"
	"github.com/jhunthrop/foreversixty/api/internal/push"
)

// wireLivePusher gives g the raid Live Activity's start and end pushes, when the app
// notifications are on and some push sender exists (APNs once its key is set, FCM as the
// runtime identity); otherwise g sends none.
func wireLivePusher(ctx context.Context, cfg config.Config, pool *pgxpool.Pool, g *guilds.Service, log *slog.Logger) {
	if !cfg.Push.Enabled || g == nil {
		return
	}
	senders := pushSenders(ctx, cfg.Push, log)
	p := &livesurface.LivePusher{Tokens: &livesurface.Store{Pool: pool}, Raids: g, Log: log}
	if apns, ok := senders[push.IOS].(*push.APNsSender); ok {
		p.APNs = apns
	}
	if fcm, ok := senders[push.Android].(*push.FCMSender); ok {
		p.FCM = fcm
	}
	if p.APNs == nil && p.FCM == nil {
		return
	}
	g.LiveRaids = p
}
