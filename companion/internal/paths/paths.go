// companion/internal/paths/paths.go
// Package paths resolves the one directory the companion owns on each
// operating system. Everything else the app writes hangs off it, so a
// test can redirect the whole application state with one variable.
package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// HomeEnv overrides the application directory. Tests set it; the
// installer never does.
const HomeEnv = "FS_COMPANION_HOME"

// Home is the application directory for this operating system:
// ~/Library/Application Support/ForeverSixty on macOS,
// %APPDATA%\ForeverSixty on Windows, ~/.config/foreversixty elsewhere.
func Home() (string, error) {
	if v := os.Getenv(HomeEnv); v != "" {
		return v, nil
	}
	base, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("locate the application directory: %w", err)
	}
	name := "foreversixty"
	if runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
		name = "ForeverSixty"
	}
	return filepath.Join(base, name), nil
}

// Dirs are the directories the companion writes into.
type Dirs struct {
	Home   string // the application directory itself
	Queue  string // pending bundles and chunks
	State  string // one file per report, holding session.State()
	Logs   string // the companion's own log files
	Update string // downloaded and verified update binaries
}

// Resolve returns the directories and creates them.
func Resolve() (Dirs, error) {
	home, err := Home()
	if err != nil {
		return Dirs{}, err
	}
	d := Dirs{
		Home:   home,
		Queue:  filepath.Join(home, "queue"),
		State:  filepath.Join(home, "state"),
		Logs:   filepath.Join(home, "logs"),
		Update: filepath.Join(home, "update"),
	}
	for _, p := range []string{d.Home, d.Queue, d.State, d.Logs, d.Update} {
		if err := os.MkdirAll(p, 0o700); err != nil {
			return Dirs{}, fmt.Errorf("create %s: %w", p, err)
		}
	}
	return d, nil
}

// ConfigFile is the path of config.json.
func (d Dirs) ConfigFile() string { return filepath.Join(d.Home, "config.json") }
