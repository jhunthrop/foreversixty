// Package linkedrole is Discord Linked Roles for the Forever Sixty server: a member authorizes
// the role connection once, and the boolean `premium` Discord reads for the Premium role then
// follows the account's Premium (on entitlement changes and in the daily reconcile).
package linkedrole

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	authorizeURL   = "https://discord.com/oauth2/authorize"
	apiBase        = "https://discord.com/api/v10"
	oauthScope     = "identify role_connections.write"
	requestTimeout = 5 * time.Second
	maxAnswerBytes = 64 << 10

	// PlatformName is the connection's name on the member's Discord profile.
	PlatformName = "Forever Sixty"
	// MetadataKey is the one metadata field; the role's condition is `premium` true.
	MetadataKey = "premium"
	// metadataBooleanEqual is Discord's BOOLEAN_EQUAL metadata type.
	metadataBooleanEqual = 7
	metadataName         = "Forever Sixty Premium"
	metadataDescription  = "Holds Forever Sixty Premium"
)

// ErrDiscord is a failed Discord call; its text carries at most a status, never a token.
var ErrDiscord = errors.New("linkedrole: discord call failed")

// ErrRevoked is a refresh token Discord no longer honours: the member removed the app, or the
// token was rotated elsewhere.
var ErrRevoked = errors.New("linkedrole: discord refused the refresh token")

// Tokens is one token answer. Both values are credentials: never log or return them.
type Tokens struct{ Access, Refresh string }

// Connection is what the role connection says about the member.
type Connection struct {
	Premium  bool
	Username string
}

// Discord is the member-facing OAuth and role-connection API, behind an interface so the
// handler and the syncer test without the network.
type Discord interface {
	AuthorizeURL(state string) string
	Exchange(ctx context.Context, code string) (Tokens, error)
	Refresh(ctx context.Context, refreshToken string) (Tokens, error)
	UserID(ctx context.Context, accessToken string) (string, error)
	Push(ctx context.Context, accessToken string, c Connection) error
}

// HTTPDiscord talks to discord.com.
type HTTPDiscord struct {
	// BaseURL overrides the API base; tests point it at httptest.
	BaseURL                             string
	clientID, clientSecret, redirectURI string
	client                              *http.Client
}

// NewHTTPDiscord returns the production client. The application id is the client id.
func NewHTTPDiscord(clientID, clientSecret, redirectURI string) *HTTPDiscord {
	return &HTTPDiscord{BaseURL: apiBase, clientID: clientID, clientSecret: clientSecret, redirectURI: redirectURI, client: newClient()}
}

// newClient is the 5 s client; a redirect would carry a credential elsewhere, so the 3xx is
// returned instead.
func newClient() *http.Client {
	return &http.Client{
		Timeout:       requestTimeout,
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
	}
}

// AuthorizeURL is Discord's consent page for the role connection.
func (d *HTTPDiscord) AuthorizeURL(state string) string {
	q := url.Values{}
	q.Set("response_type", "code")
	q.Set("client_id", d.clientID)
	q.Set("scope", oauthScope)
	q.Set("redirect_uri", d.redirectURI)
	q.Set("state", state)
	return authorizeURL + "?" + q.Encode()
}

// Exchange trades the authorization code for tokens.
func (d *HTTPDiscord) Exchange(ctx context.Context, code string) (Tokens, error) {
	form := url.Values{"grant_type": {"authorization_code"}, "code": {code}, "redirect_uri": {d.redirectURI}}
	return d.token(ctx, form)
}

// Refresh trades a refresh token for a new pair; Discord rotates the refresh token. A token it
// refuses (invalid_grant) is ErrRevoked.
func (d *HTTPDiscord) Refresh(ctx context.Context, refreshToken string) (Tokens, error) {
	form := url.Values{"grant_type": {"refresh_token"}, "refresh_token": {refreshToken}}
	return d.token(ctx, form)
}

func (d *HTTPDiscord) token(ctx context.Context, form url.Values) (Tokens, error) {
	form.Set("client_id", d.clientID)
	form.Set("client_secret", d.clientSecret)
	var out struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		Error        string `json:"error"`
	}
	status, err := d.send(ctx, http.MethodPost, "/oauth2/token", "application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()), "", &out)
	if status == http.StatusBadRequest && out.Error == "invalid_grant" {
		return Tokens{}, ErrRevoked
	}
	if err != nil {
		return Tokens{}, err
	}
	if out.AccessToken == "" || out.RefreshToken == "" {
		return Tokens{}, fmt.Errorf("%w: no tokens in the answer", ErrDiscord)
	}
	return Tokens{Access: out.AccessToken, Refresh: out.RefreshToken}, nil
}

// UserID reads the member's Discord user id.
func (d *HTTPDiscord) UserID(ctx context.Context, accessToken string) (string, error) {
	var me struct {
		ID string `json:"id"`
	}
	if _, err := d.send(ctx, http.MethodGet, "/users/@me", "", nil, "Bearer "+accessToken, &me); err != nil {
		return "", err
	}
	if me.ID == "" {
		return "", fmt.Errorf("%w: no account in the answer", ErrDiscord)
	}
	return me.ID, nil
}

// connectionBody is the role connection's JSON; metadata values are strings ("1" or "0" for a
// boolean).
type connectionBody struct {
	PlatformName     string            `json:"platform_name"`
	PlatformUsername string            `json:"platform_username,omitempty"`
	Metadata         map[string]string `json:"metadata"`
}

// Push writes the member's role connection.
func (d *HTTPDiscord) Push(ctx context.Context, accessToken string, c Connection) error {
	premium := "0"
	if c.Premium {
		premium = "1"
	}
	body, err := json.Marshal(connectionBody{
		PlatformName: PlatformName, PlatformUsername: c.Username, Metadata: map[string]string{MetadataKey: premium},
	})
	if err != nil {
		return fmt.Errorf("%w: encode", ErrDiscord)
	}
	_, err = d.send(ctx, http.MethodPut, "/users/@me/applications/"+d.clientID+"/role-connection",
		"application/json", bytes.NewReader(body), "Bearer "+accessToken, nil)
	return err
}

// send makes one call and decodes a JSON answer into out (nil skips it). It returns the status
// even on a failure, so a caller can read Discord's error field. Errors carry no body or token.
func (d *HTTPDiscord) send(ctx context.Context, method, path, contentType string, body io.Reader, auth string, out any) (int, error) {
	return doJSON(ctx, d.client, method, d.BaseURL+path, contentType, body, auth, out)
}

func doJSON(ctx context.Context, client *http.Client, method, endpoint, contentType string, body io.Reader, auth string, out any) (int, error) {
	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return 0, fmt.Errorf("%w: request", ErrDiscord)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "DiscordBot (https://foreversixty.gg, 1)")
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if auth != "" {
		req.Header.Set("Authorization", auth)
	}
	res, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("%w: no answer", ErrDiscord)
	}
	defer res.Body.Close()
	limited := io.LimitReader(res.Body, maxAnswerBytes)
	if res.StatusCode < 200 || res.StatusCode > 299 {
		if out != nil {
			_ = json.NewDecoder(limited).Decode(out)
		}
		return res.StatusCode, fmt.Errorf("%w: status %d", ErrDiscord, res.StatusCode)
	}
	if out == nil {
		return res.StatusCode, nil
	}
	if err := json.NewDecoder(limited).Decode(out); err != nil {
		return res.StatusCode, fmt.Errorf("%w: unreadable answer", ErrDiscord)
	}
	return res.StatusCode, nil
}
