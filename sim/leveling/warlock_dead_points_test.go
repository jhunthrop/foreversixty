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

// The Affliction leveling build fills Malevolence and takes Nightfall once Wrack is its filler:
// the card casts the Shadow Trance bolt between two Wracks (the Nightfall line is gated on
// knowing the talent), and the talent search at bare 60 credits Nightfall and the last three
// Malevolence points with damage where Improved Drains and the Demonology toll points carry
// none (sim-accuracy 2026-11-12, warlocks pass 3; round 17 had held Nightfall at 0 while the
// card never cast the bolt).
func TestAfflictionLevelingBuildTakesMalevolenceAndNightfall(t *testing.T) {
	_, digits, err := GuideBuildTalents(filepath.Join("..", ".."), "warlock", "affliction")
	if err != nil {
		t.Fatal(err)
	}
	checkWarlockRanks(t, "the Affliction leveling build", digits,
		map[string]int{"Nightfall": 2, "Malevolence": 5, "Improved Drains": 0})
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
