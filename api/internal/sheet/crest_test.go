package sheet

import (
	"image"
	"image/color"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
)

func TestCrestURL(t *testing.T) {
	if got := CrestURL("https://api.foreversixty.gg/", "mage"); got != "https://api.foreversixty.gg/v1/discord/tiles/crest/mage.png" {
		t.Fatal(got)
	}
	if got := CrestURL("https://api.foreversixty.gg", "monk"); got != "https://api.foreversixty.gg/v1/discord/tiles/crest/none.png" {
		t.Fatal(got)
	}
}

// crestRingPixel is the middle of the top ring; crestDiscPixel sits just inside the ring, where
// the crest art is transparent and the tinted disc shows.
var (
	crestRingPixel = image.Pt(CrestSide/2, 1)
	crestDiscPixel = image.Pt(CrestSide/2, crestRing+1)
)

// nrgbaAt is the pixel non-premultiplied, so the ring's alpha and colour read as drawn.
func nrgbaAt(img image.Image, p image.Point) color.NRGBA {
	return color.NRGBAModel.Convert(img.At(p.X, p.Y)).(color.NRGBA)
}

func near(a, b uint8) bool { d := int(a) - int(b); return d >= -3 && d <= 3 }

func TestCrestRingPerClass(t *testing.T) {
	for slug, class := range classColors {
		img := decodePNG(t, mustCrest(t, slug))
		if img.Bounds() != image.Rect(0, 0, CrestSide, CrestSide) {
			t.Fatalf("%s: %v", slug, img.Bounds())
		}
		ring := nrgbaAt(img, crestRingPixel)
		if !near(ring.R, class.R) || !near(ring.G, class.G) || !near(ring.B, class.B) || !near(ring.A, 140) {
			t.Errorf("%s: ring %v, want the class colour at 55%%", slug, ring)
		}
		if _, _, _, a := img.At(0, 0).RGBA(); a != 0 {
			t.Errorf("%s: the corner outside the circle is not transparent (no glow, no square)", slug)
		}
		if disc := nrgbaAt(img, crestDiscPixel); disc.A != 0xff {
			t.Errorf("%s: the disc under the crest is not opaque: %v", slug, disc)
		}
	}
}

func TestCrestFallbackIsTheSealInAMutedRing(t *testing.T) {
	img := decodePNG(t, mustCrest(t, "not-a-class"))
	muted := opaque(brand.ColorBaseMuted)
	ring := nrgbaAt(img, crestRingPixel)
	if !near(ring.R, muted.R) || !near(ring.G, muted.G) || !near(ring.B, muted.B) {
		t.Errorf("fallback ring %v, want muted %v", ring, muted)
	}
	if c := nrgbaAt(img, image.Pt(CrestSide/2, CrestSide/2)); c.A != 0xff {
		t.Errorf("fallback middle not drawn: %v", c)
	}
}

func TestCrestRoute(t *testing.T) {
	srv, _, _ := tileServer(t)
	for file, want := range map[string]int{"hunter.png": 200, "none.png": 200, "monk.png": 404, "hunter.webp": 404} {
		res := get(t, srv.URL+CrestPath+file)
		if res.StatusCode != want {
			t.Errorf("%s: status %d, want %d", file, res.StatusCode, want)
			continue
		}
		if want != http.StatusOK {
			continue
		}
		if res.Header.Get("Content-Type") != "image/png" || res.Header.Get("Cache-Control") != "public, max-age=31536000, immutable" {
			t.Errorf("%s: type %q cache %q", file, res.Header.Get("Content-Type"), res.Header.Get("Cache-Control"))
		}
	}
	// The item tiles beside it still route.
	if res := get(t, srv.URL+TilePath+"1.60.1.70658/16846.png"); res.StatusCode != http.StatusOK {
		t.Fatalf("tile status %d", res.StatusCode)
	}
}

// TestWriteReviewThumbs writes the review captures when CREST_REVIEW_DIR is set.
func TestWriteReviewThumbs(t *testing.T) {
	dir := os.Getenv("CREST_REVIEW_DIR")
	if dir == "" {
		t.Skip("CREST_REVIEW_DIR not set")
	}
	write := func(name string, png []byte) {
		if err := os.WriteFile(filepath.Join(dir, name), png, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("character-thumb.png", mustCrest(t, "paladin"))
	write("character-thumb-unknown.png", mustCrest(t, ""))
	icon, err := testIcons().Icon("inv_helmet_24")
	if err != nil {
		t.Fatal(err)
	}
	tile, err := Tile(icon, 4)
	if err != nil {
		t.Fatal(err)
	}
	write("item-thumb.png", tile)
}

func mustCrest(t *testing.T, slug string) []byte {
	t.Helper()
	png, err := CrestPNG(slug)
	if err != nil {
		t.Fatal(err)
	}
	return png
}

// bareCrestPattern is the web crest path a Go file must never build by hand.
const bareCrestPattern = "icons/hd/crests"

// bareCrestFiles lists the non-test Go files under root that build a bare web crest URL.
func bareCrestFiles(root string) ([]string, error) {
	var found []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return err
		}
		body, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		if strings.Contains(string(body), bareCrestPattern) {
			found = append(found, path)
		}
		return nil
	})
	return found, err
}

// No Go file outside this package builds a bare web crest URL: every Discord crest is the ringed
// PNG from CrestURL (brand audit 2026-11-10 hand-off D). The guard walks the whole api tree.
func TestNoBareCrestURLOutsideSheet(t *testing.T) {
	found, err := bareCrestFiles("../..")
	if err != nil {
		t.Fatal(err)
	}
	if len(found) > 0 {
		t.Fatalf("bare crest URLs, use sheet.CrestURL: %v", found)
	}
}

// The guard's self-check: a planted file with the pattern is found; a test file is not.
func TestBareCrestGuardCatchesItsPattern(t *testing.T) {
	dir := t.TempDir()
	planted := filepath.Join(dir, "embed.go")
	body := []byte("package x\n\nconst u = \"/" + bareCrestPattern + "/mage.png\"\n")
	if err := os.WriteFile(planted, body, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "embed_test.go"), body, 0o600); err != nil {
		t.Fatal(err)
	}
	found, err := bareCrestFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(found) != 1 || found[0] != planted {
		t.Fatalf("guard found %v, want only %s", found, planted)
	}
}
