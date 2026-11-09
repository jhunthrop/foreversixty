package creators

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	stripe "github.com/stripe/stripe-go/v82"
)

// StripeGateway is the real Gateway over stripe-go's clients, on two keys: the subscription key
// (STRIPE_SECRET_KEY) reads a player's invoices and refunds, and the payouts key
// (STRIPE_PAYOUTS_KEY, a restricted key with Connect and Transfers) makes every Connect and
// transfer call, so the subscription key is never widened. It is never built in a test; the
// sandbox check in stripe_sandbox_test.go runs it only when a sandbox key is in the environment.
type StripeGateway struct {
	client  *stripe.Client
	payouts *stripe.Client
}

// ErrPayoutsKeyMissing: no STRIPE_PAYOUTS_KEY, so no Connect or transfer call can be made.
var ErrPayoutsKeyMissing = errors.New("creators: STRIPE_PAYOUTS_KEY is not set")

// NewStripeGateway binds the gateway to the subscription key and the payouts key (” when unset).
func NewStripeGateway(secretKey, payoutsKey string) *StripeGateway {
	g := &StripeGateway{client: stripe.NewClient(secretKey)}
	if payoutsKey != "" {
		g.payouts = stripe.NewClient(payoutsKey)
	}
	return g
}

// connect is the payouts key's client, or ErrPayoutsKeyMissing.
func (g *StripeGateway) connect() (*stripe.Client, error) {
	if g.payouts == nil {
		return nil, ErrPayoutsKeyMissing
	}
	return g.payouts, nil
}

func resourceMissing(err error) bool {
	var se *stripe.Error
	return errors.As(err, &se) && se.Code == stripe.ErrorCodeResourceMissing
}

func (g *StripeGateway) CreateExpressAccount(ctx context.Context, creatorID int64) (string, error) {
	c, err := g.connect()
	if err != nil {
		return "", err
	}
	params := &stripe.AccountCreateParams{
		Type:    stripe.String(string(stripe.AccountTypeExpress)),
		Country: stripe.String("US"),
		Capabilities: &stripe.AccountCreateCapabilitiesParams{
			Transfers: &stripe.AccountCreateCapabilitiesTransfersParams{Requested: stripe.Bool(true)},
		},
		Metadata: map[string]string{"fs_creator": strconv.FormatInt(creatorID, 10)},
	}
	params.SetIdempotencyKey("creator-account-" + strconv.FormatInt(creatorID, 10))
	a, err := c.V1Accounts.Create(ctx, params)
	if err != nil {
		return "", fmt.Errorf("creators: create account: %w", err)
	}
	return a.ID, nil
}

func (g *StripeGateway) AccountLink(ctx context.Context, accountID, returnURL, refreshURL string) (string, error) {
	c, err := g.connect()
	if err != nil {
		return "", err
	}
	l, err := c.V1AccountLinks.Create(ctx, &stripe.AccountLinkCreateParams{
		Account:    stripe.String(accountID),
		Type:       stripe.String("account_onboarding"),
		ReturnURL:  stripe.String(returnURL),
		RefreshURL: stripe.String(refreshURL),
	})
	if err != nil {
		return "", fmt.Errorf("creators: account link: %w", err)
	}
	return l.URL, nil
}

func (g *StripeGateway) LoginLink(ctx context.Context, accountID string) (string, error) {
	c, err := g.connect()
	if err != nil {
		return "", err
	}
	l, err := c.V1LoginLinks.Create(ctx, &stripe.LoginLinkCreateParams{Account: stripe.String(accountID)})
	if err != nil {
		return "", fmt.Errorf("creators: login link: %w", err)
	}
	return l.URL, nil
}

func (g *StripeGateway) CheckConnect(ctx context.Context) error {
	c, err := g.connect()
	if err != nil {
		return err
	}
	params := &stripe.AccountListParams{}
	params.Limit = stripe.Int64(1)
	for _, err := range c.V1Accounts.List(ctx, params) {
		if err != nil {
			return fmt.Errorf("creators: list connected accounts: %w", err)
		}
		break
	}
	return nil
}

func (g *StripeGateway) Account(ctx context.Context, accountID string) (AccountFacts, error) {
	c, err := g.connect()
	if err != nil {
		return AccountFacts{}, err
	}
	a, err := c.V1Accounts.GetByID(ctx, accountID, nil)
	if err != nil {
		return AccountFacts{}, fmt.Errorf("creators: read account: %w", err)
	}
	facts := AccountFacts{ID: a.ID, DetailsSubmitted: a.DetailsSubmitted, PayoutsEnabled: a.PayoutsEnabled}
	if a.Capabilities != nil {
		facts.TransfersActive = a.Capabilities.Transfers == stripe.AccountCapabilityStatusActive
	}
	if a.Requirements != nil {
		facts.CurrentlyDue = len(a.Requirements.CurrentlyDue)
	}
	return facts, nil
}

