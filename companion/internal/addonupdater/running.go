// companion/internal/addonupdater/running.go
package addonupdater

import (
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// Proc is one running process: its executable's name and, where the
// system would tell, the full path of the image it runs.
type Proc struct {
	Name string
	Path string
}

// clientName matches the game's executables: Wow.exe, WowB.exe, WowT.exe,
// WowClassic.exe, WowClassicB.exe, Wow-64.exe and their kin.
var clientName = regexp.MustCompile(`(?i)^wow[a-z0-9_-]*\.exe$`)

// IsClientName reports whether an executable name is a World of Warcraft
// client's.
func IsClientName(name string) bool { return clientName.MatchString(name) }

// GameRunning reports whether a game client is running for the install, the
// flavour folder that holds WTF and Interface. The process list is the
// truth: a client counts when its image lives under the install, or sits
// beside it in the folder above (the older single-folder layout). A client
// whose path the system would not tell (an elevated process, a Wine
// prefix the path of which cannot be mapped) counts for every install,
// because an update that waited for nothing is worse than one that
// waited for the wrong game.
func GameRunning(install string, procs []Proc) bool {
	for _, p := range procs {
		if !IsClientName(p.Name) {
			continue
		}
		if p.Path == "" {
			return true
		}
		if within(install, p.Path) || sameDir(filepath.Dir(install), filepath.Dir(p.Path)) {
			return true
		}
	}
	return false
}

// within reports whether path lies under dir. Windows paths are compared
// without regard to case.
func within(dir, path string) bool {
	d, p := fold(filepath.Clean(dir)), fold(filepath.Clean(path))
	return p == d || strings.HasPrefix(p, d+string(filepath.Separator))
}

func sameDir(a, b string) bool { return fold(filepath.Clean(a)) == fold(filepath.Clean(b)) }

func fold(s string) string {
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		return strings.ToLower(s)
	}
	return s
}

// wineArg recognises a command-line argument (of a Wine loader) that is a client's Windows
// path ("Z:\home\me\wow\_classic_era_\WowClassic.exe").
func wineArg(arg string) (Proc, bool) {
	i := strings.LastIndexAny(arg, `\/`)
	name := arg[i+1:]
	if !IsClientName(name) {
		return Proc{}, false
	}
	if i < 0 {
		return Proc{Name: name}, true
	}
	if strings.ContainsRune(arg, '\\') {
		// A drive other than Z: cannot be mapped to a place on this disk, so
		// the path is left unknown and counts for every install.
		if len(arg) < 3 || (arg[0] != 'Z' && arg[0] != 'z') || arg[1] != ':' {
			return Proc{Name: name}, true
		}
		return Proc{Name: name, Path: strings.ReplaceAll(arg[2:], `\`, "/")}, true
	}
	return Proc{Name: name, Path: arg}, true
}
