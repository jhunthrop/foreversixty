// api/cmd/api/stripe_archive.go
//
// stripe-archive-prices retires the catalog's unsold prices in Stripe once
// no live subscription still bills on them. It runs last in the launch
// pricing rollout and never opens the database.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"slices"

	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/config"
)

// endedSubscriptionStatuses are the statuses that can never bill again; any
// other status (active, trialing, past_due, and also unpaid, paused or
// incomplete, which can still be revived) keeps its price in use.
var endedSubscriptionStatuses = []string{"canceled", "incomplete_expired"}

// unsoldKeysByProduct groups the catalog's unsold lookup keys by product id.
func unsoldKeysByProduct() map[string][]string {
	byProduct := map[string][]string{}
	for _, p := range catalog.Prices() {
		if !p.Sold {
			product := p.Offer.ProductID()
			byProduct[product] = append(byProduct[product], p.LookupKey)
		}
	}
	return byProduct
}

// liveOnUnsoldPrices returns "subscription (status, price)" for every live
// subscription billing on an unsold price.
func liveOnUnsoldPrices(ctx context.Context, gw billing.Gateway) ([]string, error) {
	var live []string
	for product, keys := range unsoldKeysByProduct() {
		subs, err := gw.ListSubscriptions(ctx, product)
		if err != nil {
			return nil, fmt.Errorf("list subscriptions on %s: %w", product, err)
		}
		for _, s := range subs {
			if slices.Contains(keys, s.PriceLookupKey) && !slices.Contains(endedSubscriptionStatuses, s.Status) {
				live = append(live, fmt.Sprintf("%s (%s, %s)", s.ID, s.Status, s.PriceLookupKey))
			}
		}
	}
	slices.Sort(live)
	return live, nil
}

// archiveUnsoldPrices refuses while a live subscription uses an unsold
// price; otherwise it prints the keys and, only when apply is set, archives
// them.
func archiveUnsoldPrices(ctx context.Context, gw billing.Gateway, apply bool, out io.Writer) error {
	live, err := liveOnUnsoldPrices(ctx, gw)
	if err != nil {
		return err
	}
	if len(live) > 0 {
		for _, s := range live {
			fmt.Fprintf(out, "live subscription on an unsold price: %s\n", s)
		}
		return fmt.Errorf("stripe-archive-prices: refusing, %d live subscription(s) still use unsold prices", len(live))
	}
	verb := "would archive"
	if apply {
		verb = "archiving"
	}
	for _, p := range catalog.Prices() {
		if p.Sold {
			continue
		}
		fmt.Fprintf(out, "%s %s\n", verb, p.LookupKey)
		if apply {
			if err := gw.ArchivePrice(ctx, p.LookupKey); err != nil {
				return fmt.Errorf("stripe-archive-prices: %s: %w", p.LookupKey, err)
			}
		}
	}
	if !apply {
		fmt.Fprintln(out, "dry run: pass --apply to archive")
	}
	return nil
}

func runStripeArchivePrices(ctx context.Context, log *slog.Logger, args []string) error {
	fs := flag.NewFlagSet("stripe-archive-prices", flag.ContinueOnError)
	apply := fs.Bool("apply", false, "archive the prices (default is a dry run)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	cfg, err := config.Load(os.Getenv)
	if err != nil {
		return err
	}
	if err := cfg.ValidateStripeKeyEnvironment(); err != nil {
		return err
	}
	if !cfg.StripeConfigured() {
		return errors.New("stripe-archive-prices: STRIPE_SECRET_KEY and STRIPE_WEBHOOK_SECRET are required")
	}
	log.Info("stripe-archive-prices", "apply", *apply)
	return archiveUnsoldPrices(ctx, billing.NewStripeGateway(cfg.StripeSecretKey), *apply, os.Stdout)
}
