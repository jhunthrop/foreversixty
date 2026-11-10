//go:build !windows

// companion/internal/updater/icons_other.go
package updater

// shellNotifyIconChange has nothing to tell outside Windows: the Mac's
// refresh is the bundle's new modification time (touchBundle).
func shellNotifyIconChange() error { return nil }
