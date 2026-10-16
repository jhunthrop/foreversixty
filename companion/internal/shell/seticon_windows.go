//go:build windows && !nogui

// companion/internal/shell/seticon_windows.go
// The webview library registers its window class without an icon, so the
// title bar, Alt-Tab and the taskbar show the generic one even when the
// exe carries the seal. This sends WM_SETICON with the icons from the
// exe's own resource (icon group 1, which go-winres writes from
// seal.ico; see cmd/foreversixty-companion/winres), at the sizes the
// system asks for on this display.
package shell

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	wmSetIcon    = 0x0080
	iconSmall    = 0
	iconBig      = 1
	imageIcon    = 1
	smCxIcon     = 11
	smCyIcon     = 12
	smCxSmIcon   = 49
	smCySmIcon   = 50
	appIconID    = 1 // RT_GROUP_ICON "#1" in winres.json
	lrDefaultClr = 0
)

var (
	sysUser32        = windows.NewLazySystemDLL("user32.dll")
	loadImage        = sysUser32.NewProc("LoadImageW")
	getSystemMetrics = sysUser32.NewProc("GetSystemMetrics")
	sendMessage      = sysUser32.NewProc("SendMessageW")
	getModuleHandle  = windows.NewLazySystemDLL("kernel32.dll").NewProc("GetModuleHandleW")
)

// setWindowIcons gives hwnd the seal for its title bar and taskbar. It
// does nothing when the window or the resource is missing (a `go run`
// build has no resource), and the library's default stays.
func setWindowIcons(hwnd unsafe.Pointer) {
	if hwnd == nil {
		return
	}
	inst, _, _ := getModuleHandle.Call(0)
	if inst == 0 {
		return
	}
	load := func(cx, cy int32) uintptr {
		h, _, _ := loadImage.Call(inst, uintptr(appIconID), imageIcon,
			uintptr(cx), uintptr(cy), lrDefaultClr)
		return h
	}
	metric := func(i uintptr) int32 { v, _, _ := getSystemMetrics.Call(i); return int32(v) }
	if h := load(metric(smCxSmIcon), metric(smCySmIcon)); h != 0 {
		sendMessage.Call(uintptr(hwnd), wmSetIcon, iconSmall, h)
	}
	if h := load(metric(smCxIcon), metric(smCyIcon)); h != 0 {
		sendMessage.Call(uintptr(hwnd), wmSetIcon, iconBig, h)
	}
}
