// Package watch turns a Logs directory into a stream of report
// boundaries and appended bytes. It owns no goroutine and reads no
// clock: Poll is called with the time, which is what makes a
// thirty-minute idle rule testable in a millisecond.
package watch

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/wow"
)

// The timing rules from the interface contract: a size increase after
// an idle gap longer than NewReportIdle starts a new report, and a
// report is completed once the file has been quiet for CompleteIdle.
const (
	NewReportIdle = 30 * time.Minute
	CompleteIdle  = 15 * time.Minute
)

// ReadChunk is how many bytes one Poll reads at most. It bounds the
// work per tick; the next tick takes the rest.
const ReadChunk = 1 << 20

// Kind is what one event tells the pipeline to do.
type Kind int

// The three boundaries a tail produces.
const (
	// Start opens a report at Offset in Path.
	Start Kind = iota
	// Append carries bytes at Offset in Path.
	Append
	// Complete closes the open report; Offset is its final offset.
	Complete
)

func (k Kind) String() string {
	switch k {
	case Start:
		return "start"
	case Append:
		return "append"
	case Complete:
		return "complete"
	}
	return "unknown"
}

// Event is one thing that happened to the log directory.
type Event struct {
	Kind   Kind
	Path   string
	Offset int64
	Data   []byte
	At     time.Time
}

// Options configures a Watcher.
type Options struct {
	// Dir is the install's Logs directory.
	Dir string
	// NewReportIdle and CompleteIdle default to the contract's values.
	NewReportIdle time.Duration
	CompleteIdle  time.Duration
	// ReadChunk defaults to ReadChunk.
	ReadChunk int
}

// Watcher follows the newest combat log in one directory.
type Watcher struct {
	o Options

	started    bool
	path       string
	info       os.FileInfo
	offset     int64
	open       bool // a report is open on path
	lastAppend time.Time
}

// New builds a watcher. Nothing is read until the first Poll.
func New(o Options) *Watcher {
	if o.NewReportIdle == 0 {
		o.NewReportIdle = NewReportIdle
	}
	if o.CompleteIdle == 0 {
		o.CompleteIdle = CompleteIdle
	}
	if o.ReadChunk == 0 {
		o.ReadChunk = ReadChunk
	}
	return &Watcher{o: o}
}

// Resume continues an open report rather than starting a new one. The
// pipeline calls it at startup for the report its state files say was
// in progress, so a restart mid-raid does not split the night in two.
func (w *Watcher) Resume(path string, offset int64, lastAppend time.Time) error {
	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("resume %s: %w", path, err)
	}
	if fi.Size() < offset {
		return fmt.Errorf("resume %s: the file is %d bytes, shorter than the saved offset %d",
			path, fi.Size(), offset)
	}
	w.started, w.path, w.info = true, path, fi
	w.offset, w.open, w.lastAppend = offset, true, lastAppend
	return nil
}

// Offset is the next byte the watcher will read.
func (w *Watcher) Offset() int64 { return w.offset }

// Path is the file being followed, empty before the first Poll.
func (w *Watcher) Path() string { return w.path }

// active picks the log to follow: the most recently modified match,
// with the path as a tiebreak so two files written in the same second
// resolve the same way on every machine.
func (w *Watcher) active() (string, os.FileInfo, error) {
	matches, err := filepath.Glob(filepath.Join(w.o.Dir, wow.LogGlob))
	if err != nil {
		return "", nil, err
	}
	sort.Strings(matches)
	var bestPath string
	var bestInfo os.FileInfo
	for _, m := range matches {
		fi, err := os.Stat(m)
		if err != nil || fi.IsDir() {
			continue
		}
		if bestInfo == nil || fi.ModTime().After(bestInfo.ModTime()) ||
			(fi.ModTime().Equal(bestInfo.ModTime()) && m > bestPath) {
			bestPath, bestInfo = m, fi
		}
	}
	if bestInfo == nil {
		return "", nil, nil
	}
	return bestPath, bestInfo, nil
}

// Poll reads whatever has happened since the last call and returns the
// events in the order the pipeline must apply them.
func (w *Watcher) Poll(now time.Time) ([]Event, error) {
	path, info, err := w.active()
	if err != nil {
		return nil, fmt.Errorf("scan %s: %w", w.o.Dir, err)
	}
	first := !w.started
	w.started = true

	if info == nil {
		return w.idle(now), nil
	}

	var out []Event
	rotated := w.info != nil && (path != w.path || !os.SameFile(info, w.info) || info.Size() < w.offset)
	if rotated || w.info == nil {
		if w.open {
			out = append(out, Event{Kind: Complete, Path: w.path, Offset: w.offset, At: now})
			w.open = false
		}
		w.path, w.info = path, info
		// A file that was already there when the companion started
		// holds a night we did not watch: begin at its end. A file
		// that appeared while we were running begins at zero.
		if first {
			w.offset = info.Size()
		} else {
			w.offset = 0
		}
	} else {
		w.info = info
	}

	if info.Size() <= w.offset {
		return append(out, w.idle(now)...), nil
	}

	data, err := w.read(info.Size())
	if err != nil {
		return out, err
	}
	if len(data) == 0 {
		return append(out, w.idle(now)...), nil
	}
	if w.open && now.Sub(w.lastAppend) > w.o.NewReportIdle {
		out = append(out, Event{Kind: Complete, Path: w.path, Offset: w.offset, At: now})
		w.open = false
	}
	if !w.open {
		out = append(out, Event{Kind: Start, Path: w.path, Offset: w.offset, At: now})
		w.open = true
	}
	out = append(out, Event{Kind: Append, Path: w.path, Offset: w.offset, Data: data, At: now})
	w.offset += int64(len(data))
	w.lastAppend = now
	return out, nil
}

// idle emits the completion when the file has been quiet long enough.
func (w *Watcher) idle(now time.Time) []Event {
	if !w.open || now.Sub(w.lastAppend) < w.o.CompleteIdle {
		return nil
	}
	w.open = false
	return []Event{{Kind: Complete, Path: w.path, Offset: w.offset, At: now}}
}

// read pulls at most ReadChunk bytes from the current offset.
func (w *Watcher) read(size int64) ([]byte, error) {
	f, err := os.Open(w.path)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil // rotated away between the stat and the open
	}
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", w.path, err)
	}
	defer f.Close()
	n := size - w.offset
	if n > int64(w.o.ReadChunk) {
		n = int64(w.o.ReadChunk)
	}
	buf := make([]byte, n)
	read, err := f.ReadAt(buf, w.offset)
	if err != nil && read == 0 {
		return nil, fmt.Errorf("read %s at %d: %w", w.path, w.offset, err)
	}
	return buf[:read], nil
}

// ReadRange reads the bytes of one fight back out of the log so the
// bundle can carry their hash. A range the file no longer holds
// returns an error and the caller sends the bundle without a hash
// rather than dropping the fight.
func ReadRange(path string, start, end int64) ([]byte, error) {
	if end < start {
		return nil, fmt.Errorf("read %s: range %d..%d runs backwards", path, start, end)
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := make([]byte, end-start)
	n, err := f.ReadAt(buf, start)
	if err != nil && int64(n) < end-start {
		return nil, fmt.Errorf("read %s at %d..%d: %w", path, start, end, err)
	}
	return buf[:n], nil
}
