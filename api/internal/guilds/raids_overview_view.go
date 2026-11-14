// api/internal/guilds/raids_overview_view.go
//
// The guild Overview's halves of a raid view (spec 2026-11-05-web-layout §7.1): every member
// reads the comp target; an officer also reads who has not answered, one row per account, and
// when the raid's non-signers were last pinged on Discord.
package guilds

import (
	"context"
	"fmt"
	"time"
)

// RaidPingState is when the raid's non-signers were last pinged and when the next ping may go.
// Both are null before the first ping.
type RaidPingState struct {
	LastPingedAt  *time.Time `json:"last_pinged_at"`
	NextAllowedAt *time.Time `json:"next_allowed_at"`
}

// pingEvery is the ping's throttle: one per raid per six hours.
const pingEvery = 6 * time.Hour

// fillOverview sets every view's comp target, and for an officer the not-signed raiders and
// the ping state.
func (s *Service) fillOverview(ctx context.Context, guildID, viewerID int64, views []RaidView) error {
	ids := raidIDsOf(views)
	targets, err := s.Store.compTargetsOf(ctx, ids)
	if err != nil {
		return err
	}
	for i := range views {
		var own *RoleCounts
		if c, ok := targets[views[i].ID]; ok {
			own = &c
		}
		views[i].CompTarget = compTargetOf(zoneOf(views[i].Zone), own)
	}
	officer, err := s.seesEveryAttendanceRow(ctx, guildID, viewerID)
	if err != nil {
		return err
	}
	if !officer {
		hidePickers(views)
		return nil
	}
	return s.fillOfficerOverview(ctx, guildID, views, ids)
}

// hidePickers drops the Discord names of who picked each character: they are for officers, who
// free a wrong pick, as on the guild roster (home.go).
func hidePickers(views []RaidView) {
	for i := range views {
		for j := range views[i].Signups {
			views[i].Signups[j].DiscordPickedBy = ""
		}
		for j := range views[i].NotSigned {
			views[i].NotSigned[j].DiscordPickedBy = ""
		}
	}
}

// fillOfficerOverview is fillOverview's officer half.
func (s *Service) fillOfficerOverview(ctx context.Context, guildID int64, views []RaidView, ids []int64) error {
	facts, err := s.Store.officerRosterFacts(ctx, guildID)
	if err != nil {
		return err
	}
	pings, err := s.Store.lastPings(ctx, ids)
	if err != nil {
		return err
	}
	for i := range views {
		rows := notSignedRaiders(views[i], facts)
		if err := s.Store.fillLastWeek(ctx, guildID, views[i], rows); err != nil {
			return err
		}
		views[i].NotSignedRaiders = &rows
		views[i].Ping = pingStateOf(pings[views[i].ID])
	}
	return nil
}

// pingStateOf is the ping state for the last ping at last (nil: never pinged).
func pingStateOf(last *time.Time) *RaidPingState {
	if last == nil {
		return &RaidPingState{}
	}
	at, next := last.UTC(), last.Add(pingEvery).UTC()
	return &RaidPingState{LastPingedAt: &at, NextAllowedAt: &next}
}

// lastPings reads the last ping of each of raidIDs; a raid never pinged is absent.
func (s *Store) lastPings(ctx context.Context, raidIDs []int64) (map[int64]*time.Time, error) {
	out := map[int64]*time.Time{}
	if len(raidIDs) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx, `select raid_id, last_pinged_at from guild_raid_pings where raid_id = any($1)`, raidIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: raid pings: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var at time.Time
		if err := rows.Scan(&id, &at); err != nil {
			return nil, fmt.Errorf("guilds: raid pings: %w", err)
		}
		out[id] = &at
	}
	return out, rows.Err()
}
