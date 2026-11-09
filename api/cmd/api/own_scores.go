package main

import (
	"context"
	"errors"

	"github.com/jhunthrop/foreversixty/api/internal/addon"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
)

// ownScores answers the addon inbox's own_score from the score's one home, rating.Store's
// ScoreV2, with no visibility gate: the inbox goes only to the characters' own account.
type ownScores struct{ store *rating.Store }

// OwnScore is the character's record; ok is false for a character with no rated fight.
func (o ownScores) OwnScore(ctx context.Context, key string) (addon.OwnScore, bool, error) {
	sc, err := o.store.ScoreV2(ctx, key)
	if errors.Is(err, rating.ErrNoLogs) {
		return addon.OwnScore{}, false, nil
	}
	if err != nil {
		return addon.OwnScore{}, false, err
	}
	c := sc.Components
	return addon.OwnScore{Score: sc.Score, Tier: sc.Tier, AsOf: sc.AsOf, Fights: sc.Fights,
		Gear: &c.Gear, Experience: &c.Experience, Parses: &c.Parses, Attendance: &c.Attendance,
		Discipline: &c.Discipline}, true, nil
}
