package builds

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func fixtureData(t *testing.T) *trees.Data {
	t.Helper()
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	return data
}

// warriorBuild is the fixture build addoncode_test.go and the addon
// package's own tests share: class 1 (warrior), two points in talent
// 101 then one in 102 (both tier 0 of the Arms tree, columns 0 and 1),
// and a single gear slot whose item carries three stats.
func warriorBuild(t *testing.T) Build {
	t.Helper()
	b, err := New(Input{
		ClassID: 1, RaceID: 1, TreeVersion: "test-1",
		PointOrder: []int{101, 101, 102},
		Gear:       map[string]int{"head": 12640},
		Title:      "Arms leveling",
	})
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestAddonCodeRendersTheOrderAndGearFromTheStoredBuild(t *testing.T) {
	code, err := AddonCode(fixtureData(t), warriorBuild(t))
	if err != nil {
		t.Fatal(err)
	}
	want := "FSB1:test-1:warrior:000000001:head=12640:crit=2;hit=2;strength=18"
	if code != want {
		t.Fatalf("code = %q, want %q", code, want)
	}
}

func TestAddonCodeOmitsStatsForAnItemTheTreeDataDoesNotKnow(t *testing.T) {
	b := warriorBuild(t)
	b.Gear = map[string]int{"head": 999999}
	code, err := AddonCode(fixtureData(t), b)
	if err != nil {
		t.Fatal(err)
	}
	want := "FSB1:test-1:warrior:000000001:head=999999"
	if code != want {
		t.Fatalf("code = %q, want %q (no stats for an unknown item)", code, want)
	}
}

func TestAddonCodeOrdersGearBySlotNotByMapIteration(t *testing.T) {
	b := warriorBuild(t)
	b.Gear = map[string]int{"finger1": 19325, "head": 12640}
	code, err := AddonCode(fixtureData(t), b)
	if err != nil {
		t.Fatal(err)
	}
	// Slots lists head before finger1, so head must render first
	// regardless of the map's own iteration order.
	if !strings.Contains(code, "head=12640:crit=2;hit=2;strength=18,finger1=19325:attack_power=24") {
		t.Fatalf("code = %q, want head before finger1 in the planner's slot order", code)
	}
}

func TestAddonCodeRefusesABuildWithNoTreeData(t *testing.T) {
	b := warriorBuild(t)
	b.TreeVersion = "no-such-build"
	if _, err := AddonCode(fixtureData(t), b); err == nil {
		t.Fatal("want an error for a tree version with no data")
	}
}

func TestAddonCodeRefusesAnUnknownClass(t *testing.T) {
	b := warriorBuild(t)
	b.ClassID = 99
	if _, err := AddonCode(fixtureData(t), b); err == nil {
		t.Fatal("want an error for a class the tree data does not have")
	}
}

func TestAddonCodeRefusesAnUnknownTalent(t *testing.T) {
	b := warriorBuild(t)
	b.PointOrder = []int{999}
	if _, err := AddonCode(fixtureData(t), b); err == nil {
		t.Fatal("want an error for a spend order the tree data cannot resolve")
	}
}

func TestAddonCodeHandlesAnEmptyBuild(t *testing.T) {
	b := warriorBuild(t)
	b.PointOrder = nil
	b.Gear = nil
	code, err := AddonCode(fixtureData(t), b)
	if err != nil {
		t.Fatal(err)
	}
	if code != "FSB1:test-1:warrior::" {
		t.Fatalf("code = %q, want the order and gear segments empty but present", code)
	}
}
