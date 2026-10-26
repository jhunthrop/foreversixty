package guilds

import "context"

// ReportCompleted is the reports.Completed hook: a report has just completed, or a complete
// report has just been given its guild. Lane G2 Task 2 matches the report to the guild's raid
// here; until then the call is only logged.
func (s *Service) ReportCompleted(_ context.Context, reportID string) {
	s.logger().Info("guilds", "op", "report_completed", "report", reportID)
}
