package config

import "testing"

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
