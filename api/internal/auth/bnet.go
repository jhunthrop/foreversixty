package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

// BnetUser is the part of Battle.net's userinfo the API keeps: the
// subject, which never changes, and the battletag, which can.
type BnetUser struct {
	Sub       string `json:"sub"`
	Battletag string `json:"battletag"`
}

// BattleNet is the Battle.net half of sign-in.
type BattleNet struct {
	Config      *oauth2.Config
	UserInfoURL string
	// HTTP is the client used for the token exchange and userinfo. Nil
	// means http.DefaultClient.
	HTTP *http.Client
}

// NewBattleNet builds the production configuration.
func NewBattleNet(clientID, clientSecret, redirectURL string) *BattleNet {
	return &BattleNet{
		Config: &oauth2.Config{
			ClientID: clientID, ClientSecret: clientSecret, RedirectURL: redirectURL,
			Scopes:   []string{"openid", "wow.profile"},
			Endpoint: oauth2.Endpoint{AuthURL: BnetAuthURL, TokenURL: BnetTokenURL},
		},
		UserInfoURL: BnetUserInfoURL,
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
		return BnetUser{}, fmt.Errorf("auth: battle.net exchange: %w", err)
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
	return u, nil
}

// safeNext keeps an open redirect out of the sign-in flow: the caller's
// next must be a path on our own site, so "//evil.example" and
// "https://evil.example" both fall back to the logs page.
func safeNext(next string) string {
	if !strings.HasPrefix(next, "/") || strings.HasPrefix(next, "//") {
		return "/logs"
	}
	return next
}

// encodeState packs the CSRF state and the return path into one cookie
// value. The state is random, so the separator can never appear in it.
func encodeState(state, next string) string { return state + " " + next }

func decodeState(v string) (state, next string) {
	state, next, _ = strings.Cut(v, " ")
	return state, next
}
