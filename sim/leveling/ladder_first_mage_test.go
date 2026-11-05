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
		// Sim-accuracy round 23: Fire's level 60 build takes Accelerant 5, Missile Barrage
		// and Arcane Blast; its ladder keeps Ignite, Pyroblast, Heating Up at 30 and
		// Combustion at 40, spends Accelerant 2 to 4 at 22 to 24 in place of Flame Throwing
		// and Burning Soul. Sim-accuracy round 24 (platform): Wand Specialization 2 comes at
		// 41 and 42, right after Combustion (its ladder-order row), and Accelerant 5 and
		// Burning Soul 2 at 43 to 45; 45 is the same build either way.
		{"mage-fire", "fire", 1, 20, "000000000000000000-51400001000000000-000000000000000000"},
		{"mage-fire", "fire", 1, 24, "000000000000000000-54500001000000000-000000000000000000"},
		{"mage-fire", "fire", 1, 30, "000000000000000000-54500001230100000-000000000000000000"},
		{"mage-fire", "fire", 1, 40, "000000000000000000-54500001230133031-000000000000000000"},
		{"mage-fire", "fire", 1, 41, "100000000000000000-54500001230133031-000000000000000000"},
		{"mage-fire", "fire", 1, 42, "200000000000000000-54500001230133031-000000000000000000"},
		{"mage-fire", "fire", 1, 43, "200000000000000000-55500001230133031-000000000000000000"},
		{"mage-fire", "fire", 1, 45, "200000000000000000-55500201230133031-000000000000000000"},
		{"mage-fire", "fire", 1, 60, "230105200100000000-55500201230133031-003000000000000000"},
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

// Sim-accuracy round 24 (platform): a reservation is spent in its own order across trees. An
// off-tree first-row rank listed after own-tree ranks waits for them (it was always spent
// before the own tree, so Fire's Wand Specialization could only come first or last); one
// listed first is still taken first (Shadow's Wand Specialization).
func TestReservationSpendsAnOffTreeRankInListOrder(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "mage")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, "mage", "fire")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	after := withFirst([]LadderFirstRank{{Talent: "Incineration", Rank: 5}, {Talent: "Wand Specialization", Rank: 1}})
	before := withFirst([]LadderFirstRank{{Talent: "Wand Specialization", Rank: 1}, {Talent: "Incineration", Rank: 5}})
	cases := []struct {
		name  string
		opt   LadderOption
		level int
		want  string
	}{
		{"listed after, level 14", after, 14, "000000000000000000-50000000000000000-000000000000000000"},
		{"listed after, level 15", after, 15, "100000000000000000-50000000000000000-000000000000000000"},
		{"listed first, level 10", before, 10, "100000000000000000-00000000000000000-000000000000000000"},
		{"listed first, level 15", before, 15, "100000000000000000-50000000000000000-000000000000000000"},
	}
	for _, c := range cases {
		if got := LadderTalentString(trees, targets, 1, c.level, c.opt); got != c.want {
			t.Errorf("%s: %s, want %s", c.name, got, c.want)
		}
	}
}
