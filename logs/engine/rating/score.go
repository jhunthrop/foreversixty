// logs/engine/rating/score.go
package rating

import "github.com/jhunthrop/foreversixty/logs/engine/summary"

// Score computes one player's Card for one closed fight. Every input is
// already-loaded data; Score makes no network or database call, which is
// what makes it unit-testable with golden fixtures (spec §8) the same way
// summary.Snapshot is.
func Score(
	fight summary.Summary,
	player string,
	tables CuratedTables,
	percentiles PercentileSource,
	assignments []Assignment,
	now ModelInfo,
) Card {
	row, found := rosterRow(fight, player)
	role := RoleDPS
	spec := ""
	if found {
		role = row.Role
		spec = specSlug(row.Class, row.Spec)
	}
	w, ok := now.Weights.For(role)
	if !ok {
		w, _ = now.Weights.For(RoleDPS)
	}

	bracket := Bracket{
		EncounterID: fight.EncounterID, Difficulty: fight.Difficulty,
		Spec: spec, Role: role, Kill: fight.Kill,
	}
	if band, ok := percentiles.KillTimeBand(fight.EncounterID, fight.Difficulty, fight.DurationMS); ok {
		bracket.KillTimeBand = band
	} else {
		bracket.KillTimeBand = BandTypical
	}
	if !fight.Kill {
		// A wipe never classifies a kill-time band (§1.2 buckets kill
		// durations only).
		bracket.KillTimeBand = ""
	}

	playerAssignments := assignmentsFor(assignments, player)

	outputC := scoreOutput(fight, row, bracket, percentiles)
	survivalC, deathScoreZero := scoreSurvival(fight, player, role, bracket, tables.Mechanics, assignments, percentiles)
	mechanicsC := scoreMechanicsComponent(fight, player, role, bracket, tables.Mechanics, percentiles)
	utilityC := scoreUtility(fight, player, bracket, tables.Utility, percentiles)
	preparationC := scorePreparation(fight, player, bracket, tables.Consumables, percentiles)
	activityC := scoreActivity(fight, player, bracket, tables.Mechanics, playerAssignments, percentiles)

	components := [6]Component{outputC, survivalC, mechanicsC, utilityC, preparationC, activityC}
	overallUncapped, overall, capped, basis := combine(&components, w, deathScoreZero, now.CapEnabled, now.CapThreshold)

	return Card{
		Overall: overall, OverallUncapped: overallUncapped, OverallCapped: capped,
		Components: components, Basis: basis,
		ModelVersion: now.ModelVersion, KillTimeBand: bracket.KillTimeBand,
	}
}

// assignmentsFor filters to one player's own assignment windows (spec §2).
func assignmentsFor(assignments []Assignment, player string) []Assignment {
	var out []Assignment
	for _, a := range assignments {
		if a.PlayerKey == player {
			out = append(out, a)
		}
	}
	return out
}
