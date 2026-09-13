package config

import (
	"fmt"
	"strconv"
)

// defaultMailFrom is used when MAIL_FROM is not set.
const defaultMailFrom = "Forever Sixty <hello@foreversixty.gg>"

// defaultTrustedProxyHops is used when TRUSTED_PROXY_HOPS is not set. 1
// matches a single reverse proxy (e.g. Cloud Run) sitting directly in front
// of this service.
const defaultTrustedProxyHops = 1

type Config struct {
	Port     string
	MailFrom string

	// DatabaseURL is used for normal runtime connections; it may point at a
	// connection pooler (e.g. Neon's PgBouncer endpoint).
	DatabaseURL string
	// MigrateDatabaseURL is used only for running migrations, which need a
	// session-level advisory lock that a transaction-pooled connection
	// cannot hold. Falls back to DatabaseURL when unset, so a direct
	// (non-pooled) database needs no extra configuration.
	MigrateDatabaseURL string

	ResendAPIKey  string
	PublicBaseURL string
	APIBaseURL    string

	// TrustedProxyHops is how many reverse proxies in front of this service
	// are trusted to append to X-Forwarded-For; see httpx.RateLimit. 0
	// ignores X-Forwarded-For entirely and rate-limits by RemoteAddr.
	TrustedProxyHops int
}

func Load(getenv func(string) string) (Config, error) {
	c := Config{
		Port:               getenv("PORT"),
		MailFrom:           getenv("MAIL_FROM"),
		DatabaseURL:        getenv("DATABASE_URL"),
		MigrateDatabaseURL: getenv("MIGRATE_DATABASE_URL"),
		ResendAPIKey:       getenv("RESEND_API_KEY"),
		PublicBaseURL:      getenv("PUBLIC_BASE_URL"),
		APIBaseURL:         getenv("API_BASE_URL"),
	}
	if c.Port == "" {
		c.Port = "8080"
	}
	if c.MailFrom == "" {
		c.MailFrom = defaultMailFrom
	}
	for name, v := range map[string]string{
		"DATABASE_URL": c.DatabaseURL, "RESEND_API_KEY": c.ResendAPIKey,
		"PUBLIC_BASE_URL": c.PublicBaseURL, "API_BASE_URL": c.APIBaseURL,
	} {
		if v == "" {
			return Config{}, fmt.Errorf("config: %s is required", name)
		}
	}
	if c.MigrateDatabaseURL == "" {
		c.MigrateDatabaseURL = c.DatabaseURL
	}

	c.TrustedProxyHops = defaultTrustedProxyHops
	if v := getenv("TRUSTED_PROXY_HOPS"); v != "" {
		hops, err := strconv.Atoi(v)
		if err != nil || hops < 0 {
			return Config{}, fmt.Errorf("config: TRUSTED_PROXY_HOPS must be an integer >= 0, got %q", v)
		}
		c.TrustedProxyHops = hops
	}

	return c, nil
}
