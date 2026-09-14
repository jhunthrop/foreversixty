package client

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// newTestClient builds a client with no real waiting and no randomness,
// recording what it slept for.
func newTestClient(t *testing.T, base string, slept *[]time.Duration) *Client {
	t.Helper()
	c, err := New(Options{
		BaseURL: base,
		Token:   func() string { return "fsd_test" },
		Retry:   Retry{MaxAttempts: 4, Base: time.Second, Max: 8 * time.Second},
		Frac:    func() float64 { return 1 },
		Sleep: func(_ context.Context, d time.Duration) error {
			*slept = append(*slept, d)
			return nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestTheEnvelopeIsUnwrappedIntoTheOutValue(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer fsd_test" {
			t.Errorf("Authorization = %q", got)
		}
		w.Write([]byte(`{"ok":true,"data":{"id":"abc123"},"error":null,"request_id":"r1"}`))
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	var out struct {
		ID string `json:"id"`
	}
	if _, err := c.do(t.Context(), request{Method: http.MethodGet, Path: "/v1/x", Out: &out}); err != nil {
		t.Fatal(err)
	}
	if out.ID != "abc123" {
		t.Fatalf("id = %q", out.ID)
	}
}

func TestAFailedEnvelopeBecomesATypedError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"ok":false,"data":null,"error":{"message":"verification failed",` +
			`"fields":{"metrics":"dps differs by 4%"}},"request_id":"r2"}`))
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	_, err := c.do(t.Context(), request{Method: http.MethodPut, Path: "/v1/x"})
	var ae *Error
	if !errors.As(err, &ae) {
		t.Fatalf("err = %v, want *Error", err)
	}
	if ae.Status != http.StatusConflict || ae.Fields["metrics"] != "dps differs by 4%" || ae.RequestID != "r2" {
		t.Fatalf("error = %+v", ae)
	}
	if ae.Retryable() {
		t.Error("a 409 must not be retried")
	}
	if len(slept) != 0 {
		t.Errorf("slept %v on a 409", slept)
	}
}

func TestServerErrorsAreRetriedWithExponentialBackoff(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if hits.Add(1) < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"ok":false,"error":{"message":"boom"}}`))
			return
		}
		w.Write([]byte(`{"ok":true,"data":null}`))
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	if _, err := c.do(t.Context(), request{Method: http.MethodGet, Path: "/v1/x"}); err != nil {
		t.Fatal(err)
	}
	if hits.Load() != 3 {
		t.Fatalf("server saw %d requests, want 3", hits.Load())
	}
	want := []time.Duration{time.Second, 2 * time.Second}
	if len(slept) != len(want) || slept[0] != want[0] || slept[1] != want[1] {
		t.Fatalf("slept %v, want %v", slept, want)
	}
}

func TestTheBackoffIsCappedAndJittered(t *testing.T) {
	r := Retry{MaxAttempts: 8, Base: time.Second, Max: 10 * time.Second}
	for _, tc := range []struct {
		attempt int
		frac    float64
		want    time.Duration
	}{
		{1, 1, time.Second},
		{2, 1, 2 * time.Second},
		{5, 1, 10 * time.Second},
		{9, 1, 10 * time.Second},
		{3, 0.5, 2 * time.Second},
	} {
		if got := r.Delay(tc.attempt, tc.frac); got != tc.want {
			t.Errorf("Delay(%d, %v) = %v, want %v", tc.attempt, tc.frac, got, tc.want)
		}
	}
}

