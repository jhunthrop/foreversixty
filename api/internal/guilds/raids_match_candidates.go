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

// candidateCap is the most logs the picker lists.
const candidateCap = 20

// Candidate is one log the match picker offers for a raid.
type Candidate struct {
	ReportID       string
	ZoneName       string
	UploadedAt     time.Time
	UploadedByName string
	Pulls, Kills   int
	DurationMS     int64
	// Closest marks the one selectable log (not matched to another raid) whose anchor is nearest
	// the raid's start.
	Closest     bool
	MatchedRaid *RaidRef
}

// MatchCandidatesResult is the picker's rows and whether the raid's zone has a known log name.
type MatchCandidatesResult struct {
	Rows      []Candidate
	ZoneKnown bool
}

// MatchCandidates lists the guild's evidence reports from the last candidateWindowDays days at
// now, nearest the raid's start first, at most candidateCap, the first selectable one (not
// another raid's) marked Closest and any raid already holding a log named. A raid zone whose log name is not known yet lists every
// zone's logs and answers ZoneKnown false. ErrNotFound when the guild has no such raid.
func (s *Store) MatchCandidates(ctx context.Context, guildID, raidID int64, now time.Time) (MatchCandidatesResult, error) {
	raid, found, err := s.GetRaid(ctx, guildID, raidID)
	if err != nil {
		return MatchCandidatesResult{}, err
	}
	if !found {
		return MatchCandidatesResult{}, ErrNotFound
	}
	logZone := ""
	if zone, ok := ZoneByKey(raid.Zone); ok && zone.LogZone != nil {
		logZone = *zone.LogZone
	}
	rows, err := s.guildLogs(ctx, guildID, logZone, raid.StartsAt, now)
	if err != nil {
		return MatchCandidatesResult{}, err
	}
	for i := range rows {
		if rows[i].MatchedRaid == nil || rows[i].MatchedRaid.ID == raidID {
			rows[i].Closest = true
			break
		}
	}
	return MatchCandidatesResult{Rows: rows, ZoneKnown: logZone != ""}, nil
}

func (s *Store) guildLogs(ctx context.Context, guildID int64, logZone string, startsAt, now time.Time) ([]Candidate, error) {
	rows, err := s.Pool.Query(ctx, `
		select r.id, coalesce(r.log_zone, ''), coalesce(r.completed_at, r.created_at),
		       coalesce(nullif(c.name, ''), split_part(coalesce(r.logging_character, ''), '/', 3)),
		       f.pulls, f.kills, f.duration_ms, mr.id, mr.title, mr.starts_at
		from reports r
		left join characters c on c.key = r.logging_character
		left join lateral (
		  select count(*) filter (where encounter_id is not null) as pulls,
		         count(*) filter (where encounter_id is not null and kill) as kills,
		         coalesce(max(start_ms + duration_ms) - min(start_ms), 0) as duration_ms
		  from fights where report_id = r.id) f on true
		left join guild_raids mr on mr.report_id = r.id and mr.guild_id = $1
		where `+matchGuildIs("r", "$1")+` and `+presence.EvidenceReport("r")+` and ($2 = '' or r.log_zone = $2)
		  and coalesce(r.completed_at, r.created_at) between $3 and $4
		order by abs(extract(epoch from (coalesce(r.completed_at, r.created_at) - $5::timestamptz))), r.id
		limit $6`,
		guildID, logZone, now.AddDate(0, 0, -candidateWindowDays), now, startsAt, candidateCap)
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
