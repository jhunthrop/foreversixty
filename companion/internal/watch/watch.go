// Package watch turns a Logs directory into a stream of report
// boundaries and appended bytes. It owns no goroutine and reads no
// clock: Poll is called with the time, which is what makes a
// thirty-minute idle rule testable in a millisecond.
package watch

import (
	"bytes"
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

// HeaderScanLimit bounds how far back LastHeaderBefore reads for a header: a raid night of
// logging is a few hundred megabytes, and a header further back than this belongs to a
// client session the player has long since left.
const HeaderScanLimit = 256 << 20

// headerScanWindow is how much of the file one backwards step reads. A variable so a test
// can make the window small enough to split a header line across two steps.
var headerScanWindow = 1 << 20

// headerMark is what a COMBAT_LOG_VERSION line carries after its timestamp.
const headerMark = "  COMBAT_LOG_VERSION,"

// LastHeaderBefore finds the last whole COMBAT_LOG_VERSION line that ends at or before offset
// in the log at path, reading backwards from offset in windows and at most HeaderScanLimit
// bytes. A report that begins partway through a file (the file was already there and
// growing when the companion started, or when the tail moved to another game folder) never
// sees the header the game wrote when logging began, and the engine would have to guess the
// layout from field counts; the header is in the file, a few megabytes back, and this reads
// it. It reports false when there is none within reach.
func LastHeaderBefore(path string, offset int64) (string, bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", false, err
	}
	defer f.Close()
	end := offset
	floor := max(offset-HeaderScanLimit, 0)
	var carry []byte // the head of a line cut by the previous window, which began at end
	for end > floor {
		start := max(end-int64(headerScanWindow), floor)
		buf := make([]byte, end-start)
		if _, err := f.ReadAt(buf, start); err != nil {
			return "", false, fmt.Errorf("read %s at %d..%d: %w", path, start, end, err)
		}
		window := append(buf, carry...)
		if line, ok := lastHeaderLine(window, start > 0); ok {
			return line, true, nil
		}
		// Everything up to and including the first newline is the tail of a line that
		// starts earlier; the newline comes along so the joined line reads as terminated.
		if i := bytes.IndexByte(window, '\n'); i >= 0 {
			carry = append([]byte(nil), window[:i+1]...)
		} else {
			carry = append([]byte(nil), window...)
		}
		end = start
	}
	return "", false, nil
}

// lastHeaderLine finds the last header line among the whole lines of window. The first line
// of the window is whole only when the window begins at the file's start (partial false).
func lastHeaderLine(window []byte, partial bool) (string, bool) {
	at := len(window)
	for {
		i := bytes.LastIndex(window[:at], []byte(headerMark))
		if i < 0 {
			return "", false
		}
		lineStart := bytes.LastIndexByte(window[:i], '\n') + 1
		lineEnd := bytes.IndexByte(window[i:], '\n')
		if lineEnd < 0 {
			// The line's end lies past offset: the header belongs to the report itself.
			at = i
			continue
		}
		if lineStart == 0 && partial {
			return "", false // its start lies in an earlier window; the caller carries it back
		}
		return string(window[lineStart : i+lineEnd]), true
	}
}
