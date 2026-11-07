package leveling

import (
	"path/filepath"
	"testing"
)

// Shadow priest node ids on 1.60.1.70755 (data/builds/1.60.1.70755/talents/priest.json).
const (
	shadowBlackout           = 105832
	shadowSpiritTap          = 105833
	shadowDevouringContagion = 105819
	shadowSilence            = 105824
	shadowReach              = 105829
	shadowDarkInspiration    = 105823
)

// Sim-accuracy round 15 (priest lane, guide-accuracy round 1 PRIEST-1/B1): the leveling
// build takes Spirit Tap 5/5 (100% on a kill; 3/5 is 60%) and Blackout 2/5. Neither talent
// changes a number on the band fight (no kills, a stun), so the band builds are unchanged in
// the sim; every Forever leveling guide takes Spirit Tap 5/5 and none takes Blackout 4/5.
func TestShadowLevelingBuildTakesSpiritTapFiveBlackoutTwo(t *testing.T) {
	_, targets := guideTargets(t, "priest", "shadow", GuideBuildTalents)
	if targets[shadowSpiritTap] != 5 || targets[shadowBlackout] != 2 {
		t.Errorf("Spirit Tap %d, Blackout %d; want 5 and 2", targets[shadowSpiritTap], targets[shadowBlackout])
	}
}

// The raid build fights a boss: Blackout's stun never lands (bosses are immune) and Spirit
// Tap needs a kill. Its five points go to Devouring Contagion 2, Silence 1, Shadow Reach 2/2
// and Dark Inspiration 1, the strong players' raid picks; the tiers below Shadowform keep 25.
func TestShadowRaidBuildSpendsNothingOnBossDeadTalents(t *testing.T) {
	_, targets := guideTargets(t, "priest", "shadow", GuideRaidBuildTalents)
	if targets[shadowBlackout] != 0 || targets[shadowSpiritTap] != 0 {
		t.Errorf("raid build Blackout %d, Spirit Tap %d; want 0 and 0", targets[shadowBlackout], targets[shadowSpiritTap])
	}
	want := map[int]int{shadowDevouringContagion: 2, shadowSilence: 1, shadowReach: 2, shadowDarkInspiration: 1}
	for id, rank := range want {
		if targets[id] != rank {
			t.Errorf("raid build talent %d at %d, want %d", id, targets[id], rank)
		}
	}
}

// Wand Specialization 2/2 first, then Spirit Tap 5/5 (sim-accuracy round 15 wave 6, an
// owner call delegated to the lane): the wand is 37 to 41% of a leveling priest's damage
// on the band fight, so its +25% gains 10.1% at band 20 and 5.7% at 30, and every Forever
// guide that orders it takes it first (Icy Veins Forever Shadow leveling, Kargoz's Forever
// 1-60 guide). A ladder only adds points, so the two Discipline points move Shadowform from
// 40 to 42 (band 40 -36%); a priest spends 30 levels before 40 and two after it, and the
// guides accept the same trade. Shadow Focus, Improved Mind Blast, Improved Shadow Word:
// Pain, Mind Flay and Improved Mind Flay follow ahead of Blackout (which the band fight does
// not model), so band 30 holds Improved Mind Flay (+5.9% / +4.8% there); Mind Flay lands at
// 28, past the band-20 fight where a lone rank-1 Mind Flay loses on mana.
func TestShadowLevelsWithWandSpecializationThenSpiritTap(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "priest")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, "priest", "shadow")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	const shadowTree = 2
	const wand = "020000000000000000-00000000000000000-"
	want := map[int]string{
		10: "010000000000000000-00000000000000000-00000000000000000",
		11: wand + "00000000000000000",
		16: wand + "00500000000000000",
		20: wand + "40500000000000000",
		25: wand + "50540000000000000",
		30: wand + "50542001200000000",
		40: wand + "52542001201300040",
		41: wand + "52542001201300050",
		42: wand + "52542001201300051",
		// Sim-accuracy round 29 (walks-dps): Mental Strength before Silent Resolve's threat point.
		50: "025003000000000000-00000000000000000-52542001201300051",
	}
	for level, w := range want {
		got := LadderTalentString(trees, targets, shadowTree, level, ForSpec("priest-shadow"))
		if got != w {
			t.Errorf("level %d Shadow build %q, want %q", level, got, w)
		}
	}
}
