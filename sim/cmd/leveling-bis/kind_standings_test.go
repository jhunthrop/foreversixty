package main

import (
	"fmt"
	"testing"
)

// TestARepPickNamesTheStandingOfEveryOwnKind: Protector's Band (P2 review,
// 2026-11-09) is a rep item the Silverwing Sentinels quartermaster also
// sells, so its kinds are rep and vendor and both need Honored. The page's
// leveling default caps reputation at Friendly; with only the rep kind's
// standing on the row it took the rep kind out and kept the item through
// its vendor kind. The row names each own kind's standing.
func TestARepPickNamesTheStandingOfEveryOwnKind(t *testing.T) {
	band := fixtureItem(20439, "Protector's Band", fixtureFinger, 10,
		itemSource{Kind: "rep", Label: "Silverwing Sentinels", Standing: "honored"},
		itemSource{Kind: "vendor", Label: "Illiyana Moonblaze", Standing: "honored"})
	plain := fixtureItem(1156, "Plain Ring", fixtureFinger, 5,
		src("dungeon", "Deadmines"), src("vendor", "Some Vendor"))

	pk := slotPick{Item: &band}
	row, _ := sourceKindRow(slotRow{Slot: "finger1", ItemID: band.ID}, pk, []scored{band, plain}, false, nil, 0.5)
	if fmt.Sprint(row.KindStandings) != "map[rep:honored vendor:honored]" {
		t.Errorf("kind_standings = %v, want rep and vendor at honored", row.KindStandings)
	}

	pk = slotPick{Item: &plain}
	row, _ = sourceKindRow(slotRow{Slot: "finger1", ItemID: plain.ID}, pk, []scored{plain}, false, nil, 0.5)
	if row.KindStandings != nil {
		t.Errorf("a pick no standing gates names kind_standings %v", row.KindStandings)
	}
}
