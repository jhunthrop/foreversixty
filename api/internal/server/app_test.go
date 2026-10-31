package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/app"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

// The app routes land dark: unmounted without an app service, and mounted beside the
// character page's routes (no pattern conflict) with one.
func TestTheAppRoutesAreDarkUntilMounted(t *testing.T) {
	paths := []string{"/v1/app/config", "/v1/characters/us/normal/velaren/home"}
	dark := NewRouter(Deps{Version: "test"})
	for _, p := range paths {
		rec := httptest.NewRecorder()
		dark.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, p, nil))
		if rec.Code != http.StatusNotFound {
			t.Errorf("%s answered %d while dark", p, rec.Code)
		}
	}
	cards := &rankings.Service{}
	lit := NewRouter(Deps{Version: "test", Rankings: cards, App: &app.Service{Characters: cards}})
	rec := httptest.NewRecorder()
	// An unknown region is the handler's own 404 envelope, not the mux's bare one.
	lit.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/v1/characters/xx/normal/velaren/home", nil))
	if rec.Code != http.StatusNotFound || !strings.Contains(rec.Body.String(), `"code":"not_found"`) {
		t.Fatalf("the home route is not mounted: %d %s", rec.Code, rec.Body)
	}
}
