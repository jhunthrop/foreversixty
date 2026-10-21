//go:build !windows || nogui

package shell

import "unsafe"

// prepareChrome and applyChrome are Windows-only: macOS and Linux windows
// take their title bar from the system, and the page's own background
// covers the webview's first moments there.
func prepareChrome() {}

func applyChrome(unsafe.Pointer) {}
