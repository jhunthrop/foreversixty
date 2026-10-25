package config

import (
	"strings"
	"testing"
)

func TestLoadDiscordApplicationIsAPair(t *testing.T) {
	base := map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
	}
	with := func(extra map[string]string) func(string) string {
		m := map[string]string{}
		for k, v := range base {
			m[k] = v
		}
		for k, v := range extra {
			m[k] = v
		}
		return env(m)
	}
	c, err := Load(with(nil))
	if err != nil || c.DiscordClientID != "" || c.DiscordClientSecret != "" {
		t.Fatalf("neither: %+v, %v", c, err)
	}
	c, err = Load(with(map[string]string{"DISCORD_CLIENT_ID": "123", "DISCORD_CLIENT_SECRET": "s3cret"}))
	if err != nil || c.DiscordClientID != "123" || c.DiscordClientSecret != "s3cret" {
		t.Fatalf("both: %+v, %v", c, err)
	}
	for variable, other := range map[string]string{"DISCORD_CLIENT_ID": "DISCORD_CLIENT_SECRET", "DISCORD_CLIENT_SECRET": "DISCORD_CLIENT_ID"} {
		_, err := Load(with(map[string]string{variable: "valueXYZ"}))
		if err == nil || !strings.Contains(err.Error(), other) || strings.Contains(err.Error(), "valueXYZ") {
			t.Fatalf("only %s: err = %v", variable, err)
		}
	}
}
