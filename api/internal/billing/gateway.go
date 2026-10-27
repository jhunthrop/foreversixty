// api/internal/billing/gateway.go
//
// Package billing is the Stripe integration: checkout, the customer
// portal, the webhook, and the CLI jobs (stripe-setup, stripe-reconcile).
// It never writes the entitlements table directly — every write goes
// through entitlements.Store, so that package stays the single writer
// (see webhook.go's upsertSubscription). See
// docs/superpowers/specs/2026-09-21-entitlements-and-payments-design.md §2.
package billing

import (
	"context"
	"errors"
	"time"

	stripe "github.com/stripe/stripe-go/v82"
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
	// TrialPeriodDays and TrialEnd put a trial on subscription_data. At
	// most one is set; the real gateway rejects both.
	TrialPeriodDays int64
	TrialEnd        *time.Time
}

// ProductTaxCodeSaaSPersonal is Stripe's tax code for software as a
// service sold for personal use, which is what a Premium or guild plan
// is. Managed Payments (on by default for the account, Stripe as
// merchant of record) refuses a Checkout whose product carries no tax
// code, and Stripe Tax needs one to classify the sale.
const ProductTaxCodeSaaSPersonal = "txcd_10103001"

// PriceSpec is one (product, price) pair stripe-setup ensures exists
// (spec §2.1's table).
type PriceSpec struct {
	ProductID       string
	ProductName     string
	TaxCode         string // ProductTaxCodeSaaSPersonal for every product we sell
	LookupKey       string
	Interval        string // "month" | "year"
	UnitAmountCents int64
}

// PortalSpec is the Customer Portal configuration stripe-setup ensures
// exists (spec §2.5 left it to a Dashboard click-through; code keeps
// test and live mode identical). Products lists what a subscriber may
// switch between, by lookup key.
type PortalSpec struct {
	// Key is the fs_portal metadata value that names this configuration;
	// CreatePortalSession resolves a configuration by it.
	Key               string
	Headline          string
	PrivacyPolicyURL  string
	TermsOfServiceURL string
	DefaultReturnURL  string
	Products          []PortalProduct
}

// PortalProduct is one product the portal lets a subscriber move within:
// monthly to yearly and back.
type PortalProduct struct {
	ProductID  string
	LookupKeys []string
}

// PortalCall is one CreatePortalSession invocation, recorded by
// FakeGateway for test assertions.
type PortalCall struct {
	CustomerID       string
	ReturnURL        string
	ConfigurationKey string
}

// PriceChangePreview is what a subscription price change would charge:
// the prorated amount due now, the amount of the next regular invoice,
// and the end of the current period.
//
// AmountDueCents is Stripe's amount_due for the preview invoice, so it
// includes tax and any customer balance; NextAmountCents is the new
// price's pre-tax unit amount. The two are therefore not on the same
// footing and a caller must label them accordingly. ProrationDate is the
// instant the proration was computed at; pass it to ChangePrice so the
// amount charged equals the amount shown.
type PriceChangePreview struct {
	AmountDueCents  int64
	NextAmountCents int64
	PeriodEnd       time.Time
	ProrationDate   time.Time
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
	// PriceLookupKey is the lookup key of the subscription's single item.
	PriceLookupKey string
	// ItemLookupKeys are the lookup keys of every item on the subscription
	// (empty for an item whose price has none), so a caller can refuse a
	// subscription that carries a price as a second item.
	ItemLookupKeys []string
	// CancelAtPeriodEnd is true while the subscription is set to end at
	// its period end.
	CancelAtPeriodEnd bool
}

// Gateway is every Stripe API call this package makes. StripeGateway
// wraps stripe-go's unified client; FakeGateway is an in-memory
// implementation of the same interface — every test in this codebase
// uses FakeGateway, never StripeGateway, so no test ever reaches the
// real Stripe API (the coordinator's lane constraint).
type Gateway interface {
	CreateCustomer(ctx context.Context, email string) (customerID string, err error)
	CreateCheckoutSession(ctx context.Context, p CheckoutParams) (CheckoutSession, error)
	// CreatePortalSession opens the portal configuration whose fs_portal
	// metadata equals configurationKey; a missing configuration is an error.
	CreatePortalSession(ctx context.Context, customerID, returnURL, configurationKey string) (portalURL string, err error)
	// GetSubscription re-fetches a Subscription fresh by id — the one
	// call every webhook event handler makes instead of trusting the
	// event's own embedded snapshot (spec RULING 8).
	GetSubscription(ctx context.Context, id string) (*stripe.Subscription, error)
	PriceIDForLookupKey(ctx context.Context, lookupKey string) (priceID string, err error)
	// EnsurePrice is stripe-setup's only call: create the product/price
	// when spec.LookupKey resolves to none, a no-op otherwise (spec
	// §2.1 — never mutates an existing price's amount).
	EnsurePrice(ctx context.Context, spec PriceSpec) error
	// EnsurePortalConfiguration is stripe-setup's other call: create the
	// account's default Customer Portal configuration from spec when
	// none exists, a no-op otherwise, so a Dashboard edit made later is
	// never overwritten.
	EnsurePortalConfiguration(ctx context.Context, spec PortalSpec) error
	// CancelSubscriptionAtPeriodEnd sets cancel_at_period_end on a live
	// Stripe subscription — the newcomer side of the duplicate-
	// subscription guard (entitlements.OutcomeDuplicate,
	// webhook.go's upsertSubscription): the officer who
	// accidentally paid twice keeps what they already paid for through
	// the current period; the operator's runbook (spec §2.7) covers
	// refunding that charge by hand.
	CancelSubscriptionAtPeriodEnd(ctx context.Context, subscriptionID string) error
	// ListSubscriptions lists every subscription Stripe currently
	// considers live against productID — stripe-reconcile's
	// two-directional check (spec §2.8): Stripe → our database, the
	// direction the original design deferred and this fix implements.
	ListSubscriptions(ctx context.Context, productID string) ([]StripeSubscriptionSummary, error)
	// PreviewPriceChange previews moving the subscription's single item to
	// priceID with an immediate prorated invoice.
	PreviewPriceChange(ctx context.Context, subscriptionID, priceID string) (PriceChangePreview, error)
	// ChangePrice moves the subscription's single item to priceID, invoicing
	// the proration at once and failing rather than leaving a payment pending.
	// prorationDate is the preview's ProrationDate; the zero time means now.
	ChangePrice(ctx context.Context, subscriptionID, priceID string, prorationDate time.Time) error
	// ArchivePrice sets active=false on the price with that lookup key. It is
	// idempotent: a missing or already archived price is not an error.
	ArchivePrice(ctx context.Context, lookupKey string) error
	// DisableDefaultPortalSwitching turns subscription_update off on the
	// account's default portal configuration (the one stripe-setup does not
	// own), so no portal session can switch plans across products.
	DisableDefaultPortalSwitching(ctx context.Context) error
}
