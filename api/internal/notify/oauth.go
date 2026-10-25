package notify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	discordAuthorizeURL = "https://discord.com/oauth2/authorize"
	discordTokenURL     = "https://discord.com/api/oauth2/token"
	// connectScope is the one scope the connect asks for: Discord creates an
	// incoming webhook in the channel the member picks.
	connectScope = "webhook.incoming"
	// maxTokenResponseBytes bounds the token answer we read.
	maxTokenResponseBytes = 64 << 10
)

// OAuthWebhook is the webhook Discord hands back from the code exchange. It
// holds the credential (Token, URL): never log or return it.
type OAuthWebhook struct {
	ID, Token, URL, ChannelID, GuildID, Name string
}

// OAuthClient is Discord's OAuth, behind an interface so handlers test
// without the network.
type OAuthClient interface {
	// AuthURL is where the browser goes to pick a channel.
	AuthURL(state string) string
	// Exchange trades the authorization code for the webhook. The error never
	// carries the code or any response body.
	Exchange(ctx context.Context, code string) (OAuthWebhook, error)
}

// HTTPOAuthClient talks to discord.com.
type HTTPOAuthClient struct {
	ClientID, ClientSecret, RedirectURI string
	// HTTP is the client; NewHTTPOAuthClient gives it the 5 s timeout and no redirects.
	HTTP *http.Client
	// TokenURL overrides discordTokenURL; tests point it at httptest.
	TokenURL string
}

// NewHTTPOAuthClient returns the production client.
func NewHTTPOAuthClient(clientID, clientSecret, redirectURI string) *HTTPOAuthClient {
	return &HTTPOAuthClient{
		ClientID: clientID, ClientSecret: clientSecret, RedirectURI: redirectURI,
		HTTP: newNoRedirectClient(),
	}
}

// AuthURL builds the authorize URL with every value query-encoded.
func (c *HTTPOAuthClient) AuthURL(state string) string {
	q := url.Values{}
	q.Set("response_type", "code")
	q.Set("client_id", c.ClientID)
	q.Set("scope", connectScope)
	q.Set("redirect_uri", c.RedirectURI)
	q.Set("state", state)
	return discordAuthorizeURL + "?" + q.Encode()
}

// ErrExchange is a failed code exchange; its text carries at most a status.
var ErrExchange = errors.New("notify: discord code exchange failed")

// Exchange posts the code form-encoded. Discord's answer also holds access
// and refresh tokens; only the webhook object is decoded.
func (c *HTTPOAuthClient) Exchange(ctx context.Context, code string) (OAuthWebhook, error) {
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("redirect_uri", c.RedirectURI)
	form.Set("client_id", c.ClientID)
	form.Set("client_secret", c.ClientSecret)
	endpoint := c.TokenURL
	if endpoint == "" {
		endpoint = discordTokenURL
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return OAuthWebhook{}, ErrExchange
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	res, err := c.HTTP.Do(req)
	if err != nil {
		// net/url errors embed the URL; none of it is useful here.
		return OAuthWebhook{}, fmt.Errorf("%w: no answer", ErrExchange)
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return OAuthWebhook{}, fmt.Errorf("%w: status %d", ErrExchange, res.StatusCode)
	}
	var out struct {
		Webhook *struct {
			ID        string `json:"id"`
			Token     string `json:"token"`
			URL       string `json:"url"`
			ChannelID string `json:"channel_id"`
			GuildID   string `json:"guild_id"`
			Name      string `json:"name"`
		} `json:"webhook"`
	}
	if err := json.NewDecoder(io.LimitReader(res.Body, maxTokenResponseBytes)).Decode(&out); err != nil || out.Webhook == nil {
		return OAuthWebhook{}, fmt.Errorf("%w: no webhook in the answer", ErrExchange)
	}
	w := out.Webhook
	return OAuthWebhook{ID: w.ID, Token: w.Token, URL: w.URL, ChannelID: w.ChannelID, GuildID: w.GuildID, Name: w.Name}, nil
}

// FakeOAuthClient is the test double: it records the codes it was given.
type FakeOAuthClient struct {
	Webhook OAuthWebhook
	Err     error

	mu    sync.Mutex
	codes []string
}

// AuthURL is a recognisable stand-in carrying the state.
func (f *FakeOAuthClient) AuthURL(state string) string {
	return discordAuthorizeURL + "?state=" + url.QueryEscape(state)
}

// Exchange records the code and answers the scripted webhook or error.
func (f *FakeOAuthClient) Exchange(_ context.Context, code string) (OAuthWebhook, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.codes = append(f.codes, code)
	return f.Webhook, f.Err
}

// Codes returns the codes Exchange was called with.
func (f *FakeOAuthClient) Codes() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string(nil), f.codes...)
}
