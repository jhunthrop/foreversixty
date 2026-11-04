package leveling

import (
	"path/filepath"
	"testing"
)

// Sim-accuracy round 15 (shaman): Enhancement's leveling build spent three
// points on Elemental Alacrity, which shortens Lightning Bolt and Chain
// Lightning casts; the rotation casts Lightning Bolt only at five Maelstrom
// Weapon stacks, when it is already instant, and never Chain Lightning, so
// the points did nothing (bare 60 identical with them taken out). They go
// to Elemental Devastation (melee crit after a spell crit): +1.1 to +1.2%
// at bare 60, both factions, seeds 1 and 50001.
func TestEnhancementLevelingBuildTakesNoElementalAlacrity(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "shaman")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, "shaman", "enhancement")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	ranks := map[string]int{}
	for _, tree := range trees {
		for _, talent := range tree.Talents {
			ranks[talent.Name] = targets[talent.ID]
		}
	}
	if ranks["Elemental Alacrity"] != 0 {
		t.Errorf("the Enhancement build takes Elemental Alacrity %d", ranks["Elemental Alacrity"])
	}
	if ranks["Elemental Devastation"] != 3 {
		t.Errorf("the Enhancement build takes Elemental Devastation %d, want 3", ranks["Elemental Devastation"])
	}
}
