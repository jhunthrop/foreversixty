package site

import (
	"encoding/json"
	"errors"
	"html"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func TestBuildPageSitsInTheSitesPageColumn(t *testing.T) {
	record := sampleRecord(t, "Arms leveling")
	body := get(t, testPage(t, record, nil), "/b/"+record.ID).Body.String()
	column := strings.Index(body, `<div class="page-column" data-page-body>`)
	main := strings.Index(body, `<main id="main" tabindex="-1">`)
	planner := strings.Index(body, `<div id="planner"`)
	if column < 0 || main < column || planner < main {
		t.Fatalf("the planner must mount inside <main> inside the one .page-column:\n%s", body)
	}
	if strings.Count(body, `class="page-column"`) != 1 {
		t.Fatal("the page declares the site column exactly once, as Base.astro does")
	}
	if !strings.Contains(body, `<div class="site-top ui-band ui-band-end" data-art="band">`) {
		t.Fatal("the bar floats on the site's band, as on every site page")
	}
	// Until the island mounts, a screen-high reserve keeps the footer out of the first screen.
	if !strings.Contains(body, `<div class="b-reserve" aria-hidden="true"></div></div>`) {
		t.Fatal("the planner mount point must hold the reserve")
	}
}

func TestBuildPagePassesTheRecordsTitleAndViews(t *testing.T) {
	record := sampleRecord(t, "Arms leveling")
	record.Views = 41
	body := get(t, testPage(t, record, nil), "/b/"+record.ID).Body.String()
	m := dataBuildAttr.FindStringSubmatch(body)
	if m == nil {
		t.Fatal("no data-build attribute")
	}
	var got builds.Build
	if err := json.Unmarshal([]byte(html.UnescapeString(m[1])), &got); err != nil {
		t.Fatal(err)
	}
	if got.Title != "Arms leveling" || got.Views != 41 {
		t.Fatalf("the island reads the title and views from data-build; got %q, %d", got.Title, got.Views)
	}
	if !strings.Contains(body, "<title>Arms leveling · 6/1 · Forever Sixty</title>") {
		t.Fatal("the document title is the record's title")
	}
}

func TestBuildPageEscapesTheRecordsTitle(t *testing.T) {
	record := sampleRecord(t, `<script>alert('x')</script> & "Arms"`)
	body := get(t, testPage(t, record, nil), "/b/"+record.ID).Body.String()
	if strings.Contains(body, "<script>alert") {
		t.Fatalf("a title must never reach the page as markup:\n%s", body)
	}
	m := dataBuildAttr.FindStringSubmatch(body)
	if m == nil {
		t.Fatal("data-build must survive a title with quotes in it")
	}
	var got builds.Build
	if err := json.Unmarshal([]byte(html.UnescapeString(m[1])), &got); err != nil {
		t.Fatal(err)
	}
	if got.Title != record.Title {
		t.Fatalf("title round trip = %q, want %q", got.Title, record.Title)
	}
}

func TestNotFoundPageOffersThePlannerTheGuidesAndAsksTheSender(t *testing.T) {
	record := sampleRecord(t, "Arms leveling")
	rec := get(t, testPage(t, record, nil), "/b/nosuchid")
	body := rec.Body.String()
	for _, want := range []string{
		`<div class="page-column" data-page-body>`,
		`<div class="b-band ui-band ui-band-end full-bleed" data-art="band">`,
		`<h1 class="b-title" data-testid="page-title">Build not found</h1>`,
		`<div class="b-panel ui-surface ui-framed" data-region="plate" data-testid="b-message-panel">`,
		`data-testid="b-message-sentence">That link does not point at a saved build. ` +
			`<span data-testid="b-message-ask">Ask whoever sent it for a fresh link.</span></p>`,
		`<a href="https://foreversixty.gg/planner" class="` + primaryButtonClass + `">Open the build planner</a>`,
		`<a class="` + linkClass + `" href="https://foreversixty.gg/guides">Talent builds in the class guides ›</a>`,
	} {
		if !strings.Contains(body, want) {
			t.Errorf("not-found page is missing:\n%s", want)
		}
	}
	for _, retired := range []string{"gold-text", "font-display", "linear-gradient", "text-bg", "--color-strong"} {
		if strings.Contains(body, retired) {
			t.Errorf("the Cinzel gold heading and the gold button (%s) are retired", retired)
		}
	}
}

func TestErrorPageHasTheActionsButNoAskLine(t *testing.T) {
	d := Deps{PublicBaseURL: "https://foreversixty.gg"}
	rec := httptest.NewRecorder()
	d.writeMessage(rec, httptest.NewRequest(http.MethodGet, "/b/znorjmts", nil), http.StatusInternalServerError, "Build unavailable",
		"That build could not be loaded just now.")
	body := rec.Body.String()
	if !strings.Contains(body, ">Open the build planner</a>") || !strings.Contains(body, "Talent builds in the class guides ›") {
		t.Fatal("the error page keeps both actions")
	}
	if strings.Contains(body, askSenderLine) {
		t.Fatal("only a missing build asks for a fresh link; a failure on our side does not")
	}
}

func TestMessagePageEscapesItsText(t *testing.T) {
	page, err := (Deps{}).messageHTML("<b>Gone</b>", `"quoted" & <i>`, "")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(page, "<b>Gone</b>") || strings.Contains(page, "<i>") {
		t.Fatalf("message text must be escaped:\n%s", page)
	}
	if !strings.Contains(page, "&lt;b&gt;Gone&lt;/b&gt;") {
		t.Fatal("the heading is shown as text")
	}
}

// The message page's button and link are the web system's recipes (lib/ui/button.ts,
// lib/ui/primitives.ts LINK), whose Tailwind utilities planner-island.css carries because the site
// uses them. Every token here must still be in those files; skipped where web/ is not checked out.
func TestMessageActionClassesAreTheWebRecipes(t *testing.T) {
	for _, c := range []struct {
		file    string
		classes string
	}{
		{"../../../web/src/lib/ui/button.ts", primaryButtonClass},
		{"../../../web/src/lib/ui/primitives.ts", linkClass},
	} {
		src, err := os.ReadFile(c.file)
		if errors.Is(err, fs.ErrNotExist) {
			t.Skip(c.file + " is not checked out")
		}
		if err != nil {
			t.Fatal(err)
		}
		for _, token := range strings.Fields(c.classes) {
			// button.ts writes its pointer-hover variants as ${HOVER}<utility>.
			token = strings.TrimPrefix(token, "[@media(hover:hover)_and_(pointer:fine)]:hover:")
			if !strings.Contains(string(src), token) {
				t.Errorf("%s no longer has %q; copy the recipe again", c.file, token)
			}
		}
	}
}

// The /b/:id parity fixtures (web/tests/e2e/parity-b.spec.ts, premium spec 2026-11-08 G7): the
// page this package serves for the mock's Survival record (testdata/shared-survival.json: the
// planner mocks' `guide` state, titled and viewed as web-b) and for an unknown id, rendered with an
// empty PublicBaseURL so the island and its stylesheet resolve against the preview server. The
// record's point order is the mock's (design/mocks/data/planner-facts.json states.guide.points).
// Refresh after a template or chrome change with
//
//	go test ./internal/site -run TestSharedBuildParityFixtures -update
const parityFixtureDir = "../../../web/tests/e2e/fixtures/shared-build"

func TestSharedBuildParityFixtures(t *testing.T) {
	raw, err := os.ReadFile("testdata/shared-survival.json")
	if err != nil {
		t.Fatal(err)
	}
	var record builds.Build
	if err := json.Unmarshal(raw, &record); err != nil {
		t.Fatal(err)
	}
	src := filepath.Join("../../../data/builds", record.TreeVersion)
	if _, err := os.Stat(src); errors.Is(err, fs.ErrNotExist) {
		t.Skip(src + " is not checked out; the committed fixtures stand")
	}
	dir := t.TempDir()
	if err := trees.LinkBuilds(dir, "../../../data/builds", record.TreeVersion); err != nil {
		t.Fatal(err)
	}
	if err := trees.WriteActiveBuildFile(dir, record.TreeVersion); err != nil {
		t.Fatal(err)
	}
	data, err := trees.Load(dir)
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, Deps{Store: fakeGetter{rows: map[string]builds.Build{record.ID: record}}, Data: data})
	for name, path := range map[string]string{"survival.html": "/b/" + record.ID, "missing.html": "/b/nosuchbuild"} {
		got := get(t, mux, path).Body.String()
		file := filepath.Join(parityFixtureDir, name)
		if *update {
			if err := os.WriteFile(file, []byte(got), 0o644); err != nil {
				t.Fatal(err)
			}
			continue
		}
		want, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("%v; generate it with -update", err)
		}
		if string(want) != got {
			t.Errorf("%s is stale: rerun go test ./internal/site -run TestSharedBuildParityFixtures -update", file)
		}
	}
}
