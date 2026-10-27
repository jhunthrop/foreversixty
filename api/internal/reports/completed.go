package reports

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"
)

const (
	// completionTimeout bounds the completion hook a request starts and no longer waits for.
	completionTimeout = 15 * time.Second
	// recordTimeout bounds the two inline completion writes (log zone, attributed guild).
	recordTimeout = 5 * time.Second
)

// completions runs a request's completion hook detached from it, so a slow match never stalls
// the caller's answer: the hook keeps the request's values but not its cancellation, and gets
// completionTimeout of its own. A panic is recovered and logged; Drain waits for the hooks still
// running (the API's shutdown).
type completions struct{ wg sync.WaitGroup }

func (c *completions) run(parent context.Context, log *slog.Logger, reportID string, work func(ctx context.Context)) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		defer func() {
			if p := recover(); p != nil {
				logger(log).Error("report_completed", "report", reportID, "step", "hook", "panic", fmt.Sprint(p))
			}
		}()
		ctx, cancel := context.WithTimeout(context.WithoutCancel(parent), completionTimeout)
		defer cancel()
		work(ctx)
	}()
}

// drain waits for every detached hook, or for ctx; it reports ctx's error when hooks remain.
func (c *completions) drain(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		c.wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("reports: completion hooks still running: %w", ctx.Err())
	}
}

// wait blocks until every detached completion has finished (tests).
func (c *completions) wait() { c.wg.Wait() }

func logger(log *slog.Logger) *slog.Logger {
	if log == nil {
		return slog.Default()
	}
	return log
}

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
// recordCompletion, then the hook. Each step is best effort - a failure is logged and the next
// step still runs - because the report is complete whatever happens here.
func FinishCompletion(ctx context.Context, s *Store, hook Completed, log *slog.Logger, reportID, zone string) {
	recordCompletion(ctx, s, log, reportID, zone)
	callCompleted(ctx, hook, reportID)
}

// recordCompletion records the zone the log named (empty when it named none) and gives the
// report its guild (AttributeGuild): the two writes the matcher and the picker read, so a
// request runs them before it answers and detaches only the hook.
func recordCompletion(ctx context.Context, s *Store, log *slog.Logger, reportID, zone string) {
	log = logger(log)
	if err := s.SetLogZone(ctx, reportID, zone); err != nil {
		log.Error("report_completed", "report", reportID, "step", "log_zone", "err", err)
	}
	if _, err := s.AttributeGuild(ctx, reportID); err != nil {
		log.Error("report_completed", "report", reportID, "step", "attribute_guild", "err", err)
	}
}

// recordCompletionBounded is recordCompletion on parent's values with recordTimeout of its own,
// so a companion that hangs up does not cancel the writes.
func recordCompletionBounded(parent context.Context, s *Store, log *slog.Logger, reportID, zone string) {
	ctx, cancel := context.WithTimeout(context.WithoutCancel(parent), recordTimeout)
	defer cancel()
	recordCompletion(ctx, s, log, reportID, zone)
}
