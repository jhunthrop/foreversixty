package notify

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func testOAuthClient(tokenURL string) *HTTPOAuthClient {
	c := NewHTTPOAuthClient("client-1", "shh-secret", "https://api.foreversixty.gg/v1/notifications/discord/callback")
	c.TokenURL = tokenURL
	return c
}

func TestAuthURLIsExactlyDiscordsAuthorizeURL(t *testing.T) {
	got := testOAuthClient("").AuthURL("st ate/1")
	want := "https://discord.com/oauth2/authorize?client_id=client-1&redirect_uri=https%3A%2F%2Fapi.foreversixty.gg%2Fv1%2Fnotifications%2Fdiscord%2Fcallback" +
		"&response_type=code&scope=webhook.incoming&state=st+ate%2F1"
	if got != want {
		t.Fatalf("AuthURL = %s", got)
	}
	u, _ := url.Parse(got)
	if q := u.Query(); q.Get("response_type") != "code" || q.Get("scope") != "webhook.incoming" || q.Get("state") != "st ate/1" {
		t.Fatalf("query = %v", q)
	}
}

func TestExchangePostsAFormAndDecodesOnlyTheWebhook(t *testing.T) {
	var form url.Values
	var contentType, method string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method, contentType = r.Method, r.Header.Get("Content-Type")
		_ = r.ParseForm()
		form = r.PostForm
		_, _ = w.Write([]byte(`{"access_token":"AT","refresh_token":"RT","webhook":{"id":"1","token":"tok","url":"https://discord.com/api/webhooks/1/tok","channel_id":"2","guild_id":"3","name":"n"}}`))
	}))
	defer srv.Close()
	got, err := testOAuthClient(srv.URL).Exchange(context.Background(), "the-code")
	if err != nil {
		t.Fatal(err)
	}
	if method != "POST" || contentType != "application/x-www-form-urlencoded" {
		t.Fatalf("%s %s", method, contentType)
	}
	for k, v := range map[string]string{
		"grant_type": "authorization_code", "code": "the-code", "client_id": "client-1", "client_secret": "shh-secret",
		"redirect_uri": "https://api.foreversixty.gg/v1/notifications/discord/callback",
	} {
		if form.Get(k) != v {
			t.Fatalf("form %s = %q", k, form.Get(k))
		}
	}
	want := OAuthWebhook{ID: "1", Token: "tok", URL: "https://discord.com/api/webhooks/1/tok", ChannelID: "2", GuildID: "3", Name: "n"}
	if got != want {
		t.Fatalf("webhook = %+v", got)
	}
}

func TestExchangeErrorsCarryTheStatusNeverTheBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, `{"error":"invalid_grant","echo":"the-code"}`, http.StatusBadRequest)
	}))
	defer srv.Close()
	_, err := testOAuthClient(srv.URL).Exchange(context.Background(), "the-code")
	if err == nil || !strings.Contains(err.Error(), "400") || strings.Contains(err.Error(), "the-code") || strings.Contains(err.Error(), "invalid_grant") {
		t.Fatalf("err = %v", err)
	}
}

func TestExchangeWithoutAWebhookIsAnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"access_token":"AT"}`))
	}))
	defer srv.Close()
	if _, err := testOAuthClient(srv.URL).Exchange(context.Background(), "c"); err == nil {
		t.Fatal("an answer with no webhook was accepted")
	}
}

func TestExchangeFollowsNoRedirectAndHasAFiveSecondTimeout(t *testing.T) {
	hit := false
	target := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { hit = true }))
	defer target.Close()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL, http.StatusTemporaryRedirect)
	}))
	defer srv.Close()
	client := testOAuthClient(srv.URL)
	if _, err := client.Exchange(context.Background(), "c"); err == nil || hit {
		t.Fatalf("redirect followed (hit=%v) or accepted (err=%v)", hit, err)
	}
	if client.HTTP.Timeout != 5*time.Second {
		t.Fatalf("timeout = %v", client.HTTP.Timeout)
	}
}
