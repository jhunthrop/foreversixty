package main

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// errSkipped marks a spec this search does not answer for (a healer,
// a tank, no written rotation); main exits 3 for it, mirroring
// sim/cmd/talent-search, so a driver loop can tell a skip from a
// failure.
var errSkipped = errors.New("skipped")

const exitSkipped = 3

// inputs is everything a search reads before its first sim.
type inputs struct {
	setup       simSetup
	clientBuild string
	base        rotation // the curated rotation this search starts from
	candidates  []learnedCandidate
	tickSeconds map[int]float64
}

func prepare(o options) (inputs, error) {
	spec, err := loadSpec(o.repoRoot, o.spec)
	if err != nil {
		return inputs{}, err
	}
	if spec.Role != dpsRole {
		return inputs{}, fmt.Errorf("%w: %s is a %s spec; DPS is not its objective", errSkipped, o.spec, spec.Role)
	}
	curated, err := loadCuratedAPL(o.repoRoot, o.spec)
	if err != nil {
		return inputs{}, err
	}
	if curated.State != writtenAPL {
		return inputs{}, fmt.Errorf("%w: %s has no written rotation (apl state %q)", errSkipped, o.spec, curated.State)
	}
	base, err := parseRotation(curated.Rotation)
	if err != nil {
		return inputs{}, err
	}

	clientBuild := o.build
	if clientBuild == "" {
		if clientBuild, err = leveling.ReadActiveBuild(o.repoRoot); err != nil {
			return inputs{}, err
		}
	}
	engineTalents, err := guideEngineTalents(o, spec, clientBuild)
	if err != nil {
		return inputs{}, err
	}
	band, err := loadBISBand(filepath.Join(o.repoRoot, "data", "builds", clientBuild), o.spec, o.level, o.faction)
	if err != nil {
		return inputs{}, err
	}

	authored := authoredCastIDs(base)
	candidates, err := learnedCandidates(o.repoRoot, clientBuild, spec.ClassSlug, o.level, authored)
	if err != nil {
		return inputs{}, err
	}
	ticks, err := tickSecondsByID(o.repoRoot, clientBuild, spec.ClassSlug, o.level)
	if err != nil {
		return inputs{}, err
	}

	return inputs{
		setup: simSetup{
			spec:          spec,
			band:          band,
			level:         o.level,
			seed:          o.seed,
			engineTalents: engineTalents,
		},
		clientBuild: clientBuild,
		base:        base,
		candidates:  candidates,
		tickSeconds: ticks,
	}, nil
}

// guideEngineTalents is the guide's own FS1 build, read by stable
// talent id (leveling.GuideTalentTargets), truncated to o.level the
// way the ranker and the ladder truncate it
// (leveling.LadderTalentString), then repositioned onto the engine's
// own field order (enginetalents.Layout.Reposition) - the exact two-
// step conversion sim/request/ladder_test.go's own ladder run uses, so
// this search's character carries precisely the talents the guide (and
// the ladder) credits it with, with no talent-search-style variation.
func guideEngineTalents(o options, spec specInfo, clientBuild string) (string, error) {
	activeTrees, err := leveling.LoadTalentTrees(o.repoRoot, clientBuild, spec.ClassSlug)
	if err != nil {
		return "", err
	}
	guideClient, digits, err := leveling.GuideBuildTalents(o.repoRoot, spec.ClassSlug, spec.SpecSlug)
	if err != nil {
		return "", err
	}
	guideTrees, err := leveling.LoadTalentTrees(o.repoRoot, guideClient, spec.ClassSlug)
	if err != nil {
		return "", err
	}
	targets := leveling.GuideTalentTargets(guideTrees, digits)
	talents := leveling.LadderTalentString(activeTrees, targets, spec.TreeIndex, o.level)

	engineDir := o.engineSrc
	if engineDir == "" {
		if engineDir, err = enginetalents.SourceDir(filepath.Join(o.repoRoot, "sim")); err != nil {
			return "", err
		}
	}
	layout, err := enginetalents.ForClass(engineDir, spec.ClassSlug)
	if err != nil {
		return "", err
	}
	return layout.Reposition(activeTrees, talents)
}