// InvoiceMoney reads the invoice's tax, its one paid payment, and that payment's charge and
// balance transaction (Stripe allows four expansion levels, so the payment intent is a second read).
func (g *StripeGateway) InvoiceMoney(ctx context.Context, invoiceID string) (InvoiceMoney, error) {
	inv, err := g.client.V1Invoices.Retrieve(ctx, invoiceID, nil)
	if err != nil {
		return InvoiceMoney{}, fmt.Errorf("creators: read invoice: %w", err)
	}
	if inv.AmountPaid <= 0 {
		return InvoiceMoney{}, nil // a trial's $0 invoice has no payment to read
	}
	var tax int64
	for _, t := range inv.TotalTaxes {
		tax += t.Amount
	}
	piID, paidAt, err := g.paidIntent(ctx, invoiceID)
	if err != nil {
		return InvoiceMoney{}, err
	}
	params := &stripe.PaymentIntentRetrieveParams{}
	params.AddExpand("latest_charge.balance_transaction")
	pi, err := g.client.V1PaymentIntents.Retrieve(ctx, piID, params)
	if err != nil {
		return InvoiceMoney{}, fmt.Errorf("creators: read payment: %w", err)
	}
	ch := pi.LatestCharge
	if ch == nil || ch.BalanceTransaction == nil {
		return InvoiceMoney{}, fmt.Errorf("creators: invoice %s: payment has no settled charge yet", invoiceID)
	}
	bt := ch.BalanceTransaction
	money := InvoiceMoney{ChargeID: ch.ID, PaidAt: paidAt, AmountCents: ch.Amount, TaxCents: tax, FeeCents: bt.Fee, NetCents: bt.Net}
	for _, f := range bt.FeeDetails {
		if f.Type == "tax" {
			money.TaxInFees = true
		}
	}
	return money, nil
}

// paidIntent is the invoice's paid payment intent and when it was paid.
func (g *StripeGateway) paidIntent(ctx context.Context, invoiceID string) (string, time.Time, error) {
	list := g.client.V1InvoicePayments.List(ctx, &stripe.InvoicePaymentListParams{
		Invoice: stripe.String(invoiceID), Status: stripe.String("paid"),
	})
	for p, err := range list {
		if err != nil {
			return "", time.Time{}, fmt.Errorf("creators: invoice payments: %w", err)
		}
		if p.Payment == nil || p.Payment.PaymentIntent == nil {
			continue
		}
		paid := time.Unix(p.Created, 0).UTC()
		if p.StatusTransitions != nil && p.StatusTransitions.PaidAt > 0 {
			paid = time.Unix(p.StatusTransitions.PaidAt, 0).UTC()
		}
		return p.Payment.PaymentIntent.ID, paid, nil
	}
	return "", time.Time{}, fmt.Errorf("creators: invoice %s has no paid payment", invoiceID)
}

func (g *StripeGateway) Refunds(ctx context.Context, chargeID string) ([]Refund, error) {
	var out []Refund
	for r, err := range g.client.V1Refunds.List(ctx, &stripe.RefundListParams{Charge: stripe.String(chargeID)}) {
		if err != nil {
			return nil, fmt.Errorf("creators: refunds: %w", err)
		}
		if r.Status != stripe.RefundStatusSucceeded {
			continue
		}
		out = append(out, Refund{ID: r.ID, AmountCents: r.Amount, Created: time.Unix(r.Created, 0).UTC()})
	}
	return out, nil
}

func (g *StripeGateway) AvailableUSD(ctx context.Context) (int64, error) {
	c, err := g.connect()
	if err != nil {
		return 0, err
	}
	b, err := c.V1Balance.Retrieve(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("creators: read balance: %w", err)
	}
	var cents int64
	for _, a := range b.Available {
		if a.Currency == stripe.CurrencyUSD {
			cents += a.Amount
		}
	}
	return cents, nil
}

func (g *StripeGateway) FindTransfer(ctx context.Context, group string) (string, error) {
	c, err := g.connect()
	if err != nil {
		return "", err
	}
	for t, err := range c.V1Transfers.List(ctx, &stripe.TransferListParams{TransferGroup: stripe.String(group)}) {
		if err != nil {
			return "", fmt.Errorf("creators: find transfer: %w", err)
		}
		return t.ID, nil
	}
	return "", nil
}

func (g *StripeGateway) CreateTransfer(ctx context.Context, p TransferParams) (string, error) {
	c, err := g.connect()
	if err != nil {
		return "", err
	}
	params := &stripe.TransferCreateParams{
		Amount:        stripe.Int64(p.AmountCents),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		Destination:   stripe.String(p.Destination),
		TransferGroup: stripe.String(p.Group),
		Metadata:      p.Metadata,
	}
	params.SetIdempotencyKey(p.Group)
	t, err := c.V1Transfers.Create(ctx, params)
	var se *stripe.Error
	if errors.As(err, &se) && se.Code == stripe.ErrorCodeBalanceInsufficient {
		return "", ErrInsufficientBalance
	}
	if err != nil {
		return "", fmt.Errorf("creators: create transfer: %w", err)
	}
	return t.ID, nil
}
