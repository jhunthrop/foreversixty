package icon

import (
	"bytes"
	"encoding/binary"
	"image/png"
	"os"
	"regexp"
	"strconv"
	"testing"
)

// tokenRGB reads a hex colour token from the window's generated tokens
// (ui/tokens.css), which the tray dots share with the Status headline.
func tokenRGB(t *testing.T, name string) [3]uint32 {
	t.Helper()
	css, err := os.ReadFile("../../ui/tokens.css")
	if err != nil {
		t.Fatal(err)
	}
	m := regexp.MustCompile(`(?m)^\s*` + regexp.QuoteMeta(name) + `:\s*#([0-9a-f]{6});`).FindSubmatch(css)
	if m == nil {
		t.Fatalf("ui/tokens.css has no hex %s", name)
	}
	v, _ := strconv.ParseUint(string(m[1]), 16, 32)
	return [3]uint32{uint32(v >> 16), uint32(v >> 8 & 0xff), uint32(v & 0xff)}
}

func decode(t *testing.T, b []byte, want int) {
	t.Helper()
	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	if r := img.Bounds(); r.Dx() != want || r.Dy() != want {
		t.Fatalf("bounds = %v, want %d square", r, want)
	}
}

func TestEverySizeDecodesAtItsSize(t *testing.T) {
	for _, s := range Sizes {
		decode(t, PNGSize(s), s)
	}
	if PNGSize(33) != nil {
		t.Error("PNGSize(33) is not nil")
	}
	decode(t, PNG(), 32)
}

func TestTheSealIsNotTransparentAtTheCentre(t *testing.T) {
	img, _ := png.Decode(bytes.NewReader(PNGSize(256)))
	if _, _, _, a := img.At(128, 128).RGBA(); a == 0 {
		t.Error("the middle of the seal is transparent")
	}
	if _, _, _, a := img.At(0, 0).RGBA(); a != 0 {
		t.Error("the corner of the seal is not transparent")
	}
}

func TestICOHasEverySizeAsPNG(t *testing.T) {
	b := ICO()
	if len(b) < 6 || binary.LittleEndian.Uint16(b[0:]) != 0 || binary.LittleEndian.Uint16(b[2:]) != 1 {
		t.Fatal("not an ICO header")
	}
	n := int(binary.LittleEndian.Uint16(b[4:]))
	want := []int{16, 20, 24, 32, 40, 48, 64, 128, 256}
	if n != len(want) {
		t.Fatalf("%d entries, want %d", n, len(want))
	}
	for i, w := range want {
		e := b[6+16*i:]
		got := int(e[0])
		if got == 0 {
			got = 256
		}
		if got != w {
			t.Errorf("entry %d is %d px, want %d", i, got, w)
		}
		size := binary.LittleEndian.Uint32(e[8:])
		off := binary.LittleEndian.Uint32(e[12:])
		decode(t, b[off:off+size], w)
	}
}

func TestICNSTableCoversSixteenToFiveTwelveAtTwoX(t *testing.T) {
	b := ICNS()
	if string(b[:4]) != "icns" || int(binary.BigEndian.Uint32(b[4:])) != len(b) {
		t.Fatal("bad icns header")
	}
	want := map[string]int{
		"icp4": 16, "icp5": 32, "icp6": 64, "ic07": 128, "ic08": 256, "ic09": 512,
		"ic10": 1024, "ic11": 32, "ic12": 64, "ic13": 256, "ic14": 512,
	}
	for p := 8; p < len(b); {
		typ := string(b[p : p+4])
		n := int(binary.BigEndian.Uint32(b[p+4:]))
		w, ok := want[typ]
		if !ok {
			t.Fatalf("unexpected element %q", typ)
		}
		decode(t, b[p+8:p+n], w)
		delete(want, typ)
		p += n
	}
	if len(want) != 0 {
		t.Errorf("missing elements %v", want)
	}
}

func TestTheTrayIconsCarryADotInTheToneAndOffIsThePlainSeal(t *testing.T) {
	// The dot sits in the lower right corner of the 32 px icon, centred on (24, 24),
	// and is the headline's tone, the same token the window's --tone-* reads
	// (ui/styles.css); the upper left is the seal and not the dot.
	want := map[string][3]uint32{
		"on":   tokenRGB(t, "--color-kill"),
		"warn": tokenRGB(t, "--color-gold-hover"),
		"bad":  tokenRGB(t, "--color-wipe"),
	}
	ring := tokenRGB(t, "--color-raised")
	plain, _ := png.Decode(bytes.NewReader(PNG()))
	for _, tone := range Tones {
		img, err := png.Decode(bytes.NewReader(TrayPNG(tone)))
		if err != nil {
			t.Fatalf("%s: %v", tone, err)
		}
		if r := img.Bounds(); r.Dx() != 32 || r.Dy() != 32 {
			t.Fatalf("%s: bounds %v, want 32 square", tone, r)
		}
		r, g, b, a := img.At(25, 25).RGBA()
		got := [3]uint32{r >> 8, g >> 8, b >> 8}
		if got != want[tone] || a>>8 != 0xff {
			t.Errorf("%s: dot pixel = %v alpha %d, want %v opaque", tone, got, a>>8, want[tone])
		}
		pr, pg, pb, pa := plain.At(25, 25).RGBA()
		if pr>>8 == got[0] && pg>>8 == got[1] && pb>>8 == got[2] && pa>>8 == 0xff {
			t.Errorf("%s: the plain seal already has this colour at the dot", tone)
		}
		// The ring is the panel colour, between the dot and the seal.
		rr, rg, rb, _ := img.At(24-8, 24).RGBA()
		if got := [3]uint32{rr >> 8, rg >> 8, rb >> 8}; got != ring {
			t.Errorf("%s: ring pixel = %v, want %v", tone, got, ring)
		}
	}
	for _, tone := range []string{"off", "", "unknown"} {
		if !bytes.Equal(TrayPNG(tone), PNG()) || !bytes.Equal(TrayICO(tone), ICO()) {
			t.Errorf("tone %q is not the plain seal", tone)
		}
	}
}

// Windows draws the tray at 16 px a scale step: 100, 125, 150, 200, 250 and 300 %.
func TestEachTrayICOHasEveryWindowsTrayScale(t *testing.T) {
	for _, tone := range Tones {
		b := TrayICO(tone)
		if len(b) < 6 || binary.LittleEndian.Uint16(b[0:]) != 0 || binary.LittleEndian.Uint16(b[2:]) != 1 {
			t.Fatalf("%s: not an ICO header", tone)
		}
		want := []int{16, 20, 24, 32, 40, 48}
		if n := int(binary.LittleEndian.Uint16(b[4:])); n != len(want) {
			t.Fatalf("%s: %d entries, want %d", tone, n, len(want))
		}
		for i, w := range want {
			e := b[6+16*i:]
			if int(e[0]) != w {
				t.Errorf("%s: entry %d is %d px, want %d", tone, i, e[0], w)
			}
			size := binary.LittleEndian.Uint32(e[8:])
			off := binary.LittleEndian.Uint32(e[12:])
			decode(t, b[off:off+size], w)
		}
	}
}
