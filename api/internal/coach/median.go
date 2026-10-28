package coach

import "slices"

// MedianProfile is the per-key median of each measurement over the profiles that carry
// it: a key present in only some profiles takes the median of those. A consumable counts
// as present when more than half of the profiles that record it have it.
func MedianProfile(ps []Profile) Profile {
	if len(ps) == 0 {
		return Profile{}
	}
	return Profile{
		AliveMS:         median(collect(ps, func(p Profile) int64 { return p.AliveMS })),
		CountedMS:       median(collect(ps, func(p Profile) int64 { return p.CountedMS })),
		ActivePct:       median(collect(ps, func(p Profile) float64 { return p.ActivePct })),
		CastsPerMin:     medianMaps(ps, func(p Profile) map[string]float64 { return p.CastsPerMin }),
		Uptime:          medianMaps(ps, func(p Profile) map[string]float64 { return p.Uptime }),
		CappedPct:       median(collect(ps, func(p Profile) float64 { return p.CappedPct })),
		ManaLowMS:       median(collect(ps, func(p Profile) int64 { return p.ManaLowMS })),
		ManaLowFirstMS:  medianOfPresent(ps, func(p Profile) *int64 { return p.ManaLowFirstMS }),
		OverhealByRank:  medianMaps(ps, func(p Profile) map[string]float64 { return p.OverhealByRank }),
		DefensiveUptime: medianMaps(ps, func(p Profile) map[string]float64 { return p.DefensiveUptime }),
		Adherence:       medianOfPresent(ps, func(p Profile) *float64 { return p.Adherence }),
		Consumables:     majorityFlags(ps),
	}
}

type ordered interface{ ~int64 | ~float64 }

func collect[T any](ps []Profile, get func(Profile) T) []T {
	out := make([]T, len(ps))
	for i, p := range ps {
		out[i] = get(p)
	}
	return out
}

// median is the middle value, or the mean of the two middle values (integer division for
// integers); the zero value for none.
func median[T ordered](vs []T) T {
	if len(vs) == 0 {
		return 0
	}
	s := slices.Sorted(slices.Values(vs))
	mid := len(s) / 2
	if len(s)%2 == 1 {
		return s[mid]
	}
	return (s[mid-1] + s[mid]) / 2
}

func medianMaps(ps []Profile, get func(Profile) map[string]float64) map[string]float64 {
	byKey := map[string][]float64{}
	for _, p := range ps {
		for k, v := range get(p) {
			byKey[k] = append(byKey[k], v)
		}
	}
	if len(byKey) == 0 {
		return nil
	}
	out := make(map[string]float64, len(byKey))
	for k, vs := range byKey {
		out[k] = median(vs)
	}
	return out
}

// medianOfPresent is the median over the profiles whose optional value is set; nil when
// none sets it.
func medianOfPresent[T ordered](ps []Profile, get func(Profile) *T) *T {
	var vs []T
	for _, p := range ps {
		if v := get(p); v != nil {
			vs = append(vs, *v)
		}
	}
	if len(vs) == 0 {
		return nil
	}
	m := median(vs)
	return &m
}

func majorityFlags(ps []Profile) map[string]bool {
	have, seen := map[string]int{}, map[string]int{}
	for _, p := range ps {
		for k, v := range p.Consumables {
			seen[k]++
			if v {
				have[k]++
			}
		}
	}
	if len(seen) == 0 {
		return nil
	}
	out := make(map[string]bool, len(seen))
	for k, n := range seen {
		out[k] = have[k]*2 > n
	}
	return out
}
