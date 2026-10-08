package leveling

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// BisDir is the bis/ directory that serves a build's BiS files.
//
// The nightly ranks a build only after the site activates it, so for up to a
// day the active build has no bis/ files. Tools that read a finished BiS set
// as their input (the gear a talent or rotation search wears) then read the
// newest OTHER build that has them; each file records the build it was ranked
// on in its own "build" field. A build with its own bis/ files, or a
// directory that is not a build at all, answers its own bis/ path unchanged.
// It mirrors web/scripts/bis-source-build.mjs, which the site uses.
func BisDir(buildDir string) string {
	own := filepath.Join(buildDir, "bis")
	if hasBisFiles(own) {
		return own
	}
	if _, err := os.Stat(buildDir); err != nil {
		return own
	}
	siblings, err := os.ReadDir(filepath.Dir(buildDir))
	if err != nil {
		return own
	}
	var ranked []string
	for _, entry := range siblings {
		candidate := filepath.Join(filepath.Dir(buildDir), entry.Name())
		if entry.IsDir() && candidate != buildDir && hasBisFiles(filepath.Join(candidate, "bis")) {
			ranked = append(ranked, entry.Name())
		}
	}
	if len(ranked) == 0 {
		return own
	}
	sort.Slice(ranked, func(i, j int) bool { return buildIDNewer(ranked[i], ranked[j]) })
	return filepath.Join(filepath.Dir(buildDir), ranked[0], "bis")
}

// hasBisFiles reports whether dir holds at least one published <spec>.json
// (the sibling <spec>.md reports do not count).
func hasBisFiles(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".json") {
			return true
		}
	}
	return false
}

// buildIDNewer orders dotted client-build ids numerically, field by field
// (1.60.1.70291 is newer than 1.60.1.70009).
func buildIDNewer(a, b string) bool {
	left, right := strings.Split(a, "."), strings.Split(b, ".")
	for i := 0; i < len(left) || i < len(right); i++ {
		if delta := buildField(left, i) - buildField(right, i); delta != 0 {
			return delta > 0
		}
	}
	return false
}

func buildField(fields []string, i int) int {
	if i >= len(fields) {
		return 0
	}
	n, err := strconv.Atoi(fields[i])
	if err != nil {
		return 0
	}
	return n
}
