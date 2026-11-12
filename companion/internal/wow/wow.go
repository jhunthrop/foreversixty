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
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/copygen"
)

// LogGlob matches the combat logs the game writes. The plain
// WoWCombatLog.txt is the live one; the game may also leave
// timestamped siblings behind.
const LogGlob = "WoWCombatLog*.txt"

// LiveFlavor is the live game's flavour folder (live since 4 November
// 2026), the one every message tells the player to pick, and the single
// source of truth for which folder is Forever's.
const LiveFlavor = "_forever_"

// ForeverFlavors are the flavour folders Forever Sixty runs on, and the
// only ones the companion lists, watches, uploads from or writes to: the
// live game's alone. The Classic beta's _classic_beta_ ended at launch and
// its characters are not live ones, so it is another game now, like
// Retail, Anniversary and Era; a private server's folder is not ours.
var ForeverFlavors = []string{LiveFlavor}

// FlavorInfo is the file the game writes in every flavour folder, naming
// the client's product: "Product Flavor!STRING:0" and then the product,
// "wowf" for the live game and "wow_classic_beta" for the ended beta.
const FlavorInfo = ".flavor.info"

// foreverProducts maps each Forever product to the flavour it is, so an
// install in a renamed or custom folder is still the live game.
var foreverProducts = map[string]string{"wowf": LiveFlavor}

// otherProducts are the other games a player may have beside Forever, so
// a refusal names the game the file says, not the folder it sits in. The
// ended Classic beta is one of them.
var otherProducts = map[string]string{
	"wow": "_retail_", "wow_anniversary": "_anniversary_", "wow_classic_era": "_classic_era_",
	"wow_classic_beta": "_classic_beta_",
}

// productAt is the product a flavour folder's .flavor.info names, lower
// case, or "" when the file is missing or says nothing.
func productAt(dir string) string {
	b, err := os.ReadFile(filepath.Join(dir, FlavorInfo))
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.Contains(line, "!") { // the header, "Product Flavor!STRING:0"
			continue
		}
		return strings.ToLower(line)
	}
	return ""
}

// IsForever reports whether a flavour folder name is a Forever variant.
func IsForever(folder string) bool {
	for _, f := range ForeverFlavors {
		if strings.EqualFold(f, folder) {
			return true
		}
	}
	return false
}

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
	// Product is what the folder's .flavor.info names ("wowf"), or "" where
	// it has none. A Forever product sets Flavor to its folder, whatever
	// the directory is called.
	Product string `json:"product,omitempty"`
}

// Forever reports whether the install is a Forever variant: by its
// .flavor.info product where the folder has one, by its folder name where
// it does not. The product wins, so a folder named _forever_ that holds
// another game is not ours.
func (i Install) Forever() bool {
	if i.Product != "" {
		_, ok := foreverProducts[i.Product]
		return ok
	}
	return IsForever(i.Flavor)
}

// Candidates are the paths the companion looks in before asking the
// player to pick a folder.
func Candidates() []string {
	home, _ := os.UserHomeDir()
	return candidatesFor(runtime.GOOS, home, os.Getenv)
}

