// api/internal/billing/fake_gateway.go
package billing

import (
	"context"
	"fmt"
	"sync"

	stripe "github.com/stripe/stripe-go/v82"
)

// FakeGateway is the in-memory Gateway every test in this codebase uses.
// Zero value is usable; tests populate the exported fields directly
// before exercising a handler.
type FakeGateway struct {
	mu sync.Mutex

	// Err, when non-nil, is returned by every method instead of doing
	// anything, for testing an upstream-failure path.
	Err error

	nextCustomerID int
	// Subscriptions is looked up by GetSubscription — tests register a
	// hand-built *stripe.Subscription fixture here before calling a
	// handler that re-fetches it.
	Subscriptions map[string]*stripe.Subscription
	// Prices maps a lookup key to a price id — pre-populate to simulate
	// stripe-setup having already run; EnsurePrice adds to it.
	Prices map[string]string

	Customers    []string // emails passed to CreateCustomer, in order
	Checkouts    []CheckoutParams
	Portals      []PortalCall
	EnsuredSpecs []PriceSpec

	NextCheckoutURL string
	NextPortalURL   string
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

func (f *FakeGateway) CreateCheckoutSession(_ context.Context, p CheckoutParams) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return "", f.Err
	}
	f.Checkouts = append(f.Checkouts, p)
	if f.NextCheckoutURL != "" {
		return f.NextCheckoutURL, nil
	}
	return "https://checkout.stripe.com/fake/session", nil
}

func (f *FakeGateway) CreatePortalSession(_ context.Context, customerID, returnURL string) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return "", f.Err
	}
	f.Portals = append(f.Portals, PortalCall{CustomerID: customerID, ReturnURL: returnURL})
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
	sub, ok := f.Subscriptions[id]
	if !ok {
		return nil, fmt.Errorf("billing: fake gateway: no such subscription %s", id)
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
	if _, ok := f.Prices[spec.LookupKey]; !ok {
		f.Prices[spec.LookupKey] = "price_fake_" + spec.LookupKey
	}
	return nil
}
