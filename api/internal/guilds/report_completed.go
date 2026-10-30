package guilds

import (
	"context"
	"fmt"
)

// ReportCompleted is the reports.Completed hook: a report has just completed, or a complete
// report has just been given its guild. It matches the report to the guild's raid and, on a
// match, computes the raid's attendance and attaches its loot (all in AutoMatch's one
// transaction). Best effort: a failure is logged, never returned.
func (s *Service) ReportCompleted(ctx context.Context, reportID string) {
	raid, matched, err := s.Store.AutoMatch(ctx, reportID)
	if err != nil {
		s.logger().Error("guilds", "op", "report_completed_match", "report", reportID, "err", err)
		return
	}
	if !matched {
		s.logger().Info("guilds", "op", "report_completed", "report", reportID, "matched", false)
		return
	}
	s.markRaid(ctx, raid.ID)
	s.logger().Info("guilds", "op", "report_completed", "report", reportID, "raid", raid.ID, "matched", true)
	if s.Events != nil {
		s.publishAll(ctx, s.nightReadyEvents(ctx, raid.ID, reportID))
	}
}

// ReportGuildMoved is the reports.GuildMoved hook: an officer's PATCH has just set reportID's
// guild. A raid of another guild still holding the report releases it through ClearMatch (by is
// the officer), so that raid's report, night and attendance go and it never re-picks the log.
// Best effort: a failure is logged, never returned.
func (s *Service) ReportGuildMoved(ctx context.Context, reportID string, by int64) {
	held, err := s.Store.foreignMatches(ctx, reportID)
	if err != nil {
		s.logger().Error("guilds", "op", "report_guild_moved", "report", reportID, "err", err)
		return
	}
	for _, m := range held {
		if _, err := s.Store.ClearMatch(ctx, m.guildID, m.raidID, by); err != nil {
			s.logger().Error("guilds", "op", "report_guild_moved", "report", reportID, "raid", m.raidID, "err", err)
			continue
		}
		s.markRaid(ctx, m.raidID)
		s.logger().Info("guilds", "op", "report_guild_moved", "report", reportID, "raid", m.raidID, "user_id", by)
	}
}

// raidRefKey names one raid by its guild.
type raidRefKey struct{ guildID, raidID int64 }

// foreignMatches is the raids holding reportID whose guild is not the report's guild for
// matching.
func (s *Store) foreignMatches(ctx context.Context, reportID string) ([]raidRefKey, error) {
	rows, err := s.Pool.Query(ctx, `
		select g.guild_id, g.id from guild_raids g join reports r on r.id = g.report_id
		where g.report_id = $1 and g.guild_id is distinct from `+matchGuildOf("r"), reportID)
	if err != nil {
		return nil, fmt.Errorf("guilds: foreign matches: %w", err)
	}
	defer rows.Close()
	var out []raidRefKey
	for rows.Next() {
		var k raidRefKey
		if err := rows.Scan(&k.guildID, &k.raidID); err != nil {
			return nil, fmt.Errorf("guilds: foreign matches: %w", err)
		}
		out = append(out, k)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: foreign matches: %w", err)
	}
	return out, nil
}
