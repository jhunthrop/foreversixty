//go:build nogui

// companion/internal/shell/shell_nogui.go
// The headless build. CI compiles and tests with -tags nogui so no
// runner needs WebKit, GTK or WebView2 installed; the companion still
// logs, uploads and serves its local UI over loopback.
package shell

import (
	"context"
	"errors"
)

// Run blocks until the context ends. The local UI server is already
// listening, so a player can open the printed URL in a browser.
func Run(ctx context.Context, o Options) error { return wait(ctx) }

// ErrNoWindow says this build has no window to show.
var ErrNoWindow = errors.New("shell: this build has no window")

// PromptWindow is what a binding may do to its window (prompt_gui.go).
type PromptWindow interface {
	Close()
	Eval(js string)
}

// PromptOptions configures a Prompt window (prompt_gui.go).
type PromptOptions struct {
	URL, Title    string
	Width, Height int
	Bind          func(win PromptWindow) map[string]any
}

// Prompt has no window to show in the headless build.
func Prompt(PromptOptions) error { return ErrNoWindow }
