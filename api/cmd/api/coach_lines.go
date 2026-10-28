package main

import (
	"context"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/discordbot"
)

// coachLines is the bot's view of the coach: the service's night lines copied field for
// field. The words live in discordbot.
type coachLines struct{ svc *coach.Service }

func (c coachLines) NightLinesFor(ctx context.Context, userID int64, reportID string, budget time.Duration) ([]discordbot.CoachLine, error) {
	lines, err := c.svc.NightLinesFor(ctx, userID, reportID, budget)
	if err != nil {
		return nil, err
	}
	out := make([]discordbot.CoachLine, 0, len(lines))
	for _, l := range lines {
		out = append(out, discordbot.CoachLine{Character: l.Character, Stem: l.Line.Stem,
			On: l.Line.On, Of: l.Line.Of, URL: l.URL, Pending: l.Pending})
	}
	return out, nil
}

// coachLinesOf keeps a nil service from becoming a non-nil discordbot.CoachLines.
func coachLinesOf(svc *coach.Service) discordbot.CoachLines {
	if svc == nil {
		return nil
	}
	return coachLines{svc: svc}
}
