package main

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
	"github.com/jhunthrop/foreversixty/api/internal/seats"
	"github.com/jhunthrop/foreversixty/api/internal/server"
)

// billingWiring is what wireBilling reads from serve.
type billingWiring struct {
	cfg        config.Config
	pool       *pgxpool.Pool
	deps       *server.Deps
	entitle    *entitlements.Store
	authStore  *auth.Store
	guildStore *guilds.Store
	mailer     mail.Mailer
	log        *slog.Logger
}

// wireBilling builds the Stripe gateway (nil, with one warn line, when Stripe
// is not configured) and the billing and seats services that share it.
func wireBilling(w billingWiring) {
	cfg, pool, deps, entStore := w.cfg, w.pool, w.deps, w.entitle
	authStore, guildStore, mailer, log := w.authStore, w.guildStore, w.mailer, w.log
	var gateway billing.Gateway
	if cfg.StripeConfigured() {
		gateway = billing.NewStripeGateway(cfg.StripeSecretKey)
	} else {
		log.Warn("billing", "state", "stripe is not configured",
			"effect", "checkout, portal, and the webhook answer 503")
	}
	seatStore := &seats.Store{Pool: pool, Entitlements: entStore}
	deps.Seats = &seats.Service{
		Store:  seatStore,
		Mailer: mailer, Guilds: guildStore, PublicBaseURL: cfg.PublicBaseURL, Log: log,
	}
	deps.Billing = &billing.Service{
		Store: &billing.Store{Pool: pool}, Entitlements: entStore,
		Accounts: authStore, Users: authStore, Guilds: guildStore, Gateway: gateway,
		Seats:         seats.BillingReader{Store: seatStore, Sweeper: deps.Seats},
		PublicBaseURL: cfg.PublicBaseURL, WebhookSecret: cfg.StripeWebhookSecret, Log: log,
	}
	// A nil gateway must stay a nil interface so Accept answers 503.
	if gateway != nil {
		deps.Seats.Canceller = gateway
	}
}
