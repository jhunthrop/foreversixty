package coach

import "strconv"

// comparison follows §11's order for one measurement: top logs (the pool's median), then
// the sim (only when simAllowed: a validated DPS spec, and never for a family on
// WorstActions), else none. The first that exists wins. The basis label is the prefix the
// caller appends its value to (BasisTop, BasisSimLabel).
//
// thin is the "Couldn't check" line when nothing answers: ThinPool filled with the real
// count when there is no pool, or, when a pool exists but its median lacks this
// measurement, PoolLacksMeasure (never ThinPool, which would be false); ""
// when one answers. Either way the caller makes no finding.
func comparison(in Input, pool func(Profile) (float64, bool), sim func(SimRef) (float64, bool), simAllowed bool) (ref float64, b *Basis, thin string) {
	if p := in.Refs.Pool; p != nil && pool != nil {
		if v, ok := pool(p.Median); ok {
			return v, &Basis{Label: Fill(BasisTop, topFields(in, p.N)), Kind: BasisTopLogs, N: p.N}, ""
		}
	}
	if s := in.Refs.Sim; simAllowed && s != nil && sim != nil {
		if v, ok := sim(*s); ok {
			label := Fill(BasisSimLabel, map[string]string{"Spec": in.Row.Spec, "band": strconv.Itoa(s.Band)})
			return v, &Basis{Label: label, Kind: BasisSim}, ""
		}
	}
	if in.Refs.Pool != nil {
		return 0, nil, Fill(PoolLacksMeasure, topFields(in, in.Refs.Pool.N))
	}
	return 0, nil, Fill(ThinPool, topFields(in, in.Refs.PoolCount))
}

// topFields fills BasisTop and ThinPool for in's spec, class, boss and size.
func topFields(in Input, n int) map[string]string {
	return map[string]string{
		"n":           strconv.Itoa(n),
		"Spec":        in.Row.Spec,
		"classPlural": ClassPlural(in.Row.Class),
		"Boss":        in.BossName,
		"size":        strconv.Itoa(in.Size),
	}
}
