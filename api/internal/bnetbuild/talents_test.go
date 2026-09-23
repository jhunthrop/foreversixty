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

	result, err := encodeTalents(raw, table)
	if err != nil {
		t.Fatal(err)
	}
	// The active group (is_active: true) is Fury (34 points) + Arms (17 points), 51 total,
	// 15 talents. Name matching (the primary key, see talents.go's doc comment) resolves 13
	// of them: the two misses are "Improved Battle Shout" and "Tactical Mastery", which
	// Forever renamed or removed from the warrior trees — this site's 1.60.1.69893 data has
	// no talent by either name, in any tree, so both id keys (also known not to agree with
	// Era's classic1x ids) never get a chance to try. That is the correct, honest outcome
	// for a talent Forever no longer has under that name, not a bug this lane owes a fix
	// for.
	wantUnmatched := []string{"Improved Battle Shout", "Tactical Mastery"}
	if len(result.Unmatched) != len(wantUnmatched) {
		t.Fatalf("unmatched = %v, want exactly %v", result.Unmatched, wantUnmatched)
	}
	for i, name := range wantUnmatched {
		if result.Unmatched[i] != name {
			t.Fatalf("unmatched = %v, want exactly %v", result.Unmatched, wantUnmatched)
		}
	}
	if len(result.Clamped) != 0 {
		t.Fatalf("clamped = %v, want none (every rank in the fixture is within max_rank)", result.Clamped)
	}
	if result.MatchedByName != 13 {
		t.Fatalf("MatchedByName = %d, want 13 (15 talents minus the two renamed/removed)", result.MatchedByName)
	}
	if result.MatchedByID != 0 || result.MatchedBySpell != 0 {
		t.Fatalf("MatchedByID = %d, MatchedBySpell = %d, want both 0: Era's talent.id and per-rank spell id are"+
			" both known not to agree with this build's data, so name matching alone accounts for every match",
			result.MatchedByID, result.MatchedBySpell)
	}
	totalRanks := 0
	for _, tree := range result.TreeRanks {
		for _, r := range tree {
			totalRanks += r
		}
	}
	// 51 points spent in the game, minus the 5+5 the two unmatched (both rank 5) talents
	// carried: this site simply has nowhere to put those 10 points.
	if totalRanks != 41 {
		t.Fatalf("total encoded ranks = %d, want 41 (51 spent minus the two unmatched talents' 5+5)", totalRanks)
	}
}

func TestEncodeTalentsPicksTheFirstGroupWhenNoneIsActive(t *testing.T) {
	table := loadWarriorTable(t)
	raw := []byte(`{"specialization_groups":[
		{"is_active":false,"specializations":[{"specialization_name":"Fury","spent_points":1,
			"talents":[{"talent":{"id":167},"spell_tooltip":{"spell":{"id":23881,"name":"Bloodthirst"}},"talent_rank":1}]}]},
		{"is_active":false,"specializations":[{"specialization_name":"Arms","spent_points":0,"talents":[]}]}
	]}`)
	result, err := encodeTalents(raw, table)
	if err != nil {
		t.Fatal(err)
	}
	total := 0
	for _, tree := range result.TreeRanks {
		for _, r := range tree {
			total += r
		}
	}
	if total == 0 {
		t.Fatal("no group is active: the first group (Bloodthirst rank 1) must still be used")
	}
	if result.MatchedByName != 1 {
		t.Fatalf("MatchedByName = %d, want 1 (Bloodthirst matches by name)", result.MatchedByName)
	}
}

// TestMatchTalentFallsBackToIDThenSpellWhenNameMisses pins the fallback order directly,
// independent of the real fixture's own (all-name) matches: a name that matches nothing in
// this build's data falls back to talent.id, and only then to the spell id.
func TestMatchTalentFallsBackToIDThenSpellWhenNameMisses(t *testing.T) {
	table := loadWarriorTable(t)
	classTrees := table.Build.Trees(table.ClassID)
	byName := talentsByName(classTrees)

	// Bloodthirst's real id/spell id in this build's data, found by its real name, so the
	// test does not hardcode ids that would break if the fixture data regenerates.
	bloodthirst, ok := byName[normalizeName("Bloodthirst")]
	if !ok {
		t.Fatal("Bloodthirst not found in this build's warrior data")
	}

	byID := blizzardTalentEntry{}
	byID.Talent.ID = bloodthirst.ID
	byID.SpellTooltip.Spell.Name = "Not A Real Talent Name"
	byID.SpellTooltip.Spell.ID = 0
	if _, key, ok := matchTalent(byID, table, byName); !ok || key != matchByID {
		t.Fatalf("matchTalent(by id) = ok=%v key=%v, want matchByID", ok, key)
	}

	bySpell := blizzardTalentEntry{}
	bySpell.Talent.ID = 0
	bySpell.SpellTooltip.Spell.Name = "Not A Real Talent Name"
	bySpell.SpellTooltip.Spell.ID = bloodthirst.SpellID
	if _, key, ok := matchTalent(bySpell, table, byName); !ok || key != matchBySpell {
		t.Fatalf("matchTalent(by spell) = ok=%v key=%v, want matchBySpell", ok, key)
	}

	none := blizzardTalentEntry{}
	none.Talent.ID = 0
	none.SpellTooltip.Spell.Name = "Not A Real Talent Name"
	none.SpellTooltip.Spell.ID = 0
	if _, _, ok := matchTalent(none, table, byName); ok {
		t.Fatal("matchTalent with no matching name, id or spell id must report no match")
	}
}
