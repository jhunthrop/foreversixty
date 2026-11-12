package creators

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

// TestStripeSandbox runs the real gateway against Stripe's Connect sandbox, only when
// CREATORS_STRIPE_SANDBOX_KEY holds a test key (never a live one; never in CI). It makes an Express
// account and its onboarding link, reads the fully onboarded test creator, and pays it $0.50 twice
// with one transfer group: the second call must answer the same transfer.
func TestStripeSandbox(t *testing.T) {
	key := os.Getenv("CREATORS_STRIPE_SANDBOX_KEY")
	if key == "" {
		t.Skip("CREATORS_STRIPE_SANDBOX_KEY not set")
	}
	if !strings.HasPrefix(key, "sk_test_") && !strings.HasPrefix(key, "rk_test_") {
		t.Fatal("refusing a key that is not a test key")
	}
	creator := os.Getenv("CREATORS_STRIPE_SANDBOX_ACCOUNT")
	// In test the sandbox key stands for both the subscription key and the payouts key.
	g := NewStripeGateway(key, key)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	account, err := g.CreateExpressAccount(ctx, time.Now().UnixNano())
	if err != nil {
		t.Fatal(err)
	}
	link, err := g.AccountLink(ctx, account, "https://foreversixty.gg/creators/onboarding?done=1", "https://foreversixty.gg/creators/onboarding?refresh=1")
	if err != nil || !strings.HasPrefix(link, "https://connect.stripe.com/") {
		t.Fatalf("account link: %v (%d chars)", err, len(link))
	}
	if facts, err := g.Account(ctx, account); err != nil || StripeStatusOf(facts) == StatusReady {
		t.Fatalf("a new account reads ready: %+v %v", facts, err)
	}
	// The probe invoice of 2026-11-08 (spec section 2): $4.50 paid, $0.43 fee, $2.00 refunded.
	if invoice := os.Getenv("CREATORS_STRIPE_SANDBOX_INVOICE"); invoice != "" {
		m, err := g.InvoiceMoney(ctx, invoice)
		if err != nil || m.AmountCents != 450 || m.FeeCents != 43 || m.NetCents != 407 || m.TaxInFees {
			t.Fatalf("invoice money = %+v %v", m, err)
		}
		refunds, err := g.Refunds(ctx, m.ChargeID)
		if err != nil || len(refunds) != 1 || refunds[0].AmountCents != 200 {
			t.Fatalf("refunds = %+v %v", refunds, err)
		}
	}
	if creator == "" {
		return
	}
	facts, err := g.Account(ctx, creator)
	if err != nil || StripeStatusOf(facts) != StatusReady {
		t.Fatalf("the onboarded test creator = %+v %v", facts, err)
	}
	if _, err := g.AvailableUSD(ctx); err != nil {
		t.Fatalf("balance: %v", err)
	}
	if _, err := g.LoginLink(ctx, creator); err != nil {
		t.Fatalf("login link: %v", err)
	}
	group := "creator-sandbox-" + time.Now().UTC().Format("20060102150405")
	params := TransferParams{AmountCents: 50, Destination: creator, Group: group, Metadata: map[string]string{"test": "creators"}}
	first, err := g.CreateTransfer(ctx, params)
	if err != nil {
		t.Fatal(err)
	}
	again, err := g.CreateTransfer(ctx, params)
	if err != nil || again != first {
		t.Fatalf("second transfer = %s %v, want %s", again, err, first)
	}
	found, err := g.FindTransfer(ctx, group)
	if err != nil || found != first {
		t.Fatalf("find = %s %v, want %s", found, err, first)
	}
}
