// logs/engine/rating/percentile.go
// Shared plumbing for spec §1.2's percentile-vs-absolute rule, reused by
// every component file.
package rating

// percentileComponent finishes a percentile-scored Component: the score is
// the percentile itself on a 0-100 scale (spec §1.2: "use percentile-
// within-bracket").
func percentileComponent(c Component, pct float64, n int64) Component {
	score := round2(pct * 100)
	c.Score = score
	c.Basis = BasisPercentile
	p := round2(pct * 100)
	c.Percentile = &p
	c.BracketN = n
	return c
}

// placeOrExclude implements §1.2's rule for a component whose raw value has
// no defined absolute standard (an unbounded rate: raw DPS, prevented-
// damage per second, and similar): percentile when the bracket has enough
// samples, excluded otherwise.
func placeOrExclude(c Component, src PercentileSource, bracket Bracket, value float64, reason string) Component {
	if pct, n, ok := src.Placement(bracket, value); ok && n >= MinSample {
		return percentileComponent(c, pct, n)
	}
	c.Excluded, c.Reason = true, reason
	return c
}

// placeBoundedOrSelf implements §1.2's rule for a component whose raw value
// is already a bounded 0-100 share (RULING R5): percentile when the
// bracket has enough samples, else the value itself, scaled to 0-100, as
// the absolute standard -- never excluded, since a bounded share is always
// a meaningful score on its own.
func placeBoundedOrSelf(c Component, src PercentileSource, bracket Bracket, value0to100 float64) Component {
	if pct, n, ok := src.Placement(bracket, value0to100); ok && n >= MinSample {
		return percentileComponent(c, pct, n)
	}
	c.Score = round2(clamp(value0to100, 0, 100))
	c.Basis = BasisAbsolute
	return c
}

func clamp(x, lo, hi float64) float64 {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
}
