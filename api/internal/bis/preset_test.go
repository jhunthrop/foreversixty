package bis

import (
	"os"
	"path/filepath"
	"testing"
)

const presetFixture = `{"bands": [
 {"band": 60, "faction": "horde", "slots": [{"slot": "head", "item_id": 1, "item_name": "Old Bare Helm"}]},
 {"band": 60, "faction": "alliance", "preset": "bare", "slots": [{"slot": "head", "item_id": 2, "item_name": "Bare Helm", "source": "Blacksmithing"}]},
 {"band": 60, "faction": "alliance", "preset": "raid", "slots": [{"slot": "head", "item_id": 3, "item_name": "Raid Helm", "source": "Onyxia",
   "alternatives": [{"item_id": 4, "item_name": "Second Helm", "source": "PvP rank 14 · Champion · Horde", "dps_delta": -15}]}]},
 {"band": 40, "faction": "alliance", "preset": "bare", "slots": [{"slot": "head", "item_id": 5, "item_name": "Level Helm"}]}
]}`

func writeBisFixture(t *testing.T, dir, build string) {
	t.Helper()
	path := filepath.Join(dir, build, "bis")
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(path, "warrior-fury.json"), []byte(presetFixture), 0o600); err != nil {
		t.Fatal(err)
	}
}

func TestLoadPresetBandMirrorsTheSite(t *testing.T) {
	dir := t.TempDir()
	writeBisFixture(t, dir, "b1")
	cases := []struct {
		name           string
		band           int
		faction        string
		preset         string
		wantID         int
		wantFound      bool
		wantSource     string
		wantAltSource  string
		wantAltDpsDiff float64
	}{
		{"raid preset at 60", 60, "alliance", RaidPreset, 3, true, "Onyxia", "PvP rank 14 · Champion · Horde", -15},
		{"bare preset at 60", 60, "alliance", BarePreset, 2, true, "Blacksmithing", "", 0},
		{"raid asked, only bare", 40, "alliance", RaidPreset, 5, true, "", "", 0},
		{"an entry with no preset is bare", 60, "horde", BarePreset, 1, true, "", "", 0},
		{"raid asked, no-preset entry", 60, "horde", RaidPreset, 1, true, "", "", 0},
		{"no band", 30, "alliance", BarePreset, 0, false, "", "", 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			band, found, err := LoadPresetBand(dir, "b1", "warrior-fury", c.band, c.faction, c.preset)
			if err != nil || found != c.wantFound {
				t.Fatalf("found %v err %v", found, err)
			}
			if !found {
				return
			}
			head := band.Slots[0]
			if head.ItemID != c.wantID || head.Source != c.wantSource {
				t.Fatalf("head = %+v", head)
			}
			if c.wantAltSource != "" && (head.Alternatives[0].Source != c.wantAltSource || head.Alternatives[0].DpsDelta != c.wantAltDpsDiff) {
				t.Fatalf("alternative = %+v", head.Alternatives[0])
			}
		})
	}
	if _, found, err := LoadPresetBand(dir, "b1", "", 60, "alliance", BarePreset); found || err != nil {
		t.Fatalf("empty slug = %v %v", found, err)
	}
}

func TestSourceBuildFallsBackToTheNewestRankedBuild(t *testing.T) {
	dir := t.TempDir()
	writeBisFixture(t, dir, "1.60.1.70009")
	writeBisFixture(t, dir, "1.60.1.70291")
	if err := os.MkdirAll(filepath.Join(dir, "1.60.1.70630"), 0o755); err != nil {
		t.Fatal(err)
	}
	if got := SourceBuild(dir, "1.60.1.70630"); got != "1.60.1.70291" {
		t.Fatalf("unranked active = %q", got)
	}
	if got := SourceBuild(dir, "1.60.1.70009"); got != "1.60.1.70009" {
		t.Fatalf("ranked active = %q", got)
	}
	if got := SourceBuild(dir, "9.9"); got != "9.9" {
		t.Fatalf("unknown active = %q", got)
	}
}
