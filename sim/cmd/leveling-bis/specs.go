package main

import (
	"fmt"
	"slices"
	"strings"
)

// -all's list of specs. bis.yml ranks as a matrix of legs, one per class
// (tools/bis-matrix.sh splits the written specs by class), and each leg
// names its specs with -specs. With no -specs, -all ranks every written
// spec, as a local `make bis` does. This replaced data/curated/bis-shard.json
// (2026-11-12): the matrix is what fits a night inside a job's cap, so one
// way to narrow a run is enough.

// parseSpecsFlag splits -specs's comma list; blanks are dropped so a
// trailing comma names nothing.
func parseSpecsFlag(value string) []string {
	var out []string
	for _, s := range strings.Split(value, ",") {
		if s = strings.TrimSpace(s); s != "" {
			out = append(out, s)
		}
	}
	return out
}

// keepWritten keeps the written specs that names lists, in written order. A
// name that is not written fails: a typo would otherwise rank nothing and
// leave that spec's old file standing.
func keepWritten(written, names []string) ([]string, error) {
	for _, s := range names {
		if !slices.Contains(written, s) {
			return nil, fmt.Errorf("-specs names %q, which is not a spec with a written rotation", s)
		}
	}
	out := make([]string, 0, len(names))
	for _, s := range written {
		if slices.Contains(names, s) {
			out = append(out, s)
		}
	}
	return out, nil
}

// specsToRank is -all's list: the written specs, narrowed to only when a
// leg names its specs.
func specsToRank(repoRoot string, only []string) ([]string, error) {
	written, err := writtenSpecs(repoRoot)
	if err != nil {
		return nil, err
	}
	if len(only) == 0 {
		return written, nil
	}
	return keepWritten(written, only)
}
