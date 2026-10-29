package sheet

import (
	"image"
	"image/color"
	"strconv"
	"strings"
	"sync"
)

// The tile (spec §A.4): a 72 x 72 PNG, a 4 px ring in the quality colour around the icon at
// 64 x 64, 1:1, never resampled. The corners round as the mock draws them: 8 px outside, 2 px
// on the icon.
const (
	TileSide  = 72
	tileRing  = 4
	tileIcon  = TileSide - 2*tileRing
	tileOuter = 8
	tileInner = 2
)

// TilePath is the tiles route's prefix.
const TilePath = "/v1/discord/tiles/"

// TileURL is the tile's public URL: {apiBase}/v1/discord/tiles/{build}/{item}.png. It is the
// Thumbnail URL the panel's item rows carry.
func TileURL(apiBase, build string, itemID int) string {
	return strings.TrimRight(apiBase, "/") + TilePath + build + "/" + strconv.Itoa(itemID) + ".png"
}

var tileMasks = sync.OnceValues(func() (*image.Alpha, *image.Alpha) {
	return roundRectMask(TileSide, TileSide, tileOuter), roundRectMask(tileIcon, tileIcon, tileInner)
})

// Tile draws icon in its quality ring; an unknown quality gets the empty ring (#262e40).
func Tile(icon image.Image, quality int) ([]byte, error) {
	ring := colorEmptyRing
	if knownQuality(quality) {
		ring = qualityRing[quality]
	}
	return encodePNG(tileImage(icon, ring))
}

func tileImage(icon image.Image, ring color.RGBA) *image.NRGBA {
	outer, inner := tileMasks()
	img := image.NewNRGBA(image.Rect(0, 0, TileSide, TileSide))
	fillMask(img, image.Point{}, outer, ring)
	p := image.Pt(tileRing, tileRing)
	fillMask(img, p, inner, colorBlack)
	placeImage(img, p, tileIcon, tileIcon, icon, inner)
	return img
}
