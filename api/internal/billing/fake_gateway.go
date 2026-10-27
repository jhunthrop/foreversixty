// api/internal/billing/fake_gateway.go
package billing

import (
	"context"
	"fmt"
	"sync"
	"time"

	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

// fakeCheckoutExpiry is CreateCheckoutSession's default ExpiresAt when a
// test does not set NextCheckoutExpiresAt — long enough that ordinary
// tests never trip PendingGuildCheckout's expiry check by accident.
const fakeCheckoutExpiry = 24 * time.Hour

// FakeGateway is the in-memory Gateway every test in this codebase uses.
// Zero value is usable; tests populate the exported fields directly
// before exercising a handler.
type FakeGateway struct {
	mu sync.Mutex

	// Err, when non-nil, is returned by every method instead of doing
	// anything, for testing an upstream-failure path.
	Err error
	// CancelErr, when non-nil, is returned by CancelSubscriptionAtPeriodEnd
	// alone, for testing a failed duplicate cancel.
	CancelErr error
	// ChangePriceErr, when non-nil, is returned by ChangePrice alone.
	ChangePriceErr error
	// NextPreview is PreviewPriceChange's canned answer.
	NextPreview PriceChangePreview

	// archivedKeys are lookup keys ArchivePrice archived; EnsurePrice treats
	// them as present, like Stripe does.
	archivedKeys map[string]bool

	nextCustomerID int
	nextSessionID  int
	// Subscriptions is looked up by GetSubscription — tests register a
	// hand-built *stripe.Subscription fixture here before calling a
	// handler that re-fetches it.
	Subscriptions map[string]*stripe.Subscription
	// Prices maps a lookup key to a price id — pre-populate to simulate
	// stripe-setup having already run; EnsurePrice adds to it.
	Prices map[string]string
	// EnsuredPortals records every EnsurePortalConfiguration call.
	EnsuredPortals []PortalSpec
	// ListSubscriptionsResult is ListSubscriptions' answer, keyed by
	// product id — a test populates this directly (independent of
	// Subscriptions, which GetSubscription reads) since the whole point
	// of an orphan-subscription test is a Stripe subscription our
	// database has never heard of.
	ListSubscriptionsResult map[string][]StripeSubscriptionSummary

	Customers             []string // emails passed to CreateCustomer, in order
	Checkouts             []CheckoutParams
	Portals               []PortalCall
	PriceChanges          []PriceChangeCall
	Previews              []PriceChangeCall // PreviewPriceChange calls (PriceID is the target price)
	DisabledDefaultPortal int               // DisableDefaultPortalSwitching calls
	Archived              []string          // lookup keys passed to ArchivePrice, in order
	EnsuredSpecs          []PriceSpec
	CanceledAtPeriodEnd   []string // subscription ids passed to CancelSubscriptionAtPeriodEnd, in order
	Resumed               []string // subscription ids passed to ResumeSubscription, in order
	// GetSubscriptionCalls records the ids passed to GetSubscription, in order.
	GetSubscriptionCalls []string
	// OnGetSubscription, when set, runs after each GetSubscription call is
	// recorded and before the registered object is returned, so it may
	// mutate that object. call counts from 1. It runs under the gateway
	// mutex and must not call back into the gateway.
	OnGetSubscription func(id string, call int)

	NextCheckoutURL       string
	NextCheckoutExpiresAt time.Time
	NextPortalURL         string
}

func (f *FakeGateway) CreateCustomer(_ context.Context, email string) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return "", f.Err
	}
	f.nextCustomerID++
	f.Customers = append(f.Customers, email)
	return fmt.Sprintf("cus_fake_%d", f.nextCustomerID), nil
}

func (f *FakeGateway) CreateCheckoutSession(_ context.Context, p CheckoutParams) (CheckoutSession, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return CheckoutSession{}, f.Err
	}
	f.Checkouts = append(f.Checkouts, p)
	f.nextSessionID++
	url := f.NextCheckoutURL
	if url == "" {
		url = "https://checkout.stripe.com/fake/session"
	}
	expiresAt := f.NextCheckoutExpiresAt
	if expiresAt.IsZero() {
		expiresAt = time.Now().Add(fakeCheckoutExpiry)
	}
	return CheckoutSession{
		URL:       url,
		SessionID: fmt.Sprintf("cs_fake_%d", f.nextSessionID),
		ExpiresAt: expiresAt,
	}, nil
}

// PriceChangeCall is one ChangePrice invocation.
type PriceChangeCall struct {
	SubscriptionID string
	PriceID        string
	ProrationDate  time.Time
}

