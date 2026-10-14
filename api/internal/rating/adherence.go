package rating

// Rotation adherence, the last-raid fact contract 3 calls adherence
// (docs/contracts/addon-v2.md, last_raid.bosses[].adherence; L6.5): the share of a player's
// ability casts in a fight that are abilities the spec's shipped rotation lists for the
// character's level band. Worked out by the rater beside the other last-raid facts
// (last_raid.go) and stored on the player's rating_scores row (migration 0037_adherence).
// It is a fact the addon shows, never a scored component: it does not enter the card or
// ScoreV2.

import (
	"context"
	"math"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	// AdherenceMinCasts is the fewest ability casts adherence is worked out from; under it
	// there is too little to judge and adherence is absent.
	AdherenceMinCasts = 10
	// adherenceDefaultLevel is the level band used when the site stores no level for the
	// character: last_raid is a raid's message, and raids are level-60 content.
	adherenceDefaultLevel = 60
)

// fightAdherence is the adherence of row's player in fight to rot, rounded to two decimals:
// casts of abilities rot lists (at any rank) over every ability cast. Ability casts are the
// player's own successful casts (a pet's or a guardian's rows carry the pet's GUID and are
// not the player's), less the casts adherence_exclusions.go sets aside (white hits, items,
// racials, consumables, dispels, interrupts, taunts, resurrections, self-buffs, mounts,
// movement, utility); a cast of an ability rot lists always counts.
// nil under AdherenceMinCasts ability casts.
func fightAdherence(fight summary.Summary, row summary.RosterRow, rot sims.Rotation) *float64 {
	var listed, total int64
	for _, c := range fight.Casts {
		if c.GUID != row.GUID || c.Succeeded <= 0 {
			continue
		}
		switch {
		case rot.Lists(c.SpellID):
			listed += c.Succeeded
			total += c.Succeeded
		case excludedFromAdherence(c, rot):
		default:
			total += c.Succeeded
		}
	}
	if total < AdherenceMinCasts {
		return nil
	}
	v := math.Round(float64(listed)/float64(total)*100) / 100
	return &v
}

// rosterLevels reads every roster player's level for the adherence band, in one query on the
// pool before the rating transaction begins: rankings.CharacterLevels, the character page's
// own rule (the newer of the Battle.net import and the addon export). known is false when the
// read failed; that is logged, and every player's adherence is then absent (the band cannot
// be known), but the rating itself goes on.
func (s *Store) rosterLevels(ctx context.Context, f RatedFight) (levels map[string]*int, known bool) {
	keys := make([]string, 0, len(f.Summary.Roster))
	for _, row := range f.Summary.Roster {
		keys = append(keys, character.KeyFromUnit(f.Region, f.Ruleset, row.Name))
	}
	var q character.Querier = s.Pool
	if s.levels != nil {
		q = s.levels
	}
	levels, err := rankings.CharacterLevels(ctx, q, keys)
	if err != nil {
		s.logger().Warn("rating", "op", "adherence_levels", "report", f.ReportID, "fight", f.FightIndex, "err", err)
		return nil, false
	}
	return levels, true
}

// adherence is row's adherence in fight: nil when the level could not be read
// (levelKnown false), the card's spec has no shipped rotation, or the player cast too few
// abilities.
//
//   - The spec is the rated card's own (row's class and spec, through DataSpecSlug's
//     aliases), never the character's current spec.
//   - The level band is the character's level as the site stores it at rating time
//     (rosterLevels), banded by the addon's rule (sims.RotationFor); with no level stored,
//     band 60. The level is a current-character fact: a backfill re-rate of an old fight
//     uses today's level. That is acceptable because adherence is not part of any score.
//   - The rotation is the newest loaded data build's.
func adherence(fight summary.Summary, row summary.RosterRow, level *int, levelKnown bool) *float64 {
	if !levelKnown {
		return nil
	}
	spec := DataSpecSlug(row.Class, row.Spec)
	if spec == "" {
		return nil
	}
	band := adherenceDefaultLevel
	if level != nil {
		band = *level
	}
	rot, ok := sims.RotationFor("", spec, band)
	if !ok {
		return nil
	}
	return fightAdherence(fight, row, rot)
}
