// companion/internal/install/offer.go
// The first-launch window and what follows it. The window is ui/install.html
// in the companion's own look, served over loopback like the main window, and
// its one button calls Move through a webview binding. Move installs the app,
// starts the installed copy with Finish's flags and closes the window; this
// process then ends, and the installed copy ejects the disk image, trashes the
// download, registers itself to open at login and goes straight into the
// companion's own setup (sign-in, the game folder).
package install

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/companion/ui"
)

// PromptPage is the first-launch window's page among the embedded assets.
const PromptPage = "install.html"

// WindowWidth and WindowHeight are the first-launch window's fixed size:
// wide enough that the band keeps the name beside the seal (ui/styles.css
// drops it at 640 and under), tall enough for the question, the toggle and
// the button without scrolling.
const (
	WindowWidth  = 660
	WindowHeight = 400
)

// System is what installing touches outside the file system's renames; the
// darwin build's is the real one (NewSystem) and tests bring their own.
type System interface {
	// Copy copies an app bundle, signature and all, and clears the download
	// quarantine on the copy (the player opened this very app already).
	Copy(src, dst string) error
	// Relaunch opens the installed app with args, as a new process.
	Relaunch(app string, args []string) error
	// LoginItems reports whether the system can open the app at login
	// (SMAppService, macOS 13 and later).
	LoginItems() bool
	// SetLoginItem registers or unregisters the running app to open at login.
	SetLoginItem(on bool) error
	// Eject unmounts a disk image's volume.
	Eject(volume string) error
	// Trash moves a file or folder to the Trash.
	Trash(path string) error
}

// Window is what a binding may do to its window, from any goroutine.
type Window interface {
	// Close ends the window.
	Close()
	// Eval runs JavaScript in the page (the progress the Go side reports).
	Eval(js string)
}

// Prompter shows the window at url until it is closed. bind gives the
// functions the page may call, by name.
type Prompter func(url string, bind func(win Window) map[string]any) error

// The page's modes (ui/install.js): the Mac's move, the Windows setup and
// its uninstall.
const (
	ModeMove      = "move"
	ModeSetup     = "setup"
	ModeUninstall = "uninstall"
)

// OfferOptions is what Offer needs.
type OfferOptions struct {
	Plan   Plan
	Home   string
	Log    *slog.Logger
	System System
	Prompt Prompter
	// BeforeRelaunch runs just before the installed copy starts: the
	// companion lets go of its one-copy claim there.
	BeforeRelaunch func()
	// Destinations are where the app may go: Destinations(Home) unless a
	// test says otherwise.
	Destinations []string
}

// Offer shows the window and reports whether the app was installed and the
// installed copy started, in which case this process should end. It returns
// false and the window's error when the window could not be shown (a
// headless build): the companion then runs from where it is.
func Offer(o OfferOptions) (bool, error) {
	srv, url, err := serve(ModeMove)
	if err != nil {
		return false, err
	}
	defer srv.Close()
	installed := false
	err = o.Prompt(url, func(win Window) map[string]any {
		return map[string]any{
			"fsInstallInfo": func() map[string]any {
				return map[string]any{"loginItems": o.System.LoginItems()}
			},
			"fsInstallMove": func(openAtLogin bool) (map[string]any, error) {
				target, err := Move(o, openAtLogin)
				if err != nil {
					o.Log.Warn("could not install the app", "component", "install", "err", err.Error())
					return nil, err
				}
				installed = true
				win.Close()
				return map[string]any{"path": target}, nil
			},
		}
	})
	return installed, err
}

// Move installs the app and starts the installed copy, telling it what to
// clean up.
func Move(o OfferOptions, openAtLogin bool) (string, error) {
	dests := o.Destinations
	if dests == nil {
		dests = Destinations(o.Home)
	}
	target, err := Install(o.Plan.Source, dests, o.System.Copy)
	if err != nil {
		return "", err
	}
	o.Log.Info("installed the app", "component", "install", "from", o.Plan.Original, "to", target)
	if o.BeforeRelaunch != nil {
		o.BeforeRelaunch()
	}
	if err := o.System.Relaunch(target, FinishArgs(o.Plan, openAtLogin)); err != nil {
		return "", fmt.Errorf("start the app from %s: %w", target, err)
	}
	return target, nil
}

