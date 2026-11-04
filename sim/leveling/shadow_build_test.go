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

// Spirit Tap first: every guide levels with it from level 10 (it is how a priest drinks less
// between pulls), and on the band fight the order is free, since it and Blackout change no
// number there. Shadow Focus follows, then the walk; Mind Flay stays at 28, because on the
// band-20 fight a rank-1 Mind Flay without Improved Mind Flay loses to wanding (mana).
func TestShadowLevelsSpiritTapFirst(t *testing.T) {
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
	const none = "000000000000000000-00000000000000000-"
	want := map[int]string{
		14: none + "00500000000000000",
		19: none + "50500000000000000",
		20: none + "51500000000000000",
		30: none + "52542001200000000",
		40: none + "52542001201300051",
		50: "025102000000000000-00000000000000000-52542001201300051",
	}
	for level, w := range want {
		got := LadderTalentString(trees, targets, shadowTree, level, ForSpec("priest-shadow"))
		if got != w {
			t.Errorf("level %d Shadow build %q, want %q", level, got, w)
		}
	}
}
