package subscribe

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PLACEHOLDER/forever/api/internal/mail"
)

func newTestHandler() http.Handler {
	s := &Service{Store: &memStore{rows: map[string]string{"known@example.com": "tok"}}, Mailer: &mail.Fake{}, PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg"}
	mux := http.NewServeMux()
	Mount(mux, s)
	return mux
}

func TestPostSubscribeAlways202OnValidInput(t *testing.T) {
	h := newTestHandler()
	for _, email := range []string{"new@example.com", "known@example.com"} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest("POST", "/v1/subscribe", strings.NewReader(`{"email":"`+email+`"}`)))
		if rec.Code != 202 {
			t.Fatalf("%s: code = %d body=%s", email, rec.Code, rec.Body.String())
		}
	}
}

func TestPostSubscribe400OnInvalid(t *testing.T) {
	h := newTestHandler()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("POST", "/v1/subscribe", strings.NewReader(`{"email":"nope"}`)))
	if rec.Code != 400 || !strings.Contains(rec.Body.String(), `"email"`) {
		t.Fatalf("code=%d body=%s", rec.Code, rec.Body.String())
	}
}

func TestConfirmRedirects(t *testing.T) {
	h := newTestHandler()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/v1/subscribe/confirm?token=tok", nil))
	if rec.Code != 302 || rec.Header().Get("Location") != "https://foreversixty.gg/subscribed" {
		t.Fatalf("code=%d loc=%s", rec.Code, rec.Header().Get("Location"))
	}
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/v1/subscribe/confirm?token=bad", nil))
	if rec.Header().Get("Location") != "https://foreversixty.gg/subscribe-invalid" {
		t.Fatalf("loc=%s", rec.Header().Get("Location"))
	}
}

func TestUnsubscribeRedirects(t *testing.T) {
	h := newTestHandler()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/v1/subscribe/unsubscribe?token=tok", nil))
	if rec.Code != 302 || rec.Header().Get("Location") != "https://foreversixty.gg/unsubscribed" {
		t.Fatalf("code=%d loc=%s", rec.Code, rec.Header().Get("Location"))
	}
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/v1/subscribe/unsubscribe?token=bad", nil))
	if rec.Header().Get("Location") != "https://foreversixty.gg/subscribe-invalid" {
		t.Fatalf("loc=%s", rec.Header().Get("Location"))
	}
}

func TestPostSubscribe500OnMailError(t *testing.T) {
	s := &Service{Store: &memStore{rows: map[string]string{}}, Mailer: &mail.Fake{Err: errors.New("boom")}, PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg"}
	mux := http.NewServeMux()
	Mount(mux, s)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("POST", "/v1/subscribe", strings.NewReader(`{"email":"new@example.com"}`)))
	if rec.Code != 500 {
		t.Fatalf("code=%d body=%s", rec.Code, rec.Body.String())
	}
}

var _ = context.Background
