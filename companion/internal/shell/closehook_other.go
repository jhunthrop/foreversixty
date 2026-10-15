//go:build !windows

package shell

import "unsafe"

// installCloseHook is Windows-only: the webview library has no close
// event, and macOS and Linux have no equivalent that does not need a
// native toolkit binding of our own. There the window's close button
// ends the window, as the library does it.
func installCloseHook(unsafe.Pointer, func() bool) (show, hide func(), ok bool) {
	return nil, nil, false
}
