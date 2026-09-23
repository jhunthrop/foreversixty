package httpx

import (
	"net/http/httptest"
	"testing"
	"time"
)

func TestCachePublic(t *testing.T) {
	cases := []struct {
		name                 string
		maxAge, staleRevalid time.Duration
		want                 string
	}{
		{"no swr", 3600 * time.Second, 0, "public, max-age=3600"},
		{"with swr", 30 * time.Second, 300 * time.Second, "public, max-age=30, stale-while-revalidate=300"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			CachePublic(rec, tc.maxAge, tc.staleRevalid)
			if got := rec.Header().Get("Cache-Control"); got != tc.want {
				t.Errorf("Cache-Control = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestCachePrivate(t *testing.T) {
	rec := httptest.NewRecorder()
	CachePrivate(rec)
	if got := rec.Header().Get("Cache-Control"); got != "private, no-cache" {
		t.Errorf("Cache-Control = %q", got)
	}
	if got := rec.Header().Get("Vary"); got != "Cookie, Authorization" {
		t.Errorf("Vary = %q", got)
	}
}

func TestSetPrivateListCacheDelegatesToCachePrivate(t *testing.T) {
	rec := httptest.NewRecorder()
	SetPrivateListCache(rec)
	if got := rec.Header().Get("Cache-Control"); got != "private, no-cache" {
		t.Errorf("Cache-Control = %q", got)
	}
	if got := rec.Header().Get("Vary"); got != "Cookie, Authorization" {
		t.Errorf("Vary = %q", got)
	}
}
