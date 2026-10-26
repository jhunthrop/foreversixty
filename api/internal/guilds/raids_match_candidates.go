// api/internal/guilds/raids_match_candidates.go
//
// The logs an officer may pick for a raid: the picker's rows.
package guilds

import (
	"context"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// candidateWindowDays is how far back the picker lists a guild's logs.
const candidateWindowDays = 14

// Candidate is one log the match picker offers for a raid.
type Candidate struct {
	ReportID       string
	ZoneName       string
	UploadedAt     time.Time
	UploadedByName string
	Pulls, Kills   int
	DurationMS     int64
	// Closest marks the one log whose anchor is nearest the raid's start.
	Closest     bool
	MatchedRaid *RaidRef
}

// MatchCandidates lists the guild's evidence reports of the raid's zone from the last
// candidateWindowDays days at now, newest first, with the closest to the raid's start marked
// and any raid already holding a log named. A zone with no log name has no candidates.
// ErrNotFound when the guild has no such raid.
func (s *Store) MatchCandidates(ctx context.Context, guildID, raidID int64, now time.Time) ([]Candidate, error) {
	raid, found, err := s.GetRaid(ctx, guildID, raidID)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ErrNotFound
	}
	zone, ok := ZoneByKey(raid.Zone)
	if !ok || zone.LogZone == nil {
		return []Candidate{}, nil
	}
	out, err := s.guildLogs(ctx, guildID, *zone.LogZone, now)
	if err != nil {
		return nil, err
	}
	markClosest(out, raid.StartsAt)
	return out, nil
}

func (s *Store) guildLogs(ctx context.Context, guildID int64, logZone string, now time.Time) ([]Candidate, error) {
	rows, err := s.Pool.Query(ctx, `
		select r.id, r.log_zone, coalesce(r.completed_at, r.created_at),
		       coalesce(nullif(c.name, ''), split_part(coalesce(r.logging_character, ''), '/', 3)),
		       f.pulls, f.kills, f.duration_ms, mr.id, mr.title, mr.starts_at
		from reports r
		left join characters c on c.key = r.logging_character
		left join lateral (
		  select count(*) filter (where encounter_id is not null) as pulls,
		         count(*) filter (where encounter_id is not null and kill) as kills,
		         coalesce(max(start_ms + duration_ms) - min(start_ms), 0) as duration_ms
		  from fights where report_id = r.id) f on true
		left join guild_raids mr on mr.report_id = r.id
		where r.guild_id = $1 and `+presence.EvidenceReport("r")+` and r.log_zone = $2
		  and coalesce(r.completed_at, r.created_at) between $3 and $4
		order by coalesce(r.completed_at, r.created_at) desc, r.id`,
		guildID, logZone, now.AddDate(0, 0, -candidateWindowDays), now)
	if err != nil {
		return nil, fmt.Errorf("guilds: match candidates: %w", err)
	}
	defer rows.Close()
	out := []Candidate{}
	for rows.Next() {
		var c Candidate
		var mrID *int64
		var mrTitle *string
		var mrStart *time.Time
		if err := rows.Scan(&c.ReportID, &c.ZoneName, &c.UploadedAt, &c.UploadedByName,
			&c.Pulls, &c.Kills, &c.DurationMS, &mrID, &mrTitle, &mrStart); err != nil {
			return nil, fmt.Errorf("guilds: match candidates: %w", err)
		}
		if mrID != nil {
			c.MatchedRaid = &RaidRef{ID: *mrID, Title: *mrTitle, StartsAt: *mrStart}
		}
		out = append(out, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: match candidates: %w", err)
	}
	return out, nil
}

// markClosest sets Closest on the one candidate whose upload is nearest the raid's start.
func markClosest(cs []Candidate, startsAt time.Time) {
	best := -1
	var bestGap time.Duration
	for i, c := range cs {
		gap := c.UploadedAt.Sub(startsAt).Abs()
		if best < 0 || gap < bestGap {
			best, bestGap = i, gap
		}
	}
	if best >= 0 {
		cs[best].Closest = true
	}
}
