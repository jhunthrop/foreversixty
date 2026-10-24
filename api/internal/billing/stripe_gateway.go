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

func (g *StripeGateway) CreateCheckoutSession(ctx context.Context, p CheckoutParams) (CheckoutSession, error) {
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
		return CheckoutSession{}, fmt.Errorf("billing: create checkout session: %w", err)
	}
	return CheckoutSession{URL: sess.URL, SessionID: sess.ID, ExpiresAt: time.Unix(sess.ExpiresAt, 0).UTC()}, nil
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

// ensureProduct is EnsurePrice's product half: create the product by
// its fixed id when it is missing, and set its tax code when it carries
// none or a different one, so a product made before tax codes were
// required still sells (Managed Payments refuses a Checkout otherwise).
func (g *StripeGateway) ensureProduct(ctx context.Context, spec PriceSpec) error {
	product, err := g.client.V1Products.Retrieve(ctx, spec.ProductID, nil)
	if err != nil {
		params := &stripe.ProductCreateParams{ID: stripe.String(spec.ProductID), Name: stripe.String(spec.ProductName)}
		if spec.TaxCode != "" {
			params.TaxCode = stripe.String(spec.TaxCode)
		}
		if _, cerr := g.client.V1Products.Create(ctx, params); cerr != nil {
			return fmt.Errorf("billing: create product %s: %w", spec.ProductID, cerr)
		}
		return nil
	}
	if spec.TaxCode == "" || (product.TaxCode != nil && product.TaxCode.ID == spec.TaxCode) {
		return nil
	}
	if _, err := g.client.V1Products.Update(ctx, spec.ProductID, &stripe.ProductUpdateParams{
		TaxCode: stripe.String(spec.TaxCode),
	}); err != nil {
		return fmt.Errorf("billing: set tax code on product %s: %w", spec.ProductID, err)
	}
	return nil
}

// EnsurePortalConfiguration makes the account's default Customer Portal
// configuration match spec: created when there is none, otherwise its
// legal links, return page and the prices a subscriber may switch
// between are set again (the Dashboard is not the source of truth, spec
// RULING 5). A subscriber can read invoices, change the card, cancel at
// the period's end, and move between a product's own prices.
func (g *StripeGateway) EnsurePortalConfiguration(ctx context.Context, spec PortalSpec) error {
	products, err := g.portalProducts(ctx, spec)
	if err != nil {
		return err
	}
	for existing, err := range g.client.V1BillingPortalConfigurations.List(ctx, &stripe.BillingPortalConfigurationListParams{
		IsDefault: stripe.Bool(true), Active: stripe.Bool(true),
	}) {
		if err != nil {
			return fmt.Errorf("billing: list portal configurations: %w", err)
		}
		if _, err := g.client.V1BillingPortalConfigurations.Update(ctx, existing.ID, &stripe.BillingPortalConfigurationUpdateParams{
			BusinessProfile: &stripe.BillingPortalConfigurationUpdateBusinessProfileParams{
				Headline:          stripe.String(spec.Headline),
				PrivacyPolicyURL:  stripe.String(spec.PrivacyPolicyURL),
				TermsOfServiceURL: stripe.String(spec.TermsOfServiceURL),
			},
			DefaultReturnURL: stripe.String(spec.DefaultReturnURL),
			Features: &stripe.BillingPortalConfigurationUpdateFeaturesParams{
				SubscriptionUpdate: &stripe.BillingPortalConfigurationUpdateFeaturesSubscriptionUpdateParams{
					Enabled: stripe.Bool(true), DefaultAllowedUpdates: []*string{stripe.String("price")},
					Products: products, ProrationBehavior: stripe.String("create_prorations"),
				},
			},
		}); err != nil {
			return fmt.Errorf("billing: update portal configuration %s: %w", existing.ID, err)
		}
		return nil
	}
	createProducts := make([]*stripe.BillingPortalConfigurationCreateFeaturesSubscriptionUpdateProductParams, 0, len(products))
	for _, p := range products {
		createProducts = append(createProducts, &stripe.BillingPortalConfigurationCreateFeaturesSubscriptionUpdateProductParams{
			Product: p.Product, Prices: p.Prices,
		})
	}
	if _, err := g.client.V1BillingPortalConfigurations.Create(ctx, &stripe.BillingPortalConfigurationCreateParams{
		BusinessProfile: &stripe.BillingPortalConfigurationCreateBusinessProfileParams{
			Headline:          stripe.String(spec.Headline),
			PrivacyPolicyURL:  stripe.String(spec.PrivacyPolicyURL),
			TermsOfServiceURL: stripe.String(spec.TermsOfServiceURL),
		},
		DefaultReturnURL: stripe.String(spec.DefaultReturnURL),
		Features: &stripe.BillingPortalConfigurationCreateFeaturesParams{
			CustomerUpdate: &stripe.BillingPortalConfigurationCreateFeaturesCustomerUpdateParams{
				Enabled: stripe.Bool(true), AllowedUpdates: []*string{stripe.String("email"), stripe.String("address")},
			},
			InvoiceHistory:      &stripe.BillingPortalConfigurationCreateFeaturesInvoiceHistoryParams{Enabled: stripe.Bool(true)},
			PaymentMethodUpdate: &stripe.BillingPortalConfigurationCreateFeaturesPaymentMethodUpdateParams{Enabled: stripe.Bool(true)},
			SubscriptionCancel: &stripe.BillingPortalConfigurationCreateFeaturesSubscriptionCancelParams{
				Enabled: stripe.Bool(true), Mode: stripe.String("at_period_end"),
				CancellationReason: &stripe.BillingPortalConfigurationCreateFeaturesSubscriptionCancelCancellationReasonParams{
					Enabled: stripe.Bool(true),
					Options: []*string{stripe.String("too_expensive"), stripe.String("missing_features"), stripe.String("unused"), stripe.String("other")},
				},
			},
			SubscriptionUpdate: &stripe.BillingPortalConfigurationCreateFeaturesSubscriptionUpdateParams{
				Enabled: stripe.Bool(true), DefaultAllowedUpdates: []*string{stripe.String("price")},
				Products: createProducts, ProrationBehavior: stripe.String("create_prorations"),
			},
		},
	}); err != nil {
		return fmt.Errorf("billing: create portal configuration: %w", err)
	}
	return nil
}

