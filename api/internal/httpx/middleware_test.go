package httpx

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
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
	h := Chain(okHandler(), RequestID(), RateLimit(2, 1))
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

func TestRateLimitTrustsOnlyConfiguredHopsFromRight(t *testing.T) {
	// hops=1: only the rightmost X-Forwarded-For entry is trusted. A client
	// spoofing the leftmost entry must not be able to dodge the bucket by
	// varying it.
	h := Chain(okHandler(), RequestID(), RateLimit(1, 1))

	req1 := httptest.NewRequest("GET", "/", nil)
	req1.RemoteAddr = "203.0.113.10:1"
	req1.Header.Set("X-Forwarded-For", "1.1.1.1, 9.9.9.9")
	rec1 := httptest.NewRecorder()
	h.ServeHTTP(rec1, req1)
	if rec1.Code != 200 {
		t.Fatalf("first request code = %d, want 200", rec1.Code)
	}

	req2 := httptest.NewRequest("GET", "/", nil)
	req2.RemoteAddr = "203.0.113.10:1"
	req2.Header.Set("X-Forwarded-For", "2.2.2.2, 9.9.9.9") // spoofed leftmost hop
	rec2 := httptest.NewRecorder()
	h.ServeHTTP(rec2, req2)
	if rec2.Code != 429 {
		t.Fatalf("second request (spoofed leftmost hop, same real client) code = %d, want 429 (shared bucket)", rec2.Code)
	}
}

func TestClientIPUsesRightmostHopWithHopsOne(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.RemoteAddr = "10.0.0.9:1"
	r.Header.Set("X-Forwarded-For", "1.2.3.4, 5.6.7.8")
	if ip := clientIP(r, 1); ip != "5.6.7.8" {
		t.Fatalf("clientIP = %q, want 5.6.7.8", ip)
	}
}

func TestClientIPJoinsMultipleXFFHeaderLines(t *testing.T) {
	// A client can send its own X-Forwarded-For header line ahead of the one
	// the trusted proxy appends; net/http keeps repeated header lines
	// separate (only visible via Header.Values, not Header.Get). clientIP
	// must join every line before counting hops from the right, so a
	// client-supplied line can't shift the trusted-hop count.
	r := httptest.NewRequest("GET", "/", nil)
	r.RemoteAddr = "10.0.0.9:1"
	r.Header.Add("X-Forwarded-For", "9.9.9.9, 8.8.8.8") // client-supplied line
	r.Header.Add("X-Forwarded-For", "1.2.3.4, 5.6.7.8") // proxy-supplied line
	if ip := clientIP(r, 1); ip != "5.6.7.8" {
		t.Fatalf("clientIP = %q, want 5.6.7.8 (last hop of the proxy-supplied line)", ip)
	}
}

func TestClientIPIgnoresXFFWhenHopsZero(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.RemoteAddr = "10.0.0.5:555"
	r.Header.Set("X-Forwarded-For", "1.2.3.4")
	if ip := clientIP(r, 0); ip != "10.0.0.5" {
		t.Fatalf("clientIP = %q, want 10.0.0.5 (XFF ignored)", ip)
	}
}

func TestClientIPFallsBackToRemoteAddrWhenListShorterThanHops(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.RemoteAddr = "10.0.0.7:1"
	r.Header.Set("X-Forwarded-For", "1.2.3.4") // one hop, but two trusted
	if ip := clientIP(r, 2); ip != "10.0.0.7" {
		t.Fatalf("clientIP = %q, want fallback to RemoteAddr 10.0.0.7", ip)
	}

	r2 := httptest.NewRequest("GET", "/", nil)
	r2.RemoteAddr = "10.0.0.8:1"
	// no X-Forwarded-For header at all
	if ip := clientIP(r2, 1); ip != "10.0.0.8" {
		t.Fatalf("clientIP = %q, want fallback to RemoteAddr 10.0.0.8", ip)
	}
}

func TestIPLimiterSweepEvictsIdleEntryAfterTenMinutes(t *testing.T) {
	l := newIPLimiter(10, time.Minute)
	now := time.Now()

	// Seed one entry that is already 11 minutes idle relative to `now`.
	l.entries["stale"] = &limiterEntry{
		lim:  nil,
		seen: now.Add(-11 * time.Minute),
	}

	// Drive the call counter to exactly the sweep boundary; every call in
	// between uses a distinct, fresh IP so it can't itself be the stale one.
	for i := 0; i < limiterSweepEvery; i++ {
		l.allow(fmt.Sprintf("fresh-%d", i), now)
	}

	l.mu.Lock()
	_, stillThere := l.entries["stale"]
	l.mu.Unlock()
	if stillThere {
		t.Fatal("stale entry (idle > 10m) should have been evicted by the periodic sweep")
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

func TestRateLimitPerHourAllowsTheBudgetThenReturns429(t *testing.T) {
	h := Chain(okHandler(), RequestID(), RateLimitPer(3, time.Hour, 1))
	for i := 0; i < 3; i++ {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/v1/builds", nil)
		req.RemoteAddr = "10.0.0.1:1234"
		h.ServeHTTP(rec, req)
		if rec.Code != 200 {
			t.Fatalf("request %d: code = %d, want 200", i+1, rec.Code)
		}
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/v1/builds", nil)
	req.RemoteAddr = "10.0.0.1:1234"
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusTooManyRequests {
		t.Fatalf("fourth request: code = %d, want 429", rec.Code)
	}

	// A different IP has its own hourly budget.
	rec = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/v1/builds", nil)
	req.RemoteAddr = "10.0.0.2:1234"
	h.ServeHTTP(rec, req)
	if rec.Code != 200 {
		t.Fatalf("other IP: code = %d, want 200", rec.Code)
	}
}
