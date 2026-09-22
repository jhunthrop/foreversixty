// api/internal/billing/fake_gateway_test.go
package billing

import (
	"context"
	"errors"
	"testing"

	stripe "github.com/stripe/stripe-go/v82"
)

func TestFakeGatewaySatisfiesGateway(t *testing.T) {
	var g Gateway = &FakeGateway{}
	ctx := context.Background()
	id, err := g.CreateCustomer(ctx, "a@example.com")
	if err != nil || id == "" {
		t.Fatalf("CreateCustomer: %v, %v", id, err)
	}
	if _, err := g.PriceIDForLookupKey(ctx, LookupKeyPremiumMonthly); err != ErrPriceNotFound {
		t.Fatalf("PriceIDForLookupKey before EnsurePrice: %v, want ErrPriceNotFound", err)
	}
	if err := g.EnsurePrice(ctx, PriceSpec{LookupKey: LookupKeyPremiumMonthly}); err != nil {
		t.Fatal(err)
	}
	priceID, err := g.PriceIDForLookupKey(ctx, LookupKeyPremiumMonthly)
	if err != nil || priceID == "" {
		t.Fatalf("PriceIDForLookupKey after EnsurePrice: %v, %v", priceID, err)
	}
	session, err := g.CreateCheckoutSession(ctx, CheckoutParams{CustomerID: id, PriceID: priceID})
	if err != nil || session.URL == "" || session.SessionID == "" || session.ExpiresAt.IsZero() {
		t.Fatalf("CreateCheckoutSession: %+v, %v", session, err)
	}
	if err := g.CancelSubscriptionAtPeriodEnd(ctx, "sub_1"); err != nil {
		t.Fatal(err)
	}
	if _, err := g.ListSubscriptions(ctx, ProductIDPremium); err != nil {
		t.Fatal(err)
	}
	portalURL, err := g.CreatePortalSession(ctx, id, "https://foreversixty.gg/account")
	if err != nil || portalURL == "" {
		t.Fatalf("CreatePortalSession: %v, %v", portalURL, err)
	}
}

func TestFakeGatewayReturnsConfiguredSubscription(t *testing.T) {
	f := &FakeGateway{Subscriptions: map[string]*stripe.Subscription{
		"sub_1": {ID: "sub_1", Status: stripe.SubscriptionStatusActive},
	}}
	sub, err := f.GetSubscription(context.Background(), "sub_1")
	if err != nil || sub.Status != stripe.SubscriptionStatusActive {
		t.Fatalf("GetSubscription: %+v, %v", sub, err)
	}
	if _, err := f.GetSubscription(context.Background(), "sub_missing"); err == nil {
		t.Fatal("an unregistered subscription id should error, not panic")
	}
}

func TestFakeGatewayErrPropagatesToEveryMethod(t *testing.T) {
	want := errors.New("upstream down")
	f := &FakeGateway{Err: want}
	if _, err := f.CreateCustomer(context.Background(), "a@example.com"); err != want {
		t.Fatal(err)
	}
	if _, err := f.CreateCheckoutSession(context.Background(), CheckoutParams{}); err != want {
		t.Fatal(err)
	}
	if _, err := f.CreatePortalSession(context.Background(), "cus_1", ""); err != want {
		t.Fatal(err)
	}
	if err := f.CancelSubscriptionAtPeriodEnd(context.Background(), "sub_1"); err != want {
		t.Fatal(err)
	}
	if _, err := f.ListSubscriptions(context.Background(), ProductIDPremium); err != want {
		t.Fatal(err)
	}
}

func TestFakeGatewayRecordsCancelAndListSubscriptions(t *testing.T) {
	f := &FakeGateway{ListSubscriptionsResult: map[string][]StripeSubscriptionSummary{
		ProductIDGuild: {{ID: "sub_orphan", Status: "active"}},
	}}
	if err := f.CancelSubscriptionAtPeriodEnd(context.Background(), "sub_dupe"); err != nil {
		t.Fatal(err)
	}
	if len(f.CanceledAtPeriodEnd) != 1 || f.CanceledAtPeriodEnd[0] != "sub_dupe" {
		t.Fatalf("CanceledAtPeriodEnd = %v", f.CanceledAtPeriodEnd)
	}
	subs, err := f.ListSubscriptions(context.Background(), ProductIDGuild)
	if err != nil || len(subs) != 1 || subs[0].ID != "sub_orphan" {
		t.Fatalf("ListSubscriptions = %+v, %v", subs, err)
	}
	if subs, err := f.ListSubscriptions(context.Background(), ProductIDPremium); err != nil || len(subs) != 0 {
		t.Fatalf("ListSubscriptions for an unconfigured product = %+v, %v, want empty", subs, err)
	}
}
