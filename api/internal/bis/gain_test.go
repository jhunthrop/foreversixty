package bis

import (
	"path/filepath"
	"runtime"
	"testing"
)

func pick() Slot {
	return Slot{
		Slot: "chest", ItemID: 100, ItemName: "Pick",
		Alternatives: []Alternative{
			// A real file's own sign convention: dps_delta is the alternative's own
			// deviation from the pick, negative for a worse runner-up (confirmed
			// against data/builds/1.60.1.70009/bis/hunter-marksmanship.json's own
			// alternatives, e.g. -1.05) - switching from it to the pick gains
			// -dps_delta, a positive number.
			{ItemID: 200, ItemName: "Runner-up", DpsDelta: -5},
			{ItemID: 300, ItemName: "Tie", DpsDelta: 0},
		},
	}
}

func TestVerdictForWornIsThePick(t *testing.T) {
	v := VerdictFor(pick(), 100, true)
	if v.Upgrade || v.GainDps != nil || v.NotSimChecked {
		t.Fatalf("verdict = %+v, want no upgrade", v)
	}
}

func TestVerdictForWornIsAZeroDeltaAlternative(t *testing.T) {
	v := VerdictFor(pick(), 300, true)
	if v.Upgrade || v.GainDps != nil || v.NotSimChecked {
		t.Fatalf("verdict = %+v, want a tie treated as no upgrade", v)
	}
}

func TestVerdictForWornIsAListedAlternative(t *testing.T) {
	v := VerdictFor(pick(), 200, true)
	if !v.Upgrade || v.NotSimChecked {
		t.Fatalf("verdict = %+v, want an upgrade, sim-checked", v)
	}
	if v.GainDps == nil || *v.GainDps != 5 {
		t.Fatalf("gainDps = %v, want 5 (-DpsDelta of -5)", v.GainDps)
	}
}

func TestVerdictForWornIsUnknown(t *testing.T) {
	v := VerdictFor(pick(), 999, true)
	if !v.Upgrade || !v.NotSimChecked || v.GainDps != nil {
		t.Fatalf("verdict = %+v, want upgrade+notSimChecked, no gain", v)
	}
}

func TestVerdictForNothingEquipped(t *testing.T) {
	v := VerdictFor(pick(), 0, false)
	if !v.Upgrade || !v.NotSimChecked {
		t.Fatalf("verdict = %+v, want upgrade+notSimChecked for an empty slot", v)
	}
}

func TestGapForSumsAcrossSlotsAndSkipsSlotsWithNoPick(t *testing.T) {
	band := Band{Slots: []Slot{
		pick(),                    // worn is the listed alternative below: +5 gain
		{Slot: "legs", ItemID: 0}, // no known source: skipped outright
		{Slot: "feet", ItemID: 50, Alternatives: []Alternative{}}, // worn unknown: not sim-checked
	}}
	gear := map[string]int{"chest": 200, "feet": 999}
	gap := GapFor(band, gear)
	if gap.Upgrades != 2 {
		t.Fatalf("upgrades = %d, want 2", gap.Upgrades)
	}
	if gap.GainDps != 5 {
		t.Fatalf("gainDps = %v, want 5", gap.GainDps)
	}
	if gap.NotSimChecked != 1 {
		t.Fatalf("notSimChecked = %d, want 1", gap.NotSimChecked)
	}
}

func TestFileSlugForMatchesTheRealCatalogue(t *testing.T) {
	slug, ok := FileSlugFor("hunter", "Marksmanship")
	if !ok || slug != "hunter-marksmanship" {
		t.Fatalf("slug = %q ok=%v, want hunter-marksmanship/true", slug, ok)
	}
	if _, ok := FileSlugFor("hunter", "NotASpec"); ok {
		t.Fatal("ok = true for an unknown spec, want false")
	}
}

// repoDataDir resolves ../../../../data from this test file's own location, the same
// pattern api/cmd/seedguild/gear.go's repoRoot uses, so this test needs no environment
// variable or working-directory assumption.
func repoDataDir(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("could not resolve this test file's own path")
	}
	return filepath.Join(filepath.Dir(thisFile), "..", "..", "..", "data", "builds")
}

func TestLoadBandReadsTheRealHunterMarksmanshipFile(t *testing.T) {
	band, ok, err := LoadBand(repoDataDir(t), "1.60.1.70009", "hunter-marksmanship", 60, "alliance")
	if err != nil {
		t.Fatalf("LoadBand: %v", err)
	}
	if !ok {
		t.Fatal("ok = false, want true for a real spec/band/faction combination")
	}
	if len(band.Slots) == 0 {
		t.Fatal("band has no slots, want the real file's own slot list")
	}
}

func TestLoadBandMissingSpecFileIsNotAnError(t *testing.T) {
	band, ok, err := LoadBand(repoDataDir(t), "1.60.1.70009", "warrior-protection-tank-nonexistent", 60, "alliance")
	if err != nil {
		t.Fatalf("LoadBand: %v, want no error for a missing file", err)
	}
	if ok {
		t.Fatalf("ok = true, want false; band = %+v", band)
	}
}
