//go:build darwin && !nogui

// companion/internal/shell/closehook_darwin.go
// The Mac window behaves like the Windows one: the close button hides it,
// the menu-bar icon's "Open Forever Sixty" (or a click on the Dock icon)
// shows it again, and only Quit stops the companion. The webview library
// has no close event, so the window delegate it installs is taught
// windowShouldClose: (macapp_darwin.m). The same file gives the app the
// menu bar a Mac app is expected to have: without an Edit menu, Command-V
// does nothing in a web view, and pasting a pairing code is half the
// Device page.
package shell

/*
#cgo CFLAGS: -x objective-c
#include <stdbool.h>
bool fsInstallCloseHook(void *window);
void fsShowWindow(void *window);
void fsHideWindow(void *window);
void fsInstallAppHooks(void);
*/
import "C"

import (
	"sync"
	"unsafe"
)

// macHooks are what the Objective-C side calls back into. There is one
// window at a time, so one set; every call arrives on the main thread,
// and the lock keeps the race detector's view of that honest.
var macHooks struct {
	sync.Mutex
	onClose      func() bool
	quit, reopen func()
}

// installCloseHook teaches the window to hide instead of closing. onClose
// is called on the main thread for each close request: true hides the
// window, false leaves it up (the caller hides it later with hide). ok is
// false when the window has no delegate to teach, and the caller keeps the
// library's behaviour.
func installCloseHook(window unsafe.Pointer, onClose func() bool) (show, hide func(), ok bool) {
	if window == nil {
		return nil, nil, false
	}
	macHooks.Lock()
	macHooks.onClose = onClose
	macHooks.Unlock()
	if !bool(C.fsInstallCloseHook(window)) {
		return nil, nil, false
	}
	show = func() { C.fsShowWindow(window) }
	hide = func() { C.fsHideWindow(window) }
	return show, hide, true
}

// bringToFront makes the app active and the window key: a window opened
// before the app has a menu bar or a tray (the first-launch prompt) is
// otherwise left behind whatever the player clicked last.
func bringToFront(window unsafe.Pointer) {
	if window != nil {
		C.fsShowWindow(window)
	}
}

// installAppHooks gives the app its menu bar and routes the system's Quit
// (Command-Q, the Dock's Quit) through quit, so the companion stops the
// way the tray's Quit stops it, and a click on the Dock icon through
// reopen. It must run on the main thread, before the first window runs.
func installAppHooks(quit, reopen func()) {
	macHooks.Lock()
	macHooks.quit, macHooks.reopen = quit, reopen
	macHooks.Unlock()
	C.fsInstallAppHooks()
}

//export fsGoWindowShouldClose
func fsGoWindowShouldClose() C.bool {
	macHooks.Lock()
	f := macHooks.onClose
	macHooks.Unlock()
	return C.bool(f == nil || f())
}

//export fsGoQuit
func fsGoQuit() {
	macHooks.Lock()
	f := macHooks.quit
	macHooks.Unlock()
	if f != nil {
		f()
	}
}

//export fsGoReopen
func fsGoReopen() {
	macHooks.Lock()
	f := macHooks.reopen
	macHooks.Unlock()
	if f != nil {
		f()
	}
}
