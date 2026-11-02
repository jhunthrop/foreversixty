package main

import (
	"reflect"
	"testing"
)

// TestDroppedToNoSourceNamesItemsWhosePvpSourceWasTheirOnlyOne: an item
// whose only source was a pvp row that dropNonRankRewards took off has no
// source left, and the ranker must say so rather than drop it in silence
// (round 8: Forever's PvP rewards need no rank, so a relabelled pvp source
// can now land on an item classic-db calls unranked).
func TestDroppedToNoSourceNamesItemsWhosePvpSourceWasTheirOnlyOne(t *testing.T) {
	idx := lootIndex{
		1: {{Kind: "pvp", Rank: 0}},
		2: {{Kind: "pvp", Rank: 7}, {Kind: "vendor", Rank: 7}},
		3: {{Kind: "pvp", Rank: 9}},
		4: {{Kind: "quest"}},
	}
	unranked := map[int]bool{1: true, 2: true}
	got := droppedToNoSource(idx, dropNonRankRewards(idx, unranked))
	if want := []int{1}; !reflect.DeepEqual(got, want) {
		t.Errorf("droppedToNoSource = %v, want %v (item 2 keeps its vendor row; 3 is ranked; 4 has no pvp source)", got, want)
	}
}
