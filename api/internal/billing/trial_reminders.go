// api/internal/billing/trial_reminders.go
//
// The trial reminder (spec 4.1, 4.2): billing-daily tells the payer of every free trial, Premium
// or guild, a few days before its first charge, once per trial, by email to the Stripe customer's
// address (read at send time, never stored or logged) and in the app. A cancelled trial gets none.
package billing

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/billingmail"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// Reminder leads (spec 4.2): 3 days before the charge for a trial of 6 days or more, 2 for a
// shorter one. With a daily run the reminder lands 3 to 4 days ahead, so it names the date.
const (
	trialReminderLead      = 3 * 24 * time.Hour
	shortTrialReminderLead = 2 * 24 * time.Hour
	shortTrialBelow        = 6 * 24 * time.Hour
)

// trialEventDayLayout is the notification title's date: "15 Nov".
const trialEventDayLayout = "2 Jan"

// DueTrial is one live free trial inside the longest reminder lead. Guild is empty for Premium.
type DueTrial struct {
	EntitlementID   int64
	BillingUserID   int64
	Plan            string
	Tier            string
	LookupKey       string
	UnitAmountCents int64
	TrialEnd        time.Time
	StartedAt       time.Time
	SubscriptionID  string
	Region          string
	Ruleset         string
	Guild           string
}

// reminderLead is how far ahead of its end a trial is reminded.
func reminderLead(length time.Duration) time.Duration {
	if length < shortTrialBelow {
		return shortTrialReminderLead
	}
	return trialReminderLead
}

// dueTrialQuery selects the free trials (recorded in premium_trials or guild_trials, so a billing
// transfer's trial is never one) still trialing, not set to cancel, ending in (now, now+lead].
const dueTrialQuery = `
	select e.id, e.billing_user_id, e.plan, coalesce(e.tier, ''), coalesce(e.price_lookup_key, ''),
	       coalesce(e.unit_amount_cents, 0), e.trial_end, coalesce(pt.created_at, gt.started_at),
	       e.stripe_subscription_id, coalesce(g.region, ''), coalesce(g.ruleset, ''), coalesce(g.name, '')
	from entitlements e
	left join guilds g on g.id = e.guild_id
	left join premium_trials pt on e.plan = 'premium' and pt.user_id = e.user_id
	     and pt.stripe_subscription_id = e.stripe_subscription_id
	left join guild_trials gt on e.plan = 'guild' and gt.guild_id = e.guild_id
	     and gt.stripe_subscription_id = e.stripe_subscription_id
	where e.source = 'stripe' and e.status = 'trialing' and not e.cancel_at_period_end
	  and e.billing_user_id is not null and e.trial_end is not null
	  and e.trial_end > $1 and e.trial_end <= $1 + $2::interval
	  and (pt.user_id is not null or gt.guild_id is not null)
	order by e.id`

// DueTrialReminders lists the trials whose reminder is due at now.
func (s *Store) DueTrialReminders(ctx context.Context, now time.Time) ([]DueTrial, error) {
	rows, err := s.Pool.Query(ctx, dueTrialQuery, now, fmt.Sprintf("%d seconds", int64(trialReminderLead.Seconds())))
	if err != nil {
		return nil, fmt.Errorf("billing: due trials: %w", err)
	}
	all, err := pgx.CollectRows(rows, func(r pgx.CollectableRow) (DueTrial, error) {
		var d DueTrial
		err := r.Scan(&d.EntitlementID, &d.BillingUserID, &d.Plan, &d.Tier, &d.LookupKey, &d.UnitAmountCents,
			&d.TrialEnd, &d.StartedAt, &d.SubscriptionID, &d.Region, &d.Ruleset, &d.Guild)
		return d, err
	})
	if err != nil {
		return nil, fmt.Errorf("billing: due trials: %w", err)
	}
	due := all[:0]
	for _, d := range all {
		if d.TrialEnd.Sub(now) <= reminderLead(d.TrialEnd.Sub(d.StartedAt)) {
			due = append(due, d)
		}
	}
	return due, nil
}

// sendTrialReminders sends every due reminder; it returns how many trials were told and the
// failures joined.
func (s *Service) sendTrialReminders(ctx context.Context, now time.Time) (int, error) {
	dues, err := s.Store.DueTrialReminders(ctx, now)
	if err != nil {
		return 0, err
	}
	sent := 0
	var errs []error
	for _, due := range dues {
		ok, err := s.sendTrialReminder(ctx, due)
		if err != nil {
			s.logger().Error("billing notice", "kind", NoticeTrialEnding, "entitlement", due.EntitlementID, "err", err)
			errs = append(errs, fmt.Errorf("trial reminder for entitlement %d: %w", due.EntitlementID, err))
		} else if ok {
			sent++
		}
	}
	return sent, errors.Join(errs...)
}

