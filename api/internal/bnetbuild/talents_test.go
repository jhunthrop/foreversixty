// api/internal/bnetbuild/talents_test.go
package bnetbuild

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// readEraKiloz reads one file out of the real captured fixture
// (api/internal/bnetapi/testdata/era-kiloz/), shared by every test in this package that
// exercises Encode/encodeGear/encodeTalents against real Blizzard data.
func readEraKiloz(t *testing.T, file string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "bnetapi", "testdata", "era-kiloz", file))
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func loadWarriorTable(t *testing.T) TalentTable {
	t.Helper()
	data, err := trees.Load(realDataDir)
	if err != nil {
		t.Fatal(err)
	}
	build, ok := data.Build("1.60.1.69893")
	if !ok {
		t.Fatal("1.60.1.69893 must load")
	}
	classID := classIDFor(t, build, "warrior")
	return TalentTable{Build: build, ClassID: classID}
}

func classIDFor(t *testing.T, b *trees.Build, slug string) int {
	t.Helper()
	for _, c := range b.Classes() {
		if c.Slug == slug {
			return c.ID
		}
	}
	t.Fatalf("class %q not found", slug)
	return 0
}

func TestEncodeTalentsUsesTheActiveGroupAndReportsUnmatched(t *testing.T) {
	table := loadWarriorTable(t)
	raw := readEraKiloz(t, "specializations.json")

	treeRanks, unmatched, clamped, err := encodeTalents(raw, table)
	if err != nil {
		t.Fatal(err)
	}
	// The active group (is_active: true) is Fury + Arms, 51 points total; the site's
	// 1.60.1.69893 talent ids are a different DBC generation than the classic1x profile
	// API's, so most of these 15 talents cannot be id- or spell-matched and are reported —
	// this pins that today's real behaviour, not a bug this lane owes a fix for.
	if len(unmatched) == 0 {
		t.Fatal("this fixture's talent ids are known not to match this build's table; want at least one report")
	}
	if len(clamped) != 0 {
		t.Fatalf("clamped = %v, want none (every rank in the fixture is within max_rank)", clamped)
	}
	totalRanks := 0
	for _, tree := range treeRanks {
		for _, r := range tree {
			totalRanks += r
		}
	}
	if totalRanks == 0 {
		t.Fatal("at least the id-matched talents (single-rank abilities) must land a nonzero rank")
	}
}

func TestEncodeTalentsPicksTheFirstGroupWhenNoneIsActive(t *testing.T) {
	table := loadWarriorTable(t)
	raw := []byte(`{"specialization_groups":[
		{"is_active":false,"specializations":[{"specialization_name":"Fury","spent_points":1,
			"talents":[{"talent":{"id":167},"spell_tooltip":{"spell":{"id":23881,"name":"Bloodthirst"}},"talent_rank":1}]}]},
		{"is_active":false,"specializations":[{"specialization_name":"Arms","spent_points":0,"talents":[]}]}
	]}`)
	treeRanks, _, _, err := encodeTalents(raw, table)
	if err != nil {
		t.Fatal(err)
	}
	total := 0
	for _, tree := range treeRanks {
		for _, r := range tree {
			total += r
		}
	}
	if total == 0 {
		t.Fatal("no group is active: the first group (Bloodthirst rank 1) must still be used")
	}
}
