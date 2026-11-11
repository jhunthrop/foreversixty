// api/cmd/api/stripe_reconcile.go
//
// stripe-reconcile is the nightly backstop (spec §2.8), run the same way
// sim-validate already is: a Cloud Scheduler job hitting this image.
package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jhunthrop/foreversixty/api/internal/billing"
)

func runStripeReconcile(ctx context.Context, log *slog.Logger) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	if !cfg.StripeConfigured() {
		log.Warn("stripe-reconcile", "state", "stripe not configured", "effect", "nothing to reconcile")
		return nil
	}
	svc := &billing.Service{
		Store: &billing.Store{Pool: pool}, Entitlements: entitlementsStore(pool, linkedRoleSyncer(cfg, pool, log)),
		Gateway: billing.NewStripeGateway(cfg.StripeSecretKey), Log: log,
	}
	n, err := svc.Reconcile(ctx)
	log.Info("stripe-reconcile", "reconciled", n)
	if err != nil {
		return fmt.Errorf("stripe-reconcile: %w", err)
	}
	return nil
}
