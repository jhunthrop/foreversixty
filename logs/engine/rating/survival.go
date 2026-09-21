// logs/engine/rating/survival.go
package rating

import (
	"strconv"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The death-penalty cause multipliers and the 65/35 split, spec §1.3.
const (
	deathCauseAvoidable     = 1.00
	deathCauseUnavoidable   = 0.15
	deathCauseUnclassified  = 0.50
	deathPenaltyBase        = 70.0
	survivalDeathWeight     = 0.65
	survivalAvoidableWeight = 0.35
)

// scoreSurvival implements spec §1.3's Survival component:
// 0.65*DeathScore + 0.35*AvoidableHitScore. Returns the Component and,
// separately, deathScoreZero -- true only when the component was actually
// scored (not excluded) and DeathScore alone floored to zero, spec §1.5's
// cap condition, which Component itself carries no field for.
func scoreSurvival(fight summary.Summary, player, role string, bracket Bracket, table *mechanics.Table, assignments []Assignment, src PercentileSource) (Component, bool) {
	c := Component{Name: ComponentNameSurvival}
	if table == nil {
		// §1.1/§2: no table means no killing-blow classification, and a
		// death-timing-only signal is not the "avoidable damage taken"
		// component the brief asks for.
		c.Excluded, c.Reason = true, ReasonNoMechanicsTable
		return c, false
	}

	deathScore, deathMoments := deathScoreFor(fight, player, table)
	deathScoreZero := deathScore == 0

	bracket.Component = ComponentSurvivalAvoidableHit
	avoidablePerSec := avoidableDamagePerSecond(fight, player, role, assignments)
	pct, n, ok := src.Placement(bracket, avoidablePerSec)
	if !ok || n < MinSample {
		// RULING R4: avoidable-damage-per-second is an unbounded rate with
		// no defined absolute standard (spec §1.3 gives none); when the
		// bracket lacks samples, Survival falls back to DeathScore alone
		// rather than inventing a numeric standard.
		c.Score = round2(deathScore)
		c.Basis = BasisAbsolute
		c.Moments = deathMoments
		return c, deathScoreZero
	}

	avoidableScore := round2(100 - pct*100)
	c.Score = round2(survivalDeathWeight*deathScore + survivalAvoidableWeight*avoidableScore)
	c.Basis = BasisPercentile
	p := round2(pct * 100)
	c.Percentile = &p
	c.BracketN = n
	c.Moments = deathMoments
	return c, deathScoreZero
}

// deathScoreFor is 100 minus the sum of every death's penalty, floored at
// zero, plus the death Moments the card links to.
func deathScoreFor(fight summary.Summary, player string, table *mechanics.Table) (float64, []Moment) {
	score := 100.0
	var moments []Moment
	for _, d := range fight.Deaths {
		if d.GUID != player {
			continue
		}
		timeRemaining := 1.0
		if fight.DurationMS > 0 {
			timeRemaining = 1 - float64(d.AtMS)/float64(fight.DurationMS)
		}
		if timeRemaining < 0 {
			timeRemaining = 0
		}
		var spellID int64
		var spellName string
		avoidable := false
		cause := deathCauseUnclassified
		if d.KillingBlow != nil {
			spellID, spellName = d.KillingBlow.SpellID, d.KillingBlow.SpellName
			if m, ok := table.Lookup(spellID); ok {
				switch m.Kind {
				case mechanics.Avoidable:
					cause, avoidable = deathCauseAvoidable, true
				case mechanics.Unavoidable:
					cause = deathCauseUnavoidable
				}
			}
		}
		score -= deathPenaltyBase * cause * timeRemaining
		moments = append(moments, Moment{
			Kind: "death", AtMS: d.AtMS, SpellID: spellID, SpellName: spellName,
			Avoidable: avoidable, Anchor: "death-" + player + "-" + strconv.FormatInt(d.AtMS, 10),
		})
	}
	if score < 0 {
		score = 0
	}
	return score, moments
}

// avoidableDamagePerSecond sums avoidable-mechanic damage taken over the
// fight. A row whose Role does not match the player's own role always
// counts in full -- it was never this player's job. A Role-tagged row
// matching the player's own role goes through excusedByAssignment: the
// curated table alone cannot say WHICH player of that role a given hit was
// "for" (spec §1.3's original text wrongly excused every same-role hit
// unconditionally, which credited the off tank standing in a cleave meant
// for the main tank the same as the main tank doing their job); an
// Assignment naming who actually had it during that hit is what tells
// them apart (RULING, whole-branch review).
func avoidableDamagePerSecond(fight summary.Summary, player, role string, assignments []Assignment) float64 {
	var total int64
	for _, row := range fight.Mechanics.Rows {
		if row.Kind != mechanics.Avoidable {
			continue
		}
		for _, hit := range row.Players {
			if hit.GUID != player {
				continue
			}
			if row.Role == role && excusedByAssignment(fight, assignments, row.Role, player, hit) {
				continue
			}
			total += hit.Damage
		}
	}
	seconds := float64(fight.DurationMS) / 1000
	if seconds <= 0 {
		return 0
	}
	return float64(total) / seconds
}

// excusedByAssignment decides whether one Role-tagged avoidable hit is
// excused for player: with no Assignment overlapping the hit's own window
// for any player of mechanicRole, every player of that role is excused (we
// cannot tell the main tank doing their job from the off tank standing in
// it, and wrongly punishing the main tank for every cleave is the worse
// error); with at least one such assignment, player is excused only if one
// of them names player specifically -- otherwise the hit was demonstrably
// someone else's to take, and counts in full.
func excusedByAssignment(fight summary.Summary, assignments []Assignment, mechanicRole, player string, hit summary.MechanicHit) bool {
	anyForRole := false
	hitStart, hitEnd := hit.FirstMS, hitEndMS(hit)
	for _, a := range assignments {
		if !overlapsMS(a.FromMS, a.ToMS, hitStart, hitEnd) {
			continue
		}
		row, ok := rosterRow(fight, a.PlayerKey)
		if !ok || row.Role != mechanicRole {
			continue
		}
		anyForRole = true
		if a.PlayerKey == player {
			return true
		}
	}
	return !anyForRole
}

// hitEndMS is the hit's own window end: LastMS when the track recorded a
// span, or FirstMS nudged forward by one millisecond for a single instant
// so overlapsMS's half-open interval still contains it.
func hitEndMS(hit summary.MechanicHit) int64 {
	if hit.LastMS > hit.FirstMS {
		return hit.LastMS
	}
	return hit.FirstMS + 1
}
