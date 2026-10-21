// companion/internal/shell/chrome.go
// The window chrome's numbers, kept apart from the Win32 calls that use
// them so a test on any platform can pin them to the design tokens in
// ui/styles.css. The title bar is the header's colour, so the header
// reads as one band with it; the page behind the webview is the page's
// colour, so nothing flashes white while the webview starts.
package shell

// rect is a screen rectangle in pixels, as Win32 has it.
type rect struct{ Left, Top, Right, Bottom int32 }

// colorRef is a Win32 COLORREF, 0x00BBGGRR.
func colorRef(r, g, b uint8) uint32 {
	return uint32(b)<<16 | uint32(g)<<8 | uint32(r)
}

// The tokens, as ui/styles.css has them (--header, --nav, --line, --bg).
var (
	captionColor = colorRef(0x0b, 0x10, 0x18) // #0b1018
	captionText  = colorRef(0xb9, 0xb3, 0xa4) // #b9b3a4
	borderColor  = colorRef(0x26, 0x2e, 0x40) // #262e40
	pageColor    = colorRef(0x07, 0x09, 0x0d) // #07090d
)

// pageARGB is the page colour as the webview's own default background
// wants it, 0xAARRGGBB as hex digits (WEBVIEW2_DEFAULT_BACKGROUND_COLOR).
const pageARGB = "FF07090D"

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
