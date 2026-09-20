package runner

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// DefaultBinary is where the image puts the engine binary CI builds
// from the pinned sha. See the api lane's Task 15.
const DefaultBinary = "/engine/forever-sim"

// exitBadInput is forever-sim's exit code for a request it could not
// decode or could not build: 0 is success, 1 an engine failure, 2 bad
// input.
const exitBadInput = 2

// exitAborted is forever-sim's exit code when the run was stopped on
// request (a Cloud Run preemption, an operator's --task-timeout, or
// SIGINT/SIGTERM forwarded to the child) rather than run to
// completion. It is the shell's own convention for "killed by
// SIGINT" (128 + signal 2), which is also what sim/cmd/forever-sim's
// main.go exits with; it is redeclared here rather than imported
// because that package pulls in the engine, which this one must not.
const exitAborted = 130

// maxStderr bounds how much of a chatty binary's stderr is kept for
// the error message.
const maxStderr = 8 << 10

// maxStderrLine bounds the largest single stderr line the scanner will
// accept. forever-sim wraps the engine's own log output as a single
// {"log":"..."} line on the same stream as our progress ticks; one
// line past bufio.Scanner's 64KB default (a panic stack, a long
// engine warning) made Scan return false with bufio.ErrTooLong, which
// silently killed the goroutine and every progress tick after it. 1MB
// comfortably covers a pathological log line without buffering the
// whole stream.
const maxStderrLine = 1 << 20

// tick is one progress line on stderr, as forever-sim writes it with
// -progress. The three bulk fields are absent for a plain run.
type tick struct {
	Completed   int     `json:"completed"`
	Total       int     `json:"total"`
	DPS         float64 `json:"dps"`
	Stage       int     `json:"stage"`
	CombosDone  int     `json:"combos_done"`
	CombosTotal int     `json:"combos_total"`
}

// Native runs the forever-sim binary the image carries: the request
// as JSON on stdin, the result as JSON on stdout, progress as JSON
// lines on stderr.
type Native struct {
	// Binary is the engine executable. Empty means DefaultBinary.
	Binary string
}

// Run executes one sim, reporting the iteration count and the running
// mean. It is RunStaged with the three bulk fields dropped.
func (n *Native) Run(ctx context.Context, req api.SimRequest, onProgress Progress) (api.SimResult, error) {
	var staged StageProgress
	if onProgress != nil {
		staged = func(p api.Progress) { onProgress(p.IterationsRun, p.DPS.Mean) }
	}
	return n.RunStaged(ctx, req, staged)
}

