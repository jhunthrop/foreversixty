package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	posterTimeout  = 5 * time.Second
	maxErrorRunes  = 200
	maxResponseLen = 4096
)

// Embed is the one message shape the channel receives.
type Embed struct {
	Title       string
	Description string
	Fields      []EmbedField
	URL         string
	Footer      string
}

// EmbedField is one name/value pair of an Embed.
type EmbedField struct {
	Name, Value string
	Inline      bool
}

// Poster is Discord, behind an interface so jobs and handlers test without
// the network. webhookURL is the full credential URL.
type Poster interface {
	Post(ctx context.Context, webhookURL string, embed Embed) error
	// Lookup reads the webhook's channel id (GET on the webhook URL).
	Lookup(ctx context.Context, webhookURL string) (channelID string, err error)
}

// DiscordError is a non-2xx answer from Discord. Body is already sanitised.
type DiscordError struct {
	Status int
	Body   string
}

func (e *DiscordError) Error() string { return fmt.Sprintf("discord: status %d: %s", e.Status, e.Body) }

// RateLimited reports a 429, which never counts toward the failure limit.
func (e *DiscordError) RateLimited() bool { return e.Status == http.StatusTooManyRequests }

var (
	urlPattern     = regexp.MustCompile(`https?://\S*`)
	webhookSegment = regexp.MustCompile(`/?api/webhooks/\S*`)
)

// sanitize strips any URL and webhook path segment from Discord's text and
// truncates it, so a credential never reaches a row or the page.
func sanitize(text string) string {
	text = webhookSegment.ReplaceAllString(urlPattern.ReplaceAllString(text, ""), "")
	text = strings.Join(strings.Fields(text), " ")
	if runes := []rune(text); len(runes) > maxErrorRunes {
		text = string(runes[:maxErrorRunes])
	}
	return text
}

// HTTPPoster posts to Discord with a short timeout and no retry.
type HTTPPoster struct{ client *http.Client }

// NewHTTPPoster returns a Poster with the 5 s timeout.
func NewHTTPPoster() *HTTPPoster { return &HTTPPoster{client: &http.Client{Timeout: posterTimeout}} }

type embedJSON struct {
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	URL         string      `json:"url,omitempty"`
	Fields      []fieldJSON `json:"fields,omitempty"`
	Footer      *footerJSON `json:"footer,omitempty"`
}

type fieldJSON struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type footerJSON struct {
	Text string `json:"text"`
}

func (e Embed) wire() embedJSON {
	out := embedJSON{Title: e.Title, Description: e.Description, URL: e.URL}
	for _, f := range e.Fields {
		out.Fields = append(out.Fields, fieldJSON(f))
	}
	if e.Footer != "" {
		out.Footer = &footerJSON{Text: e.Footer}
	}
	return out
}

// Post sends one embed that mentions no one; any 2xx is success.
func (p *HTTPPoster) Post(ctx context.Context, webhookURL string, embed Embed) error {
	payload := map[string]any{
		"embeds":           []embedJSON{embed.wire()},
		"allowed_mentions": map[string]any{"parse": []string{}},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("notify: encode embed: %w", err)
	}
	res, err := p.do(ctx, http.MethodPost, webhookURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return statusError(res)
}

// Lookup returns the webhook's channel_id.
func (p *HTTPPoster) Lookup(ctx context.Context, webhookURL string) (string, error) {
	res, err := p.do(ctx, http.MethodGet, webhookURL, nil)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if err := statusError(res); err != nil {
		return "", err
	}
	var out struct {
		ChannelID string `json:"channel_id"`
	}
	if err := json.NewDecoder(io.LimitReader(res.Body, maxResponseLen)).Decode(&out); err != nil || out.ChannelID == "" {
		return "", errors.New("notify: webhook answered without a channel id")
	}
	return out.ChannelID, nil
}

// do sends the request; a transport failure is reported without the URL
// (net/url errors embed it).
func (p *HTTPPoster) do(ctx context.Context, method, webhookURL string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, webhookURL, body)
	if err != nil {
		return nil, errors.New("notify: bad webhook request")
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := p.client.Do(req)
	if err != nil {
		if ctx.Err() != nil || isTimeout(err) {
			return nil, errors.New("notify: discord did not answer in time")
		}
		return nil, errors.New("notify: discord unreachable")
	}
	return res, nil
}

func isTimeout(err error) bool {
	var t interface{ Timeout() bool }
	return errors.As(err, &t) && t.Timeout()
}

func statusError(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return nil
	}
	raw, _ := io.ReadAll(io.LimitReader(res.Body, maxResponseLen))
	return &DiscordError{Status: res.StatusCode, Body: sanitize(string(raw))}
}

// PostCall is one recorded FakePoster.Post.
type PostCall struct {
	URL   string
	Embed Embed
}

// FakePoster is the test double: it records calls and returns scripted errors.
type FakePoster struct {
	mu sync.Mutex
	// PostErr, when set, is returned by every Post (a *DiscordError with
	// Status 429 scripts a rate limit).
	PostErr error
	// LookupErr, when set, is returned by Lookup; otherwise LookupChannelID.
	LookupErr       error
	LookupChannelID string
	posts           []PostCall
}

func (f *FakePoster) Post(_ context.Context, webhookURL string, embed Embed) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.posts = append(f.posts, PostCall{URL: webhookURL, Embed: embed})
	return f.PostErr
}

func (f *FakePoster) Lookup(context.Context, string) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.LookupChannelID, f.LookupErr
}

// Posts returns a copy of the recorded Post calls.
func (f *FakePoster) Posts() []PostCall {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]PostCall(nil), f.posts...)
}
