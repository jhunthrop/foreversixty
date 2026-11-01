// Command addon-builds writes every spec's default talent builds for the
// in-game addon: the guide's leveling and raid builds, each as the order
// a character learns its 51 points in (leveling.LadderPointOrder), to
// data/builds/<active build>/addon-builds.json. The data pipeline reads
// that file into Data.lua's defaultBuilds (docs/contracts/addon-v2.md,
// contract 1).
//
//	go run ./sim/cmd/addon-builds -repo-root .
//
// Every spec in sim/specs must have a guide whose builds are stamped with
// the active build and have a legal 51-point order; the command fails
// listing every build that does not, and writes nothing.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

// build is one default build as addon-builds.json carries it: order is
// one [tab, tier, column] triple per point, 1-based.
type build struct {
	Name  string   `json:"name"`
	Order [][3]int `json:"order"`
}

type guideReader func(repoRoot, class, specSlug string) (string, [3]string, error)

func run(args []string) error {
	fs := flag.NewFlagSet("addon-builds", flag.ContinueOnError)
	repoRoot := fs.String("repo-root", ".", "the site repository root")
	if err := fs.Parse(args); err != nil {
		return err
	}
	active, b, err := generate(*repoRoot)
	if err != nil {
		return err
	}
	path := filepath.Join(*repoRoot, "data", "builds", active, "addon-builds.json")
	if err := os.WriteFile(path, b, 0o644); err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}

// generate orders every spec's guide builds for the active client build
// and renders addon-builds.json's bytes.
func generate(repoRoot string) (string, []byte, error) {
	active, err := leveling.ReadActiveBuild(repoRoot)
	if err != nil {
		return "", nil, err
	}
	out := make(map[string]map[string]build, len(specs.All))
	var failed []string
	for _, spec := range specs.All {
		levelingBuild, levelingErr := specBuild(repoRoot, active, spec, leveling.GuideBuildTalents, spec.Name+" leveling")
		if levelingErr != nil {
			failed = append(failed, fmt.Sprintf("%s leveling: %v", spec.Spec, levelingErr))
		}
		raidBuild, raidErr := specBuild(repoRoot, active, spec, leveling.GuideRaidBuildTalents, spec.Name+" raid")
		if raidErr != nil {
			failed = append(failed, fmt.Sprintf("%s raid: %v", spec.Spec, raidErr))
		}
		if levelingErr == nil && raidErr == nil {
			out[spec.Spec] = map[string]build{"leveling": levelingBuild, "raid": raidBuild}
		}
	}
	if len(failed) > 0 {
		return "", nil, fmt.Errorf("addon-builds: %d guide build(s) have no usable point order for %s:\n  %s", len(failed), active, strings.Join(failed, "\n  "))
	}
	b, err := render(out)
	return active, b, err
}

// specBuild reads one of spec's guide builds against the active build's
// trees and orders its points.
func specBuild(repoRoot, active string, spec specs.Spec, read guideReader, name string) (build, error) {
	stamp, digits, err := read(repoRoot, spec.ClassSlug, spec.SpecSlug)
	if err != nil {
		return build{}, err
	}
	if err := leveling.RequireGuideBuildMatchesActive(stamp, active); err != nil {
		return build{}, err
	}
	trees, err := leveling.LoadTalentTrees(repoRoot, active, spec.ClassSlug)
	if err != nil {
		return build{}, err
	}
	points, err := leveling.LadderPointOrder(trees, leveling.GuideTalentTargets(trees, digits), spec.TreeIndex)
	if err != nil {
		return build{}, err
	}
	order := make([][3]int, len(points))
	for i, p := range points {
		order[i] = [3]int{p.Tab, p.Tier, p.Column}
	}
	return build{Name: name, Order: order}, nil
}

// render writes one spec per block and one build per line, specs in
// sim/specs' own order, so a regenerated file diffs by build.
func render(out map[string]map[string]build) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{\n")
	for i, spec := range specs.All {
		key, _ := json.Marshal(spec.Spec)
		fmt.Fprintf(&buf, "  %s: {\n", key)
		for j, kind := range []string{"leveling", "raid"} {
			b, err := json.Marshal(out[spec.Spec][kind])
			if err != nil {
				return nil, err
			}
			sep := ","
			if j == 1 {
				sep = ""
			}
			fmt.Fprintf(&buf, "    %q: %s%s\n", kind, b, sep)
		}
		sep := ","
		if i == len(specs.All)-1 {
			sep = ""
		}
		fmt.Fprintf(&buf, "  }%s\n", sep)
	}
	buf.WriteString("}\n")
	return buf.Bytes(), nil
}
