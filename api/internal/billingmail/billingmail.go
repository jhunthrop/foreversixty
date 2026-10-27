package billingmail

import (
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// Plan is the plan name as it reads in an email.
type Plan string

// The plans a renewal email can name.
const (
	PlanPremium   Plan = "Premium"
	PlanGuild     Plan = "Guild"
	PlanRaidGuild Plan = "Raid guild"
)

// Renewal is the email sent 7 days before a yearly renewal. Guild is empty
// for Premium.
type Renewal struct {
	To, Guild   string
	Plan        Plan
	RenewsOn    time.Time
	AmountCents int64
	Interval    string
	Founder     bool
	ManageURL   string
}

// Size is the email sent 14 days before renewal when a Guild plan has more
// than 25 verified members.
type Size struct {
	To, Guild      string
	Members        int
	RenewsOn       time.Time
	RaidGuildCents int64
	Interval       string
	SettingsURL    string
}

// SeatOffered goes to a member who pays for their own Premium.
type SeatOffered struct {
	To, Guild, From string
	OwnEndsOn       time.Time
	AccountURL      string
}

// SeatGiven goes to a member who does not pay.
type SeatGiven struct {
	To, Guild, From, AccountURL string
}

// SeatRemoved goes to someone who gave up their own Premium for a seat.
type SeatRemoved struct {
	To, Guild   string
	EndedOn     time.Time
	AmountCents int64
	Interval    string
	ComebackURL string
}

// compose joins the first sentence and the single link into the body shape.
func compose(to, subject, sentence, label, url string) mail.Message {
	return mail.Message{To: to, Subject: subject, Text: fmt.Sprintf("%s\n\n%s: %s", sentence, label, url)}
}

// RenewalEmail builds the renewal notice.
func RenewalEmail(r Renewal) (mail.Message, error) {
	price, err := FormatPrice(r.AmountCents, r.Interval)
	if err != nil {
		return mail.Message{}, err
	}
	subject := fmt.Sprintf("Your %s plan renews on %s", r.Plan, EmailDayMonth(r.RenewsOn))
	forGuild, founder := "", ""
	if r.Guild != "" {
		forGuild = " for " + r.Guild
	}
	if r.Founder {
		founder = ", your founder price"
	}
	sentence := fmt.Sprintf("Your %s plan%s renews on %s for %s%s; to stop it, cancel before then and nothing more is charged.",
		r.Plan, forGuild, EmailDate(r.RenewsOn), price, founder)
	return compose(r.To, subject, sentence, "Manage or cancel", r.ManageURL), nil
}

// SizeEmail builds the over-25 notice.
func SizeEmail(s Size) (mail.Message, error) {
	price, err := FormatPrice(s.RaidGuildCents, s.Interval)
	if err != nil {
		return mail.Message{}, err
	}
	subject := fmt.Sprintf("%s has %d members; Guild covers 25", s.Guild, s.Members)
	sentence := fmt.Sprintf("%s has %d verified members and its Guild plan covers up to 25, so when it renews on %s you may want Raid guild at %s; nothing switches off either way.",
		s.Guild, s.Members, EmailDate(s.RenewsOn), price)
	return compose(s.To, subject, sentence, "Look at Raid guild", s.SettingsURL), nil
}

// SeatOfferedEmail builds the offer to a paying member.
func SeatOfferedEmail(o SeatOffered) mail.Message {
	subject := o.Guild + " offers you a Premium seat"
	sentence := fmt.Sprintf("%s offered you a Premium seat in %s; if you accept, your own Premium stops renewing and ends on %s, the seat starts then, and you are not charged twice.",
		o.From, o.Guild, EmailDate(o.OwnEndsOn))
	return compose(o.To, subject, sentence, "Accept or decline", o.AccountURL)
}

// SeatGivenEmail builds the notice to a non-paying member.
func SeatGivenEmail(g SeatGiven) mail.Message {
	subject := "You have Premium from " + g.Guild
	sentence := fmt.Sprintf("%s gave you a Premium seat in %s, so you have Premium for as long as the guild's plan lasts, with nothing billed to you.",
		g.From, g.Guild)
	return compose(g.To, subject, sentence, "See your plan", g.AccountURL)
}

// SeatRemovedEmail builds the come-back notice.
func SeatRemovedEmail(r SeatRemoved) (mail.Message, error) {
	price, err := FormatPrice(r.AmountCents, r.Interval)
	if err != nil {
		return mail.Message{}, err
	}
	subject := fmt.Sprintf("Your Premium seat from %s has ended", r.Guild)
	sentence := fmt.Sprintf("Your Premium seat in %s ended on %s, and because you gave up your own Premium for it, this link brings it back at %s, the price you had.",
		r.Guild, EmailDate(r.EndedOn), price)
	return compose(r.To, subject, sentence, "Come back at "+price, r.ComebackURL), nil
}
