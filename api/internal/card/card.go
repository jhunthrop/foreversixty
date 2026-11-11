// Package card draws the 1200x630 preview images shown when a build or a report link is
// unfurled (Discord, Slack, any Open Graph reader). They are the site's link preview in the
// round-5 look (web/src/lib/og.ts, design/specs/2026-11-09-directions.md): the owner's
// road-and-glow art under the night veil, the painted seal and "Forever Sixty" in Marcellus at
// the top, then the build's or the report's own lines over the fan-made line. The faces, colours
// and art are internal/brand's, the site's own.
//
// It draws directly onto an RGBA image with golang.org/x/image rather than going through SVG:
// the site's renderer uses satori and resvg, a JavaScript toolchain this service has no reason
// to carry.
package card

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	// Width and Height are the Open Graph card size the contract fixes.
	Width  = 1200
	Height = 630

	marginX = 64 // the site card's padding
	marginY = 64

	sealPage  = 88  // the seal on a build or report card
	sealBrand = 116 // the seal on the brand card
	sealGap   = 22

	// The night veil over the art: the site preview's (art_crops.OG_VEIL), the night token at 42%.
	veilAlpha = 0.42
	// The scrim under the lines: the night token from y 220 down to 45% at the foot.
	scrimTop   = 220
	scrimAlpha = 0.45
)

// The copy every card carries (web/src/lib/og.ts): the site's name and the fan-made line, so a
// shared link never reads as an official page.
const (
	copyBrand   = "Forever Sixty"
	copyFanLine = "Fan-made player tool for WoW: Forever"
	copyBuild   = "Talent build"
	copyReport  = "Raid report"
	lineJoin    = " · "
)

// The inks, all site tokens.
var (
	inkCream  = brand.ColorForeverCream
	inkCream2 = brand.ColorForeverCream2
	inkText   = brand.ColorBaseText
	inkKill   = brand.ColorOutcomeKill
	inkWipe   = brand.ColorOutcomeWipe
)

// Input is everything a build card shows. Split is points per tree in client order and Trees
// their names in the same order; Level is the build's level as builds.Describe computes it;
// ClassSlug picks the crest ("" draws none); DPS is the simmed DPS line, "" when unsimmed.
type Input struct {
	Title      string
	Race       string
	Class      string
	ClassSlug  string
	ClassColor string
	Split      []int
	Trees      []string
	Level      int
	DPS        string
}

// The build card's crest, the site's ringed HD crest, beside the title and the line under it.
const (
	crestPx  = 112
	crestGap = 28
)

// faces is one render's set. A font.Face may not be shared across goroutines (opentype's
// carries a rasteriser every draw overwrites), so each render builds its own from the shared
// parsed fonts; the parse is the costly half and brand does it once.
type faces struct {
	brand    font.Face // "Forever Sixty" beside the seal
	brandBig font.Face // the brand card's name
	title    font.Face // the build's or report's title
	titleSm  font.Face // a title too long for the large face
	lead     font.Face // the line under the title
	body     font.Face // the bosses, the fan-made line
	label    font.Face // "Killed", "Wiped"
}

func newFaces() (*faces, error) {
	f := &faces{}
	for _, spec := range []struct {
		dst  *font.Face
		face brand.Face
		px   float64
	}{
		{&f.brand, brand.Marcellus400, 40},
		{&f.brandBig, brand.Marcellus400, 64},
		{&f.title, brand.Marcellus400, 72},
		{&f.titleSm, brand.Marcellus400, 54},
		{&f.lead, brand.OpenSans600, 34},
		{&f.body, brand.OpenSans400, 26},
		{&f.label, brand.OpenSans700, 26},
	} {
		face, err := brand.NewFace(spec.face, spec.px)
		if err != nil {
			return nil, fmt.Errorf("card: %w", err)
		}
		*spec.dst = face
	}
	return f, nil
}

