package guilds

import "context"

// ReportCompleted is the reports.Completed hook: a report has just completed, or a complete
// report has just been given its guild. It matches the report to the guild's raid and, on a
// match, computes the raid's attendance. Best effort: a failure is logged, never returned.
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
	if err := s.Store.RecomputeAttendance(ctx, raid.ID); err != nil {
		s.logger().Error("guilds", "op", "report_completed_attendance", "report", reportID, "raid", raid.ID, "err", err)
		return
	}
	s.logger().Info("guilds", "op", "report_completed", "report", reportID, "raid", raid.ID, "matched", true)
}
