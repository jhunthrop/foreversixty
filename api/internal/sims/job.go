package sims

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/runner"
)

// RunTimeout bounds one premium run. Ten thousand iterations is about
// eight CPU-seconds natively; ten minutes is room for the staged Top
// Gear runs that come later and still a bound.
const RunTimeout = 10 * time.Minute

// JobDeps is everything `api sim-run <sim_id>` needs. There is no
// Getter here: the request is in the row, and the only object this job
// touches it writes.
type JobDeps struct {
	Store  *Store
	Put    store.Putter
	Engine runner.Runner
	Log    *slog.Logger
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
	stored, err := d.Store.Get(ctx, simID)
	if err != nil {
		// Nothing to mark: there is no row.
		return err
	}
	req := stored.Request

	runCtx, cancel := context.WithTimeout(ctx, RunTimeout)
	defer cancel()
	started := time.Now()
	res, err := d.Engine.Run(runCtx, req, func(done int, mean float64) {
		// A progress write that fails is logged and the run carries
		// on: the figure on the page is a courtesy, the result is not.
		if err := d.Store.Advance(ctx, simID, done, mean); err != nil {
			d.logger().Error("sims", "op", "progress", "sim", simID, "err", err)
		}
	})
	if err != nil {
		return d.failed(ctx, simID, err)
	}
	res.SimID, res.Lane = simID, simapi.LaneServer
	if res.DurationMS == 0 {
		res.DurationMS = time.Since(started).Milliseconds()
	}

	// The bucket and the row both get the whole result, so the page
	// needs one read and the history needs no bucket at all.
	body, err := json.Marshal(res)
	if err != nil {
		return d.failed(ctx, simID, fmt.Errorf("sims: encode result %s: %w", simID, err))
	}
	if err := d.Put.Put(ctx, Keys{SimID: simID}.Result(), body, ResultPut); err != nil {
		return d.failed(ctx, simID, err)
	}
	return d.Store.Finish(ctx, simID, res)
}

// failed records why a run did not finish and returns the original
// error, so the job exits non-zero and the page stops polling.
func (d JobDeps) failed(ctx context.Context, simID string, cause error) error {
	d.logger().Error("sims", "op", "run", "sim", simID, "err", cause)
	if err := d.Store.Fail(ctx, simID, cause.Error()); err != nil {
		return fmt.Errorf("%w (and the row could not be marked failed: %v)", cause, err)
	}
	return cause
}
