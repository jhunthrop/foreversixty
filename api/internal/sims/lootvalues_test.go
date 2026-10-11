package sims

import (
	"context"
	"math"
	"testing"
)

const lootBuild = "1.60.1.70291"

func TestDPSPerPointReadsTheBuildsFile(t *testing.T) {
	useRepoGearData(t)
	if got, ok := DPSPerPoint(lootBuild, "warrior-arms"); !ok || got != 0.2342 {
		t.Fatalf("warrior-arms = %v, %v; want 0.2342", got, ok)
	}
	// The newest loaded build answers for a run that names no build.
	if got, ok := DPSPerPoint("", "mage-fire"); !ok || got <= 0 {
		t.Fatalf("mage-fire in the newest build = %v, %v", got, ok)
	}
	// The tank specs have no entry.
	if got, ok := DPSPerPoint(lootBuild, "warrior-protection"); ok {
		t.Fatalf("warrior-protection = %v, want no entry", got)
	}
	if _, ok := DPSPerPoint("0.0.0.0", "warrior-arms"); ok {
		t.Fatal("a build with no file has no entries")
	}
}

func TestRaidKeyMatchesAZoneHoweverItIsWritten(t *testing.T) {
	for _, zone := range []string{"Onyxia's Lair", "onyxias lair", "  ONYXIA'S  LAIR "} {
		if raidKey(zone) != raidKey("Onyxia's Lair") {
			t.Errorf("raidKey(%q) = %q", zone, raidKey(zone))
		}
	}
	if raidKey("Molten Core") == raidKey("Onyxia's Lair") {
		t.Fatal("two raids share a key")
	}
}

// A Molten Core chest for a warrior wearing a 10-strength chest: the shipped warrior-arms
// weights score both stat lines, and the difference is turned into DPS with the spec's
// DPS per attack power.
func TestLootValuesPricesTheRaidsLootAgainstWhatIsWorn(t *testing.T) {
	useRepoGearData(t)
	const code = "FS1:1.60.1.70291:warrior:human:0/0/0:chest=6627"
	set, ok, err := LootValues(context.Background(), "Molten Core", code, "warrior-arms", nil)
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	if set.Raid != "Molten Core" || len(set.Items) == 0 {
		t.Fatalf("set = %s with %d items", set.Raid, len(set.Items))
	}

	src := currentGearData()
	b, _ := src.trees.Build(lootBuild)
	classID, _ := classIDBySlug(b, "warrior")
	weights, err := src.specWeights(lootBuild, "warrior-arms")
	if err != nil {
		t.Fatal(err)
	}
	dpp, _ := DPSPerPoint(lootBuild, "warrior-arms")
	worn, _ := b.Item(classID, 6627)
	breastplate, ok := b.Item(classID, 16865) // Breastplate of Might, Molten Core
	if !ok {
		t.Fatal("the warrior item pool has no Breastplate of Might")
	}
	want := math.Round((statScore(breastplate.Stats, weights)-statScore(worn.Stats, weights))*dpp*10) / 10
	got, ok := set.Items[16865]
	if !ok || got.Value != want || got.Unit != UnitDPS || got.Planned {
		t.Fatalf("Breastplate of Might = %+v (present %v), want %v dps, not planned", got, ok, want)
	}
	if want <= 0 {
		t.Fatalf("a 20-strength chest over a 10-strength one must be worth something, got %v", want)
	}
	// A ring against two empty finger slots is its whole stat line.
	for id := range set.Items {
		item, _ := b.Item(classID, id)
		if item.Slot != "finger" {
			continue
		}
		ring := math.Round(statScore(item.Stats, weights)*dpp*10) / 10
		if set.Items[id].Value != ring {
			t.Fatalf("ring %d = %v, want %v", id, set.Items[id].Value, ring)
		}
		break
	}
	// Every priced item is one the warrior pool has a stat line for.
	for id := range set.Items {
		if item, ok := b.Item(classID, id); !ok || len(item.Stats) == 0 {
			t.Fatalf("item %d is priced with no stat line", id)
		}
	}
}

func TestLootValuesUseTheCharactersOwnWeightsRunWhenGiven(t *testing.T) {
	useRepoGearData(t)
	const code = "FS1:1.60.1.70291:warrior:human:0/0/0:chest=6627"
	run := map[string]float64{"strength": 1}
	set, ok, err := LootValues(context.Background(), "Molten Core", code, "warrior-arms", run)
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	dpp, _ := DPSPerPoint(lootBuild, "warrior-arms")
	// Breastplate of Might has 20 strength, the worn chest 10: ten points of strength.
	if got, want := set.Items[16865].Value, math.Round(10*dpp*10)/10; got != want {
		t.Fatalf("Breastplate of Might = %v, want %v", got, want)
	}
}

func TestLootValuesAreInScoreUnitsWithoutADPSPerPoint(t *testing.T) {
	useRepoGearData(t)
	const code = "FS1:1.60.1.70291:warrior:human:0/0/0:chest=6627"
	set, ok, err := LootValues(context.Background(), "Molten Core", code, "warrior-protection", nil)
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	for id, v := range set.Items {
		if v.Unit != UnitScore {
			t.Fatalf("item %d unit = %q, want score", id, v.Unit)
		}
	}
}

func TestLootValuesOfAZoneNoRaidNamesIsNothing(t *testing.T) {
	useRepoGearData(t)
	const code = "FS1:1.60.1.70291:warrior:human:0/0/0:chest=6627"
	if _, ok, err := LootValues(context.Background(), "Elwynn Forest", code, "warrior-arms", nil); ok || err != nil {
		t.Fatalf("ok=%v err=%v, want no raid", ok, err)
	}
}

// planned is the spec's level-60 best-in-slot picks: no raid item is one in today's data
// (the band-60 sets are pre-raid), so the picks themselves are what is checked.
func TestPlannedItemsAreTheLevelSixtyPicks(t *testing.T) {
	useRepoGearData(t)
	planned, err := plannedItems(currentGearData().dir, lootBuild, "warrior-arms")
	if err != nil {
		t.Fatal(err)
	}
	if !planned[12640] { // Lionheart Helm, the band-60 head pick
		t.Fatalf("planned = %v, want the Lionheart Helm among them", planned)
	}
	if none, err := plannedItems(currentGearData().dir, lootBuild, "no-such-spec"); err != nil || len(none) != 0 {
		t.Fatalf("a spec with no best-in-slot file = %v, %v; want nothing planned", none, err)
	}
}
