package coach

// F5 fixtures: a tank casting one Sunder Armor rank, and a Holy priest whose one heal
// rank overhealed at a chosen percentage.

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	sunderRank4ID    = 11596
	healTotalFixture = 10_000
)

// protSunderRank is the Protection warrior at level, casting casts of the Sunder Armor
// rank whose spell id is the trainable table's rank number.
func protSunderRank(t *testing.T, level int64, rank int, casts int) Input {
	t.Helper()
	in := protInput(t)
	in.Fight = withLevel(in.Fight, in.Row.GUID, level)
	in.Row = rosterRowOf(t, in.Fight, in.Row.GUID)
	var id int64
	for _, r := range in.Tables.Trainable["Sunder Armor"] {
		if r.Rank == rank {
			id = r.ID
		}
	}
	if id == 0 {
		t.Fatalf("no Sunder Armor rank %d", rank)
	}
	in.Fight = withCasts(withoutCasts(in.Fight, in.Row.GUID), in.Row.GUID, id, "Sunder Armor", seqEvery(1_000, 2_000, casts)...)
	return in
}

// withDownrank is in with family listed as a deliberate downrank.
func withDownrank(in Input, family string) Input {
	in.Tables.Rules.Downranks = append(slices.Clone(in.Tables.Rules.Downranks),
		RuleEntry{Family: family, SpellIDs: []int64{sunderRank4ID}, Verified: fixtureVerified})
	return in
}

// holyOverheal is the Holy priest casting casts of spellID, which overhealed overhealPct
// percent of healTotalFixture, with refs as the comparison.
func holyOverheal(t *testing.T, spellID int64, casts int, overhealPct int64, refs Refs) Input {
	t.Helper()
	in := withPool(holyInput(t), refs)
	guid := in.Row.GUID
	in.Fight = withCasts(in.Fight, guid, spellID, "Greater Heal", seqEvery(1_000, 2_000, casts)...)
	total := int64(healTotalFixture)
	ability := summary.Ability{SpellID: spellID, Name: "Greater Heal", Total: total,
		Overheal: total * overhealPct / 100, Effective: total - total*overhealPct/100, Hits: int64(casts)}
	in.Fight.Healing = append(slices.DeleteFunc(slices.Clone(in.Fight.Healing), func(a summary.Actor) bool {
		return a.GUID == guid
	}), summary.Actor{GUID: guid, Total: total, Abilities: []summary.Ability{ability}})
	return in
}
