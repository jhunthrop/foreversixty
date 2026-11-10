// Package brand is the round-5 look for everything the API draws as an image (design/specs/
// 2026-11-09-directions.md): the Discord sheet and item tiles (internal/sheet) and the build and
// report link previews (internal/card). It holds one copy of each thing those renderers share:
//
//   - tokens_gen.go, the site's colours, radii, strokes and spaces, generated from
//     design/tokens/tokens.json (node design/tokens/build.mjs), never copied by hand;
//   - fonts/, the site's own subset faces, Open Sans 400/600/700 and Marcellus 400, as TTF
//     (make brand-fonts);
//   - art/, the seal v4, the brown ground tile, the blue band and the lead panel's corner
//     ornament, copied byte for byte from the site (make sheet-assets);
//   - the paint those share: the ground, the band, the panel with its frame and corners.
package brand

import (
	"embed"
	"fmt"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed fonts/*.ttf
var fontFS embed.FS

// Face is one of the site's faces.
type Face int

// The faces: Open Sans for every word, Marcellus for titles and big figures.
const (
	OpenSans400 Face = iota
	OpenSans600
	OpenSans700
	Marcellus400
)

var faceFiles = [...]string{
	OpenSans400:  "fonts/OpenSans-Regular.ttf",
	OpenSans600:  "fonts/OpenSans-SemiBold.ttf",
	OpenSans700:  "fonts/OpenSans-Bold.ttf",
	Marcellus400: "fonts/Marcellus-Regular.ttf",
}

// parsed holds every face parsed once; a parsed *opentype.Font is read-only and shared.
var parsed = sync.OnceValues(func() ([]*opentype.Font, error) {
	out := make([]*opentype.Font, len(faceFiles))
	for i, name := range faceFiles {
		raw, err := fontFS.ReadFile(name)
		if err != nil {
			return nil, fmt.Errorf("brand: read %s: %w", name, err)
		}
		f, err := opentype.Parse(raw)
		if err != nil {
			return nil, fmt.Errorf("brand: parse %s: %w", name, err)
		}
		out[i] = f
	}
	return out, nil
})

// FontFile is face f's TTF bytes, for an HTML page that inlines the site's own faces.
func FontFile(f Face) ([]byte, error) {
	if f < 0 || int(f) >= len(faceFiles) {
		return nil, fmt.Errorf("brand: no face %d", f)
	}
	raw, err := fontFS.ReadFile(faceFiles[f])
	if err != nil {
		return nil, fmt.Errorf("brand: read %s: %w", faceFiles[f], err)
	}
	return raw, nil
}

// Font is the parsed face f.
func Font(f Face) (*opentype.Font, error) {
	fonts, err := parsed()
	if err != nil {
		return nil, err
	}
	if f < 0 || int(f) >= len(fonts) {
		return nil, fmt.Errorf("brand: no face %d", f)
	}
	return fonts[f], nil
}

// NewFace is f at px device pixels. A font.Face is not safe for concurrent use (opentype's
// faces share a rasteriser), so each render builds its own; the parse is the shared, costly
// half.
func NewFace(f Face, px float64) (font.Face, error) {
	src, err := Font(f)
	if err != nil {
		return nil, err
	}
	face, err := opentype.NewFace(src, &opentype.FaceOptions{Size: px, DPI: 72, Hinting: font.HintingNone})
	if err != nil {
		return nil, fmt.Errorf("brand: face at %.0fpx: %w", px, err)
	}
	return face, nil
}
