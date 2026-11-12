// api/internal/billing/creator_code.go
package billing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/webhook"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// Creators is the creator program's side of billing (creators.Service satisfies it): a creator's
// code turns a checkout's trial into a free month, and the webhook tells it what was bought,
// paid and refunded. Nil leaves codes refused and the program unaware of payments.
type Creators interface {
	CodeIsLive(ctx context.Context, code string) (bool, error)
	CodeTrialDays() int64
	CheckoutCompleted(ctx context.Context, userID int64, code string, at time.Time, invoiceID string) error
	InvoicePaid(ctx context.Context, userID int64, invoiceID string) error
	ChargeRefunded(ctx context.Context, chargeID string) error
	AccountUpdated(ctx context.Context, accountID string) error
}

// metadataCreatorCode is the subscription metadata key that carries a redeemed creator code.
const metadataCreatorCode = "creator_code"

var (
	errBadCreatorCode = &checkoutRefusal{status: http.StatusBadRequest, code: "creator_code",
		message: "that isn't a creator's code"}
	errTrialUsed = &checkoutRefusal{status: http.StatusConflict, code: "trial_used",
		message: "this account or guild already had its free trial, so a code can't give another"}
)

// creatorCodeTrial validates a code a checkout carries against the plan's default trial and
// answers the trial the checkout gets: a code that is not live or meets a spent trial refuses.
func (s *Service) creatorCodeTrial(ctx context.Context, code string, trialUsed bool, defaultDays int64) (*checkoutTrial, error) {
	d, err := s.decideTrial(ctx, code, trialUsed, defaultDays)
	if err != nil {
		return nil, err
	}
	switch d.code {
	case codeUnknown:
		return nil, errBadCreatorCode
	case codeTrialUsed:
		return nil, errTrialUsed
	}
	if d.days == 0 {
		return nil, nil
	}
	return &checkoutTrial{days: d.days}, nil
}

// RecordPremiumTrial notes that an account has used its one Premium trial; a second is ignored.
func (s *Store) RecordPremiumTrial(ctx context.Context, userID int64, subscriptionID string) error {
	if _, err := s.Pool.Exec(ctx, `insert into premium_trials (user_id, stripe_subscription_id) values ($1, $2)
		on conflict (user_id) do nothing`, userID, subscriptionID); err != nil {
		return fmt.Errorf("billing: record premium trial: %w", err)
	}
	return nil
}

// PremiumTrialSpent reports whether the account's one Premium trial is gone: it had the trial, or
// it already paid for Premium before (its first Premium was that subscription).
func (s *Store) PremiumTrialSpent(ctx context.Context, userID int64) (bool, error) {
	var spent bool
	err := s.Pool.QueryRow(ctx, `select exists (select 1 from premium_trials where user_id = $1)
		or exists (select 1 from entitlements where user_id = $1 and plan = 'premium' and source = 'stripe')`, userID).Scan(&spent)
	if err != nil {
		return false, fmt.Errorf("billing: premium trial spent %d: %w", userID, err)
	}
	return spent, nil
}

// creatorsCheckoutCompleted tells the program about a completed Checkout: the buyer and the code
// the subscription carries, and the session's first invoice.
func (s *Service) creatorsCheckoutCompleted(ctx context.Context, event stripe.Event, sub *stripe.Subscription) error {
	if s.Creators == nil {
		return nil
	}
	uid, ok := metadataUserID(sub)
	if !ok {
		return nil
	}
	at := time.Unix(event.Created, 0).UTC()
	// Logged, never failed: a retry would replay the whole checkout event, and the session's
	// invoice is recorded again by its own invoice.paid.
	if err := s.Creators.CheckoutCompleted(ctx, uid, sub.Metadata[metadataCreatorCode], at, stripeRefID(event.Data.Object["invoice"])); err != nil {
		s.logger().Error("billing", "op", "creators_checkout", "event", event.ID, "err", err)
	}
	return nil
}

// creatorsInvoicePaid tells the program about a paid invoice of an account's subscription.
func (s *Service) creatorsInvoicePaid(ctx context.Context, userID int64, invoiceID string) error {
	if s.Creators == nil || invoiceID == "" {
		return nil
	}
	return s.Creators.InvoicePaid(ctx, userID, invoiceID)
}

// onChargeRefunded is charge.refunded: a refund of a creator's player's payment becomes a line.
func (s *Service) onChargeRefunded(ctx context.Context, event stripe.Event) error {
	if s.Creators == nil {
		return nil
	}
	return s.Creators.ChargeRefunded(ctx, stripeRefID(event.Data.Object["id"]))
}

// onAccountUpdated is a connected account's account.updated (the Connect endpoint).
func (s *Service) onAccountUpdated(ctx context.Context, event stripe.Event) error {
	if s.Creators == nil {
		return nil
	}
	return s.Creators.AccountUpdated(ctx, event.Account)
}

// recordPremiumTrial is afterWritten's premium half: a subscription that started on a trial
// spends the account's one trial.
func (s *Service) recordPremiumTrial(ctx context.Context, p entitlements.StripeUpsert, sub *stripe.Subscription) error {
	if p.TrialEnd == nil {
		return nil
	}
	return s.Store.RecordPremiumTrial(ctx, p.BillingUserID, sub.ID)
}

// verifySecrets are the endpoint secrets the webhook accepts a signature from: the account's own
// endpoint and, when configured, the Connect endpoint on the same URL (account.updated).
func (s *Service) verifySecrets() []string {
	if s.ConnectWebhookSecret == "" {
		return []string{s.WebhookSecret}
	}
	return []string{s.WebhookSecret, s.ConnectWebhookSecret}
}

// errNoSecret answers a webhook with no secret configured to check it against.
var errNoSecret = errors.New("billing: no webhook secret")

// constructEvent verifies the signature against each secret in turn, answering the first secret's
// error when every one refuses it (that is the reason worth logging).
func constructEvent(payload []byte, signature string, secrets []string) (stripe.Event, error) {
	first := errNoSecret
	for i, secret := range secrets {
		event, err := webhook.ConstructEvent(payload, signature, secret)
		if err == nil {
			return event, nil
		}
		if i == 0 {
			first = err
		}
	}
	return stripe.Event{}, first
}

// creatorsNormalize is a code as typed, as the program stores it (creators.NormalizeCode).
func creatorsNormalize(code string) string { return strings.ToUpper(strings.TrimSpace(code)) }

// premiumTrial is a Premium checkout's trial: the first Premium of an account trials
// catalog.PremiumTrialDays while PremiumTrial is on, a creator's code makes it a free month, and a
// comeback (a seat holder returning at their own price) has had its Premium already.
func (s *Service) premiumTrial(ctx context.Context, req premiumCheckoutRequest) (*checkoutTrial, error) {
	spent := req.comebackID != nil
	if !spent {
		var err error
		if spent, err = s.Store.PremiumTrialSpent(ctx, req.userID); err != nil {
			return nil, err
		}
	}
	return s.creatorCodeTrial(ctx, req.creatorCode, spent, s.premiumTrialDays())
}
