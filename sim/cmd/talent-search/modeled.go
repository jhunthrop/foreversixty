package main

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
)

// talentRefRE is how engine code reads a talent field, in two shapes.
// Chained: any selector chain that ends in Talents and then names the
// field - druid.Talents.Furor, hunter.Talents.GetSurefooted(),
// c.Character.Talents.Subtlety. Aliased: an identifier holding a Talents
// proto, such as mage's own applyDeclarativeTalents, which does
// `t := mage.Talents` and then reads `t.FirePower`.
//
// The two are separate patterns because one Go regexp cannot see both: the
// aliased pattern consumes "druid.Talents" in `druid.Talents.Furor` and
// leaves no receiver for ".Furor", which is how every chained read went
// uncounted and the report called Furor and Natural Shapeshifter unmodeled
// while the cat rotation reads them. modeledTalents only keeps the names
// that are actually a talent's Go name, so an unrelated "x.FirePower" on
// some other struct would have to collide with a real talent's generated
// name to misfire.
var talentRefREs = []*regexp.Regexp{
	regexp.MustCompile(`\bTalents\.(?:Get)?([A-Z][A-Za-z0-9]*)\b`),
	regexp.MustCompile(`\b\w+\.(?:Get)?([A-Z][A-Za-z0-9]*)\b`),
}

// ignoredByGoTool is a file or directory the go tool never compiles: a
// name starting with "_" or ".". The fork parks retired druid code in
// _maul.go and _tank/, which still mention talents.
func ignoredByGoTool(name string) bool {
	return strings.HasPrefix(name, "_") || strings.HasPrefix(name, ".")
}

// skipEngineFile is generated, test or uncompiled code, which reads
// talents without modeling them.
func skipEngineFile(name string) bool {
	return ignoredByGoTool(name) || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") ||
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

// blankAssignRE is a statement that only discards its values: `_ = x` or
// `_, _ = x, y`. The fork's engine writes `_ = druid.Talents.Subtlety`
// beside a comment saying why the talent changes nothing here; that read
// acknowledges the field for the compiler and models nothing, so it must
// not count as the engine reading the talent.
var blankAssignRE = regexp.MustCompile(`(?m)^\s*_(?:\s*,\s*_)*\s*=.*$`)

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
		if d.IsDir() {
			if path != root && ignoredByGoTool(d.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if skipEngineFile(d.Name()) {
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		src := blankAssignRE.ReplaceAllString(stripLineComments(string(b)), "")
		for _, re := range talentRefREs {
			for _, m := range re.FindAllStringSubmatch(src, -1) {
				out[m[1]] = true
			}
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

// talentClass is the final, three-way classification a report shows
// for one talent: how much the engine's own numbers back its tooltip.
type talentClass int

const (
	// classUnmodeled is a talent the static scan found no code reading
	// AND whose probe moved DPS no further than the combined error:
	// the engine gives it nothing, as far as either signal can tell.
	classUnmodeled talentClass = iota
	// classModeledNoDamage is a talent the static scan found code
	// reading, but whose probe did not move DPS beyond the combined
	// error (a utility talent, most often).
	classModeledNoDamage
	// classDamage is a talent whose probe moved DPS beyond the
	// combined error, whatever the static scan found. The probe
	// outranks the scan: the scan can always miss a field read through
	// an aliased receiver, but a probe that moves DPS beyond error is
	// the engine crediting the talent, full stop.
	classDamage
)

func (c talentClass) String() string {
	switch c {
	case classDamage:
		return "damage"
	case classModeledNoDamage:
		return "modeled, no damage"
	default:
		return "unmodeled"
	}
}

// classify applies the precedence rule between the static scan and
// the probe: a probe beyond the combined error is damage whatever the
// static scan said; short of that, the static scan finding some code
// reading the field is at worst "modeled, no damage"; "unmodeled" only
// when the scan finds nothing and the probe is within error.
func classify(staticModeled bool, c credit) talentClass {
	switch {
	case math.Abs(c.Diff) > c.Err:
		return classDamage
	case staticModeled:
		return classModeledNoDamage
	default:
		return classUnmodeled
	}
}
