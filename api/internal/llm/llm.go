// Package llm is a plain net/http client for an OpenAI-compatible chat
// completions endpoint (the Vercel AI Gateway), with per-attempt timeouts,
// bounded retries, token usage and cost capture. It never logs or returns
// the API key, the request body or the response body.
package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"
)

const (
	defaultTimeout     = 60 * time.Second
	defaultMaxRetries  = 2
	defaultTemperature = 0.3
	defaultMaxTokens   = 1200
	tokensPerMTok      = 1e6
	// charsPerTokenWorstCase is the deliberately low chars-per-token ratio used
	// to over-estimate prompt tokens when the gateway reports no usage block.
	charsPerTokenWorstCase = 3
)

// Secret is a credential that never prints: %s, %v, %+v, %#v and slog all
// show a placeholder, and even a struct embedding it by an unexported field
// prints only a pointer (fmt cannot call methods there). Read the value with
// Reveal; the zero Secret is empty.
type Secret struct{ v *string }

// NewSecret wraps a credential.
func NewSecret(value string) Secret { return Secret{v: &value} }

const redacted = "[redacted]"

// String redacts the secret.
func (Secret) String() string { return redacted }

// GoString redacts the secret under %#v.
func (Secret) GoString() string { return redacted }

// LogValue redacts the secret in structured logs.
func (Secret) LogValue() slog.Value { return slog.StringValue(redacted) }

// Reveal returns the secret value; call it only to build an auth header.
func (s Secret) Reveal() string {
	if s.v == nil {
		return ""
	}
	return *s.v
}

// Config configures an HTTPClient. It is copied by New and never mutated.
type Config struct {
	BaseURL string
	Model   string
	APIKey  Secret
	// Providers pins the gateway's provider routing (providerOptions.gateway.only).
	Providers []string
	// PriceInPerMTok and PriceOutPerMTok (USD) are required: the gateway
	// documents no inline cost, so estimation is the normal path and the
	// spend caps rely on it.
	PriceInPerMTok, PriceOutPerMTok float64
	// Timeout bounds each attempt; zero means 60 s.
	Timeout time.Duration
	// MaxRetries is the retries after the first attempt; zero means 2,
	// a negative value means none.
	MaxRetries int
}

// String describes the config without the key (the key is also a Secret).
func (c Config) String() string {
	return fmt.Sprintf("llm.Config{BaseURL:%q Model:%q Providers:%v Timeout:%s MaxRetries:%d APIKey:%s}",
		c.BaseURL, c.Model, c.Providers, c.Timeout, c.MaxRetries, redacted)
}

// GoString matches String under %#v.
func (c Config) GoString() string { return c.String() }

// Message is one chat message; Role is "system", "user" or "assistant".
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Request is one completion request.
type Request struct {
	Messages []Message
	// JSONSchema, when set, selects the json_schema response format; otherwise JSON mode.
	JSONSchema  json.RawMessage
	Temperature float64 // zero means 0.3
	MaxTokens   int     // zero means 1200
}

// Usage is the token usage the gateway reported.
type Usage struct {
	PromptTokens, CompletionTokens int
}

// Response is one completion.
type Response struct {
	Content       string
	Usage         Usage
	CostUSD       float64
	CostEstimated bool
	RequestID     string
	Model         string
	Attempts      int
}

// Client is what callers depend on; HTTPClient and Fake implement it.
type Client interface {
	Complete(ctx context.Context, req Request) (Response, error)
}
