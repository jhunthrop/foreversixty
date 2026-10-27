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
	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/config"
)

// productNames are the Stripe product names, by offer.
var productNames = map[catalog.Offer]string{
	catalog.OfferPremium:   "Forever Sixty Premium",
	catalog.OfferGuild:     "Forever Sixty Guild",
	catalog.OfferRaidGuild: "Forever Sixty Raid Guild",
}

// stripeIntervals maps the catalog's intervals to Stripe's recurring ones.
var stripeIntervals = map[catalog.Interval]string{catalog.Monthly: "month", catalog.Yearly: "year"}

// priceSpecs is derived from the catalog, the only place an amount, a
// lookup key or a product id is written (Ruling D: fixed product ids so
// this is idempotent with no Search API dependency).
var priceSpecs = derivePriceSpecs(catalog.Prices())

func derivePriceSpecs(prices []catalog.Price) []billing.PriceSpec {
	specs := make([]billing.PriceSpec, 0, len(prices))
	for _, p := range prices {
		specs = append(specs, billing.PriceSpec{
			ProductID: p.Offer.ProductID(), ProductName: productNames[p.Offer],
			TaxCode: billing.ProductTaxCodeSaaSPersonal, LookupKey: p.LookupKey,
			Interval: stripeIntervals[p.Interval], UnitAmountCents: p.AmountCents,
		})
	}
	return specs
}

// portalSpecs is the six Customer Portal configurations stripe-setup
// ensures (one per family, so a founder keeps the founder prices when
// switching monthly to yearly): the site's own legal pages, a return to
// the account page, and the family's two prices under its one product.
func portalSpecs(publicBaseURL string) []billing.PortalSpec {
	families := catalog.PortalFamilies()
	specs := make([]billing.PortalSpec, 0, len(families))
	for _, family := range families {
		keys := catalog.FamilyLookupKeys(family)
		if len(keys) == 0 {
			panic("stripe-setup: portal family " + family + " has no lookup keys in the catalog")
		}
		price, ok := catalog.Lookup(keys[0])
		if !ok {
			panic("stripe-setup: portal family " + family + " lists " + keys[0] + ", which the catalog does not know")
		}
		specs = append(specs, billing.PortalSpec{
			Key:               family,
			Headline:          "Forever Sixty",
			PrivacyPolicyURL:  publicBaseURL + "/privacy",
			TermsOfServiceURL: publicBaseURL + "/terms",
			DefaultReturnURL:  publicBaseURL + "/account",
			Products:          []billing.PortalProduct{{ProductID: price.Offer.ProductID(), LookupKeys: keys}},
		})
	}
	return specs
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
	for _, spec := range portalSpecs(cfg.PublicBaseURL) {
		if err := gw.EnsurePortalConfiguration(ctx, spec); err != nil {
			return fmt.Errorf("stripe-setup: portal %s: %w", spec.Key, err)
		}
		log.Info("stripe-setup", "portal", spec.Key)
	}
	if err := gw.DisableDefaultPortalSwitching(ctx); err != nil {
		return fmt.Errorf("stripe-setup: default portal: %w", err)
	}
	log.Info("stripe-setup", "portal", "default switching off")
	return nil
}
