package builds

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

func serveGet(h http.Handler, path string, userID int64) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	if userID != 0 {
		req = signedInAs(req, userID)
	}
	h.ServeHTTP(rec, req)
	return rec
}

func assertRecorder(t *testing.T, name string, rec *httptest.ResponseRecorder, status int, opts contract.Options) {
	t.Helper()
	if rec.Code != status {
		t.Fatalf("%s = %d, want %d: %s", name, rec.Code, status, rec.Body.String())
	}
	contract.AssertFixture(t, name, rec.Body.Bytes(), opts)
}

var (
	buildOptional = contract.Options{Optional: []string{"data.title"}}
	buildOpen     = contract.Options{Optional: []string{"data.title"}, Open: []string{"data.gear"}}
	pageOptional  = contract.Options{
		Optional: []string{"data.rows[].title"}, Open: []string{"data.rows[].gear"},
	}
)

func TestAppContractBuildSaveAndFetch(t *testing.T) {
	h := testRouter(t, newFakeStore())

	assertRecorder(t, "builds.save", postBuild(t, h, validBody), http.StatusCreated, contract.Options{})
	assertRecorder(t, "builds.save.error-invalid",
		postBuild(t, h, `{"class_id":1,"race_id":1,"tree_version":"test-1","point_order":[999]}`),
		http.StatusBadRequest, contract.Options{})

	rec := serveGet(h, "/v1/builds/znorjmts", 0)
	if got, want := rec.Header().Get("Cache-Control"), "public, max-age=86400"; got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
	assertRecorder(t, "builds.get", rec, http.StatusOK, buildOpen)
	assertRecorder(t, "builds.get.error-not-found", serveGet(h, "/v1/builds/aaaaaaaa", 0), http.StatusNotFound, contract.Options{})
}

func TestAppContractBuildList(t *testing.T) {
	store := newFakeStore()
	store.seed(Build{
		ID: "znorjmts", ClassID: 1, RaceID: 1, TreeVersion: "test-1", PointOrder: []int{101, 101, 201},
		Gear: map[string]int{"head": 12640}, Title: "Arms leveling",
		CreatedAt: time.Date(2026, 10, 28, 19, 0, 0, 0, time.UTC), Views: 3,
	}, 1)
	h := testRouter(t, store)

	rec := serveGet(h, "/v1/builds?mine=1", 1)
	if got, want := rec.Header().Get("Cache-Control"), "private, no-cache"; got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
	assertRecorder(t, "builds.list", rec, http.StatusOK, pageOptional)
	assertRecorder(t, "builds.list.empty", serveGet(h, "/v1/builds?mine=1", 2), http.StatusOK, pageOptional)
	assertRecorder(t, "builds.list.error-invalid", serveGet(h, "/v1/builds", 1), http.StatusBadRequest, contract.Options{})
}
