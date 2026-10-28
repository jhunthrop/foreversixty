package coach

// Helpers for Review's tests (Task 14). Each builds on the shared fixtures and the F
// tasks' helpers; none is a known-good value.

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	reviewLongMS        = 120_000
	reviewShortMS       = 60_000
	bloodthirstRank1ID  = int64(23881)
	bloodthirstRank4ID  = int64(23894)
	bloodthirstEveryMS  = 6_000
	reviewRageHalf      = int64(500)
	reviewCapFromMS     = 100_000
	reviewCapToMS       = 112_000 // 12 s capped of 120 s counted: 10%, a minor F4
	reviewSparseEveryMS = 30_000  // four Bloodthirsts in 120 s: a major F1
)

// reviewBase is the Fury warrior over durationMS, active every second, alive, with only
// the casts and rage the caller adds, fully enchanted, alone in the roster but for rows,
// with Battle Shout, the given seeds and a combat potion at the pull.
func reviewBase(t *testing.T, durationMS int64, ss seeds, rows ...summary.RosterRow) Input {
	t.Helper()
	in := idleOf(t, furyInput(t), durationMS)
	in.Fight = withoutCasts(f4WithoutResources(in.Fight, furyGUID), furyGUID)
	in.Fight = withRoster(in.Fight, append([]summary.RosterRow{in.Row}, rows...)...)
	in.Fight = seedPotion(withSeeds(in.Fight, furyGUID, ss), furyGUID)
	in = withRaidBuff(t, withGear(in, preparedGear(true)), "Battle Shout")
	return in
}

// withBloodthirst is in with Bloodthirst of spellID cast at each instant and rage at
// half the bar except at its cap in each window.
func withBloodthirst(in Input, spellID int64, casts []int64, capped ...summary.Segment) Input {
	in.Fight = withCasts(in.Fight, furyGUID, spellID, "Bloodthirst", casts...)
	in.Fight = withResource(in.Fight, furyGUID, PowerRage, fixtureRageMax,
		f4Series(in.Fight.DurationMS, reviewRageHalf, fixtureRageMax, capped...)...)
	return in
}

// furyThreeFindings: no flask at 40 (a major F6), four Bloodthirsts in two minutes (a
// major F1) and rage capped 10% of the time (a minor F4).
func furyThreeFindings(t *testing.T) Input {
	t.Helper()
	in := reviewBase(t, reviewLongMS, seededAt(food("Grilled Squid")))
	return withBloodthirst(in, bloodthirstRank4ID, seqEvery(0, reviewSparseEveryMS, 4), span(reviewCapFromMS, reviewCapToMS))
}

// raidBuffOnly: everything brought, Bloodthirst on cooldown, and a priest in the roster
// whose Power Word: Fortitude the warrior lacks.
func raidBuffOnly(t *testing.T) Input {
	t.Helper()
	in := reviewBase(t, reviewShortMS, fullKit(), rosterClass("priest", "Priest"))
	return withBloodthirst(in, bloodthirstRank4ID, seqEvery(0, bloodthirstEveryMS, int(reviewShortMS/bloodthirstEveryMS)))
}

// raidBuffAndMinorRanks is raidBuffOnly with every Bloodthirst cast at Rank 1, where
// Rank 4 is trainable: a minor F5.
func raidBuffAndMinorRanks(t *testing.T) Input {
	t.Helper()
	in := reviewBase(t, reviewShortMS, fullKit(), rosterClass("priest", "Priest"))
	return withBloodthirst(in, bloodthirstRank1ID, seqEvery(0, bloodthirstEveryMS, int(reviewShortMS/bloodthirstEveryMS)))
}

// holyIdleAndOffRotation is a Holy priest idle for 50 s with the rotation's adherence
// rule in reach; a healer gets neither finding.
func holyIdleAndOffRotation(t *testing.T) Input {
	t.Helper()
	in := holyIdle(t, reviewLongMS, span(10_000, 60_000))
	in.Tables.RotationOK = true
	return in
}

// allFixtureInputs is every review fixture, by name.
func allFixtureInputs(t *testing.T) map[string]Input {
	t.Helper()
	return map[string]Input{
		"fury":                furyInput(t),
		"holy":                holyInput(t),
		"prot":                protInput(t),
		"fury three findings": furyThreeFindings(t),
		"raid buff only":      raidBuffOnly(t),
		"raid buff and ranks": raidBuffAndMinorRanks(t),
		"fury dead at 40 s":   furyFightDeadAt(t, reviewShortMS, bloodthirstRank4ID, []int64{0, 30_000}, repeat(reviewRageHalf, 60), 40_000),
		"fury idle":           furyIdle(t, reviewLongMS, span(10_000, 40_000)),
		"holy out of mana":    holyManaLow(t, reviewLongMS, 60_000, reviewLongMS),
		"holy overheal":       holyOverheal(t, 10965, 23, 44, poolOverheal(10965, 22, 21)),
		"prot sunder rank 4":  protSunderRank(t, fixtureLevel, 4, 22),
		"holy idle":           holyIdleAndOffRotation(t),
	}
}

// withoutAura is s with target's track of spellID removed.
func withoutAura(s summary.Summary, target string, spellID int64) summary.Summary {
	s.Auras = slices.DeleteFunc(slices.Clone(s.Auras), func(a summary.AuraTrack) bool {
		return a.TargetGUID == target && a.SpellID == spellID
	})
	return s
}

// anchorExists reports an anchor id in the pack's anchors.
func anchorExists(p Pack, id string) bool {
	return slices.ContainsFunc(p.Anchors, func(a Anchor) bool { return a.ID == id })
}
