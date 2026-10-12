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
// not the player's), less the casts adherence_exclusions.go sets aside (white hits,
// consumables, mounts, movement, utility); a cast of an ability rot lists always counts.
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
		case excludedFromAdherence(c):
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

// adherence is row's adherence in fight, for the character key: nil when the card's spec
// has no shipped rotation or the player cast too few abilities.
//
//   - The spec is the rated card's own (row's class and spec, through DataSpecSlug's
//     aliases), never the character's current spec.
//   - The level band is the character's level as the site stores it at rating time
//     (rankings.CharacterLevel: the Battle.net import or the newest export's level=
//     section, the character page's own rule), banded by the addon's rule
//     (sims.RotationFor); with no level stored, band 60. The level is a current-character
//     fact: a backfill re-rate of an old fight uses today's level. That is acceptable
//     because adherence is not part of any score.
//   - The rotation is the newest loaded data build's.
func adherence(ctx context.Context, q rankings.RowQuerier, key string, fight summary.Summary, row summary.RosterRow) (*float64, error) {
	spec := DataSpecSlug(row.Class, row.Spec)
	if spec == "" {
		return nil, nil
	}
	if _, ok := sims.RotationFor("", spec, adherenceDefaultLevel); !ok {
		return nil, nil
	}
	level := adherenceDefaultLevel
	stored, err := rankings.CharacterLevel(ctx, q, key)
	if err != nil {
		return nil, err
	}
	if stored != nil {
		level = *stored
	}
	rot, ok := sims.RotationFor("", spec, level)
	if !ok {
		return nil, nil
	}
	return fightAdherence(fight, row, rot), nil
}
