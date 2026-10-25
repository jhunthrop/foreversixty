package notify

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

func TestHTTPPosterSendsEmbedWithoutMentions(t *testing.T) {
	var body map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewDecoder(r.Body).Decode(&body)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()
	embed := Embed{Title: "t", Fields: []EmbedField{{Name: "n", Value: "v", Inline: true}}}
	if err := NewHTTPPoster("").Post(context.Background(), srv.URL, embed); err != nil {
		t.Fatal(err)
	}
	mentions := body["allowed_mentions"].(map[string]any)["parse"].([]any)
	embeds := body["embeds"].([]any)
	if len(mentions) != 0 || len(embeds) != 1 || embeds[0].(map[string]any)["title"] != "t" {
		t.Fatalf("payload = %v", body)
	}
}

func TestHTTPPosterErrorIsTypedAndSanitised(t *testing.T) {
	leak := "bad https://discord.com/api/webhooks/1/SECRET and /api/webhooks/2/SECRET2 and api/v10/webhooks/3/SECRET3 " + strings.Repeat("x", 400)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(leak))
	}))
	defer srv.Close()
	err := NewHTTPPoster("").Post(context.Background(), srv.URL, Embed{})
	var de *DiscordError
	if !errors.As(err, &de) || de.Status != 429 || !de.RateLimited() {
		t.Fatalf("err = %v", err)
	}
	if strings.Contains(de.Body, "SECRET") || strings.Contains(de.Body, "webhooks") || len([]rune(de.Body)) > 200 {
		t.Fatalf("body not sanitised: %q", de.Body)
	}
	if strings.Contains(err.Error(), "SECRET") {
		t.Fatalf("Error() leaks: %q", err.Error())
	}
}

func TestHTTPPosterDoesNotFollowRedirects(t *testing.T) {
	var hit atomic.Bool
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hit.Store(true)
	}))
	defer target.Close()
	redirector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL, http.StatusTemporaryRedirect)
	}))
	defer redirector.Close()
	err := NewHTTPPoster("").Post(context.Background(), redirector.URL, Embed{})
	var de *DiscordError
	if !errors.As(err, &de) || de.Status != http.StatusTemporaryRedirect {
		t.Fatalf("err = %v; want *DiscordError 307", err)
	}
	if hit.Load() {
		t.Fatal("the redirect target was contacted")
	}
}

func TestHTTPPosterTransportErrorHoldsNoURL(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	url := srv.URL + "/api/webhooks/1/SECRET"
	srv.Close()
	err := NewHTTPPoster("").Post(context.Background(), url, Embed{})
	if err == nil || strings.Contains(err.Error(), "SECRET") || strings.Contains(err.Error(), "127.0.0.1") {
		t.Fatalf("err = %v", err)
	}
}

func TestHTTPPosterLookupReadsTheChannelID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"id":"1","channel_id":"999"}`))
	}))
	defer srv.Close()
	id, err := NewHTTPPoster("").Lookup(context.Background(), srv.URL)
	if err != nil || id != "999" {
		t.Fatalf("Lookup = %q, %v", id, err)
	}
	bad := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(404) }))
	defer bad.Close()
	if _, err := NewHTTPPoster("").Lookup(context.Background(), bad.URL); err == nil {
		t.Fatal("a 404 lookup succeeded")
	}
}

func TestFakePosterRecordsAndScripts(t *testing.T) {
	f := &FakePoster{}
	_ = f.Post(context.Background(), "u", Embed{Title: "a"})
	f.PostErr = &DiscordError{Status: 500}
	if err := f.Post(context.Background(), "u", Embed{Title: "b"}); err == nil {
		t.Fatal("scripted error not returned")
	}
	if got := f.Posts(); len(got) != 2 || got[1].Embed.Title != "b" {
		t.Fatalf("posts = %+v", got)
	}
}

func postErrorFor(t *testing.T, status int, body string) *DiscordError {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	defer srv.Close()
	var de *DiscordError
	if err := NewHTTPPoster("").Post(context.Background(), srv.URL, Embed{}); !errors.As(err, &de) {
		t.Fatalf("err = %v; want *DiscordError", err)
	}
	return de
}

func TestDiscordErrorBodyIsItsMessage(t *testing.T) {
	cases := []struct {
		name   string
		status int
		body   string
		want   string
	}{
		{"json message", 404, `{"message": "Unknown Webhook", "code": 10015}`, "Unknown Webhook"},
		{"html from a proxy", 502, "<html><body>bad gateway page</body></html>", "Bad Gateway"},
		{"empty message", 500, `{"message": "", "code": 1}`, "Internal Server Error"},
		{"url in message", 400, `{"message": "bad https://discord.com/api/webhooks/1/SECRET here"}`, "bad here"},
	}
	for _, c := range cases {
		if got := postErrorFor(t, c.status, c.body).Body; got != c.want {
			t.Errorf("%s: Body = %q, want %q", c.name, got, c.want)
		}
	}
}

func TestTransportErrorsAreTheSentinels(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	url := srv.URL
	srv.Close()
	if err := NewHTTPPoster("").Post(context.Background(), url, Embed{}); !errors.Is(err, ErrUnreachable) {
		t.Fatalf("refused connection = %v; want ErrUnreachable", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := NewHTTPPoster("").Post(ctx, url, Embed{}); !errors.Is(err, ErrTimeout) {
		t.Fatalf("cancelled = %v; want ErrTimeout", err)
	}
}
