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

// Both loaders carry the file's generated_at, the band's own scalar fields and a slot's flags.
func TestLoadersCarryTheFileStampAndTheBandFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "b1", "bis")
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatal(err)
	}
	raw := `{"generated_at": "2026-10-29T13:34:26Z", "bands": [{"band": 60, "faction": "alliance", "preset": "raid",
	 "score_unit": "tank_score", "reference_dps_per_point": 0.25,
	 "slots": [{"slot": "head", "item_id": 3, "score": 96.5, "effect_unmodelled": true},
	           {"slot": "off_hand", "empty_reason": "two_hand_equipped"}]}]}`
	if err := os.WriteFile(filepath.Join(path, "warrior-fury.json"), []byte(raw), 0o600); err != nil {
		t.Fatal(err)
	}
	plain, ok1, err1 := LoadBand(dir, "b1", "warrior-fury", 60, "alliance")
	preset, ok2, err2 := LoadPresetBand(dir, "b1", "warrior-fury", 60, "alliance", RaidPreset)
	for _, b := range []Band{plain, preset} {
		if !ok1 || !ok2 || err1 != nil || err2 != nil || b.GeneratedAt != "2026-10-29T13:34:26Z" || b.ScoreUnit != TankScoreUnit ||
			b.ReferenceDpsPerPoint == nil || *b.ReferenceDpsPerPoint != 0.25 {
			t.Fatalf("band = %+v (%v %v %v %v)", b, ok1, ok2, err1, err2)
		}
		head, off := b.BySlot()["head"], b.BySlot()["off_hand"]
		if head.Score == nil || *head.Score != 96.5 || !head.EffectUnmodelled || off.EmptyReason != EmptyTwoHand {
			t.Fatalf("slots = %+v %+v", head, off)
		}
	}
}
