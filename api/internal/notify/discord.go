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
	// Thumbnail is an absolute image URL; empty omits it.
	Thumbnail string
	// Timestamp is the moment the event finished; the zero time omits it.
	Timestamp time.Time
	// Color is the embed's side stripe; zero omits it.
	Color int
}

// EmbedField is one name/value pair of an Embed.
type EmbedField struct {
	Name, Value string
	Inline      bool
}

// Poster is Discord, behind an interface so jobs and handlers test without
// the network. webhookURL is the full credential URL. Callers pass only
// ParsedURL.URL() or Sealer.Open(box) output, which are host-pinned to
// Discord; the HTTP poster follows no redirects.
type Poster interface {
	Post(ctx context.Context, webhookURL string, embed Embed) error
	// Lookup reads the webhook's channel id and name (GET on the webhook URL).
	Lookup(ctx context.Context, webhookURL string) (LookupResult, error)
	// Delete removes the webhook at Discord (DELETE on its URL). A 404 is
	// success: it is already gone.
	Delete(ctx context.Context, webhookURL string) error
}

// Transport failures, reported without the URL. The deliverer turns them
// into the plain sentences a member reads.
var (
	ErrTimeout     = errors.New("notify: discord did not answer in time")
	ErrUnreachable = errors.New("notify: discord unreachable")
)

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
	webhookSegment = regexp.MustCompile(`/?api/(?:v\d+/)?webhooks/\S*`)
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
type HTTPPoster struct {
	client    *http.Client
	avatarURL string
}

// NewHTTPPoster returns a Poster with the 5 s timeout. avatarURL is the
// absolute URL of the Forever Sixty seal every message carries; empty
// leaves the webhook's own avatar.
func NewHTTPPoster(avatarURL string) *HTTPPoster {
	return &HTTPPoster{avatarURL: avatarURL, client: newNoRedirectClient()}
}

// newNoRedirectClient is the 5 s client every Discord call uses. A redirect
// would carry a credential to another host; the 3xx is returned instead.
func newNoRedirectClient() *http.Client {
	return &http.Client{
		Timeout:       posterTimeout,
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
	}
}

type embedJSON struct {
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	URL         string      `json:"url,omitempty"`
	Fields      []fieldJSON `json:"fields,omitempty"`
	Footer      *footerJSON `json:"footer,omitempty"`
	Thumbnail   *imageJSON  `json:"thumbnail,omitempty"`
	Timestamp   string      `json:"timestamp,omitempty"`
	Color       int         `json:"color,omitempty"`
}

type imageJSON struct {
	URL string `json:"url"`
}

// messageJSON is the one request body every post sends: one embed, no
// content, and a mention list that pings no one.
type messageJSON struct {
	Username        string          `json:"username"`
	AvatarURL       string          `json:"avatar_url,omitempty"`
	Embeds          []embedJSON     `json:"embeds"`
	AllowedMentions allowedMentions `json:"allowed_mentions"`
}

type allowedMentions struct {
	Parse []string `json:"parse"`
}

// Username is the name every message is posted under.
const Username = "Forever Sixty"

func newMessage(avatarURL string, embed Embed) messageJSON {
	return messageJSON{
		Username: Username, AvatarURL: avatarURL, Embeds: []embedJSON{embed.wire()},
		AllowedMentions: allowedMentions{Parse: []string{}},
	}
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
	out := embedJSON{Title: e.Title, Description: e.Description, URL: e.URL, Color: e.Color}
	if e.Thumbnail != "" {
		out.Thumbnail = &imageJSON{URL: e.Thumbnail}
	}
	if !e.Timestamp.IsZero() {
		out.Timestamp = e.Timestamp.UTC().Format(time.RFC3339)
	}
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
	body, err := json.Marshal(newMessage(p.avatarURL, embed))
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

// LookupResult is what Discord says about a webhook. Name may be empty.
type LookupResult struct{ ChannelID, GuildID, Name string }

// Lookup returns the webhook's channel_id and name.
func (p *HTTPPoster) Lookup(ctx context.Context, webhookURL string) (LookupResult, error) {
	res, err := p.do(ctx, http.MethodGet, webhookURL, nil)
	if err != nil {
		return LookupResult{}, err
	}
	defer res.Body.Close()
	if err := statusError(res); err != nil {
		return LookupResult{}, err
	}
	var out struct {
		ChannelID string `json:"channel_id"`
		GuildID   string `json:"guild_id"`
		Name      string `json:"name"`
	}
	if err := json.NewDecoder(io.LimitReader(res.Body, maxResponseLen)).Decode(&out); err != nil || out.ChannelID == "" {
		return LookupResult{}, errors.New("notify: webhook answered without a channel id")
	}
	return LookupResult{ChannelID: out.ChannelID, GuildID: out.GuildID, Name: out.Name}, nil
}

// Delete removes the webhook; any 2xx or a 404 is success.
func (p *HTTPPoster) Delete(ctx context.Context, webhookURL string) error {
	res, err := p.do(ctx, http.MethodDelete, webhookURL, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return nil
	}
	return statusError(res)
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
			return nil, fmt.Errorf("%w: %w", ErrTimeout, context.DeadlineExceeded)
		}
		return nil, ErrUnreachable
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
	return &DiscordError{Status: res.StatusCode, Body: sanitize(discordMessage(raw, res.StatusCode))}
}

// discordMessage is the "message" of Discord's JSON error body, or the
// standard status text when the body is not that (an HTML proxy page).
func discordMessage(raw []byte, status int) string {
	var body struct {
		Message string `json:"message"`
	}
	if json.Unmarshal(raw, &body) == nil && strings.TrimSpace(body.Message) != "" {
		return body.Message
	}
	return http.StatusText(status)
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
	LookupName      string
	LookupGuildID   string
	// DeleteErr, when set, is returned by every Delete.
	DeleteErr error
	posts     []PostCall
	deletes   []string
}

func (f *FakePoster) Post(_ context.Context, webhookURL string, embed Embed) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.posts = append(f.posts, PostCall{URL: webhookURL, Embed: embed})
	return f.PostErr
}

func (f *FakePoster) Lookup(context.Context, string) (LookupResult, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return LookupResult{ChannelID: f.LookupChannelID, GuildID: f.LookupGuildID, Name: f.LookupName}, f.LookupErr
}

// Posts returns a copy of the recorded Post calls.
func (f *FakePoster) Posts() []PostCall {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]PostCall(nil), f.posts...)
}

// Delete records the URL and returns DeleteErr.
func (f *FakePoster) Delete(_ context.Context, webhookURL string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.deletes = append(f.deletes, webhookURL)
	return f.DeleteErr
}

// Deletes returns the URLs Delete was called with.
func (f *FakePoster) Deletes() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string(nil), f.deletes...)
}
