package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/push"
)

// The app notification routes ship dark: unmounted unless their handlers are wired.
func TestAppNotificationRoutesAreUnmountedByDefault(t *testing.T) {
	routes := []struct{ method, path string }{
		{http.MethodGet, "/v1/me/notifications"},
		{http.MethodPatch, "/v1/me/notifications"},
		{http.MethodPost, "/v1/me/push-tokens"},
		{http.MethodDelete, "/v1/me/push-tokens/abc"},
	}
	dark := NewRouter(Deps{})
	lit := NewRouter(Deps{Notifications: &events.Handler{}, PushTokens: &push.Handler{}})
	for _, r := range routes {
		rec := httptest.NewRecorder()
		dark.ServeHTTP(rec, httptest.NewRequest(r.method, r.path, nil))
		if rec.Code != http.StatusNotFound {
			t.Errorf("dark %s %s = %d", r.method, r.path, rec.Code)
		}
		rec = httptest.NewRecorder()
		lit.ServeHTTP(rec, httptest.NewRequest(r.method, r.path, nil))
		if rec.Code != http.StatusUnauthorized {
			t.Errorf("lit anonymous %s %s = %d", r.method, r.path, rec.Code)
		}
	}
}
