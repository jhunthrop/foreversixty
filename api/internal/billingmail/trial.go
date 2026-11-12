package billingmail

import (
	"fmt"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// The trial emails (design/specs/2026-11-08-premium-trial.md 3.6, rulings 0B and 0C). Every date is
// the instant's UTC day in the house style, as every trial surface on the site says it, and the
// text keeps straight apostrophes: it is plain text.

const (
	trialDateLayout = "2 Jan 2006"
	// companyName signs every trial email; never a person's name.
	companyName = "Commish LLC"
	// stripeAlsoReminds is said once, so our reminder never contradicts Stripe's own.
	stripeAlsoReminds = "You may also get a reminder from Stripe, our payment processor."
)

// TrialDate is a trial's date as the site and the emails say it: "15 Nov 2026" (UTC).
func TrialDate(t time.Time) string { return t.UTC().Format(trialDateLayout) }

// TrialReminder is the email a few days before a free trial's first charge. Guild is empty for
// Premium.
type TrialReminder struct {
	To, Guild   string
	Plan        Plan
	EndsOn      time.Time
	AmountCents int64
	Interval    string
	CancelURL   string
}

// TrialReceipt is the email after a trial's first charge.
type TrialReceipt struct {
	To, Guild   string
	Plan        Plan
	ChargedOn   time.Time
	NextOn      time.Time
	AmountCents int64
	Interval    string
	ReceiptURL  string
	ManageURL   string
}

// planName is "Premium", or "Guild for Olympus XXVII".
func planName(plan Plan, guild string) string {
	if guild == "" {
		return string(plan)
	}
	return fmt.Sprintf("%s for %s", plan, guild)
}

// firstCharge is what the first charge buys: "$40 for the next 12 months", or "$5 a month".
func firstCharge(cents int64, interval string) (string, error) {
	if interval == intervalYearly {
		amount, err := FormatPrice(cents, "")
		return amount + " for the next 12 months", err
	}
	return FormatPrice(cents, interval)
}

// whySent is the closing line and the company.
func whySent(reason string) string {
	return fmt.Sprintf("You're getting this because %s on Forever Sixty.\n%s", reason, companyName)
}

// TrialReminderEmail builds the reminder.
func TrialReminderEmail(r TrialReminder) (mail.Message, error) {
	charge, err := firstCharge(r.AmountCents, r.Interval)
	if err != nil {
		return mail.Message{}, err
	}
	amount, err := FormatPrice(r.AmountCents, "")
	if err != nil {
		return mail.Message{}, err
	}
	date, name := TrialDate(r.EndsOn), planName(r.Plan, r.Guild)
	whose, keep, reason := "Your free "+string(r.Plan), "To keep Premium, do nothing.",
		"you started a free Premium trial"
	if r.Guild != "" {
		whose = fmt.Sprintf("%s's free %s", r.Guild, r.Plan)
		keep = "To keep the plan, do nothing."
		reason = fmt.Sprintf("you started a free %s trial for %s", r.Plan, r.Guild)
	}
	text := strings.Join([]string{
		fmt.Sprintf("%s ends on %s. On that day you'll be charged %s. Cancel before %s (UTC) and you pay nothing.",
			whose, date, charge, date),
		fmt.Sprintf("Plan: %s, %s: %s", name, r.Interval, amount),
		"Cancel in one click: " + r.CancelURL,
		keep + " " + stripeAlsoReminds,
		whySent(reason),
	}, "\n\n")
	return mail.Message{To: r.To, Subject: fmt.Sprintf("%s ends on %s", whose, date), Text: text}, nil
}

// TrialReceiptEmail builds the receipt.
func TrialReceiptEmail(r TrialReceipt) (mail.Message, error) {
	price, err := FormatPrice(r.AmountCents, r.Interval)
	if err != nil {
		return mail.Message{}, err
	}
	amount, err := FormatPrice(r.AmountCents, "")
	if err != nil {
		return mail.Message{}, err
	}
	name := planName(r.Plan, r.Guild)
	reason := fmt.Sprintf("your free %s trial became a paid plan", r.Plan)
	if r.Guild != "" {
		reason = fmt.Sprintf("the free %s trial for %s became a paid plan", r.Plan, r.Guild)
	}
	text := strings.Join([]string{
		fmt.Sprintf("Your free trial ended and we charged %s on %s. Your receipt: %s", price, TrialDate(r.ChargedOn), r.ReceiptURL),
		fmt.Sprintf("The next charge is %s on %s. Change or cancel any time: %s", amount, TrialDate(r.NextOn), r.ManageURL),
		whySent(reason),
	}, "\n\n")
	return mail.Message{To: r.To, Subject: fmt.Sprintf("Receipt: %s, %s", name, price), Text: text}, nil
}
