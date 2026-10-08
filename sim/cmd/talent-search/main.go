// Command talent-search sims talent build variants for one spec at one
// level and reports which beat the spec's guide build beyond error.
//
// Gear is the band's committed BiS set, so only talents vary. Candidates
// are viable, not exhaustive: every talent is first probed for the DPS
// the engine credits it with around the guide build; then single swaps
// move 1-5 points out of talents the engine does not credit with damage
// into ones it does, one archetype per deep tree is filled greedily by
// those credits, and the guide is re-spent with its non-damage points
// moved. Every candidate is screened at a modest iteration count, the
// best are re-simmed at the ranker's weights level, and a markdown
// report goes to design/reviews/talent-search/<spec>.md. Nothing under
// data/ or web/ is ever written.
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// errSkipped marks a spec this search does not answer for (a healer, a
// tank, no written rotation); the command exits 3 for it so a driver
// loop can tell a skip from a failure.
var errSkipped = errors.New("skipped")

const exitSkipped = 3

func main() {
	err := run(os.Args[1:])
	switch {
	case errors.Is(err, errSkipped):
		log.Print(err)
		os.Exit(exitSkipped)
	case err != nil:
		log.Fatal(err)
	}
}

// options is one search's parameters.
type options struct {
	repoRoot, spec, faction, preset, build, out, engineSrc, keep string
	level                                                        int
	probeIters, screenIters, finalIters                          int
	top, limit, refineTop                                        int
	seed                                                         int64
}

func parseOptions(args []string) (options, error) {
	var o options
	fs := flag.NewFlagSet("talent-search", flag.ContinueOnError)
	fs.StringVar(&o.repoRoot, "repo-root", ".", "the site repository root")
	fs.StringVar(&o.spec, "spec", "", "spec slug from data/curated/specs.json (required)")
	fs.IntVar(&o.level, "level", 60, "level band (a band data/builds/<build>/bis/<spec>.json carries, e.g. 40 or 60)")
	fs.StringVar(&o.faction, "faction", "alliance", "faction whose BiS band (gear and race) to wear")
	fs.StringVar(&o.preset, "preset", request.BarePreset, "sim preset: bare (gear and class kit) or raid (the Phase 1 raid buffs and consumables on the raid band's gear)")
	fs.StringVar(&o.build, "build", "", "client build; defaults to web/src/data/active-build.json's")
	fs.IntVar(&o.probeIters, "probe-iterations", 200, "iterations for each talent's credit probe")
	fs.IntVar(&o.screenIters, "screen-iterations", 300, "iterations for screening every candidate")
	fs.IntVar(&o.finalIters, "final-iterations", 100*api.WeightsIterationsFactor, "iterations for the finalists (default: the ranker's weights level, 100 per direction x WeightsIterationsFactor)")
	fs.IntVar(&o.top, "top", 10, "finalists re-simmed at -final-iterations")
	fs.IntVar(&o.limit, "cap", 60, "most candidates screened")
	fs.IntVar(&o.refineTop, "refine", 3, "swap refinements kept per structural build")
	fs.Int64Var(&o.seed, "seed", 7, "one seed for every sim (paired comparisons)")
	fs.StringVar(&o.out, "out", "", "report directory; defaults to design/reviews/talent-search under -repo-root")
	fs.StringVar(&o.keep, "keep", "", "comma-separated talent names no candidate may take a point from (a raid search protects what the rotation casts, the raid-wide auras, threat reduction and survivability cooldowns)")
	fs.StringVar(&o.engineSrc, "engine-src", "", "engine module source directory; defaults to go list -m's")
	if err := fs.Parse(args); err != nil {
		return o, err
	}
	if o.spec == "" {
		return o, errors.New("-spec is required")
	}
	if o.level < 10 || o.level > 60 {
		return o, fmt.Errorf("-level %d is outside 10..60", o.level)
	}
	if o.out == "" {
		o.out = filepath.Join(o.repoRoot, "design", "reviews", "talent-search")
	}
	return o, nil
}

