package main

import (
	"fmt"
	"strings"
)

// resolveKeep maps the -keep talent names onto active-build talent ids.
// A name no tree of the class carries is an error, so a typo cannot
// silently protect nothing; a name the guide does not take protects
// nothing and is allowed (a spec's keep list may name a talent only some
// of its builds hold).
func resolveKeep(t talentTrees, names []string) (map[int]bool, error) {
	byName := make(map[string]int, len(t.byID))
	for id, n := range t.byID {
		byName[n.Name] = id
	}
	keep := make(map[int]bool, len(names))
	for _, name := range names {
		id, ok := byName[name]
		if !ok {
			return nil, fmt.Errorf("-keep: no talent named %q in this class's trees", name)
		}
		keep[id] = true
	}
	return keep, nil
}

// splitKeep is the -keep flag's comma-separated names, trimmed.
func splitKeep(s string) []string {
	var out []string
	for _, name := range strings.Split(s, ",") {
		if name = strings.TrimSpace(name); name != "" {
			out = append(out, name)
		}
	}
	return out
}
