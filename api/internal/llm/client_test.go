package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

const testKey = "test-key-SECRET"

var testSecret = NewSecret(testKey)

func gateway(t *testing.T, h http.HandlerFunc) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	return srv
}

func decode(t *testing.T, r *http.Request) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		t.Errorf("decode: %v", err)
	}
	return m
}

// writeCompletion answers 200; cost < 0 omits the cost field, tokens < 0 omits the usage block.
func writeCompletion(w http.ResponseWriter, content string, in, out int, cost float64, id string) {
	usage := ""
	if in >= 0 {
		usage = fmt.Sprintf(`"usage":{"prompt_tokens":%d,"completion_tokens":%d`, in, out)
		if cost >= 0 {
			usage += fmt.Sprintf(`,"cost":%v`, cost)
		}
		usage += "},"
	}
	c, _ := json.Marshal(content)
	fmt.Fprintf(w, `{"id":%q,"model":"deepseek/test",%s"choices":[{"message":{"role":"assistant","content":%s}}]}`, id, usage, c)
}

func mustNew(t *testing.T, cfg Config) (*HTTPClient, *[]time.Duration) {
	t.Helper()
	if cfg.PriceInPerMTok == 0 && cfg.PriceOutPerMTok == 0 {
		cfg.PriceInPerMTok, cfg.PriceOutPerMTok = 1, 1
	}
	c, err := New(cfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	var sleeps []time.Duration
	c.sleep = func(_ context.Context, d time.Duration) error { sleeps = append(sleeps, d); return nil }
	c.jitter = func(d time.Duration) time.Duration { return d }
	return c, &sleeps
}

var userMsg = Request{Messages: []Message{{Role: "user", Content: "{}"}}}

func TestComplete_SendsJSONModeAndReadsUsageAndCost(t *testing.T) {
	srv := gateway(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/chat/completions" || r.Header.Get("Authorization") != "Bearer "+testKey {
			t.Errorf("path %q auth %q", r.URL.Path, r.Header.Get("Authorization"))
		}
		body := decode(t, r)
		rf, _ := body["response_format"].(map[string]any)
		if body["model"] != "deepseek/test" || rf["type"] != "json_object" || body["max_tokens"] != float64(1200) || body["temperature"] != 0.3 {
			t.Errorf("body %v", body)
		}
		if _, ok := body["providerOptions"]; ok {
			t.Errorf("providerOptions sent without pin: %v", body)
		}
		writeCompletion(w, `{"opening":"x"}`, 7000, 800, 0.0021, "gen_123")
	})
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "deepseek/test", APIKey: testSecret})
	res, err := c.Complete(context.Background(), userMsg)
	if err != nil || res.Content != `{"opening":"x"}` || res.Usage.PromptTokens != 7000 || res.Usage.CompletionTokens != 800 ||
		res.CostUSD != 0.0021 || res.CostEstimated || res.RequestID != "gen_123" || res.Attempts != 1 || res.Model != "deepseek/test" {
		t.Fatalf("%+v %v", res, err)
	}
}

func TestComplete_JSONSchemaWhenGiven(t *testing.T) {
	srv := gateway(t, func(w http.ResponseWriter, r *http.Request) {
		rf, _ := decode(t, r)["response_format"].(map[string]any)
		js, _ := rf["json_schema"].(map[string]any)
		schema, _ := js["schema"].(map[string]any)
		if rf["type"] != "json_schema" || schema["type"] != "object" {
			t.Errorf("response_format %v", rf)
		}
		writeCompletion(w, "{}", 1, 1, -1, "g")
	})
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret})
	req := userMsg
	req.JSONSchema = json.RawMessage(`{"type":"object"}`)
	if _, err := c.Complete(context.Background(), req); err != nil {
		t.Fatal(err)
	}
}

func TestComplete_RetriesOn429HonouringRetryAfter(t *testing.T) {
	var n atomic.Int32
	srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) {
		if n.Add(1) == 1 {
			w.Header().Set("Retry-After", "2")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		writeCompletion(w, "{}", 1, 1, 0.1, "g")
	})
	c, sleeps := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret})
	res, err := c.Complete(context.Background(), userMsg)
	if err != nil || res.Attempts != 2 || len(*sleeps) != 1 || (*sleeps)[0] != 2*time.Second {
		t.Fatalf("%+v %v sleeps %v", res, err, *sleeps)
	}
}

