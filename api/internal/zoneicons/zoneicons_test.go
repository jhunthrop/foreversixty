package zoneicons

import "testing"

func TestZoneIcons(t *testing.T) {
	if For("raid:onyxias-lair") != "inv_misc_head_dragon_black" || ForName("Onyxia's Lair") != "inv_misc_head_dragon_black" {
		t.Fatal("Onyxia's icon")
	}
	if For("dungeon:scholomance") != Generic || ForName("Scholomance") != Generic {
		t.Fatal("a zone with no icon reads as Generic")
	}
}
