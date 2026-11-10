// companion/cmd/foreversixty-companion/main.go
// Command foreversixty-companion watches the game's combat log,
// uploads fights as they close, keeps the addon in step with the
// site, and shows all of it in a tray window.
//
//	foreversixty-companion              tray and window (closing the window keeps it in the tray)
//	foreversixty-companion -headless    no window; the URL is printed
//	foreversixty-companion -version     print the version and exit
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/browser"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/install"
	"github.com/jhunthrop/foreversixty/companion/internal/instance"
	"github.com/jhunthrop/foreversixty/companion/internal/logging"
	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/secret"
	"github.com/jhunthrop/foreversixty/companion/internal/shell"
	"github.com/jhunthrop/foreversixty/companion/internal/updater"
)

// openURL shows a link in the system browser; tests replace it.
var openURL = defaultOpen

// startNew starts the new build after an update; tests replace it.
var startNew updater.Starter = updater.StartDetached

func defaultOpen(u string) error { return browser.Open(u) }

// Every desktop toolkit here wants its loop on the main thread, and the
// first-launch window (install) opens before the shell's.
func init() { runtime.LockOSThread() }

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "foreversixty-companion:", err)
		os.Exit(1)
	}
}

func run(args []string, out *os.File) error {
	fs := flag.NewFlagSet("foreversixty-companion", flag.ContinueOnError)
	fs.SetOutput(out)
	headless := fs.Bool("headless", false, "run without a tray icon or a window")
	debug := fs.Bool("debug", false, "log at debug level")
	version := fs.Bool("version", false, "print the version and exit")
	finish := install.RegisterFlags(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *version {
		fmt.Fprintln(out, updater.Version)
		return nil
	}

	dirs, err := paths.Resolve()
	if err != nil {
		return err
	}
	// One companion per home directory: two would tail the same log and
	// upload every fight twice. The claim comes before the update swap,
	// which cannot happen under a running copy on Windows anyway.
	// The log comes first: whatever happens after this line, a launch
	// leaves a trace, so a silent exit cannot go unnoticed.
	log, logFile, err := logging.New(dirs.Logs, *debug)
	if err != nil {
		return err
	}
	defer logFile.Close()
	log.Info("launching", "component", "main", "version", updater.Version, "pid", os.Getpid())

	// One companion per home directory: two would tail the same log and
	// upload every fight twice. The claim comes before the update swap,
	// which cannot happen under a running copy on Windows anyway.
	lock, err := instance.Acquire(dirs.Home)
	if errors.Is(err, instance.ErrRunning) {
		return handOff(log, dirs)
	}
	if err != nil {
		return err
	}
	defer func() { lock.Release() }()
	// A Mac copy opened from the download moves itself to Applications and
	// starts again from there, before an update could be swapped into the
	// download.
	if !*headless && offerInstall(log, func() { lock.Release() }) {
		return nil
	}
	// Before anything else: a staged update becomes the program, and
	// the new build is started here, so one double-click after an update
	// is enough. The claim is let go just before, or the new build would
	// find it held and take this for a running copy.
	if exe, err := os.Executable(); err == nil {
		wd, _ := os.Getwd()
		if updater.ApplyAndRelaunch(updater.LaunchOptions{
			Dir: dirs.Update, PublicKey: updater.PublicKey, Exe: exe,
			Args: args, WorkDir: wd, Log: log, Start: startNew,
			BeforeStart: func() { lock.Release() },
			AfterFailedStart: func() {
				if l, err := instance.Acquire(dirs.Home); err == nil {
					lock = l
				}
			},
		}) {
			return nil
		}
	}

	cfg, err := config.Load(dirs.ConfigFile())
	if err != nil {
		return err
	}
	store := secret.Open(dirs.ConfigFile())
	if err := secret.Migrate(store, dirs.ConfigFile()); err != nil {
		log.Warn("could not move the device token into the keychain",
			"component", "main", "err", err.Error())
	}

	a, err := app.New(app.Options{Dirs: dirs, Config: cfg, Secret: store, Log: log})
	if err != nil {
		return err
	}
	defer a.Close()

	ui, err := a.Serve()
	if err != nil {
		return err
	}
	defer ui.Close()
	// The address without the token: the token in the path is the
	// whole authentication for the settings and pairing routes, and
	// this file is one the troubleshooting section asks players to
	// open and paste.
	if err := lock.PublishURL(ui.URL); err != nil {
		log.Warn("could not publish the window address for a second launch",
			"component", "main", "err", err.Error())
	}
	log.Info("the companion is running", "component", "main",
		"version", updater.Version, "home", dirs.Home, "ui", ui.Addr)
	if *headless {
		fmt.Fprintln(out, "Companion UI:", ui.URL)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	done := make(chan struct{})
	go func() {
		defer close(done)
		_ = a.Run(ctx)
	}()
	if finish.Moved && install.Supported {
		home, _ := os.UserHomeDir()
		go func() { _ = finish.Complete(ctx, install.NewSystem(), home, log) }()
	}

	// The shell owns the main goroutine: every desktop toolkit here
	// requires its loop to run on it.
	err = shell.Run(ctx, shell.Options{
		URL: ui.URL, Headless: *headless, OnQuit: stop,
		// The tray says what the Status headline says, and the window
		// closing leaves the companion running (the shell hides it).
		State:       func() string { return a.Headline().Title },
		Tone:        func() string { return a.Headline().Tone },
		OpenReports: func() { _ = browser.Open(a.ReportsURL()) },
		FirstClose:  a.FirstClose,
		Notice:      app.CopyTrayNotice,
	})
	// Where the window's close cannot be caught, closing it returns from
	// Run without ever calling OnQuit, so the cancel happens here too,
	// and the loop is waited for: the deferred a.Close releases the zstd
	// encoder, and Step must not still be inside EncodeAll.
	stop()
	<-done
	return err
}

// offerInstall shows the Mac's "Move Forever Sixty to Applications?" when
// the app runs from anywhere else, and reports whether this process should
// end: the installed copy was started, or the player closed the window.
func offerInstall(log *slog.Logger, release func()) bool {
	if !install.Supported {
		return false
	}
	exe, err := os.Executable()
	if err != nil {
		return false
	}
	// hdiutil and App Translocation name real paths (/private/var, not /var).
	if real, err := filepath.EvalSymlinks(exe); err == nil {
		exe = real
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	plan := install.Decide(updater.BundleRoot(exe), install.Detect(home))
	if !plan.Move {
		return false
	}
	log.Info("the app runs outside Applications; offering to move it", "component", "main", "from", plan.Original)
	moved, err := install.Offer(install.OfferOptions{
		Plan: plan, Home: home, Log: log, System: install.NewSystem(),
		Prompt: promptWindow, BeforeRelaunch: release,
	})
	if err != nil {
		log.Warn("could not show the move window; running from here", "component", "main", "err", err.Error())
		return false
	}
	if !moved {
		log.Info("the move window was closed; quitting until the next launch", "component", "main")
	}
	return true
}

// promptWindow is install's window: small, fixed, and the companion's own look.
func promptWindow(url string, bind func(closeWindow func()) map[string]any) error {
	return shell.Prompt(shell.PromptOptions{
		URL: url, Title: "Forever Sixty", Width: install.WindowWidth, Height: install.WindowHeight, Bind: bind,
	})
}

// handOff is what a second launch does: say so in the log, open the
// running copy's window and succeed. Nothing of the companion starts.
func handOff(log *slog.Logger, dirs paths.Dirs) error {
	log.Info("the companion is already running; showing its window", "component", "main")
	url := instance.RunningURL(dirs.Home)
	if url == "" {
		return nil // the first copy is still starting; its window is coming
	}
	if err := openURL(url); err != nil {
		log.Warn("could not open the running companion's window", "component", "main", "err", err.Error())
	}
	return nil
}
