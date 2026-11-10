//go:build !windows && (!darwin || nogui)

package shell

import "unsafe"

// installCloseHook is Windows and macOS only (closehook_darwin.go): the
// webview library has no close event, and Linux has no equivalent that
// does not need a GTK binding of our own. There the window's close button
// ends the window, as the library does it.
// bringToFront is macOS only: elsewhere a new window opens in front.
func bringToFront(unsafe.Pointer) {}

func installCloseHook(unsafe.Pointer, func() bool) (show, hide func(), ok bool) {
	return nil, nil, false
}