func TestComplete_RetryAfterHTTPDate(t *testing.T) {
	now := time.Date(2026, 10, 27, 12, 0, 0, 0, time.UTC)
	d, ok := retryAfter(now.Add(3*time.Second).Format(http.TimeFormat), now)
	if !ok || d != 3*time.Second {
		t.Fatalf("%v %v", d, ok)
	}
	if _, ok := retryAfter("soon", now); ok {
		t.Fatal("garbage accepted")
	}
}

func TestComplete_RetriesTwiceOn5xxThenFails(t *testing.T) {
	var n atomic.Int32
	statuses := []int{500, 502, 503}
	srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("X-Vercel-Id", "req_9")
		w.WriteHeader(statuses[n.Add(1)-1])
		fmt.Fprint(w, "UPSTREAM-BODY-LEAK")
	})
	c, sleeps := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret})
	_, err := c.Complete(context.Background(), userMsg)
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Status != 503 || n.Load() != 3 {
		t.Fatalf("err %v attempts %d", err, n.Load())
	}
	if !strings.Contains(err.Error(), "503") || !strings.Contains(err.Error(), "req_9") || strings.Contains(err.Error(), "LEAK") {
		t.Fatalf("error text %q", err)
	}
	want := []time.Duration{500 * time.Millisecond, time.Second}
	if len(*sleeps) != 2 || (*sleeps)[0] != want[0] || (*sleeps)[1] != want[1] {
		t.Fatalf("sleeps %v", *sleeps)
	}
}

func TestComplete_NoRetryOn400(t *testing.T) {
	var n atomic.Int32
	srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) {
		n.Add(1)
		w.WriteHeader(http.StatusBadRequest)
	})
	c, sleeps := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret})
	if _, err := c.Complete(context.Background(), userMsg); err == nil || n.Load() != 1 || len(*sleeps) != 0 {
		t.Fatalf("err %v attempts %d", err, n.Load())
	}
}

func TestComplete_TimeoutPerAttempt(t *testing.T) {
	var n atomic.Int32
	srv := gateway(t, func(_ http.ResponseWriter, r *http.Request) {
		n.Add(1)
		select {
		case <-r.Context().Done():
		case <-time.After(2 * time.Second):
		}
	})
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret, Timeout: 50 * time.Millisecond})
	_, err := c.Complete(context.Background(), userMsg)
	if !errors.Is(err, context.DeadlineExceeded) || n.Load() != 3 {
		t.Fatalf("err %v attempts %d", err, n.Load())
	}
}

func TestComplete_MalformedJSONFailsWithoutRetry(t *testing.T) {
	var n atomic.Int32
	srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) {
		n.Add(1)
		fmt.Fprint(w, `{"choices": [`)
	})
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret})
	if _, err := c.Complete(context.Background(), userMsg); !errors.Is(err, errMalformed) || n.Load() != 1 {
		t.Fatalf("err %v attempts %d", err, n.Load())
	}
}

func TestComplete_NoChoicesIsMalformed(t *testing.T) {
	srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, `{"id":"g","choices":[]}`) })
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret})
	if _, err := c.Complete(context.Background(), userMsg); !errors.Is(err, errMalformed) {
		t.Fatal(err)
	}
}

func TestComplete_EstimatesCostWhenNoneReported(t *testing.T) {
	srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) { writeCompletion(w, "{}", 7000, 800, -1, "g") })
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret, PriceInPerMTok: 0.2, PriceOutPerMTok: 0.8})
	res, err := c.Complete(context.Background(), userMsg)
	if err != nil || !res.CostEstimated || res.CostUSD < 0.002039 || res.CostUSD > 0.002041 {
		t.Fatalf("%+v %v", res, err)
	}
}

