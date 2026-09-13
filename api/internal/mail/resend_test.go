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
	m := NewResend("key", "Forever Sixty <hello@foreversixty.gg>", srv.Client())
	m.endpoint = srv.URL
	err := m.Send(context.Background(), Message{To: "a@b.c", Subject: "Confirm", Text: "hi"})
	if err != nil {
		t.Fatal(err)
	}
	if got["subject"] != "Confirm" || got["text"] != "hi" {
		t.Fatalf("payload = %v", got)
	}
	if got["from"] != "Forever Sixty <hello@foreversixty.gg>" {
		t.Errorf("from = %q", got["from"])
	}
	to, ok := got["to"].([]any)
	if !ok || len(to) != 1 || to[0] != "a@b.c" {
		t.Errorf("to = %v", got["to"])
	}
}

func TestResendNon2xxIsError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer srv.Close()
	m := NewResend("key", "sender@example.com", srv.Client())
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
	m := NewResend("key", "sender@example.com", http.DefaultClient)
	m.endpoint = "http://invalid.example.test:1"
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
