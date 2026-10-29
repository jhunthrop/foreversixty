package sheet

import (
	"image"
	"image/color"
)

// The quality dot (spec §A.4.1) at Discord's 128 px emoji size: the spec's 64 px drawing at
// scale 2, a filled circle of diameter 72 in the quality colour with a 4 px outline of
// #000000 at 45% opacity centred on its edge, so Common's white dot holds on Discord's light theme.
const (
	dotSide    = 128
	dotRadius  = 36
	dotOutline = 4
	dotShadeA  = 115 // 45% of 255
)

// QualityDot is one quality's emoji: its application-emoji name and PNG.
type QualityDot struct {
	Name string
	PNG  []byte
}

// QualityDots draws the six quality dots, Poor to Legendary.
func QualityDots() ([]QualityDot, error) {
	out := make([]QualityDot, 0, len(qualityRing))
	for q := range qualityRing {
		png, err := encodePNG(dotImage(qualityRing[q]))
		if err != nil {
			return nil, err
		}
		out = append(out, QualityDot{Name: QualityEmojiName(q), PNG: png})
	}
	return out, nil
}

func dotImage(fill color.RGBA) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, dotSide, dotSide))
	shade := color.NRGBA{A: dotShadeA}
	fillMask(img, image.Point{}, discMask(dotSide, dotRadius+dotOutline/2), shade)
	fillMask(img, image.Point{}, discMask(dotSide, dotRadius-dotOutline/2), fill)
	return img
}
