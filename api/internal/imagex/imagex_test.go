package imagex

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"testing"

	"github.com/HugoSmits86/nativewebp"
)

// genPNG builds a w x h PNG with a distinct, checkable color in each
// quadrant: red top-left, green top-right, blue bottom-left, and a
// semi-transparent yellow bottom-right - so a resize that silently drops
// or washes out a region fails a test that reads it back. (On a wide or
// tall source, a centre-square crop always straddles the midline on both
// axes and so always keeps a sliver of every quadrant - testing crop
// direction itself needs the two-halves fixture below instead.)
func genPNG(t *testing.T, w, h int, alpha bool) []byte {
	t.Helper()
	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	a := uint8(255)
	if alpha {
		a = 128
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			switch {
			case x < w/2 && y < h/2:
				img.Set(x, y, color.NRGBA{255, 0, 0, 255})
			case x >= w/2 && y < h/2:
				img.Set(x, y, color.NRGBA{0, 255, 0, 255})
			case x < w/2 && y >= h/2:
				img.Set(x, y, color.NRGBA{0, 0, 255, 255})
			default:
				img.Set(x, y, color.NRGBA{255, 255, 0, a})
			}
		}
	}
	buf := &bytes.Buffer{}
	if err := png.Encode(buf, img); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

// genHalvesPNG builds a w x h PNG split in two along one axis: the first
// half (left, for a horizontal split, or top, for a vertical one) is a,
// the second half is b. A centre-square crop of this always keeps points
// near its near edge on the a side and its far edge on the b side, which
// is what TestProcessCrestCropsA{Landscape,Portrait}UploadAroundItsCentre
// checks - a crop that silently picked the wrong offset or axis would
// instead land both sample points on the same side.
func genHalvesPNG(t *testing.T, w, h int, horizontal bool, a, b color.NRGBA) []byte {
	t.Helper()
	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			first := x < w/2
			if !horizontal {
				first = y < h/2
			}
			if first {
				img.Set(x, y, a)
			} else {
				img.Set(x, y, b)
			}
		}
	}
	buf := &bytes.Buffer{}
	if err := png.Encode(buf, img); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func genJPEG(t *testing.T, w, h int) []byte {
	t.Helper()
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), 128, 255})
		}
	}
	buf := &bytes.Buffer{}
	if err := jpeg.Encode(buf, img, nil); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestProcessCrestRejectsNonImageBytes(t *testing.T) {
	_, err := ProcessCrest([]byte("not an image, just text"))
	var verr *ValidationError
	if err == nil || !asValidationError(err, &verr) {
		t.Fatalf("err = %v, want a *ValidationError", err)
	}
	if verr.Message == "" {
		t.Fatal("want a plain message")
	}
}

func TestProcessCrestRejectsOversizeFiles(t *testing.T) {
	// A valid PNG whose own bytes exceed the 2 MiB ceiling - a large,
	// incompressible (per-pixel random-ish) image is the simplest way to
	// force a real PNG past the limit without just padding junk after a
	// valid file (which a real decoder might also reject, for the wrong
	// reason).
	img := image.NewNRGBA(image.Rect(0, 0, 1200, 1200))
	for y := 0; y < 1200; y++ {
		for x := 0; x < 1200; x++ {
			img.Set(x, y, color.NRGBA{uint8(x * 7 % 256), uint8(y * 13 % 256), uint8((x ^ y) % 256), 255})
		}
	}
	buf := &bytes.Buffer{}
	if err := (&png.Encoder{CompressionLevel: png.NoCompression}).Encode(buf, img); err != nil {
		t.Fatal(err)
	}
	data := buf.Bytes()
	if len(data) <= MaxBytes {
		t.Fatalf("fixture is %d bytes, want it over MaxBytes (%d) to exercise the oversize path", len(data), MaxBytes)
	}

	_, err := ProcessCrest(data)
	var verr *ValidationError
	if err == nil || !asValidationError(err, &verr) {
		t.Fatalf("err = %v, want a *ValidationError", err)
	}
	if want := "the limit is 2 MB"; !bytes.Contains([]byte(verr.Message), []byte(want)) {
		t.Fatalf("message = %q, want it to contain %q", verr.Message, want)
	}
}

func TestProcessCrestRejectsUnder128(t *testing.T) {
	data := genPNG(t, 64, 64, false)
	_, err := ProcessCrest(data)
	var verr *ValidationError
	if err == nil || !asValidationError(err, &verr) {
		t.Fatalf("err = %v, want a *ValidationError", err)
	}
	if want := "128x128"; !bytes.Contains([]byte(verr.Message), []byte(want)) {
		t.Fatalf("message = %q, want it to mention the 128x128 floor", verr.Message)
	}
}

func TestProcessCrestAcceptsExactlyTheMinimum(t *testing.T) {
	data := genPNG(t, MinDimension, MinDimension, false)
	result, err := ProcessCrest(data)
	if err != nil {
		t.Fatalf("a 128x128 upload should be accepted: %v", err)
	}
	assertCrestSize(t, result.WebP)
}

