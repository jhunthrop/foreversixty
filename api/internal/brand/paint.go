package brand

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	xdraw "golang.org/x/image/draw"
)

// supersample is the anti-aliasing grid per pixel for drawn shapes: 4 x 4 samples, enough that
// edges match a browser's to within the golden tolerances.
const supersample = 4

// RGB is c as Discord's integer colour (an embed's color, a container's accent_color).
func RGB(c color.NRGBA) int { return int(c.R)<<16 | int(c.G)<<8 | int(c.B) }

// RoundRectMask is the coverage of a w x h rectangle with corner radius r (device px).
func RoundRectMask(w, h int, r float64) *image.Alpha {
	m := image.NewAlpha(image.Rect(0, 0, w, h))
	paintRoundRect(m, 0, 0, float64(w), float64(h), r, +1)
	return m
}

// RingMask is the coverage of a line `width` wide just inside the edge of a w x h rectangle
// with corner radius r: a rounded outline, the outer shape less the inner one.
func RingMask(w, h int, r, width float64) *image.Alpha {
	m := RoundRectMask(w, h, r)
	paintRoundRect(m, width, width, float64(w)-width, float64(h)-width, math.Max(r-width, 0), -1)
	return m
}

// paintRoundRect adds (sign +1) or removes (sign -1) a rounded rectangle's coverage in m. Only
// the pixels on its edge are sampled; the rest are wholly in or out, which keeps a card-sized
// panel cheap.
func paintRoundRect(m *image.Alpha, x0, y0, x1, y1, r float64, sign int) {
	b := m.Bounds()
	for py := b.Min.Y; py < b.Max.Y; py++ {
		fy := float64(py)
		for px := b.Min.X; px < b.Max.X; px++ {
			fx := float64(px)
			if fx+1 <= x0 || fx >= x1 || fy+1 <= y0 || fy >= y1 {
				continue
			}
			cov := 255
			whole := (fx >= x0 && fx+1 <= x1 && fy >= y0+r && fy+1 <= y1-r) ||
				(fx >= x0+r && fx+1 <= x1-r && fy >= y0 && fy+1 <= y1)
			if !whole {
				cov = sampleCoverage(fx, fy, func(x, y float64) bool { return inRoundRect(x, y, x0, y0, x1, y1, r) })
			}
			i := m.PixOffset(px, py)
			m.Pix[i] = uint8(max(0, min(255, int(m.Pix[i])+sign*cov)))
		}
	}
}

// sampleCoverage is the share of the pixel at (px, py) inside, on the supersample grid, 0 to 255.
func sampleCoverage(px, py float64, inside func(x, y float64) bool) int {
	const n = supersample
	hits := 0
	for sy := range n {
		for sx := range n {
			if inside(px+(float64(sx)+0.5)/n, py+(float64(sy)+0.5)/n) {
				hits++
			}
		}
	}
	return (hits*255 + n*n/2) / (n * n)
}

func inRoundRect(x, y, x0, y0, x1, y1, r float64) bool {
	if x < x0 || y < y0 || x > x1 || y > y1 {
		return false
	}
	cx := math.Max(x0+r, math.Min(x1-r, x))
	cy := math.Max(y0+r, math.Min(y1-r, y))
	dx, dy := x-cx, y-cy
	return dx*dx+dy*dy <= r*r
}

// CoverageMask samples inside on a supersample grid in each pixel.
func CoverageMask(w, h int, inside func(x, y float64) bool) *image.Alpha {
	m := image.NewAlpha(image.Rect(0, 0, w, h))
	for py := range h {
		for px := range w {
			m.Pix[py*m.Stride+px] = uint8(sampleCoverage(float64(px), float64(py), inside))
		}
	}
	return m
}

// FillMask paints col through mask with its top-left at p.
func FillMask(dst draw.Image, p image.Point, mask *image.Alpha, col color.Color) {
	r := mask.Bounds().Add(p)
	draw.DrawMask(dst, r, image.NewUniform(col), image.Point{}, mask, image.Point{}, draw.Over)
}

// FillRect paints an axis-aligned rectangle.
func FillRect(dst draw.Image, r image.Rectangle, col color.Color) {
	draw.Draw(dst, r, image.NewUniform(col), image.Point{}, draw.Over)
}

// PlaceImage draws src into the w x h box at p: 1:1 when the sizes match, else scaled with
// Catmull-Rom, through clip when one is given.
func PlaceImage(dst draw.Image, p image.Point, w, h int, src image.Image, clip *image.Alpha) {
	box := image.Rect(0, 0, w, h)
	tile := src
	if src.Bounds().Dx() != w || src.Bounds().Dy() != h {
		scaled := image.NewNRGBA(box)
		xdraw.CatmullRom.Scale(scaled, box, src, src.Bounds(), xdraw.Src, nil)
		tile = scaled
	}
	sp := tile.Bounds().Min
	if clip == nil {
		draw.Draw(dst, box.Add(p), tile, sp, draw.Over)
		return
	}
	draw.DrawMask(dst, box.Add(p), tile, sp, clip, image.Point{}, draw.Over)
}

