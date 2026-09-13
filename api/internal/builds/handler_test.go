package builds

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// fakeStore is an in-memory Storer: the handlers' contract with the store is
// small enough that a map is a better test double than a live Postgres.
type fakeStore struct {
	rows  map[string]Build
	saves int
	err   error
}

func newFakeStore() *fakeStore { return &fakeStore{rows: map[string]Build{}} }

func (f *fakeStore) Save(_ context.Context, b Build) (Build, bool, error) {
	if f.err != nil {
		return Build{}, false, f.err
	}
	f.saves++
	if existing, ok := f.rows[b.ID]; ok {
		return existing, false, nil
	}
	f.rows[b.ID] = b
	return b, true, nil
}

func (f *fakeStore) Get(_ context.Context, id string) (Build, error) {
	if f.err != nil {
		return Build{}, f.err
	}
	b, ok := f.rows[id]
	if !ok {
		return Build{}, ErrNotFound
	}
	return b, nil
}

func testRouter(t *testing.T, store Storer) http.Handler {
	t.Helper()
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, &Service{Store: store, Data: data, PublicBaseURL: "https://foreversixty.gg"}, 1)
	return mux
}

func postBuild(t *testing.T, h http.Handler, body string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/builds", bytes.NewBufferString(body))
	req.RemoteAddr = "10.0.0.1:1234"
	h.ServeHTTP(rec, req)
	return rec
}

const validBody = `{"class_id":1,"race_id":1,"tree_version":"test-1","point_order":[101,101,101,101,101,201],"title":"Arms leveling"}`

func TestSaveReturns201ThenDedupesTo200(t *testing.T) {
	h := testRouter(t, newFakeStore())

	rec := postBuild(t, h, validBody)
	if rec.Code != http.StatusCreated {
		t.Fatalf("code = %d body = %s", rec.Code, rec.Body.String())
	}
	var env struct {
		OK   bool              `json:"ok"`
		Data map[string]string `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if !env.OK || env.Data["id"] != "znorjmts" || env.Data["url"] != "https://foreversixty.gg/b/znorjmts" {
		t.Fatalf("data = %v", env.Data)
	}

	rec = postBuild(t, h, validBody)
	if rec.Code != http.StatusOK {
		t.Fatalf("second save: code = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"id":"znorjmts"`) {
		t.Fatalf("second save body = %s", rec.Body.String())
	}
}

func TestSaveReturns400WithPerFieldMessages(t *testing.T) {
	h := testRouter(t, newFakeStore())
	rec := postBuild(t, h, `{"class_id":1,"race_id":1,"tree_version":"test-1","point_order":[104]}`)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("code = %d", rec.Code)
	}
	var env struct {
		Error struct {
			Code   string            `json:"code"`
			Fields map[string]string `json:"fields"`
		} `json:"error"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if env.Error.Code != "invalid" {
		t.Fatalf("code = %q", env.Error.Code)
	}
	if got := env.Error.Fields["point_order[0]"]; got != "Tier 2 of Arms needs 10 points in Arms first" {
		t.Fatalf("fields = %v", env.Error.Fields)
	}
}

func TestSaveRejectsMalformedAndOversizedBodies(t *testing.T) {
	h := testRouter(t, newFakeStore())

	if rec := postBuild(t, h, `not json`); rec.Code != http.StatusBadRequest {
		t.Fatalf("malformed body: code = %d", rec.Code)
	}

	// 8 KB is the cap; a body past it must be refused, not parsed.
	big := `{"class_id":1,"race_id":1,"tree_version":"test-1","title":"` + strings.Repeat("x", 9000) + `"}`
	rec := postBuild(t, h, big)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("oversized body: code = %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "8 KB") {
		t.Fatalf("oversized body message = %s", rec.Body.String())
	}
}

func TestSaveRateLimitsAtTwentyPerIP(t *testing.T) {
	h := testRouter(t, newFakeStore())
	for i := 0; i < savesPerHour; i++ {
		if rec := postBuild(t, h, validBody); rec.Code == http.StatusTooManyRequests {
			t.Fatalf("save %d was rate limited before the budget ran out", i+1)
		}
	}
	if rec := postBuild(t, h, validBody); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("save %d: code = %d, want 429", savesPerHour+1, rec.Code)
	}
}

func TestFetchReturnsTheRecordAndCacheHeader(t *testing.T) {
	store := newFakeStore()
	h := testRouter(t, store)
	postBuild(t, h, validBody)

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/v1/builds/znorjmts", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("code = %d body = %s", rec.Code, rec.Body.String())
	}
	if got := rec.Header().Get("Cache-Control"); got != "public, max-age=86400" {
		t.Fatalf("cache-control = %q", got)
	}
	var env struct {
		Data Build `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if env.Data.ID != "znorjmts" || env.Data.Title != "Arms leveling" || len(env.Data.PointOrder) != 6 {
		t.Fatalf("record = %+v", env.Data)
	}
}

func TestFetchReturns404InTheEnvelope(t *testing.T) {
	h := testRouter(t, newFakeStore())
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/v1/builds/nosuchid", nil))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("code = %d", rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, `"ok":false`) || !strings.Contains(body, `"code":"not_found"`) {
		t.Fatalf("body = %s", body)
	}
}
