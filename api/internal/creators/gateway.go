package creators

import (
	"context"
	"time"
)

// AccountFacts is what the pages need of a creator's Express account.
type AccountFacts struct {
	ID               string
	DetailsSubmitted bool
	PayoutsEnabled   bool
	TransfersActive  bool
	CurrentlyDue     int
}

// InvoiceMoney is what a paid invoice earns a share on.
type InvoiceMoney struct {
	ChargeID    string
	PaidAt      time.Time
	AmountCents int64
	TaxCents    int64
	FeeCents    int64
	NetCents    int64
	// TaxInFees: the balance transaction's fees already took the tax (Stripe as merchant of record).
	TaxInFees bool
}

// Refund is one refund of a charge.
type Refund struct {
	ID          string
	AmountCents int64
	Created     time.Time
}

// TransferParams is one payout's transfer. Group doubles as the idempotency key.
type TransferParams struct {
	AmountCents int64
	Destination string
	Group       string
	Metadata    map[string]string
}

// Gateway is every Stripe call the program makes. StripeGateway is the real one; tests use a
// fake. Every call runs on the deployment's own key (the sandbox in tests and previews).
type Gateway interface {
	// CreateExpressAccount makes the creator's Express account (transfers only) and answers its id.
	CreateExpressAccount(ctx context.Context, creatorID int64) (string, error)
	// AccountLink is a fresh onboarding link for the account.
	AccountLink(ctx context.Context, accountID, returnURL, refreshURL string) (string, error)
	// LoginLink is a one-time link into the account's Express dashboard. Never stored.
	LoginLink(ctx context.Context, accountID string) (string, error)
	// Account re-reads the account.
	Account(ctx context.Context, accountID string) (AccountFacts, error)
	// InvoiceMoney reads a paid invoice's charge and its balance transaction.
	InvoiceMoney(ctx context.Context, invoiceID string) (InvoiceMoney, error)
	// Refunds lists a charge's refunds.
	Refunds(ctx context.Context, chargeID string) ([]Refund, error)
	// FindTransfer answers the transfer already made for a group, "" when there is none.
	FindTransfer(ctx context.Context, group string) (string, error)
	// CreateTransfer makes the transfer, idempotent on the group, and answers its id.
	CreateTransfer(ctx context.Context, p TransferParams) (string, error)
}