func run(args []string) error {
	o, err := parseOptions(args)
	if err != nil {
		return err
	}
	start := time.Now()
	in, err := prepare(o)
	if err != nil {
		return err
	}
	run := engineRun(in.setup)
	credits, err := probeCredits(in.setup.trees, in.guide, run, o.probeIters)
	if err != nil {
		return err
	}
	pool := generate(in.setup.trees, in.guide, credits, in.modeled, in.keep, o.level-9, o.refineTop)
	cands := capCandidates(pool, o.limit)
	log.Printf("talent-search: %s: %d candidates generated, %d screened", o.spec, len(pool), len(cands))
	clean := func(b build) bool { return len(removedUnmodeled(in, credits, b)) == 0 }
	ev, err := evaluate(in.guide, cands, run, o.screenIters, o.finalIters, o.top, clean)
	if err != nil {
		return err
	}
	rep := report{opts: o, in: in, credits: credits, eval: ev, poolSize: len(pool), elapsed: time.Since(start)}
	path := filepath.Join(o.out, fmt.Sprintf("%s.md", reportName(o)))
	if err := os.MkdirAll(o.out, 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(path, []byte(rep.markdown()), 0o644); err != nil {
		return err
	}
	log.Printf("talent-search: %s: wrote %s in %s", o.spec, path, time.Since(start).Round(time.Second))
	return nil
}

// reportName is <spec> at level 60 and <spec>-<level> otherwise.
func reportName(o options) string {
	name := o.spec
	if o.level != 60 {
		name = fmt.Sprintf("%s-%d", name, o.level)
	}
	if o.preset != "" && o.preset != request.BarePreset {
		name += "-" + o.preset
	}
	return name
}

// inputs is everything a search reads before its first sim.
type inputs struct {
	setup       simSetup
	clientBuild string
	guide       build
	guideCode   string // the guide frontmatter's own trees and client build
	guideIssue  string // why the guide is not a legal active-build build, if it is not
	modeled     map[int]bool
	keep        map[int]bool // talents no candidate may take a point from (-keep)
}

func prepare(o options) (inputs, error) {
	spec, err := loadSpec(o.repoRoot, o.spec)
	if err != nil {
		return inputs{}, err
	}
	if spec.Role != dpsRole {
		return inputs{}, fmt.Errorf("%w: %s is a %s spec; DPS is not its objective", errSkipped, o.spec, spec.Role)
	}
	if state, err := aplState(o.repoRoot, o.spec); err != nil {
		return inputs{}, err
	} else if state != writtenAPL {
		return inputs{}, fmt.Errorf("%w: %s has no written rotation (apl state %q)", errSkipped, o.spec, state)
	}
	clientBuild := o.build
	if clientBuild == "" {
		if clientBuild, err = leveling.ReadActiveBuild(o.repoRoot); err != nil {
			return inputs{}, err
		}
	}
	activeTrees, err := leveling.LoadTalentTrees(o.repoRoot, clientBuild, spec.ClassSlug)
	if err != nil {
		return inputs{}, err
	}
	trees := newTalentTrees(activeTrees)
	engineDir := o.engineSrc
	if engineDir == "" {
		if engineDir, err = enginetalents.SourceDir(filepath.Join(o.repoRoot, "sim")); err != nil {
			return inputs{}, err
		}
	}
	layout, err := enginetalents.ForClass(engineDir, spec.ClassSlug)
	if err != nil {
		return inputs{}, err
	}
	if _, err := layout.Encode(activeTrees, build{}); err != nil {
		return inputs{}, fmt.Errorf("the engine cannot read %s's %s trees: %w", spec.ClassSlug, clientBuild, err)
	}
	guide, guideCode, err := guideBuild(o, spec, activeTrees, trees)
	if err != nil {
		return inputs{}, err
	}
	in := inputs{clientBuild: clientBuild, guide: guide, guideCode: guideCode}
	if err := trees.legal(guide, o.level-9); err != nil {
		in.guideIssue = err.Error()
	}
	band, err := loadBISBand(filepath.Join(o.repoRoot, "data", "builds", clientBuild), o.spec, o.level, o.faction, o.preset)
	if err != nil {
		return inputs{}, err
	}
	applied, err := request.ResolveFromFile(
		filepath.Join(o.repoRoot, "data", "curated", "presets.json"), o.preset,
		specs.Spec{Spec: spec.Spec, ClassSlug: spec.ClassSlug, ReferenceStat: spec.ReferenceStat})
	if err != nil {
		return inputs{}, err
	}
	in.setup = simSetup{preset: applied, spec: spec, band: band, level: o.level, seed: o.seed, layout: layout, trees: trees}
	names, err := modeledGoNames(engineDir, spec.ClassSlug)
	if err != nil {
		return inputs{}, err
	}
	if in.modeled, err = modeledTalents(trees, layout, names); err != nil {
		return inputs{}, err
	}
	in.keep, err = resolveKeep(trees, splitKeep(o.keep))
	return in, err
}

// guideBuild is the guide's level-60 build (its raid build under the raid preset) read by stable talent id
// (leveling.GuideTalentTargets) and truncated to this level the way
// the ranker and the ladder truncate it (leveling.LadderTalentString),
// so a guide still exported from an older client build lands on the
// right talents of the active one.
func guideBuild(o options, spec specInfo, activeTrees []leveling.TalentTree, trees talentTrees) (build, string, error) {
	readBuild := leveling.GuideBuildTalents
	if o.preset == request.RaidPreset {
		// A raid search starts from the build the raid-ready entry is
		// simmed on (the guide's raidBuild, else its leveling build).
		readBuild = leveling.GuideRaidBuildTalents
	}
	guideClient, digits, err := readBuild(o.repoRoot, spec.ClassSlug, spec.SpecSlug)
	if err != nil {
		return nil, "", err
	}
	guideTrees, err := leveling.LoadTalentTrees(o.repoRoot, guideClient, spec.ClassSlug)
	if err != nil {
		return nil, "", err
	}
	targets := leveling.GuideTalentTargets(guideTrees, digits)
	s := leveling.LadderTalentString(activeTrees, targets, spec.TreeIndex, o.level)
	b, err := trees.decodeActive(s)
	if err != nil {
		return nil, "", err
	}
	code := fmt.Sprintf("%s/%s/%s (client %s)", digits[0], digits[1], digits[2], guideClient)
	return b, code, nil
}
