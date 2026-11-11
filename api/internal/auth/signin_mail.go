package auth

import (
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// signInEmail is the emailed sign-in link (POST /v1/auth/email), in the shell every email shares.
func signInEmail(to, link string, now time.Time) (mail.Message, error) {
	minutes := strconv.Itoa(int(LoginTokenTTL / time.Minute))
	return mail.Compose(mail.Letter{
		To:        to,
		Subject:   copygen.MailSignInSubject,
		Preheader: copygen.Fill(copygen.MailSignInPreheader, "{minutes}", minutes),
		Heading:   copygen.MailSignInHeading,
		Lead:      copygen.Fill(copygen.MailSignInLead, "{minutes}", minutes),
		Action:    mail.Action{Label: copygen.MailSignInAction, URL: link},
	}, now)
}
