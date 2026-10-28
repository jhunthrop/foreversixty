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
// is not configured) and the billing and seats services that share it. It
// returns the seats store the account read needs.
func wireBilling(w billingWiring) *seats.Store {
	cfg, pool, deps, entStore := w.cfg, w.pool, w.deps, w.entitle
	authStore, guildStore, mailer, log := w.authStore, w.guildStore, w.mailer, w.log
	gateway := stripeGatewayFor(cfg)
	if gateway == nil {
		log.Warn("billing", "state", "stripe is not configured",
			"effect", "checkout, portal, and the webhook answer 503")
	}
	seatStore := &seats.Store{Pool: pool, Entitlements: entStore}
	deps.Seats = &seats.Service{
		Store:  seatStore,
		Mailer: mailer, Guilds: guildStore, PublicBaseURL: cfg.PublicBaseURL, Log: log,
	}
	billingStore := &billing.Store{Pool: pool}
	// The subscribe step on GET /v1/me reads the same checkout facts billing refuses on.
	guildStore.Billing, guildStore.Plans = billingStore, entStore
	deps.Billing = &billing.Service{
		Store: billingStore, Entitlements: entStore,
		Accounts: authStore, Users: authStore, Guilds: guildStore, Gateway: gateway,
		Seats:         seats.BillingReader{Store: seatStore, Sweeper: deps.Seats},
		PublicBaseURL: cfg.PublicBaseURL, WebhookSecret: cfg.StripeWebhookSecret, Log: log,
	}
	deps.Seats.Canceller = cancellerFor(gateway)
	return seatStore
}

// stripeGatewayFor is the Stripe gateway, nil when Stripe is not configured.
func stripeGatewayFor(cfg config.Config) billing.Gateway {
	if !cfg.StripeConfigured() {
		return nil
	}
	return billing.NewStripeGateway(cfg.StripeSecretKey)
}

// cancellerFor is the seats service's Canceller. A nil gateway stays a nil
// interface, so Accept answers 503 and a release falls back to the comeback.
func cancellerFor(gateway billing.Gateway) seats.Canceller {
	if gateway == nil {
		return nil
	}
	return gateway
}
