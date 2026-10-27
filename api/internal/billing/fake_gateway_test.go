// api/internal/billing/fake_gateway_test.go
package billing

import (
	"context"
	"errors"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	stripe "github.com/stripe/stripe-go/v82"
)

func TestFakeGatewaySatisfiesGateway(t *testing.T) {
	var g Gateway = &FakeGateway{}
	ctx := context.Background()
	id, err := g.CreateCustomer(ctx, "a@example.com")
	if err != nil || id == "" {
		t.Fatalf("CreateCustomer: %v, %v", id, err)
	}
	if _, err := g.PriceIDForLookupKey(ctx, catalog.LookupKeyPremiumMonthly); err != ErrPriceNotFound {
		t.Fatalf("PriceIDForLookupKey before EnsurePrice: %v, want ErrPriceNotFound", err)
	}
	if err := g.EnsurePrice(ctx, PriceSpec{LookupKey: catalog.LookupKeyPremiumMonthly}); err != nil {
		t.Fatal(err)
	}
	priceID, err := g.PriceIDForLookupKey(ctx, catalog.LookupKeyPremiumMonthly)
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
	if _, err := g.ListSubscriptions(ctx, catalog.ProductIDPremium); err != nil {
		t.Fatal(err)
	}
	portalURL, err := g.CreatePortalSession(ctx, id, "https://foreversixty.gg/account", "premium")
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
	if _, err := f.CreatePortalSession(context.Background(), "cus_1", "", "premium"); err != want {
		t.Fatal(err)
	}
	if err := f.CancelSubscriptionAtPeriodEnd(context.Background(), "sub_1"); err != want {
		t.Fatal(err)
	}
	if _, err := f.ListSubscriptions(context.Background(), catalog.ProductIDPremium); err != want {
		t.Fatal(err)
	}
}

func TestFakeGatewayRecordsCancelAndListSubscriptions(t *testing.T) {
	f := &FakeGateway{ListSubscriptionsResult: map[string][]StripeSubscriptionSummary{
		catalog.ProductIDGuild: {{ID: "sub_orphan", Status: "active"}},
	}}
	if err := f.CancelSubscriptionAtPeriodEnd(context.Background(), "sub_dupe"); err != nil {
		t.Fatal(err)
	}
	if len(f.CanceledAtPeriodEnd) != 1 || f.CanceledAtPeriodEnd[0] != "sub_dupe" {
		t.Fatalf("CanceledAtPeriodEnd = %v", f.CanceledAtPeriodEnd)
	}
	subs, err := f.ListSubscriptions(context.Background(), catalog.ProductIDGuild)
	if err != nil || len(subs) != 1 || subs[0].ID != "sub_orphan" {
		t.Fatalf("ListSubscriptions = %+v, %v", subs, err)
	}
	if subs, err := f.ListSubscriptions(context.Background(), catalog.ProductIDPremium); err != nil || len(subs) != 0 {
		t.Fatalf("ListSubscriptions for an unconfigured product = %+v, %v, want empty", subs, err)
	}
}

func TestFakeGatewayRecordsPortalKeyTrialAndPriceChanges(t *testing.T) {
	ctx := context.Background()
	f := &FakeGateway{NextPreview: PriceChangePreview{AmountDueCents: 1000, NextAmountCents: 2500}}
	if _, err := f.CreatePortalSession(ctx, "cus_1", "https://x/account", "guild_founder"); err != nil {
		t.Fatal(err)
	}
	if len(f.Portals) != 1 || f.Portals[0].ConfigurationKey != "guild_founder" {
		t.Fatalf("Portals = %+v", f.Portals)
	}
	if _, err := f.CreateCheckoutSession(ctx, CheckoutParams{PriceID: "price_1", TrialPeriodDays: 7}); err != nil {
		t.Fatal(err)
	}
	if f.Checkouts[0].TrialPeriodDays != 7 {
		t.Fatalf("Checkouts = %+v", f.Checkouts)
	}
	pv, err := f.PreviewPriceChange(ctx, "sub_1", "price_2")
	if err != nil || pv.AmountDueCents != 1000 || pv.NextAmountCents != 2500 {
		t.Fatalf("PreviewPriceChange = %+v, %v", pv, err)
	}
	if err := f.ChangePrice(ctx, "sub_1", "price_2"); err != nil {
		t.Fatal(err)
	}
	if len(f.PriceChanges) != 1 || f.PriceChanges[0] != (PriceChangeCall{"sub_1", "price_2"}) {
		t.Fatalf("PriceChanges = %+v", f.PriceChanges)
	}
	if err := f.ArchivePrice(ctx, catalog.LookupKeyPremiumMonthly); err != nil || len(f.Archived) != 1 {
		t.Fatalf("Archived = %v, %v", f.Archived, err)
	}
	want := errors.New("declined")
	f.ChangePriceErr = want
	if err := f.ChangePrice(ctx, "sub_1", "price_3"); err != want {
		t.Fatalf("ChangePrice = %v, want %v", err, want)
	}
	if len(f.PriceChanges) != 1 {
		t.Fatal("a failed change must not be recorded")
	}
}
