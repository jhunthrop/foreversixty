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
	if c.MailFrom != defaultMailFrom {
		t.Fatalf("default MailFrom = %q", c.MailFrom)
	}
	if c.TrustedProxyHops != 1 {
		t.Fatalf("default TrustedProxyHops = %d, want 1", c.TrustedProxyHops)
	}
	if c.MigrateDatabaseURL != c.DatabaseURL {
		t.Fatalf("default MigrateDatabaseURL = %q, want it to fall back to DatabaseURL %q", c.MigrateDatabaseURL, c.DatabaseURL)
	}
}

func TestLoadMailFromOverride(t *testing.T) {
	c, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
		"MAIL_FROM": "Test <test@example.com>",
	}))
	if err != nil {
		t.Fatal(err)
	}
	if c.MailFrom != "Test <test@example.com>" {
		t.Fatalf("MailFrom = %q", c.MailFrom)
	}
}

func TestLoadMigrateDatabaseURLOverride(t *testing.T) {
	c, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://pooled", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
		"MIGRATE_DATABASE_URL": "postgres://direct",
	}))
	if err != nil {
		t.Fatal(err)
	}
	if c.MigrateDatabaseURL != "postgres://direct" {
		t.Fatalf("MigrateDatabaseURL = %q, want postgres://direct", c.MigrateDatabaseURL)
	}
	if c.DatabaseURL != "postgres://pooled" {
		t.Fatalf("DatabaseURL should be unaffected, got %q", c.DatabaseURL)
	}
}

func TestLoadTrustedProxyHopsOverride(t *testing.T) {
	c, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
		"TRUSTED_PROXY_HOPS": "0",
	}))
	if err != nil {
		t.Fatal(err)
	}
	if c.TrustedProxyHops != 0 {
		t.Fatalf("TrustedProxyHops = %d, want 0", c.TrustedProxyHops)
	}
}

func TestLoadTrustedProxyHopsRejectsNonInteger(t *testing.T) {
	_, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
		"TRUSTED_PROXY_HOPS": "nope",
	}))
	if err == nil {
		t.Fatal("expected error for non-integer TRUSTED_PROXY_HOPS")
	}
}

func TestLoadTrustedProxyHopsRejectsNegative(t *testing.T) {
	_, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
		"TRUSTED_PROXY_HOPS": "-1",
	}))
	if err == nil {
		t.Fatal("expected error for negative TRUSTED_PROXY_HOPS")
	}
}
