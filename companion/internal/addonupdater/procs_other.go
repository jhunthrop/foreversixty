//go:build !windows && !linux

package addonupdater

import (
	"os/exec"
	"path/filepath"
	"strings"
)

// ListProcesses asks ps for each process's executable; on macOS comm is the
// full path of the image.
func ListProcesses() ([]Proc, error) {
	out, err := exec.Command("ps", "-axww", "-o", "comm=").Output()
	if err != nil {
		return nil, err
	}
	var procs []Proc
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		procs = append(procs, Proc{Name: filepath.Base(line), Path: line})
	}
	return procs, nil
}
