package shell

import "testing"

// The title bar is the header's colour, so a change to a token in
// ui/styles.css is a change here; these are the values the design spec
// gives as COLORREFs (0x00BBGGRR).
func TestTheTitleBarColoursAreTheHeaderTokens(t *testing.T) {
	for name, c := range map[string][2]uint32{
		"caption #0b1018": {captionColor, 0x0018100B},
		"text #b9b3a4":    {captionText, 0x00A4B3B9},
		"border #262e40":  {borderColor, 0x00402E26},
		"page #07090d":    {pageColor, 0x000D0907},
	} {
		if c[0] != c[1] {
			t.Errorf("%s = %#08x, want %#08x", name, c[0], c[1])
		}
	}
	if pageARGB != "FF07090D" {
		t.Errorf("pageARGB = %s", pageARGB)
	}
}

func TestTheWindowIsCentredOnTheWorkAreaAndNeverOffIt(t *testing.T) {
	// A 1920x1080 screen with a 48 px taskbar at the bottom.
	work := rect{0, 0, 1920, 1032}
	if x, y := centred(work, 800, 620); x != 560 || y != 206 {
		t.Errorf("centred = %d,%d, want 560,206", x, y)
	}
	// A second monitor to the left, with a taskbar at the top.
	left := rect{-1280, 40, 0, 1024}
	if x, y := centred(left, 800, 620); x != -1040 || y != 222 {
		t.Errorf("centred on the left monitor = %d,%d, want -1040,222", x, y)
	}
	// A work area smaller than the window pins it to the corner.
	if x, y := centred(rect{0, 0, 600, 400}, 800, 620); x != 0 || y != 0 {
		t.Errorf("centred in a small area = %d,%d, want 0,0", x, y)
	}
}
