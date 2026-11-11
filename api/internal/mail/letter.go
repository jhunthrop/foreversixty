package mail

import (
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/copygen"
)

// Letter is one email's content, before it is drawn: every message the server sends is a
// Letter composed into the one branded shell (Compose), so the HTML and the plain text say the
// same thing in the same order.
type Letter struct {
	To, Subject string
	// Preheader is the inbox preview line under the subject; it never shows in the body.
	Preheader string
	// Heading is the panel's title; Lead the sentence under it.
	Heading, Lead string
	// Facts are the figure block under the lead (plan, date, price), the account page's order.
	Facts []Fact
	// Note is the closing sentence, after the facts; empty draws nothing.
	Note string
	// Action is the one button. Its URL is also printed under it for clients that drop links.
	Action Action
}

// Fact is one row of the figure block.
type Fact struct{ Label, Value string }

// Action is the letter's one link.
type Action struct{ Label, URL string }

// Compose draws the letter as a message with an HTML part and its plain-text alternative.
// now dates the footer's copyright line, the site footer's year rule.
func Compose(l Letter, now time.Time) (Message, error) {
	html, err := renderHTML(l, now)
	if err != nil {
		return Message{}, err
	}
	return Message{To: l.To, Subject: l.Subject, Text: renderText(l, now), HTML: html}, nil
}

// footer is the shell's closing lines: the site footer's own (shared.footer.*), then how to reach
// a person, whose address the HTML part links.
type footer struct {
	FanLine, Operator string
	// ContactLead and ContactTail are the contact line around the address.
	ContactLead, ContactTail, Email string
}

func footerOf(now time.Time) footer {
	lead, tail, _ := strings.Cut(copygen.MailFooterContact, "{email}")
	return footer{
		FanLine:     copygen.SharedFooterFanLine,
		Operator:    copygen.Fill(copygen.SharedFooterOperator, "{year}", strconv.Itoa(now.Year())),
		ContactLead: lead, ContactTail: tail, Email: copygen.MailFooterEmail,
	}
}

func (f footer) lines() []string {
	return []string{f.FanLine, f.Operator, f.ContactLead + f.Email + f.ContactTail}
}

// renderText is the plain-text alternative: the same blocks, one blank line apart.
func renderText(l Letter, now time.Time) string {
	blocks := []string{l.Heading, l.Lead}
	if len(l.Facts) > 0 {
		rows := make([]string, len(l.Facts))
		for i, f := range l.Facts {
			rows[i] = f.Label + ": " + f.Value
		}
		blocks = append(blocks, strings.Join(rows, "\n"))
	}
	if l.Note != "" {
		blocks = append(blocks, l.Note)
	}
	blocks = append(blocks, l.Action.Label+": "+l.Action.URL, textRule+"\n"+strings.Join(footerOf(now).lines(), "\n"))
	return strings.Join(blocks, "\n\n") + "\n"
}

// textRule sets the footer off in the plain-text part, the usual signature separator.
const textRule = "--"
