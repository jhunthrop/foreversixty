//go:build !windows || nogui

package shell

import "unsafe"

// setWindowIcons is Windows-only: macOS takes the dock icon from the app
// bundle and Linux window managers from the desktop entry.
func setWindowIcons(unsafe.Pointer) {}
