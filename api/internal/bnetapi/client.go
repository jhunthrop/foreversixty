// api/internal/bnetapi/client.go

// Package bnetapi is a small client for Blizzard's game-data and profile
// APIs — separate from auth.BattleNet, which is the OAuth login flow and
// is unchanged by this package. See
// docs/superpowers/specs/2026-09-22-battlenet-character-import-design.md
// §3 for the endpoints and response shapes this client reads.
//
// A client-credentials (app) token is held in memory only and is never
// persisted or logged, the same rule that governs a user's OAuth access
// token (spec §1 rule 1).
package bnetapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	// DefaultTokenURL is Blizzard's client-credentials token endpoint.
	DefaultTokenURL = "https://oauth.battle.net/token"
	// httpTimeout bounds every outbound call this client makes, so a
	// stalled Blizzard endpoint cannot pin a caller's goroutine forever.
	httpTimeout = 10 * time.Second
	// tokenExpiryMargin is how far ahead of a client-credentials token's
	// real expiry this client treats it as already stale, so a call
	// straddling the boundary never sends a token Blizzard is about to
	// reject.
	tokenExpiryMargin = 60 * time.Second
	// defaultRetryBackoff is how long a single retry waits after a 5xx
	// status or a network error, before trying once more (spec §3: never
	// on a 4xx). Tests override Client.RetryBackoff to keep runs fast.
	defaultRetryBackoff = 500 * time.Millisecond
	// realmCacheTTL is how long Realms caches one region's realm list in
	// memory (spec §3: "cached 24 h in memory").
	realmCacheTTL = 24 * time.Hour
	// maxResponseBytes bounds every response body this client reads.
	maxResponseBytes = 1 << 20
)

// Config builds a Client via New.
type Config struct {
	// HTTP is the client used for every outbound call. Nil gets one
	// bounded by httpTimeout.
	HTTP *http.Client
	// TokenURL is Blizzard's client-credentials endpoint. Empty gets
	// DefaultTokenURL.
	TokenURL string
	// APIHost maps a region to its regional API host. Nil gets
	// DefaultAPIHost.
	APIHost func(region string) string
	// ClientID and ClientSecret are the app's Battle.net API client
	// credentials, used only for the client-credentials token exchange.
	ClientID, ClientSecret string
	// Game is the profile-game segment of a namespace (e.g.
	// "classic1x") — BNET_PROFILE_GAME.
	Game string
	// Regions is which regional hosts to try for an account — BNET_REGIONS.
	Regions []string
	Log     *slog.Logger
}

// Client is a small client for Blizzard's game-data and profile APIs.
type Client struct {
	HTTP                   *http.Client
	TokenURL               string
	APIHost                func(region string) string
	ClientID, ClientSecret string
	Game                   string
	Regions                []string
	Log                    *slog.Logger
	// RetryBackoff overrides defaultRetryBackoff; tests set it small.
	RetryBackoff time.Duration

	now func() time.Time

	mu          sync.Mutex
	token       string
	tokenExpiry time.Time

	realmMu    sync.Mutex
	realmCache map[string]realmCacheEntry
}

type realmCacheEntry struct {
	realms   []Realm
	cachedAt time.Time
}

// DefaultAPIHost is the production Blizzard regional host mapping.
func DefaultAPIHost(region string) string {
	return "https://" + strings.ToLower(region) + ".api.blizzard.com"
}

// New builds a Client from cfg, filling in every production default a
// test fixture would want to override.
func New(cfg Config) *Client {
	c := &Client{
		HTTP: cfg.HTTP, TokenURL: cfg.TokenURL, APIHost: cfg.APIHost,
		ClientID: cfg.ClientID, ClientSecret: cfg.ClientSecret,
		Game: cfg.Game, Regions: cfg.Regions, Log: cfg.Log,
		now: time.Now, realmCache: map[string]realmCacheEntry{},
	}
	if c.HTTP == nil {
		c.HTTP = &http.Client{Timeout: httpTimeout}
	}
	if c.TokenURL == "" {
		c.TokenURL = DefaultTokenURL
	}
	if c.APIHost == nil {
		c.APIHost = DefaultAPIHost
	}
	return c
}

func (c *Client) logger() *slog.Logger {
	if c.Log != nil {
		return c.Log
	}
	return slog.Default()
}

func (c *Client) clock() time.Time {
	if c.now != nil {
		return c.now()
	}
	return time.Now()
}

func (c *Client) retryBackoff() time.Duration {
	if c.RetryBackoff > 0 {
		return c.RetryBackoff
	}
	return defaultRetryBackoff
}

// ProfileNamespace is "profile-<game>-<region>", e.g. "profile-classic1x-us".
func (c *Client) ProfileNamespace(region string) string {
	return "profile-" + c.Game + "-" + strings.ToLower(region)
}

