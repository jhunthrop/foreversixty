package leveling

import (
	"path/filepath"
	"testing"
)

// Sim-accuracy round 13 (walk lane): the band builds the plain top-row-down walk made left
// out a talent the rotation casts, though the band's points could legally hold it, and
// taking it gained beyond error and at least 1% at seeds 1 and 50001 on the band's own
// fight. Each want is the build the lane measured
// (design/reviews/2026-11-03-sim-audit-round13/walk.md).
func TestBandBuildsTakeTheCastTalentsThatFit(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		spec, class, slug string
		tree, level       int
		talent, want      string
	}{
		// Sim-accuracy round 28 (hybrids): Predator's Edge 5 before Deterrence
		// (TestSurvivalTakesPredatorsEdgeBeforeDeterrence).
		{"hunter-survival", "hunter", "survival", 2, 30, "Strider Kick", "0000000000000000-0000000000000000-500230230050000100"},
		{"mage-fire", "mage", "fire", 1, 20, "Pyroblast", "000000000000000000-51400001000000000-000000000000000000"},
		{"rogue-assassination", "rogue", "assassination", 0, 40, "Mutilate", "32530200251511001-0000000000000000000-000000000000000000"},
		// Sim-accuracy round 26 (melee): Initiative 3 now comes before Camouflage's last
		// two points and Elusiveness (TestSubtletyTakesInitiativeBeforeTheTollPoints).
		{"rogue-subtlety", "rogue", "subtlety", 2, 30, "Ghostly Strike and Hemorrhage", "00000000000000000-0000000000000000000-035220031301000100"},
		// Sim-accuracy round 29 (walks-heal): Restorative Totems and Healing Way before Healing
		// Focus and Tidal Mastery (TestHealerWalksTakeTheirLiveTalentsFirst).
		{"shaman-restoration", "shaman", "restoration", 2, 30, "Mana Tide Totem", "0000000000000000-000000000000000000-5530001010411000"},
		// Sim-accuracy round 15 (shaman): Stormstrike sits in the 15-point row and beta
		// logs show Enhancement shamans of level 25 to 35 casting it; band 30 left it out.
		{"shaman-enhancement", "shaman", "enhancement", 1, 30, "Stormstrike", "0000000000000000-055030310031000000-0000000000000000"},
		// Sim-accuracy round 29 (walks-dps): Amplify Curse, Pandemic and Malevolence before
		// Soul Harvest and Improved Drains (TestAfflictionTakesPandemicBeforeTheDrainTalents).
		{"warlock-affliction", "warlock", "affliction", 0, 30, "Siphon Life", "25520000132001000-0000000000000000000-0000000000000000"},
		{"warlock-demonology", "warlock", "demonology", 1, 30, "Soul Link", "00000000000000000-2035113101210001000-0000000000000000"},
		{"warrior-arms", "warrior", "arms", 0, 40, "Mortal Strike", "12323213132214001-000000000000000000-000000000000000000"},
		// Sim-accuracy round 15 (warrior): the guide build dropped Lingering Rage and Piercing
		// Howl for Improved Berserker Rage, so band 30 spends those two points in Dual Wield
		// Specialization (5, not 3) and still takes Death Wish (+3.5% at seeds 1 and 50001).
		{"warrior-fury", "warrior", "fury", 1, 30, "Death Wish", "00000000000000000-050530025000100000-000000000000000000"},
	}
	for _, c := range cases {
		trees, err := LoadTalentTrees(repoRoot, active, c.class)
		if err != nil {
			t.Fatal(err)
		}
		_, digits, err := GuideBuildTalents(repoRoot, c.class, c.slug)
		if err != nil {
			t.Fatal(err)
		}
		targets := GuideTalentTargets(trees, digits)
		if got := LadderTalentString(trees, targets, c.tree, c.level, ForSpec(c.spec)); got != c.want {
			t.Errorf("%s level %d: %s, want %s (with %s)", c.spec, c.level, got, c.want, c.talent)
		}
	}
}

// Demonic Sacrifice is Soul Link's prerequisite, but the rotation never casts it without
// Demonic Pact, so the reservation never pulls it ahead of the walk: below 25 no band takes
// it (round 13). Since round 28 Unholy Power and Fel Vitality come first below 25
// (TestDemonologyTakesUnholyPowerBeforeHealthFunnel), so this holds Demonic Sacrifice only.
func TestDemonologyKeepsDemonicSacrificeOutBelowTwentyFive(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "warlock")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, "warlock", "demonology")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	const demonicSacrifice = 9 // Demonology's tree position of Demonic Sacrifice
	for level := 10; level <= 24; level++ {
		got := LadderTalentString(trees, targets, 1, level, ForSpec("warlock-demonology"))
		if demo := got[len("00000000000000000-"):]; demo[demonicSacrifice] != '0' {
			t.Errorf("level %d: %s takes Demonic Sacrifice", level, got)
		}
	}
}
