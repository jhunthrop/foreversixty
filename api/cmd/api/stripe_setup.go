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
	{ProductID: billing.ProductIDPremium, ProductName: "Forever Sixty Premium", TaxCode: billing.ProductTaxCodeSaaSPersonal,
		LookupKey: billing.LookupKeyPremiumMonthly, Interval: "month", UnitAmountCents: 4 * centsPerDollar},
	{ProductID: billing.ProductIDPremium, ProductName: "Forever Sixty Premium", TaxCode: billing.ProductTaxCodeSaaSPersonal,
		LookupKey: billing.LookupKeyPremiumYearly, Interval: "year", UnitAmountCents: 40 * centsPerDollar},
	{ProductID: billing.ProductIDGuild, ProductName: "Forever Sixty Guild", TaxCode: billing.ProductTaxCodeSaaSPersonal,
		LookupKey: billing.LookupKeyGuildMonthly, Interval: "month", UnitAmountCents: 15 * centsPerDollar},
	{ProductID: billing.ProductIDGuild, ProductName: "Forever Sixty Guild", TaxCode: billing.ProductTaxCodeSaaSPersonal,
		LookupKey: billing.LookupKeyGuildYearly, Interval: "year", UnitAmountCents: 150 * centsPerDollar},
	// The launch prices (docs/superpowers/specs/2026-10-27-pricing-design.md §3), sold once
	// the pricing lane switches the site over.
	premiumSpec(billing.LookupKeyPremiumMonthlyLaunch, "month", 6*centsPerDollar),
	premiumSpec(billing.LookupKeyPremiumYearlyLaunch, "year", 50*centsPerDollar),
	premiumSpec(billing.LookupKeyPremiumFounderMonthly, "month", 5*centsPerDollar),
	premiumSpec(billing.LookupKeyPremiumFounderYearly, "year", 40*centsPerDollar),
	guildSpec(billing.LookupKeyGuildFounderMonthly, "month", 12*centsPerDollar),
	guildSpec(billing.LookupKeyGuildFounderYearly, "year", 120*centsPerDollar),
	raidGuildSpec(billing.LookupKeyRaidGuildMonthly, "month", 25*centsPerDollar),
	raidGuildSpec(billing.LookupKeyRaidGuildYearly, "year", 250*centsPerDollar),
	raidGuildSpec(billing.LookupKeyRaidGuildFounderMonthly, "month", 20*centsPerDollar),
	raidGuildSpec(billing.LookupKeyRaidGuildFounderYearly, "year", 200*centsPerDollar),
}

func premiumSpec(lookupKey, interval string, cents int64) billing.PriceSpec {
	return billing.PriceSpec{ProductID: billing.ProductIDPremium, ProductName: "Forever Sixty Premium",
		TaxCode: billing.ProductTaxCodeSaaSPersonal, LookupKey: lookupKey, Interval: interval, UnitAmountCents: cents}
}

func guildSpec(lookupKey, interval string, cents int64) billing.PriceSpec {
	return billing.PriceSpec{ProductID: billing.ProductIDGuild, ProductName: "Forever Sixty Guild",
		TaxCode: billing.ProductTaxCodeSaaSPersonal, LookupKey: lookupKey, Interval: interval, UnitAmountCents: cents}
}

func raidGuildSpec(lookupKey, interval string, cents int64) billing.PriceSpec {
	return billing.PriceSpec{ProductID: billing.ProductIDRaidGuild, ProductName: "Forever Sixty Raid Guild",
		TaxCode: billing.ProductTaxCodeSaaSPersonal, LookupKey: lookupKey, Interval: interval, UnitAmountCents: cents}
}

// portalSpec is the Customer Portal stripe-setup ensures (spec §2.5):
// the site's own legal pages, a return to the account page, and each
// product's two prices as the only switch a subscriber can make there.
func portalSpec(publicBaseURL string) billing.PortalSpec {
	return billing.PortalSpec{
		Headline:          "Forever Sixty",
		PrivacyPolicyURL:  publicBaseURL + "/privacy",
		TermsOfServiceURL: publicBaseURL + "/terms",
		DefaultReturnURL:  publicBaseURL + "/account",
		Products: []billing.PortalProduct{
			{ProductID: billing.ProductIDPremium, LookupKeys: []string{billing.LookupKeyPremiumMonthly, billing.LookupKeyPremiumYearly}},
			{ProductID: billing.ProductIDGuild, LookupKeys: []string{billing.LookupKeyGuildMonthly, billing.LookupKeyGuildYearly}},
		},
	}
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
	if err := gw.EnsurePortalConfiguration(ctx, portalSpec(cfg.PublicBaseURL)); err != nil {
		return fmt.Errorf("stripe-setup: portal: %w", err)
	}
	log.Info("stripe-setup", "portal", "ensured")
	return nil
}
