package brand

import (
	"crypto/sha256"
	"flag"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

// preview writes testdata-free previews of the paint to the given directory, for review:
// go test ./internal/brand -run Preview -preview /tmp/out
var preview = flag.String("preview", "", "write paint previews into this directory")

func TestArtIsTheSitesBytes(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	for embedded, site := range ArtSources {
		a, err := ArtFile(embedded)
		if err != nil {
			t.Fatal(err)
		}
		b, err := os.ReadFile(filepath.Join(root, site))
		if err != nil {
			t.Fatal(err)
		}
		if sha256.Sum256(a) != sha256.Sum256(b) {
			t.Errorf("%s differs from %s: run make sheet-assets", embedded, site)
		}
	}
}

func TestEveryFaceParses(t *testing.T) {
	for f := OpenSans400; f <= Marcellus400; f++ {
		face, err := NewFace(f, 16)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := face.GlyphAdvance('Ä'); !ok {
			t.Errorf("face %d has no Latin-1 letters", f)
		}
	}
	if _, err := Font(Marcellus400 + 1); err == nil {
		t.Error("an unknown face parsed")
	}
}

func TestArtDecodes(t *testing.T) {
	for _, get := range []func() (image.Image, error){Seal, groundImage, bandImage} {
		img, err := get()
		if err != nil {
			t.Fatal(err)
		}
		if img.Bounds().Empty() {
			t.Fatal("empty art")
		}
	}
}

func TestCornerMirrors(t *testing.T) {
	const side = 64
	tl, err := CornerMask(side, TopLeft)
	if err != nil {
		t.Fatal(err)
	}
	br, err := CornerMask(side, BottomRight)
	if err != nil {
		t.Fatal(err)
	}
	ink := 0
	for y := range side {
		for x := range side {
			a, b := tl.AlphaAt(x, y).A, br.AlphaAt(side-1-x, side-1-y).A
			if diff := int(a) - int(b); diff > 24 || diff < -24 {
				t.Fatalf("(%d,%d): top left %d, bottom right mirrored %d", x, y, a, b)
			}
			if a > 128 {
				ink++
			}
		}
	}
	if share := float64(ink) / side / side; share < 0.02 || share > 0.6 {
		t.Fatalf("the ornament inks %.0f%% of its square", share*100)
	}
}

func TestParsePolygonsRejectsCurves(t *testing.T) {
	if _, err := parsePolygons("M0,0 C1,1 2,2 3,3Z"); err == nil {
		t.Fatal("a curve parsed")
	}
}

func TestRGB(t *testing.T) {
	if got := RGB(ColorForeverAccent); got != 0x0093b2 {
		t.Fatalf("accent %06x", got)
	}
}

func TestPreviewPanel(t *testing.T) {
	if *preview == "" {
		t.Skip("no -preview directory")
	}
	img := image.NewRGBA(image.Rect(0, 0, 1200, 800))
	if err := Ground(img, img.Bounds()); err != nil {
		t.Fatal(err)
	}
	if err := Band(img, image.Rect(0, 0, 1200, 140)); err != nil {
		t.Fatal(err)
	}
	if err := Panel(img, image.Rect(40, 200, 560, 760), 2, PanelStyle{Ornate: true}); err != nil {
		t.Fatal(err)
	}
	if err := Panel(img, image.Rect(640, 200, 1160, 760), 2, PanelStyle{Framed: true}); err != nil {
		t.Fatal(err)
	}
	f, err := os.Create(filepath.Join(*preview, "brand-panel.png"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		t.Fatal(err)
	}
}
