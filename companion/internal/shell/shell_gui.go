//go:build !nogui

// companion/internal/shell/shell_gui.go
// The real window. Both libraries here need cgo and platform
// toolkits: WebKit on macOS, WebView2 on Windows, WebKitGTK on Linux
// (libgtk-3-dev and libwebkit2gtk-4.1-dev). The release workflow
// installs them per runner; `go test -tags nogui` needs none of it.
package shell

import (
	"context"
	"encoding/json"
	"runtime"
	"sync"
	"time"

	"github.com/getlantern/systray"
	webview "github.com/webview/webview_go"

	"github.com/jhunthrop/foreversixty/companion/internal/icon"
)

// Run shows the tray and the window and blocks until the player quits
// or the context ends. It must be called from the main goroutine:
// every desktop toolkit here requires it.
//
// The tray outlives the window. On Windows the close button hides the
// window (installCloseHook) and "Open Forever Sixty" shows it again;
// the first close leaves it up for NoticeFor with a bar saying the
// companion is still running in the tray. Where the close cannot be
// caught (macOS, Linux: the toolkits have no close event and the tray
// there shares the window's main loop) closing the window ends the
// process, as before.
func Run(ctx context.Context, o Options) error {
	o = o.withDefaults()
	if o.Headless {
		return wait(ctx)
	}
	runtime.LockOSThread()

	var (
		mu       sync.Mutex
		cur      webview.WebView
		show     func()
		quitOnce sync.Once
	)
	quit := make(chan struct{})
	openReq := make(chan struct{}, 1)
	stopAll := func() {
		quitOnce.Do(func() { close(quit) })
		mu.Lock()
		w := cur
		mu.Unlock()
		if w != nil {
			w.Terminate()
		}
	}
	// openWindow brings the window back: a hidden one is shown, a
	// destroyed one is asked for by the main loop below.
	openWindow := func() {
		mu.Lock()
		w, s := cur, show
		mu.Unlock()
		if w != nil && s != nil {
			w.Dispatch(s)
			return
		}
		select {
		case openReq <- struct{}{}:
		default:
		}
	}

	// Register rather than Run: systray drives its own loop on some
	// platforms and hands the main loop back on others, and the
	// webview owns the main loop here.
	systray.Register(func() {
		view := func() trayView { return trayView{Line: o.State(), Tone: o.Tone()} }
		first := view()
		// The icon and the tooltip are set before the menu is built, so the tray
		// is never up without them.
		setTrayIcon(first.Tone)
		systray.SetTooltip(tooltip(first.Line))
		state := systray.AddMenuItem(first.Line, first.Line)
		state.Disable()
		open := systray.AddMenuItem("Open Forever Sixty", "Show the companion window")
		reports := systray.AddMenuItem("Open my reports on the site", "Open your reports in the browser")
		systray.AddSeparator()
		quitItem := systray.AddMenuItem("Quit", "Stop logging and quit")
		// Each is rewritten only when what it shows changes.
		shown := &first
		go watchLine(ctx, TrayEvery, view, func(now trayView) {
			applyTray(shown, now, trayApply{
				Tooltip: systray.SetTooltip,
				Title:   state.SetTitle,
				Icon:    setTrayIcon,
			})
			shown = &now
		})
		go func() {
			for {
				select {
				case <-ctx.Done():
					stopAll()
					return
				case <-quit:
					return
				case <-open.ClickedCh:
					openWindow()
				case <-reports.ClickedCh:
					if o.OpenReports != nil {
						o.OpenReports()
					}
				case <-quitItem.ClickedCh:
					if o.OnQuit != nil {
						o.OnQuit()
					}
					stopAll()
					return
				}
			}
		}()
	}, func() {})
	go func() {
		select {
		case <-ctx.Done():
			stopAll()
		case <-quit:
		}
	}()

	// Where the close button cannot be caught, the window ending ends
	// the process: the tray cannot be trusted to outlive the toolkit's
	// main loop. Windows runs the tray on a thread of its own.
	trayOutlives := runtime.GOOS == "windows"
	pendingNotice := false
loop:
	for {
		prepareChrome()
		w := webview.New(false)
		w.SetTitle(o.Title)
		setWindowIcons(w.Window())
		w.SetSize(o.Width, o.Height, webview.HintNone)
		w.SetSize(MinWidth, MinHeight, webview.HintMin)
		// After the size, which the centring needs.
		applyChrome(w.Window())
		if pendingNotice && o.Notice != "" {
			w.Init("window.__trayNotice = " + jsString(o.Notice) + ";")
			pendingNotice = false
		}
		var hide func()
		var showFn func()
		if trayOutlives {
			var ok bool
			showFn, hide, ok = installCloseHook(w.Window(), func() bool {
				if o.FirstClose != nil && o.FirstClose() && o.Notice != "" {
					w.Eval("window.trayNotice && window.trayNotice(" + jsString(o.Notice) + ");")
					time.AfterFunc(NoticeFor, func() { w.Dispatch(func() { hide() }) })
					return false
				}
				return true
			})
			if !ok {
				showFn, hide = nil, nil
			}
		}
		w.Navigate(o.URL)
		mu.Lock()
		cur, show = w, showFn
		mu.Unlock()
		select {
		case <-quit:
			// Asked to quit while the window was being made.
			w.Terminate()
		default:
		}
		w.Run()
		mu.Lock()
		cur, show = nil, nil
		mu.Unlock()
		w.Destroy()

		select {
		case <-quit:
			break loop
		case <-ctx.Done():
			break loop
		default:
		}
		if !trayOutlives {
			break loop
		}
		// The window was destroyed (the close could not be caught):
		// the companion keeps running until the player opens it again
		// or quits, and the next window says so.
		if o.FirstClose != nil && o.FirstClose() {
			pendingNotice = true
		}
		select {
		case <-openReq:
		case <-quit:
			break loop
		case <-ctx.Done():
			break loop
		}
	}
	stopAll()
	systray.Quit()
	return nil
}

// setTrayIcon is the seal, with a dot in the headline's tone when it has
// one. Windows takes an ICO; macOS and Linux take a colour PNG, not a
// template image: a template is drawn from its alpha channel alone, which
// would turn the seal's navy disc into a solid blob in the menu bar.
func setTrayIcon(tone string) {
	if runtime.GOOS == "windows" {
		systray.SetIcon(icon.TrayICO(tone))
		return
	}
	systray.SetIcon(icon.TrayPNG(tone))
}

// jsString is s as a JavaScript string literal.
func jsString(s string) string {
	b, err := json.Marshal(s)
	if err != nil {
		return `""`
	}
	return string(b)
}
