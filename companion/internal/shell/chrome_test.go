package shell

import (
	"fmt"
	"image/color"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// tokenRef is a hex token of ui/tokens.css as a COLORREF (0x00BBGGRR).
func tokenRef(t *testing.T, css []byte, name string) uint32 {
	t.Helper()
	m := regexp.MustCompile(`(?m)^\s*` + regexp.QuoteMeta(name) + `:\s*#([0-9a-f]{6});`).FindSubmatch(css)
	if m == nil {
		t.Fatalf("ui/tokens.css has no hex %s", name)
	}
	v, _ := strconv.ParseUint(string(m[1]), 16, 32)
	return colorRef(color.NRGBA{R: uint8(v >> 16), G: uint8(v >> 8), B: uint8(v), A: 0xFF})
}

// The title bar and the page behind the webview are the window's tokens,
// so a change in design/tokens is a change here.
func TestTheTitleBarColoursAreTheWindowTokens(t *testing.T) {
	css, err := os.ReadFile("../../ui/tokens.css")
	if err != nil {
		t.Fatal(err)
	}
	for name, got := range map[string]uint32{
		"--color-nav-bar": captionColor,
		"--color-cream":   captionText,
		"--color-foot":    borderColor,
		"--color-ground":  pageColor,
	} {
		if want := tokenRef(t, css, name); got != want {
			t.Errorf("%s = %#08x, want %#08x", name, got, want)
		}
	}
	r, g, b := pageColor&0xff, pageColor>>8&0xff, pageColor>>16&0xff
	if want := fmt.Sprintf("FF%02X%02X%02X", r, g, b); !strings.EqualFold(pageARGB, want) {
		t.Errorf("pageARGB = %s, want %s", pageARGB, want)
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
