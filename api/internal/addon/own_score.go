// api/internal/addon/own_score.go
//
// Contract 3's own_score message (docs/contracts/addon-v2.md, added 2026-11-08): each of the
// account's characters' own Forever Score record, sent only to the owning account, so a player
// who hid the score (contract 4 export.score_visible false) still sees it although the nightly
// Data addon (contract 5) leaves them out. Sent for every character with a record, hidden or
// not. Flat, strings and numbers only; a part the record lacks is left out.
package addon

import (
	"context"
	"time"
)

const messageOwnScore = "own_score"

// OwnScore is one character's own score record: the figure, its tier, the newest rated fight's
// time, how many rated fights it rests on, and its five parts (nil for a part it lacks).
type OwnScore struct {
	Score                                            int
	Tier                                             string
	AsOf                                             time.Time
	Fights                                           int
	Gear, Experience, Parses, Attendance, Discipline *int
}

// OwnScoreSource reads a character's own record, whoever may see it; ok is false with none.
// rating.Store's ScoreV2 answers it (api/cmd/api).
type OwnScoreSource interface {
	OwnScore(ctx context.Context, characterKey string) (OwnScore, bool, error)
}

// ownScoreMessages is one own_score per key with a record; a failed read is logged and that
// character sends none (the sync never fails over it).
func (s *Service) ownScoreMessages(ctx context.Context, keys []string) []MessageEntry {
	if s.OwnScores == nil || len(keys) == 0 {
		return nil
	}
	var out []MessageEntry
	for _, key := range keys {
		r, ok, err := s.OwnScores.OwnScore(ctx, key)
		if err != nil {
			s.logger().Warn("addon", "op", "messages_own_score", "character", key, "err", err)
			continue
		}
		if !ok {
			continue
		}
		score, fights := r.Score, r.Fights
		out = append(out, MessageEntry{Type: messageOwnScore, Character: key, Score: &score, Tier: r.Tier,
			ScoreAsOf: r.AsOf.UTC().Format(time.DateOnly), Fights: &fights, Gear: r.Gear, Experience: r.Experience,
			Parses: r.Parses, Attendance: r.Attendance, Discipline: r.Discipline})
	}
	return out
}
