//go:build windows

// companion/internal/icon/tray_windows.go
package icon

import "embed"

// The Windows tray takes an ICO; the binary carries the four tray ICOs and
// nothing else.
//
//go:embed assets/seal-tray-*.ico
var tray embed.FS

const trayExt = "ico"
