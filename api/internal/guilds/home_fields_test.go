package guilds

import (
	"encoding/json"
	"net/http"
	"testing"
)

// The raid view names its zone's icon and carries answer_by, null while no guild stores a
// sign-up cutoff.
func TestRaidViewCarriesZoneIconAndAnswerBy(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	res := w.h.do(http.MethodGet, w.path("/raids/%d", raid.ID), "")
	var env struct {
		Data map[string]json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(readBody(res)), &env); err != nil {
		t.Fatal(err)
	}
	if string(env.Data["zone_icon"]) != `"inv_misc_head_dragon_black"` {
		t.Fatalf("zone_icon %s", env.Data["zone_icon"])
	}
	if raw, ok := env.Data["answer_by"]; !ok || string(raw) != "null" {
		t.Fatalf("answer_by %s (present %v)", raw, ok)
	}
	night, code := w.raidNightOf(t, w.member, raid.ID)
	if code != http.StatusOK || night.ZoneIcon != "inv_misc_head_dragon_black" {
		t.Fatalf("raid night %d %q", code, night.ZoneIcon)
	}
}

// A night's loot gain is the Loot tab's upgrade gain for the character's band: the band's pick
// over an empty slot gains, in the band's unit; an item no band names has none.
func TestNightLootGainIsTheLootTabsGain(t *testing.T) {
	s := realTreesStore(t)
	helm := furyHeadPick(t, s, 60)
	item, ok := lootItemByID(helm)
	if !ok {
		item = lootItem{ItemID: helm, Slot: "head", RequiredLevel: 1}
	}
	row := furyRow("us/pvp/gainer", "Gainer", 60, shirtOnly())
	gain, unit := s.nightLootGain(row, item)
	if gain != nil && (unit == nil || *unit != "DPS") {
		t.Fatalf("gain %v unit %v", gain, unit)
	}
	if g, u := s.nightLootGain(row, lootItem{ItemID: 1, Slot: "head"}); g != nil || u != nil {
		t.Fatalf("an unnamed item gains %v %v", g, u)
	}
}
