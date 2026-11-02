package spellranks

import (
	"github.com/wowsims/classic/sim/core"
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

func TestLearnedSpellIDResolvesByLevel(t *testing.T) {
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
			gotID, gotOK := learnedIn(hunter, c.id, c.level)
			if gotID != c.wantID || gotOK != c.wantKnown {
				t.Errorf("learnedIn(%d, %d) = (%d, %v), want (%d, %v)", c.id, c.level, gotID, gotOK, c.wantID, c.wantKnown)
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

// The active build's embedded table loads and LearnedSpellID
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
	// Serpent Sting's top trainer rank (13555) is unchanged at level 60; its
	// rank 9 (25295) is an Ahn'Qiraj book rank, which resolves to the trainer
	// rank while the engine leaves the book out.
	if id, ok := LearnedSpellID("hunter", 13555, 60); !ok || id != 13555 {
		t.Errorf(`LearnedSpellID("hunter", 13555, 60) = (%d, %v), want (13555, true)`, id, ok)
	}
	wantBook := int32(13555)
	if core.IncludeAQ {
		wantBook = 25295
	}
	if id, ok := LearnedSpellID("hunter", 25295, 60); !ok || id != wantBook {
		t.Errorf(`LearnedSpellID("hunter", 25295, 60) = (%d, %v), want (%d, true)`, id, ok, wantBook)
	}
}

func TestATierWithSeveralIDsStaysInTheRotationsBand(t *testing.T) {
	// Earth Shock rank 5 at level 36: the classic 10412 beside two Forever
	// copies. A rotation naming the classic rank 7 (10414) lands on 10412;
	// one naming a Forever copy stays with the Forever ids.
	ids := []int32{10412, 408688, 1220749}
	if got := sameBandOrFirst(ids, 10414); got != 10412 {
		t.Errorf("classic original: got %d, want 10412", got)
	}
	if got := sameBandOrFirst(ids, 408690); got != 408688 {
		t.Errorf("forever original: got %d, want 408688", got)
	}
	if got := sameBandOrFirst([]int32{408688}, 10414); got != 408688 {
		t.Errorf("no classic candidate: got %d, want the first id", got)
	}
}

// bookRankExcerpt is Aspect of the Hawk's top two ranks: the trainer's rank 6
// (14322, level 58) and the Ahn'Qiraj book's rank 7 (25296, flagged book).
const bookRankExcerpt = `{"build":"x","classes":{"hunter":{"Aspect of the Hawk":[
  {"id":14321,"rank":5,"level":50},
  {"id":14322,"rank":6,"level":58},
  {"id":25296,"rank":7,"level":60,"book":true}]}}}`

// A book rank is learnable exactly when the engine's own IncludeAQ says so:
// the resolved id at level 60 is the book rank with it on, the trainer rank
// without it, never a disagreement with what the engine registers.
func TestBookRanksAreLearnedOnlyWhileTheEngineIncludesThem(t *testing.T) {
	table, err := parseSpellRanks([]byte(bookRankExcerpt))
	if err != nil {
		t.Fatal(err)
	}
	want := int32(14322)
	if core.IncludeAQ {
		want = 25296
	}
	got, ok := learnedIn(table["hunter"], 14322, 60)
	if !ok || got != want {
		t.Errorf("Aspect of the Hawk at level 60 resolves to %d (ok=%v), want %d for IncludeAQ=%v", got, ok, want, core.IncludeAQ)
	}
}

func TestRankAvailableFollowsTheEnginesConstant(t *testing.T) {
	if !RankAvailable(false) {
		t.Error("a trainer rank must always be available")
	}
	if RankAvailable(true) != core.IncludeAQ {
		t.Errorf("a book rank is available = %v, want core.IncludeAQ = %v", RankAvailable(true), core.IncludeAQ)
	}
}

// inferiorRankExcerpt is Aspect of the Hawk's real top ranks with the AQ book off:
// rank 6 (14322) gives 55 ranged attack power against rank 5's 90, so it is
// flagged inferior (data/curated/inferior-ranks.json).
const inferiorRankExcerpt = `{"build":"x","classes":{"hunter":{"Aspect of the Hawk":[
  {"id":14320,"rank":4,"level":38},
  {"id":14321,"rank":5,"level":48},
  {"id":14322,"rank":6,"level":58,"inferior":true}]}}}`

// An inferior rank is never the highest learned rank: a rotation naming the
// inferior top rank, or the strongest rank, casts the strongest learnable
// rank. One naming a lower rank on purpose keeps it (a downrank is never
// lifted).
func TestInferiorRanksAreNeverTheHighestLearnedRank(t *testing.T) {
	table, err := parseSpellRanks([]byte(inferiorRankExcerpt))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ named, want int32 }{{14320, 14320}, {14321, 14321}, {14322, 14321}} {
		for _, level := range []int{48, 58, 60} {
			got, ok := learnedIn(table["hunter"], c.named, level)
			if !ok || got != c.want {
				t.Errorf("named %d at level %d resolves to %d (ok=%v), want %d", c.named, level, got, ok, c.want)
			}
		}
	}
}

func TestInferiorRankStillGatesByTheRanksBeforeIt(t *testing.T) {
	table, err := parseSpellRanks([]byte(inferiorRankExcerpt))
	if err != nil {
		t.Fatal(err)
	}
	if got, ok := learnedIn(table["hunter"], 14322, 40); !ok || got != 14320 {
		t.Errorf("named 14322 at level 40 resolves to %d (ok=%v), want 14320", got, ok)
	}
	if _, ok := learnedIn(table["hunter"], 14322, 30); ok {
		t.Error("named 14322 at level 30 resolves, want unlearned")
	}
}

// A tier that holds an inferior id beside a sound one (Hunter's Mark rank 4:
// 14325 at 71 against the Forever copy 1213268 at 110) resolves to the sound id.
func TestInferiorIdInAMixedTierResolvesToItsSoundSibling(t *testing.T) {
	const mixed = `{"build":"x","classes":{"hunter":{"Hunter's Mark":[
  {"id":14324,"rank":3,"level":40},
  {"id":14325,"rank":4,"level":58,"inferior":true},
  {"id":1213268,"rank":4,"level":58}]}}}`
	table, err := parseSpellRanks([]byte(mixed))
	if err != nil {
		t.Fatal(err)
	}
	if got, ok := learnedIn(table["hunter"], 14325, 60); !ok || got != 1213268 {
		t.Errorf("named 14325 at level 60 resolves to %d (ok=%v), want 1213268", got, ok)
	}
}

// A book rank's id still belongs to its spell: a rotation authored at the
// book rank must resolve to the highest trainer rank the character has
// learned, not be left as written for an engine that has not registered it.
func TestABookRankIDResolvesToTheHighestTrainerRankLearned(t *testing.T) {
	if core.IncludeAQ {
		t.Skip("the book rank is a real rank while IncludeAQ is on")
	}
	table, err := parseSpellRanks([]byte(bookRankExcerpt))
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		level  int
		want   int32
		wantOK bool
	}{
		{60, 14322, true},
		{58, 14322, true},
		{50, 14321, true},
		{40, 0, false}, // nothing of the spell learned yet: the line is dropped
	} {
		got, ok := learnedIn(table["hunter"], 25296, tc.level)
		if ok != tc.wantOK || (ok && got != tc.want) {
			t.Errorf("book id 25296 at level %d: got (%d, %v), want (%d, %v)", tc.level, got, ok, tc.want, tc.wantOK)
		}
	}
}

