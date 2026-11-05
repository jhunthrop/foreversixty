package guilds

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
)

// The readiness board, the roster and /char read one band at 60: the raid preset the site and
// /char use (upgrades.PresetFor), never the file's first entry. Round 3: a Fire Mage in the bare
// picks read "No upgrades" on the board while /char listed raid upgrades.
func TestReadinessMeasuresAgainstTheSitesBand(t *testing.T) {
	s := realTreesStore(t)
	band, ok := s.loadBandFor("mage", "Fire", "alliance")
	if !ok {
		t.Fatal("no Fire Mage band at 60")
	}
	if want := upgrades.PresetFor(bisBandLevel); band.Preset != want {
		t.Fatalf("band preset = %q, want %q", band.Preset, want)
	}
	bare, ok, err := bis.LoadPresetBand(s.DataDir, s.BisBuild(), "mage-fire", bisBandLevel, "alliance", bis.BarePreset)
	if err != nil || !ok {
		t.Fatalf("bare band: %v %v", ok, err)
	}
	gear := map[string]int{}
	for _, sl := range bare.Slots {
		if sl.ItemID > 0 {
			gear[sl.Slot] = sl.ItemID
		}
	}
	level := 60
	row := RosterRow{className: "mage", specStr: "Fire", faction: "alliance", Consent: "gear", Level: &level, gear: gear}
	if cr := s.readinessOf(row, nil); cr.GearGap.Upgrades == 0 {
		t.Fatal("the bare picks read no upgrades against the raid band")
	}
	gap := s.rosterGearGap(row)
	cr := s.readinessOf(row, nil)
	if gap == nil || float64(*gap) < cr.GearGap.GainDps-0.5 || float64(*gap) > cr.GearGap.GainDps+0.5 {
		t.Fatalf("roster gap = %v, want the board's %.1f DPS rounded", gap, cr.GearGap.GainDps)
	}
	if healer := (RosterRow{className: "priest", specStr: "Holy", faction: "alliance", Consent: "gear", Level: &level, gear: gear}); s.rosterGearGap(healer) != nil {
		t.Fatal("a healer's roster row carries a DPS gap")
	}
}

// The roster's item level is the equipped average from the export's gear, rounded, the figure
// /char prints; with no gear the row keeps the log's.
func TestRosterItemLevelIsTheEquippedAverage(t *testing.T) {
	s := realTreesStore(t)
	// Helm of Wrath (76) and an empty rest: the mean of the one worn item.
	if got, ok := s.equippedItemLevel(RosterRow{className: "warrior", gear: map[string]int{"head": 16963}}); !ok || got != 76 {
		t.Fatalf("equipped item level = %d %v, want 76", got, ok)
	}
	if _, ok := s.equippedItemLevel(RosterRow{className: "warrior"}); ok {
		t.Fatal("no gear read an item level")
	}
	if _, ok := s.equippedItemLevel(RosterRow{className: "warrior", gear: map[string]int{"head": 16963, "chest": 999999999}}); ok {
		t.Fatal("an unknown item still gave a figure")
	}
}
