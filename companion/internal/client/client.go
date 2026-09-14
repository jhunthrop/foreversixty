// companion/internal/client/client.go
// Package client is the companion's side of the ingest contract: the
// envelope, the device-token header, and a retry policy that knows
// which failures are worth repeating. Every request body is a byte
// slice so a retry can send it again without the caller rewinding
// anything.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// MaxResponseBody bounds what one response may cost in memory. The
// API answers every route with a small JSON envelope; anything near
// this is a proxy's error page or a server that has lost its mind.
const MaxResponseBody = 8 << 20

// MaxErrorBody is how much of an unparseable error response is kept for
// the message. A server that answers HTML must not fill the log file.
const MaxErrorBody = 2048

// Envelope is the API's response shape for every route.
type Envelope struct {
	OK        bool            `json:"ok"`
	Data      json.RawMessage `json:"data"`
	Error     *ErrorBody      `json:"error"`
	RequestID string          `json:"request_id"`
}

// ErrorBody is the envelope's error member.
type ErrorBody struct {
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
}

// Error is one failed request. The status and the field errors are kept
// because the pipeline reacts differently to a verification mismatch
// than to a dead network.
type Error struct {
	Status    int
	Message   string
	Fields    map[string]string
	RequestID string
	// RetryAfter is the server's own hint, from the Retry-After
	// header. Zero means it sent none.
	RetryAfter time.Duration
}

func (e *Error) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("api: %d %s (request %s)", e.Status, e.Message, e.RequestID)
	}
	return fmt.Sprintf("api: %d %s", e.Status, e.Message)
}

// Retryable reports whether sending the same request again could
// succeed. A 4xx other than 408 and 429 is the companion's own fault
// and repeating it only burns the queue.
func (e *Error) Retryable() bool {
	return e.Status == http.StatusRequestTimeout ||
		e.Status == http.StatusTooManyRequests ||
		e.Status >= 500
}

// Retryable reports whether err is worth another attempt. Transport
// errors always are: the player closed a laptop lid, not a contract.
func Retryable(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	var ae *Error
	if errors.As(err, &ae) {
		return ae.Retryable()
	}
	return true
}

// Unauthorized reports whether err says the device token is gone, which
// is the one failure that must stop the uploader and ask the player to
// pair again rather than retry forever.
func Unauthorized(err error) bool {
	var ae *Error
	return errors.As(err, &ae) &&
		(ae.Status == http.StatusUnauthorized || ae.Status == http.StatusForbidden)
}

// Retry is the backoff policy: exponential from Base, capped at Max,
// with full jitter so a thousand companions reconnecting after an
// outage do not arrive together.
type Retry struct {
	MaxAttempts int
	Base        time.Duration
	Max         time.Duration
}

// DefaultRetry is five attempts from one second up to thirty.
func DefaultRetry() Retry {
	return Retry{MaxAttempts: 5, Base: time.Second, Max: 30 * time.Second}
}

// Delay is how long to wait before attempt n, counting from 1.
func (r Retry) Delay(attempt int, frac float64) time.Duration {
	if attempt < 1 {
		attempt = 1
	}
	d := float64(r.Base) * math.Pow(2, float64(attempt-1))
	if d > float64(r.Max) {
		d = float64(r.Max)
	}
	return time.Duration(d * frac)
}

// Options configures a Client.
type Options struct {
	BaseURL string
	HTTP    *http.Client
	// Token returns the device token, or the empty string when the
	// companion is not paired. It is a function because pairing can
	// happen while the uploader is running.
	Token func() string
	Retry Retry
	Log   *slog.Logger
	// Sleep waits, honouring cancellation. Tests replace it so a
	// backoff test does not take thirty seconds.
	Sleep func(ctx context.Context, d time.Duration) error
	// Frac returns the jitter fraction in [0,1). Tests pin it.
	Frac func() float64
}

// Client talks to the API.
type Client struct {
	base  *url.URL
	http  *http.Client
	token func() string
	retry Retry
	log   *slog.Logger
	sleep func(ctx context.Context, d time.Duration) error
	frac  func() float64
}

// New builds a client. An unparseable base URL is a configuration error
// and is reported now rather than on the first upload.
func New(o Options) (*Client, error) {
	if o.BaseURL == "" {
		return nil, errors.New("client: base URL is empty")
	}
	u, err := url.Parse(strings.TrimRight(o.BaseURL, "/"))
	if err != nil {
		return nil, fmt.Errorf("client: base URL: %w", err)
	}
	c := &Client{
		base:  u,
		http:  o.HTTP,
		token: o.Token,
		retry: o.Retry,
		log:   o.Log,
		sleep: o.Sleep,
		frac:  o.Frac,
	}
	if c.http == nil {
		c.http = &http.Client{Timeout: 2 * time.Minute}
	}
	if c.token == nil {
		c.token = func() string { return "" }
	}
	if c.retry.MaxAttempts == 0 {
		c.retry = DefaultRetry()
	}
	if c.log == nil {
		c.log = slog.New(slog.DiscardHandler)
	}
	if c.sleep == nil {
		c.sleep = sleep
	}
	if c.frac == nil {
		c.frac = rand.Float64
	}
	return c, nil
}

func sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// request is one call to the API.
type request struct {
	Method string
	Path   string // beginning with a slash, already escaped
	Query  url.Values
	Body   []byte
	Type   string            // Content-Type, empty for no body
	Header map[string]string // extra headers
	// Anonymous skips the Authorization header. Only pairing uses it:
	// a device has no token until the claim succeeds.
	Anonymous bool
	// Out receives the envelope's data member when it is not nil.
	Out any
}

// do sends a request, retrying the failures worth retrying, and decodes
// the envelope. The returned status is the last response's, so a caller
// can tell 201 (stored) from 200 (already stored).
func (c *Client) do(ctx context.Context, r request) (int, error) {
	var lastErr error
	for attempt := 1; attempt <= c.retry.MaxAttempts; attempt++ {
		status, err := c.attempt(ctx, r)
		if err == nil {
			return status, nil
		}
		lastErr = err
		if !Retryable(err) || attempt == c.retry.MaxAttempts {
			return status, err
		}
		d := c.retry.Delay(attempt, c.frac())
		var ae *Error
		if errors.As(err, &ae) && ae.RetryAfter > d {
			d = ae.RetryAfter
		}
		c.log.Warn("retrying", "component", "client", "method", r.Method, "path", r.Path,
			"attempt", attempt, "in", d.String(), "err", err.Error())
		if serr := c.sleep(ctx, d); serr != nil {
			return status, errors.Join(err, serr)
		}
	}
	return 0, lastErr
}

func (c *Client) attempt(ctx context.Context, r request) (int, error) {
	u := *c.base
	u.Path = c.base.Path + r.Path
	if len(r.Query) > 0 {
		u.RawQuery = r.Query.Encode()
	}
	var body io.Reader
	if r.Body != nil {
		body = bytes.NewReader(r.Body)
	}
	req, err := http.NewRequestWithContext(ctx, r.Method, u.String(), body)
	if err != nil {
		return 0, err
	}
	if r.Type != "" {
		req.Header.Set("Content-Type", r.Type)
	}
	for k, v := range r.Header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Accept", "application/json")
	if !r.Anonymous {
		tok := c.token()
		if tok == "" {
			return 0, &Error{Status: http.StatusUnauthorized, Message: "this device is not paired"}
		}
		req.Header.Set("Authorization", "Bearer "+tok)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return 0, fmt.Errorf("%s %s: %w", r.Method, r.Path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		// A 204 has no body to read; draining it is only so the
		// connection can go back in the pool, and a failure there
		// costs one pooled connection and nothing the caller can act
		// on — the request itself already succeeded.
		_, _ = io.Copy(io.Discard, resp.Body)
		return resp.StatusCode, nil
	}
	raw, err := io.ReadAll(io.LimitReader(resp.Body, MaxResponseBody))
	if err != nil {
		return resp.StatusCode, fmt.Errorf("%s %s: read body: %w", r.Method, r.Path, err)
	}
	var env Envelope
	if jerr := json.Unmarshal(raw, &env); jerr != nil || (!env.OK && env.Error == nil) {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp.StatusCode, nil // a 2xx with no envelope is success
		}
		after, _ := retryAfter(resp.Header)
		return resp.StatusCode, &Error{
			Status: resp.StatusCode, Message: snippet(raw), RetryAfter: after}
	}
	if !env.OK {
		after, _ := retryAfter(resp.Header)
		return resp.StatusCode, &Error{
			Status:     resp.StatusCode,
			Message:    env.Error.Message,
			Fields:     env.Error.Fields,
			RequestID:  env.RequestID,
			RetryAfter: after,
		}
	}
	if r.Out != nil && len(env.Data) > 0 {
		if err := json.Unmarshal(env.Data, r.Out); err != nil {
			return resp.StatusCode, fmt.Errorf("%s %s: decode data: %w", r.Method, r.Path, err)
		}
	}
	return resp.StatusCode, nil
}

func snippet(b []byte) string {
	s := strings.TrimSpace(string(b))
	if s == "" {
		return "empty response"
	}
	if len(s) > MaxErrorBody {
		s = s[:MaxErrorBody] + "…"
	}
	return s
}

// retryAfter is the server's own backoff hint, honoured over ours when
// it is longer. Only seconds are supported; the API sends nothing else.
func retryAfter(h http.Header) (time.Duration, bool) {
	v := h.Get("Retry-After")
	if v == "" {
		return 0, false
	}
	n, err := strconv.Atoi(v)
	if err != nil || n < 0 {
		return 0, false
	}
	return time.Duration(n) * time.Second, true
}
