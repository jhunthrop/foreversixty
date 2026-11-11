package subscribe

import (
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	mailer "github.com/jhunthrop/foreversixty/api/internal/mail"
)

// confirmationEmail is the double opt-in mail, in the shell every email shares.
func confirmationEmail(to, confirmURL string, now time.Time) (mailer.Message, error) {
	return mailer.Compose(mailer.Letter{
		To:        to,
		Subject:   copygen.MailSubscribeSubject,
		Preheader: copygen.MailSubscribePreheader,
		Heading:   copygen.MailSubscribeHeading,
		Lead:      copygen.MailSubscribeLead,
		Note:      copygen.MailSubscribeNote,
		Action:    mailer.Action{Label: copygen.MailSubscribeAction, URL: confirmURL},
	}, now)
}
