package ui

import (
	"bytes"
	"io/fs"
	"mime"
	"os"
	"regexp"
	"strings"
	"testing"
)

// The six faces, the art and the seal are what the stylesheet and the
// page ask for; a missing file would be a silent fall back to a system face.
func TestTheFacesArtAndSealAreEmbedded(t *testing.T) {
	for _, name := range []string{
		"fonts/open-sans-400-subset.woff2",
		"fonts/open-sans-600-subset.woff2",
		"fonts/open-sans-700-subset.woff2",
		"fonts/marcellus-400-subset.woff2",
		"fonts/archivo-600-subset.woff2",
		"fonts/jetbrains-mono-500-subset.woff2",
		"art/ground.webp", "art/band-1200.webp", "seal.svg",
		"index.html", "styles.css", "tokens.css", "icons.js", "app.js", "copy_gen.js", "copy.js", "format.js",
	} {
		b, err := fs.ReadFile(Assets, name)
		if err != nil || len(b) == 0 {
			t.Errorf("%s: %d bytes, %v", name, len(b), err)
		}
	}
}

func TestEveryFontFileIsAWoff2(t *testing.T) {
	files, err := fs.Glob(Assets, "fonts/*.woff2")
	if err != nil || len(files) != 6 {
		t.Fatalf("fonts = %v, %v", files, err)
	}
	for _, f := range files {
		b, _ := fs.ReadFile(Assets, f)
		if len(b) < 4 || string(b[:4]) != "wOF2" {
			t.Errorf("%s does not start with the wOF2 signature", f)
		}
	}
}

// The header seal is the site's small flat seal, byte for byte (tools/icons/render.py copies it).
func TestTheSealIsTheSitesFlatSeal(t *testing.T) {
	ours, _ := fs.ReadFile(Assets, "seal.svg")
	theirs, err := os.ReadFile("../../web/public/art/seal-flat-small.svg")
	if err != nil || !bytes.Equal(ours, theirs) {
		t.Errorf("seal.svg is not the site's art/seal-flat-small.svg (%v)", err)
	}
}

// The faces and the art are the site's own files, byte for byte.
func TestTheFacesAndArtAreTheSitesFiles(t *testing.T) {
	files, _ := fs.Glob(Assets, "fonts/*.woff2")
	art, _ := fs.Glob(Assets, "art/*.webp")
	site := map[string]string{"fonts/": "../../web/src/assets/fonts/", "art/": "../../web/public/art/"}
	for _, f := range append(files, art...) {
		dir := f[:strings.Index(f, "/")+1]
		ours, _ := fs.ReadFile(Assets, f)
		theirs, err := os.ReadFile(site[dir] + strings.TrimPrefix(f, dir))
		if err != nil || !bytes.Equal(ours, theirs) {
			t.Errorf("%s is not the site's file (%v)", f, err)
		}
	}
}

func TestWoff2IsServedAsAFont(t *testing.T) {
	if got := mime.TypeByExtension(".woff2"); !strings.HasPrefix(got, "font/woff2") {
		t.Fatalf("type = %q", got)
	}
}

// Every file the page or the stylesheet names ships.
func TestTheStylesheetAndPageNameOnlyFilesThatShip(t *testing.T) {
	ref := regexp.MustCompile(`(?:href|src)="([^"#:]+)"|url\('([^']+)'\)`)
	for _, page := range []string{"styles.css", "index.html", "install.html"} {
		b, _ := fs.ReadFile(Assets, page)
		for _, m := range ref.FindAllStringSubmatch(string(b), -1) {
			name := m[1] + m[2]
			if _, err := fs.Stat(Assets, name); err != nil {
				t.Errorf("%s names %s: %v", page, name, err)
			}
		}
	}
}

// The stylesheet's own literals: colours, font sizes and radii. Each pattern
// names what it forbids; stylesheetLiterals is the scan, so the self-check
// below proves it catches each kind.
var stylesheetLiterals = []struct {
	what, fix string
	re        *regexp.Regexp
}{
	{"a hex colour", "a --color-* token", regexp.MustCompile(`#[0-9a-fA-F]{3,8}(?:[^-\w]|$)`)},
	{"an rgb or rgba colour", "a --color-* veil token", regexp.MustCompile(`rgba?\([^)]*\)`)},
	{"a named black or white", "--color-edge or --color-strong", regexp.MustCompile(`(?i)(?:^|[\s:])(?:black|white)(?:[\s;!]|$)`)},
	{"a px font size", "a --text-* role", regexp.MustCompile(`font-size:\s*\d+px|font:\s*(?:\d{3}\s+)?\d+px`)},
	{"a px radius", "a --radius-* token", regexp.MustCompile(`border-radius:\s*[^;]*\d+px`)},
}

// stylesheetFindings is every own literal in css, comments left out. A radius
// derived from a token (calc(var(--radius-plate) - 4px)) is the token's, not
// a literal of its own.
func stylesheetFindings(css string) []string {
	css = regexp.MustCompile(`(?s)/\*.*?\*/`).ReplaceAllString(css, "")
	var out []string
	for _, line := range strings.Split(css, "\n") {
		for _, l := range stylesheetLiterals {
			for _, m := range l.re.FindAllString(line, -1) {
				if l.what == "a px radius" && strings.Contains(line, "calc(var(--radius-") {
					continue
				}
				out = append(out, l.what+" ("+strings.TrimSpace(m)+"); use "+l.fix)
			}
		}
	}
	return out
}

// Colour, type size and corner come from the tokens alone: styles.css composes
// tokens.css (generated from design/tokens/tokens.json) and declares none of
// its own.
func TestTheStylesheetDeclaresNoColourSizeOrRadiusOfItsOwn(t *testing.T) {
	b, _ := fs.ReadFile(Assets, "styles.css")
	for _, f := range stylesheetFindings(string(b)) {
		t.Errorf("styles.css has %s", f)
	}
}

// The scan catches each kind it names, and leaves token uses and comments alone.
func TestTheStylesheetScanCatchesEachLiteral(t *testing.T) {
	planted := ".a { color: #fff; }\n.b { background: rgba(255, 255, 255, 0.1); }\n" +
		".c { border: 1px solid black; }\n.d { font-size: 14px; }\n.e { font: 600 14px/1 x; }\n" +
		".f { border-radius: 11px; }\n"
	if got := len(stylesheetFindings(planted)); got != 6 {
		t.Errorf("planted 6 literals, the scan found %d: %v", got, stylesheetFindings(planted))
	}
	clean := "/* brown-black, white at .72, 14px */\n.a { font: 600 var(--text-small)/18px var(--font-body); " +
		"border-radius: calc(var(--radius-plate) - 4px); color: var(--color-edge); white-space: nowrap; }\n"
	if f := stylesheetFindings(clean); len(f) > 0 {
		t.Errorf("token uses were flagged: %v", f)
	}
}

// tokens.css and icons.js are the generators' output, never a hand copy.
func TestTheTokensAndIconsAreGenerated(t *testing.T) {
	for name, line := range map[string]string{
		"tokens.css": "GENERATED by design/tokens/build.mjs",
		"icons.js":   "GENERATED by web/scripts/gen-ui-icons.mjs",
	} {
		b, _ := fs.ReadFile(Assets, name)
		if !strings.Contains(string(b[:min(len(b), 200)]), line) {
			t.Errorf("%s lost its %q line; rerun its generator", name, line)
		}
	}
}
