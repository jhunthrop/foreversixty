package mail

import "context"

// Message is one email as it is sent: Text is the plain-text alternative of HTML, both drawn by
// Compose from one Letter.
type Message struct {
	To      string
	Subject string
	Text    string
	HTML    string
}

type Mailer interface {
	Send(ctx context.Context, m Message) error
}
