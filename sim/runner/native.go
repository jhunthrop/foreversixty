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

// maxStderr bounds how much of a chatty binary's stderr is kept for
// the error message.
const maxStderr = 8 << 10

// tick is one progress line on stderr, as forever-sim writes it with
// -progress.
type tick struct {
	Completed int     `json:"completed"`
	Total     int     `json:"total"`
	DPS       float64 `json:"dps"`
}

// Native runs the forever-sim binary the image carries: the request
// as JSON on stdin, the result as JSON on stdout, progress as JSON
// lines on stderr.
type Native struct {
	// Binary is the engine executable. Empty means DefaultBinary.
	Binary string
}

// Run executes one sim. A non-zero exit, or output that does not
// decode, is an error; a result carrying its own Error field is
// returned as a result, because the engine refusing a character is an
// answer the caller stores rather than a failure of this process.
func (n *Native) Run(ctx context.Context, req api.SimRequest, onProgress Progress) (api.SimResult, error) {
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
	// deadlock the other.
	var (
		wg   sync.WaitGroup
		said strings.Builder
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		scan := bufio.NewScanner(stderr)
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
				onProgress(t.Completed, t.DPS)
			}
		}
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
		if errors.As(waitErr, &exit) && exit.ExitCode() == exitBadInput {
			return api.SimResult{}, fmt.Errorf("%w: %s: %s", ErrBadInput, bin, strings.TrimSpace(said.String()))
		}
		return api.SimResult{}, fmt.Errorf("runner: %s: %w: %s", bin, waitErr, strings.TrimSpace(said.String()))
	}
	if readErr != nil {
		return api.SimResult{}, fmt.Errorf("runner: read %s: %w", bin, readErr)
	}
	if len(bytes.TrimSpace(out)) == 0 {
		return api.SimResult{}, fmt.Errorf("runner: %s produced no result: %s", bin, strings.TrimSpace(said.String()))
	}
	var res api.SimResult
	if err := json.Unmarshal(out, &res); err != nil {
		return api.SimResult{}, fmt.Errorf("runner: %s wrote something that is not a SimResult: %w", bin, err)
	}
	res.Request, res.Lane = req, api.LaneServer
	if res.EngineVersion == "" {
		res.EngineVersion = req.EngineVersion
	}
	return res, nil
}
