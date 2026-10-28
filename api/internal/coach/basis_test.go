package coach

import "testing"

func TestComparison_Order(t *testing.T) {
	poolCPMOf := func(p Profile) (float64, bool) { return p.CastsPerMinOf("Bloodthirst") }
	simCPMOf := func(s SimRef) (float64, bool) { v, ok := s.CastsPerMin["Bloodthirst"]; return v, ok }
	sim := &SimRef{Band: 60, CastsPerMin: map[string]float64{"Bloodthirst": 10.2}}
	withSim := func(r Refs) Refs { r.Sim = sim; return r }
	base := furyInput(t)

	cases := []struct {
		name       string
		refs       Refs
		simAllowed bool
		wantRef    float64
		wantKind   BasisKind
		wantLabel  string
		wantN      int
		wantThin   string
	}{
		{name: "a pool of 25 is top logs", refs: poolCPM("Bloodthirst", 9.1, 25), simAllowed: true,
			wantRef: 9.1, wantKind: BasisTopLogs, wantLabel: "Top 25 Fury warriors on Onyxia (40): ", wantN: 25},
		{name: "a pool and a sim: the pool wins", refs: withSim(poolCPM("Bloodthirst", 9.1, 30)), simAllowed: true,
			wantRef: 9.1, wantKind: BasisTopLogs, wantLabel: "Top 25 Fury warriors on Onyxia (40): ", wantN: 25},
		{name: "no pool, a validated sim, allowed", refs: withSim(noPool(7)), simAllowed: true,
			wantRef: 10.2, wantKind: BasisSim, wantLabel: "The written Fury rotation, simulated at level 60: "},
		{name: "no pool, sim not allowed", refs: withSim(noPool(7)), simAllowed: false,
			wantThin: "not enough logs yet — 7 Fury kills of Onyxia (40) on the site, the coach needs 20."},
		{name: "no pool, no sim", refs: noPool(19), simAllowed: true,
			wantThin: "not enough logs yet — 19 Fury kills of Onyxia (40) on the site, the coach needs 20."},
		// Ruling A3: no finding, and a draft line, never the thin-pool line (it would be false).
		{name: "a pool without this measurement says so, never thin", refs: poolUptime(6673, 97, 25), simAllowed: false,
			wantThin: "Top Fury logs on Onyxia (40) don't record this yet."},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ref, b, thin := comparison(withPool(base, c.refs), poolCPMOf, simCPMOf, c.simAllowed)
			if thin != c.wantThin {
				t.Errorf("thin = %q, want %q", thin, c.wantThin)
			}
			if c.wantKind == "" {
				if b != nil {
					t.Fatalf("basis %+v, want none", b)
				}
				return
			}
			if b == nil || b.Kind != c.wantKind || b.Label != c.wantLabel || b.N != c.wantN || ref != c.wantRef {
				t.Fatalf("got %v %+v, want %v %s %q n=%d", ref, b, c.wantRef, c.wantKind, c.wantLabel, c.wantN)
			}
		})
	}
}
