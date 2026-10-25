package notify

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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
	if err := NewHTTPPoster().Post(context.Background(), srv.URL, embed); err != nil {
		t.Fatal(err)
	}
	mentions := body["allowed_mentions"].(map[string]any)["parse"].([]any)
	embeds := body["embeds"].([]any)
	if len(mentions) != 0 || len(embeds) != 1 || embeds[0].(map[string]any)["title"] != "t" {
		t.Fatalf("payload = %v", body)
	}
}

func TestHTTPPosterErrorIsTypedAndSanitised(t *testing.T) {
	leak := "bad https://discord.com/api/webhooks/1/SECRET and /api/webhooks/2/SECRET2 " + strings.Repeat("x", 400)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(leak))
	}))
	defer srv.Close()
	err := NewHTTPPoster().Post(context.Background(), srv.URL, Embed{})
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

func TestHTTPPosterTransportErrorHoldsNoURL(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	url := srv.URL + "/api/webhooks/1/SECRET"
	srv.Close()
	err := NewHTTPPoster().Post(context.Background(), url, Embed{})
	if err == nil || strings.Contains(err.Error(), "SECRET") || strings.Contains(err.Error(), "127.0.0.1") {
		t.Fatalf("err = %v", err)
	}
}

func TestHTTPPosterLookupReadsTheChannelID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"id":"1","channel_id":"999"}`))
	}))
	defer srv.Close()
	id, err := NewHTTPPoster().Lookup(context.Background(), srv.URL)
	if err != nil || id != "999" {
		t.Fatalf("Lookup = %q, %v", id, err)
	}
	bad := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(404) }))
	defer bad.Close()
	if _, err := NewHTTPPoster().Lookup(context.Background(), bad.URL); err == nil {
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
