package sheet

import (
	"image"
	"image/color"
	"strings"
	"sync"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
)

// The crest thumbnail (brand audit 2026-11-10 hand-off D): the web's one crest (web
// styles/crest.css) as a PNG for Discord, which cannot draw CSS. A ring in the class colour at
// 55 % around a disc of the page ground tinted 12 % by the class, the 128 px HD crest clipped to
// the disc, no glow. A slug that is not a class draws the LX seal in a muted ring, never a person
// glyph. 128 px with a 4 px ring: the web's 2 px ring at Discord's ~64 px thumbnail, at 2x.
const (
	CrestSide = 128
	crestRing = 4
	crestDisc = CrestSide - 2*crestRing
	// crestRingAlpha and crestTint are crest.css's 55 % ring and 12 % disc.
	crestRingAlpha = 0.55
	crestTint      = 0.12
)

// CrestPath is the crest route's prefix, under the tiles' so it shares their rate-limit rules.
const CrestPath = TilePath + "crest/"

// crestNone is the file the LX fallback is served as.
const crestNone = "none"

// CrestURL is a class's ringed crest thumbnail: {apiBase}/v1/discord/tiles/crest/{slug}.png, the
// LX fallback (none.png) for a slug that is not a class.
func CrestURL(apiBase, classSlug string) string {
	if _, ok := ClassColor(classSlug); !ok {
		classSlug = crestNone
	}
	return strings.TrimRight(apiBase, "/") + CrestPath + classSlug + ".png"
}

var crestMasks = sync.OnceValues(func() (*image.Alpha, *image.Alpha) {
	return brand.RingMask(CrestSide, CrestSide, CrestSide/2, crestRing),
		brand.RoundRectMask(crestDisc, crestDisc, crestDisc/2)
})

// CrestPNG draws classSlug's ringed crest, or the LX seal in a muted ring for any other slug.
func CrestPNG(classSlug string) ([]byte, error) {
	img, err := CrestImage(classSlug)
	if err != nil {
		return nil, err
	}
	return encodePNG(img)
}

// CrestImage is the ringed crest as an image, for a card that composes it onto its own canvas:
// the only class image any Go surface draws (the bare art stays inside this package).
func CrestImage(classSlug string) (*image.NRGBA, error) {
	ring, known := ClassColor(classSlug)
	art, err := crestArt(classSlug, known)
	if err != nil {
		return nil, err
	}
	disc := colorEmptyFill
	if known {
		disc = brand.Over(colorEmptyFill, ring, crestTint)
	} else {
		ring = opaque(brand.ColorBaseMuted)
	}
	ringMask, discMask := crestMasks()
	img := image.NewNRGBA(image.Rect(0, 0, CrestSide, CrestSide))
	brand.FillMask(img, image.Point{}, ringMask, withAlpha(ring, crestRingAlpha))
	p := image.Pt(crestRing, crestRing)
	brand.FillMask(img, p, discMask, disc)
	brand.PlaceImage(img, p, crestDisc, crestDisc, art, discMask)
	return img, nil
}

func crestArt(classSlug string, known bool) (image.Image, error) {
	if known {
		return embeddedCrest(classSlug)
	}
	return seal()
}

// withAlpha is c at alpha a (0 to 1), non-premultiplied.
func withAlpha(c color.RGBA, a float64) color.NRGBA {
	return color.NRGBA{R: c.R, G: c.G, B: c.B, A: uint8(a*255 + 0.5)}
}
