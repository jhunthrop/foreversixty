package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/creators"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/server"
)

// connectCheckBudget bounds the startup Connect check.
const connectCheckBudget = 20 * time.Second

// creatorsWiring is what wireCreators reads from serve.
type creatorsWiring struct {
	cfg           config.Config
	pool          *pgxpool.Pool
	deps          *server.Deps
	entitle       *entitlements.Store
	accounts      *auth.Service
	authenticator *auth.Authenticator
	log           *slog.Logger
}

// wireCreators builds the creator program (design/specs/2026-11-08-creator-program.md) on the
// deployment's own Stripe key and hooks it into the tracked links, sign-in visits, /v1/me and
// billing's checkout and webhook. Run it after wireBilling.
func wireCreators(w creatorsWiring) {
	store := &creators.Store{Pool: w.pool}
	svc := &creators.Service{
		Store: store, Entitlements: w.entitle, PublicBaseURL: w.cfg.PublicBaseURL,
		PayoutsEnabled: w.cfg.CreatorPayouts, Log: w.log,
	}
	if w.cfg.StripeConfigured() {
		svc.Gateway = creators.NewStripeGateway(w.cfg.StripeSecretKey, w.cfg.StripePayoutsKey)
	}
	w.log.Info("creators", "stripe", svc.Gateway != nil, "payouts_key", w.cfg.StripePayoutsKey != "", "payouts", svc.PayoutsEnabled)
	// Payouts on: check the key reaches Connect now, in the background so startup never waits on Stripe.
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), connectCheckBudget)
		defer cancel()
		svc.CheckConnect(ctx)
	}()
	w.deps.Creators = svc
	if w.deps.GoLinks != nil {
		w.deps.GoLinks.Arrivals = svc
	}
	if w.deps.Billing != nil {
		w.deps.Billing.Creators = svc
		w.deps.Billing.ConnectWebhookSecret = w.cfg.StripeConnectWebhookSecret
	}
	w.accounts.Creators = store
	w.authenticator.Visits = auth.AddVisitHook(w.authenticator.Visits, svc)
}
