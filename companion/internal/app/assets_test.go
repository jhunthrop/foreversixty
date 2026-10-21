package app_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// The window's typefaces and seal come from the same loopback server as
// its page, under the session token. A font served as anything but
// font/woff2 is refused by the webview's preload, and the page then shows
// Georgia and Arial.
func TestTheTypefacesAndTheSealAreServedWithTheirTypes(t *testing.T) {
	h := newHarness(t)
	for path, want := range map[string]string{
		"/fonts/cinzel-700-subset.woff2":         "font/woff2",
		"/fonts/barlow-400-subset.woff2":         "font/woff2",
		"/fonts/barlow-600-subset.woff2":         "font/woff2",
		"/fonts/jetbrains-mono-500-subset.woff2": "font/woff2",
		"/seal.svg":                              "image/svg+xml",
		"/styles.css":                            "text/css",
	} {
		w := httptest.NewRecorder()
		h.app.Handler().ServeHTTP(w, httptest.NewRequest(http.MethodGet, h.url(path), nil))
		if w.Code != http.StatusOK || !strings.HasPrefix(w.Header().Get("Content-Type"), want) || w.Body.Len() == 0 {
			t.Errorf("%s: %d %q (%d bytes), want 200 %s", path, w.Code, w.Header().Get("Content-Type"), w.Body.Len(), want)
		}
	}
}
