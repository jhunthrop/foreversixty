package httpx

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func okHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { WriteOK(w, r, 200, nil) })
}

func TestRequestIDIsSetAndEchoed(t *testing.T) {
	h := Chain(okHandler(), RequestID())
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	if rec.Header().Get("X-Request-Id") == "" {
		t.Fatal("missing X-Request-Id header")
	}
}

func TestRecoverTurnsPanicInto500(t *testing.T) {
	buf := &bytes.Buffer{}
	log := slog.New(slog.NewTextHandler(buf, nil))
	h := Chain(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { panic("boom") }), RequestID(), Recover(log))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	if rec.Code != 500 {
		t.Fatalf("code = %d", rec.Code)
	}
	logOutput := buf.String()
	if !strings.Contains(logOutput, "method=GET") {
		t.Fatalf("log missing method=GET: %s", logOutput)
	}
	if !strings.Contains(logOutput, "path=/") {
		t.Fatalf("log missing path=/: %s", logOutput)
	}
}

func TestRateLimitReturns429AfterBudget(t *testing.T) {
	h := Chain(okHandler(), RequestID(), RateLimit(2))
	var last int
	for i := 0; i < 3; i++ {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		req.RemoteAddr = "10.0.0.1:1234"
		h.ServeHTTP(rec, req)
		last = rec.Code
	}
	if last != 429 {
		t.Fatalf("third request code = %d, want 429", last)
	}
}

func TestCORSAllowsConfiguredOrigin(t *testing.T) {
	h := Chain(okHandler(), CORS("https://foreversixty.gg"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("OPTIONS", "/v1/subscribe", nil)
	req.Header.Set("Origin", "https://foreversixty.gg")
	h.ServeHTTP(rec, req)
	if rec.Code != 204 || rec.Header().Get("Access-Control-Allow-Origin") != "https://foreversixty.gg" {
		t.Fatalf("code=%d origin=%q", rec.Code, rec.Header().Get("Access-Control-Allow-Origin"))
	}
}
