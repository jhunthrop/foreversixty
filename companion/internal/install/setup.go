// companion/internal/install/setup.go
// The Windows setup and uninstall, in the companion itself. The download
// /setup offers (ForeverSixtySetup.exe) is the signed companion under a setup
// name: run from there it shows one window in the companion's own look
// (ui/install.html, mode "setup"): a Start with Windows toggle and Install.
// Install puts the companion in the player's own Programs folder, adds the
// Start menu shortcut, the startup entry and the Apps & Features entry, starts
// the installed copy (which opens into its own first run) and closes. Apps &
// Features runs the installed copy with -uninstall, which shows the same window
// asking "Remove Forever Sixty?".
//
// The steps are here, with the system behind an interface (WinSystem), so a
// test on any platform runs them; setup_windows.go is the real system.
package install

import (
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"
)

// ExeName is the installed companion's file name, as the updater's swap and
// the startup entry name it.
const ExeName = "foreversixty-companion.exe"

// ProductName is the name in the Start menu, Apps & Features and the startup list.
const ProductName = "Forever Sixty"

// Publisher and AboutURL fill the Apps & Features entry.
const (
	Publisher = "Commish LLC"
	AboutURL  = "https://foreversixty.gg"
)

// FlagUninstall is the flag Apps & Features runs the installed copy with.
const FlagUninstall = "uninstall"

// InstallDir is where the setup installs for one Windows account:
// %LOCALAPPDATA%\Programs\Forever Sixty, which needs no admin rights.
func InstallDir(localAppData string) string {
	return filepath.Join(localAppData, "Programs", ProductName)
}

// SetupMode reports whether the exe at exe is the setup download: a setup name
// (ForeverSixtySetup.exe, or the release's foreversixty-companion-setup.exe, or
// a browser's "ForeverSixtySetup (1).exe"), run from anywhere but the install folder.
func SetupMode(exe, installDir string) bool {
	return strings.Contains(strings.ToLower(filepath.Base(exe)), "setup") &&
		!strings.EqualFold(filepath.Clean(filepath.Dir(exe)), filepath.Clean(installDir))
}

// Entry is the Apps & Features entry's values.
type Entry struct {
	DisplayName, DisplayIcon, DisplayVersion, Publisher string
	InstallLocation, UninstallString, URLInfoAbout      string
	// EstimatedSize is in kilobytes, as Apps & Features reads it.
	EstimatedSize uint32
}

// NewEntry is the entry for the companion installed at exe.
func NewEntry(exe, version string, sizeBytes int64) Entry {
	return Entry{
		DisplayName:     ProductName,
		DisplayIcon:     exe,
		DisplayVersion:  version,
		Publisher:       Publisher,
		InstallLocation: filepath.Dir(exe),
		UninstallString: quoted(exe) + " -" + FlagUninstall,
		URLInfoAbout:    AboutURL,
		EstimatedSize:   uint32((sizeBytes + 1023) / 1024),
	}
}

// WinSystem is everything the setup and the uninstall touch.
type WinSystem interface {
	// StopCompanions ends every running companion of this account but this process.
	StopCompanions() error
	// CopyExe copies src over dst, retrying while a stopped companion lets go of it,
	// and returns the copy's size.
	CopyExe(src, dst string) (int64, error)
	// CreateShortcut writes a .lnk at lnk that starts target.
	CreateShortcut(lnk, target string) error
	// SetUninstallEntry writes the Apps & Features entry; DeleteUninstallEntry removes it.
	SetUninstallEntry(Entry) error
	DeleteUninstallEntry() error
	// SetStartup writes the startup entry that starts command at sign-in;
	// DeleteStartup removes it.
	SetStartup(command string) error
	DeleteStartup() error
	// StartupEnabled reports whether the startup entry is there: an update keeps
	// the player's choice.
	StartupEnabled() bool
	// Start starts exe as a process of its own.
	Start(exe string) error
	// Remove deletes a file or a folder and what is in it; a missing one is no error.
	Remove(path string) error
	// RemoveAfterExit deletes dir once this process has ended.
	RemoveAfterExit(dir string) error
	// ForgetToken deletes the device token from Credential Manager.
	ForgetToken() error
}

