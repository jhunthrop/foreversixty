// api/internal/billing/gateway.go
//
// Package billing is the Stripe integration: checkout, the customer
// portal, the webhook, and the CLI jobs (stripe-setup, stripe-reconcile).
// It never writes the entitlements table directly — every write goes
// through entitlements.Store, so that package stays the single writer
// (see webhook.go's upsertFromSubscription). See
// docs/superpowers/specs/2026-09-21-entitlements-and-payments-design.md §2.
package billing

import (
	"context"
	"errors"
	"time"

	stripe "github.com/stripe/stripe-go/v82"
)

// The four lookup keys spec §2.1's table fixes — the only Stripe price
// identifiers application code ever hardcodes; every checkout resolves a
// Price id from one of these at request time.
const (
	LookupKeyPremiumMonthly = "premium_monthly"
	LookupKeyPremiumYearly  = "premium_yearly"
	LookupKeyGuildMonthly   = "guild_monthly"
	LookupKeyGuildYearly    = "guild_yearly"
)

// ErrPriceNotFound is PriceIDForLookupKey's answer when no active Price
// carries that lookup key yet (stripe-setup has not run, or is stale).
var ErrPriceNotFound = errors.New("billing: no active price for that lookup key")

// CheckoutParams is what CreateCheckoutSession turns into a Stripe
// Checkout Session (spec §2.2). Metadata is set identically on both the
// session and subscription_data.metadata by the real implementation —
// callers pass it once.
type CheckoutParams struct {
	CustomerID        string
	PriceID           string
	ClientReferenceID string
	Metadata          map[string]string
	SuccessURL        string
	CancelURL         string
}

// ProductIDPremium and ProductIDGuild are the two fixed Stripe Product
// ids stripe-setup creates (cmd/api/stripe_setup.go's priceSpecs) — the
// single source of truth for both that command and stripe-reconcile's
// two-directional check (spec §2.8, security review fix 2026-09-21),
// which lists Stripe's own live subscriptions against exactly these ids.
const (
	ProductIDPremium = "prod_fs_premium"
	ProductIDGuild   = "prod_fs_guild"
)

// PriceSpec is one (product, price) pair stripe-setup ensures exists
// (spec §2.1's table).
type PriceSpec struct {
	ProductID       string
	ProductName     string
	LookupKey       string
	Interval        string // "month" | "year"
	UnitAmountCents int64
}

// PortalCall is one CreatePortalSession invocation, recorded by
// FakeGateway for test assertions.
type PortalCall struct {
	CustomerID string
	ReturnURL  string
}

// CheckoutSession is what CreateCheckoutSession returns: the checkout
// URL, plus the Stripe Session id and its own expiry — both needed by
// the guild-plan path's pending_checkouts row (security review fix,
// 2026-09-21, handler.go's guildCheckout): SessionID is what
// DeletePendingCheckout later matches against, and ExpiresAt is what
// that row's own expires_at mirrors, so an abandoned checkout stops
// blocking a later real one on the same guild without waiting on a
// webhook that will never arrive.
type CheckoutSession struct {
	URL       string
	SessionID string
	ExpiresAt time.Time
}

// StripeSubscriptionSummary is what ListSubscriptions returns for each
// live Stripe subscription against one of our own products — just
// enough for stripe-reconcile's orphan check (spec §2.8's last
// paragraph, security review fix 2026-09-21), never the full
// Subscription object.
type StripeSubscriptionSummary struct {
	ID     string
	Status string
}

// Gateway is every Stripe API call this package makes. StripeGateway
// wraps stripe-go's unified client; FakeGateway is an in-memory
// implementation of the same interface — every test in this codebase
// uses FakeGateway, never StripeGateway, so no test ever reaches the
// real Stripe API (the coordinator's lane constraint).
type Gateway interface {
	CreateCustomer(ctx context.Context, email string) (customerID string, err error)
	CreateCheckoutSession(ctx context.Context, p CheckoutParams) (CheckoutSession, error)
	CreatePortalSession(ctx context.Context, customerID, returnURL string) (portalURL string, err error)
	// GetSubscription re-fetches a Subscription fresh by id — the one
	// call every webhook event handler makes instead of trusting the
	// event's own embedded snapshot (spec RULING 8).
	GetSubscription(ctx context.Context, id string) (*stripe.Subscription, error)
	PriceIDForLookupKey(ctx context.Context, lookupKey string) (priceID string, err error)
	// EnsurePrice is stripe-setup's only call: create the product/price
	// when spec.LookupKey resolves to none, a no-op otherwise (spec
	// §2.1 — never mutates an existing price's amount).
	EnsurePrice(ctx context.Context, spec PriceSpec) error
	// CancelSubscriptionAtPeriodEnd sets cancel_at_period_end on a live
	// Stripe subscription — the newcomer side of the duplicate-
	// subscription guard (entitlements.UpsertResult.Duplicate,
	// webhook.go's upsertFromSubscription): the officer who
	// accidentally paid twice keeps what they already paid for through
	// the current period; the operator's runbook (spec §2.7) covers
	// refunding that charge by hand.
	CancelSubscriptionAtPeriodEnd(ctx context.Context, subscriptionID string) error
	// ListSubscriptions lists every subscription Stripe currently
	// considers live against productID — stripe-reconcile's
	// two-directional check (spec §2.8): Stripe → our database, the
	// direction the original design deferred and this fix implements.
	ListSubscriptions(ctx context.Context, productID string) ([]StripeSubscriptionSummary, error)
}
