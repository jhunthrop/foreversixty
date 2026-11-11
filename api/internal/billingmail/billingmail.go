package billingmail

import (
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// Plan is the plan name as it reads in an email.
type Plan string

// The plans a renewal email can name, in the site's words.
const (
	PlanPremium   Plan = copygen.PlanPremium
	PlanGuild     Plan = copygen.MailPlanGuild
	PlanRaidGuild Plan = copygen.MailPlanRaidGuild
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
	To, Guild string
	Members   int
	// MemberLimit is the member count the Guild tier covers (catalog.GuildTierMemberLimit).
	MemberLimit    int
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

// RenewalEmail builds the renewal notice: the plan and date in the heading, the price (and the
// guild a guild plan is for) in the figure block, and how to stop it. Each fact is said once.
func RenewalEmail(r Renewal, now time.Time) (mail.Message, error) {
	price, err := FormatPrice(r.AmountCents, r.Interval)
	if err != nil {
		return mail.Message{}, err
	}
	date := EmailDate(r.RenewsOn)
	plan := string(r.Plan)
	var facts []mail.Fact
	if r.Guild != "" {
		facts = append(facts, mail.Fact{Label: copygen.MailFactGuild, Value: r.Guild})
	}
	facts = append(facts, mail.Fact{Label: copygen.MailFactPrice, Value: price})
	note := ""
	if r.Founder {
		note = copygen.PlanFounderNote
	}
	return mail.Compose(mail.Letter{
		To:        r.To,
		Subject:   copygen.Fill(copygen.MailRenewalSubject, "{plan}", plan, "{date}", date),
		Preheader: copygen.Fill(copygen.MailRenewalPreheader, "{price}", price, "{date}", date),
		Heading:   copygen.Fill(copygen.MailRenewalHeading, "{plan}", plan, "{date}", date),
		Lead:      copygen.MailRenewalLead,
		Facts:     facts,
		Note:      note,
		Action:    mail.Action{Label: copygen.MailRenewalAction, URL: r.ManageURL},
	}, now)
}

// SizeEmail builds the notice for a Guild-tier guild over its member limit: the count and the
// limit, then what Raid guild costs and when the plan renews.
func SizeEmail(s Size, now time.Time) (mail.Message, error) {
	price, err := FormatPrice(s.RaidGuildCents, s.Interval)
	if err != nil {
		return mail.Message{}, err
	}
	if s.MemberLimit <= 0 {
		return mail.Message{}, ErrNoMemberLimit
	}
	members, limit := strconv.Itoa(s.Members), strconv.Itoa(s.MemberLimit)
	return mail.Compose(mail.Letter{
		To:        s.To,
		Subject:   copygen.Fill(copygen.MailSizeSubject, "{guild}", s.Guild, "{members}", members, "{limit}", limit),
		Preheader: copygen.Fill(copygen.MailSizePreheader, "{price}", price),
		Heading:   copygen.Fill(copygen.MailSizeHeading, "{guild}", s.Guild, "{members}", members),
		Lead:      copygen.Fill(copygen.MailSizeLead, "{limit}", limit),
		Facts: []mail.Fact{
			{Label: string(PlanRaidGuild), Value: price},
			{Label: copygen.MailFactRenews, Value: EmailDate(s.RenewsOn)},
		},
		Action: mail.Action{Label: copygen.MailSizeAction, URL: s.SettingsURL},
	}, now)
}

// SeatOfferedEmail builds the offer to a paying member.
func SeatOfferedEmail(o SeatOffered, now time.Time) (mail.Message, error) {
	return mail.Compose(mail.Letter{
		To:        o.To,
		Subject:   copygen.Fill(copygen.MailSeatOfferedSubject, "{guild}", o.Guild),
		Preheader: copygen.MailSeatOfferedPreheader,
		Heading:   copygen.Fill(copygen.MailSeatOfferedHeading, "{guild}", o.Guild),
		Lead:      copygen.Fill(copygen.MailSeatOfferedLead, "{from}", o.From),
		Note:      copygen.Fill(copygen.MailSeatOfferedNote, "{date}", EmailDate(o.OwnEndsOn)),
		Action:    mail.Action{Label: copygen.MailSeatOfferedAction, URL: o.AccountURL},
	}, now)
}

// SeatGivenEmail builds the notice to a non-paying member, in the account page's own words for a
// given seat (plan.seatGivenBy).
func SeatGivenEmail(g SeatGiven, now time.Time) (mail.Message, error) {
	return mail.Compose(mail.Letter{
		To:        g.To,
		Subject:   copygen.Fill(copygen.MailSeatGivenSubject, "{guild}", g.Guild),
		Preheader: copygen.MailSeatGivenPreheader,
		Heading:   copygen.Fill(copygen.MailSeatGivenHeading, "{guild}", g.Guild),
		Lead:      copygen.Fill(copygen.PlanSeatGivenBy, "{name}", g.From),
		Action:    mail.Action{Label: copygen.MailSeatGivenAction, URL: g.AccountURL},
	}, now)
}

// SeatRemovedEmail builds the come-back notice; its button carries the price.
func SeatRemovedEmail(r SeatRemoved, now time.Time) (mail.Message, error) {
	price, err := FormatPrice(r.AmountCents, r.Interval)
	if err != nil {
		return mail.Message{}, err
	}
	return mail.Compose(mail.Letter{
		To:        r.To,
		Subject:   copygen.Fill(copygen.MailSeatRemovedSubject, "{guild}", r.Guild),
		Preheader: copygen.MailSeatRemovedPreheader,
		Heading:   copygen.MailSeatRemovedHeading,
		Lead:      copygen.Fill(copygen.MailSeatRemovedLead, "{guild}", r.Guild, "{date}", EmailDate(r.EndedOn)),
		Note:      copygen.MailSeatRemovedNote,
		Action:    mail.Action{Label: copygen.Fill(copygen.MailSeatRemovedAction, "{price}", price), URL: r.ComebackURL},
	}, now)
}
