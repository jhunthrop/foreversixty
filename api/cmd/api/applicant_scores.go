package main

import (
	"context"

	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

// applicantScores reads the applicant sheet's Forever Score from the character page's scorer
// (rating.Service), so the sheet and the character page show one number.
type applicantScores struct{ scorer rankings.Scorer }

// ScoreOf is the character's Forever Score, nil without one.
func (a applicantScores) ScoreOf(ctx context.Context, key string) (*int, error) {
	sc, err := a.scorer.ForeverScore(ctx, key)
	if err != nil {
		return nil, err
	}
	return sc.Score, nil
}
