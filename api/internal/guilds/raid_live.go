// api/internal/guilds/raid_live.go
//
// The raid's live surfaces (spec 2026-11-08-app-flows F2): the hooks the raid writes call so the
// Live Activity starts on raid.starting and ends when the raid is logged, and who the start
// goes to. livesurface.LivePusher implements LiveRaids.
package guilds

import (
	"context"
	"fmt"
)

// LiveRaids starts and ends the raid's Live Activities; nil sends nothing.
type LiveRaids interface {
	RaidStarting(ctx context.Context, raidID int64)
	RaidEnded(ctx context.Context, raidID int64)
}

// RaidAttendee is one account In or Late for a raid, by the character of its best answer.
type RaidAttendee struct {
	UserID       int64
	CharacterKey string
	Name         string
}

// RaidAttendees is every account In or Late for raidID, one row each (its In before its Late).
func (s *Service) RaidAttendees(ctx context.Context, raidID int64) ([]RaidAttendee, error) {
	rows, err := s.Store.Pool.Query(ctx, `
		select distinct on (signed_by) signed_by, character_key, name
		from guild_raid_signups where raid_id = $1 and state in ('in', 'late')
		order by signed_by, state = 'late', name`, raidID)
	if err != nil {
		return nil, fmt.Errorf("guilds: raid attendees: %w", err)
	}
	defer rows.Close()
	var out []RaidAttendee
	for rows.Next() {
		var a RaidAttendee
		if err := rows.Scan(&a.UserID, &a.CharacterKey, &a.Name); err != nil {
			return nil, fmt.Errorf("guilds: raid attendees: %w", err)
		}
		out = append(out, a)
	}
	return out, rows.Err()
}

// liveRaidEnded tells the live surfaces raidID is logged.
func (s *Service) liveRaidEnded(ctx context.Context, raidID int64) {
	if s.LiveRaids != nil {
		s.LiveRaids.RaidEnded(context.WithoutCancel(ctx), raidID)
	}
}
