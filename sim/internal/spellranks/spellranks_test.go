package spellranks

import (
	"os"
	"testing"
)

// spellRanksExcerpt is Serpent Sting's real chain (nine ranks, level 4
// through 60, plus its rank-9 NPC-copy noise at level 0), Eagle Eye (a
// single-rank, level-gated hunter ability with no progression), and
// Thorns (a druid ability with both rank-0 noise at level 20/0 and a
// real six-rank player chain) - cut from the shape
// data/normalize is expected to emit.
func spellRanksExcerpt(t *testing.T) map[string]map[int32]*rankChain {
	t.Helper()
	b, err := os.ReadFile("testdata/spellranks.excerpt.json")
	if err != nil {
		t.Fatal(err)
	}
	table, err := parseSpellRanks(b)
	if err != nil {
		t.Fatal(err)
	}
	return table
}

func TestParseSpellRanksChainsRanksAndDropsRankZeroDuplicates(t *testing.T) {
	table := spellRanksExcerpt(t)
	hunter, ok := table["hunter"]
	if !ok {
		t.Fatal(`no "hunter" class in the parsed table`)
	}

	// Every id across Serpent Sting's nine ranks - including its rank-9
	// NPC-copy duplicates at level 0 - resolves to ONE chain object.
	serpentStingIDs := []int32{
		1978, 425728, 13549, 425729, 13550, 425730, 13551, 425732,
		13552, 425733, 13553, 425734, 13554, 425735, 13555, 425736,
		25295, 425737, 1232979, 1232980, 1232981, 1232982,
	}
	var chain *rankChain
	for _, id := range serpentStingIDs {
		c, ok := hunter[id]
		if !ok {
			t.Fatalf("Serpent Sting id %d is missing from the hunter table", id)
		}
		if chain == nil {
			chain = c
		} else if c != chain {
			t.Errorf("Serpent Sting id %d resolved to a different chain than id %d", id, serpentStingIDs[0])
		}
	}
	if len(chain.tiers) != 9 {
		t.Fatalf("Serpent Sting has %d tiers, want 9", len(chain.tiers))
	}
	// The rank-9 tier's level is 60, from the FIRST rank-9 row seen
	// (25295) - never lowered to 0 by the NPC-copy duplicates that
	// follow it in the file.
	top := chain.tiers[8]
	if top.level != 60 {
		t.Errorf("Serpent Sting's rank-9 tier is learned at %d, want 60 (the NPC-copy duplicates at level 0 must not corrupt it)", top.level)
	}
	wantTopIDs := []int32{25295, 425737, 1232979, 1232980, 1232981, 1232982}
	if len(top.ids) != len(wantTopIDs) {
		t.Fatalf("Serpent Sting's rank-9 tier carries %d ids, want %d: %v", len(top.ids), len(wantTopIDs), top.ids)
	}

	// Eagle Eye has exactly one rank: a single-tier chain, not "no
	// chain at all" - it still gates on a learn level.
	eagleEye, ok := hunter[6197]
	if !ok {
		t.Fatal("Eagle Eye (6197) is missing from the hunter table")
	}
	if len(eagleEye.tiers) != 1 || eagleEye.tiers[0].level != 20 {
		t.Errorf("Eagle Eye's chain = %+v, want one tier at level 20", eagleEye.tiers)
	}

	druid, ok := table["druid"]
	if !ok {
		t.Fatal(`no "druid" class in the parsed table`)
	}
	// Thorns' rank-0 duplicates (21335, 21337, 1236308) are never
	// treated as ranks of the spell: the reference caveat this package
	// exists to honor.
	for _, id := range []int32{21335, 21337, 1236308} {
		if _, ok := druid[id]; ok {
			t.Errorf("Thorns' rank-0 duplicate %d was chained as if it were a player rank", id)
		}
	}
	thorns, ok := druid[467]
	if !ok {
		t.Fatal("Thorns rank 1 (467) is missing from the druid table")
	}
	if len(thorns.tiers) != 6 {
		t.Fatalf("Thorns has %d tiers, want 6 (its rank-0 duplicates must not add tiers)", len(thorns.tiers))
	}
	if thorns.tiers[5].level != 54 || len(thorns.tiers[5].ids) != 2 {
		t.Errorf("Thorns' rank-6 tier = %+v, want level 54 with both ids 9910 and 16877", thorns.tiers[5])
	}
}

func TestHighestLearnedSpellIDResolvesByLevel(t *testing.T) {
	hunter := spellRanksExcerpt(t)["hunter"]
	cases := []struct {
		name      string
		id        int32
		level     int
		wantID    int32
		wantKnown bool
	}{
		{"level 60: the top rank, unchanged", 25295, 60, 25295, true},
		{"level 38: Serpent Sting resolves down to rank 5 (learned at 34)", 25295, 38, 13552, true},
		{"level 42: exactly rank 6's learn level", 25295, 42, 13553, true},
		{"level 3: below Serpent Sting's first rank (4) - unlearned", 25295, 3, 0, false},
		{"Eagle Eye at its learn level", 6197, 20, 6197, true},
		{"Eagle Eye below its learn level - unlearned", 6197, 19, 0, false},
		{"an id this table has never heard of stays untouched", 999999, 60, 999999, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotID, gotOK := highestLearnedIn(hunter, c.id, c.level)
			if gotID != c.wantID || gotOK != c.wantKnown {
				t.Errorf("highestLearnedIn(%d, %d) = (%d, %v), want (%d, %v)", c.id, c.level, gotID, gotOK, c.wantID, c.wantKnown)
			}
		})
	}
}

func TestParseSpellRanksRejectsAnEmptyTable(t *testing.T) {
	if _, err := parseSpellRanks([]byte(`{"build":"x","classes":{}}`)); err == nil {
		t.Error("a table with no classes was accepted")
	}
	if _, err := parseSpellRanks([]byte(`not json`)); err == nil {
		t.Error("unreadable JSON was accepted")
	}
}

// The active build's embedded table loads and HighestLearnedSpellID
// answers something for a real class/id pair. This is the same
// "the embed loads" guard simdb_test.go's TestDatabaseLoads gives
// simdb.bin and enchants.json.
func TestEmbeddedSpellRanksLoad(t *testing.T) {
	table, err := spellRankTable()
	if err != nil {
		t.Fatal(err)
	}
	if len(table) == 0 {
		t.Fatal("the embedded spell rank table carries no classes")
	}
	hunter, ok := table["hunter"]
	if !ok || len(hunter) == 0 {
		t.Fatal(`the embedded table carries no "hunter" entries`)
	}
	// Serpent Sting's max-rank id (25295), unchanged at level 60.
	if id, ok := HighestLearnedSpellID("hunter", 25295, 60); !ok || id != 25295 {
		t.Errorf(`HighestLearnedSpellID("hunter", 25295, 60) = (%d, %v), want (25295, true)`, id, ok)
	}
}
