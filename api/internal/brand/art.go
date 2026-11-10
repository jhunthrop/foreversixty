package brand

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/image/vector"
	"golang.org/x/image/webp"
)

// The site's art, copied byte for byte by make sheet-assets (art_test.go compares the hashes).
const (
	sealFile   = "art/seal-400.png"
	groundFile = "art/ground.webp"
	bandFile   = "art/band-1200.webp"
	cornerFile = "art/corner-filigree.svg"
	linkFile   = "art/og-art.jpg"
)

// ArtSources maps each embedded file to the site file it copies, from the repo root.
var ArtSources = map[string]string{
	sealFile:   "web/public/seal-400.png",
	groundFile: "web/public/art/ground.webp",
	bandFile:   "web/public/art/band-1200.webp",
	cornerFile: "design/assets/ornaments/corner-filigree.svg",
	linkFile:   "web/public/art/og-art.jpg",
}

//go:embed art
var artFS embed.FS

// ArtFile is an embedded art file's bytes, for the hash test.
func ArtFile(name string) ([]byte, error) { return artFS.ReadFile(name) }

func decodeArt(name string) (image.Image, error) {
	raw, err := artFS.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("brand: read %s: %w", name, err)
	}
	var img image.Image
	switch {
	case strings.HasSuffix(name, ".webp"):
		img, err = webp.Decode(bytes.NewReader(raw))
	case strings.HasSuffix(name, ".jpg"):
		img, err = jpeg.Decode(bytes.NewReader(raw))
	default:
		img, err = png.Decode(bytes.NewReader(raw))
	}
	if err != nil {
		return nil, fmt.Errorf("brand: decode %s: %w", name, err)
	}
	return img, nil
}

// The decoded art, each decoded once and only ever read.
var (
	sealImage   = sync.OnceValues(func() (image.Image, error) { return decodeArt(sealFile) })
	groundImage = sync.OnceValues(func() (image.Image, error) { return decodeArt(groundFile) })
	bandImage   = sync.OnceValues(func() (image.Image, error) { return decodeArt(bandFile) })
	linkImage   = sync.OnceValues(func() (image.Image, error) { return decodeArt(linkFile) })
)

// Seal is the LX seal v4 at 400 px, with real transparency.
func Seal() (image.Image, error) { return sealImage() }

// LinkArt is the link preview's art, 1200 x 630: the owner's valley road under its glow, cropped
// by tools/art/build_art.py.
func LinkArt() (image.Image, error) { return linkImage() }

// cornerShapes is the corner ornament's outline: closed polygons in its 1761-unit view box.
// The traced file draws with straight segments only (M, L, Z).
var cornerShapes = sync.OnceValues(func() ([][][2]float32, error) {
	raw, err := artFS.ReadFile(cornerFile)
	if err != nil {
		return nil, fmt.Errorf("brand: read %s: %w", cornerFile, err)
	}
	m := regexp.MustCompile(` d="([^"]+)"`).FindSubmatch(raw)
	if m == nil {
		return nil, fmt.Errorf("brand: %s has no path", cornerFile)
	}
	return parsePolygons(string(m[1]))
})

// cornerView is the ornament's view box side.
const cornerView = 1761

var pathToken = regexp.MustCompile(`[MLZ]|-?\d+(?:\.\d+)?`)

func parsePolygons(d string) ([][][2]float32, error) {
	var out [][][2]float32
	var cur [][2]float32
	tokens := pathToken.FindAllString(d, -1)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "Z":
			if len(cur) > 0 {
				out = append(out, cur)
			}
			cur = nil
		case "M", "L":
			if i+2 >= len(tokens) {
				return nil, fmt.Errorf("brand: corner path ends inside a %s", tokens[i])
			}
			x, errX := strconv.ParseFloat(tokens[i+1], 32)
			y, errY := strconv.ParseFloat(tokens[i+2], 32)
			if errX != nil || errY != nil {
				return nil, fmt.Errorf("brand: corner path: bad point %s,%s", tokens[i+1], tokens[i+2])
			}
			cur = append(cur, [2]float32{float32(x), float32(y)})
			i += 2
		default:
			return nil, fmt.Errorf("brand: corner path: unexpected %q", tokens[i])
		}
	}
	if len(cur) > 0 {
		out = append(out, cur)
	}
	return out, nil
}

// Corner is which corner of a panel an ornament sits in; the file draws the top left, and the
// others mirror it, as the site's FrameCorners does.
type Corner int

// The four corners.
const (
	TopLeft Corner = iota
	TopRight
	BottomRight
	BottomLeft
)

type cornerKey struct {
	side   int
	corner Corner
}

var (
	cornerMu    sync.Mutex
	cornerCache = map[cornerKey]*image.Alpha{}
)

// CornerMask is the ornament's coverage in a side x side square, mirrored for corner.
func CornerMask(side int, corner Corner) (*image.Alpha, error) {
	cornerMu.Lock()
	defer cornerMu.Unlock()
	key := cornerKey{side, corner}
	if m, ok := cornerCache[key]; ok {
		return m, nil
	}
	shapes, err := cornerShapes()
	if err != nil {
		return nil, err
	}
	k := float32(side) / cornerView
	flipX := corner == TopRight || corner == BottomRight
	flipY := corner == BottomRight || corner == BottomLeft
	pt := func(p [2]float32) (float32, float32) {
		x, y := p[0]*k, p[1]*k
		if flipX {
			x = float32(side) - x
		}
		if flipY {
			y = float32(side) - y
		}
		return x, y
	}
	z := vector.NewRasterizer(side, side)
	for _, poly := range shapes {
		z.MoveTo(pt(poly[0]))
		for _, p := range poly[1:] {
			z.LineTo(pt(p))
		}
		z.ClosePath()
	}
	m := image.NewAlpha(image.Rect(0, 0, side, side))
	z.Draw(m, m.Bounds(), image.Opaque, image.Point{})
	cornerCache[key] = m
	return m, nil
}
