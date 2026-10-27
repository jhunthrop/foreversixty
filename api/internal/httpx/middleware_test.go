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

func TestRateLimitByKeyAllowsTheBudgetThenReturns429PerKey(t *testing.T) {
	keyOf := func(r *http.Request) (string, bool) {
		return r.Header.Get("X-Account"), r.Header.Get("X-Account") != ""
	}
	h := Chain(okHandler(), RequestID(), RateLimitByKey(3, time.Hour, keyOf))
	post := func(account string) int {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/v1/sims/run", nil)
		req.RemoteAddr = "10.0.0.1:1234" // same address for every call: the key is the account, not the IP
		req.Header.Set("X-Account", account)
		h.ServeHTTP(rec, req)
		return rec.Code
	}
	for i := 0; i < 3; i++ {
		if code := post("42"); code != 200 {
			t.Fatalf("request %d for account 42: code = %d, want 200", i+1, code)
		}
	}
	if code := post("42"); code != http.StatusTooManyRequests {
		t.Fatalf("fourth request for account 42: code = %d, want 429", code)
	}
	// A different account, same address, has its own budget.
	if code := post("7"); code != 200 {
		t.Fatalf("other account: code = %d, want 200", code)
	}
}

func TestRateLimitByKeyExemptsAKeylessRequestEntirely(t *testing.T) {
	calls := 0
	keyOf := func(r *http.Request) (string, bool) { calls++; return "", false }
	h := Chain(okHandler(), RequestID(), RateLimitByKey(1, time.Hour, keyOf))
	for i := 0; i < 5; i++ {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/v1/uploads", nil)
		h.ServeHTTP(rec, req)
		if rec.Code != 200 {
			t.Fatalf("request %d: code = %d, want 200 (no key means no budget to exceed)", i+1, rec.Code)
		}
	}
	if calls != 5 {
		t.Fatalf("keyOf was called %d times, want 5", calls)
	}
}

func TestRateLimitExceptLetsAnExemptPathPastTheSpentBudget(t *testing.T) {
	exempt := func(r *http.Request) bool {
		return r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/b/")
	}
	h := Chain(okHandler(), RequestID(), RateLimitExcept(120, 1, exempt))
	get := func(path string) int {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", path, nil)
		req.RemoteAddr = "10.0.0.7:1234"
		h.ServeHTTP(rec, req)
		return rec.Code
	}

	// Spend the whole 120-per-minute budget from one address.
	for i := 0; i < 120; i++ {
		if code := get("/version"); code != 200 {
			t.Fatalf("request %d code = %d, want 200", i+1, code)
		}
	}
	// The 121st request is served when it is exempt...
	if code := get("/b/x"); code != 200 {
		t.Fatalf("exempt request past the budget code = %d, want 200", code)
	}
	// ...and rejected when it is not.
	if code := get("/version"); code != 429 {
		t.Fatalf("non-exempt request past the budget code = %d, want 429", code)
	}
}

// credReq builds a request from one address, optionally carrying a device bearer.
func credReq(ip, token string) *http.Request {
	req := httptest.NewRequest("GET", "/v1/x", nil)
	req.RemoteAddr = ip + ":1234"
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return req
}

func credCode(h http.Handler, ip, token string) int {
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, credReq(ip, token))
	return rec.Code
}

func frozenClock(t *testing.T) {
	t.Helper()
	frozen := time.Date(2026, time.October, 19, 0, 0, 0, 0, time.UTC)
	Now = func() time.Time { return frozen }
	t.Cleanup(func() { Now = time.Now })
}

func TestRateLimitByCredentialDeviceTrafficLeavesTheIPBucketAlone(t *testing.T) {
	frozenClock(t)
	h := Chain(okHandler(), RequestID(), RateLimitByCredential(120, 120, 10, 1, nil))
	for i := 0; i < 121; i++ {
		credCode(h, "10.1.0.1", "fsd_aaaaaaaa")
	}
	if code := credCode(h, "10.1.0.1", ""); code != 200 {
		t.Fatalf("browser GET after 121 device requests = %d, want 200", code)
	}
}