func TestComplete_NoUsageBlockChargesWorstCase(t *testing.T) {
	srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) { writeCompletion(w, "{}", -1, 0, -1, "g") })
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret, PriceInPerMTok: 1, PriceOutPerMTok: 2})
	req := Request{Messages: []Message{{Role: "user", Content: strings.Repeat("x", 296)}}, MaxTokens: 500}
	res, err := c.Complete(context.Background(), req)
	// "user" (4) + 296 chars = 300 chars -> 100 prompt tokens at 3 chars per token.
	wantCost := (100*1 + 500*2) / 1e6
	if err != nil || res.Usage != (Usage{PromptTokens: 100, CompletionTokens: 500}) || !res.CostEstimated || math.Abs(res.CostUSD-wantCost) > 1e-12 {
		t.Fatalf("%+v %v", res, err)
	}
}

func TestComplete_IgnoresNegativeOrNaNCost(t *testing.T) {
	for _, cost := range []string{"-0.5"} {
		srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintf(w, `{"id":"g","choices":[{"message":{"content":"{}"}}],"usage":{"prompt_tokens":1000000,"completion_tokens":0,"cost":%s}}`, cost)
		})
		c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret, PriceInPerMTok: 3, PriceOutPerMTok: 1})
		res, err := c.Complete(context.Background(), userMsg)
		if err != nil || !res.CostEstimated || res.CostUSD != 3 {
			t.Fatalf("cost %s: %+v %v", cost, res, err)
		}
	}
}

func TestReportedCost_ZeroIsTrusted(t *testing.T) {
	// A gateway-reported 0 (BYOK or a free model) is real, not a missing value.
	if v, ok := reportedCost(nil, []byte(`{"usage":{"cost":0}}`)); !ok || v != 0 {
		t.Fatalf("%v %v", v, ok)
	}
}

func TestSecretNeverPrints(t *testing.T) {
	cfg := Config{BaseURL: "http://x", Model: "m", APIKey: testSecret, PriceInPerMTok: 1, PriceOutPerMTok: 1}
	c, err := New(cfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	type holder struct{ cfg Config }
	for _, out := range []string{
		fmt.Sprintf("%v %+v %#v %s", cfg, cfg, cfg, cfg), fmt.Sprintf("%v %+v %#v", c, c, c),
		fmt.Sprintf("%+v", holder{cfg}), fmt.Sprintf("%+v %#v %s", cfg.APIKey, cfg.APIKey, cfg.APIKey),
	} {
		if strings.Contains(out, testKey) {
			t.Fatalf("key printed: %s", out)
		}
	}
}

func TestComplete_AttemptsOnError(t *testing.T) {
	srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusBadGateway) })
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret})
	res, err := c.Complete(context.Background(), userMsg)
	if err == nil || res.Attempts != 3 {
		t.Fatalf("%+v %v", res, err)
	}
}

func TestComplete_PinsProviders(t *testing.T) {
	srv := gateway(t, func(w http.ResponseWriter, r *http.Request) {
		gw, _ := decode(t, r)["providerOptions"].(map[string]any)["gateway"].(map[string]any)
		only, _ := gw["only"].([]any)
		if len(only) != 2 || only[0] != "a" || only[1] != "b" {
			t.Errorf("gateway options %v", gw)
		}
		writeCompletion(w, "{}", 1, 1, -1, "g")
	})
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret, Providers: []string{"a", "b"}})
	if _, err := c.Complete(context.Background(), userMsg); err != nil {
		t.Fatal(err)
	}
}

