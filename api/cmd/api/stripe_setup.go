// api/cmd/api/stripe_setup.go
//
// stripe-setup is an idempotent Products/Prices setup command (spec
// RULING 5), safe to run repeatedly and against either Stripe mode.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/config"
)

// centsPerDollar keeps priceSpecs' amounts readable as dollars.
const centsPerDollar = 100

// priceSpecs is spec §2.1's table verbatim — the only place a Stripe
// amount or product name is hardcoded (Ruling D: fixed product ids so
// this is idempotent with no Search API dependency).
var priceSpecs = []billing.PriceSpec{
	{ProductID: "prod_fs_premium", ProductName: "Forever Sixty Premium",
		LookupKey: billing.LookupKeyPremiumMonthly, Interval: "month", UnitAmountCents: 4 * centsPerDollar},
	{ProductID: "prod_fs_premium", ProductName: "Forever Sixty Premium",
		LookupKey: billing.LookupKeyPremiumYearly, Interval: "year", UnitAmountCents: 40 * centsPerDollar},
	{ProductID: "prod_fs_guild", ProductName: "Forever Sixty Guild",
		LookupKey: billing.LookupKeyGuildMonthly, Interval: "month", UnitAmountCents: 15 * centsPerDollar},
	{ProductID: "prod_fs_guild", ProductName: "Forever Sixty Guild",
		LookupKey: billing.LookupKeyGuildYearly, Interval: "year", UnitAmountCents: 150 * centsPerDollar},
}

func runStripeSetup(ctx context.Context, log *slog.Logger) error {
	cfg, err := config.Load(os.Getenv)
	if err != nil {
		return err
	}
	if err := cfg.ValidateStripeKeyEnvironment(); err != nil {
		return err
	}
	if !cfg.StripeConfigured() {
		return fmt.Errorf("stripe-setup: STRIPE_SECRET_KEY and STRIPE_WEBHOOK_SECRET are required")
	}
	gw := billing.NewStripeGateway(cfg.StripeSecretKey)
	for _, spec := range priceSpecs {
		if err := gw.EnsurePrice(ctx, spec); err != nil {
			return fmt.Errorf("stripe-setup: %s: %w", spec.LookupKey, err)
		}
		log.Info("stripe-setup", "lookup_key", spec.LookupKey, "product", spec.ProductID)
	}
	return nil
}
