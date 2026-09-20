package builds

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// fakeStore is an in-memory Storer: the handlers' contract with the store is
// small enough that a map is a better test double than a live Postgres.
type fakeStore struct {
	rows map[string]Build
	// owners is who saved each row, the column 0015 added. Mine honours
	// it, because that handoff — the handler's actor reaching the
	// store's filter — is the whole authorization surface of the list.
	owners map[string]*int64
	saves  int
	err    error
}

func newFakeStore() *fakeStore {
	return &fakeStore{rows: map[string]Build{}, owners: map[string]*int64{}}
}

// seed puts a row in the store already owned, so a list test does not
// have to post one build per user through the save route.
func (f *fakeStore) seed(b Build, owner int64) {
	f.rows[b.ID] = b
	f.owners[b.ID] = &owner
}

func (f *fakeStore) Save(_ context.Context, b Build, owner *int64) (Build, bool, error) {
	if f.err != nil {
		return Build{}, false, f.err
	}
	f.saves++
	if existing, ok := f.rows[b.ID]; ok {
		return existing, false, nil
	}
	f.rows[b.ID] = b
	f.owners[b.ID] = owner
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

func (f *fakeStore) Mine(_ context.Context, userID int64, page int) (Page, error) {
	if f.err != nil {
		return Page{}, f.err
	}
	out := Page{Rows: []Build{}, Page: clampPage(page), PerPage: PerPage}
	for id, b := range f.rows {
		if owner := f.owners[id]; owner != nil && *owner == userID {
			out.Rows = append(out.Rows, b)
		}
	}
	slices.SortFunc(out.Rows, func(a, b Build) int { return strings.Compare(a.ID, b.ID) })
	out.Total = len(out.Rows)
	return out, nil
}

func testRouter(t *testing.T, store Storer) http.Handler {
	t.Helper()
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, &Service{
		Store:         store,
		Data:          data,
		PublicBaseURL: "https://foreversixty.gg",
		Log:           slog.New(slog.NewTextHandler(io.Discard, nil)),
	}, 1)
	return mux
}

// signedIn wraps r with a signed-in actor, the way api/internal/sims's
// harness carries identity through its test server.
func signedIn(r *http.Request) *http.Request { return signedInAs(r, 1) }

// signedInAs is signedIn for a named account, which is what the list
// test needs: two accounts, one list each.
func signedInAs(r *http.Request, userID int64) *http.Request {
	return r.WithContext(auth.WithActor(r.Context(),
		auth.Actor{UserID: userID, Role: "user", Method: "session"}))
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

// The two tests below drive the handlers' 500 branches through fakeStore's
// err field, which is the only way either branch is reached: every other
// path either succeeds or reports ErrNotFound.

func TestSaveReports500WhenTheStoreFails(t *testing.T) {
	store := newFakeStore()
	store.err = errors.New("database is down")
	h := testRouter(t, store)

	rec := postBuild(t, h, validBody)
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("code = %d body = %s", rec.Code, rec.Body.String())
	}
	if code := errorCode(t, rec); code != "internal" {
		t.Fatalf("error.code = %q, want internal", code)
	}
	if store.saves != 0 {
		t.Fatalf("saves = %d, want 0: a failing store stores nothing", store.saves)
	}
}

func TestFetchReports500WhenTheStoreFails(t *testing.T) {
	store := newFakeStore()
	store.err = errors.New("database is down")
	h := testRouter(t, store)

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/v1/builds/znorjmts", nil))
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("code = %d body = %s", rec.Code, rec.Body.String())
	}
	if code := errorCode(t, rec); code != "internal" {
		t.Fatalf("error.code = %q, want internal", code)
	}
}

