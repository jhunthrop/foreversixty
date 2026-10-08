package main

import (
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// passTargets is a spec's two level-60 talent targets, each addressed by
// stable talent id (leveling.GuideTalentTargets): the guide's leveling
// build, which every leveling band and the bare band 60 spend from, and
// its raid build, which only the raid preset's band 60 spends from. A
// guide with no separate raid build carries the leveling targets twice.
type passTargets struct {
	leveling map[int]int
	raid     map[int]int
}

// forPass picks the targets a ranking pass spends from: the raid
// preset's pass wears the raid build, every other pass the leveling one.
func (t passTargets) forPass(pass string) map[int]int {
	if pass == presetRaid {
		return t.raid
	}
	return t.leveling
}

// guideBuildReader is leveling.GuideBuildTalents' shape, which the
// leveling and raid readers both satisfy.
type guideBuildReader func(repoRoot, class, specSlug string) (string, [3]string, error)

// loadPassTargets reads both of a spec's guide builds against the active
// build's trees. Each build's stamp must name the active build (the
// positional digit read is only right for the tree shape it was authored
// against; leveling.RequireGuideBuildMatchesActive's own doc).
func loadPassTargets(repoRoot, activeBuild string, spec specInfo) (passTargets, error) {
	levelingTargets, err := guideTargets(repoRoot, activeBuild, spec, leveling.GuideBuildTalents)
	if err != nil {
		return passTargets{}, err
	}
	raidTargets, err := guideTargets(repoRoot, activeBuild, spec, leveling.GuideRaidBuildTalents)
	if err != nil {
		return passTargets{}, err
	}
	return passTargets{leveling: levelingTargets, raid: raidTargets}, nil
}

func guideTargets(repoRoot, activeBuild string, spec specInfo, read guideBuildReader) (map[int]int, error) {
	build, digits, err := read(repoRoot, spec.ClassSlug, spec.SpecSlug)
	if err != nil {
		return nil, err
	}
	if err := leveling.RequireGuideBuildMatchesActive(build, activeBuild); err != nil {
		return nil, err
	}
	trees, err := leveling.LoadTalentTrees(repoRoot, build, spec.ClassSlug)
	if err != nil {
		return nil, err
	}
	return leveling.GuideTalentTargets(trees, digits), nil
}
