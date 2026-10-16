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
