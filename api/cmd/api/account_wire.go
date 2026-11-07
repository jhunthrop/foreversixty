package main

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/account"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/linkedrole"
	"github.com/jhunthrop/foreversixty/api/internal/r2"
)

// accountHandler is DELETE /v1/me over the deployment's billing service, linked-role syncer
// and object storage. The syncer and the client may be nil (not configured); each becomes a
// nil interface, never a typed nil, so the deletion skips that step instead of panicking.
func accountHandler(pool *pgxpool.Pool, bill *billing.Service, roles *linkedrole.Syncer, objects *r2.Client,
	sessions *auth.Authenticator, log *slog.Logger) *account.Handler {
	svc := &account.Service{Store: &account.Store{Pool: pool}, Billing: bill, Log: log}
	if roles != nil {
		svc.Discord = roles
	}
	if objects != nil {
		svc.Objects = objects
	}
	return &account.Handler{Service: svc, Sessions: sessions}
}