// Primal Bite rank 2 (1238069) is learned at 50 by the pinned engine; the
// table's row says 36.
func TestRankLevelOverrideGatesARankAtTheEnginesLevel(t *testing.T) {
	const excerpt = `{"build":"x","classes":{"druid":{"Primal Bite":[
	  {"id":407995,"rank":1,"level":25},
	  {"id":1238069,"rank":2,"level":36}]}}}`
	table, err := parseSpellRanks([]byte(excerpt))
	if err != nil {
		t.Fatal(err)
	}
	if got, ok := learnedIn(table["druid"], 1238069, 49); !ok || got != 407995 {
		t.Errorf("Primal Bite at level 49: got (%d, %v), want rank 1 (407995, true)", got, ok)
	}
	if got, ok := learnedIn(table["druid"], 1238069, 50); !ok || got != 1238069 {
		t.Errorf("Primal Bite at level 50: got (%d, %v), want (1238069, true)", got, ok)
	}
}

// The single-tier overrides name ids the engine gates behind a real level.
func TestSingleTierOverridesGateTheirIDs(t *testing.T) {
	for _, tc := range []struct {
		class string
		id    int32
		level int
	}{
		{"druid", 29166, 40},    // Innervate
		{"paladin", 407632, 40}, // Hammer of the Righteous
		{"warlock", 603, 60},    // Bane of Doom
		{"warrior", 1680, 36},   // Whirlwind
		{"warrior", 2458, 30},   // Berserker Stance
	} {
		if _, ok := LearnedSpellID(tc.class, tc.id, tc.level-1); ok {
			t.Errorf("%s %d is learned at level %d, not %d", tc.class, tc.id, tc.level, tc.level-1)
		}
		if got, ok := LearnedSpellID(tc.class, tc.id, tc.level); !ok || got != tc.id {
			t.Errorf("%s %d at level %d: got (%d, %v)", tc.class, tc.id, tc.level, got, ok)
		}
	}
}

