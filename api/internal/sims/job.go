package sims

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/measure"
	"github.com/jhunthrop/foreversixty/sim/runner"
)

// RunTimeout bounds one premium run. Ten thousand iterations is about
// eight CPU-seconds natively; ten minutes is room for the staged Top
// Gear runs that come later and still a bound.
const RunTimeout = 10 * time.Minute

// BulkBudget bounds one bulk or weights run: every stage of the ladder,
// not one sim. measure.NativeJobSeconds fixes it at 840 seconds
// (contract A2). api/README.md creates the sim-run job with
// --task-timeout 15m, and a job killed by the platform dies between
// the bucket write and the row write, leaving the page polling
// "running" forever - so the in-process bound stops a minute short of
// it and fails the row on its way out. It is also the budget the
// submit-time estimate is refused against, so a run that is accepted
// is a run that can finish.
const BulkBudget = measure.NativeJobSeconds * time.Second

// timeoutFor is the bound one request's run gets. A plain run keeps the
// tighter one: ten minutes for a single sim is already an outlier worth
// failing.
func timeoutFor(req simapi.SimRequest) time.Duration {
	if req.Kind() == simapi.KindRun {
		return RunTimeout
	}
	return BulkBudget
}

// JobDeps is everything `api sim-run <sim_id>` needs. There is no
// Getter here: the request is in the row, and the only object this job
// touches it writes.
type JobDeps struct {
	Store  *Store
	Put    store.Putter
	Engine runner.StageRunner
	Log    *slog.Logger
	// Notify, when set, is told of every run that finishes in state done.
	Notify notify.SimEvents
}

func (d JobDeps) logger() *slog.Logger {
	if d.Log != nil {
		return d.Log
	}
	return slog.Default()
}

// Run is the premium lane's Cloud Run job: read the queued request off
// the row, run the engine natively, stream progress back into the row
// as iterations land, and write the finished result to Postgres and
// the bucket.
//
// Every failure after the row is found is recorded on the row as well
// as returned, so the page polling it is told what happened instead
// of waiting forever; the error is still returned so the job's exit
// code says it failed and the platform's logs show why.
func Run(ctx context.Context, d JobDeps, simID string) error {
	stored, _, err := d.Store.Get(ctx, simID)
	if err != nil {
		// Nothing to mark: there is no row.
		return err
	}
	req := stored.Request

	runCtx, cancel := context.WithTimeout(ctx, timeoutFor(req))
	defer cancel()
	started := time.Now()
	res, err := d.Engine.RunStaged(runCtx, req, func(p simapi.Progress) {
		// A progress write that fails is logged and the run carries
		// on: the figure on the page is a courtesy, the result is not.
		if err := d.Store.Advance(ctx, simID, Tick{
			IterationsDone: p.IterationsRun, Mean: p.DPS.Mean,
			Stage: p.Stage, CombosDone: p.CombosDone, CombosTotal: p.CombosTotal,
		}); err != nil {
			d.logger().Error("sims", "op", "progress", "sim", simID, "err", err)
		}
	})
	if err != nil {
		if errors.Is(err, runner.ErrAborted) {
			// The engine wrote a partial result before it stopped, but
			// the runner already discards it here: a stopped run is
			// never stored as a finished one, so the row is failed
			// with why an operator can read, not the engine's own
			// abort message.
			return d.failed(ctx, simID, err, "stopped before completion")
		}
		return d.failed(ctx, simID, err, "")
	}
	res.SimID, res.Lane = simID, simapi.LaneServer
	if res.DurationMS == 0 {
		res.DurationMS = time.Since(started).Milliseconds()
	}

	// The bucket and the row both get the whole result, so the page
	// needs one read and the history needs no bucket at all.
	body, err := json.Marshal(res)
	if err != nil {
		return d.failed(ctx, simID, fmt.Errorf("sims: encode result %s: %w", simID, err), "")
	}
	if err := d.Put.Put(ctx, Keys{SimID: simID}.Result(), body, ResultPut); err != nil {
		return d.failed(ctx, simID, err, "")
	}
	if err := d.Store.Finish(ctx, simID, res); err != nil {
		return err
	}
	if res.Error == "" {
		d.announce(ctx, simID, res)
	}
	return nil
}

// announce tells the owner's channel the run is done. The run already
// succeeded, so a failure here is logged and never fails the job.
func (d JobDeps) announce(ctx context.Context, simID string, res simapi.SimResult) {
	if d.Notify == nil {
		return
	}
	row, err := d.Store.Finished(ctx, simID)
	if err != nil {
		d.logger().Error("sims", "op", "notify", "sim", simID, "err", err)
		return
	}
	if row.UserID == nil {
		return
	}
	d.Notify.SimFinished(ctx, notify.SimFinished{
		UserID: *row.UserID, SimID: simID, Spec: res.Request.Spec, DPS: res.DPS.Mean,
		Headline: row.Headline, Title: row.Title, CreatedAt: row.CreatedAt, FinishedAt: row.FinishedAt,
	})
}

// failureUnknown is what a failed run's row says when the job has no
// plainer sentence for the member.
const failureUnknown = "the run could not be completed"

// failed records why a run did not finish and returns the original
// error, so the job exits non-zero and the page stops polling. reason
// is the plain sentence the row keeps and the member reads; an empty
// one becomes failureUnknown. The internal cause is never stored: it
// goes to the log and to the returned error only. (Finish is the one
// place a sentence from the engine itself is stored.)
func (d JobDeps) failed(ctx context.Context, simID string, cause error, reason string) error {
	if reason == "" {
		reason = failureUnknown
	}
	d.logger().Error("sims", "op", "run", "sim", simID, "err", cause)
	if err := d.Store.Fail(ctx, simID, reason); err != nil {
		return fmt.Errorf("%w (and the row could not be marked failed: %v)", cause, err)
	}
	return cause
}