// DynamicNamespace is "dynamic-<game>-<region>", e.g. "dynamic-classic1x-us".
func (c *Client) DynamicNamespace(region string) string {
	return "dynamic-" + c.Game + "-" + strings.ToLower(region)
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// AppToken returns a cached client-credentials token, refreshing it once
// it is within tokenExpiryMargin of expiry. Held in memory only.
func (c *Client) AppToken(ctx context.Context) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.token != "" && c.clock().Before(c.tokenExpiry.Add(-tokenExpiryMargin)) {
		return c.token, nil
	}
	form := url.Values{"grant_type": {"client_credentials"}}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.TokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("bnetapi: token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.ClientID, c.ClientSecret)
	res, err := c.do(req)
	if err != nil {
		return "", fmt.Errorf("bnetapi: token: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(io.LimitReader(res.Body, maxResponseBytes))
	if err != nil {
		return "", fmt.Errorf("bnetapi: token: read: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bnetapi: token: status %d", res.StatusCode)
	}
	var tr tokenResponse
	if err := json.Unmarshal(body, &tr); err != nil {
		return "", fmt.Errorf("bnetapi: token: decode: %w", err)
	}
	if tr.AccessToken == "" {
		return "", fmt.Errorf("bnetapi: token response carried no access_token")
	}
	c.token = tr.AccessToken
	c.tokenExpiry = c.clock().Add(time.Duration(tr.ExpiresIn) * time.Second)
	return c.token, nil
}

// do sends req, retrying exactly once after retryBackoff on a 5xx status
// or a network error — never on a 4xx (spec §3). A request whose body
// was already read on the first attempt is replayed via req.GetBody,
// which net/http populates automatically for a strings.Reader/
// bytes.Reader/bytes.Buffer body (every body this package sends).
func (c *Client) do(req *http.Request) (*http.Response, error) {
	res, err := c.HTTP.Do(req)
	if err == nil && res.StatusCode < 500 {
		return res, nil
	}
	if res != nil {
		res.Body.Close()
	}
	select {
	case <-req.Context().Done():
		if err == nil {
			err = req.Context().Err()
		}
		return nil, err
	case <-time.After(c.retryBackoff()):
	}
	retry := req
	if req.GetBody != nil {
		body, gerr := req.GetBody()
		if gerr != nil {
			return nil, fmt.Errorf("bnetapi: retry: %w", gerr)
		}
		retry = req.Clone(req.Context())
		retry.Body = body
	}
	return c.HTTP.Do(retry)
}

// getJSON performs an app-token-authenticated GET and decodes the JSON
// body into out.
func (c *Client) getJSON(ctx context.Context, op, rawURL string, out any) error {
	token, err := c.AppToken(ctx)
	if err != nil {
		return fmt.Errorf("bnetapi: %s: %w", op, err)
	}
	return c.getJSONWithToken(ctx, op, rawURL, token, out)
}

// getJSONWithToken performs a bearer-authenticated GET with an arbitrary
// token (the app token, or a user's own OAuth access token) and decodes
// the JSON body into out.
func (c *Client) getJSONWithToken(ctx context.Context, op, rawURL, token string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, withLocale(rawURL), nil)
	if err != nil {
		return fmt.Errorf("bnetapi: %s: request: %w", op, err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := c.do(req)
	if err != nil {
		return fmt.Errorf("bnetapi: %s: %w", op, err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(io.LimitReader(res.Body, maxResponseBytes))
	if err != nil {
		return fmt.Errorf("bnetapi: %s: read: %w", op, err)
	}
	if res.StatusCode != http.StatusOK {
		return statusError(op, res.StatusCode)
	}
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("bnetapi: %s: decode: %w", op, err)
	}
	return nil
}

// classSlug lowercases a Blizzard class/name field into the site's own
// slug vocabulary — every WoW class name is a single word, so a bare
// lowercase is exact (matches the FS1 export's classSlug field, see
// web/src/lib/planner/fs1.ts).
func classSlug(name string) string { return strings.ToLower(strings.TrimSpace(name)) }

// apiLocale pins every game-data and profile response to one language.
// Without a locale Blizzard answers every "name" field as an object keyed
// by locale ({"en_US": "Whitemane", ...}), which the narrow response
// structs decode as a string and reject; seen on the first production
// import, 2026-09-22.
const apiLocale = "en_US"

// withLocale appends locale=en_US to a namespaced Blizzard API URL. A
// URL with no namespace (the token endpoint) is returned unchanged.
func withLocale(rawURL string) string {
	if !strings.Contains(rawURL, "namespace=") || strings.Contains(rawURL, "locale=") {
		return rawURL
	}
	return rawURL + "&locale=" + apiLocale
}