func TestRetryAfterOverridesAShorterBackoff(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if hits.Add(1) == 1 {
			w.Header().Set("Retry-After", "7")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"ok":false,"error":{"message":"slow down"}}`))
			return
		}
		w.Write([]byte(`{"ok":true,"data":null}`))
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	if _, err := c.do(t.Context(), request{Method: http.MethodGet, Path: "/v1/x"}); err != nil {
		t.Fatal(err)
	}
	if len(slept) != 1 || slept[0] != 7*time.Second {
		t.Fatalf("slept %v, want [7s]", slept)
	}
}

func TestAnUnpairedDeviceFailsWithoutSendingAnything(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("the client sent a request without a token")
	}))
	defer srv.Close()
	c, err := New(Options{BaseURL: srv.URL, Token: func() string { return "" }})
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.do(t.Context(), request{Method: http.MethodGet, Path: "/v1/x"})
	if !Unauthorized(err) {
		t.Fatalf("err = %v, want an unauthorized error", err)
	}
}

func TestANonEnvelopeErrorBodyIsKeptAsASnippet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<html>gateway</html>"))
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	_, err := c.do(t.Context(), request{Method: http.MethodGet, Path: "/v1/x"})
	var ae *Error
	if !errors.As(err, &ae) || ae.Status != http.StatusBadGateway || ae.Message != "<html>gateway</html>" {
		t.Fatalf("err = %v", err)
	}
}

func TestAnEmptyBaseURLIsRefused(t *testing.T) {
	if _, err := New(Options{}); err == nil {
		t.Fatal("New accepted an empty base URL")
	}
}

func TestACancelledContextIsNotRetried(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if _, err := c.do(ctx, request{Method: http.MethodGet, Path: "/v1/x"}); err == nil {
		t.Fatal("a cancelled request succeeded")
	}
	if len(slept) != 0 {
		t.Errorf("slept %v on a cancelled context", slept)
	}
}

func TestErrorMessagesNameTheRequestWhenTheServerGaveOne(t *testing.T) {
	with := (&Error{Status: 409, Message: "nope", RequestID: "r9"}).Error()
	if !strings.Contains(with, "409") || !strings.Contains(with, "r9") {
		t.Errorf("with a request id: %q", with)
	}
	without := (&Error{Status: 500, Message: "boom"}).Error()
	if strings.Contains(without, "request") {
		t.Errorf("without a request id: %q", without)
	}
}

func TestRetryableClassifiesFailures(t *testing.T) {
	if Retryable(nil) {
		t.Error("nil is retryable")
	}
	if Retryable(context.Canceled) {
		t.Error("a cancelled context is retryable")
	}
	if !Retryable(errors.New("connection reset by peer")) {
		t.Error("a transport error is not retryable")
	}
	for status, want := range map[int]bool{
		408: true, 429: true, 500: true, 503: true,
		400: false, 401: false, 404: false, 409: false, 413: false,
	} {
		if got := (&Error{Status: status}).Retryable(); got != want {
			t.Errorf("%d retryable = %v, want %v", status, got, want)
		}
	}
	if !Unauthorized(&Error{Status: 403}) || Unauthorized(&Error{Status: 409}) {
		t.Error("Unauthorized misclassified a status")
	}
}

func TestDefaultsAreFilledIn(t *testing.T) {
	c, err := New(Options{BaseURL: "https://api.foreversixty.gg/"})
	if err != nil {
		t.Fatal(err)
	}
	if c.retry != DefaultRetry() || c.token() != "" || c.frac == nil || c.sleep == nil || c.log == nil {
		t.Fatalf("defaults = %+v", c.retry)
	}
	if c.base.String() != "https://api.foreversixty.gg" {
		t.Errorf("base = %q", c.base.String())
	}
	if _, err := New(Options{BaseURL: "://nonsense"}); err == nil {
		t.Error("an unparseable base URL was accepted")
	}
}

func TestSleepHonoursCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := sleep(ctx, time.Hour); err == nil {
		t.Fatal("sleep ignored a cancelled context")
	}
	if err := sleep(context.Background(), time.Millisecond); err != nil {
		t.Fatal(err)
	}
}

func TestSnippetTrimsAndDescribesAnEmptyBody(t *testing.T) {
	if got := snippet([]byte("   ")); got != "empty response" {
		t.Errorf("empty = %q", got)
	}
	long := snippet([]byte(strings.Repeat("x", MaxErrorBody+50)))
	if len(long) != MaxErrorBody+len("…") {
		t.Errorf("long snippet is %d characters", len(long))
	}
}
