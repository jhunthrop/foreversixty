package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// fakeEngine is the engineRunner test double every test that would
// otherwise reach for the real engine (weights/verify runs - the
// slowest, least deterministic part of this command) uses instead.
// It never touches wowsims-classic; it answers from DPSByGear/
// WeightsResult, which the test sets up, so runSpec/verifyBand/
// rankTrinketSlot can be exercised without the real 13-minute ranking
// run this lane's brief explicitly forbids.
type fakeEngine struct {
	// DPSByGear maps a gear fingerprint (gearKey below) to the DPS
	// RunPlainDPS should return for that exact gear list. A fingerprint
	// this map does not carry falls back to DefaultDPS.
	DPSByGear map[string]float64
	// DefaultDPS is returned for any gear fingerprint DPSByGear does
	// not name.
	DefaultDPS float64
	// WeightsResult is what RunWeights always returns.
	WeightsResult map[string]api.StatWeight
	// FailGear, if non-empty, makes RunPlainDPS return an error for
	// exactly that gear fingerprint - how tests exercise this
	// command's per-candidate/per-slot error tolerance.
	FailGear string
	// FailWeights, if true, makes RunWeights return an error.
	FailWeights bool

	// DPSFunc, when set, computes RunPlainDPS's return value directly
	// from the full request rather than DPSByGear/DefaultDPS - the only
	// way a test can make the fake DPS depend on req.Character.Talents,
	// since DPSByGear/gearKey fingerprint gear alone. Used by
	// TestRunSpecPublishesDifferentSetDPSForDifferentTalentStrings
	// (main_test.go) to prove Talents actually reaches the engine
	// (this lane's brief, defect 1).
	DPSFunc func(req api.SimRequest) (float64, error)

	// Calls records every RunPlainDPS gear fingerprint seen, in order,
	// so a test can assert what this command actually tried (e.g. "did
	// it drop the pair-mate", "did it skip the failing candidate and
	// keep going").
	Calls []string
	// TalentsSeen records req.Character.Talents for every RunPlainDPS
	// call, parallel to Calls - so a test can assert every site that
	// builds a plain-DPS request (verifyBand, applySwaps,
	// rankTrinketSlot, rankSlotWithEffects, trySetCompletion) carried
	// the band's own talent string, not the zero value (this lane's
	// brief, defect 1).
	TalentsSeen []string
}

// gearKey fingerprints a gear list as a stable, comparable string:
// slot=itemID pairs, sorted by slot so call order never matters.
func gearKey(gear []api.GearSlot) string {
	byID := make(map[string]int, len(gear))
	var slots []string
	for _, g := range gear {
		byID[g.Slot] = g.ItemID
		slots = append(slots, g.Slot)
	}
	// slotOrder is this package's own canonical slot order - reusing it
	// keeps the fingerprint deterministic without importing sort just
	// for this helper.
	key := ""
	for _, s := range slotOrder {
		if id, ok := byID[s]; ok {
			key += fmt.Sprintf("%s=%d;", s, id)
		}
	}
	return key
}

func (f *fakeEngine) RunPlainDPS(req api.SimRequest) (float64, error) {
	key := gearKey(req.Character.Gear)
	f.Calls = append(f.Calls, key)
	f.TalentsSeen = append(f.TalentsSeen, req.Character.Talents)
	if f.FailGear != "" && key == f.FailGear {
		return 0, fmt.Errorf("fakeEngine: forced failure for gear %s", key)
	}
	if f.DPSFunc != nil {
		return f.DPSFunc(req)
	}
	if dps, ok := f.DPSByGear[key]; ok {
		return dps, nil
	}
	return f.DefaultDPS, nil
}

func (f *fakeEngine) RunWeights(req api.SimRequest) (map[string]api.StatWeight, error) {
	if f.FailWeights {
		return nil, fmt.Errorf("fakeEngine: forced weights failure")
	}
	return f.WeightsResult, nil
}

// writeFile creates path (and its parent directories) with contents -
// a small shared helper for tests that build their own scratch
// repoRoot-shaped fixtures under t.TempDir() rather than the committed
// testdata/reporoot fixture (used when a test needs a layout the
// committed fixture does not have, such as a malformed file).
func writeFile(t *testing.T, path, contents string) error {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(contents), 0o644)
}
