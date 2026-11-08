package guilds

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

// onyxiaItem is one item the loot table drops in Onyxia's Lair.
func onyxiaItem(t *testing.T) int {
	t.Helper()
	for _, b := range lootBosses {
		if b.RaidKey == raidZoneOny && len(b.Items) > 0 {
			return b.Items[0].ItemID
		}
	}
	t.Fatal("the loot table has no Onyxia item")
	return 0
}

// The item sheet's next raid: the guild's next raid in the zone that drops the item, not over
// and not cancelled; null for an item no table raid drops or with none planned.
func TestItemNeedsNamesTheNextRaidForTheItemsZone(t *testing.T) {
	w := newRaidWorld(t, true)
	item := onyxiaItem(t)
	read := func(itemID int) map[string]json.RawMessage {
		t.Helper()
		w.as(w.member)
		res := w.h.do(http.MethodGet, w.path("/items/%d/needs", itemID), "")
		if res.StatusCode != http.StatusOK {
			t.Fatalf("needs = %d: %s", res.StatusCode, readBody(res))
		}
		var got map[string]json.RawMessage
		w.h.data(res, &got)
		return got
	}
	if got := read(item); string(got["next_raid"]) != "null" {
		t.Fatalf("none planned: %s", got["next_raid"])
	}
	later := w.createRaid(t)
	soonAt := time.Now().Add(2 * 24 * time.Hour).UTC().Truncate(time.Minute)
	soon := w.createRaids(t, raidBody(soonAt.Format(time.RFC3339), 0))[0]
	// A cancelled raid sooner still is not the next raid.
	cancelled := w.createRaids(t, raidBody(time.Now().Add(24*time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339), 0))[0]
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/cancel", cancelled.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("cancel = %d: %s", res.StatusCode, readBody(res))
	}

	var next ItemNextRaid
	if err := json.Unmarshal(read(item)["next_raid"], &next); err != nil {
		t.Fatal(err)
	}
	if next.RaidID != soon.ID || next.ZoneKey != raidZoneOny || !next.StartsAt.Equal(soonAt) {
		t.Fatalf("next raid %+v, want %d at %s (later %d)", next, soon.ID, soonAt, later.ID)
	}
	if got := read(999999); string(got["next_raid"]) != "null" {
		t.Fatalf("an item no raid drops: %s", got["next_raid"])
	}
}
