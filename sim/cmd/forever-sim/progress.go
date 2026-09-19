package main

import (
	"encoding/json"
	"io"
	"strings"
	"sync"
)

// progressStderr is stderr for a -progress run, where the contract is
// that every line is JSON.
//
// Two writers share that stream. The progress ticks are ours, written
// from the goroutine that drains the reporter channel. The engine's own
// commentary - "Running 3000 iterations on 14 concurrent sims", and the
// panic trace when a thread dies - goes through the standard log
// package, from whichever goroutine hit it. Left alone the two
// interleave, and the reader on the other end (sim/runner.Native) has to
// guess which lines are its own.
//
// So both go through here: one mutex, so a line is never cut in half,
// and the log's output wrapped as {"log": "..."} so it parses like
// everything else rather than being thrown away. A run that fails still
// says why on stderr; it just says it in the stream's own shape.
type progressStderr struct {
	mu sync.Mutex
	w  io.Writer
}

func newProgressStderr(w io.Writer) *progressStderr { return &progressStderr{w: w} }

// Write takes one already-encoded JSON line from the tick encoder.
func (s *progressStderr) Write(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.w.Write(p)
}

// LogOutput is what to hand log.SetOutput.
func (s *progressStderr) LogOutput() io.Writer { return logLines{s} }

type logLines struct{ out *progressStderr }

// Write wraps one log write - which is often several lines, a panic
// trace most of all - and emits it as ONE write on the shared stream.
// Taking the mutex per line let a tick land in the middle of a trace,
// which is the interleaving this file exists to prevent.
func (l logLines) Write(p []byte) (int, error) {
	var buf []byte
	for _, line := range strings.Split(strings.TrimRight(string(p), "\n"), "\n") {
		if line == "" {
			continue
		}
		b, err := json.Marshal(struct {
			Log string `json:"log"`
		}{line})
		if err != nil {
			// Dropping it would lose the one explanation a failed run
			// has. The line is not JSON-encodable, so it is reported
			// as the error it is rather than silently skipped.
			return 0, err
		}
		buf = append(append(buf, b...), '\n')
	}
	if len(buf) == 0 {
		return len(p), nil
	}
	if _, err := l.out.Write(buf); err != nil {
		return 0, err
	}
	return len(p), nil
}
