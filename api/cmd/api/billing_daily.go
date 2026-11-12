// api/cmd/api/billing_daily.go
//
// billing-daily is the daily Cloud Scheduler job: it releases seats that no
// longer count, sends the renewal and size notices, then re-pushes the Discord
// Premium role of every linked account. It logs counts only.
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/linkedrole"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
	"github.com/jhunthrop/foreversixty/api/internal/seats"
)

func runBillingDaily(ctx context.Context, log *slog.Logger) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	mailer := mail.NewResend(cfg.ResendAPIKey, cfg.MailSender(), nil)
	return billingDaily(ctx, log, newBillingDailyServices(cfg, pool, mailer, log), time.Now())
}

// billingDailyServices is what the job runs; tests build it with a fake mailer.
type billingDailyServices struct {
	Seats   *seats.Service
	Billing *billing.Service
	// Roles reconciles the Discord linked role; nil (the job lacks the Discord
	// application or the sealing key) skips it.
	Roles *linkedrole.Syncer
}

func newBillingDailyServices(cfg config.Config, pool *pgxpool.Pool, mailer mail.Mailer, log *slog.Logger) billingDailyServices {
	roles := linkedRoleSyncer(cfg, pool, log)
	entStore := entitlementsStore(pool, roles)
	return billingDailyServices{
		Roles: roles,
		// The Canceller lets the sweep resume a holder's own Premium when it
		// releases a seat that never started.
		Seats: &seats.Service{
			Store:  &seats.Store{Pool: pool, Entitlements: entStore},
			Mailer: mailer, Canceller: cancellerFor(stripeGatewayFor(cfg)), PublicBaseURL: cfg.PublicBaseURL, Log: log,
		},
		// The gateway reads the Stripe customer's email for a trial reminder; the router puts the
		// reminder in the app (nil without push: email alone).
		Billing: &billing.Service{
			Store: &billing.Store{Pool: pool}, Users: &auth.Store{Pool: pool}, Gateway: stripeGatewayFor(cfg),
			Events: jobEvents(cfg, pool, log), RenewalEmails: cfg.OurRenewalEmails,
			Mailer: mailer, PublicBaseURL: cfg.PublicBaseURL, Log: log,
		},
	}
}

// jobEvents is a job's event publisher: the router when app notifications are on, else nil.
func jobEvents(cfg config.Config, pool *pgxpool.Pool, log *slog.Logger) events.Publisher {
	if !cfg.Push.Enabled {
		return nil
	}
	if router := newEventsRouter(context.Background(), cfg, pool, nil, log); router != nil {
		return router
	}
	return nil
}

// billingDaily sweeps seats, then sends notices. Per-guild sweep failures are
// logged by the sweep; the job fails on them only when every guild failed
// (nothing was swept), and always on a notice failure.
func billingDaily(ctx context.Context, log *slog.Logger, svc billingDailyServices, now time.Time) error {
	swept, sweepErr := svc.Seats.Sweep(ctx, now)
	sent, noticeErr := svc.Billing.SendNotices(ctx, now)
	roles, rolesErr := reconcileRoles(ctx, svc.Roles)
	log.Info("billing-daily", "seats_released", swept.Released, "guilds_failed", swept.Failed, "notices_sent", sent,
		"roles_synced", roles.Synced, "roles_failed", roles.Failed)
	var errs []error
	if rolesErr != nil {
		errs = append(errs, fmt.Errorf("linked roles: %w", rolesErr))
	}
	if sweepErr != nil && swept.TotalFailure() {
		errs = append(errs, fmt.Errorf("sweep: %w", sweepErr))
	}
	if noticeErr != nil {
		errs = append(errs, fmt.Errorf("notices: %w", noticeErr))
	}
	if err := errors.Join(errs...); err != nil {
		return fmt.Errorf("billing-daily: %w", err)
	}
	return nil
}

// reconcileRoles runs the linked-role reconcile when the job has it.
func reconcileRoles(ctx context.Context, roles *linkedrole.Syncer) (linkedrole.ReconcileResult, error) {
	if roles == nil {
		return linkedrole.ReconcileResult{}, nil
	}
	return roles.Reconcile(ctx)
}
