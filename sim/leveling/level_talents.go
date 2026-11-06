package leveling

import (
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

// GuideLevelTalents is the talent string a player of spec holds at level on
// the guide's published order: the guide build (its raid build when raid)
// read by stable talent id and climbed to level exactly as the ladder, the
// ranker and the addon's leveling order climb it (GuideTalentTargets,
// LadderTalentString), in activeBuild's own (tier, column) layout, the one a
// published BiS entry's talents use.
func GuideLevelTalents(repoRoot, activeBuild string, spec specs.Spec, level int, raid bool) (string, error) {
	read := GuideBuildTalents
	if raid {
		read = GuideRaidBuildTalents
	}
	guideBuild, digits, err := read(repoRoot, spec.ClassSlug, spec.SpecSlug)
	if err != nil {
		return "", fmt.Errorf("%s guide build: %w", spec.Spec, err)
	}
	guideTrees, err := LoadTalentTrees(repoRoot, guideBuild, spec.ClassSlug)
	if err != nil {
		return "", fmt.Errorf("%s guide talent trees: %w", spec.Spec, err)
	}
	activeTrees, err := LoadTalentTrees(repoRoot, activeBuild, spec.ClassSlug)
	if err != nil {
		return "", fmt.Errorf("%s active talent trees: %w", spec.Spec, err)
	}
	targets := GuideTalentTargets(guideTrees, digits)
	return LadderTalentString(activeTrees, targets, spec.TreeIndex, level, ForSpec(spec.Spec)), nil
}
