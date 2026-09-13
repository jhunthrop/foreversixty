package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthAndVersion(t *testing.T) {
	h := NewRouter(Deps{Version: "test-1"})
	for _, tc := range []struct{ path, want string }{
		{"/healthz", `"status":"ok"`},
		{"/version", `"version":"test-1"`},
	} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, tc.path, nil))
		if rec.Code != 200 || !strings.Contains(rec.Body.String(), tc.want) {
			t.Fatalf("%s: code=%d body=%s", tc.path, rec.Code, rec.Body.String())
		}
	}
}
