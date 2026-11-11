package mail

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResendPostsExpectedJSON(t *testing.T) {
	var got map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer key" {
			t.Errorf("auth = %q", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type = %q", r.Header.Get("Content-Type"))
		}
		_ = json.NewDecoder(r.Body).Decode(&got)
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"id":"1"}`))
	}))
	defer srv.Close()
	sender := Sender{From: "Forever Sixty <help@foreversixty.gg>", ReplyTo: "help@foreversixty.gg"}
	m := NewResend("key", sender, srv.Client())
	m.endpoint = srv.URL
	err := m.Send(context.Background(), Message{To: "a@b.c", Subject: "Confirm", Text: "hi", HTML: "<p>hi</p>"})
	if err != nil {
		t.Fatal(err)
	}
	if got["subject"] != "Confirm" || got["text"] != "hi" || got["html"] != "<p>hi</p>" {
		t.Fatalf("payload = %v", got)
	}
	if got["from"] != sender.From || got["reply_to"] != sender.ReplyTo {
		t.Errorf("from = %q, reply_to = %q", got["from"], got["reply_to"])
	}
	to, ok := got["to"].([]any)
	if !ok || len(to) != 1 || to[0] != "a@b.c" {
		t.Errorf("to = %v", got["to"])
	}
}

func TestResendLeavesOutAnEmptyReplyToAndHTML(t *testing.T) {
	var got map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewDecoder(r.Body).Decode(&got)
	}))
	defer srv.Close()
	m := NewResend("key", Sender{From: "sender@example.com"}, srv.Client())
	m.endpoint = srv.URL
	if err := m.Send(context.Background(), Message{To: "a@b.c", Subject: "s", Text: "t"}); err != nil {
		t.Fatal(err)
	}
	if _, ok := got["reply_to"]; ok {
		t.Errorf("reply_to sent empty: %v", got)
	}
	if _, ok := got["html"]; ok {
		t.Errorf("html sent empty: %v", got)
	}
}

func TestResendNon2xxIsError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer srv.Close()
	m := NewResend("key", Sender{From: "sender@example.com"}, srv.Client())
	m.endpoint = srv.URL
	err := m.Send(context.Background(), Message{To: "user@example.com", Subject: "Test", Text: "body"})
	if err == nil {
		t.Fatal("expected error for 500 status")
	}
	if !contains(err.Error(), "status 500") {
		t.Errorf("error = %v", err)
	}
}

func TestResendTransportErrorIsWrapped(t *testing.T) {
	// A closed local server refuses the connection; an unresolvable host name is not hermetic
	// (a proxy can answer it with a 502 instead of failing the dial).
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	closedURL := srv.URL
	srv.Close()
	m := NewResend("key", Sender{From: "sender@example.com"}, http.DefaultClient)
	m.endpoint = closedURL
	err := m.Send(context.Background(), Message{To: "user@example.com", Subject: "Test", Text: "body"})
	if err == nil {
		t.Fatal("expected transport error")
	}
	if !contains(err.Error(), "mail: send") {
		t.Errorf("error = %v", err)
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