// portalProducts resolves spec's lookup keys to the price ids the portal
// lets a subscriber switch between, product by product.
func (g *StripeGateway) portalProducts(ctx context.Context, spec PortalSpec) ([]*stripe.BillingPortalConfigurationUpdateFeaturesSubscriptionUpdateProductParams, error) {
	products := make([]*stripe.BillingPortalConfigurationUpdateFeaturesSubscriptionUpdateProductParams, 0, len(spec.Products))
	for _, p := range spec.Products {
		prices := make([]*string, 0, len(p.LookupKeys))
		for _, key := range p.LookupKeys {
			id, err := g.PriceIDForLookupKey(ctx, key)
			if err != nil {
				return nil, fmt.Errorf("billing: portal price %s: %w", key, err)
			}
			prices = append(prices, stripe.String(id))
		}
		products = append(products, &stripe.BillingPortalConfigurationUpdateFeaturesSubscriptionUpdateProductParams{
			Product: stripe.String(p.ProductID), Prices: prices,
		})
	}
	return products, nil
}

// CancelSubscriptionAtPeriodEnd implements Gateway.CancelSubscriptionAtPeriodEnd.
func (g *StripeGateway) CancelSubscriptionAtPeriodEnd(ctx context.Context, subscriptionID string) error {
	if _, err := g.client.V1Subscriptions.Update(ctx, subscriptionID, &stripe.SubscriptionUpdateParams{
		CancelAtPeriodEnd: stripe.Bool(true),
	}); err != nil {
		return fmt.Errorf("billing: cancel subscription at period end %s: %w", subscriptionID, err)
	}
	return nil
}

// ListSubscriptions implements Gateway.ListSubscriptions. Stripe's
// Subscriptions List API filters by Price, not Product, so productID is
// applied client-side against each subscription's own line items —
// every subscription this integration creates has exactly one item
// (spec §2.2), so this is a cheap per-subscription check, not an N+1
// query. No Status filter is passed: by default the List API already
// excludes canceled subscriptions, which is exactly "currently live."
func (g *StripeGateway) ListSubscriptions(ctx context.Context, productID string) ([]StripeSubscriptionSummary, error) {
	params := &stripe.SubscriptionListParams{}
	params.AddExpand("data.items.data.price.product")
	var out []StripeSubscriptionSummary
	for sub, err := range g.client.V1Subscriptions.List(ctx, params) {
		if err != nil {
			return nil, fmt.Errorf("billing: list subscriptions for product %s: %w", productID, err)
		}
		if !subscriptionHasProduct(sub, productID) {
			continue
		}
		out = append(out, StripeSubscriptionSummary{ID: sub.ID, Status: string(sub.Status)})
	}
	return out, nil
}

// subscriptionHasProduct reports whether any of sub's line items bills
// against productID — ListSubscriptions' client-side filter.
func subscriptionHasProduct(sub *stripe.Subscription, productID string) bool {
	if sub.Items == nil {
		return false
	}
	for _, item := range sub.Items.Data {
		if item.Price != nil && item.Price.Product != nil && item.Price.Product.ID == productID {
			return true
		}
	}
	return false
}

// EnsurePrice is stripe-setup's price call (Ruling D): retrieve-or-create
// the product by its fixed id, give it its tax code, then list-or-create
// the price by its lookup key. Never mutates an existing price's amount.
func (g *StripeGateway) EnsurePrice(ctx context.Context, spec PriceSpec) error {
	if err := g.ensureProduct(ctx, spec); err != nil {
		return err
	}
	if _, err := g.PriceIDForLookupKey(ctx, spec.LookupKey); err == nil {
		return nil
	} else if err != ErrPriceNotFound {
		return err
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
