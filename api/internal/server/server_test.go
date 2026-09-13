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
		{"/health", `"status":"ok"`},
		{"/version", `"version":"test-1"`},
	} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, tc.path, nil))
		if rec.Code != 200 || !strings.Contains(rec.Body.String(), tc.want) {
			t.Fatalf("%s: code=%d body=%s", tc.path, rec.Code, rec.Body.String())
		}
	}
}

func TestUnknownRouteReturnsEnvelope(t *testing.T) {
	h := NewRouter(Deps{Version: "test-1"})
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/nope", nil))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("code = %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.HasPrefix(ct, "application/json") {
		t.Fatalf("content-type = %q", ct)
	}
	body := rec.Body.String()
	if !strings.Contains(body, `"ok":false`) || !strings.Contains(body, `"code":"not_found"`) {
		t.Fatalf("unexpected body: %s", body)
	}
}
