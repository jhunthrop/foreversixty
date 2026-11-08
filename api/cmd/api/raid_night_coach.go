package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/coach"
)

// raidNightCoachBudget bounds the coach read inside one raid night answer.
const raidNightCoachBudget = 2 * time.Second

// raidNightCoach is the raid night's view of the coach (guilds.RaidNightCoach): the character's
// night lines as the teaser's sentences, the finding's stem and how many bosses had it first.
type raidNightCoach struct{ svc *coach.Service }

func (c raidNightCoach) CoachNotes(ctx context.Context, userID int64, reportID, characterKey string) ([]string, error) {
	lines, err := c.svc.NightLinesFor(ctx, userID, reportID, raidNightCoachBudget)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, l := range lines {
		if l.PlayerKey != characterKey || l.Pending {
			continue
		}
		note := l.Line.Stem
		if l.Line.Of > 1 {
			note = fmt.Sprintf("%s First fix on %d of %d bosses.", note, l.Line.On, l.Line.Of)
		}
		out = append(out, note)
	}
	return out, nil
}
