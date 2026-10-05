// Package imagex turns a raw guild crest upload (PNG, JPEG, or WebP) into the
// contract's own crest shape (docs/contracts/2026-10-05-guild-crest-api.md): centre-cropped
// to square, resized to a fixed side, and encoded as WebP with alpha kept. The original is
// never stored - only ProcessCrest's output ever reaches R2.
package imagex

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"

	"github.com/HugoSmits86/nativewebp"
	ximage "golang.org/x/image/draw"
	"golang.org/x/image/webp"
)

const (
	// MinDimension is the contract's floor: an upload under this on either axis is
	// refused outright rather than upscaled into something soft.
	MinDimension = 128
	// MaxBytes is the contract's ceiling on the raw upload, before any decoding.
	MaxBytes = 2 << 20 // 2 MiB
	// MaxDimension bounds the decoded image's resolution - not a contract
	// requirement (the contract accepts "any aspect"), but a defensive ceiling
	// against a small, highly-compressed file that decodes to an enormous pixel
	// buffer. 8192 on a side is far beyond anything a real crest needs.
	MaxDimension = 8192
	// CrestSize is the square side every processed crest is resized to.
	CrestSize = 256
)

// ValidationError is a rejection the contract itself describes with a plain,
// user-facing message - the exact text PUT /v1/guilds/{id}/crest answers with,
// under the "invalid" error code.
type ValidationError struct{ Message string }

func (e *ValidationError) Error() string { return e.Message }

func reject(format string, args ...any) error {
	return &ValidationError{Message: fmt.Sprintf(format, args...)}
}

// Result is one processed crest: the WebP bytes ready to store, and their own
// sha256 - the content hash the stored object's key is named from.
type Result struct {
	WebP   []byte
	SHA256 string
}

// ProcessCrest validates and transforms a raw upload into the contract's crest.
// It never trusts the upload's declared content type: the bytes are sniffed and
// decoded directly. A rejection for type, size, or dimensions is a *ValidationError
// carrying the contract's own plain message; any other error is this package's own
// failure to encode, which the caller should treat as internal.
func ProcessCrest(data []byte) (Result, error) {
	if len(data) > MaxBytes {
		return Result{}, reject("That file is %s; the limit is %s.", formatMB(len(data)), formatWholeMB(MaxBytes))
	}
	img, err := decode(data)
	if err != nil {
		return Result{}, reject("That is not a PNG, JPEG, or WebP image.")
	}
	b := img.Bounds()
	w, h := b.Dx(), b.Dy()
	if w < MinDimension || h < MinDimension {
		return Result{}, reject("That image is %dx%d; the minimum is %dx%d.", w, h, MinDimension, MinDimension)
	}
	if w > MaxDimension || h > MaxDimension {
		return Result{}, reject("That image is %dx%d; the maximum is %dx%d.", w, h, MaxDimension, MaxDimension)
	}

	resized := resizeTo(cropToSquare(img), CrestSize)

	buf := &bytes.Buffer{}
	if err := nativewebp.Encode(buf, resized, nil); err != nil {
		return Result{}, fmt.Errorf("imagex: encode webp: %w", err)
	}
	sum := sha256.Sum256(buf.Bytes())
	return Result{WebP: buf.Bytes(), SHA256: hex.EncodeToString(sum[:])}, nil
}

// decode sniffs data's own magic bytes and decodes it as PNG, JPEG, or WebP -
// never by trusting a filename or a declared Content-Type, both of which a
// caller controls and neither of which this package is ever given anyway.
func decode(data []byte) (image.Image, error) {
	switch {
	case len(data) >= 8 && bytes.Equal(data[:8], []byte{0x89, 'P', 'N', 'G', '\r', '\n', 0x1a, '\n'}):
		return png.Decode(bytes.NewReader(data))
	case len(data) >= 2 && data[0] == 0xFF && data[1] == 0xD8:
		return jpeg.Decode(bytes.NewReader(data))
	case len(data) >= 12 && string(data[0:4]) == "RIFF" && string(data[8:12]) == "WEBP":
		return webp.Decode(bytes.NewReader(data))
	default:
		return nil, errors.New("imagex: unrecognized image type")
	}
}

// cropToSquare returns the centred square crop of img's longer axis - a plain
// translated copy, no resampling, so it introduces no blur of its own before
// resizeTo does the one resampling pass the contract calls for.
func cropToSquare(img image.Image) image.Image {
	b := img.Bounds()
	side := b.Dx()
	if b.Dy() < side {
		side = b.Dy()
	}
	x0 := b.Min.X + (b.Dx()-side)/2
	y0 := b.Min.Y + (b.Dy()-side)/2
	src := image.Rect(x0, y0, x0+side, y0+side)
	out := image.NewNRGBA(image.Rect(0, 0, side, side))
	draw.Draw(out, out.Bounds(), img, src.Min, draw.Src)
	return out
}

// resizeTo resamples img to a size x size square with the Catmull-Rom kernel -
// the contract's own choice, a good general-purpose filter for shrinking a crest
// down from whatever square the upload cropped to.
func resizeTo(img image.Image, size int) image.Image {
	dst := image.NewNRGBA(image.Rect(0, 0, size, size))
	ximage.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), ximage.Over, nil)
	return dst
}

// formatMB renders a byte count to one decimal place, the contract's own
// "3.1 MB" shape.
func formatMB(n int) string {
	return fmt.Sprintf("%.1f MB", float64(n)/(1<<20))
}

// formatWholeMB renders a byte count known to be a whole number of
// mebibytes (the fixed limits this package itself defines) without a
// trailing ".0" - the contract's own "2 MB", not "2.0 MB".
func formatWholeMB(n int) string {
	return fmt.Sprintf("%d MB", n/(1<<20))
}
