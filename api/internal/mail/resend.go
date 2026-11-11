package mail

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// resendClientTimeout bounds a Resend API call when NewResend is not given
// an explicit client, so a slow or hanging provider can't tie up a caller
// indefinitely.
const resendClientTimeout = 10 * time.Second

// Sender is who a message comes from (config MAIL_FROM) and where a reply goes (MAIL_REPLY_TO);
// an empty ReplyTo leaves replies to the From address.
type Sender struct {
	From, ReplyTo string
}

type Resend struct {
	apiKey   string
	sender   Sender
	client   *http.Client
	endpoint string
}

func NewResend(apiKey string, sender Sender, client *http.Client) *Resend {
	if client == nil {
		client = &http.Client{Timeout: resendClientTimeout}
	}
	return &Resend{apiKey: apiKey, sender: sender, client: client, endpoint: "https://api.resend.com/emails"}
}

// resendEmail is the body of Resend's POST /emails; html and reply_to are left out when empty.
type resendEmail struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	ReplyTo string   `json:"reply_to,omitempty"`
	Subject string   `json:"subject"`
	Text    string   `json:"text"`
	HTML    string   `json:"html,omitempty"`
}

func (r *Resend) Send(ctx context.Context, m Message) error {
	body, err := json.Marshal(resendEmail{
		From: r.sender.From, To: []string{m.To}, ReplyTo: r.sender.ReplyTo,
		Subject: m.Subject, Text: m.Text, HTML: m.HTML,
	})
	if err != nil {
		return fmt.Errorf("mail: encode: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("mail: request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+r.apiKey)
	req.Header.Set("Content-Type", "application/json")
	res, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("mail: send: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		return fmt.Errorf("mail: resend status %d", res.StatusCode)
	}
	return nil
}
