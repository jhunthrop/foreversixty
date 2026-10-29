package sheet

import (
	"bytes"
	"encoding/json"
	"flag"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

// update rewrites the goldens: go test ./internal/sheet -run Golden -update
var update = flag.Bool("update", false, "rewrite the golden PNGs under testdata")

// The golden comparison's tolerance: a pixel differs when any channel is more than
// pixelTolerance apart, and the render matches when at most maxDiffShare of pixels differ.
// The slack is for rasterising (font hinting and anti-aliasing across Go versions), never for
// a moved box: a shifted row or a wrong colour differs in far more pixels.
const (
	pixelTolerance = 24
	maxDiffShare   = 0.002
)

// testIcons are the mock's icons, copied from the build into testdata so the goldens never move
// with the active build.
func testIcons() IconSource { return NewFileIcons(filepath.Join("testdata", "icons")) }

func loadFixture(t *testing.T, name string) Sheet {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("testdata", name+".json"))
	if err != nil {
		t.Fatal(err)
	}
	var s Sheet
	if err := json.Unmarshal(raw, &s); err != nil {
		t.Fatal(err)
	}
	return s
}

func decodePNG(t *testing.T, b []byte) image.Image {
	t.Helper()
	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	return img
}

func TestGoldenSheets(t *testing.T) {
	for _, name := range []string{"sheet-bis", "sheet-char"} {
		t.Run(name, func(t *testing.T) {
			got, err := Render(loadFixture(t, name), testIcons())
			if err != nil {
				t.Fatal(err)
			}
			golden := filepath.Join("testdata", name+".golden.png")
			if *update {
				if err := os.WriteFile(golden, got, 0o600); err != nil {
					t.Fatal(err)
				}
			}
			want, err := os.ReadFile(golden)
			if err != nil {
				t.Fatalf("%v (run with -update to write it)", err)
			}
			g, w := decodePNG(t, got), decodePNG(t, want)
			if g.Bounds() != image.Rect(0, 0, Width, Height) || w.Bounds() != g.Bounds() {
				t.Fatalf("bounds %v, golden %v, want %dx%d", g.Bounds(), w.Bounds(), Width, Height)
			}
			if share := diffShare(g, w); share > maxDiffShare {
				t.Fatalf("%.4f%% of pixels differ from %s, over %.2f%%", share*100, golden, maxDiffShare*100)
			}
		})
	}
}

// diffShare is the share of pixels whose channels differ by more than pixelTolerance.
func diffShare(a, b image.Image) float64 {
	r := a.Bounds()
	diff := 0
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			if !closeColor(a.At(x, y), b.At(x, y)) {
				diff++
			}
		}
	}
	return float64(diff) / float64(r.Dx()*r.Dy())
}

func closeColor(a, b color.Color) bool {
	ar, ag, ab, aa := a.RGBA()
	br, bg, bb, ba := b.RGBA()
	for _, p := range [][2]uint32{{ar, br}, {ag, bg}, {ab, bb}, {aa, ba}} {
		d := int(p[0]>>8) - int(p[1]>>8)
		if d < -pixelTolerance || d > pixelTolerance {
			return false
		}
	}
	return true
}

func readJSON(t *testing.T, path string, v any) {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, v); err != nil {
		t.Fatal(err)
	}
}
