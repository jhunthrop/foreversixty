//go:build linux

package addonupdater

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ListProcesses reads /proc. A native process is named by its executable's
// link; a game under Wine is the wine loader, so every argument of its
// command line is read too, and one that names a client executable stands
// for it, its Windows path mapped onto the Z: drive's view of the disk.
func ListProcesses() ([]Proc, error) {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}
	var out []Proc
	for _, e := range entries {
		if _, err := strconv.Atoi(e.Name()); err != nil {
			continue
		}
		dir := filepath.Join("/proc", e.Name())
		if exe, err := os.Readlink(filepath.Join(dir, "exe")); err == nil {
			out = append(out, Proc{Name: filepath.Base(exe), Path: exe})
		}
		cmd, err := os.ReadFile(filepath.Join(dir, "cmdline"))
		if err != nil {
			continue
		}
		for _, arg := range strings.Split(string(cmd), "\x00") {
			if p, ok := wineArg(arg); ok {
				out = append(out, p)
			}
		}
	}
	return out, nil
}
