//go:build nogui

// companion/internal/shell/shell_nogui.go
// The headless build. CI compiles and tests with -tags nogui so no
// runner needs WebKit, GTK or WebView2 installed; the companion still
// logs, uploads and serves its local UI over loopback.
package shell

import "context"

// Run blocks until the context ends. The local UI server is already
// listening, so a player can open the printed URL in a browser.
func Run(ctx context.Context, o Options) error { return wait(ctx) }
