package main

import (
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The sample ships as static files, so its size is a budget. A crowded world fight carries
// hundreds of enemy rows whose per-second series nobody reads one by one; the trim keeps
// every player and pet row whole and only shortens the long tails the report page lists
// below the first screen. The kept rows are untouched, so every figure they show is the
// engine's own.
const (
	// keepEnemyRows is how many enemy rows (by total) each of damage done and damage taken keeps.
	keepEnemyRows = 40
	// keepThreatPairs is how many player-on-enemy threat pairs a fight keeps, by threat.
	keepThreatPairs = 50
)

func isCreature(guid string) bool {
	return strings.HasPrefix(guid, "Creature-") || strings.HasPrefix(guid, "Vehicle-")
}

// trim shortens a fight's summary in place.
func trim(s *summary.Summary) {
	s.DamageDone = trimActors(s.DamageDone)
	s.DamageTaken = trimActors(s.DamageTaken)

	// Resource graphs are drawn for the people in the fight, not for every mob's mana.
	kept := s.Resources[:0:0]
	for _, r := range s.Resources {
		if !isCreature(r.GUID) {
			kept = append(kept, r)
		}
	}
	s.Resources = kept

	if len(s.ThreatByTarget) > keepThreatPairs {
		pairs := append([]summary.ThreatPair(nil), s.ThreatByTarget...)
		sort.SliceStable(pairs, func(i, j int) bool { return pairs[i].Threat > pairs[j].Threat })
		s.ThreatByTarget = pairs[:keepThreatPairs]
	}
}

func trimActors(rows []summary.Actor) []summary.Actor {
	var enemies []int
	for i, a := range rows {
		if isCreature(a.GUID) {
			enemies = append(enemies, i)
		}
	}
	if len(enemies) <= keepEnemyRows {
		return rows
	}
	sort.SliceStable(enemies, func(i, j int) bool { return rows[enemies[i]].Total > rows[enemies[j]].Total })
	drop := map[int]bool{}
	for _, i := range enemies[keepEnemyRows:] {
		drop[i] = true
	}
	out := make([]summary.Actor, 0, len(rows)-len(drop))
	for i, a := range rows {
		if !drop[i] {
			out = append(out, a)
		}
	}
	return out
}
