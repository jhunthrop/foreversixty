// api/internal/billing/trial_reminders_test.go
package billing

import (
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/events"
)

var trialStart = time.Date(2026, 11, 8, 10, 0, 0, 0, time.UTC)

func days(n float64) time.Duration { return time.Duration(n * 24 * float64(time.Hour)) }

// A 7-day trial is reminded once, on the run 3 days before its first charge: not the run before,
// not again the next day. The email goes to the Stripe customer's address, never the account's.
func TestTrialReminderFiresOnceThreeDaysBefore(t *testing.T) {
	h := newTrialHarness(t)
	uid := h.seedUser(t, "account@example.com")
	h.seedCustomer(t, uid, "stripe-payer@example.com")
	tr := h.seedTrial(t, 0, uid, catalog.LookupKeyPremiumFounderMonthly, trialStart, trialStart.Add(days(7)))

	if sent, published := h.sendAt(t, tr.end.Add(-days(3)-time.Hour)); len(sent) != 0 || published != 0 {
		t.Fatalf("the run before sent %d and published %d", len(sent), published)
	}
	sent, published := h.sendAt(t, tr.end.Add(-days(3)+time.Hour))
	if len(sent) != 1 || published != 1 {
		t.Fatalf("the run 3 days before sent %d and published %d, want 1 and 1", len(sent), published)
	}
	m := sent[0]
	if m.To != "stripe-payer@example.com" || m.Subject != "Your free Premium ends on 15 Nov 2026" {
		t.Fatalf("sent to %q: %q", m.To, m.Subject)
	}
	if !strings.Contains(m.Text, "you'll be charged $5 a month") || !strings.Contains(m.Text, "https://foreversixty.gg/premium/cancel?t=") {
		t.Fatalf("text %q", m.Text)
	}
	e := h.events.got[0]
	if e.Kind != events.BillingTrialEnding || e.Audience.UserIDs[0] != uid || e.Title != "Free Premium ends 15 Nov" ||
		e.Body != "On that day you'll be charged $5 a month. Tap to cancel." || e.Link != "https://foreversixty.gg/account#plan" {
		t.Fatalf("event %+v", e)
	}
	if sent, published := h.sendAt(t, tr.end.Add(-days(2)+time.Hour)); len(sent) != 0 || published != 0 {
		t.Fatalf("the next day sent %d and published %d", len(sent), published)
	}
}

// A trial cancelled before the run gets no reminder.
func TestTrialReminderSkipsACancelledTrial(t *testing.T) {
	h := newTrialHarness(t)
	uid := h.seedUser(t, "cancelled@example.com")
	h.seedCustomer(t, uid, "stripe-cancelled@example.com")
	tr := h.seedTrial(t, 0, uid, catalog.LookupKeyPremiumFounderYearly, trialStart, trialStart.Add(days(7)))
	if _, err := h.pool.Exec(t.Context(), `update entitlements set cancel_at_period_end = true where id = $1`, tr.entitlementID); err != nil {
		t.Fatal(err)
	}
	if sent, published := h.sendAt(t, tr.end.Add(-days(2.5))); len(sent) != 0 || published != 0 {
		t.Fatalf("a cancelled trial sent %d and published %d", len(sent), published)
	}
}

// A trial shorter than 6 days is reminded 2 days ahead.
func TestTrialReminderGivesAShortTrialTwoDays(t *testing.T) {
	h := newTrialHarness(t)
	uid := h.seedUser(t, "short@example.com")
	h.seedCustomer(t, uid, "stripe-short@example.com")
	tr := h.seedTrial(t, 0, uid, catalog.LookupKeyPremiumFounderMonthly, trialStart, trialStart.Add(days(5)))
	if sent, _ := h.sendAt(t, tr.end.Add(-days(2.5))); len(sent) != 0 {
		t.Fatalf("a 5-day trial was reminded 2.5 days ahead")
	}
	if sent, _ := h.sendAt(t, tr.end.Add(-days(1.5))); len(sent) != 1 {
		t.Fatalf("a 5-day trial sent %d inside 2 days, want 1", len(sent))
	}
}

// The guild trial is reminded the same way, naming the guild and its price.
func TestTrialReminderForAGuild(t *testing.T) {
	h := newTrialHarness(t)
	uid := h.seedUser(t, "officer@example.com")
	h.seedCustomer(t, uid, "stripe-officer@example.com")
	g := h.seedGuild(t, "Olympus XXVII", true)
	tr := h.seedTrial(t, g, uid, catalog.LookupKeyGuildFounderYearly, trialStart, trialStart.Add(days(7)))
	sent, published := h.sendAt(t, tr.end.Add(-days(2.5)))
	if len(sent) != 1 || published != 1 {
		t.Fatalf("sent %d, published %d", len(sent), published)
	}
	if sent[0].Subject != "Olympus XXVII's free Guild ends on 15 Nov 2026" || !strings.Contains(sent[0].Text, "for the next 12 months") {
		t.Fatalf("%q\n%s", sent[0].Subject, sent[0].Text)
	}
	if link := h.events.got[0].Link; link != "https://foreversixty.gg/guild/us/hardcore/olympus-xxvii/settings#billing" {
		t.Fatalf("event link %q", link)
	}
}

