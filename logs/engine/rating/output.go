// logs/engine/rating/output.go
package rating

import (
	"math"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// scoreOutput implements spec §1.3's Output component.
func scoreOutput(fight summary.Summary, row summary.RosterRow, bracket Bracket, src PercentileSource) Component {
	bracket.Component = ComponentNameOutput
	c := Component{Name: ComponentNameOutput}

	if !fight.Kill {
		// §2: a wipe's Output is excluded outright -- its damage window is
		// truncated by definition, and comparing partial output to a full
		// kill's is not a fair standard, absolute or percentile.
		c.Excluded, c.Reason = true, ReasonWipe
		return c
	}

	if row.ExecutionScore != nil {
		value := *row.ExecutionScore
		if pct, n, ok := src.Placement(bracket, value); ok && n >= MinSample {
			return percentileComponent(c, pct, n)
		}
		// RULING: capped at 100, not scaled past it -- a pull with
		// favourable crit variance can post execution_score > 1.0, and
		// scoring it above 100 would make "how good was this player"
		// indistinguishable from "how lucky was this pull" (spec §1.3).
		c.Score = round2(math.Min(100, value*100))
		c.Basis = BasisAbsolute
		return c
	}

	value := row.DPS
	if row.Role == RoleHealer {
		value = row.HPS
	}
	// Raw DPS/HPS fallback has no absolute standard: unbounded and
	// gear-dependent (spec §1.3) -- excluded rather than scored off an
	// ungrounded number when the bracket lacks samples.
	return placeOrExclude(c, src, bracket, value, ReasonNotEnoughSamples)
}
