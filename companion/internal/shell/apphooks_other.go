//go:build !darwin && !nogui

package shell

// installAppHooks is macOS-only (closehook_darwin.go): Windows and Linux
// have no application menu, no Dock to click and no system Quit to catch.
func installAppHooks(quit, reopen func()) {}
