//go:build !windows

// companion/internal/icon/tray_other.go
package icon

import "embed"

// The macOS menu bar and the Linux tray take a PNG; the binary carries the
// four tray PNGs and nothing else.
//
//go:embed assets/seal-tray-*.png
var tray embed.FS

const trayExt = "png"
