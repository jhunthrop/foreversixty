package discordbot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	authorizeURL = "https://discord.com/oauth2/authorize"
	// InstallPermissions is View Channel, Send Messages, Embed Links and Read
	// Message History.
	InstallPermissions = PermViewChannel | PermSendMessages | PermEmbedLinks | PermReadMessageHistory

	installScope  = "bot applications.commands"
	identityScope = "identify"
)

// ErrExchange is a failed code exchange; its text carries at most a status.
var ErrExchange = errors.New("discordbot: discord code exchange failed")

// InstalledServer is the server a member added the bot to.
type InstalledServer struct{ ID, Name string }

// Identity is the Discord account a member linked.
type Identity struct{ ID, DisplayName string }

// OAuth is Discord's OAuth as the bot uses it, behind an interface so
// handlers test without the network.
type OAuth interface {
	InstallURL(state string) string
	LinkURL(state string) string
	ExchangeInstall(ctx context.Context, code string) (InstalledServer, error)
	ExchangeIdentity(ctx context.Context, code string) (Identity, error)
}

// HTTPOAuth talks to discord.com.
type HTTPOAuth struct {
	// BaseURL overrides the API base; tests point it at httptest.
	BaseURL      string
	clientID     string
	clientSecret string
	redirectURI  string
	client       *http.Client
}

// NewHTTPOAuth returns the production client.
func NewHTTPOAuth(clientID, clientSecret, redirectURI string) *HTTPOAuth {
	return &HTTPOAuth{BaseURL: apiBase, clientID: clientID, clientSecret: clientSecret, redirectURI: redirectURI, client: newClient()}
}

func (o *HTTPOAuth) authorize(scope, state string, extra url.Values) string {
	q := url.Values{}
	q.Set("response_type", "code")
	q.Set("client_id", o.clientID)
	q.Set("scope", scope)
	q.Set("redirect_uri", o.redirectURI)
	q.Set("state", state)
	for k, v := range extra {
		q[k] = v
	}
	return authorizeURL + "?" + q.Encode()
}

// InstallURL is where a member adds the bot to a server.
func (o *HTTPOAuth) InstallURL(state string) string {
	return o.authorize(installScope, state, url.Values{"permissions": {fmt.Sprint(InstallPermissions)}})
}

// LinkURL is where a member proves which Discord account is theirs.
func (o *HTTPOAuth) LinkURL(state string) string { return o.authorize(identityScope, state, nil) }

// tokenAnswer is what the token endpoint returns; the guild is present on an
// install.
type tokenAnswer struct {
	AccessToken string `json:"access_token"`
	Guild       *struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"guild"`
}

// exchange trades the code form-encoded. Errors carry no code or body.
func (o *HTTPOAuth) exchange(ctx context.Context, code string) (tokenAnswer, error) {
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("redirect_uri", o.redirectURI)
	form.Set("client_id", o.clientID)
	form.Set("client_secret", o.clientSecret)
	var out tokenAnswer
	if err := o.send(ctx, http.MethodPost, "/oauth2/token", strings.NewReader(form.Encode()), "application/x-www-form-urlencoded", "", &out); err != nil {
		return tokenAnswer{}, err
	}
	return out, nil
}

func (o *HTTPOAuth) send(ctx context.Context, method, path string, body io.Reader, contentType, bearer string, out any) error {
	req, err := http.NewRequestWithContext(ctx, method, o.BaseURL+path, body)
	if err != nil {
		return ErrExchange
	}
	req.Header.Set("Accept", "application/json")
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if bearer != "" {
		req.Header.Set("Authorization", "Bearer "+bearer)
	}
	res, err := o.client.Do(req)
	if err != nil {
		return fmt.Errorf("%w: no answer", ErrExchange)
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("%w: status %d", ErrExchange, res.StatusCode)
	}
	if err := json.NewDecoder(io.LimitReader(res.Body, maxAnswerBytes)).Decode(out); err != nil {
		return fmt.Errorf("%w: unreadable answer", ErrExchange)
	}
	return nil
}

// ExchangeInstall returns the server named in the token answer's guild.
func (o *HTTPOAuth) ExchangeInstall(ctx context.Context, code string) (InstalledServer, error) {
	answer, err := o.exchange(ctx, code)
	if err != nil {
		return InstalledServer{}, err
	}
	if answer.Guild == nil || answer.Guild.ID == "" {
		return InstalledServer{}, fmt.Errorf("%w: no server in the answer", ErrExchange)
	}
	return InstalledServer{ID: answer.Guild.ID, Name: answer.Guild.Name}, nil
}

// ExchangeIdentity reads the member's account once; the bearer token never
// leaves this function.
func (o *HTTPOAuth) ExchangeIdentity(ctx context.Context, code string) (Identity, error) {
	answer, err := o.exchange(ctx, code)
	if err != nil {
		return Identity{}, err
	}
	if answer.AccessToken == "" {
		return Identity{}, fmt.Errorf("%w: no token in the answer", ErrExchange)
	}
	var me struct {
		ID         string  `json:"id"`
		Username   string  `json:"username"`
		GlobalName *string `json:"global_name"`
	}
	if err := o.send(ctx, http.MethodGet, "/users/@me", nil, "", answer.AccessToken, &me); err != nil {
		return Identity{}, err
	}
	if me.ID == "" {
		return Identity{}, fmt.Errorf("%w: no account in the answer", ErrExchange)
	}
	name := me.Username
	if me.GlobalName != nil && *me.GlobalName != "" {
		name = *me.GlobalName
	}
	return Identity{ID: me.ID, DisplayName: name}, nil
}
