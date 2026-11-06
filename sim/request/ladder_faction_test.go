package request

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

// Round 26 platform: the ladder's own pickers were faction-blind. Round 25
// (P5) filtered only the below-band refill, so the ladder still handed a
// character the other faction's quest rewards and faction-restricted items.
// Every ladder pick, for either faction, is one that faction can earn.
func TestTheLadderGearIsOneTheFactionCanEarn(t *testing.T) {
	build := activeBuild(t)
	for _, faction := range []string{"horde", "alliance"} {
		other, err := loadOtherFactionItems(repoRoot, build, faction)
		if err != nil {
			t.Fatal(err)
		}
		for _, spec := range specs.All {
			items, err := loadClassItems(repoRoot, build, spec.ClassSlug)
			if err != nil {
				t.Fatal(err)
			}
			known, err := obtainableItemIDs(repoRoot, build)
			if err != nil {
				t.Fatal(err)
			}
			floors, err := loadRequiredLevelFloors(repoRoot, build, items)
			if err != nil {
				t.Fatal(err)
			}
			for _, level := range ladderLevels {
				for _, g := range ladderGear(items, known, floors, other, spec.Spec, level) {
					if other[g.ItemID] {
						t.Errorf("%s %s %d: %s %d is only the other faction's", faction, spec.Spec, level, g.Slot, g.ItemID)
					}
				}
			}
		}
	}
}

// A quest whose faction loot.json could not read ("unknown") is not the other
// faction's: the BiS ranker reads it as open to both (questFactionSide), and
// so does the gear filter. Scholarly Pendant (277203) is the reward of Friend
// of the Library, which both factions take; round 25's filter kept it from
// both.
func TestAQuestOfUnknownFactionIsOpenToBoth(t *testing.T) {
	build := activeBuild(t)
	for _, faction := range []string{"horde", "alliance"} {
		other, err := loadOtherFactionItems(repoRoot, build, faction)
		if err != nil {
			t.Fatal(err)
		}
		if other[277203] {
			t.Errorf("%s: Scholarly Pendant (Friend of the Library, faction unknown) reads as the other faction's", faction)
		}
	}
}
