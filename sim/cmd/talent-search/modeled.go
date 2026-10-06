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

// talentRefRE is how engine code reads a talent field: any identifier
// holding (or aliasing) a Talents proto, followed by a dot and the
// field's Go name - paladin.Talents.Vengeance,
// hunter.Talents.GetSurefooted(), or an alias such as mage's own
// applyDeclarativeTalents, which does `t := mage.Talents` and then
// reads `t.FirePower`. A scan anchored on the literal receiver name
// "Talents" misses every field read through such an alias, so this
// matches any identifier before the dot, not just "Talents" itself;
// modeledTalents only keeps the names that are actually a talent's
// Go name, so an unrelated "x.FirePower" on some other struct would
// have to collide with a real talent's generated name to misfire.
var talentRefRE = regexp.MustCompile(`\b\w+\.(?:Get)?([A-Z][A-Za-z0-9]*)\b`)

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
