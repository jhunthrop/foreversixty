package sheet

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"

	xdraw "golang.org/x/image/draw"
)

// supersample is the anti-aliasing grid per pixel for the drawn shapes (rounded rectangles,
// discs): 4 x 4 samples, enough that edges match a browser's to within the golden tolerance.
const supersample = 4

// roundRectMask is the coverage of a w x h rectangle with corner radius r (device px).
func roundRectMask(w, h int, r float64) *image.Alpha {
	return coverageMask(w, h, func(x, y float64) bool {
		cx := math.Max(r, math.Min(float64(w)-r, x))
		cy := math.Max(r, math.Min(float64(h)-r, y))
		dx, dy := x-cx, y-cy
		return dx*dx+dy*dy <= r*r
	})
}

// discMask is the coverage of a disc of radius r centred in a side x side square.
func discMask(side int, r float64) *image.Alpha {
	c := float64(side) / 2
	return coverageMask(side, side, func(x, y float64) bool {
		dx, dy := x-c, y-c
		return dx*dx+dy*dy <= r*r
	})
}

// coverageMask samples inside on a supersample grid in each pixel.
func coverageMask(w, h int, inside func(x, y float64) bool) *image.Alpha {
	m := image.NewAlpha(image.Rect(0, 0, w, h))
	const n = supersample
	for py := range h {
		for px := range w {
			hits := 0
			for sy := range n {
				for sx := range n {
					if inside(float64(px)+(float64(sx)+0.5)/n, float64(py)+(float64(sy)+0.5)/n) {
						hits++
					}
				}
			}
			m.Pix[py*m.Stride+px] = uint8((hits*255 + n*n/2) / (n * n))
		}
	}
	return m
}

// fillMask paints col through mask with its top-left at p.
func fillMask(dst draw.Image, p image.Point, mask *image.Alpha, col color.Color) {
	r := mask.Bounds().Add(p)
	draw.DrawMask(dst, r, image.NewUniform(col), image.Point{}, mask, image.Point{}, draw.Over)
}

// fillRect paints an axis-aligned rectangle.
func fillRect(dst draw.Image, r image.Rectangle, col color.Color) {
	draw.Draw(dst, r, image.NewUniform(col), image.Point{}, draw.Over)
}

// placeImage draws src into the w x h box at p: 1:1 when the sizes match, else scaled down
// with Catmull-Rom (game art is never scaled up, spec B4), through clip when one is given.
func placeImage(dst draw.Image, p image.Point, w, h int, src image.Image, clip *image.Alpha) {
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

// over blends c at alpha a (0 to 1) onto base.
func over(base, c color.RGBA, a float64) color.RGBA {
	mix := func(b, f uint8) uint8 { return uint8(math.Round(float64(b)*(1-a) + float64(f)*a)) }
	return color.RGBA{R: mix(base.R, c.R), G: mix(base.G, c.G), B: mix(base.B, c.B), A: 0xff}
}

func encodePNG(img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	enc := png.Encoder{CompressionLevel: png.BestSpeed}
	if err := enc.Encode(&buf, img); err != nil {
		return nil, fmt.Errorf("sheet: encode png: %w", err)
	}
	return buf.Bytes(), nil
}