// TestARankAboveTheLevelCapDoesNotSetTheTiersLevel: spellranks.json lists
// Arcane Blast rank 3 twice, the NPC copy 42896 at level 76 first and the
// player's 1239697 at level 40 (SpellLevels 148351) second. The tier
// trusted the first row, so a level 40 to 49 mage "learned" rank 3 at 76
// and cast rank 2 (1239696, 78 base points against rank 3's 101).
func TestARankAboveTheLevelCapDoesNotSetTheTiersLevel(t *testing.T) {
	table, err := parseSpellRanks([]byte(`{"build": "test", "classes": {"mage": {"Arcane Blast": [
		{"id": 400574, "level": 20, "rank": 1},
		{"id": 1239696, "level": 30, "rank": 2},
		{"id": 42896, "level": 76, "rank": 3},
		{"id": 1239697, "level": 40, "rank": 3},
		{"id": 1239699, "level": 50, "rank": 4},
		{"id": 1239700, "level": 60, "rank": 5}
	]}}}`))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		level int
		want  int32
	}{{39, 1239696}, {40, 1239697}, {49, 1239697}, {50, 1239699}} {
		if got, ok := learnedIn(table["mage"], 1239700, c.level); !ok || got != c.want {
			t.Errorf("Arcane Blast at level %d = (%d, %v), want (%d, true)", c.level, got, ok, c.want)
		}
	}
	if got, ok := LearnedSpellID("mage", 1239700, 45); !ok || got != 1239697 {
		t.Errorf(`embedded table: LearnedSpellID("mage", 1239700, 45) = (%d, %v), want (1239697, true)`, got, ok)
	}
}

// Mutilate rank 1 is 1310707, SpellLevels BaseLevel 30 and a trainable
// SkillLineAbility row (AcquireMethod 0). The tier's other id, 1329, is
// Era's: no SkillLineAbility row and no BaseLevel, but listed first at 40.
func TestMutilateRankOneIsLearnedAtThirty(t *testing.T) {
	if _, ok := LearnedSpellID("rogue", 1241584, 29); ok {
		t.Error("Mutilate is learned at 29")
	}
	if got, ok := LearnedSpellID("rogue", 1241584, 30); !ok || got != 1310707 {
		t.Errorf("Mutilate at 30: got (%d, %v), want (1310707, true)", got, ok)
	}
	if got, _ := LearnedSpellID("rogue", 1241584, 40); got != 399956 {
		t.Errorf("Mutilate at 40: got %d, want 399956", got)
	}
}

// Shred rank 1 (5221) is learned at 20: the client row, Wowhead Forever's
// "Requires level 20" and Blizzard's 15 October beta notes agree, and the
// engine registers it there since lane/r8-hybrids.
func TestShredRankOneIsLearnedAtTwenty(t *testing.T) {
	if got, ok := LearnedSpellID("druid", 9830, 20); !ok || got != 5221 {
		t.Errorf("Shred at level 20: got (%d, %v), want (5221, true)", got, ok)
	}
	if _, ok := LearnedSpellID("druid", 9830, 19); ok {
		t.Error("Shred is learned at level 19")
	}
}

// An authored rank below the top is a deliberate downrank: the rewrite
// never lifts it. Below the level that rank is learned at, the highest
// rank learned under it stands in.
func TestLearnedUpToKeepsAnAuthoredDownrank(t *testing.T) {
	hunter := spellRanksExcerpt(t)["hunter"]
	cases := []struct {
		name      string
		id        int32
		level     int
		wantID    int32
		wantKnown bool
	}{
		{"rank 5 at 60 stays rank 5", 13552, 60, 13552, true},
		{"rank 5 at 50 stays rank 5, not rank 7", 13552, 50, 13552, true},
		{"rank 5 below its own level falls to rank 4", 13552, 30, 13551, true},
		{"the top rank still resolves to the highest learned", 25295, 38, 13552, true},
		{"a Forever copy of rank 5 stays in its band", 425733, 50, 425733, true},
		{"nothing learned yet", 13552, 3, 0, false},
		{"an untracked id is untouched", 999999, 20, 999999, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotID, gotOK := learnedIn(hunter, c.id, c.level)
			if gotID != c.wantID || gotOK != c.wantKnown {
				t.Errorf("learnedIn(%d, %d) = (%d, %v), want (%d, %v)", c.id, c.level, gotID, gotOK, c.wantID, c.wantKnown)
			}
		})
	}
}

// A per-band rank choice names a rank number of the authored spell's chain.
func TestRankOfResolvesAChosenRank(t *testing.T) {
	hunter := spellRanksExcerpt(t)["hunter"]
	cases := []struct {
		name      string
		id        int32
		rank      int
		level     int
		wantID    int32
		wantKnown bool
	}{
		{"rank 6 of the top rank at 50", 25295, 6, 50, 13553, true},
		{"rank 6 is not learned at 40", 25295, 6, 40, 0, false},
		{"a rank the chain does not have", 25295, 12, 60, 0, false},
		{"a Forever id picks the Forever copy", 425737, 6, 50, 425734, true},
		{"an untracked id has no ranks", 999999, 2, 60, 0, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotID, gotOK := rankOfIn(hunter, c.id, c.rank, c.level)
			if gotID != c.wantID || gotOK != c.wantKnown {
				t.Errorf("rankOfIn(%d, %d, %d) = (%d, %v), want (%d, %v)", c.id, c.rank, c.level, gotID, gotOK, c.wantID, c.wantKnown)
			}
		})
	}
}
