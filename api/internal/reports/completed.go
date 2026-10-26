package reports

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

// completionTimeout bounds the completion work a request starts and no longer waits for.
const completionTimeout = 15 * time.Second

// completions runs a request's completion work detached from it, so a slow database never
// stalls the caller's answer: the work keeps the request's values but not its cancellation, and
// gets completionTimeout of its own. Each step logs its own failure. wait is for tests.
type completions struct{ wg sync.WaitGroup }

func (c *completions) run(parent context.Context, work func(ctx context.Context)) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		ctx, cancel := context.WithTimeout(context.WithoutCancel(parent), completionTimeout)
		defer cancel()
		work(ctx)
	}()
}

// wait blocks until every detached completion has finished.
func (c *completions) wait() { c.wg.Wait() }

// Completed is told of a report that has just become complete, or whose guild was just set
// on a complete report. It re-reads the report itself and never fails its caller: the report
// is already complete. guilds.Service satisfies it; the interface lives here so reports and
// parse never import guilds.
type Completed interface {
	ReportCompleted(ctx context.Context, reportID string)
}

// GuildMoved is told when an officer's PATCH sets a report's guild, before the completion hook,
// so a raid of another guild still holding the report lets it go. guilds.Service satisfies it.
type GuildMoved interface {
	ReportGuildMoved(ctx context.Context, reportID string, by int64)
}

// GuildHooks is both hooks, as the guilds service provides them.
type GuildHooks interface {
	Completed
	GuildMoved
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
