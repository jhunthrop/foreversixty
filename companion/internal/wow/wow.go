// Package wow finds the game on disk. An install is a flavour
// directory — the folder holding WTF and Logs — because that is the
// unit the game rotates logs and saves variables in, and because a
// private server's folder may be named anything at all.
package wow

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

// LogGlob matches the combat logs the game writes. The plain
// WoWCombatLog.txt is the live one; the game may also leave
// timestamped siblings behind.
const LogGlob = "WoWCombatLog*.txt"

// Install is one flavour directory of the game.
type Install struct {
	// Path is the flavour directory itself.
	Path string `json:"path"`
	// Flavor is its folder name, "_classic_era_" and the like, or
	// the folder name of a private-server install.
	Flavor string `json:"flavor"`
	// Logs is <Path>/Logs and WTF is <Path>/WTF.
	Logs string `json:"logs"`
	WTF  string `json:"wtf"`
}

// Candidates are the paths the companion looks in before asking the
// player to pick a folder.
func Candidates() []string {
	home, _ := os.UserHomeDir()
	switch runtime.GOOS {
	case "darwin":
		return []string{
			"/Applications/World of Warcraft",
			filepath.Join(home, "Applications", "World of Warcraft"),
		}
	case "windows":
		var out []string
		for _, env := range []string{"ProgramFiles(x86)", "ProgramFiles"} {
			if v := os.Getenv(env); v != "" {
				out = append(out, filepath.Join(v, "World of Warcraft"))
			}
		}
		return append(out, `C:\Program Files (x86)\World of Warcraft`)
	default:
		return []string{
			filepath.Join(home, "Games", "world-of-warcraft", "drive_c",
				"Program Files (x86)", "World of Warcraft"),
			filepath.Join(home, ".wine", "drive_c",
				"Program Files (x86)", "World of Warcraft"),
		}
	}
}

// isInstall reports whether a directory is a flavour directory. WTF is
// the marker: every install has one and nothing else does.
func isInstall(dir string) bool {
	fi, err := os.Stat(filepath.Join(dir, "WTF"))
	return err == nil && fi.IsDir()
}

func installAt(dir string) Install {
	return Install{
		Path:   dir,
		Flavor: filepath.Base(dir),
		Logs:   filepath.Join(dir, "Logs"),
		WTF:    filepath.Join(dir, "WTF"),
	}
}

// Scan returns every install at or one level below root, sorted by
// path so two runs agree.
func Scan(root string) []Install {
	var out []Install
	if isInstall(root) {
		out = append(out, installAt(root))
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		return out
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		dir := filepath.Join(root, e.Name())
		if isInstall(dir) {
			out = append(out, installAt(dir))
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	return out
}

// Detect scans every candidate path.
func Detect() []Install {
	var out []Install
	seen := map[string]bool{}
	for _, root := range Candidates() {
		for _, in := range Scan(root) {
			if !seen[in.Path] {
				seen[in.Path] = true
				out = append(out, in)
			}
		}
	}
	return out
}

// ErrNotAnInstall says a picked folder is not the game.
var ErrNotAnInstall = errors.New("that folder is not a World of Warcraft install: " +
	"pick the folder that contains WTF and Logs, such as _classic_era_")

// Pick validates a folder the player chose. It accepts either a
// flavour directory or the folder above one, which is what a file
// dialog most often returns.
func Pick(path string) ([]Install, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	if !fi.IsDir() {
		return nil, ErrNotAnInstall
	}
	found := Scan(path)
	if len(found) == 0 {
		return nil, ErrNotAnInstall
	}
	return found, nil
}

// EnsureLogs creates the Logs directory if the player has never
// logged, so the watcher has something to watch from the first run.
func (i Install) EnsureLogs() error {
	return os.MkdirAll(i.Logs, 0o755)
}

// CombatLogs lists the combat logs in the install, newest last.
func (i Install) CombatLogs() ([]string, error) {
	matches, err := filepath.Glob(filepath.Join(i.Logs, LogGlob))
	if err != nil {
		return nil, err
	}
	sort.Strings(matches)
	return matches, nil
}

// AdvancedLoggingPath is the settings file the checkbox writes to.
func (i Install) AdvancedLoggingPath() string { return filepath.Join(i.WTF, "Config.wtf") }

// AdvancedLogging reports whether advanced combat logging is on. A
// missing setting means off: the game omits the line until the box is
// ticked once.
func (i Install) AdvancedLogging() (bool, error) {
	f, err := os.Open(i.AdvancedLoggingPath())
	if err != nil {
		return false, fmt.Errorf("read %s: %w", i.AdvancedLoggingPath(), err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) == 3 && fields[0] == "SET" && fields[1] == "advancedCombatLogging" {
			return strings.Trim(fields[2], `"`) == "1", nil
		}
	}
	if err := sc.Err(); err != nil {
		return false, err
	}
	return false, nil
}

// AdvancedLoggingHelp is the exact menu path the status page shows when
// advanced logging is off, per spec section 5.
const AdvancedLoggingHelp = "Options → Network → Advanced Combat Logging"

// Accounts lists the account folders under WTF/Account, sorted. The
// game names them after the account, upper-cased.
func (i Install) Accounts() ([]string, error) {
	dir := filepath.Join(i.WTF, "Account")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", dir, err)
	}
	var out []string
	for _, e := range entries {
		if e.IsDir() && e.Name() != "SavedVariables" {
			out = append(out, filepath.Join(dir, e.Name()))
		}
	}
	sort.Strings(out)
	return out, nil
}

// SavedVariables returns the path of one addon's saved-variables file
// in each account folder, whether or not it exists yet: the addon-sync
// loop watches paths that may appear later.
func (i Install) SavedVariables(addon string) ([]string, error) {
	accounts, err := i.Accounts()
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(accounts))
	for _, a := range accounts {
		out = append(out, filepath.Join(a, "SavedVariables", addon+".lua"))
	}
	return out, nil
}
