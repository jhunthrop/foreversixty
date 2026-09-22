package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

// The Battle.net endpoints. They are fields on BattleNet rather than
// constants used directly so the tests can point the whole flow at an
// httptest server.
const (
	BnetAuthURL     = "https://oauth.battle.net/authorize"
	BnetTokenURL    = "https://oauth.battle.net/token"
	BnetUserInfoURL = "https://oauth.battle.net/oauth/userinfo"
)

// bnetStateCookie carries the OAuth state and the page to return to.
const bnetStateCookie = "fs_bnet_state"

// bnetStateTTL is how long a sign-in may sit half-finished.
const bnetStateTTL = 10 * time.Minute

// bnetHTTPTimeout bounds both outbound legs of Identify. Without it,
// http.DefaultClient has no deadline and a Battle.net endpoint that
// accepts a connection and then stalls would pin a handler goroutine
// indefinitely.
const bnetHTTPTimeout = 10 * time.Second

// BnetUser is the part of Battle.net's userinfo the API keeps: the
// subject, which never changes, and the battletag, which can.
type BnetUser struct {
	Sub       string `json:"sub"`
	Battletag string `json:"battletag"`
	// AccessToken is the OAuth access token from the code exchange,
	// carried alongside the identity so bnetCallback can hand it to the
	// Battle.net importer. It is used inside that one request and then
	// dropped — it is never persisted, never logged, and never
	// marshalled (no json tag), per spec §1 rule 1.
	AccessToken string `json:"-"`
}

// BattleNet is the Battle.net half of sign-in.
type BattleNet struct {
	Config      *oauth2.Config
	UserInfoURL string
	// HTTP is the client used for the token exchange and userinfo. Nil
	// means http.DefaultClient.
	HTTP *http.Client
}

// NewBattleNet builds the production configuration, with an HTTP client
// bounded by bnetHTTPTimeout so a stalled Battle.net endpoint cannot pin
// a handler goroutine forever.
func NewBattleNet(clientID, clientSecret, redirectURL string) *BattleNet {
	return &BattleNet{
		Config: &oauth2.Config{
			ClientID: clientID, ClientSecret: clientSecret, RedirectURL: redirectURL,
			Scopes:   []string{"openid", "wow.profile"},
			Endpoint: oauth2.Endpoint{AuthURL: BnetAuthURL, TokenURL: BnetTokenURL},
		},
		UserInfoURL: BnetUserInfoURL,
		HTTP:        &http.Client{Timeout: bnetHTTPTimeout},
	}
}

// AuthURL is where the browser is sent to approve the sign-in.
func (b *BattleNet) AuthURL(state string) string {
	return b.Config.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

// Identify exchanges an authorization code and reads the account behind
// it.
func (b *BattleNet) Identify(ctx context.Context, code string) (BnetUser, error) {
	if b.HTTP != nil {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, b.HTTP)
	}
	tok, err := b.Config.Exchange(ctx, code)
	if err != nil {
		return BnetUser{}, exchangeError(err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, b.UserInfoURL, nil)
	if err != nil {
		return BnetUser{}, fmt.Errorf("auth: battle.net userinfo: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+tok.AccessToken)
	client := b.HTTP
	if client == nil {
		client = http.DefaultClient
	}
	res, err := client.Do(req)
	if err != nil {
		return BnetUser{}, fmt.Errorf("auth: battle.net userinfo: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(io.LimitReader(res.Body, 1<<16))
	if err != nil {
		return BnetUser{}, fmt.Errorf("auth: battle.net userinfo: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return BnetUser{}, fmt.Errorf("auth: battle.net userinfo: status %d", res.StatusCode)
	}
	var u BnetUser
	if err := json.Unmarshal(body, &u); err != nil {
		return BnetUser{}, fmt.Errorf("auth: battle.net userinfo: %w", err)
	}
	if u.Sub == "" {
		return BnetUser{}, fmt.Errorf("auth: battle.net userinfo carried no subject")
	}
	u.AccessToken = tok.AccessToken
	return u, nil
}

// exchangeError turns a failed code exchange into an error safe to log
// or return: oauth2.RetrieveError.Error() embeds the provider's raw
// response body when the provider returned no structured error code, so
// wrapping it with %w would let an upstream response we do not control
// reach a log line or the API's error envelope. The status code (when
// there is one) is kept; the body is not.
func exchangeError(err error) error {
	var rerr *oauth2.RetrieveError
	if errors.As(err, &rerr) {
		if rerr.Response != nil {
			return fmt.Errorf("auth: battle.net exchange failed (status %d)", rerr.Response.StatusCode)
		}
		return errors.New("auth: battle.net exchange failed")
	}
	return fmt.Errorf("auth: battle.net exchange: %w", err)
}

// safeNext keeps an open redirect out of the sign-in flow: the caller's
// next must be a same-site path, with no scheme and no host. It is
// checked structurally rather than by prefix, because a browser treats
// a backslash as a path separator in the authority position of a
// special-scheme URL: a denylist of "//" and "https://" alone still
// lets "/\evil.example" resolve to https://evil.example. The result is
// also validated against RFC 6265's cookie-octet set, since safeNext's
// caller packs it into the state cookie and a stray ';', '"' or '\'
// would make net/http's cookie sanitiser drop bytes and log a warning
// on every request.
func safeNext(next string) string {
	const fallback = "/logs"
	u, err := url.Parse(next)
	if err != nil || u.Scheme != "" || u.Host != "" || u.Opaque != "" {
		return fallback
	}
	if !strings.HasPrefix(u.Path, "/") || strings.HasPrefix(u.Path, "//") || strings.ContainsRune(u.Path, '\\') {
		return fallback
	}
	safe := u.EscapedPath()
	if u.RawQuery != "" {
		safe += "?" + u.RawQuery
	}
	if !cookieSafe(safe) {
		return fallback
	}
	return safe
}

// cookieSafe reports whether every byte of s is a valid cookie-octet
// (RFC 6265 §4.1.1): printable US-ASCII, excluding space, DQUOTE,
// comma, semicolon and backslash.
func cookieSafe(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < 0x21 || c > 0x7e || c == '"' || c == ',' || c == ';' || c == '\\' {
			return false
		}
	}
	return true
}

// encodeState packs the CSRF state and the return path into one cookie
// value. The state is random, so the separator can never appear in it.
func encodeState(state, next string) string { return state + " " + next }

func decodeState(v string) (state, next string) {
	state, next, _ = strings.Cut(v, " ")
	return state, next
}