// Render draws the card for one build. It is safe to call concurrently.
func Render(in Input) ([]byte, error) {
	f, err := newFaces()
	if err != nil {
		return nil, err
	}
	img, err := canvas(f)
	if err != nil {
		return nil, err
	}
	title := strings.TrimSpace(in.Title)
	if title == "" {
		title = copyBuild
	}
	// The lines from the foot up: the fan-made line, the trees, the race, class and level, the
	// title; the crest stands beside the title and the line under it.
	const fanY, treesY, leadY, titleY = Height - marginY - 34, Height - marginY - 92, Height - marginY - 138, Height - marginY - 204
	drawText(img, f.body, inkText, marginX, fanY, fit(f.body, copyBuild+lineJoin+copyFanLine, Width-2*marginX))
	x := marginX
	if crest, err := sheet.EmbeddedCrest(in.ClassSlug); err == nil {
		top := (titleY - 54 + leadY - crestPx + 8) / 2
		brand.PlaceImage(img, image.Pt(marginX, top), crestPx, crestPx, crest, nil)
		x += crestPx + crestGap
	}
	maxW := Width - marginX - x
	drawText(img, f.body, inkCream2, x, treesY, fit(f.body, treesText(in.Split, in.Trees), maxW))
	who := strings.TrimSpace(in.Race + " " + in.Class)
	stats := joinNonEmpty(lineJoin, "Level "+strconv.Itoa(in.Level), in.DPS)
	lx := x
	if who != "" {
		lx = drawText(img, f.lead, parseHexColor(in.ClassColor), lx, leadY, fit(f.lead, who, maxW/2))
		lx = drawText(img, f.lead, inkCream2, lx, leadY, lineJoin)
	}
	drawText(img, f.lead, inkCream2, lx, leadY, fit(f.lead, stats, Width-marginX-lx))
	drawTitle(img, f, x, titleY, title, maxW)
	return encode(img)
}

// treesText is the talent split as players read it, each tree named: "31 Arms · 0 Fury · 20
// Protection"; the bare split ("31/0/20") when the trees are not known.
func treesText(split []int, trees []string) string {
	if len(trees) != len(split) || len(split) == 0 {
		return splitText(split)
	}
	parts := make([]string, len(split))
	for i, n := range split {
		parts[i] = strconv.Itoa(n) + " " + trees[i]
	}
	return strings.Join(parts, lineJoin)
}

// joinNonEmpty joins the parts that are not empty.
func joinNonEmpty(sep string, parts ...string) string {
	kept := parts[:0:0]
	for _, p := range parts {
		if p != "" {
			kept = append(kept, p)
		}
	}
	return strings.Join(kept, sep)
}

// fallback is rendered once and reused: it is the image served when a card cannot be drawn,
// so it must never itself fail at request time.
var fallback = sync.OnceValue(func() []byte {
	out, err := brandCard()
	if err != nil {
		return nil
	}
	return out
})

// Fallback is the brand card, served when a build's or report's own card cannot be drawn: the
// seal beside the name, the fan-made line under it, at the foot of the art. It returns nil only
// when the embedded fonts or art cannot be used at all, which the package's tests rule out.
func Fallback() []byte { return fallback() }

func brandCard() ([]byte, error) {
	f, err := newFaces()
	if err != nil {
		return nil, err
	}
	img, err := background()
	if err != nil {
		return nil, err
	}
	top := Height - marginY - sealBrand
	if err := drawSeal(img, marginX, top, sealBrand); err != nil {
		return nil, err
	}
	x := marginX + sealBrand + 26
	drawText(img, f.brandBig, inkCream2, x, top+62, copyBrand)
	drawText(img, f.body, inkText, x, top+106, copyFanLine)
	return encode(img)
}

