// Package llm is a plain net/http client for an OpenAI-compatible chat
// completions endpoint (the Vercel AI Gateway), with per-attempt timeouts,
// bounded retries, token usage and cost capture. It never logs or returns
// the API key, the request body or the response body.
package llm

import (
	"context"
	"encoding/json"
	"time"
)

const (
	defaultTimeout     = 60 * time.Second
	defaultMaxRetries  = 2
	defaultTemperature = 0.3
	defaultMaxTokens   = 1200
	tokensPerMTok      = 1e6
)

// Config configures an HTTPClient. It is copied by New and never mutated.
type Config struct {
	BaseURL string
	Model   string
	APIKey  string
	// Providers pins the gateway's provider routing (providerOptions.gateway.only).
	Providers []string
	// PriceInPerMTok and PriceOutPerMTok (USD) are used only when the gateway
	// reports no cost for a request.
	PriceInPerMTok, PriceOutPerMTok float64
	// Timeout bounds each attempt; zero means 60 s.
	Timeout time.Duration
	// MaxRetries is the retries after the first attempt; zero means 2,
	// a negative value means none.
	MaxRetries int
}

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
