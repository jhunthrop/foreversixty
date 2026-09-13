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

type Resend struct {
	apiKey   string
	from     string
	client   *http.Client
	endpoint string
}

func NewResend(apiKey, from string, client *http.Client) *Resend {
	if client == nil {
		client = &http.Client{Timeout: resendClientTimeout}
	}
	return &Resend{apiKey: apiKey, from: from, client: client, endpoint: "https://api.resend.com/emails"}
}

func (r *Resend) Send(ctx context.Context, m Message) error {
	body, _ := json.Marshal(map[string]any{"from": r.from, "to": []string{m.To}, "subject": m.Subject, "text": m.Text})
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
