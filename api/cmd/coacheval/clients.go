package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

const (
	envKey       = "LLM_API_KEY"
	envBaseURL   = "LLM_BASE_URL"
	envPriceIn   = "LLM_PRICE_IN_PER_MTOK"
	envPriceOut  = "LLM_PRICE_OUT_PER_MTOK"
	envProviders = "LLM_PROVIDERS"

	fakeTokensIn  = 3000
	fakeTokensOut = 600
	fakeCostUSD   = 0.0005
)

// gatewayClients builds one HTTP client per model from the environment and the options.
// The key is read here, handed to llm.New wrapped in a Secret and never printed.
func gatewayClients(o Options, getenv func(string) string) (map[string]llm.Client, error) {
	key := strings.TrimSpace(getenv(envKey))
	if key == "" {
		return nil, fmt.Errorf("%s is not set", envKey)
	}
	base := getenv(envBaseURL)
	if base == "" {
		base = defaultGatewayURL
	}
	clients := make(map[string]llm.Client, len(o.Models))
	for _, slug := range o.Models {
		p, err := priceFor(o, slug, getenv)
		if err != nil {
			return nil, err
		}
		c, err := llm.New(llm.Config{BaseURL: base, Model: slug, APIKey: llm.NewSecret(key),
			Providers: providersFor(o, slug, getenv), PriceInPerMTok: p.In, PriceOutPerMTok: p.Out,
			Reasoning: o.ReasoningOf(slug)}, nil)
		if err != nil {
			return nil, fmt.Errorf("model %s: %w", slug, err)
		}
		clients[slug] = c
	}
	return clients, nil
}

func priceFor(o Options, slug string, getenv func(string) string) (price, error) {
	if p, ok := o.Prices[slug]; ok {
		return p, nil
	}
	in, err1 := strconv.ParseFloat(getenv(envPriceIn), 64)
	out, err2 := strconv.ParseFloat(getenv(envPriceOut), 64)
	if err1 != nil || err2 != nil {
		return price{}, fmt.Errorf("model %s needs -price %s=in,out or %s and %s", slug, slug, envPriceIn, envPriceOut)
	}
	return price{In: in, Out: out}, nil
}

func providersFor(o Options, slug string, getenv func(string) string) []string {
	if p, ok := o.Providers[slug]; ok {
		return p
	}
	var out []string
	for _, p := range strings.Split(getenv(envProviders), ",") {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

// sharedFactory answers every pack with the model's shared client.
func sharedFactory(clients map[string]llm.Client) ClientFactory {
	return func(m ModelSpec, _ Built) (llm.Client, error) {
		c, ok := clients[m.Slug]
		if !ok {
			return nil, fmt.Errorf("no client for %s", m.Slug)
		}
		return c, nil
	}
}

// fakeFactory scripts an llm.Fake per pack with the rules rendering as the model's answer,
// so a smoke run exercises build, prompt, parse, validate and the writers without a gateway.
func fakeFactory(m ModelSpec, b Built) (llm.Client, error) {
	doc, err := json.Marshal(coach.RenderRules(b.Pack, b.Persona.Drills))
	if err != nil {
		return nil, fmt.Errorf("fake answer: %w", err)
	}
	resp := llm.Response{Content: string(doc), Model: m.Slug, CostUSD: fakeCostUSD, CostEstimated: true,
		Usage: llm.Usage{PromptTokens: fakeTokensIn, CompletionTokens: fakeTokensOut}}
	return &llm.Fake{Responses: []llm.Response{resp, resp}}, nil
}
