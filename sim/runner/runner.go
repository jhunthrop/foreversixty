// Package runner runs one sim natively, by invoking the forever-sim
// binary. The premium lane's job, the fight-close scorer, and the
// nightly validation job all go through it, so there is one place
// that knows how the engine is started.
//
// Nothing here imports the engine. That is deliberate: the api module
// imports this package, and a module that imports sim/adapter or
// sim/request drags github.com/wowsims/classic into its own build,
// where sim/go.mod's replace does not reach it.
package runner

import (
	"context"
	"errors"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// Progress is called as iterations land, so a caller can refine the
// figure it is showing. It is never called after Run returns.
type Progress func(done int, mean float64)

// Runner runs one request to completion. Implementations must be safe
// for concurrent use.
type Runner interface {
	Run(ctx context.Context, req api.SimRequest, onProgress Progress) (api.SimResult, error)
}

// ErrBadInput is what a runner returns when the binary refused the
// request itself (exit 2): retrying it unchanged will fail the same
// way, so a caller records it rather than requeueing.
var ErrBadInput = errors.New("runner: the engine refused the request")

func isBadInput(err error) bool { return errors.Is(err, ErrBadInput) }

// ErrAborted is what a runner returns when the engine stopped a run
// before it finished, rather than failing it: a Cloud Run preemption,
// an operator's --task-timeout, or SIGINT/SIGTERM forwarded to the
// child. The result returned alongside it is not the zero value —
// forever-sim writes a partial SimResult with Aborted set before it
// exits, and Run decodes and returns that result together with this
// error, so a caller sees both what happened and how far the run got.
var ErrAborted = errors.New("runner: the engine stopped before the run finished")
