package guilds

import "testing"

// A picked raid spec wins over the logs on the roster, as on the raid page; a withheld one
// (nil) or a key the catalogue does not name changes nothing.
func TestRosterRowTakesTheRaidSpec(t *testing.T) {
	frost, dps := "Frost", "dps"
	row := RosterRow{Class: strp("mage"), Spec: &frost, Role: &dps, specStr: frost, className: "mage"}
	if got := row.withRaidSpec(strp("mage-fire")); *got.Spec != "Fire" || got.specStr != "Fire" || *got.Role != "dps" {
		t.Fatalf("picked = %+v", got)
	}
	if got := row.withRaidSpec(nil); *got.Spec != "Frost" {
		t.Fatalf("withheld = %q", *got.Spec)
	}
	if got := row.withRaidSpec(strp("mage-nonsense")); *got.Spec != "Frost" {
		t.Fatalf("unknown key = %q", *got.Spec)
	}
	c := raidRosterChar{raidSpec: "mage-fire"}
	c.Class, c.Spec = "mage", "Frost"
	if c.specKey() != "mage-fire" {
		t.Fatalf("raid roster key = %q", c.specKey())
	}
	c.raidSpec = ""
	if c.specKey() != "mage-frost" {
		t.Fatalf("logged key = %q", c.specKey())
	}
}
