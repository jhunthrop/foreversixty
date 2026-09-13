package config

import "testing"

func env(m map[string]string) func(string) string {
	return func(k string) string { return m[k] }
}

func TestLoadRequiresDatabaseURL(t *testing.T) {
	_, err := Load(env(map[string]string{"PORT": "8080"}))
	if err == nil {
		t.Fatal("expected error for missing DATABASE_URL")
	}
}

func TestLoadDefaults(t *testing.T) {
	c, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
	}))
	if err != nil {
		t.Fatal(err)
	}
	if c.Port != "8080" {
		t.Fatalf("default port = %q", c.Port)
	}
}
