package notify

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGuildOfficerGateIsExported(t *testing.T) {
	g := newGuildAPI(t, officerWithPlan())
	mux := http.NewServeMux()
	mux.HandleFunc("GET /gate/{id}", func(w http.ResponseWriter, r *http.Request) {
		if _, _, ok := GuildOfficerGate(w, r, g.access, true, func(string, error) {}); ok {
			w.WriteHeader(http.StatusNoContent)
		}
	})
	for path, want := range map[string]int{
		"/gate/" + strconv.FormatInt(g.gid, 10): http.StatusNoContent,
		"/gate/zero":                            http.StatusNotFound,
	} {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
		if rec.Code != want {
			t.Fatalf("%s: status %d, want %d", path, rec.Code, want)
		}
	}
}
