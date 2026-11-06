package main

import (
	"reflect"
	"testing"
)

func TestTrinketRankingIsTheTournamentAgainstThePick(t *testing.T) {
	a := fixtureItem(1, "A", fixtureTrinket, 5, src("raid", "MC"))
	b := fixtureItem(2, "B", fixtureTrinket, 4, src("dungeon", "UBRS"))
	c := fixtureItem(3, "C", fixtureTrinket, 3, src("rep", "AD"))
	tour := []tournamentEntry{{Item: b, DPS: 205.004}, {Item: a, DPS: 200}, {Item: c, DPS: 190.126}}
	// The pick is A (a modelled trinket kept over B within the margin): B's gap is positive.
	got := trinketRanking(slotPick{Item: &a, Tournament: tour})
	want := []rankingRow{{ItemID: 2, SourceKind: "dungeon", DPSDelta: 5}, {ItemID: 3, SourceKind: "rep", DPSDelta: -9.87}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ranking = %+v, want %+v", got, want)
	}
	other := fixtureItem(9, "Swapped in", fixtureTrinket, 1, src("vendor", "V"))
	if got := trinketRanking(slotPick{Item: &other, Tournament: tour}); got != nil {
		t.Errorf("a pick the tournament never measured published %+v", got)
	}
	long := make([]tournamentEntry, 0, 12)
	for i := 0; i < 12; i++ {
		long = append(long, tournamentEntry{Item: fixtureItem(100+i, "T", fixtureTrinket, 1, src("vendor", "V")), DPS: float64(100 - i)})
	}
	if got := trinketRanking(slotPick{Item: &long[0].Item, Tournament: long}); len(got) != trinketRankingLimit {
		t.Errorf("ranking has %d rows, want %d", len(got), trinketRankingLimit)
	}
}
