// Package card draws the 1200x630 preview image shown when a build link is
// unfurled. It draws directly onto an RGBA image with golang.org/x/image
// rather than going through SVG: the site's TypeScript OG renderer
// (web/src/lib/og.ts) uses satori and resvg, which is a JavaScript
// toolchain this service has no reason to carry.
//
// The two faces under fonts/ are the site's display and body faces, taken
// as TTF from the Fontsource CDN at the same version web/package.json pins
// (cinzel@5.3.0, barlow@5.3.0). Both are SIL Open Font License 1.1; the
// license text sits beside them as OFL-Cinzel.txt and OFL-Barlow.txt.
package card

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

//go:embed fonts/*.ttf
var fontFS embed.FS

const (
	// Width and Height are the Open Graph card size the contract fixes.
	Width  = 1200
	Height = 630

	bandHeight = 24
	marginX    = 72

	titleBaseline    = 160
	subtitleBaseline = 222
	splitBaseline    = 410
	levelBaseline    = 474
	wordmarkBaseline = 566
)

// The palette is the site's: bg, text, muted, and gold from
// design/DESIGN-SYSTEM.md, the same values web/src/lib/og.ts draws with.
var (
	bgColor    = color.RGBA{R: 0x07, G: 0x0b, B: 0x12, A: 0xff}
	textColor  = color.RGBA{R: 0xf2, G: 0xee, B: 0xe4, A: 0xff}
	mutedColor = color.RGBA{R: 0x9a, G: 0x94, B: 0x84, A: 0xff}
	goldColor  = color.RGBA{R: 0xe5, G: 0xb9, B: 0x55, A: 0xff}
)

// Input is everything the card shows. Split is points per tree in client
// order; Level is the build's level as builds.Describe computes it.
type Input struct {
	Title      string
	Race       string
	Class      string
	ClassColor string
	Split      []int
	Level      int
}

type faces struct {
	title    font.Face
	numerals font.Face
	body     font.Face
	wordmark font.Face
}

// loadFaces parses the embedded fonts once. Parsing can only fail if the
// embedded files are corrupt, which is a build problem, not a request one.
var loadFaces = sync.OnceValues(func() (*faces, error) {
	cinzel, err := parseFont("fonts/Cinzel-Bold.ttf")
	if err != nil {
		return nil, err
	}
	barlow, err := parseFont("fonts/Barlow-Regular.ttf")
	if err != nil {
		return nil, err
	}
	f := &faces{}
	for _, spec := range []struct {
		dst  *font.Face
		src  *opentype.Font
		size float64
	}{
		{&f.title, cinzel, 64},
		{&f.numerals, cinzel, 96},
		{&f.body, barlow, 34},
		{&f.wordmark, barlow, 26},
	} {
		face, err := opentype.NewFace(spec.src, &opentype.FaceOptions{Size: spec.size, DPI: 72, Hinting: font.HintingFull})
		if err != nil {
			return nil, fmt.Errorf("card: face at %.0fpx: %w", spec.size, err)
		}
		*spec.dst = face
	}
	return f, nil
})

func parseFont(name string) (*opentype.Font, error) {
	b, err := fontFS.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("card: read %s: %w", name, err)
	}
	f, err := opentype.Parse(b)
	if err != nil {
		return nil, fmt.Errorf("card: parse %s: %w", name, err)
	}
	return f, nil
}

// Render draws the card for one build.
func Render(in Input) ([]byte, error) {
	f, err := loadFaces()
	if err != nil {
		return nil, err
	}
	accent := parseHexColor(in.ClassColor)
	img := canvas(accent)

	drawText(img, f.title, goldColor, marginX, titleBaseline, fit(f.title, in.Title, Width-2*marginX))
	drawText(img, f.body, textColor, marginX, subtitleBaseline,
		fit(f.body, strings.TrimSpace(in.Race+" "+in.Class), Width-2*marginX))
	drawText(img, f.numerals, accent, marginX, splitBaseline, splitText(in.Split))
	drawText(img, f.body, mutedColor, marginX, levelBaseline, "Level "+strconv.Itoa(in.Level))
	drawText(img, f.wordmark, mutedColor, marginX, wordmarkBaseline, "foreversixty.gg")

	return encode(img)
}

// fallback is rendered once and reused: it is the image served when a
// build's own card cannot be drawn, so it must never itself fail at
// request time.
var fallback = sync.OnceValue(func() []byte {
	f, err := loadFaces()
	if err != nil {
		return nil
	}
	img := canvas(goldColor)
	drawText(img, f.title, goldColor, marginX, titleBaseline, "Forever Sixty")
	drawText(img, f.body, textColor, marginX, subtitleBaseline, "Build planner")
	drawText(img, f.wordmark, mutedColor, marginX, wordmarkBaseline, "foreversixty.gg")
	out, err := encode(img)
	if err != nil {
		return nil
	}
	return out
})

// Fallback is the static card served when a build's card cannot be drawn.
// It returns nil only when the embedded fonts cannot be used at all, which
// the package's tests rule out.
func Fallback() []byte { return fallback() }

func canvas(band color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	draw.Draw(img, img.Bounds(), image.NewUniform(bgColor), image.Point{}, draw.Src)
	draw.Draw(img, image.Rect(0, 0, Width, bandHeight), image.NewUniform(band), image.Point{}, draw.Src)
	return img
}

func encode(img *image.RGBA) ([]byte, error) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, fmt.Errorf("card: encode: %w", err)
	}
	return buf.Bytes(), nil
}

func drawText(dst draw.Image, face font.Face, col color.Color, x, baseline int, s string) {
	d := &font.Drawer{Dst: dst, Src: image.NewUniform(col), Face: face, Dot: fixed.P(x, baseline)}
	d.DrawString(s)
}

// splitText renders the per-tree point counts, spaced for the large face.
func splitText(split []int) string {
	if len(split) == 0 {
		return "0"
	}
	parts := make([]string, len(split))
	for i, n := range split {
		parts[i] = strconv.Itoa(n)
	}
	return strings.Join(parts, " / ")
}

// fit trims s with an ellipsis until it draws within maxWidth. The faces
// are the latin subsets, so the ellipsis is three periods rather than the
// single character, which those subsets may not carry.
func fit(face font.Face, s string, maxWidth int) string {
	s = strings.TrimSpace(s)
	limit := fixed.I(maxWidth)
	if font.MeasureString(face, s) <= limit {
		return s
	}
	r := []rune(s)
	for len(r) > 0 {
		r = r[:len(r)-1]
		if font.MeasureString(face, string(r)+"...") <= limit {
			return strings.TrimRight(string(r), " ") + "..."
		}
	}
	return ""
}

// parseHexColor reads "#rrggbb" and falls back to the site's gold.
func parseHexColor(s string) color.RGBA {
	if len(s) == 7 && s[0] == '#' {
		if v, err := strconv.ParseUint(s[1:], 16, 32); err == nil {
			return color.RGBA{R: uint8(v >> 16), G: uint8(v >> 8), B: uint8(v), A: 0xff}
		}
	}
	return goldColor
}