// Ground fills r with the brown ground: the site's tile at its own pixels, repeated, on the
// ground colour.
func Ground(dst draw.Image, r image.Rectangle) error {
	tile, err := groundImage()
	if err != nil {
		return err
	}
	FillRect(dst, r, ColorForeverGround)
	tb := tile.Bounds()
	for y := r.Min.Y; y < r.Max.Y; y += tb.Dy() {
		for x := r.Min.X; x < r.Max.X; x += tb.Dx() {
			cell := image.Rect(x, y, x+tb.Dx(), y+tb.Dy()).Intersect(r)
			draw.Draw(dst, cell, tile, tb.Min, draw.Src)
		}
	}
	return nil
}

// bandArtTop is how far down the painted band its first visible row sits, as a share of the
// art's width: the site draws it 96 px up at 1440 px wide (ui/band.css).
const bandArtTop = 96.0 / 1440

// Band paints the section band into r: the owner's blue band art stretched to r's width, under
// the band veil, on the band colour, with the band's 1 px black foot.
func Band(dst draw.Image, r image.Rectangle) error {
	art, err := bandImage()
	if err != nil {
		return err
	}
	FillRect(dst, r, ColorForeverBand)
	ab := art.Bounds()
	w := r.Dx()
	h := int(math.Round(float64(ab.Dy()) * float64(w) / float64(ab.Dx())))
	scaled := image.NewNRGBA(image.Rect(0, 0, w, h))
	xdraw.CatmullRom.Scale(scaled, scaled.Bounds(), art, ab, xdraw.Src, nil)
	top := int(math.Round(bandArtTop * float64(w)))
	draw.Draw(dst, r, scaled, image.Pt(0, top), draw.Over)
	FillRect(dst, r, ColorForeverVeilBand)
	FillRect(dst, image.Rect(r.Min.X, r.Max.Y-1, r.Max.X, r.Max.Y), color.Black)
	return nil
}

// PanelStyle is how a panel is dressed: Framed draws the cream inset line; Ornate adds the
// corner ornaments and implies Framed. Spec: two or three framed panels a page, one ornate.
type PanelStyle struct {
	Framed, Ornate bool
	// CornerSide is the ornament's side in CSS px: 0 is the site's 64, and 44 its small form.
	CornerSide float64
}

// The panel's measures in CSS px (ui/panel.css): the frame line's inset, the ornament's inset
// and side, and the ornament's opacity.
const (
	frameInset   = 6
	cornerInset  = 4
	cornerSide   = 64
	cornerAlpha  = 0.8
	frameRadiusD = 4 // the frame line's radius is the panel's less this
)

// Panel draws the round-5 surface into r at scale s (device px per CSS px): the card brown-black
// (raised), a black hairline inside a warm outer ring, the plate radius, and the style's frame
// and corners. The outer ring sits just outside r, as the site's box-shadow does.
func Panel(dst draw.Image, r image.Rectangle, s float64, st PanelStyle) error {
	rad := RadiusPlate * s
	line := math.Max(1, math.Round(StrokeHairline*s))
	ring := r.Inset(-int(line))
	FillMask(dst, ring.Min, RoundRectMask(ring.Dx(), ring.Dy(), rad+line), ColorForeverPanelEdge)
	// The edge ring is drawn over the ground; the black hairline and the fill replace it inside.
	FillMask(dst, r.Min, RoundRectMask(r.Dx(), r.Dy(), rad), color.Black)
	in := r.Inset(int(line))
	FillMask(dst, in.Min, RoundRectMask(in.Dx(), in.Dy(), rad-line), ColorBaseRaised)
	if !st.Framed && !st.Ornate {
		return nil
	}
	fr := r.Inset(int(math.Round(frameInset * s)))
	FillMask(dst, fr.Min, RingMask(fr.Dx(), fr.Dy(), (RadiusPlate-frameRadiusD)*s, line), ColorForeverFrameLine)
	if !st.Ornate {
		return nil
	}
	side := st.CornerSide
	if side == 0 {
		side = cornerSide
	}
	return corners(dst, r, s, side)
}

// corners draws the four ornaments, side CSS px, in the cream at the ornament's opacity.
func corners(dst draw.Image, r image.Rectangle, s float64, side float64) error {
	px := int(math.Round(side * s))
	in := int(math.Round(cornerInset * s))
	ink := ColorForeverCream
	ink.A = uint8(math.Round(cornerAlpha * 255))
	at := map[Corner]image.Point{
		TopLeft:     {r.Min.X + in, r.Min.Y + in},
		TopRight:    {r.Max.X - in - px, r.Min.Y + in},
		BottomRight: {r.Max.X - in - px, r.Max.Y - in - px},
		BottomLeft:  {r.Min.X + in, r.Max.Y - in - px},
	}
	for c, p := range at {
		m, err := CornerMask(px, c)
		if err != nil {
			return err
		}
		FillMask(dst, p, m, ink)
	}
	return nil
}

// Over blends c at alpha a (0 to 1) onto base, opaque.
func Over(base, c color.Color, a float64) color.RGBA {
	b, f := color.RGBAModel.Convert(base).(color.RGBA), color.RGBAModel.Convert(c).(color.RGBA)
	mix := func(x, y uint8) uint8 { return uint8(math.Round(float64(x)*(1-a) + float64(y)*a)) }
	return color.RGBA{R: mix(b.R, f.R), G: mix(b.G, f.G), B: mix(b.B, f.B), A: 0xff}
}
