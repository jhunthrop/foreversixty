// api/internal/dataaddon/guild_aggregate.go
package dataaddon

import (
	"sort"
)

// guildRecord is one guild's entry in ForeverSixtyData.guilds (contract 5), keyed by the
// guild's name as the game's GetGuildInfo returns it.
type guildRecord struct {
	// Region is which file the record goes to when the render splits. It is not written.
	Region string
	// Standing is the guild's rank among every guild with a verified member: by distinct
	// raid bosses killed in the window, ties by attendance; guilds still tied share a rank.
	Standing int
	// Nights is the distinct UTC dates of the guild's public reports in the window.
	Nights int
	// Roster is the verified member count, anonymised members included (a count names no
	// one).
	Roster int
	// Attendance is the mean of the listed members' attendance fractions (ScoreV2's own);
	// nil when no member has a score, and then it is not written.
	Attendance *float64
	// Members is the Name-Realm keys of the guild's characters that have a record, sorted.
	Members []string
}

// guildIdentity names one guild with at least one verified member -- the only kind this
// job publishes a record for.
type guildIdentity struct {
	ID                    int64
	Region, Ruleset, Name string
}

// meanAttendance is the mean of the fractions, or nil for none.
func meanAttendance(fractions []float64) *float64 {
	if len(fractions) == 0 {
		return nil
	}
	var sum float64
	for _, f := range fractions {
		sum += f
	}
	mean := sum / float64(len(fractions))
	return &mean
}

// guildStanding is what a guild is ranked by.
type guildStanding struct {
	ID         int64
	Kills      int
	Attendance *float64
}

// standings ranks guilds by distinct raid bosses killed, most first, ties broken by
// attendance (no attendance counting as none); guilds tied on both share a rank, and the
// next rank skips past them (1, 2, 2, 4).
func standings(guilds []guildStanding) map[int64]int {
	attendance := func(g guildStanding) float64 {
		if g.Attendance == nil {
			return 0
		}
		return *g.Attendance
	}
	sorted := append([]guildStanding(nil), guilds...)
	sort.SliceStable(sorted, func(i, j int) bool {
		if sorted[i].Kills != sorted[j].Kills {
			return sorted[i].Kills > sorted[j].Kills
		}
		return attendance(sorted[i]) > attendance(sorted[j])
	})
	out := make(map[int64]int, len(sorted))
	for i, g := range sorted {
		if i > 0 && g.Kills == sorted[i-1].Kills && attendance(g) == attendance(sorted[i-1]) {
			out[g.ID] = out[sorted[i-1].ID]
			continue
		}
		out[g.ID] = i + 1
	}
	return out
}
