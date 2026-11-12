package guilds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/reportguild"
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
	s.fillPullDigests(ctx, reportID)
	s.liveRaidEnded(ctx, raid.ID)
	s.logger().Info("guilds", "op", "report_completed", "report", reportID, "raid", raid.ID, "matched", true)
	if s.Events != nil {
		s.publishAll(ctx, s.summaryEvents(ctx, raid.ID, reportID))
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

// ReportVisibilityChanged is the reports.VisibilityChanged hook: a PATCH has just changed
// reportID's visibility. Every raid holding the report recomputes its attendance (a log made
// public again counts the sign-ups corrected while it was private); a report no raid holds runs
// the completion hook, so a log made visible after it completed still finds its raid. Best
// effort: a failure is logged, never returned.
func (s *Service) ReportVisibilityChanged(ctx context.Context, reportID string) {
	raids, err := s.Store.raidsHolding(ctx, reportID)
	if err != nil {
		s.logger().Error("guilds", "op", "report_visibility", "report", reportID, "err", err)
		return
	}
	if len(raids) == 0 {
		complete, err := s.Store.reportComplete(ctx, reportID)
		if err != nil {
			s.logger().Error("guilds", "op", "report_visibility", "report", reportID, "err", err)
			return
		}
		if complete {
			s.ReportCompleted(ctx, reportID)
		}
		return
	}
	for _, id := range raids {
		s.recomputeAttendance(ctx, id)
	}
}

// recomputeAttendance rebuilds raidID's attendance and marks its post. Best effort: a failure is
// logged, never returned (the write that asked for it already stands).
func (s *Service) recomputeAttendance(ctx context.Context, raidID int64) {
	if err := s.Store.RecomputeAttendance(ctx, raidID); err != nil {
		s.logger().Error("guilds", "op", "recompute_attendance", "raid_id", raidID, "err", err)
	}
	s.markRaid(ctx, raidID)
}

// raidsHolding is the ids of the raids whose log is reportID.
func (s *Store) raidsHolding(ctx context.Context, reportID string) ([]int64, error) {
	rows, err := s.Pool.Query(ctx, `select id from guild_raids where report_id = $1 order by id`, reportID)
	if err != nil {
		return nil, fmt.Errorf("guilds: raids holding report: %w", err)
	}
	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return nil, fmt.Errorf("guilds: raids holding report: %w", err)
	}
	return ids, nil
}

// reportComplete reports whether reportID is a complete report.
func (s *Store) reportComplete(ctx context.Context, reportID string) (bool, error) {
	var ok bool
	if err := s.Pool.QueryRow(ctx,
		`select exists (select 1 from reports where id = $1 and status = 'complete')`, reportID).Scan(&ok); err != nil {
		return false, fmt.Errorf("guilds: report complete: %w", err)
	}
	return ok, nil
}

// raidRefKey names one raid by its guild.
type raidRefKey struct{ guildID, raidID int64 }

// foreignMatches is the raids holding reportID whose guild is not the report's guild for
// matching.
func (s *Store) foreignMatches(ctx context.Context, reportID string) ([]raidRefKey, error) {
	rows, err := s.Pool.Query(ctx, `
		select g.guild_id, g.id from guild_raids g join reports r on r.id = g.report_id
		where g.report_id = $1 and g.guild_id is distinct from `+reportguild.Of("r"), reportID)
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
