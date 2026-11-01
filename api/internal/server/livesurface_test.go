package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/livesurface"
)

// The live-surface routes ship dark: unmounted (the catch-all 404) until APP_WIDGETS builds the
// service, then behind a session (401 for an anonymous call).
func TestLiveSurfaceRoutesAreDarkUntilSwitchedOn(t *testing.T) {
	cases := []struct {
		name string
		deps Deps
		want int
	}{
		{"off", Deps{Version: "test-1"}, http.StatusNotFound},
		{"on", Deps{Version: "test-1", LiveSurface: &livesurface.Service{}}, http.StatusUnauthorized},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			h := NewRouter(c.deps)
			for _, r := range []*http.Request{
				httptest.NewRequest(http.MethodGet, "/v1/me/widget", nil),
				httptest.NewRequest(http.MethodPost, "/v1/me/live-activities", nil),
				httptest.NewRequest(http.MethodDelete, "/v1/me/live-activities/4", nil),
			} {
				rec := httptest.NewRecorder()
				h.ServeHTTP(rec, r)
				if rec.Code != c.want {
					t.Errorf("%s %s = %d, want %d", r.Method, r.URL.Path, rec.Code, c.want)
				}
			}
		})
	}
}
