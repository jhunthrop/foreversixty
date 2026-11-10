package sheet

import (
	"bytes"
	"fmt"
	"image"
	"image/png"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
)

// discMask is the coverage of a disc of radius r centred in a side x side square.
func discMask(side int, r float64) *image.Alpha {
	c := float64(side) / 2
	return brand.CoverageMask(side, side, func(x, y float64) bool {
		dx, dy := x-c, y-c
		return dx*dx+dy*dy <= r*r
	})
}

func encodePNG(img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	enc := png.Encoder{CompressionLevel: png.BestSpeed}
	if err := enc.Encode(&buf, img); err != nil {
		return nil, fmt.Errorf("sheet: encode png: %w", err)
	}
	return buf.Bytes(), nil
}