func (f *FakeGateway) CreatePortalSession(_ context.Context, customerID, returnURL, configurationKey string) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return "", f.Err
	}
	f.Portals = append(f.Portals, PortalCall{CustomerID: customerID, ReturnURL: returnURL, ConfigurationKey: configurationKey})
	if f.NextPortalURL != "" {
		return f.NextPortalURL, nil
	}
	return "https://billing.stripe.com/fake/session", nil
}

func (f *FakeGateway) GetSubscription(_ context.Context, id string) (*stripe.Subscription, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return nil, f.Err
	}
	f.GetSubscriptionCalls = append(f.GetSubscriptionCalls, id)
	if f.OnGetSubscription != nil {
		f.OnGetSubscription(id, len(f.GetSubscriptionCalls))
	}
	sub, ok := f.Subscriptions[id]
	if !ok {
		return nil, fmt.Errorf("billing: fake gateway: %s: %w", id, ErrSubscriptionNotFound)
	}
	return sub, nil
}

func (f *FakeGateway) PriceIDForLookupKey(_ context.Context, lookupKey string) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return "", f.Err
	}
	id, ok := f.Prices[lookupKey]
	if !ok {
		return "", ErrPriceNotFound
	}
	return id, nil
}

func (f *FakeGateway) CancelSubscriptionAtPeriodEnd(_ context.Context, subscriptionID string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return f.Err
	}
	if f.CancelErr != nil {
		return f.CancelErr
	}
	f.CanceledAtPeriodEnd = append(f.CanceledAtPeriodEnd, subscriptionID)
	return nil
}

func (f *FakeGateway) ResumeSubscription(_ context.Context, subscriptionID string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return f.Err
	}
	f.Resumed = append(f.Resumed, subscriptionID)
	return nil
}

func (f *FakeGateway) ListSubscriptions(_ context.Context, productID string) ([]StripeSubscriptionSummary, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return nil, f.Err
	}
	return f.ListSubscriptionsResult[productID], nil
}

// EnsurePortalConfiguration records the spec; the fake has no portal.
func (f *FakeGateway) EnsurePortalConfiguration(_ context.Context, spec PortalSpec) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return f.Err
	}
	f.EnsuredPortals = append(f.EnsuredPortals, spec)
	return nil
}

func (f *FakeGateway) EnsurePrice(_ context.Context, spec PriceSpec) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return f.Err
	}
	f.EnsuredSpecs = append(f.EnsuredSpecs, spec)
	if f.Prices == nil {
		f.Prices = map[string]string{}
	}
	if _, ok := f.Prices[spec.LookupKey]; !ok && !f.archivedKeys[spec.LookupKey] {
		f.Prices[spec.LookupKey] = "price_fake_" + spec.LookupKey
	}
	return nil
}

func (f *FakeGateway) PreviewPriceChange(_ context.Context, subscriptionID, priceID string) (PriceChangePreview, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return PriceChangePreview{}, f.Err
	}
	f.Previews = append(f.Previews, PriceChangeCall{SubscriptionID: subscriptionID, PriceID: priceID})
	return f.NextPreview, nil
}

func (f *FakeGateway) ChangePrice(_ context.Context, subscriptionID, priceID string, prorationDate time.Time) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return f.Err
	}
	if f.ChangePriceErr != nil {
		return f.ChangePriceErr
	}
	f.PriceChanges = append(f.PriceChanges, PriceChangeCall{SubscriptionID: subscriptionID, PriceID: priceID, ProrationDate: prorationDate})
	f.swapRegisteredPrice(subscriptionID, priceID)
	return nil
}

// swapRegisteredPrice moves a registered subscription's single item onto
// priceID, as Stripe would, so a later GetSubscription reads the new price.
// Callers hold f.mu.
func (f *FakeGateway) swapRegisteredPrice(subscriptionID, priceID string) {
	sub, ok := f.Subscriptions[subscriptionID]
	if !ok || sub.Items == nil || len(sub.Items.Data) != 1 {
		return
	}
	for key, id := range f.Prices {
		if id != priceID {
			continue
		}
		price, _ := catalog.Lookup(key)
		sub.Items.Data[0].Price = &stripe.Price{ID: priceID, LookupKey: key, UnitAmount: price.AmountCents}
		return
	}
}

func (f *FakeGateway) ArchivePrice(_ context.Context, lookupKey string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return f.Err
	}
	f.Archived = append(f.Archived, lookupKey)
	if _, ok := f.Prices[lookupKey]; ok {
		delete(f.Prices, lookupKey)
		if f.archivedKeys == nil {
			f.archivedKeys = map[string]bool{}
		}
		f.archivedKeys[lookupKey] = true
	}
	return nil
}

func (f *FakeGateway) DisableDefaultPortalSwitching(_ context.Context) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return f.Err
	}
	f.DisabledDefaultPortal++
	return nil
}
