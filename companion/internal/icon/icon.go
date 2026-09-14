// companion/internal/icon/icon.go
// Package icon draws the tray icon rather than shipping a binary
// asset: one gold ring on transparency, in the site's accent colour,
// encoded as PNG for macOS and Linux and as a PNG inside an ICO
// container for Windows, which is the format its tray wants.
package icon

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/color"
	"image/png"
	"math"
	"sync"
)

// Size is the icon's edge in pixels. 32 is what every tray asks for
// and scales acceptably on a retina menu bar.
const Size = 32

// Gold is --color-gold from web/src/styles/tokens.css.
var Gold = color.NRGBA{R: 0xe5, G: 0xb9, B: 0x55, A: 0xff}

var (
	once sync.Once
	png_ []byte
	ico_ []byte
)

// draw paints a ring, antialiased by supersampling the coverage of
// each pixel four times over.
func draw() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, Size, Size))
	const outer, inner = Size/2 - 1.5, Size/2 - 6.5
	centre := float64(Size) / 2
	for y := range Size {
		for x := range Size {
			var hits int
			for _, dy := range []float64{0.25, 0.75} {
				for _, dx := range []float64{0.25, 0.75} {
					px, py := float64(x)+dx-centre, float64(y)+dy-centre
					d := math.Hypot(px, py)
					if d <= outer && d >= inner {
						hits++
					}
				}
			}
			if hits == 0 {
				continue
			}
			c := Gold
			c.A = uint8(hits * 255 / 4)
			img.SetNRGBA(x, y, c)
		}
	}
	return img
}

func build() {
	var buf bytes.Buffer
	if err := png.Encode(&buf, draw()); err != nil {
		panic(err) // encoding a 32×32 image cannot fail
	}
	png_ = buf.Bytes()

	// ICONDIR, one ICONDIRENTRY, then the PNG itself. Windows has
	// accepted PNG-in-ICO since Vista.
	var ico bytes.Buffer
	binary.Write(&ico, binary.LittleEndian, [3]uint16{0, 1, 1})
	ico.Write([]byte{Size, Size, 0, 0})
	binary.Write(&ico, binary.LittleEndian, [2]uint16{1, 32})
	binary.Write(&ico, binary.LittleEndian, uint32(len(png_)))
	binary.Write(&ico, binary.LittleEndian, uint32(22))
	ico.Write(png_)
	ico_ = ico.Bytes()
}

// PNG is the icon for macOS and Linux.
func PNG() []byte { once.Do(build); return png_ }

// ICO is the icon for Windows.
func ICO() []byte { once.Do(build); return ico_ }
