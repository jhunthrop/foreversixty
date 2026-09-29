package request

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/specs"
	engine "github.com/wowsims/classic/sim"
)

// TestLadderIsDeterministic guards the failure mode the
// night-determinism audit went looking for: an unrelated code change in
// one package (a spell registered for a different spec, a map ranged
// somewhere in registration or attack-table setup) reordering how the
// ladder's fixed RandomSeed is consumed, which shifts DPS for a spec
// that never itself changed. druid-feral and druid-balance are the pair
// the audit's own symptom named (Moonfire/Wrath added to Feral moved
// Balance's ladder DPS 127.3 -> 123.1, byte-stable per binary, reverting
// the one file restored it exactly) -- so this test runs Feral
// immediately before Balance, twice, and fails on the first spec/level
// whose signature differs between the two passes. A plain "run the same
// spec twice" would not catch cross-spec ordering leaks; running Feral
// first both times, back to back, gives that leak a chance to surface
// without depending on any particular unrelated file being present to
// trigger it.
//
// As investigated 2026-09-28 (night-determinism lane): at engine pin
// 84575e84b / site 98b8a711 this test passes -- in-process repetition,
// this cross-spec ordering, 5 repeated `go test -count=1` process
// invocations, and `-race` were all byte-stable/clean. Nothing here
// found a live ordering bug to fix in the engine; this test exists so a
// regression is caught immediately rather than rediscovered by another
// audit.
func TestLadderIsDeterministic(t *testing.T) {
	registerEngine.Do(engine.RegisterAll)
	build := activeBuild(t)
	ranks, err := loadSpellRanks(repoRoot, build)
	if err != nil {
		t.Fatal(err)
	}

	firstFeral := ladderDeterminismSignature(t, build, "druid-feral", ranks)
	firstBalance := ladderDeterminismSignature(t, build, "druid-balance", ranks)
	secondFeral := ladderDeterminismSignature(t, build, "druid-feral", ranks)
	secondBalance := ladderDeterminismSignature(t, build, "druid-balance", ranks)

	if firstFeral != secondFeral {
		t.Errorf("druid-feral's ladder signature differs across two runs in the same process:\n--- run 1 ---\n%s--- run 2 ---\n%s", firstFeral, secondFeral)
	}
	if firstBalance != secondBalance {
		t.Errorf("druid-balance's ladder signature differs across two runs in the same process (druid-feral ran immediately before it both times):\n--- run 1 ---\n%s--- run 2 ---\n%s", firstBalance, secondBalance)
	}
}

// ladderDeterminismSignature runs one spec's full ladder and renders a
// signature at full float64 precision (DPS, distinct-cast count, the
// top-casts summary, and any unresolved ids) for every level -- deliberately
// NOT renderLadderGolden's markdown, whose DPS column is rounded to one
// decimal and would mask exactly the kind of least-significant-bit drift
// this test exists to catch.
func ladderDeterminismSignature(t *testing.T, build, specName string, ranks spellRanksFile) string {
	t.Helper()
	spec := ladderSpecByName(t, specName)
	curated, err := loadLadderCurated(repoRoot, spec.Spec)
	if err != nil {
		t.Fatal(err)
	}
	rows, _, _ := runLadderSpec(t, build, spec, curated, ranks)

	var b strings.Builder
	for _, r := range rows {
		fmt.Fprintf(&b, "%s level=%d dps=%.17g casts=%d top=%s unresolved=%s\n",
			spec.Spec, r.Level, r.DPS, r.DistinctCasts, r.TopCasts, strings.Join(r.Unresolved, ","))
	}
	return b.String()
}

// ladderSpecByName finds one spec by its full name (e.g. "druid-feral")
// in specs.All.
func ladderSpecByName(t *testing.T, name string) specs.Spec {
	t.Helper()
	for _, s := range specs.All {
		if s.Spec == name {
			return s
		}
	}
	t.Fatalf("no spec named %s", name)
	return specs.Spec{}
}
