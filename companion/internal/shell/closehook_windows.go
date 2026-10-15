//go:build windows

// companion/internal/shell/closehook_windows.go
// Closing the window must not end the companion: the tray is what keeps
// it logging. The webview library offers no close event, so on Windows
// the window procedure is wrapped: WM_CLOSE is answered by hiding the
// window instead of destroying it, and the tray's "Open Forever Sixty"
// shows it again. Everything else goes to the library's own procedure.
package shell

import (
	"syscall"
	"unsafe"
)

const (
	wmClose   = 0x0010
	swHide    = 0
	swShow    = 5
	swRestore = 9
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	setWindowLongPtr = user32.NewProc("SetWindowLongPtrW")
	callWindowProc   = user32.NewProc("CallWindowProcW")
	showWindow       = user32.NewProc("ShowWindow")
	isIconic         = user32.NewProc("IsIconic")
	setForeground    = user32.NewProc("SetForegroundWindow")
)

// installCloseHook wraps the window procedure of hwnd. onClose is
// called on the window's thread for each close request: true hides the
// window, false leaves it up (the caller hides it later with hide).
// show brings the window back and to the front. ok is false when the
// window could not be wrapped, and the caller keeps the library's
// behaviour.
func installCloseHook(hwnd unsafe.Pointer, onClose func() bool) (show, hide func(), ok bool) {
	if hwnd == nil {
		return nil, nil, false
	}
	h := uintptr(hwnd)
	var old uintptr
	cb := syscall.NewCallback(func(w, msg, wparam, lparam uintptr) uintptr {
		if msg == wmClose && w == h && onClose() {
			showWindow.Call(h, swHide)
			return 0
		}
		r, _, _ := callWindowProc.Call(old, w, msg, wparam, lparam)
		return r
	})
	gwlpWndProc := -4
	prev, _, _ := setWindowLongPtr.Call(h, uintptr(gwlpWndProc), cb)
	if prev == 0 {
		return nil, nil, false
	}
	old = prev
	hide = func() { showWindow.Call(h, swHide) }
	show = func() {
		flag := uintptr(swShow)
		if min, _, _ := isIconic.Call(h); min != 0 {
			flag = swRestore
		}
		showWindow.Call(h, flag)
		setForeground.Call(h)
	}
	return show, hide, true
}
