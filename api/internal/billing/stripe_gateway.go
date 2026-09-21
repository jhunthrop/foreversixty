// api/internal/billing/stripe_gateway.go
package billing

import (
	"context"
	"fmt"
	"sync"
	"time"

	stripe "github.com/stripe/stripe-go/v82"
)

// priceCacheTTL is how long StripeGateway caches a lookup-key → price id
// resolution in-process (spec §2.1: "the four values change only when
// the owner deliberately reprices, not per request").
const priceCacheTTL = 5 * time.Minute

type cachedPrice struct {
	id        string
	expiresAt time.Time
}

// StripeGateway is the real Gateway, wrapping stripe-go's unified
// client. Every method call reaches the live Stripe API (test or live
// mode, whichever the key belongs to) — this type is never constructed
// in a test.
type StripeGateway struct {
	client *stripe.Client

	mu    sync.Mutex
	cache map[string]cachedPrice
}

// NewStripeGateway builds a Gateway bound to secretKey (test or
// live — the caller's config.ValidateStripeKeyEnvironment already
// checked which).
func NewStripeGateway(secretKey string) *StripeGateway {
	return &StripeGateway{client: stripe.NewClient(secretKey), cache: map[string]cachedPrice{}}
}

func (g *StripeGateway) CreateCustomer(ctx context.Context, email string) (string, error) {
	params := &stripe.CustomerCreateParams{}
	if email != "" {
		params.Email = stripe.String(email)
	}
	c, err := g.client.V1Customers.Create(ctx, params)
	if err != nil {
		return "", fmt.Errorf("billing: create customer: %w", err)
	}
	return c.ID, nil
}

func (g *StripeGateway) CreateCheckoutSession(ctx context.Context, p CheckoutParams) (string, error) {
	sess, err := g.client.V1CheckoutSessions.Create(ctx, &stripe.CheckoutSessionCreateParams{
		Mode:              stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		Customer:          stripe.String(p.CustomerID),
		ClientReferenceID: stripe.String(p.ClientReferenceID),
		Metadata:          p.Metadata,
		LineItems: []*stripe.CheckoutSessionCreateLineItemParams{
			{Price: stripe.String(p.PriceID), Quantity: stripe.Int64(1)},
		},
		SubscriptionData:    &stripe.CheckoutSessionCreateSubscriptionDataParams{Metadata: p.Metadata},
		AutomaticTax:        &stripe.CheckoutSessionCreateAutomaticTaxParams{Enabled: stripe.Bool(true)},
		AllowPromotionCodes: stripe.Bool(true),
		SuccessURL:          stripe.String(p.SuccessURL),
		CancelURL:           stripe.String(p.CancelURL),
	})
	if err != nil {
		return "", fmt.Errorf("billing: create checkout session: %w", err)
	}
	return sess.URL, nil
}

func (g *StripeGateway) CreatePortalSession(ctx context.Context, customerID, returnURL string) (string, error) {
	sess, err := g.client.V1BillingPortalSessions.Create(ctx, &stripe.BillingPortalSessionCreateParams{
		Customer: stripe.String(customerID), ReturnURL: stripe.String(returnURL),
	})
	if err != nil {
		return "", fmt.Errorf("billing: create portal session: %w", err)
	}
	return sess.URL, nil
}

func (g *StripeGateway) GetSubscription(ctx context.Context, id string) (*stripe.Subscription, error) {
	sub, err := g.client.V1Subscriptions.Retrieve(ctx, id, nil)
	if err != nil {
		return nil, fmt.Errorf("billing: get subscription %s: %w", id, err)
	}
	return sub, nil
}

func (g *StripeGateway) PriceIDForLookupKey(ctx context.Context, lookupKey string) (string, error) {
	g.mu.Lock()
	if c, ok := g.cache[lookupKey]; ok && time.Now().Before(c.expiresAt) {
		g.mu.Unlock()
		return c.id, nil
	}
	g.mu.Unlock()

	for price, err := range g.client.V1Prices.List(ctx, &stripe.PriceListParams{
		LookupKeys: []*string{stripe.String(lookupKey)}, Active: stripe.Bool(true),
	}) {
		if err != nil {
			return "", fmt.Errorf("billing: list prices for %s: %w", lookupKey, err)
		}
		g.mu.Lock()
		g.cache[lookupKey] = cachedPrice{id: price.ID, expiresAt: time.Now().Add(priceCacheTTL)}
		g.mu.Unlock()
		return price.ID, nil
	}
	return "", ErrPriceNotFound
}

// EnsurePrice is stripe-setup's only call (Ruling D): retrieve-or-create
// the product by its fixed id, then list-or-create the price by its
// lookup key. Never mutates an existing price's amount.
func (g *StripeGateway) EnsurePrice(ctx context.Context, spec PriceSpec) error {
	if _, err := g.PriceIDForLookupKey(ctx, spec.LookupKey); err == nil {
		return nil
	} else if err != ErrPriceNotFound {
		return err
	}
	if _, err := g.client.V1Products.Retrieve(ctx, spec.ProductID, nil); err != nil {
		if _, cerr := g.client.V1Products.Create(ctx, &stripe.ProductCreateParams{
			ID: stripe.String(spec.ProductID), Name: stripe.String(spec.ProductName),
		}); cerr != nil {
			return fmt.Errorf("billing: create product %s: %w", spec.ProductID, cerr)
		}
	}
	if _, err := g.client.V1Prices.Create(ctx, &stripe.PriceCreateParams{
		Product: stripe.String(spec.ProductID), Currency: stripe.String("usd"),
		UnitAmount: stripe.Int64(spec.UnitAmountCents), LookupKey: stripe.String(spec.LookupKey),
		Recurring: &stripe.PriceCreateRecurringParams{Interval: stripe.String(spec.Interval)},
	}); err != nil {
		return fmt.Errorf("billing: create price %s: %w", spec.LookupKey, err)
	}
	return nil
}
