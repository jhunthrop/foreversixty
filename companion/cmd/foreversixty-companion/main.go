// companion/cmd/foreversixty-companion/main.go
// Command foreversixty-companion watches the game's combat log,
// uploads fights as they close, keeps the addon in step with the
// site, and shows all of it in a tray window.
//
//	foreversixty-companion              tray and window
//	foreversixty-companion -headless    no window; the URL is printed
//	foreversixty-companion -version     print the version and exit
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/logging"
	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/secret"
	"github.com/jhunthrop/foreversixty/companion/internal/shell"
	"github.com/jhunthrop/foreversixty/companion/internal/updater"
)

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
	// Before anything else: a staged update becomes the program. It
	// cannot happen while the binary is running on Windows, so it
	// happens here and takes effect on the next launch.
	if exe, err := os.Executable(); err == nil {
		if swapped, err := updater.ApplyPending(dirs.Update, updater.PublicKey, exe); err != nil {
			fmt.Fprintln(os.Stderr, "foreversixty-companion: the staged update was refused:", err)
		} else if swapped {
			fmt.Fprintln(out, "An update was applied. Start the companion again to run it.")
			return nil
		}
	}

	log, logFile, err := logging.New(dirs.Logs, *debug)
	if err != nil {
		return err
	}
	defer logFile.Close()

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

	url, ln, err := a.Serve()
	if err != nil {
		return err
	}
	defer ln.Close()
	log.Info("the companion is running", "component", "main",
		"version", updater.Version, "home", dirs.Home, "ui", url)
	if *headless {
		fmt.Fprintln(out, "Companion UI:", url)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go a.Run(ctx)

	// The shell owns the main goroutine: every desktop toolkit here
	// requires its loop to run on it.
	return shell.Run(ctx, shell.Options{URL: url, Headless: *headless, OnQuit: stop})
}
