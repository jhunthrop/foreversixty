package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/livesurface"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
)

// liveSweepEvery is how often ended live registrations are removed.
const liveSweepEvery = time.Hour

// liveSurfaceWiring is what the mobile live surfaces read.
type liveSurfaceWiring struct {
	cfg     config.Config
	pool    *pgxpool.Pool
	auth    *auth.Store
	guilds  *guilds.Service
	sims    *sims.Store
	entitle *entitlements.Store
	log     *slog.Logger
}

// newLiveSurface builds the widget snapshot and live-registration routes, or nil while
// APP_WIDGETS is off. The snapshot's upgrades stay null until the Home answer's builder (lane
// A3, api/internal/app) is on main to back livesurface.Upgrades. APNs and FCM senders for the
// live.* and widget.refresh events arrive with lane A2; until then the registrations are kept
// and swept, and the apps update their surfaces locally.
func newLiveSurface(ctx context.Context, w liveSurfaceWiring) *livesurface.Service {
	if !w.cfg.AppWidgets {
		return nil
	}
	tokens := &livesurface.Store{Pool: w.pool}
	go sweepLiveRegistrations(ctx, tokens, w.log)
	return &livesurface.Service{
		Accounts: w.auth, Raids: w.guilds, Sims: w.sims, Entitlements: w.entitle, Tokens: tokens, Log: w.log,
	}
}

// sweepLiveRegistrations removes registrations whose raid or run has ended, hourly.
func sweepLiveRegistrations(ctx context.Context, tokens *livesurface.Store, log *slog.Logger) {
	tick := time.NewTicker(liveSweepEvery)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case now := <-tick.C:
			if n, err := tokens.DeleteExpired(ctx, now); err != nil {
				log.Error("livesurface", "op", "sweep", "err", err)
			} else if n > 0 {
				log.Info("livesurface", "op", "sweep", "removed", n)
			}
		}
	}
}
