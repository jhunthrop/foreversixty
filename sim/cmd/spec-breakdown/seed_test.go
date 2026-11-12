package main

import (
	"encoding/json"
	"testing"
)

// breakdownDPS runs the whole breakdown path (flags to engine to report)
// and returns the DPS it printed.
func breakdownDPS(t *testing.T, o options) float64 {
	t.Helper()
	o.jsonOut = true
	out, err := captureStdout(t, func() error { return run(o) })
	if err != nil {
		t.Fatal(err)
	}
	var got report
	if err := json.Unmarshal([]byte(out), &got); err != nil {
		t.Fatalf("not JSON: %v\n%s", err, out)
	}
	return got.DPS
}

// The 2026-11-12 shamans lane found -seed 50001 reproducing seed 1 to the
// decimal, which made every "mean of two seeds" in the audits one seed.
// Two seeds must give two answers; one seed twice must give one.
func TestSeedReachesTheEngine(t *testing.T) {
	cases := []options{
		realOptions(testSpec),
		{repoRoot: realRoot, spec: "shaman-elemental", preset: "bare", band: 30, level: 39, faction: "horde", race: "tauren"},
	}
	for _, base := range cases {
		t.Run(base.spec, func(t *testing.T) {
			base.iterations = 40
			base.seed = 1
			first := breakdownDPS(t, base)
			again := breakdownDPS(t, base)
			if first != again {
				t.Fatalf("seed 1 twice gave %v and %v", first, again)
			}
			other := base
			other.seed = 50001
			if second := breakdownDPS(t, other); second == first {
				t.Fatalf("seeds 1 and 50001 both gave %v: the seed never reached the engine", first)
			}
		})
	}
}
