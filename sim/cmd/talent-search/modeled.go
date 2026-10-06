package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
)

// talentRefRE is how engine code reads a talent field:
// paladin.Talents.Vengeance, hunter.Talents.GetSurefooted().
var talentRefRE = regexp.MustCompile(`\bTalents\.(?:Get)?([A-Z][A-Za-z0-9]*)`)

// skipEngineFile is generated or test code, which reads talents
// without modeling them.
func skipEngineFile(name string) bool {
	return !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") ||
		strings.HasSuffix(name, ".pb.go") || strings.HasSuffix(name, "_auto_gen.go")
}

// stripLineComments drops // comments, which in this engine carry the
// FOREVER fork's commented-out vanilla talent code.
func stripLineComments(src string) string {
	lines := strings.Split(src, "\n")
	for i, l := range lines {
		if j := strings.Index(l, "//"); j >= 0 {
			lines[i] = l[:j]
		}
	}
	return strings.Join(lines, "\n")
}

// modeledGoNames is every talent Go field name the engine's class
// package (and its spec sub-packages) reads outside comments, tests
// and generated code. A talent nothing reads sims as zero whatever
// its tooltip says.
func modeledGoNames(engineDir, class string) (map[string]bool, error) {
	root := filepath.Join(engineDir, "sim", class)
	out := map[string]bool{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || skipEngineFile(d.Name()) {
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		for _, m := range talentRefRE.FindAllStringSubmatch(stripLineComments(string(b)), -1) {
			out[m[1]] = true
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("scanning %s for talent use: %w", root, err)
	}
	return out, nil
}

// modeledTalents is the active-build talent ids whose engine field
// the engine's code reads.
func modeledTalents(t talentTrees, layout enginetalents.Layout, goNames map[string]bool) (map[int]bool, error) {
	out := map[int]bool{}
	for ti, tree := range t.trees {
		for _, n := range tree.Talents {
			f, err := layout.FieldFor(ti, n)
			if err != nil {
				return nil, err
			}
			if goNames[f.GoName()] {
				out[n.ID] = true
			}
		}
	}
	return out, nil
}