// RunStaged executes one sim, reporting the wider stage progress. A
// non-zero exit, or output that does not decode, is an error; a
// result carrying its own Error field is returned as a result,
// because the engine refusing a character is an answer the caller
// stores rather than a failure of this process.
func (n *Native) RunStaged(ctx context.Context, req api.SimRequest, onProgress StageProgress) (api.SimResult, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return api.SimResult{}, fmt.Errorf("runner: encode request: %w", err)
	}
	bin := n.Binary
	if bin == "" {
		bin = DefaultBinary
	}
	cmd := exec.CommandContext(ctx, bin, "-in", "-", "-out", "-", "-progress")
	cmd.Stdin = bytes.NewReader(body)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return api.SimResult{}, fmt.Errorf("runner: stdout: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return api.SimResult{}, fmt.Errorf("runner: stderr: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return api.SimResult{}, fmt.Errorf("runner: start %s: %w", bin, err)
	}

	// stderr is read in its own goroutine: the binary interleaves
	// progress with the result, and a full pipe on either side would
	// deadlock the other. scanErr is written here and read only after
	// wg.Wait() returns below, which happens-after this goroutine's
	// Done call — the same safe handoff said already relies on.
	var (
		wg      sync.WaitGroup
		said    strings.Builder
		scanErr error
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		scan := bufio.NewScanner(stderr)
		scan.Buffer(make([]byte, 0, 64<<10), maxStderrLine)
		for scan.Scan() {
			line := scan.Bytes()
			var t tick
			if err := json.Unmarshal(line, &t); err != nil || t.Total == 0 {
				// Not a progress line: a log message, a warning. Keep
				// a bounded amount of it for the error message.
				if said.Len() < maxStderr {
					said.Write(line)
					said.WriteByte('\n')
				}
				continue
			}
			if onProgress != nil {
				onProgress(api.Progress{
					IterationsRun: t.Completed,
					DPS:           api.Estimate{Mean: t.DPS},
					Stage:         t.Stage,
					CombosDone:    t.CombosDone,
					CombosTotal:   t.CombosTotal,
				})
			}
		}
		scanErr = scan.Err()
	}()

	out, readErr := io.ReadAll(stdout)
	wg.Wait()
	waitErr := cmd.Wait()
	if waitErr != nil {
		// A cancelled or expired context is why the process died, not
		// what it died of: exec reports a killed child as the opaque
		// "signal: killed", which would otherwise mask the caller's
		// own deadline behind an engine-failure-shaped error.
		if ctxErr := ctx.Err(); ctxErr != nil {
			return api.SimResult{}, fmt.Errorf("runner: %s: %w", bin, ctxErr)
		}
		var exit *exec.ExitError
		if errors.As(waitErr, &exit) {
			switch exit.ExitCode() {
			case exitBadInput:
				return api.SimResult{}, fmt.Errorf("%w: %s: %s", ErrBadInput, bin, strings.TrimSpace(said.String()))
			case exitAborted:
				// forever-sim writes the partial result before it
				// exits 130. If it decodes, the caller gets both the
				// result and ErrAborted; if it does not (a crash
				// before anything was written), fall through to the
				// generic error below.
				if res, decErr := decodeResult(out, req); decErr == nil {
					return res, fmt.Errorf("%w: %s", ErrAborted, bin)
				}
			}
		}
		if scanErr != nil {
			// The process also failed for its own reason; a torn
			// stderr stream may be why said is incomplete, so both are
			// worth telling the caller.
			return api.SimResult{}, fmt.Errorf("runner: %s: %w: %s (reading its progress stream also failed: %w)", bin, waitErr, strings.TrimSpace(said.String()), scanErr)
		}
		return api.SimResult{}, fmt.Errorf("runner: %s: %w: %s", bin, waitErr, strings.TrimSpace(said.String()))
	}
	if readErr != nil {
		return api.SimResult{}, fmt.Errorf("runner: read %s: %w", bin, readErr)
	}
	if len(bytes.TrimSpace(out)) == 0 {
		return api.SimResult{}, fmt.Errorf("runner: %s produced no result: %s", bin, strings.TrimSpace(said.String()))
	}
	res, decErr := decodeResult(out, req)
	if decErr != nil {
		return api.SimResult{}, fmt.Errorf("runner: %s wrote something that is not a SimResult: %w", bin, decErr)
	}
	// scanErr is dropped here on purpose: the process exited 0 and
	// stdout — a separate pipe — decoded to a real result, so a torn
	// stderr progress stream does not invalidate what the engine
	// actually produced. The result wins.
	return res, nil
}

// planBody is writePlan's JSON, minus its "stage" field: Plan only
// ever wants the count. json.Decoder still has to scan past "stage"'s
// bytes to find the object's end - that part of a request's payload
// cannot be skipped when talking to a subprocess over a pipe, and at
// the server cap it is roughly 22MB - but omitting the field here
// means the decoder never allocates a Go value for any of it: no
// bulk.StageRequests, no api.SimRequest per combination, the way
// unmarshalling into the real type would. It also keeps sim/bulk's
// types, and everything they pull in, out of this package's imports,
// which the file doc above says nothing here may do.
//
// The same struct reads both of writePlan's shapes: a plan
// ("combinations" and the now-ignored "stage") and a cap breach
// ("error", "cap" and "combinations" again, this time the count that
// breached).
type planBody struct {
	Error        string `json:"error"`
	Cap          int    `json:"cap"`
	Combinations int    `json:"combinations"`
}

