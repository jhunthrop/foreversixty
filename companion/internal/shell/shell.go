// companion/internal/shell/shell.go
// Package shell is the window: a tray icon and a native webview over
// the local UI server. It exists in two builds — the real one, which
// needs cgo and each platform's webview libraries, and a headless one
// selected with -tags nogui for CI and for servers. Both export the
// same Run, and both honour the -headless flag at runtime, so a
// player on a machine with no tray can still run the companion.
package shell

import "context"

// Options configures the window.
type Options struct {
	// URL is the local UI server's address, token and all.
	URL string
	// Title is the window title.
	Title string
	// Headless skips the tray and the window entirely.
	Headless bool
	// Width and Height are the window's starting size.
	Width, Height int
	// OnQuit is called when the player quits from the tray.
	OnQuit func()
}

// Defaults fills in the window size.
func (o Options) withDefaults() Options {
	if o.Title == "" {
		o.Title = "Forever Sixty"
	}
	if o.Width == 0 {
		o.Width = 960
	}
	if o.Height == 0 {
		o.Height = 680
	}
	return o
}

// wait blocks until the context ends. It is what headless mode does,
// and what the real shell falls back to when the tray is off.
func wait(ctx context.Context) error {
	<-ctx.Done()
	return ctx.Err()
}
