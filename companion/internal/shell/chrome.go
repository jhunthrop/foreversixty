// companion/internal/shell/chrome.go
// The window chrome's numbers, kept apart from the Win32 calls that use
// them so a test on any platform can pin them to the design tokens in
// ui/tokens.css (generated from design/tokens/tokens.json). The title bar
// is the site's nav-bar teal, the colour the window's painted band opens
// with, so the caption reads as one band with it; the page behind the
// webview is the brown ground, so nothing flashes white while it starts.
package shell

// rect is a screen rectangle in pixels, as Win32 has it.
type rect struct{ Left, Top, Right, Bottom int32 }

// colorRef is a Win32 COLORREF, 0x00BBGGRR.
func colorRef(r, g, b uint8) uint32 {
	return uint32(b)<<16 | uint32(g)<<8 | uint32(r)
}

// The tokens, as ui/tokens.css has them (--color-nav-bar, --color-cream,
// --color-foot, --color-ground).
var (
	captionColor = colorRef(0x0b, 0x2f, 0x3b) // #0b2f3b
	captionText  = colorRef(0xf3, 0xee, 0xe2) // #f3eee2
	borderColor  = colorRef(0x17, 0x0e, 0x09) // #170e09
	pageColor    = colorRef(0x67, 0x56, 0x3a) // #67563a
)

// pageARGB is the page colour as the webview's own default background
// wants it, 0xAARRGGBB as hex digits (WEBVIEW2_DEFAULT_BACKGROUND_COLOR).
const pageARGB = "FF67563A"

// MinWidth and MinHeight are the smallest the window may be dragged to;
// the page is laid out for 560 wide and up.
const (
	MinWidth  = 560
	MinHeight = 480
)

// centred is the top-left corner that puts a w by h window in the middle
// of the work area, the screen less the taskbar. A work area smaller than
// the window pins the window to its top-left corner rather than off it.
func centred(work rect, w, h int32) (x, y int32) {
	x = work.Left + (work.Right-work.Left-w)/2
	y = work.Top + (work.Bottom-work.Top-h)/2
	if x < work.Left {
		x = work.Left
	}
	if y < work.Top {
		y = work.Top
	}
	return x, y
}
