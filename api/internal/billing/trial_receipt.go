// api/internal/billing/trial_receipt.go
//
// The receipt after a free trial's first charge (spec 3.6, 0C): invoice.paid with an amount on a
// subscription that started as a free trial sends it once, to the Stripe customer's email, with
// Stripe's hosted invoice and the next charge's date. Stripe's own successful-payment receipt
// stays off.
package billing

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/billingmail"
)

// billingReasonCycle is the invoice reason of a subscription's charge at its period's end, which
// a trial's first charge is.
const billingReasonCycle = "subscription_cycle"

// receiptTarget is a free trial's subscription as its receipt reads it.
type receiptTarget struct {
	DueTrial
	NextOn time.Time
}

// receiptTargetQuery reads the entitlement of a subscription recorded as a free trial; $1 is the
// subscription id.
const receiptTargetQuery = `
	select e.id, e.billing_user_id, e.plan, coalesce(e.tier, ''), coalesce(e.price_lookup_key, ''),
	       coalesce(e.unit_amount_cents, 0), e.trial_end, e.current_period_end,
	       coalesce(g.region, ''), coalesce(g.ruleset, ''), coalesce(g.name, '')
	from entitlements e
	left join guilds g on g.id = e.guild_id
	where e.source = 'stripe' and e.stripe_subscription_id = $1 and e.trial_end is not null
	  and e.billing_user_id is not null and e.current_period_end is not null
	  and (exists (select 1 from premium_trials pt where pt.stripe_subscription_id = $1)
	       or exists (select 1 from guild_trials gt where gt.stripe_subscription_id = $1))`

// TrialReceiptTarget is the free trial a subscription started as; ok false for any other.
func (s *Store) TrialReceiptTarget(ctx context.Context, subscriptionID string) (receiptTarget, bool, error) {
	t := receiptTarget{DueTrial: DueTrial{SubscriptionID: subscriptionID}}
	err := s.Pool.QueryRow(ctx, receiptTargetQuery, subscriptionID).Scan(&t.EntitlementID, &t.BillingUserID, &t.Plan,
		&t.Tier, &t.LookupKey, &t.UnitAmountCents, &t.TrialEnd, &t.NextOn, &t.Region, &t.Ruleset, &t.Guild)
	if errors.Is(err, pgx.ErrNoRows) {
		return receiptTarget{}, false, nil
	}
	if err != nil {
		return receiptTarget{}, false, fmt.Errorf("billing: trial receipt %s: %w", subscriptionID, err)
	}
	return t, true, nil
}

// invoiceInt reads an integer field of a webhook's invoice object.
func invoiceInt(invoice map[string]any, key string) int64 {
	if v, ok := invoice[key].(float64); ok {
		return int64(v)
	}
	return 0
}

// invoiceString reads a string field of a webhook's invoice object.
func invoiceString(invoice map[string]any, key string) string {
	v, _ := invoice[key].(string)
	return v
}

// sendTrialReceipt answers a paid invoice: the receipt when it is a free trial's first charge.
// The entitlement was written from the fresh subscription just before, so its period end is the
// next charge's date. A failed send releases the claim and fails the event, so Stripe redelivers.
func (s *Service) sendTrialReceipt(ctx context.Context, event stripe.Event, subscriptionID string) error {
	invoice := event.Data.Object
	if s.Mailer == nil {
		return nil
	}
	if invoiceInt(invoice, "amount_paid") <= 0 || invoiceString(invoice, "billing_reason") != billingReasonCycle {
		return nil
	}
	t, ok, err := s.Store.TrialReceiptTarget(ctx, subscriptionID)
	if err != nil || !ok {
		return err
	}
	claimed, err := s.Store.ClaimNotice(ctx, NoticeTrialReceipt, t.EntitlementID, t.TrialEnd)
	if err != nil || !claimed {
		return err
	}
	err = s.mailTrialReceipt(ctx, t, time.Unix(event.Created, 0), invoiceString(invoice, "hosted_invoice_url"))
	if err != nil {
		if relErr := s.Store.ReleaseNotice(ctx, NoticeTrialReceipt, t.EntitlementID, t.TrialEnd); relErr != nil {
			err = errors.Join(err, relErr)
		}
		return fmt.Errorf("billing: trial receipt for entitlement %d: %w", t.EntitlementID, err)
	}
	return nil
}

// mailTrialReceipt builds and sends the receipt; nothing goes when the customer has no email.
func (s *Service) mailTrialReceipt(ctx context.Context, t receiptTarget, chargedOn time.Time, receiptURL string) error {
	facts := catalog.FactsOf(t.LookupKey, t.UnitAmountCents)
	if facts.AmountCents == nil || facts.Interval == nil {
		return fmt.Errorf("lookup key %q is not in the catalog", t.LookupKey)
	}
	to, err := s.payerEmail(ctx, NoticeTrialReceipt, t.EntitlementID, t.BillingUserID)
	if err != nil || to == "" {
		return err
	}
	manage := s.siteURL("/account#plan")
	if t.Guild != "" {
		manage = s.guildSettingsURL(t.Region, t.Ruleset, t.Guild)
	}
	msg, err := billingmail.TrialReceiptEmail(billingmail.TrialReceipt{
		To: to, Guild: t.Guild, Plan: trialPlan(t.Plan, t.Tier), ChargedOn: chargedOn, NextOn: t.NextOn,
		AmountCents: *facts.AmountCents, Interval: *facts.Interval, ReceiptURL: receiptURL, ManageURL: manage,
	})
	if err != nil {
		return err
	}
	return s.Mailer.Send(ctx, msg)
}
