//go:build !nogui

// companion/internal/shell/prompt_gui.go
// A small fixed window over one page, with no tray: the Mac's first-launch
// "Move Forever Sixty to Applications?" (internal/install). It runs on the
// main goroutine, before the companion's own window, and returns when the
// page's binding closes it or the player does.
package shell

import (
	"runtime"

	webview "github.com/webview/webview_go"
)

// PromptOptions configures a Prompt window.
type PromptOptions struct {
	URL, Title    string
	Width, Height int
	// Bind gives the functions the page may call, by name. It gets the
	// window's close, which any goroutine may call.
	Bind func(closeWindow func()) map[string]any
}

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
	if o.Bind != nil {
		closeWindow := func() { w.Dispatch(w.Terminate) }
		for name, fn := range o.Bind(closeWindow) {
			if err := w.Bind(name, fn); err != nil {
				return err
			}
		}
	}
	w.Navigate(o.URL)
	w.Run()
	return nil
}
