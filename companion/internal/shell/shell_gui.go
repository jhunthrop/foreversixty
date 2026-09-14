//go:build !nogui

// companion/internal/shell/shell_gui.go
// The real window. Both libraries here need cgo and platform
// toolkits: WebKit on macOS, WebView2 on Windows, WebKitGTK on Linux
// (libgtk-3-dev and libwebkit2gtk-4.1-dev). The release workflow
// installs them per runner; `go test -tags nogui` needs none of it.
package shell

import (
	"context"
	"runtime"

	"github.com/getlantern/systray"
	webview "github.com/webview/webview_go"

	"github.com/jhunthrop/foreversixty/companion/internal/icon"
)

// Run shows the tray and the window and blocks until the player quits
// or the context ends. It must be called from the main goroutine:
// every desktop toolkit here requires it.
func Run(ctx context.Context, o Options) error {
	o = o.withDefaults()
	if o.Headless {
		return wait(ctx)
	}
	runtime.LockOSThread()

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle(o.Title)
	w.SetSize(o.Width, o.Height, webview.HintNone)
	w.Navigate(o.URL)

	// Register rather than Run: systray drives its own loop on some
	// platforms and hands the main loop back on others, and the
	// webview owns the main loop here.
	systray.Register(func() {
		if runtime.GOOS == "windows" {
			systray.SetIcon(icon.ICO())
		} else {
			systray.SetTemplateIcon(icon.PNG(), icon.PNG())
		}
		systray.SetTooltip(o.Title)
		open := systray.AddMenuItem("Open Forever Sixty", "Show the companion window")
		systray.AddSeparator()
		quit := systray.AddMenuItem("Quit", "Stop logging and quit")
		go func() {
			for {
				select {
				case <-ctx.Done():
					w.Terminate()
					return
				case <-open.ClickedCh:
					w.Dispatch(func() { w.Navigate(o.URL) })
				case <-quit.ClickedCh:
					if o.OnQuit != nil {
						o.OnQuit()
					}
					w.Terminate()
					return
				}
			}
		}()
	}, func() {})

	go func() {
		<-ctx.Done()
		w.Terminate()
	}()

	w.Run()
	systray.Quit()
	return nil
}
