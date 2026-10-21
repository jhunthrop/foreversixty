// companion/internal/shell/shell.go
// Package shell is the window: a tray icon and a native webview over
// the local UI server. It exists in two builds — the real one, which
// needs cgo and each platform's webview libraries, and a headless one
// selected with -tags nogui for CI and for servers. Both export the
// same Run, and both honour the -headless flag at runtime, so a
// player on a machine with no tray can still run the companion.
package shell

import (
	"context"
	"time"
)

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

	// State is the tray's state line: the Status headline's title,
	// word for word. It becomes the tooltip and the menu's first
	// (disabled) item, and is asked for again every TrayEvery.
	State func() string
	// Tone is the headline's tone (on, warn, bad or off). The tray icon is
	// the seal with a dot in that colour; off is the plain seal. It is
	// asked for again every TrayEvery, and the icon is replaced only when
	// the tone changes.
	Tone func() string
	// OpenReports opens the player's reports on the site; the tray's
	// "Open my reports on the site".
	OpenReports func()
	// FirstClose reports whether this is the first time the player has
	// closed the window, and remembers that it was. The window then
	// tells them the companion is still running in the tray.
	FirstClose func() bool
	// Notice is that sentence.
	Notice string
}

// TrayEvery is how often the tray looks at the state line. The window
// polls every two seconds; the tray keeps the same pace.
const TrayEvery = 2 * time.Second

// NoticeFor is how long the first close leaves the window up, with its
// "still running in the tray" bar showing, before hiding it.
const NoticeFor = 4 * time.Second

// withDefaults fills in the window size: 800 by 620, which shows a
// Status headline and its button without scrolling.
func (o Options) withDefaults() Options {
	if o.Title == "" {
		o.Title = "Forever Sixty"
	}
	if o.Width == 0 {
		o.Width = 800
	}
	if o.Height == 0 {
		o.Height = 620
	}
	if o.State == nil {
		o.State = func() string { return o.Title }
	}
	if o.Tone == nil {
		o.Tone = func() string { return "off" }
	}
	return o
}

// wait blocks until the context ends. It is what headless mode does,
// and what the real shell falls back to when the tray is off.
func wait(ctx context.Context) error {
	<-ctx.Done()
	return ctx.Err()
}

// watchLine calls apply with the value now and again whenever it
// changes, until the context ends. The tray uses it so a menu item, a
// tooltip and an icon are rewritten only when what they show changes.
func watchLine[T comparable](ctx context.Context, every time.Duration, read func() T, apply func(T)) {
	last := read()
	apply(last)
	t := time.NewTicker(every)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			if line := read(); line != last {
				last = line
				apply(line)
			}
		}
	}
}
