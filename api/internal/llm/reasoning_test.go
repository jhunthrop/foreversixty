package llm

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func reasoningOf(t *testing.T, cfg Config) any {
	t.Helper()
	var got any
	srv := gateway(t, func(w http.ResponseWriter, r *http.Request) {
		got = decode(t, r)["reasoning"]
		writeCompletion(w, "{}", 1, 1, -1, "g")
	})
	cfg.BaseURL, cfg.Model, cfg.APIKey = srv.URL, "m", testSecret
	c, _ := mustNew(t, cfg)
	if _, err := c.Complete(context.Background(), userMsg); err != nil {
		t.Fatal(err)
	}
	return got
}

func TestComplete_ReasoningOffByDefault(t *testing.T) {
	got, _ := reasoningOf(t, Config{}).(map[string]any)
	if len(got) != 1 || got["enabled"] != false {
		t.Fatalf("reasoning %v", got)
	}
}

func TestComplete_ReasoningEffortLevels(t *testing.T) {
	for _, level := range []string{"low", "medium", "high"} {
		got, _ := reasoningOf(t, Config{Reasoning: level}).(map[string]any)
		if len(got) != 1 || got["effort"] != level {
			t.Fatalf("%s: reasoning %v", level, got)
		}
	}
}

func TestNew_RejectsUnknownReasoningNamingTheAllowedValues(t *testing.T) {
	_, err := New(Config{BaseURL: "http://x", Model: "m", APIKey: testSecret, PriceInPerMTok: 1, PriceOutPerMTok: 1, Reasoning: "xhigh"}, nil)
	if err == nil || !strings.Contains(err.Error(), "LLM_REASONING") {
		t.Fatalf("%v", err)
	}
	for _, level := range ReasoningLevels {
		if !strings.Contains(err.Error(), level) {
			t.Errorf("error does not name %q: %v", level, err)
		}
	}
}

func TestComplete_ReadsReasoningTokensAndIgnoresReasoningText(t *testing.T) {
	srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `{"id":"g","model":"m","usage":{"prompt_tokens":20,"completion_tokens":209,"completion_tokens_details":{"reasoning_tokens":164}},`+
			`"choices":[{"message":{"role":"assistant","content":"ok","reasoning":"PROMPT-QUOTE"}}]}`)
	})
	c, _ := mustNew(t, Config{BaseURL: srv.URL, Model: "m", APIKey: testSecret})
	res, err := c.Complete(context.Background(), userMsg)
	if err != nil || res.Usage.ReasoningTokens != 164 || res.Usage.CompletionTokens != 209 || strings.Contains(fmt.Sprintf("%+v", res), "PROMPT-QUOTE") {
		t.Fatalf("%+v %v", res, err)
	}
}
