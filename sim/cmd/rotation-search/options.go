package main

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/sim/request"
)

// options is one search's parameters, named and defaulted to match
// sim/cmd/talent-search's own flag set wherever the two commands share
// a concern.
type options struct {
	repoRoot, spec, faction, preset, build, buildCode, out, engineSrc string
	level                                                             int
	iterations                                                        int // screening and probe iterations
	confirmIterations                                                 int
	rounds                                                            int
	seed                                                              int64
	probeOnly                                                         bool
}

func parseOptions(args []string) (options, error) {
	var o options
	fs := flag.NewFlagSet("rotation-search", flag.ContinueOnError)
	fs.StringVar(&o.repoRoot, "repo-root", ".", "the site repository root")
	fs.StringVar(&o.spec, "spec", "", "spec slug from data/curated/specs.json (required)")
	fs.IntVar(&o.level, "level", 60, "level band; this command only supports 60 today")
	fs.StringVar(&o.faction, "faction", "alliance", "faction whose BiS band (gear and race) to wear")
	fs.StringVar(&o.preset, "preset", request.BarePreset, "sim preset: bare (gear and class kit) or raid (the Phase 1 raid buffs and consumables on the raid band's gear)")
	fs.StringVar(&o.buildCode, "build-code", "", "FS1 build code to wear instead of the guide's (FS1:<build>:<class>:<race>:<t1>/<t2>/<t3>:)")
	fs.StringVar(&o.build, "build", "", "client build; defaults to web/src/data/active-build.json's")
	fs.IntVar(&o.iterations, "iterations", 300, "iterations for screening every mutation and every probe")
	fs.IntVar(&o.confirmIterations, "confirm", 800, "iterations for confirming the final rotation")
	fs.IntVar(&o.rounds, "rounds", 6, "greedy hill-climb rounds, at most")
	fs.Int64Var(&o.seed, "seed", 7, "one seed for every sim (paired comparisons)")
	fs.BoolVar(&o.probeOnly, "probe-only", false, "write only the baseline and the action-probe table (steps 1-2); skip the mutation search")
	fs.StringVar(&o.out, "out", "", "report directory; defaults to design/reviews/rotation-search under -repo-root")
	fs.StringVar(&o.engineSrc, "engine-src", "", "engine module source directory; defaults to go list -m's")
	if err := fs.Parse(args); err != nil {
		return o, err
	}
	if o.spec == "" {
		return o, errors.New("-spec is required")
	}
	if o.level != 60 {
		return o, fmt.Errorf("-level %d: rotation-search only supports level 60 - a curated rotation's ranks are authored for max level and this tool does not rewrite them down", o.level)
	}
	if o.out == "" {
		o.out = filepath.Join(o.repoRoot, "design", "reviews", "rotation-search")
	}
	return o, nil
}
