package llm

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
	"strings"
	"time"
)

const (
	baseBackoff = 500 * time.Millisecond
	maxBackoff  = 30 * time.Second
	maxBodyRead = 4 << 20
)

// HTTPClient talks to an OpenAI-compatible endpoint. sleep and jitter are
// fields so tests run with no real wait.
type HTTPClient struct {
	cfg    Config
	hc     *http.Client
	sleep  func(ctx context.Context, d time.Duration) error
	jitter func(d time.Duration) time.Duration
	now    func() time.Time
}

// retryHint tells Complete whether and when to try again.
type retryHint struct {
	retry    bool
	after    time.Duration
	hasAfter bool
}

// New validates cfg, applies defaults and returns a client. It fails fast on
// a missing BaseURL, Model or APIKey, naming the environment variable.
func New(cfg Config, hc *http.Client) (*HTTPClient, error) {
	for _, req := range []struct{ value, name string }{
		{cfg.BaseURL, "LLM_BASE_URL"}, {cfg.Model, "LLM_MODEL"}, {cfg.APIKey.Reveal(), "LLM_API_KEY"},
	} {
		if strings.TrimSpace(req.value) == "" {
			return nil, fmt.Errorf("llm: %s is not set", req.name)
		}
	}
	if err := requirePrice(cfg.PriceInPerMTok, "LLM_PRICE_IN_PER_MTOK"); err != nil {
		return nil, err
	}
	if err := requirePrice(cfg.PriceOutPerMTok, "LLM_PRICE_OUT_PER_MTOK"); err != nil {
		return nil, err
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = defaultTimeout
	}
	switch {
	case cfg.MaxRetries == 0:
		cfg.MaxRetries = defaultMaxRetries
	case cfg.MaxRetries < 0:
		cfg.MaxRetries = 0
	}
	cfg.BaseURL = strings.TrimRight(cfg.BaseURL, "/")
	cfg.Providers = append([]string(nil), cfg.Providers...)
	if hc == nil {
		hc = &http.Client{}
	}
	return &HTTPClient{cfg: cfg, hc: hc, sleep: sleepCtx, jitter: halfJitter, now: time.Now}, nil
}

// requirePrice fails fast on a missing fallback price: a zero or invalid
// price would let an estimated cost silently under-count against the caps.
func requirePrice(price float64, name string) error {
	if !(price > 0) || math.IsInf(price, 0) {
		return fmt.Errorf("llm: %s must be set to a positive price (USD per million tokens)", name)
	}
	return nil
}

// String describes the client without its credentials.
func (c *HTTPClient) String() string {
	return fmt.Sprintf("llm.HTTPClient{%s}", c.cfg)
}

// GoString matches String under %#v.
func (c *HTTPClient) GoString() string { return c.String() }

func sleepCtx(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// halfJitter spreads a backoff over [d/2, d].
func halfJitter(d time.Duration) time.Duration {
	return d/2 + time.Duration(rand.Int64N(int64(d/2)+1))
}

// Complete sends one chat completion, retrying 429, 5xx and network errors.
// On failure the returned Response still carries Attempts (and RequestID when
// the gateway sent one).
func (c *HTTPClient) Complete(ctx context.Context, req Request) (Response, error) {
	req = withDefaults(req)
	body, err := json.Marshal(c.wire(req))
	if err != nil {
		return Response{}, fmt.Errorf("llm: encode request: %w", err)
	}
	fallback := worstCaseUsage(req)
	var lastErr error
	attempts := 0
	for attempt := 1; attempt <= c.cfg.MaxRetries+1; attempt++ {
		res, hint, err := c.attempt(ctx, body, fallback)
		attempts = attempt
		res.Attempts = attempt
		if err == nil {
			c.logResult(ctx, res)
			return res, nil
		}
		lastErr = err
		if !hint.retry || attempt > c.cfg.MaxRetries || ctx.Err() != nil {
			break
		}
		if serr := c.sleep(ctx, c.wait(hint, attempt)); serr != nil {
			lastErr = serr
			break
		}
	}
	return c.failure(ctx, attempts, lastErr), lastErr
}

// failure builds the Response returned alongside a final error and logs it.
func (c *HTTPClient) failure(ctx context.Context, attempts int, err error) Response {
	res := Response{Model: c.cfg.Model, Attempts: attempts}
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		res.RequestID = apiErr.RequestID
	}
	slog.WarnContext(ctx, "llm request failed", "model", c.cfg.Model, "request_id", res.RequestID, "attempts", attempts, "error", err.Error())
	return res
}

