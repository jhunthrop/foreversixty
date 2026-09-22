package rating

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func TestCuratedTablesForSelectsBySpecAndRole(t *testing.T) {
	row := summary.RosterRow{GUID: "g1", Name: "Baelgrim", Class: "Warrior", Spec: "Protection", Role: "tank"}
	tables := curatedTablesFor(row, 667) // Shazzrah, a Molten Core boss with a curated table
	if tables.Mechanics == nil {
		t.Error("expected a mechanics table for encounter 667")
	}
	if tables.Utility == nil || tables.Utility.Spec != "warrior-protection" {
		t.Errorf("expected the warrior-protection utility table, got %+v", tables.Utility)
	}
	if tables.Consumables == nil {
		t.Error("expected the tank consumable catalogue row")
	}
}

func TestCuratedTablesForHandlesNoEncounterTable(t *testing.T) {
	row := summary.RosterRow{GUID: "g1", Name: "Nobody", Class: "Mage", Spec: "Frost", Role: "dps"}
	tables := curatedTablesFor(row, 99999999) // no such encounter
	if tables.Mechanics != nil {
		t.Error("expected no mechanics table for an uncurated encounter")
	}
	if tables.Utility == nil {
		t.Error("expected the mage-frost utility table regardless of the encounter")
	}
}

func TestSpecSlugMatchesEveryClassSpecPattern(t *testing.T) {
	cases := []struct{ class, spec, want string }{
		{"Warrior", "Protection", "warrior-protection"},
		{"Death Knight", "Blood", "death-knight-blood"},
		{"Priest", "Holy", "priest-holy"},
	}
	for _, c := range cases {
		if got := specSlug(c.class, c.spec); got != c.want {
			t.Errorf("specSlug(%q, %q) = %q, want %q", c.class, c.spec, got, c.want)
		}
	}
}
