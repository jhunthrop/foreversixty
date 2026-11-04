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
	_, digits, err := GuideRaidBuildTalents(filepath.Join("..", ".."), "warlock", "demonology")
	if err != nil {
		t.Fatal(err)
	}
	checkWarlockRanks(t, "the Demonology raid build", digits,
		map[string]int{"Improved Voidwalker": 0, "Improved Imp": 0, "Improved Sayaad": 2})
}

// From 40 the Affliction leveling rotation channels Wrack as its filler and never casts
// Shadow Bolt, so Nightfall's Shadow Trance (an instant Shadow Bolt) is never spent, and
// spending it in place of Wrack loses at every band (round 15 A4; round 17 at bare 40, 50
// and 60). The two points take Malevolence 2 in the same tier: +1.4% to +1.5% at 40, 45
// and 60, and level or better from 35 to 39, where Shadow Bolt is still the filler
// (sim-accuracy round 17, seeds 1 and 50001).
func TestAfflictionLevelingBuildTakesMalevolenceOverNightfall(t *testing.T) {
	_, digits, err := GuideBuildTalents(filepath.Join("..", ".."), "warlock", "affliction")
	if err != nil {
		t.Fatal(err)
	}
	checkWarlockRanks(t, "the Affliction leveling build", digits,
		map[string]int{"Nightfall": 0, "Malevolence": 2})
}

// checkWarlockRanks fails unless the guide build's digits put exactly the wanted rank in
// each named warlock talent.
func checkWarlockRanks(t *testing.T, build string, digits [3]string, want map[string]int) {
	t.Helper()
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "warlock")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	seen := 0
	for _, tree := range trees {
		for _, node := range tree.Talents {
			if w, ok := want[node.Name]; ok {
				seen++
				if got := targets[node.ID]; got != w {
					t.Errorf("%s has %s %d, want %d", build, node.Name, got, w)
				}
			}
		}
	}
	if seen != len(want) {
		t.Fatalf("found %d of %d talents in the warlock trees", seen, len(want))
	}
}