func (c *HTTPClient) logResult(ctx context.Context, r Response) {
	slog.InfoContext(ctx, "llm request", "model", r.Model, "request_id", r.RequestID, "attempts", r.Attempts,
		"prompt_tokens", r.Usage.PromptTokens, "completion_tokens", r.Usage.CompletionTokens,
		"cost_usd", r.CostUSD, "cost_estimated", r.CostEstimated)
}

// withDefaults fills the zero-valued temperature and token ceiling.
func withDefaults(req Request) Request {
	if req.Temperature == 0 {
		req.Temperature = defaultTemperature
	}
	if req.MaxTokens == 0 {
		req.MaxTokens = defaultMaxTokens
	}
	return req
}

// worstCaseUsage is the charge assumed when the gateway reports no usage:
// the full output ceiling plus a low-ratio estimate of the prompt tokens.
func worstCaseUsage(req Request) Usage {
	chars := len(req.JSONSchema)
	for _, m := range req.Messages {
		chars += len(m.Role) + len(m.Content)
	}
	prompt := (chars + charsPerTokenWorstCase - 1) / charsPerTokenWorstCase
	return Usage{PromptTokens: prompt, CompletionTokens: req.MaxTokens}
}

// wait is the delay before the retry that follows the given attempt.
func (c *HTTPClient) wait(h retryHint, attempt int) time.Duration {
	if h.hasAfter {
		return min(h.after, maxBackoff)
	}
	return c.jitter(min(baseBackoff<<(attempt-1), maxBackoff))
}

// wire builds the request body from a request that has its defaults applied.
func (c *HTTPClient) wire(req Request) wireRequest {
	w := wireRequest{
		Model: c.cfg.Model, Messages: req.Messages, Temperature: req.Temperature,
		MaxTokens: req.MaxTokens, ResponseFormat: responseFormat(req.JSONSchema),
	}
	if len(c.cfg.Providers) > 0 {
		w.ProviderOpts = &providerOption{Gateway: gatewayOption{Only: c.cfg.Providers}}
	}
	return w
}

// attempt makes one request under the per-attempt timeout.
func (c *HTTPClient) attempt(ctx context.Context, body []byte, fallback Usage) (Response, retryHint, error) {
	actx, cancel := context.WithTimeout(ctx, c.cfg.Timeout)
	defer cancel()
	httpReq, err := http.NewRequestWithContext(actx, http.MethodPost, c.cfg.BaseURL+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return Response{}, retryHint{}, errors.New("llm: build request failed")
	}
	httpReq.Header.Set("Authorization", "Bearer "+c.cfg.APIKey.Reveal())
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.hc.Do(httpReq)
	if err != nil {
		return Response{}, retryHint{retry: true}, scrub(err)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, maxBodyRead))
	if err != nil {
		return Response{}, retryHint{retry: true}, scrub(err)
	}
	return c.interpret(resp, raw, fallback)
}

// interpret turns an HTTP answer into a Response or a classified error.
func (c *HTTPClient) interpret(resp *http.Response, raw []byte, fallback Usage) (Response, retryHint, error) {
	if resp.StatusCode/100 != 2 {
		hint := retryHint{retry: retryableStatus(resp.StatusCode)}
		hint.after, hint.hasAfter = retryAfter(resp.Header.Get("Retry-After"), c.now())
		return Response{}, hint, &APIError{Status: resp.StatusCode, RequestID: resp.Header.Get(requestIDHeader)}
	}
	var w wireResponse
	if err := json.Unmarshal(raw, &w); err != nil || len(w.Choices) == 0 {
		return Response{}, retryHint{}, errMalformed
	}
	res := Response{
		Content: w.Choices[0].Message.Content, RequestID: w.ID, Model: w.Model,
	}
	if res.RequestID == "" {
		res.RequestID = resp.Header.Get(requestIDHeader)
	}
	res.Usage = fallback // no usage block: charge the worst case
	if w.Usage != nil {
		res.Usage = Usage{PromptTokens: w.Usage.PromptTokens, CompletionTokens: w.Usage.CompletionTokens}
	}
	res.CostUSD, res.CostEstimated = c.cost(resp, raw, res.Usage)
	return res, retryHint{}, nil
}

// cost prefers the gateway-reported cost, else estimates from the fallback prices.
func (c *HTTPClient) cost(resp *http.Response, raw []byte, u Usage) (float64, bool) {
	if v, ok := reportedCost(resp, raw); ok {
		return v, false
	}
	in := float64(u.PromptTokens) * c.cfg.PriceInPerMTok
	out := float64(u.CompletionTokens) * c.cfg.PriceOutPerMTok
	return (in + out) / tokensPerMTok, true
}

// scrub strips the URL wrapper from a transport error, keeping its cause.
func scrub(err error) error {
	var ue *url.Error
	if errors.As(err, &ue) {
		err = ue.Err
	}
	return fmt.Errorf("llm: request failed: %w", err)
}
