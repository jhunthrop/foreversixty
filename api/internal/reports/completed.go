package reports

import (
	"context"
	"log/slog"
)

// Completed is told of a report that has just become complete, or whose guild was just set
// on a complete report. It re-reads the report itself and never fails its caller: the report
// is already complete. guilds.Service satisfies it; the interface lives here so reports and
// parse never import guilds.
type Completed interface {
	ReportCompleted(ctx context.Context, reportID string)
}

// callCompleted tells hook of reportID; a nil hook is a deployment that listens to nothing.
func callCompleted(ctx context.Context, hook Completed, reportID string) {
	if hook == nil {
		return
	}
	hook.ReportCompleted(ctx, reportID)
}

// FinishCompletion is what every completion point runs once the report's status is complete:
// record the zone the log named (empty when it named none), give the report its guild
// (AttributeGuild), then call the hook. Each step is best effort - a failure is logged and the
// next step still runs - because the report is complete whatever happens here.
func FinishCompletion(ctx context.Context, s *Store, hook Completed, log *slog.Logger, reportID, zone string) {
	if log == nil {
		log = slog.Default()
	}
	if err := s.SetLogZone(ctx, reportID, zone); err != nil {
		log.Error("report_completed", "report", reportID, "step", "log_zone", "err", err)
	}
	if _, err := s.AttributeGuild(ctx, reportID); err != nil {
		log.Error("report_completed", "report", reportID, "step", "attribute_guild", "err", err)
	}
	callCompleted(ctx, hook, reportID)
}