func TestRateLimitByCredentialDeviceOverBudgetDoesNotStarveTheBrowser(t *testing.T) {
	frozenClock(t)
	h := Chain(okHandler(), RequestID(), RateLimitByCredential(120, 120, 10, 1, nil))
	for i := 0; i < 120; i++ {
		if code := credCode(h, "10.1.0.2", "fsd_bbbbbbbb"); code != 200 {
			t.Fatalf("device request %d = %d, want 200", i+1, code)
		}
	}
	if code := credCode(h, "10.1.0.2", "fsd_bbbbbbbb"); code != 429 {
		t.Fatalf("device request past its budget = %d, want 429", code)
	}
	if code := credCode(h, "10.1.0.2", ""); code != 200 {
		t.Fatalf("browser from the same IP = %d, want 200", code)
	}
}

func TestRateLimitByCredentialBrowserBucketIsStillPerIP(t *testing.T) {
	frozenClock(t)
	h := Chain(okHandler(), RequestID(), RateLimitByCredential(3, 120, 10, 1, nil))
	for i := 0; i < 3; i++ {
		credCode(h, "10.1.0.3", "")
	}
	if code := credCode(h, "10.1.0.3", ""); code != 429 {
		t.Fatalf("browser past the IP budget = %d, want 429", code)
	}
	if code := credCode(h, "10.1.0.3", "fsd_cccccccc"); code != 200 {
		t.Fatalf("device from the spent IP = %d, want 200", code)
	}
}

func TestRateLimitByCredentialTwoDevicesOnOneIPHaveSeparateBudgets(t *testing.T) {
	frozenClock(t)
	h := Chain(okHandler(), RequestID(), RateLimitByCredential(120, 5, 2, 1, nil))
	for i := 0; i < 5; i++ {
		credCode(h, "10.1.0.4", "fsd_dddddddd")
	}
	if code := credCode(h, "10.1.0.4", "fsd_dddddddd"); code != 429 {
		t.Fatalf("first device past budget = %d, want 429", code)
	}
	if code := credCode(h, "10.1.0.4", "fsd_eeeeeeee"); code != 200 {
		t.Fatalf("second device = %d, want 200", code)
	}
}

func TestRateLimitByCredentialADeadTokenDropsToTheDeadBudget(t *testing.T) {
	frozenClock(t)
	h := Chain(okHandler(), RequestID(), RateLimitByCredential(120, 120, 3, 1, nil))
	tok := "fsd_deadtoken1"
	if code := credCode(h, "10.1.0.5", tok); code != 200 {
		t.Fatalf("first request = %d, want 200", code)
	}
	MarkDeadToken(sha256Sum(tok))
	n := 0
	for i := 0; i < 10; i++ {
		if credCode(h, "10.1.0.5", tok) == 200 {
			n++
		}
	}
	if n != 3 {
		t.Fatalf("dead token served %d requests, want its dead budget of 3", n)
	}
	if code := credCode(h, "10.1.0.5", "fsd_livetoken1"); code != 200 {
		t.Fatalf("other token = %d, want 200", code)
	}
}

func TestRateLimitByCredentialMalformedTokensCountAgainstTheirOwnBucket(t *testing.T) {
	frozenClock(t)
	h := Chain(okHandler(), RequestID(), RateLimitByCredential(2, 2, 1, 1, nil))
	for i := 0; i < 5; i++ {
		credCode(h, "10.1.0.6", "fsd_")
	}
	if code := credCode(h, "10.1.0.6", ""); code != 200 {
		t.Fatalf("browser after a flood of bad tokens = %d, want 200", code)
	}
}

func TestRateLimitByCredentialExemptRequestsPass(t *testing.T) {
	frozenClock(t)
	exempt := func(r *http.Request) bool { return strings.HasPrefix(r.URL.Path, "/b/") }
	h := Chain(okHandler(), RequestID(), RateLimitByCredential(1, 1, 1, 1, exempt))
	credCode(h, "10.1.0.7", "")
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/b/x", nil)
	req.RemoteAddr = "10.1.0.7:1"
	h.ServeHTTP(rec, req)
	if rec.Code != 200 {
		t.Fatalf("exempt request = %d, want 200", rec.Code)
	}
}

func TestKeyLimiterAllowsNPerWindow(t *testing.T) {
	frozenClock(t)
	l := NewKeyLimiter(30, time.Minute)
	for i := 1; i <= 30; i++ {
		if !l.Allow("k") {
			t.Fatalf("Allow #%d = false, want true", i)
		}
	}
	if l.Allow("k") {
		t.Fatal("31st Allow = true, want false")
	}
	if !l.Allow("other") {
		t.Fatal("another key must have its own bucket")
	}
}