// serve answers the embedded assets on a loopback port. The page holds
// nothing private: what it can do is the bindings', not the server's.
func serve(mode string) (*http.Server, string, error) {
	assets, err := fs.Sub(ui.Assets, ".")
	if err != nil {
		return nil, "", err
	}
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, "", fmt.Errorf("listen for the install window: %w", err)
	}
	srv := &http.Server{Handler: http.FileServerFS(assets), ReadHeaderTimeout: 5 * time.Second}
	go func() { _ = srv.Serve(ln) }()
	return srv, "http://" + ln.Addr().String() + "/" + PromptPage + "?mode=" + mode, nil
}

// Finish is what the installed copy is told to do after a move.
type Finish struct {
	// Moved is true when this launch is the one a move started.
	Moved bool
	// OpenAtLogin is the first window's toggle.
	OpenAtLogin bool
	// Eject is the disk image's volume, Trash the downloaded file.
	Eject, Trash string
}

// The flags a move passes to the installed copy.
const (
	flagMoved       = "installed"
	flagOpenAtLogin = "open-at-login"
	flagEject       = "eject"
	flagTrash       = "trash-download"
)

// RegisterFlags adds Finish's flags to the companion's flag set.
func RegisterFlags(fs *flag.FlagSet) *Finish {
	f := &Finish{}
	fs.BoolVar(&f.Moved, flagMoved, false, "set by the first launch after it installs the app")
	fs.BoolVar(&f.OpenAtLogin, flagOpenAtLogin, false, "with -installed: open the app at login")
	fs.StringVar(&f.Eject, flagEject, "", "with -installed: the disk image volume to eject")
	fs.StringVar(&f.Trash, flagTrash, "", "with -installed: the download to put in the Trash")
	return f
}

// FinishArgs are the installed copy's arguments for plan.
func FinishArgs(plan Plan, openAtLogin bool) []string {
	args := []string{"-" + flagMoved, fmt.Sprintf("-%s=%t", flagOpenAtLogin, openAtLogin)}
	if plan.Volume != "" {
		args = append(args, "-"+flagEject+"="+plan.Volume)
	}
	if plan.Trash != "" {
		args = append(args, "-"+flagTrash+"="+plan.Trash)
	}
	return args
}

// EjectTries and EjectEvery bound the wait for the first copy to let go of
// the disk image: it ends a moment after it starts this one.
const (
	EjectTries = 20
	EjectEvery = 500 * time.Millisecond
)

// ErrNotADownload says a trash path was not one a move would hand over.
var ErrNotADownload = errors.New("install: not a downloaded disk image or app in the home folder")

// Complete does what a move asked: registers the login item, ejects the
// disk image once the first copy has let go of it, and trashes the download.
// Each step is logged and none stops the others; it returns what failed.
func (f Finish) Complete(ctx context.Context, sys System, home string, log *slog.Logger) error {
	if !f.Moved {
		return nil
	}
	var errs []error
	if sys.LoginItems() {
		if err := sys.SetLoginItem(f.OpenAtLogin); err != nil {
			errs = append(errs, fmt.Errorf("open at login: %w", err))
		} else {
			log.Info("set the login item", "component", "install", "on", f.OpenAtLogin)
		}
	}
	if f.Eject != "" {
		if err := ejectWhenFree(ctx, sys, f.Eject); err != nil {
			errs = append(errs, fmt.Errorf("eject %s: %w", f.Eject, err))
		} else {
			log.Info("ejected the disk image", "component", "install", "volume", f.Eject)
		}
	}
	if f.Trash != "" {
		if err := trashDownload(sys, f.Trash, home); err != nil {
			errs = append(errs, fmt.Errorf("trash %s: %w", f.Trash, err))
		} else {
			log.Info("put the download in the Trash", "component", "install", "path", f.Trash)
		}
	}
	err := errors.Join(errs...)
	if err != nil {
		log.Warn("the install's clean-up did not all happen", "component", "install", "err", err.Error())
	}
	return err
}

func ejectWhenFree(ctx context.Context, sys System, volume string) error {
	var err error
	for range EjectTries {
		if err = sys.Eject(volume); err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(EjectEvery):
		}
	}
	return err
}

// trashDownload trashes only what a move hands over: a disk image or an app
// in the player's home folder, never Applications.
func trashDownload(sys System, path, home string) error {
	clean := filepath.Clean(path)
	isDownload := strings.HasSuffix(clean, ".dmg") || strings.HasSuffix(clean, ".app")
	if !isDownload || !within(clean, home) || InApplications(clean, home) {
		return ErrNotADownload
	}
	return sys.Trash(clean)
}
