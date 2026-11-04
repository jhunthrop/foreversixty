package leveling

import (
	"path/filepath"
	"testing"
)

// Sim-accuracy round 15 (mage wave 2, from guide-accuracy round 1): Frost takes Improved
// Frostbolt before Frost Warding, so its level-10 point speeds the Frostbolt it casts
// (+2.0 to +2.9% at band 10); Arcane takes Arcane Blast at 20, in place of the armor point in
// Arcane Resilience (+0.2% at 20, +17 to +21% at 25, where the plain walk had no Blast
// until 27). Every other checkpoint is unchanged
// (design/reviews/2026-11-04-sim-audit-round15/mage.md, Wave 2).
func TestMageLadderTakesItsCastTalentsEarly(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		spec, slug  string
		tree, level int
		want        string
	}{
		{"mage-frost", "frost", 2, 10, "000000000000000000-00000000000000000-010000000000000000"},
		{"mage-frost", "frost", 2, 20, "000000000000000000-00000000000000000-253100000000000000"},
		{"mage-arcane", "arcane", 0, 20, "140005000100000000-00000000000000000-000000000000000000"},
		{"mage-arcane", "arcane", 0, 25, "141005211100000000-00000000000000000-000000000000000000"},
		{"mage-arcane", "arcane", 0, 30, "141005213100300000-00000000000000000-000000000000000000"},
	}
	trees, err := LoadTalentTrees(repoRoot, active, "mage")
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		_, digits, err := GuideBuildTalents(repoRoot, "mage", c.slug)
		if err != nil {
			t.Fatal(err)
		}
		targets := GuideTalentTargets(trees, digits)
		if got := LadderTalentString(trees, targets, c.tree, c.level, ForSpec(c.spec)); got != c.want {
			t.Errorf("%s level %d: %s, want %s", c.spec, c.level, got, c.want)
		}
	}
}
