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

// RenewalEmail builds the renewal notice: the plan, the date and the price as the account page
// lists them, and how to stop it.
func RenewalEmail(r Renewal, now time.Time) (mail.Message, error) {
	price, err := FormatPrice(r.AmountCents, r.Interval)
	if err != nil {
		return mail.Message{}, err
	}
	date := EmailDate(r.RenewsOn)
	plan := string(r.Plan)
	planFact := plan
	if r.Guild != "" {
		planFact = copygen.Fill(copygen.MailFactPlanForGuild, "{plan}", plan, "{guild}", r.Guild)
	}
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
		Facts: []mail.Fact{
			{Label: copygen.MailFactPlan, Value: planFact},
			{Label: copygen.MailFactRenews, Value: date},
			{Label: copygen.MailFactPrice, Value: price},
		},
		Note:   note,
		Action: mail.Action{Label: copygen.MailRenewalAction, URL: r.ManageURL},
	}, now)
}

// SizeEmail builds the notice for a Guild-tier guild over its member limit.
func SizeEmail(s Size, now time.Time) (mail.Message, error) {
	price, err := FormatPrice(s.RaidGuildCents, s.Interval)
	if err != nil {
		return mail.Message{}, err
	}
	if s.MemberLimit <= 0 {
		return mail.Message{}, ErrNoMemberLimit
	}
	date := EmailDate(s.RenewsOn)
	members, limit := strconv.Itoa(s.Members), strconv.Itoa(s.MemberLimit)
	return mail.Compose(mail.Letter{
		To:        s.To,
		Subject:   copygen.Fill(copygen.MailSizeSubject, "{guild}", s.Guild, "{members}", members, "{limit}", limit),
		Preheader: copygen.MailSizePreheader,
		Heading:   copygen.Fill(copygen.MailSizeHeading, "{guild}", s.Guild),
		Lead:      copygen.Fill(copygen.MailSizeLead, "{guild}", s.Guild, "{members}", members, "{limit}", limit),
		Facts: []mail.Fact{
			{Label: copygen.MailFactGuild, Value: s.Guild},
			{Label: copygen.MailFactMembers, Value: copygen.Fill(copygen.MailFactMembersOf, "{members}", members, "{limit}", limit)},
			{Label: copygen.MailFactRenews, Value: date},
			{Label: string(PlanRaidGuild), Value: price},
		},
		Note:   copygen.Fill(copygen.MailSizeNote, "{date}", date, "{price}", price),
		Action: mail.Action{Label: copygen.MailSizeAction, URL: s.SettingsURL},
	}, now)
}

// SeatOfferedEmail builds the offer to a paying member.
func SeatOfferedEmail(o SeatOffered, now time.Time) (mail.Message, error) {
	ends := EmailDate(o.OwnEndsOn)
	return mail.Compose(mail.Letter{
		To:        o.To,
		Subject:   copygen.Fill(copygen.MailSeatOfferedSubject, "{guild}", o.Guild),
		Preheader: copygen.MailSeatOfferedPreheader,
		Heading:   copygen.Fill(copygen.MailSeatOfferedHeading, "{guild}", o.Guild),
		Lead:      copygen.Fill(copygen.MailSeatOfferedLead, "{from}", o.From, "{guild}", o.Guild),
		Facts: []mail.Fact{
			{Label: copygen.MailFactFrom, Value: o.From},
			{Label: copygen.MailFactOwnEnds, Value: ends},
		},
		Note:   copygen.Fill(copygen.MailSeatOfferedNote, "{date}", ends),
		Action: mail.Action{Label: copygen.MailSeatOfferedAction, URL: o.AccountURL},
	}, now)
}

// SeatGivenEmail builds the notice to a non-paying member.
func SeatGivenEmail(g SeatGiven, now time.Time) (mail.Message, error) {
	return mail.Compose(mail.Letter{
		To:        g.To,
		Subject:   copygen.Fill(copygen.MailSeatGivenSubject, "{guild}", g.Guild),
		Preheader: copygen.MailSeatGivenPreheader,
		Heading:   copygen.Fill(copygen.MailSeatGivenHeading, "{guild}", g.Guild),
		Lead:      copygen.Fill(copygen.MailSeatGivenLead, "{from}", g.From, "{guild}", g.Guild),
		Note:      copygen.MailSeatGivenNote,
		Action:    mail.Action{Label: copygen.MailSeatGivenAction, URL: g.AccountURL},
	}, now)
}

// SeatRemovedEmail builds the come-back notice.
func SeatRemovedEmail(r SeatRemoved, now time.Time) (mail.Message, error) {
	price, err := FormatPrice(r.AmountCents, r.Interval)
	if err != nil {
		return mail.Message{}, err
	}
	ended := EmailDate(r.EndedOn)
	return mail.Compose(mail.Letter{
		To:        r.To,
		Subject:   copygen.Fill(copygen.MailSeatRemovedSubject, "{guild}", r.Guild),
		Preheader: copygen.MailSeatRemovedPreheader,
		Heading:   copygen.MailSeatRemovedHeading,
		Lead:      copygen.Fill(copygen.MailSeatRemovedLead, "{guild}", r.Guild, "{date}", ended),
		Facts: []mail.Fact{
			{Label: copygen.MailFactGuild, Value: r.Guild},
			{Label: copygen.MailFactEnded, Value: ended},
			{Label: copygen.MailFactPrice, Value: price},
		},
		Note:   copygen.MailSeatRemovedNote,
		Action: mail.Action{Label: copygen.Fill(copygen.MailSeatRemovedAction, "{price}", price), URL: r.ComebackURL},
	}, now)
}
