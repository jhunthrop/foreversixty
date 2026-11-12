package billingmail

import (
	"encoding/json"
	"flag"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// updateTrialFixture rewrites testdata/trial-emails.json, the contract the site's trial-copy.ts
// is held to (spec 3.6: the site and the email never drift).
var updateTrialFixture = flag.Bool("update", false, "rewrite testdata/trial-emails.json")

const trialFixture = "testdata/trial-emails.json"

var (
	trialEnds = time.Date(2026, time.November, 15, 10, 0, 0, 0, time.UTC)
	trialNext = time.Date(2026, time.December, 15, 10, 0, 0, 0, time.UTC)
)

func trialEmails() map[string]mail.Message {
	return map[string]mail.Message{
		"reminder_premium_yearly": must(TrialReminderEmail(TrialReminder{To: "a@b.c", Plan: PlanPremium, EndsOn: trialEnds,
			AmountCents: 4000, Interval: "yearly", CancelURL: "https://foreversixty.gg/premium/cancel?t=TOKEN"})),
		"reminder_premium_monthly": must(TrialReminderEmail(TrialReminder{To: "a@b.c", Plan: PlanPremium, EndsOn: trialEnds,
			AmountCents: 500, Interval: "monthly", CancelURL: "https://foreversixty.gg/premium/cancel?t=TOKEN"})),
		"reminder_guild_yearly": must(TrialReminderEmail(TrialReminder{To: "a@b.c", Guild: guild, Plan: PlanGuild, EndsOn: trialEnds,
			AmountCents: 12000, Interval: "yearly", CancelURL: "https://foreversixty.gg/premium/cancel?t=TOKEN"})),
		"receipt_premium_monthly": must(TrialReceiptEmail(TrialReceipt{To: "a@b.c", Plan: PlanPremium, ChargedOn: trialEnds,
			NextOn: trialNext, AmountCents: 500, Interval: "monthly", ReceiptURL: "https://invoice.stripe.com/i/x",
			ManageURL: "https://foreversixty.gg/account#plan"})),
		"receipt_guild_yearly": must(TrialReceiptEmail(TrialReceipt{To: "a@b.c", Guild: guild, Plan: PlanGuild, ChargedOn: trialEnds,
			NextOn: trialNext.AddDate(0, 11, 0), AmountCents: 12000, Interval: "yearly", ReceiptURL: "https://invoice.stripe.com/i/x",
			ManageURL: "https://foreversixty.gg/guild/us/forever/olympus-xxvii/settings#billing"})),
	}
}

type fixtureEmail struct {
	Subject string `json:"subject"`
	Text    string `json:"text"`
}

func TestTrialEmailsMatchTheContract(t *testing.T) {
	got := map[string]fixtureEmail{}
	for k, m := range trialEmails() {
		got[k] = fixtureEmail{Subject: m.Subject, Text: m.Text}
	}
	if *updateTrialFixture {
		b, err := json.MarshalIndent(got, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(trialFixture, append(b, '\n'), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	raw, err := os.ReadFile(trialFixture)
	if err != nil {
		t.Fatal(err)
	}
	var want map[string]fixtureEmail
	if err := json.Unmarshal(raw, &want); err != nil {
		t.Fatal(err)
	}
	for k, w := range want {
		if got[k] != w {
			t.Errorf("%s drifted from %s (rerun with -update after a copy ruling):\n got %+v\nwant %+v", k, trialFixture, got[k], w)
		}
	}
	if len(want) != len(got) {
		t.Errorf("fixture has %d emails, the builders %d", len(want), len(got))
	}
}

func TestTrialEmailsSaySpecWords(t *testing.T) {
	e := trialEmails()
	cases := []struct{ key, want string }{
		{"reminder_premium_yearly", "Your free Premium ends on 15 Nov 2026. On that day you'll be charged $40 for the next 12 months. Cancel before 15 Nov 2026 (UTC) and you pay nothing."},
		{"reminder_premium_yearly", "Plan: Premium, yearly: $40"},
		{"reminder_premium_yearly", "You may also get a reminder from Stripe, our payment processor."},
		{"reminder_premium_yearly", "You're getting this because you started a free Premium trial on Forever Sixty.\nCommish LLC"},
		{"reminder_premium_monthly", "you'll be charged $5 a month."},
		{"reminder_guild_yearly", "Olympus XXVII's free Guild ends on 15 Nov 2026."},
		{"receipt_premium_monthly", "Your free trial ended and we charged $5 a month on 15 Nov 2026. Your receipt: https://invoice.stripe.com/i/x"},
		{"receipt_premium_monthly", "The next charge is $5 on 15 Dec 2026."},
	}
	for _, c := range cases {
		if !strings.Contains(e[c.key].Text, c.want) {
			t.Errorf("%s lacks %q:\n%s", c.key, c.want, e[c.key].Text)
		}
	}
	if s := e["reminder_premium_yearly"].Subject; s != "Your free Premium ends on 15 Nov 2026" {
		t.Errorf("reminder subject %q", s)
	}
	if s := e["receipt_premium_monthly"].Subject; s != "Receipt: Premium, $5 a month" {
		t.Errorf("receipt subject %q", s)
	}
	for k, m := range e {
		if strings.ContainsAny(m.Text+m.Subject, "\u2019\u2018\u2014") {
			t.Errorf("%s has a curly apostrophe or an em dash", k)
		}
	}
}