// canvas is the art under the veil with the seal and the site's name across the top.
func canvas(f *faces) (*image.RGBA, error) {
	img, err := background()
	if err != nil {
		return nil, err
	}
	if err := drawSeal(img, marginX, marginY, sealPage); err != nil {
		return nil, err
	}
	drawText(img, f.brand, inkCream2, marginX+sealPage+sealGap, marginY+sealPage/2+14, copyBrand)
	return img, nil
}

func background() (*image.RGBA, error) {
	art, err := brand.LinkArt()
	if err != nil {
		return nil, err
	}
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	brand.PlaceImage(img, image.Point{}, Width, Height, art, nil)
	veil := brand.ColorBaseNight
	veil.A = uint8(math.Round(veilAlpha * 255))
	brand.FillRect(img, img.Bounds(), veil)
	scrim(img)
	return img, nil
}

// scrim darkens the art's lower part under the card's lines, from nothing at scrimTop to
// scrimAlpha at the foot, so the words hold over a bright patch of the road.
func scrim(img *image.RGBA) {
	shade := brand.ColorBaseNight
	for y := scrimTop; y < Height; y++ {
		shade.A = uint8(math.Round(scrimAlpha * 255 * float64(y-scrimTop) / float64(Height-scrimTop)))
		brand.FillRect(img, image.Rect(0, y, Width, y+1), shade)
	}
}

func drawSeal(img draw.Image, x, y, px int) error {
	seal, err := brand.Seal()
	if err != nil {
		return err
	}
	brand.PlaceImage(img, image.Pt(x, y), px, px, seal, nil)
	return nil
}

func encode(img *image.RGBA) ([]byte, error) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, fmt.Errorf("card: encode: %w", err)
	}
	return buf.Bytes(), nil
}

// drawTitle draws the title in the large face, or the smaller one when it does not fit, cut
// with "…" only when even that runs out of room.
func drawTitle(img draw.Image, f *faces, x, baseline int, title string, maxW int) {
	face := f.title
	if measure(face, title) > maxW {
		face = f.titleSm
	}
	drawText(img, face, inkCream, x, baseline, fit(face, title, maxW))
}

// drawText draws s with its left edge at x on the baseline, and returns the x it ends at.
func drawText(dst draw.Image, face font.Face, col color.Color, x, baseline int, s string) int {
	d := &font.Drawer{Dst: dst, Src: image.NewUniform(col), Face: face, Dot: fixed.P(x, baseline)}
	d.DrawString(s)
	return d.Dot.X.Ceil()
}

// splitText renders the per-tree point counts as players write them: 21/30/0.
func splitText(split []int) string {
	if len(split) == 0 {
		return "0"
	}
	parts := make([]string, len(split))
	for i, n := range split {
		parts[i] = strconv.Itoa(n)
	}
	return strings.Join(parts, "/")
}

// measure is the drawn width of s in a face, in pixels.
func measure(face font.Face, s string) int {
	return font.MeasureString(face, s).Ceil()
}

// ellipsis ends a trimmed line, the one glyph the rest of the product uses (every face carries
// U+2026).
const ellipsis = "…"

// fit trims s with an ellipsis until it draws within maxWidth.
func fit(face font.Face, s string, maxWidth int) string {
	s = strings.TrimSpace(s)
	limit := fixed.I(maxWidth)
	if font.MeasureString(face, s) <= limit {
		return s
	}
	r := []rune(s)
	for len(r) > 0 {
		r = r[:len(r)-1]
		if font.MeasureString(face, string(r)+ellipsis) <= limit {
			return strings.TrimRight(string(r), " ") + ellipsis
		}
	}
	return ""
}

// parseHexColor reads "#rrggbb" and falls back to the cream.
func parseHexColor(s string) color.NRGBA {
	if len(s) == 7 && s[0] == '#' {
		if v, err := strconv.ParseUint(s[1:], 16, 32); err == nil {
			return color.NRGBA{R: uint8(v >> 16), G: uint8(v >> 8), B: uint8(v), A: 0xff}
		}
	}
	return inkCream
}
