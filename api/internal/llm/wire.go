package llm

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

// requestIDHeader is the gateway's per-request header, used when an error
// answer carries no completion id.
const requestIDHeader = "X-Vercel-Id"

type wireRequest struct {
	Model          string          `json:"model"`
	Messages       []Message       `json:"messages"`
	Temperature    float64         `json:"temperature"`
	MaxTokens      int             `json:"max_tokens"`
	ResponseFormat map[string]any  `json:"response_format"`
	ProviderOpts   *providerOption `json:"providerOptions,omitempty"`
	Reasoning      reasoningOption `json:"reasoning"`
}

// reasoningOption is the gateway's top-level reasoning object: enabled or effort, never both.
type reasoningOption struct {
	Enabled *bool  `json:"enabled,omitempty"`
	Effort  string `json:"effort,omitempty"`
}

func reasoningFor(level string) reasoningOption {
	if level == DefaultReasoning {
		off := false
		return reasoningOption{Enabled: &off}
	}
	return reasoningOption{Effort: level}
}

type providerOption struct {
	Gateway gatewayOption `json:"gateway"`
}

type gatewayOption struct {
	Only []string `json:"only"`
}

type wireResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
	Usage *struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		Details          *struct {
			ReasoningTokens int `json:"reasoning_tokens"`
		} `json:"completion_tokens_details"`
		Cost *float64 `json:"cost"`
	} `json:"usage"`
}

// APIError is a non-2xx answer: the status and request id only, never a body.
type APIError struct {
	Status    int
	RequestID string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("llm: gateway answered status %d (request %q)", e.Status, e.RequestID)
}

var errMalformed = errors.New("llm: gateway answered with a malformed completion")

func responseFormat(schema json.RawMessage) map[string]any {
	if len(schema) == 0 {
		return map[string]any{"type": "json_object"}
	}
	return map[string]any{
		"type": "json_schema",
		"json_schema": map[string]any{
			"name": "review", "strict": true, "schema": schema,
		},
	}
}

// reportedCost reads the gateway-reported cost of a request: usage.cost in
// the completion body. The bool is false when nothing is reported.
func reportedCost(_ *http.Response, body []byte) (float64, bool) {
	var w wireResponse
	if json.Unmarshal(body, &w) != nil || w.Usage == nil || w.Usage.Cost == nil {
		return 0, false
	}
	// A reported 0 is trusted (BYOK or free models); negative or NaN is not.
	if c := *w.Usage.Cost; c >= 0 && !math.IsNaN(c) && !math.IsInf(c, 0) {
		return c, true
	}
	return 0, false
}

// retryAfter parses a Retry-After header: delta seconds or an HTTP date.
func retryAfter(h string, now time.Time) (time.Duration, bool) {
	if h == "" {
		return 0, false
	}
	if secs, err := strconv.Atoi(h); err == nil && secs >= 0 {
		return time.Duration(secs) * time.Second, true
	}
	if t, err := http.ParseTime(h); err == nil {
		return max(t.Sub(now), 0), true
	}
	return 0, false
}

func retryableStatus(status int) bool {
	return status == http.StatusTooManyRequests || status >= http.StatusInternalServerError
}
