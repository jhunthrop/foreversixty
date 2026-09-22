// logs/engine/rating/mechanics_component.go
package rating

import (
	"strconv"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// scoreMechanicsComponent implements spec §1.3's Mechanics component: three
// independently-excludable parts (interrupts made, dispels made, own
// dispellable-debuff uptime), averaged with equal weight among the parts
// that have data.
func scoreMechanicsComponent(fight summary.Summary, player, role string, bracket Bracket, table *mechanics.Table, src PercentileSource) Component {
	c := Component{Name: ComponentNameMechanics}
	if table == nil {
		c.Excluded, c.Reason = true, ReasonNoMechanicsTable
		return c
	}
	hasInterrupt, hasDispel := false, false
	for _, m := range table.Mechanics {
		switch m.Kind {
		case mechanics.Interrupt:
			hasInterrupt = true
		case mechanics.Dispel:
			hasDispel = true
		}
	}
	if !hasInterrupt && !hasDispel {
		// §1.1's table row: "no mechanics table, or table has no
		// interrupt/dispel rows."
		c.Excluded, c.Reason = true, ReasonNoMechanicsTable
		return c
	}

	aliveSeconds := float64(timeAliveMS(fight, player)) / 1000
	var scores []float64
	var moments []Moment
	usedPercentile, usedAbsolute := false, false

	if hasInterrupt {
		credit, ms := interruptCredit(fight, player, table)
		perSec := perSecond(credit, aliveSeconds)
		b := bracket
		b.Component = ComponentMechanicsInterrupt
		if pct, n, ok := src.Placement(b, perSec); ok && n >= MinSample {
			scores = append(scores, round2(pct*100))
			usedPercentile = true
			moments = append(moments, ms...)
		}
	}
	if hasDispel {
		credit, ms := dispelCredit(fight, player, table)
		perSec := perSecond(credit, aliveSeconds)
		b := bracket
		b.Component = ComponentMechanicsDispel
		if pct, n, ok := src.Placement(b, perSec); ok && n >= MinSample {
			scores = append(scores, round2(pct*100))
			usedPercentile = true
			moments = append(moments, ms...)
		}

		uptimeShare, tracked, dms := dispellableDebuffUptimeShare(fight, player, table)
		if tracked {
			b.Component = ComponentMechanicsDebuffUptime
			if pct, n, ok := src.Placement(b, uptimeShare); ok && n >= MinSample {
				scores = append(scores, round2(100-pct*100))
				usedPercentile = true
			} else {
				// RULING R5: a bounded 0-100 share is its own absolute
				// standard when the bracket lacks samples.
				scores = append(scores, round2(clamp(100-uptimeShare, 0, 100)))
				usedAbsolute = true
			}
			moments = append(moments, dms...)
		}
	}

	if len(scores) == 0 {
		c.Excluded, c.Reason = true, ReasonNotEnoughSamples
		return c
	}
	var sum float64
	for _, s := range scores {
		sum += s
	}
	c.Score = round2(sum / float64(len(scores)))
	switch {
	case usedPercentile && usedAbsolute:
		c.Basis = BasisPercentile // the richer signal; Card.Basis carries "mixed" at the fight level
	case usedAbsolute:
		c.Basis = BasisAbsolute
	default:
		c.Basis = BasisPercentile
	}
	c.Moments = moments
	return c
}

func perSecond(value, seconds float64) float64 {
	if seconds <= 0 {
		return 0
	}
	return value / seconds
}

// interruptCredit sums, over every interrupt this player made, the average
// cost of one uninterrupted cast of that spell -- "damage prevented" (spec
// §1.3, part 1).
func interruptCredit(fight summary.Summary, player string, table *mechanics.Table) (float64, []Moment) {
	var credit float64
	var moments []Moment
	for _, x := range fight.Interrupts {
		if x.SourceGUID != player {
			continue
		}
		row := lookupMechanicRow(fight.Mechanics.Rows, mechanics.Interrupt, x.ExtraSpellID)
		if row == nil {
			continue
		}
		casts := row.Casts
		if casts < 1 {
			casts = 1
		}
		credit += float64(x.Count) * float64(row.Damage) / float64(casts)
		moments = append(moments, Moment{
			Kind: "interrupt", SpellID: x.ExtraSpellID, SpellName: x.ExtraSpellName,
			Anchor: "exchange-" + player + "-" + strconv.FormatInt(x.SpellID, 10),
		})
	}
	return credit, moments
}

// dispelCredit sums, over every dispel this player made, the average value
// of one dispel of that debuff -- the enemy healing it denied plus any
// damage the debuff itself did, divided by how many times it was
// dispelled overall (RULING R6: mirrors interrupts' own division by
// Casts, since the spec gives no divisor for dispels explicitly, and
// crediting the fight's whole prevented total on every single dispel a
// player made would double-count a debuff dispelled more than once).
func dispelCredit(fight summary.Summary, player string, table *mechanics.Table) (float64, []Moment) {
	var credit float64
	var moments []Moment
	for _, x := range fight.Dispels {
		if x.SourceGUID != player {
			continue
		}
		row := lookupMechanicRow(fight.Mechanics.Rows, mechanics.Dispel, x.ExtraSpellID)
		if row == nil {
			continue
		}
		dispelled := row.Dispelled
		if dispelled < 1 {
			dispelled = 1
		}
		credit += float64(x.Count) * float64(row.Healed+row.Damage) / float64(dispelled)
		moments = append(moments, Moment{
			Kind: "dispel", SpellID: x.ExtraSpellID, SpellName: x.ExtraSpellName,
			Anchor: "exchange-" + player + "-" + strconv.FormatInt(x.SpellID, 10),
		})
	}
	return credit, moments
}

// dispellableDebuffUptimeShare is this player's own uptime, as a 0-100
// share of the fight, on every dispellable debuff they carried (spec
// §1.3, part 3: "the target, not the healer who should have dispelled
// them"). tracked is false when this player was never afflicted by any
// dispel-classified debuff at all this fight -- distinct from a debuff
// applied and immediately cleansed (0% uptime, still tracked): with no
// track at all there is nothing to measure, so this sub-part drops out of
// Mechanics' average rather than crediting a full, untested 100.
func dispellableDebuffUptimeShare(fight summary.Summary, player string, table *mechanics.Table) (float64, bool, []Moment) {
	var totalMS int64
	var moments []Moment
	tracked := false
	for _, tr := range fight.Auras {
		if tr.Type != "DEBUFF" || tr.TargetGUID != player {
			continue
		}
		m, ok := table.Lookup(tr.SpellID)
		if !ok || m.Kind != mechanics.Dispel {
			continue
		}
		tracked = true
		totalMS += tr.UptimeMS
		if tr.UptimeMS > 0 {
			moments = append(moments, Moment{Kind: "dropped_debuff", SpellID: tr.SpellID, SpellName: tr.Name})
		}
	}
	if !tracked || fight.DurationMS <= 0 {
		return 0, tracked, moments
	}
	return clamp(float64(totalMS)/float64(fight.DurationMS)*100, 0, 100), tracked, moments
}

func lookupMechanicRow(rows []summary.MechanicRow, kind mechanics.Kind, spellID int64) *summary.MechanicRow {
	for i := range rows {
		if rows[i].Kind == kind && rows[i].SpellID == spellID {
			return &rows[i]
		}
	}
	return nil
}
