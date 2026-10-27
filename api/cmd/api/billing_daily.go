// api/cmd/api/billing_daily.go
//
// billing-daily is the daily Cloud Scheduler job: it releases seats that no
// longer count, then sends the renewal and size notices. It logs counts only.
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
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
	"github.com/jhunthrop/foreversixty/api/internal/seats"
)

func runBillingDaily(ctx context.Context, log *slog.Logger) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	mailer := mail.NewResend(cfg.ResendAPIKey, cfg.MailFrom, nil)
	return billingDaily(ctx, log, newBillingDailyServices(cfg, pool, mailer, log), time.Now())
}

// billingDailyServices is what the job runs; tests build it with a fake mailer.
type billingDailyServices struct {
	Seats   *seats.Service
	Billing *billing.Service
}

func newBillingDailyServices(cfg config.Config, pool *pgxpool.Pool, mailer mail.Mailer, log *slog.Logger) billingDailyServices {
	entStore := &entitlements.Store{Pool: pool}
	return billingDailyServices{
		Seats: &seats.Service{
			Store:  &seats.Store{Pool: pool, Entitlements: entStore},
			Mailer: mailer, PublicBaseURL: cfg.PublicBaseURL, Log: log,
		},
		Billing: &billing.Service{
			Store: &billing.Store{Pool: pool}, Users: &auth.Store{Pool: pool},
			Mailer: mailer, PublicBaseURL: cfg.PublicBaseURL, Log: log,
		},
	}
}

// billingDaily sweeps seats, then sends notices. Per-guild sweep failures are
// logged by the sweep; the job fails on them only when every guild failed
// (nothing was swept), and always on a notice failure.
func billingDaily(ctx context.Context, log *slog.Logger, svc billingDailyServices, now time.Time) error {
	swept, sweepErr := svc.Seats.Sweep(ctx, now)
	sent, noticeErr := svc.Billing.SendNotices(ctx, now)
	log.Info("billing-daily", "seats_released", swept.Released, "guilds_failed", swept.Failed, "notices_sent", sent)
	var errs []error
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
