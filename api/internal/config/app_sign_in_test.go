package config

import (
	"testing"
	"time"
)

// The native app sign-in lands dark: off unless APP_SIGN_IN is exactly true.
func TestAppSignInIsOffUnlessExactlyTrue(t *testing.T) {
	base := map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "http://localhost:4321", "API_BASE_URL": "http://localhost:8080",
	}
	for v, want := range map[string]bool{"": false, "0": false, "1": false, "yes": false, "TRUE": false, "true": true} {
		m := map[string]string{"APP_SIGN_IN": v}
		for k, val := range base {
			m[k] = val
		}
		c, err := Load(env(m))
		if err != nil {
			t.Fatalf("APP_SIGN_IN=%q: %v", v, err)
		}
		if c.AppSignIn != want {
			t.Errorf("APP_SIGN_IN=%q: AppSignIn = %v, want %v", v, c.AppSignIn, want)
		}
	}
}

func appEnv(extra map[string]string) func(string) string {
	m := map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "http://localhost:4321", "API_BASE_URL": "http://localhost:8080",
	}
	for k, v := range extra {
		m[k] = v
	}
	return env(m)
}

// An app token lapses after 30 idle days and 180 days of age unless the deployment says otherwise.
func TestAppTokenLifetimesDefaultAndOverride(t *testing.T) {
	c, err := Load(appEnv(nil))
	if err != nil {
		t.Fatal(err)
	}
	if c.AppTokenIdle != 30*24*time.Hour || c.AppTokenMaxAge != 180*24*time.Hour {
		t.Fatalf("defaults = %s idle, %s max age", c.AppTokenIdle, c.AppTokenMaxAge)
	}
	c, err = Load(appEnv(map[string]string{"APP_TOKEN_IDLE_DAYS": "7", "APP_TOKEN_MAX_AGE_DAYS": "90"}))
	if err != nil {
		t.Fatal(err)
	}
	if c.AppTokenIdle != 7*24*time.Hour || c.AppTokenMaxAge != 90*24*time.Hour {
		t.Fatalf("overrides = %s idle, %s max age", c.AppTokenIdle, c.AppTokenMaxAge)
	}
}

func TestAppTokenLifetimesRefuseNonsense(t *testing.T) {
	for _, extra := range []map[string]string{
		{"APP_TOKEN_IDLE_DAYS": "0"}, {"APP_TOKEN_IDLE_DAYS": "-1"}, {"APP_TOKEN_IDLE_DAYS": "thirty"},
		{"APP_TOKEN_MAX_AGE_DAYS": "0"},
		// An idle limit longer than the maximum age could never apply.
		{"APP_TOKEN_IDLE_DAYS": "200", "APP_TOKEN_MAX_AGE_DAYS": "100"},
	} {
		if _, err := Load(appEnv(extra)); err == nil {
			t.Errorf("%v should be refused", extra)
		}
	}
}
