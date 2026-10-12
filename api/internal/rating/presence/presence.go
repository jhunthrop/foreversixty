// api/internal/rating/presence/presence.go

// Package presence is the one reader of "was this character there": the Forever Score's
// attendance component (rating.Store.guildNights) and the guild's attendance, log
// verification, moderator evidence and raid present list (api/internal/guilds) all call In.
//
// It is its own package under rating, not a function in rating itself, because guilds cannot
// import rating: rating imports rankings, which imports guilds (for the crest URL). This
// package imports only character.
package presence

import (
	"context"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// Row is one character at one report: it has a rated row (rating_scores) or a ranked row
// (fight_metrics, not removed) there.
//
// Presence is read by character key, never from fights.players: that column holds the log's
// GUIDs (api/internal/engine/engine.go), which no character key ever equals. The ranked and
// rated rows exist only for boss fights in reports that are not private (reports/ingest.go's
// rank, and the rater), so a private or trash-only report has nobody present.
type Row struct {
	Key      string
	ReportID string
	// Night is the UTC date of the row's fought_at.
	Night time.Time
}

// In lists every character present in the reports reportIDs, one row per character, report
// and UTC night, ordered by key, report and night. keys narrows it to those characters; nil
// means everyone. No reports gives nothing, without a query. q is a pool or a transaction.
func In(ctx context.Context, q character.Querier, reportIDs, keys []string) ([]Row, error) {
	if len(reportIDs) == 0 {
		return nil, nil
	}
	rows, err := q.Query(ctx, `
		select distinct player_key, report_id, (fought_at at time zone 'UTC')::date
		from (
		  select rs.player_key, rs.report_id, rs.fought_at
		  from rating_scores rs
		  where rs.report_id = any($1) and ($2::text[] is null or rs.player_key = any($2))
		  union all
		  select m.player_key, m.report_id, m.fought_at
		  from fight_metrics m
		  where m.report_id = any($1) and m.state <> 'removed' and ($2::text[] is null or m.player_key = any($2))
		) present
		order by 1, 2, 3`, reportIDs, keys)
	if err != nil {
		return nil, fmt.Errorf("presence: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var r Row
		if err := rows.Scan(&r.Key, &r.ReportID, &r.Night); err != nil {
			return nil, fmt.Errorf("presence: %w", err)
		}
		out = append(out, r)
	}
	return out, rows.Err()
}