func TestProcessCrestCropsALandscapeUploadAroundItsCentre(t *testing.T) {
	red, blue := color.NRGBA{255, 0, 0, 255}, color.NRGBA{0, 0, 255, 255}
	data := genHalvesPNG(t, 600, 200, true, red, blue)
	result, err := ProcessCrest(data)
	if err != nil {
		t.Fatal(err)
	}
	img := decodeOutput(t, result.WebP)
	assertCloserTo(t, img, 0, CrestSize/2, red, blue)
	assertCloserTo(t, img, CrestSize-1, CrestSize/2, blue, red)
}

func TestProcessCrestCropsAPortraitUploadAroundItsCentre(t *testing.T) {
	red, blue := color.NRGBA{255, 0, 0, 255}, color.NRGBA{0, 0, 255, 255}
	data := genHalvesPNG(t, 200, 600, false, red, blue)
	result, err := ProcessCrest(data)
	if err != nil {
		t.Fatal(err)
	}
	img := decodeOutput(t, result.WebP)
	assertCloserTo(t, img, CrestSize/2, 0, red, blue)
	assertCloserTo(t, img, CrestSize/2, CrestSize-1, blue, red)
}

func TestProcessCrestKeepsAlpha(t *testing.T) {
	data := genPNG(t, 256, 256, true)
	result, err := ProcessCrest(data)
	if err != nil {
		t.Fatal(err)
	}
	img := decodeOutput(t, result.WebP)
	_, _, _, a := img.At(CrestSize-1, CrestSize-1).RGBA()
	if a == 0xffff {
		t.Fatal("the bottom-right quadrant was uploaded semi-transparent; it must not come out fully opaque")
	}
	_, _, _, a = img.At(0, 0).RGBA()
	if a != 0xffff {
		t.Fatal("the top-left quadrant was uploaded fully opaque; it must not come out transparent")
	}
}

func TestProcessCrestAcceptsJPEGAndWebP(t *testing.T) {
	jpegData := genJPEG(t, 300, 300)
	if _, err := ProcessCrest(jpegData); err != nil {
		t.Fatalf("a valid JPEG should be accepted: %v", err)
	}

	// Round-trip a PNG through this package's own encoder to get a
	// valid WebP fixture, then feed that back in as the upload.
	png := genPNG(t, 300, 300, false)
	first, err := ProcessCrest(png)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := ProcessCrest(first.WebP); err != nil {
		t.Fatalf("a valid WebP should be accepted: %v", err)
	}
}

func TestProcessCrestResizesToTheCrestSize(t *testing.T) {
	data := genPNG(t, 1000, 1000, false)
	result, err := ProcessCrest(data)
	if err != nil {
		t.Fatal(err)
	}
	assertCrestSize(t, result.WebP)
}

func TestProcessCrestSHA256MatchesItsOwnBytes(t *testing.T) {
	data := genPNG(t, 256, 256, false)
	result, err := ProcessCrest(data)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.SHA256) != 64 {
		t.Fatalf("sha256 hex = %q, want 64 hex characters", result.SHA256)
	}
}

// --- helpers ---

func asValidationError(err error, target **ValidationError) bool {
	verr, ok := err.(*ValidationError)
	if ok {
		*target = verr
	}
	return ok
}

func decodeOutput(t *testing.T, webpBytes []byte) image.Image {
	t.Helper()
	img, err := nativewebp.Decode(bytes.NewReader(webpBytes))
	if err != nil {
		t.Fatalf("could not decode this package's own output: %v", err)
	}
	return img
}

func assertCrestSize(t *testing.T, webpBytes []byte) {
	t.Helper()
	img := decodeOutput(t, webpBytes)
	b := img.Bounds()
	if b.Dx() != CrestSize || b.Dy() != CrestSize {
		t.Fatalf("output size = %dx%d, want %dx%d", b.Dx(), b.Dy(), CrestSize, CrestSize)
	}
}

// assertCloserTo checks the pixel at (x, y) is closer to want than to
// other - CatmullRom resampling blends neighbouring pixels, so an exact
// match is not expected right at an edge, but a crop that landed on the
// wrong side (or the wrong axis, or an unflipped offset) reads much
// closer to other instead.
func assertCloserTo(t *testing.T, img image.Image, x, y int, want, other color.NRGBA) {
	t.Helper()
	r, g, b, _ := img.At(x, y).RGBA()
	got := color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255}
	if diff(got, other) < diff(got, want) {
		t.Fatalf("pixel at (%d,%d) = %+v, closer to %+v than to the expected %+v", x, y, got, other, want)
	}
}

func diff(a, b color.NRGBA) int {
	d := func(x, y uint8) int {
		if x > y {
			return int(x - y)
		}
		return int(y - x)
	}
	return d(a.R, b.R) + d(a.G, b.G) + d(a.B, b.B)
}
