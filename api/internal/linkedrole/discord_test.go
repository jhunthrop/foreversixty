package linkedrole

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func newTestDiscord(t *testing.T, h http.HandlerFunc) *HTTPDiscord {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	d := NewHTTPDiscord("app1", "s3cret", "https://api.test/cb")
	d.BaseURL = srv.URL
	return d
}

func TestAuthorizeURLAsksForTheRoleConnectionScope(t *testing.T) {
	u, err := url.Parse(NewHTTPDiscord("app1", "x", "https://api.test/cb").AuthorizeURL("st"))
	if err != nil {
		t.Fatal(err)
	}
	q := u.Query()
	if q.Get("scope") != "identify role_connections.write" || q.Get("state") != "st" || q.Get("redirect_uri") != "https://api.test/cb" {
		t.Fatalf("query = %v", q)
	}
}

func TestRefreshPostsTheTokenAndReadsTheRotatedPair(t *testing.T) {
	d := newTestDiscord(t, func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		if r.URL.Path != "/oauth2/token" || r.PostForm.Get("grant_type") != "refresh_token" || r.PostForm.Get("refresh_token") != "r0" {
			t.Errorf("request = %s %v", r.URL.Path, r.PostForm)
		}
		_, _ = io.WriteString(w, `{"access_token":"a1","refresh_token":"r1"}`)
	})
	got, err := d.Refresh(context.Background(), "r0")
	if err != nil || got != (Tokens{Access: "a1", Refresh: "r1"}) {
		t.Fatalf("got %+v, %v", got, err)
	}
}

func TestInvalidGrantIsRevokedAndNoErrorCarriesASecret(t *testing.T) {
	d := newTestDiscord(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, `{"error":"invalid_grant"}`)
	})
	if _, err := d.Refresh(context.Background(), "r0"); !errors.Is(err, ErrRevoked) {
		t.Fatalf("err = %v, want ErrRevoked", err)
	}
	d = newTestDiscord(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = io.WriteString(w, `{"error":"invalid_client","detail":"s3cret"}`)
	})
	_, err := d.Refresh(context.Background(), "r0")
	if err == nil || errors.Is(err, ErrRevoked) || strings.Contains(err.Error(), "s3cret") || strings.Contains(err.Error(), "r0") {
		t.Fatalf("err = %v", err)
	}
}

func TestPushWritesTheRoleConnection(t *testing.T) {
	d := newTestDiscord(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.URL.Path != "/users/@me/applications/app1/role-connection" || r.Header.Get("Authorization") != "Bearer a1" {
			t.Errorf("request = %s %s %q", r.Method, r.URL.Path, r.Header.Get("Authorization"))
		}
		var body connectionBody
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body.PlatformName != "Forever Sixty" || body.PlatformUsername != "a#1" || body.Metadata["premium"] != "1" {
			t.Errorf("body = %+v", body)
		}
		_, _ = io.WriteString(w, `{}`)
	})
	if err := d.Push(context.Background(), "a1", Connection{Premium: true, Username: "a#1"}); err != nil {
		t.Fatal(err)
	}
}

func TestSetupRegistersTheMetadataAndTheVerificationURL(t *testing.T) {
	seen := map[string]string{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		if r.Header.Get("Authorization") != "Bot tok" {
			t.Errorf("auth = %q", r.Header.Get("Authorization"))
		}
		seen[r.Method+" "+r.URL.Path] = string(b)
		_, _ = io.WriteString(w, `{}`)
	}))
	t.Cleanup(srv.Close)
	app := NewApp("app1", "tok")
	app.BaseURL = srv.URL
	if err := app.Setup(context.Background(), "https://api.test/v1/discord/linked-role"); err != nil {
		t.Fatal(err)
	}
	meta := seen["PUT /applications/app1/role-connections/metadata"]
	if !strings.Contains(meta, `"key":"premium"`) || !strings.Contains(meta, `"type":7`) {
		t.Fatalf("metadata = %q", meta)
	}
	if got := seen["PATCH /applications/@me"]; got != `{"role_connections_verification_url":"https://api.test/v1/discord/linked-role"}` {
		t.Fatalf("patch = %q", got)
	}
}

func TestRevokeTokenPostsTheRefreshTokenWithTheClientCredentials(t *testing.T) {
	d := newTestDiscord(t, func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		f := r.PostForm
		if r.URL.Path != "/oauth2/token/revoke" || f.Get("token") != "r0" || f.Get("token_type_hint") != "refresh_token" ||
			f.Get("client_id") != "app1" || f.Get("client_secret") != "s3cret" {
			t.Errorf("request = %s %v", r.URL.Path, f)
		}
		_, _ = io.WriteString(w, `{}`)
	})
	if err := d.RevokeToken(context.Background(), "r0"); err != nil {
		t.Fatal(err)
	}
}
