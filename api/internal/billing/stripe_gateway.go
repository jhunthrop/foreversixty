// api/internal/billing/stripe_gateway.go
package billing

import (
	"context"
	"errors"
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

	mu          sync.Mutex
	cache       map[string]cachedPrice
	portalCache map[string]cachedPrice
}

// NewStripeGateway builds a Gateway bound to secretKey (test or
// live — the caller's config.ValidateStripeKeyEnvironment already
// checked which).
func NewStripeGateway(secretKey string) *StripeGateway {
	return &StripeGateway{client: stripe.NewClient(secretKey), cache: map[string]cachedPrice{}, portalCache: map[string]cachedPrice{}}
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

// checkoutSessionLifetime is how long a Checkout page stays valid: the
// minimum Stripe allows, so an abandoned one stops blocking a guild soon.
const checkoutSessionLifetime = 30 * time.Minute

func (g *StripeGateway) CreateCheckoutSession(ctx context.Context, p CheckoutParams) (CheckoutSession, error) {
	subscriptionData, err := checkoutSubscriptionData(p)
	if err != nil {
		return CheckoutSession{}, err
	}
	sess, err := g.client.V1CheckoutSessions.Create(ctx, &stripe.CheckoutSessionCreateParams{
		Mode:              stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		Customer:          stripe.String(p.CustomerID),
		ClientReferenceID: stripe.String(p.ClientReferenceID),
		Metadata:          p.Metadata,
		LineItems: []*stripe.CheckoutSessionCreateLineItemParams{
			{Price: stripe.String(p.PriceID), Quantity: stripe.Int64(1)},
		},
		SubscriptionData:    subscriptionData,
		AutomaticTax:        &stripe.CheckoutSessionCreateAutomaticTaxParams{Enabled: stripe.Bool(true)},
		AllowPromotionCodes: stripe.Bool(true),
		// Stripe's shortest expiry. A guild checkout that is abandoned on
		// the Checkout page holds a pending_checkouts row until the session
		// expires, and every other officer sees "already has the plan"
		// until then; thirty minutes instead of Stripe's default day.
		ExpiresAt:  stripe.Int64(time.Now().Add(checkoutSessionLifetime).Unix()),
		SuccessURL: stripe.String(p.SuccessURL),
		CancelURL:  stripe.String(p.CancelURL),
	})
	if err != nil {
		return CheckoutSession{}, fmt.Errorf("billing: create checkout session: %w", err)
	}
	return CheckoutSession{URL: sess.URL, SessionID: sess.ID, ExpiresAt: time.Unix(sess.ExpiresAt, 0).UTC()}, nil
}

// checkoutSubscriptionData is the Checkout session's subscription_data:
// the metadata and, when asked, one trial (days or an end instant).
func checkoutSubscriptionData(p CheckoutParams) (*stripe.CheckoutSessionCreateSubscriptionDataParams, error) {
	if p.TrialPeriodDays != 0 && p.TrialEnd != nil {
		return nil, errors.New("billing: checkout takes TrialPeriodDays or TrialEnd, not both")
	}
	data := &stripe.CheckoutSessionCreateSubscriptionDataParams{Metadata: p.Metadata}
	if p.TrialPeriodDays != 0 {
		data.TrialPeriodDays = stripe.Int64(p.TrialPeriodDays)
	}
	if p.TrialEnd != nil {
		data.TrialEnd = stripe.Int64(p.TrialEnd.Unix())
	}
	return data, nil
}

func (g *StripeGateway) CreatePortalSession(ctx context.Context, customerID, returnURL, configurationKey string) (string, error) {
	configID, err := g.portalConfigurationID(ctx, configurationKey)
	if err != nil {
		return "", err
	}
	sess, err := g.client.V1BillingPortalSessions.Create(ctx, &stripe.BillingPortalSessionCreateParams{
		Customer: stripe.String(customerID), ReturnURL: stripe.String(returnURL),
		Configuration: stripe.String(configID),
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

// portalMetadataKey tags each portal configuration with its family key.
const portalMetadataKey = "fs_portal"

// errNoPortalConfiguration marks portalConfigurationID's "none tagged" answer.
var errNoPortalConfiguration = errors.New("billing: no portal configuration")

// portalConfigurationID finds the active configuration tagged fs_portal=key.
func (g *StripeGateway) portalConfigurationID(ctx context.Context, key string) (string, error) {
	if key == "" {
		return "", errors.New("billing: portal configuration key is empty")
	}
	g.mu.Lock()
	if c, ok := g.portalCache[key]; ok && time.Now().Before(c.expiresAt) {
		g.mu.Unlock()
		return c.id, nil
	}
	g.mu.Unlock()
	for cfg, err := range g.client.V1BillingPortalConfigurations.List(ctx, &stripe.BillingPortalConfigurationListParams{
		Active: stripe.Bool(true),
	}) {
		if err != nil {
			return "", fmt.Errorf("billing: list portal configurations: %w", err)
		}
		if cfg.Metadata[portalMetadataKey] == key {
			g.mu.Lock()
			g.portalCache[key] = cachedPrice{id: cfg.ID, expiresAt: time.Now().Add(priceCacheTTL)}
			g.mu.Unlock()
			return cfg.ID, nil
		}
	}
	return "", fmt.Errorf("%w tagged %s=%s (stripe-setup has not run)", errNoPortalConfiguration, portalMetadataKey, key)
}

// EnsurePortalConfiguration creates the portal configuration tagged
// fs_portal=spec.Key when none carries that key, and does nothing
// otherwise, so a Dashboard edit made later is never overwritten. A
// subscriber can read invoices, change the card, cancel at the period's
// end, and move between the family's own prices.
func (g *StripeGateway) EnsurePortalConfiguration(ctx context.Context, spec PortalSpec) error {
	if _, err := g.portalConfigurationID(ctx, spec.Key); err == nil {
		return nil
	} else if !errors.Is(err, errNoPortalConfiguration) {
		return err
	}
	products, err := g.portalProducts(ctx, spec)
	if err != nil {
		return err
	}
	if _, err := g.client.V1BillingPortalConfigurations.Create(ctx, &stripe.BillingPortalConfigurationCreateParams{
		BusinessProfile: &stripe.BillingPortalConfigurationCreateBusinessProfileParams{
			Headline:          stripe.String(spec.Headline),
			PrivacyPolicyURL:  stripe.String(spec.PrivacyPolicyURL),
			TermsOfServiceURL: stripe.String(spec.TermsOfServiceURL),
		},
		DefaultReturnURL: stripe.String(spec.DefaultReturnURL),
		Metadata:         map[string]string{portalMetadataKey: spec.Key},
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
				Products: products, ProrationBehavior: stripe.String("create_prorations"),
			},
		},
	}); err != nil {
		return fmt.Errorf("billing: create portal configuration %s: %w", spec.Key, err)
	}
	return nil
}

// portalProducts resolves spec's lookup keys to the price ids the portal
// lets a subscriber switch between, product by product.
func (g *StripeGateway) portalProducts(ctx context.Context, spec PortalSpec) ([]*stripe.BillingPortalConfigurationCreateFeaturesSubscriptionUpdateProductParams, error) {
	products := make([]*stripe.BillingPortalConfigurationCreateFeaturesSubscriptionUpdateProductParams, 0, len(spec.Products))
	for _, p := range spec.Products {
		prices := make([]*string, 0, len(p.LookupKeys))
		for _, key := range p.LookupKeys {
			id, err := g.PriceIDForLookupKey(ctx, key)
			if err != nil {
				return nil, fmt.Errorf("billing: portal price %s: %w", key, err)
			}
			prices = append(prices, stripe.String(id))
		}
		products = append(products, &stripe.BillingPortalConfigurationCreateFeaturesSubscriptionUpdateProductParams{
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
	// No expansion: an unexpanded price.product still decodes to a
	// Product carrying its id, which is all subscriptionHasProduct reads,
	// and data.items.data.price.product is five levels, one more than
	// Stripe expands (the first live reconcile failed on exactly that).
	var out []StripeSubscriptionSummary
	for sub, err := range g.client.V1Subscriptions.List(ctx, params) {
		if err != nil {
			return nil, fmt.Errorf("billing: list subscriptions for product %s: %w", productID, err)
		}
		if !subscriptionHasProduct(sub, productID) {
			continue
		}
		out = append(out, StripeSubscriptionSummary{ID: sub.ID, Status: string(sub.Status), PriceLookupKey: itemLookupKey(sub)})
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

// findPrice resolves a lookup key to its price whether active or archived.
func (g *StripeGateway) findPrice(ctx context.Context, lookupKey string) (*stripe.Price, error) {
	for price, err := range g.client.V1Prices.List(ctx, &stripe.PriceListParams{
		LookupKeys: []*string{stripe.String(lookupKey)},
	}) {
		if err != nil {
			return nil, fmt.Errorf("billing: list prices for %s: %w", lookupKey, err)
		}
		return price, nil
	}
	return nil, ErrPriceNotFound
}

// EnsurePrice is stripe-setup's price call (Ruling D): retrieve-or-create
// the product by its fixed id, give it its tax code, then list-or-create
// the price by its lookup key. Never mutates an existing price's amount.
func (g *StripeGateway) EnsurePrice(ctx context.Context, spec PriceSpec) error {
	if err := g.ensureProduct(ctx, spec); err != nil {
		return err
	}
	// An archived price still owns its lookup key (Stripe keeps keys unique
	// across archived prices), so it counts as present: recreating it would
	// fail, or with transfer_lookup_key would start selling it again.
	if _, err := g.findPrice(ctx, spec.LookupKey); err == nil {
		return nil
	} else if !errors.Is(err, ErrPriceNotFound) {
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

// itemLookupKey is the lookup key of sub's first line item, "" when none.
func itemLookupKey(sub *stripe.Subscription) string {
	if sub.Items == nil || len(sub.Items.Data) == 0 || sub.Items.Data[0].Price == nil {
		return ""
	}
	return sub.Items.Data[0].Price.LookupKey
}

// singleItem is the one line item every subscription of ours has.
func singleItem(sub *stripe.Subscription) (*stripe.SubscriptionItem, error) {
	if sub.Items == nil || len(sub.Items.Data) != 1 {
		return nil, fmt.Errorf("billing: subscription %s does not have exactly one item", sub.ID)
	}
	return sub.Items.Data[0], nil
}

// priceChangeProration is the proration behaviour of both the preview and
// the change, so the preview shows exactly what the change charges.
const priceChangeProration = "always_invoice"

// PreviewPriceChange implements Gateway.PreviewPriceChange with a preview
// invoice of the item swapped under always_invoice proration, computed at a
// fixed ProrationDate that ChangePrice then reuses.
func (g *StripeGateway) PreviewPriceChange(ctx context.Context, subscriptionID, priceID string) (PriceChangePreview, error) {
	sub, err := g.GetSubscription(ctx, subscriptionID)
	if err != nil {
		return PriceChangePreview{}, err
	}
	item, err := singleItem(sub)
	if err != nil {
		return PriceChangePreview{}, err
	}
	price, err := g.client.V1Prices.Retrieve(ctx, priceID, nil)
	if err != nil {
		return PriceChangePreview{}, fmt.Errorf("billing: get price %s: %w", priceID, err)
	}
	prorationDate := time.Now().UTC().Truncate(time.Second)
	inv, err := g.client.V1Invoices.CreatePreview(ctx, &stripe.InvoiceCreatePreviewParams{
		Subscription: stripe.String(subscriptionID),
		SubscriptionDetails: &stripe.InvoiceCreatePreviewSubscriptionDetailsParams{
			Items: []*stripe.InvoiceCreatePreviewSubscriptionDetailsItemParams{
				{ID: stripe.String(item.ID), Price: stripe.String(priceID)},
			},
			ProrationBehavior: stripe.String(priceChangeProration),
			ProrationDate:     stripe.Int64(prorationDate.Unix()),
		},
	})
	if err != nil {
		return PriceChangePreview{}, fmt.Errorf("billing: preview price change on %s: %w", subscriptionID, err)
	}
	return PriceChangePreview{
		AmountDueCents:  inv.AmountDue,
		NextAmountCents: price.UnitAmount,
		PeriodEnd:       time.Unix(item.CurrentPeriodEnd, 0).UTC(),
		ProrationDate:   prorationDate,
	}, nil
}

// ChangePrice implements Gateway.ChangePrice.
func (g *StripeGateway) ChangePrice(ctx context.Context, subscriptionID, priceID string, prorationDate time.Time) error {
	sub, err := g.GetSubscription(ctx, subscriptionID)
	if err != nil {
		return err
	}
	item, err := singleItem(sub)
	if err != nil {
		return err
	}
	params := &stripe.SubscriptionUpdateParams{
		Items:             []*stripe.SubscriptionUpdateItemParams{{ID: stripe.String(item.ID), Price: stripe.String(priceID)}},
		ProrationBehavior: stripe.String(priceChangeProration),
		PaymentBehavior:   stripe.String("error_if_incomplete"),
	}
	if !prorationDate.IsZero() {
		params.ProrationDate = stripe.Int64(prorationDate.Unix())
	}
	if _, err := g.client.V1Subscriptions.Update(ctx, subscriptionID, params); err != nil {
		return fmt.Errorf("billing: change price on subscription %s: %w", subscriptionID, err)
	}
	return nil
}

// ArchivePrice implements Gateway.ArchivePrice. An archived price keeps
// serving its existing subscribers; it only stops being sold. A missing or
// already archived price is already in the wanted state.
func (g *StripeGateway) ArchivePrice(ctx context.Context, lookupKey string) error {
	price, err := g.findPrice(ctx, lookupKey)
	if errors.Is(err, ErrPriceNotFound) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("billing: archive price %s: %w", lookupKey, err)
	}
	if !price.Active {
		return nil
	}
	if _, err := g.client.V1Prices.Update(ctx, price.ID, &stripe.PriceUpdateParams{Active: stripe.Bool(false)}); err != nil {
		return fmt.Errorf("billing: archive price %s: %w", lookupKey, err)
	}
	g.mu.Lock()
	delete(g.cache, lookupKey)
	g.mu.Unlock()
	return nil
}

// DisableDefaultPortalSwitching implements Gateway.DisableDefaultPortalSwitching.
// Configurations tagged fs_portal are ours and keep their switching; the
// untagged default is the legacy one with its Premium and Guild switch.
func (g *StripeGateway) DisableDefaultPortalSwitching(ctx context.Context) error {
	for cfg, err := range g.client.V1BillingPortalConfigurations.List(ctx, &stripe.BillingPortalConfigurationListParams{
		IsDefault: stripe.Bool(true), Active: stripe.Bool(true),
	}) {
		if err != nil {
			return fmt.Errorf("billing: list default portal configuration: %w", err)
		}
		if cfg.Metadata[portalMetadataKey] != "" {
			continue
		}
		if cfg.Features == nil || cfg.Features.SubscriptionUpdate == nil || !cfg.Features.SubscriptionUpdate.Enabled {
			continue
		}
		if _, err := g.client.V1BillingPortalConfigurations.Update(ctx, cfg.ID, &stripe.BillingPortalConfigurationUpdateParams{
			Features: &stripe.BillingPortalConfigurationUpdateFeaturesParams{
				SubscriptionUpdate: &stripe.BillingPortalConfigurationUpdateFeaturesSubscriptionUpdateParams{Enabled: stripe.Bool(false)},
			},
		}); err != nil {
			return fmt.Errorf("billing: disable switching on default portal configuration %s: %w", cfg.ID, err)
		}
	}
	return nil
}
