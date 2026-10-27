// api/cmd/api/stripe_archive.go
//
// stripe-archive-prices retires the two legacy Premium prices in Stripe once
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

// retiredPriceKeys are the two legacy Premium prices this command archives,
// written out on purpose: a catalog flag flipping must never widen what it
// archives to a launch, founder or Guild price.
var retiredPriceKeys = []string{catalog.LookupKeyPremiumMonthly, catalog.LookupKeyPremiumYearly}

// retiredProductID is the product that holds the retired prices.
const retiredProductID = catalog.ProductIDPremium

// endedSubscriptionStatuses are the statuses that can never bill again; any
// other status (active, trialing, past_due, and also unpaid, paused or
// incomplete, which can still be revived) keeps its price in use.
var endedSubscriptionStatuses = []string{"canceled", "incomplete_expired"}

// itemKeys are the lookup keys of every item on the subscription.
func itemKeys(s billing.StripeSubscriptionSummary) []string {
	if len(s.ItemLookupKeys) > 0 {
		return s.ItemLookupKeys
	}
	return []string{s.PriceLookupKey}
}

// blocksArchive explains why a not-yet-ended subscription keeps the retired
// prices in use, or returns "" when it does not. A subscription with more
// than one item is refused outright: ours have exactly one.
func blocksArchive(s billing.StripeSubscriptionSummary) string {
	keys := itemKeys(s)
	if len(keys) > 1 {
		return "has more than one item"
	}
	if slices.Contains(retiredPriceKeys, keys[0]) {
		return "bills on " + keys[0]
	}
	return ""
}

// liveOnRetiredPrices returns one line per subscription that blocks archiving.
func liveOnRetiredPrices(ctx context.Context, gw billing.Gateway) ([]string, error) {
	subs, err := gw.ListSubscriptions(ctx, retiredProductID)
	if err != nil {
		return nil, fmt.Errorf("stripe-archive-prices: list subscriptions on %s: %w", retiredProductID, err)
	}
	var live []string
	for _, s := range subs {
		if slices.Contains(endedSubscriptionStatuses, s.Status) {
			continue
		}
		if reason := blocksArchive(s); reason != "" {
			live = append(live, fmt.Sprintf("%s (%s, %s)", s.ID, s.Status, reason))
		}
	}
	slices.Sort(live)
	return live, nil
}

// archiveRetiredPrices refuses while a live subscription uses a retired
// price; otherwise it prints the keys and, only when apply is set, archives
// them, reporting each one after Stripe accepted it.
func archiveRetiredPrices(ctx context.Context, gw billing.Gateway, apply bool, out io.Writer) error {
	live, err := liveOnRetiredPrices(ctx, gw)
	if err != nil {
		return err
	}
	if len(live) > 0 {
		for _, s := range live {
			fmt.Fprintf(out, "live subscription blocks archiving: %s\n", s)
		}
		return fmt.Errorf("stripe-archive-prices: refusing, %d live subscription(s) still use the retired prices", len(live))
	}
	for _, key := range retiredPriceKeys {
		if !apply {
			fmt.Fprintf(out, "would archive %s\n", key)
			continue
		}
		if err := gw.ArchivePrice(ctx, key); err != nil {
			return fmt.Errorf("stripe-archive-prices: %s: %w", key, err)
		}
		fmt.Fprintf(out, "archived %s\n", key)
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
	if cfg.StripeSecretKey == "" {
		return errors.New("stripe-archive-prices: STRIPE_SECRET_KEY is required")
	}
	log.Info("stripe-archive-prices", "apply", *apply)
	return archiveRetiredPrices(ctx, billing.NewStripeGateway(cfg.StripeSecretKey), *apply, os.Stdout)
}
