package sheet

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"strings"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

// fontFS holds the sheet's three faces, full Latin TTFs from the site's @fontsource packages
// (`make sheet-fonts`), with their OFL licences beside them.
//
//go:embed fonts/*.ttf
var fontFS embed.FS

// The faces, by the weight the spec names.
type typeface int

const (
	barlow400 typeface = iota
	barlow600
	cinzel700
)

var typefaceFiles = [...]string{
	barlow400: "fonts/Barlow-Regular.ttf",
	barlow600: "fonts/Barlow-SemiBold.ttf",
	cinzel700: "fonts/Cinzel-Bold.ttf",
}

// loadFonts parses the embedded faces once; a parsed font is read-only and shared.
var loadFonts = sync.OnceValues(func() ([]*opentype.Font, error) {
	out := make([]*opentype.Font, len(typefaceFiles))
	for i, name := range typefaceFiles {
		raw, err := fontFS.ReadFile(name)
		if err != nil {
			return nil, fmt.Errorf("sheet: read %s: %w", name, err)
		}
		f, err := opentype.Parse(raw)
		if err != nil {
			return nil, fmt.Errorf("sheet: parse %s: %w", name, err)
		}
		out[i] = f
	}
	return out, nil
})

// style is one text style of the spec, in logical px: face, size, line height, letter
// spacing (em) and whether it is drawn uppercase.
type style struct {
	face       typeface
	size       float64
	lineHeight float64
	trackingEm float64
	upper      bool
}

// faceSet builds faces on demand for one render. A font.Face is not safe for concurrent use
// (opentype's faces share a rasteriser), so each render owns its set.
type faceSet struct {
	fonts []*opentype.Font
	faces map[faceKey]font.Face
}

type faceKey struct {
	face typeface
	px   float64
}

func newFaceSet() (*faceSet, error) {
	fonts, err := loadFonts()
	if err != nil {
		return nil, err
	}
	return &faceSet{fonts: fonts, faces: map[faceKey]font.Face{}}, nil
}

// get is the face for st at the device scale.
func (fs *faceSet) get(st style) (font.Face, error) {
	key := faceKey{st.face, st.size * scale}
	if f, ok := fs.faces[key]; ok {
		return f, nil
	}
	f, err := opentype.NewFace(fs.fonts[st.face], &opentype.FaceOptions{Size: key.px, DPI: 72, Hinting: font.HintingNone})
	if err != nil {
		return nil, fmt.Errorf("sheet: face at %.0fpx: %w", key.px, err)
	}
	fs.faces[key] = f
	return f, nil
}

// text is a run of one style: what draws and what measures.
type text struct {
	face     font.Face
	st       style
	tracking fixed.Int26_6 // device px between letters
}

func (fs *faceSet) text(st style) (text, error) {
	f, err := fs.get(st)
	if err != nil {
		return text{}, err
	}
	return text{face: f, st: st, tracking: fixed.Int26_6(math.Round(st.trackingEm * st.size * scale * 64))}, nil
}

func (t text) prepare(s string) string {
	if t.st.upper {
		return strings.ToUpper(s)
	}
	return s
}

// width is s's drawn advance in device px, letter spacing included (CSS adds it after every
// letter, the last one too).
func (t text) width(s string) float64 {
	s = t.prepare(s)
	var w fixed.Int26_6
	prev := rune(-1)
	for _, r := range s {
		if prev >= 0 {
			w += t.face.Kern(prev, r)
		}
		adv, _ := t.face.GlyphAdvance(r)
		w += adv + t.tracking
		prev = r
	}
	return float64(w) / 64
}

// baseline is the baseline's device y for a line box whose top is at device y top: CSS
// centres the face's content area (ascent plus descent) in the line height.
func (t text) baseline(top float64) float64 {
	m := t.face.Metrics()
	asc, desc := float64(m.Ascent)/64, float64(m.Descent)/64
	lh := t.st.lineHeight * scale
	return top + (lh-(asc+desc))/2 + asc
}

// draw puts s on dst with its left edge at device x and its line box's top at device y.
func (t text) draw(dst draw.Image, col color.Color, x, top float64, s string) {
	s = t.prepare(s)
	src := image.NewUniform(col)
	dot := fixed.Point26_6{X: fixed.Int26_6(math.Round(x * 64)), Y: fixed.Int26_6(math.Round(t.baseline(top) * 64))}
	prev := rune(-1)
	for _, r := range s {
		if prev >= 0 {
			dot.X += t.face.Kern(prev, r)
		}
		dr, mask, mp, adv, ok := t.face.Glyph(dot, r)
		if ok {
			draw.DrawMask(dst, dr, src, image.Point{}, mask, mp, draw.Over)
		}
		dot.X += adv + t.tracking
		prev = r
	}
}

// align is where a line sits in its box.
type align int

const (
	alignLeft align = iota
	alignRight
	alignCentre
)

// drawAligned draws s in the box from device x to x+w.
func (t text) drawAligned(dst draw.Image, col color.Color, x, w, top float64, a align, s string) {
	switch a {
	case alignRight:
		x += w - t.width(s)
	case alignCentre:
		x += (w - t.width(s)) / 2
	}
	t.draw(dst, col, x, top, s)
}

// wrap breaks s at spaces into lines no wider than maxW device px, greedily, as a browser
// does. fits is false when a single word is wider than maxW on its own.
func (t text) wrap(s string, maxW float64) (lines []string, fits bool) {
	words := strings.Fields(s)
	fits = true
	var line string
	for _, w := range words {
		if t.width(w) > maxW {
			fits = false
		}
		next := w
		if line != "" {
			next = line + " " + w
		}
		if line != "" && t.width(next) > maxW {
			lines = append(lines, line)
			line = w
			continue
		}
		line = next
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines, fits
}
