// companion/internal/install/setup_window.go
// The window around RunSetup and RunUninstall: the page asks, the steps run
// on a goroutine of their own, and each step's start reaches the page as it
// happens, so the progress is real.
package install

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

// DoneFor is how long the uninstall's last line stays up before the window closes.
const DoneFor = 1500 * time.Millisecond

// SetupOptions is what ShowSetup needs.
type SetupOptions struct {
	// Mode is ModeSetup or ModeUninstall.
	Mode    string
	System  WinSystem
	Folders Folders
	// Self is the running exe, Version its version.
	Self, Version string
	Log           *slog.Logger
	Prompt        Prompter
}

// ShowSetup shows the setup or uninstall window and returns once it closes.
func ShowSetup(o SetupOptions) error {
	srv, url, err := serve(o.Mode)
	if err != nil {
		return err
	}
	defer srv.Close()
	return o.Prompt(url, func(win Window) map[string]any {
		return map[string]any{
			"fsSetupInfo": func() map[string]any {
				_, statErr := os.Stat(o.Folders.Exe())
				return map[string]any{"upgrade": statErr == nil}
			},
			// The page's one toggle: Start with Windows, or Keep my settings.
			"fsSetupRun": func(toggle bool) {
				go runInWindow(o, win, toggle)
			},
		}
	})
}

func runInWindow(o SetupOptions, win Window, toggle bool) {
	progress := func(step, of int, key string) {
		win.Eval(fmt.Sprintf("window.fsProgress(%d, %d, %q)", step, of, key))
	}
	var err error
	if o.Mode == ModeUninstall {
		err = RunUninstall(o.System, o.Folders, toggle, progress, o.Log)
	} else {
		err = RunSetup(o.System, o.Folders, o.Self, o.Version, toggle, progress, o.Log)
	}
	if err != nil {
		o.Log.Warn("the "+o.Mode+" did not finish", "component", "setup", "err", err.Error())
		win.Eval("window.fsFailed()")
		return
	}
	if o.Mode == ModeUninstall {
		win.Eval("window.fsDone()")
		time.Sleep(DoneFor)
	}
	win.Close()
}
