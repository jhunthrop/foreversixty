package coach

// Helpers for F1's tests only (the shared ones live in fixtures_test.go).

import (
	"maps"
	"testing"
)

const (
	fixtureBloodthirstID = int64(23881) // Bloodthirst: category cooldown 6 s, cost 300 (30 rage)
	fixtureRagePower     = int64(1)
	fixtureRageMax       = int64(1000)
)

// furyFight is the Fury warrior's input over durationMS with one ability cast at casts and
// the rage series (one reading per second; nil is a rage track with no readings).
func furyFight(t *testing.T, durationMS, spellID int64, casts, rage []int64) Input {
	t.Helper()
	in := furyInput(t)
	fight := withoutCasts(withDuration(in.Fight, durationMS), furyGUID)
	fight = withCasts(fight, furyGUID, spellID, "Bloodthirst", casts...)
	in.Fight = withResource(fight, furyGUID, fixtureRagePower, fixtureRageMax, rage...)
	return in
}

// furyFightDeadAt is furyFight with the warrior dying at deathMS.
func furyFightDeadAt(t *testing.T, durationMS, spellID int64, casts, rage []int64, deathMS int64) Input {
	t.Helper()
	in := furyFight(t, durationMS, spellID, casts, rage)
	in.Fight = withDeath(in.Fight, furyGUID, deathMS, 0, "Melee")
	return in
}

// withUnknownCost is in with the data stating spellID's cost only as a share of the bar.
func withUnknownCost(in Input, spellID int64) Input {
	spells := maps.Clone(in.Tables.Spells)
	sc := spells[spellID]
	sc.Cost, sc.CostKnown = 0, false
	spells[spellID] = sc
	in.Tables.Spells = spells
	return in
}

// furyTalentUncast is a fight where the warrior never casts the talent ability Bloodthirst.
func furyTalentUncast(t *testing.T) Input {
	t.Helper()
	in := furyInput(t)
	in.Fight = withDuration(withoutCasts(in.Fight, furyGUID), 60_000)
	in.Fight = withResource(in.Fight, furyGUID, fixtureRagePower, fixtureRageMax, repeat(500, 60)...)
	return in
}

// holyWithBloodthirstRule is a Holy priest whose rule file lists Bloodthirst as a core
// ability and who casts it five times, with no efficiency cooldown at all.
func holyWithBloodthirstRule(t *testing.T) Input {
	t.Helper()
	in := holyInput(t)
	in.Fight = withCasts(withDuration(in.Fight, 60_000), priestGUID, fixtureBloodthirstID, "Bloodthirst", 0, 6000, 12000, 18000, 24000)
	in.Tables.Rules.Core = furyRules.Core
	in.Tables.Spells = maps.Clone(warriorSpells)
	return in
}

// simCPM is Refs with no pool and a sim reference at band with family at cpm.
func simCPM(family string, cpm float64, band int) Refs {
	return Refs{Sim: &SimRef{Band: band, CastsPerMin: map[string]float64{family: cpm}}}
}

// withWorstAction is in with family on the spec's sims WorstActions list.
func withWorstAction(in Input, family string) Input {
	worst := maps.Clone(in.Tables.WorstActions)
	if worst == nil {
		worst = map[string]bool{}
	}
	worst[family] = true
	in.Tables.WorstActions = worst
	return in
}
