package billingmail

import (
	"errors"
	"html"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

var (
	fixtureDate = time.Date(2027, time.November, 18, 20, 0, 0, 0, time.UTC)
	sentAt      = time.Date(2027, time.November, 11, 16, 0, 0, 0, time.UTC)
)

const (
	guild = "Olympus XXVII"
	name  = "Kraggor"
)

func must(m mail.Message, err error) mail.Message {
	if err != nil {
		panic(err)
	}
	return m
}

func allEmails() map[string]mail.Message {
	return map[string]mail.Message{
		"renewal": must(RenewalEmail(Renewal{To: "a@b.c", Plan: PlanGuild, Guild: guild, RenewsOn: fixtureDate,
			AmountCents: 12000, Interval: "yearly", Founder: true, ManageURL: "https://x/m"}, sentAt)),
		"size": must(SizeEmail(Size{To: "a@b.c", Guild: guild, Members: 31, MemberLimit: 25, RenewsOn: fixtureDate,
			RaidGuildCents: 25000, Interval: "yearly", SettingsURL: "https://x/s"}, sentAt)),
		"offered": must(SeatOfferedEmail(SeatOffered{To: "a@b.c", Guild: guild, From: name, OwnEndsOn: fixtureDate,
			AccountURL: "https://x/a"}, sentAt)),
		"given": must(SeatGivenEmail(SeatGiven{To: "a@b.c", Guild: guild, From: name, AccountURL: "https://x/a"}, sentAt)),
		"removed": must(SeatRemovedEmail(SeatRemoved{To: "a@b.c", Guild: guild, EndedOn: fixtureDate,
			AmountCents: 500, Interval: "monthly", ComebackURL: "https://x/c"}, sentAt)),
	}
}

func TestEmailSubjectsAndLinks(t *testing.T) {
	want := map[string]struct{ subject, action string }{
		"renewal": {"Your Guild plan renews on 18 Nov 2027", "Manage or cancel: https://x/m"},
		"size":    {"Olympus XXVII has 31 members; Guild covers 25", "See Raid guild: https://x/s"},
		"offered": {"Olympus XXVII offers you a Premium seat", "Accept or decline: https://x/a"},
		"given":   {"You have Premium from Olympus XXVII", "See your plan: https://x/a"},
		"removed": {"Your Premium seat from Olympus XXVII has ended", "Come back at $5 a month: https://x/c"},
	}
	for key, m := range allEmails() {
		if m.To != "a@b.c" || m.Subject != want[key].subject {
			t.Errorf("%s: to %q, subject %q", key, m.To, m.Subject)
		}
		if !strings.Contains(m.Text, "\n\n"+want[key].action+"\n\n") {
			t.Errorf("%s: text has no %q:\n%s", key, want[key].action, m.Text)
		}
		if !strings.Contains(m.HTML, `href="`+strings.SplitN(want[key].action, ": ", 2)[1]+`"`) {
			t.Errorf("%s: html has no button link", key)
		}
	}
}

// The renewal says each fact once: the plan and date in the heading, the guild and price in the
// figure block, then the founder note.
func TestRenewalText(t *testing.T) {
	m := allEmails()["renewal"]
	want := "Your Guild plan renews on 18 Nov 2027\n\n" +
		"To stop it, cancel before then and nothing more is charged.\n\n" +
		"Guild: Olympus XXVII\nPrice: $120 a year\n\n" +
		"Founder price, kept while your plan stays active.\n\n" +
		"Manage or cancel: https://x/m\n\n" +
		"--\nA fan-made player tool. Not affiliated with Blizzard.\n" +
		"© 2027 Commish LLC d/b/a Forever Sixty\n" +
		"Reply to this email or write to help@foreversixty.gg.\n"
	if m.Text != want {
		t.Errorf("text =\n%s\nwant\n%s", m.Text, want)
	}
}

func TestPremiumRenewalAtListPrice(t *testing.T) {
	m := must(RenewalEmail(Renewal{To: "a@b.c", Plan: PlanPremium, RenewsOn: fixtureDate,
		AmountCents: 4000, Interval: "yearly", ManageURL: "https://x/m"}, sentAt))
	if strings.Contains(m.Text, "Guild:") || !strings.Contains(m.Text, "Price: $40 a year\n") {
		t.Errorf("text = %q", m.Text)
	}
	if strings.Contains(m.Text, "Founder") {
		t.Errorf("a list price carries the founder note: %q", m.Text)
	}
}

func TestNoExclamationMarks(t *testing.T) {
	for key, m := range allEmails() {
		if strings.Contains(m.Subject+m.Text, "!") {
			t.Errorf("%s contains an exclamation mark", key)
		}
	}
}

func TestMonthlySizeEmail(t *testing.T) {
	m := must(SizeEmail(Size{To: "a@b.c", Guild: guild, Members: 40, MemberLimit: 25, RenewsOn: fixtureDate,
		RaidGuildCents: 2500, Interval: "monthly", SettingsURL: "https://x/s"}, sentAt))
	if !strings.Contains(m.Text, "Raid guild: $25 a month\n") {
		t.Errorf("text = %q", m.Text)
	}
}

// The member limit is the caller's (the catalog's), never typed here.
func TestSizeEmailNamesTheGivenLimit(t *testing.T) {
	m := must(SizeEmail(Size{To: "a@b.c", Guild: guild, Members: 40, MemberLimit: 30, RenewsOn: fixtureDate,
		RaidGuildCents: 2500, Interval: "monthly", SettingsURL: "https://x/s"}, sentAt))
	if !strings.Contains(m.Subject, "Guild covers 30") || !strings.Contains(m.Text, "covers up to 30.") {
		t.Errorf("subject = %q, text = %q", m.Subject, m.Text)
	}
	if _, err := SizeEmail(Size{RaidGuildCents: 2500, Interval: "monthly"}, sentAt); !errors.Is(err, ErrNoMemberLimit) {
		t.Errorf("no limit err = %v", err)
	}
}

// A guild name reaches the subject and the text as written, and the HTML escaped.
func TestGuildNamePassesThroughVerbatim(t *testing.T) {
	odd := "A <b> & C"
	m := must(SeatGivenEmail(SeatGiven{To: "a@b.c", Guild: odd, From: name, AccountURL: "https://x/a"}, sentAt))
	if !strings.Contains(m.Subject, odd) || !strings.Contains(m.Text, odd) {
		t.Errorf("name altered: %q / %q", m.Subject, m.Text)
	}
	if strings.Contains(m.HTML, odd) || !strings.Contains(m.HTML, html.EscapeString(odd)) {
		t.Errorf("html does not escape the name")
	}
}

func TestNegativeAmountsAreRejected(t *testing.T) {
	_, err := RenewalEmail(Renewal{Plan: PlanGuild, AmountCents: -1, Interval: "yearly"}, sentAt)
	if !errors.Is(err, ErrNegativeAmount) {
		t.Errorf("renewal err = %v", err)
	}
	_, err = SizeEmail(Size{RaidGuildCents: -1}, sentAt)
	if !errors.Is(err, ErrNegativeAmount) {
		t.Errorf("size err = %v", err)
	}
	_, err = SeatRemovedEmail(SeatRemoved{AmountCents: -1}, sentAt)
	if !errors.Is(err, ErrNegativeAmount) {
		t.Errorf("removed err = %v", err)
	}
}