// Folders are the account's folders the setup writes into.
type Folders struct {
	// Install is InstallDir; StartMenu is the account's Start menu Programs
	// folder; Data is the companion's own folder (internal/paths).
	Install, StartMenu, Data string
}

// Shortcut is the Start menu shortcut's path.
func (f Folders) Shortcut() string { return filepath.Join(f.StartMenu, ProductName+".lnk") }

// Exe is the installed companion's path.
func (f Folders) Exe() string { return filepath.Join(f.Install, ExeName) }

// Progress reports a step as it starts: its number from 1, the count, and
// its key in the page's copy (ui/copy.js, install.steps).
type Progress func(step, of int, key string)

// The setup's steps, by their keys in the page's copy.
var setupSteps = []string{"stop", "copy", "shortcut", "register", "start"}

// RunSetup installs the companion running as self into f, with or without the
// startup entry, and starts the installed copy.
func RunSetup(sys WinSystem, f Folders, self, version string, startWithWindows bool, progress Progress, log *slog.Logger) error {
	step := stepper(setupSteps, progress)
	step("stop")
	if err := sys.StopCompanions(); err != nil {
		log.Warn("could not stop a running companion", "component", "setup", "err", err.Error())
	}
	step("copy")
	exe := f.Exe()
	size, err := sys.CopyExe(self, exe)
	if err != nil {
		return fmt.Errorf("copy the companion to %s: %w", f.Install, err)
	}
	step("shortcut")
	if err := sys.CreateShortcut(f.Shortcut(), exe); err != nil {
		return fmt.Errorf("add the Start menu shortcut: %w", err)
	}
	step("register")
	if err := sys.SetUninstallEntry(NewEntry(exe, version, size)); err != nil {
		return fmt.Errorf("add the Apps & Features entry: %w", err)
	}
	if startWithWindows {
		err = sys.SetStartup(quoted(exe))
	} else {
		err = sys.DeleteStartup()
	}
	if err != nil {
		return fmt.Errorf("set Start with Windows: %w", err)
	}
	step("start")
	if err := sys.Start(exe); err != nil {
		return fmt.Errorf("start the installed companion: %w", err)
	}
	log.Info("installed the companion", "component", "setup", "to", exe, "version", version, "startup", startWithWindows)
	return nil
}

// The uninstall's steps.
var uninstallSteps = []string{"stop", "remove"}

// RunUninstall removes what RunSetup made, and the companion's own folder too
// unless keepSettings, then has the install folder removed once this process ends.
func RunUninstall(sys WinSystem, f Folders, keepSettings bool, progress Progress, log *slog.Logger) error {
	step := stepper(uninstallSteps, progress)
	step("stop")
	if err := sys.StopCompanions(); err != nil {
		log.Warn("could not stop a running companion", "component", "setup", "err", err.Error())
	}
	step("remove")
	var errs []error
	errs = append(errs, sys.DeleteStartup(), sys.Remove(f.Shortcut()), sys.DeleteUninstallEntry())
	if !keepSettings {
		errs = append(errs, sys.Remove(f.Data), sys.ForgetToken())
	}
	errs = append(errs, sys.RemoveAfterExit(f.Install))
	if err := errors.Join(errs...); err != nil {
		return err
	}
	log.Info("removed the companion", "component", "setup", "kept settings", keepSettings)
	return nil
}

// quoted is path in double quotes, as a Windows command line takes a path
// with spaces (Go's %q would double every backslash).
func quoted(path string) string { return `"` + path + `"` }

func stepper(keys []string, progress Progress) func(string) {
	return func(key string) {
		if progress == nil {
			return
		}
		for i, k := range keys {
			if k == key {
				progress(i+1, len(keys), key)
				return
			}
		}
	}
}
