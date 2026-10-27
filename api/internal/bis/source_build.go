package bis

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// SourceBuild is the build whose bis/ directory serves active, the site's own rule
// (web/scripts/bis-source-build.mjs): active itself when it has BiS files, else the newest
// other build that has them, else active. An active build with no directory never falls back.
func SourceBuild(dataDir, active string) string {
	if !dirExists(filepath.Join(dataDir, active)) || hasBisFiles(dataDir, active) {
		return active
	}
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return active
	}
	best := ""
	for _, e := range entries {
		name := e.Name()
		if !e.IsDir() || name == active || !hasBisFiles(dataDir, name) {
			continue
		}
		if best == "" || buildNewer(name, best) {
			best = name
		}
	}
	if best == "" {
		return active
	}
	return best
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// hasBisFiles reports whether build's bis/ holds a published <spec>.json.
func hasBisFiles(dataDir, build string) bool {
	entries, err := os.ReadDir(filepath.Join(dataDir, build, "bis"))
	if err != nil {
		return false
	}
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".json") {
			return true
		}
	}
	return false
}

// buildNewer compares dotted build ids numerically, field by field (a missing or
// non-numeric field counts as 0, as the site's compareBuildsDescending does).
func buildNewer(a, b string) bool {
	left, right := strings.Split(a, "."), strings.Split(b, ".")
	for i := 0; i < max(len(left), len(right)); i++ {
		l, r := buildField(left, i), buildField(right, i)
		if l != r {
			return l > r
		}
	}
	return false
}

func buildField(parts []string, i int) int {
	if i >= len(parts) {
		return 0
	}
	n, _ := strconv.Atoi(parts[i])
	return n
}
