package adapter

// The summary's clock.
//
// A summary carries one duration, but a sim measures two different
// things that could each claim it: AvgIterationDuration, the mean length
// of an iteration, and DPS(res).Mean, the mean of the per-iteration DPS
// values. Every damage figure in the summary is itself a third kind of
// mean - the across-iterations total divided by IterationsDone - so
// "table total / AvgIterationDuration" and "DPS(res).Mean" are two
// different means of the encounter length interacting with two
// different means of the damage, and they disagree by a fraction of a
// percent whenever the encounter length varies. deriveDurationMS instead
// picks the one duration that makes the table and the headline agree
// exactly: see the package comment in adapter.go for the fuller account.

import (
	"math"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// sumActorTotals adds up every actor's Total in a damage table - the
// player and every pet - because RaidMetrics.Dps is the RAID's DPS, and
// a hunter's or warlock's pet damage is folded into it (sim/core adds a
// pet's dps.Total into its owner's before the raid sums the owners:
// character.doneIteration -> AddFinalPetMetrics). A duration derived
// from the player's total alone would under-count for any spec with a
// pet and the table still would not match the headline.
func sumActorTotals(actors []summary.Actor) int64 {
	var total int64
	for _, a := range actors {
		total += a.Total
	}
	return total
}

// deriveDurationMS is the fight length that makes the summary's own
// total damage, divided by this duration, equal meanDPS - to the
// precision an integer-millisecond field allows. It is computed from
// totalDamage, the sum of the summary's already-rounded actor totals,
// rather than from the engine's raw float damage sum, because deriving
// it from anything the summary does not already carry would reopen the
// disagreement between the table and the headline that this function
// exists to close.
//
// Zero total damage and a zero or negative meanDPS are both degenerate:
// dividing by DPS gives no duration at all in the first case (nothing
// happened) and a meaningless or infinite one in the second (a
// non-positive rate has no well-defined "time to deal this damage").
// fallbackMS - the mean iteration length the engine measured directly -
// is the sane answer either way, and it must be positive: DurationMS
// feeds a per-second division everywhere else in the summary
// (logs/engine/summary/roster.go's DPS and activity-% arithmetic, the
// report's own per-second table), and a zero duration would turn every
// one of those into a division by zero.
func deriveDurationMS(totalDamage int64, meanDPS float64, fallbackMS int64) int64 {
	if totalDamage <= 0 || meanDPS <= 0 {
		return fallbackMS
	}
	return int64(math.Round(float64(totalDamage) / meanDPS * 1000))
}