// With our renewal emails on, a free trial still gets the trial reminder alone, never a renewal.
func TestAFreeTrialIsNeverARenewal(t *testing.T) {
	h := newTrialHarness(t)
	h.svc.RenewalEmails = true
	uid := h.seedUser(t, "yearly-trial@example.com")
	h.seedCustomer(t, uid, "stripe-yearly-trial@example.com")
	tr := h.seedTrial(t, 0, uid, catalog.LookupKeyPremiumFounderYearly, trialStart, trialStart.Add(days(7)))
	sent, _ := h.sendAt(t, tr.end.Add(-days(2.5)))
	if len(sent) != 1 || !strings.HasPrefix(sent[0].Subject, "Your free Premium ends") {
		t.Fatalf("sent %+v", sent)
	}
}

// A billing transfer trials to its first charge too, but it is no free trial: no reminder.
func TestTrialReminderSkipsATransferTrial(t *testing.T) {
	h := newTrialHarness(t)
	uid := h.seedUser(t, "transfer@example.com")
	h.seedCustomer(t, uid, "stripe-transfer@example.com")
	g := h.seedGuild(t, "Handover", true)
	tr := h.seedTrial(t, g, uid, catalog.LookupKeyGuildFounderYearly, trialStart, trialStart.Add(days(7)))
	if _, err := h.pool.Exec(t.Context(), `delete from guild_trials where guild_id = $1`, g); err != nil {
		t.Fatal(err)
	}
	if sent, published := h.sendAt(t, tr.end.Add(-days(2.5))); len(sent) != 0 || published != 0 {
		t.Fatalf("a transfer sent %d and published %d", len(sent), published)
	}
}

// A customer with no email gets the notification alone; the miss is logged with ids, and the
// logs never carry an email address or a cancel token.
func TestTrialReminderWithoutAnEmailAndNoPIIInLogs(t *testing.T) {
	h := newTrialHarness(t)
	quiet := h.seedUser(t, "account-quiet@example.com")
	customer := h.seedCustomer(t, quiet, "")
	h.seedTrial(t, 0, quiet, catalog.LookupKeyPremiumFounderMonthly, trialStart, trialStart.Add(days(7)))
	loud := h.seedUser(t, "account-loud@example.com")
	h.seedCustomer(t, loud, "stripe-loud@example.com")
	h.seedTrial(t, 0, loud, catalog.LookupKeyPremiumFounderMonthly, trialStart, trialStart.Add(days(7)))

	sent, published := h.sendAt(t, trialStart.Add(days(4.5)))
	if len(sent) != 1 || published != 2 {
		t.Fatalf("sent %d (want 1) and published %d (want 2)", len(sent), published)
	}
	logs := h.logs.String()
	if !strings.Contains(logs, "billing notice skipped") || !strings.Contains(logs, customer) {
		t.Fatalf("the miss was not logged with the customer id:\n%s", logs)
	}
	for _, secret := range []string{"@example.com", tokenOf(t, sent[0])} {
		if strings.Contains(logs, secret) {
			t.Fatalf("the logs carry %q:\n%s", secret, logs)
		}
	}
}

// A failed send releases the reminder and drops its link, so the next run tries again with one
// notification in the app.
func TestTrialReminderRetriesAfterAFailedSend(t *testing.T) {
	h := newTrialHarness(t)
	uid := h.seedUser(t, "retry@example.com")
	h.seedCustomer(t, uid, "stripe-retry@example.com")
	tr := h.seedTrial(t, 0, uid, catalog.LookupKeyPremiumFounderMonthly, trialStart, trialStart.Add(days(7)))
	h.mailer.Err = errMailDown
	if _, err := h.svc.SendNotices(t.Context(), tr.end.Add(-days(2.5))); err == nil {
		t.Fatal("a failed send was not an error")
	}
	var tokens int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from billing_cancel_tokens`).Scan(&tokens); err != nil || tokens != 0 {
		t.Fatalf("tokens left %d (%v)", tokens, err)
	}
	h.mailer.Err = nil
	if sent, _ := h.sendAt(t, tr.end.Add(-days(2.4))); len(sent) != 1 {
		t.Fatalf("the retry sent %d", len(sent))
	}
	if len(h.events.got) != 2 || h.events.got[0].DedupeKey != h.events.got[1].DedupeKey {
		t.Fatalf("the retry's notification must share the first's dedupe key: %+v", h.events.got)
	}
}
