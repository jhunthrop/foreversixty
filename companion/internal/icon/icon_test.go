package icon

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"testing"
)

// The files are read from assets/ here, so every platform's test checks
// every platform's icons; Tray is checked against them for this one.

func asset(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("assets", name))
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func decode(t *testing.T, b []byte, want int) image.Image {
	t.Helper()
	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	if r := img.Bounds(); r.Dx() != want || r.Dy() != want {
		t.Fatalf("bounds = %v, want %d square", r, want)
	}
	return img
}

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

// icoSizes checks an ICO's header and decodes each PNG entry at its size.
func icoSizes(t *testing.T, b []byte, want []int) {
	t.Helper()
	if len(b) < 6 || binary.LittleEndian.Uint16(b[0:]) != 0 || binary.LittleEndian.Uint16(b[2:]) != 1 {
		t.Fatal("not an ICO header")
	}
	if n := int(binary.LittleEndian.Uint16(b[4:])); n != len(want) {
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

// The exe's icon resource: every size Windows draws, 16 to 64 at each
// scale step, and 256 for Explorer.
func TestTheWindowsAppIconHasEverySizeWindowsDraws(t *testing.T) {
	icoSizes(t, asset(t, "seal.ico"), []int{16, 20, 24, 32, 40, 48, 64, 256})
}

// The bundle icon: one PNG per pixel size macOS draws, 16 to 512, and
// nothing repeated.
func TestTheMacAppIconHasOneEntryPerPixelSize(t *testing.T) {
	b := asset(t, "seal.icns")
	if string(b[:4]) != "icns" || int(binary.BigEndian.Uint32(b[4:])) != len(b) {
		t.Fatal("bad icns header")
	}
	want := map[string]int{"icp4": 16, "icp5": 32, "icp6": 64, "ic07": 128, "ic08": 256, "ic09": 512}
	for p := 8; p < len(b); {
		typ := string(b[p : p+4])
		n := int(binary.BigEndian.Uint32(b[p+4:]))
		w, ok := want[typ]
		if !ok {
			t.Fatalf("unexpected element %q", typ)
		}
		img := decode(t, b[p+8:p+n], w)
		// Apple's grid: the corner is clear, the centre is the seal.
		if _, _, _, a := img.At(w/2, w/2).RGBA(); a == 0 {
			t.Errorf("%s: the middle is transparent", typ)
		}
		if _, _, _, a := img.At(0, 0).RGBA(); a != 0 {
			t.Errorf("%s: the corner is not transparent", typ)
		}
		delete(want, typ)
		p += n
	}
	if len(want) != 0 {
		t.Errorf("missing elements %v", want)
	}
}

// The icons are the small files the owner asked for (2026-11-09: the
// first cut added about 2.5 MB to every binary).
func TestTheIconFilesStaySmall(t *testing.T) {
	for name, max := range map[string]int{"seal.icns": 160 << 10, "seal.ico": 64 << 10} {
		if n := len(asset(t, name)); n > max {
			t.Errorf("%s is %d bytes, over %d", name, n, max)
		}
	}
	tray, _ := filepath.Glob("assets/seal-tray-*")
	total := 0
	for _, f := range tray {
		fi, _ := os.Stat(f)
		total += int(fi.Size())
	}
	if len(tray) != 8 || total > 64<<10 {
		t.Errorf("tray files = %d, %d bytes", len(tray), total)
	}
}

// Windows draws the tray at 16 px a scale step: 100, 125, 150, 200, 250 and 300 %.
func TestEachTrayICOHasEveryWindowsTrayScale(t *testing.T) {
	for _, tone := range append([]string{Off}, Tones...) {
		t.Run(tone, func(t *testing.T) {
			icoSizes(t, asset(t, "seal-tray-"+tone+".ico"), []int{16, 20, 24, 32, 40, 48})
		})
	}
}

func TestTheTrayIconsCarryADotInTheToneAndOffIsThePlainSeal(t *testing.T) {
	// The dot sits in the lower right corner of the 32 px icon, centred on (26, 26),
	// and is the headline's tone, the same token the window's --tone-* reads
	// (ui/styles.css); off has no dot there.
	want := map[string][3]uint32{
		"on":   tokenRGB(t, "--color-kill"),
		"warn": tokenRGB(t, "--color-gold-hover"),
		"bad":  tokenRGB(t, "--color-wipe"),
	}
	ring := tokenRGB(t, "--color-raised")
	near := func(a, b [3]uint32) bool {
		for i := range a {
			if d := int(a[i]) - int(b[i]); d < -6 || d > 6 { // the palette may move a channel a step
				return false
			}
		}
		return true
	}
	rgb := func(img image.Image, x, y int) [3]uint32 {
		r, g, b, _ := img.At(x, y).RGBA()
		return [3]uint32{r >> 8, g >> 8, b >> 8}
	}
	plain := decode(t, asset(t, "seal-tray-off.png"), 32)
	for _, tone := range Tones {
		img := decode(t, asset(t, "seal-tray-"+tone+".png"), 32)
		if got := rgb(img, 25, 25); !near(got, want[tone]) {
			t.Errorf("%s: dot pixel = %v, want %v", tone, got, want[tone])
		}
		if near(rgb(plain, 25, 25), want[tone]) {
			t.Errorf("%s: the plain seal already has this colour at the dot", tone)
		}
		if got := rgb(img, 20, 25); !near(got, ring) {
			t.Errorf("%s: ring pixel = %v, want %v", tone, got, ring)
		}
	}
}

// Tray hands this platform its own image for every tone, and the plain
// seal for off and anything unknown.
func TestTrayIsThisPlatformsImage(t *testing.T) {
	ext := "png"
	if runtime.GOOS == "windows" {
		ext = "ico"
	}
	for _, tone := range append([]string{Off}, Tones...) {
		if !bytes.Equal(Tray(tone), asset(t, "seal-tray-"+tone+"."+ext)) {
			t.Errorf("Tray(%q) is not seal-tray-%s.%s", tone, tone, ext)
		}
	}
	for _, tone := range []string{"", "unknown"} {
		if !bytes.Equal(Tray(tone), Tray(Off)) {
			t.Errorf("tone %q is not the plain seal", tone)
		}
	}
}
