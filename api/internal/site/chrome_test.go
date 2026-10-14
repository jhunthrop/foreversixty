package site

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"
)

var update = flag.Bool("update", false,
	"rewrite chrome.html and testdata/chrome.html from the built site page (-site)")

// sitePage is any built content page: all of them carry the same chrome. It defaults to
// this checkout's own web build; -site points it at another worktree's build, when the
// site is built on a different branch than the api.
var sitePage = flag.String("site", "../../../web/dist/about.html", "the built site page to read the chrome from")

const snapshotPath = "testdata/chrome.html"

// chromeNote opens chrome.html and its snapshot, saying where the copy comes from and how to
// regenerate it; splitChrome drops it.
const chromeNote = `<!-- The site's own header and footer, copied from a built page of the web site
(web/dist/about.html) so that /b/{id} wears the site's chrome. Generated: do not edit by hand.
Regenerate both chrome.html and testdata/chrome.html from the repository root with
  (cd web && npm run build) && (cd api && go test ./internal/site -run TestChromeSnapshot -update)
or, when the site is built in another worktree, from api/ with
  go test ./internal/site -run TestChromeSnapshot -update -site <that worktree>/web/dist/about.html
then run go test ./internal/site again without -update. -->`

// extractChrome pulls the site's chrome out of a built page. The site header
// is the first <header> in the document; the site footer is the last
// <footer>, because a content page can carry its own (the /about page ends
// with a sources <footer> inside its article).
func extractChrome(page string) (header, footer string, err error) {
	headerStart := strings.Index(page, "<header")
	headerEnd := strings.Index(page, "</header>")
	footerStart := strings.LastIndex(page, "<footer")
	footerEnd := strings.LastIndex(page, "</footer>")
	if headerStart < 0 || headerEnd < headerStart || footerStart < 0 || footerEnd < footerStart {
		return "", "", fmt.Errorf("site: no <header>...</header> and <footer>...</footer> pair in the page")
	}
	return page[headerStart : headerEnd+len("</header>")], page[footerStart : footerEnd+len("</footer>")], nil
}

// TestChromeTemplateMatchesTheSnapshot always runs: it proves the markup the
// binary serves is the markup that was recorded from the site.
func TestChromeTemplateMatchesTheSnapshot(t *testing.T) {
	snap, err := os.ReadFile(snapshotPath)
	if err != nil {
		t.Fatal(err)
	}
	wantHeader, wantFooter := splitChrome(string(snap))
	if Header != wantHeader {
		t.Errorf("chrome.html header differs from the snapshot\n got: %s\nwant: %s", Header, wantHeader)
	}
	if Footer != wantFooter {
		t.Errorf("chrome.html footer differs from the snapshot\n got: %s\nwant: %s", Footer, wantFooter)
	}
}

// TestChromeSnapshotMatchesTheBuiltSite runs only where the site has been
// built, so the api suite never depends on the web build. Refresh both files
// with: (cd web && npm run build) && cd api && go test ./internal/site -run TestChromeSnapshot -update
func TestChromeSnapshotMatchesTheBuiltSite(t *testing.T) {
	page, err := os.ReadFile(*sitePage)
	if errors.Is(err, fs.ErrNotExist) {
		t.Skip(*sitePage + " is not built; the checked-in snapshot is the reference")
	}
	if err != nil {
		t.Fatal(err)
	}
	header, footer, err := extractChrome(string(page))
	if err != nil {
		t.Fatal(err)
	}
	if *update {
		body := chromeNote + "\n" + header + "\n" + mainMarker + "\n" + footer + "\n"
		if err := os.WriteFile(snapshotPath, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile("chrome.html", []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
		t.Log("rewrote chrome.html and testdata/chrome.html; rebuild and rerun without -update")
		return
	}
	snap, err := os.ReadFile(snapshotPath)
	if err != nil {
		t.Fatal(err)
	}
	gotHeader, gotFooter := splitChrome(string(snap))
	if string(gotHeader) != header {
		t.Errorf("the site's header drifted from the snapshot; rerun with -update\n site: %s\nsnap: %s", header, gotHeader)
	}
	if string(gotFooter) != footer {
		t.Errorf("the site's footer drifted from the snapshot; rerun with -update\n site: %s\nsnap: %s", footer, gotFooter)
	}
}

func TestSplitChromeDropsTheLeadingNote(t *testing.T) {
	header, footer := splitChrome(chromeNote + "\n<header>h</header>\n" + mainMarker + "\n<footer>f</footer>\n")
	if header != "<header>h</header>" || footer != "<footer>f</footer>" {
		t.Fatalf("header %q footer %q; want the chrome without the note", header, footer)
	}
	if !strings.HasPrefix(chromeHTML, "<!--") || strings.Contains(string(Header), "Regenerate") {
		t.Fatal("chrome.html must open with its regeneration note, and the served header must not carry it")
	}
}

func TestSplitChromePanicsWithoutTheMarker(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("a chrome document without the marker must panic")
		}
	}()
	splitChrome("<header></header><footer></footer>")
}
