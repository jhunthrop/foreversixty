package billingmail

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

var fixtureDate = time.Date(2027, time.November, 18, 20, 0, 0, 0, time.UTC)

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
			AmountCents: 12000, Interval: "yearly", Founder: true, ManageURL: "https://x/m"})),
		"size": must(SizeEmail(Size{To: "a@b.c", Guild: guild, Members: 31, RenewsOn: fixtureDate,
			RaidGuildCents: 25000, Interval: "yearly", SettingsURL: "https://x/s"})),
		"offered": SeatOfferedEmail(SeatOffered{To: "a@b.c", Guild: guild, From: name, OwnEndsOn: fixtureDate, AccountURL: "https://x/a"}),
		"given":   SeatGivenEmail(SeatGiven{To: "a@b.c", Guild: guild, From: name, AccountURL: "https://x/a"}),
		"removed": must(SeatRemovedEmail(SeatRemoved{To: "a@b.c", Guild: guild, EndedOn: fixtureDate,
			AmountCents: 500, Interval: "monthly", ComebackURL: "https://x/c"})),
	}
}

func TestEmailsMatchSpec(t *testing.T) {
	want := map[string]struct{ subject, text string }{
		"renewal": {"Your Guild plan renews on 18 November",
			"Your Guild plan for Olympus XXVII renews on 18 November 2027 for $120 a year, your founder price; to stop it, cancel before then and nothing more is charged.\n\nManage or cancel: https://x/m"},
		"size": {"Olympus XXVII has 31 members; Guild covers 25",
			"Olympus XXVII has 31 verified members and its Guild plan covers up to 25, so when it renews on 18 November 2027 you may want Raid guild at $250 a year; nothing switches off either way.\n\nLook at Raid guild: https://x/s"},
		"offered": {"Olympus XXVII offers you a Premium seat",
			"Kraggor offered you a Premium seat in Olympus XXVII; if you accept, your own Premium stops renewing and ends on 18 November 2027, the seat starts then, and you are not charged twice.\n\nAccept or decline: https://x/a"},
		"given": {"You have Premium from Olympus XXVII",
			"Kraggor gave you a Premium seat in Olympus XXVII, so you have Premium for as long as the guild's plan lasts, with nothing billed to you.\n\nSee your plan: https://x/a"},
		"removed": {"Your Premium seat from Olympus XXVII has ended",
			"Your Premium seat in Olympus XXVII ended on 18 November 2027, and because you gave up your own Premium for it, this link brings it back at $5 a month, the price you had.\n\nCome back at $5 a month: https://x/c"},
	}
	for key, m := range allEmails() {
		if m.To != "a@b.c" {
			t.Errorf("%s: To = %q", key, m.To)
		}
		if m.Subject != want[key].subject {
			t.Errorf("%s: subject = %q", key, m.Subject)
		}
		if m.Text != want[key].text {
			t.Errorf("%s: text = %q\nwant %q", key, m.Text, want[key].text)
		}
	}
}

func TestPremiumRenewalAtListPrice(t *testing.T) {
	m := must(RenewalEmail(Renewal{To: "a@b.c", Plan: PlanPremium, RenewsOn: fixtureDate,
		AmountCents: 4000, Interval: "yearly", ManageURL: "https://x/m"}))
	want := "Your Premium plan renews on 18 November 2027 for $40 a year; to stop it, cancel before then and nothing more is charged."
	if first := strings.SplitN(m.Text, "\n", 2)[0]; first != want {
		t.Errorf("first line = %q", first)
	}
	if m.Subject != "Your Premium plan renews on 18 November" {
		t.Errorf("subject = %q", m.Subject)
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
	m := must(SizeEmail(Size{To: "a@b.c", Guild: guild, Members: 40, RenewsOn: fixtureDate,
		RaidGuildCents: 2500, Interval: "monthly", SettingsURL: "https://x/s"}))
	if !strings.Contains(m.Text, "Raid guild at $25 a month;") {
		t.Errorf("text = %q", m.Text)
	}
}

func TestGuildNamePassesThroughVerbatim(t *testing.T) {
	odd := "A <b> & C\nD"
	m := SeatGivenEmail(SeatGiven{To: "a@b.c", Guild: odd, From: name, AccountURL: "https://x/a"})
	if !strings.Contains(m.Subject, odd) || !strings.Contains(m.Text, odd) {
		t.Errorf("name altered: %q / %q", m.Subject, m.Text)
	}
}

func TestNegativeAmountsAreRejected(t *testing.T) {
	_, err := RenewalEmail(Renewal{Plan: PlanGuild, AmountCents: -1, Interval: "yearly"})
	if !errors.Is(err, ErrNegativeAmount) {
		t.Errorf("renewal err = %v", err)
	}
	_, err = SizeEmail(Size{RaidGuildCents: -1})
	if !errors.Is(err, ErrNegativeAmount) {
		t.Errorf("size err = %v", err)
	}
	_, err = SeatRemovedEmail(SeatRemoved{AmountCents: -1})
	if !errors.Is(err, ErrNegativeAmount) {
		t.Errorf("removed err = %v", err)
	}
}
