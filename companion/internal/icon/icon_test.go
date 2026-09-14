package icon

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/png"
	"testing"
)

func TestThePNGDecodesAtTheRightSize(t *testing.T) {
	img, err := png.Decode(bytes.NewReader(PNG()))
	if err != nil {
		t.Fatal(err)
	}
	if b := img.Bounds(); b.Dx() != Size || b.Dy() != Size {
		t.Fatalf("bounds = %v", b)
	}
	// The middle is transparent and the ring is gold.
	if _, _, _, a := img.At(Size/2, Size/2).RGBA(); a != 0 {
		t.Error("the middle of the ring is not transparent")
	}
	// RGBA() is alpha-premultiplied, so the colour is read straight
	// off the NRGBA image the encoder was given.
	rgba, ok := img.(*image.NRGBA)
	if !ok {
		t.Fatalf("decoded image is %T", img)
	}
	top := rgba.NRGBAAt(Size/2, 2)
	if top.A == 0 {
		t.Fatal("the top of the ring is transparent")
	}
	if top.R != Gold.R || top.G != Gold.G || top.B != Gold.B {
		t.Errorf("ring colour = %+v, want %+v", top, Gold)
	}
}

func TestTheICOWrapsThePNG(t *testing.T) {
	ico := ICO()
	if len(ico) != 22+len(PNG()) {
		t.Fatalf("ICO is %d bytes, want 22 + %d", len(ico), len(PNG()))
	}
	var head [3]uint16
	if err := binary.Read(bytes.NewReader(ico), binary.LittleEndian, &head); err != nil {
		t.Fatal(err)
	}
	if head != [3]uint16{0, 1, 1} {
		t.Fatalf("ICONDIR = %v, want reserved 0, type 1, count 1", head)
	}
	if ico[6] != Size || ico[7] != Size {
		t.Errorf("entry size = %d×%d", ico[6], ico[7])
	}
	if !bytes.Equal(ico[22:], PNG()) {
		t.Error("the embedded image is not the PNG")
	}
}
