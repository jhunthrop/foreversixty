//go:build windows && !nogui

// companion/internal/shell/chrome_windows.go
// What makes the window look like part of the app rather than a web page
// in a box: a dark title bar in the header's colour, a page-coloured
// backing so no white shows while the webview starts, and a place in the
// middle of the screen it opens on. Every call ignores its error:
// Windows 10 and older builds of Windows 11 do not know the title bar
// attributes and simply leave the system's bar, which is the correct
// fallback.
package shell

import (
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	dwmwaBorderColor          = 34
	dwmwaCaptionColor         = 35
	dwmwaTextColor            = 36
	dwmwaUseImmersiveDarkMode = 20

	gclpHbrBackground = ^uintptr(10 - 1) // -10 as GCLP_HBRBACKGROUND
	monitorNearest    = 2                // MONITOR_DEFAULTTONEAREST
	swpNoSize         = 0x0001
	swpNoZOrder       = 0x0004
	swpNoActivate     = 0x0010
)

var (
	dwmapi                = windows.NewLazySystemDLL("dwmapi.dll")
	dwmSetWindowAttribute = dwmapi.NewProc("DwmSetWindowAttribute")
	sysGdi32              = windows.NewLazySystemDLL("gdi32.dll")
	createSolidBrush      = sysGdi32.NewProc("CreateSolidBrush")
	setClassLongPtr       = sysUser32.NewProc("SetClassLongPtrW")
	monitorFromWindow     = sysUser32.NewProc("MonitorFromWindow")
	getMonitorInfo        = sysUser32.NewProc("GetMonitorInfoW")
	getWindowRect         = sysUser32.NewProc("GetWindowRect")
	setWindowPos          = sysUser32.NewProc("SetWindowPos")
)

type monitorInfo struct {
	Size    uint32
	Monitor rect
	Work    rect
	Flags   uint32
}

// prepareChrome runs before the webview is made. WebView2 paints its own
// default background (white) until the first page paints; this is the
// documented environment override for that colour. The page's own
// html background and the window class brush below are the other two
// measures, and all three are wanted.
func prepareChrome() {
	if os.Getenv("WEBVIEW2_DEFAULT_BACKGROUND_COLOR") == "" {
		_ = os.Setenv("WEBVIEW2_DEFAULT_BACKGROUND_COLOR", pageARGB)
	}
}

// applyChrome dresses the window hwnd: title bar colours, the class
// background brush, and a position centred on the work area of the
// monitor it opened on. The size is already set when this is called.
func applyChrome(hwnd unsafe.Pointer) {
	if hwnd == nil {
		return
	}
	h := uintptr(hwnd)
	setDwm := func(attr uintptr, v uint32) {
		dwmSetWindowAttribute.Call(h, attr, uintptr(unsafe.Pointer(&v)), unsafe.Sizeof(v))
	}
	setDwm(dwmwaUseImmersiveDarkMode, 1)
	setDwm(dwmwaCaptionColor, captionColor)
	setDwm(dwmwaTextColor, captionText)
	setDwm(dwmwaBorderColor, borderColor)

	if brush, _, _ := createSolidBrush.Call(uintptr(pageColor)); brush != 0 {
		setClassLongPtr.Call(h, gclpHbrBackground, brush)
	}

	mon, _, _ := monitorFromWindow.Call(h, monitorNearest)
	var mi monitorInfo
	mi.Size = uint32(unsafe.Sizeof(mi))
	if mon == 0 {
		return
	}
	if ok, _, _ := getMonitorInfo.Call(mon, uintptr(unsafe.Pointer(&mi))); ok == 0 {
		return
	}
	var wr rect
	if ok, _, _ := getWindowRect.Call(h, uintptr(unsafe.Pointer(&wr))); ok == 0 {
		return
	}
	x, y := centred(mi.Work, wr.Right-wr.Left, wr.Bottom-wr.Top)
	setWindowPos.Call(h, 0, uintptr(x), uintptr(y), 0, 0, swpNoSize|swpNoZOrder|swpNoActivate)
}
