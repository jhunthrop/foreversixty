// api/internal/guilds/raids_nudge.go
//
// raid.signup_nudge at each player's raid-day reminder (spec 2026-11-08-app-flows F1, appendix
// B): on the raid's day in the player's own time zone, at their reminder hour (notification
// settings raid_reminder, 15:00 by default), a verified member who has not answered is nudged.
// A reminder at or after the start falls back to nudgeFallback before it. The raid-posts job
// runs every ten minutes; each nudge fires while its due time is inside the catch-up window,
// and the event's dedupe key holds it to once per raid and player.
package guilds

import (
	"context"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/textfmt"
)

// nudgeFallback is how long before the start a player whose reminder falls at or after it is
// nudged.
const nudgeFallback = 2 * time.Hour

// nudgeWindow is how far ahead a raid can be and still be today for someone: a reminder at
// 00:00 for a raid at 23:59, plus the catch-up.
const nudgeWindow = 24*time.Hour + eventCatchUp

// NudgeReader finds who a raid's nudge goes to and their settings (events.Store).
type NudgeReader interface {
	Recipients(ctx context.Context, a events.Audience) ([]int64, error)
	PrefsOf(ctx context.Context, userIDs []int64) (map[int64]events.Prefs, error)
}

// nudgeDue is when p's raid-day reminder for a raid at start comes due.
func nudgeDue(start time.Time, p events.Prefs) time.Time {
	loc := textfmt.Location(p.TimeZone)
	local := start.In(loc)
	due := time.Date(local.Year(), local.Month(), local.Day(), 0, p.RaidReminder, 0, 0, loc)
	if !due.Before(start) {
		return start.Add(-nudgeFallback)
	}
	return due
}

// nudgeable reports whether r can be nudged at now: planned and starting inside the window.
func nudgeable(r scheduledRaid, now time.Time) bool {
	return r.Status == RaidPlanned && r.StartsAt.After(now) && !r.StartsAt.After(now.Add(nudgeWindow))
}

// publishNudge nudges r's unanswered accounts whose reminder is due at now; true when it
// published one.
func (s *Service) publishNudge(ctx context.Context, r scheduledRaid, now time.Time) (bool, error) {
	facts, ok := s.raidFacts(ctx, r.ID, "event_nudge")
	if !ok {
		return false, nil
	}
	users, err := s.Nudges.Recipients(ctx, events.Audience{Kind: events.ToUnanswered, GuildID: facts.GuildID, RaidID: r.ID})
	if err != nil || len(users) == 0 {
		return false, err
	}
	prefs, err := s.Nudges.PrefsOf(ctx, users)
	if err != nil {
		return false, fmt.Errorf("guilds: nudge prefs: %w", err)
	}
	var due []int64
	for _, uid := range users {
		if dueWithin(nudgeDue(r.StartsAt, prefs[uid]), now) {
			due = append(due, uid)
		}
	}
	if len(due) == 0 {
		return false, nil
	}
	s.publishAll(ctx, []events.Event{raidNudgeEvent(facts.GuildID, facts.Post.RaidPosted, due)})
	return true, nil
}
