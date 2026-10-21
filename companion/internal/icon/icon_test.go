package icon

import (
	"bytes"
	"encoding/binary"
	"image/png"
	"testing"
)

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
	want := []int{16, 24, 32, 48, 64, 128, 256}
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
	// The dot sits at the lower right of the 32 px icon, centred on (25, 25),
	// and is the headline's tone; the upper left is the seal and not the dot.
	want := map[string][3]uint32{"on": {0x7b, 0xff, 0x5c}, "warn": {0xe8, 0x88, 0x3a}, "bad": {0xff, 0x6b, 0x5c}}
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
		// The ring is the page colour, between the dot and the seal.
		rr, rg, rb, _ := img.At(25-7, 25).RGBA()
		if rr>>8 != 0x07 || rg>>8 != 0x09 || rb>>8 != 0x0d {
			t.Errorf("%s: ring pixel = %02x%02x%02x, want 07090d", tone, rr>>8, rg>>8, rb>>8)
		}
	}
	for _, tone := range []string{"off", "", "unknown"} {
		if !bytes.Equal(TrayPNG(tone), PNG()) || !bytes.Equal(TrayICO(tone), ICO()) {
			t.Errorf("tone %q is not the plain seal", tone)
		}
	}
}

func TestEachTrayICOHasSixteenTwentyFourThirtyTwoAndFortyEight(t *testing.T) {
	for _, tone := range Tones {
		b := TrayICO(tone)
		if len(b) < 6 || binary.LittleEndian.Uint16(b[0:]) != 0 || binary.LittleEndian.Uint16(b[2:]) != 1 {
			t.Fatalf("%s: not an ICO header", tone)
		}
		want := []int{16, 24, 32, 48}
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
