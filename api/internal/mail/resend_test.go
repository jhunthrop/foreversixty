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
}
