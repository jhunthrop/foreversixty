package subscribe

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/PLACEHOLDER/forever/api/internal/mail"
)

func discardLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

func newTestHandler() http.Handler {
	s := &Service{Store: &memStore{rows: map[string]memRow{"known@example.com": {token: "tok", unsubscribeToken: "utok"}}}, Mailer: &mail.Fake{}, PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg"}
	mux := http.NewServeMux()
	Mount(mux, s, discardLogger(), 0)
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
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/v1/subscribe/unsubscribe?token=utok", nil))
	if rec.Code != 302 || rec.Header().Get("Location") != "https://foreversixty.gg/unsubscribed" {
		t.Fatalf("code=%d loc=%s", rec.Code, rec.Header().Get("Location"))
	}
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/v1/subscribe/unsubscribe?token=bad", nil))
	if rec.Header().Get("Location") != "https://foreversixty.gg/subscribe-invalid" {
		t.Fatalf("loc=%s", rec.Header().Get("Location"))
	}
}

func TestPostSubscribeReturns202EvenWhenMailFailsAndLogsAsync(t *testing.T) {
	var buf bytes.Buffer
	log := slog.New(slog.NewTextHandler(&buf, nil))
	s := &Service{Store: &memStore{rows: map[string]memRow{}}, Mailer: &mail.Fake{Err: errors.New("boom")}, PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg", Logger: log}
	mux := http.NewServeMux()
	Mount(mux, s, log, 0)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("POST", "/v1/subscribe", strings.NewReader(`{"email":"new@example.com"}`)))
	if rec.Code != 202 {
		t.Fatalf("code=%d body=%s", rec.Code, rec.Body.String())
	}
	s.Wait()
	logOutput := buf.String()
	if !strings.Contains(logOutput, "op=subscribe") || !strings.Contains(logOutput, "stage=mail") {
		t.Fatalf("expected async log with op=subscribe stage=mail, got %q", logOutput)
	}
}

func TestPostSubscribe500OnStoreErrorLogsIt(t *testing.T) {
	var buf bytes.Buffer
	log := slog.New(slog.NewTextHandler(&buf, nil))
	s := &Service{Store: &erroringStore{err: errors.New("db down")}, Mailer: &mail.Fake{}, PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg"}
	mux := http.NewServeMux()
	Mount(mux, s, log, 0)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("POST", "/v1/subscribe", strings.NewReader(`{"email":"new@example.com"}`)))
	if rec.Code != 500 {
		t.Fatalf("code=%d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(buf.String(), "op=subscribe") {
		t.Fatalf("expected log with op=subscribe, got %q", buf.String())
	}
}

type erroringStore struct{ err error }

func (e *erroringStore) Upsert(_ context.Context, _, _, _, _ string) (UpsertResult, error) {
	return UpsertResult{}, e.err
}
func (e *erroringStore) ClaimConfirmationSend(_ context.Context, _ string, _ time.Duration) (bool, error) {
	return false, e.err
}
func (e *erroringStore) Confirm(_ context.Context, _ string) (bool, error)     { return false, e.err }
func (e *erroringStore) Unsubscribe(_ context.Context, _ string) (bool, error) { return false, e.err }

func TestConfirmStoreErrorRedirectsInvalidAndLogs(t *testing.T) {
	var buf bytes.Buffer
	log := slog.New(slog.NewTextHandler(&buf, nil))
	s := &Service{Store: &erroringStore{err: errors.New("db down")}, Mailer: &mail.Fake{}, PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg"}
	mux := http.NewServeMux()
	Mount(mux, s, log, 0)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("GET", "/v1/subscribe/confirm?token=tok", nil))
	if rec.Code != 302 || rec.Header().Get("Location") != "https://foreversixty.gg/subscribe-invalid" {
		t.Fatalf("code=%d loc=%s", rec.Code, rec.Header().Get("Location"))
	}
	if !strings.Contains(buf.String(), "op=confirm") {
		t.Fatalf("expected log with op=confirm, got %q", buf.String())
	}
}

func TestUnsubscribeStoreErrorRedirectsInvalidAndLogs(t *testing.T) {
	var buf bytes.Buffer
	log := slog.New(slog.NewTextHandler(&buf, nil))
	s := &Service{Store: &erroringStore{err: errors.New("db down")}, Mailer: &mail.Fake{}, PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg"}
	mux := http.NewServeMux()
	Mount(mux, s, log, 0)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("GET", "/v1/subscribe/unsubscribe?token=tok", nil))
	if rec.Code != 302 || rec.Header().Get("Location") != "https://foreversixty.gg/subscribe-invalid" {
		t.Fatalf("code=%d loc=%s", rec.Code, rec.Header().Get("Location"))
	}
	if !strings.Contains(buf.String(), "op=unsubscribe") {
		t.Fatalf("expected log with op=unsubscribe, got %q", buf.String())
	}
}

var _ = context.Background
