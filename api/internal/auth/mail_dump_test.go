//go:build payloaddump

package auth

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billingmail"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// TestDumpEmails writes every email the server sends, the HTML part and its plain text, to
// $REVIEW_PAYLOADS_DIR/mail for the review captures (go test -tags payloaddump). The subscribe
// confirmation is dumped by its own package.
func TestDumpEmails(t *testing.T) {
	dir := os.Getenv("REVIEW_PAYLOADS_DIR")
	if dir == "" {
		t.Skip("REVIEW_PAYLOADS_DIR is not set")
	}
	sent := time.Date(2026, time.November, 10, 18, 0, 0, 0, time.UTC)
	renews := time.Date(2026, time.December, 12, 20, 0, 0, 0, time.UTC)
	site := "https://foreversixty.gg"
	must := func(m mail.Message, err error) mail.Message {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
		return m
	}
	emails := map[string]mail.Message{
		"renewal-premium": must(billingmail.RenewalEmail(billingmail.Renewal{To: "kraggor@example.com",
			Plan: billingmail.PlanPremium, RenewsOn: renews, AmountCents: 5000, Interval: "yearly", Founder: true,
			ManageURL: site + "/account#plan"}, sent)),
		"renewal-guild": must(billingmail.RenewalEmail(billingmail.Renewal{To: "officer@example.com",
			Plan: billingmail.PlanGuild, Guild: "Olympus XXVII", RenewsOn: renews, AmountCents: 15000, Interval: "yearly",
			ManageURL: site + "/guild/us/normal/olympus-xxvii/settings#billing"}, sent)),
		"size": must(billingmail.SizeEmail(billingmail.Size{To: "officer@example.com", Guild: "Olympus XXVII",
			Members: 31, MemberLimit: 25, RenewsOn: renews, RaidGuildCents: 25000, Interval: "yearly",
			SettingsURL: site + "/guild/us/normal/olympus-xxvii/settings#billing"}, sent)),
		"seat-offered": must(billingmail.SeatOfferedEmail(billingmail.SeatOffered{To: "zulmara@example.com",
			Guild: "Olympus XXVII", From: "Kraggor", OwnEndsOn: renews, AccountURL: site + "/account#plan"}, sent)),
		"seat-given": must(billingmail.SeatGivenEmail(billingmail.SeatGiven{To: "zulmara@example.com",
			Guild: "Olympus XXVII", From: "Kraggor", AccountURL: site + "/account#plan"}, sent)),
		"seat-removed": must(billingmail.SeatRemovedEmail(billingmail.SeatRemoved{To: "zulmara@example.com",
			Guild: "Olympus XXVII", EndedOn: sent, AmountCents: 600, Interval: "monthly",
			ComebackURL: site + "/checkout?plan=premium&interval=monthly&comeback=12"}, sent)),
		"sign-in": must(signInEmail("kraggor@example.com",
			"https://api.foreversixty.gg/v1/auth/email/callback?token=0f3c9a1e7b2d4c6a8e0f1b3d5c7a9e1f", sent)),
	}
	writeDumpedEmails(t, filepath.Join(dir, "mail"), emails)
}

// writeDumpedEmails writes each message as <name>.html and <name>.txt, with its subject on the
// text's first line.
func writeDumpedEmails(t *testing.T, dir string, emails map[string]mail.Message) {
	t.Helper()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	for name, m := range emails {
		if err := os.WriteFile(filepath.Join(dir, name+".html"), []byte(m.HTML), 0o644); err != nil {
			t.Fatal(err)
		}
		text := "Subject: " + m.Subject + "\n\n" + m.Text
		if err := os.WriteFile(filepath.Join(dir, name+".txt"), []byte(text), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}
