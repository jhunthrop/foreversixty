package measure

import (
	"fmt"
	"strings"
)

// insufficient is what a figure prints as when it is under the sample
// floor. Never a number: a number goes into a constants file and nobody
// re-checks it.
const insufficient = "insufficient data"

// unknownLevel labels a target whose level no line in the log states. It
// gets its own row rather than being folded in with a known one: two
// levels in one row is a wrong number, and a missing level is a known
// unknown.
const unknownLevel = "unknown level"

// Table renders a report for a human reading a terminal on beta day.
// The validation job uses the struct, not this.
func (r Report) Table() string {
	var b strings.Builder
	fmt.Fprintf(&b, "forever-measure: %s, %d events, %d fights\n", r.Actor, r.Events, r.FightCount)
	r.writeSheet(&b)
	r.writePeriodic(&b)
	r.writeAttackTable(&b)
	r.writeProcs(&b)
	r.writeCoefficients(&b)
	r.writeIncomplete(&b)
	return b.String()
}

// writeSheet prints the character sheet the log itself reported, because
// every figure below is only interpretable against it.
func (r Report) writeSheet(b *strings.Builder) {
	if r.Sheet.Samples == 0 {
		fmt.Fprintf(b, "sheet: none; no advanced block named this actor\n\n")
		return
	}
	fmt.Fprintf(b, "sheet: attack power %.0f, spell power %.0f, armour %.0f, max health %.0f, level field %d (item level for a player), from %d advanced blocks\n\n",
		r.Sheet.AttackPower, r.Sheet.SpellPower, r.Sheet.Armor, r.Sheet.MaxHP, r.Sheet.Level, r.Sheet.Samples)
}

func (r Report) writePeriodic(b *strings.Builder) {
	fmt.Fprintf(b, "PERIODIC DAMAGE: does it crit, and by how much\n")
	fmt.Fprintf(b, "  %-24s %-8s %6s %6s %8s %9s %9s %8s\n",
		"spell", "school", "ticks", "crits", "can crit", "mean", "mean crit", "mult")
	for _, p := range r.Periodic {
		school := "magic"
		if p.Physical {
			school = "physical"
		}
		if !p.Enough {
			fmt.Fprintf(b, "  %-24s %-8s %6d %6d   %s\n", trunc(p.SpellName, 24), school, p.Ticks, p.CritTicks, insufficient)
			continue
		}
		fmt.Fprintf(b, "  %-24s %-8s %6d %6d %8v %9.1f %9.1f %8.3f\n",
			trunc(p.SpellName, 24), school, p.Ticks, p.CritTicks, p.CanCrit, p.MeanNormal, p.MeanCrit, p.Multiplier)
	}
}

func (r Report) writeAttackTable(b *strings.Builder) {
	fmt.Fprintf(b, "\nATTACK TABLE: the unified-hit and weapon-skill interaction\n")
	fmt.Fprintf(b, "  %-9s %-16s %13s %7s %7s %7s %7s %7s %7s\n",
		"type", "target", "lvl", "swings", "miss", "dodge", "parry", "glance", "crit")
	for _, t := range r.AttackTable {
		level := fmt.Sprintf("%d", t.TargetLevel)
		if t.TargetLevel == 0 {
			level = unknownLevel
		}
		if !t.Enough {
			fmt.Fprintf(b, "  %-9s %-16s %13s %7d   %s\n", t.AttackType, trunc(t.TargetName, 16), level, t.Swings, insufficient)
			continue
		}
		fmt.Fprintf(b, "  %-9s %-16s %13s %7d %6.2f%% %6.2f%% %6.2f%% %6.2f%% %6.2f%%\n",
			t.AttackType, trunc(t.TargetName, 16), level, t.Swings,
			t.Miss*100, t.Dodge*100, t.Parry*100, t.Glance*100, t.Crit*100)
	}
}

func (r Report) writeProcs(b *strings.Builder) {
	fmt.Fprintf(b, "\nPROCS: rate and internal cooldown\n")
	fmt.Fprintf(b, "  %-24s %6s %7s %6s %10s %9s %9s\n",
		"aura", "procs", "swings", "casts", "per swing", "per cast", "min gap")
	for _, p := range r.Procs {
		if !p.Enough {
			fmt.Fprintf(b, "  %-24s %6d %7d %6d   %s\n", trunc(p.SpellName, 24), p.Procs, p.Swings, p.Casts, insufficient)
			continue
		}
		fmt.Fprintf(b, "  %-24s %6d %7d %6d %9.3f%% %8.3f%% %9s\n",
			trunc(p.SpellName, 24), p.Procs, p.Swings, p.Casts, p.PerSwing*100, p.PerCast*100, p.MinGap)
	}
}

func (r Report) writeCoefficients(b *strings.Builder) {
	fmt.Fprintf(b, "\nDAMAGE: observed means, for coefficient fitting\n")
	fmt.Fprintf(b, "  %-24s %6s %10s %9s %12s\n", "spell", "hits", "mean", "stddev", "per sp")
	for _, c := range r.Coefficients {
		if !c.Enough {
			fmt.Fprintf(b, "  %-24s %6d   %s\n", trunc(c.SpellName, 24), c.Hits, insufficient)
			continue
		}
		perSP := "n/a"
		if c.Observed > 0 {
			perSP = fmt.Sprintf("%.4f", c.Observed)
		}
		fmt.Fprintf(b, "  %-24s %6d %10.1f %9.1f %12s\n", trunc(c.SpellName, 24), c.Hits, c.MeanDamage, c.StdDev, perSP)
	}
}

// writeIncomplete is the honest half of the tool: the constants this log
// cannot yield, in words, so a reader stops looking for them here and a
// logger knows what to go and record instead.
func (r Report) writeIncomplete(b *strings.Builder) {
	if len(r.Incomplete) == 0 {
		fmt.Fprintf(b, "\nNOT MEASURABLE FROM THIS LOG: nothing; this log supports every figure above.\n")
		return
	}
	fmt.Fprintf(b, "\nNOT MEASURABLE FROM THIS LOG\n")
	for _, s := range r.Incomplete {
		fmt.Fprintf(b, "  - %s\n", s)
	}
}

func trunc(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-1] + "…"
}
