package request

import (
	"reflect"
	"testing"
)

const activeBuildDir = "../../data/builds/1.60.1.70790"

// The legal races come from the client's race-class combinations, Forever's
// added ones (Skyborne, Human hunter, Troll mage, Gnome shaman) included.
func TestLegalRacesReadTheClientCombinationsPerFaction(t *testing.T) {
	cases := []struct {
		class, faction string
		want           []string
	}{
		{"warrior", "horde", []string{"orc", "undead", "tauren", "troll", "windshaper-skyborne"}},
		{"druid", "alliance", []string{"night-elf", "high-order-skyborne"}},
		{"hunter", "alliance", []string{"human", "dwarf", "night-elf", "high-order-skyborne"}},
		{"shaman", "alliance", []string{"dwarf"}},
		{"paladin", "horde", []string{"undead"}},
	}
	for _, c := range cases {
		got, err := LegalRaces(activeBuildDir, c.class, c.faction)
		if err != nil {
			t.Fatalf("LegalRaces(%s, %s): %v", c.class, c.faction, err)
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("LegalRaces(%s, %s) = %v, want %v", c.class, c.faction, got, c.want)
		}
	}
}
