package sims

import (
	"context"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

const lootBuild = "1.60.1.70291"

func TestDPSPerPointReadsTheBuildsFile(t *testing.T) {
	useRepoGearData(t)
	if got, ok := DPSPerPoint(lootBuild, "warrior-arms"); !ok || got != (DPSPoint{Reference: "attack_power", DPS: 0.2222}) {
		t.Fatalf("warrior-arms = %v, %v; want 0.2222", got, ok)
	}
	// The active build answers for a run that names no build.
	if got, ok := DPSPerPoint("", "mage-fire"); !ok || got.DPS <= 0 || got.Reference != "spell_power" {
		t.Fatalf("mage-fire in the active build = %v, %v", got, ok)
	}
	// The tank specs have no entry.
	if got, ok := DPSPerPoint(lootBuild, "warrior-protection"); ok {
		t.Fatalf("warrior-protection = %v, want no entry", got)
	}
	if _, ok := DPSPerPoint("0.0.0.0", "warrior-arms"); ok {
		t.Fatal("a build with no file has no entries")
	}
}

// A build data.yml fetched but no activation moved to (70685 on disk, with no
// dps-per-point.json yet, active-build.json still naming 70658) must not change the build
// DPSPerPoint reads by default; activating it does.
func TestDPSPerPointDefaultsToTheActiveBuildNotTheNewestFolder(t *testing.T) {
	const active, fetched = "1.60.1.70658", "1.60.1.70685"
	root := t.TempDir()
	if err := trees.LinkBuilds(root, repoBuilds, active, fetched); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { SetGearData(nil, "") })
	wire := func(build string) {
		t.Helper()
		if err := trees.WriteActiveBuildFile(root, build); err != nil {
			t.Fatal(err)
		}
		data, err := trees.Load(root)
		if err != nil {
			t.Fatal(err)
		}
		SetGearData(data, root)
	}

	wire(active)
	want, wantOK := DPSPerPoint(active, "mage-fire")
	if got, ok := DPSPerPoint("", "mage-fire"); !wantOK || ok != wantOK || got != want {
		t.Fatalf("default = %v, %v; want %s's %v, %v", got, ok, active, want, wantOK)
	}

	wire(fetched)
	want, wantOK = DPSPerPoint(fetched, "mage-fire")
	if got, ok := DPSPerPoint("", "mage-fire"); ok != wantOK || got != want {
		t.Fatalf("default after activation = %v, %v; want %s's %v, %v", got, ok, fetched, want, wantOK)
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
	p, _ := DPSPerPoint(lootBuild, "warrior-arms")
	dpp := p.DPS
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
	// The run's weight-1 stat is attack power, warrior-arms' reference in dps-per-point.json.
	run := map[string]float64{"attack_power": 1, "strength": 2}
	set, ok, err := LootValues(context.Background(), "Molten Core", code, "warrior-arms", run)
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	p, _ := DPSPerPoint(lootBuild, "warrior-arms")
	// Breastplate of Might has 20 strength, the worn chest 10: ten strength, twenty points.
	if got, want := set.Items[16865], (LootValue{Value: math.Round(20*p.DPS*10) / 10, Unit: UnitDPS}); got != want {
		t.Fatalf("Breastplate of Might = %+v, want %+v", got, want)
	}
}

// A run normalised to some other stat than the file's reference is not in that
// reference's points, so DPS per point cannot turn it into DPS: score units.
func TestLootValuesStayInScoreUnitsWhenTheRunsReferenceDiffers(t *testing.T) {
	useRepoGearData(t)
	const code = "FS1:1.60.1.70291:warrior:human:0/0/0:chest=6627"
	set, ok, err := LootValues(context.Background(), "Molten Core", code, "warrior-arms",
		map[string]float64{"strength": 1, "attack_power": 0.5})
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	if got, want := set.Items[16865], (LootValue{Value: 10, Unit: UnitScore}); got != want {
		t.Fatalf("Breastplate of Might = %+v, want %+v", got, want)
	}
}

// A mage with a staff: an off-hand item cannot be worn without giving up the staff for a
// one-hander it does not bring, so it is skipped; a two-hander is weighed against the
// staff (and the empty off hand).
func TestLootValuesUnderAWornTwoHander(t *testing.T) {
	useRepoGearData(t)
	const code = "FS1:1.60.1.70291:mage:human:0/0/0:main_hand=13249" // Argent Crusader, a staff
	set, ok, err := LootValues(context.Background(), "Molten Core", code, "mage-fire", nil)
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	if got, ok := set.Items[19142]; ok { // Fire Runed Grimoire
		t.Fatalf("Fire Runed Grimoire = %+v, want no entry under a staff", got)
	}
	src := currentGearData()
	b, _ := src.trees.Build(lootBuild)
	classID, _ := classIDBySlug(b, "mage")
	weights, err := src.specWeights(lootBuild, "mage-fire")
	if err != nil {
		t.Fatal(err)
	}
	p, _ := DPSPerPoint(lootBuild, "mage-fire")
	staff, _ := b.Item(classID, 13249)
	dominance, _ := b.Item(classID, 18842) // Staff of Dominance
	want := math.Round((statScore(dominance.Stats, weights)-statScore(staff.Stats, weights))*p.DPS*10) / 10
	if got := set.Items[18842]; got.Value != want {
		t.Fatalf("Staff of Dominance = %+v, want %v", got, want)
	}
}

// A warrior with a sword and a shield: a two-hander replaces both together.
func TestLootValuesWeighATwoHanderAgainstBothHands(t *testing.T) {
	useRepoGearData(t)
	// Sword of Omen in the main hand, Drillborer Disk in the off hand.
	const code = "FS1:1.60.1.70291:warrior:human:0/0/0:main_hand=6802,off_hand=17066"
	set, ok, err := LootValues(context.Background(), "Molten Core", code, "warrior-arms", nil)
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	src := currentGearData()
	b, _ := src.trees.Build(lootBuild)
	classID, _ := classIDBySlug(b, "warrior")
	weights, err := src.specWeights(lootBuild, "warrior-arms")
	if err != nil {
		t.Fatal(err)
	}
	p, _ := DPSPerPoint(lootBuild, "warrior-arms")
	sword, _ := b.Item(classID, 6802)
	disk, _ := b.Item(classID, 17066)
	blade, _ := b.Item(classID, 18822) // Obsidian Edged Blade, two-handed
	if !blade.TwoHand {
		t.Fatal("Obsidian Edged Blade must be two-handed for this case")
	}
	want := math.Round((statScore(blade.Stats, weights)-statScore(sword.Stats, weights)-statScore(disk.Stats, weights))*
		p.DPS*10) / 10
	if got := set.Items[18822]; got.Value != want {
		t.Fatalf("Obsidian Edged Blade = %+v, want %v against sword and shield together", got, want)
	}
	// With a one-hander worn, an off-hand item is still weighed, against the worn one.
	if got, ok := set.Items[17066]; !ok || got.Value != 0 {
		t.Fatalf("Drillborer Disk against itself = %+v (present %v), want 0", got, ok)
	}
}

// Only a successful read of dps-per-point.json is kept: a build whose file is missing is
// asked again, and answers once the file is there.
func TestDPSPerPointKeepsOnlyASuccessfulRead(t *testing.T) {
	dir := t.TempDir()
	useRepoGearData(t)
	SetGearData(currentGearData().trees, dir)
	t.Cleanup(func() { SetGearData(nil, "") })
	if _, ok := DPSPerPoint(lootBuild, "warrior-arms"); ok {
		t.Fatal("no file yet, want no entry")
	}
	if err := os.MkdirAll(filepath.Join(dir, lootBuild), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, lootBuild, "dps-per-point.json"),
		[]byte(`{"warrior-arms":{"reference":"attack_power","dps":0.5,"healing":false}}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if p, ok := DPSPerPoint(lootBuild, "warrior-arms"); !ok || p != (DPSPoint{Reference: "attack_power", DPS: 0.5}) {
		t.Fatalf("after the file appears = %+v, %v", p, ok)
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
	planned, err := currentGearData().plannedItems(lootBuild, "warrior-arms")
	if err != nil {
		t.Fatal(err)
	}
	if !planned[12640] { // Lionheart Helm, the band-60 head pick
		t.Fatalf("planned = %v, want the Lionheart Helm among them", planned)
	}
	if none, err := currentGearData().plannedItems(lootBuild, "no-such-spec"); err != nil || len(none) != 0 {
		t.Fatalf("a spec with no best-in-slot file = %v, %v; want nothing planned", none, err)
	}
}
