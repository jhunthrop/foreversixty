package billingmail

import (
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// The trial emails (design/specs/2026-11-08-premium-trial.md 3.6, rulings 0B and 0C). Every date is
// the instant's UTC day in the house style, as every trial surface on the site says it, and the
// text keeps straight apostrophes: it is plain text.

// trialDateLayout is the house date: "15 Nov 2026".
const trialDateLayout = "2 Jan 2006"

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
	return copygen.Fill(copygen.MailTrialPlanFor, "{plan}", string(plan), "{guild}", guild)
}

// firstCharge is what the first charge buys: "$40 for the next 12 months", or "$5 a month".
func firstCharge(cents int64, interval string) (string, error) {
	if interval == intervalYearly {
		amount, err := FormatPrice(cents, "")
		return copygen.Fill(copygen.MailTrialYearlyCharge, "{amount}", amount), err
	}
	return FormatPrice(cents, interval)
}

// whySent is the closing line and the company, which signs every trial email (never a person).
func whySent(reason string) string {
	return copygen.Fill(copygen.MailTrialWhySent, "{reason}", reason) + "\n" + copygen.MailTrialCompany
}

// byGuild picks the guild form of a text when the email is for a guild's plan.
func byGuild(guild, premium, forGuild string) string {
	if guild == "" {
		return premium
	}
	return forGuild
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
	date, plan := TrialDate(r.EndsOn), string(r.Plan)
	fill := func(text string) string {
		return copygen.Fill(text, "{plan}", plan, "{guild}", r.Guild)
	}
	whose := fill(byGuild(r.Guild, copygen.MailTrialReminderWhose, copygen.MailTrialReminderWhoseGuild))
	text := strings.Join([]string{
		copygen.Fill(copygen.MailTrialReminderLead, "{whose}", whose, "{date}", date, "{charge}", charge),
		copygen.Fill(copygen.MailTrialReminderPlan, "{name}", planName(r.Plan, r.Guild), "{interval}", r.Interval, "{amount}", amount),
		copygen.Fill(copygen.MailTrialReminderCancel, "{url}", r.CancelURL),
		byGuild(r.Guild, copygen.MailTrialReminderKeep, copygen.MailTrialReminderKeepGuild) + " " + copygen.MailTrialStripeAlsoReminds,
		whySent(fill(byGuild(r.Guild, copygen.MailTrialReminderReason, copygen.MailTrialReminderReasonGuild))),
	}, "\n\n")
	subject := copygen.Fill(copygen.MailTrialReminderSubject, "{whose}", whose, "{date}", date)
	return mail.Message{To: r.To, Subject: subject, Text: text}, nil
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
	reason := copygen.Fill(byGuild(r.Guild, copygen.MailTrialReceiptReason, copygen.MailTrialReceiptReasonGuild),
		"{plan}", string(r.Plan), "{guild}", r.Guild)
	text := strings.Join([]string{
		copygen.Fill(copygen.MailTrialReceiptLead, "{price}", price, "{date}", TrialDate(r.ChargedOn), "{url}", r.ReceiptURL),
		copygen.Fill(copygen.MailTrialReceiptNext, "{amount}", amount, "{date}", TrialDate(r.NextOn), "{url}", r.ManageURL),
		whySent(reason),
	}, "\n\n")
	subject := copygen.Fill(copygen.MailTrialReceiptSubject, "{name}", planName(r.Plan, r.Guild), "{price}", price)
	return mail.Message{To: r.To, Subject: subject, Text: text}, nil
}
