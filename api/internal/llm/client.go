package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
		{cfg.BaseURL, "LLM_BASE_URL"}, {cfg.Model, "LLM_MODEL"}, {cfg.APIKey, "LLM_API_KEY"},
	} {
		if strings.TrimSpace(req.value) == "" {
			return nil, fmt.Errorf("llm: %s is not set", req.name)
		}
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
func (c *HTTPClient) Complete(ctx context.Context, req Request) (Response, error) {
	body, err := json.Marshal(c.wire(req))
	if err != nil {
		return Response{}, fmt.Errorf("llm: encode request: %w", err)
	}
	var lastErr error
	for attempt := 1; attempt <= c.cfg.MaxRetries+1; attempt++ {
		res, hint, err := c.attempt(ctx, body)
		if err == nil {
			res.Attempts = attempt
			return res, nil
		}
		lastErr = err
		if !hint.retry || attempt > c.cfg.MaxRetries || ctx.Err() != nil {
			break
		}
		if serr := c.sleep(ctx, c.wait(hint, attempt)); serr != nil {
			return Response{}, serr
		}
	}
	return Response{}, lastErr
}

// wait is the delay before the retry that follows the given attempt.
func (c *HTTPClient) wait(h retryHint, attempt int) time.Duration {
	if h.hasAfter {
		return min(h.after, maxBackoff)
	}
	return c.jitter(min(baseBackoff<<(attempt-1), maxBackoff))
}

func (c *HTTPClient) wire(req Request) wireRequest {
	temp, maxTokens := req.Temperature, req.MaxTokens
	if temp == 0 {
		temp = defaultTemperature
	}
	if maxTokens == 0 {
		maxTokens = defaultMaxTokens
	}
	w := wireRequest{
		Model: c.cfg.Model, Messages: req.Messages, Temperature: temp,
		MaxTokens: maxTokens, ResponseFormat: responseFormat(req.JSONSchema),
	}
	if len(c.cfg.Providers) > 0 {
		w.ProviderOpts = &providerOption{Gateway: gatewayOption{Only: c.cfg.Providers}}
	}
	return w
}

// attempt makes one request under the per-attempt timeout.
func (c *HTTPClient) attempt(ctx context.Context, body []byte) (Response, retryHint, error) {
	actx, cancel := context.WithTimeout(ctx, c.cfg.Timeout)
	defer cancel()
	httpReq, err := http.NewRequestWithContext(actx, http.MethodPost, c.cfg.BaseURL+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return Response{}, retryHint{}, errors.New("llm: build request failed")
	}
	httpReq.Header.Set("Authorization", "Bearer "+c.cfg.APIKey)
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
	return c.interpret(resp, raw)
}

// interpret turns an HTTP answer into a Response or a classified error.
func (c *HTTPClient) interpret(resp *http.Response, raw []byte) (Response, retryHint, error) {
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