func TestNew_FailsFastNamingTheVariable(t *testing.T) {
	cases := []struct {
		cfg  Config
		name string
	}{
		{Config{Model: "m", APIKey: testSecret, PriceInPerMTok: 1, PriceOutPerMTok: 1}, "LLM_BASE_URL"},
		{Config{BaseURL: "http://x", APIKey: testSecret, PriceInPerMTok: 1, PriceOutPerMTok: 1}, "LLM_MODEL"},
		{Config{BaseURL: "http://x", Model: "m", PriceInPerMTok: 1, PriceOutPerMTok: 1}, "LLM_API_KEY"},
		{Config{BaseURL: "http://x", Model: "m", APIKey: testSecret, PriceOutPerMTok: 1}, "LLM_PRICE_IN_PER_MTOK"},
		{Config{BaseURL: "http://x", Model: "m", APIKey: testSecret, PriceInPerMTok: 1}, "LLM_PRICE_OUT_PER_MTOK"},
		{Config{BaseURL: "http://x", Model: "m", APIKey: testSecret, PriceInPerMTok: math.NaN(), PriceOutPerMTok: 1}, "LLM_PRICE_IN_PER_MTOK"},
	}
	for _, tc := range cases {
		_, err := New(tc.cfg, nil)
		if err == nil || !strings.Contains(err.Error(), tc.name) || strings.Contains(err.Error(), testKey) {
			t.Errorf("%s: %v", tc.name, err)
		}
	}
}

func TestNew_Defaults(t *testing.T) {
	c, err := New(Config{BaseURL: "http://x/", Model: "m", APIKey: testSecret, PriceInPerMTok: 1, PriceOutPerMTok: 1}, nil)
	if err != nil || c.cfg.Timeout != defaultTimeout || c.cfg.MaxRetries != defaultMaxRetries || c.cfg.BaseURL != "http://x" {
		t.Fatalf("%+v %v", c.cfg, err)
	}
	c, _ = New(Config{BaseURL: "http://x", Model: "m", APIKey: testSecret, MaxRetries: -1, PriceInPerMTok: 1, PriceOutPerMTok: 1}, nil)
	if c.cfg.MaxRetries != 0 {
		t.Fatalf("retries %d", c.cfg.MaxRetries)
	}
}

func TestComplete_StopsWhenContextCancelled(t *testing.T) {
	srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusServiceUnavailable) })
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret})
	ctx, cancel := context.WithCancel(context.Background())
	c.sleep = func(ctx context.Context, _ time.Duration) error { cancel(); return ctx.Err() }
	if _, err := c.Complete(ctx, userMsg); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}

func TestNoSecretInLogsOrErrors(t *testing.T) {
	var logs bytes.Buffer
	prev := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(&logs, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { slog.SetDefault(prev) })

	srv := gateway(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "bad key "+strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
	})
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret})
	_, err := c.Complete(context.Background(), userMsg)
	if err == nil {
		t.Fatal("want error")
	}
	secret := testKey
	if strings.Contains(err.Error(), secret) || strings.Contains(logs.String(), secret) {
		t.Fatalf("secret leaked: %q / %q", err, logs.String())
	}
	if !strings.Contains(logs.String(), "attempts=1") || !strings.Contains(logs.String(), "model=m") {
		t.Fatalf("expected a structured failure line, got %q", logs.String())
	}
	ok := gateway(t, func(w http.ResponseWriter, _ *http.Request) { writeCompletion(w, "{}", 10, 5, 0.01, "gen_1") })
	c2, _ := mustNew(t, Config{BaseURL: ok.URL, Model: "m", APIKey: testSecret})
	logs.Reset()
	if _, err := c2.Complete(context.Background(), userMsg); err != nil {
		t.Fatal(err)
	}
	line := logs.String()
	for _, want := range []string{"model=deepseek/test", "attempts=1", "prompt_tokens=10", "completion_tokens=5"} {
		if !strings.Contains(line, want) {
			t.Errorf("log line lacks %s: %q", want, line)
		}
	}
	if strings.Contains(line, secret) {
		t.Fatalf("secret in log: %q", line)
	}
	_, nerr := New(Config{BaseURL: "x", APIKey: testSecret, PriceInPerMTok: 1, PriceOutPerMTok: 1}, nil)
	if nerr == nil || strings.Contains(nerr.Error(), secret) {
		t.Fatalf("New leaked: %v", nerr)
	}
}

func TestNoSecretInTransportErrors(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	url := srv.URL
	srv.Close()
	c, _ := mustNew(t, Config{BaseURL: url, Model: "m", APIKey: testSecret})
	_, err := c.Complete(context.Background(), userMsg)
	if err == nil || strings.Contains(err.Error(), testKey) || strings.Contains(err.Error(), url) {
		t.Fatalf("%v", err)
	}
}