// candidatesFor is Candidates for one system, home folder and
// environment, so every system's list is testable on any of them. The
// Battle.net launcher installs into /Applications on a Mac unless the
// player chose their own Applications folder; a Mac flavour folder holds
// WTF and Logs beside the client's .app exactly as on Windows.
func candidatesFor(goos, home string, getenv func(string) string) []string {
	switch goos {
	case "darwin":
		return []string{
			"/Applications/World of Warcraft",
			filepath.Join(home, "Applications", "World of Warcraft"),
		}
	case "windows":
		var out []string
		for _, env := range []string{"ProgramFiles(x86)", "ProgramFiles"} {
			if v := getenv(env); v != "" {
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

// Detect scans every candidate path.
func Detect() []Install { return detectIn(Candidates()) }

// detectIn is Detect over the given roots.
func detectIn(roots []string) []Install {
	var out []Install
	seen := map[string]bool{}
	for _, root := range roots {
		for _, in := range Scan(root) {
			if !seen[in.Path] {
				seen[in.Path] = true
				out = append(out, in)
			}
		}
	}
	return out
}

// isInstall reports whether a directory is a flavour directory. WTF is
// the marker: every install has one and nothing else does.
func isInstall(dir string) bool {
	fi, err := os.Stat(filepath.Join(dir, "WTF"))
	return err == nil && fi.IsDir()
}

func installAt(dir string) Install {
	in := Install{
		Path:    dir,
		Flavor:  filepath.Base(dir),
		Logs:    filepath.Join(dir, "Logs"),
		WTF:     filepath.Join(dir, "WTF"),
		Product: productAt(dir),
	}
	if flavor, ok := foreverProducts[in.Product]; ok {
		in.Flavor = flavor
	} else if flavor, ok := otherProducts[in.Product]; ok {
		in.Flavor = flavor
	}
	return in
}

// Scan returns the Forever installs at or one level below root, by path,
// so two runs agree. Other games' folders, the ended beta's among them,
// are not returned.
func Scan(root string) []Install {
	var out []Install
	for _, in := range scanAll(root) {
		if in.Forever() {
			out = append(out, in)
		}
	}
	return out
}

// scanAll returns every install at or one level below root, Forever
// variant or not.
func scanAll(root string) []Install {
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

// ErrNotAnInstall says a picked folder is not the game.
var ErrNotAnInstall = errors.New("that folder is not a World of Warcraft install: " +
	"pick the folder that contains WTF and Logs, such as " + LiveFlavor)

// Refusal says a folder is a World of Warcraft install that is not a
// Forever variant.
type Refusal struct {
	// Folder is the refused flavour folder's name and Path where it is.
	Folder string
	Path   string
	// Known is true when Folder is a game the player may well have
	// (Retail, Classic Anniversary), so the words can name it.
	Known bool
}

func (r *Refusal) Error() string { return "not a Forever variant: " + r.Path }

// PickWords names the folder to pick, in the player's words.
const PickWords = copygen.FolderPick

// Words is the refusal in the player's words.
func (r *Refusal) Words() string {
	if r.Known {
		return copygen.Fill(copygen.FolderThatIs, "{flavor}", FlavorName(r.Folder)) + " " + PickWords
	}
	return PickWords
}

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
		if other := scanAll(path); len(other) > 0 {
			r := &Refusal{Folder: other[0].Flavor, Path: other[0].Path}
			// A Forever folder name that its file says is some other, unknown
			// product is not named: "That is WoW Forever" would be wrong.
			r.Known = len(other) == 1 && FlavorName(r.Folder) != r.Folder && !IsForever(r.Folder)
			return nil, r
		}
		return nil, ErrNotAnInstall
	}
	return found, nil
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

// FlavorName is a flavour folder in the player's words. The folders a
// player has are named by the game ("_anniversary_"), and the status
// page and the settings list show what they mean; a private server's
// folder, or one this table does not know, keeps its own name.
func FlavorName(folder string) string {
	switch folder {
	case "_anniversary_":
		return copygen.FlavorAnniversary
	case "_classic_era_":
		return copygen.FlavorClassicEra
	case LiveFlavor:
		return copygen.FlavorForever
	case "_classic_beta_":
		return copygen.FlavorClassicBeta
	case "_retail_":
		return copygen.FlavorRetail
	}
	return folder
}

// LastActivity is the newest modification time among the regular files
// directly under the install's Logs directory. The game rewrites its
// General, gx, cpu and threadpool logs at launch and the combat log
// while logging, so this is when the install was last played. It does
// not look into subfolders; a missing or empty Logs directory is false.
func LastActivity(in Install) (time.Time, bool) {
	entries, err := os.ReadDir(in.Logs)
	if err != nil {
		return time.Time{}, false
	}
	var newest time.Time
	found := false
	for _, e := range entries {
		if !e.Type().IsRegular() {
			continue
		}
		fi, err := e.Info()
		if err != nil {
			continue
		}
		if !found || fi.ModTime().After(newest) {
			newest, found = fi.ModTime(), true
		}
	}
	return newest, found
}
