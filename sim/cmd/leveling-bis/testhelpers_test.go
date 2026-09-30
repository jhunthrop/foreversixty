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
	// ReferenceDPSPerPoint is RunWeights' second return - this lane's
	// brief, item 2 (bandReport.ReferenceDPSPerPoint's own source).
	ReferenceDPSPerPoint float64
	// FailGear, if non-empty, makes RunPlainDPS return an error for
	// exactly that gear fingerprint - how tests exercise this
	// command's per-candidate/per-slot error tolerance.
	FailGear string

	// StdErrByGear/DefaultStdErr/StdErrFunc are RunPlainDPSWithError's
	// own standard-error return - the DPSByGear/DefaultDPS/DPSFunc
	// trio, one level up (bis-ranker-integrity-11's brief, item 2).
	// Every field defaults to 0 ("no measurement noise"), so a test
	// that never sets one gets an always-significant gain, the same
	// as this command's real 0-error fixtures before this method
	// existed.
	StdErrByGear  map[string]float64
	DefaultStdErr float64
	StdErrFunc    func(req api.SimRequest) float64
	// FailWeights, if true, makes RunWeights return an error.
	FailWeights bool

	// WeightsFunc, when set, computes RunWeights' return value
	// directly from the full request rather than
	// WeightsResult/ReferenceDPSPerPoint - the only way a test can vary
	// the sweep's own result by req.Character.Level (which band) or
	// req.Iterations (main.go's own retry runs at
	// weightsRetryIterationsFactor iterations), the same shape DPSFunc
	// already gives RunPlainDPS below.
	WeightsFunc func(req api.SimRequest) (map[string]api.StatWeight, float64, error)

	// WeightsIterationsSeen records req.Iterations for every RunWeights
	// call, in order - so a test can assert main.go's own retry guard
	// (bis-ranker-integrity-11's brief, item 1) actually re-ran at
	// weightsRetryIterationsFactor iterations, not just that its log
	// line claims to have.
	WeightsIterationsSeen []int

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

// RunPlainDPSWithError delegates its mean to RunPlainDPS (so every
// existing test's DPSByGear/DefaultDPS/DPSFunc/FailGear/Calls/
// TalentsSeen behaviour is exercised identically whichever method a
// call site uses), then adds a standard error from StdErrFunc/
// StdErrByGear/DefaultStdErr - defaulting to 0 (a "no measurement
// noise at all" fixture) for every test that does not set one, so
// this lane's brief, item 2's own significance check (trinkets.go)
// can be driven deterministically by a test without depending on the
// real engine's own noise.
func (f *fakeEngine) RunPlainDPSWithError(req api.SimRequest) (float64, float64, error) {
	mean, err := f.RunPlainDPS(req)
	if err != nil {
		return 0, 0, err
	}
	if f.StdErrFunc != nil {
		return mean, f.StdErrFunc(req), nil
	}
	key := gearKey(req.Character.Gear)
	if stdErr, ok := f.StdErrByGear[key]; ok {
		return mean, stdErr, nil
	}
	return mean, f.DefaultStdErr, nil
}

func (f *fakeEngine) RunWeights(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
	f.WeightsIterationsSeen = append(f.WeightsIterationsSeen, req.Iterations)
	if f.FailWeights {
		return nil, 0, fmt.Errorf("fakeEngine: forced weights failure")
	}
	if f.WeightsFunc != nil {
		return f.WeightsFunc(req)
	}
	return f.WeightsResult, f.ReferenceDPSPerPoint, nil
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
