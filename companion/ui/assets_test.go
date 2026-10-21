package ui

import (
	"io/fs"
	"mime"
	"strings"
	"testing"
)

// The five faces and the seal are what the stylesheet and the header
// ask for; a missing file would be a silent fall back to Georgia.
func TestTheFacesAndTheSealAreEmbedded(t *testing.T) {
	for _, name := range []string{
		"fonts/cinzel-700-subset.woff2",
		"fonts/barlow-400-subset.woff2",
		"fonts/barlow-600-subset.woff2",
		"fonts/barlow-700-subset.woff2",
		"fonts/jetbrains-mono-500-subset.woff2",
		"seal.svg", "index.html", "styles.css", "app.js", "copy.js", "format.js",
	} {
		b, err := fs.ReadFile(Assets, name)
		if err != nil || len(b) == 0 {
			t.Errorf("%s: %d bytes, %v", name, len(b), err)
		}
	}
}

func TestEveryFontFileIsAWoff2(t *testing.T) {
	files, err := fs.Glob(Assets, "fonts/*.woff2")
	if err != nil || len(files) != 5 {
		t.Fatalf("fonts = %v, %v", files, err)
	}
	for _, f := range files {
		b, _ := fs.ReadFile(Assets, f)
		if len(b) < 4 || string(b[:4]) != "wOF2" {
			t.Errorf("%s does not start with the wOF2 signature", f)
		}
	}
}

func TestWoff2IsServedAsAFont(t *testing.T) {
	if got := mime.TypeByExtension(".woff2"); !strings.HasPrefix(got, "font/woff2") {
		t.Fatalf("type = %q", got)
	}
}

// Every face the page preloads or declares is a file that ships.
func TestTheStylesheetAndPageNameOnlyFilesThatShip(t *testing.T) {
	for _, page := range []string{"styles.css", "index.html"} {
		b, _ := fs.ReadFile(Assets, page)
		for _, part := range strings.Split(string(b), "fonts/")[1:] {
			name := part[:strings.IndexAny(part, "')\" ")]
			if _, err := fs.Stat(Assets, "fonts/"+name); err != nil {
				t.Errorf("%s names fonts/%s: %v", page, name, err)
			}
		}
	}
}
