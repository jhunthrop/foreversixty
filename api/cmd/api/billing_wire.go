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

// wireBilling builds the Stripe gateway (nil, with one warn line, when Stripe
// is not configured) and the billing and seats services that share it.
func wireBilling(cfg config.Config, pool *pgxpool.Pool, deps *server.Deps, entStore *entitlements.Store,
	authStore *auth.Store, guildStore *guilds.Store, mailer mail.Mailer, log *slog.Logger) {
	var gateway billing.Gateway
	if cfg.StripeConfigured() {
		gateway = billing.NewStripeGateway(cfg.StripeSecretKey)
	} else {
		log.Warn("billing", "state", "stripe is not configured",
			"effect", "checkout, portal, and the webhook answer 503")
	}
	deps.Billing = &billing.Service{
		Store: &billing.Store{Pool: pool}, Entitlements: entStore,
		Accounts: authStore, Users: authStore, Guilds: guildStore, Gateway: gateway,
		PublicBaseURL: cfg.PublicBaseURL, WebhookSecret: cfg.StripeWebhookSecret, Log: log,
	}
	deps.Seats = &seats.Service{
		Store:  &seats.Store{Pool: pool, Entitlements: entStore},
		Mailer: mailer, Guilds: guildStore, PublicBaseURL: cfg.PublicBaseURL, Log: log,
	}
	// A nil gateway must stay a nil interface so Accept answers 503.
	if gateway != nil {
		deps.Seats.Canceller = gateway
	}
}