// errorCode reads the envelope's error.code.
func errorCode(t *testing.T, rec *httptest.ResponseRecorder) string {
	t.Helper()
	var env struct {
		OK    bool `json:"ok"`
		Error struct {
			Code string `json:"code"`
		} `json:"error"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatalf("body is not an envelope: %v (%s)", err, rec.Body.String())
	}
	if env.OK {
		t.Fatalf("ok = true on an error response: %s", rec.Body.String())
	}
	return env.Error.Code
}

func TestMyOwnBuildsNeedASessionAndMineEqualsOne(t *testing.T) {
	// The handler tests run against fakeStore, so this one checks the
	// route's shape: the parameter is required and the session is.
	store := &fakeStore{}
	h := testRouter(t, store)

	w := httptest.NewRecorder()
	h.ServeHTTP(w, signedIn(httptest.NewRequest(http.MethodGet, "/v1/builds", nil)))
	if w.Code != http.StatusBadRequest {
		t.Errorf("status %d, want 400 without mine=1", w.Code)
	}

	w = httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/v1/builds?mine=1", nil))
	if w.Code != http.StatusUnauthorized {
		t.Errorf("status %d, want 401 without a session", w.Code)
	}
}

// TestMyOwnBuildsAreOnlyMine drives the handler's actor through to the
// store's owner filter. It is the builds side of what the sims package
// proves against Postgres: the list is an authorization surface, and
// the handoff from the signed-in actor to the store's user id is the
// part that has to be right.
func TestMyOwnBuildsAreOnlyMine(t *testing.T) {
	const mine, theirs = int64(7), int64(8)
	store := newFakeStore()
	store.seed(Build{ID: "aaaaaaaa", Title: "my arms"}, mine)
	store.seed(Build{ID: "bbbbbbbb", Title: "my fury"}, mine)
	store.seed(Build{ID: "cccccccc", Title: "their frost"}, theirs)
	h := testRouter(t, store)

	for _, c := range []struct {
		user int64
		want []string
	}{
		{mine, []string{"aaaaaaaa", "bbbbbbbb"}},
		{theirs, []string{"cccccccc"}},
	} {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, signedInAs(httptest.NewRequest(http.MethodGet, "/v1/builds?mine=1", nil), c.user))
		if w.Code != http.StatusOK {
			t.Fatalf("user %d: status %d, body %s", c.user, w.Code, w.Body.String())
		}
		if got := w.Header().Get("Cache-Control"); got != "private, no-store" {
			t.Errorf("user %d: cache-control = %q", c.user, got)
		}
		var env struct {
			OK   bool `json:"ok"`
			Data Page `json:"data"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &env); err != nil {
			t.Fatalf("user %d: %v (%s)", c.user, err, w.Body.String())
		}
		got := make([]string, 0, len(env.Data.Rows))
		for _, b := range env.Data.Rows {
			got = append(got, b.ID)
		}
		if !slices.Equal(got, c.want) {
			t.Errorf("user %d: rows = %v, want %v", c.user, got, c.want)
		}
		if env.Data.Total != len(c.want) {
			t.Errorf("user %d: total = %d, want %d", c.user, env.Data.Total, len(c.want))
		}
	}
}

// TestMyOwnBuildsClampAnAbsurdPage pins the overflow: (page-1)*PerPage
// past the int range would hand Postgres a negative OFFSET, which is a
// 500 where an empty page is the true answer.
func TestMyOwnBuildsClampAnAbsurdPage(t *testing.T) {
	store := newFakeStore()
	store.seed(Build{ID: "aaaaaaaa"}, 7)
	h := testRouter(t, store)

	w := httptest.NewRecorder()
	h.ServeHTTP(w, signedInAs(httptest.NewRequest(http.MethodGet,
		"/v1/builds?mine=1&page=9223372036854775807", nil), 7))
	if w.Code != http.StatusOK {
		t.Fatalf("status %d, body %s", w.Code, w.Body.String())
	}
	var env struct {
		Data Page `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if env.Data.Page != MaxPage {
		t.Errorf("page = %d, want it clamped to %d", env.Data.Page, MaxPage)
	}
}
