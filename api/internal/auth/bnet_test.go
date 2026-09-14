package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/oauth2"
)

func TestNewBattleNetUsesTheProductionEndpoints(t *testing.T) {
	b := NewBattleNet("id", "secret", "https://api.foreversixty.gg/v1/auth/battlenet/callback")
	if b.Config.Endpoint.AuthURL != BnetAuthURL || b.Config.Endpoint.TokenURL != BnetTokenURL {
		t.Fatalf("endpoints = %+v", b.Config.Endpoint)
	}
	if b.UserInfoURL != BnetUserInfoURL {
		t.Fatalf("userinfo = %q", b.UserInfoURL)
	}
	if !strings.Contains(b.AuthURL("state"), "state=state") {
		t.Fatalf("auth url = %q", b.AuthURL("state"))
	}
}

func TestSafeNextKeepsRedirectsOnSite(t *testing.T) {
	for in, want := range map[string]string{
		"/logs": "/logs", "": "/logs", "//evil.example": "/logs",
		"https://evil.example": "/logs", "/account": "/account",
	} {
		if got := safeNext(in); got != want {
			t.Errorf("safeNext(%q) = %q, want %q", in, got, want)
		}
	}
}

// provider is a Battle.net stand-in: the token endpoint and the
// userinfo endpoint, which is all Identify touches.
func provider(t *testing.T, userinfo string, status int) *BattleNet {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /token", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Error(err)
		}
		if r.Form.Get("code") != "the-code" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"at","token_type":"bearer","expires_in":3600}`))
	})
	mux.HandleFunc("GET /oauth/userinfo", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer at" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(userinfo))
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return &BattleNet{
		Config: &oauth2.Config{
			ClientID: "id", ClientSecret: "secret",
			RedirectURL: "https://api.foreversixty.gg/v1/auth/battlenet/callback",
			Scopes:      []string{"openid", "wow.profile"},
			Endpoint:    oauth2.Endpoint{AuthURL: srv.URL + "/authorize", TokenURL: srv.URL + "/token"},
		},
		UserInfoURL: srv.URL + "/oauth/userinfo",
		HTTP:        srv.Client(),
	}
}

func TestIdentifyExchangesTheCodeAndReadsTheAccount(t *testing.T) {
	b := provider(t, `{"sub":"12345","battletag":"Baelgrim#1234"}`, http.StatusOK)
	u, err := b.Identify(context.Background(), "the-code")
	if err != nil {
		t.Fatal(err)
	}
	if u.Sub != "12345" || u.Battletag != "Baelgrim#1234" {
		t.Fatalf("user = %+v", u)
	}
}

func TestIdentifyRefusesWhatItCannotTrust(t *testing.T) {
	for name, b := range map[string]*BattleNet{
		"no subject":   provider(t, `{"battletag":"Baelgrim#1234"}`, http.StatusOK),
		"not json":     provider(t, `<html>`, http.StatusOK),
		"upstream 500": provider(t, `{}`, http.StatusInternalServerError),
	} {
		if _, err := b.Identify(context.Background(), "the-code"); err == nil {
			t.Errorf("%s should have been refused", name)
		}
	}
	b := provider(t, `{"sub":"1"}`, http.StatusOK)
	if _, err := b.Identify(context.Background(), "the-wrong-code"); err == nil {
		t.Error("a code the provider rejects must surface")
	}
}

func TestAuthURLCarriesTheScopeAndState(t *testing.T) {
	b := provider(t, `{"sub":"1"}`, http.StatusOK)
	got := b.AuthURL("the-state")
	for _, want := range []string{"state=the-state", "scope=openid+wow.profile", "client_id=id"} {
		if !strings.Contains(got, want) {
			t.Errorf("auth url = %s, want %s", got, want)
		}
	}
}

func TestEncodeAndDecodeState(t *testing.T) {
	state, next := decodeState(encodeState("abc", "/account"))
	if state != "abc" || next != "/account" {
		t.Fatalf("state = %q, next = %q", state, next)
	}
}
