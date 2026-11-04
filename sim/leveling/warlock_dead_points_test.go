package leveling

import (
	"path/filepath"
	"testing"
)

// The Demonology raid build plays with the Succubus out and sacrifices the Imp before the
// pull (Demonic Pact keeps the sacrifice), so it never summons the Voidwalker and never
// casts a Firebolt. Improved Voidwalker ("Torment, Consume Shadows, Sacrifice, and
// Suffering") and Improved Imp ("Firebolt ... Fire Shield") do nothing for it; the point
// goes to Improved Sayaad, 10% more Lash of Pain (sim-accuracy round 16).
func TestDemonologyRaidBuildSpendsNothingOnDeadDemons(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "warlock")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideRaidBuildTalents(repoRoot, "warlock", "demonology")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	want := map[string]int{"Improved Voidwalker": 0, "Improved Imp": 0, "Improved Sayaad": 2}
	seen := 0
	for _, tree := range trees {
		for _, node := range tree.Talents {
			if w, ok := want[node.Name]; ok {
				seen++
				if got := targets[node.ID]; got != w {
					t.Errorf("the Demonology raid build has %s %d, want %d", node.Name, got, w)
				}
			}
		}
	}
	if seen != len(want) {
		t.Fatalf("found %d of %d talents in the warlock trees", seen, len(want))
	}
}
