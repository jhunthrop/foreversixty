package notify

import (
	"errors"
	"testing"
)

func TestParseWebhookURL(t *testing.T) {
	accept := map[string]ParsedURL{
		"https://discord.com/api/webhooks/123/abc-DEF_9": {WebhookID: "123", Token: "abc-DEF_9"},
		"https://discordapp.com/api/webhooks/456/tok":    {WebhookID: "456", Token: "tok"},
		"https://discord.com/api/v6/webhooks/123/abc":    {WebhookID: "123", Token: "abc"},
		"https://discord.com/api/v7/webhooks/123/abc":    {WebhookID: "123", Token: "abc"},
		"https://discord.com/api/v8/webhooks/123/abc":    {WebhookID: "123", Token: "abc"},
		"https://discord.com/api/v9/webhooks/123/abc":    {WebhookID: "123", Token: "abc"},
		"https://discord.com/api/v10/webhooks/123/abc":   {WebhookID: "123", Token: "abc"},
		"  HTTPS://DISCORD.COM/api/webhooks/9/t \n":      {WebhookID: "9", Token: "t"},
	}
	for raw, want := range accept {
		got, err := ParseWebhookURL(raw)
		if err != nil || got != want {
			t.Errorf("ParseWebhookURL(%q) = %+v, %v; want %+v", raw, got, err, want)
		}
	}
	refuse := []string{
		"http://discord.com/api/webhooks/123/abc",
		"https://example.com/api/webhooks/123/abc",
		"https://ptb.discord.com/api/webhooks/123/abc",
		"https://www.discord.com/api/webhooks/123/abc",
		"https://discord.com.evil.test/api/webhooks/123/abc",
		"https://discord.com/api/webhooks/123/",
		"https://discord.com/api/webhooks/123",
		"https://discord.com/api/webhooks/123/abc?wait=true",
		"https://discord.com/api/webhooks/123/abc#x",
		"https://discord.com/api/webhooks/123/abc/extra",
		"https://discord.com/api/webhooks/12x/abc",
		"https://user:pw@discord.com/api/webhooks/123/abc",
		"https://discord.com:8443/api/webhooks/123/abc",
		"https://discord.com/api/v5/webhooks/123/abc",
		"https://discord.com/api/v11/webhooks/123/abc",
		"https://discord.com/api/v1x/webhooks/123/abc",
		"https://discord.com/api/v10/webhooks/123/abc/extra",
		"",
		"not a url",
	}
	for _, raw := range refuse {
		if _, err := ParseWebhookURL(raw); !errors.Is(err, ErrNotAWebhook) {
			t.Errorf("ParseWebhookURL(%q) err = %v; want ErrNotAWebhook", raw, err)
		}
	}
}

func TestParsedURLRoundTrips(t *testing.T) {
	const raw = "https://discord.com/api/webhooks/123/abc"
	p, err := ParseWebhookURL(raw)
	if err != nil || p.URL() != raw {
		t.Fatalf("URL() = %q, %v; want %q", p.URL(), err, raw)
	}
}
