package builds

import (
	"encoding/json"
	"os"
	"path/filepath"
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
// 101 then one in 102 (both tier 0 of the Arms tree, columns 0 and 1; the fixture is 0-based like talents.json, the code 1-based),
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
	want := "FSB1:test-1:warrior:111111112:head=12640:crit=2;hit=2;strength=18"
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
	want := "FSB1:test-1:warrior:111111112:head=999999"
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

// TestAddonCodeWritesOneBasedTriplesForTheRealHunterTree pins the bug
// lifecycle round 3 found: talents.json is 0-based (104960 is tier 0,
// column 1 of the first tab) while the format's triples are 1-based, so
// the 20-29 Beast Mastery build must render with no 0 in its order.
func TestAddonCodeWritesOneBasedTriplesForTheRealHunterTree(t *testing.T) {
	data, err := trees.Load(filepath.Join("..", "..", "..", "data", "builds"))
	if err != nil {
		t.Skipf("pipeline tree data not loadable here: %v", err)
	}
	b := Build{
		ClassID: 3, RaceID: 1, TreeVersion: "1.60.1.70450",
		PointOrder: []int{104960, 104960, 104960, 104960, 104960,
			104976, 104976, 104976, 104976, 104975, 104975},
	}
	code, err := AddonCode(data, b)
	if err != nil {
		t.Fatal(err)
	}
	want := "FSB1:1.60.1.70450:hunter:112112112112112113113113113121121:"
	if code != want {
		t.Fatalf("code = %q, want %q", code, want)
	}
}

// TestAddonCodeOrderMatchesTheWebCodecVector keeps the API and the web
// planner from drifting again: the shared codec vector's three points are
// the first two cells of tier 0 and the first of tier 1 in the first
// tab, which is warrior talents 101, 102 and 103 in the fixture tree.
func TestAddonCodeOrderMatchesTheWebCodecVector(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "..", "..", "web", "src", "fixtures", "addon", "codec-vectors.json"))
	if err != nil {
		t.Fatal(err)
	}
	var vectors struct {
		FSB1 []struct {
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"fsb1"`
	}
	if err := json.Unmarshal(raw, &vectors); err != nil {
		t.Fatal(err)
	}
	var wantOrder string
	for _, v := range vectors.FSB1 {
		if v.Name == "three points, one item with two stats" {
			wantOrder = strings.Split(v.Code, ":")[3]
		}
	}
	if wantOrder == "" {
		t.Fatal("the web codec vector is missing")
	}
	b := warriorBuild(t)
	b.PointOrder = []int{101, 102, 103}
	b.Gear = nil
	code, err := AddonCode(fixtureData(t), b)
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.Split(code, ":")[3]; got != wantOrder {
		t.Fatalf("order = %q, web vector has %q", got, wantOrder)
	}
}
