package adapter

// Apportioning a rounded total across parts whose unrounded shares are
// known, so the parts sum to EXACTLY the total rather than to whatever
// independently rounding each part happens to add up to.
//
// This is the largest-remainder method (also called Hamilton
// apportionment, from the same arithmetic used to divide legislature
// seats among states by population): give every part its floor, then
// hand the few leftover units - the total minus the sum of the floors -
// to the parts with the largest fractional remainder, one unit each.
// Every part ends within one unit of its raw share, and the parts sum to
// the total exactly, by construction: the floors already account for
// everything but the leftover, and the leftover is handed out unit for
// unit.

import "sort"

// apportion splits total across len(shares) parts, in proportion to
// shares, so that the parts sum to exactly total. shares need not sum to
// total themselves - only their PROPORTIONS are used - which is what
// lets a caller pass unrounded per-part figures against an already
// rounded total and still get an exact split.
//
// An empty shares has nothing to apportion onto and returns an empty
// result; total then goes nowhere, which is correct - there is no part
// left to carry it. A shares that sums to zero or less has no proportion
// to follow (every part claims none of the total, or the notion of
// "share" is meaningless for it), so the total is spread evenly instead:
// evenly is exactly as defensible as any other split when nothing in the
// input favours one part over another.
func apportion(total int64, shares []float64) []int64 {
	out := make([]int64, len(shares))
	if len(shares) == 0 {
		return out
	}
	var sum float64
	for _, s := range shares {
		sum += s
	}
	if sum <= 0 {
		equal := make([]float64, len(shares))
		for i := range equal {
			equal[i] = 1
		}
		return apportion(total, equal)
	}

	type quota struct {
		part      int
		remainder float64
	}
	floors := make([]int64, len(shares))
	quotas := make([]quota, len(shares))
	var sumFloors int64
	for i, s := range shares {
		exact := s / sum * float64(total)
		floor := int64(exact)
		if exact < 0 && float64(floor) != exact {
			// Go truncates toward zero, not down, for a negative
			// float->int conversion; floor a negative share the same
			// way math.Floor would, so a caller passing signed shares
			// still gets floors that sum below the exact total rather
			// than above it. Damage shares are never negative in
			// practice; this keeps the function correct regardless.
			floor--
		}
		floors[i] = floor
		sumFloors += floor
		quotas[i] = quota{part: i, remainder: exact - float64(floor)}
	}

	// Largest remainder first; ties keep the parts' original order, so
	// the result is deterministic rather than dependent on sort's
	// internal tie-breaking.
	sort.SliceStable(quotas, func(i, j int) bool { return quotas[i].remainder > quotas[j].remainder })

	copy(out, floors)
	leftover := total - sumFloors
	for i := int64(0); i < leftover && int(i) < len(quotas); i++ {
		out[quotas[i].part]++
	}
	return out
}
