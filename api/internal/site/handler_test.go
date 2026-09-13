package site

import (
	"context"
	"encoding/json"
	"errors"
	"html"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

type fakeGetter struct{ rows map[string]builds.Build }

func (f fakeGetter) Get(_ context.Context, id string) (builds.Build, error) {
	b, ok := f.rows[id]
	if !ok {
		return builds.Build{}, builds.ErrNotFound
	}
	return b, nil
}

// failingGetter stands in for a store that is down.
type failingGetter struct{ err error }

func (f failingGetter) Get(context.Context, string) (builds.Build, error) {
	return builds.Build{}, f.err
}

type recordingViews struct{ ids []string }

func (r *recordingViews) Record(id string) { r.ids = append(r.ids, id) }

func testPage(t *testing.T, record builds.Build, views ViewRecorder) http.Handler {
	t.Helper()
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, Deps{
		Store:         fakeGetter{rows: map[string]builds.Build{record.ID: record}},
		Data:          data,
		PublicBaseURL: "https://foreversixty.gg",
		Views:         views,
	})
	return mux
}

func sampleRecord(t *testing.T, title string) builds.Build {
	t.Helper()
	b, err := builds.New(builds.Input{
		ClassID: 1, RaceID: 1, TreeVersion: "test-1",
		PointOrder: []int{101, 101, 101, 101, 101, 103, 201},
		Title:      title,
	})
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func get(t *testing.T, h http.Handler, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	return rec
}

func TestBuildPageCarriesTheContractMetaTags(t *testing.T) {
	record := sampleRecord(t, "Arms leveling")
	rec := get(t, testPage(t, record, nil), "/b/"+record.ID)
	if rec.Code != http.StatusOK {
		t.Fatalf("code = %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
		t.Fatalf("content-type = %q", ct)
	}
	body := rec.Body.String()
	for _, want := range []string{
		"<title>Arms leveling · 6/1 · Forever Sixty</title>",
		`<meta name="description" content="Human Warrior build, 6/1 at level 16. Forever Sixty build planner.">`,
		`<link rel="canonical" href="https://foreversixty.gg/b/` + record.ID + `">`,
		`<meta property="og:title" content="Arms leveling · 6/1 · Forever Sixty">`,
		`<meta property="og:description" content="Human Warrior build, 6/1 at level 16. Forever Sixty build planner.">`,
		`<meta property="og:image" content="https://foreversixty.gg/b/` + record.ID + `/card.png">`,
		`<meta property="og:url" content="https://foreversixty.gg/b/` + record.ID + `">`,
		`<meta name="twitter:card" content="summary_large_image">`,
		`<script type="module" src="https://foreversixty.gg/planner-island.js"></script>`,
		`<main id="main">`,
	} {
		if !strings.Contains(body, want) {
			t.Errorf("page is missing:\n%s", want)
		}
	}
	if !strings.Contains(body, string(Header)) || !strings.Contains(body, string(Footer)) {
		t.Error("page is missing the site header or footer")
	}
}

func TestBuildPageFallsBackToRaceAndClassInTheTitle(t *testing.T) {
	record := sampleRecord(t, "")
	body := get(t, testPage(t, record, nil), "/b/"+record.ID).Body.String()
	if !strings.Contains(body, "<title>Human Warrior · 6/1 · Forever Sixty</title>") {
		t.Fatalf("title tag missing from:\n%s", body)
	}
}

var dataBuildAttr = regexp.MustCompile(`data-build='([^']*)'`)

func TestBuildPageInlinesTheRecord(t *testing.T) {
	record := sampleRecord(t, "Arms leveling")
	body := get(t, testPage(t, record, nil), "/b/"+record.ID).Body.String()

	m := dataBuildAttr.FindStringSubmatch(body)
	if m == nil {
		t.Fatalf("no data-build attribute in:\n%s", body)
	}
	var got builds.Build
	if err := json.Unmarshal([]byte(html.UnescapeString(m[1])), &got); err != nil {
		t.Fatalf("data-build is not the record JSON: %v", err)
	}
	if got.ID != record.ID || got.Title != "Arms leveling" || len(got.PointOrder) != 7 {
		t.Fatalf("inlined record = %+v", got)
	}
	if !strings.Contains(body, `data-tree-version="test-1"`) {
		t.Error("page is missing data-tree-version")
	}
}

func TestBuildPageRecordsAView(t *testing.T) {
	record := sampleRecord(t, "Arms leveling")
	views := &recordingViews{}
	get(t, testPage(t, record, views), "/b/"+record.ID)
	if len(views.ids) != 1 || views.ids[0] != record.ID {
		t.Fatalf("recorded views = %v", views.ids)
	}
}

func TestLoggerFallsBackToTheDefault(t *testing.T) {
	if (Deps{}).logger() == nil {
		t.Fatal("a Deps with no Log must still return a logger")
	}
}

func TestBuildPageReportsAStoreFailure(t *testing.T) {
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, Deps{
		Store:         failingGetter{err: errors.New("database is down")},
		Data:          data,
		PublicBaseURL: "https://foreversixty.gg",
		Log:           slog.New(slog.NewTextHandler(io.Discard, nil)),
	})

	rec := get(t, mux, "/b/znorjmts")
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("code = %d, want 500", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "<title>Build unavailable · Forever Sixty</title>") {
		t.Fatalf("body = %s", rec.Body.String())
	}
	if cc := rec.Header().Get("Cache-Control"); cc != "no-store" {
		t.Fatalf("cache-control = %q, want no-store on an error page", cc)
	}
}

func TestUnknownBuildRendersThe404Page(t *testing.T) {
	record := sampleRecord(t, "Arms leveling")
	views := &recordingViews{}
	rec := get(t, testPage(t, record, views), "/b/nosuchid")
	if rec.Code != http.StatusNotFound {
		t.Fatalf("code = %d", rec.Code)
	}
	body := rec.Body.String()
	for _, want := range []string{
		"<title>Build not found · Forever Sixty</title>",
		`href="https://foreversixty.gg/planner"`,
		string(Header),
		string(Footer),
	} {
		if !strings.Contains(body, want) {
			t.Errorf("404 page is missing:\n%s", want)
		}
	}
	if len(views.ids) != 0 {
		t.Fatalf("a missing build must not be counted as a view: %v", views.ids)
	}
}
