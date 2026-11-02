package request

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// The raid preset uses a buff's improved form when the live tree carries the
// talent and a raid would take it. Restorative Totems (16187, Restoration
// tier 3: "Increases the effect of your Mana Spring Totem by 25%") is live
// (forever-bugs #215), and our own Restoration shaman raid build takes it
// 5/5, so the raid's Mana Spring is the improved one (round 11, auras lane).
func TestRaidManaSpringIsTheRestorationShamansImprovedTotem(t *testing.T) {
	const restorativeTotemsIndex = 10 // Restoration tree, tier 3 column 1
	_, trees, err := leveling.GuideRaidBuildTalents(repoRoot, "shaman", "restoration")
	if err != nil {
		t.Fatal(err)
	}
	if got := trees[2][restorativeTotemsIndex]; got != '5' {
		t.Fatalf("the Restoration raid build takes Restorative Totems %c/5; this test assumes 5/5", got)
	}
	resolved, err := loadRealPresets(t).Resolve(RaidPreset, specs.ByKey["mage-arcane"])
	if err != nil {
		t.Fatal(err)
	}
	have := buffAndDebuffIDs(resolved)
	if !slices.Contains(have, "mana_spring_totem"+improvedSuffix) || slices.Contains(have, "mana_spring_totem") {
		t.Errorf("the raid preset's Mana Spring is not the improved totem: %v", have)
	}
}
