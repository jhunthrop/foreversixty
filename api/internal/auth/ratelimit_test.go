// api/internal/auth/ratelimit_test.go
package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccountRateLimitKeyKeysBySignedInUserID(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/v1/sims/run", nil)
	r = r.WithContext(WithActor(r.Context(), Actor{UserID: 42, Role: "user", Method: "session"}))
	key, ok := AccountRateLimitKey(r)
	if !ok || key != "42" {
		t.Fatalf("AccountRateLimitKey = %q, %v, want \"42\", true", key, ok)
	}
}

func TestAccountRateLimitKeyExemptsAnUnsignedRequest(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/v1/sims/run", nil)
	key, ok := AccountRateLimitKey(r)
	if ok || key != "" {
		t.Fatalf("AccountRateLimitKey = %q, %v, want \"\", false for an unsigned request", key, ok)
	}
}

func TestAccountRateLimitKeyGivesDifferentAccountsDifferentKeys(t *testing.T) {
	a := httptest.NewRequest(http.MethodPost, "/v1/sims/run", nil)
	a = a.WithContext(WithActor(a.Context(), Actor{UserID: 1, Role: "user", Method: "session"}))
	b := httptest.NewRequest(http.MethodPost, "/v1/sims/run", nil)
	b = b.WithContext(WithActor(b.Context(), Actor{UserID: 2, Role: "user", Method: "session"}))
	keyA, _ := AccountRateLimitKey(a)
	keyB, _ := AccountRateLimitKey(b)
	if keyA == keyB {
		t.Fatalf("accounts 1 and 2 got the same key %q", keyA)
	}
}
