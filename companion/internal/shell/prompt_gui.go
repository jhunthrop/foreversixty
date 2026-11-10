//go:build !nogui

// companion/internal/shell/prompt_gui.go
// A small fixed window over one page, with no tray: the first-launch and
// setup windows (internal/install), the Mac's "Move Forever Sixty to
// Applications?" and the Windows setup and uninstall. It runs on the main
// goroutine, before the companion's own window, and returns when the page's
// binding closes it or the player does.
package shell

import (
	"runtime"

	webview "github.com/webview/webview_go"
)

// PromptWindow is what a binding may do to its window, from any goroutine.
type PromptWindow interface {
	// Close ends the window; Prompt then returns.
	Close()
	// Eval runs JavaScript in the page: progress the Go side reports.
	Eval(js string)
}

// PromptOptions configures a Prompt window.
type PromptOptions struct {
	URL, Title    string
	Width, Height int
	// Bind gives the functions the page may call, by name.
	Bind func(win PromptWindow) map[string]any
}

type promptWindow struct{ w webview.WebView }

func (p promptWindow) Close()         { p.w.Dispatch(p.w.Terminate) }
func (p promptWindow) Eval(js string) { p.w.Dispatch(func() { p.w.Eval(js) }) }

// Prompt shows the window and blocks until it is closed.
func Prompt(o PromptOptions) error {
	runtime.LockOSThread()
	prepareChrome()
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle(o.Title)
	setWindowIcons(w.Window())
	w.SetSize(o.Width, o.Height, webview.HintFixed)
	applyChrome(w.Window())
	dialogChrome(w.Window())
	if o.Bind != nil {
		for name, fn := range o.Bind(promptWindow{w}) {
			if err := w.Bind(name, fn); err != nil {
				return err
			}
		}
	}
	w.Navigate(o.URL)
	// Once the loop runs: the app is only then able to become active.
	w.Dispatch(func() { bringToFront(w.Window()) })
	w.Run()
	return nil
}