// Plan shells "forever-sim -plan": it prints simCount's answer and
// the first stage without running anything (contract 10.2), which is
// how a caller here counts a bulk request the same way the API does,
// without importing sim/internal.
//
// A cap breach comes back as api.ErrCapExceeded via errors.As, the
// same type bulk.Count itself returns, so a caller cannot tell this
// apart from having called bulk.Count directly. The binary exits 2
// for both a cap breach and ordinary bad input - writePlan wraps
// errBadInput either way - so the body is what tells them apart, not
// the exit code: a cap breach writes {"error":"cap_exceeded",...} to
// stdout, and ordinary bad input writes nothing there at all, the
// message going to stderr instead.
func (n *Native) Plan(ctx context.Context, req api.SimRequest) (api.PlanSummary, error) {
	if req.Bulk == nil {
		return api.PlanSummary{}, fmt.Errorf("%w: Plan takes a bulk request; this one has none", ErrBadInput)
	}
	body, err := json.Marshal(req)
	if err != nil {
		return api.PlanSummary{}, fmt.Errorf("runner: encode request: %w", err)
	}
	bin := n.Binary
	if bin == "" {
		bin = DefaultBinary
	}
	cmd := exec.CommandContext(ctx, bin, "-plan", "-in", "-", "-out", "-")
	cmd.Stdin = bytes.NewReader(body)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return api.PlanSummary{}, fmt.Errorf("runner: stdout: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return api.PlanSummary{}, fmt.Errorf("runner: stderr: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return api.PlanSummary{}, fmt.Errorf("runner: start %s: %w", bin, err)
	}

	// stderr is read in its own goroutine, the same reason RunStaged's
	// is: -plan writes little to it, but a full pipe buffer on either
	// side would still deadlock the other side's write.
	var (
		wg   sync.WaitGroup
		said strings.Builder
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		kept, _ := io.ReadAll(io.LimitReader(stderr, maxStderr))
		said.Write(kept)
		// Drain whatever maxStderr didn't keep, so a chattier-than-
		// expected binary still exits instead of blocking on a full
		// pipe nobody is reading.
		io.Copy(io.Discard, stderr)
	}()

	var out planBody
	decodeErr := json.NewDecoder(stdout).Decode(&out)
	// writePlan makes exactly one Write call with nothing before or
	// after it, so nothing should be left; draining anyway is the
	// same defensive read RunStaged gives stdout.
	io.Copy(io.Discard, stdout)
	wg.Wait()
	waitErr := cmd.Wait()

	if waitErr != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return api.PlanSummary{}, fmt.Errorf("runner: %s: %w", bin, ctxErr)
		}
		var exit *exec.ExitError
		if errors.As(waitErr, &exit) && exit.ExitCode() == exitBadInput {
			if decodeErr == nil && out.Error == "cap_exceeded" {
				return api.PlanSummary{}, api.ErrCapExceeded{Cap: out.Cap, Combinations: out.Combinations}
			}
			return api.PlanSummary{}, fmt.Errorf("%w: %s: %s", ErrBadInput, bin, strings.TrimSpace(said.String()))
		}
		return api.PlanSummary{}, fmt.Errorf("runner: %s: %w: %s", bin, waitErr, strings.TrimSpace(said.String()))
	}
	if decodeErr != nil {
		return api.PlanSummary{}, fmt.Errorf("runner: %s wrote something -plan does not: %w", bin, decodeErr)
	}
	ladder, ok := api.Ladders[req.Bulk.Precision]
	if !ok {
		return api.PlanSummary{}, fmt.Errorf("runner: no ladder for precision %q", req.Bulk.Precision)
	}
	return api.PlanSummary{
		Kind:            req.Kind(),
		Combinations:    out.Combinations,
		Cap:             req.Bulk.Cap,
		IterationsTotal: api.LadderIterations(ladder, out.Combinations),
	}, nil
}

// decodeResult unmarshals out as a SimResult and fills in the fields
// every successful path sets the same way, so Run's normal-exit and
// exit-130 branches share one implementation instead of two that can
// drift apart.
func decodeResult(out []byte, req api.SimRequest) (api.SimResult, error) {
	var res api.SimResult
	if err := json.Unmarshal(out, &res); err != nil {
		return api.SimResult{}, err
	}
	res.Request, res.Lane = req, api.LaneServer
	if res.EngineVersion == "" {
		res.EngineVersion = req.EngineVersion
	}
	return res, nil
}