// sendTrialReminder claims the reminder, tells the app, then emails the payer with a fresh cancel
// link. A failed email releases the claim and drops its token, so the next run tries again; the
// notification's dedupe key keeps that from repeating in the app.
func (s *Service) sendTrialReminder(ctx context.Context, due DueTrial) (bool, error) {
	claimed, err := s.Store.ClaimNotice(ctx, NoticeTrialEnding, due.EntitlementID, due.TrialEnd)
	if err != nil || !claimed {
		return false, err
	}
	facts := catalog.FactsOf(due.LookupKey, due.UnitAmountCents)
	if facts.AmountCents == nil || facts.Interval == nil {
		return false, s.releaseTrialNotice(ctx, due, fmt.Errorf("lookup key %q is not in the catalog", due.LookupKey))
	}
	s.publishTrialEnding(ctx, due, *facts.AmountCents, *facts.Interval)
	to, err := s.payerEmail(ctx, NoticeTrialEnding, due.EntitlementID, due.BillingUserID)
	if err != nil || to == "" {
		if err != nil {
			return false, s.releaseTrialNotice(ctx, due, err)
		}
		return true, nil
	}
	token, err := s.Store.CreateCancelToken(ctx, due.EntitlementID, due.SubscriptionID, due.TrialEnd)
	if err != nil {
		return false, s.releaseTrialNotice(ctx, due, err)
	}
	msg, err := billingmail.TrialReminderEmail(billingmail.TrialReminder{
		To: to, Guild: due.Guild, Plan: trialPlan(due.Plan, due.Tier), EndsOn: due.TrialEnd,
		AmountCents: *facts.AmountCents, Interval: *facts.Interval, CancelURL: s.cancelLinkURL(token),
	})
	if err == nil {
		err = s.Mailer.Send(ctx, msg)
	}
	if err != nil {
		if delErr := s.Store.DeleteCancelToken(ctx, token); delErr != nil {
			err = errors.Join(err, delErr)
		}
		return false, s.releaseTrialNotice(ctx, due, err)
	}
	return true, nil
}

// releaseTrialNotice undoes the claim after err, answering err (joined with a failed release).
func (s *Service) releaseTrialNotice(ctx context.Context, due DueTrial, err error) error {
	if relErr := s.Store.ReleaseNotice(ctx, NoticeTrialEnding, due.EntitlementID, due.TrialEnd); relErr != nil {
		return errors.Join(err, relErr)
	}
	return err
}

// payerEmail is the Stripe customer's email for one send, "" when there is none: no Stripe on this
// deployment, no customer, or a customer without an email. A miss is logged with ids only.
func (s *Service) payerEmail(ctx context.Context, kind string, entitlementID, userID int64) (string, error) {
	customerID, ok, err := s.Store.CustomerID(ctx, userID)
	if err != nil {
		return "", err
	}
	if !ok || s.Gateway == nil {
		s.logger().Info("billing notice skipped", "kind", kind, "entitlement", entitlementID, "reason", "no stripe customer")
		return "", nil
	}
	email, err := s.Gateway.CustomerEmail(ctx, customerID)
	if err != nil {
		return "", err
	}
	if email == "" {
		s.logger().Info("billing notice skipped", "kind", kind, "entitlement", entitlementID, "customer", customerID, "reason", "no email")
	}
	return email, nil
}

// trialPlan is the plan as an email names it.
func trialPlan(plan, tier string) billingmail.Plan {
	if plan == entitlements.PlanGuild {
		return noticePlan(tier)
	}
	return billingmail.PlanPremium
}

// trialEndingCopy is billing.trial_ending's words (spec 3.6).
func trialEndingCopy(due DueTrial, price string) (title, body string) {
	day := due.TrialEnd.UTC().Format(trialEventDayLayout)
	title = "Free Premium ends " + day
	if due.Guild != "" {
		title = fmt.Sprintf("%s's free %s ends %s", due.Guild, trialPlan(due.Plan, due.Tier), day)
	}
	return title, fmt.Sprintf("On that day you'll be charged %s. Tap to cancel.", price)
}

// publishTrialEnding tells the payer in the app, once per trial (its dedupe key).
func (s *Service) publishTrialEnding(ctx context.Context, due DueTrial, cents int64, interval string) {
	if s.Events == nil {
		return
	}
	price, err := billingmail.FormatPrice(cents, interval)
	if err != nil {
		s.logger().Error("billing notice", "kind", NoticeTrialEnding, "entitlement", due.EntitlementID, "err", err)
		return
	}
	title, body := trialEndingCopy(due, price)
	link := s.siteURL("/account#plan")
	if due.Guild != "" {
		link = s.guildSettingsURL(due.Region, due.Ruleset, due.Guild)
	}
	events.Quietly(ctx, s.Events, s.logger(), events.Event{
		Kind: events.BillingTrialEnding, Audience: events.Users(due.BillingUserID),
		DedupeKey: "billing.trial_ending:" + strconv.FormatInt(due.EntitlementID, 10) + ":" + strconv.FormatInt(due.TrialEnd.Unix(), 10),
		Title:     title, Body: body, Link: link,
		Data: map[string]string{"user_id": strconv.FormatInt(due.BillingUserID, 10)},
	})
}

// guildSettingsURL is the web's guildSettingsHref plus the #billing anchor.
func (s *Service) guildSettingsURL(region, ruleset, guild string) string {
	return s.siteURL(fmt.Sprintf("/guild/%s/%s/%s/settings#billing", region, ruleset, character.Slug(guild)))
}
